// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package common

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Seagate/seagate-exos-x-api-go/pkg/client"

	"k8s.io/klog/v2"
)

var (
	configuration = &client.Configuration{}
	apiClient     *client.APIClient
)

// GetAddress: Return ip address, handles the case where addr contains protocol
func GetAddress(addr string) string {

	// addr may include protocol and ip address or hostname, or only ip address
	// Example:
	//   addr == "http://1.2.3.4" OR "https://1.2.3.4" OR "1.2.3.4"
	//   addr == "http://host.name" OR "https://host.name" OR "host.name"
	// client.Addr only contains ip address
	ipaddress := addr
	if strings.HasPrefix(addr, "http://") {
		ipaddress = strings.Replace(addr, "http://", "", 1)
	} else if strings.HasPrefix(addr, "https://") {
		ipaddress = strings.Replace(addr, "https://", "", 1)
	}

	return ipaddress
}

// GetAddressAndProtocol: Return ip address and protocol, handling concatenated string
func GetAddressAndProtocol(addr string, protocol string) (string, string) {

	// addr may include protocol and ip address or hostname, or only ip address
	// Example:
	//   addr == "http://1.2.3.4" OR "https://1.2.3.4" OR "1.2.3.4"
	//   addr == "http://host.name" OR "https://host.name" OR "host.name"
	ipaddress := addr
	usingProtocol := ""
	if strings.HasPrefix(addr, "http://") {
		ipaddress = strings.Replace(addr, "http://", "", 1)
		usingProtocol = "http"
	} else if strings.HasPrefix(addr, "https://") {
		ipaddress = strings.Replace(addr, "https://", "", 1)
		usingProtocol = "https"
	}

	if protocol == "" && usingProtocol == "" {
		// Use a default protocol if not supplied
		protocol = "https"
		klog.V(0).InfoS("protocol not provided, using default", "protocol", protocol)
	} else if protocol == "" {
		// Use the protocol supplied with addr
		protocol = usingProtocol
	} else if protocol != usingProtocol {
		klog.V(0).InfoS("Address protocol does not match passed in protocol", "addr", addr, "protocol", protocol, "usingProtocol", usingProtocol)
	}

	return ipaddress, protocol
}

// Login: Perform the needed steps to configure a connection and login to the MC
func Login(ctx context.Context, config *Config) (*client.APIClient, error) {

	if apiClient != nil {
		return apiClient, nil
	}

	logger := klog.FromContext(ctx)

	configuration := &client.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "MC OpenAPI",
		Debug:         false,
		Servers: client.ServerConfigurations{
			{
				URL:         fmt.Sprintf("%s://%s/api", config.MCProtocol, config.MCIpAddress),
				Description: config.MCDescription,
			},
		},
		OperationServers: map[string]client.ServerConfigurations{},
	}

	// Add 'datatype: json' in the default header so the MC returns JSON data
	configuration.DefaultHeader["datatype"] = "json"

	apiClient = client.NewAPIClient(configuration)

	var sessionKey string = ""

	logger.V(3).Info("================================================================================")
	resp1, httpRes, err := apiClient.DefaultApi.LoginGet(ctx).Execute()
	if err == nil && httpRes != nil && httpRes.StatusCode == http.StatusOK {
		logger.V(3).Info("++ LoginGet",
			"ResponseType", *resp1.Status[0].ResponseType,
			"ResponseTypeNumeric", *resp1.Status[0].ResponseTypeNumeric)

		if *resp1.Status[0].ResponseTypeNumeric == 0 {
			logger.V(0).Info("++ MC Login SUCCESS", "ipaddress", config.MCIpAddress, "protocol", config.MCProtocol)
			sessionKey = *resp1.Status[0].Response
		} else {
			logger.V(0).Info("++ MC Login FAILURE", "response", *resp1.Status[0].Response)
			return nil, fmt.Errorf("++ MC Login FAILURE", "response", *resp1.Status[0].Response)
		}
	} else {
		logger.V(0).Info("-- LoginGet", "status", httpRes.Status, "err", err, "body", httpRes.Body)
	}
	logger.V(3).Info("================================================================================")

	// Store the 'sessionKey' in the default header
	configuration.DefaultHeader["sessionKey"] = sessionKey

	return apiClient, nil
}
