# FcPortResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**ConfiguredTopology** | Pointer to **string** |  | [optional] 
**ConfiguredTopologyNumeric** | Pointer to **int64** |  | [optional] 
**PrimaryLoopId** | Pointer to **string** | Primary loop ID for this port | [optional] 
**SfpPartNumber** | Pointer to **string** |  | [optional] 
**SfpPresent** | Pointer to **string** |  | [optional] 
**SfpPresentNumeric** | Pointer to **int64** |  | [optional] 
**SfpRevision** | Pointer to **string** |  | [optional] 
**SfpStatus** | Pointer to **string** |  | [optional] 
**SfpStatusNumeric** | Pointer to **int64** |  | [optional] 
**SfpSupportedSpeeds** | Pointer to **string** |  | [optional] 
**SfpSupportedSpeedsNumeric** | Pointer to **int64** |  | [optional] 
**SfpVendor** | Pointer to **string** |  | [optional] 

## Methods

### NewFcPortResourceInner

`func NewFcPortResourceInner() *FcPortResourceInner`

NewFcPortResourceInner instantiates a new FcPortResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFcPortResourceInnerWithDefaults

`func NewFcPortResourceInnerWithDefaults() *FcPortResourceInner`

NewFcPortResourceInnerWithDefaults instantiates a new FcPortResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *FcPortResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *FcPortResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *FcPortResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *FcPortResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *FcPortResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FcPortResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FcPortResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FcPortResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetConfiguredTopology

`func (o *FcPortResourceInner) GetConfiguredTopology() string`

GetConfiguredTopology returns the ConfiguredTopology field if non-nil, zero value otherwise.

### GetConfiguredTopologyOk

`func (o *FcPortResourceInner) GetConfiguredTopologyOk() (*string, bool)`

GetConfiguredTopologyOk returns a tuple with the ConfiguredTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredTopology

`func (o *FcPortResourceInner) SetConfiguredTopology(v string)`

SetConfiguredTopology sets ConfiguredTopology field to given value.

### HasConfiguredTopology

`func (o *FcPortResourceInner) HasConfiguredTopology() bool`

HasConfiguredTopology returns a boolean if a field has been set.

### GetConfiguredTopologyNumeric

`func (o *FcPortResourceInner) GetConfiguredTopologyNumeric() int64`

GetConfiguredTopologyNumeric returns the ConfiguredTopologyNumeric field if non-nil, zero value otherwise.

### GetConfiguredTopologyNumericOk

`func (o *FcPortResourceInner) GetConfiguredTopologyNumericOk() (*int64, bool)`

GetConfiguredTopologyNumericOk returns a tuple with the ConfiguredTopologyNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredTopologyNumeric

`func (o *FcPortResourceInner) SetConfiguredTopologyNumeric(v int64)`

SetConfiguredTopologyNumeric sets ConfiguredTopologyNumeric field to given value.

### HasConfiguredTopologyNumeric

`func (o *FcPortResourceInner) HasConfiguredTopologyNumeric() bool`

HasConfiguredTopologyNumeric returns a boolean if a field has been set.

### GetPrimaryLoopId

`func (o *FcPortResourceInner) GetPrimaryLoopId() string`

GetPrimaryLoopId returns the PrimaryLoopId field if non-nil, zero value otherwise.

### GetPrimaryLoopIdOk

`func (o *FcPortResourceInner) GetPrimaryLoopIdOk() (*string, bool)`

GetPrimaryLoopIdOk returns a tuple with the PrimaryLoopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLoopId

`func (o *FcPortResourceInner) SetPrimaryLoopId(v string)`

SetPrimaryLoopId sets PrimaryLoopId field to given value.

### HasPrimaryLoopId

`func (o *FcPortResourceInner) HasPrimaryLoopId() bool`

HasPrimaryLoopId returns a boolean if a field has been set.

### GetSfpPartNumber

`func (o *FcPortResourceInner) GetSfpPartNumber() string`

GetSfpPartNumber returns the SfpPartNumber field if non-nil, zero value otherwise.

### GetSfpPartNumberOk

`func (o *FcPortResourceInner) GetSfpPartNumberOk() (*string, bool)`

GetSfpPartNumberOk returns a tuple with the SfpPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpPartNumber

`func (o *FcPortResourceInner) SetSfpPartNumber(v string)`

SetSfpPartNumber sets SfpPartNumber field to given value.

### HasSfpPartNumber

`func (o *FcPortResourceInner) HasSfpPartNumber() bool`

HasSfpPartNumber returns a boolean if a field has been set.

### GetSfpPresent

`func (o *FcPortResourceInner) GetSfpPresent() string`

GetSfpPresent returns the SfpPresent field if non-nil, zero value otherwise.

### GetSfpPresentOk

`func (o *FcPortResourceInner) GetSfpPresentOk() (*string, bool)`

GetSfpPresentOk returns a tuple with the SfpPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpPresent

`func (o *FcPortResourceInner) SetSfpPresent(v string)`

SetSfpPresent sets SfpPresent field to given value.

### HasSfpPresent

`func (o *FcPortResourceInner) HasSfpPresent() bool`

HasSfpPresent returns a boolean if a field has been set.

### GetSfpPresentNumeric

`func (o *FcPortResourceInner) GetSfpPresentNumeric() int64`

GetSfpPresentNumeric returns the SfpPresentNumeric field if non-nil, zero value otherwise.

### GetSfpPresentNumericOk

`func (o *FcPortResourceInner) GetSfpPresentNumericOk() (*int64, bool)`

GetSfpPresentNumericOk returns a tuple with the SfpPresentNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpPresentNumeric

`func (o *FcPortResourceInner) SetSfpPresentNumeric(v int64)`

SetSfpPresentNumeric sets SfpPresentNumeric field to given value.

### HasSfpPresentNumeric

`func (o *FcPortResourceInner) HasSfpPresentNumeric() bool`

HasSfpPresentNumeric returns a boolean if a field has been set.

### GetSfpRevision

`func (o *FcPortResourceInner) GetSfpRevision() string`

GetSfpRevision returns the SfpRevision field if non-nil, zero value otherwise.

### GetSfpRevisionOk

`func (o *FcPortResourceInner) GetSfpRevisionOk() (*string, bool)`

GetSfpRevisionOk returns a tuple with the SfpRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpRevision

`func (o *FcPortResourceInner) SetSfpRevision(v string)`

SetSfpRevision sets SfpRevision field to given value.

### HasSfpRevision

`func (o *FcPortResourceInner) HasSfpRevision() bool`

HasSfpRevision returns a boolean if a field has been set.

### GetSfpStatus

`func (o *FcPortResourceInner) GetSfpStatus() string`

GetSfpStatus returns the SfpStatus field if non-nil, zero value otherwise.

### GetSfpStatusOk

`func (o *FcPortResourceInner) GetSfpStatusOk() (*string, bool)`

GetSfpStatusOk returns a tuple with the SfpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpStatus

`func (o *FcPortResourceInner) SetSfpStatus(v string)`

SetSfpStatus sets SfpStatus field to given value.

### HasSfpStatus

`func (o *FcPortResourceInner) HasSfpStatus() bool`

HasSfpStatus returns a boolean if a field has been set.

### GetSfpStatusNumeric

`func (o *FcPortResourceInner) GetSfpStatusNumeric() int64`

GetSfpStatusNumeric returns the SfpStatusNumeric field if non-nil, zero value otherwise.

### GetSfpStatusNumericOk

`func (o *FcPortResourceInner) GetSfpStatusNumericOk() (*int64, bool)`

GetSfpStatusNumericOk returns a tuple with the SfpStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpStatusNumeric

`func (o *FcPortResourceInner) SetSfpStatusNumeric(v int64)`

SetSfpStatusNumeric sets SfpStatusNumeric field to given value.

### HasSfpStatusNumeric

`func (o *FcPortResourceInner) HasSfpStatusNumeric() bool`

HasSfpStatusNumeric returns a boolean if a field has been set.

### GetSfpSupportedSpeeds

`func (o *FcPortResourceInner) GetSfpSupportedSpeeds() string`

GetSfpSupportedSpeeds returns the SfpSupportedSpeeds field if non-nil, zero value otherwise.

### GetSfpSupportedSpeedsOk

`func (o *FcPortResourceInner) GetSfpSupportedSpeedsOk() (*string, bool)`

GetSfpSupportedSpeedsOk returns a tuple with the SfpSupportedSpeeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpSupportedSpeeds

`func (o *FcPortResourceInner) SetSfpSupportedSpeeds(v string)`

SetSfpSupportedSpeeds sets SfpSupportedSpeeds field to given value.

### HasSfpSupportedSpeeds

`func (o *FcPortResourceInner) HasSfpSupportedSpeeds() bool`

HasSfpSupportedSpeeds returns a boolean if a field has been set.

### GetSfpSupportedSpeedsNumeric

`func (o *FcPortResourceInner) GetSfpSupportedSpeedsNumeric() int64`

GetSfpSupportedSpeedsNumeric returns the SfpSupportedSpeedsNumeric field if non-nil, zero value otherwise.

### GetSfpSupportedSpeedsNumericOk

`func (o *FcPortResourceInner) GetSfpSupportedSpeedsNumericOk() (*int64, bool)`

GetSfpSupportedSpeedsNumericOk returns a tuple with the SfpSupportedSpeedsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpSupportedSpeedsNumeric

`func (o *FcPortResourceInner) SetSfpSupportedSpeedsNumeric(v int64)`

SetSfpSupportedSpeedsNumeric sets SfpSupportedSpeedsNumeric field to given value.

### HasSfpSupportedSpeedsNumeric

`func (o *FcPortResourceInner) HasSfpSupportedSpeedsNumeric() bool`

HasSfpSupportedSpeedsNumeric returns a boolean if a field has been set.

### GetSfpVendor

`func (o *FcPortResourceInner) GetSfpVendor() string`

GetSfpVendor returns the SfpVendor field if non-nil, zero value otherwise.

### GetSfpVendorOk

`func (o *FcPortResourceInner) GetSfpVendorOk() (*string, bool)`

GetSfpVendorOk returns a tuple with the SfpVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpVendor

`func (o *FcPortResourceInner) SetSfpVendor(v string)`

SetSfpVendor sets SfpVendor field to given value.

### HasSfpVendor

`func (o *FcPortResourceInner) HasSfpVendor() bool`

HasSfpVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


