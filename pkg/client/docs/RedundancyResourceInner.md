# RedundancyResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**ControllerASerialNumber** | Pointer to **string** |  | [optional] 
**ControllerAStatus** | Pointer to **string** |  | [optional] 
<<<<<<< HEAD
**ControllerAStatusNumeric** | Pointer to **int64** |  | [optional] 
**ControllerBSerialNumber** | Pointer to **string** |  | [optional] 
**ControllerBStatus** | Pointer to **string** |  | [optional] 
**ControllerBStatusNumeric** | Pointer to **int64** |  | [optional] 
**LocalReady** | Pointer to **string** |  | [optional] 
**LocalReadyNumeric** | Pointer to **int64** |  | [optional] 
**LocalReason** | Pointer to **string** |  | [optional] 
**OtherMCStatus** | Pointer to **string** | Identifies the availability of the partner MC | [optional] 
**OtherMCStatusNumeric** | Pointer to **int64** | Identifies the availability of the partner MC( In numeric form ) | [optional] 
**OtherReady** | Pointer to **string** |  | [optional] 
**OtherReadyNumeric** | Pointer to **int64** |  | [optional] 
**OtherReason** | Pointer to **string** |  | [optional] 
**RedundancyMode** | Pointer to **string** | Mode in which the controllers are operating | [optional] 
**RedundancyModeNumeric** | Pointer to **int64** | Mode in which the controllers are operating( In numeric form ) | [optional] 
**RedundancyStatus** | Pointer to **string** | Current operational state of the controllers | [optional] 
**RedundancyStatusNumeric** | Pointer to **int64** | Current operational state of the controllers( In numeric form ) | [optional] 
**SystemReady** | Pointer to **string** |  | [optional] 
**SystemReadyNumeric** | Pointer to **int64** |  | [optional] 
=======
**ControllerAStatusNumeric** | Pointer to **int32** |  | [optional] 
**ControllerBSerialNumber** | Pointer to **string** |  | [optional] 
**ControllerBStatus** | Pointer to **string** |  | [optional] 
**ControllerBStatusNumeric** | Pointer to **int32** |  | [optional] 
**LocalReady** | Pointer to **string** |  | [optional] 
**LocalReadyNumeric** | Pointer to **int32** |  | [optional] 
**LocalReason** | Pointer to **string** |  | [optional] 
**OtherMCStatus** | Pointer to **string** | Identifies the availability of the partner MC | [optional] 
**OtherMCStatusNumeric** | Pointer to **int32** | Identifies the availability of the partner MC( In numeric form ) | [optional] 
**OtherReady** | Pointer to **string** |  | [optional] 
**OtherReadyNumeric** | Pointer to **int32** |  | [optional] 
**OtherReason** | Pointer to **string** |  | [optional] 
**RedundancyMode** | Pointer to **string** | Mode in which the controllers are operating | [optional] 
**RedundancyModeNumeric** | Pointer to **int32** | Mode in which the controllers are operating( In numeric form ) | [optional] 
**RedundancyStatus** | Pointer to **string** | Current operational state of the controllers | [optional] 
**RedundancyStatusNumeric** | Pointer to **int32** | Current operational state of the controllers( In numeric form ) | [optional] 
**SystemReady** | Pointer to **string** |  | [optional] 
**SystemReadyNumeric** | Pointer to **int32** |  | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

## Methods

### NewRedundancyResourceInner

`func NewRedundancyResourceInner() *RedundancyResourceInner`

NewRedundancyResourceInner instantiates a new RedundancyResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedundancyResourceInnerWithDefaults

`func NewRedundancyResourceInnerWithDefaults() *RedundancyResourceInner`

NewRedundancyResourceInnerWithDefaults instantiates a new RedundancyResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *RedundancyResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *RedundancyResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *RedundancyResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *RedundancyResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *RedundancyResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RedundancyResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RedundancyResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RedundancyResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetControllerASerialNumber

`func (o *RedundancyResourceInner) GetControllerASerialNumber() string`

GetControllerASerialNumber returns the ControllerASerialNumber field if non-nil, zero value otherwise.

### GetControllerASerialNumberOk

`func (o *RedundancyResourceInner) GetControllerASerialNumberOk() (*string, bool)`

GetControllerASerialNumberOk returns a tuple with the ControllerASerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerASerialNumber

`func (o *RedundancyResourceInner) SetControllerASerialNumber(v string)`

SetControllerASerialNumber sets ControllerASerialNumber field to given value.

### HasControllerASerialNumber

`func (o *RedundancyResourceInner) HasControllerASerialNumber() bool`

HasControllerASerialNumber returns a boolean if a field has been set.

### GetControllerAStatus

`func (o *RedundancyResourceInner) GetControllerAStatus() string`

GetControllerAStatus returns the ControllerAStatus field if non-nil, zero value otherwise.

### GetControllerAStatusOk

`func (o *RedundancyResourceInner) GetControllerAStatusOk() (*string, bool)`

GetControllerAStatusOk returns a tuple with the ControllerAStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerAStatus

`func (o *RedundancyResourceInner) SetControllerAStatus(v string)`

SetControllerAStatus sets ControllerAStatus field to given value.

### HasControllerAStatus

`func (o *RedundancyResourceInner) HasControllerAStatus() bool`

HasControllerAStatus returns a boolean if a field has been set.

### GetControllerAStatusNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetControllerAStatusNumeric() int64`
=======
`func (o *RedundancyResourceInner) GetControllerAStatusNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetControllerAStatusNumeric returns the ControllerAStatusNumeric field if non-nil, zero value otherwise.

### GetControllerAStatusNumericOk

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetControllerAStatusNumericOk() (*int64, bool)`
=======
`func (o *RedundancyResourceInner) GetControllerAStatusNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetControllerAStatusNumericOk returns a tuple with the ControllerAStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerAStatusNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) SetControllerAStatusNumeric(v int64)`
=======
`func (o *RedundancyResourceInner) SetControllerAStatusNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetControllerAStatusNumeric sets ControllerAStatusNumeric field to given value.

### HasControllerAStatusNumeric

`func (o *RedundancyResourceInner) HasControllerAStatusNumeric() bool`

HasControllerAStatusNumeric returns a boolean if a field has been set.

### GetControllerBSerialNumber

`func (o *RedundancyResourceInner) GetControllerBSerialNumber() string`

GetControllerBSerialNumber returns the ControllerBSerialNumber field if non-nil, zero value otherwise.

### GetControllerBSerialNumberOk

`func (o *RedundancyResourceInner) GetControllerBSerialNumberOk() (*string, bool)`

GetControllerBSerialNumberOk returns a tuple with the ControllerBSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerBSerialNumber

`func (o *RedundancyResourceInner) SetControllerBSerialNumber(v string)`

SetControllerBSerialNumber sets ControllerBSerialNumber field to given value.

### HasControllerBSerialNumber

`func (o *RedundancyResourceInner) HasControllerBSerialNumber() bool`

HasControllerBSerialNumber returns a boolean if a field has been set.

### GetControllerBStatus

`func (o *RedundancyResourceInner) GetControllerBStatus() string`

GetControllerBStatus returns the ControllerBStatus field if non-nil, zero value otherwise.

### GetControllerBStatusOk

`func (o *RedundancyResourceInner) GetControllerBStatusOk() (*string, bool)`

GetControllerBStatusOk returns a tuple with the ControllerBStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerBStatus

`func (o *RedundancyResourceInner) SetControllerBStatus(v string)`

SetControllerBStatus sets ControllerBStatus field to given value.

### HasControllerBStatus

`func (o *RedundancyResourceInner) HasControllerBStatus() bool`

HasControllerBStatus returns a boolean if a field has been set.

### GetControllerBStatusNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetControllerBStatusNumeric() int64`
=======
`func (o *RedundancyResourceInner) GetControllerBStatusNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetControllerBStatusNumeric returns the ControllerBStatusNumeric field if non-nil, zero value otherwise.

### GetControllerBStatusNumericOk

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetControllerBStatusNumericOk() (*int64, bool)`
=======
`func (o *RedundancyResourceInner) GetControllerBStatusNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetControllerBStatusNumericOk returns a tuple with the ControllerBStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerBStatusNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) SetControllerBStatusNumeric(v int64)`
=======
`func (o *RedundancyResourceInner) SetControllerBStatusNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetControllerBStatusNumeric sets ControllerBStatusNumeric field to given value.

### HasControllerBStatusNumeric

`func (o *RedundancyResourceInner) HasControllerBStatusNumeric() bool`

HasControllerBStatusNumeric returns a boolean if a field has been set.

### GetLocalReady

`func (o *RedundancyResourceInner) GetLocalReady() string`

GetLocalReady returns the LocalReady field if non-nil, zero value otherwise.

### GetLocalReadyOk

`func (o *RedundancyResourceInner) GetLocalReadyOk() (*string, bool)`

GetLocalReadyOk returns a tuple with the LocalReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalReady

`func (o *RedundancyResourceInner) SetLocalReady(v string)`

SetLocalReady sets LocalReady field to given value.

### HasLocalReady

`func (o *RedundancyResourceInner) HasLocalReady() bool`

HasLocalReady returns a boolean if a field has been set.

### GetLocalReadyNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetLocalReadyNumeric() int64`
=======
`func (o *RedundancyResourceInner) GetLocalReadyNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLocalReadyNumeric returns the LocalReadyNumeric field if non-nil, zero value otherwise.

### GetLocalReadyNumericOk

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetLocalReadyNumericOk() (*int64, bool)`
=======
`func (o *RedundancyResourceInner) GetLocalReadyNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLocalReadyNumericOk returns a tuple with the LocalReadyNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalReadyNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) SetLocalReadyNumeric(v int64)`
=======
`func (o *RedundancyResourceInner) SetLocalReadyNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetLocalReadyNumeric sets LocalReadyNumeric field to given value.

### HasLocalReadyNumeric

`func (o *RedundancyResourceInner) HasLocalReadyNumeric() bool`

HasLocalReadyNumeric returns a boolean if a field has been set.

### GetLocalReason

`func (o *RedundancyResourceInner) GetLocalReason() string`

GetLocalReason returns the LocalReason field if non-nil, zero value otherwise.

### GetLocalReasonOk

`func (o *RedundancyResourceInner) GetLocalReasonOk() (*string, bool)`

GetLocalReasonOk returns a tuple with the LocalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalReason

`func (o *RedundancyResourceInner) SetLocalReason(v string)`

SetLocalReason sets LocalReason field to given value.

### HasLocalReason

`func (o *RedundancyResourceInner) HasLocalReason() bool`

HasLocalReason returns a boolean if a field has been set.

### GetOtherMCStatus

`func (o *RedundancyResourceInner) GetOtherMCStatus() string`

GetOtherMCStatus returns the OtherMCStatus field if non-nil, zero value otherwise.

### GetOtherMCStatusOk

`func (o *RedundancyResourceInner) GetOtherMCStatusOk() (*string, bool)`

GetOtherMCStatusOk returns a tuple with the OtherMCStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherMCStatus

`func (o *RedundancyResourceInner) SetOtherMCStatus(v string)`

SetOtherMCStatus sets OtherMCStatus field to given value.

### HasOtherMCStatus

`func (o *RedundancyResourceInner) HasOtherMCStatus() bool`

HasOtherMCStatus returns a boolean if a field has been set.

### GetOtherMCStatusNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetOtherMCStatusNumeric() int64`
=======
`func (o *RedundancyResourceInner) GetOtherMCStatusNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetOtherMCStatusNumeric returns the OtherMCStatusNumeric field if non-nil, zero value otherwise.

### GetOtherMCStatusNumericOk

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetOtherMCStatusNumericOk() (*int64, bool)`
=======
`func (o *RedundancyResourceInner) GetOtherMCStatusNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetOtherMCStatusNumericOk returns a tuple with the OtherMCStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherMCStatusNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) SetOtherMCStatusNumeric(v int64)`
=======
`func (o *RedundancyResourceInner) SetOtherMCStatusNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetOtherMCStatusNumeric sets OtherMCStatusNumeric field to given value.

### HasOtherMCStatusNumeric

`func (o *RedundancyResourceInner) HasOtherMCStatusNumeric() bool`

HasOtherMCStatusNumeric returns a boolean if a field has been set.

### GetOtherReady

`func (o *RedundancyResourceInner) GetOtherReady() string`

GetOtherReady returns the OtherReady field if non-nil, zero value otherwise.

### GetOtherReadyOk

`func (o *RedundancyResourceInner) GetOtherReadyOk() (*string, bool)`

GetOtherReadyOk returns a tuple with the OtherReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherReady

`func (o *RedundancyResourceInner) SetOtherReady(v string)`

SetOtherReady sets OtherReady field to given value.

### HasOtherReady

`func (o *RedundancyResourceInner) HasOtherReady() bool`

HasOtherReady returns a boolean if a field has been set.

### GetOtherReadyNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetOtherReadyNumeric() int64`
=======
`func (o *RedundancyResourceInner) GetOtherReadyNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetOtherReadyNumeric returns the OtherReadyNumeric field if non-nil, zero value otherwise.

### GetOtherReadyNumericOk

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetOtherReadyNumericOk() (*int64, bool)`
=======
`func (o *RedundancyResourceInner) GetOtherReadyNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetOtherReadyNumericOk returns a tuple with the OtherReadyNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherReadyNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) SetOtherReadyNumeric(v int64)`
=======
`func (o *RedundancyResourceInner) SetOtherReadyNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetOtherReadyNumeric sets OtherReadyNumeric field to given value.

### HasOtherReadyNumeric

`func (o *RedundancyResourceInner) HasOtherReadyNumeric() bool`

HasOtherReadyNumeric returns a boolean if a field has been set.

### GetOtherReason

`func (o *RedundancyResourceInner) GetOtherReason() string`

GetOtherReason returns the OtherReason field if non-nil, zero value otherwise.

### GetOtherReasonOk

`func (o *RedundancyResourceInner) GetOtherReasonOk() (*string, bool)`

GetOtherReasonOk returns a tuple with the OtherReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherReason

`func (o *RedundancyResourceInner) SetOtherReason(v string)`

SetOtherReason sets OtherReason field to given value.

### HasOtherReason

`func (o *RedundancyResourceInner) HasOtherReason() bool`

HasOtherReason returns a boolean if a field has been set.

### GetRedundancyMode

`func (o *RedundancyResourceInner) GetRedundancyMode() string`

GetRedundancyMode returns the RedundancyMode field if non-nil, zero value otherwise.

### GetRedundancyModeOk

`func (o *RedundancyResourceInner) GetRedundancyModeOk() (*string, bool)`

GetRedundancyModeOk returns a tuple with the RedundancyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyMode

`func (o *RedundancyResourceInner) SetRedundancyMode(v string)`

SetRedundancyMode sets RedundancyMode field to given value.

### HasRedundancyMode

`func (o *RedundancyResourceInner) HasRedundancyMode() bool`

HasRedundancyMode returns a boolean if a field has been set.

### GetRedundancyModeNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetRedundancyModeNumeric() int64`
=======
`func (o *RedundancyResourceInner) GetRedundancyModeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRedundancyModeNumeric returns the RedundancyModeNumeric field if non-nil, zero value otherwise.

### GetRedundancyModeNumericOk

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetRedundancyModeNumericOk() (*int64, bool)`
=======
`func (o *RedundancyResourceInner) GetRedundancyModeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRedundancyModeNumericOk returns a tuple with the RedundancyModeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyModeNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) SetRedundancyModeNumeric(v int64)`
=======
`func (o *RedundancyResourceInner) SetRedundancyModeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetRedundancyModeNumeric sets RedundancyModeNumeric field to given value.

### HasRedundancyModeNumeric

`func (o *RedundancyResourceInner) HasRedundancyModeNumeric() bool`

HasRedundancyModeNumeric returns a boolean if a field has been set.

### GetRedundancyStatus

`func (o *RedundancyResourceInner) GetRedundancyStatus() string`

GetRedundancyStatus returns the RedundancyStatus field if non-nil, zero value otherwise.

### GetRedundancyStatusOk

`func (o *RedundancyResourceInner) GetRedundancyStatusOk() (*string, bool)`

GetRedundancyStatusOk returns a tuple with the RedundancyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyStatus

`func (o *RedundancyResourceInner) SetRedundancyStatus(v string)`

SetRedundancyStatus sets RedundancyStatus field to given value.

### HasRedundancyStatus

`func (o *RedundancyResourceInner) HasRedundancyStatus() bool`

HasRedundancyStatus returns a boolean if a field has been set.

### GetRedundancyStatusNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetRedundancyStatusNumeric() int64`
=======
`func (o *RedundancyResourceInner) GetRedundancyStatusNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRedundancyStatusNumeric returns the RedundancyStatusNumeric field if non-nil, zero value otherwise.

### GetRedundancyStatusNumericOk

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetRedundancyStatusNumericOk() (*int64, bool)`
=======
`func (o *RedundancyResourceInner) GetRedundancyStatusNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRedundancyStatusNumericOk returns a tuple with the RedundancyStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyStatusNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) SetRedundancyStatusNumeric(v int64)`
=======
`func (o *RedundancyResourceInner) SetRedundancyStatusNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetRedundancyStatusNumeric sets RedundancyStatusNumeric field to given value.

### HasRedundancyStatusNumeric

`func (o *RedundancyResourceInner) HasRedundancyStatusNumeric() bool`

HasRedundancyStatusNumeric returns a boolean if a field has been set.

### GetSystemReady

`func (o *RedundancyResourceInner) GetSystemReady() string`

GetSystemReady returns the SystemReady field if non-nil, zero value otherwise.

### GetSystemReadyOk

`func (o *RedundancyResourceInner) GetSystemReadyOk() (*string, bool)`

GetSystemReadyOk returns a tuple with the SystemReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemReady

`func (o *RedundancyResourceInner) SetSystemReady(v string)`

SetSystemReady sets SystemReady field to given value.

### HasSystemReady

`func (o *RedundancyResourceInner) HasSystemReady() bool`

HasSystemReady returns a boolean if a field has been set.

### GetSystemReadyNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetSystemReadyNumeric() int64`
=======
`func (o *RedundancyResourceInner) GetSystemReadyNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSystemReadyNumeric returns the SystemReadyNumeric field if non-nil, zero value otherwise.

### GetSystemReadyNumericOk

<<<<<<< HEAD
`func (o *RedundancyResourceInner) GetSystemReadyNumericOk() (*int64, bool)`
=======
`func (o *RedundancyResourceInner) GetSystemReadyNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSystemReadyNumericOk returns a tuple with the SystemReadyNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemReadyNumeric

<<<<<<< HEAD
`func (o *RedundancyResourceInner) SetSystemReadyNumeric(v int64)`
=======
`func (o *RedundancyResourceInner) SetSystemReadyNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSystemReadyNumeric sets SystemReadyNumeric field to given value.

### HasSystemReadyNumeric

`func (o *RedundancyResourceInner) HasSystemReadyNumeric() bool`

HasSystemReadyNumeric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


