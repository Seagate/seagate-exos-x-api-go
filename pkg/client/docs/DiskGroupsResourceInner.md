# DiskGroupsResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**ActualSpareCapacity** | Pointer to **string** |  | [optional] 
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

`func (o *DiskGroupsResourceInner) GetActualSpareCapacityNumeric() int64`

GetActualSpareCapacityNumeric returns the ActualSpareCapacityNumeric field if non-nil, zero value otherwise.

### GetActualSpareCapacityNumericOk

`func (o *DiskGroupsResourceInner) GetActualSpareCapacityNumericOk() (*int64, bool)`

GetActualSpareCapacityNumericOk returns a tuple with the ActualSpareCapacityNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSpareCapacityNumeric

`func (o *DiskGroupsResourceInner) SetActualSpareCapacityNumeric(v int64)`

SetActualSpareCapacityNumeric sets ActualSpareCapacityNumeric field to given value.

### HasActualSpareCapacityNumeric

`func (o *DiskGroupsResourceInner) HasActualSpareCapacityNumeric() bool`

HasActualSpareCapacityNumeric returns a boolean if a field has been set.

### GetAllocatedPages

`func (o *DiskGroupsResourceInner) GetAllocatedPages() int64`

GetAllocatedPages returns the AllocatedPages field if non-nil, zero value otherwise.

### GetAllocatedPagesOk

`func (o *DiskGroupsResourceInner) GetAllocatedPagesOk() (*int64, bool)`

GetAllocatedPagesOk returns a tuple with the AllocatedPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedPages

`func (o *DiskGroupsResourceInner) SetAllocatedPages(v int64)`

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

`func (o *DiskGroupsResourceInner) GetArrayDriveTypeNumeric() int64`

GetArrayDriveTypeNumeric returns the ArrayDriveTypeNumeric field if non-nil, zero value otherwise.

### GetArrayDriveTypeNumericOk

`func (o *DiskGroupsResourceInner) GetArrayDriveTypeNumericOk() (*int64, bool)`

GetArrayDriveTypeNumericOk returns a tuple with the ArrayDriveTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayDriveTypeNumeric

`func (o *DiskGroupsResourceInner) SetArrayDriveTypeNumeric(v int64)`

SetArrayDriveTypeNumeric sets ArrayDriveTypeNumeric field to given value.

### HasArrayDriveTypeNumeric

`func (o *DiskGroupsResourceInner) HasArrayDriveTypeNumeric() bool`

HasArrayDriveTypeNumeric returns a boolean if a field has been set.

### GetAvailablePages

`func (o *DiskGroupsResourceInner) GetAvailablePages() int64`

GetAvailablePages returns the AvailablePages field if non-nil, zero value otherwise.

### GetAvailablePagesOk

`func (o *DiskGroupsResourceInner) GetAvailablePagesOk() (*int64, bool)`

GetAvailablePagesOk returns a tuple with the AvailablePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePages

`func (o *DiskGroupsResourceInner) SetAvailablePages(v int64)`

SetAvailablePages sets AvailablePages field to given value.

### HasAvailablePages

`func (o *DiskGroupsResourceInner) HasAvailablePages() bool`

HasAvailablePages returns a boolean if a field has been set.

### GetBlocks

`func (o *DiskGroupsResourceInner) GetBlocks() int64`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *DiskGroupsResourceInner) GetBlocksOk() (*int64, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *DiskGroupsResourceInner) SetBlocks(v int64)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *DiskGroupsResourceInner) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetBlocksize

`func (o *DiskGroupsResourceInner) GetBlocksize() int64`

GetBlocksize returns the Blocksize field if non-nil, zero value otherwise.

### GetBlocksizeOk

`func (o *DiskGroupsResourceInner) GetBlocksizeOk() (*int64, bool)`

GetBlocksizeOk returns a tuple with the Blocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocksize

`func (o *DiskGroupsResourceInner) SetBlocksize(v int64)`

SetBlocksize sets Blocksize field to given value.

### HasBlocksize

`func (o *DiskGroupsResourceInner) HasBlocksize() bool`

HasBlocksize returns a boolean if a field has been set.

### GetCacheFlushPeriod

`func (o *DiskGroupsResourceInner) GetCacheFlushPeriod() int64`

GetCacheFlushPeriod returns the CacheFlushPeriod field if non-nil, zero value otherwise.

### GetCacheFlushPeriodOk

`func (o *DiskGroupsResourceInner) GetCacheFlushPeriodOk() (*int64, bool)`

GetCacheFlushPeriodOk returns a tuple with the CacheFlushPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheFlushPeriod

`func (o *DiskGroupsResourceInner) SetCacheFlushPeriod(v int64)`

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

`func (o *DiskGroupsResourceInner) GetCacheReadAheadNumeric() int64`

GetCacheReadAheadNumeric returns the CacheReadAheadNumeric field if non-nil, zero value otherwise.

### GetCacheReadAheadNumericOk

`func (o *DiskGroupsResourceInner) GetCacheReadAheadNumericOk() (*int64, bool)`

GetCacheReadAheadNumericOk returns a tuple with the CacheReadAheadNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadAheadNumeric

`func (o *DiskGroupsResourceInner) SetCacheReadAheadNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetCreateDateNumeric() int64`

GetCreateDateNumeric returns the CreateDateNumeric field if non-nil, zero value otherwise.

### GetCreateDateNumericOk

`func (o *DiskGroupsResourceInner) GetCreateDateNumericOk() (*int64, bool)`

GetCreateDateNumericOk returns a tuple with the CreateDateNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDateNumeric

`func (o *DiskGroupsResourceInner) SetCreateDateNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetCriticalCapacityNumeric() int64`

GetCriticalCapacityNumeric returns the CriticalCapacityNumeric field if non-nil, zero value otherwise.

### GetCriticalCapacityNumericOk

`func (o *DiskGroupsResourceInner) GetCriticalCapacityNumericOk() (*int64, bool)`

GetCriticalCapacityNumericOk returns a tuple with the CriticalCapacityNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalCapacityNumeric

`func (o *DiskGroupsResourceInner) SetCriticalCapacityNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetCurrentJobNumeric() int64`

GetCurrentJobNumeric returns the CurrentJobNumeric field if non-nil, zero value otherwise.

### GetCurrentJobNumericOk

`func (o *DiskGroupsResourceInner) GetCurrentJobNumericOk() (*int64, bool)`

GetCurrentJobNumericOk returns a tuple with the CurrentJobNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentJobNumeric

`func (o *DiskGroupsResourceInner) SetCurrentJobNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetDegradedCapacityNumeric() int64`

GetDegradedCapacityNumeric returns the DegradedCapacityNumeric field if non-nil, zero value otherwise.

### GetDegradedCapacityNumericOk

`func (o *DiskGroupsResourceInner) GetDegradedCapacityNumericOk() (*int64, bool)`

GetDegradedCapacityNumericOk returns a tuple with the DegradedCapacityNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDegradedCapacityNumeric

`func (o *DiskGroupsResourceInner) SetDegradedCapacityNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetDiskDescriptionNumeric() int64`

GetDiskDescriptionNumeric returns the DiskDescriptionNumeric field if non-nil, zero value otherwise.

### GetDiskDescriptionNumericOk

`func (o *DiskGroupsResourceInner) GetDiskDescriptionNumericOk() (*int64, bool)`

GetDiskDescriptionNumericOk returns a tuple with the DiskDescriptionNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDescriptionNumeric

`func (o *DiskGroupsResourceInner) SetDiskDescriptionNumeric(v int64)`

SetDiskDescriptionNumeric sets DiskDescriptionNumeric field to given value.

### HasDiskDescriptionNumeric

`func (o *DiskGroupsResourceInner) HasDiskDescriptionNumeric() bool`

HasDiskDescriptionNumeric returns a boolean if a field has been set.

### GetDiskDsdDelayVdisk

`func (o *DiskGroupsResourceInner) GetDiskDsdDelayVdisk() int64`

GetDiskDsdDelayVdisk returns the DiskDsdDelayVdisk field if non-nil, zero value otherwise.

### GetDiskDsdDelayVdiskOk

`func (o *DiskGroupsResourceInner) GetDiskDsdDelayVdiskOk() (*int64, bool)`

GetDiskDsdDelayVdiskOk returns a tuple with the DiskDsdDelayVdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDsdDelayVdisk

`func (o *DiskGroupsResourceInner) SetDiskDsdDelayVdisk(v int64)`

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

`func (o *DiskGroupsResourceInner) GetDiskDsdEnableVdiskNumeric() int64`

GetDiskDsdEnableVdiskNumeric returns the DiskDsdEnableVdiskNumeric field if non-nil, zero value otherwise.

### GetDiskDsdEnableVdiskNumericOk

`func (o *DiskGroupsResourceInner) GetDiskDsdEnableVdiskNumericOk() (*int64, bool)`

GetDiskDsdEnableVdiskNumericOk returns a tuple with the DiskDsdEnableVdiskNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDsdEnableVdiskNumeric

`func (o *DiskGroupsResourceInner) SetDiskDsdEnableVdiskNumeric(v int64)`

SetDiskDsdEnableVdiskNumeric sets DiskDsdEnableVdiskNumeric field to given value.

### HasDiskDsdEnableVdiskNumeric

`func (o *DiskGroupsResourceInner) HasDiskDsdEnableVdiskNumeric() bool`

HasDiskDsdEnableVdiskNumeric returns a boolean if a field has been set.

### GetDiskcount

`func (o *DiskGroupsResourceInner) GetDiskcount() int64`

GetDiskcount returns the Diskcount field if non-nil, zero value otherwise.

### GetDiskcountOk

`func (o *DiskGroupsResourceInner) GetDiskcountOk() (*int64, bool)`

GetDiskcountOk returns a tuple with the Diskcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskcount

`func (o *DiskGroupsResourceInner) SetDiskcount(v int64)`

SetDiskcount sets Diskcount field to given value.

### HasDiskcount

`func (o *DiskGroupsResourceInner) HasDiskcount() bool`

HasDiskcount returns a boolean if a field has been set.

### GetExtendedStatus

`func (o *DiskGroupsResourceInner) GetExtendedStatus() int64`

GetExtendedStatus returns the ExtendedStatus field if non-nil, zero value otherwise.

### GetExtendedStatusOk

`func (o *DiskGroupsResourceInner) GetExtendedStatusOk() (*int64, bool)`

GetExtendedStatusOk returns a tuple with the ExtendedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedStatus

`func (o *DiskGroupsResourceInner) SetExtendedStatus(v int64)`

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

`func (o *DiskGroupsResourceInner) GetFreespaceNumeric() int64`

GetFreespaceNumeric returns the FreespaceNumeric field if non-nil, zero value otherwise.

### GetFreespaceNumericOk

`func (o *DiskGroupsResourceInner) GetFreespaceNumericOk() (*int64, bool)`

GetFreespaceNumericOk returns a tuple with the FreespaceNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreespaceNumeric

`func (o *DiskGroupsResourceInner) SetFreespaceNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetHealthNumeric() int64`

GetHealthNumeric returns the HealthNumeric field if non-nil, zero value otherwise.

### GetHealthNumericOk

`func (o *DiskGroupsResourceInner) GetHealthNumericOk() (*int64, bool)`

GetHealthNumericOk returns a tuple with the HealthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthNumeric

`func (o *DiskGroupsResourceInner) SetHealthNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetHealthReasonNumeric() int64`

GetHealthReasonNumeric returns the HealthReasonNumeric field if non-nil, zero value otherwise.

### GetHealthReasonNumericOk

`func (o *DiskGroupsResourceInner) GetHealthReasonNumericOk() (*int64, bool)`

GetHealthReasonNumericOk returns a tuple with the HealthReasonNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReasonNumeric

`func (o *DiskGroupsResourceInner) SetHealthReasonNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetHealthRecommendationNumeric() int64`

GetHealthRecommendationNumeric returns the HealthRecommendationNumeric field if non-nil, zero value otherwise.

### GetHealthRecommendationNumericOk

`func (o *DiskGroupsResourceInner) GetHealthRecommendationNumericOk() (*int64, bool)`

GetHealthRecommendationNumericOk returns a tuple with the HealthRecommendationNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecommendationNumeric

`func (o *DiskGroupsResourceInner) SetHealthRecommendationNumeric(v int64)`

SetHealthRecommendationNumeric sets HealthRecommendationNumeric field to given value.

### HasHealthRecommendationNumeric

`func (o *DiskGroupsResourceInner) HasHealthRecommendationNumeric() bool`

HasHealthRecommendationNumeric returns a boolean if a field has been set.

### GetInterleavedVolumeCount

`func (o *DiskGroupsResourceInner) GetInterleavedVolumeCount() int64`

GetInterleavedVolumeCount returns the InterleavedVolumeCount field if non-nil, zero value otherwise.

### GetInterleavedVolumeCountOk

`func (o *DiskGroupsResourceInner) GetInterleavedVolumeCountOk() (*int64, bool)`

GetInterleavedVolumeCountOk returns a tuple with the InterleavedVolumeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterleavedVolumeCount

`func (o *DiskGroupsResourceInner) SetInterleavedVolumeCount(v int64)`

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

`func (o *DiskGroupsResourceInner) GetIsJobAutoAbortableNumeric() int64`

GetIsJobAutoAbortableNumeric returns the IsJobAutoAbortableNumeric field if non-nil, zero value otherwise.

### GetIsJobAutoAbortableNumericOk

`func (o *DiskGroupsResourceInner) GetIsJobAutoAbortableNumericOk() (*int64, bool)`

GetIsJobAutoAbortableNumericOk returns a tuple with the IsJobAutoAbortableNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJobAutoAbortableNumeric

`func (o *DiskGroupsResourceInner) SetIsJobAutoAbortableNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetLargestFreePartitionSpaceNumeric() int64`

GetLargestFreePartitionSpaceNumeric returns the LargestFreePartitionSpaceNumeric field if non-nil, zero value otherwise.

### GetLargestFreePartitionSpaceNumericOk

`func (o *DiskGroupsResourceInner) GetLargestFreePartitionSpaceNumericOk() (*int64, bool)`

GetLargestFreePartitionSpaceNumericOk returns a tuple with the LargestFreePartitionSpaceNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargestFreePartitionSpaceNumeric

`func (o *DiskGroupsResourceInner) SetLargestFreePartitionSpaceNumeric(v int64)`

SetLargestFreePartitionSpaceNumeric sets LargestFreePartitionSpaceNumeric field to given value.

### HasLargestFreePartitionSpaceNumeric

`func (o *DiskGroupsResourceInner) HasLargestFreePartitionSpaceNumeric() bool`

HasLargestFreePartitionSpaceNumeric returns a boolean if a field has been set.

### GetLinearVolumeBoundary

`func (o *DiskGroupsResourceInner) GetLinearVolumeBoundary() int64`

GetLinearVolumeBoundary returns the LinearVolumeBoundary field if non-nil, zero value otherwise.

### GetLinearVolumeBoundaryOk

`func (o *DiskGroupsResourceInner) GetLinearVolumeBoundaryOk() (*int64, bool)`

GetLinearVolumeBoundaryOk returns a tuple with the LinearVolumeBoundary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinearVolumeBoundary

`func (o *DiskGroupsResourceInner) SetLinearVolumeBoundary(v int64)`

SetLinearVolumeBoundary sets LinearVolumeBoundary field to given value.

### HasLinearVolumeBoundary

`func (o *DiskGroupsResourceInner) HasLinearVolumeBoundary() bool`

HasLinearVolumeBoundary returns a boolean if a field has been set.

### GetLun

`func (o *DiskGroupsResourceInner) GetLun() int64`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *DiskGroupsResourceInner) GetLunOk() (*int64, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *DiskGroupsResourceInner) SetLun(v int64)`

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

`func (o *DiskGroupsResourceInner) GetMetadataSizeNumeric() int64`

GetMetadataSizeNumeric returns the MetadataSizeNumeric field if non-nil, zero value otherwise.

### GetMetadataSizeNumericOk

`func (o *DiskGroupsResourceInner) GetMetadataSizeNumericOk() (*int64, bool)`

GetMetadataSizeNumericOk returns a tuple with the MetadataSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSizeNumeric

`func (o *DiskGroupsResourceInner) SetMetadataSizeNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetMinDriveSizeNumeric() int64`

GetMinDriveSizeNumeric returns the MinDriveSizeNumeric field if non-nil, zero value otherwise.

### GetMinDriveSizeNumericOk

`func (o *DiskGroupsResourceInner) GetMinDriveSizeNumericOk() (*int64, bool)`

GetMinDriveSizeNumericOk returns a tuple with the MinDriveSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDriveSizeNumeric

`func (o *DiskGroupsResourceInner) SetMinDriveSizeNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetNewPartitionLbaNumeric() int64`

GetNewPartitionLbaNumeric returns the NewPartitionLbaNumeric field if non-nil, zero value otherwise.

### GetNewPartitionLbaNumericOk

`func (o *DiskGroupsResourceInner) GetNewPartitionLbaNumericOk() (*int64, bool)`

GetNewPartitionLbaNumericOk returns a tuple with the NewPartitionLbaNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPartitionLbaNumeric

`func (o *DiskGroupsResourceInner) SetNewPartitionLbaNumeric(v int64)`

SetNewPartitionLbaNumeric sets NewPartitionLbaNumeric field to given value.

### HasNewPartitionLbaNumeric

`func (o *DiskGroupsResourceInner) HasNewPartitionLbaNumeric() bool`

HasNewPartitionLbaNumeric returns a boolean if a field has been set.

### GetNumArrayPartitions

`func (o *DiskGroupsResourceInner) GetNumArrayPartitions() int64`

GetNumArrayPartitions returns the NumArrayPartitions field if non-nil, zero value otherwise.

### GetNumArrayPartitionsOk

`func (o *DiskGroupsResourceInner) GetNumArrayPartitionsOk() (*int64, bool)`

GetNumArrayPartitionsOk returns a tuple with the NumArrayPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumArrayPartitions

`func (o *DiskGroupsResourceInner) SetNumArrayPartitions(v int64)`

SetNumArrayPartitions sets NumArrayPartitions field to given value.

### HasNumArrayPartitions

`func (o *DiskGroupsResourceInner) HasNumArrayPartitions() bool`

HasNumArrayPartitions returns a boolean if a field has been set.

### GetNumDrivesPerLowLevelArray

`func (o *DiskGroupsResourceInner) GetNumDrivesPerLowLevelArray() int64`

GetNumDrivesPerLowLevelArray returns the NumDrivesPerLowLevelArray field if non-nil, zero value otherwise.

### GetNumDrivesPerLowLevelArrayOk

`func (o *DiskGroupsResourceInner) GetNumDrivesPerLowLevelArrayOk() (*int64, bool)`

GetNumDrivesPerLowLevelArrayOk returns a tuple with the NumDrivesPerLowLevelArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDrivesPerLowLevelArray

`func (o *DiskGroupsResourceInner) SetNumDrivesPerLowLevelArray(v int64)`

SetNumDrivesPerLowLevelArray sets NumDrivesPerLowLevelArray field to given value.

### HasNumDrivesPerLowLevelArray

`func (o *DiskGroupsResourceInner) HasNumDrivesPerLowLevelArray() bool`

HasNumDrivesPerLowLevelArray returns a boolean if a field has been set.

### GetNumExpansionPartitions

`func (o *DiskGroupsResourceInner) GetNumExpansionPartitions() int64`

GetNumExpansionPartitions returns the NumExpansionPartitions field if non-nil, zero value otherwise.

### GetNumExpansionPartitionsOk

`func (o *DiskGroupsResourceInner) GetNumExpansionPartitionsOk() (*int64, bool)`

GetNumExpansionPartitionsOk returns a tuple with the NumExpansionPartitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumExpansionPartitions

`func (o *DiskGroupsResourceInner) SetNumExpansionPartitions(v int64)`

SetNumExpansionPartitions sets NumExpansionPartitions field to given value.

### HasNumExpansionPartitions

`func (o *DiskGroupsResourceInner) HasNumExpansionPartitions() bool`

HasNumExpansionPartitions returns a boolean if a field has been set.

### GetNumPartitionSegments

`func (o *DiskGroupsResourceInner) GetNumPartitionSegments() int64`

GetNumPartitionSegments returns the NumPartitionSegments field if non-nil, zero value otherwise.

### GetNumPartitionSegmentsOk

`func (o *DiskGroupsResourceInner) GetNumPartitionSegmentsOk() (*int64, bool)`

GetNumPartitionSegmentsOk returns a tuple with the NumPartitionSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPartitionSegments

`func (o *DiskGroupsResourceInner) SetNumPartitionSegments(v int64)`

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

`func (o *DiskGroupsResourceInner) GetOverheadNumeric() int64`

GetOverheadNumeric returns the OverheadNumeric field if non-nil, zero value otherwise.

### GetOverheadNumericOk

`func (o *DiskGroupsResourceInner) GetOverheadNumericOk() (*int64, bool)`

GetOverheadNumericOk returns a tuple with the OverheadNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverheadNumeric

`func (o *DiskGroupsResourceInner) SetOverheadNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetOwnerNumeric() int64`

GetOwnerNumeric returns the OwnerNumeric field if non-nil, zero value otherwise.

### GetOwnerNumericOk

`func (o *DiskGroupsResourceInner) GetOwnerNumericOk() (*int64, bool)`

GetOwnerNumericOk returns a tuple with the OwnerNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerNumeric

`func (o *DiskGroupsResourceInner) SetOwnerNumeric(v int64)`

SetOwnerNumeric sets OwnerNumeric field to given value.

### HasOwnerNumeric

`func (o *DiskGroupsResourceInner) HasOwnerNumeric() bool`

HasOwnerNumeric returns a boolean if a field has been set.

### GetPerformanceRank

`func (o *DiskGroupsResourceInner) GetPerformanceRank() int64`

GetPerformanceRank returns the PerformanceRank field if non-nil, zero value otherwise.

### GetPerformanceRankOk

`func (o *DiskGroupsResourceInner) GetPerformanceRankOk() (*int64, bool)`

GetPerformanceRankOk returns a tuple with the PerformanceRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceRank

`func (o *DiskGroupsResourceInner) SetPerformanceRank(v int64)`

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

`func (o *DiskGroupsResourceInner) GetPoolPercentage() int64`

GetPoolPercentage returns the PoolPercentage field if non-nil, zero value otherwise.

### GetPoolPercentageOk

`func (o *DiskGroupsResourceInner) GetPoolPercentageOk() (*int64, bool)`

GetPoolPercentageOk returns a tuple with the PoolPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolPercentage

`func (o *DiskGroupsResourceInner) SetPoolPercentage(v int64)`

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

`func (o *DiskGroupsResourceInner) GetPoolSectorFormatNumeric() int64`

GetPoolSectorFormatNumeric returns the PoolSectorFormatNumeric field if non-nil, zero value otherwise.

### GetPoolSectorFormatNumericOk

`func (o *DiskGroupsResourceInner) GetPoolSectorFormatNumericOk() (*int64, bool)`

GetPoolSectorFormatNumericOk returns a tuple with the PoolSectorFormatNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSectorFormatNumeric

`func (o *DiskGroupsResourceInner) SetPoolSectorFormatNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetPreferredOwnerNumeric() int64`

GetPreferredOwnerNumeric returns the PreferredOwnerNumeric field if non-nil, zero value otherwise.

### GetPreferredOwnerNumericOk

`func (o *DiskGroupsResourceInner) GetPreferredOwnerNumericOk() (*int64, bool)`

GetPreferredOwnerNumericOk returns a tuple with the PreferredOwnerNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredOwnerNumeric

`func (o *DiskGroupsResourceInner) SetPreferredOwnerNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetRaidtypeNumeric() int64`

GetRaidtypeNumeric returns the RaidtypeNumeric field if non-nil, zero value otherwise.

### GetRaidtypeNumericOk

`func (o *DiskGroupsResourceInner) GetRaidtypeNumericOk() (*int64, bool)`

GetRaidtypeNumericOk returns a tuple with the RaidtypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidtypeNumeric

`func (o *DiskGroupsResourceInner) SetRaidtypeNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetRawSizeNumeric() int64`

GetRawSizeNumeric returns the RawSizeNumeric field if non-nil, zero value otherwise.

### GetRawSizeNumericOk

`func (o *DiskGroupsResourceInner) GetRawSizeNumericOk() (*int64, bool)`

GetRawSizeNumericOk returns a tuple with the RawSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSizeNumeric

`func (o *DiskGroupsResourceInner) SetRawSizeNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetReadAheadEnabledNumeric() int64`

GetReadAheadEnabledNumeric returns the ReadAheadEnabledNumeric field if non-nil, zero value otherwise.

### GetReadAheadEnabledNumericOk

`func (o *DiskGroupsResourceInner) GetReadAheadEnabledNumericOk() (*int64, bool)`

GetReadAheadEnabledNumericOk returns a tuple with the ReadAheadEnabledNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadEnabledNumeric

`func (o *DiskGroupsResourceInner) SetReadAheadEnabledNumeric(v int64)`

SetReadAheadEnabledNumeric sets ReadAheadEnabledNumeric field to given value.

### HasReadAheadEnabledNumeric

`func (o *DiskGroupsResourceInner) HasReadAheadEnabledNumeric() bool`

HasReadAheadEnabledNumeric returns a boolean if a field has been set.

### GetScrubDurationGoal

`func (o *DiskGroupsResourceInner) GetScrubDurationGoal() int64`

GetScrubDurationGoal returns the ScrubDurationGoal field if non-nil, zero value otherwise.

### GetScrubDurationGoalOk

`func (o *DiskGroupsResourceInner) GetScrubDurationGoalOk() (*int64, bool)`

GetScrubDurationGoalOk returns a tuple with the ScrubDurationGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrubDurationGoal

`func (o *DiskGroupsResourceInner) SetScrubDurationGoal(v int64)`

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

`func (o *DiskGroupsResourceInner) GetSizeNumeric() int64`

GetSizeNumeric returns the SizeNumeric field if non-nil, zero value otherwise.

### GetSizeNumericOk

`func (o *DiskGroupsResourceInner) GetSizeNumericOk() (*int64, bool)`

GetSizeNumericOk returns a tuple with the SizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeNumeric

`func (o *DiskGroupsResourceInner) SetSizeNumeric(v int64)`

SetSizeNumeric sets SizeNumeric field to given value.

### HasSizeNumeric

`func (o *DiskGroupsResourceInner) HasSizeNumeric() bool`

HasSizeNumeric returns a boolean if a field has been set.

### GetSparecount

`func (o *DiskGroupsResourceInner) GetSparecount() int64`

GetSparecount returns the Sparecount field if non-nil, zero value otherwise.

### GetSparecountOk

`func (o *DiskGroupsResourceInner) GetSparecountOk() (*int64, bool)`

GetSparecountOk returns a tuple with the Sparecount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparecount

`func (o *DiskGroupsResourceInner) SetSparecount(v int64)`

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

`func (o *DiskGroupsResourceInner) GetSpearNumeric() int64`

GetSpearNumeric returns the SpearNumeric field if non-nil, zero value otherwise.

### GetSpearNumericOk

`func (o *DiskGroupsResourceInner) GetSpearNumericOk() (*int64, bool)`

GetSpearNumericOk returns a tuple with the SpearNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpearNumeric

`func (o *DiskGroupsResourceInner) SetSpearNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetStatusNumeric() int64`

GetStatusNumeric returns the StatusNumeric field if non-nil, zero value otherwise.

### GetStatusNumericOk

`func (o *DiskGroupsResourceInner) GetStatusNumericOk() (*int64, bool)`

GetStatusNumericOk returns a tuple with the StatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusNumeric

`func (o *DiskGroupsResourceInner) SetStatusNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetStorageTierNumeric() int64`

GetStorageTierNumeric returns the StorageTierNumeric field if non-nil, zero value otherwise.

### GetStorageTierNumericOk

`func (o *DiskGroupsResourceInner) GetStorageTierNumericOk() (*int64, bool)`

GetStorageTierNumericOk returns a tuple with the StorageTierNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTierNumeric

`func (o *DiskGroupsResourceInner) SetStorageTierNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetStorageTypeNumeric() int64`

GetStorageTypeNumeric returns the StorageTypeNumeric field if non-nil, zero value otherwise.

### GetStorageTypeNumericOk

`func (o *DiskGroupsResourceInner) GetStorageTypeNumericOk() (*int64, bool)`

GetStorageTypeNumericOk returns a tuple with the StorageTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypeNumeric

`func (o *DiskGroupsResourceInner) SetStorageTypeNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetStripeWidthNumeric() int64`

GetStripeWidthNumeric returns the StripeWidthNumeric field if non-nil, zero value otherwise.

### GetStripeWidthNumericOk

`func (o *DiskGroupsResourceInner) GetStripeWidthNumericOk() (*int64, bool)`

GetStripeWidthNumericOk returns a tuple with the StripeWidthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeWidthNumeric

`func (o *DiskGroupsResourceInner) SetStripeWidthNumeric(v int64)`

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

`func (o *DiskGroupsResourceInner) GetTargetSpareCapacityNumeric() int64`

GetTargetSpareCapacityNumeric returns the TargetSpareCapacityNumeric field if non-nil, zero value otherwise.

### GetTargetSpareCapacityNumericOk

`func (o *DiskGroupsResourceInner) GetTargetSpareCapacityNumericOk() (*int64, bool)`

GetTargetSpareCapacityNumericOk returns a tuple with the TargetSpareCapacityNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSpareCapacityNumeric

`func (o *DiskGroupsResourceInner) SetTargetSpareCapacityNumeric(v int64)`

SetTargetSpareCapacityNumeric sets TargetSpareCapacityNumeric field to given value.

### HasTargetSpareCapacityNumeric

`func (o *DiskGroupsResourceInner) HasTargetSpareCapacityNumeric() bool`

HasTargetSpareCapacityNumeric returns a boolean if a field has been set.

### GetTotalPages

`func (o *DiskGroupsResourceInner) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *DiskGroupsResourceInner) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *DiskGroupsResourceInner) SetTotalPages(v int64)`

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

`func (o *DiskGroupsResourceInner) GetWriteBackEnabledNumeric() int64`

GetWriteBackEnabledNumeric returns the WriteBackEnabledNumeric field if non-nil, zero value otherwise.

### GetWriteBackEnabledNumericOk

`func (o *DiskGroupsResourceInner) GetWriteBackEnabledNumericOk() (*int64, bool)`

GetWriteBackEnabledNumericOk returns a tuple with the WriteBackEnabledNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackEnabledNumeric

`func (o *DiskGroupsResourceInner) SetWriteBackEnabledNumeric(v int64)`

SetWriteBackEnabledNumeric sets WriteBackEnabledNumeric field to given value.

### HasWriteBackEnabledNumeric

`func (o *DiskGroupsResourceInner) HasWriteBackEnabledNumeric() bool`

HasWriteBackEnabledNumeric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


