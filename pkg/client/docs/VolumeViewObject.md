# VolumeViewObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**VolumeGroupView** | Pointer to [**[]VolumeGroupViewResourceInner**](VolumeGroupViewResourceInner.md) |  | [optional] 
**VolumeView** | Pointer to [**[]VolumeViewResourceInner**](VolumeViewResourceInner.md) |  | [optional] 

## Methods

### NewVolumeViewObject

`func NewVolumeViewObject() *VolumeViewObject`

NewVolumeViewObject instantiates a new VolumeViewObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeViewObjectWithDefaults

`func NewVolumeViewObjectWithDefaults() *VolumeViewObject`

NewVolumeViewObjectWithDefaults instantiates a new VolumeViewObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VolumeViewObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumeViewObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumeViewObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumeViewObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVolumeGroupView

`func (o *VolumeViewObject) GetVolumeGroupView() []VolumeGroupViewResourceInner`

GetVolumeGroupView returns the VolumeGroupView field if non-nil, zero value otherwise.

### GetVolumeGroupViewOk

`func (o *VolumeViewObject) GetVolumeGroupViewOk() (*[]VolumeGroupViewResourceInner, bool)`

GetVolumeGroupViewOk returns a tuple with the VolumeGroupView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupView

`func (o *VolumeViewObject) SetVolumeGroupView(v []VolumeGroupViewResourceInner)`

SetVolumeGroupView sets VolumeGroupView field to given value.

### HasVolumeGroupView

`func (o *VolumeViewObject) HasVolumeGroupView() bool`

HasVolumeGroupView returns a boolean if a field has been set.

### GetVolumeView

`func (o *VolumeViewObject) GetVolumeView() []VolumeViewResourceInner`

GetVolumeView returns the VolumeView field if non-nil, zero value otherwise.

### GetVolumeViewOk

`func (o *VolumeViewObject) GetVolumeViewOk() (*[]VolumeViewResourceInner, bool)`

GetVolumeViewOk returns a tuple with the VolumeView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeView

`func (o *VolumeViewObject) SetVolumeView(v []VolumeViewResourceInner)`

SetVolumeView sets VolumeView field to given value.

### HasVolumeView

`func (o *VolumeViewObject) HasVolumeView() bool`

HasVolumeView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


