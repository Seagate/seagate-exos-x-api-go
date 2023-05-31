// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package mcapi

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/Seagate/seagate-exos-x-api-go/pkg/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

const (
	ApiMaximumLUN         = 255
	ApiTierAffinity       = "no-affinity"
	ApiSuccess            = 0
	ApiVolumeDoesNotExist = 1
	ApiVolumeDeleted      = 2
)

// Volume : volume-view representation
type Volume struct {
	LUN int
}

type Volumes []Volume

func (v Volumes) Len() int {
	return len(v)
}

func (v Volumes) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v Volumes) Less(i, j int) bool {
	return v[i].LUN < v[j].LUN
}

// CreateVolume : creates a volume with the given name, capacity in the given pool
func (client *Client) CreateVolume(name, size, pool, poolType string) (*common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)
	response, httpRes, err := client.apiClient.DefaultApi.CreateVolumePoolSizeTierAffinityNameGet(client.Ctx, pool, size, ApiTierAffinity, name).Execute()
	logger.V(2).Info("create volume", "name", name, "http", httpRes.Status)
	status := CreateCommonStatus(response)

	return status, err
}

// CheckVolumeExists: Return true if a volume already exists
func (client *Client) CheckVolumeExists(volumeId string, size int64) (bool, error) {

	logger := klog.FromContext(client.Ctx)

	volumes, responseStatus, err := client.ShowVolumes(volumeId)
	logger.V(2).Info("check volume exists", "volume", volumeId, "size", size)

	if err != nil && responseStatus.ReturnCode != common.BadInputParam {
		return false, err
	}

	for _, v := range volumes {
		if v.VolumeName == volumeId {
			logger.V(2).Info("volume exists", "volume", volumeId, "size", size, "blocks", v.Blocks, "blocksize", v.BlockSize)
			if v.Blocks*v.BlockSize == size {
				return true, nil
			} else {
				return true, status.Error(codes.AlreadyExists, "cannot create volume with same name but different capacity than the existing one")
			}
		}
	}

	return false, nil
}

// ShowVolumes : get information about volumes
func (client *Client) ShowVolumes(volume string) ([]common.VolumeObject, *common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)

	// show volumes
	logger.V(4).Info("")
	logger.V(4).Info("================================================================================")
	logger.V(4).Info(">> ShowVolumesNamesGet")

	response, httpRes, err := client.apiClient.DefaultApi.ShowVolumesNamesGet(client.Ctx, volume).Execute()

	if httpRes.StatusCode == http.StatusOK && response != nil {

		// Extract Status resource information
		logger.V(4).Info("++ ShowVolumesNamesGet.Status[0]",
			"Response", *response.Status[0].Response,
			"ReturnCode", *response.Status[0].ReturnCode,
		)

		for i, volume := range response.Volumes {
			logger.V(4).Info("-- volumes",
				"index", i,
				"VolumeName", *volume.VolumeName,
				"TotalSize", *volume.TotalSize,
				"StoragePoolName", *volume.StoragePoolName,
				"SerialNumber", *volume.SerialNumber,
			)
		}
	} else {
		logger.V(4).Info("-- ShowVolumesNamesGet", "status", httpRes.Status, "err", err, "body", httpRes.Body)
	}

	logger.V(4).Info("================================================================================")

	returnVolumes := []common.VolumeObject{}
	status := CreateCommonStatusFromStatus(&response.Status[0])

	if response != nil {
		for _, v := range response.Volumes {
			volume := common.VolumeObject{}
			volume.ObjectName = v.GetObjectName()
			volume.Blocks = v.GetBlocks()
			volume.BlockSize = v.GetBlocksize()
			volume.Health = v.GetHealth()
			volume.SizeNumeric = v.GetSizeNumeric()
			volume.StoragePoolName = v.GetStoragePoolName()
			volume.StorageType = v.GetStorageType()
			volume.TierAffinity = v.GetTierAffinity()
			volume.TotalSize = v.GetTotalSize()
			volume.VolumeName = v.GetVolumeName()
			volume.VolumeType = v.GetVolumeType()
			volume.Wwn = v.GetWwn()

			returnVolumes = append(returnVolumes, volume)
		}
	}

	return returnVolumes, status, err
}

// GetVolumeWwn: Retrieve the WWN for a volume, very useful for host operating system device mapping
func (client *Client) GetVolumeWwn(volumeName string) (string, error) {

	logger := klog.FromContext(client.Ctx)

	wwn := ""
	response, status, err := client.ShowVolumes(volumeName)
	if err == nil && status.ResponseTypeNumeric == 0 {
		if len(response) > 0 && response[0].VolumeName == volumeName {
			wwn = strings.ToLower(response[0].Wwn)
		}
	}

	logger.V(3).Info("GetVolumeWwn", "volume", volumeName, "wwn", wwn)
	return wwn, err
}

// GetVolumeMapsHostNames: Use /show/maps to retrieve host names
func (client *Client) GetVolumeMapsHostNames(name string) ([]string, *common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)

	logger.V(2).Info("get volume maps host names", "volume", name)
	response, httpRes, err := client.apiClient.DefaultApi.ShowMapsNamesGet(client.Ctx, name).Execute()

	status := CreateCommonStatusFromStatus(&response.Status[0])

	if err != nil {
		return []string{}, status, err
	}

	hostNames := []string{}
	if httpRes.StatusCode == http.StatusOK && response != nil {
		for _, vv := range response.GetVolumeView() {
			for _, vvm := range vv.GetVolumeViewMappings() {
				hostName := vvm.GetIdentifier()
				if vvm.GetObjectName() == "host-view" && hostName != "all other initiators" {
					logger.V(2).Info("++ host-view", "name", hostName)
					hostNames = append(hostNames, hostName)
				}

			}
		}

	}

	return hostNames, status, err
}

// ShowHostMaps: list the volume mappings for given host. If host is an empty string, mapping for all hosts is shown
func (client *Client) ShowHostMaps(host string) ([]Volume, *common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)

	// We don't use "/show/maps/initiator/<host>" here because
	// the maps for an initiator with a nickname or in a host group will not
	// be returned. Instead we get all initiator mappings and filter by initiator
	logger.Info("++ ShowHostMaps", "host", host)

	response, _, err := client.apiClient.DefaultApi.ShowMapsInitiatorNamesGet(client.Ctx, "").Execute()
	status := CreateCommonStatusFromStatus(&response.Status[0])

	if err != nil {
		return nil, status, err
	}

	mappings := make([]Volume, 0)

	for _, hv := range response.GetHostsView() {
		if host != "" {
			if hv.GetHostName() != host {
				continue
			}
		}

		for _, hvm := range hv.GetHostViewMappings() {
			if hvm.GetObjectName() == "volume-view" {
				vol := Volume{}
				vol.LUN, _ = strconv.Atoi(hvm.GetLun())
				logger.Info("++ volume", "volume", hvm.GetVolumeName(), "lun", vol.LUN)
				mappings = append(mappings, vol)
			}
		}
	}

	return mappings, status, err
}

// chooseLUN: Choose the next available LUN for a given initiator
func (client *Client) chooseLUN(initiators []string) (int, error) {

	logger := klog.FromContext(client.Ctx)
	logger.Info("listing all LUN mappings")

	var allvolumes []Volume
	for _, initiatorName := range initiators {
		volumes, responseStatus, err := client.ShowHostMaps(initiatorName)
		if err != nil {
			logger.Error(err, "error looking for host maps", "initiator", initiatorName)
		}
		if responseStatus.ReturnCode == common.HostMapDoesNotExistsErrorCode {
			logger.Info("initiator does not exist", "initiator", initiatorName)
		}
		if volumes != nil {
			allvolumes = append(allvolumes, volumes...)
		}
	}

	sort.Sort(Volumes(allvolumes))

	logger.V(5).Info("use LUN 1 when volumes slice is empty")
	if len(allvolumes) == 0 {
		return 1, nil
	}

	logger.V(5).Info("use the next highest LUN number, until the end is reached")
	if allvolumes[len(allvolumes)-1].LUN+1 < ApiMaximumLUN {
		return allvolumes[len(allvolumes)-1].LUN + 1, nil
	}

	logger.V(5).Info("use LUN 1 when not in use")
	if allvolumes[0].LUN > 1 {
		return 1, nil
	}

	logger.V(5).Info("use the next available LUN, searching from LUN 1 towards the maximum")
	for index := 1; index < len(allvolumes); index++ {
		// Find a gap between used LUNs
		if allvolumes[index].LUN-allvolumes[index-1].LUN > 1 {
			return allvolumes[index-1].LUN + 1, nil
		}
	}

	logger.V(0).Info("no available LUN", "total", len(allvolumes), "volumes", allvolumes)

	return -1, status.Error(codes.ResourceExhausted, "no more available LUNs")
}

// MapVolume : map a volume to an initiator using a specified LUN
func (client *Client) MapVolume(name, initiator, access string, lun int) (*common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)
	response, httpRes, err := client.apiClient.DefaultApi.MapVolumeAccessLunInitiatorNamesGet(client.Ctx, access, strconv.Itoa(lun), initiator, name).Execute()
	logger.V(2).Info("map volume", "name", name, "lun", lun, "initiator", initiator, "http", httpRes.Status)
	return CreateCommonStatus(response), err
}

// CreateNickname : Create a nickname for an initiator. The Storage API policy is to prohibit mapping of initiators which are not either
// (a) presently connected to the array or (b) represented by an entry in the initiator nickname table.
func (client *Client) CreateNickname(name, iqn string) (*common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)
	response, httpRes, err := client.apiClient.DefaultApi.SetInitiatorIdNicknameGet(client.Ctx, iqn, name).Execute()
	logger.V(2).Info("create nickname", "name", name, "nickname", iqn, "http", httpRes.Status)
	return CreateCommonStatus(response), err
}

// mapVolumeProcess: Map a volume to an initiator and create a nickname when required by the storage array
func (client *Client) mapVolumeProcess(volumeName, initiatorName string, lun int) error {

	logger := klog.FromContext(client.Ctx)
	logger.Info("trying to map volume", "volume", volumeName, "initiator", initiatorName, "lun", lun)
	metadata, err := client.MapVolume(volumeName, initiatorName, "rw", lun)
	if err != nil && metadata == nil {
		return err
	}

	logger.Info("status", "ReturnCode", metadata.ReturnCode)
	if metadata.ReturnCode == common.InitiatorNicknameOrIdentifierNotFound {
		nodeIDParts := strings.Split(initiatorName, ":")
		if len(nodeIDParts) < 2 {
			return status.Error(codes.NotFound, "specified node ID is not a valid IQN")
		}

		nickname := strings.Join(nodeIDParts[1:], ":")
		nickname = strings.ReplaceAll(nickname, ".", "-")

		logger.Info("initiator does not exist, creating it", "nickname", nickname)
		_, err = client.CreateNickname(nickname, initiatorName)
		if err != nil {
			return err
		}
		klog.Info("retrying to map volume")
		_, err = client.MapVolume(volumeName, initiatorName, "rw", lun)
		if err != nil {
			return err
		}
	} else if metadata.ReturnCode == common.VolumeNotFoundErrorCode {
		return status.Errorf(codes.NotFound, "volume %s not found", volumeName)
	} else if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

// UnmapVolume : unmap a volume from an initiator
func (client *Client) UnmapVolume(name, initiator string) (*common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)

	if initiator == "" {
		response, httpRes, err := client.apiClient.DefaultApi.UnmapVolumeNamesGet(client.Ctx, name).Execute()
		logger.V(2).Info("unmap volume", "name", name, "initiator", initiator, "http", httpRes.Status)
		return CreateCommonStatus(response), err
	}

	response, httpRes, err := client.apiClient.DefaultApi.UnmapVolumeInitiatorNamesGet(client.Ctx, initiator, name).Execute()
	logger.V(2).Info("unmap volume", "name", name, "initiator", initiator, "http", httpRes.Status)
	return CreateCommonStatus(response), err
}

// ExpandVolume : extend a volume if there is enough space in the disk group
func (client *Client) ExpandVolume(name, size string) (*common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)
	response, httpRes, err := client.apiClient.DefaultApi.ExpandVolumeSizeNameGet(client.Ctx, size, name).Execute()
	logger.V(2).Info("expand volume", "name", name, "size", size, "http", httpRes.Status)
	return CreateCommonStatus(response), err
}

// CopyVolume : create an new volume by copying another one or a snapshot
func (client *Client) CopyVolume(sourceName string, destinationName string, pool string) (*common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)
	response, httpRes, err := client.apiClient.DefaultApi.CopyVolumeDestinationPoolNameSourceGet(client.Ctx, pool, destinationName, sourceName).Execute()
	logger.V(2).Info("copy volume", "destination", destinationName, "source", sourceName, "pool", pool, "http", httpRes.Status)
	return CreateCommonStatus(response), err
}

// DeleteVolume : deletes a volume
func (client *Client) DeleteVolume(name string) (*common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)
	response, httpRes, err := client.apiClient.DefaultApi.DeleteVolumesNamesGet(client.Ctx, name).Execute()
	logger.V(2).Info("delete volume", "name", name, "http", httpRes.Status)
	return CreateCommonStatus(response), err
}

// PublishVolume: Attach a volume to an initiator
func (client *Client) PublishVolume(volumeId string, initiators []string) (string, error) {

	logger := klog.FromContext(client.Ctx)

	hostNames, apistatus, err := client.GetVolumeMapsHostNames(volumeId)
	if err != nil {
		if apistatus != nil && apistatus.ReturnCode == common.VolumeNotFoundErrorCode {
			return "", status.Errorf(codes.NotFound, "The specified volume (%s) was not found.", volumeId)
		} else {
			return "", err
		}
	}
	for _, hostName := range hostNames {
		for _, initiator := range initiators {
			if hostName == initiator {
				logger.Info("volume is already mapped to initiator", "volume", volumeId, "initiator", initiators)
			}
		}
	}

	lun, err := client.chooseLUN(initiators)
	if err != nil {
		return "", err
	}

	logger.Info("using LUN", "lun", lun)

	mappingSuccessful := false
	for _, initiator := range initiators {
		if err = client.mapVolumeProcess(volumeId, initiator, lun); err != nil {
			klog.Errorf("error mapping volume (%s) for initiator (%s) using LUN (%d): %v", volumeId, initiators, lun, err)
		} else {
			mappingSuccessful = true
			logger.Info("successfully mapped", "volume", volumeId, "initiator", initiators, "lun", lun)
		}
	}

	if mappingSuccessful {
		return strconv.Itoa(lun), nil
	} else {
		return "", fmt.Errorf("error mapping volume (%s), no initiators were mapped successfully", volumeId)
	}
}
