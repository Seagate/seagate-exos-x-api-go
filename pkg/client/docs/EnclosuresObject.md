# EnclosuresObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**Enclosures** | Pointer to [**[]EnclosuresResourceInner**](EnclosuresResourceInner.md) |  | [optional] 

## Methods

### NewEnclosuresObject

`func NewEnclosuresObject() *EnclosuresObject`

NewEnclosuresObject instantiates a new EnclosuresObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnclosuresObjectWithDefaults

`func NewEnclosuresObjectWithDefaults() *EnclosuresObject`

NewEnclosuresObjectWithDefaults instantiates a new EnclosuresObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EnclosuresObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnclosuresObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnclosuresObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnclosuresObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnclosures

`func (o *EnclosuresObject) GetEnclosures() []EnclosuresResourceInner`

GetEnclosures returns the Enclosures field if non-nil, zero value otherwise.

### GetEnclosuresOk

`func (o *EnclosuresObject) GetEnclosuresOk() (*[]EnclosuresResourceInner, bool)`

GetEnclosuresOk returns a tuple with the Enclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosures

`func (o *EnclosuresObject) SetEnclosures(v []EnclosuresResourceInner)`

SetEnclosures sets Enclosures field to given value.

### HasEnclosures

`func (o *EnclosuresObject) HasEnclosures() bool`

HasEnclosures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


