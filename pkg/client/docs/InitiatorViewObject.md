# InitiatorViewObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**InitiatorView** | Pointer to [**[]InitiatorViewResourceInner**](InitiatorViewResourceInner.md) |  | [optional] 

## Methods

### NewInitiatorViewObject

`func NewInitiatorViewObject() *InitiatorViewObject`

NewInitiatorViewObject instantiates a new InitiatorViewObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiatorViewObjectWithDefaults

`func NewInitiatorViewObjectWithDefaults() *InitiatorViewObject`

NewInitiatorViewObjectWithDefaults instantiates a new InitiatorViewObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InitiatorViewObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InitiatorViewObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InitiatorViewObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InitiatorViewObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInitiatorView

`func (o *InitiatorViewObject) GetInitiatorView() []InitiatorViewResourceInner`

GetInitiatorView returns the InitiatorView field if non-nil, zero value otherwise.

### GetInitiatorViewOk

`func (o *InitiatorViewObject) GetInitiatorViewOk() (*[]InitiatorViewResourceInner, bool)`

GetInitiatorViewOk returns a tuple with the InitiatorView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorView

`func (o *InitiatorViewObject) SetInitiatorView(v []InitiatorViewResourceInner)`

SetInitiatorView sets InitiatorView field to given value.

### HasInitiatorView

`func (o *InitiatorViewObject) HasInitiatorView() bool`

HasInitiatorView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


