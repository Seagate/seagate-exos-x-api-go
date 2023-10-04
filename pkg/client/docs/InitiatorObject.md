# InitiatorObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**Initiator** | Pointer to [**[]InitiatorResourceInner**](InitiatorResourceInner.md) |  | [optional] 

## Methods

### NewInitiatorObject

`func NewInitiatorObject() *InitiatorObject`

NewInitiatorObject instantiates a new InitiatorObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiatorObjectWithDefaults

`func NewInitiatorObjectWithDefaults() *InitiatorObject`

NewInitiatorObjectWithDefaults instantiates a new InitiatorObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InitiatorObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InitiatorObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InitiatorObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InitiatorObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInitiator

`func (o *InitiatorObject) GetInitiator() []InitiatorResourceInner`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *InitiatorObject) GetInitiatorOk() (*[]InitiatorResourceInner, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *InitiatorObject) SetInitiator(v []InitiatorResourceInner)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *InitiatorObject) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


