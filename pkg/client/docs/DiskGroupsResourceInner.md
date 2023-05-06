# DiskGroupsResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**ActualSpareCapacity** | Pointer to **string** |  | [optional] 
<<<<<<< HEAD
**ActualSpareCapacityNumeric** | Pointer to **int64** |  | [optional] 
**AllocatedPages** | Pointer to **int64** | Number of pages allocated | [optional] 
**ArrayDriveType** | Pointer to **string** | Disk interface type | [optional] 
**ArrayDriveTypeNumeric** | Pointer to **int64** | Disk interface type( In numeric form ) | [optional] 
**AvailablePages** | Pointer to **int64** | Available pages | [optional] 
**Blocks** | Pointer to **int64** | The size in blocks | [optional] 
**Blocksize** | Pointer to **int64** |  | [optional] 
**CacheFlushPeriod** | Pointer to **int64** |  | [optional] 
**CacheReadAhead** | Pointer to **string** |  | [optional] 
**CacheReadAheadNumeric** | Pointer to **int64** |  | [optional] 
**Chunksize** | Pointer to **string** | Smallest block of usable space | [optional] 
**CreateDate** | Pointer to **string** |  | [optional] 
**CreateDateNumeric** | Pointer to **int64** |  | [optional] 
**CriticalCapacity** | Pointer to **string** |  | [optional] 
**CriticalCapacityNumeric** | Pointer to **int64** |  | [optional] 
**CurrentJob** | Pointer to **string** |  | [optional] 
**CurrentJobCompletion** | Pointer to **string** |  | [optional] 
**CurrentJobNumeric** | Pointer to **int64** |  | [optional] 
**DegradedCapacity** | Pointer to **string** |  | [optional] 
**DegradedCapacityNumeric** | Pointer to **int64** |  | [optional] 
**DiskDescription** | Pointer to **string** | Disk interface Description | [optional] 
**DiskDescriptionNumeric** | Pointer to **int64** | Disk interface Description( In numeric form ) | [optional] 
**DiskDsdDelayVdisk** | Pointer to **int64** |  | [optional] 
**DiskDsdEnableVdisk** | Pointer to **string** |  | [optional] 
**DiskDsdEnableVdiskNumeric** | Pointer to **int64** |  | [optional] 
**Diskcount** | Pointer to **int64** | Number of disks | [optional] 
**ExtendedStatus** | Pointer to **int64** | Extended status (bits) | [optional] 
**Freespace** | Pointer to **string** | Amount of free space in the vdisk | [optional] 
**FreespaceNumeric** | Pointer to **int64** | Amount of free space in the vdisk( In numeric form ) | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**HealthNumeric** | Pointer to **int64** |  | [optional] 
**HealthReason** | Pointer to **string** |  | [optional] 
**HealthReasonNumeric** | Pointer to **int64** |  | [optional] 
**HealthRecommendation** | Pointer to **string** |  | [optional] 
**HealthRecommendationNumeric** | Pointer to **int64** |  | [optional] 
**InterleavedVolumeCount** | Pointer to **int64** |  | [optional] 
**IsJobAutoAbortable** | Pointer to **string** |  | [optional] 
**IsJobAutoAbortableNumeric** | Pointer to **int64** |  | [optional] 
**JobRunning** | Pointer to **string** |  | [optional] 
**LargestFreePartitionSpace** | Pointer to **string** |  | [optional] 
**LargestFreePartitionSpaceNumeric** | Pointer to **int64** |  | [optional] 
**LinearVolumeBoundary** | Pointer to **int64** |  | [optional] 
**Lun** | Pointer to **int64** | Logical Unit Number | [optional] 
**MetadataSize** | Pointer to **string** | Disk Group Metadata Capacity | [optional] 
**MetadataSizeNumeric** | Pointer to **int64** | Disk Group Metadata Capacity( In numeric form ) | [optional] 
**MinDriveSize** | Pointer to **string** | Smallest disk size that can be used for this disk group | [optional] 
**MinDriveSizeNumeric** | Pointer to **int64** | Smallest disk size that can be used for this disk group( In numeric form ) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NewPartitionLba** | Pointer to **string** | LBA used for the next volume that will be created | [optional] 
**NewPartitionLbaNumeric** | Pointer to **int64** | LBA used for the next volume that will be created( In numeric form ) | [optional] 
**NumArrayPartitions** | Pointer to **int64** | Number of volumes in this vdisk | [optional] 
**NumDrivesPerLowLevelArray** | Pointer to **int64** | Number of disks in the RAID 10 or 50 subgroup | [optional] 
**NumExpansionPartitions** | Pointer to **int64** |  | [optional] 
**NumPartitionSegments** | Pointer to **int64** |  | [optional] 
**Overhead** | Pointer to **string** |  | [optional] 
**OverheadNumeric** | Pointer to **int64** |  | [optional] 
**Owner** | Pointer to **string** | Controller owning the component | [optional] 
**OwnerNumeric** | Pointer to **int64** | Controller owning the component( In numeric form ) | [optional] 
**PerformanceRank** | Pointer to **int64** | Disk group performance rank within the virtual pool | [optional] 
**Pool** | Pointer to **string** | Pool | [optional] 
**PoolPercentage** | Pointer to **int64** | Portion of the virtual pool used by this disk group | [optional] 
**PoolSectorFormat** | Pointer to **string** | Pool Sector Format | [optional] 
**PoolSectorFormatNumeric** | Pointer to **int64** | Pool Sector Format( In numeric form ) | [optional] 
**PoolSerialNumber** | Pointer to **string** | Serial number of the pool | [optional] 
**PoolsUrl** | Pointer to **string** | URL for associated Storage Pool | [optional] 
**PreferredOwner** | Pointer to **string** | Configured owner | [optional] 
**PreferredOwnerNumeric** | Pointer to **int64** | Configured owner( In numeric form ) | [optional] 
**Raidtype** | Pointer to **string** | RAID level | [optional] 
**RaidtypeNumeric** | Pointer to **int64** | RAID level( In numeric form ) | [optional] 
**RawSize** | Pointer to **string** |  | [optional] 
**RawSizeNumeric** | Pointer to **int64** |  | [optional] 
**ReadAheadEnabled** | Pointer to **string** | Indicates whether read-ahead cache mode is enabled | [optional] 
**ReadAheadEnabledNumeric** | Pointer to **int64** | Indicates whether read-ahead cache mode is enabled( In numeric form ) | [optional] 
**ScrubDurationGoal** | Pointer to **int64** | Preferred duration for vdisk scrub utility | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** | The size or capacity formatted with the current session base, precision, and units | [optional] 
**SizeNumeric** | Pointer to **int64** | The size or capacity formatted with the current session base, precision, and units( In numeric form ) | [optional] 
**Sparecount** | Pointer to **int64** | Number of spare disks currently configured | [optional] 
**Spear** | Pointer to **string** |  | [optional] 
**SpearNumeric** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusNumeric** | Pointer to **int64** |  | [optional] 
**StorageTier** | Pointer to **string** | Disk group tier assignment for tiered migration | [optional] 
**StorageTierNumeric** | Pointer to **int64** | Disk group tier assignment for tiered migration( In numeric form ) | [optional] 
**StorageType** | Pointer to **string** | Storage type | [optional] 
**StorageTypeNumeric** | Pointer to **int64** | Storage type( In numeric form ) | [optional] 
**StripeWidth** | Pointer to **string** |  | [optional] 
**StripeWidthNumeric** | Pointer to **int64** |  | [optional] 
**TargetSpareCapacity** | Pointer to **string** |  | [optional] 
**TargetSpareCapacityNumeric** | Pointer to **int64** |  | [optional] 
**TotalPages** | Pointer to **int64** |  | [optional] 
**Url** | Pointer to **string** | The resource URL | [optional] 
**WriteBackEnabled** | Pointer to **string** | Indicates whether disk write-back cache is enabled | [optional] 
**WriteBackEnabledNumeric** | Pointer to **int64** | Indicates whether disk write-back cache is enabled( In numeric form ) | [optional] 
=======
**ActualSpareCapacityNumeric** | Pointer to **int32** |  | [optional] 
**AllocatedPages** | Pointer to **int32** | Number of pages allocated | [optional] 
**ArrayDriveType** | Pointer to **string** | Disk interface type | [optional] 
**ArrayDriveTypeNumeric** | Pointer to **int32** | Disk interface type( In numeric form ) | [optional] 
**AvailablePages** | Pointer to **int32** | Available pages | [optional] 
**Blocks** | Pointer to **int32** | The size in blocks | [optional] 
**Blocksize** | Pointer to **int32** |  | [optional] 
**CacheFlushPeriod** | Pointer to **int32** |  | [optional] 
**CacheReadAhead** | Pointer to **string** |  | [optional] 
**CacheReadAheadNumeric** | Pointer to **int32** |  | [optional] 
**Chunksize** | Pointer to **string** | Smallest block of usable space | [optional] 
**CreateDate** | Pointer to **string** |  | [optional] 
**CreateDateNumeric** | Pointer to **int32** |  | [optional] 
**CriticalCapacity** | Pointer to **string** |  | [optional] 
**CriticalCapacityNumeric** | Pointer to **int32** |  | [optional] 
**CurrentJob** | Pointer to **string** |  | [optional] 
**CurrentJobCompletion** | Pointer to **string** |  | [optional] 
**CurrentJobNumeric** | Pointer to **int32** |  | [optional] 
**DegradedCapacity** | Pointer to **string** |  | [optional] 
**DegradedCapacityNumeric** | Pointer to **int32** |  | [optional] 
**DiskDescription** | Pointer to **string** | Disk interface Description | [optional] 
**DiskDescriptionNumeric** | Pointer to **int32** | Disk interface Description( In numeric form ) | [optional] 
**DiskDsdDelayVdisk** | Pointer to **int32** |  | [optional] 
**DiskDsdEnableVdisk** | Pointer to **string** |  | [optional] 
**DiskDsdEnableVdiskNumeric** | Pointer to **int32** |  | [optional] 
**Diskcount** | Pointer to **int32** | Number of disks | [optional] 
**ExtendedStatus** | Pointer to **int32** | Extended status (bits) | [optional] 
**Freespace** | Pointer to **string** | Amount of free space in the vdisk | [optional] 
**FreespaceNumeric** | Pointer to **int32** | Amount of free space in the vdisk( In numeric form ) | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**HealthNumeric** | Pointer to **int32** |  | [optional] 
**HealthReason** | Pointer to **string** |  | [optional] 
**HealthReasonNumeric** | Pointer to **int32** |  | [optional] 
**HealthRecommendation** | Pointer to **string** |  | [optional] 
**HealthRecommendationNumeric** | Pointer to **int32** |  | [optional] 
**InterleavedVolumeCount** | Pointer to **int32** |  | [optional] 
**IsJobAutoAbortable** | Pointer to **string** |  | [optional] 
**IsJobAutoAbortableNumeric** | Pointer to **int32** |  | [optional] 
**JobRunning** | Pointer to **string** |  | [optional] 
**LargestFreePartitionSpace** | Pointer to **string** |  | [optional] 
**LargestFreePartitionSpaceNumeric** | Pointer to **int32** |  | [optional] 
**LinearVolumeBoundary** | Pointer to **int32** |  | [optional] 
**Lun** | Pointer to **int32** | Logical Unit Number | [optional] 
**MetadataSize** | Pointer to **string** | Disk Group Metadata Capacity | [optional] 
**MetadataSizeNumeric** | Pointer to **int32** | Disk Group Metadata Capacity( In numeric form ) | [optional] 
**MinDriveSize** | Pointer to **string** | Smallest disk size that can be used for this disk group | [optional] 
**MinDriveSizeNumeric** | Pointer to **int32** | Smallest disk size that can be used for this disk group( In numeric form ) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NewPartitionLba** | Pointer to **string** | LBA used for the next volume that will be created | [optional] 
**NewPartitionLbaNumeric** | Pointer to **int32** | LBA used for the next volume that will be created( In numeric form ) | [optional] 
**NumArrayPartitions** | Pointer to **int32** | Number of volumes in this vdisk | [optional] 
**NumDrivesPerLowLevelArray** | Pointer to **int32** | Number of disks in the RAID 10 or 50 subgroup | [optional] 
**NumExpansionPartitions** | Pointer to **int32** |  | [optional] 
**NumPartitionSegments** | Pointer to **int32** |  | [optional] 
**Overhead** | Pointer to **string** |  | [optional] 
**OverheadNumeric** | Pointer to **int32** |  | [optional] 
**Owner** | Pointer to **string** | Controller owning the component | [optional] 
**OwnerNumeric** | Pointer to **int32** | Controller owning the component( In numeric form ) | [optional] 
**PerformanceRank** | Pointer to **int32** | Disk group performance rank within the virtual pool | [optional] 
**Pool** | Pointer to **string** | Pool | [optional] 
**PoolPercentage** | Pointer to **int32** | Portion of the virtual pool used by this disk group | [optional] 
**PoolSectorFormat** | Pointer to **string** | Pool Sector Format | [optional] 
**PoolSectorFormatNumeric** | Pointer to **int32** | Pool Sector Format( In numeric form ) | [optional] 
**PoolSerialNumber** | Pointer to **string** | Serial number of the pool | [optional] 
**PoolsUrl** | Pointer to **string** | URL for associated Storage Pool | [optional] 
**PreferredOwner** | Pointer to **string** | Configured owner | [optional] 
**PreferredOwnerNumeric** | Pointer to **int32** | Configured owner( In numeric form ) | [optional] 
**Raidtype** | Pointer to **string** | RAID level | [optional] 
**RaidtypeNumeric** | Pointer to **int32** | RAID level( In numeric form ) | [optional] 
**RawSize** | Pointer to **string** |  | [optional] 
**RawSizeNumeric** | Pointer to **int32** |  | [optional] 
**ReadAheadEnabled** | Pointer to **string** | Indicates whether read-ahead cache mode is enabled | [optional] 
**ReadAheadEnabledNumeric** | Pointer to **int32** | Indicates whether read-ahead cache mode is enabled( In numeric form ) | [optional] 
**ScrubDurationGoal** | Pointer to **int32** | Preferred duration for vdisk scrub utility | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** | The size or capacity formatted with the current session base, precision, and units | [optional] 
**SizeNumeric** | Pointer to **int32** | The size or capacity formatted with the current session base, precision, and units( In numeric form ) | [optional] 
**Sparecount** | Pointer to **int32** | Number of spare disks currently configured | [optional] 
**Spear** | Pointer to **string** |  | [optional] 
**SpearNumeric** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusNumeric** | Pointer to **int32** |  | [optional] 
**StorageTier** | Pointer to **string** | Disk group tier assignment for tiered migration | [optional] 
**StorageTierNumeric** | Pointer to **int32** | Disk group tier assignment for tiered migration( In numeric form ) | [optional] 
**StorageType** | Pointer to **string** | Storage type | [optional] 
**StorageTypeNumeric** | Pointer to **int32** | Storage type( In numeric form ) | [optional] 
**StripeWidth** | Pointer to **string** |  | [optional] 
**StripeWidthNumeric** | Pointer to **int32** |  | [optional] 
**TargetSpareCapacity** | Pointer to **string** |  | [optional] 
**TargetSpareCapacityNumeric** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** | The resource URL | [optional] 
**WriteBackEnabled** | Pointer to **string** | Indicates whether disk write-back cache is enabled | [optional] 
**WriteBackEnabledNumeric** | Pointer to **int32** | Indicates whether disk write-back cache is enabled( In numeric form ) | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

## Methods

### NewDiskGroupsResourceInner

`func NewDiskGroupsResourceInner() *DiskGroupsResourceInner`

NewDiskGroupsResourceInner instantiates a new DiskGroupsResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskGroupsResourceInnerWithDefaults

`func NewDiskGroupsResourceInnerWithDefaults() *DiskGroupsResourceInner`

NewDiskGroupsResourceInnerWithDefaults instantiates a new DiskGroupsResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *DiskGroupsResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *DiskGroupsResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *DiskGroupsResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *DiskGroupsResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *DiskGroupsResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DiskGroupsResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DiskGroupsResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DiskGroupsResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetActualSpareCapacity

`func (o *DiskGroupsResourceInner) GetActualSpareCapacity() string`

GetActualSpareCapacity returns the ActualSpareCapacity field if non-nil, zero value otherwise.

### GetActualSpareCapacityOk

`func (o *DiskGroupsResourceInner) GetActualSpareCapacityOk() (*string, bool)`

GetActualSpareCapacityOk returns a tuple with the ActualSpareCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSpareCapacity

`func (o *DiskGroupsResourceInner) SetActualSpareCapacity(v string)`

SetActualSpareCapacity sets ActualSpareCapacity field to given value.

### HasActualSpareCapacity

`func (o *DiskGroupsResourceInner) HasActualSpareCapacity() bool`

HasActualSpareCapacity returns a boolean if a field has been set.

### GetActualSpareCapacityNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetActualSpareCapacityNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetActualSpareCapacityNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetActualSpareCapacityNumeric returns the ActualSpareCapacityNumeric field if non-nil, zero value otherwise.

### GetActualSpareCapacityNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetActualSpareCapacityNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetActualSpareCapacityNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetActualSpareCapacityNumericOk returns a tuple with the ActualSpareCapacityNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSpareCapacityNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetActualSpareCapacityNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetActualSpareCapacityNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetActualSpareCapacityNumeric sets ActualSpareCapacityNumeric field to given value.

### HasActualSpareCapacityNumeric

`func (o *DiskGroupsResourceInner) HasActualSpareCapacityNumeric() bool`

HasActualSpareCapacityNumeric returns a boolean if a field has been set.

### GetAllocatedPages

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetAllocatedPages() int64`
=======
`func (o *DiskGroupsResourceInner) GetAllocatedPages() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAllocatedPages returns the AllocatedPages field if non-nil, zero value otherwise.

### GetAllocatedPagesOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetAllocatedPagesOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetAllocatedPagesOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAllocatedPagesOk returns a tuple with the AllocatedPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedPages

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetAllocatedPages(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetAllocatedPages(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetAllocatedPages sets AllocatedPages field to given value.

### HasAllocatedPages

`func (o *DiskGroupsResourceInner) HasAllocatedPages() bool`

HasAllocatedPages returns a boolean if a field has been set.

### GetArrayDriveType

`func (o *DiskGroupsResourceInner) GetArrayDriveType() string`

GetArrayDriveType returns the ArrayDriveType field if non-nil, zero value otherwise.

### GetArrayDriveTypeOk

`func (o *DiskGroupsResourceInner) GetArrayDriveTypeOk() (*string, bool)`

GetArrayDriveTypeOk returns a tuple with the ArrayDriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayDriveType

`func (o *DiskGroupsResourceInner) SetArrayDriveType(v string)`

SetArrayDriveType sets ArrayDriveType field to given value.

### HasArrayDriveType

`func (o *DiskGroupsResourceInner) HasArrayDriveType() bool`

HasArrayDriveType returns a boolean if a field has been set.

### GetArrayDriveTypeNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetArrayDriveTypeNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetArrayDriveTypeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetArrayDriveTypeNumeric returns the ArrayDriveTypeNumeric field if non-nil, zero value otherwise.

### GetArrayDriveTypeNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetArrayDriveTypeNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetArrayDriveTypeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetArrayDriveTypeNumericOk returns a tuple with the ArrayDriveTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayDriveTypeNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetArrayDriveTypeNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetArrayDriveTypeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetArrayDriveTypeNumeric sets ArrayDriveTypeNumeric field to given value.

### HasArrayDriveTypeNumeric

`func (o *DiskGroupsResourceInner) HasArrayDriveTypeNumeric() bool`

HasArrayDriveTypeNumeric returns a boolean if a field has been set.

### GetAvailablePages

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetAvailablePages() int64`
=======
`func (o *DiskGroupsResourceInner) GetAvailablePages() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAvailablePages returns the AvailablePages field if non-nil, zero value otherwise.

### GetAvailablePagesOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetAvailablePagesOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetAvailablePagesOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAvailablePagesOk returns a tuple with the AvailablePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePages

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetAvailablePages(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetAvailablePages(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetAvailablePages sets AvailablePages field to given value.

### HasAvailablePages

`func (o *DiskGroupsResourceInner) HasAvailablePages() bool`

HasAvailablePages returns a boolean if a field has been set.

### GetBlocks

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetBlocks() int64`
=======
`func (o *DiskGroupsResourceInner) GetBlocks() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetBlocksOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetBlocksOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetBlocks(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetBlocks(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *DiskGroupsResourceInner) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetBlocksize

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetBlocksize() int64`
=======
`func (o *DiskGroupsResourceInner) GetBlocksize() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetBlocksize returns the Blocksize field if non-nil, zero value otherwise.

### GetBlocksizeOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetBlocksizeOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetBlocksizeOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetBlocksizeOk returns a tuple with the Blocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocksize

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetBlocksize(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetBlocksize(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetBlocksize sets Blocksize field to given value.

### HasBlocksize

`func (o *DiskGroupsResourceInner) HasBlocksize() bool`

HasBlocksize returns a boolean if a field has been set.

### GetCacheFlushPeriod

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetCacheFlushPeriod() int64`
=======
`func (o *DiskGroupsResourceInner) GetCacheFlushPeriod() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCacheFlushPeriod returns the CacheFlushPeriod field if non-nil, zero value otherwise.

### GetCacheFlushPeriodOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetCacheFlushPeriodOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetCacheFlushPeriodOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCacheFlushPeriodOk returns a tuple with the CacheFlushPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheFlushPeriod

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetCacheFlushPeriod(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetCacheFlushPeriod(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetCacheFlushPeriod sets CacheFlushPeriod field to given value.

### HasCacheFlushPeriod

`func (o *DiskGroupsResourceInner) HasCacheFlushPeriod() bool`

HasCacheFlushPeriod returns a boolean if a field has been set.

### GetCacheReadAhead

`func (o *DiskGroupsResourceInner) GetCacheReadAhead() string`

GetCacheReadAhead returns the CacheReadAhead field if non-nil, zero value otherwise.

### GetCacheReadAheadOk

`func (o *DiskGroupsResourceInner) GetCacheReadAheadOk() (*string, bool)`

GetCacheReadAheadOk returns a tuple with the CacheReadAhead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadAhead

`func (o *DiskGroupsResourceInner) SetCacheReadAhead(v string)`

SetCacheReadAhead sets CacheReadAhead field to given value.

### HasCacheReadAhead

`func (o *DiskGroupsResourceInner) HasCacheReadAhead() bool`

HasCacheReadAhead returns a boolean if a field has been set.

### GetCacheReadAheadNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetCacheReadAheadNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetCacheReadAheadNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCacheReadAheadNumeric returns the CacheReadAheadNumeric field if non-nil, zero value otherwise.

### GetCacheReadAheadNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetCacheReadAheadNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetCacheReadAheadNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCacheReadAheadNumericOk returns a tuple with the CacheReadAheadNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadAheadNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetCacheReadAheadNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetCacheReadAheadNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetCacheReadAheadNumeric sets CacheReadAheadNumeric field to given value.

### HasCacheReadAheadNumeric

`func (o *DiskGroupsResourceInner) HasCacheReadAheadNumeric() bool`

HasCacheReadAheadNumeric returns a boolean if a field has been set.

### GetChunksize

`func (o *DiskGroupsResourceInner) GetChunksize() string`

GetChunksize returns the Chunksize field if non-nil, zero value otherwise.

### GetChunksizeOk

`func (o *DiskGroupsResourceInner) GetChunksizeOk() (*string, bool)`

GetChunksizeOk returns a tuple with the Chunksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunksize

`func (o *DiskGroupsResourceInner) SetChunksize(v string)`

SetChunksize sets Chunksize field to given value.

### HasChunksize

`func (o *DiskGroupsResourceInner) HasChunksize() bool`

HasChunksize returns a boolean if a field has been set.

### GetCreateDate

`func (o *DiskGroupsResourceInner) GetCreateDate() string`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *DiskGroupsResourceInner) GetCreateDateOk() (*string, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *DiskGroupsResourceInner) SetCreateDate(v string)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *DiskGroupsResourceInner) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetCreateDateNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetCreateDateNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetCreateDateNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCreateDateNumeric returns the CreateDateNumeric field if non-nil, zero value otherwise.

### GetCreateDateNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetCreateDateNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetCreateDateNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCreateDateNumericOk returns a tuple with the CreateDateNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDateNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetCreateDateNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetCreateDateNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetCreateDateNumeric sets CreateDateNumeric field to given value.

### HasCreateDateNumeric

`func (o *DiskGroupsResourceInner) HasCreateDateNumeric() bool`

HasCreateDateNumeric returns a boolean if a field has been set.

### GetCriticalCapacity

`func (o *DiskGroupsResourceInner) GetCriticalCapacity() string`

GetCriticalCapacity returns the CriticalCapacity field if non-nil, zero value otherwise.

### GetCriticalCapacityOk

`func (o *DiskGroupsResourceInner) GetCriticalCapacityOk() (*string, bool)`

GetCriticalCapacityOk returns a tuple with the CriticalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalCapacity

`func (o *DiskGroupsResourceInner) SetCriticalCapacity(v string)`

SetCriticalCapacity sets CriticalCapacity field to given value.

### HasCriticalCapacity

`func (o *DiskGroupsResourceInner) HasCriticalCapacity() bool`

HasCriticalCapacity returns a boolean if a field has been set.

### GetCriticalCapacityNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetCriticalCapacityNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetCriticalCapacityNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCriticalCapacityNumeric returns the CriticalCapacityNumeric field if non-nil, zero value otherwise.

### GetCriticalCapacityNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetCriticalCapacityNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetCriticalCapacityNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCriticalCapacityNumericOk returns a tuple with the CriticalCapacityNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalCapacityNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetCriticalCapacityNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetCriticalCapacityNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetCriticalCapacityNumeric sets CriticalCapacityNumeric field to given value.

### HasCriticalCapacityNumeric

`func (o *DiskGroupsResourceInner) HasCriticalCapacityNumeric() bool`

HasCriticalCapacityNumeric returns a boolean if a field has been set.

### GetCurrentJob

`func (o *DiskGroupsResourceInner) GetCurrentJob() string`

GetCurrentJob returns the CurrentJob field if non-nil, zero value otherwise.

### GetCurrentJobOk

`func (o *DiskGroupsResourceInner) GetCurrentJobOk() (*string, bool)`

GetCurrentJobOk returns a tuple with the CurrentJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentJob

`func (o *DiskGroupsResourceInner) SetCurrentJob(v string)`

SetCurrentJob sets CurrentJob field to given value.

### HasCurrentJob

`func (o *DiskGroupsResourceInner) HasCurrentJob() bool`

HasCurrentJob returns a boolean if a field has been set.

### GetCurrentJobCompletion

`func (o *DiskGroupsResourceInner) GetCurrentJobCompletion() string`

GetCurrentJobCompletion returns the CurrentJobCompletion field if non-nil, zero value otherwise.

### GetCurrentJobCompletionOk

`func (o *DiskGroupsResourceInner) GetCurrentJobCompletionOk() (*string, bool)`

GetCurrentJobCompletionOk returns a tuple with the CurrentJobCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentJobCompletion

`func (o *DiskGroupsResourceInner) SetCurrentJobCompletion(v string)`

SetCurrentJobCompletion sets CurrentJobCompletion field to given value.

### HasCurrentJobCompletion

`func (o *DiskGroupsResourceInner) HasCurrentJobCompletion() bool`

HasCurrentJobCompletion returns a boolean if a field has been set.

### GetCurrentJobNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetCurrentJobNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetCurrentJobNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCurrentJobNumeric returns the CurrentJobNumeric field if non-nil, zero value otherwise.

### GetCurrentJobNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetCurrentJobNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetCurrentJobNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCurrentJobNumericOk returns a tuple with the CurrentJobNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentJobNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetCurrentJobNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetCurrentJobNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetCurrentJobNumeric sets CurrentJobNumeric field to given value.

### HasCurrentJobNumeric

`func (o *DiskGroupsResourceInner) HasCurrentJobNumeric() bool`

HasCurrentJobNumeric returns a boolean if a field has been set.

### GetDegradedCapacity

`func (o *DiskGroupsResourceInner) GetDegradedCapacity() string`

GetDegradedCapacity returns the DegradedCapacity field if non-nil, zero value otherwise.

### GetDegradedCapacityOk

`func (o *DiskGroupsResourceInner) GetDegradedCapacityOk() (*string, bool)`

GetDegradedCapacityOk returns a tuple with the DegradedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedCapacity

`func (o *DiskGroupsResourceInner) SetDegradedCapacity(v string)`

SetDegradedCapacity sets DegradedCapacity field to given value.

### HasDegradedCapacity

`func (o *DiskGroupsResourceInner) HasDegradedCapacity() bool`

HasDegradedCapacity returns a boolean if a field has been set.

### GetDegradedCapacityNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetDegradedCapacityNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetDegradedCapacityNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDegradedCapacityNumeric returns the DegradedCapacityNumeric field if non-nil, zero value otherwise.

### GetDegradedCapacityNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetDegradedCapacityNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetDegradedCapacityNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDegradedCapacityNumericOk returns a tuple with the DegradedCapacityNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedCapacityNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetDegradedCapacityNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetDegradedCapacityNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetDegradedCapacityNumeric sets DegradedCapacityNumeric field to given value.

### HasDegradedCapacityNumeric

`func (o *DiskGroupsResourceInner) HasDegradedCapacityNumeric() bool`

HasDegradedCapacityNumeric returns a boolean if a field has been set.

### GetDiskDescription

`func (o *DiskGroupsResourceInner) GetDiskDescription() string`

GetDiskDescription returns the DiskDescription field if non-nil, zero value otherwise.

### GetDiskDescriptionOk

`func (o *DiskGroupsResourceInner) GetDiskDescriptionOk() (*string, bool)`

GetDiskDescriptionOk returns a tuple with the DiskDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDescription

`func (o *DiskGroupsResourceInner) SetDiskDescription(v string)`

SetDiskDescription sets DiskDescription field to given value.

### HasDiskDescription

`func (o *DiskGroupsResourceInner) HasDiskDescription() bool`

HasDiskDescription returns a boolean if a field has been set.

### GetDiskDescriptionNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetDiskDescriptionNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetDiskDescriptionNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDiskDescriptionNumeric returns the DiskDescriptionNumeric field if non-nil, zero value otherwise.

### GetDiskDescriptionNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetDiskDescriptionNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetDiskDescriptionNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDiskDescriptionNumericOk returns a tuple with the DiskDescriptionNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDescriptionNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetDiskDescriptionNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetDiskDescriptionNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetDiskDescriptionNumeric sets DiskDescriptionNumeric field to given value.

### HasDiskDescriptionNumeric

`func (o *DiskGroupsResourceInner) HasDiskDescriptionNumeric() bool`

HasDiskDescriptionNumeric returns a boolean if a field has been set.

### GetDiskDsdDelayVdisk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetDiskDsdDelayVdisk() int64`
=======
`func (o *DiskGroupsResourceInner) GetDiskDsdDelayVdisk() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDiskDsdDelayVdisk returns the DiskDsdDelayVdisk field if non-nil, zero value otherwise.

### GetDiskDsdDelayVdiskOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetDiskDsdDelayVdiskOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetDiskDsdDelayVdiskOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDiskDsdDelayVdiskOk returns a tuple with the DiskDsdDelayVdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDsdDelayVdisk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetDiskDsdDelayVdisk(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetDiskDsdDelayVdisk(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetDiskDsdDelayVdisk sets DiskDsdDelayVdisk field to given value.

### HasDiskDsdDelayVdisk

`func (o *DiskGroupsResourceInner) HasDiskDsdDelayVdisk() bool`

HasDiskDsdDelayVdisk returns a boolean if a field has been set.

### GetDiskDsdEnableVdisk

`func (o *DiskGroupsResourceInner) GetDiskDsdEnableVdisk() string`

GetDiskDsdEnableVdisk returns the DiskDsdEnableVdisk field if non-nil, zero value otherwise.

### GetDiskDsdEnableVdiskOk

`func (o *DiskGroupsResourceInner) GetDiskDsdEnableVdiskOk() (*string, bool)`

GetDiskDsdEnableVdiskOk returns a tuple with the DiskDsdEnableVdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDsdEnableVdisk

`func (o *DiskGroupsResourceInner) SetDiskDsdEnableVdisk(v string)`

SetDiskDsdEnableVdisk sets DiskDsdEnableVdisk field to given value.

### HasDiskDsdEnableVdisk

`func (o *DiskGroupsResourceInner) HasDiskDsdEnableVdisk() bool`

HasDiskDsdEnableVdisk returns a boolean if a field has been set.

### GetDiskDsdEnableVdiskNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetDiskDsdEnableVdiskNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetDiskDsdEnableVdiskNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDiskDsdEnableVdiskNumeric returns the DiskDsdEnableVdiskNumeric field if non-nil, zero value otherwise.

### GetDiskDsdEnableVdiskNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetDiskDsdEnableVdiskNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetDiskDsdEnableVdiskNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDiskDsdEnableVdiskNumericOk returns a tuple with the DiskDsdEnableVdiskNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDsdEnableVdiskNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetDiskDsdEnableVdiskNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetDiskDsdEnableVdiskNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetDiskDsdEnableVdiskNumeric sets DiskDsdEnableVdiskNumeric field to given value.

### HasDiskDsdEnableVdiskNumeric

`func (o *DiskGroupsResourceInner) HasDiskDsdEnableVdiskNumeric() bool`

HasDiskDsdEnableVdiskNumeric returns a boolean if a field has been set.

### GetDiskcount

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetDiskcount() int64`
=======
`func (o *DiskGroupsResourceInner) GetDiskcount() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDiskcount returns the Diskcount field if non-nil, zero value otherwise.

### GetDiskcountOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetDiskcountOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetDiskcountOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDiskcountOk returns a tuple with the Diskcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskcount

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetDiskcount(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetDiskcount(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetDiskcount sets Diskcount field to given value.

### HasDiskcount

`func (o *DiskGroupsResourceInner) HasDiskcount() bool`

HasDiskcount returns a boolean if a field has been set.

### GetExtendedStatus

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetExtendedStatus() int64`
=======
`func (o *DiskGroupsResourceInner) GetExtendedStatus() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetExtendedStatus returns the ExtendedStatus field if non-nil, zero value otherwise.

### GetExtendedStatusOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetExtendedStatusOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetExtendedStatusOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetExtendedStatusOk returns a tuple with the ExtendedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedStatus

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetExtendedStatus(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetExtendedStatus(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetExtendedStatus sets ExtendedStatus field to given value.

### HasExtendedStatus

`func (o *DiskGroupsResourceInner) HasExtendedStatus() bool`

HasExtendedStatus returns a boolean if a field has been set.

### GetFreespace

`func (o *DiskGroupsResourceInner) GetFreespace() string`

GetFreespace returns the Freespace field if non-nil, zero value otherwise.

### GetFreespaceOk

`func (o *DiskGroupsResourceInner) GetFreespaceOk() (*string, bool)`

GetFreespaceOk returns a tuple with the Freespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreespace

`func (o *DiskGroupsResourceInner) SetFreespace(v string)`

SetFreespace sets Freespace field to given value.

### HasFreespace

`func (o *DiskGroupsResourceInner) HasFreespace() bool`

HasFreespace returns a boolean if a field has been set.

### GetFreespaceNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetFreespaceNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetFreespaceNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetFreespaceNumeric returns the FreespaceNumeric field if non-nil, zero value otherwise.

### GetFreespaceNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetFreespaceNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetFreespaceNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetFreespaceNumericOk returns a tuple with the FreespaceNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreespaceNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetFreespaceNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetFreespaceNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetFreespaceNumeric sets FreespaceNumeric field to given value.

### HasFreespaceNumeric

`func (o *DiskGroupsResourceInner) HasFreespaceNumeric() bool`

HasFreespaceNumeric returns a boolean if a field has been set.

### GetHealth

`func (o *DiskGroupsResourceInner) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *DiskGroupsResourceInner) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *DiskGroupsResourceInner) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *DiskGroupsResourceInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHealthNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetHealthNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetHealthNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHealthNumeric returns the HealthNumeric field if non-nil, zero value otherwise.

### GetHealthNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetHealthNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetHealthNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHealthNumericOk returns a tuple with the HealthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetHealthNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetHealthNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetHealthNumeric sets HealthNumeric field to given value.

### HasHealthNumeric

`func (o *DiskGroupsResourceInner) HasHealthNumeric() bool`

HasHealthNumeric returns a boolean if a field has been set.

### GetHealthReason

`func (o *DiskGroupsResourceInner) GetHealthReason() string`

GetHealthReason returns the HealthReason field if non-nil, zero value otherwise.

### GetHealthReasonOk

`func (o *DiskGroupsResourceInner) GetHealthReasonOk() (*string, bool)`

GetHealthReasonOk returns a tuple with the HealthReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReason

`func (o *DiskGroupsResourceInner) SetHealthReason(v string)`

SetHealthReason sets HealthReason field to given value.

### HasHealthReason

`func (o *DiskGroupsResourceInner) HasHealthReason() bool`

HasHealthReason returns a boolean if a field has been set.

### GetHealthReasonNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetHealthReasonNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetHealthReasonNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHealthReasonNumeric returns the HealthReasonNumeric field if non-nil, zero value otherwise.

### GetHealthReasonNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetHealthReasonNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetHealthReasonNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHealthReasonNumericOk returns a tuple with the HealthReasonNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReasonNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetHealthReasonNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetHealthReasonNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetHealthReasonNumeric sets HealthReasonNumeric field to given value.

### HasHealthReasonNumeric

`func (o *DiskGroupsResourceInner) HasHealthReasonNumeric() bool`

HasHealthReasonNumeric returns a boolean if a field has been set.

### GetHealthRecommendation

`func (o *DiskGroupsResourceInner) GetHealthRecommendation() string`

GetHealthRecommendation returns the HealthRecommendation field if non-nil, zero value otherwise.

### GetHealthRecommendationOk

`func (o *DiskGroupsResourceInner) GetHealthRecommendationOk() (*string, bool)`

GetHealthRecommendationOk returns a tuple with the HealthRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecommendation

`func (o *DiskGroupsResourceInner) SetHealthRecommendation(v string)`

SetHealthRecommendation sets HealthRecommendation field to given value.

### HasHealthRecommendation

`func (o *DiskGroupsResourceInner) HasHealthRecommendation() bool`

HasHealthRecommendation returns a boolean if a field has been set.

### GetHealthRecommendationNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetHealthRecommendationNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetHealthRecommendationNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHealthRecommendationNumeric returns the HealthRecommendationNumeric field if non-nil, zero value otherwise.

### GetHealthRecommendationNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetHealthRecommendationNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetHealthRecommendationNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHealthRecommendationNumericOk returns a tuple with the HealthRecommendationNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecommendationNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetHealthRecommendationNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetHealthRecommendationNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetHealthRecommendationNumeric sets HealthRecommendationNumeric field to given value.

### HasHealthRecommendationNumeric

`func (o *DiskGroupsResourceInner) HasHealthRecommendationNumeric() bool`

HasHealthRecommendationNumeric returns a boolean if a field has been set.

### GetInterleavedVolumeCount

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetInterleavedVolumeCount() int64`
=======
`func (o *DiskGroupsResourceInner) GetInterleavedVolumeCount() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetInterleavedVolumeCount returns the InterleavedVolumeCount field if non-nil, zero value otherwise.

### GetInterleavedVolumeCountOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetInterleavedVolumeCountOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetInterleavedVolumeCountOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetInterleavedVolumeCountOk returns a tuple with the InterleavedVolumeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterleavedVolumeCount

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetInterleavedVolumeCount(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetInterleavedVolumeCount(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetInterleavedVolumeCount sets InterleavedVolumeCount field to given value.

### HasInterleavedVolumeCount

`func (o *DiskGroupsResourceInner) HasInterleavedVolumeCount() bool`

HasInterleavedVolumeCount returns a boolean if a field has been set.

### GetIsJobAutoAbortable

`func (o *DiskGroupsResourceInner) GetIsJobAutoAbortable() string`

GetIsJobAutoAbortable returns the IsJobAutoAbortable field if non-nil, zero value otherwise.

### GetIsJobAutoAbortableOk

`func (o *DiskGroupsResourceInner) GetIsJobAutoAbortableOk() (*string, bool)`

GetIsJobAutoAbortableOk returns a tuple with the IsJobAutoAbortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJobAutoAbortable

`func (o *DiskGroupsResourceInner) SetIsJobAutoAbortable(v string)`

SetIsJobAutoAbortable sets IsJobAutoAbortable field to given value.

### HasIsJobAutoAbortable

`func (o *DiskGroupsResourceInner) HasIsJobAutoAbortable() bool`

HasIsJobAutoAbortable returns a boolean if a field has been set.

### GetIsJobAutoAbortableNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetIsJobAutoAbortableNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetIsJobAutoAbortableNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetIsJobAutoAbortableNumeric returns the IsJobAutoAbortableNumeric field if non-nil, zero value otherwise.

### GetIsJobAutoAbortableNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetIsJobAutoAbortableNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetIsJobAutoAbortableNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetIsJobAutoAbortableNumericOk returns a tuple with the IsJobAutoAbortableNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJobAutoAbortableNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetIsJobAutoAbortableNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetIsJobAutoAbortableNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetIsJobAutoAbortableNumeric sets IsJobAutoAbortableNumeric field to given value.

### HasIsJobAutoAbortableNumeric

`func (o *DiskGroupsResourceInner) HasIsJobAutoAbortableNumeric() bool`

HasIsJobAutoAbortableNumeric returns a boolean if a field has been set.

### GetJobRunning

`func (o *DiskGroupsResourceInner) GetJobRunning() string`

GetJobRunning returns the JobRunning field if non-nil, zero value otherwise.

### GetJobRunningOk

`func (o *DiskGroupsResourceInner) GetJobRunningOk() (*string, bool)`

GetJobRunningOk returns a tuple with the JobRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunning

`func (o *DiskGroupsResourceInner) SetJobRunning(v string)`

SetJobRunning sets JobRunning field to given value.

### HasJobRunning

`func (o *DiskGroupsResourceInner) HasJobRunning() bool`

HasJobRunning returns a boolean if a field has been set.

### GetLargestFreePartitionSpace

`func (o *DiskGroupsResourceInner) GetLargestFreePartitionSpace() string`

GetLargestFreePartitionSpace returns the LargestFreePartitionSpace field if non-nil, zero value otherwise.

### GetLargestFreePartitionSpaceOk

`func (o *DiskGroupsResourceInner) GetLargestFreePartitionSpaceOk() (*string, bool)`

GetLargestFreePartitionSpaceOk returns a tuple with the LargestFreePartitionSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargestFreePartitionSpace

`func (o *DiskGroupsResourceInner) SetLargestFreePartitionSpace(v string)`

SetLargestFreePartitionSpace sets LargestFreePartitionSpace field to given value.

### HasLargestFreePartitionSpace

`func (o *DiskGroupsResourceInner) HasLargestFreePartitionSpace() bool`

HasLargestFreePartitionSpace returns a boolean if a field has been set.

### GetLargestFreePartitionSpaceNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetLargestFreePartitionSpaceNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetLargestFreePartitionSpaceNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLargestFreePartitionSpaceNumeric returns the LargestFreePartitionSpaceNumeric field if non-nil, zero value otherwise.

### GetLargestFreePartitionSpaceNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetLargestFreePartitionSpaceNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetLargestFreePartitionSpaceNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLargestFreePartitionSpaceNumericOk returns a tuple with the LargestFreePartitionSpaceNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargestFreePartitionSpaceNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetLargestFreePartitionSpaceNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetLargestFreePartitionSpaceNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetLargestFreePartitionSpaceNumeric sets LargestFreePartitionSpaceNumeric field to given value.

### HasLargestFreePartitionSpaceNumeric

`func (o *DiskGroupsResourceInner) HasLargestFreePartitionSpaceNumeric() bool`

HasLargestFreePartitionSpaceNumeric returns a boolean if a field has been set.

### GetLinearVolumeBoundary

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetLinearVolumeBoundary() int64`
=======
`func (o *DiskGroupsResourceInner) GetLinearVolumeBoundary() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLinearVolumeBoundary returns the LinearVolumeBoundary field if non-nil, zero value otherwise.

### GetLinearVolumeBoundaryOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetLinearVolumeBoundaryOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetLinearVolumeBoundaryOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLinearVolumeBoundaryOk returns a tuple with the LinearVolumeBoundary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinearVolumeBoundary

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetLinearVolumeBoundary(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetLinearVolumeBoundary(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetLinearVolumeBoundary sets LinearVolumeBoundary field to given value.

### HasLinearVolumeBoundary

`func (o *DiskGroupsResourceInner) HasLinearVolumeBoundary() bool`

HasLinearVolumeBoundary returns a boolean if a field has been set.

### GetLun

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetLun() int64`
=======
`func (o *DiskGroupsResourceInner) GetLun() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetLunOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetLunOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetLun(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetLun(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetLun sets Lun field to given value.

### HasLun

`func (o *DiskGroupsResourceInner) HasLun() bool`

HasLun returns a boolean if a field has been set.

### GetMetadataSize

`func (o *DiskGroupsResourceInner) GetMetadataSize() string`

GetMetadataSize returns the MetadataSize field if non-nil, zero value otherwise.

### GetMetadataSizeOk

`func (o *DiskGroupsResourceInner) GetMetadataSizeOk() (*string, bool)`

GetMetadataSizeOk returns a tuple with the MetadataSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSize

`func (o *DiskGroupsResourceInner) SetMetadataSize(v string)`

SetMetadataSize sets MetadataSize field to given value.

### HasMetadataSize

`func (o *DiskGroupsResourceInner) HasMetadataSize() bool`

HasMetadataSize returns a boolean if a field has been set.

### GetMetadataSizeNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetMetadataSizeNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetMetadataSizeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMetadataSizeNumeric returns the MetadataSizeNumeric field if non-nil, zero value otherwise.

### GetMetadataSizeNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetMetadataSizeNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetMetadataSizeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMetadataSizeNumericOk returns a tuple with the MetadataSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSizeNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetMetadataSizeNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetMetadataSizeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetMetadataSizeNumeric sets MetadataSizeNumeric field to given value.

### HasMetadataSizeNumeric

`func (o *DiskGroupsResourceInner) HasMetadataSizeNumeric() bool`

HasMetadataSizeNumeric returns a boolean if a field has been set.

### GetMinDriveSize

`func (o *DiskGroupsResourceInner) GetMinDriveSize() string`

GetMinDriveSize returns the MinDriveSize field if non-nil, zero value otherwise.

### GetMinDriveSizeOk

`func (o *DiskGroupsResourceInner) GetMinDriveSizeOk() (*string, bool)`

GetMinDriveSizeOk returns a tuple with the MinDriveSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDriveSize

`func (o *DiskGroupsResourceInner) SetMinDriveSize(v string)`

SetMinDriveSize sets MinDriveSize field to given value.

### HasMinDriveSize

`func (o *DiskGroupsResourceInner) HasMinDriveSize() bool`

HasMinDriveSize returns a boolean if a field has been set.

### GetMinDriveSizeNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetMinDriveSizeNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetMinDriveSizeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMinDriveSizeNumeric returns the MinDriveSizeNumeric field if non-nil, zero value otherwise.

### GetMinDriveSizeNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetMinDriveSizeNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetMinDriveSizeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMinDriveSizeNumericOk returns a tuple with the MinDriveSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDriveSizeNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetMinDriveSizeNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetMinDriveSizeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetMinDriveSizeNumeric sets MinDriveSizeNumeric field to given value.

### HasMinDriveSizeNumeric

`func (o *DiskGroupsResourceInner) HasMinDriveSizeNumeric() bool`

HasMinDriveSizeNumeric returns a boolean if a field has been set.

### GetName

`func (o *DiskGroupsResourceInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiskGroupsResourceInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiskGroupsResourceInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiskGroupsResourceInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewPartitionLba

`func (o *DiskGroupsResourceInner) GetNewPartitionLba() string`

GetNewPartitionLba returns the NewPartitionLba field if non-nil, zero value otherwise.

### GetNewPartitionLbaOk

`func (o *DiskGroupsResourceInner) GetNewPartitionLbaOk() (*string, bool)`

GetNewPartitionLbaOk returns a tuple with the NewPartitionLba field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPartitionLba

`func (o *DiskGroupsResourceInner) SetNewPartitionLba(v string)`

SetNewPartitionLba sets NewPartitionLba field to given value.

### HasNewPartitionLba

`func (o *DiskGroupsResourceInner) HasNewPartitionLba() bool`

HasNewPartitionLba returns a boolean if a field has been set.

### GetNewPartitionLbaNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetNewPartitionLbaNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetNewPartitionLbaNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetNewPartitionLbaNumeric returns the NewPartitionLbaNumeric field if non-nil, zero value otherwise.

### GetNewPartitionLbaNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetNewPartitionLbaNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetNewPartitionLbaNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetNewPartitionLbaNumericOk returns a tuple with the NewPartitionLbaNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPartitionLbaNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetNewPartitionLbaNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetNewPartitionLbaNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetNewPartitionLbaNumeric sets NewPartitionLbaNumeric field to given value.

### HasNewPartitionLbaNumeric

`func (o *DiskGroupsResourceInner) HasNewPartitionLbaNumeric() bool`

HasNewPartitionLbaNumeric returns a boolean if a field has been set.

### GetNumArrayPartitions

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetNumArrayPartitions() int64`
=======
`func (o *DiskGroupsResourceInner) GetNumArrayPartitions() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetNumArrayPartitions returns the NumArrayPartitions field if non-nil, zero value otherwise.

### GetNumArrayPartitionsOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetNumArrayPartitionsOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetNumArrayPartitionsOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetNumArrayPartitionsOk returns a tuple with the NumArrayPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumArrayPartitions

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetNumArrayPartitions(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetNumArrayPartitions(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetNumArrayPartitions sets NumArrayPartitions field to given value.

### HasNumArrayPartitions

`func (o *DiskGroupsResourceInner) HasNumArrayPartitions() bool`

HasNumArrayPartitions returns a boolean if a field has been set.

### GetNumDrivesPerLowLevelArray

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetNumDrivesPerLowLevelArray() int64`
=======
`func (o *DiskGroupsResourceInner) GetNumDrivesPerLowLevelArray() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetNumDrivesPerLowLevelArray returns the NumDrivesPerLowLevelArray field if non-nil, zero value otherwise.

### GetNumDrivesPerLowLevelArrayOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetNumDrivesPerLowLevelArrayOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetNumDrivesPerLowLevelArrayOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetNumDrivesPerLowLevelArrayOk returns a tuple with the NumDrivesPerLowLevelArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDrivesPerLowLevelArray

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetNumDrivesPerLowLevelArray(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetNumDrivesPerLowLevelArray(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetNumDrivesPerLowLevelArray sets NumDrivesPerLowLevelArray field to given value.

### HasNumDrivesPerLowLevelArray

`func (o *DiskGroupsResourceInner) HasNumDrivesPerLowLevelArray() bool`

HasNumDrivesPerLowLevelArray returns a boolean if a field has been set.

### GetNumExpansionPartitions

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetNumExpansionPartitions() int64`
=======
`func (o *DiskGroupsResourceInner) GetNumExpansionPartitions() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetNumExpansionPartitions returns the NumExpansionPartitions field if non-nil, zero value otherwise.

### GetNumExpansionPartitionsOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetNumExpansionPartitionsOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetNumExpansionPartitionsOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetNumExpansionPartitionsOk returns a tuple with the NumExpansionPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumExpansionPartitions

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetNumExpansionPartitions(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetNumExpansionPartitions(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetNumExpansionPartitions sets NumExpansionPartitions field to given value.

### HasNumExpansionPartitions

`func (o *DiskGroupsResourceInner) HasNumExpansionPartitions() bool`

HasNumExpansionPartitions returns a boolean if a field has been set.

### GetNumPartitionSegments

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetNumPartitionSegments() int64`
=======
`func (o *DiskGroupsResourceInner) GetNumPartitionSegments() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetNumPartitionSegments returns the NumPartitionSegments field if non-nil, zero value otherwise.

### GetNumPartitionSegmentsOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetNumPartitionSegmentsOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetNumPartitionSegmentsOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetNumPartitionSegmentsOk returns a tuple with the NumPartitionSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPartitionSegments

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetNumPartitionSegments(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetNumPartitionSegments(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetNumPartitionSegments sets NumPartitionSegments field to given value.

### HasNumPartitionSegments

`func (o *DiskGroupsResourceInner) HasNumPartitionSegments() bool`

HasNumPartitionSegments returns a boolean if a field has been set.

### GetOverhead

`func (o *DiskGroupsResourceInner) GetOverhead() string`

GetOverhead returns the Overhead field if non-nil, zero value otherwise.

### GetOverheadOk

`func (o *DiskGroupsResourceInner) GetOverheadOk() (*string, bool)`

GetOverheadOk returns a tuple with the Overhead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverhead

`func (o *DiskGroupsResourceInner) SetOverhead(v string)`

SetOverhead sets Overhead field to given value.

### HasOverhead

`func (o *DiskGroupsResourceInner) HasOverhead() bool`

HasOverhead returns a boolean if a field has been set.

### GetOverheadNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetOverheadNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetOverheadNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetOverheadNumeric returns the OverheadNumeric field if non-nil, zero value otherwise.

### GetOverheadNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetOverheadNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetOverheadNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetOverheadNumericOk returns a tuple with the OverheadNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverheadNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetOverheadNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetOverheadNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetOverheadNumeric sets OverheadNumeric field to given value.

### HasOverheadNumeric

`func (o *DiskGroupsResourceInner) HasOverheadNumeric() bool`

HasOverheadNumeric returns a boolean if a field has been set.

### GetOwner

`func (o *DiskGroupsResourceInner) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DiskGroupsResourceInner) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DiskGroupsResourceInner) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DiskGroupsResourceInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnerNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetOwnerNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetOwnerNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetOwnerNumeric returns the OwnerNumeric field if non-nil, zero value otherwise.

### GetOwnerNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetOwnerNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetOwnerNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetOwnerNumericOk returns a tuple with the OwnerNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetOwnerNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetOwnerNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetOwnerNumeric sets OwnerNumeric field to given value.

### HasOwnerNumeric

`func (o *DiskGroupsResourceInner) HasOwnerNumeric() bool`

HasOwnerNumeric returns a boolean if a field has been set.

### GetPerformanceRank

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetPerformanceRank() int64`
=======
`func (o *DiskGroupsResourceInner) GetPerformanceRank() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPerformanceRank returns the PerformanceRank field if non-nil, zero value otherwise.

### GetPerformanceRankOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetPerformanceRankOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetPerformanceRankOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPerformanceRankOk returns a tuple with the PerformanceRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceRank

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetPerformanceRank(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetPerformanceRank(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetPerformanceRank sets PerformanceRank field to given value.

### HasPerformanceRank

`func (o *DiskGroupsResourceInner) HasPerformanceRank() bool`

HasPerformanceRank returns a boolean if a field has been set.

### GetPool

`func (o *DiskGroupsResourceInner) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *DiskGroupsResourceInner) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *DiskGroupsResourceInner) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *DiskGroupsResourceInner) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolPercentage

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetPoolPercentage() int64`
=======
`func (o *DiskGroupsResourceInner) GetPoolPercentage() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPoolPercentage returns the PoolPercentage field if non-nil, zero value otherwise.

### GetPoolPercentageOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetPoolPercentageOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetPoolPercentageOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPoolPercentageOk returns a tuple with the PoolPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolPercentage

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetPoolPercentage(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetPoolPercentage(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetPoolPercentage sets PoolPercentage field to given value.

### HasPoolPercentage

`func (o *DiskGroupsResourceInner) HasPoolPercentage() bool`

HasPoolPercentage returns a boolean if a field has been set.

### GetPoolSectorFormat

`func (o *DiskGroupsResourceInner) GetPoolSectorFormat() string`

GetPoolSectorFormat returns the PoolSectorFormat field if non-nil, zero value otherwise.

### GetPoolSectorFormatOk

`func (o *DiskGroupsResourceInner) GetPoolSectorFormatOk() (*string, bool)`

GetPoolSectorFormatOk returns a tuple with the PoolSectorFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSectorFormat

`func (o *DiskGroupsResourceInner) SetPoolSectorFormat(v string)`

SetPoolSectorFormat sets PoolSectorFormat field to given value.

### HasPoolSectorFormat

`func (o *DiskGroupsResourceInner) HasPoolSectorFormat() bool`

HasPoolSectorFormat returns a boolean if a field has been set.

### GetPoolSectorFormatNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetPoolSectorFormatNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetPoolSectorFormatNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPoolSectorFormatNumeric returns the PoolSectorFormatNumeric field if non-nil, zero value otherwise.

### GetPoolSectorFormatNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetPoolSectorFormatNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetPoolSectorFormatNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPoolSectorFormatNumericOk returns a tuple with the PoolSectorFormatNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSectorFormatNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetPoolSectorFormatNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetPoolSectorFormatNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetPoolSectorFormatNumeric sets PoolSectorFormatNumeric field to given value.

### HasPoolSectorFormatNumeric

`func (o *DiskGroupsResourceInner) HasPoolSectorFormatNumeric() bool`

HasPoolSectorFormatNumeric returns a boolean if a field has been set.

### GetPoolSerialNumber

`func (o *DiskGroupsResourceInner) GetPoolSerialNumber() string`

GetPoolSerialNumber returns the PoolSerialNumber field if non-nil, zero value otherwise.

### GetPoolSerialNumberOk

`func (o *DiskGroupsResourceInner) GetPoolSerialNumberOk() (*string, bool)`

GetPoolSerialNumberOk returns a tuple with the PoolSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSerialNumber

`func (o *DiskGroupsResourceInner) SetPoolSerialNumber(v string)`

SetPoolSerialNumber sets PoolSerialNumber field to given value.

### HasPoolSerialNumber

`func (o *DiskGroupsResourceInner) HasPoolSerialNumber() bool`

HasPoolSerialNumber returns a boolean if a field has been set.

### GetPoolsUrl

`func (o *DiskGroupsResourceInner) GetPoolsUrl() string`

GetPoolsUrl returns the PoolsUrl field if non-nil, zero value otherwise.

### GetPoolsUrlOk

`func (o *DiskGroupsResourceInner) GetPoolsUrlOk() (*string, bool)`

GetPoolsUrlOk returns a tuple with the PoolsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolsUrl

`func (o *DiskGroupsResourceInner) SetPoolsUrl(v string)`

SetPoolsUrl sets PoolsUrl field to given value.

### HasPoolsUrl

`func (o *DiskGroupsResourceInner) HasPoolsUrl() bool`

HasPoolsUrl returns a boolean if a field has been set.

### GetPreferredOwner

`func (o *DiskGroupsResourceInner) GetPreferredOwner() string`

GetPreferredOwner returns the PreferredOwner field if non-nil, zero value otherwise.

### GetPreferredOwnerOk

`func (o *DiskGroupsResourceInner) GetPreferredOwnerOk() (*string, bool)`

GetPreferredOwnerOk returns a tuple with the PreferredOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredOwner

`func (o *DiskGroupsResourceInner) SetPreferredOwner(v string)`

SetPreferredOwner sets PreferredOwner field to given value.

### HasPreferredOwner

`func (o *DiskGroupsResourceInner) HasPreferredOwner() bool`

HasPreferredOwner returns a boolean if a field has been set.

### GetPreferredOwnerNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetPreferredOwnerNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetPreferredOwnerNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPreferredOwnerNumeric returns the PreferredOwnerNumeric field if non-nil, zero value otherwise.

### GetPreferredOwnerNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetPreferredOwnerNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetPreferredOwnerNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPreferredOwnerNumericOk returns a tuple with the PreferredOwnerNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredOwnerNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetPreferredOwnerNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetPreferredOwnerNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetPreferredOwnerNumeric sets PreferredOwnerNumeric field to given value.

### HasPreferredOwnerNumeric

`func (o *DiskGroupsResourceInner) HasPreferredOwnerNumeric() bool`

HasPreferredOwnerNumeric returns a boolean if a field has been set.

### GetRaidtype

`func (o *DiskGroupsResourceInner) GetRaidtype() string`

GetRaidtype returns the Raidtype field if non-nil, zero value otherwise.

### GetRaidtypeOk

`func (o *DiskGroupsResourceInner) GetRaidtypeOk() (*string, bool)`

GetRaidtypeOk returns a tuple with the Raidtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidtype

`func (o *DiskGroupsResourceInner) SetRaidtype(v string)`

SetRaidtype sets Raidtype field to given value.

### HasRaidtype

`func (o *DiskGroupsResourceInner) HasRaidtype() bool`

HasRaidtype returns a boolean if a field has been set.

### GetRaidtypeNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetRaidtypeNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetRaidtypeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRaidtypeNumeric returns the RaidtypeNumeric field if non-nil, zero value otherwise.

### GetRaidtypeNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetRaidtypeNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetRaidtypeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRaidtypeNumericOk returns a tuple with the RaidtypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidtypeNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetRaidtypeNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetRaidtypeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetRaidtypeNumeric sets RaidtypeNumeric field to given value.

### HasRaidtypeNumeric

`func (o *DiskGroupsResourceInner) HasRaidtypeNumeric() bool`

HasRaidtypeNumeric returns a boolean if a field has been set.

### GetRawSize

`func (o *DiskGroupsResourceInner) GetRawSize() string`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *DiskGroupsResourceInner) GetRawSizeOk() (*string, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *DiskGroupsResourceInner) SetRawSize(v string)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *DiskGroupsResourceInner) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### GetRawSizeNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetRawSizeNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetRawSizeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRawSizeNumeric returns the RawSizeNumeric field if non-nil, zero value otherwise.

### GetRawSizeNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetRawSizeNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetRawSizeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRawSizeNumericOk returns a tuple with the RawSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSizeNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetRawSizeNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetRawSizeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetRawSizeNumeric sets RawSizeNumeric field to given value.

### HasRawSizeNumeric

`func (o *DiskGroupsResourceInner) HasRawSizeNumeric() bool`

HasRawSizeNumeric returns a boolean if a field has been set.

### GetReadAheadEnabled

`func (o *DiskGroupsResourceInner) GetReadAheadEnabled() string`

GetReadAheadEnabled returns the ReadAheadEnabled field if non-nil, zero value otherwise.

### GetReadAheadEnabledOk

`func (o *DiskGroupsResourceInner) GetReadAheadEnabledOk() (*string, bool)`

GetReadAheadEnabledOk returns a tuple with the ReadAheadEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadEnabled

`func (o *DiskGroupsResourceInner) SetReadAheadEnabled(v string)`

SetReadAheadEnabled sets ReadAheadEnabled field to given value.

### HasReadAheadEnabled

`func (o *DiskGroupsResourceInner) HasReadAheadEnabled() bool`

HasReadAheadEnabled returns a boolean if a field has been set.

### GetReadAheadEnabledNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetReadAheadEnabledNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetReadAheadEnabledNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetReadAheadEnabledNumeric returns the ReadAheadEnabledNumeric field if non-nil, zero value otherwise.

### GetReadAheadEnabledNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetReadAheadEnabledNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetReadAheadEnabledNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetReadAheadEnabledNumericOk returns a tuple with the ReadAheadEnabledNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadEnabledNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetReadAheadEnabledNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetReadAheadEnabledNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetReadAheadEnabledNumeric sets ReadAheadEnabledNumeric field to given value.

### HasReadAheadEnabledNumeric

`func (o *DiskGroupsResourceInner) HasReadAheadEnabledNumeric() bool`

HasReadAheadEnabledNumeric returns a boolean if a field has been set.

### GetScrubDurationGoal

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetScrubDurationGoal() int64`
=======
`func (o *DiskGroupsResourceInner) GetScrubDurationGoal() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetScrubDurationGoal returns the ScrubDurationGoal field if non-nil, zero value otherwise.

### GetScrubDurationGoalOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetScrubDurationGoalOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetScrubDurationGoalOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetScrubDurationGoalOk returns a tuple with the ScrubDurationGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrubDurationGoal

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetScrubDurationGoal(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetScrubDurationGoal(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetScrubDurationGoal sets ScrubDurationGoal field to given value.

### HasScrubDurationGoal

`func (o *DiskGroupsResourceInner) HasScrubDurationGoal() bool`

HasScrubDurationGoal returns a boolean if a field has been set.

### GetSerialNumber

`func (o *DiskGroupsResourceInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DiskGroupsResourceInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DiskGroupsResourceInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DiskGroupsResourceInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSize

`func (o *DiskGroupsResourceInner) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DiskGroupsResourceInner) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DiskGroupsResourceInner) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *DiskGroupsResourceInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetSizeNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetSizeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSizeNumeric returns the SizeNumeric field if non-nil, zero value otherwise.

### GetSizeNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetSizeNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetSizeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSizeNumericOk returns a tuple with the SizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetSizeNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetSizeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSizeNumeric sets SizeNumeric field to given value.

### HasSizeNumeric

`func (o *DiskGroupsResourceInner) HasSizeNumeric() bool`

HasSizeNumeric returns a boolean if a field has been set.

### GetSparecount

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetSparecount() int64`
=======
`func (o *DiskGroupsResourceInner) GetSparecount() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSparecount returns the Sparecount field if non-nil, zero value otherwise.

### GetSparecountOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetSparecountOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetSparecountOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSparecountOk returns a tuple with the Sparecount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparecount

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetSparecount(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetSparecount(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSparecount sets Sparecount field to given value.

### HasSparecount

`func (o *DiskGroupsResourceInner) HasSparecount() bool`

HasSparecount returns a boolean if a field has been set.

### GetSpear

`func (o *DiskGroupsResourceInner) GetSpear() string`

GetSpear returns the Spear field if non-nil, zero value otherwise.

### GetSpearOk

`func (o *DiskGroupsResourceInner) GetSpearOk() (*string, bool)`

GetSpearOk returns a tuple with the Spear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpear

`func (o *DiskGroupsResourceInner) SetSpear(v string)`

SetSpear sets Spear field to given value.

### HasSpear

`func (o *DiskGroupsResourceInner) HasSpear() bool`

HasSpear returns a boolean if a field has been set.

### GetSpearNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetSpearNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetSpearNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSpearNumeric returns the SpearNumeric field if non-nil, zero value otherwise.

### GetSpearNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetSpearNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetSpearNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSpearNumericOk returns a tuple with the SpearNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpearNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetSpearNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetSpearNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSpearNumeric sets SpearNumeric field to given value.

### HasSpearNumeric

`func (o *DiskGroupsResourceInner) HasSpearNumeric() bool`

HasSpearNumeric returns a boolean if a field has been set.

### GetStatus

`func (o *DiskGroupsResourceInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiskGroupsResourceInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiskGroupsResourceInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiskGroupsResourceInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetStatusNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetStatusNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetStatusNumeric returns the StatusNumeric field if non-nil, zero value otherwise.

### GetStatusNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetStatusNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetStatusNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetStatusNumericOk returns a tuple with the StatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetStatusNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetStatusNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetStatusNumeric sets StatusNumeric field to given value.

### HasStatusNumeric

`func (o *DiskGroupsResourceInner) HasStatusNumeric() bool`

HasStatusNumeric returns a boolean if a field has been set.

### GetStorageTier

`func (o *DiskGroupsResourceInner) GetStorageTier() string`

GetStorageTier returns the StorageTier field if non-nil, zero value otherwise.

### GetStorageTierOk

`func (o *DiskGroupsResourceInner) GetStorageTierOk() (*string, bool)`

GetStorageTierOk returns a tuple with the StorageTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTier

`func (o *DiskGroupsResourceInner) SetStorageTier(v string)`

SetStorageTier sets StorageTier field to given value.

### HasStorageTier

`func (o *DiskGroupsResourceInner) HasStorageTier() bool`

HasStorageTier returns a boolean if a field has been set.

### GetStorageTierNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetStorageTierNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetStorageTierNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetStorageTierNumeric returns the StorageTierNumeric field if non-nil, zero value otherwise.

### GetStorageTierNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetStorageTierNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetStorageTierNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetStorageTierNumericOk returns a tuple with the StorageTierNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTierNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetStorageTierNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetStorageTierNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetStorageTierNumeric sets StorageTierNumeric field to given value.

### HasStorageTierNumeric

`func (o *DiskGroupsResourceInner) HasStorageTierNumeric() bool`

HasStorageTierNumeric returns a boolean if a field has been set.

### GetStorageType

`func (o *DiskGroupsResourceInner) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *DiskGroupsResourceInner) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *DiskGroupsResourceInner) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *DiskGroupsResourceInner) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetStorageTypeNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetStorageTypeNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetStorageTypeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetStorageTypeNumeric returns the StorageTypeNumeric field if non-nil, zero value otherwise.

### GetStorageTypeNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetStorageTypeNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetStorageTypeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetStorageTypeNumericOk returns a tuple with the StorageTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypeNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetStorageTypeNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetStorageTypeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetStorageTypeNumeric sets StorageTypeNumeric field to given value.

### HasStorageTypeNumeric

`func (o *DiskGroupsResourceInner) HasStorageTypeNumeric() bool`

HasStorageTypeNumeric returns a boolean if a field has been set.

### GetStripeWidth

`func (o *DiskGroupsResourceInner) GetStripeWidth() string`

GetStripeWidth returns the StripeWidth field if non-nil, zero value otherwise.

### GetStripeWidthOk

`func (o *DiskGroupsResourceInner) GetStripeWidthOk() (*string, bool)`

GetStripeWidthOk returns a tuple with the StripeWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeWidth

`func (o *DiskGroupsResourceInner) SetStripeWidth(v string)`

SetStripeWidth sets StripeWidth field to given value.

### HasStripeWidth

`func (o *DiskGroupsResourceInner) HasStripeWidth() bool`

HasStripeWidth returns a boolean if a field has been set.

### GetStripeWidthNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetStripeWidthNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetStripeWidthNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetStripeWidthNumeric returns the StripeWidthNumeric field if non-nil, zero value otherwise.

### GetStripeWidthNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetStripeWidthNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetStripeWidthNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetStripeWidthNumericOk returns a tuple with the StripeWidthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeWidthNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetStripeWidthNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetStripeWidthNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetStripeWidthNumeric sets StripeWidthNumeric field to given value.

### HasStripeWidthNumeric

`func (o *DiskGroupsResourceInner) HasStripeWidthNumeric() bool`

HasStripeWidthNumeric returns a boolean if a field has been set.

### GetTargetSpareCapacity

`func (o *DiskGroupsResourceInner) GetTargetSpareCapacity() string`

GetTargetSpareCapacity returns the TargetSpareCapacity field if non-nil, zero value otherwise.

### GetTargetSpareCapacityOk

`func (o *DiskGroupsResourceInner) GetTargetSpareCapacityOk() (*string, bool)`

GetTargetSpareCapacityOk returns a tuple with the TargetSpareCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSpareCapacity

`func (o *DiskGroupsResourceInner) SetTargetSpareCapacity(v string)`

SetTargetSpareCapacity sets TargetSpareCapacity field to given value.

### HasTargetSpareCapacity

`func (o *DiskGroupsResourceInner) HasTargetSpareCapacity() bool`

HasTargetSpareCapacity returns a boolean if a field has been set.

### GetTargetSpareCapacityNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetTargetSpareCapacityNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetTargetSpareCapacityNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetTargetSpareCapacityNumeric returns the TargetSpareCapacityNumeric field if non-nil, zero value otherwise.

### GetTargetSpareCapacityNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetTargetSpareCapacityNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetTargetSpareCapacityNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetTargetSpareCapacityNumericOk returns a tuple with the TargetSpareCapacityNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSpareCapacityNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetTargetSpareCapacityNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetTargetSpareCapacityNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetTargetSpareCapacityNumeric sets TargetSpareCapacityNumeric field to given value.

### HasTargetSpareCapacityNumeric

`func (o *DiskGroupsResourceInner) HasTargetSpareCapacityNumeric() bool`

HasTargetSpareCapacityNumeric returns a boolean if a field has been set.

### GetTotalPages

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetTotalPages() int64`
=======
`func (o *DiskGroupsResourceInner) GetTotalPages() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetTotalPagesOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetTotalPagesOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetTotalPages(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetTotalPages(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *DiskGroupsResourceInner) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetUrl

`func (o *DiskGroupsResourceInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DiskGroupsResourceInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DiskGroupsResourceInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DiskGroupsResourceInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWriteBackEnabled

`func (o *DiskGroupsResourceInner) GetWriteBackEnabled() string`

GetWriteBackEnabled returns the WriteBackEnabled field if non-nil, zero value otherwise.

### GetWriteBackEnabledOk

`func (o *DiskGroupsResourceInner) GetWriteBackEnabledOk() (*string, bool)`

GetWriteBackEnabledOk returns a tuple with the WriteBackEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackEnabled

`func (o *DiskGroupsResourceInner) SetWriteBackEnabled(v string)`

SetWriteBackEnabled sets WriteBackEnabled field to given value.

### HasWriteBackEnabled

`func (o *DiskGroupsResourceInner) HasWriteBackEnabled() bool`

HasWriteBackEnabled returns a boolean if a field has been set.

### GetWriteBackEnabledNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetWriteBackEnabledNumeric() int64`
=======
`func (o *DiskGroupsResourceInner) GetWriteBackEnabledNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetWriteBackEnabledNumeric returns the WriteBackEnabledNumeric field if non-nil, zero value otherwise.

### GetWriteBackEnabledNumericOk

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) GetWriteBackEnabledNumericOk() (*int64, bool)`
=======
`func (o *DiskGroupsResourceInner) GetWriteBackEnabledNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetWriteBackEnabledNumericOk returns a tuple with the WriteBackEnabledNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackEnabledNumeric

<<<<<<< HEAD
`func (o *DiskGroupsResourceInner) SetWriteBackEnabledNumeric(v int64)`
=======
`func (o *DiskGroupsResourceInner) SetWriteBackEnabledNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetWriteBackEnabledNumeric sets WriteBackEnabledNumeric field to given value.

### HasWriteBackEnabledNumeric

`func (o *DiskGroupsResourceInner) HasWriteBackEnabledNumeric() bool`

HasWriteBackEnabledNumeric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


