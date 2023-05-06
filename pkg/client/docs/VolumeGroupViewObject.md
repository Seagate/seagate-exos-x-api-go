# VolumeGroupViewObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**VolumeGroupView** | Pointer to [**[]VolumeGroupViewResourceInner**](VolumeGroupViewResourceInner.md) |  | [optional] 

## Methods

### NewVolumeGroupViewObject

`func NewVolumeGroupViewObject() *VolumeGroupViewObject`

NewVolumeGroupViewObject instantiates a new VolumeGroupViewObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupViewObjectWithDefaults

`func NewVolumeGroupViewObjectWithDefaults() *VolumeGroupViewObject`

NewVolumeGroupViewObjectWithDefaults instantiates a new VolumeGroupViewObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VolumeGroupViewObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumeGroupViewObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumeGroupViewObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumeGroupViewObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVolumeGroupView

`func (o *VolumeGroupViewObject) GetVolumeGroupView() []VolumeGroupViewResourceInner`

GetVolumeGroupView returns the VolumeGroupView field if non-nil, zero value otherwise.

### GetVolumeGroupViewOk

`func (o *VolumeGroupViewObject) GetVolumeGroupViewOk() (*[]VolumeGroupViewResourceInner, bool)`

GetVolumeGroupViewOk returns a tuple with the VolumeGroupView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupView

`func (o *VolumeGroupViewObject) SetVolumeGroupView(v []VolumeGroupViewResourceInner)`

SetVolumeGroupView sets VolumeGroupView field to given value.

### HasVolumeGroupView

`func (o *VolumeGroupViewObject) HasVolumeGroupView() bool`

HasVolumeGroupView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


