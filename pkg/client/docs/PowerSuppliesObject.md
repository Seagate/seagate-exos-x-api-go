# PowerSuppliesObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**PowerSupplies** | Pointer to [**[]PowerSuppliesResourceInner**](PowerSuppliesResourceInner.md) |  | [optional] 

## Methods

### NewPowerSuppliesObject

`func NewPowerSuppliesObject() *PowerSuppliesObject`

NewPowerSuppliesObject instantiates a new PowerSuppliesObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerSuppliesObjectWithDefaults

`func NewPowerSuppliesObjectWithDefaults() *PowerSuppliesObject`

NewPowerSuppliesObjectWithDefaults instantiates a new PowerSuppliesObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PowerSuppliesObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PowerSuppliesObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PowerSuppliesObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PowerSuppliesObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPowerSupplies

`func (o *PowerSuppliesObject) GetPowerSupplies() []PowerSuppliesResourceInner`

GetPowerSupplies returns the PowerSupplies field if non-nil, zero value otherwise.

### GetPowerSuppliesOk

`func (o *PowerSuppliesObject) GetPowerSuppliesOk() (*[]PowerSuppliesResourceInner, bool)`

GetPowerSuppliesOk returns a tuple with the PowerSupplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSupplies

`func (o *PowerSuppliesObject) SetPowerSupplies(v []PowerSuppliesResourceInner)`

SetPowerSupplies sets PowerSupplies field to given value.

### HasPowerSupplies

`func (o *PowerSuppliesObject) HasPowerSupplies() bool`

HasPowerSupplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


