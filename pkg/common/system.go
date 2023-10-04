// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package common

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

// System: Information stored for a single storage array controller
type SystemInfo struct {
	IPAddress     string
	Protocol      string
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

// SystemsData: Information stored for multiple storage array controllers
type SystemsData struct {
	Systems []*SystemInfo
}
