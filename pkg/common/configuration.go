// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package common

import (
	"github.com/namsral/flag"
)

const (
	DefaultSpecFile = "generated-mc-openapi.yaml" // Default OpenAPI specification file
	DefaultProtocol = "https"                     // Default MC API Protocol
)

type Config struct {
	Verbosity         string `json:"verbosity"`
	ConfigurationFile string `json:"configuration"`
	SpecificationFile string `json:"specification"`
	MCIpAddress       string `json:"mc-ip"`
	MCProtocol        string `json:"mc-protocol"` // User can specify http, https
	MCUsername        string `json:"mc-username"`
	MCPassword        string `json:"mc-password"`
	MCDescription     string `json:"mc-description"`
}

// Init: initialize the configuration data using command line args, ENV, or a file
func InitConfig(args []string) (*Config, error) {

	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	flags.String(flag.DefaultConfigFlagname, "", "Path to config file")

	var (
		verbosity     = flags.String("verbosity", "0", "Log level verbosity")
		configuration = flags.String("configuration", "", "MC Configuration YAML File which lists commands to add to spec")
		specification = flags.String("specification", DefaultSpecFile, "Generated MC OpenAPI Specification File")
		mcipaddress   = flags.String("mcipaddress", "", "MC IP Address")
		mcprotocol    = flags.String("mcprotocol", DefaultProtocol, "MC IP Protocol: http, https")
		mcusername    = flags.String("mcusername", "", "MC Username")
		mcpassword    = flags.String("mcpassword", "", "MC Password")
		description   = flags.String("description", "MC Storage System", "Brief name of the storage system for reference")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return nil, err
	}

	c := &Config{
		Verbosity:         *verbosity,
		ConfigurationFile: *configuration,
		SpecificationFile: *specification,
		MCIpAddress:       *mcipaddress,
		MCProtocol:        *mcprotocol,
		MCUsername:        *mcusername,
		MCPassword:        *mcpassword,
		MCDescription:     *description,
	}

	return c, nil
}
