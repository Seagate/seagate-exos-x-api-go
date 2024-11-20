// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package mcapi

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/Seagate/seagate-exos-x-api-go/v2/pkg/client"
	"github.com/Seagate/seagate-exos-x-api-go/v2/pkg/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

const (
	ApiMaximumLUN            = 255
	ApiTierAffinity          = "no-affinity"
	ApiSuccess               = 0
	ApiError                 = 1
	ApiInfo                  = 2
	ApiWarning               = 3
	UngroupedHostsGroupID    = "HGU"
	UngroupedInitiatorHostID = "HU"
)

// Volume : a mapped volume
type Volume struct {
	LUN       int
	Name      string
	Initiator string
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

// UpdateVolumeObject: transfer results to the common Volume object
func UpdateVolumeObject(target *common.VolumeObject, source *client.VolumesResourceInner) {
	target.ObjectName = source.GetObjectName()
	target.Blocks = source.GetBlocks()
	target.BlockSize = source.GetBlocksize()
	target.Health = source.GetHealth()
	target.SizeNumeric = source.GetSizeNumeric()
	target.StoragePoolName = source.GetStoragePoolName()
	target.StorageType = source.GetStorageType()
	target.TierAffinity = source.GetTierAffinity()
	target.TotalSize = source.GetTotalSize()
	target.VolumeName = source.GetVolumeName()
	target.VolumeType = source.GetVolumeType()
	target.Wwn = strings.ToLower(source.GetWwn())
}

// CreateVolume : creates a volume with the given name, capacity in the given pool
func (client *Client) CreateVolume(name, size, pool string) (*common.VolumeObject, *common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)
	response, status, httpRes, err := ExecuteWithFailover(client.apiClient.DefaultApi.CreateVolumePoolSizeTierAffinityNameGet(client.Ctx, pool, size, ApiTierAffinity, name).Execute, client)
	logger.V(2).Info("create volume", "name", name, "http", httpRes.Status)

	volume := common.VolumeObject{}
	if response != nil && len(response.GetVolumes()) > 0 {
		UpdateVolumeObject(&volume, &response.Volumes[0])
	}

	// For API versions that do not return a representation of the volume object, use ShowVolumes to fill in the data
	if err == nil && status.ResponseTypeNumeric == 0 && volume.Wwn == "" {
		volumes, status, err := client.ShowVolumes(name)
		if err == nil && status.ResponseTypeNumeric == 0 {
			if len(volumes) > 0 && volumes[0].ObjectName == "volume" {
				volume = common.VolumeObject(volumes[0])
			}
		}
	}

	return &volume, status, err
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

	response, responseStatus, httpRes, err := ExecuteWithFailover(client.apiClient.DefaultApi.ShowVolumesNamesGet(client.Ctx, volume).Execute, client)

	if httpRes.StatusCode == http.StatusOK && response != nil && err == nil {
		if logger.V(4).Enabled() {
			// Extract Status resource information
			logger.V(4).Info("++ ShowVolumesNamesGet.Status[0]",
				"Response", responseStatus.Response,
				"ReturnCode", responseStatus.ReturnCode,
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
		}
	} else {
		logger.V(4).Info("-- ShowVolumesNamesGet", "status", httpRes.Status, "err", err, "body", httpRes.Body)
		return nil, &common.ResponseStatus{}, err
	}

	logger.V(4).Info("================================================================================")

	returnVolumes := []common.VolumeObject{}

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
		volume.Wwn = strings.ToLower(v.GetWwn())

		returnVolumes = append(returnVolumes, volume)
	}

	return returnVolumes, responseStatus, err
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

// Given an initiator ID or nickname, get associated host and hostgroups if they exist
func (myclient *Client) GetInitiatorHostGroup(initiator string) (hostGroup string, host string, err error) {
	response, _, httpRes, err := ExecuteWithFailover(myclient.apiClient.DefaultApi.ShowHostGroupsGet(myclient.Ctx).Execute, myclient)
	if err != nil {
		return
	}
	if httpRes.StatusCode == http.StatusOK && response != nil {
		for _, hg := range response.GetHostGroup() {
			for _, h := range hg.GetHost() {
				for _, i := range h.GetInitiator() {
					if i.GetId() == initiator || i.GetNickname() == initiator {
						// return nil for host group if host is in the "ungrouped hosts" group, host group ID == "HGU"
						if hg.GetDurableId() != UngroupedHostsGroupID {
							hostGroup = hg.GetName()
						}
						// return nil for host if the initiator host id is the ungrouped initiators id == "HU"
						if h.GetDurableId() != UngroupedInitiatorHostID {
							host = h.GetName()
						}
						return
					}
				}
			}
		}
	}
	err = fmt.Errorf("initiator not found in show host groups (%s)", initiator)
	return
}

// ShowHostMaps: list the volume mappings for given host
func (myclient *Client) ShowHostMaps(initiator string) ([]Volume, *common.ResponseStatus, error) {

	logger := klog.FromContext(myclient.Ctx)

	logger.Info("++ ShowHostMaps", "initiator", initiator)

	hostgroup, hostname, err := myclient.GetInitiatorHostGroup(initiator)
	if err != nil {
		return nil, nil, err
	}

	//If an initiator is part of a host or host-group, show maps must be called with the Host or Host-Group name
	var host string
	if hostgroup != "" {
		host = fmt.Sprintf("%s.*.*", hostgroup)
	} else if hostname != "" {
		host = fmt.Sprintf("%s.*", hostname)
	} else {
		host = initiator
	}
	rawResponse, apistatus, _, err := ExecuteWithFailover(myclient.apiClient.DefaultApi.ShowMapsInitiatorNamesGet(myclient.Ctx, host).Execute, myclient)
	if err != nil {
		return nil, apistatus, err
	}

	response := rawResponse.GetActualInstance()

	// Get the HostViewMappings (present in all return types from ShowMapsInitiator)
	// For each volume view in the HVM, append to the mappings list we will return
	mappings := make([]Volume, 0, 32)
	appendMappingsFn := func(viewObject interface {
		GetHostViewMappings() []client.HostViewMappingsResourceInner
	}) {
		for _, hvm := range viewObject.GetHostViewMappings() {
			if hvm.GetObjectName() == "volume-view" {
				vol := Volume{}
				vol.LUN, _ = strconv.Atoi(hvm.GetLun())
				vol.Name = hvm.GetVolume()
				// don't use the hvm initiator name/id here, keep the raw initiator ID passed in
				vol.Initiator = initiator
				logger.Info("++ volume", "volume", hvm.GetVolumeName(), "lun", vol.LUN)
				mappings = append(mappings, vol)
			}
		}
	}

	switch resp := response.(type) {
	// No mappings on this system, so only a status was returned
	case *client.StatusObject:
		break
	case *client.HostGroupViewObject:
		for _, hgv := range resp.GetHostGroupView() {
			appendMappingsFn(&hgv)
		}
	case *client.HostsViewObject:
		for _, hv := range resp.GetHostsView() {
			appendMappingsFn(&hv)
		}
	case *client.InitiatorViewObject:
		for _, hv := range resp.GetInitiatorView() {
			appendMappingsFn(&hv)
		}
	}
	return mappings, apistatus, err
}

// chooseLUN: Choose the next available LUN for a given initiator. Optional volumeID to look for existing LUNs for that volume
func (client *Client) chooseLUN(initiators []string, volumeID string) (int, error) {

	logger := klog.FromContext(client.Ctx)
	logger.Info("listing all LUN mappings")

	var allvolumes []Volume
	for _, initiator := range initiators {
		logger.V(4).Info("Searching for maps for initiator", "initiatorName", initiator)
		volumes, _, err := client.ShowHostMaps(initiator)
		if err != nil {
			logger.Error(err, "error looking for host maps", "initiator", initiator)
		}
		if volumes != nil {
			allvolumes = append(allvolumes, volumes...)
		}
	}

	sort.Sort(Volumes(allvolumes))

	if volumeID != "" {
		lun := -1
		for _, vol := range allvolumes {
			if vol.Name == volumeID {
				if lun < 1 {
					lun = vol.LUN
				} else if lun != vol.LUN {
					return -1, fmt.Errorf("found multiple LUNs (%d, %d) for volume %s", lun, vol.LUN, volumeID)
				}
			}
		}

		if lun > 0 {
			logger.V(5).Info("using LUN from existing mapping", "lun", lun)
			return lun, nil
		}
	}

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
	_, respStatus, httpRes, err := ExecuteWithFailover(client.apiClient.DefaultApi.MapVolumeAccessLunInitiatorNamesGet(client.Ctx, access, strconv.Itoa(lun), initiator, name).Execute, client)
	logger.V(2).Info("map volume", "name", name, "lun", lun, "initiator", initiator, "http", httpRes.Status)
	return respStatus, err
}

// CreateNickname : Create a nickname for an initiator. The Storage API policy is to prohibit mapping of initiators which are not either
// (a) presently connected to the array or (b) represented by an entry in the initiator nickname table.
func (client *Client) CreateNickname(name, iqn string) (*common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)
	_, response, httpRes, err := ExecuteWithFailover(client.apiClient.DefaultApi.SetInitiatorIdNicknameGet(client.Ctx, iqn, name).Execute, client)
	logger.V(2).Info("create nickname", "name", name, "nickname", iqn, "http", httpRes.Status)
	return response, err
}

// mapVolumeProcess: Map a volume to an initiator and create a nickname when required by the storage array
func (client *Client) mapVolumeProcess(volumeName, initiatorName string, lun int) error {

	logger := klog.FromContext(client.Ctx)
	logger.V(0).Info("trying to map volume", "volume", volumeName, "initiator", initiatorName, "lun", lun)
	metadata, err := client.MapVolume(volumeName, initiatorName, "rw", lun)
	if err != nil && metadata == nil {
		return err
	}

	logger.Info("status", "ReturnCode", metadata.ReturnCode)
	if metadata.ReturnCode == common.VolumeNotFoundErrorCode {
		return status.Errorf(codes.NotFound, "volume %s not found", volumeName)
	} else if metadata.ReturnCode == common.LUNOverlapErrorCode {
		return status.Errorf(codes.AlreadyExists, "lun overlap for lun: %d", lun)
	} else if metadata.ReturnCode == common.InitiatorNicknameOrIdentifierNotFound {
		return status.Errorf(codes.AlreadyExists, "specified initiator for mapping not found: %s", initiatorName)
	} else if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

// UnmapVolume : unmap a volume from an initiator
func (client *Client) UnmapVolume(name, initiator string) (*common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)

	if initiator == "" {
		_, status, httpRes, err := ExecuteWithFailover(client.apiClient.DefaultApi.UnmapVolumeNamesGet(client.Ctx, name).Execute, client)
		logger.V(2).Info("unmap volume", "name", name, "initiator", initiator, "http", httpRes.Status)
		return status, err
	}

	_, status, httpRes, err := ExecuteWithFailover(client.apiClient.DefaultApi.UnmapVolumeInitiatorNamesGet(client.Ctx, initiator, name).Execute, client)
	logger.V(2).Info("unmap volume", "name", name, "initiator", initiator, "http", httpRes.Status)
	return status, err
}

// ExpandVolume : extend a volume if there is enough space in the disk group
func (client *Client) ExpandVolume(name, size string) (*common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)
	_, status, httpRes, err := ExecuteWithFailover(client.apiClient.DefaultApi.ExpandVolumeSizeNameGet(client.Ctx, size, name).Execute, client)
	logger.V(2).Info("expand volume", "name", name, "size", size, "http", httpRes.Status)
	return status, err
}

// CopyVolume : create an new volume by copying another one or a snapshot
func (client *Client) CopyVolume(sourceName string, destinationName string, pool string) (*common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)
	_, status, httpRes, err := ExecuteWithFailover(client.apiClient.DefaultApi.CopyVolumeDestinationPoolNameSourceGet(client.Ctx, pool, destinationName, sourceName).Execute, client)
	logger.V(2).Info("copy volume", "destination", destinationName, "source", sourceName, "pool", pool, "http", httpRes.Status)
	return status, err
}

// DeleteVolume : deletes a volume
func (client *Client) DeleteVolume(name string) (*common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)
	_, status, httpRes, err := ExecuteWithFailover(client.apiClient.DefaultApi.DeleteVolumesNamesGet(client.Ctx, name).Execute, client)
	logger.V(2).Info("delete volume", "name", name, "http", httpRes.Status)
	return status, err
}

// PublishVolume: Attach a volume to an initiator, return mapped lun or error
func (client *Client) PublishVolume(volumeId string, initiators []string) (string, error) {

	logger := klog.FromContext(client.Ctx)

	lun, err := client.chooseLUN(initiators, volumeId)
	if err != nil {
		return "", err
	}

	logger.Info("using LUN", "lun", lun)

	mappingSuccessful := false
	for _, initiator := range initiators {
		if err = client.mapVolumeProcess(volumeId, initiator, lun); err != nil {
			logger.Error(err, "mapping error", "volume", volumeId, "initiator", initiator, "LUN", lun)
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
