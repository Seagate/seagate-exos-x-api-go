# ExpandersResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**DomId** | Pointer to **int32** |  | [optional] 
**DrawerId** | Pointer to **int32** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**EnclosureId** | Pointer to **int32** |  | [optional] 
**ExtendedStatus** | Pointer to **string** | Extended status (bits) | [optional] 
**FwRevision** | Pointer to **string** | Firmware version of the FRU | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**HealthNumeric** | Pointer to **int32** |  | [optional] 
**HealthReason** | Pointer to **string** |  | [optional] 
**HealthRecommendation** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PathId** | Pointer to **string** |  | [optional] 
**PathIdNumeric** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusNumeric** | Pointer to **int32** |  | [optional] 

## Methods

### NewExpandersResourceInner

`func NewExpandersResourceInner() *ExpandersResourceInner`

NewExpandersResourceInner instantiates a new ExpandersResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandersResourceInnerWithDefaults

`func NewExpandersResourceInnerWithDefaults() *ExpandersResourceInner`

NewExpandersResourceInnerWithDefaults instantiates a new ExpandersResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *ExpandersResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *ExpandersResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *ExpandersResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *ExpandersResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *ExpandersResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ExpandersResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ExpandersResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ExpandersResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetDomId

`func (o *ExpandersResourceInner) GetDomId() int32`

GetDomId returns the DomId field if non-nil, zero value otherwise.

### GetDomIdOk

`func (o *ExpandersResourceInner) GetDomIdOk() (*int32, bool)`

GetDomIdOk returns a tuple with the DomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomId

`func (o *ExpandersResourceInner) SetDomId(v int32)`

SetDomId sets DomId field to given value.

### HasDomId

`func (o *ExpandersResourceInner) HasDomId() bool`

HasDomId returns a boolean if a field has been set.

### GetDrawerId

`func (o *ExpandersResourceInner) GetDrawerId() int32`

GetDrawerId returns the DrawerId field if non-nil, zero value otherwise.

### GetDrawerIdOk

`func (o *ExpandersResourceInner) GetDrawerIdOk() (*int32, bool)`

GetDrawerIdOk returns a tuple with the DrawerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawerId

`func (o *ExpandersResourceInner) SetDrawerId(v int32)`

SetDrawerId sets DrawerId field to given value.

### HasDrawerId

`func (o *ExpandersResourceInner) HasDrawerId() bool`

HasDrawerId returns a boolean if a field has been set.

### GetDurableId

`func (o *ExpandersResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *ExpandersResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *ExpandersResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *ExpandersResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetEnclosureId

`func (o *ExpandersResourceInner) GetEnclosureId() int32`

GetEnclosureId returns the EnclosureId field if non-nil, zero value otherwise.

### GetEnclosureIdOk

`func (o *ExpandersResourceInner) GetEnclosureIdOk() (*int32, bool)`

GetEnclosureIdOk returns a tuple with the EnclosureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureId

`func (o *ExpandersResourceInner) SetEnclosureId(v int32)`

SetEnclosureId sets EnclosureId field to given value.

### HasEnclosureId

`func (o *ExpandersResourceInner) HasEnclosureId() bool`

HasEnclosureId returns a boolean if a field has been set.

### GetExtendedStatus

`func (o *ExpandersResourceInner) GetExtendedStatus() string`

GetExtendedStatus returns the ExtendedStatus field if non-nil, zero value otherwise.

### GetExtendedStatusOk

`func (o *ExpandersResourceInner) GetExtendedStatusOk() (*string, bool)`

GetExtendedStatusOk returns a tuple with the ExtendedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedStatus

`func (o *ExpandersResourceInner) SetExtendedStatus(v string)`

SetExtendedStatus sets ExtendedStatus field to given value.

### HasExtendedStatus

`func (o *ExpandersResourceInner) HasExtendedStatus() bool`

HasExtendedStatus returns a boolean if a field has been set.

### GetFwRevision

`func (o *ExpandersResourceInner) GetFwRevision() string`

GetFwRevision returns the FwRevision field if non-nil, zero value otherwise.

### GetFwRevisionOk

`func (o *ExpandersResourceInner) GetFwRevisionOk() (*string, bool)`

GetFwRevisionOk returns a tuple with the FwRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwRevision

`func (o *ExpandersResourceInner) SetFwRevision(v string)`

SetFwRevision sets FwRevision field to given value.

### HasFwRevision

`func (o *ExpandersResourceInner) HasFwRevision() bool`

HasFwRevision returns a boolean if a field has been set.

### GetHealth

`func (o *ExpandersResourceInner) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ExpandersResourceInner) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ExpandersResourceInner) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ExpandersResourceInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHealthNumeric

`func (o *ExpandersResourceInner) GetHealthNumeric() int32`

GetHealthNumeric returns the HealthNumeric field if non-nil, zero value otherwise.

### GetHealthNumericOk

`func (o *ExpandersResourceInner) GetHealthNumericOk() (*int32, bool)`

GetHealthNumericOk returns a tuple with the HealthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthNumeric

`func (o *ExpandersResourceInner) SetHealthNumeric(v int32)`

SetHealthNumeric sets HealthNumeric field to given value.

### HasHealthNumeric

`func (o *ExpandersResourceInner) HasHealthNumeric() bool`

HasHealthNumeric returns a boolean if a field has been set.

### GetHealthReason

`func (o *ExpandersResourceInner) GetHealthReason() string`

GetHealthReason returns the HealthReason field if non-nil, zero value otherwise.

### GetHealthReasonOk

`func (o *ExpandersResourceInner) GetHealthReasonOk() (*string, bool)`

GetHealthReasonOk returns a tuple with the HealthReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReason

`func (o *ExpandersResourceInner) SetHealthReason(v string)`

SetHealthReason sets HealthReason field to given value.

### HasHealthReason

`func (o *ExpandersResourceInner) HasHealthReason() bool`

HasHealthReason returns a boolean if a field has been set.

### GetHealthRecommendation

`func (o *ExpandersResourceInner) GetHealthRecommendation() string`

GetHealthRecommendation returns the HealthRecommendation field if non-nil, zero value otherwise.

### GetHealthRecommendationOk

`func (o *ExpandersResourceInner) GetHealthRecommendationOk() (*string, bool)`

GetHealthRecommendationOk returns a tuple with the HealthRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecommendation

`func (o *ExpandersResourceInner) SetHealthRecommendation(v string)`

SetHealthRecommendation sets HealthRecommendation field to given value.

### HasHealthRecommendation

`func (o *ExpandersResourceInner) HasHealthRecommendation() bool`

HasHealthRecommendation returns a boolean if a field has been set.

### GetLocation

`func (o *ExpandersResourceInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ExpandersResourceInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ExpandersResourceInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ExpandersResourceInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *ExpandersResourceInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpandersResourceInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpandersResourceInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExpandersResourceInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPathId

`func (o *ExpandersResourceInner) GetPathId() string`

GetPathId returns the PathId field if non-nil, zero value otherwise.

### GetPathIdOk

`func (o *ExpandersResourceInner) GetPathIdOk() (*string, bool)`

GetPathIdOk returns a tuple with the PathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathId

`func (o *ExpandersResourceInner) SetPathId(v string)`

SetPathId sets PathId field to given value.

### HasPathId

`func (o *ExpandersResourceInner) HasPathId() bool`

HasPathId returns a boolean if a field has been set.

### GetPathIdNumeric

`func (o *ExpandersResourceInner) GetPathIdNumeric() int32`

GetPathIdNumeric returns the PathIdNumeric field if non-nil, zero value otherwise.

### GetPathIdNumericOk

`func (o *ExpandersResourceInner) GetPathIdNumericOk() (*int32, bool)`

GetPathIdNumericOk returns a tuple with the PathIdNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathIdNumeric

`func (o *ExpandersResourceInner) SetPathIdNumeric(v int32)`

SetPathIdNumeric sets PathIdNumeric field to given value.

### HasPathIdNumeric

`func (o *ExpandersResourceInner) HasPathIdNumeric() bool`

HasPathIdNumeric returns a boolean if a field has been set.

### GetStatus

`func (o *ExpandersResourceInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExpandersResourceInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExpandersResourceInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExpandersResourceInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusNumeric

`func (o *ExpandersResourceInner) GetStatusNumeric() int32`

GetStatusNumeric returns the StatusNumeric field if non-nil, zero value otherwise.

### GetStatusNumericOk

`func (o *ExpandersResourceInner) GetStatusNumericOk() (*int32, bool)`

GetStatusNumericOk returns a tuple with the StatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusNumeric

`func (o *ExpandersResourceInner) SetStatusNumeric(v int32)`

SetStatusNumeric sets StatusNumeric field to given value.

### HasStatusNumeric

`func (o *ExpandersResourceInner) HasStatusNumeric() bool`

HasStatusNumeric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


