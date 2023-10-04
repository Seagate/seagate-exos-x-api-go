// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package generator

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/Seagate/seagate-exos-x-api-go/pkg/common"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"k8s.io/klog/v2"
)

const (
	maxArrayCount = 1024
)

type Section struct {
	lines []string
}

type Specification struct {
	header             Section         // openapi, info, servers
	paths              Section         // paths
	parameters         Section         // components: parameters
	resources          Section         // components: schemas, resources
	objects            Section         // components: schemas, objects
	generatedResources map[string]bool // tracks which resources have been generated
	generatedObjects   map[string]bool // tracks which objects have been generated
	generatedOptions   map[string]bool // tracks which command options have been generated
}

// Header: add a new line to the openapi: section of the specification
func (s *Specification) Header(level int, line string) {
	s.header.lines = append(s.header.lines, fmt.Sprintf("%s%s", strings.Repeat("  ", level), line))
}

// Paths: add a new line to the paths: section of the specification
func (s *Specification) Paths(level int, line string) {
	s.paths.lines = append(s.paths.lines, fmt.Sprintf("%s%s", strings.Repeat("  ", level), line))
}

// Parameters: add a new line to the components/parameters section of the specification
func (s *Specification) Parameters(level int, line string) {
	s.parameters.lines = append(s.parameters.lines, fmt.Sprintf("%s%s", strings.Repeat("  ", level), line))
}

// Resources: add a new line to the components/schemas-resources section of the specification
func (s *Specification) Resources(level int, line string) {
	s.resources.lines = append(s.resources.lines, fmt.Sprintf("%s%s", strings.Repeat("  ", level), line))
}

// Objects: add a new line to the components/schemas-objects section of the specification
func (s *Specification) Objects(level int, line string) {
	s.objects.lines = append(s.objects.lines, fmt.Sprintf("%s%s", strings.Repeat("  ", level), line))
}

// Reset: reset all sections of the specification object
func (s *Specification) Reset() {
	s.header.lines = nil
	s.paths.lines = nil
	s.parameters.lines = nil
	s.resources.lines = nil
	s.objects.lines = nil
	s.generatedResources = make(map[string]bool)
	s.generatedObjects = make(map[string]bool)
	s.generatedOptions = make(map[string]bool)
}

// Write: write all sections to create an openapi specification yaml file
func (s *Specification) Write(ctx context.Context, filename string, append bool) error {

	var f *os.File
	var err error

	logger := klog.FromContext(ctx)
	logger.V(1).Info("write specification", "filename", filename)

	if append {
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		f, err = os.Create(filename)
	}
	if err != nil {
		return fmt.Errorf("failure to create/open file (%s), err: %v", filename, err)
	}

	defer f.Close()

	// Write out the openapi: header section
	for line, value := range s.header.lines {
		logger.V(4).Info(fmt.Sprintf("[%d][%s] %s", line, "header", value))
		//f.WriteString(fmt.Sprintf("%s\n", value))
		fmt.Fprintln(f, value)
	}

	// Write out the paths section
	if len(s.paths.lines) > 0 {
		fmt.Fprintln(f, "")
		fmt.Fprintln(f, "paths:")
		for line, value := range s.paths.lines {
			logger.V(4).Info(fmt.Sprintf("[%d][%s] %s", line, "paths", value))
			fmt.Fprintln(f, value)
		}
	}

	// Write out the components/parameters section
	if len(s.parameters.lines) > 0 {
		fmt.Fprintln(f, "")
		fmt.Fprintln(f, "# Security, common parameters, and schemas used by this API")
		fmt.Fprintln(f, "security:")
		fmt.Fprintln(f, "- basicAuth: []")
		fmt.Fprintln(f, "")
		fmt.Fprintln(f, "components:")
		fmt.Fprintln(f, "  securitySchemes:")
		fmt.Fprintln(f, "    basicAuth:")
		fmt.Fprintln(f, "      type: http")
		fmt.Fprintln(f, "      scheme: basic")
		fmt.Fprintln(f, "  parameters:")
		for line, value := range s.parameters.lines {
			logger.V(4).Info(fmt.Sprintf("[%d][%s] %s", line, "parameters", value))
			fmt.Fprintln(f, value)
		}
	}

	// The schema is separated into two parts:
	//   resources that are typically an array of properties, named <tag>Resource
	//   objects that reference a resource schema, named <tag>Object

	// Write out the components/schemas-resources section
	if len(s.resources.lines) > 0 {
		fmt.Fprintln(f, "")
		fmt.Fprintln(f, "  schemas:")
		for line, value := range s.resources.lines {
			logger.V(4).Info(fmt.Sprintf("[%d][%s] %s", line, "resources", value))
			fmt.Fprintln(f, value)
		}
	}

	// Write out the components/schemas-objects section
	if len(s.objects.lines) > 0 {
		fmt.Fprintln(f, "")
		for line, value := range s.objects.lines {
			logger.V(4).Info(fmt.Sprintf("[%d][%s] %s", line, "objects", value))
			fmt.Fprintln(f, value)
		}
	}

	f.Sync()

	return nil
}

// GenerateOpenAPIHeader: Add all required lines for the openapi header section
func (s *Specification) GenerateHeader(ctx context.Context) error {

	s.Header(0, "openapi: 3.0.0")
	s.Header(0, "")
	s.Header(0, "info:")
	s.Header(1, "title: Seagate Management Controller (MC) OpenAPI")
	s.Header(1, "description: |")
	s.Header(2, "This API allows users to interact through the MC API to provision and manage storage.")
	s.Header(1, "license:")
	s.Header(2, "name: Apache-2.0")
	s.Header(2, "url: https://www.apache.org/licenses/LICENSE-2.0.html")
	s.Header(1, "version: 1.0.0")
	s.Header(0, "")
	s.Header(0, "servers:")
	s.Header(0, "- url: http://localhost:8080/api")
	s.Header(0, "- url: https://localhost:8080/api")

	return nil
}

// GenerateLogin: Add all required lines for MC login path
func (s *Specification) GenerateLogin(ctx context.Context) error {

	s.Paths(1, "# Login")
	s.Paths(1, "/login:")
	s.Paths(2, "get:")
	s.Paths(3, "description: Log in to the storage array management controller")
	s.Paths(3, "operationId: LoginGet")
	s.Paths(3, "security:")
	s.Paths(3, "- basicAuth: []")
	s.Paths(3, "responses:")
	s.Paths(4, "'200':")
	s.Paths(5, "description: OK")
	s.Paths(5, "content:")
	s.Paths(6, "application/json:")
	s.Paths(7, "schema:")
	s.Paths(8, "$ref: '#/components/schemas/statusObject'")

	s.generatedResources["status"] = true
	s.Resources(2, "statusResource:")
	s.Resources(3, "type: array")
	s.Resources(3, "minItems: 1")
	s.Resources(3, fmt.Sprintf("maxItems: %d", maxArrayCount))
	s.Resources(3, "items:")
	s.Resources(4, "type: object")
	s.Resources(4, "properties:")

	s.Resources(5, "object-name:")
	s.Resources(6, "type: string")
	s.Resources(5, "meta:")
	s.Resources(6, "type: string")
	s.Resources(5, "component-id:")
	s.Resources(6, "type: string")
	s.Resources(5, "response:")
	s.Resources(6, "type: string")
	s.Resources(5, "response-type:")
	s.Resources(6, "type: string")
	s.Resources(6, "enum: [Success, Error, Info, Warning]")
	s.Resources(6, "description: Indicates whether the command returned a Success, Error, Info, or Warning message")
	s.Resources(5, "response-type-numeric:")
	s.Resources(6, "type: integer")
	s.Resources(6, "description: Indicates whether the command returned a Success (0), Error (1), Info (2), or Warning (3) message (numeric form)")
	s.Resources(5, "return-code:")
	s.Resources(6, "type: integer")
	s.Resources(6, "description: Numeric return code for the command")
	s.Resources(5, "time-stamp:")
	s.Resources(6, "type: string")
	s.Resources(6, "description: Time at which this event was detected")
	s.Resources(5, "time-stamp-numeric:")
	s.Resources(6, "type: integer")
	s.Resources(6, "description: Time at which this event was detected( In numeric form )")

	s.generatedObjects["status"] = true
	s.Objects(2, "statusObject:")
	s.Objects(3, "type: object")
	s.Objects(3, "properties:")
	s.Objects(4, "status:")
	s.Objects(5, "$ref: '#/components/schemas/statusResource'")

	return nil
}

// GenerateMetaSchema: Add required lines to support the meta API command which provides a JSON representation of a resource
func (s *Specification) GenerateMetaSchema(ctx context.Context) error {

	s.Paths(1, "# Meta")
	s.Paths(1, "/meta/{schemaId}:")
	s.Paths(2, "get:")
	s.Paths(3, "description: Retrieve schema definition for a resource")
	s.Paths(3, "operationId: SchemaGet")
	s.Paths(3, "parameters:")
	s.Paths(3, "- $ref: '#/components/parameters/schemaId'")
	s.Paths(3, "responses:")
	s.Paths(4, "'200':")
	s.Paths(5, "description: OK")
	s.Paths(5, "content:")
	s.Paths(6, "application/json:")
	s.Paths(7, "schema:")
	s.Paths(8, "$ref: '#/components/schemas/metaObject'")

	s.Parameters(2, "schemaId:")
	s.Parameters(3, "description: A string label for a resource")
	s.Parameters(3, "name: schemaId")
	s.Parameters(3, "in: path")
	s.Parameters(3, "required: true")
	s.Parameters(3, "schema:")
	s.Parameters(4, "type: string")
	s.Parameters(4, "example: 'volumes'")

	s.Objects(2, "metaObject:")
	s.Objects(3, "description: 'Can be anything: string, number, array, object, etc.'")

	return nil
}

// AddMetaProperties: retrieve meta data for the schema and adds properties to specification object in memory
// This routine is called recursively for each level of nested resources, until the last Nested slice is nil
func (s *Specification) AddMetaProperties(ctx context.Context, config *common.Config, meta string, nested []NestedMeta, exceptions []ExceptionInformation) error {

	// Look up the meta string name in a table of booleans storing whether the meta data has been generated or not
	generated := s.generatedResources[meta]
	logger := klog.FromContext(ctx)
	logger.V(1).Info("add meta properties", "meta", meta, "generated", generated, "nested", nested)

	// If generated already, return success
	if generated {
		return nil
	}

	// If specified in the input yaml, add all definitions for all nested resources
	if len(nested) > 0 {
		for _, n := range nested {
			logger.V(3).Info("insert nested", "meta", n)
			err := s.AddMetaProperties(ctx, config, n.Name, n.Nested, exceptions)
			if err != nil {
				return fmt.Errorf("failed to include properties for (%s)", n)
			}
		}
	}

	// Log in to the MC
	apiClient, err := common.Login(ctx, config)
	if err != nil || apiClient == nil {
		logger.Error(err, "login", "ipaddress", config.MCIpAddress)
		return fmt.Errorf("login failed for (%s) at (%s)", config.MCDescription, config.MCIpAddress)
	}

	// Run http get /api/meta/{metaId} and process the JSON response data
	response, httpRes, err := apiClient.DefaultApi.SchemaGet(ctx, meta).Execute()

	if httpRes.StatusCode != http.StatusOK {
		logger.V(2).Info("-- SchemaGet", "status", httpRes.Status, "err", err, "body", httpRes.Body)
		return fmt.Errorf("SchemaGet failure, status=%v, err=%v", httpRes.Status, err)
	}

	// Validate that we did receive a 'map[string]interface{}' value
	value, ok := response.(map[string]interface{})
	if !ok {
		return fmt.Errorf("unexpected response from SchemaGet, type=%v", reflect.TypeOf(response))
	}

	// Parse the JSON data and extract all object properties
	// Example:
	// {
	// 	"components": {
	// 		"schemas": {
	// 			"status": {
	// 				"type": "object",
	// 				"properties": {
	// 					"response-type": {
	// 						"type": "string",
	// 						"description": "Indicates whether the command returned a success, failure, or informational message"
	// 					},
	// 					"response-type-numeric": {
	// 						"type": "integer",
	// 						"description": "Indicates whether the command returned a success, failure, or informational message( In numeric form )"
	// 					},
	// 					"response": {
	// 						"type": "string"
	// 					},
	// 					"return-code": {
	// 						"type": "integer",
	// 						"description": "Numeric return code for the command"
	// 					},
	// 					"component-id": {
	// 						"type": "string"
	// 					},
	// 					"time-stamp": {
	// 						"type": "string",
	// 						"description": "Time at which this event was detected"
	// 					},
	// 					"time-stamp-numeric": {
	// 						"type": "integer",
	// 						"description": "Time at which this event was detected( In numeric form )"
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	schemas := value["components"].(map[string]interface{})
	if schemas == nil {
		return fmt.Errorf("unable to extract (components) from /meta command for (%s)", meta)
	}

	object := schemas["schemas"].(map[string]interface{})
	if schemas == nil {
		return fmt.Errorf("unable to extract (schemas) from /meta command for (%s)", meta)
	}

	metadata := object[meta].(map[string]interface{})
	if metadata == nil {
		return fmt.Errorf("unable to extract (%s) from /meta command for (%s)", meta, meta)
	}

	properties := metadata["properties"].(map[string]interface{})
	if properties == nil {
		return fmt.Errorf("unable to extract (properties) from /meta command for (%s)", meta)
	}

	s.generatedResources[meta] = true

	s.Resources(2, fmt.Sprintf("%sResource:", meta))
	s.Resources(3, "type: array")
	s.Resources(3, "minItems: 1")
	s.Resources(3, fmt.Sprintf("maxItems: %d", maxArrayCount))
	s.Resources(3, "items:")
	s.Resources(4, "type: object")
	s.Resources(4, "properties:")

	// By default, all resource contain these two properties: object-name and meta
	s.Resources(5, "object-name:")
	s.Resources(6, "type: string")
	s.Resources(5, "meta:")
	s.Resources(6, "type: string")

	// Next export all the properties, but do so in a sorted manner
	// The meta command does not produce the same property order every time, hence the desire to sort for consistency
	if len(properties) > 0 {

		// Determine if there is an exception for this resource
		hasException := false
		for _, e := range exceptions {
			if (e.Resource == meta) && (e.Type == exceptionRename) {
				hasException = true
				break
			}
		}

		keys := make([]string, 0, len(properties))

		for k := range properties {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		for _, key := range keys {
			newkey := key
			if hasException {
				for _, e := range exceptions {
					if (e.Resource == meta) && (e.OldProperty == key) {
						newkey = e.NewProperty
					}
				}
			}
			data := properties[key].(map[string]interface{})
			datatype := data["type"]
			description := data["description"]
			logger.V(4).Info("properties", "key", newkey, "type", datatype, "description", description)
			s.Resources(5, fmt.Sprintf("%s:", newkey))
			s.Resources(6, fmt.Sprintf("type: %s", datatype))
			if datatype == "integer" {
				// add special formatting to support very large integers
				s.Resources(6, fmt.Sprintf("format: %s", "int64"))
			}
			if description != nil {
				s.Resources(6, fmt.Sprintf("description: %s", description))
			}
		}
	}

	// If specified in the input yaml, add a reference link to all nested resources
	if len(nested) > 0 {
		for _, n := range nested {
			logger.V(3).Info("nested", "meta", n)
			s.Resources(5, fmt.Sprintf("%s:", n.Name))
			s.Resources(6, fmt.Sprintf("$ref: '#/components/schemas/%sResource'", n.Name))
		}
	}

	return nil
}

// AddCommand: retrieve meta data for the command and return schema and add to the specification object in memory
func (s *Specification) AddCommand(ctx context.Context, config *common.Config, command *CommandInformation, exceptions []ExceptionInformation) error {

	logger := klog.FromContext(ctx)
	logger.V(1).Info("add command and meta data", "command", command.Command)

	capitalize := cases.Title(language.English)

	// Create the full path with all options
	cmd := command.Command
	for _, option := range command.Options {
		if option.KeywordRequired == "true" {
			cmd += fmt.Sprintf("/%s/{%sOption}", option.Flag, option.Flag)
		}
		if option.KeywordRequired == "false" {
			cmd += fmt.Sprintf("/{%sOption}", option.Flag)
		}
	}

	// Create a variable name for the path
	variable := ""
	tokens := strings.Split(command.Command, "/")
	for _, token := range tokens {
		variable += capitalize.String(token)
	}
	for _, option := range command.Options {
		variable += capitalize.String(option.Flag)
	}

	// Add path and get option
	s.Paths(1, fmt.Sprintf("# %s", command.Command))
	s.Paths(1, fmt.Sprintf("%s:", cmd))
	s.Paths(2, "get:")
	s.Paths(3, fmt.Sprintf("description: Execute %s command", command.Command))
	s.Paths(3, fmt.Sprintf("operationId: %sGet", variable))

	// Add any parameters
	if len(command.Options) > 0 {
		s.Paths(3, "parameters:")
		for _, option := range command.Options {
			s.Paths(4, fmt.Sprintf("- $ref: '#/components/parameters/%sOption'", option.Flag))
		}
	}

	// Add responses
	s.Paths(3, "responses:")
	s.Paths(4, "'200':")
	s.Paths(5, "description: OK")
	s.Paths(5, "content:")
	s.Paths(6, "application/json:")
	s.Paths(7, "schema:")
	s.Paths(8, fmt.Sprintf("$ref: '#/components/schemas/%sObject'", command.Meta))
	s.Paths(4, "'401':")
	s.Paths(5, "description: Unauthorized")
	s.Paths(5, "content:")
	s.Paths(6, "application/json:")
	s.Paths(7, "schema:")
	s.Paths(8, "$ref: '#/components/schemas/statusObject'")

	// If specified in the input yaml, generate spec lines for all include data types
	if len(command.Include) > 0 {
		for _, include := range command.Include {
			err := s.AddMetaProperties(ctx, config, include, nil, exceptions)
			if err != nil {
				return fmt.Errorf("unable to add meta (properties) from /meta command for (%s), error: %v", include, err)
			}
		}
	}

	// If specified in the input yaml, generate spec lines for the resource and add a reference link
	err := s.AddMetaProperties(ctx, config, command.Meta, command.Nested, exceptions)
	if err != nil {
		return fmt.Errorf("unable to add meta (properties) from /meta command for (%s)", command.Meta)
	}

	// Add the meta object, if not already generated
	generated := s.generatedObjects[command.Meta]
	if !generated {
		s.generatedObjects[command.Meta] = true

		logger.V(3).Info("add object", "name", fmt.Sprintf("%sObject", command.Meta))
		s.Objects(2, fmt.Sprintf("%sObject:", command.Meta))
		s.Objects(3, "type: object")
		s.Objects(3, "properties:")

		// If specified in the input yaml, add a link to any included resource before defining this resource's object properties
		if len(command.Include) > 0 {
			for _, include := range command.Include {
				logger.V(3).Info("add include ref", "name", fmt.Sprintf("%sResource:", include))
				s.Objects(4, fmt.Sprintf("%s:", include))
				s.Objects(5, fmt.Sprintf("$ref: '#/components/schemas/%sResource'", include))
			}
		}

		// Add link to meta resource reference
		s.Objects(4, fmt.Sprintf("%s:", command.Meta))
		s.Objects(5, fmt.Sprintf("$ref: '#/components/schemas/%sResource'", command.Meta))
	}

	// Add all options as parameters
	if len(command.Options) > 0 {
		for _, option := range command.Options {
			logger.V(3).Info("option", "flag", option.Flag, "type", option.Type, "values", option.Values)
			if !s.generatedOptions[option.Flag] {
				s.generatedOptions[option.Flag] = true
				s.Parameters(2, fmt.Sprintf("%sOption:", option.Flag))
				s.Parameters(3, fmt.Sprintf("name: %sOption", option.Flag))
				s.Parameters(3, "in: path")
				if strings.EqualFold(option.Required, "true") {
					s.Parameters(3, "required: true")
				} else {
					// TODO: in order to NOT be required, need to omit form path
					s.Parameters(3, "required: true")
				}
				s.Parameters(3, "schema:")

				if strings.EqualFold(option.Type, "enum") {
					s.Parameters(4, "type: string")
					s.Parameters(4, "enum:")
					tokens := strings.Split(option.Values, ",")
					for _, token := range tokens {
						s.Parameters(4, fmt.Sprintf("- '%s'", token))
					}
				} else {
					s.Parameters(4, fmt.Sprintf("type: %s", strings.ToLower(option.Type)))
				}
			}
		}
	}

	return nil
}
