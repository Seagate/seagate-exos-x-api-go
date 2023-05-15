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

package test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/Seagate/seagate-exos-x-api-go/pkg/common"
	storageapi "github.com/Seagate/seagate-exos-x-api-go/pkg/exosx"
	. "github.com/onsi/gomega"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"k8s.io/klog/v2"
)

//
// An API Test Suite
//
// Goal: Execute various Storage Array API calls used by other drivers and validating
//       those calls against all supported storage array API versions. Also validate
//       the acquistion and use of system information needed for various operations.
//
// API Calls Tested:
//     /login
//     /show/versions/detail
//     /show/controllers
//     /show/pools
//     /create/volume
//     /show/volumes
//     /show/initiators/
//     /show/maps
//     /show/maps/initiators/
//     /map/volume
//     /expand/volume
//     /copy/volume
//     /show/snapshots
//     /create/snapshots
//     /delete/snapshot
//     /unmap/volume
//     /delete/volumes
//     /set/initiator - called if an initiator nickname is needed
//

var size = "1GiB"
var sizeint int64 = 1024 * 1024 * 1024
var volname1 = "apitest_1"
var volname2 = "apitest_2"
var expandSize = "1GiB"
var snap1 = "snap1"
var snap2 = "snap2"
var loginFail = false
var poolType = ""
var initiatorNick = "test-nickname"

// ShowVolume: Display useful data from a volume object
func ShowVolume(t *testing.T, volumeName string) {
	g := NewWithT(t)

	response, status, err := client.ShowVolumes(volumeName)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	if len(response) > 0 {
		p := message.NewPrinter(language.English)
		p.Printf("\n")
		p.Printf("    volume-name       = %v\n", response[0].VolumeName)
		p.Printf("    storage-pool-name = %v\n", response[0].StoragePoolName)
		p.Printf("    blocksize         = %v\n", response[0].BlockSize)
		p.Printf("    blocks            = %d\n", response[0].Blocks)
		p.Printf("    current size      = %d\n", response[0].BlockSize*response[0].Blocks)
		p.Printf("    tier-affinity     = %v\n", response[0].TierAffinity)
		p.Printf("\n")
	}
}

// ShowVolumes: Display useful information for all volume objects allocated
func ShowVolumes(t *testing.T) {
	g := NewWithT(t)

	response, status, err := client.ShowVolumes()
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	if err == nil {
		fmt.Printf("\n")
		fmt.Printf("Volumes:\n")
		for _, object := range response {
			if object.ObjectName == "volume" {
				fmt.Printf("%8v, %32v, %10v, %8v, %10v, %v, %8v, %v\n",
					object.StoragePoolName,
					object.VolumeName,
					object.TotalSize,
					object.Blocks,
					object.BlockSize,
					object.StorageType,
					object.VolumeType,
					object.Health,
				)
			}
		}
	}

	fmt.Printf("\n")
}

// ShowSnapshot: Display useful information for a snapshot object
func ShowSnapshot(t *testing.T, name string) {
	g := NewWithT(t)

	response, status, err := client.ShowSnapshots(name, "")
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	if err == nil {
		fmt.Printf("\n")
		fmt.Printf("Snapshots:\n")
		for _, object := range response {
			if object.Name == "snapshots" {
				fmt.Printf("%8v, %32v, %32v\n",
					object.StoragePoolName,
					object.Name,
					object.VolumeParent,
				)
			}
		}
	}

	fmt.Printf("\n")
}

// ConditionalSkip: Skip test case if the log in failed
func ConditionalSkip(t *testing.T) {
	if loginFail {
		t.Skip()
	}
}
func TestAPILogin(t *testing.T) {
	g := NewWithT(t)
	err := client.Login()
	if err != nil {
		loginFail = true
	}
	g.Expect(err).To(BeNil())
}

func TestAPISystemInfo(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	err := storageapi.AddSystem(client.Addr, client)
	g.Expect(err).To(BeNil())

	client.Info, err = storageapi.GetSystem(client.Addr)
	g.Expect(err).To(BeNil())

	err = client.Info.Log()
	g.Expect(err).To(BeNil())

	// Store the pool type for use throughout the test cases
	poolType, _ = client.Info.GetPoolType(client.PoolName)

}

func TestAPICreateVolume(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	fmt.Printf("create volume (%s)\n", volname1)

	status, err := client.CreateVolume(volname1, size, client.PoolName, poolType)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	ShowVolume(t, volname1)
	ShowVolumes(t)

	var exists bool
	exists, err = client.CheckVolumeExists(volname1, sizeint)
	fmt.Printf("CheckVolumeExists(%s, %d) returned %v, %v\n", volname1, sizeint, exists, err)
	g.Expect(err).To(BeNil())
	g.Expect(exists).To(Equal(true))

	exists, err = client.CheckVolumeExists("unknown-volume", sizeint)
	fmt.Printf("CheckVolumeExists(%s, %d) returned %v, %v\n", volname1, sizeint, exists, err)
	g.Expect(err).To(BeNil())
	g.Expect(exists).To(Equal(false))
}

func TestAPIShowInitiators(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	response, status, err := client.FormattedRequest("/show/initiators/")
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	fmt.Printf("\n")
	fmt.Printf("%-22s %-11s %-7s %-9s %-11s %-11s\n", "Nickname", "Discovered", "Mapped", "Profile", "Host", "ID")

	if err == nil {
		for _, obj := range response.Objects {
			if obj.Name == "initiator" {
				fmt.Printf("%-22s %-11s %-7s %-9s %-11s %-11s\n",
					obj.PropertiesMap["nickname"].Data,
					obj.PropertiesMap["discovered"].Data,
					obj.PropertiesMap["mapped"].Data,
					obj.PropertiesMap["profile"].Data,
					obj.PropertiesMap["host-bus-type"].Data,
					obj.PropertiesMap["id"].Data)
			}
		}
	}

	fmt.Printf("\n")
}

func volumeSlicesEqual(vol1 []storageapi.Volume, vol2 []storageapi.Volume) bool {
	if len(vol1) != len(vol2) {
		return false
	}
	sort.Sort(storageapi.Volumes(vol1))
	sort.Sort(storageapi.Volumes(vol2))
	for i, vol := range vol1 {
		if vol.LUN != vol2[i].LUN {
			return false
		}
	}
	return true
}

// Test that "show host maps" returns the same set of LUNS with and
// without an initiator nickname
func TestAPIShowHostMapsWithNickname(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	response, status, err := client.ShowHostMaps(client.Initiator)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	fmt.Printf("Show Host Maps Response: %v", response)
	fmt.Printf("\n")

	// Retrieve initial initiator nickname if it has one
	resp, status, err := client.FormattedRequest(fmt.Sprintf("/show/initiator/%s", client.Initiator))
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	originalNickname := resp.ObjectsMap["initiator"].PropertiesMap["nickname"].Data

	client.FormattedRequest(fmt.Sprintf("/set/initiator/id/%s/nickname/%s", client.Initiator, initiatorNick))
	// If the initiator had a nickname before this test, restore it. If not, delete the test nickname
	if originalNickname != "" {
		defer client.FormattedRequest(fmt.Sprintf("/set/initiator/id/%s/nickname/%s", client.Initiator, originalNickname))
	} else {
		defer client.FormattedRequest(fmt.Sprintf("/delete/initiator-nickname/%s", initiatorNick))
	}

	response2, status, err := client.ShowHostMaps(client.Initiator)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	g.Expect(volumeSlicesEqual(response, response2)).To(BeTrue())
	fmt.Printf("Show Host Maps Response: %v", response)
	fmt.Printf("\n")
}

// Test that "show host maps" returns the same set of LUNS with and
// without host groups defined
func TestAPIShowHostMapsWithHostGroups(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	response, status, err := client.ShowHostMaps(client.Initiator)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	fmt.Printf("Show Host Maps Response: %v", response)
	fmt.Printf("\n")

	client.FormattedRequest(fmt.Sprintf("/set/initiator/id/%s/nickname/%s", client.Initiator, initiatorNick))
	defer client.FormattedRequest(fmt.Sprintf("/delete/initiator-nickname/%s", initiatorNick))

	response2, status, err := client.ShowHostMaps(client.Initiator)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	g.Expect(volumeSlicesEqual(response, response2)).To(BeTrue())
	fmt.Printf("Show Host Maps Response: %v", response)
	fmt.Printf("\n")
}

func TestAPIGetMaps(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	respone, status, err := client.FormattedRequest("/show/maps/%s", volname1)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	hostNames := []string{}

	if err == nil {
		for _, rootObj := range respone.Objects {
			if rootObj.Name != "volume-view" {
				continue
			}

			for _, object := range rootObj.Objects {
				hostName := object.PropertiesMap["identifier"].Data
				if object.Name == "host-view" && hostName != "all other initiators" {
					hostNames = append(hostNames, hostName)
				}
			}
		}
	}

	fmt.Printf("volume %q host names:\n", volname1)
	for i, h := range hostNames {
		fmt.Printf("    [%d] %s\n", i, h)
	}
}

func TestAPIMapVolume(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)
	lun, err := client.PublishVolume(volname1, []string{client.Initiator})
	g.Expect(err).To(BeNil())
	g.Expect(lun).ToNot(Equal(0))
}

func TestAPIExpandVolume(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	klog.Infof("expand volume (%s) from original size (%s) to new size (%s)", volname1, size, expandSize)
	status, err := client.ExpandVolume(volname1, expandSize)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("expand volume failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	klog.Infof("successfully expanded volume (%s)", volname1)
	ShowVolume(t, volname1)
}

func TestAPICreateSnapshots(t *testing.T) {
	ConditionalSkip(t)

	if poolType == "Linear" {
		fmt.Printf("Linear snapshots are not supported\n")
		return
	}

	g := NewWithT(t)

	klog.Infof("snapshot volume (%s) using name (%s)", volname1, snap1)
	status, err := client.CreateSnapshot(volname1, snap1)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("snapshot volume failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	klog.Infof("successfully snapped volume (%s)", snap1)
	ShowSnapshot(t, snap1)
	ShowVolumes(t)

	klog.Infof("snapshot volume (%s) using name (%s)", volname1, snap2)
	status, err = client.CreateSnapshot(volname1, snap2)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("snapshot volume failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	klog.Infof("successfully snapped volume (%s)", snap2)
	ShowSnapshot(t, snap2)
	ShowVolumes(t)
}

func TestAPIDeleteSnapshots(t *testing.T) {
	ConditionalSkip(t)

	if poolType == "Linear" {
		fmt.Printf("Linear snapshots are not supported\n")
		return
	}

	g := NewWithT(t)

	klog.Infof("delete snapshot (%s)", snap1)
	status, err := client.DeleteSnapshot(snap1)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("delete snapshot failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	klog.Infof("successfully deleted snapshot (%s)", snap1)

	klog.Infof("delete snapshot (%s)", snap2)
	status, err = client.DeleteSnapshot(snap2)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("delete snapshot failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	klog.Infof("successfully deleted snapshot (%s)", snap2)
	ShowVolumes(t)
}

func TestAPIUnmapVolume(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	klog.Infof("unmapping volume %s from initiator %s", volname1, client.Initiator)
	status, err := client.UnmapVolume(volname1, client.Initiator)
	if err != nil {
		if status != nil && status.ReturnCode == common.UnmapFailedErrorCode {
			fmt.Printf("unmap failed, assuming volume is already unmapped")
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	klog.Infof("successfully unmapped volume %s from initiator %s", volname1, client.Initiator)
}

func TestAPICopyVolume(t *testing.T) {
	ConditionalSkip(t)

	if poolType == "Linear" {
		fmt.Printf("Linear snapshots are not supported\n")
		return
	}

	g := NewWithT(t)

	klog.Infof("copy volume (%s) to (%s) using pool (%s)", volname1, volname1, client.PoolName)

	status, err := client.CopyVolume(volname1, volname2, client.PoolName)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("copy volume failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	klog.Infof("successfully copied volume to (%s)", volname2)
	ShowVolumes(t)
}

func TestAPIDeleteVolumes(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	status, err := client.DeleteVolume(volname1)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	if poolType != "Linear" {
		status, err := client.DeleteVolume(volname2)
		g.Expect(err).To(BeNil())
		g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	}

	ShowVolumes(t)
}
