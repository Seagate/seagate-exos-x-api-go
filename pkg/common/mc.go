// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package common

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Seagate/seagate-exos-x-api-go/pkg/client"

	"k8s.io/klog/v2"
)

var (
	configuration = &client.Configuration{}
	apiClient     *client.APIClient
)

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
