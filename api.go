//
// Copyright (c) 2021 Seagate Technology LLC and/or its Affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// For any questions about this software or licensing,
// please email opensource@seagate.com or cortx-questions@seagate.com.

package exosx

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog"
)

// Configuration constants
const (
	MaximumLUN = 255
)

// Exos X Storage API Error Codes
const (
	snapshotNotFoundErrorCode             = -10050
	badInputParam                         = -10058
	hostMapDoesNotExistsErrorCode         = -10074
	volumeNotFoundErrorCode               = -10075
	volumeHasSnapshot                     = -10183
	snapshotAlreadyExists                 = -10186
	initiatorNicknameOrIdentifierNotFound = -10386
	unmapFailedErrorCode                  = -10509
)

type VolumeMapInfo struct {
	Volume       string
	Exists       bool
	SerialNumber string
	Mappings     []VolumeMapItem
}

type VolumeMapItem struct {
	InitiatorId string
	LUN         string
	Access      string
	Ports       string
	Nickname    string
	Profile     string
}

type InitiatorMapInfo struct {
	InitiatorId string
	Nickname    string
	Profile     string
	Mappings    []InitiatorMapItem
}

type InitiatorMapItem struct {
	Volume       string
	SerialNumber string
	LUN          string
	Access       string
	Ports        string
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

// InitSystemInfo: Retrieve and store system information for this client
func (client *Client) InitSystemInfo() error {

	err := AddSystem(client.Addr, client)
	if err != nil {
		return fmt.Errorf("unable to add system info for ip (%s) ", client.Addr)
	}

	client.Info, err = GetSystem(client.Addr)
	if err == nil {
		_ = client.Info.Log()
	}

	return err
}

// GetVolumeMaps2: Return a list of mapped initiators for a specified volume
func (client *Client) GetVolumeMaps2(volume string) (VolumeMapInfo, *ResponseStatus, error) {

	m := VolumeMapInfo{Volume: volume}
	original := volume

	if len(volume) > 0 {
		volume = fmt.Sprintf("\"%s\"", volume)
	}

	results, status, _ := client.FormattedRequest("/show/maps/%s", volume)

	if status.ReturnCode != 0 {
		klog.V(0).Infof("volume (%s) was not found", volume)
		m.Exists = false
		return m, status, nil
	}

	for _, rootObj := range results.Objects {
		if rootObj.Name == "volume-view" {
			id := rootObj.PropertiesMap["volume-name"].Data
			if id != original {
				klog.Warningf("Map data volume-name (%s) did not match parameter volume (%s)\n", id, volume)
			}
			m.Exists = true
			m.SerialNumber = rootObj.PropertiesMap["volume-serial"].Data

			m.Mappings = make([]VolumeMapItem, 0)
			var mi VolumeMapItem

			for _, object := range rootObj.Objects {
				if object.Name == "host-view" && object.PropertiesMap["identifier"].Data != "all other initiators" {
					mi = VolumeMapItem{
						InitiatorId: object.PropertiesMap["identifier"].Data,
						LUN:         object.PropertiesMap["lun"].Data,
						Access:      object.PropertiesMap["access"].Data,
						Ports:       object.PropertiesMap["ports"].Data,
						Nickname:    object.PropertiesMap["nickname"].Data,
						Profile:     object.PropertiesMap["host-profile"].Data,
					}
					m.Mappings = append(m.Mappings, mi)
				}
			}
		}
	}

	return m, status, nil
}

// LogVolumeMaps: Log all map information
func (client *Client) LogVolumeMaps(maps VolumeMapInfo) error {

	klog.V(0).Infof("Mapping for (%s)\n", maps.Volume)
	klog.V(0).Infof("-- SerialNumber: %s\n", maps.SerialNumber)
	klog.V(0).Infof("-- Exists: %v\n", maps.Exists)
	klog.V(0).Infof("-- Mappings:\n")

	klog.V(0).Infof("--       %-46s  %-4s  %-16s  %-8s  %-16s  %-12s\n", "Initiator", "LUN", "Access", "Ports", "Nickname", "Profile")
	klog.V(0).Infof("--       %-46s  %-4s  %-16s  %-8s  %-16s  %-12s\n", strings.Repeat("-", 46), strings.Repeat("-", 4), strings.Repeat("-", 16), strings.Repeat("-", 8), strings.Repeat("-", 16), strings.Repeat("-", 12))

	for i, m := range maps.Mappings {
		klog.V(0).Infof("-- [%3d] %-46s  %-4s  %-16s  %-8s  %-16s  %-12s\n", i, m.InitiatorId, m.LUN, m.Access, m.Ports, m.Ports, m.Nickname)
	}

	klog.V(0).Infof("\n")

	return nil
}

// GetInitiatorMaps: Return a list of mapped volumes for a specified initiator
func (client *Client) GetInitiatorMaps(initiator string) (InitiatorMapInfo, *ResponseStatus, error) {

	m := InitiatorMapInfo{InitiatorId: initiator}
	original := initiator

	if len(initiator) > 0 {
		initiator = fmt.Sprintf("\"%s\"", initiator)
	}

	results, status, err := client.FormattedRequest("/show/maps/%s", initiator)
	if err != nil {
		return m, status, err
	}

	for _, rootObj := range results.Objects {
		if rootObj.Name == "initiator-view" {
			id := rootObj.PropertiesMap["id"].Data
			if id != original {
				klog.Warningf("Map data id (%s) did not match parameter initiator (%s)\n", id, initiator)
			}
			m.Nickname = rootObj.PropertiesMap["hba-nickname"].Data
			m.Profile = rootObj.PropertiesMap["host-profile"].Data

			m.Mappings = make([]InitiatorMapItem, 0)
			var mi InitiatorMapItem

			for _, object := range rootObj.Objects {
				if object.Name == "volume-view" {
					mi = InitiatorMapItem{
						Volume:       object.PropertiesMap["volume"].Data,
						SerialNumber: object.PropertiesMap["volume-serial"].Data,
						LUN:          object.PropertiesMap["lun"].Data,
						Access:       object.PropertiesMap["access"].Data,
						Ports:        object.PropertiesMap["ports"].Data,
					}
					m.Mappings = append(m.Mappings, mi)
				}
			}
		}
	}

	return m, status, nil
}

// LogInitiatorMaps: Log all map information
func (client *Client) LogInitiatorMaps(maps InitiatorMapInfo) error {

	klog.V(0).Infof("Mapping for (%s)\n", maps.InitiatorId)
	klog.V(0).Infof("-- Nickname: %s\n", maps.Nickname)
	klog.V(0).Infof("-- Profile : %s\n", maps.Profile)
	klog.V(0).Infof("-- Mappings:\n")

	klog.V(0).Infof("--       %-32s  %-32s  %-4s  %-16s  %-8s\n", "Volume", "SerialNumber", "LUN", "Access", "Ports")
	klog.V(0).Infof("--       %-32s  %-32s  %-4s  %-16s  %-8s\n", strings.Repeat("-", 32), strings.Repeat("-", 32), strings.Repeat("-", 4), strings.Repeat("-", 16), strings.Repeat("-", 8))

	for i, m := range maps.Mappings {
		klog.V(0).Infof("-- [%3d] %-32s  %-32s  %-4s  %-16s  %-8s\n", i, m.Volume, m.SerialNumber, m.LUN, m.Access, m.Ports)
	}

	klog.V(0).Infof("\n")

	return nil
}

// GetVolumeMaps: Return a slice of mapped initiators for a specified volume
func (client *Client) GetVolumeMaps(volume string) ([]string, []string, *ResponseStatus, error) {
	if volume != "" {
		volume = fmt.Sprintf("\"%s\"", volume)
	}
	res, status, err := client.FormattedRequest("/show/maps/%s", volume)

	if err != nil {
		return []string{}, []string{}, status, err
	}

	initiators := []string{}
	luns := []string{}

	for _, rootObj := range res.Objects {
		if rootObj.Name != "volume-view" {
			continue
		}

		for _, object := range rootObj.Objects {
			//klog.Infof("%v", object)
			initiatorName := object.PropertiesMap["identifier"].Data
			lun := object.PropertiesMap["lun"].Data

			if object.Name == "host-view" && initiatorName != "all other initiators" {
				klog.V(2).Infof("map: volume (%s) --> initiator (%s) lun (%s)", volume, initiatorName, lun)
				initiators = append(initiators, initiatorName)
				luns = append(luns, lun)
			}
		}
	}

	return initiators, luns, status, err
}

// nextLUN: Determine the next available Loginal Unit Number (LUN), which is used for mapping avolume to an initiator
func (client *Client) nextLUN(maps InitiatorMapInfo) (int, error) {

	if len(maps.InitiatorId) == 0 {
		klog.V(0).Infof("initiator does not exist, no LUN mappings yet, using LUN 1")
		return 1, nil
	}

	luns := make([]int, 0)
	for _, m := range maps.Mappings {
		lun, err := strconv.Atoi(m.LUN)
		if err == nil {
			luns = append(luns, lun)
		}
	}
	sort.Ints(luns)

	klog.V(2).Infof("checking if LUN 1 is in use")
	if len(luns) == 0 || luns[0] > 1 {
		return 1, nil
	}

	klog.V(2).Infof("searching for an available LUN between LUNs in use")
	for index := 1; index < len(luns); index++ {
		if luns[index]-luns[index-1] > 1 {
			return luns[index-1] + 1, nil
		}
	}

	klog.V(2).Infof("checking if next LUN is not above maximum LUNs limit")
	if luns[len(luns)-1]+1 < MaximumLUN {
		return luns[len(luns)-1] + 1, nil
	}

	return -1, status.Error(codes.ResourceExhausted, "no LUN is available")
}

// chooseLUN: Choose the next available LUN for a given initiator
func (client *Client) chooseLUN(initiators []string) (int, error) {
	klog.Infof("listing all LUN mappings")

	var allvolumes []Volume
	for _, initiatorName := range initiators {
		volumes, responseStatus, err := client.ShowHostMaps(initiatorName)
		if err != nil {
			klog.Infof("error looking for host maps for initiator %s: %s", initiatorName, err)
		}
		if responseStatus.ReturnCode == hostMapDoesNotExistsErrorCode {
			klog.Infof("initiator %s does not exist", initiatorName)
		}
		if volumes != nil {
			allvolumes = append(allvolumes, volumes...)
		}
	}

	sort.Sort(Volumes(allvolumes))

	klog.V(5).Infof("use LUN 1 when volumes slice is empty")
	if len(allvolumes) == 0 {
		return 1, nil
	}

	klog.V(5).Infof("use the next highest LUN number, until the end is reached")
	if allvolumes[len(allvolumes)-1].LUN+1 < MaximumLUN {
		return allvolumes[len(allvolumes)-1].LUN + 1, nil
	}

	klog.V(5).Infof("use LUN 1 when not in use")
	if allvolumes[0].LUN > 1 {
		return 1, nil
	}

	klog.V(5).Infof("use the next available LUN, searching from LUN 1 towards the maximum")
	for index := 1; index < len(allvolumes); index++ {
		// Find a gap between used LUNs
		if allvolumes[index].LUN-allvolumes[index-1].LUN > 1 {
			return allvolumes[index-1].LUN + 1, nil
		}
	}

	klog.Errorf("no available LUN: [%d] luns=%v", len(allvolumes), allvolumes)

	return -1, status.Error(codes.ResourceExhausted, "no more available LUNs")
}

// mapVolumeProcess: Map a volume to an initiator and create a nickname when required by the storage array
func (client *Client) mapVolumeProcess(volumeName, initiatorName string, lun int) error {
	klog.Infof("trying to map volume %s for initiator %s on LUN %d", volumeName, initiatorName, lun)
	_, metadata, err := client.MapVolume(volumeName, initiatorName, "rw", lun)
	if err != nil && metadata == nil {
		return err
	}

	klog.Infof("status: metadata.ReturnCode=%v", metadata.ReturnCode)
	if metadata.ReturnCode == initiatorNicknameOrIdentifierNotFound {
		nodeIDParts := strings.Split(initiatorName, ":")
		if len(nodeIDParts) < 2 {
			return status.Error(codes.NotFound, "specified node ID is not a valid IQN")
		}

		nickname := strings.Join(nodeIDParts[1:], ":")
		nickname = strings.ReplaceAll(nickname, ".", "-")

		klog.Infof("initiator does not exist, creating it with nickname %s", nickname)
		_, _, err = client.CreateNickname(nickname, initiatorName)
		if err != nil {
			return err
		}
		klog.Info("retrying to map volume")
		_, _, err = client.MapVolume(volumeName, initiatorName, "rw", lun)
		if err != nil {
			return err
		}
	} else if metadata.ReturnCode == volumeNotFoundErrorCode {
		return status.Errorf(codes.NotFound, "volume %s not found", volumeName)
	} else if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

// CheckVolumeExists: Return true if a volume already exists
func (client *Client) CheckVolumeExists(volumeID string, size int64) (bool, error) {
	data, responseStatus, err := client.ShowVolumes(volumeID)
	if err != nil && responseStatus.ReturnCode != badInputParam {
		return false, err
	}

	for _, object := range data.Objects {
		if object.Name == "volume" {
			if object.PropertiesMap["volume-name"].Data == volumeID {
				blocks, _ := strconv.ParseInt(object.PropertiesMap["blocks"].Data, 10, 64)
				blocksize, _ := strconv.ParseInt(object.PropertiesMap["blocksize"].Data, 10, 64)
				klog.V(3).Infof("volume exists: checking (%s) size (%d) against blocksize (%d)", volumeID, size, blocksize)

				if blocks*blocksize == size {
					return true, nil
				}
				return true, status.Error(codes.AlreadyExists, "cannot create volume with same name but different capacity than the existing one")
			}
		}
	}

	return false, nil
}

// PublishVolume: Attach a volume to an initiator
func (client *Client) PublishVolume(volumeId string, initiators []string) (string, error) {

	hostNames, apistatus, err := client.GetVolumeMapsHostNames(volumeId)
	if err != nil {
		if apistatus != nil && apistatus.ReturnCode == volumeNotFoundErrorCode {
			return "", status.Errorf(codes.NotFound, "The specified volume (%s) was not found.", volumeId)
		} else {
			return "", err
		}
	}
	for _, hostName := range hostNames {
		for _, initiator := range initiators {
			if hostName == initiator {
				klog.Infof("volume %s is already mapped to initiator %s", volumeId, initiators)
			}
		}
	}

	lun, err := client.chooseLUN(initiators)
	if err != nil {
		return "", err
	}

	klog.Infof("using LUN %d", lun)

	for _, initiator := range initiators {
		//TODO optimize this for fewer API calls
		if err = client.mapVolumeProcess(volumeId, initiator, lun); err != nil {
			return "", err
		}
	}

	klog.Infof("successfully mapped volume (%s) for initiator (%s) using LUN (%d)", volumeId, initiators, lun)
	return strconv.Itoa(lun), nil
}

// GetVolumeWwn: Retrieve the WWN for a volume, very useful for host operating system device mapping
func (client *Client) GetVolumeWwn(volumeName string) (string, error) {

	wwn := ""
	response, status, err := client.ShowVolumes(volumeName)
	if err == nil && status.ResponseTypeNumeric == 0 {
		if response.ObjectsMap["volume"] != nil {
			wwn = strings.ToLower(response.ObjectsMap["volume"].PropertiesMap["wwn"].Data)
		}
	}

	klog.V(3).Infof("GetVolumeWwn (%s) returning wwn (%s)", volumeName, wwn)
	return wwn, err
}
