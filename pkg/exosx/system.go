//
// Copyright (c) 2021 Seagate Technology LLC and/or its Affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// For any questions about this software or licensing,
// please email opensource@seagate.com or cortx-questions@seagate.com.

package exosx

import (
	"fmt"
	"strings"

	"k8s.io/klog/v2"
)

//
// System Information Storage and Creation
//

// PoolType: Linear or virtual pool attributes
type PoolType struct {
	Name         string
	SerialNumber string
	Type         string
}

// PortType: Storage system port attributes
type PortType struct {
	Label      string
	Type       string
	TargetId   string
	IPAddress  string
	Present    string
	Compliance string
}

// System: Information stored for a storage array controller
type SystemInfo struct {
	IPAddress     string
	HTTP          string
	URL           string
	Controller    string
	Platform      string
	SerialNumber  string
	Status        string
	MCCodeVersion string
	MCBaseVersion string
	Pools         []PoolType
	Ports         []PortType
}

// SystemsData: Information stored multiple storage array controllers
type SystemsData struct {
	Systems []*SystemInfo
}

// systems: Data object storing all system information for all added controllers
var systems SystemsData

// AddSystem: Uses the client to query and store system data
func AddSystem(url string, client *Client) error {

	// Create the system record with the designated ip address
	var s SystemInfo
	s.URL = strings.ToLower(url)
	if strings.HasPrefix(s.URL, "http://") {
		parts := strings.Split(s.URL, "http://")
		if len(parts) >= 2 {
			s.HTTP = "http://"
			s.IPAddress = parts[1]
		}
	} else if strings.HasPrefix(s.URL, "https://") {
		parts := strings.Split(s.URL, "https://")
		if len(parts) >= 2 {
			s.HTTP = "https://"
			s.IPAddress = parts[1]
		}
	} else {
		s.IPAddress = s.URL
		s.HTTP = "http://"
	}
	systems.Systems = append(systems.Systems, &s)
	s.Ports = nil

	// Extract and store controller data, including ports
	response, status, err := client.FormattedRequest("/show/controllers")
	if err == nil && status.ResponseTypeNumeric == 0 {
		for _, obj1 := range response.Objects {
			if obj1.Name == "controller" || obj1.Name == "controllers" {
				klog.V(2).Infof("++ Processing controller (%s)\n", obj1.PropertiesMap["controller-id"].Data)

				if obj1.PropertiesMap["ip-address"].Data == s.IPAddress {
					klog.V(2).Infof("++ Saving controller (%s)\n", obj1.PropertiesMap["controller-id"].Data)
					s.Controller = obj1.PropertiesMap["controller-id"].Data
					s.Platform = obj1.PropertiesMap["platform-type"].Data
					s.SerialNumber = obj1.PropertiesMap["serial-number"].Data
					s.Status = obj1.PropertiesMap["status"].Data
				}

				for _, obj2 := range obj1.Objects {
					if obj2.Name == "ports" {
						klog.V(2).Infof("++ Adding port (%s)\n", obj2.PropertiesMap["port"].Data)
						p := PortType{
							Label:    obj2.PropertiesMap["port"].Data,
							Type:     obj2.PropertiesMap["port-type"].Data,
							TargetId: obj2.PropertiesMap["target-id"].Data,
						}
						if obj2.PropertiesMap["port-type"].Data == "iSCSI" {
							for _, obj3 := range obj2.Objects {
								if obj3.Name == "port-details" {
									p.IPAddress = obj3.PropertiesMap["ip-address"].Data
									p.Present = obj3.PropertiesMap["sfp-present"].Data
									p.Compliance = obj3.PropertiesMap["sfp-ethernet-compliance"].Data
								}
							}
						}
						s.Ports = append(s.Ports, p)
					}
				}
			}
		}
	}

	// Extract and store controller firmware versions
	response, status, err = client.FormattedRequest("/show/versions/detail")
	if err == nil && status.ResponseTypeNumeric == 0 {
		for _, obj1 := range response.Objects {
			if strings.EqualFold("A", s.Controller) && obj1.Name == "controller-a-versions" {
				s.MCCodeVersion = obj1.PropertiesMap["mc-fw"].Data
				s.MCBaseVersion = obj1.PropertiesMap["mc-base-fw"].Data
			}
			if strings.EqualFold("B", s.Controller) && obj1.Name == "controller-b-versions" {
				s.MCCodeVersion = obj1.PropertiesMap["mc-fw"].Data
				s.MCBaseVersion = obj1.PropertiesMap["mc-base-fw"].Data
			}
		}
	}

	// Extract and store pool data
	s.Pools = nil
	response, status, err = client.FormattedRequest("/show/pools")
	if err == nil && status.ResponseTypeNumeric == 0 {
		for _, obj1 := range response.Objects {
			if obj1.Name == "pools" || obj1.Name == "pool" {
				klog.V(2).Infof("++ Adding pool (%s)\n", obj1.PropertiesMap["name"].Data)
				s.Pools = append(s.Pools,
					PoolType{
						Name:         obj1.PropertiesMap["name"].Data,
						SerialNumber: obj1.PropertiesMap["serial-number"].Data,
						Type:         obj1.PropertiesMap["storage-type"].Data,
					})
			}
		}
	}

	return nil
}

// GetSystem: Return the System data object correspoinding to the IP Address
func GetSystem(url string) (*SystemInfo, error) {

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
func (system *SystemInfo) Log() error {

	klog.Infof("\n")
	klog.Infof("System Information:")

	klog.Infof("\n")
	klog.Infof("=== Controller ===")
	klog.Infof("IPAddress:     %v\n", system.IPAddress)
	klog.Infof("HTTP:          %v\n", system.HTTP)
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

// GetPoolType: Return the pool type for a given pool
func (system *SystemInfo) GetPoolType(pool string) (string, error) {
	if system == nil {
		return "", fmt.Errorf("system pointer is nil")
	}

	for _, p := range system.Pools {
		if p.Name == pool {
			klog.V(2).Infof("++ pool (%s) type is (%s)\n", p.Name, p.Type)
			return p.Type, nil
		}
	}

	return "", fmt.Errorf("pool (%s) was not found in (%d) system information pools", pool, len(system.Pools))
}

// GetTargetId: Return the target id value for this storage system
func (system *SystemInfo) GetTargetId(portType string) (string, error) {
	if system == nil {
		return "", fmt.Errorf("system pointer is nil")
	}

	for _, p := range system.Ports {
		if p.Type == portType && p.TargetId != "" {
			klog.V(2).Infof("++ TargetId (%s) for (%s) type (%s)\n", p.TargetId, system.IPAddress, p.Type)
			return p.TargetId, nil
		}
	}

	return "", fmt.Errorf("TargetId was not found for system (%s) with (%d) ports", system.IPAddress, len(system.Ports))
}

// GetPortals: Return a list of iSCSI portals for the storage system
func (system *SystemInfo) GetPortals() (string, error) {
	if system == nil {
		return "", fmt.Errorf("system pointer is nil")
	}

	portals := ""

	for _, p := range system.Ports {
		if p.IPAddress != "0.0.0.0" {
			klog.V(2).Infof("++ Add portal (%s) for (%s)\n", p.IPAddress, system.IPAddress)
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

	return "", fmt.Errorf("No portals found for system (%s) with (%d) ports", system.IPAddress, len(system.Ports))
}
