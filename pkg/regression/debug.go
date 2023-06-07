// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package regression

import (
	"github.com/Seagate/seagate-exos-x-api-go/pkg/common"
	"github.com/go-logr/logr"
)

// ShowVolumes: Display useful data from volume objects
func ShowVolumes(logger logr.Logger, volumes []common.VolumeObject) {

	logger.V(3).Info("============================================================")
	logger.V(3).Info("Show Volumes:", "count", len(volumes))
	for _, volume := range volumes {
		logger.V(3).Info("volume",
			"object", volume.ObjectName,
			"volume-name", volume.VolumeName,
			"storage-pool-name", volume.StoragePoolName,
			"blocks", volume.Blocks,
			"total-size", volume.TotalSize,
			"volume-type", volume.VolumeType,
		)
	}
}

// ShowSnapshots: Display useful information for snapshot objects
func ShowSnapshots(logger logr.Logger, snapshots []common.SnapshotObject) {

	logger.V(3).Info("============================================================")
	logger.V(3).Info("Show Snapshots:", "count", len(snapshots))
	for _, snapshot := range snapshots {
		logger.V(3).Info("volume",
			"object", snapshot.ObjectName,
			"name", snapshot.Name,
			"parent", snapshot.VolumeParent,
		)
	}
}
