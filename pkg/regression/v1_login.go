// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	storageapi "github.com/Seagate/seagate-exos-x-api-go/pkg/v1"
	"k8s.io/klog/v2"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeRegression("Login Testing (v1)", func(tc *TestContext) {

	var (
		client *storageapi.Client
	)

	Describe("v1LoginTest", func() {

		It("should be able to log into storage controller", func() {
			client = storageapi.NewClient()
			client.StoreCredentials(tc.Config.StorageController.Ip, tc.Config.StorageController.Username, tc.Config.StorageController.Password)
			err := client.Login()
			logger := klog.FromContext(tc.Config.Ctx)
			logger.V(3).Info("Login", "ip", client.Addr, "username", client.Username, "err", err)
			Expect(err).To(BeNil())
		})

		It("should be a valid session", func() {
			valid := client.SessionValid(client.Addr, client.Username)
			Expect(valid).To(Equal(true))
		})

	})
})
