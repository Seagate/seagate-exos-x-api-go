# PoolsResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**AllocatedPages** | Pointer to **int32** | Number of pages allocated | [optional] 
**AvailablePages** | Pointer to **int32** | Available pages | [optional] 
**AvailableRfcSize** | Pointer to **string** |  | [optional] 
**AvailableRfcSizeNumeric** | Pointer to **int32** |  | [optional] 
**Blocksize** | Pointer to **int32** |  | [optional] 
**CompressionEfficiency** | Pointer to **string** |  | [optional] 
**DiskGroupsCount** | Pointer to **int32** |  | [optional] 
**ExtendedStatus** | Pointer to **int32** | Extended status (bits) | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**HealthNumeric** | Pointer to **int32** |  | [optional] 
**HealthReason** | Pointer to **string** |  | [optional] 
**HealthReasonNumeric** | Pointer to **int32** |  | [optional] 
**HealthRecommendation** | Pointer to **string** |  | [optional] 
**HealthRecommendationNumeric** | Pointer to **int32** |  | [optional] 
**HighThreshold** | Pointer to **string** | High threshold percentage | [optional] 
**IdlePageCheck** | Pointer to **string** |  | [optional] 
**IdlePageCheckNumeric** | Pointer to **int32** |  | [optional] 
**LowThreshold** | Pointer to **string** | Low threshold percentage | [optional] 
**MetadataAllocated** | Pointer to **string** | Pool Metadata currently in use | [optional] 
**MetadataAllocatedNumeric** | Pointer to **int32** | Pool Metadata currently in use( In numeric form ) | [optional] 
**MetadataAvailable** | Pointer to **string** | Pool Metadata available capacity | [optional] 
**MetadataAvailableNumeric** | Pointer to **int32** | Pool Metadata available capacity( In numeric form ) | [optional] 
**MetadataTotalSize** | Pointer to **string** | Disk Group Metadata Total Size | [optional] 
**MetadataTotalSizeNumeric** | Pointer to **int32** | Disk Group Metadata Total Size( In numeric form ) | [optional] 
**MetadataVolSize** | Pointer to **string** | Size of the storage pool metadata volume | [optional] 
**MetadataVolSizeNumeric** | Pointer to **int32** | Size of the storage pool metadata volume( In numeric form ) | [optional] 
**MiddleThreshold** | Pointer to **string** | Middle threshold percentage | [optional] 
**Migration** | Pointer to **string** | Migration between tiers | [optional] 
**MigrationNumeric** | Pointer to **int32** | Migration between tiers( In numeric form ) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OverCommitted** | Pointer to **string** |  | [optional] 
**OverCommittedNumeric** | Pointer to **int32** |  | [optional] 
**Overcommit** | Pointer to **string** |  | [optional] 
**OvercommitNumeric** | Pointer to **int32** |  | [optional] 
**Owner** | Pointer to **string** | Controller owning the component | [optional] 
**OwnerNumeric** | Pointer to **int32** | Controller owning the component( In numeric form ) | [optional] 
**PageSize** | Pointer to **string** | Page size in blocks | [optional] 
**PageSizeNumeric** | Pointer to **int32** | Page size in blocks( In numeric form ) | [optional] 
**PoolSectorFormat** | Pointer to **string** | Pool Sector Format | [optional] 
**PoolSectorFormatNumeric** | Pointer to **int32** | Pool Sector Format( In numeric form ) | [optional] 
**PreferredOwner** | Pointer to **string** | Configured owner | [optional] 
**PreferredOwnerNumeric** | Pointer to **int32** | Configured owner( In numeric form ) | [optional] 
**ReadFlashCache** | Pointer to **string** |  | [optional] 
**ReadFlashCacheNumeric** | Pointer to **int32** |  | [optional] 
**Rebalance** | Pointer to **string** | Rebalance within the tier | [optional] 
**RebalanceNumeric** | Pointer to **int32** | Rebalance within the tier( In numeric form ) | [optional] 
**ReservedSize** | Pointer to **string** |  | [optional] 
**ReservedSizeNumeric** | Pointer to **int32** |  | [optional] 
**ReservedUnallocSize** | Pointer to **string** |  | [optional] 
**ReservedUnallocSizeNumeric** | Pointer to **int32** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**SnapSize** | Pointer to **string** |  | [optional] 
**SnapSizeNumeric** | Pointer to **int32** |  | [optional] 
**StorageType** | Pointer to **string** | Storage type | [optional] 
**StorageTypeNumeric** | Pointer to **int32** | Storage type( In numeric form ) | [optional] 
**TotalAvail** | Pointer to **string** |  | [optional] 
**TotalAvailNumeric** | Pointer to **int32** |  | [optional] 
**TotalRfcSize** | Pointer to **string** |  | [optional] 
**TotalRfcSizeNumeric** | Pointer to **int32** |  | [optional] 
**TotalSize** | Pointer to **string** | The total size formatted using the session settings for base, precision, and units | [optional] 
**TotalSizeNumeric** | Pointer to **int32** | The total size formatted using the session settings for base, precision, and units( In numeric form ) | [optional] 
**Url** | Pointer to **string** | The resource URL | [optional] 
**UtilityRunning** | Pointer to **string** |  | [optional] 
**UtilityRunningNumeric** | Pointer to **int32** |  | [optional] 
**Volumes** | Pointer to **int32** |  | [optional] 
**ZeroScan** | Pointer to **string** | Scan for zero-value pages | [optional] 
**ZeroScanNumeric** | Pointer to **int32** | Scan for zero-value pages( In numeric form ) | [optional] 
**DiskGroups** | Pointer to [**[]DiskGroupsResourceInner**](DiskGroupsResourceInner.md) |  | [optional] 
**Tiers** | Pointer to [**[]TiersResourceInner**](TiersResourceInner.md) |  | [optional] 

## Methods

### NewPoolsResourceInner

`func NewPoolsResourceInner() *PoolsResourceInner`

NewPoolsResourceInner instantiates a new PoolsResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolsResourceInnerWithDefaults

`func NewPoolsResourceInnerWithDefaults() *PoolsResourceInner`

NewPoolsResourceInnerWithDefaults instantiates a new PoolsResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *PoolsResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *PoolsResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *PoolsResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *PoolsResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *PoolsResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PoolsResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PoolsResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PoolsResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetAllocatedPages

`func (o *PoolsResourceInner) GetAllocatedPages() int32`

GetAllocatedPages returns the AllocatedPages field if non-nil, zero value otherwise.

### GetAllocatedPagesOk

`func (o *PoolsResourceInner) GetAllocatedPagesOk() (*int32, bool)`

GetAllocatedPagesOk returns a tuple with the AllocatedPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedPages

`func (o *PoolsResourceInner) SetAllocatedPages(v int32)`

SetAllocatedPages sets AllocatedPages field to given value.

### HasAllocatedPages

`func (o *PoolsResourceInner) HasAllocatedPages() bool`

HasAllocatedPages returns a boolean if a field has been set.

### GetAvailablePages

`func (o *PoolsResourceInner) GetAvailablePages() int32`

GetAvailablePages returns the AvailablePages field if non-nil, zero value otherwise.

### GetAvailablePagesOk

`func (o *PoolsResourceInner) GetAvailablePagesOk() (*int32, bool)`

GetAvailablePagesOk returns a tuple with the AvailablePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePages

`func (o *PoolsResourceInner) SetAvailablePages(v int32)`

SetAvailablePages sets AvailablePages field to given value.

### HasAvailablePages

`func (o *PoolsResourceInner) HasAvailablePages() bool`

HasAvailablePages returns a boolean if a field has been set.

### GetAvailableRfcSize

`func (o *PoolsResourceInner) GetAvailableRfcSize() string`

GetAvailableRfcSize returns the AvailableRfcSize field if non-nil, zero value otherwise.

### GetAvailableRfcSizeOk

`func (o *PoolsResourceInner) GetAvailableRfcSizeOk() (*string, bool)`

GetAvailableRfcSizeOk returns a tuple with the AvailableRfcSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableRfcSize

`func (o *PoolsResourceInner) SetAvailableRfcSize(v string)`

SetAvailableRfcSize sets AvailableRfcSize field to given value.

### HasAvailableRfcSize

`func (o *PoolsResourceInner) HasAvailableRfcSize() bool`

HasAvailableRfcSize returns a boolean if a field has been set.

### GetAvailableRfcSizeNumeric

`func (o *PoolsResourceInner) GetAvailableRfcSizeNumeric() int32`

GetAvailableRfcSizeNumeric returns the AvailableRfcSizeNumeric field if non-nil, zero value otherwise.

### GetAvailableRfcSizeNumericOk

`func (o *PoolsResourceInner) GetAvailableRfcSizeNumericOk() (*int32, bool)`

GetAvailableRfcSizeNumericOk returns a tuple with the AvailableRfcSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableRfcSizeNumeric

`func (o *PoolsResourceInner) SetAvailableRfcSizeNumeric(v int32)`

SetAvailableRfcSizeNumeric sets AvailableRfcSizeNumeric field to given value.

### HasAvailableRfcSizeNumeric

`func (o *PoolsResourceInner) HasAvailableRfcSizeNumeric() bool`

HasAvailableRfcSizeNumeric returns a boolean if a field has been set.

### GetBlocksize

`func (o *PoolsResourceInner) GetBlocksize() int32`

GetBlocksize returns the Blocksize field if non-nil, zero value otherwise.

### GetBlocksizeOk

`func (o *PoolsResourceInner) GetBlocksizeOk() (*int32, bool)`

GetBlocksizeOk returns a tuple with the Blocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocksize

`func (o *PoolsResourceInner) SetBlocksize(v int32)`

SetBlocksize sets Blocksize field to given value.

### HasBlocksize

`func (o *PoolsResourceInner) HasBlocksize() bool`

HasBlocksize returns a boolean if a field has been set.

### GetCompressionEfficiency

`func (o *PoolsResourceInner) GetCompressionEfficiency() string`

GetCompressionEfficiency returns the CompressionEfficiency field if non-nil, zero value otherwise.

### GetCompressionEfficiencyOk

`func (o *PoolsResourceInner) GetCompressionEfficiencyOk() (*string, bool)`

GetCompressionEfficiencyOk returns a tuple with the CompressionEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionEfficiency

`func (o *PoolsResourceInner) SetCompressionEfficiency(v string)`

SetCompressionEfficiency sets CompressionEfficiency field to given value.

### HasCompressionEfficiency

`func (o *PoolsResourceInner) HasCompressionEfficiency() bool`

HasCompressionEfficiency returns a boolean if a field has been set.

### GetDiskGroupsCount

`func (o *PoolsResourceInner) GetDiskGroupsCount() int32`

GetDiskGroupsCount returns the DiskGroupsCount field if non-nil, zero value otherwise.

### GetDiskGroupsCountOk

`func (o *PoolsResourceInner) GetDiskGroupsCountOk() (*int32, bool)`

GetDiskGroupsCountOk returns a tuple with the DiskGroupsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroupsCount

`func (o *PoolsResourceInner) SetDiskGroupsCount(v int32)`

SetDiskGroupsCount sets DiskGroupsCount field to given value.

### HasDiskGroupsCount

`func (o *PoolsResourceInner) HasDiskGroupsCount() bool`

HasDiskGroupsCount returns a boolean if a field has been set.

### GetExtendedStatus

`func (o *PoolsResourceInner) GetExtendedStatus() int32`

GetExtendedStatus returns the ExtendedStatus field if non-nil, zero value otherwise.

### GetExtendedStatusOk

`func (o *PoolsResourceInner) GetExtendedStatusOk() (*int32, bool)`

GetExtendedStatusOk returns a tuple with the ExtendedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedStatus

`func (o *PoolsResourceInner) SetExtendedStatus(v int32)`

SetExtendedStatus sets ExtendedStatus field to given value.

### HasExtendedStatus

`func (o *PoolsResourceInner) HasExtendedStatus() bool`

HasExtendedStatus returns a boolean if a field has been set.

### GetHealth

`func (o *PoolsResourceInner) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *PoolsResourceInner) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *PoolsResourceInner) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *PoolsResourceInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHealthNumeric

`func (o *PoolsResourceInner) GetHealthNumeric() int32`

GetHealthNumeric returns the HealthNumeric field if non-nil, zero value otherwise.

### GetHealthNumericOk

`func (o *PoolsResourceInner) GetHealthNumericOk() (*int32, bool)`

GetHealthNumericOk returns a tuple with the HealthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthNumeric

`func (o *PoolsResourceInner) SetHealthNumeric(v int32)`

SetHealthNumeric sets HealthNumeric field to given value.

### HasHealthNumeric

`func (o *PoolsResourceInner) HasHealthNumeric() bool`

HasHealthNumeric returns a boolean if a field has been set.

### GetHealthReason

`func (o *PoolsResourceInner) GetHealthReason() string`

GetHealthReason returns the HealthReason field if non-nil, zero value otherwise.

### GetHealthReasonOk

`func (o *PoolsResourceInner) GetHealthReasonOk() (*string, bool)`

GetHealthReasonOk returns a tuple with the HealthReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReason

`func (o *PoolsResourceInner) SetHealthReason(v string)`

SetHealthReason sets HealthReason field to given value.

### HasHealthReason

`func (o *PoolsResourceInner) HasHealthReason() bool`

HasHealthReason returns a boolean if a field has been set.

### GetHealthReasonNumeric

`func (o *PoolsResourceInner) GetHealthReasonNumeric() int32`

GetHealthReasonNumeric returns the HealthReasonNumeric field if non-nil, zero value otherwise.

### GetHealthReasonNumericOk

`func (o *PoolsResourceInner) GetHealthReasonNumericOk() (*int32, bool)`

GetHealthReasonNumericOk returns a tuple with the HealthReasonNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReasonNumeric

`func (o *PoolsResourceInner) SetHealthReasonNumeric(v int32)`

SetHealthReasonNumeric sets HealthReasonNumeric field to given value.

### HasHealthReasonNumeric

`func (o *PoolsResourceInner) HasHealthReasonNumeric() bool`

HasHealthReasonNumeric returns a boolean if a field has been set.

### GetHealthRecommendation

`func (o *PoolsResourceInner) GetHealthRecommendation() string`

GetHealthRecommendation returns the HealthRecommendation field if non-nil, zero value otherwise.

### GetHealthRecommendationOk

`func (o *PoolsResourceInner) GetHealthRecommendationOk() (*string, bool)`

GetHealthRecommendationOk returns a tuple with the HealthRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecommendation

`func (o *PoolsResourceInner) SetHealthRecommendation(v string)`

SetHealthRecommendation sets HealthRecommendation field to given value.

### HasHealthRecommendation

`func (o *PoolsResourceInner) HasHealthRecommendation() bool`

HasHealthRecommendation returns a boolean if a field has been set.

### GetHealthRecommendationNumeric

`func (o *PoolsResourceInner) GetHealthRecommendationNumeric() int32`

GetHealthRecommendationNumeric returns the HealthRecommendationNumeric field if non-nil, zero value otherwise.

### GetHealthRecommendationNumericOk

`func (o *PoolsResourceInner) GetHealthRecommendationNumericOk() (*int32, bool)`

GetHealthRecommendationNumericOk returns a tuple with the HealthRecommendationNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecommendationNumeric

`func (o *PoolsResourceInner) SetHealthRecommendationNumeric(v int32)`

SetHealthRecommendationNumeric sets HealthRecommendationNumeric field to given value.

### HasHealthRecommendationNumeric

`func (o *PoolsResourceInner) HasHealthRecommendationNumeric() bool`

HasHealthRecommendationNumeric returns a boolean if a field has been set.

### GetHighThreshold

`func (o *PoolsResourceInner) GetHighThreshold() string`

GetHighThreshold returns the HighThreshold field if non-nil, zero value otherwise.

### GetHighThresholdOk

`func (o *PoolsResourceInner) GetHighThresholdOk() (*string, bool)`

GetHighThresholdOk returns a tuple with the HighThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighThreshold

`func (o *PoolsResourceInner) SetHighThreshold(v string)`

SetHighThreshold sets HighThreshold field to given value.

### HasHighThreshold

`func (o *PoolsResourceInner) HasHighThreshold() bool`

HasHighThreshold returns a boolean if a field has been set.

### GetIdlePageCheck

`func (o *PoolsResourceInner) GetIdlePageCheck() string`

GetIdlePageCheck returns the IdlePageCheck field if non-nil, zero value otherwise.

### GetIdlePageCheckOk

`func (o *PoolsResourceInner) GetIdlePageCheckOk() (*string, bool)`

GetIdlePageCheckOk returns a tuple with the IdlePageCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdlePageCheck

`func (o *PoolsResourceInner) SetIdlePageCheck(v string)`

SetIdlePageCheck sets IdlePageCheck field to given value.

### HasIdlePageCheck

`func (o *PoolsResourceInner) HasIdlePageCheck() bool`

HasIdlePageCheck returns a boolean if a field has been set.

### GetIdlePageCheckNumeric

`func (o *PoolsResourceInner) GetIdlePageCheckNumeric() int32`

GetIdlePageCheckNumeric returns the IdlePageCheckNumeric field if non-nil, zero value otherwise.

### GetIdlePageCheckNumericOk

`func (o *PoolsResourceInner) GetIdlePageCheckNumericOk() (*int32, bool)`

GetIdlePageCheckNumericOk returns a tuple with the IdlePageCheckNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdlePageCheckNumeric

`func (o *PoolsResourceInner) SetIdlePageCheckNumeric(v int32)`

SetIdlePageCheckNumeric sets IdlePageCheckNumeric field to given value.

### HasIdlePageCheckNumeric

`func (o *PoolsResourceInner) HasIdlePageCheckNumeric() bool`

HasIdlePageCheckNumeric returns a boolean if a field has been set.

### GetLowThreshold

`func (o *PoolsResourceInner) GetLowThreshold() string`

GetLowThreshold returns the LowThreshold field if non-nil, zero value otherwise.

### GetLowThresholdOk

`func (o *PoolsResourceInner) GetLowThresholdOk() (*string, bool)`

GetLowThresholdOk returns a tuple with the LowThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowThreshold

`func (o *PoolsResourceInner) SetLowThreshold(v string)`

SetLowThreshold sets LowThreshold field to given value.

### HasLowThreshold

`func (o *PoolsResourceInner) HasLowThreshold() bool`

HasLowThreshold returns a boolean if a field has been set.

### GetMetadataAllocated

`func (o *PoolsResourceInner) GetMetadataAllocated() string`

GetMetadataAllocated returns the MetadataAllocated field if non-nil, zero value otherwise.

### GetMetadataAllocatedOk

`func (o *PoolsResourceInner) GetMetadataAllocatedOk() (*string, bool)`

GetMetadataAllocatedOk returns a tuple with the MetadataAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataAllocated

`func (o *PoolsResourceInner) SetMetadataAllocated(v string)`

SetMetadataAllocated sets MetadataAllocated field to given value.

### HasMetadataAllocated

`func (o *PoolsResourceInner) HasMetadataAllocated() bool`

HasMetadataAllocated returns a boolean if a field has been set.

### GetMetadataAllocatedNumeric

`func (o *PoolsResourceInner) GetMetadataAllocatedNumeric() int32`

GetMetadataAllocatedNumeric returns the MetadataAllocatedNumeric field if non-nil, zero value otherwise.

### GetMetadataAllocatedNumericOk

`func (o *PoolsResourceInner) GetMetadataAllocatedNumericOk() (*int32, bool)`

GetMetadataAllocatedNumericOk returns a tuple with the MetadataAllocatedNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataAllocatedNumeric

`func (o *PoolsResourceInner) SetMetadataAllocatedNumeric(v int32)`

SetMetadataAllocatedNumeric sets MetadataAllocatedNumeric field to given value.

### HasMetadataAllocatedNumeric

`func (o *PoolsResourceInner) HasMetadataAllocatedNumeric() bool`

HasMetadataAllocatedNumeric returns a boolean if a field has been set.

### GetMetadataAvailable

`func (o *PoolsResourceInner) GetMetadataAvailable() string`

GetMetadataAvailable returns the MetadataAvailable field if non-nil, zero value otherwise.

### GetMetadataAvailableOk

`func (o *PoolsResourceInner) GetMetadataAvailableOk() (*string, bool)`

GetMetadataAvailableOk returns a tuple with the MetadataAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataAvailable

`func (o *PoolsResourceInner) SetMetadataAvailable(v string)`

SetMetadataAvailable sets MetadataAvailable field to given value.

### HasMetadataAvailable

`func (o *PoolsResourceInner) HasMetadataAvailable() bool`

HasMetadataAvailable returns a boolean if a field has been set.

### GetMetadataAvailableNumeric

`func (o *PoolsResourceInner) GetMetadataAvailableNumeric() int32`

GetMetadataAvailableNumeric returns the MetadataAvailableNumeric field if non-nil, zero value otherwise.

### GetMetadataAvailableNumericOk

`func (o *PoolsResourceInner) GetMetadataAvailableNumericOk() (*int32, bool)`

GetMetadataAvailableNumericOk returns a tuple with the MetadataAvailableNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataAvailableNumeric

`func (o *PoolsResourceInner) SetMetadataAvailableNumeric(v int32)`

SetMetadataAvailableNumeric sets MetadataAvailableNumeric field to given value.

### HasMetadataAvailableNumeric

`func (o *PoolsResourceInner) HasMetadataAvailableNumeric() bool`

HasMetadataAvailableNumeric returns a boolean if a field has been set.

### GetMetadataTotalSize

`func (o *PoolsResourceInner) GetMetadataTotalSize() string`

GetMetadataTotalSize returns the MetadataTotalSize field if non-nil, zero value otherwise.

### GetMetadataTotalSizeOk

`func (o *PoolsResourceInner) GetMetadataTotalSizeOk() (*string, bool)`

GetMetadataTotalSizeOk returns a tuple with the MetadataTotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataTotalSize

`func (o *PoolsResourceInner) SetMetadataTotalSize(v string)`

SetMetadataTotalSize sets MetadataTotalSize field to given value.

### HasMetadataTotalSize

`func (o *PoolsResourceInner) HasMetadataTotalSize() bool`

HasMetadataTotalSize returns a boolean if a field has been set.

### GetMetadataTotalSizeNumeric

`func (o *PoolsResourceInner) GetMetadataTotalSizeNumeric() int32`

GetMetadataTotalSizeNumeric returns the MetadataTotalSizeNumeric field if non-nil, zero value otherwise.

### GetMetadataTotalSizeNumericOk

`func (o *PoolsResourceInner) GetMetadataTotalSizeNumericOk() (*int32, bool)`

GetMetadataTotalSizeNumericOk returns a tuple with the MetadataTotalSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataTotalSizeNumeric

`func (o *PoolsResourceInner) SetMetadataTotalSizeNumeric(v int32)`

SetMetadataTotalSizeNumeric sets MetadataTotalSizeNumeric field to given value.

### HasMetadataTotalSizeNumeric

`func (o *PoolsResourceInner) HasMetadataTotalSizeNumeric() bool`

HasMetadataTotalSizeNumeric returns a boolean if a field has been set.

### GetMetadataVolSize

`func (o *PoolsResourceInner) GetMetadataVolSize() string`

GetMetadataVolSize returns the MetadataVolSize field if non-nil, zero value otherwise.

### GetMetadataVolSizeOk

`func (o *PoolsResourceInner) GetMetadataVolSizeOk() (*string, bool)`

GetMetadataVolSizeOk returns a tuple with the MetadataVolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataVolSize

`func (o *PoolsResourceInner) SetMetadataVolSize(v string)`

SetMetadataVolSize sets MetadataVolSize field to given value.

### HasMetadataVolSize

`func (o *PoolsResourceInner) HasMetadataVolSize() bool`

HasMetadataVolSize returns a boolean if a field has been set.

### GetMetadataVolSizeNumeric

`func (o *PoolsResourceInner) GetMetadataVolSizeNumeric() int32`

GetMetadataVolSizeNumeric returns the MetadataVolSizeNumeric field if non-nil, zero value otherwise.

### GetMetadataVolSizeNumericOk

`func (o *PoolsResourceInner) GetMetadataVolSizeNumericOk() (*int32, bool)`

GetMetadataVolSizeNumericOk returns a tuple with the MetadataVolSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataVolSizeNumeric

`func (o *PoolsResourceInner) SetMetadataVolSizeNumeric(v int32)`

SetMetadataVolSizeNumeric sets MetadataVolSizeNumeric field to given value.

### HasMetadataVolSizeNumeric

`func (o *PoolsResourceInner) HasMetadataVolSizeNumeric() bool`

HasMetadataVolSizeNumeric returns a boolean if a field has been set.

### GetMiddleThreshold

`func (o *PoolsResourceInner) GetMiddleThreshold() string`

GetMiddleThreshold returns the MiddleThreshold field if non-nil, zero value otherwise.

### GetMiddleThresholdOk

`func (o *PoolsResourceInner) GetMiddleThresholdOk() (*string, bool)`

GetMiddleThresholdOk returns a tuple with the MiddleThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleThreshold

`func (o *PoolsResourceInner) SetMiddleThreshold(v string)`

SetMiddleThreshold sets MiddleThreshold field to given value.

### HasMiddleThreshold

`func (o *PoolsResourceInner) HasMiddleThreshold() bool`

HasMiddleThreshold returns a boolean if a field has been set.

### GetMigration

`func (o *PoolsResourceInner) GetMigration() string`

GetMigration returns the Migration field if non-nil, zero value otherwise.

### GetMigrationOk

`func (o *PoolsResourceInner) GetMigrationOk() (*string, bool)`

GetMigrationOk returns a tuple with the Migration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigration

`func (o *PoolsResourceInner) SetMigration(v string)`

SetMigration sets Migration field to given value.

### HasMigration

`func (o *PoolsResourceInner) HasMigration() bool`

HasMigration returns a boolean if a field has been set.

### GetMigrationNumeric

`func (o *PoolsResourceInner) GetMigrationNumeric() int32`

GetMigrationNumeric returns the MigrationNumeric field if non-nil, zero value otherwise.

### GetMigrationNumericOk

`func (o *PoolsResourceInner) GetMigrationNumericOk() (*int32, bool)`

GetMigrationNumericOk returns a tuple with the MigrationNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationNumeric

`func (o *PoolsResourceInner) SetMigrationNumeric(v int32)`

SetMigrationNumeric sets MigrationNumeric field to given value.

### HasMigrationNumeric

`func (o *PoolsResourceInner) HasMigrationNumeric() bool`

HasMigrationNumeric returns a boolean if a field has been set.

### GetName

`func (o *PoolsResourceInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolsResourceInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolsResourceInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoolsResourceInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverCommitted

`func (o *PoolsResourceInner) GetOverCommitted() string`

GetOverCommitted returns the OverCommitted field if non-nil, zero value otherwise.

### GetOverCommittedOk

`func (o *PoolsResourceInner) GetOverCommittedOk() (*string, bool)`

GetOverCommittedOk returns a tuple with the OverCommitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverCommitted

`func (o *PoolsResourceInner) SetOverCommitted(v string)`

SetOverCommitted sets OverCommitted field to given value.

### HasOverCommitted

`func (o *PoolsResourceInner) HasOverCommitted() bool`

HasOverCommitted returns a boolean if a field has been set.

### GetOverCommittedNumeric

`func (o *PoolsResourceInner) GetOverCommittedNumeric() int32`

GetOverCommittedNumeric returns the OverCommittedNumeric field if non-nil, zero value otherwise.

### GetOverCommittedNumericOk

`func (o *PoolsResourceInner) GetOverCommittedNumericOk() (*int32, bool)`

GetOverCommittedNumericOk returns a tuple with the OverCommittedNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverCommittedNumeric

`func (o *PoolsResourceInner) SetOverCommittedNumeric(v int32)`

SetOverCommittedNumeric sets OverCommittedNumeric field to given value.

### HasOverCommittedNumeric

`func (o *PoolsResourceInner) HasOverCommittedNumeric() bool`

HasOverCommittedNumeric returns a boolean if a field has been set.

### GetOvercommit

`func (o *PoolsResourceInner) GetOvercommit() string`

GetOvercommit returns the Overcommit field if non-nil, zero value otherwise.

### GetOvercommitOk

`func (o *PoolsResourceInner) GetOvercommitOk() (*string, bool)`

GetOvercommitOk returns a tuple with the Overcommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvercommit

`func (o *PoolsResourceInner) SetOvercommit(v string)`

SetOvercommit sets Overcommit field to given value.

### HasOvercommit

`func (o *PoolsResourceInner) HasOvercommit() bool`

HasOvercommit returns a boolean if a field has been set.

### GetOvercommitNumeric

`func (o *PoolsResourceInner) GetOvercommitNumeric() int32`

GetOvercommitNumeric returns the OvercommitNumeric field if non-nil, zero value otherwise.

### GetOvercommitNumericOk

`func (o *PoolsResourceInner) GetOvercommitNumericOk() (*int32, bool)`

GetOvercommitNumericOk returns a tuple with the OvercommitNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvercommitNumeric

`func (o *PoolsResourceInner) SetOvercommitNumeric(v int32)`

SetOvercommitNumeric sets OvercommitNumeric field to given value.

### HasOvercommitNumeric

`func (o *PoolsResourceInner) HasOvercommitNumeric() bool`

HasOvercommitNumeric returns a boolean if a field has been set.

### GetOwner

`func (o *PoolsResourceInner) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PoolsResourceInner) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PoolsResourceInner) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PoolsResourceInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnerNumeric

`func (o *PoolsResourceInner) GetOwnerNumeric() int32`

GetOwnerNumeric returns the OwnerNumeric field if non-nil, zero value otherwise.

### GetOwnerNumericOk

`func (o *PoolsResourceInner) GetOwnerNumericOk() (*int32, bool)`

GetOwnerNumericOk returns a tuple with the OwnerNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerNumeric

`func (o *PoolsResourceInner) SetOwnerNumeric(v int32)`

SetOwnerNumeric sets OwnerNumeric field to given value.

### HasOwnerNumeric

`func (o *PoolsResourceInner) HasOwnerNumeric() bool`

HasOwnerNumeric returns a boolean if a field has been set.

### GetPageSize

`func (o *PoolsResourceInner) GetPageSize() string`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PoolsResourceInner) GetPageSizeOk() (*string, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PoolsResourceInner) SetPageSize(v string)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PoolsResourceInner) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageSizeNumeric

`func (o *PoolsResourceInner) GetPageSizeNumeric() int32`

GetPageSizeNumeric returns the PageSizeNumeric field if non-nil, zero value otherwise.

### GetPageSizeNumericOk

`func (o *PoolsResourceInner) GetPageSizeNumericOk() (*int32, bool)`

GetPageSizeNumericOk returns a tuple with the PageSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSizeNumeric

`func (o *PoolsResourceInner) SetPageSizeNumeric(v int32)`

SetPageSizeNumeric sets PageSizeNumeric field to given value.

### HasPageSizeNumeric

`func (o *PoolsResourceInner) HasPageSizeNumeric() bool`

HasPageSizeNumeric returns a boolean if a field has been set.

### GetPoolSectorFormat

`func (o *PoolsResourceInner) GetPoolSectorFormat() string`

GetPoolSectorFormat returns the PoolSectorFormat field if non-nil, zero value otherwise.

### GetPoolSectorFormatOk

`func (o *PoolsResourceInner) GetPoolSectorFormatOk() (*string, bool)`

GetPoolSectorFormatOk returns a tuple with the PoolSectorFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSectorFormat

`func (o *PoolsResourceInner) SetPoolSectorFormat(v string)`

SetPoolSectorFormat sets PoolSectorFormat field to given value.

### HasPoolSectorFormat

`func (o *PoolsResourceInner) HasPoolSectorFormat() bool`

HasPoolSectorFormat returns a boolean if a field has been set.

### GetPoolSectorFormatNumeric

`func (o *PoolsResourceInner) GetPoolSectorFormatNumeric() int32`

GetPoolSectorFormatNumeric returns the PoolSectorFormatNumeric field if non-nil, zero value otherwise.

### GetPoolSectorFormatNumericOk

`func (o *PoolsResourceInner) GetPoolSectorFormatNumericOk() (*int32, bool)`

GetPoolSectorFormatNumericOk returns a tuple with the PoolSectorFormatNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSectorFormatNumeric

`func (o *PoolsResourceInner) SetPoolSectorFormatNumeric(v int32)`

SetPoolSectorFormatNumeric sets PoolSectorFormatNumeric field to given value.

### HasPoolSectorFormatNumeric

`func (o *PoolsResourceInner) HasPoolSectorFormatNumeric() bool`

HasPoolSectorFormatNumeric returns a boolean if a field has been set.

### GetPreferredOwner

`func (o *PoolsResourceInner) GetPreferredOwner() string`

GetPreferredOwner returns the PreferredOwner field if non-nil, zero value otherwise.

### GetPreferredOwnerOk

`func (o *PoolsResourceInner) GetPreferredOwnerOk() (*string, bool)`

GetPreferredOwnerOk returns a tuple with the PreferredOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredOwner

`func (o *PoolsResourceInner) SetPreferredOwner(v string)`

SetPreferredOwner sets PreferredOwner field to given value.

### HasPreferredOwner

`func (o *PoolsResourceInner) HasPreferredOwner() bool`

HasPreferredOwner returns a boolean if a field has been set.

### GetPreferredOwnerNumeric

`func (o *PoolsResourceInner) GetPreferredOwnerNumeric() int32`

GetPreferredOwnerNumeric returns the PreferredOwnerNumeric field if non-nil, zero value otherwise.

### GetPreferredOwnerNumericOk

`func (o *PoolsResourceInner) GetPreferredOwnerNumericOk() (*int32, bool)`

GetPreferredOwnerNumericOk returns a tuple with the PreferredOwnerNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredOwnerNumeric

`func (o *PoolsResourceInner) SetPreferredOwnerNumeric(v int32)`

SetPreferredOwnerNumeric sets PreferredOwnerNumeric field to given value.

### HasPreferredOwnerNumeric

`func (o *PoolsResourceInner) HasPreferredOwnerNumeric() bool`

HasPreferredOwnerNumeric returns a boolean if a field has been set.

### GetReadFlashCache

`func (o *PoolsResourceInner) GetReadFlashCache() string`

GetReadFlashCache returns the ReadFlashCache field if non-nil, zero value otherwise.

### GetReadFlashCacheOk

`func (o *PoolsResourceInner) GetReadFlashCacheOk() (*string, bool)`

GetReadFlashCacheOk returns a tuple with the ReadFlashCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadFlashCache

`func (o *PoolsResourceInner) SetReadFlashCache(v string)`

SetReadFlashCache sets ReadFlashCache field to given value.

### HasReadFlashCache

`func (o *PoolsResourceInner) HasReadFlashCache() bool`

HasReadFlashCache returns a boolean if a field has been set.

### GetReadFlashCacheNumeric

`func (o *PoolsResourceInner) GetReadFlashCacheNumeric() int32`

GetReadFlashCacheNumeric returns the ReadFlashCacheNumeric field if non-nil, zero value otherwise.

### GetReadFlashCacheNumericOk

`func (o *PoolsResourceInner) GetReadFlashCacheNumericOk() (*int32, bool)`

GetReadFlashCacheNumericOk returns a tuple with the ReadFlashCacheNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadFlashCacheNumeric

`func (o *PoolsResourceInner) SetReadFlashCacheNumeric(v int32)`

SetReadFlashCacheNumeric sets ReadFlashCacheNumeric field to given value.

### HasReadFlashCacheNumeric

`func (o *PoolsResourceInner) HasReadFlashCacheNumeric() bool`

HasReadFlashCacheNumeric returns a boolean if a field has been set.

### GetRebalance

`func (o *PoolsResourceInner) GetRebalance() string`

GetRebalance returns the Rebalance field if non-nil, zero value otherwise.

### GetRebalanceOk

`func (o *PoolsResourceInner) GetRebalanceOk() (*string, bool)`

GetRebalanceOk returns a tuple with the Rebalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebalance

`func (o *PoolsResourceInner) SetRebalance(v string)`

SetRebalance sets Rebalance field to given value.

### HasRebalance

`func (o *PoolsResourceInner) HasRebalance() bool`

HasRebalance returns a boolean if a field has been set.

### GetRebalanceNumeric

`func (o *PoolsResourceInner) GetRebalanceNumeric() int32`

GetRebalanceNumeric returns the RebalanceNumeric field if non-nil, zero value otherwise.

### GetRebalanceNumericOk

`func (o *PoolsResourceInner) GetRebalanceNumericOk() (*int32, bool)`

GetRebalanceNumericOk returns a tuple with the RebalanceNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebalanceNumeric

`func (o *PoolsResourceInner) SetRebalanceNumeric(v int32)`

SetRebalanceNumeric sets RebalanceNumeric field to given value.

### HasRebalanceNumeric

`func (o *PoolsResourceInner) HasRebalanceNumeric() bool`

HasRebalanceNumeric returns a boolean if a field has been set.

### GetReservedSize

`func (o *PoolsResourceInner) GetReservedSize() string`

GetReservedSize returns the ReservedSize field if non-nil, zero value otherwise.

### GetReservedSizeOk

`func (o *PoolsResourceInner) GetReservedSizeOk() (*string, bool)`

GetReservedSizeOk returns a tuple with the ReservedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedSize

`func (o *PoolsResourceInner) SetReservedSize(v string)`

SetReservedSize sets ReservedSize field to given value.

### HasReservedSize

`func (o *PoolsResourceInner) HasReservedSize() bool`

HasReservedSize returns a boolean if a field has been set.

### GetReservedSizeNumeric

`func (o *PoolsResourceInner) GetReservedSizeNumeric() int32`

GetReservedSizeNumeric returns the ReservedSizeNumeric field if non-nil, zero value otherwise.

### GetReservedSizeNumericOk

`func (o *PoolsResourceInner) GetReservedSizeNumericOk() (*int32, bool)`

GetReservedSizeNumericOk returns a tuple with the ReservedSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedSizeNumeric

`func (o *PoolsResourceInner) SetReservedSizeNumeric(v int32)`

SetReservedSizeNumeric sets ReservedSizeNumeric field to given value.

### HasReservedSizeNumeric

`func (o *PoolsResourceInner) HasReservedSizeNumeric() bool`

HasReservedSizeNumeric returns a boolean if a field has been set.

### GetReservedUnallocSize

`func (o *PoolsResourceInner) GetReservedUnallocSize() string`

GetReservedUnallocSize returns the ReservedUnallocSize field if non-nil, zero value otherwise.

### GetReservedUnallocSizeOk

`func (o *PoolsResourceInner) GetReservedUnallocSizeOk() (*string, bool)`

GetReservedUnallocSizeOk returns a tuple with the ReservedUnallocSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedUnallocSize

`func (o *PoolsResourceInner) SetReservedUnallocSize(v string)`

SetReservedUnallocSize sets ReservedUnallocSize field to given value.

### HasReservedUnallocSize

`func (o *PoolsResourceInner) HasReservedUnallocSize() bool`

HasReservedUnallocSize returns a boolean if a field has been set.

### GetReservedUnallocSizeNumeric

`func (o *PoolsResourceInner) GetReservedUnallocSizeNumeric() int32`

GetReservedUnallocSizeNumeric returns the ReservedUnallocSizeNumeric field if non-nil, zero value otherwise.

### GetReservedUnallocSizeNumericOk

`func (o *PoolsResourceInner) GetReservedUnallocSizeNumericOk() (*int32, bool)`

GetReservedUnallocSizeNumericOk returns a tuple with the ReservedUnallocSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedUnallocSizeNumeric

`func (o *PoolsResourceInner) SetReservedUnallocSizeNumeric(v int32)`

SetReservedUnallocSizeNumeric sets ReservedUnallocSizeNumeric field to given value.

### HasReservedUnallocSizeNumeric

`func (o *PoolsResourceInner) HasReservedUnallocSizeNumeric() bool`

HasReservedUnallocSizeNumeric returns a boolean if a field has been set.

### GetSerialNumber

`func (o *PoolsResourceInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PoolsResourceInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PoolsResourceInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *PoolsResourceInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSnapSize

`func (o *PoolsResourceInner) GetSnapSize() string`

GetSnapSize returns the SnapSize field if non-nil, zero value otherwise.

### GetSnapSizeOk

`func (o *PoolsResourceInner) GetSnapSizeOk() (*string, bool)`

GetSnapSizeOk returns a tuple with the SnapSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapSize

`func (o *PoolsResourceInner) SetSnapSize(v string)`

SetSnapSize sets SnapSize field to given value.

### HasSnapSize

`func (o *PoolsResourceInner) HasSnapSize() bool`

HasSnapSize returns a boolean if a field has been set.

### GetSnapSizeNumeric

`func (o *PoolsResourceInner) GetSnapSizeNumeric() int32`

GetSnapSizeNumeric returns the SnapSizeNumeric field if non-nil, zero value otherwise.

### GetSnapSizeNumericOk

`func (o *PoolsResourceInner) GetSnapSizeNumericOk() (*int32, bool)`

GetSnapSizeNumericOk returns a tuple with the SnapSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapSizeNumeric

`func (o *PoolsResourceInner) SetSnapSizeNumeric(v int32)`

SetSnapSizeNumeric sets SnapSizeNumeric field to given value.

### HasSnapSizeNumeric

`func (o *PoolsResourceInner) HasSnapSizeNumeric() bool`

HasSnapSizeNumeric returns a boolean if a field has been set.

### GetStorageType

`func (o *PoolsResourceInner) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *PoolsResourceInner) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *PoolsResourceInner) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *PoolsResourceInner) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetStorageTypeNumeric

`func (o *PoolsResourceInner) GetStorageTypeNumeric() int32`

GetStorageTypeNumeric returns the StorageTypeNumeric field if non-nil, zero value otherwise.

### GetStorageTypeNumericOk

`func (o *PoolsResourceInner) GetStorageTypeNumericOk() (*int32, bool)`

GetStorageTypeNumericOk returns a tuple with the StorageTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypeNumeric

`func (o *PoolsResourceInner) SetStorageTypeNumeric(v int32)`

SetStorageTypeNumeric sets StorageTypeNumeric field to given value.

### HasStorageTypeNumeric

`func (o *PoolsResourceInner) HasStorageTypeNumeric() bool`

HasStorageTypeNumeric returns a boolean if a field has been set.

### GetTotalAvail

`func (o *PoolsResourceInner) GetTotalAvail() string`

GetTotalAvail returns the TotalAvail field if non-nil, zero value otherwise.

### GetTotalAvailOk

`func (o *PoolsResourceInner) GetTotalAvailOk() (*string, bool)`

GetTotalAvailOk returns a tuple with the TotalAvail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAvail

`func (o *PoolsResourceInner) SetTotalAvail(v string)`

SetTotalAvail sets TotalAvail field to given value.

### HasTotalAvail

`func (o *PoolsResourceInner) HasTotalAvail() bool`

HasTotalAvail returns a boolean if a field has been set.

### GetTotalAvailNumeric

`func (o *PoolsResourceInner) GetTotalAvailNumeric() int32`

GetTotalAvailNumeric returns the TotalAvailNumeric field if non-nil, zero value otherwise.

### GetTotalAvailNumericOk

`func (o *PoolsResourceInner) GetTotalAvailNumericOk() (*int32, bool)`

GetTotalAvailNumericOk returns a tuple with the TotalAvailNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAvailNumeric

`func (o *PoolsResourceInner) SetTotalAvailNumeric(v int32)`

SetTotalAvailNumeric sets TotalAvailNumeric field to given value.

### HasTotalAvailNumeric

`func (o *PoolsResourceInner) HasTotalAvailNumeric() bool`

HasTotalAvailNumeric returns a boolean if a field has been set.

### GetTotalRfcSize

`func (o *PoolsResourceInner) GetTotalRfcSize() string`

GetTotalRfcSize returns the TotalRfcSize field if non-nil, zero value otherwise.

### GetTotalRfcSizeOk

`func (o *PoolsResourceInner) GetTotalRfcSizeOk() (*string, bool)`

GetTotalRfcSizeOk returns a tuple with the TotalRfcSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRfcSize

`func (o *PoolsResourceInner) SetTotalRfcSize(v string)`

SetTotalRfcSize sets TotalRfcSize field to given value.

### HasTotalRfcSize

`func (o *PoolsResourceInner) HasTotalRfcSize() bool`

HasTotalRfcSize returns a boolean if a field has been set.

### GetTotalRfcSizeNumeric

`func (o *PoolsResourceInner) GetTotalRfcSizeNumeric() int32`

GetTotalRfcSizeNumeric returns the TotalRfcSizeNumeric field if non-nil, zero value otherwise.

### GetTotalRfcSizeNumericOk

`func (o *PoolsResourceInner) GetTotalRfcSizeNumericOk() (*int32, bool)`

GetTotalRfcSizeNumericOk returns a tuple with the TotalRfcSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRfcSizeNumeric

`func (o *PoolsResourceInner) SetTotalRfcSizeNumeric(v int32)`

SetTotalRfcSizeNumeric sets TotalRfcSizeNumeric field to given value.

### HasTotalRfcSizeNumeric

`func (o *PoolsResourceInner) HasTotalRfcSizeNumeric() bool`

HasTotalRfcSizeNumeric returns a boolean if a field has been set.

### GetTotalSize

`func (o *PoolsResourceInner) GetTotalSize() string`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *PoolsResourceInner) GetTotalSizeOk() (*string, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *PoolsResourceInner) SetTotalSize(v string)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *PoolsResourceInner) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetTotalSizeNumeric

`func (o *PoolsResourceInner) GetTotalSizeNumeric() int32`

GetTotalSizeNumeric returns the TotalSizeNumeric field if non-nil, zero value otherwise.

### GetTotalSizeNumericOk

`func (o *PoolsResourceInner) GetTotalSizeNumericOk() (*int32, bool)`

GetTotalSizeNumericOk returns a tuple with the TotalSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeNumeric

`func (o *PoolsResourceInner) SetTotalSizeNumeric(v int32)`

SetTotalSizeNumeric sets TotalSizeNumeric field to given value.

### HasTotalSizeNumeric

`func (o *PoolsResourceInner) HasTotalSizeNumeric() bool`

HasTotalSizeNumeric returns a boolean if a field has been set.

### GetUrl

`func (o *PoolsResourceInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PoolsResourceInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PoolsResourceInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PoolsResourceInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUtilityRunning

`func (o *PoolsResourceInner) GetUtilityRunning() string`

GetUtilityRunning returns the UtilityRunning field if non-nil, zero value otherwise.

### GetUtilityRunningOk

`func (o *PoolsResourceInner) GetUtilityRunningOk() (*string, bool)`

GetUtilityRunningOk returns a tuple with the UtilityRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilityRunning

`func (o *PoolsResourceInner) SetUtilityRunning(v string)`

SetUtilityRunning sets UtilityRunning field to given value.

### HasUtilityRunning

`func (o *PoolsResourceInner) HasUtilityRunning() bool`

HasUtilityRunning returns a boolean if a field has been set.

### GetUtilityRunningNumeric

`func (o *PoolsResourceInner) GetUtilityRunningNumeric() int32`

GetUtilityRunningNumeric returns the UtilityRunningNumeric field if non-nil, zero value otherwise.

### GetUtilityRunningNumericOk

`func (o *PoolsResourceInner) GetUtilityRunningNumericOk() (*int32, bool)`

GetUtilityRunningNumericOk returns a tuple with the UtilityRunningNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilityRunningNumeric

`func (o *PoolsResourceInner) SetUtilityRunningNumeric(v int32)`

SetUtilityRunningNumeric sets UtilityRunningNumeric field to given value.

### HasUtilityRunningNumeric

`func (o *PoolsResourceInner) HasUtilityRunningNumeric() bool`

HasUtilityRunningNumeric returns a boolean if a field has been set.

### GetVolumes

`func (o *PoolsResourceInner) GetVolumes() int32`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *PoolsResourceInner) GetVolumesOk() (*int32, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *PoolsResourceInner) SetVolumes(v int32)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *PoolsResourceInner) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetZeroScan

`func (o *PoolsResourceInner) GetZeroScan() string`

GetZeroScan returns the ZeroScan field if non-nil, zero value otherwise.

### GetZeroScanOk

`func (o *PoolsResourceInner) GetZeroScanOk() (*string, bool)`

GetZeroScanOk returns a tuple with the ZeroScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroScan

`func (o *PoolsResourceInner) SetZeroScan(v string)`

SetZeroScan sets ZeroScan field to given value.

### HasZeroScan

`func (o *PoolsResourceInner) HasZeroScan() bool`

HasZeroScan returns a boolean if a field has been set.

### GetZeroScanNumeric

`func (o *PoolsResourceInner) GetZeroScanNumeric() int32`

GetZeroScanNumeric returns the ZeroScanNumeric field if non-nil, zero value otherwise.

### GetZeroScanNumericOk

`func (o *PoolsResourceInner) GetZeroScanNumericOk() (*int32, bool)`

GetZeroScanNumericOk returns a tuple with the ZeroScanNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroScanNumeric

`func (o *PoolsResourceInner) SetZeroScanNumeric(v int32)`

SetZeroScanNumeric sets ZeroScanNumeric field to given value.

### HasZeroScanNumeric

`func (o *PoolsResourceInner) HasZeroScanNumeric() bool`

HasZeroScanNumeric returns a boolean if a field has been set.

### GetDiskGroups

`func (o *PoolsResourceInner) GetDiskGroups() []DiskGroupsResourceInner`

GetDiskGroups returns the DiskGroups field if non-nil, zero value otherwise.

### GetDiskGroupsOk

`func (o *PoolsResourceInner) GetDiskGroupsOk() (*[]DiskGroupsResourceInner, bool)`

GetDiskGroupsOk returns a tuple with the DiskGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroups

`func (o *PoolsResourceInner) SetDiskGroups(v []DiskGroupsResourceInner)`

SetDiskGroups sets DiskGroups field to given value.

### HasDiskGroups

`func (o *PoolsResourceInner) HasDiskGroups() bool`

HasDiskGroups returns a boolean if a field has been set.

### GetTiers

`func (o *PoolsResourceInner) GetTiers() []TiersResourceInner`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *PoolsResourceInner) GetTiersOk() (*[]TiersResourceInner, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *PoolsResourceInner) SetTiers(v []TiersResourceInner)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *PoolsResourceInner) HasTiers() bool`

HasTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


