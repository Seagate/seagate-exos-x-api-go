// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	storageapi "github.com/Seagate/seagate-exos-x-api-go/pkg/mcapi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/klog/v2"
)

var _ = DescribeRegression("Snapshot Testing (v2)", func(tc *TestContext) {
	var (
		client   *storageapi.Client = nil
		volname1                    = "apitest_1"
		size                        = "1GiB"
		poolType                    = "Virtual"
		snap1                       = "snap1"
		snap2                       = "snap2"
	)

	BeforeEach(func() {
		By("setup client for testing")

		logger := klog.FromContext(tc.Config.Ctx)
		client = storageapi.NewClient()
		client.StoreCredentials(tc.Config.StorageController.Ip, tc.Config.StorageController.Username, tc.Config.StorageController.Password)
		err := client.Login(tc.Config.Ctx)
		logger.V(3).Info("Login", "ipaddress", client.Addr, "username", client.Username, "err", err)
		Expect(err).To(BeNil())
	})

	Describe("v2SnapshotTest", func() {
		It("should successfully create the volume", func() {

			logger := klog.FromContext(tc.Config.Ctx)
			status, err := client.CreateVolume(volname1, size, tc.Config.StorageController.Pool, poolType)
			logger.V(3).Info("CreateVolume", "name", volname1, "size", size, "mc-response", status.ResponseTypeNumeric)

			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))

			// Dump volumes
			response, status, err := client.ShowVolumes(volname1)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))
			ShowVolumes(logger, response)
		})

		It("should successfully create a snapshot", func() {

			logger := klog.FromContext(tc.Config.Ctx)

			status, err := client.CreateSnapshot(volname1, snap1)
			logger.V(3).Info("CreateSnapshot", "name", volname1, "snap1", snap1, "mc-response", status.ResponseTypeNumeric)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))

			// Show snapshots
			snapshots, status, err := client.ShowSnapshots(snap1, "")
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))
			ShowSnapshots(logger, snapshots)

			// Show volumes
			volumes, status, err := client.ShowVolumes(volname1)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))
			ShowVolumes(logger, volumes)
		})

		It("should successfully create a 2nd snapshot", func() {

			logger := klog.FromContext(tc.Config.Ctx)

			status, err := client.CreateSnapshot(volname1, snap2)
			logger.V(3).Info("CreateSnapshot", "name", volname1, "snap2", snap2, "mc-response", status.ResponseTypeNumeric)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))

			// Show snapshots
			snapshots, status, err := client.ShowSnapshots(snap1, "")
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))
			ShowSnapshots(logger, snapshots)

			// Show volumes
			volumes, status, err := client.ShowVolumes(volname1)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))
			ShowVolumes(logger, volumes)
		})

		It("should successfully delete both snapshots", func() {

			logger := klog.FromContext(tc.Config.Ctx)

			status, err := client.DeleteSnapshot(snap1)
			logger.V(3).Info("delete snapshot", "name", snap1, "mc-response", status.ResponseTypeNumeric, "err", err)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))

			status, err = client.DeleteSnapshot(snap2)
			logger.V(3).Info("delete snapshot", "name", snap2, "mc-response", status.ResponseTypeNumeric, "err", err)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))

			// Show snapshots
			snapshots, status, err := client.ShowSnapshots(snap1, "")
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))
			ShowSnapshots(logger, snapshots)
		})

		It("should successfully delete the volume", func() {

			logger := klog.FromContext(tc.Config.Ctx)

			status, err := client.DeleteVolume(volname1)
			logger.V(3).Info("delete volume", "name", volname1, "mc-response", status.ResponseTypeNumeric, "err", err)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiVolumeDeleted))

			// Show volumes
			response, status, err := client.ShowVolumes(volname1)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiVolumeDoesNotExist))
			ShowVolumes(logger, response)
		})
	})
})
