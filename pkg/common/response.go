// Copyright (c) 2023 Seagate Technology LLC and/or its Affiliates

package common

import (
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// Exos X Storage API Error Codes
const (
	InvalidSessionKey                     = 2
	SnapshotNotFoundErrorCode             = -10050
	BadInputParam                         = -10058
	InvalidArgumentErrorCode              = -10058
	HostMapDoesNotExistsErrorCode         = -10074
	VolumeNotFoundErrorCode               = -10075
	VolumeHasSnapshot                     = -10183
	SnapshotAlreadyExists                 = -10186
	InitiatorNicknameOrIdentifierNotFound = -10386
	UnmapFailedErrorCode                  = -10509
)

// ResponseStatus: Final representation of the "status" object in every API response
type ResponseStatus struct {
	ResponseType        string
	ResponseTypeNumeric int
	Response            string
	ReturnCode          int
	Time                time.Time
}

// VolumeObject: Schema data for each volume object returned from 'show volumes'
type VolumeObject struct {
	ObjectName      string // volume
	Blocks          int64  // The size in blocks, blocks
	BlockSize       int64  // Block Size
	Health          string
	SizeNumeric     int64  // size-numeric
	StoragePoolName string // The disk group used to create this volume.
	StorageType     string
	TierAffinity    string // The tier affinity used to create this volume.
	TotalSize       string
	VolumeName      string // User-defined name for the volume, volume-name
	VolumeType      string
	Wwn             string // World Wide Name
}

// SnapshotObject: Schema data for each snapshot object returned from 'show snapshots'
type SnapshotObject struct {
	ObjectName              string               // snapshot
	CreationDateTime        string               //  creation-date-time, created using ptypes.TimestampProto
	CreationDateTimeNumeric int64                // creation-date-time-numeric
	CreationTime            *timestamp.Timestamp // Creation timestamp
	MasterVolumeName        string               // master-volume-name,omitempty
	Name                    string               // name
	StoragePoolName         string               // Disk group used to create the snapshot
	TotalSize               string               // The total size (total-size) formatted using the session settings for base, precision, and units
	TotalSizeNumeric        int64                // The total size (total-size-numeric) formatted using the session settings for base, precision, and units( In numeric form )
	VolumeParent            string               // The name of volume used to create the snapshot
}

func CreationTimeFromString(creationTime string) (*timestamp.Timestamp, error) {
	creationTimestamp, err := strconv.ParseInt(creationTime, 10, 64)
	if err != nil {
		return nil, err
	}

	return ptypes.TimestampProto(time.Unix(creationTimestamp, 0))
}
