# SystemObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**System** | Pointer to [**[]SystemResourceInner**](SystemResourceInner.md) |  | [optional] 

## Methods

### NewSystemObject

`func NewSystemObject() *SystemObject`

NewSystemObject instantiates a new SystemObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemObjectWithDefaults

`func NewSystemObjectWithDefaults() *SystemObject`

NewSystemObjectWithDefaults instantiates a new SystemObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SystemObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SystemObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SystemObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SystemObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *SystemObject) GetSystem() []SystemResourceInner`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *SystemObject) GetSystemOk() (*[]SystemResourceInner, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *SystemObject) SetSystem(v []SystemResourceInner)`

SetSystem sets System field to given value.

### HasSystem

`func (o *SystemObject) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


