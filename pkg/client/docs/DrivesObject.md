# DrivesObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**Drives** | Pointer to [**[]DrivesResourceInner**](DrivesResourceInner.md) |  | [optional] 

## Methods

### NewDrivesObject

`func NewDrivesObject() *DrivesObject`

NewDrivesObject instantiates a new DrivesObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrivesObjectWithDefaults

`func NewDrivesObjectWithDefaults() *DrivesObject`

NewDrivesObjectWithDefaults instantiates a new DrivesObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DrivesObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DrivesObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DrivesObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DrivesObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDrives

`func (o *DrivesObject) GetDrives() []DrivesResourceInner`

GetDrives returns the Drives field if non-nil, zero value otherwise.

### GetDrivesOk

`func (o *DrivesObject) GetDrivesOk() (*[]DrivesResourceInner, bool)`

GetDrivesOk returns a tuple with the Drives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrives

`func (o *DrivesObject) SetDrives(v []DrivesResourceInner)`

SetDrives sets Drives field to given value.

### HasDrives

`func (o *DrivesObject) HasDrives() bool`

HasDrives returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


