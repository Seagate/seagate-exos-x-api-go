// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Seagate/seagate-exos-x-api-go/v2/pkg/regression"

	"k8s.io/klog/v2"
)

const VERSION = "v1.0.0"

func stringVar(p *string, name string, usage string) {
	flag.StringVar(p, name, *p, usage)
}

func intVar(p *int, name string, usage string) {
	flag.IntVar(p, name, *p, usage)
}

type testing struct {
	result int
}

func (t *testing) Fail() {
	t.result = 1
}

func main() {

	// Flags for version and debug level
	version := flag.Bool("version", false, "print the version of this program")
	debug := flag.String("debug", "", "set the logging debug level")
	config := flag.String("config", "", "set the configuration file")

	// Parse flags
	//klog.InitFlags(nil)
	flag.Parse()

	if *version {
		fmt.Printf("api-regression (%s)\n", VERSION)
		os.Exit(0)
	}

	if *debug != "" {
		fmt.Printf("setting debug level to (%s)\n", *debug)
		level := klog.Level(0)
		level.Set(*debug)
	}

	// Use default configuration file or user specified value
	configFilename := "api-regression.conf"
	if *config != "" {
		configFilename = *config
	}
	fmt.Printf("using config file (%s)\n", configFilename)

	// Read config info from yaml file and set up configuration instance
	apiConfig, err := regression.NewTestConfig(configFilename)
	if err != nil {
		fmt.Printf("ERROR: reading config file (%s), err: %v\n", configFilename, err)
		os.Exit(1)
	}

	// Create testing object
	t := testing{}

	// Run regression tests
	regression.Test(&t, apiConfig, klog.FromContext(apiConfig.Ctx))

	os.Exit(t.result)
}
