// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package validator

import (
	"context"
	"net/http"

	"github.com/Seagate/seagate-exos-x-api-go/pkg/client"

	"github.com/go-logr/logr"
	"k8s.io/klog/v2"
)

// parseJsonMap: Display a JSON Map
func parseJsonMap(logger logr.Logger, level int, aMap map[string]interface{}) {
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			logger.V(1).Info("JSON dic", "level", level, "key", key)
			parseJsonMap(logger, level+1, val.(map[string]interface{}))
		case []interface{}:
			logger.V(1).Info("JSON arr", "level", level, "key", key)
			parseJsonArray(logger, level+1, val.([]interface{}))
		default:
			logger.V(1).Info("JSON obj", "level", level, "key", key, "value", concreteVal)
		}
	}
}

// parseJsonArray: Display a JSON Array object
func parseJsonArray(logger logr.Logger, level int, anArray []interface{}) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			logger.V(1).Info("JSON array", "level", level, "index", i)
			parseJsonMap(logger, level+1, val.(map[string]interface{}))
		case []interface{}:
			logger.V(1).Info("JSON array", "level", level, "index", i)
			parseJsonArray(logger, level+1, val.([]interface{}))
		default:
			logger.V(1).Info("JSON array", "level", level, "index", i, "value", concreteVal)
		}
	}
}

// Meta: Execute a meta command and display the JSON map
func Meta(ctx context.Context, apiClient *client.APIClient, resource string) error {

	// meta status - demonstrates handling generic JSON object
	logger := klog.FromContext(ctx)
	logger.V(0).Info("")
	logger.V(0).Info("================================================================================")
	logger.V(0).Info(">> SchemaGet", "resource", resource)

	response, httpRes, err := apiClient.DefaultApi.SchemaGet(ctx, "status").Execute()
	if httpRes.StatusCode == http.StatusOK {
		logger.V(5).Info("++ SchemaGet", "resp", response)
		value, ok := response.(map[string]interface{})
		if ok {
			parseJsonMap(logger, 0, value)
		}
	} else {
		logger.V(0).Info("-- SystemGet", "status", httpRes.Status, "err", err, "body", httpRes.Body)
	}

	return nil
}

// ShowSystems: Execute a show systems command and display the results
func ShowSystems(ctx context.Context, apiClient *client.APIClient) error {

	logger := klog.FromContext(ctx)

	// show systems
	logger.V(0).Info("")
	logger.V(0).Info("================================================================================")
	logger.V(0).Info(">> SystemGet")
	response, httpRes, err := apiClient.DefaultApi.ShowSystemGet(ctx).Execute()
	if httpRes.StatusCode == http.StatusOK {
		logger.V(5).Info("++ SystemGet", "resp", response)
		logger.V(5).Info("++ SystemGet", "SystemResource", response.System[0])

		// Extract Status resource information
		logger.V(0).Info("++ SystemGet.Status[0]",
			"Response", *response.Status[0].Response,
			"ReturnCode", *response.Status[0].ReturnCode,
		)

		// Extract System resource information
		logger.V(0).Info("++ SystemGet.System[0]",
			"MidplaneSerialNumber", *response.System[0].MidplaneSerialNumber,
			"ScsiVendorId", *response.System[0].ScsiVendorId,
			"ProductId", *response.System[0].ProductId,
		)

		// Extract System.Redundancy resource information
		logger.V(0).Info("++ SystemGet.System[0].Redundancy[0]",
			"RedundancyMode", *response.System[0].Redundancy[0].RedundancyMode,
		)
		logger.V(0).Info("++ SystemGet.System[0]",
			"ControllerAStatus", *response.System[0].Redundancy[0].ControllerAStatus,
			"ControllerASerialNumber", *response.System[0].Redundancy[0].ControllerASerialNumber,
		)
		logger.V(0).Info("++ SystemGet.System[0]",
			"ControllerBStatus", *response.System[0].Redundancy[0].ControllerBStatus,
			"ControllerBSerialNumber", *response.System[0].Redundancy[0].ControllerBSerialNumber,
		)

	} else {
		logger.V(0).Info("-- SystemGet", "status", httpRes.Status, "err", err, "body", httpRes.Body)
	}

	logger.V(0).Info("================================================================================")

	return nil
}

// ShowVersionsDetail: Execute a show versions command and display the results
func ShowVersionsDetail(ctx context.Context, apiClient *client.APIClient) error {

	logger := klog.FromContext(ctx)

	// show systems
	logger.V(0).Info("")
	logger.V(0).Info("================================================================================")
	logger.V(0).Info(">> ShowVersionsDetailGet")
	response, httpRes, err := apiClient.DefaultApi.ShowVersionsDetailGet(ctx).Execute()

	if httpRes.StatusCode == http.StatusOK {
		logger.V(5).Info("++ ShowVersionsDetailGet", "resp", response)

		// Extract Status resource information
		logger.V(0).Info("++ ShowVersionsDetailGet.Status[0]",
			"Response", *response.Status[0].Response,
			"ReturnCode", *response.Status[0].ReturnCode,
		)

		for i, version := range response.Versions {
			// Extract version information
			logger.V(0).Info("++ Versions",
				"index", i,
				"BundleVersion", *version.BundleVersion,
				"BuildDate", *version.BuildDate,
				"McFw", *version.McFw,
			)
		}

	} else {
		logger.V(0).Info("-- ShowVersionsDetailGet", "status", httpRes.Status, "err", err, "body", httpRes.Body)
	}

	logger.V(0).Info("================================================================================")

	return nil
}

// ShowHostGroups: Execute a show host-groups command and display the results
func ShowHostGroups(ctx context.Context, apiClient *client.APIClient) error {

	logger := klog.FromContext(ctx)

	// show host-groups
	logger.V(0).Info("")
	logger.V(0).Info("================================================================================")
	logger.V(0).Info(">> HostGroupGet")
	response, httpRes, err := apiClient.DefaultApi.ShowHostGroupsGet(ctx).Execute()
	if httpRes.StatusCode == http.StatusOK {

		// Extract Status resource information
		logger.V(0).Info("++ HostGroupGet.Status[0]",
			"Response", *response.Status[0].Response,
			"ReturnCode", *response.Status[0].ReturnCode,
		)

		// Extract System resource information
		logger.V(0).Info("++ HostGroupGet.HostGroup[0]",
			"DurableId", *response.HostGroup[0].DurableId,
			"Name", *response.HostGroup[0].Name,
			"SerialNumber", *response.HostGroup[0].SerialNumber,
			"MemberCount", *response.HostGroup[0].MemberCount,
		)

		for i, host := range response.HostGroup[0].Host {
			// Extract Host resource information
			logger.V(0).Info("++ HostGroupGet.HostGroup[0].Host[]",
				"index", i,
				"Name", *host.Name,
				"SerialNumber", *host.SerialNumber,
				"HostGroup", *host.HostGroup,
				"MemberCount", *&host.MemberCount,
			)
		}
	} else {
		logger.V(0).Info("-- HostGroupGet", "status", httpRes.Status, "err", err, "body", httpRes.Body)
	}

	logger.V(0).Info("================================================================================")

	return nil
}

// ShowVolumes: Execute a show volumes ({$name}) command and display the results
func ShowVolumes(ctx context.Context, apiClient *client.APIClient, names string) error {

	logger := klog.FromContext(ctx)

	// show volumes
	logger.V(0).Info("")
	logger.V(0).Info("================================================================================")
	logger.V(0).Info(">> ShowVolumesNamesGet")
	response, httpRes, err := apiClient.DefaultApi.ShowVolumesNamesGet(ctx, names).Execute()
	if httpRes.StatusCode == http.StatusOK {

		// Extract Status resource information
		logger.V(0).Info("++ ShowVolumesNamesGet.Status[0]",
			"Response", *response.Status[0].Response,
			"ReturnCode", *response.Status[0].ReturnCode,
		)

		for i, volume := range response.Volumes {
			logger.V(0).Info("-- volumes",
				"index", i,
				"VolumeName", *volume.VolumeName,
				"TotalSize", *volume.TotalSize,
				"StoragePoolName", *volume.StoragePoolName,
				"SerialNumber", *volume.SerialNumber,
			)
		}
	} else {
		logger.V(0).Info("-- ShowVolumesNamesGet", "status", httpRes.Status, "err", err, "body", httpRes.Body)
	}

	logger.V(0).Info("================================================================================")

	return nil
}

// CreateVolume: Execute a create volumes ({$pool}, {$size}, {$name}) command and display the results
func CreateVolume(ctx context.Context, apiClient *client.APIClient, pool string, size string, tierAffinity string, name string) error {

	logger := klog.FromContext(ctx)

	// show volumes
	logger.V(0).Info("")
	logger.V(0).Info("================================================================================")
	logger.V(0).Info(">> CreateVolumeGet")
	response, httpRes, err := apiClient.DefaultApi.CreateVolumePoolSizeTierAffinityNameGet(ctx, pool, size, tierAffinity, name).Execute()

	if httpRes.StatusCode == http.StatusOK {

		// Extract Status resource information
		logger.V(0).Info("++ CreateVolumeGet.Status[0]",
			"Response", *response.Status[0].Response,
			"ReturnCode", *response.Status[0].ReturnCode,
		)
	} else {
		logger.V(0).Info("-- CreateVolumeGet", "status", httpRes.Status, "err", err, "body", httpRes.Body)
	}

	logger.V(0).Info("================================================================================")

	return nil
}

// DeleteVolumes: Delete one or more volumes
// names - A comma-separated list of the names or serial numbers of the volumes to delete. A name that includes a space must be enclosed in double quotes.
func DeleteVolumes(ctx context.Context, apiClient *client.APIClient, names string) error {

	logger := klog.FromContext(ctx)

	// delete volumes names
	logger.V(0).Info("")
	logger.V(0).Info("================================================================================")
	logger.V(0).Info(">> DeleteVolumesNamesGet", "names", names)

	response, httpRes, err := apiClient.DefaultApi.DeleteVolumesNamesGet(ctx, names).Execute()
	logger.V(0).Info("-- DeleteVolumesNamesGet", "http status", httpRes.Status, "err", err)

	if httpRes.StatusCode == http.StatusOK {
		// Extract Status resource information
		logger.V(0).Info("++ DeleteVolumesNamesGet.Status[0]",
			"Response", *response.Status[0].Response,
			"ReturnCode", *response.Status[0].ReturnCode,
		)
	}
	logger.V(0).Info("================================================================================")

	return nil
}
