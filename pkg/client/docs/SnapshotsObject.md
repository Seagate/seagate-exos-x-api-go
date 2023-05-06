# SnapshotsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**Snapshots** | Pointer to [**[]SnapshotsResourceInner**](SnapshotsResourceInner.md) |  | [optional] 

## Methods

### NewSnapshotsObject

`func NewSnapshotsObject() *SnapshotsObject`

NewSnapshotsObject instantiates a new SnapshotsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotsObjectWithDefaults

`func NewSnapshotsObjectWithDefaults() *SnapshotsObject`

NewSnapshotsObjectWithDefaults instantiates a new SnapshotsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SnapshotsObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SnapshotsObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SnapshotsObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SnapshotsObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSnapshots

`func (o *SnapshotsObject) GetSnapshots() []SnapshotsResourceInner`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *SnapshotsObject) GetSnapshotsOk() (*[]SnapshotsResourceInner, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *SnapshotsObject) SetSnapshots(v []SnapshotsResourceInner)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *SnapshotsObject) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


