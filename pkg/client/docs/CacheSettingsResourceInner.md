# CacheSettingsResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**CacheBlockSize** | Pointer to **int64** |  | [optional] 
**OperationMode** | Pointer to **string** | Storage system operation mode (redundancy mode) | [optional] 
**OperationModeNumeric** | Pointer to **int64** | Storage system operation mode (redundancy mode)( In numeric form ) | [optional] 
**PiFormat** | Pointer to **string** | Used to describe the Host Protection Information | [optional] 
**PiFormatNumeric** | Pointer to **int64** | Used to describe the Host Protection Information( In numeric form ) | [optional] 
**ControllerCacheParameters** | Pointer to [**[]ControllerCacheParametersResourceInner**](ControllerCacheParametersResourceInner.md) |  | [optional] 

## Methods

### NewCacheSettingsResourceInner

`func NewCacheSettingsResourceInner() *CacheSettingsResourceInner`

NewCacheSettingsResourceInner instantiates a new CacheSettingsResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingsResourceInnerWithDefaults

`func NewCacheSettingsResourceInnerWithDefaults() *CacheSettingsResourceInner`

NewCacheSettingsResourceInnerWithDefaults instantiates a new CacheSettingsResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *CacheSettingsResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *CacheSettingsResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *CacheSettingsResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *CacheSettingsResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *CacheSettingsResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CacheSettingsResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CacheSettingsResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CacheSettingsResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetCacheBlockSize

`func (o *CacheSettingsResourceInner) GetCacheBlockSize() int64`

GetCacheBlockSize returns the CacheBlockSize field if non-nil, zero value otherwise.

### GetCacheBlockSizeOk

`func (o *CacheSettingsResourceInner) GetCacheBlockSizeOk() (*int64, bool)`

GetCacheBlockSizeOk returns a tuple with the CacheBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheBlockSize

`func (o *CacheSettingsResourceInner) SetCacheBlockSize(v int64)`

SetCacheBlockSize sets CacheBlockSize field to given value.

### HasCacheBlockSize

`func (o *CacheSettingsResourceInner) HasCacheBlockSize() bool`

HasCacheBlockSize returns a boolean if a field has been set.

### GetOperationMode

`func (o *CacheSettingsResourceInner) GetOperationMode() string`

GetOperationMode returns the OperationMode field if non-nil, zero value otherwise.

### GetOperationModeOk

`func (o *CacheSettingsResourceInner) GetOperationModeOk() (*string, bool)`

GetOperationModeOk returns a tuple with the OperationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationMode

`func (o *CacheSettingsResourceInner) SetOperationMode(v string)`

SetOperationMode sets OperationMode field to given value.

### HasOperationMode

`func (o *CacheSettingsResourceInner) HasOperationMode() bool`

HasOperationMode returns a boolean if a field has been set.

### GetOperationModeNumeric

`func (o *CacheSettingsResourceInner) GetOperationModeNumeric() int64`

GetOperationModeNumeric returns the OperationModeNumeric field if non-nil, zero value otherwise.

### GetOperationModeNumericOk

`func (o *CacheSettingsResourceInner) GetOperationModeNumericOk() (*int64, bool)`

GetOperationModeNumericOk returns a tuple with the OperationModeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationModeNumeric

`func (o *CacheSettingsResourceInner) SetOperationModeNumeric(v int64)`

SetOperationModeNumeric sets OperationModeNumeric field to given value.

### HasOperationModeNumeric

`func (o *CacheSettingsResourceInner) HasOperationModeNumeric() bool`

HasOperationModeNumeric returns a boolean if a field has been set.

### GetPiFormat

`func (o *CacheSettingsResourceInner) GetPiFormat() string`

GetPiFormat returns the PiFormat field if non-nil, zero value otherwise.

### GetPiFormatOk

`func (o *CacheSettingsResourceInner) GetPiFormatOk() (*string, bool)`

GetPiFormatOk returns a tuple with the PiFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiFormat

`func (o *CacheSettingsResourceInner) SetPiFormat(v string)`

SetPiFormat sets PiFormat field to given value.

### HasPiFormat

`func (o *CacheSettingsResourceInner) HasPiFormat() bool`

HasPiFormat returns a boolean if a field has been set.

### GetPiFormatNumeric

`func (o *CacheSettingsResourceInner) GetPiFormatNumeric() int64`

GetPiFormatNumeric returns the PiFormatNumeric field if non-nil, zero value otherwise.

### GetPiFormatNumericOk

`func (o *CacheSettingsResourceInner) GetPiFormatNumericOk() (*int64, bool)`

GetPiFormatNumericOk returns a tuple with the PiFormatNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiFormatNumeric

`func (o *CacheSettingsResourceInner) SetPiFormatNumeric(v int64)`

SetPiFormatNumeric sets PiFormatNumeric field to given value.

### HasPiFormatNumeric

`func (o *CacheSettingsResourceInner) HasPiFormatNumeric() bool`

HasPiFormatNumeric returns a boolean if a field has been set.

### GetControllerCacheParameters

`func (o *CacheSettingsResourceInner) GetControllerCacheParameters() []ControllerCacheParametersResourceInner`

GetControllerCacheParameters returns the ControllerCacheParameters field if non-nil, zero value otherwise.

### GetControllerCacheParametersOk

`func (o *CacheSettingsResourceInner) GetControllerCacheParametersOk() (*[]ControllerCacheParametersResourceInner, bool)`

GetControllerCacheParametersOk returns a tuple with the ControllerCacheParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerCacheParameters

`func (o *CacheSettingsResourceInner) SetControllerCacheParameters(v []ControllerCacheParametersResourceInner)`

SetControllerCacheParameters sets ControllerCacheParameters field to given value.

### HasControllerCacheParameters

`func (o *CacheSettingsResourceInner) HasControllerCacheParameters() bool`

HasControllerCacheParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


