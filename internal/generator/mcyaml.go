// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package generator

import (
	"context"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
	"k8s.io/klog/v2"
)

const (
	exceptionRename = "rename"
)

type GeneralInformation struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

type NestedMeta struct {
	Name   string       `yaml:"name"`   // The name of the meta object, such as 'hosts-groups'
	Nested []NestedMeta `yaml:"nested"` // When one or more response objects are nested into the main response object, such as 'redundancy' is nested under 'system'
}

type OptionInformation struct {
	Flag            string `yaml:"flag"`             // The command line parameter option
	Type            string `yaml:"type"`             // The option type: string, enum, abs
	Values          string `yaml:"values"`           // Optional. Allowed enumeration values
	Required        string `yaml:"required"`         // Required is true or false
	KeywordRequired string `yaml:"keyword-required"` // Bool. Indicates keyword should be used in path, or must not be used in path
	Description     string `yaml:"description"`      // Optional. Command description
}

type CommandInformation struct {
	Path    string              `yaml:"path"`    // The command path, such as '/show/system'
	Meta    string              `yaml:"meta"`    // The meta resource, which may be different from the main command name, such as 'drives'
	Include string              `yaml:"include"` // When a response object includes other response types at the root level, such as 'status'
	Nested  []NestedMeta        `yaml:"nested"`  // When one or more response objects are nested into the main response object, such as 'redundancy' is nested under 'system'
	Options []OptionInformation `yaml:"options"` // One or more command line options
}

type ExceptionInformation struct {
	Resource    string `yaml:"resource"`     // The meta resource which has an issue
	Type        string `yaml:"type"`         // The type of exception to handle
	OldProperty string `yaml:"old-property"` // For a rename exception, the old property name
	NewProperty string `yaml:"new-property"` // For a rename exception, the new property name
}

type ConfigurationYaml struct {
	Header     string                 `yaml:"mc-configuration"`
	Info       GeneralInformation     `yaml:"info"`
	Commands   []CommandInformation   `yaml:"commands"`
	Exceptions []ExceptionInformation `yaml:"exceptions"`
}

// ReadYamlConfigurationFile: Read configuration.yaml
func ReadYamlConfigurationFile(ctx context.Context, filename string) (*ConfigurationYaml, error) {

	logger := klog.FromContext(ctx)
	logger.V(3).Info("read mc-configuration.yaml", "filename", filename)

	var yamlc = ConfigurationYaml{}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read yaml configuration file (%s)", filename)
	}

	err = yaml.Unmarshal(yamlFile, &yamlc)
	if err != nil {
		return nil, fmt.Errorf("unable to Unmarshal yaml configuration file (%s)", filename)
	}

	return &yamlc, nil
}
