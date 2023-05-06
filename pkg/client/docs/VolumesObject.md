# VolumesObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**Volumes** | Pointer to [**[]VolumesResourceInner**](VolumesResourceInner.md) |  | [optional] 

## Methods

### NewVolumesObject

`func NewVolumesObject() *VolumesObject`

NewVolumesObject instantiates a new VolumesObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumesObjectWithDefaults

`func NewVolumesObjectWithDefaults() *VolumesObject`

NewVolumesObjectWithDefaults instantiates a new VolumesObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VolumesObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VolumesObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VolumesObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VolumesObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVolumes

`func (o *VolumesObject) GetVolumes() []VolumesResourceInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *VolumesObject) GetVolumesOk() (*[]VolumesResourceInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *VolumesObject) SetVolumes(v []VolumesResourceInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *VolumesObject) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


