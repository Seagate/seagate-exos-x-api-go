// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package mcapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/Seagate/seagate-exos-x-api-go/v2/pkg/client"
	"github.com/Seagate/seagate-exos-x-api-go/v2/pkg/common"
	"github.com/go-logr/logr"
	"k8s.io/klog/v2"
)

// Client : Can be used to request the API
type Client struct {
	Ctx        context.Context
	Username   string
	Password   string
	Addr       string
	Protocol   string
	HTTPClient http.Client
	Collector  *common.Collector
	SessionKey string
	Initiator  []string
	PoolName   string
	Info       *common.SystemInfo
	apiClient  *client.APIClient
}

// NewClient : Creates an API client by setting up its HTTP transport
func NewClient() *Client {
	return &Client{
		HTTPClient: http.Client{
			Timeout: time.Duration(15 * time.Second),
			Transport: &http.Transport{
				// Proxy: http.ProxyURL(proxy),
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		Collector: common.NewCollector(),
	}
}

// StoreCredentials : Called to store ip address, username, and password for the client
func (client *Client) StoreCredentials(addr string, protocol string, username string, password string) {

	// addr may include protocol and ip address or hostname, or only ip address
	ipaddress, usingProtocol := common.GetAddressAndProtocol(addr, protocol)

	// Store the login credentials in the Client object
	client.Username = username
	client.Password = password
	client.Addr = ipaddress
	client.Protocol = usingProtocol
}

// Login: Called to log into the storage controller API
func (client *Client) Login(ctx context.Context) error {
	client.Ctx = ctx

	config := &common.Config{
		MCIpAddress: client.Addr,
		MCProtocol:  client.Protocol,
		MCUsername:  client.Username,
		MCPassword:  client.Password,
	}

	apiClient, err := common.Login(ctx, config)
	if err == nil && apiClient != nil {
		client.apiClient = apiClient
		configuration := apiClient.GetConfig()
		client.SessionKey = configuration.DefaultHeader["sessionKey"]
	}

	return err
}

// Login: Called to log into the storage controller API
func (client *Client) Logout() error {
	if client.Ctx == nil {
		client.Ctx = context.Background()
	}
	_, _, err := client.apiClient.DefaultApi.LogoutGet(client.Ctx).Execute()
	return err
}

// SessionValid : Determine if a session is valid, if not a login is required
// Makes the 'show controller-date' API call to validate the session is still valid
//
// Deprecated: This function will be made redundant by retry logic in a future update
func (client *Client) SessionValid(addr, username string) bool {
	if client.Ctx == nil {
		return false
	}
	logger := klog.FromContext(client.Ctx)

	// addr may include protocol and ip address or hostname, or only ip address
	ipaddress := common.GetAddress(addr)

	if client.Addr == ipaddress && client.Username == username {
		if client.SessionKey == "" {
			logger.V(2).Info("session invalid", "sessionkey", client.SessionKey)
			return false
		}
	}

	//run an API call to test that the session is still valid
	_, _, err := client.apiClient.DefaultApi.ShowControllerDateGet(client.Ctx).Execute()
	return err == nil
}

// InitSystemInfo: Retrieve and store system information for this client
func (client *Client) InitSystemInfo() error {

	err := AddSystem(client.Addr, client)
	if err != nil {
		return fmt.Errorf("unable to add system info for ip (%s) ", client.Addr)
	}

	client.Info, err = GetSystem(client.Addr)
	if err == nil {
		_ = Log(client.Info)
	}

	return err
}

// CreateCommonStatus : create a common API status object based on the OpenAPI client response
func CreateCommonStatus(logger logr.Logger, response *[]client.StatusResourceInner) *common.ResponseStatus {

	status := common.ResponseStatus{}

	// Track if this status contains one or more non-info response types
	found := false

	if response != nil {

		// The status response may contain multiple entries, so skip over an "Info" types
		for _, s := range *response {

			if *s.ResponseType == "Info" {
				logger.V(2).Info("create common status (info)",
					"ResponseType", *s.ResponseType,
					"ResponseTypeNumeric", *s.ResponseTypeNumeric,
					"Response", *s.Response,
				)
			} else {
				if found {
					// print warning since there are multiple non-info messages
					logger.V(0).Info("create common status multiple non-info responses",
						"ResponseType", *s.ResponseType,
						"ResponseTypeNumeric", *s.ResponseTypeNumeric,
						"Response", *s.Response,
						"ReturnCode", *s.ReturnCode,
					)
				}
				logger.V(2).Info("create common status",
					"ResponseType", *s.ResponseType,
					"ResponseTypeNumeric", *s.ResponseTypeNumeric,
					"Response", *s.Response,
				)
				found = true
				status.ResponseType = s.GetResponseType()
				status.ResponseTypeNumeric = int(s.GetResponseTypeNumeric())
				status.Response = s.GetResponse()
				status.ReturnCode = int(s.GetReturnCode())
				status.Time = time.Unix(int64(s.GetTimeStampNumeric()), 0)
			}
		}
	}

	return &status
}
