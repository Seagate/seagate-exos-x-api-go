# HostGroupViewObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**HostGroupView** | Pointer to [**[]HostGroupViewResourceInner**](HostGroupViewResourceInner.md) |  | [optional] 

## Methods

### NewHostGroupViewObject

`func NewHostGroupViewObject() *HostGroupViewObject`

NewHostGroupViewObject instantiates a new HostGroupViewObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostGroupViewObjectWithDefaults

`func NewHostGroupViewObjectWithDefaults() *HostGroupViewObject`

NewHostGroupViewObjectWithDefaults instantiates a new HostGroupViewObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HostGroupViewObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostGroupViewObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostGroupViewObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HostGroupViewObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHostGroupView

`func (o *HostGroupViewObject) GetHostGroupView() []HostGroupViewResourceInner`

GetHostGroupView returns the HostGroupView field if non-nil, zero value otherwise.

### GetHostGroupViewOk

`func (o *HostGroupViewObject) GetHostGroupViewOk() (*[]HostGroupViewResourceInner, bool)`

GetHostGroupViewOk returns a tuple with the HostGroupView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroupView

`func (o *HostGroupViewObject) SetHostGroupView(v []HostGroupViewResourceInner)`

SetHostGroupView sets HostGroupView field to given value.

### HasHostGroupView

`func (o *HostGroupViewObject) HasHostGroupView() bool`

HasHostGroupView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


