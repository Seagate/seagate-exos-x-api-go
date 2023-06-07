# CacheSettingsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**CacheSettings** | Pointer to [**[]CacheSettingsResourceInner**](CacheSettingsResourceInner.md) |  | [optional] 

## Methods

### NewCacheSettingsObject

`func NewCacheSettingsObject() *CacheSettingsObject`

NewCacheSettingsObject instantiates a new CacheSettingsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheSettingsObjectWithDefaults

`func NewCacheSettingsObjectWithDefaults() *CacheSettingsObject`

NewCacheSettingsObjectWithDefaults instantiates a new CacheSettingsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CacheSettingsObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CacheSettingsObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CacheSettingsObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CacheSettingsObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCacheSettings

`func (o *CacheSettingsObject) GetCacheSettings() []CacheSettingsResourceInner`

GetCacheSettings returns the CacheSettings field if non-nil, zero value otherwise.

### GetCacheSettingsOk

`func (o *CacheSettingsObject) GetCacheSettingsOk() (*[]CacheSettingsResourceInner, bool)`

GetCacheSettingsOk returns a tuple with the CacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSettings

`func (o *CacheSettingsObject) SetCacheSettings(v []CacheSettingsResourceInner)`

SetCacheSettings sets CacheSettings field to given value.

### HasCacheSettings

`func (o *CacheSettingsObject) HasCacheSettings() bool`

HasCacheSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


