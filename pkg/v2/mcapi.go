// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package mcapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/Seagate/seagate-exos-x-api-go/pkg/client"
	openapiclient "github.com/Seagate/seagate-exos-x-api-go/pkg/client"
	"github.com/Seagate/seagate-exos-x-api-go/pkg/common"
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

	if client.Ctx == nil {
		client.Ctx = ctx
	}

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

// SessionValid : Determine if a session is valid, if not a login is required
func (client *Client) SessionValid(addr, username string) bool {

	logger := klog.FromContext(client.Ctx)

	// addr may include protocol and ip address or hostname, or only ip address
	ipaddress := common.GetAddress(addr)

	if client.Addr == ipaddress && client.Username == username {
		if client.SessionKey == "" {
			logger.V(2).Info("session invalid", "sessionkey", client.SessionKey)
			return false
		}
		logger.V(2).Info("client is configured", "ipaddress", addr)
		return true
	}

	return false
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
func CreateCommonStatus(response *client.StatusObject) *common.ResponseStatus {

	status := common.ResponseStatus{}

	if response != nil {
		status.ResponseType = response.Status[0].GetResponseType()
		status.ResponseTypeNumeric = int(response.Status[0].GetResponseTypeNumeric())
		status.Response = response.Status[0].GetResponse()
		status.ReturnCode = int(response.Status[0].GetReturnCode())
		status.Time = time.Unix(int64(response.Status[0].GetTimeStampNumeric()), 0)
	}

	return &status
}

// CreateCommonStatusFromStatus : create a common API status object based on the OpenAPI client response
func CreateCommonStatusFromStatus(logger logr.Logger, response *openapiclient.StatusResourceInner) *common.ResponseStatus {

	status := common.ResponseStatus{}

	if response != nil {
		status.ResponseType = response.GetResponseType()
		status.ResponseTypeNumeric = int(response.GetResponseTypeNumeric())
		status.Response = response.GetResponse()
		status.ReturnCode = int(response.GetReturnCode())
		status.Time = time.Unix(int64(response.GetTimeStampNumeric()), 0)
		if status.ReturnCode != 0 {
			logger.V(1).Info("status",
				"ResponseType", status.ResponseType,
				"ResponseTypeNumeric", status.ResponseTypeNumeric,
				"Response", status.Response,
				"ReturnCode", status.ReturnCode,
			)
		}
	}

	return &status
}
