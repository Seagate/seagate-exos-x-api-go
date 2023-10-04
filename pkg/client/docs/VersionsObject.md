# VersionsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**Versions** | Pointer to [**[]VersionsResourceInner**](VersionsResourceInner.md) |  | [optional] 

## Methods

### NewVersionsObject

`func NewVersionsObject() *VersionsObject`

NewVersionsObject instantiates a new VersionsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionsObjectWithDefaults

`func NewVersionsObjectWithDefaults() *VersionsObject`

NewVersionsObjectWithDefaults instantiates a new VersionsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VersionsObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VersionsObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VersionsObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VersionsObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersions

`func (o *VersionsObject) GetVersions() []VersionsResourceInner`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *VersionsObject) GetVersionsOk() (*[]VersionsResourceInner, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *VersionsObject) SetVersions(v []VersionsResourceInner)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *VersionsObject) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


