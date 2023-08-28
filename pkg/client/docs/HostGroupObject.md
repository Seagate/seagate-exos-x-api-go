# HostGroupObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**HostGroup** | Pointer to [**[]HostGroupResourceInner**](HostGroupResourceInner.md) |  | [optional] 

## Methods

### NewHostGroupObject

`func NewHostGroupObject() *HostGroupObject`

NewHostGroupObject instantiates a new HostGroupObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostGroupObjectWithDefaults

`func NewHostGroupObjectWithDefaults() *HostGroupObject`

NewHostGroupObjectWithDefaults instantiates a new HostGroupObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HostGroupObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostGroupObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostGroupObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HostGroupObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHostGroup

`func (o *HostGroupObject) GetHostGroup() []HostGroupResourceInner`

GetHostGroup returns the HostGroup field if non-nil, zero value otherwise.

### GetHostGroupOk

`func (o *HostGroupObject) GetHostGroupOk() (*[]HostGroupResourceInner, bool)`

GetHostGroupOk returns a tuple with the HostGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroup

`func (o *HostGroupObject) SetHostGroup(v []HostGroupResourceInner)`

SetHostGroup sets HostGroup field to given value.

### HasHostGroup

`func (o *HostGroupObject) HasHostGroup() bool`

HasHostGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


