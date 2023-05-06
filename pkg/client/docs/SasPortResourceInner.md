# SasPortResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**ConfiguredTopology** | Pointer to **string** |  | [optional] 
<<<<<<< HEAD
**ConfiguredTopologyNumeric** | Pointer to **int64** |  | [optional] 
**SasActiveLanes** | Pointer to **int64** | Number of currently active PHY lanes for the SAS port | [optional] 
**SasDisabledLanes** | Pointer to **int64** | Number of disabled PHY lanes for the SAS port | [optional] 
**SasLanesExpected** | Pointer to **int64** | Expected number of active PHY lanes for the SAS port | [optional] 
**Width** | Pointer to **int64** | Number of PHY lanes for the SAS port | [optional] 
=======
**ConfiguredTopologyNumeric** | Pointer to **int32** |  | [optional] 
**SasActiveLanes** | Pointer to **int32** | Number of currently active PHY lanes for the SAS port | [optional] 
**SasDisabledLanes** | Pointer to **int32** | Number of disabled PHY lanes for the SAS port | [optional] 
**SasLanesExpected** | Pointer to **int32** | Expected number of active PHY lanes for the SAS port | [optional] 
**Width** | Pointer to **int32** | Number of PHY lanes for the SAS port | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

## Methods

### NewSasPortResourceInner

`func NewSasPortResourceInner() *SasPortResourceInner`

NewSasPortResourceInner instantiates a new SasPortResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSasPortResourceInnerWithDefaults

`func NewSasPortResourceInnerWithDefaults() *SasPortResourceInner`

NewSasPortResourceInnerWithDefaults instantiates a new SasPortResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *SasPortResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *SasPortResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *SasPortResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *SasPortResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *SasPortResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SasPortResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SasPortResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SasPortResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetConfiguredTopology

`func (o *SasPortResourceInner) GetConfiguredTopology() string`

GetConfiguredTopology returns the ConfiguredTopology field if non-nil, zero value otherwise.

### GetConfiguredTopologyOk

`func (o *SasPortResourceInner) GetConfiguredTopologyOk() (*string, bool)`

GetConfiguredTopologyOk returns a tuple with the ConfiguredTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredTopology

`func (o *SasPortResourceInner) SetConfiguredTopology(v string)`

SetConfiguredTopology sets ConfiguredTopology field to given value.

### HasConfiguredTopology

`func (o *SasPortResourceInner) HasConfiguredTopology() bool`

HasConfiguredTopology returns a boolean if a field has been set.

### GetConfiguredTopologyNumeric

<<<<<<< HEAD
`func (o *SasPortResourceInner) GetConfiguredTopologyNumeric() int64`
=======
`func (o *SasPortResourceInner) GetConfiguredTopologyNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetConfiguredTopologyNumeric returns the ConfiguredTopologyNumeric field if non-nil, zero value otherwise.

### GetConfiguredTopologyNumericOk

<<<<<<< HEAD
`func (o *SasPortResourceInner) GetConfiguredTopologyNumericOk() (*int64, bool)`
=======
`func (o *SasPortResourceInner) GetConfiguredTopologyNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetConfiguredTopologyNumericOk returns a tuple with the ConfiguredTopologyNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredTopologyNumeric

<<<<<<< HEAD
`func (o *SasPortResourceInner) SetConfiguredTopologyNumeric(v int64)`
=======
`func (o *SasPortResourceInner) SetConfiguredTopologyNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetConfiguredTopologyNumeric sets ConfiguredTopologyNumeric field to given value.

### HasConfiguredTopologyNumeric

`func (o *SasPortResourceInner) HasConfiguredTopologyNumeric() bool`

HasConfiguredTopologyNumeric returns a boolean if a field has been set.

### GetSasActiveLanes

<<<<<<< HEAD
`func (o *SasPortResourceInner) GetSasActiveLanes() int64`
=======
`func (o *SasPortResourceInner) GetSasActiveLanes() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSasActiveLanes returns the SasActiveLanes field if non-nil, zero value otherwise.

### GetSasActiveLanesOk

<<<<<<< HEAD
`func (o *SasPortResourceInner) GetSasActiveLanesOk() (*int64, bool)`
=======
`func (o *SasPortResourceInner) GetSasActiveLanesOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSasActiveLanesOk returns a tuple with the SasActiveLanes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasActiveLanes

<<<<<<< HEAD
`func (o *SasPortResourceInner) SetSasActiveLanes(v int64)`
=======
`func (o *SasPortResourceInner) SetSasActiveLanes(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSasActiveLanes sets SasActiveLanes field to given value.

### HasSasActiveLanes

`func (o *SasPortResourceInner) HasSasActiveLanes() bool`

HasSasActiveLanes returns a boolean if a field has been set.

### GetSasDisabledLanes

<<<<<<< HEAD
`func (o *SasPortResourceInner) GetSasDisabledLanes() int64`
=======
`func (o *SasPortResourceInner) GetSasDisabledLanes() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSasDisabledLanes returns the SasDisabledLanes field if non-nil, zero value otherwise.

### GetSasDisabledLanesOk

<<<<<<< HEAD
`func (o *SasPortResourceInner) GetSasDisabledLanesOk() (*int64, bool)`
=======
`func (o *SasPortResourceInner) GetSasDisabledLanesOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSasDisabledLanesOk returns a tuple with the SasDisabledLanes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasDisabledLanes

<<<<<<< HEAD
`func (o *SasPortResourceInner) SetSasDisabledLanes(v int64)`
=======
`func (o *SasPortResourceInner) SetSasDisabledLanes(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSasDisabledLanes sets SasDisabledLanes field to given value.

### HasSasDisabledLanes

`func (o *SasPortResourceInner) HasSasDisabledLanes() bool`

HasSasDisabledLanes returns a boolean if a field has been set.

### GetSasLanesExpected

<<<<<<< HEAD
`func (o *SasPortResourceInner) GetSasLanesExpected() int64`
=======
`func (o *SasPortResourceInner) GetSasLanesExpected() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSasLanesExpected returns the SasLanesExpected field if non-nil, zero value otherwise.

### GetSasLanesExpectedOk

<<<<<<< HEAD
`func (o *SasPortResourceInner) GetSasLanesExpectedOk() (*int64, bool)`
=======
`func (o *SasPortResourceInner) GetSasLanesExpectedOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSasLanesExpectedOk returns a tuple with the SasLanesExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasLanesExpected

<<<<<<< HEAD
`func (o *SasPortResourceInner) SetSasLanesExpected(v int64)`
=======
`func (o *SasPortResourceInner) SetSasLanesExpected(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSasLanesExpected sets SasLanesExpected field to given value.

### HasSasLanesExpected

`func (o *SasPortResourceInner) HasSasLanesExpected() bool`

HasSasLanesExpected returns a boolean if a field has been set.

### GetWidth

<<<<<<< HEAD
`func (o *SasPortResourceInner) GetWidth() int64`
=======
`func (o *SasPortResourceInner) GetWidth() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

<<<<<<< HEAD
`func (o *SasPortResourceInner) GetWidthOk() (*int64, bool)`
=======
`func (o *SasPortResourceInner) GetWidthOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

<<<<<<< HEAD
`func (o *SasPortResourceInner) SetWidth(v int64)`
=======
`func (o *SasPortResourceInner) SetWidth(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetWidth sets Width field to given value.

### HasWidth

`func (o *SasPortResourceInner) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


