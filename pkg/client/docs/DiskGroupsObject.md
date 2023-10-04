# DiskGroupsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**DiskGroups** | Pointer to [**[]DiskGroupsResourceInner**](DiskGroupsResourceInner.md) |  | [optional] 

## Methods

### NewDiskGroupsObject

`func NewDiskGroupsObject() *DiskGroupsObject`

NewDiskGroupsObject instantiates a new DiskGroupsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskGroupsObjectWithDefaults

`func NewDiskGroupsObjectWithDefaults() *DiskGroupsObject`

NewDiskGroupsObjectWithDefaults instantiates a new DiskGroupsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DiskGroupsObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiskGroupsObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiskGroupsObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiskGroupsObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDiskGroups

`func (o *DiskGroupsObject) GetDiskGroups() []DiskGroupsResourceInner`

GetDiskGroups returns the DiskGroups field if non-nil, zero value otherwise.

### GetDiskGroupsOk

`func (o *DiskGroupsObject) GetDiskGroupsOk() (*[]DiskGroupsResourceInner, bool)`

GetDiskGroupsOk returns a tuple with the DiskGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroups

`func (o *DiskGroupsObject) SetDiskGroups(v []DiskGroupsResourceInner)`

SetDiskGroups sets DiskGroups field to given value.

### HasDiskGroups

`func (o *DiskGroupsObject) HasDiskGroups() bool`

HasDiskGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


