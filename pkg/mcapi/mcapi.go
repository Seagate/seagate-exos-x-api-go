// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package mcapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Seagate/seagate-exos-x-api-go/pkg/client"
	"github.com/Seagate/seagate-exos-x-api-go/pkg/common"
	"k8s.io/klog/v2"
)

// Client : Can be used to request the API
type Client struct {
	Ctx        context.Context
	Username   string
	Password   string
	Addr       string
	HTTPClient http.Client
	Collector  *common.Collector
	SessionKey string
	Initiator  string
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
func (client *Client) StoreCredentials(ipaddress string, username string, password string) {

	// Store the login credentials in the Client object
	client.Username = username
	client.Password = password
	if strings.HasPrefix(ipaddress, "https") {
		client.Addr = strings.Replace(ipaddress, "https://", "", 1)
	} else {
		client.Addr = strings.Replace(ipaddress, "http://", "", 1)
	}
}

// Login: Called to log into the storage controller API
func (client *Client) Login(ctx context.Context) error {

	config := &common.Config{
		MCIpAddress: client.Addr,
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

	if client.Addr == addr && client.Username == username {
		if client.SessionKey == "" {
			klog.V(2).InfoS("session invalid", "sessionkey", client.SessionKey)
			return false
		}
		klog.V(2).InfoS("client is configured", "ipaddress", addr)
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
