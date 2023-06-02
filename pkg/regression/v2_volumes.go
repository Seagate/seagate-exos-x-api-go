// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	storageapi "github.com/Seagate/seagate-exos-x-api-go/pkg/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/klog/v2"
)

var _ = DescribeRegression("Volume Testing (v2)", func(tc *TestContext) {
	var (
		client     *storageapi.Client = nil
		volname1                      = "apitest_1"
		volname2                      = "apitest_2"
		size                          = "1GiB"
		expandSize                    = "1GiB"
		poolType                      = "Virtual"
		sizeValue  int64              = 1024 * 1024 * 1024
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

	Describe("v2VolumeTest", func() {
		It("should successfully create the volume", func() {

			logger := klog.FromContext(tc.Config.Ctx)
			volume, status, err := client.CreateVolume(volname1, size, tc.Config.StorageController.Pool, poolType)
			logger.V(3).Info("CreateVolume", "name", volname1, "size", size, "wwn", volume.Wwn, "mc-response", status.ResponseTypeNumeric)

			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))

			// Dump volumes
			response, status, err := client.ShowVolumes(volname1)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))
			ShowVolumes(logger, response)
		})

		It("created volume should exist", func() {

			logger := klog.FromContext(tc.Config.Ctx)
			exists, err := client.CheckVolumeExists(volname1, sizeValue)
			logger.V(3).Info("CheckVolumeExists", "name", volname1, "exists", exists)
			Expect(err).To(BeNil())
			Expect(exists).To(Equal(true))
		})

		It("unknown volume should NOT exist", func() {

			logger := klog.FromContext(tc.Config.Ctx)
			exists, err := client.CheckVolumeExists("unknown-volume", sizeValue)
			logger.V(3).Info("CheckVolumeExists", "name", volname1, "exists", exists)
			Expect(err).To(BeNil())
			Expect(exists).To(Equal(false))
		})

		It("created volume should have valid WWN", func() {

			logger := klog.FromContext(tc.Config.Ctx)
			wwn, err := client.GetVolumeWwn(volname1)
			logger.V(3).Info("GetVolumeWwn", "name", volname1, "wwn", wwn)
			Expect(err).To(BeNil())
			Expect(wwn).ShouldNot(Equal(""))
		})

		It("should be able to publish/map volume", func() {

			logger := klog.FromContext(tc.Config.Ctx)
			lun, err := client.PublishVolume(volname1, []string{client.Initiator})
			logger.V(3).Info("PublishVolume", "name", volname1, "lun", lun)
			Expect(err).To(BeNil())
			Expect(lun).ToNot(Equal(0))
		})

		It("should be able to unmap volume", func() {

			logger := klog.FromContext(tc.Config.Ctx)
			status, err := client.UnmapVolume(volname1, tc.Config.StorageController.Initiator)
			logger.V(3).Info("UnmapVolume", "name", volname1, "mc-response", status.ResponseTypeNumeric)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))
		})

		It("should be able to expand the volume", func() {

			logger := klog.FromContext(tc.Config.Ctx)
			status, err := client.ExpandVolume(volname1, expandSize)
			logger.V(3).Info("ExpandVolume", "name", volname1, "mc-response", status.ResponseTypeNumeric)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))

			// Show volumes
			response, status, err := client.ShowVolumes(volname1)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))
			ShowVolumes(logger, response)
		})

		It("should be able to copy the volume", func() {

			logger := klog.FromContext(tc.Config.Ctx)
			status, err := client.CopyVolume(volname1, volname2, tc.Config.StorageController.Pool)
			logger.V(3).Info("CopyVolume", "from", volname1, "to", volname2, "mc-response", status.ResponseTypeNumeric)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))

			// Show volumes
			response, status, err := client.ShowVolumes(volname2)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiSuccess))
			ShowVolumes(logger, response)
		})

		It("should successfully delete both volumes", func() {

			logger := klog.FromContext(tc.Config.Ctx)

			status, err := client.DeleteVolume(volname1)
			logger.V(3).Info("delete volume", "name", volname1, "mc-response", status.ResponseTypeNumeric, "err", err)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiVolumeDeleted))

			status, err = client.DeleteVolume(volname2)
			logger.V(3).Info("delete volume", "name", volname2, "mc-response", status.ResponseTypeNumeric, "err", err)
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiVolumeDeleted))

			// Show volumes
			response, status, err := client.ShowVolumes("volname1")
			Expect(err).To(BeNil())
			Expect(status.ResponseTypeNumeric).To(Equal(storageapi.ApiVolumeDoesNotExist))
			ShowVolumes(logger, response)
		})
	})
})
