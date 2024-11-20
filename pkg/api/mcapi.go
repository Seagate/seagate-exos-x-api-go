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
	Ctx           context.Context
	Username      string
	Password      string
	Addrs         []string
	CurrentAddr   string
	NotResponding string
	Protocol      string
	HTTPClient    http.Client
	Collector     *common.Collector
	SessionKey    string
	Initiator     []string
	PoolName      string
	Info          *common.SystemInfo
	apiClient     *client.APIClient
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
func (client *Client) StoreCredentials(addrs []string, protocol string, username string, password string) {
	// addr may include protocol and ip address or hostname, or only ip address
	ipaddress, usingProtocol := common.GetAddressAndProtocol(addrs[0], protocol)

	// Store the login credentials in the Client object
	client.Username = username
	client.Password = password
	client.CurrentAddr = ipaddress
	client.Protocol = usingProtocol
	client.Addrs = addrs
}

func (client *Client) MarkCurrentControllerDown() error {
	if len(client.Addrs) < 2 {
		return fmt.Errorf("cannot mark controller down with less than 2 controller IP addresses configured")
	}
	downController := client.CurrentAddr
	for _, addr := range client.Addrs {
		if addr != downController {
			client.Addrs[0] = addr
			client.Addrs[1] = downController
			break
		}
	}
	return nil
}

// Login: Called to log into the storage controller API
func (myclient *Client) Login(ctx context.Context) (err error) {
	logger := klog.FromContext(ctx)

	var api_client *client.APIClient
	myclient.SessionKey = ""

	for _, host := range myclient.Addrs {
		myclient.CurrentAddr, _ = common.GetAddressAndProtocol(host, myclient.Protocol)
		config := &common.Config{
			MCIpAddress: myclient.CurrentAddr,
			MCProtocol:  myclient.Protocol,
			MCUsername:  myclient.Username,
			MCPassword:  myclient.Password,
		}

		myclient.Ctx = context.WithValue(ctx, client.ContextBasicAuth, client.BasicAuth{
			UserName: myclient.Username,
			Password: myclient.Password,
		})
		logger.V(4).Info("==============config=================", "IPAddr", config.MCIpAddress, "Username", config.MCUsername)
		api_client, err = common.Login(myclient.Ctx, config)

		if err == nil && api_client != nil {
			myclient.apiClient = api_client
			configuration := api_client.GetConfig()
			myclient.SessionKey = configuration.DefaultHeader["sessionKey"]
		}
		// if we have a new session key, login was successful and we can leave
		// if not, try the next controller IP
		if myclient.SessionKey != "" {
			break
		}
	}
	return err
}

// Logout: Called to log out of the storage controller API
func (client *Client) Logout() error {
	if client.Ctx == nil {
		client.Ctx = context.Background()
	}
	_, _, err := client.apiClient.DefaultApi.LogoutGet(client.Ctx).Execute()
	if err == nil {
		client.SessionKey = ""
	}
	return err
}

// InitSystemInfo: Retrieve and store system information for this client
func (client *Client) InitSystemInfo() error {

	err := AddSystem(client.CurrentAddr, client)
	if err != nil {
		return fmt.Errorf("unable to add system info for ip (%s) ", client.CurrentAddr)
	}

	client.Info, err = GetSystem(client.CurrentAddr)
	if err == nil {
		_ = Log(client.Info)
	}

	return err
}

type ResponseType interface {
	GetStatus() []client.StatusResourceInner
}

type OneOfReturnResponse interface {
	GetActualInstance() interface{}
}

// Extract "StatusResourceInner" from either an openAPI Object with GetStatus, or by first Retrieving
// the real object with GetActualInstance, in the case of an OpenAPI call with multiple possible return types
// Example: ShowMapsInitiatorGet200Response
func GetStatus(response interface{}) (returnStatus []client.StatusResourceInner, err error) {
	switch resp := response.(type) {
	case ResponseType:
		returnStatus = resp.GetStatus()
	case OneOfReturnResponse:
		rv := resp.GetActualInstance()
		if r, ok := rv.(ResponseType); ok {
			returnStatus = r.GetStatus()
		} else {
			err = fmt.Errorf("unable to extract status from response type")
		}
	default:
		err = fmt.Errorf("unable to extract status from response type")
	}
	return
}

// ExecuteWithFailover: Retry wrapper for the Execute functions of the openapi generated client library
// If the initial request fails, attempts to login on the secondary controller
func ExecuteWithFailover[R interface{}](executeFunc func() (R, *http.Response, error), myclient *Client) (R, *common.ResponseStatus, *http.Response, error) {
	logger := klog.FromContext(myclient.Ctx)
	logger.V(4).Info("Execute with failover...")
	response, httpRes, err := executeFunc()
	if err == nil {

		status, err := GetStatus(response)
		if err != nil {
			return response, nil, httpRes, err
		}
		commonStatus := CreateCommonStatus(logger, &status)
		// retry if the return code from the controller is in the list of retryable return codes
		retry := false
		for _, rc := range common.RetryableErrorCodes {
			if rc == commonStatus.ReturnCode {
				logger.V(4).Info("Retrying with return code", "return code", rc)
				retry = true
				break
			}
		}
		if !retry {
			return response, commonStatus, httpRes, err
		}
	}
	// If our first attempt resulted in an error or a retryable return code
	if len(myclient.Addrs) > 1 {
		myclient.NotResponding = myclient.CurrentAddr
		logger.V(1).Info("Retrying...", "err", err, "response", response)
		myclient.MarkCurrentControllerDown()
		myclient.Login(myclient.Ctx)
		response, httpRes, err = executeFunc()
		if err != nil {
			return response, nil, httpRes, err
		} else {
			status, err := GetStatus(response)
			commonStatus := CreateCommonStatus(logger, &status)
			return response, commonStatus, httpRes, err
		}
	} else {
		logger.V(1).Info("Cannot failover as only 1 controller address Specified")
		return response, &common.ResponseStatus{}, httpRes, err
	}
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
					"ResponseCode", s.GetReturnCode(),
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
