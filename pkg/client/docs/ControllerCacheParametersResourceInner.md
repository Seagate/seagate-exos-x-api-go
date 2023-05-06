# ControllerCacheParametersResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**CacheFlush** | Pointer to **string** |  | [optional] 
<<<<<<< HEAD
**CacheFlushNumeric** | Pointer to **int64** |  | [optional] 
**ControllerId** | Pointer to **string** |  | [optional] 
**ControllerIdNumeric** | Pointer to **int64** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**MemoryCardHealth** | Pointer to **string** |  | [optional] 
**MemoryCardHealthNumeric** | Pointer to **int64** |  | [optional] 
**MemoryCardStatus** | Pointer to **string** |  | [optional] 
**MemoryCardStatusNumeric** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**WriteBackStatus** | Pointer to **string** | Indicates whether disk write-back cache is enabled | [optional] 
**WriteBackStatusNumeric** | Pointer to **int64** | Indicates whether disk write-back cache is enabled( In numeric form ) | [optional] 
=======
**CacheFlushNumeric** | Pointer to **int32** |  | [optional] 
**ControllerId** | Pointer to **string** |  | [optional] 
**ControllerIdNumeric** | Pointer to **int32** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**MemoryCardHealth** | Pointer to **string** |  | [optional] 
**MemoryCardHealthNumeric** | Pointer to **int32** |  | [optional] 
**MemoryCardStatus** | Pointer to **string** |  | [optional] 
**MemoryCardStatusNumeric** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**WriteBackStatus** | Pointer to **string** | Indicates whether disk write-back cache is enabled | [optional] 
**WriteBackStatusNumeric** | Pointer to **int32** | Indicates whether disk write-back cache is enabled( In numeric form ) | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

## Methods

### NewControllerCacheParametersResourceInner

`func NewControllerCacheParametersResourceInner() *ControllerCacheParametersResourceInner`

NewControllerCacheParametersResourceInner instantiates a new ControllerCacheParametersResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerCacheParametersResourceInnerWithDefaults

`func NewControllerCacheParametersResourceInnerWithDefaults() *ControllerCacheParametersResourceInner`

NewControllerCacheParametersResourceInnerWithDefaults instantiates a new ControllerCacheParametersResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *ControllerCacheParametersResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ControllerCacheParametersResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ControllerCacheParametersResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ControllerCacheParametersResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *ControllerCacheParametersResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ControllerCacheParametersResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ControllerCacheParametersResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ControllerCacheParametersResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetCacheFlush

`func (o *ControllerCacheParametersResourceInner) GetCacheFlush() string`

GetCacheFlush returns the CacheFlush field if non-nil, zero value otherwise.

### GetCacheFlushOk

`func (o *ControllerCacheParametersResourceInner) GetCacheFlushOk() (*string, bool)`

GetCacheFlushOk returns a tuple with the CacheFlush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheFlush

`func (o *ControllerCacheParametersResourceInner) SetCacheFlush(v string)`

SetCacheFlush sets CacheFlush field to given value.

### HasCacheFlush

`func (o *ControllerCacheParametersResourceInner) HasCacheFlush() bool`

HasCacheFlush returns a boolean if a field has been set.

### GetCacheFlushNumeric

<<<<<<< HEAD
`func (o *ControllerCacheParametersResourceInner) GetCacheFlushNumeric() int64`
=======
`func (o *ControllerCacheParametersResourceInner) GetCacheFlushNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCacheFlushNumeric returns the CacheFlushNumeric field if non-nil, zero value otherwise.

### GetCacheFlushNumericOk

<<<<<<< HEAD
`func (o *ControllerCacheParametersResourceInner) GetCacheFlushNumericOk() (*int64, bool)`
=======
`func (o *ControllerCacheParametersResourceInner) GetCacheFlushNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCacheFlushNumericOk returns a tuple with the CacheFlushNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheFlushNumeric

<<<<<<< HEAD
`func (o *ControllerCacheParametersResourceInner) SetCacheFlushNumeric(v int64)`
=======
`func (o *ControllerCacheParametersResourceInner) SetCacheFlushNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetCacheFlushNumeric sets CacheFlushNumeric field to given value.

### HasCacheFlushNumeric

`func (o *ControllerCacheParametersResourceInner) HasCacheFlushNumeric() bool`

HasCacheFlushNumeric returns a boolean if a field has been set.

### GetControllerId

`func (o *ControllerCacheParametersResourceInner) GetControllerId() string`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *ControllerCacheParametersResourceInner) GetControllerIdOk() (*string, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *ControllerCacheParametersResourceInner) SetControllerId(v string)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *ControllerCacheParametersResourceInner) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### GetControllerIdNumeric

<<<<<<< HEAD
`func (o *ControllerCacheParametersResourceInner) GetControllerIdNumeric() int64`
=======
`func (o *ControllerCacheParametersResourceInner) GetControllerIdNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetControllerIdNumeric returns the ControllerIdNumeric field if non-nil, zero value otherwise.

### GetControllerIdNumericOk

<<<<<<< HEAD
`func (o *ControllerCacheParametersResourceInner) GetControllerIdNumericOk() (*int64, bool)`
=======
`func (o *ControllerCacheParametersResourceInner) GetControllerIdNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetControllerIdNumericOk returns a tuple with the ControllerIdNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerIdNumeric

<<<<<<< HEAD
`func (o *ControllerCacheParametersResourceInner) SetControllerIdNumeric(v int64)`
=======
`func (o *ControllerCacheParametersResourceInner) SetControllerIdNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetControllerIdNumeric sets ControllerIdNumeric field to given value.

### HasControllerIdNumeric

`func (o *ControllerCacheParametersResourceInner) HasControllerIdNumeric() bool`

HasControllerIdNumeric returns a boolean if a field has been set.

### GetDurableId

`func (o *ControllerCacheParametersResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *ControllerCacheParametersResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *ControllerCacheParametersResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *ControllerCacheParametersResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetMemoryCardHealth

`func (o *ControllerCacheParametersResourceInner) GetMemoryCardHealth() string`

GetMemoryCardHealth returns the MemoryCardHealth field if non-nil, zero value otherwise.

### GetMemoryCardHealthOk

`func (o *ControllerCacheParametersResourceInner) GetMemoryCardHealthOk() (*string, bool)`

GetMemoryCardHealthOk returns a tuple with the MemoryCardHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCardHealth

`func (o *ControllerCacheParametersResourceInner) SetMemoryCardHealth(v string)`

SetMemoryCardHealth sets MemoryCardHealth field to given value.

### HasMemoryCardHealth

`func (o *ControllerCacheParametersResourceInner) HasMemoryCardHealth() bool`

HasMemoryCardHealth returns a boolean if a field has been set.

### GetMemoryCardHealthNumeric

<<<<<<< HEAD
`func (o *ControllerCacheParametersResourceInner) GetMemoryCardHealthNumeric() int64`
=======
`func (o *ControllerCacheParametersResourceInner) GetMemoryCardHealthNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMemoryCardHealthNumeric returns the MemoryCardHealthNumeric field if non-nil, zero value otherwise.

### GetMemoryCardHealthNumericOk

<<<<<<< HEAD
`func (o *ControllerCacheParametersResourceInner) GetMemoryCardHealthNumericOk() (*int64, bool)`
=======
`func (o *ControllerCacheParametersResourceInner) GetMemoryCardHealthNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMemoryCardHealthNumericOk returns a tuple with the MemoryCardHealthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCardHealthNumeric

<<<<<<< HEAD
`func (o *ControllerCacheParametersResourceInner) SetMemoryCardHealthNumeric(v int64)`
=======
`func (o *ControllerCacheParametersResourceInner) SetMemoryCardHealthNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetMemoryCardHealthNumeric sets MemoryCardHealthNumeric field to given value.

### HasMemoryCardHealthNumeric

`func (o *ControllerCacheParametersResourceInner) HasMemoryCardHealthNumeric() bool`

HasMemoryCardHealthNumeric returns a boolean if a field has been set.

### GetMemoryCardStatus

`func (o *ControllerCacheParametersResourceInner) GetMemoryCardStatus() string`

GetMemoryCardStatus returns the MemoryCardStatus field if non-nil, zero value otherwise.

### GetMemoryCardStatusOk

`func (o *ControllerCacheParametersResourceInner) GetMemoryCardStatusOk() (*string, bool)`

GetMemoryCardStatusOk returns a tuple with the MemoryCardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCardStatus

`func (o *ControllerCacheParametersResourceInner) SetMemoryCardStatus(v string)`

SetMemoryCardStatus sets MemoryCardStatus field to given value.

### HasMemoryCardStatus

`func (o *ControllerCacheParametersResourceInner) HasMemoryCardStatus() bool`

HasMemoryCardStatus returns a boolean if a field has been set.

### GetMemoryCardStatusNumeric

<<<<<<< HEAD
`func (o *ControllerCacheParametersResourceInner) GetMemoryCardStatusNumeric() int64`
=======
`func (o *ControllerCacheParametersResourceInner) GetMemoryCardStatusNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMemoryCardStatusNumeric returns the MemoryCardStatusNumeric field if non-nil, zero value otherwise.

### GetMemoryCardStatusNumericOk

<<<<<<< HEAD
`func (o *ControllerCacheParametersResourceInner) GetMemoryCardStatusNumericOk() (*int64, bool)`
=======
`func (o *ControllerCacheParametersResourceInner) GetMemoryCardStatusNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMemoryCardStatusNumericOk returns a tuple with the MemoryCardStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCardStatusNumeric

<<<<<<< HEAD
`func (o *ControllerCacheParametersResourceInner) SetMemoryCardStatusNumeric(v int64)`
=======
`func (o *ControllerCacheParametersResourceInner) SetMemoryCardStatusNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetMemoryCardStatusNumeric sets MemoryCardStatusNumeric field to given value.

### HasMemoryCardStatusNumeric

`func (o *ControllerCacheParametersResourceInner) HasMemoryCardStatusNumeric() bool`

HasMemoryCardStatusNumeric returns a boolean if a field has been set.

### GetName

`func (o *ControllerCacheParametersResourceInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllerCacheParametersResourceInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllerCacheParametersResourceInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllerCacheParametersResourceInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWriteBackStatus

`func (o *ControllerCacheParametersResourceInner) GetWriteBackStatus() string`

GetWriteBackStatus returns the WriteBackStatus field if non-nil, zero value otherwise.

### GetWriteBackStatusOk

`func (o *ControllerCacheParametersResourceInner) GetWriteBackStatusOk() (*string, bool)`

GetWriteBackStatusOk returns a tuple with the WriteBackStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackStatus

`func (o *ControllerCacheParametersResourceInner) SetWriteBackStatus(v string)`

SetWriteBackStatus sets WriteBackStatus field to given value.

### HasWriteBackStatus

`func (o *ControllerCacheParametersResourceInner) HasWriteBackStatus() bool`

HasWriteBackStatus returns a boolean if a field has been set.

### GetWriteBackStatusNumeric

<<<<<<< HEAD
`func (o *ControllerCacheParametersResourceInner) GetWriteBackStatusNumeric() int64`
=======
`func (o *ControllerCacheParametersResourceInner) GetWriteBackStatusNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetWriteBackStatusNumeric returns the WriteBackStatusNumeric field if non-nil, zero value otherwise.

### GetWriteBackStatusNumericOk

<<<<<<< HEAD
`func (o *ControllerCacheParametersResourceInner) GetWriteBackStatusNumericOk() (*int64, bool)`
=======
`func (o *ControllerCacheParametersResourceInner) GetWriteBackStatusNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetWriteBackStatusNumericOk returns a tuple with the WriteBackStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteBackStatusNumeric

<<<<<<< HEAD
`func (o *ControllerCacheParametersResourceInner) SetWriteBackStatusNumeric(v int64)`
=======
`func (o *ControllerCacheParametersResourceInner) SetWriteBackStatusNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetWriteBackStatusNumeric sets WriteBackStatusNumeric field to given value.

### HasWriteBackStatusNumeric

`func (o *ControllerCacheParametersResourceInner) HasWriteBackStatusNumeric() bool`

HasWriteBackStatusNumeric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


