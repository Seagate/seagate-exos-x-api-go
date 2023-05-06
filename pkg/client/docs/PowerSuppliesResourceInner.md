# PowerSuppliesResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**ConfigurationSerialnumber** | Pointer to **string** |  | [optional] 
**DashLevel** | Pointer to **string** |  | [optional] 
**Dc12i** | Pointer to **int32** |  | [optional] 
**Dc12v** | Pointer to **int32** |  | [optional] 
**Dc33v** | Pointer to **int32** |  | [optional] 
**Dc5i** | Pointer to **int32** |  | [optional] 
**Dc5v** | Pointer to **int32** |  | [optional] 
**Dctemp** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DomId** | Pointer to **int32** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**EnclosureId** | Pointer to **int32** |  | [optional] 
**EnclosuresUrl** | Pointer to **string** |  | [optional] 
**FruShortname** | Pointer to **string** |  | [optional] 
**FwRevision** | Pointer to **string** | Firmware version of the FRU | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**HealthNumeric** | Pointer to **int32** |  | [optional] 
**HealthReason** | Pointer to **string** |  | [optional] 
**HealthRecommendation** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**MfgDate** | Pointer to **string** |  | [optional] 
**MfgDateNumeric** | Pointer to **int32** |  | [optional] 
**MfgLocation** | Pointer to **string** |  | [optional] 
**MfgVendorId** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PartNumber** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **string** | Position of the component in the enclosure | [optional] 
**PositionNumeric** | Pointer to **int32** | Position of the component in the enclosure( In numeric form ) | [optional] 
**Revision** | Pointer to **string** | Current revision for this FRU | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusNumeric** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** | The resource URL | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Fan** | Pointer to [**[]FanResourceInner**](FanResourceInner.md) |  | [optional] 

## Methods

### NewPowerSuppliesResourceInner

`func NewPowerSuppliesResourceInner() *PowerSuppliesResourceInner`

NewPowerSuppliesResourceInner instantiates a new PowerSuppliesResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerSuppliesResourceInnerWithDefaults

`func NewPowerSuppliesResourceInnerWithDefaults() *PowerSuppliesResourceInner`

NewPowerSuppliesResourceInnerWithDefaults instantiates a new PowerSuppliesResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *PowerSuppliesResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *PowerSuppliesResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *PowerSuppliesResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *PowerSuppliesResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *PowerSuppliesResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PowerSuppliesResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PowerSuppliesResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PowerSuppliesResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetConfigurationSerialnumber

`func (o *PowerSuppliesResourceInner) GetConfigurationSerialnumber() string`

GetConfigurationSerialnumber returns the ConfigurationSerialnumber field if non-nil, zero value otherwise.

### GetConfigurationSerialnumberOk

`func (o *PowerSuppliesResourceInner) GetConfigurationSerialnumberOk() (*string, bool)`

GetConfigurationSerialnumberOk returns a tuple with the ConfigurationSerialnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationSerialnumber

`func (o *PowerSuppliesResourceInner) SetConfigurationSerialnumber(v string)`

SetConfigurationSerialnumber sets ConfigurationSerialnumber field to given value.

### HasConfigurationSerialnumber

`func (o *PowerSuppliesResourceInner) HasConfigurationSerialnumber() bool`

HasConfigurationSerialnumber returns a boolean if a field has been set.

### GetDashLevel

`func (o *PowerSuppliesResourceInner) GetDashLevel() string`

GetDashLevel returns the DashLevel field if non-nil, zero value otherwise.

### GetDashLevelOk

`func (o *PowerSuppliesResourceInner) GetDashLevelOk() (*string, bool)`

GetDashLevelOk returns a tuple with the DashLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashLevel

`func (o *PowerSuppliesResourceInner) SetDashLevel(v string)`

SetDashLevel sets DashLevel field to given value.

### HasDashLevel

`func (o *PowerSuppliesResourceInner) HasDashLevel() bool`

HasDashLevel returns a boolean if a field has been set.

### GetDc12i

`func (o *PowerSuppliesResourceInner) GetDc12i() int32`

GetDc12i returns the Dc12i field if non-nil, zero value otherwise.

### GetDc12iOk

`func (o *PowerSuppliesResourceInner) GetDc12iOk() (*int32, bool)`

GetDc12iOk returns a tuple with the Dc12i field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc12i

`func (o *PowerSuppliesResourceInner) SetDc12i(v int32)`

SetDc12i sets Dc12i field to given value.

### HasDc12i

`func (o *PowerSuppliesResourceInner) HasDc12i() bool`

HasDc12i returns a boolean if a field has been set.

### GetDc12v

`func (o *PowerSuppliesResourceInner) GetDc12v() int32`

GetDc12v returns the Dc12v field if non-nil, zero value otherwise.

### GetDc12vOk

`func (o *PowerSuppliesResourceInner) GetDc12vOk() (*int32, bool)`

GetDc12vOk returns a tuple with the Dc12v field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc12v

`func (o *PowerSuppliesResourceInner) SetDc12v(v int32)`

SetDc12v sets Dc12v field to given value.

### HasDc12v

`func (o *PowerSuppliesResourceInner) HasDc12v() bool`

HasDc12v returns a boolean if a field has been set.

### GetDc33v

`func (o *PowerSuppliesResourceInner) GetDc33v() int32`

GetDc33v returns the Dc33v field if non-nil, zero value otherwise.

### GetDc33vOk

`func (o *PowerSuppliesResourceInner) GetDc33vOk() (*int32, bool)`

GetDc33vOk returns a tuple with the Dc33v field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc33v

`func (o *PowerSuppliesResourceInner) SetDc33v(v int32)`

SetDc33v sets Dc33v field to given value.

### HasDc33v

`func (o *PowerSuppliesResourceInner) HasDc33v() bool`

HasDc33v returns a boolean if a field has been set.

### GetDc5i

`func (o *PowerSuppliesResourceInner) GetDc5i() int32`

GetDc5i returns the Dc5i field if non-nil, zero value otherwise.

### GetDc5iOk

`func (o *PowerSuppliesResourceInner) GetDc5iOk() (*int32, bool)`

GetDc5iOk returns a tuple with the Dc5i field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc5i

`func (o *PowerSuppliesResourceInner) SetDc5i(v int32)`

SetDc5i sets Dc5i field to given value.

### HasDc5i

`func (o *PowerSuppliesResourceInner) HasDc5i() bool`

HasDc5i returns a boolean if a field has been set.

### GetDc5v

`func (o *PowerSuppliesResourceInner) GetDc5v() int32`

GetDc5v returns the Dc5v field if non-nil, zero value otherwise.

### GetDc5vOk

`func (o *PowerSuppliesResourceInner) GetDc5vOk() (*int32, bool)`

GetDc5vOk returns a tuple with the Dc5v field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc5v

`func (o *PowerSuppliesResourceInner) SetDc5v(v int32)`

SetDc5v sets Dc5v field to given value.

### HasDc5v

`func (o *PowerSuppliesResourceInner) HasDc5v() bool`

HasDc5v returns a boolean if a field has been set.

### GetDctemp

`func (o *PowerSuppliesResourceInner) GetDctemp() int32`

GetDctemp returns the Dctemp field if non-nil, zero value otherwise.

### GetDctempOk

`func (o *PowerSuppliesResourceInner) GetDctempOk() (*int32, bool)`

GetDctempOk returns a tuple with the Dctemp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDctemp

`func (o *PowerSuppliesResourceInner) SetDctemp(v int32)`

SetDctemp sets Dctemp field to given value.

### HasDctemp

`func (o *PowerSuppliesResourceInner) HasDctemp() bool`

HasDctemp returns a boolean if a field has been set.

### GetDescription

`func (o *PowerSuppliesResourceInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PowerSuppliesResourceInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PowerSuppliesResourceInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PowerSuppliesResourceInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomId

`func (o *PowerSuppliesResourceInner) GetDomId() int32`

GetDomId returns the DomId field if non-nil, zero value otherwise.

### GetDomIdOk

`func (o *PowerSuppliesResourceInner) GetDomIdOk() (*int32, bool)`

GetDomIdOk returns a tuple with the DomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomId

`func (o *PowerSuppliesResourceInner) SetDomId(v int32)`

SetDomId sets DomId field to given value.

### HasDomId

`func (o *PowerSuppliesResourceInner) HasDomId() bool`

HasDomId returns a boolean if a field has been set.

### GetDurableId

`func (o *PowerSuppliesResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *PowerSuppliesResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *PowerSuppliesResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *PowerSuppliesResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetEnclosureId

`func (o *PowerSuppliesResourceInner) GetEnclosureId() int32`

GetEnclosureId returns the EnclosureId field if non-nil, zero value otherwise.

### GetEnclosureIdOk

`func (o *PowerSuppliesResourceInner) GetEnclosureIdOk() (*int32, bool)`

GetEnclosureIdOk returns a tuple with the EnclosureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureId

`func (o *PowerSuppliesResourceInner) SetEnclosureId(v int32)`

SetEnclosureId sets EnclosureId field to given value.

### HasEnclosureId

`func (o *PowerSuppliesResourceInner) HasEnclosureId() bool`

HasEnclosureId returns a boolean if a field has been set.

### GetEnclosuresUrl

`func (o *PowerSuppliesResourceInner) GetEnclosuresUrl() string`

GetEnclosuresUrl returns the EnclosuresUrl field if non-nil, zero value otherwise.

### GetEnclosuresUrlOk

`func (o *PowerSuppliesResourceInner) GetEnclosuresUrlOk() (*string, bool)`

GetEnclosuresUrlOk returns a tuple with the EnclosuresUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosuresUrl

`func (o *PowerSuppliesResourceInner) SetEnclosuresUrl(v string)`

SetEnclosuresUrl sets EnclosuresUrl field to given value.

### HasEnclosuresUrl

`func (o *PowerSuppliesResourceInner) HasEnclosuresUrl() bool`

HasEnclosuresUrl returns a boolean if a field has been set.

### GetFruShortname

`func (o *PowerSuppliesResourceInner) GetFruShortname() string`

GetFruShortname returns the FruShortname field if non-nil, zero value otherwise.

### GetFruShortnameOk

`func (o *PowerSuppliesResourceInner) GetFruShortnameOk() (*string, bool)`

GetFruShortnameOk returns a tuple with the FruShortname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFruShortname

`func (o *PowerSuppliesResourceInner) SetFruShortname(v string)`

SetFruShortname sets FruShortname field to given value.

### HasFruShortname

`func (o *PowerSuppliesResourceInner) HasFruShortname() bool`

HasFruShortname returns a boolean if a field has been set.

### GetFwRevision

`func (o *PowerSuppliesResourceInner) GetFwRevision() string`

GetFwRevision returns the FwRevision field if non-nil, zero value otherwise.

### GetFwRevisionOk

`func (o *PowerSuppliesResourceInner) GetFwRevisionOk() (*string, bool)`

GetFwRevisionOk returns a tuple with the FwRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwRevision

`func (o *PowerSuppliesResourceInner) SetFwRevision(v string)`

SetFwRevision sets FwRevision field to given value.

### HasFwRevision

`func (o *PowerSuppliesResourceInner) HasFwRevision() bool`

HasFwRevision returns a boolean if a field has been set.

### GetHealth

`func (o *PowerSuppliesResourceInner) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *PowerSuppliesResourceInner) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *PowerSuppliesResourceInner) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *PowerSuppliesResourceInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHealthNumeric

`func (o *PowerSuppliesResourceInner) GetHealthNumeric() int32`

GetHealthNumeric returns the HealthNumeric field if non-nil, zero value otherwise.

### GetHealthNumericOk

`func (o *PowerSuppliesResourceInner) GetHealthNumericOk() (*int32, bool)`

GetHealthNumericOk returns a tuple with the HealthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthNumeric

`func (o *PowerSuppliesResourceInner) SetHealthNumeric(v int32)`

SetHealthNumeric sets HealthNumeric field to given value.

### HasHealthNumeric

`func (o *PowerSuppliesResourceInner) HasHealthNumeric() bool`

HasHealthNumeric returns a boolean if a field has been set.

### GetHealthReason

`func (o *PowerSuppliesResourceInner) GetHealthReason() string`

GetHealthReason returns the HealthReason field if non-nil, zero value otherwise.

### GetHealthReasonOk

`func (o *PowerSuppliesResourceInner) GetHealthReasonOk() (*string, bool)`

GetHealthReasonOk returns a tuple with the HealthReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReason

`func (o *PowerSuppliesResourceInner) SetHealthReason(v string)`

SetHealthReason sets HealthReason field to given value.

### HasHealthReason

`func (o *PowerSuppliesResourceInner) HasHealthReason() bool`

HasHealthReason returns a boolean if a field has been set.

### GetHealthRecommendation

`func (o *PowerSuppliesResourceInner) GetHealthRecommendation() string`

GetHealthRecommendation returns the HealthRecommendation field if non-nil, zero value otherwise.

### GetHealthRecommendationOk

`func (o *PowerSuppliesResourceInner) GetHealthRecommendationOk() (*string, bool)`

GetHealthRecommendationOk returns a tuple with the HealthRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecommendation

`func (o *PowerSuppliesResourceInner) SetHealthRecommendation(v string)`

SetHealthRecommendation sets HealthRecommendation field to given value.

### HasHealthRecommendation

`func (o *PowerSuppliesResourceInner) HasHealthRecommendation() bool`

HasHealthRecommendation returns a boolean if a field has been set.

### GetLocation

`func (o *PowerSuppliesResourceInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PowerSuppliesResourceInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PowerSuppliesResourceInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PowerSuppliesResourceInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMfgDate

`func (o *PowerSuppliesResourceInner) GetMfgDate() string`

GetMfgDate returns the MfgDate field if non-nil, zero value otherwise.

### GetMfgDateOk

`func (o *PowerSuppliesResourceInner) GetMfgDateOk() (*string, bool)`

GetMfgDateOk returns a tuple with the MfgDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgDate

`func (o *PowerSuppliesResourceInner) SetMfgDate(v string)`

SetMfgDate sets MfgDate field to given value.

### HasMfgDate

`func (o *PowerSuppliesResourceInner) HasMfgDate() bool`

HasMfgDate returns a boolean if a field has been set.

### GetMfgDateNumeric

`func (o *PowerSuppliesResourceInner) GetMfgDateNumeric() int32`

GetMfgDateNumeric returns the MfgDateNumeric field if non-nil, zero value otherwise.

### GetMfgDateNumericOk

`func (o *PowerSuppliesResourceInner) GetMfgDateNumericOk() (*int32, bool)`

GetMfgDateNumericOk returns a tuple with the MfgDateNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgDateNumeric

`func (o *PowerSuppliesResourceInner) SetMfgDateNumeric(v int32)`

SetMfgDateNumeric sets MfgDateNumeric field to given value.

### HasMfgDateNumeric

`func (o *PowerSuppliesResourceInner) HasMfgDateNumeric() bool`

HasMfgDateNumeric returns a boolean if a field has been set.

### GetMfgLocation

`func (o *PowerSuppliesResourceInner) GetMfgLocation() string`

GetMfgLocation returns the MfgLocation field if non-nil, zero value otherwise.

### GetMfgLocationOk

`func (o *PowerSuppliesResourceInner) GetMfgLocationOk() (*string, bool)`

GetMfgLocationOk returns a tuple with the MfgLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgLocation

`func (o *PowerSuppliesResourceInner) SetMfgLocation(v string)`

SetMfgLocation sets MfgLocation field to given value.

### HasMfgLocation

`func (o *PowerSuppliesResourceInner) HasMfgLocation() bool`

HasMfgLocation returns a boolean if a field has been set.

### GetMfgVendorId

`func (o *PowerSuppliesResourceInner) GetMfgVendorId() string`

GetMfgVendorId returns the MfgVendorId field if non-nil, zero value otherwise.

### GetMfgVendorIdOk

`func (o *PowerSuppliesResourceInner) GetMfgVendorIdOk() (*string, bool)`

GetMfgVendorIdOk returns a tuple with the MfgVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgVendorId

`func (o *PowerSuppliesResourceInner) SetMfgVendorId(v string)`

SetMfgVendorId sets MfgVendorId field to given value.

### HasMfgVendorId

`func (o *PowerSuppliesResourceInner) HasMfgVendorId() bool`

HasMfgVendorId returns a boolean if a field has been set.

### GetModel

`func (o *PowerSuppliesResourceInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PowerSuppliesResourceInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PowerSuppliesResourceInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PowerSuppliesResourceInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *PowerSuppliesResourceInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerSuppliesResourceInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerSuppliesResourceInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PowerSuppliesResourceInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartNumber

`func (o *PowerSuppliesResourceInner) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *PowerSuppliesResourceInner) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *PowerSuppliesResourceInner) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *PowerSuppliesResourceInner) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPosition

`func (o *PowerSuppliesResourceInner) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PowerSuppliesResourceInner) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PowerSuppliesResourceInner) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PowerSuppliesResourceInner) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPositionNumeric

`func (o *PowerSuppliesResourceInner) GetPositionNumeric() int32`

GetPositionNumeric returns the PositionNumeric field if non-nil, zero value otherwise.

### GetPositionNumericOk

`func (o *PowerSuppliesResourceInner) GetPositionNumericOk() (*int32, bool)`

GetPositionNumericOk returns a tuple with the PositionNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionNumeric

`func (o *PowerSuppliesResourceInner) SetPositionNumeric(v int32)`

SetPositionNumeric sets PositionNumeric field to given value.

### HasPositionNumeric

`func (o *PowerSuppliesResourceInner) HasPositionNumeric() bool`

HasPositionNumeric returns a boolean if a field has been set.

### GetRevision

`func (o *PowerSuppliesResourceInner) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PowerSuppliesResourceInner) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PowerSuppliesResourceInner) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *PowerSuppliesResourceInner) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSerialNumber

`func (o *PowerSuppliesResourceInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PowerSuppliesResourceInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PowerSuppliesResourceInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *PowerSuppliesResourceInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetStatus

`func (o *PowerSuppliesResourceInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PowerSuppliesResourceInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PowerSuppliesResourceInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PowerSuppliesResourceInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusNumeric

`func (o *PowerSuppliesResourceInner) GetStatusNumeric() int32`

GetStatusNumeric returns the StatusNumeric field if non-nil, zero value otherwise.

### GetStatusNumericOk

`func (o *PowerSuppliesResourceInner) GetStatusNumericOk() (*int32, bool)`

GetStatusNumericOk returns a tuple with the StatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusNumeric

`func (o *PowerSuppliesResourceInner) SetStatusNumeric(v int32)`

SetStatusNumeric sets StatusNumeric field to given value.

### HasStatusNumeric

`func (o *PowerSuppliesResourceInner) HasStatusNumeric() bool`

HasStatusNumeric returns a boolean if a field has been set.

### GetUrl

`func (o *PowerSuppliesResourceInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PowerSuppliesResourceInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PowerSuppliesResourceInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PowerSuppliesResourceInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVendor

`func (o *PowerSuppliesResourceInner) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *PowerSuppliesResourceInner) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *PowerSuppliesResourceInner) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *PowerSuppliesResourceInner) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetFan

`func (o *PowerSuppliesResourceInner) GetFan() []FanResourceInner`

GetFan returns the Fan field if non-nil, zero value otherwise.

### GetFanOk

`func (o *PowerSuppliesResourceInner) GetFanOk() (*[]FanResourceInner, bool)`

GetFanOk returns a tuple with the Fan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFan

`func (o *PowerSuppliesResourceInner) SetFan(v []FanResourceInner)`

SetFan sets Fan field to given value.

### HasFan

`func (o *PowerSuppliesResourceInner) HasFan() bool`

HasFan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


