// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	storageapi "github.com/Seagate/seagate-exos-x-api-go/pkg/v1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/klog/v2"
)

var _ = DescribeRegression("System Testing (v1)", func(tc *TestContext) {
	var (
		client *storageapi.Client = nil
	)

	BeforeEach(func() {
		By("setup client for testing")

		logger := klog.FromContext(tc.Config.Ctx)
		client = storageapi.NewClient()
		client.StoreCredentials(tc.Config.StorageController.Ip, tc.Config.StorageController.Protocol, tc.Config.StorageController.Username, tc.Config.StorageController.Password)
		err := client.Login()
		logger.V(3).Info("Login", "ipaddress", client.Addr, "username", client.Username, "err", err)
		Expect(err).To(BeNil())
	})

	Describe("v1SystemTest", func() {
		It("should successfully init the system info", func() {

			logger := klog.FromContext(tc.Config.Ctx)
			err := client.InitSystemInfo()
			logger.V(3).Info("InitSystemInfo", "err", err)
			Expect(err).To(BeNil())

			err = storageapi.Log(client.Info)
			Expect(err).To(BeNil())
		})

		It("should be able to get a valid pool type", func() {

			logger := klog.FromContext(tc.Config.Ctx)
			err := client.InitSystemInfo()
			Expect(err).To(BeNil())
			poolType, err := storageapi.GetPoolType(client.Info, tc.Config.StorageController.Pool)
			logger.V(3).Info("GetPoolType", "pool", tc.Config.StorageController.Pool, "type", poolType)
			Expect(err).To(BeNil())
			Expect(poolType).ShouldNot(Equal(""))
		})

		It("should be able to get valid portals", func() {

			logger := klog.FromContext(tc.Config.Ctx)
			err := client.InitSystemInfo()
			Expect(err).To(BeNil())
			portals, err := storageapi.GetPortals(client.Info)
			logger.V(3).Info("GetPortals", "portals", portals)
			Expect(err).To(BeNil())
			Expect(portals).ShouldNot(Equal(""))
		})

		It("should be able to get valid target id", func() {

			logger := klog.FromContext(tc.Config.Ctx)
			err := client.InitSystemInfo()
			Expect(err).To(BeNil())
			targetid, err := storageapi.GetTargetId(client.Info, "iSCSI")
			logger.V(3).Info("GetTargetId", "targetid", targetid)
			Expect(err).To(BeNil())
			Expect(targetid).ShouldNot(Equal(""))
		})

	})
})
