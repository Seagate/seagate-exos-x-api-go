# ExpanderPortsResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**Controller** | Pointer to **string** |  | [optional] 
**ControllerNumeric** | Pointer to **int64** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**EnclosureId** | Pointer to **int64** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**HealthNumeric** | Pointer to **int64** |  | [optional] 
**HealthReason** | Pointer to **string** |  | [optional] 
**HealthRecommendation** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SasPortIndex** | Pointer to **int64** | Port number for this SAS port | [optional] 
**SasPortType** | Pointer to **string** | The type of SAS port | [optional] 
**SasPortTypeNumeric** | Pointer to **int64** | The type of SAS port( In numeric form ) | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusNumeric** | Pointer to **int64** |  | [optional] 

## Methods

### NewExpanderPortsResourceInner

`func NewExpanderPortsResourceInner() *ExpanderPortsResourceInner`

NewExpanderPortsResourceInner instantiates a new ExpanderPortsResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpanderPortsResourceInnerWithDefaults

`func NewExpanderPortsResourceInnerWithDefaults() *ExpanderPortsResourceInner`

NewExpanderPortsResourceInnerWithDefaults instantiates a new ExpanderPortsResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *ExpanderPortsResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ExpanderPortsResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ExpanderPortsResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ExpanderPortsResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *ExpanderPortsResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ExpanderPortsResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ExpanderPortsResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ExpanderPortsResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetController

`func (o *ExpanderPortsResourceInner) GetController() string`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *ExpanderPortsResourceInner) GetControllerOk() (*string, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *ExpanderPortsResourceInner) SetController(v string)`

SetController sets Controller field to given value.

### HasController

`func (o *ExpanderPortsResourceInner) HasController() bool`

HasController returns a boolean if a field has been set.

### GetControllerNumeric

`func (o *ExpanderPortsResourceInner) GetControllerNumeric() int64`

GetControllerNumeric returns the ControllerNumeric field if non-nil, zero value otherwise.

### GetControllerNumericOk

`func (o *ExpanderPortsResourceInner) GetControllerNumericOk() (*int64, bool)`

GetControllerNumericOk returns a tuple with the ControllerNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerNumeric

`func (o *ExpanderPortsResourceInner) SetControllerNumeric(v int64)`

SetControllerNumeric sets ControllerNumeric field to given value.

### HasControllerNumeric

`func (o *ExpanderPortsResourceInner) HasControllerNumeric() bool`

HasControllerNumeric returns a boolean if a field has been set.

### GetDurableId

`func (o *ExpanderPortsResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *ExpanderPortsResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *ExpanderPortsResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *ExpanderPortsResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetEnclosureId

`func (o *ExpanderPortsResourceInner) GetEnclosureId() int64`

GetEnclosureId returns the EnclosureId field if non-nil, zero value otherwise.

### GetEnclosureIdOk

`func (o *ExpanderPortsResourceInner) GetEnclosureIdOk() (*int64, bool)`

GetEnclosureIdOk returns a tuple with the EnclosureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureId

`func (o *ExpanderPortsResourceInner) SetEnclosureId(v int64)`

SetEnclosureId sets EnclosureId field to given value.

### HasEnclosureId

`func (o *ExpanderPortsResourceInner) HasEnclosureId() bool`

HasEnclosureId returns a boolean if a field has been set.

### GetHealth

`func (o *ExpanderPortsResourceInner) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ExpanderPortsResourceInner) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ExpanderPortsResourceInner) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ExpanderPortsResourceInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHealthNumeric

`func (o *ExpanderPortsResourceInner) GetHealthNumeric() int64`

GetHealthNumeric returns the HealthNumeric field if non-nil, zero value otherwise.

### GetHealthNumericOk

`func (o *ExpanderPortsResourceInner) GetHealthNumericOk() (*int64, bool)`

GetHealthNumericOk returns a tuple with the HealthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthNumeric

`func (o *ExpanderPortsResourceInner) SetHealthNumeric(v int64)`

SetHealthNumeric sets HealthNumeric field to given value.

### HasHealthNumeric

`func (o *ExpanderPortsResourceInner) HasHealthNumeric() bool`

HasHealthNumeric returns a boolean if a field has been set.

### GetHealthReason

`func (o *ExpanderPortsResourceInner) GetHealthReason() string`

GetHealthReason returns the HealthReason field if non-nil, zero value otherwise.

### GetHealthReasonOk

`func (o *ExpanderPortsResourceInner) GetHealthReasonOk() (*string, bool)`

GetHealthReasonOk returns a tuple with the HealthReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReason

`func (o *ExpanderPortsResourceInner) SetHealthReason(v string)`

SetHealthReason sets HealthReason field to given value.

### HasHealthReason

`func (o *ExpanderPortsResourceInner) HasHealthReason() bool`

HasHealthReason returns a boolean if a field has been set.

### GetHealthRecommendation

`func (o *ExpanderPortsResourceInner) GetHealthRecommendation() string`

GetHealthRecommendation returns the HealthRecommendation field if non-nil, zero value otherwise.

### GetHealthRecommendationOk

`func (o *ExpanderPortsResourceInner) GetHealthRecommendationOk() (*string, bool)`

GetHealthRecommendationOk returns a tuple with the HealthRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecommendation

`func (o *ExpanderPortsResourceInner) SetHealthRecommendation(v string)`

SetHealthRecommendation sets HealthRecommendation field to given value.

### HasHealthRecommendation

`func (o *ExpanderPortsResourceInner) HasHealthRecommendation() bool`

HasHealthRecommendation returns a boolean if a field has been set.

### GetName

`func (o *ExpanderPortsResourceInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpanderPortsResourceInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpanderPortsResourceInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExpanderPortsResourceInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSasPortIndex

`func (o *ExpanderPortsResourceInner) GetSasPortIndex() int64`

GetSasPortIndex returns the SasPortIndex field if non-nil, zero value otherwise.

### GetSasPortIndexOk

`func (o *ExpanderPortsResourceInner) GetSasPortIndexOk() (*int64, bool)`

GetSasPortIndexOk returns a tuple with the SasPortIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasPortIndex

`func (o *ExpanderPortsResourceInner) SetSasPortIndex(v int64)`

SetSasPortIndex sets SasPortIndex field to given value.

### HasSasPortIndex

`func (o *ExpanderPortsResourceInner) HasSasPortIndex() bool`

HasSasPortIndex returns a boolean if a field has been set.

### GetSasPortType

`func (o *ExpanderPortsResourceInner) GetSasPortType() string`

GetSasPortType returns the SasPortType field if non-nil, zero value otherwise.

### GetSasPortTypeOk

`func (o *ExpanderPortsResourceInner) GetSasPortTypeOk() (*string, bool)`

GetSasPortTypeOk returns a tuple with the SasPortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasPortType

`func (o *ExpanderPortsResourceInner) SetSasPortType(v string)`

SetSasPortType sets SasPortType field to given value.

### HasSasPortType

`func (o *ExpanderPortsResourceInner) HasSasPortType() bool`

HasSasPortType returns a boolean if a field has been set.

### GetSasPortTypeNumeric

`func (o *ExpanderPortsResourceInner) GetSasPortTypeNumeric() int64`

GetSasPortTypeNumeric returns the SasPortTypeNumeric field if non-nil, zero value otherwise.

### GetSasPortTypeNumericOk

`func (o *ExpanderPortsResourceInner) GetSasPortTypeNumericOk() (*int64, bool)`

GetSasPortTypeNumericOk returns a tuple with the SasPortTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasPortTypeNumeric

`func (o *ExpanderPortsResourceInner) SetSasPortTypeNumeric(v int64)`

SetSasPortTypeNumeric sets SasPortTypeNumeric field to given value.

### HasSasPortTypeNumeric

`func (o *ExpanderPortsResourceInner) HasSasPortTypeNumeric() bool`

HasSasPortTypeNumeric returns a boolean if a field has been set.

### GetStatus

`func (o *ExpanderPortsResourceInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExpanderPortsResourceInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExpanderPortsResourceInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExpanderPortsResourceInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusNumeric

`func (o *ExpanderPortsResourceInner) GetStatusNumeric() int64`

GetStatusNumeric returns the StatusNumeric field if non-nil, zero value otherwise.

### GetStatusNumericOk

`func (o *ExpanderPortsResourceInner) GetStatusNumericOk() (*int64, bool)`

GetStatusNumericOk returns a tuple with the StatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusNumeric

`func (o *ExpanderPortsResourceInner) SetStatusNumeric(v int64)`

SetStatusNumeric sets StatusNumeric field to given value.

### HasStatusNumeric

`func (o *ExpanderPortsResourceInner) HasStatusNumeric() bool`

HasStatusNumeric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


