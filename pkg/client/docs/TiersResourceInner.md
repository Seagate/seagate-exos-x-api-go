# TiersResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**AffinitySize** | Pointer to **string** |  | [optional] 
**AffinitySizeNumeric** | Pointer to **int64** |  | [optional] 
**AllocatedSize** | Pointer to **string** |  | [optional] 
**AllocatedSizeNumeric** | Pointer to **int64** |  | [optional] 
**AvailableSize** | Pointer to **string** |  | [optional] 
**AvailableSizeNumeric** | Pointer to **int64** |  | [optional] 
**Diskcount** | Pointer to **int64** | Number of disks | [optional] 
**Pool** | Pointer to **string** | Pool | [optional] 
**PoolPercentage** | Pointer to **int64** | Portion of the virtual pool used by this disk group | [optional] 
**RawSize** | Pointer to **string** |  | [optional] 
**RawSizeNumeric** | Pointer to **int64** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Tier** | Pointer to **string** |  | [optional] 
**TierNumeric** | Pointer to **int64** |  | [optional] 
**TotalSize** | Pointer to **string** | The total size formatted using the session settings for base, precision, and units | [optional] 
**TotalSizeNumeric** | Pointer to **int64** | The total size formatted using the session settings for base, precision, and units( In numeric form ) | [optional] 

## Methods

### NewTiersResourceInner

`func NewTiersResourceInner() *TiersResourceInner`

NewTiersResourceInner instantiates a new TiersResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTiersResourceInnerWithDefaults

`func NewTiersResourceInnerWithDefaults() *TiersResourceInner`

NewTiersResourceInnerWithDefaults instantiates a new TiersResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *TiersResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *TiersResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *TiersResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *TiersResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *TiersResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TiersResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TiersResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TiersResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetAffinitySize

`func (o *TiersResourceInner) GetAffinitySize() string`

GetAffinitySize returns the AffinitySize field if non-nil, zero value otherwise.

### GetAffinitySizeOk

`func (o *TiersResourceInner) GetAffinitySizeOk() (*string, bool)`

GetAffinitySizeOk returns a tuple with the AffinitySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinitySize

`func (o *TiersResourceInner) SetAffinitySize(v string)`

SetAffinitySize sets AffinitySize field to given value.

### HasAffinitySize

`func (o *TiersResourceInner) HasAffinitySize() bool`

HasAffinitySize returns a boolean if a field has been set.

### GetAffinitySizeNumeric

`func (o *TiersResourceInner) GetAffinitySizeNumeric() int64`

GetAffinitySizeNumeric returns the AffinitySizeNumeric field if non-nil, zero value otherwise.

### GetAffinitySizeNumericOk

`func (o *TiersResourceInner) GetAffinitySizeNumericOk() (*int64, bool)`

GetAffinitySizeNumericOk returns a tuple with the AffinitySizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinitySizeNumeric

`func (o *TiersResourceInner) SetAffinitySizeNumeric(v int64)`

SetAffinitySizeNumeric sets AffinitySizeNumeric field to given value.

### HasAffinitySizeNumeric

`func (o *TiersResourceInner) HasAffinitySizeNumeric() bool`

HasAffinitySizeNumeric returns a boolean if a field has been set.

### GetAllocatedSize

`func (o *TiersResourceInner) GetAllocatedSize() string`

GetAllocatedSize returns the AllocatedSize field if non-nil, zero value otherwise.

### GetAllocatedSizeOk

`func (o *TiersResourceInner) GetAllocatedSizeOk() (*string, bool)`

GetAllocatedSizeOk returns a tuple with the AllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSize

`func (o *TiersResourceInner) SetAllocatedSize(v string)`

SetAllocatedSize sets AllocatedSize field to given value.

### HasAllocatedSize

`func (o *TiersResourceInner) HasAllocatedSize() bool`

HasAllocatedSize returns a boolean if a field has been set.

### GetAllocatedSizeNumeric

`func (o *TiersResourceInner) GetAllocatedSizeNumeric() int64`

GetAllocatedSizeNumeric returns the AllocatedSizeNumeric field if non-nil, zero value otherwise.

### GetAllocatedSizeNumericOk

`func (o *TiersResourceInner) GetAllocatedSizeNumericOk() (*int64, bool)`

GetAllocatedSizeNumericOk returns a tuple with the AllocatedSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedSizeNumeric

`func (o *TiersResourceInner) SetAllocatedSizeNumeric(v int64)`

SetAllocatedSizeNumeric sets AllocatedSizeNumeric field to given value.

### HasAllocatedSizeNumeric

`func (o *TiersResourceInner) HasAllocatedSizeNumeric() bool`

HasAllocatedSizeNumeric returns a boolean if a field has been set.

### GetAvailableSize

`func (o *TiersResourceInner) GetAvailableSize() string`

GetAvailableSize returns the AvailableSize field if non-nil, zero value otherwise.

### GetAvailableSizeOk

`func (o *TiersResourceInner) GetAvailableSizeOk() (*string, bool)`

GetAvailableSizeOk returns a tuple with the AvailableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSize

`func (o *TiersResourceInner) SetAvailableSize(v string)`

SetAvailableSize sets AvailableSize field to given value.

### HasAvailableSize

`func (o *TiersResourceInner) HasAvailableSize() bool`

HasAvailableSize returns a boolean if a field has been set.

### GetAvailableSizeNumeric

`func (o *TiersResourceInner) GetAvailableSizeNumeric() int64`

GetAvailableSizeNumeric returns the AvailableSizeNumeric field if non-nil, zero value otherwise.

### GetAvailableSizeNumericOk

`func (o *TiersResourceInner) GetAvailableSizeNumericOk() (*int64, bool)`

GetAvailableSizeNumericOk returns a tuple with the AvailableSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSizeNumeric

`func (o *TiersResourceInner) SetAvailableSizeNumeric(v int64)`

SetAvailableSizeNumeric sets AvailableSizeNumeric field to given value.

### HasAvailableSizeNumeric

`func (o *TiersResourceInner) HasAvailableSizeNumeric() bool`

HasAvailableSizeNumeric returns a boolean if a field has been set.

### GetDiskcount

`func (o *TiersResourceInner) GetDiskcount() int64`

GetDiskcount returns the Diskcount field if non-nil, zero value otherwise.

### GetDiskcountOk

`func (o *TiersResourceInner) GetDiskcountOk() (*int64, bool)`

GetDiskcountOk returns a tuple with the Diskcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskcount

`func (o *TiersResourceInner) SetDiskcount(v int64)`

SetDiskcount sets Diskcount field to given value.

### HasDiskcount

`func (o *TiersResourceInner) HasDiskcount() bool`

HasDiskcount returns a boolean if a field has been set.

### GetPool

`func (o *TiersResourceInner) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *TiersResourceInner) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *TiersResourceInner) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *TiersResourceInner) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPoolPercentage

`func (o *TiersResourceInner) GetPoolPercentage() int64`

GetPoolPercentage returns the PoolPercentage field if non-nil, zero value otherwise.

### GetPoolPercentageOk

`func (o *TiersResourceInner) GetPoolPercentageOk() (*int64, bool)`

GetPoolPercentageOk returns a tuple with the PoolPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolPercentage

`func (o *TiersResourceInner) SetPoolPercentage(v int64)`

SetPoolPercentage sets PoolPercentage field to given value.

### HasPoolPercentage

`func (o *TiersResourceInner) HasPoolPercentage() bool`

HasPoolPercentage returns a boolean if a field has been set.

### GetRawSize

`func (o *TiersResourceInner) GetRawSize() string`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *TiersResourceInner) GetRawSizeOk() (*string, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *TiersResourceInner) SetRawSize(v string)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *TiersResourceInner) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### GetRawSizeNumeric

`func (o *TiersResourceInner) GetRawSizeNumeric() int64`

GetRawSizeNumeric returns the RawSizeNumeric field if non-nil, zero value otherwise.

### GetRawSizeNumericOk

`func (o *TiersResourceInner) GetRawSizeNumericOk() (*int64, bool)`

GetRawSizeNumericOk returns a tuple with the RawSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSizeNumeric

`func (o *TiersResourceInner) SetRawSizeNumeric(v int64)`

SetRawSizeNumeric sets RawSizeNumeric field to given value.

### HasRawSizeNumeric

`func (o *TiersResourceInner) HasRawSizeNumeric() bool`

HasRawSizeNumeric returns a boolean if a field has been set.

### GetSerialNumber

`func (o *TiersResourceInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *TiersResourceInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *TiersResourceInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *TiersResourceInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetTier

`func (o *TiersResourceInner) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *TiersResourceInner) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *TiersResourceInner) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *TiersResourceInner) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetTierNumeric

`func (o *TiersResourceInner) GetTierNumeric() int64`

GetTierNumeric returns the TierNumeric field if non-nil, zero value otherwise.

### GetTierNumericOk

`func (o *TiersResourceInner) GetTierNumericOk() (*int64, bool)`

GetTierNumericOk returns a tuple with the TierNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierNumeric

`func (o *TiersResourceInner) SetTierNumeric(v int64)`

SetTierNumeric sets TierNumeric field to given value.

### HasTierNumeric

`func (o *TiersResourceInner) HasTierNumeric() bool`

HasTierNumeric returns a boolean if a field has been set.

### GetTotalSize

`func (o *TiersResourceInner) GetTotalSize() string`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *TiersResourceInner) GetTotalSizeOk() (*string, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *TiersResourceInner) SetTotalSize(v string)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *TiersResourceInner) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetTotalSizeNumeric

`func (o *TiersResourceInner) GetTotalSizeNumeric() int64`

GetTotalSizeNumeric returns the TotalSizeNumeric field if non-nil, zero value otherwise.

### GetTotalSizeNumericOk

`func (o *TiersResourceInner) GetTotalSizeNumericOk() (*int64, bool)`

GetTotalSizeNumericOk returns a tuple with the TotalSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeNumeric

`func (o *TiersResourceInner) SetTotalSizeNumeric(v int64)`

SetTotalSizeNumeric sets TotalSizeNumeric field to given value.

### HasTotalSizeNumeric

`func (o *TiersResourceInner) HasTotalSizeNumeric() bool`

HasTotalSizeNumeric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


