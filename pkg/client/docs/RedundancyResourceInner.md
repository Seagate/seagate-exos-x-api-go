# RedundancyResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**ControllerASerialNumber** | Pointer to **string** |  | [optional] 
**ControllerAStatus** | Pointer to **string** |  | [optional] 
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

`func (o *RedundancyResourceInner) GetControllerAStatusNumeric() int32`

GetControllerAStatusNumeric returns the ControllerAStatusNumeric field if non-nil, zero value otherwise.

### GetControllerAStatusNumericOk

`func (o *RedundancyResourceInner) GetControllerAStatusNumericOk() (*int32, bool)`

GetControllerAStatusNumericOk returns a tuple with the ControllerAStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerAStatusNumeric

`func (o *RedundancyResourceInner) SetControllerAStatusNumeric(v int32)`

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

`func (o *RedundancyResourceInner) GetControllerBStatusNumeric() int32`

GetControllerBStatusNumeric returns the ControllerBStatusNumeric field if non-nil, zero value otherwise.

### GetControllerBStatusNumericOk

`func (o *RedundancyResourceInner) GetControllerBStatusNumericOk() (*int32, bool)`

GetControllerBStatusNumericOk returns a tuple with the ControllerBStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerBStatusNumeric

`func (o *RedundancyResourceInner) SetControllerBStatusNumeric(v int32)`

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

`func (o *RedundancyResourceInner) GetLocalReadyNumeric() int32`

GetLocalReadyNumeric returns the LocalReadyNumeric field if non-nil, zero value otherwise.

### GetLocalReadyNumericOk

`func (o *RedundancyResourceInner) GetLocalReadyNumericOk() (*int32, bool)`

GetLocalReadyNumericOk returns a tuple with the LocalReadyNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalReadyNumeric

`func (o *RedundancyResourceInner) SetLocalReadyNumeric(v int32)`

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

`func (o *RedundancyResourceInner) GetOtherMCStatusNumeric() int32`

GetOtherMCStatusNumeric returns the OtherMCStatusNumeric field if non-nil, zero value otherwise.

### GetOtherMCStatusNumericOk

`func (o *RedundancyResourceInner) GetOtherMCStatusNumericOk() (*int32, bool)`

GetOtherMCStatusNumericOk returns a tuple with the OtherMCStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherMCStatusNumeric

`func (o *RedundancyResourceInner) SetOtherMCStatusNumeric(v int32)`

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

`func (o *RedundancyResourceInner) GetOtherReadyNumeric() int32`

GetOtherReadyNumeric returns the OtherReadyNumeric field if non-nil, zero value otherwise.

### GetOtherReadyNumericOk

`func (o *RedundancyResourceInner) GetOtherReadyNumericOk() (*int32, bool)`

GetOtherReadyNumericOk returns a tuple with the OtherReadyNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherReadyNumeric

`func (o *RedundancyResourceInner) SetOtherReadyNumeric(v int32)`

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

`func (o *RedundancyResourceInner) GetRedundancyModeNumeric() int32`

GetRedundancyModeNumeric returns the RedundancyModeNumeric field if non-nil, zero value otherwise.

### GetRedundancyModeNumericOk

`func (o *RedundancyResourceInner) GetRedundancyModeNumericOk() (*int32, bool)`

GetRedundancyModeNumericOk returns a tuple with the RedundancyModeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyModeNumeric

`func (o *RedundancyResourceInner) SetRedundancyModeNumeric(v int32)`

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

`func (o *RedundancyResourceInner) GetRedundancyStatusNumeric() int32`

GetRedundancyStatusNumeric returns the RedundancyStatusNumeric field if non-nil, zero value otherwise.

### GetRedundancyStatusNumericOk

`func (o *RedundancyResourceInner) GetRedundancyStatusNumericOk() (*int32, bool)`

GetRedundancyStatusNumericOk returns a tuple with the RedundancyStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancyStatusNumeric

`func (o *RedundancyResourceInner) SetRedundancyStatusNumeric(v int32)`

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

`func (o *RedundancyResourceInner) GetSystemReadyNumeric() int32`

GetSystemReadyNumeric returns the SystemReadyNumeric field if non-nil, zero value otherwise.

### GetSystemReadyNumericOk

`func (o *RedundancyResourceInner) GetSystemReadyNumericOk() (*int32, bool)`

GetSystemReadyNumericOk returns a tuple with the SystemReadyNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemReadyNumeric

`func (o *RedundancyResourceInner) SetSystemReadyNumeric(v int32)`

SetSystemReadyNumeric sets SystemReadyNumeric field to given value.

### HasSystemReadyNumeric

`func (o *RedundancyResourceInner) HasSystemReadyNumeric() bool`

HasSystemReadyNumeric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


