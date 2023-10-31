// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package mcapi

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Seagate/seagate-exos-x-api-go/v2/pkg/common"
	"k8s.io/klog/v2"
)

//
// System Information Storage and Creation
//
// These functions are used to store information for one or more storage arrays, all indexed by their URL
//

// systems: Data object storing all system information for all added controllers
var systems common.SystemsData

// AddSystem: Uses the client to query and store system data
func AddSystem(url string, client *Client) error {

	// Create the system record with the designated ip address
	var s common.SystemInfo
	s.URL = strings.ToLower(url)
	if strings.HasPrefix(s.URL, "http://") {
		parts := strings.Split(s.URL, "http://")
		if len(parts) >= 2 {
			s.Protocol = "http://"
			s.IPAddress = parts[1]
		}
	} else if strings.HasPrefix(s.URL, "https://") {
		parts := strings.Split(s.URL, "https://")
		if len(parts) >= 2 {
			s.Protocol = "https://"
			s.IPAddress = parts[1]
		}
	} else {
		s.IPAddress = s.URL
		s.Protocol = "http://"
	}
	systems.Systems = append(systems.Systems, &s)
	s.Ports = nil

	// Extract and store controller data, including ports
	response, httpRes, err := client.apiClient.DefaultApi.ShowControllersGet(client.Ctx).Execute()

	if err == nil && httpRes.StatusCode == http.StatusOK && response != nil {

		for _, controller := range response.GetControllers() {
			klog.V(2).InfoS("++ Processing controller", "controller-id", controller.GetControllerId())

			if controller.GetIpAddress() == s.IPAddress {
				klog.V(2).InfoS("++ Saving controller", "controller-id", controller.GetControllerId())
				s.Controller = controller.GetControllerId()
				s.Platform = controller.GetPlatformType()
				s.SerialNumber = controller.GetSerialNumber()
				s.Status = controller.GetStatus()
			}

			for _, port := range controller.Port {
				klog.V(2).InfoS("++ Adding", "port", port.GetPort(), "type", port.GetPortType())
				p := common.PortType{
					Label:    port.GetPort(),
					Type:     port.GetPortType(),
					TargetId: port.GetTargetId(),
				}
				if port.GetPortType() == "iSCSI" {
					for _, i := range port.GetIscsiPort() {
						p.IPAddress = i.GetIpAddress()
						p.Present = i.GetSfpPresent()
						p.Compliance = i.GetSfpEthernetCompliance()
					}
				}
				s.Ports = append(s.Ports, p)
			}
		}
	}

	// Extract and store controller firmware versions
	versions, httpRes, err := client.apiClient.DefaultApi.ShowVersionsDetailGet(client.Ctx).Execute()
	if err == nil && httpRes.StatusCode == http.StatusOK {
		for _, version := range versions.GetVersions() {
			if strings.EqualFold("A", s.Controller) && version.GetObjectName() == "controller-a-versions" {
				s.MCCodeVersion = version.GetMcFw()
				s.MCBaseVersion = version.GetMcBaseFw()
			}
			if strings.EqualFold("B", s.Controller) && version.GetObjectName() == "controller-b-versions" {
				s.MCCodeVersion = version.GetMcFw()
				s.MCBaseVersion = version.GetMcBaseFw()
			}
		}
	}

	// Extract and store pool data
	s.Pools = nil
	diskgroups, httpRes, err := client.apiClient.DefaultApi.ShowDiskGroupsGet(client.Ctx).Execute()

	if err != nil {
		klog.ErrorS(err, "show disk-groups")
		return err
	}

	if err == nil && httpRes.StatusCode == http.StatusOK {
		for _, diskgroup := range diskgroups.GetDiskGroups() {
			klog.V(2).InfoS("++ Adding", "diskgroup pool", diskgroup.GetPool())
			s.Pools = append(s.Pools,
				common.PoolType{
					Name:         diskgroup.GetPool(),
					SerialNumber: diskgroup.GetPoolSerialNumber(),
					Type:         diskgroup.GetStorageType(),
				})
		}
	}

	return nil
}

// GetSystem: Return the System data object corresponding to the IP Address
func GetSystem(url string) (*common.SystemInfo, error) {

	for _, s := range systems.Systems {
		if s.URL == url {
			return s, nil
		}
	}

	return nil, fmt.Errorf("no system data found for ip (%s) in (%d) systems", url, len(systems.Systems))
}

//
// System functions for ease of use
//

// Log: Display the contents of all system information collected
func Log(system *common.SystemInfo) error {

	klog.Infof("\n")
	klog.Infof("System Information:")

	klog.Infof("\n")
	klog.Infof("=== Controller ===")
	klog.Infof("IPAddress:     %v\n", system.IPAddress)
	klog.Infof("Protocol:      %v\n", system.Protocol)
	klog.Infof("Controller:    %v\n", system.Controller)
	klog.Infof("Platform:      %v\n", system.Platform)
	klog.Infof("SerialNumber:  %v\n", system.SerialNumber)
	klog.Infof("Status:        %v\n", system.Status)
	klog.Infof("MCCodeVersion: %v\n", system.MCCodeVersion)
	klog.Infof("MCBaseVersion: %v\n", system.MCBaseVersion)

	klog.Infof("\n")
	klog.Infof("=== Ports ===")
	for i, p := range system.Ports {
		klog.Infof("Port [%d] %v, %v, %v, %15v, %12v, %v\n",
			i, p.Label, p.Type, p.TargetId, p.IPAddress, p.Present, p.Compliance)
	}

	klog.Infof("\n")
	klog.Infof("=== Pools ===")
	for i, p := range system.Pools {
		klog.Infof("Pool [%d] %-12s  %-8s  %s\n", i, p.Name, p.Type, p.SerialNumber)
	}

	klog.Infof("\n")
	return nil
}

// GetPoolType: Return the pool type of the storage system
func (client *Client) GetPoolType(pool string) (string, error) {
	return GetPoolType(client.Info, pool)
}

// GetPoolType: Return the pool type for a given pool
func GetPoolType(system *common.SystemInfo, pool string) (string, error) {
	if system == nil {
		return "", fmt.Errorf("system pointer is nil")
	}

	for _, p := range system.Pools {
		if p.Name == pool {
			klog.V(2).InfoS("++ pool", "name", p.Name, "type", p.Type)
			return p.Type, nil
		}
	}

	return "", fmt.Errorf("pool (%s) was not found in (%d) system information pools", pool, len(system.Pools))
}

// GetTargetId: Return the target id value for this storage system
func GetTargetId(system *common.SystemInfo, portType string) (string, error) {
	if system == nil {
		return "", fmt.Errorf("system pointer is nil")
	}

	for _, p := range system.Ports {
		if p.Type == portType && p.TargetId != "" {
			klog.V(2).InfoS("++ Targets", "target", p.TargetId, "ipaddress", system.IPAddress, "type", p.Type)
			return p.TargetId, nil
		}
	}

	return "", fmt.Errorf("TargetId was not found for system (%s) with (%d) ports", system.IPAddress, len(system.Ports))
}

// GetPortals: Return a list of portals for the storage system
func (client *Client) GetPortals() (string, error) {
	return GetPortals(client.Info)
}

// GetPortals: Return a list of iSCSI portals for the storage system
func GetPortals(system *common.SystemInfo) (string, error) {
	if system == nil {
		return "", fmt.Errorf("system pointer is nil")
	}

	portals := ""

	for _, p := range system.Ports {
		if p.IPAddress != "0.0.0.0" {
			klog.V(2).InfoS("++ Add portal", "portal", p.IPAddress, "system", system.IPAddress)
			if portals != "" {
				portals = portals + "," + p.IPAddress
			} else {
				portals = p.IPAddress
			}
		}
	}

	if portals != "" {
		return portals, nil
	}

	return "", fmt.Errorf("no portals found for system (%s) with (%d) ports", system.IPAddress, len(system.Ports))
}
