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
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/Seagate/seagate-exos-x-api-go/pkg/common"
	storageapi "github.com/Seagate/seagate-exos-x-api-go/pkg/exosx"
	"github.com/joho/godotenv"
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
//       the retrieval and use of system information needed for various operations.
//
// API Calls Tested:
//
//     client.Login()
//     client.SessionValid(apiAddr, username)
//
//     client.InitSystemInfo()
//     client.GetPoolType(pool)
//     client.GetPortals()
//     client.Info.GetTargetId("iSCSI")
//
//     client.ShowVolumes(volumeName)
//     client.CheckVolumeExists(volumeName, size)
//     client.CreateVolume(volumeName, sizeStr, parameters[common.PoolConfigKey], poolType)
//     client.GetVolumeWwn(volumeName)
//     client.PublishVolume(volumeName, initiators)
//     client.CopyVolume(sourceName, volumeName, parameters[common.PoolConfigKey])
//     client.UnmapVolume(volumeName, initiator)
//     client.ExpandVolume(volumeName, expansionSizeStr)
//     client.DeleteVolume(volumeName)
//
//     client.ShowSnapshots(snapshotName, "")
//     client.ShowSnapshots(req.SnapshotId, sourceVolumeId)
//     client.CreateSnapshot(sourceVolumeId, snapshotName)
//     client.DeleteSnapshot(req.SnapshotId)

const (
	expectedPoolType = "Virtual"
)

var size = "1GiB"
var sizeint int64 = 1024 * 1024 * 1024
var volname1 = "apitest_1"
var volname2 = "apitest_2"
var expandSize = "1GiB"
var snap1 = "snap1"
var snap2 = "snap2"
var loginFail = false
var poolType = ""

var client = storageapi.NewClient()

var logging = flag.Bool("logging", true, "logging")

func init() {
	var exists bool
	settingsfile := ".env"

	fmt.Printf("Testing setup: \n")
	fmt.Printf("    logging (%v)\n", *logging)

	if *logging {
		klog.InitFlags(nil)
	}

	// Note, any defined environment variable is used over the ones defined in .env
	if _, err := os.Stat(settingsfile); err == nil {
		fmt.Printf("    loading (%s)\n", settingsfile)
		err := godotenv.Load(settingsfile)
		if err != nil {
			fmt.Printf("Error loading file (%s), error: %v\n", settingsfile, err)
			return
		}
	}

	client.Addr, exists = os.LookupEnv("TEST_STORAGEIP")
	if exists {
		fmt.Printf("    %s=%s\n", "TEST_STORAGEIP", client.Addr)
	}

	client.Username, exists = os.LookupEnv("TEST_USERNAME")
	if exists {
		fmt.Printf("    %s=%s\n", "TEST_USERNAME", client.Username)
	}

	client.Password, exists = os.LookupEnv("TEST_PASSWORD")
	if exists {
		fmt.Printf("    %s=%s\n", "TEST_PASSWORD", client.Password)
	}

	client.Initiator, exists = os.LookupEnv("TEST_INITIATOR")
	if exists {
		fmt.Printf("    %s=%s\n", "TEST_INITIATOR", client.Initiator)
	}

	client.PoolName, exists = os.LookupEnv("TEST_POOL")
	if exists {
		fmt.Printf("    %s=%s\n", "TEST_POOL", client.PoolName)
	}

	fmt.Printf("================================================================================\n")
}

// ShowVolume: Display useful data from a volume object
func ShowVolume(t *testing.T, volumeName string) {
	g := NewWithT(t)

	response, status, err := client.ShowVolumes(volumeName)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	if len(response) > 0 {
		p := message.NewPrinter(language.English)
		p.Printf("\n")
		fmt.Printf("Show Volume:\n")
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
		fmt.Printf("Show Volumes:\n")
		for _, object := range response {
			fmt.Printf("%8v, %8v, %32v, %10v, %8v, %10v, %v, %8v, %v\n",
				object.ObjectName,
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
		fmt.Printf("Show Snapshots:\n")
		for _, object := range response {
			fmt.Printf("%8v, %32v, %32v\n",
				object.StoragePoolName,
				object.Name,
				object.VolumeParent,
			)
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

	// Validate that the session is valid
	valid := client.SessionValid(client.Addr, client.Username)
	g.Expect(valid).To(Equal(true))
	g.Expect(valid).To(BeTrue())
}

func TestAPISystemInfo(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	err := client.InitSystemInfo()
	g.Expect(err).To(BeNil())

	err = client.Info.Log()
	g.Expect(err).To(BeNil())

	// Test and store the pool type for use throughout the test cases
	poolType, err = client.Info.GetPoolType(client.PoolName)
	fmt.Printf("    Pool: %s\n", poolType)
	g.Expect(err).To(BeNil())
	g.Expect(poolType).To(Equal(expectedPoolType))

	// Test getting portals and target id
	portals, err := client.Info.GetPortals()
	fmt.Printf("    Portals: %s (%v)\n", portals, err)
	//g.Expect(err).To(BeNil())

	targetid, err := client.Info.GetTargetId("iSCSI")
	fmt.Printf("    Target: %s (%v)\n", targetid, err)
	//g.Expect(err).To(BeNil())
}

func TestAPICreateVolume(t *testing.T) {
	ConditionalSkip(t)
	g := NewWithT(t)

	fmt.Printf("create volume (%s)\n", volname1)

	ShowVolumes(t)

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

	wwn, err := client.GetVolumeWwn(volname1)
	fmt.Printf("GetVolumeWwn(%s) returned %v, %v\n", volname1, wwn, err)
	g.Expect(err).To(BeNil())

	lun, err := client.PublishVolume(volname1, []string{client.Initiator})
	fmt.Printf("PublishVolume(%s) returned %v, %v\n", volname1, lun, err)
	g.Expect(err).To(BeNil())
	g.Expect(lun).ToNot(Equal(0))

	fmt.Printf("unmapping volume %s from initiator %s\n", volname1, client.Initiator)
	status, err = client.UnmapVolume(volname1, client.Initiator)
	if err != nil {
		if status != nil && status.ReturnCode == common.UnmapFailedErrorCode {
			fmt.Printf("unmap failed, assuming volume is already unmapped\n")
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	fmt.Printf("successfully unmapped volume %s from initiator %s\n", volname1, client.Initiator)

	fmt.Printf("expand volume (%s) from original size (%s) to new size (%s)\n", volname1, size, expandSize)
	status, err = client.ExpandVolume(volname1, expandSize)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("expand volume failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	fmt.Printf("successfully expanded volume (%s)\n", volname1)
	ShowVolume(t, volname1)

	fmt.Printf("copy volume (%s) to (%s) using pool (%s)\n", volname1, volname1, client.PoolName)
	status, err = client.CopyVolume(volname1, volname2, client.PoolName)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("copy volume failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	fmt.Printf("successfully copied volume to (%s)\n", volname2)
	ShowVolumes(t)

	status, err = client.DeleteVolume(volname1)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	status, err = client.DeleteVolume(volname2)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

}

func TestAPICreateSnapshots(t *testing.T) {
	ConditionalSkip(t)

	if poolType == "Linear" {
		fmt.Printf("WARNING: Linear snapshots are not supported\n")
		return
	}

	g := NewWithT(t)

	status, err := client.CreateVolume(volname1, size, client.PoolName, poolType)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))

	fmt.Printf("snapshot volume (%s) using name (%s)\n", volname1, snap1)
	status, err = client.CreateSnapshot(volname1, snap1)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("snapshot volume failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	fmt.Printf("successfully snapped volume (%s)\n", snap1)
	ShowSnapshot(t, snap1)
	ShowVolumes(t)

	fmt.Printf("snapshot volume (%s) using name (%s)\n", volname1, snap2)
	status, err = client.CreateSnapshot(volname1, snap2)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("snapshot volume failed, status.ReturnCode=%v", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	fmt.Printf("successfully snapped volume (%s)\n", snap2)
	ShowSnapshot(t, snap2)
	ShowVolumes(t)

	if poolType == "Linear" {
		fmt.Printf("Linear snapshots are not supported\n")
		return
	}

	fmt.Printf("delete snapshot (%s)\n", snap1)
	status, err = client.DeleteSnapshot(snap1)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("delete snapshot failed, status.ReturnCode=%v\n", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	fmt.Printf("successfully deleted snapshot (%s)\n", snap1)

	fmt.Printf("delete snapshot (%s)\n", snap2)
	status, err = client.DeleteSnapshot(snap2)
	if err != nil {
		if status != nil && status.ReturnCode != 0 {
			fmt.Printf("delete snapshot failed, status.ReturnCode=%v\n", status.ReturnCode)
		}
	}
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
	fmt.Printf("successfully deleted snapshot (%s)\n", snap2)
	ShowVolumes(t)

	status, err = client.DeleteVolume(volname1)
	g.Expect(err).To(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(0))
}
