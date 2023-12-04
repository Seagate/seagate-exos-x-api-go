// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Seagate/seagate-exos-x-api-go/v2/internal/generator"
	"github.com/Seagate/seagate-exos-x-api-go/v2/pkg/client"
	"github.com/Seagate/seagate-exos-x-api-go/v2/pkg/common"

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
	logger.V(0).Info("[] generator", "version", version)
	logger.V(0).Info("")

	// Open mc-configuration yaml and process
	logger.V(0).Info(">> process mc-configuration", "filename", config.ConfigurationFile)
	yamlc, err := generator.ReadYamlConfigurationFile(ctx, config.ConfigurationFile)
	if err != nil {
		logger.Error(err, "read yaml configuration file", "filename", config.ConfigurationFile)
		os.Exit(2)
	}
	logger.V(4).Info("++ mc-configuration", "contents", *yamlc)

	// Create the openapi specification file
	logger.V(0).Info(">> generate openapi spec header...")
	s := &generator.Specification{}
	s.Reset()
	s.GenerateHeader(ctx)
	s.GenerateLoginLogout(ctx)
	s.GenerateMetaSchema(ctx)

	logger.V(0).Info(">> generate openapi spec...", "commands", len(yamlc.Commands), "exceptions", len(yamlc.Exceptions))

	for _, command := range yamlc.Commands {
		err := s.AddCommand(ctx, config, &command, yamlc.Exceptions)
		if err != nil {
			logger.Error(err, "add command", "command", command)
			os.Exit(3)
		}
	}

	// Save the specification to a file
	s.Write(ctx, config.SpecificationFile, false)
}
