# PoolsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**Pools** | Pointer to [**[]PoolsResourceInner**](PoolsResourceInner.md) |  | [optional] 

## Methods

### NewPoolsObject

`func NewPoolsObject() *PoolsObject`

NewPoolsObject instantiates a new PoolsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolsObjectWithDefaults

`func NewPoolsObjectWithDefaults() *PoolsObject`

NewPoolsObjectWithDefaults instantiates a new PoolsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PoolsObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PoolsObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PoolsObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PoolsObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPools

`func (o *PoolsObject) GetPools() []PoolsResourceInner`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *PoolsObject) GetPoolsOk() (*[]PoolsResourceInner, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *PoolsObject) SetPools(v []PoolsResourceInner)`

SetPools sets Pools field to given value.

### HasPools

`func (o *PoolsObject) HasPools() bool`

HasPools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


