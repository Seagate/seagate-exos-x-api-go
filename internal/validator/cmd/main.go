// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Seagate/seagate-exos-x-api-go/internal/validator"

	"github.com/Seagate/seagate-exos-x-api-go/pkg/client"
	"github.com/Seagate/seagate-exos-x-api-go/pkg/common"

	"k8s.io/klog/v2"
)

const (
	version = "1.1.1"
)

func main() {

	config, err := common.InitConfig(os.Args)
	if err != nil {
		fmt.Printf("unable to init configuration")
		os.Exit(1)
	}

	// Use contextual logging with a context that stores basic authentication
	ctx := context.WithValue(context.Background(), client.ContextBasicAuth, client.BasicAuth{
		UserName: config.MCUsername,
		Password: config.MCPassword,
	})

	klog.EnableContextualLogging(true)
	logger := klog.FromContext(ctx)

	// Set verbosity level according to ...
	var l klog.Level
	l.Set(config.Verbosity)

	logger.V(0).Info("")
	logger.V(0).Info("[] validator", "version", version)
	logger.V(0).Info("")

	apiClient, err := common.Login(ctx, config)
	logger.V(4).Info("Login", "apiClient", apiClient, "err", err)

	if err != nil || apiClient == nil {
		logger.Error(err, "++ login FAILURE")
		os.Exit(2)
	}

	validator.Meta(ctx, apiClient, "status")
	validator.ShowVersionsDetail(ctx, apiClient)
	validator.ShowSystems(ctx, apiClient)
	validator.ShowHostGroups(ctx, apiClient)

	validator.ShowVolumes(ctx, apiClient, "")

	pool := "A"
	size := "1GiB"
	volumeId := "myVolume1"

	validator.CreateVolume(ctx, apiClient, pool, size, "no-affinity", volumeId)
	validator.ShowVolumes(ctx, apiClient, volumeId)
	validator.DeleteVolumes(ctx, apiClient, volumeId)
	validator.ShowVolumes(ctx, apiClient, "")
}
