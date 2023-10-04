# FanObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**Fan** | Pointer to [**[]FanResourceInner**](FanResourceInner.md) |  | [optional] 

## Methods

### NewFanObject

`func NewFanObject() *FanObject`

NewFanObject instantiates a new FanObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFanObjectWithDefaults

`func NewFanObjectWithDefaults() *FanObject`

NewFanObjectWithDefaults instantiates a new FanObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FanObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FanObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FanObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FanObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFan

`func (o *FanObject) GetFan() []FanResourceInner`

GetFan returns the Fan field if non-nil, zero value otherwise.

### GetFanOk

`func (o *FanObject) GetFanOk() (*[]FanResourceInner, bool)`

GetFanOk returns a tuple with the Fan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFan

`func (o *FanObject) SetFan(v []FanResourceInner)`

SetFan sets Fan field to given value.

### HasFan

`func (o *FanObject) HasFan() bool`

HasFan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


