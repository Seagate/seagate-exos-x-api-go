// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package mcapi

import (
	"net/http"

	openapiclient "github.com/Seagate/seagate-exos-x-api-go/v2/pkg/client"
	"github.com/Seagate/seagate-exos-x-api-go/v2/pkg/common"
	"k8s.io/klog/v2"
)

// CreateSnapshot : create a snapshot in a snap pool and the snap pool if it doesn't already exist
func (client *Client) CreateSnapshot(volumeName string, snapshotName string) (*common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)

	createSnapCall := client.apiClient.DefaultApi.CreateSnapshotsVolumesNamesGet(client.Ctx, volumeName, snapshotName).Execute
	_, commonStatus, httpRes, err := ExecuteWithFailover(createSnapCall, client)
	logger.V(2).Info("create snapshot", "volume", volumeName, "snapshot", snapshotName, "http", httpRes.Status)

	return commonStatus, err
}

// ShowSnapshots : Show one snaphot, or all snapshots, or all snapshots for a volume
func (client *Client) ShowSnapshots(snapshotId string, sourceVolumeId string) ([]common.SnapshotObject, *common.ResponseStatus, error) {

	logger := klog.FromContext(client.Ctx)

	var err error
	var httpRes *http.Response
	var response *openapiclient.SnapshotsObject
	var status *common.ResponseStatus

	// Call the correct OpenAPI Client function
	if sourceVolumeId != "" {
		response, status, httpRes, err = ExecuteWithFailover(client.apiClient.DefaultApi.ShowSnapshotsVolumeGet(client.Ctx, sourceVolumeId).Execute, client)
		logger.V(2).Info("show snapshots volume", "volume", sourceVolumeId, "snapshot", snapshotId, "http", httpRes.Status)
	} else if snapshotId != "" {
		response, status, httpRes, err = ExecuteWithFailover(client.apiClient.DefaultApi.ShowSnapshotsPatternGet(client.Ctx, snapshotId).Execute, client)
		logger.V(2).Info("show snapshots pattern", "volume", sourceVolumeId, "snapshot", snapshotId, "http", httpRes.Status)
	} else {
		response, status, httpRes, err = ExecuteWithFailover(client.apiClient.DefaultApi.ShowSnapshotsGet(client.Ctx).Execute, client)
		logger.V(2).Info("show snapshots", "volume", sourceVolumeId, "snapshot", snapshotId, "http", httpRes.Status)
	}

	returnSnapshots := []common.SnapshotObject{}

	// Fill in Snapshot properties for all data objects returned
	if err == nil && status.ResponseTypeNumeric == 0 {
		for _, s := range response.Snapshots {
			snapshot := common.SnapshotObject{}
			snapshot.ObjectName = s.GetObjectName()
			snapshot.CreationTime, _ = common.CreationTimeFromString(s.GetCreationDateTime())
			snapshot.CreationDateTime = s.GetCreationDateTime()
			snapshot.CreationDateTimeNumeric = s.GetCreationDateTimeNumeric()
			snapshot.MasterVolumeName = s.GetMasterVolumeName()
			snapshot.Name = s.GetName()
			snapshot.StoragePoolName = s.GetStoragePoolName()
			snapshot.TotalSize = s.GetTotalSize()
			snapshot.TotalSizeNumeric = s.GetTotalSizeNumeric()
			snapshot.VolumeParent = s.GetVolumeParent()

			returnSnapshots = append(returnSnapshots, snapshot)
		}
	}

	return returnSnapshots, status, err
}

// DeleteSnapshot : delete a snapshot
func (client *Client) DeleteSnapshot(name string) (*common.ResponseStatus, error) {
	logger := klog.FromContext(client.Ctx)
	_, responseStatus, httpRes, err := ExecuteWithFailover(client.apiClient.DefaultApi.DeleteSnapshotNamesGet(client.Ctx, name).Execute, client)
	logger.V(2).Info("delete snapshot", "name", name, "http", httpRes.Status)
	return responseStatus, err
}
