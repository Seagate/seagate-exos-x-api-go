# HostsViewObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**HostsView** | Pointer to [**[]HostsViewResourceInner**](HostsViewResourceInner.md) |  | [optional] 

## Methods

### NewHostsViewObject

`func NewHostsViewObject() *HostsViewObject`

NewHostsViewObject instantiates a new HostsViewObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostsViewObjectWithDefaults

`func NewHostsViewObjectWithDefaults() *HostsViewObject`

NewHostsViewObjectWithDefaults instantiates a new HostsViewObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HostsViewObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostsViewObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostsViewObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HostsViewObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHostsView

`func (o *HostsViewObject) GetHostsView() []HostsViewResourceInner`

GetHostsView returns the HostsView field if non-nil, zero value otherwise.

### GetHostsViewOk

`func (o *HostsViewObject) GetHostsViewOk() (*[]HostsViewResourceInner, bool)`

GetHostsViewOk returns a tuple with the HostsView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsView

`func (o *HostsViewObject) SetHostsView(v []HostsViewResourceInner)`

SetHostsView sets HostsView field to given value.

### HasHostsView

`func (o *HostsViewObject) HasHostsView() bool`

HasHostsView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


