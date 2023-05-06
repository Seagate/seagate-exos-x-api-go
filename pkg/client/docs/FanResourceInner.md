# FanResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**ExtendedStatus** | Pointer to **string** | Extended status (bits) | [optional] 
**FwRevision** | Pointer to **string** | Firmware version of the FRU | [optional] 
**Health** | Pointer to **string** |  | [optional] 
<<<<<<< HEAD
**HealthNumeric** | Pointer to **int64** |  | [optional] 
=======
**HealthNumeric** | Pointer to **int32** |  | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
**HealthReason** | Pointer to **string** |  | [optional] 
**HealthRecommendation** | Pointer to **string** |  | [optional] 
**HwRevision** | Pointer to **string** | Hardware version of the FRU | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**LocatorLed** | Pointer to **string** | Indicates whether the locator LED is on | [optional] 
<<<<<<< HEAD
**LocatorLedNumeric** | Pointer to **int64** | Indicates whether the locator LED is on( In numeric form ) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PartNumber** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **string** | Position of the component in the enclosure | [optional] 
**PositionNumeric** | Pointer to **int64** | Position of the component in the enclosure( In numeric form ) | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusNumeric** | Pointer to **int64** |  | [optional] 
**StatusSes** | Pointer to **string** | SES Common status | [optional] 
**StatusSesNumeric** | Pointer to **int64** | SES Common status( In numeric form ) | [optional] 
=======
**LocatorLedNumeric** | Pointer to **int32** | Indicates whether the locator LED is on( In numeric form ) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PartNumber** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **string** | Position of the component in the enclosure | [optional] 
**PositionNumeric** | Pointer to **int32** | Position of the component in the enclosure( In numeric form ) | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusNumeric** | Pointer to **int32** |  | [optional] 
**StatusSes** | Pointer to **string** | SES Common status | [optional] 
**StatusSesNumeric** | Pointer to **int32** | SES Common status( In numeric form ) | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
**Url** | Pointer to **string** | The resource URL | [optional] 

## Methods

### NewFanResourceInner

`func NewFanResourceInner() *FanResourceInner`

NewFanResourceInner instantiates a new FanResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFanResourceInnerWithDefaults

`func NewFanResourceInnerWithDefaults() *FanResourceInner`

NewFanResourceInnerWithDefaults instantiates a new FanResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *FanResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *FanResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *FanResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *FanResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *FanResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FanResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FanResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FanResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetDurableId

`func (o *FanResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *FanResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *FanResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *FanResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetExtendedStatus

`func (o *FanResourceInner) GetExtendedStatus() string`

GetExtendedStatus returns the ExtendedStatus field if non-nil, zero value otherwise.

### GetExtendedStatusOk

`func (o *FanResourceInner) GetExtendedStatusOk() (*string, bool)`

GetExtendedStatusOk returns a tuple with the ExtendedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedStatus

`func (o *FanResourceInner) SetExtendedStatus(v string)`

SetExtendedStatus sets ExtendedStatus field to given value.

### HasExtendedStatus

`func (o *FanResourceInner) HasExtendedStatus() bool`

HasExtendedStatus returns a boolean if a field has been set.

### GetFwRevision

`func (o *FanResourceInner) GetFwRevision() string`

GetFwRevision returns the FwRevision field if non-nil, zero value otherwise.

### GetFwRevisionOk

`func (o *FanResourceInner) GetFwRevisionOk() (*string, bool)`

GetFwRevisionOk returns a tuple with the FwRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwRevision

`func (o *FanResourceInner) SetFwRevision(v string)`

SetFwRevision sets FwRevision field to given value.

### HasFwRevision

`func (o *FanResourceInner) HasFwRevision() bool`

HasFwRevision returns a boolean if a field has been set.

### GetHealth

`func (o *FanResourceInner) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *FanResourceInner) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *FanResourceInner) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *FanResourceInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHealthNumeric

<<<<<<< HEAD
`func (o *FanResourceInner) GetHealthNumeric() int64`
=======
`func (o *FanResourceInner) GetHealthNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHealthNumeric returns the HealthNumeric field if non-nil, zero value otherwise.

### GetHealthNumericOk

<<<<<<< HEAD
`func (o *FanResourceInner) GetHealthNumericOk() (*int64, bool)`
=======
`func (o *FanResourceInner) GetHealthNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHealthNumericOk returns a tuple with the HealthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthNumeric

<<<<<<< HEAD
`func (o *FanResourceInner) SetHealthNumeric(v int64)`
=======
`func (o *FanResourceInner) SetHealthNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetHealthNumeric sets HealthNumeric field to given value.

### HasHealthNumeric

`func (o *FanResourceInner) HasHealthNumeric() bool`

HasHealthNumeric returns a boolean if a field has been set.

### GetHealthReason

`func (o *FanResourceInner) GetHealthReason() string`

GetHealthReason returns the HealthReason field if non-nil, zero value otherwise.

### GetHealthReasonOk

`func (o *FanResourceInner) GetHealthReasonOk() (*string, bool)`

GetHealthReasonOk returns a tuple with the HealthReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReason

`func (o *FanResourceInner) SetHealthReason(v string)`

SetHealthReason sets HealthReason field to given value.

### HasHealthReason

`func (o *FanResourceInner) HasHealthReason() bool`

HasHealthReason returns a boolean if a field has been set.

### GetHealthRecommendation

`func (o *FanResourceInner) GetHealthRecommendation() string`

GetHealthRecommendation returns the HealthRecommendation field if non-nil, zero value otherwise.

### GetHealthRecommendationOk

`func (o *FanResourceInner) GetHealthRecommendationOk() (*string, bool)`

GetHealthRecommendationOk returns a tuple with the HealthRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecommendation

`func (o *FanResourceInner) SetHealthRecommendation(v string)`

SetHealthRecommendation sets HealthRecommendation field to given value.

### HasHealthRecommendation

`func (o *FanResourceInner) HasHealthRecommendation() bool`

HasHealthRecommendation returns a boolean if a field has been set.

### GetHwRevision

`func (o *FanResourceInner) GetHwRevision() string`

GetHwRevision returns the HwRevision field if non-nil, zero value otherwise.

### GetHwRevisionOk

`func (o *FanResourceInner) GetHwRevisionOk() (*string, bool)`

GetHwRevisionOk returns a tuple with the HwRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwRevision

`func (o *FanResourceInner) SetHwRevision(v string)`

SetHwRevision sets HwRevision field to given value.

### HasHwRevision

`func (o *FanResourceInner) HasHwRevision() bool`

HasHwRevision returns a boolean if a field has been set.

### GetLocation

`func (o *FanResourceInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FanResourceInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FanResourceInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FanResourceInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLocatorLed

`func (o *FanResourceInner) GetLocatorLed() string`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *FanResourceInner) GetLocatorLedOk() (*string, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *FanResourceInner) SetLocatorLed(v string)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *FanResourceInner) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### GetLocatorLedNumeric

<<<<<<< HEAD
`func (o *FanResourceInner) GetLocatorLedNumeric() int64`
=======
`func (o *FanResourceInner) GetLocatorLedNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLocatorLedNumeric returns the LocatorLedNumeric field if non-nil, zero value otherwise.

### GetLocatorLedNumericOk

<<<<<<< HEAD
`func (o *FanResourceInner) GetLocatorLedNumericOk() (*int64, bool)`
=======
`func (o *FanResourceInner) GetLocatorLedNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLocatorLedNumericOk returns a tuple with the LocatorLedNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLedNumeric

<<<<<<< HEAD
`func (o *FanResourceInner) SetLocatorLedNumeric(v int64)`
=======
`func (o *FanResourceInner) SetLocatorLedNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetLocatorLedNumeric sets LocatorLedNumeric field to given value.

### HasLocatorLedNumeric

`func (o *FanResourceInner) HasLocatorLedNumeric() bool`

HasLocatorLedNumeric returns a boolean if a field has been set.

### GetName

`func (o *FanResourceInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FanResourceInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FanResourceInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FanResourceInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartNumber

`func (o *FanResourceInner) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *FanResourceInner) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *FanResourceInner) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *FanResourceInner) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPosition

`func (o *FanResourceInner) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FanResourceInner) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FanResourceInner) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *FanResourceInner) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPositionNumeric

<<<<<<< HEAD
`func (o *FanResourceInner) GetPositionNumeric() int64`
=======
`func (o *FanResourceInner) GetPositionNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPositionNumeric returns the PositionNumeric field if non-nil, zero value otherwise.

### GetPositionNumericOk

<<<<<<< HEAD
`func (o *FanResourceInner) GetPositionNumericOk() (*int64, bool)`
=======
`func (o *FanResourceInner) GetPositionNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPositionNumericOk returns a tuple with the PositionNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionNumeric

<<<<<<< HEAD
`func (o *FanResourceInner) SetPositionNumeric(v int64)`
=======
`func (o *FanResourceInner) SetPositionNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetPositionNumeric sets PositionNumeric field to given value.

### HasPositionNumeric

`func (o *FanResourceInner) HasPositionNumeric() bool`

HasPositionNumeric returns a boolean if a field has been set.

### GetSerialNumber

`func (o *FanResourceInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *FanResourceInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *FanResourceInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *FanResourceInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSpeed

<<<<<<< HEAD
`func (o *FanResourceInner) GetSpeed() int64`
=======
`func (o *FanResourceInner) GetSpeed() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

<<<<<<< HEAD
`func (o *FanResourceInner) GetSpeedOk() (*int64, bool)`
=======
`func (o *FanResourceInner) GetSpeedOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

<<<<<<< HEAD
`func (o *FanResourceInner) SetSpeed(v int64)`
=======
`func (o *FanResourceInner) SetSpeed(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *FanResourceInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *FanResourceInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FanResourceInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FanResourceInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FanResourceInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusNumeric

<<<<<<< HEAD
`func (o *FanResourceInner) GetStatusNumeric() int64`
=======
`func (o *FanResourceInner) GetStatusNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetStatusNumeric returns the StatusNumeric field if non-nil, zero value otherwise.

### GetStatusNumericOk

<<<<<<< HEAD
`func (o *FanResourceInner) GetStatusNumericOk() (*int64, bool)`
=======
`func (o *FanResourceInner) GetStatusNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetStatusNumericOk returns a tuple with the StatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusNumeric

<<<<<<< HEAD
`func (o *FanResourceInner) SetStatusNumeric(v int64)`
=======
`func (o *FanResourceInner) SetStatusNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetStatusNumeric sets StatusNumeric field to given value.

### HasStatusNumeric

`func (o *FanResourceInner) HasStatusNumeric() bool`

HasStatusNumeric returns a boolean if a field has been set.

### GetStatusSes

`func (o *FanResourceInner) GetStatusSes() string`

GetStatusSes returns the StatusSes field if non-nil, zero value otherwise.

### GetStatusSesOk

`func (o *FanResourceInner) GetStatusSesOk() (*string, bool)`

GetStatusSesOk returns a tuple with the StatusSes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSes

`func (o *FanResourceInner) SetStatusSes(v string)`

SetStatusSes sets StatusSes field to given value.

### HasStatusSes

`func (o *FanResourceInner) HasStatusSes() bool`

HasStatusSes returns a boolean if a field has been set.

### GetStatusSesNumeric

<<<<<<< HEAD
`func (o *FanResourceInner) GetStatusSesNumeric() int64`
=======
`func (o *FanResourceInner) GetStatusSesNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetStatusSesNumeric returns the StatusSesNumeric field if non-nil, zero value otherwise.

### GetStatusSesNumericOk

<<<<<<< HEAD
`func (o *FanResourceInner) GetStatusSesNumericOk() (*int64, bool)`
=======
`func (o *FanResourceInner) GetStatusSesNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetStatusSesNumericOk returns a tuple with the StatusSesNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSesNumeric

<<<<<<< HEAD
`func (o *FanResourceInner) SetStatusSesNumeric(v int64)`
=======
`func (o *FanResourceInner) SetStatusSesNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetStatusSesNumeric sets StatusSesNumeric field to given value.

### HasStatusSesNumeric

`func (o *FanResourceInner) HasStatusSesNumeric() bool`

HasStatusSesNumeric returns a boolean if a field has been set.

### GetUrl

`func (o *FanResourceInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FanResourceInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FanResourceInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FanResourceInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


