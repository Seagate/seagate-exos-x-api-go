# ControllersObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**Controllers** | Pointer to [**[]ControllersResourceInner**](ControllersResourceInner.md) |  | [optional] 

## Methods

### NewControllersObject

`func NewControllersObject() *ControllersObject`

NewControllersObject instantiates a new ControllersObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllersObjectWithDefaults

`func NewControllersObjectWithDefaults() *ControllersObject`

NewControllersObjectWithDefaults instantiates a new ControllersObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ControllersObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ControllersObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ControllersObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ControllersObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetControllers

`func (o *ControllersObject) GetControllers() []ControllersResourceInner`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *ControllersObject) GetControllersOk() (*[]ControllersResourceInner, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *ControllersObject) SetControllers(v []ControllersResourceInner)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *ControllersObject) HasControllers() bool`

HasControllers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


