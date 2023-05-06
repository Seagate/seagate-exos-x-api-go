# VolumesResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**AllocateReservedPagesFirst** | Pointer to **string** | Allocate Reserved Pages First | [optional] 
<<<<<<< HEAD
**AllocateReservedPagesFirstNumeric** | Pointer to **int64** | Allocate Reserved Pages First( In numeric form ) | [optional] 
**AllocatedSize** | Pointer to **string** |  | [optional] 
**AllocatedSizeNumeric** | Pointer to **int64** |  | [optional] 
**AllowedStorageTiers** | Pointer to **string** | Allowed Virtual Pool Tiers | [optional] 
**AllowedStorageTiersNumeric** | Pointer to **int64** | Allowed Virtual Pool Tiers( In numeric form ) | [optional] 
**Attributes** | Pointer to **string** | Indicates if the disk is single-pathed | [optional] 
**Blocks** | Pointer to **int64** | The size in blocks | [optional] 
**Blocksize** | Pointer to **int64** |  | [optional] 
**CacheOptimization** | Pointer to **string** | Whether cache is optimized for sequential and random reads, or sequential only | [optional] 
**CacheOptimizationNumeric** | Pointer to **int64** | Whether cache is optimized for sequential and random reads, or sequential only( In numeric form ) | [optional] 
**Capabilities** | Pointer to **string** |  | [optional] 
**Compression** | Pointer to **string** |  | [optional] 
**CompressionEfficiency** | Pointer to **string** |  | [optional] 
**CompressionNumeric** | Pointer to **int64** |  | [optional] 
**ContainerName** | Pointer to **string** | Container Name | [optional] 
**ContainerSerial** | Pointer to **string** | Container Serial Number | [optional] 
**CreationDateTime** | Pointer to **string** |  | [optional] 
**CreationDateTimeNumeric** | Pointer to **int64** |  | [optional] 
**CsCopyDest** | Pointer to **string** |  | [optional] 
**CsCopyDestNumeric** | Pointer to **int64** |  | [optional] 
**CsCopySrc** | Pointer to **string** |  | [optional] 
**CsCopySrcNumeric** | Pointer to **int64** |  | [optional] 
**CsPrimary** | Pointer to **string** |  | [optional] 
**CsPrimaryNumeric** | Pointer to **int64** |  | [optional] 
**CsReplicationRole** | Pointer to **string** |  | [optional] 
**CsSecondary** | Pointer to **string** |  | [optional] 
**CsSecondaryNumeric** | Pointer to **int64** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**GroupKey** | Pointer to **string** | Durable ID of a Management Group | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**HealthNumeric** | Pointer to **int64** |  | [optional] 
**HealthReason** | Pointer to **string** |  | [optional] 
**HealthRecommendation** | Pointer to **string** |  | [optional] 
**Interleaved** | Pointer to **string** |  | [optional] 
**InterleavedNumeric** | Pointer to **int64** |  | [optional] 
**LargeVirtualExtents** | Pointer to **string** | Sequentially optimize zone allocation for virtual volume | [optional] 
**LargeVirtualExtentsNumeric** | Pointer to **int64** | Sequentially optimize zone allocation for virtual volume( In numeric form ) | [optional] 
**MetadataInUse** | Pointer to **string** | Pool Metadata currently being used by volume | [optional] 
**MetadataInUseNumeric** | Pointer to **int64** | Pool Metadata currently being used by volume( In numeric form ) | [optional] 
**Owner** | Pointer to **string** | Controller owning the component | [optional] 
**OwnerNumeric** | Pointer to **int64** | Controller owning the component( In numeric form ) | [optional] 
**PiFormat** | Pointer to **string** | Used to describe the Host Protection Information | [optional] 
**PiFormatNumeric** | Pointer to **int64** | Used to describe the Host Protection Information( In numeric form ) | [optional] 
**PreferredOwner** | Pointer to **string** | Configured owner | [optional] 
**PreferredOwnerNumeric** | Pointer to **int64** | Configured owner( In numeric form ) | [optional] 
**Progress** | Pointer to **string** | Progress in percent for the volume-copy operation | [optional] 
**ProgressNumeric** | Pointer to **int64** | Progress in percent for the volume-copy operation( In numeric form ) | [optional] 
**Raidtype** | Pointer to **string** | RAID level | [optional] 
**RaidtypeNumeric** | Pointer to **int64** | RAID level( In numeric form ) | [optional] 
**ReadAheadSize** | Pointer to **string** | Read-ahead cache setting | [optional] 
**ReadAheadSizeNumeric** | Pointer to **int64** | Read-ahead cache setting( In numeric form ) | [optional] 
**ReplicationSet** | Pointer to **string** |  | [optional] 
**ReservedSizeInPages** | Pointer to **int64** | Reserved Size In Pages | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** | The size or capacity formatted with the current session base, precision, and units | [optional] 
**SizeNumeric** | Pointer to **int64** | The size or capacity formatted with the current session base, precision, and units( In numeric form ) | [optional] 
**SnapPool** | Pointer to **string** | Snap pool associated with this volume | [optional] 
**Snapshot** | Pointer to **string** |  | [optional] 
**SnapshotRetentionPriority** | Pointer to **string** | Retention priority of the snapshot | [optional] 
**SnapshotRetentionPriorityNumeric** | Pointer to **int64** | Retention priority of the snapshot( In numeric form ) | [optional] 
**StoragePoolName** | Pointer to **string** | User-defined name for the pool | [optional] 
**StoragePoolsUrl** | Pointer to **string** |  | [optional] 
**StorageType** | Pointer to **string** | Storage type | [optional] 
**StorageTypeNumeric** | Pointer to **int64** | Storage type( In numeric form ) | [optional] 
**ThresholdPercentOfPool** | Pointer to **string** | Threshold Percent of Pool | [optional] 
**TierAffinity** | Pointer to **string** |  | [optional] 
**TierAffinityNumeric** | Pointer to **int64** |  | [optional] 
**TotalSize** | Pointer to **string** | The total size formatted using the session settings for base, precision, and units | [optional] 
**TotalSizeNumeric** | Pointer to **int64** | The total size formatted using the session settings for base, precision, and units( In numeric form ) | [optional] 
=======
**AllocateReservedPagesFirstNumeric** | Pointer to **int32** | Allocate Reserved Pages First( In numeric form ) | [optional] 
**AllocatedSize** | Pointer to **string** |  | [optional] 
**AllocatedSizeNumeric** | Pointer to **int32** |  | [optional] 
**AllowedStorageTiers** | Pointer to **string** | Allowed Virtual Pool Tiers | [optional] 
**AllowedStorageTiersNumeric** | Pointer to **int32** | Allowed Virtual Pool Tiers( In numeric form ) | [optional] 
**Attributes** | Pointer to **string** | Indicates if the disk is single-pathed | [optional] 
**Blocks** | Pointer to **int32** | The size in blocks | [optional] 
**Blocksize** | Pointer to **int32** |  | [optional] 
**CacheOptimization** | Pointer to **string** | Whether cache is optimized for sequential and random reads, or sequential only | [optional] 
**CacheOptimizationNumeric** | Pointer to **int32** | Whether cache is optimized for sequential and random reads, or sequential only( In numeric form ) | [optional] 
**Capabilities** | Pointer to **string** |  | [optional] 
**Compression** | Pointer to **string** |  | [optional] 
**CompressionEfficiency** | Pointer to **string** |  | [optional] 
**CompressionNumeric** | Pointer to **int32** |  | [optional] 
**ContainerName** | Pointer to **string** | Container Name | [optional] 
**ContainerSerial** | Pointer to **string** | Container Serial Number | [optional] 
**CreationDateTime** | Pointer to **string** |  | [optional] 
**CreationDateTimeNumeric** | Pointer to **int32** |  | [optional] 
**CsCopyDest** | Pointer to **string** |  | [optional] 
**CsCopyDestNumeric** | Pointer to **int32** |  | [optional] 
**CsCopySrc** | Pointer to **string** |  | [optional] 
**CsCopySrcNumeric** | Pointer to **int32** |  | [optional] 
**CsPrimary** | Pointer to **string** |  | [optional] 
**CsPrimaryNumeric** | Pointer to **int32** |  | [optional] 
**CsReplicationRole** | Pointer to **string** |  | [optional] 
**CsSecondary** | Pointer to **string** |  | [optional] 
**CsSecondaryNumeric** | Pointer to **int32** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**GroupKey** | Pointer to **string** | Durable ID of a Management Group | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**HealthNumeric** | Pointer to **int32** |  | [optional] 
**HealthReason** | Pointer to **string** |  | [optional] 
**HealthRecommendation** | Pointer to **string** |  | [optional] 
**Interleaved** | Pointer to **string** |  | [optional] 
**InterleavedNumeric** | Pointer to **int32** |  | [optional] 
**LargeVirtualExtents** | Pointer to **string** | Sequentially optimize zone allocation for virtual volume | [optional] 
**LargeVirtualExtentsNumeric** | Pointer to **int32** | Sequentially optimize zone allocation for virtual volume( In numeric form ) | [optional] 
**MetadataInUse** | Pointer to **string** | Pool Metadata currently being used by volume | [optional] 
**MetadataInUseNumeric** | Pointer to **int32** | Pool Metadata currently being used by volume( In numeric form ) | [optional] 
**Owner** | Pointer to **string** | Controller owning the component | [optional] 
**OwnerNumeric** | Pointer to **int32** | Controller owning the component( In numeric form ) | [optional] 
**PiFormat** | Pointer to **string** | Used to describe the Host Protection Information | [optional] 
**PiFormatNumeric** | Pointer to **int32** | Used to describe the Host Protection Information( In numeric form ) | [optional] 
**PreferredOwner** | Pointer to **string** | Configured owner | [optional] 
**PreferredOwnerNumeric** | Pointer to **int32** | Configured owner( In numeric form ) | [optional] 
**Progress** | Pointer to **string** | Progress in percent for the volume-copy operation | [optional] 
**ProgressNumeric** | Pointer to **int32** | Progress in percent for the volume-copy operation( In numeric form ) | [optional] 
**Raidtype** | Pointer to **string** | RAID level | [optional] 
**RaidtypeNumeric** | Pointer to **int32** | RAID level( In numeric form ) | [optional] 
**ReadAheadSize** | Pointer to **string** | Read-ahead cache setting | [optional] 
**ReadAheadSizeNumeric** | Pointer to **int32** | Read-ahead cache setting( In numeric form ) | [optional] 
**ReplicationSet** | Pointer to **string** |  | [optional] 
**ReservedSizeInPages** | Pointer to **int32** | Reserved Size In Pages | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** | The size or capacity formatted with the current session base, precision, and units | [optional] 
**SizeNumeric** | Pointer to **int32** | The size or capacity formatted with the current session base, precision, and units( In numeric form ) | [optional] 
**SnapPool** | Pointer to **string** | Snap pool associated with this volume | [optional] 
**Snapshot** | Pointer to **string** |  | [optional] 
**SnapshotRetentionPriority** | Pointer to **string** | Retention priority of the snapshot | [optional] 
**SnapshotRetentionPriorityNumeric** | Pointer to **int32** | Retention priority of the snapshot( In numeric form ) | [optional] 
**StoragePoolName** | Pointer to **string** | User-defined name for the pool | [optional] 
**StoragePoolsUrl** | Pointer to **string** |  | [optional] 
**StorageType** | Pointer to **string** | Storage type | [optional] 
**StorageTypeNumeric** | Pointer to **int32** | Storage type( In numeric form ) | [optional] 
**ThresholdPercentOfPool** | Pointer to **string** | Threshold Percent of Pool | [optional] 
**TierAffinity** | Pointer to **string** |  | [optional] 
**TierAffinityNumeric** | Pointer to **int32** |  | [optional] 
**TotalSize** | Pointer to **string** | The total size formatted using the session settings for base, precision, and units | [optional] 
**TotalSizeNumeric** | Pointer to **int32** | The total size formatted using the session settings for base, precision, and units( In numeric form ) | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
**Url** | Pointer to **string** | The resource URL | [optional] 
**VirtualDiskName** | Pointer to **string** | The name of the pool that contains the volume | [optional] 
**VirtualDiskSerial** | Pointer to **string** | Unique serial number for the disk group | [optional] 
**VolumeClass** | Pointer to **string** |  | [optional] 
<<<<<<< HEAD
**VolumeClassNumeric** | Pointer to **int64** |  | [optional] 
=======
**VolumeClassNumeric** | Pointer to **int32** |  | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
**VolumeDescription** | Pointer to **string** | Provides a user-defined volume identifier | [optional] 
**VolumeGroup** | Pointer to **string** | Volume Group | [optional] 
**VolumeName** | Pointer to **string** | User-defined name for the volume | [optional] 
**VolumeParent** | Pointer to **string** | Serial number of the associated volume | [optional] 
**VolumeQualifier** | Pointer to **string** |  | [optional] 
<<<<<<< HEAD
**VolumeQualifierNumeric** | Pointer to **int64** |  | [optional] 
**VolumeType** | Pointer to **string** | Indicates whether the volume is being used for Snapshots | [optional] 
**VolumeTypeNumeric** | Pointer to **int64** | Indicates whether the volume is being used for Snapshots( In numeric form ) | [optional] 
**VolumeUsageType** | Pointer to **string** | Indicates the usage of volume | [optional] 
**VolumeUsageTypeNumeric** | Pointer to **int64** | Indicates the usage of volume( In numeric form ) | [optional] 
**WritePolicy** | Pointer to **string** | The write-back cache mode | [optional] 
**WritePolicyNumeric** | Pointer to **int64** | The write-back cache mode( In numeric form ) | [optional] 
**Wwn** | Pointer to **string** | World Wide Name | [optional] 
**ZeroInitPageOnAllocation** | Pointer to **string** | Zero Init Page On Allocation | [optional] 
**ZeroInitPageOnAllocationNumeric** | Pointer to **int64** | Zero Init Page On Allocation( In numeric form ) | [optional] 
=======
**VolumeQualifierNumeric** | Pointer to **int32** |  | [optional] 
**VolumeType** | Pointer to **string** | Indicates whether the volume is being used for Snapshots | [optional] 
**VolumeTypeNumeric** | Pointer to **int32** | Indicates whether the volume is being used for Snapshots( In numeric form ) | [optional] 
**VolumeUsageType** | Pointer to **string** | Indicates the usage of volume | [optional] 
**VolumeUsageTypeNumeric** | Pointer to **int32** | Indicates the usage of volume( In numeric form ) | [optional] 
**WritePolicy** | Pointer to **string** | The write-back cache mode | [optional] 
**WritePolicyNumeric** | Pointer to **int32** | The write-back cache mode( In numeric form ) | [optional] 
**Wwn** | Pointer to **string** | World Wide Name | [optional] 
**ZeroInitPageOnAllocation** | Pointer to **string** | Zero Init Page On Allocation | [optional] 
**ZeroInitPageOnAllocationNumeric** | Pointer to **int32** | Zero Init Page On Allocation( In numeric form ) | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

## Methods

### NewVolumesResourceInner

`func NewVolumesResourceInner() *VolumesResourceInner`

NewVolumesResourceInner instantiates a new VolumesResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumesResourceInnerWithDefaults

`func NewVolumesResourceInnerWithDefaults() *VolumesResourceInner`

NewVolumesResourceInnerWithDefaults instantiates a new VolumesResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *VolumesResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *VolumesResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *VolumesResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *VolumesResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *VolumesResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VolumesResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VolumesResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *VolumesResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetAllocateReservedPagesFirst

`func (o *VolumesResourceInner) GetAllocateReservedPagesFirst() string`

GetAllocateReservedPagesFirst returns the AllocateReservedPagesFirst field if non-nil, zero value otherwise.

### GetAllocateReservedPagesFirstOk

`func (o *VolumesResourceInner) GetAllocateReservedPagesFirstOk() (*string, bool)`

GetAllocateReservedPagesFirstOk returns a tuple with the AllocateReservedPagesFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocateReservedPagesFirst

`func (o *VolumesResourceInner) SetAllocateReservedPagesFirst(v string)`

SetAllocateReservedPagesFirst sets AllocateReservedPagesFirst field to given value.

### HasAllocateReservedPagesFirst

`func (o *VolumesResourceInner) HasAllocateReservedPagesFirst() bool`

HasAllocateReservedPagesFirst returns a boolean if a field has been set.

### GetAllocateReservedPagesFirstNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetAllocateReservedPagesFirstNumeric() int64`
=======
`func (o *VolumesResourceInner) GetAllocateReservedPagesFirstNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAllocateReservedPagesFirstNumeric returns the AllocateReservedPagesFirstNumeric field if non-nil, zero value otherwise.

### GetAllocateReservedPagesFirstNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetAllocateReservedPagesFirstNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetAllocateReservedPagesFirstNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAllocateReservedPagesFirstNumericOk returns a tuple with the AllocateReservedPagesFirstNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocateReservedPagesFirstNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetAllocateReservedPagesFirstNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetAllocateReservedPagesFirstNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetAllocateReservedPagesFirstNumeric sets AllocateReservedPagesFirstNumeric field to given value.

### HasAllocateReservedPagesFirstNumeric

`func (o *VolumesResourceInner) HasAllocateReservedPagesFirstNumeric() bool`

HasAllocateReservedPagesFirstNumeric returns a boolean if a field has been set.

### GetAllocatedSize

`func (o *VolumesResourceInner) GetAllocatedSize() string`

GetAllocatedSize returns the AllocatedSize field if non-nil, zero value otherwise.

### GetAllocatedSizeOk

`func (o *VolumesResourceInner) GetAllocatedSizeOk() (*string, bool)`

GetAllocatedSizeOk returns a tuple with the AllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSize

`func (o *VolumesResourceInner) SetAllocatedSize(v string)`

SetAllocatedSize sets AllocatedSize field to given value.

### HasAllocatedSize

`func (o *VolumesResourceInner) HasAllocatedSize() bool`

HasAllocatedSize returns a boolean if a field has been set.

### GetAllocatedSizeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetAllocatedSizeNumeric() int64`
=======
`func (o *VolumesResourceInner) GetAllocatedSizeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAllocatedSizeNumeric returns the AllocatedSizeNumeric field if non-nil, zero value otherwise.

### GetAllocatedSizeNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetAllocatedSizeNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetAllocatedSizeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAllocatedSizeNumericOk returns a tuple with the AllocatedSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSizeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetAllocatedSizeNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetAllocatedSizeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetAllocatedSizeNumeric sets AllocatedSizeNumeric field to given value.

### HasAllocatedSizeNumeric

`func (o *VolumesResourceInner) HasAllocatedSizeNumeric() bool`

HasAllocatedSizeNumeric returns a boolean if a field has been set.

### GetAllowedStorageTiers

`func (o *VolumesResourceInner) GetAllowedStorageTiers() string`

GetAllowedStorageTiers returns the AllowedStorageTiers field if non-nil, zero value otherwise.

### GetAllowedStorageTiersOk

`func (o *VolumesResourceInner) GetAllowedStorageTiersOk() (*string, bool)`

GetAllowedStorageTiersOk returns a tuple with the AllowedStorageTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedStorageTiers

`func (o *VolumesResourceInner) SetAllowedStorageTiers(v string)`

SetAllowedStorageTiers sets AllowedStorageTiers field to given value.

### HasAllowedStorageTiers

`func (o *VolumesResourceInner) HasAllowedStorageTiers() bool`

HasAllowedStorageTiers returns a boolean if a field has been set.

### GetAllowedStorageTiersNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetAllowedStorageTiersNumeric() int64`
=======
`func (o *VolumesResourceInner) GetAllowedStorageTiersNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAllowedStorageTiersNumeric returns the AllowedStorageTiersNumeric field if non-nil, zero value otherwise.

### GetAllowedStorageTiersNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetAllowedStorageTiersNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetAllowedStorageTiersNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAllowedStorageTiersNumericOk returns a tuple with the AllowedStorageTiersNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedStorageTiersNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetAllowedStorageTiersNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetAllowedStorageTiersNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetAllowedStorageTiersNumeric sets AllowedStorageTiersNumeric field to given value.

### HasAllowedStorageTiersNumeric

`func (o *VolumesResourceInner) HasAllowedStorageTiersNumeric() bool`

HasAllowedStorageTiersNumeric returns a boolean if a field has been set.

### GetAttributes

`func (o *VolumesResourceInner) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *VolumesResourceInner) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *VolumesResourceInner) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *VolumesResourceInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetBlocks

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetBlocks() int64`
=======
`func (o *VolumesResourceInner) GetBlocks() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetBlocksOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetBlocksOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetBlocks(v int64)`
=======
`func (o *VolumesResourceInner) SetBlocks(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *VolumesResourceInner) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetBlocksize

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetBlocksize() int64`
=======
`func (o *VolumesResourceInner) GetBlocksize() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetBlocksize returns the Blocksize field if non-nil, zero value otherwise.

### GetBlocksizeOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetBlocksizeOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetBlocksizeOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetBlocksizeOk returns a tuple with the Blocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocksize

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetBlocksize(v int64)`
=======
`func (o *VolumesResourceInner) SetBlocksize(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetBlocksize sets Blocksize field to given value.

### HasBlocksize

`func (o *VolumesResourceInner) HasBlocksize() bool`

HasBlocksize returns a boolean if a field has been set.

### GetCacheOptimization

`func (o *VolumesResourceInner) GetCacheOptimization() string`

GetCacheOptimization returns the CacheOptimization field if non-nil, zero value otherwise.

### GetCacheOptimizationOk

`func (o *VolumesResourceInner) GetCacheOptimizationOk() (*string, bool)`

GetCacheOptimizationOk returns a tuple with the CacheOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheOptimization

`func (o *VolumesResourceInner) SetCacheOptimization(v string)`

SetCacheOptimization sets CacheOptimization field to given value.

### HasCacheOptimization

`func (o *VolumesResourceInner) HasCacheOptimization() bool`

HasCacheOptimization returns a boolean if a field has been set.

### GetCacheOptimizationNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetCacheOptimizationNumeric() int64`
=======
`func (o *VolumesResourceInner) GetCacheOptimizationNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCacheOptimizationNumeric returns the CacheOptimizationNumeric field if non-nil, zero value otherwise.

### GetCacheOptimizationNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetCacheOptimizationNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetCacheOptimizationNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCacheOptimizationNumericOk returns a tuple with the CacheOptimizationNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheOptimizationNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetCacheOptimizationNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetCacheOptimizationNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetCacheOptimizationNumeric sets CacheOptimizationNumeric field to given value.

### HasCacheOptimizationNumeric

`func (o *VolumesResourceInner) HasCacheOptimizationNumeric() bool`

HasCacheOptimizationNumeric returns a boolean if a field has been set.

### GetCapabilities

`func (o *VolumesResourceInner) GetCapabilities() string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *VolumesResourceInner) GetCapabilitiesOk() (*string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *VolumesResourceInner) SetCapabilities(v string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *VolumesResourceInner) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCompression

`func (o *VolumesResourceInner) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *VolumesResourceInner) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *VolumesResourceInner) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *VolumesResourceInner) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetCompressionEfficiency

`func (o *VolumesResourceInner) GetCompressionEfficiency() string`

GetCompressionEfficiency returns the CompressionEfficiency field if non-nil, zero value otherwise.

### GetCompressionEfficiencyOk

`func (o *VolumesResourceInner) GetCompressionEfficiencyOk() (*string, bool)`

GetCompressionEfficiencyOk returns a tuple with the CompressionEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEfficiency

`func (o *VolumesResourceInner) SetCompressionEfficiency(v string)`

SetCompressionEfficiency sets CompressionEfficiency field to given value.

### HasCompressionEfficiency

`func (o *VolumesResourceInner) HasCompressionEfficiency() bool`

HasCompressionEfficiency returns a boolean if a field has been set.

### GetCompressionNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetCompressionNumeric() int64`
=======
`func (o *VolumesResourceInner) GetCompressionNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCompressionNumeric returns the CompressionNumeric field if non-nil, zero value otherwise.

### GetCompressionNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetCompressionNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetCompressionNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCompressionNumericOk returns a tuple with the CompressionNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetCompressionNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetCompressionNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetCompressionNumeric sets CompressionNumeric field to given value.

### HasCompressionNumeric

`func (o *VolumesResourceInner) HasCompressionNumeric() bool`

HasCompressionNumeric returns a boolean if a field has been set.

### GetContainerName

`func (o *VolumesResourceInner) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *VolumesResourceInner) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *VolumesResourceInner) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *VolumesResourceInner) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetContainerSerial

`func (o *VolumesResourceInner) GetContainerSerial() string`

GetContainerSerial returns the ContainerSerial field if non-nil, zero value otherwise.

### GetContainerSerialOk

`func (o *VolumesResourceInner) GetContainerSerialOk() (*string, bool)`

GetContainerSerialOk returns a tuple with the ContainerSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSerial

`func (o *VolumesResourceInner) SetContainerSerial(v string)`

SetContainerSerial sets ContainerSerial field to given value.

### HasContainerSerial

`func (o *VolumesResourceInner) HasContainerSerial() bool`

HasContainerSerial returns a boolean if a field has been set.

### GetCreationDateTime

`func (o *VolumesResourceInner) GetCreationDateTime() string`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *VolumesResourceInner) GetCreationDateTimeOk() (*string, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *VolumesResourceInner) SetCreationDateTime(v string)`

SetCreationDateTime sets CreationDateTime field to given value.

### HasCreationDateTime

`func (o *VolumesResourceInner) HasCreationDateTime() bool`

HasCreationDateTime returns a boolean if a field has been set.

### GetCreationDateTimeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetCreationDateTimeNumeric() int64`
=======
`func (o *VolumesResourceInner) GetCreationDateTimeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCreationDateTimeNumeric returns the CreationDateTimeNumeric field if non-nil, zero value otherwise.

### GetCreationDateTimeNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetCreationDateTimeNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetCreationDateTimeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCreationDateTimeNumericOk returns a tuple with the CreationDateTimeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTimeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetCreationDateTimeNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetCreationDateTimeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetCreationDateTimeNumeric sets CreationDateTimeNumeric field to given value.

### HasCreationDateTimeNumeric

`func (o *VolumesResourceInner) HasCreationDateTimeNumeric() bool`

HasCreationDateTimeNumeric returns a boolean if a field has been set.

### GetCsCopyDest

`func (o *VolumesResourceInner) GetCsCopyDest() string`

GetCsCopyDest returns the CsCopyDest field if non-nil, zero value otherwise.

### GetCsCopyDestOk

`func (o *VolumesResourceInner) GetCsCopyDestOk() (*string, bool)`

GetCsCopyDestOk returns a tuple with the CsCopyDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsCopyDest

`func (o *VolumesResourceInner) SetCsCopyDest(v string)`

SetCsCopyDest sets CsCopyDest field to given value.

### HasCsCopyDest

`func (o *VolumesResourceInner) HasCsCopyDest() bool`

HasCsCopyDest returns a boolean if a field has been set.

### GetCsCopyDestNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetCsCopyDestNumeric() int64`
=======
`func (o *VolumesResourceInner) GetCsCopyDestNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCsCopyDestNumeric returns the CsCopyDestNumeric field if non-nil, zero value otherwise.

### GetCsCopyDestNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetCsCopyDestNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetCsCopyDestNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCsCopyDestNumericOk returns a tuple with the CsCopyDestNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsCopyDestNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetCsCopyDestNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetCsCopyDestNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetCsCopyDestNumeric sets CsCopyDestNumeric field to given value.

### HasCsCopyDestNumeric

`func (o *VolumesResourceInner) HasCsCopyDestNumeric() bool`

HasCsCopyDestNumeric returns a boolean if a field has been set.

### GetCsCopySrc

`func (o *VolumesResourceInner) GetCsCopySrc() string`

GetCsCopySrc returns the CsCopySrc field if non-nil, zero value otherwise.

### GetCsCopySrcOk

`func (o *VolumesResourceInner) GetCsCopySrcOk() (*string, bool)`

GetCsCopySrcOk returns a tuple with the CsCopySrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsCopySrc

`func (o *VolumesResourceInner) SetCsCopySrc(v string)`

SetCsCopySrc sets CsCopySrc field to given value.

### HasCsCopySrc

`func (o *VolumesResourceInner) HasCsCopySrc() bool`

HasCsCopySrc returns a boolean if a field has been set.

### GetCsCopySrcNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetCsCopySrcNumeric() int64`
=======
`func (o *VolumesResourceInner) GetCsCopySrcNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCsCopySrcNumeric returns the CsCopySrcNumeric field if non-nil, zero value otherwise.

### GetCsCopySrcNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetCsCopySrcNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetCsCopySrcNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCsCopySrcNumericOk returns a tuple with the CsCopySrcNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsCopySrcNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetCsCopySrcNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetCsCopySrcNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetCsCopySrcNumeric sets CsCopySrcNumeric field to given value.

### HasCsCopySrcNumeric

`func (o *VolumesResourceInner) HasCsCopySrcNumeric() bool`

HasCsCopySrcNumeric returns a boolean if a field has been set.

### GetCsPrimary

`func (o *VolumesResourceInner) GetCsPrimary() string`

GetCsPrimary returns the CsPrimary field if non-nil, zero value otherwise.

### GetCsPrimaryOk

`func (o *VolumesResourceInner) GetCsPrimaryOk() (*string, bool)`

GetCsPrimaryOk returns a tuple with the CsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsPrimary

`func (o *VolumesResourceInner) SetCsPrimary(v string)`

SetCsPrimary sets CsPrimary field to given value.

### HasCsPrimary

`func (o *VolumesResourceInner) HasCsPrimary() bool`

HasCsPrimary returns a boolean if a field has been set.

### GetCsPrimaryNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetCsPrimaryNumeric() int64`
=======
`func (o *VolumesResourceInner) GetCsPrimaryNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCsPrimaryNumeric returns the CsPrimaryNumeric field if non-nil, zero value otherwise.

### GetCsPrimaryNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetCsPrimaryNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetCsPrimaryNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCsPrimaryNumericOk returns a tuple with the CsPrimaryNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsPrimaryNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetCsPrimaryNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetCsPrimaryNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetCsPrimaryNumeric sets CsPrimaryNumeric field to given value.

### HasCsPrimaryNumeric

`func (o *VolumesResourceInner) HasCsPrimaryNumeric() bool`

HasCsPrimaryNumeric returns a boolean if a field has been set.

### GetCsReplicationRole

`func (o *VolumesResourceInner) GetCsReplicationRole() string`

GetCsReplicationRole returns the CsReplicationRole field if non-nil, zero value otherwise.

### GetCsReplicationRoleOk

`func (o *VolumesResourceInner) GetCsReplicationRoleOk() (*string, bool)`

GetCsReplicationRoleOk returns a tuple with the CsReplicationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsReplicationRole

`func (o *VolumesResourceInner) SetCsReplicationRole(v string)`

SetCsReplicationRole sets CsReplicationRole field to given value.

### HasCsReplicationRole

`func (o *VolumesResourceInner) HasCsReplicationRole() bool`

HasCsReplicationRole returns a boolean if a field has been set.

### GetCsSecondary

`func (o *VolumesResourceInner) GetCsSecondary() string`

GetCsSecondary returns the CsSecondary field if non-nil, zero value otherwise.

### GetCsSecondaryOk

`func (o *VolumesResourceInner) GetCsSecondaryOk() (*string, bool)`

GetCsSecondaryOk returns a tuple with the CsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsSecondary

`func (o *VolumesResourceInner) SetCsSecondary(v string)`

SetCsSecondary sets CsSecondary field to given value.

### HasCsSecondary

`func (o *VolumesResourceInner) HasCsSecondary() bool`

HasCsSecondary returns a boolean if a field has been set.

### GetCsSecondaryNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetCsSecondaryNumeric() int64`
=======
`func (o *VolumesResourceInner) GetCsSecondaryNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCsSecondaryNumeric returns the CsSecondaryNumeric field if non-nil, zero value otherwise.

### GetCsSecondaryNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetCsSecondaryNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetCsSecondaryNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCsSecondaryNumericOk returns a tuple with the CsSecondaryNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsSecondaryNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetCsSecondaryNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetCsSecondaryNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetCsSecondaryNumeric sets CsSecondaryNumeric field to given value.

### HasCsSecondaryNumeric

`func (o *VolumesResourceInner) HasCsSecondaryNumeric() bool`

HasCsSecondaryNumeric returns a boolean if a field has been set.

### GetDurableId

`func (o *VolumesResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *VolumesResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *VolumesResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *VolumesResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetGroupKey

`func (o *VolumesResourceInner) GetGroupKey() string`

GetGroupKey returns the GroupKey field if non-nil, zero value otherwise.

### GetGroupKeyOk

`func (o *VolumesResourceInner) GetGroupKeyOk() (*string, bool)`

GetGroupKeyOk returns a tuple with the GroupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKey

`func (o *VolumesResourceInner) SetGroupKey(v string)`

SetGroupKey sets GroupKey field to given value.

### HasGroupKey

`func (o *VolumesResourceInner) HasGroupKey() bool`

HasGroupKey returns a boolean if a field has been set.

### GetHealth

`func (o *VolumesResourceInner) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *VolumesResourceInner) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *VolumesResourceInner) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *VolumesResourceInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHealthNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetHealthNumeric() int64`
=======
`func (o *VolumesResourceInner) GetHealthNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHealthNumeric returns the HealthNumeric field if non-nil, zero value otherwise.

### GetHealthNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetHealthNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetHealthNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHealthNumericOk returns a tuple with the HealthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetHealthNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetHealthNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetHealthNumeric sets HealthNumeric field to given value.

### HasHealthNumeric

`func (o *VolumesResourceInner) HasHealthNumeric() bool`

HasHealthNumeric returns a boolean if a field has been set.

### GetHealthReason

`func (o *VolumesResourceInner) GetHealthReason() string`

GetHealthReason returns the HealthReason field if non-nil, zero value otherwise.

### GetHealthReasonOk

`func (o *VolumesResourceInner) GetHealthReasonOk() (*string, bool)`

GetHealthReasonOk returns a tuple with the HealthReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReason

`func (o *VolumesResourceInner) SetHealthReason(v string)`

SetHealthReason sets HealthReason field to given value.

### HasHealthReason

`func (o *VolumesResourceInner) HasHealthReason() bool`

HasHealthReason returns a boolean if a field has been set.

### GetHealthRecommendation

`func (o *VolumesResourceInner) GetHealthRecommendation() string`

GetHealthRecommendation returns the HealthRecommendation field if non-nil, zero value otherwise.

### GetHealthRecommendationOk

`func (o *VolumesResourceInner) GetHealthRecommendationOk() (*string, bool)`

GetHealthRecommendationOk returns a tuple with the HealthRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecommendation

`func (o *VolumesResourceInner) SetHealthRecommendation(v string)`

SetHealthRecommendation sets HealthRecommendation field to given value.

### HasHealthRecommendation

`func (o *VolumesResourceInner) HasHealthRecommendation() bool`

HasHealthRecommendation returns a boolean if a field has been set.

### GetInterleaved

`func (o *VolumesResourceInner) GetInterleaved() string`

GetInterleaved returns the Interleaved field if non-nil, zero value otherwise.

### GetInterleavedOk

`func (o *VolumesResourceInner) GetInterleavedOk() (*string, bool)`

GetInterleavedOk returns a tuple with the Interleaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterleaved

`func (o *VolumesResourceInner) SetInterleaved(v string)`

SetInterleaved sets Interleaved field to given value.

### HasInterleaved

`func (o *VolumesResourceInner) HasInterleaved() bool`

HasInterleaved returns a boolean if a field has been set.

### GetInterleavedNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetInterleavedNumeric() int64`
=======
`func (o *VolumesResourceInner) GetInterleavedNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetInterleavedNumeric returns the InterleavedNumeric field if non-nil, zero value otherwise.

### GetInterleavedNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetInterleavedNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetInterleavedNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetInterleavedNumericOk returns a tuple with the InterleavedNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterleavedNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetInterleavedNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetInterleavedNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetInterleavedNumeric sets InterleavedNumeric field to given value.

### HasInterleavedNumeric

`func (o *VolumesResourceInner) HasInterleavedNumeric() bool`

HasInterleavedNumeric returns a boolean if a field has been set.

### GetLargeVirtualExtents

`func (o *VolumesResourceInner) GetLargeVirtualExtents() string`

GetLargeVirtualExtents returns the LargeVirtualExtents field if non-nil, zero value otherwise.

### GetLargeVirtualExtentsOk

`func (o *VolumesResourceInner) GetLargeVirtualExtentsOk() (*string, bool)`

GetLargeVirtualExtentsOk returns a tuple with the LargeVirtualExtents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeVirtualExtents

`func (o *VolumesResourceInner) SetLargeVirtualExtents(v string)`

SetLargeVirtualExtents sets LargeVirtualExtents field to given value.

### HasLargeVirtualExtents

`func (o *VolumesResourceInner) HasLargeVirtualExtents() bool`

HasLargeVirtualExtents returns a boolean if a field has been set.

### GetLargeVirtualExtentsNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetLargeVirtualExtentsNumeric() int64`
=======
`func (o *VolumesResourceInner) GetLargeVirtualExtentsNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLargeVirtualExtentsNumeric returns the LargeVirtualExtentsNumeric field if non-nil, zero value otherwise.

### GetLargeVirtualExtentsNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetLargeVirtualExtentsNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetLargeVirtualExtentsNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLargeVirtualExtentsNumericOk returns a tuple with the LargeVirtualExtentsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeVirtualExtentsNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetLargeVirtualExtentsNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetLargeVirtualExtentsNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetLargeVirtualExtentsNumeric sets LargeVirtualExtentsNumeric field to given value.

### HasLargeVirtualExtentsNumeric

`func (o *VolumesResourceInner) HasLargeVirtualExtentsNumeric() bool`

HasLargeVirtualExtentsNumeric returns a boolean if a field has been set.

### GetMetadataInUse

`func (o *VolumesResourceInner) GetMetadataInUse() string`

GetMetadataInUse returns the MetadataInUse field if non-nil, zero value otherwise.

### GetMetadataInUseOk

`func (o *VolumesResourceInner) GetMetadataInUseOk() (*string, bool)`

GetMetadataInUseOk returns a tuple with the MetadataInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataInUse

`func (o *VolumesResourceInner) SetMetadataInUse(v string)`

SetMetadataInUse sets MetadataInUse field to given value.

### HasMetadataInUse

`func (o *VolumesResourceInner) HasMetadataInUse() bool`

HasMetadataInUse returns a boolean if a field has been set.

### GetMetadataInUseNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetMetadataInUseNumeric() int64`
=======
`func (o *VolumesResourceInner) GetMetadataInUseNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMetadataInUseNumeric returns the MetadataInUseNumeric field if non-nil, zero value otherwise.

### GetMetadataInUseNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetMetadataInUseNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetMetadataInUseNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMetadataInUseNumericOk returns a tuple with the MetadataInUseNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataInUseNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetMetadataInUseNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetMetadataInUseNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetMetadataInUseNumeric sets MetadataInUseNumeric field to given value.

### HasMetadataInUseNumeric

`func (o *VolumesResourceInner) HasMetadataInUseNumeric() bool`

HasMetadataInUseNumeric returns a boolean if a field has been set.

### GetOwner

`func (o *VolumesResourceInner) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VolumesResourceInner) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VolumesResourceInner) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *VolumesResourceInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnerNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetOwnerNumeric() int64`
=======
`func (o *VolumesResourceInner) GetOwnerNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetOwnerNumeric returns the OwnerNumeric field if non-nil, zero value otherwise.

### GetOwnerNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetOwnerNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetOwnerNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetOwnerNumericOk returns a tuple with the OwnerNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetOwnerNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetOwnerNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetOwnerNumeric sets OwnerNumeric field to given value.

### HasOwnerNumeric

`func (o *VolumesResourceInner) HasOwnerNumeric() bool`

HasOwnerNumeric returns a boolean if a field has been set.

### GetPiFormat

`func (o *VolumesResourceInner) GetPiFormat() string`

GetPiFormat returns the PiFormat field if non-nil, zero value otherwise.

### GetPiFormatOk

`func (o *VolumesResourceInner) GetPiFormatOk() (*string, bool)`

GetPiFormatOk returns a tuple with the PiFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiFormat

`func (o *VolumesResourceInner) SetPiFormat(v string)`

SetPiFormat sets PiFormat field to given value.

### HasPiFormat

`func (o *VolumesResourceInner) HasPiFormat() bool`

HasPiFormat returns a boolean if a field has been set.

### GetPiFormatNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetPiFormatNumeric() int64`
=======
`func (o *VolumesResourceInner) GetPiFormatNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPiFormatNumeric returns the PiFormatNumeric field if non-nil, zero value otherwise.

### GetPiFormatNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetPiFormatNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetPiFormatNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPiFormatNumericOk returns a tuple with the PiFormatNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiFormatNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetPiFormatNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetPiFormatNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetPiFormatNumeric sets PiFormatNumeric field to given value.

### HasPiFormatNumeric

`func (o *VolumesResourceInner) HasPiFormatNumeric() bool`

HasPiFormatNumeric returns a boolean if a field has been set.

### GetPreferredOwner

`func (o *VolumesResourceInner) GetPreferredOwner() string`

GetPreferredOwner returns the PreferredOwner field if non-nil, zero value otherwise.

### GetPreferredOwnerOk

`func (o *VolumesResourceInner) GetPreferredOwnerOk() (*string, bool)`

GetPreferredOwnerOk returns a tuple with the PreferredOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredOwner

`func (o *VolumesResourceInner) SetPreferredOwner(v string)`

SetPreferredOwner sets PreferredOwner field to given value.

### HasPreferredOwner

`func (o *VolumesResourceInner) HasPreferredOwner() bool`

HasPreferredOwner returns a boolean if a field has been set.

### GetPreferredOwnerNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetPreferredOwnerNumeric() int64`
=======
`func (o *VolumesResourceInner) GetPreferredOwnerNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPreferredOwnerNumeric returns the PreferredOwnerNumeric field if non-nil, zero value otherwise.

### GetPreferredOwnerNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetPreferredOwnerNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetPreferredOwnerNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPreferredOwnerNumericOk returns a tuple with the PreferredOwnerNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredOwnerNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetPreferredOwnerNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetPreferredOwnerNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetPreferredOwnerNumeric sets PreferredOwnerNumeric field to given value.

### HasPreferredOwnerNumeric

`func (o *VolumesResourceInner) HasPreferredOwnerNumeric() bool`

HasPreferredOwnerNumeric returns a boolean if a field has been set.

### GetProgress

`func (o *VolumesResourceInner) GetProgress() string`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *VolumesResourceInner) GetProgressOk() (*string, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *VolumesResourceInner) SetProgress(v string)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *VolumesResourceInner) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProgressNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetProgressNumeric() int64`
=======
`func (o *VolumesResourceInner) GetProgressNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetProgressNumeric returns the ProgressNumeric field if non-nil, zero value otherwise.

### GetProgressNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetProgressNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetProgressNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetProgressNumericOk returns a tuple with the ProgressNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetProgressNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetProgressNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetProgressNumeric sets ProgressNumeric field to given value.

### HasProgressNumeric

`func (o *VolumesResourceInner) HasProgressNumeric() bool`

HasProgressNumeric returns a boolean if a field has been set.

### GetRaidtype

`func (o *VolumesResourceInner) GetRaidtype() string`

GetRaidtype returns the Raidtype field if non-nil, zero value otherwise.

### GetRaidtypeOk

`func (o *VolumesResourceInner) GetRaidtypeOk() (*string, bool)`

GetRaidtypeOk returns a tuple with the Raidtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidtype

`func (o *VolumesResourceInner) SetRaidtype(v string)`

SetRaidtype sets Raidtype field to given value.

### HasRaidtype

`func (o *VolumesResourceInner) HasRaidtype() bool`

HasRaidtype returns a boolean if a field has been set.

### GetRaidtypeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetRaidtypeNumeric() int64`
=======
`func (o *VolumesResourceInner) GetRaidtypeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRaidtypeNumeric returns the RaidtypeNumeric field if non-nil, zero value otherwise.

### GetRaidtypeNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetRaidtypeNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetRaidtypeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRaidtypeNumericOk returns a tuple with the RaidtypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidtypeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetRaidtypeNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetRaidtypeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetRaidtypeNumeric sets RaidtypeNumeric field to given value.

### HasRaidtypeNumeric

`func (o *VolumesResourceInner) HasRaidtypeNumeric() bool`

HasRaidtypeNumeric returns a boolean if a field has been set.

### GetReadAheadSize

`func (o *VolumesResourceInner) GetReadAheadSize() string`

GetReadAheadSize returns the ReadAheadSize field if non-nil, zero value otherwise.

### GetReadAheadSizeOk

`func (o *VolumesResourceInner) GetReadAheadSizeOk() (*string, bool)`

GetReadAheadSizeOk returns a tuple with the ReadAheadSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadSize

`func (o *VolumesResourceInner) SetReadAheadSize(v string)`

SetReadAheadSize sets ReadAheadSize field to given value.

### HasReadAheadSize

`func (o *VolumesResourceInner) HasReadAheadSize() bool`

HasReadAheadSize returns a boolean if a field has been set.

### GetReadAheadSizeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetReadAheadSizeNumeric() int64`
=======
`func (o *VolumesResourceInner) GetReadAheadSizeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetReadAheadSizeNumeric returns the ReadAheadSizeNumeric field if non-nil, zero value otherwise.

### GetReadAheadSizeNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetReadAheadSizeNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetReadAheadSizeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetReadAheadSizeNumericOk returns a tuple with the ReadAheadSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAheadSizeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetReadAheadSizeNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetReadAheadSizeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetReadAheadSizeNumeric sets ReadAheadSizeNumeric field to given value.

### HasReadAheadSizeNumeric

`func (o *VolumesResourceInner) HasReadAheadSizeNumeric() bool`

HasReadAheadSizeNumeric returns a boolean if a field has been set.

### GetReplicationSet

`func (o *VolumesResourceInner) GetReplicationSet() string`

GetReplicationSet returns the ReplicationSet field if non-nil, zero value otherwise.

### GetReplicationSetOk

`func (o *VolumesResourceInner) GetReplicationSetOk() (*string, bool)`

GetReplicationSetOk returns a tuple with the ReplicationSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSet

`func (o *VolumesResourceInner) SetReplicationSet(v string)`

SetReplicationSet sets ReplicationSet field to given value.

### HasReplicationSet

`func (o *VolumesResourceInner) HasReplicationSet() bool`

HasReplicationSet returns a boolean if a field has been set.

### GetReservedSizeInPages

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetReservedSizeInPages() int64`
=======
`func (o *VolumesResourceInner) GetReservedSizeInPages() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetReservedSizeInPages returns the ReservedSizeInPages field if non-nil, zero value otherwise.

### GetReservedSizeInPagesOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetReservedSizeInPagesOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetReservedSizeInPagesOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetReservedSizeInPagesOk returns a tuple with the ReservedSizeInPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedSizeInPages

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetReservedSizeInPages(v int64)`
=======
`func (o *VolumesResourceInner) SetReservedSizeInPages(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetReservedSizeInPages sets ReservedSizeInPages field to given value.

### HasReservedSizeInPages

`func (o *VolumesResourceInner) HasReservedSizeInPages() bool`

HasReservedSizeInPages returns a boolean if a field has been set.

### GetSerialNumber

`func (o *VolumesResourceInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *VolumesResourceInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *VolumesResourceInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *VolumesResourceInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSize

`func (o *VolumesResourceInner) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumesResourceInner) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumesResourceInner) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *VolumesResourceInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetSizeNumeric() int64`
=======
`func (o *VolumesResourceInner) GetSizeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSizeNumeric returns the SizeNumeric field if non-nil, zero value otherwise.

### GetSizeNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetSizeNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetSizeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSizeNumericOk returns a tuple with the SizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetSizeNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetSizeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSizeNumeric sets SizeNumeric field to given value.

### HasSizeNumeric

`func (o *VolumesResourceInner) HasSizeNumeric() bool`

HasSizeNumeric returns a boolean if a field has been set.

### GetSnapPool

`func (o *VolumesResourceInner) GetSnapPool() string`

GetSnapPool returns the SnapPool field if non-nil, zero value otherwise.

### GetSnapPoolOk

`func (o *VolumesResourceInner) GetSnapPoolOk() (*string, bool)`

GetSnapPoolOk returns a tuple with the SnapPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapPool

`func (o *VolumesResourceInner) SetSnapPool(v string)`

SetSnapPool sets SnapPool field to given value.

### HasSnapPool

`func (o *VolumesResourceInner) HasSnapPool() bool`

HasSnapPool returns a boolean if a field has been set.

### GetSnapshot

`func (o *VolumesResourceInner) GetSnapshot() string`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *VolumesResourceInner) GetSnapshotOk() (*string, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *VolumesResourceInner) SetSnapshot(v string)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *VolumesResourceInner) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetSnapshotRetentionPriority

`func (o *VolumesResourceInner) GetSnapshotRetentionPriority() string`

GetSnapshotRetentionPriority returns the SnapshotRetentionPriority field if non-nil, zero value otherwise.

### GetSnapshotRetentionPriorityOk

`func (o *VolumesResourceInner) GetSnapshotRetentionPriorityOk() (*string, bool)`

GetSnapshotRetentionPriorityOk returns a tuple with the SnapshotRetentionPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetentionPriority

`func (o *VolumesResourceInner) SetSnapshotRetentionPriority(v string)`

SetSnapshotRetentionPriority sets SnapshotRetentionPriority field to given value.

### HasSnapshotRetentionPriority

`func (o *VolumesResourceInner) HasSnapshotRetentionPriority() bool`

HasSnapshotRetentionPriority returns a boolean if a field has been set.

### GetSnapshotRetentionPriorityNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetSnapshotRetentionPriorityNumeric() int64`
=======
`func (o *VolumesResourceInner) GetSnapshotRetentionPriorityNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSnapshotRetentionPriorityNumeric returns the SnapshotRetentionPriorityNumeric field if non-nil, zero value otherwise.

### GetSnapshotRetentionPriorityNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetSnapshotRetentionPriorityNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetSnapshotRetentionPriorityNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSnapshotRetentionPriorityNumericOk returns a tuple with the SnapshotRetentionPriorityNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetentionPriorityNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetSnapshotRetentionPriorityNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetSnapshotRetentionPriorityNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSnapshotRetentionPriorityNumeric sets SnapshotRetentionPriorityNumeric field to given value.

### HasSnapshotRetentionPriorityNumeric

`func (o *VolumesResourceInner) HasSnapshotRetentionPriorityNumeric() bool`

HasSnapshotRetentionPriorityNumeric returns a boolean if a field has been set.

### GetStoragePoolName

`func (o *VolumesResourceInner) GetStoragePoolName() string`

GetStoragePoolName returns the StoragePoolName field if non-nil, zero value otherwise.

### GetStoragePoolNameOk

`func (o *VolumesResourceInner) GetStoragePoolNameOk() (*string, bool)`

GetStoragePoolNameOk returns a tuple with the StoragePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolName

`func (o *VolumesResourceInner) SetStoragePoolName(v string)`

SetStoragePoolName sets StoragePoolName field to given value.

### HasStoragePoolName

`func (o *VolumesResourceInner) HasStoragePoolName() bool`

HasStoragePoolName returns a boolean if a field has been set.

### GetStoragePoolsUrl

`func (o *VolumesResourceInner) GetStoragePoolsUrl() string`

GetStoragePoolsUrl returns the StoragePoolsUrl field if non-nil, zero value otherwise.

### GetStoragePoolsUrlOk

`func (o *VolumesResourceInner) GetStoragePoolsUrlOk() (*string, bool)`

GetStoragePoolsUrlOk returns a tuple with the StoragePoolsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolsUrl

`func (o *VolumesResourceInner) SetStoragePoolsUrl(v string)`

SetStoragePoolsUrl sets StoragePoolsUrl field to given value.

### HasStoragePoolsUrl

`func (o *VolumesResourceInner) HasStoragePoolsUrl() bool`

HasStoragePoolsUrl returns a boolean if a field has been set.

### GetStorageType

`func (o *VolumesResourceInner) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *VolumesResourceInner) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *VolumesResourceInner) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *VolumesResourceInner) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetStorageTypeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetStorageTypeNumeric() int64`
=======
`func (o *VolumesResourceInner) GetStorageTypeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetStorageTypeNumeric returns the StorageTypeNumeric field if non-nil, zero value otherwise.

### GetStorageTypeNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetStorageTypeNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetStorageTypeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetStorageTypeNumericOk returns a tuple with the StorageTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetStorageTypeNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetStorageTypeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetStorageTypeNumeric sets StorageTypeNumeric field to given value.

### HasStorageTypeNumeric

`func (o *VolumesResourceInner) HasStorageTypeNumeric() bool`

HasStorageTypeNumeric returns a boolean if a field has been set.

### GetThresholdPercentOfPool

`func (o *VolumesResourceInner) GetThresholdPercentOfPool() string`

GetThresholdPercentOfPool returns the ThresholdPercentOfPool field if non-nil, zero value otherwise.

### GetThresholdPercentOfPoolOk

`func (o *VolumesResourceInner) GetThresholdPercentOfPoolOk() (*string, bool)`

GetThresholdPercentOfPoolOk returns a tuple with the ThresholdPercentOfPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdPercentOfPool

`func (o *VolumesResourceInner) SetThresholdPercentOfPool(v string)`

SetThresholdPercentOfPool sets ThresholdPercentOfPool field to given value.

### HasThresholdPercentOfPool

`func (o *VolumesResourceInner) HasThresholdPercentOfPool() bool`

HasThresholdPercentOfPool returns a boolean if a field has been set.

### GetTierAffinity

`func (o *VolumesResourceInner) GetTierAffinity() string`

GetTierAffinity returns the TierAffinity field if non-nil, zero value otherwise.

### GetTierAffinityOk

`func (o *VolumesResourceInner) GetTierAffinityOk() (*string, bool)`

GetTierAffinityOk returns a tuple with the TierAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierAffinity

`func (o *VolumesResourceInner) SetTierAffinity(v string)`

SetTierAffinity sets TierAffinity field to given value.

### HasTierAffinity

`func (o *VolumesResourceInner) HasTierAffinity() bool`

HasTierAffinity returns a boolean if a field has been set.

### GetTierAffinityNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetTierAffinityNumeric() int64`
=======
`func (o *VolumesResourceInner) GetTierAffinityNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetTierAffinityNumeric returns the TierAffinityNumeric field if non-nil, zero value otherwise.

### GetTierAffinityNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetTierAffinityNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetTierAffinityNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetTierAffinityNumericOk returns a tuple with the TierAffinityNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierAffinityNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetTierAffinityNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetTierAffinityNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetTierAffinityNumeric sets TierAffinityNumeric field to given value.

### HasTierAffinityNumeric

`func (o *VolumesResourceInner) HasTierAffinityNumeric() bool`

HasTierAffinityNumeric returns a boolean if a field has been set.

### GetTotalSize

`func (o *VolumesResourceInner) GetTotalSize() string`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *VolumesResourceInner) GetTotalSizeOk() (*string, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *VolumesResourceInner) SetTotalSize(v string)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *VolumesResourceInner) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetTotalSizeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetTotalSizeNumeric() int64`
=======
`func (o *VolumesResourceInner) GetTotalSizeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetTotalSizeNumeric returns the TotalSizeNumeric field if non-nil, zero value otherwise.

### GetTotalSizeNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetTotalSizeNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetTotalSizeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetTotalSizeNumericOk returns a tuple with the TotalSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetTotalSizeNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetTotalSizeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetTotalSizeNumeric sets TotalSizeNumeric field to given value.

### HasTotalSizeNumeric

`func (o *VolumesResourceInner) HasTotalSizeNumeric() bool`

HasTotalSizeNumeric returns a boolean if a field has been set.

### GetUrl

`func (o *VolumesResourceInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VolumesResourceInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VolumesResourceInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VolumesResourceInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVirtualDiskName

`func (o *VolumesResourceInner) GetVirtualDiskName() string`

GetVirtualDiskName returns the VirtualDiskName field if non-nil, zero value otherwise.

### GetVirtualDiskNameOk

`func (o *VolumesResourceInner) GetVirtualDiskNameOk() (*string, bool)`

GetVirtualDiskNameOk returns a tuple with the VirtualDiskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskName

`func (o *VolumesResourceInner) SetVirtualDiskName(v string)`

SetVirtualDiskName sets VirtualDiskName field to given value.

### HasVirtualDiskName

`func (o *VolumesResourceInner) HasVirtualDiskName() bool`

HasVirtualDiskName returns a boolean if a field has been set.

### GetVirtualDiskSerial

`func (o *VolumesResourceInner) GetVirtualDiskSerial() string`

GetVirtualDiskSerial returns the VirtualDiskSerial field if non-nil, zero value otherwise.

### GetVirtualDiskSerialOk

`func (o *VolumesResourceInner) GetVirtualDiskSerialOk() (*string, bool)`

GetVirtualDiskSerialOk returns a tuple with the VirtualDiskSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskSerial

`func (o *VolumesResourceInner) SetVirtualDiskSerial(v string)`

SetVirtualDiskSerial sets VirtualDiskSerial field to given value.

### HasVirtualDiskSerial

`func (o *VolumesResourceInner) HasVirtualDiskSerial() bool`

HasVirtualDiskSerial returns a boolean if a field has been set.

### GetVolumeClass

`func (o *VolumesResourceInner) GetVolumeClass() string`

GetVolumeClass returns the VolumeClass field if non-nil, zero value otherwise.

### GetVolumeClassOk

`func (o *VolumesResourceInner) GetVolumeClassOk() (*string, bool)`

GetVolumeClassOk returns a tuple with the VolumeClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeClass

`func (o *VolumesResourceInner) SetVolumeClass(v string)`

SetVolumeClass sets VolumeClass field to given value.

### HasVolumeClass

`func (o *VolumesResourceInner) HasVolumeClass() bool`

HasVolumeClass returns a boolean if a field has been set.

### GetVolumeClassNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetVolumeClassNumeric() int64`
=======
`func (o *VolumesResourceInner) GetVolumeClassNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetVolumeClassNumeric returns the VolumeClassNumeric field if non-nil, zero value otherwise.

### GetVolumeClassNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetVolumeClassNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetVolumeClassNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetVolumeClassNumericOk returns a tuple with the VolumeClassNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeClassNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetVolumeClassNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetVolumeClassNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetVolumeClassNumeric sets VolumeClassNumeric field to given value.

### HasVolumeClassNumeric

`func (o *VolumesResourceInner) HasVolumeClassNumeric() bool`

HasVolumeClassNumeric returns a boolean if a field has been set.

### GetVolumeDescription

`func (o *VolumesResourceInner) GetVolumeDescription() string`

GetVolumeDescription returns the VolumeDescription field if non-nil, zero value otherwise.

### GetVolumeDescriptionOk

`func (o *VolumesResourceInner) GetVolumeDescriptionOk() (*string, bool)`

GetVolumeDescriptionOk returns a tuple with the VolumeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeDescription

`func (o *VolumesResourceInner) SetVolumeDescription(v string)`

SetVolumeDescription sets VolumeDescription field to given value.

### HasVolumeDescription

`func (o *VolumesResourceInner) HasVolumeDescription() bool`

HasVolumeDescription returns a boolean if a field has been set.

### GetVolumeGroup

`func (o *VolumesResourceInner) GetVolumeGroup() string`

GetVolumeGroup returns the VolumeGroup field if non-nil, zero value otherwise.

### GetVolumeGroupOk

`func (o *VolumesResourceInner) GetVolumeGroupOk() (*string, bool)`

GetVolumeGroupOk returns a tuple with the VolumeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroup

`func (o *VolumesResourceInner) SetVolumeGroup(v string)`

SetVolumeGroup sets VolumeGroup field to given value.

### HasVolumeGroup

`func (o *VolumesResourceInner) HasVolumeGroup() bool`

HasVolumeGroup returns a boolean if a field has been set.

### GetVolumeName

`func (o *VolumesResourceInner) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *VolumesResourceInner) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *VolumesResourceInner) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *VolumesResourceInner) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumeParent

`func (o *VolumesResourceInner) GetVolumeParent() string`

GetVolumeParent returns the VolumeParent field if non-nil, zero value otherwise.

### GetVolumeParentOk

`func (o *VolumesResourceInner) GetVolumeParentOk() (*string, bool)`

GetVolumeParentOk returns a tuple with the VolumeParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeParent

`func (o *VolumesResourceInner) SetVolumeParent(v string)`

SetVolumeParent sets VolumeParent field to given value.

### HasVolumeParent

`func (o *VolumesResourceInner) HasVolumeParent() bool`

HasVolumeParent returns a boolean if a field has been set.

### GetVolumeQualifier

`func (o *VolumesResourceInner) GetVolumeQualifier() string`

GetVolumeQualifier returns the VolumeQualifier field if non-nil, zero value otherwise.

### GetVolumeQualifierOk

`func (o *VolumesResourceInner) GetVolumeQualifierOk() (*string, bool)`

GetVolumeQualifierOk returns a tuple with the VolumeQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeQualifier

`func (o *VolumesResourceInner) SetVolumeQualifier(v string)`

SetVolumeQualifier sets VolumeQualifier field to given value.

### HasVolumeQualifier

`func (o *VolumesResourceInner) HasVolumeQualifier() bool`

HasVolumeQualifier returns a boolean if a field has been set.

### GetVolumeQualifierNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetVolumeQualifierNumeric() int64`
=======
`func (o *VolumesResourceInner) GetVolumeQualifierNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetVolumeQualifierNumeric returns the VolumeQualifierNumeric field if non-nil, zero value otherwise.

### GetVolumeQualifierNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetVolumeQualifierNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetVolumeQualifierNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetVolumeQualifierNumericOk returns a tuple with the VolumeQualifierNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeQualifierNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetVolumeQualifierNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetVolumeQualifierNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetVolumeQualifierNumeric sets VolumeQualifierNumeric field to given value.

### HasVolumeQualifierNumeric

`func (o *VolumesResourceInner) HasVolumeQualifierNumeric() bool`

HasVolumeQualifierNumeric returns a boolean if a field has been set.

### GetVolumeType

`func (o *VolumesResourceInner) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *VolumesResourceInner) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *VolumesResourceInner) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *VolumesResourceInner) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetVolumeTypeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetVolumeTypeNumeric() int64`
=======
`func (o *VolumesResourceInner) GetVolumeTypeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetVolumeTypeNumeric returns the VolumeTypeNumeric field if non-nil, zero value otherwise.

### GetVolumeTypeNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetVolumeTypeNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetVolumeTypeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetVolumeTypeNumericOk returns a tuple with the VolumeTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTypeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetVolumeTypeNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetVolumeTypeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetVolumeTypeNumeric sets VolumeTypeNumeric field to given value.

### HasVolumeTypeNumeric

`func (o *VolumesResourceInner) HasVolumeTypeNumeric() bool`

HasVolumeTypeNumeric returns a boolean if a field has been set.

### GetVolumeUsageType

`func (o *VolumesResourceInner) GetVolumeUsageType() string`

GetVolumeUsageType returns the VolumeUsageType field if non-nil, zero value otherwise.

### GetVolumeUsageTypeOk

`func (o *VolumesResourceInner) GetVolumeUsageTypeOk() (*string, bool)`

GetVolumeUsageTypeOk returns a tuple with the VolumeUsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUsageType

`func (o *VolumesResourceInner) SetVolumeUsageType(v string)`

SetVolumeUsageType sets VolumeUsageType field to given value.

### HasVolumeUsageType

`func (o *VolumesResourceInner) HasVolumeUsageType() bool`

HasVolumeUsageType returns a boolean if a field has been set.

### GetVolumeUsageTypeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetVolumeUsageTypeNumeric() int64`
=======
`func (o *VolumesResourceInner) GetVolumeUsageTypeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetVolumeUsageTypeNumeric returns the VolumeUsageTypeNumeric field if non-nil, zero value otherwise.

### GetVolumeUsageTypeNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetVolumeUsageTypeNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetVolumeUsageTypeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetVolumeUsageTypeNumericOk returns a tuple with the VolumeUsageTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUsageTypeNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetVolumeUsageTypeNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetVolumeUsageTypeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetVolumeUsageTypeNumeric sets VolumeUsageTypeNumeric field to given value.

### HasVolumeUsageTypeNumeric

`func (o *VolumesResourceInner) HasVolumeUsageTypeNumeric() bool`

HasVolumeUsageTypeNumeric returns a boolean if a field has been set.

### GetWritePolicy

`func (o *VolumesResourceInner) GetWritePolicy() string`

GetWritePolicy returns the WritePolicy field if non-nil, zero value otherwise.

### GetWritePolicyOk

`func (o *VolumesResourceInner) GetWritePolicyOk() (*string, bool)`

GetWritePolicyOk returns a tuple with the WritePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritePolicy

`func (o *VolumesResourceInner) SetWritePolicy(v string)`

SetWritePolicy sets WritePolicy field to given value.

### HasWritePolicy

`func (o *VolumesResourceInner) HasWritePolicy() bool`

HasWritePolicy returns a boolean if a field has been set.

### GetWritePolicyNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetWritePolicyNumeric() int64`
=======
`func (o *VolumesResourceInner) GetWritePolicyNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetWritePolicyNumeric returns the WritePolicyNumeric field if non-nil, zero value otherwise.

### GetWritePolicyNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetWritePolicyNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetWritePolicyNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetWritePolicyNumericOk returns a tuple with the WritePolicyNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritePolicyNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetWritePolicyNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetWritePolicyNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetWritePolicyNumeric sets WritePolicyNumeric field to given value.

### HasWritePolicyNumeric

`func (o *VolumesResourceInner) HasWritePolicyNumeric() bool`

HasWritePolicyNumeric returns a boolean if a field has been set.

### GetWwn

`func (o *VolumesResourceInner) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *VolumesResourceInner) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *VolumesResourceInner) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *VolumesResourceInner) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### GetZeroInitPageOnAllocation

`func (o *VolumesResourceInner) GetZeroInitPageOnAllocation() string`

GetZeroInitPageOnAllocation returns the ZeroInitPageOnAllocation field if non-nil, zero value otherwise.

### GetZeroInitPageOnAllocationOk

`func (o *VolumesResourceInner) GetZeroInitPageOnAllocationOk() (*string, bool)`

GetZeroInitPageOnAllocationOk returns a tuple with the ZeroInitPageOnAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroInitPageOnAllocation

`func (o *VolumesResourceInner) SetZeroInitPageOnAllocation(v string)`

SetZeroInitPageOnAllocation sets ZeroInitPageOnAllocation field to given value.

### HasZeroInitPageOnAllocation

`func (o *VolumesResourceInner) HasZeroInitPageOnAllocation() bool`

HasZeroInitPageOnAllocation returns a boolean if a field has been set.

### GetZeroInitPageOnAllocationNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetZeroInitPageOnAllocationNumeric() int64`
=======
`func (o *VolumesResourceInner) GetZeroInitPageOnAllocationNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetZeroInitPageOnAllocationNumeric returns the ZeroInitPageOnAllocationNumeric field if non-nil, zero value otherwise.

### GetZeroInitPageOnAllocationNumericOk

<<<<<<< HEAD
`func (o *VolumesResourceInner) GetZeroInitPageOnAllocationNumericOk() (*int64, bool)`
=======
`func (o *VolumesResourceInner) GetZeroInitPageOnAllocationNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetZeroInitPageOnAllocationNumericOk returns a tuple with the ZeroInitPageOnAllocationNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroInitPageOnAllocationNumeric

<<<<<<< HEAD
`func (o *VolumesResourceInner) SetZeroInitPageOnAllocationNumeric(v int64)`
=======
`func (o *VolumesResourceInner) SetZeroInitPageOnAllocationNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetZeroInitPageOnAllocationNumeric sets ZeroInitPageOnAllocationNumeric field to given value.

### HasZeroInitPageOnAllocationNumeric

`func (o *VolumesResourceInner) HasZeroInitPageOnAllocationNumeric() bool`

HasZeroInitPageOnAllocationNumeric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


