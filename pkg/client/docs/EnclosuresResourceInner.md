# EnclosuresResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**BoardModel** | Pointer to **string** | IOM Board model | [optional] 
**BoardModelNumeric** | Pointer to **int64** | IOM Board model( In numeric form ) | [optional] 
**Columns** | Pointer to **int64** |  | [optional] 
**DashLevel** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DriveOrientation** | Pointer to **string** |  | [optional] 
**DriveOrientationNumeric** | Pointer to **int64** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**EmpA** | Pointer to **string** |  | [optional] 
**EmpABusid** | Pointer to **string** |  | [optional] 
**EmpAChIdRev** | Pointer to **string** |  | [optional] 
**EmpARev** | Pointer to **string** | Enclosure Management Processor | [optional] 
**EmpATargetid** | Pointer to **string** |  | [optional] 
**EmpB** | Pointer to **string** |  | [optional] 
**EmpBBusid** | Pointer to **string** |  | [optional] 
**EmpBChIdRev** | Pointer to **string** |  | [optional] 
**EmpBRev** | Pointer to **string** | Enclosure Management Processor | [optional] 
**EmpBTargetid** | Pointer to **string** |  | [optional] 
**EnclosureArrangement** | Pointer to **string** |  | [optional] 
**EnclosureArrangementNumeric** | Pointer to **int64** |  | [optional] 
**EnclosureId** | Pointer to **int64** |  | [optional] 
**EnclosurePower** | Pointer to **string** |  | [optional] 
**EnclosureWwn** | Pointer to **string** | Enclosure World Wide Name | [optional] 
**ExtendedStatus** | Pointer to **string** | Extended status (bits) | [optional] 
**FruLocation** | Pointer to **string** |  | [optional] 
**FruShortname** | Pointer to **string** |  | [optional] 
**FruTlapn** | Pointer to **string** |  | [optional] 
**GemVersionA** | Pointer to **string** |  | [optional] 
**GemVersionB** | Pointer to **string** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**HealthNumeric** | Pointer to **int64** |  | [optional] 
**HealthReason** | Pointer to **string** |  | [optional] 
**HealthRecommendation** | Pointer to **string** |  | [optional] 
**IomType** | Pointer to **string** |  | [optional] 
**IomTypeNumeric** | Pointer to **int64** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**LocatorLed** | Pointer to **string** | Indicates whether the locator LED is on | [optional] 
**LocatorLedNumeric** | Pointer to **int64** | Indicates whether the locator LED is on( In numeric form ) | [optional] 
**MfgDate** | Pointer to **string** |  | [optional] 
**MfgDateNumeric** | Pointer to **int64** |  | [optional] 
**MfgLocation** | Pointer to **string** |  | [optional] 
**MidplaneRev** | Pointer to **int64** | Revision level of midplane | [optional] 
**MidplaneSerialNumber** | Pointer to **string** | Serial number of the enclosure | [optional] 
**MidplaneType** | Pointer to **string** | Type of midplane | [optional] 
**MidplaneTypeNumeric** | Pointer to **int64** | Type of midplane( In numeric form ) | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumberOfCoolingsElements** | Pointer to **int64** | Number of fans in the enclosure | [optional] 
**NumberOfDisks** | Pointer to **int64** | Number of disks in the enclosure | [optional] 
**NumberOfPowerSupplies** | Pointer to **int64** | Number of power supplies in the enclosure | [optional] 
**PartNumber** | Pointer to **string** |  | [optional] 
**Pcie2Capable** | Pointer to **string** | Enclosure is capable of using PCIe version 2 | [optional] 
**Pcie2CapableNumeric** | Pointer to **int64** | Enclosure is capable of using PCIe version 2( In numeric form ) | [optional] 
**PlatformType** | Pointer to **string** | HW Platform Type | [optional] 
**PlatformTypeNumeric** | Pointer to **int64** | HW Platform Type( In numeric form ) | [optional] 
**RackNumber** | Pointer to **int64** |  | [optional] 
**RackPosition** | Pointer to **int64** |  | [optional] 
**Revision** | Pointer to **string** | Current revision for this FRU | [optional] 
**Rows** | Pointer to **int64** |  | [optional] 
**Slots** | Pointer to **int64** | The number of disk slots in this enclosure | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusNumeric** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** | Type of component in this enclosure | [optional] 
**TypeNumeric** | Pointer to **int64** | Type of component in this enclosure( In numeric form ) | [optional] 
**Url** | Pointer to **string** | The resource URL | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Controllers** | Pointer to [**[]ControllersResourceInner**](ControllersResourceInner.md) |  | [optional] 
**PowerSupplies** | Pointer to [**[]PowerSuppliesResourceInner**](PowerSuppliesResourceInner.md) |  | [optional] 

## Methods

### NewEnclosuresResourceInner

`func NewEnclosuresResourceInner() *EnclosuresResourceInner`

NewEnclosuresResourceInner instantiates a new EnclosuresResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnclosuresResourceInnerWithDefaults

`func NewEnclosuresResourceInnerWithDefaults() *EnclosuresResourceInner`

NewEnclosuresResourceInnerWithDefaults instantiates a new EnclosuresResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *EnclosuresResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *EnclosuresResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *EnclosuresResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *EnclosuresResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *EnclosuresResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *EnclosuresResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *EnclosuresResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *EnclosuresResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetBoardModel

`func (o *EnclosuresResourceInner) GetBoardModel() string`

GetBoardModel returns the BoardModel field if non-nil, zero value otherwise.

### GetBoardModelOk

`func (o *EnclosuresResourceInner) GetBoardModelOk() (*string, bool)`

GetBoardModelOk returns a tuple with the BoardModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardModel

`func (o *EnclosuresResourceInner) SetBoardModel(v string)`

SetBoardModel sets BoardModel field to given value.

### HasBoardModel

`func (o *EnclosuresResourceInner) HasBoardModel() bool`

HasBoardModel returns a boolean if a field has been set.

### GetBoardModelNumeric

`func (o *EnclosuresResourceInner) GetBoardModelNumeric() int64`

GetBoardModelNumeric returns the BoardModelNumeric field if non-nil, zero value otherwise.

### GetBoardModelNumericOk

`func (o *EnclosuresResourceInner) GetBoardModelNumericOk() (*int64, bool)`

GetBoardModelNumericOk returns a tuple with the BoardModelNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardModelNumeric

`func (o *EnclosuresResourceInner) SetBoardModelNumeric(v int64)`

SetBoardModelNumeric sets BoardModelNumeric field to given value.

### HasBoardModelNumeric

`func (o *EnclosuresResourceInner) HasBoardModelNumeric() bool`

HasBoardModelNumeric returns a boolean if a field has been set.

### GetColumns

`func (o *EnclosuresResourceInner) GetColumns() int64`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *EnclosuresResourceInner) GetColumnsOk() (*int64, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *EnclosuresResourceInner) SetColumns(v int64)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *EnclosuresResourceInner) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetDashLevel

`func (o *EnclosuresResourceInner) GetDashLevel() string`

GetDashLevel returns the DashLevel field if non-nil, zero value otherwise.

### GetDashLevelOk

`func (o *EnclosuresResourceInner) GetDashLevelOk() (*string, bool)`

GetDashLevelOk returns a tuple with the DashLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashLevel

`func (o *EnclosuresResourceInner) SetDashLevel(v string)`

SetDashLevel sets DashLevel field to given value.

### HasDashLevel

`func (o *EnclosuresResourceInner) HasDashLevel() bool`

HasDashLevel returns a boolean if a field has been set.

### GetDescription

`func (o *EnclosuresResourceInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnclosuresResourceInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnclosuresResourceInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnclosuresResourceInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDriveOrientation

`func (o *EnclosuresResourceInner) GetDriveOrientation() string`

GetDriveOrientation returns the DriveOrientation field if non-nil, zero value otherwise.

### GetDriveOrientationOk

`func (o *EnclosuresResourceInner) GetDriveOrientationOk() (*string, bool)`

GetDriveOrientationOk returns a tuple with the DriveOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveOrientation

`func (o *EnclosuresResourceInner) SetDriveOrientation(v string)`

SetDriveOrientation sets DriveOrientation field to given value.

### HasDriveOrientation

`func (o *EnclosuresResourceInner) HasDriveOrientation() bool`

HasDriveOrientation returns a boolean if a field has been set.

### GetDriveOrientationNumeric

`func (o *EnclosuresResourceInner) GetDriveOrientationNumeric() int64`

GetDriveOrientationNumeric returns the DriveOrientationNumeric field if non-nil, zero value otherwise.

### GetDriveOrientationNumericOk

`func (o *EnclosuresResourceInner) GetDriveOrientationNumericOk() (*int64, bool)`

GetDriveOrientationNumericOk returns a tuple with the DriveOrientationNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveOrientationNumeric

`func (o *EnclosuresResourceInner) SetDriveOrientationNumeric(v int64)`

SetDriveOrientationNumeric sets DriveOrientationNumeric field to given value.

### HasDriveOrientationNumeric

`func (o *EnclosuresResourceInner) HasDriveOrientationNumeric() bool`

HasDriveOrientationNumeric returns a boolean if a field has been set.

### GetDurableId

`func (o *EnclosuresResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *EnclosuresResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *EnclosuresResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *EnclosuresResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetEmpA

`func (o *EnclosuresResourceInner) GetEmpA() string`

GetEmpA returns the EmpA field if non-nil, zero value otherwise.

### GetEmpAOk

`func (o *EnclosuresResourceInner) GetEmpAOk() (*string, bool)`

GetEmpAOk returns a tuple with the EmpA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpA

`func (o *EnclosuresResourceInner) SetEmpA(v string)`

SetEmpA sets EmpA field to given value.

### HasEmpA

`func (o *EnclosuresResourceInner) HasEmpA() bool`

HasEmpA returns a boolean if a field has been set.

### GetEmpABusid

`func (o *EnclosuresResourceInner) GetEmpABusid() string`

GetEmpABusid returns the EmpABusid field if non-nil, zero value otherwise.

### GetEmpABusidOk

`func (o *EnclosuresResourceInner) GetEmpABusidOk() (*string, bool)`

GetEmpABusidOk returns a tuple with the EmpABusid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpABusid

`func (o *EnclosuresResourceInner) SetEmpABusid(v string)`

SetEmpABusid sets EmpABusid field to given value.

### HasEmpABusid

`func (o *EnclosuresResourceInner) HasEmpABusid() bool`

HasEmpABusid returns a boolean if a field has been set.

### GetEmpAChIdRev

`func (o *EnclosuresResourceInner) GetEmpAChIdRev() string`

GetEmpAChIdRev returns the EmpAChIdRev field if non-nil, zero value otherwise.

### GetEmpAChIdRevOk

`func (o *EnclosuresResourceInner) GetEmpAChIdRevOk() (*string, bool)`

GetEmpAChIdRevOk returns a tuple with the EmpAChIdRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpAChIdRev

`func (o *EnclosuresResourceInner) SetEmpAChIdRev(v string)`

SetEmpAChIdRev sets EmpAChIdRev field to given value.

### HasEmpAChIdRev

`func (o *EnclosuresResourceInner) HasEmpAChIdRev() bool`

HasEmpAChIdRev returns a boolean if a field has been set.

### GetEmpARev

`func (o *EnclosuresResourceInner) GetEmpARev() string`

GetEmpARev returns the EmpARev field if non-nil, zero value otherwise.

### GetEmpARevOk

`func (o *EnclosuresResourceInner) GetEmpARevOk() (*string, bool)`

GetEmpARevOk returns a tuple with the EmpARev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpARev

`func (o *EnclosuresResourceInner) SetEmpARev(v string)`

SetEmpARev sets EmpARev field to given value.

### HasEmpARev

`func (o *EnclosuresResourceInner) HasEmpARev() bool`

HasEmpARev returns a boolean if a field has been set.

### GetEmpATargetid

`func (o *EnclosuresResourceInner) GetEmpATargetid() string`

GetEmpATargetid returns the EmpATargetid field if non-nil, zero value otherwise.

### GetEmpATargetidOk

`func (o *EnclosuresResourceInner) GetEmpATargetidOk() (*string, bool)`

GetEmpATargetidOk returns a tuple with the EmpATargetid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpATargetid

`func (o *EnclosuresResourceInner) SetEmpATargetid(v string)`

SetEmpATargetid sets EmpATargetid field to given value.

### HasEmpATargetid

`func (o *EnclosuresResourceInner) HasEmpATargetid() bool`

HasEmpATargetid returns a boolean if a field has been set.

### GetEmpB

`func (o *EnclosuresResourceInner) GetEmpB() string`

GetEmpB returns the EmpB field if non-nil, zero value otherwise.

### GetEmpBOk

`func (o *EnclosuresResourceInner) GetEmpBOk() (*string, bool)`

GetEmpBOk returns a tuple with the EmpB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpB

`func (o *EnclosuresResourceInner) SetEmpB(v string)`

SetEmpB sets EmpB field to given value.

### HasEmpB

`func (o *EnclosuresResourceInner) HasEmpB() bool`

HasEmpB returns a boolean if a field has been set.

### GetEmpBBusid

`func (o *EnclosuresResourceInner) GetEmpBBusid() string`

GetEmpBBusid returns the EmpBBusid field if non-nil, zero value otherwise.

### GetEmpBBusidOk

`func (o *EnclosuresResourceInner) GetEmpBBusidOk() (*string, bool)`

GetEmpBBusidOk returns a tuple with the EmpBBusid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpBBusid

`func (o *EnclosuresResourceInner) SetEmpBBusid(v string)`

SetEmpBBusid sets EmpBBusid field to given value.

### HasEmpBBusid

`func (o *EnclosuresResourceInner) HasEmpBBusid() bool`

HasEmpBBusid returns a boolean if a field has been set.

### GetEmpBChIdRev

`func (o *EnclosuresResourceInner) GetEmpBChIdRev() string`

GetEmpBChIdRev returns the EmpBChIdRev field if non-nil, zero value otherwise.

### GetEmpBChIdRevOk

`func (o *EnclosuresResourceInner) GetEmpBChIdRevOk() (*string, bool)`

GetEmpBChIdRevOk returns a tuple with the EmpBChIdRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpBChIdRev

`func (o *EnclosuresResourceInner) SetEmpBChIdRev(v string)`

SetEmpBChIdRev sets EmpBChIdRev field to given value.

### HasEmpBChIdRev

`func (o *EnclosuresResourceInner) HasEmpBChIdRev() bool`

HasEmpBChIdRev returns a boolean if a field has been set.

### GetEmpBRev

`func (o *EnclosuresResourceInner) GetEmpBRev() string`

GetEmpBRev returns the EmpBRev field if non-nil, zero value otherwise.

### GetEmpBRevOk

`func (o *EnclosuresResourceInner) GetEmpBRevOk() (*string, bool)`

GetEmpBRevOk returns a tuple with the EmpBRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpBRev

`func (o *EnclosuresResourceInner) SetEmpBRev(v string)`

SetEmpBRev sets EmpBRev field to given value.

### HasEmpBRev

`func (o *EnclosuresResourceInner) HasEmpBRev() bool`

HasEmpBRev returns a boolean if a field has been set.

### GetEmpBTargetid

`func (o *EnclosuresResourceInner) GetEmpBTargetid() string`

GetEmpBTargetid returns the EmpBTargetid field if non-nil, zero value otherwise.

### GetEmpBTargetidOk

`func (o *EnclosuresResourceInner) GetEmpBTargetidOk() (*string, bool)`

GetEmpBTargetidOk returns a tuple with the EmpBTargetid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpBTargetid

`func (o *EnclosuresResourceInner) SetEmpBTargetid(v string)`

SetEmpBTargetid sets EmpBTargetid field to given value.

### HasEmpBTargetid

`func (o *EnclosuresResourceInner) HasEmpBTargetid() bool`

HasEmpBTargetid returns a boolean if a field has been set.

### GetEnclosureArrangement

`func (o *EnclosuresResourceInner) GetEnclosureArrangement() string`

GetEnclosureArrangement returns the EnclosureArrangement field if non-nil, zero value otherwise.

### GetEnclosureArrangementOk

`func (o *EnclosuresResourceInner) GetEnclosureArrangementOk() (*string, bool)`

GetEnclosureArrangementOk returns a tuple with the EnclosureArrangement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureArrangement

`func (o *EnclosuresResourceInner) SetEnclosureArrangement(v string)`

SetEnclosureArrangement sets EnclosureArrangement field to given value.

### HasEnclosureArrangement

`func (o *EnclosuresResourceInner) HasEnclosureArrangement() bool`

HasEnclosureArrangement returns a boolean if a field has been set.

### GetEnclosureArrangementNumeric

`func (o *EnclosuresResourceInner) GetEnclosureArrangementNumeric() int64`

GetEnclosureArrangementNumeric returns the EnclosureArrangementNumeric field if non-nil, zero value otherwise.

### GetEnclosureArrangementNumericOk

`func (o *EnclosuresResourceInner) GetEnclosureArrangementNumericOk() (*int64, bool)`

GetEnclosureArrangementNumericOk returns a tuple with the EnclosureArrangementNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureArrangementNumeric

`func (o *EnclosuresResourceInner) SetEnclosureArrangementNumeric(v int64)`

SetEnclosureArrangementNumeric sets EnclosureArrangementNumeric field to given value.

### HasEnclosureArrangementNumeric

`func (o *EnclosuresResourceInner) HasEnclosureArrangementNumeric() bool`

HasEnclosureArrangementNumeric returns a boolean if a field has been set.

### GetEnclosureId

`func (o *EnclosuresResourceInner) GetEnclosureId() int64`

GetEnclosureId returns the EnclosureId field if non-nil, zero value otherwise.

### GetEnclosureIdOk

`func (o *EnclosuresResourceInner) GetEnclosureIdOk() (*int64, bool)`

GetEnclosureIdOk returns a tuple with the EnclosureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureId

`func (o *EnclosuresResourceInner) SetEnclosureId(v int64)`

SetEnclosureId sets EnclosureId field to given value.

### HasEnclosureId

`func (o *EnclosuresResourceInner) HasEnclosureId() bool`

HasEnclosureId returns a boolean if a field has been set.

### GetEnclosurePower

`func (o *EnclosuresResourceInner) GetEnclosurePower() string`

GetEnclosurePower returns the EnclosurePower field if non-nil, zero value otherwise.

### GetEnclosurePowerOk

`func (o *EnclosuresResourceInner) GetEnclosurePowerOk() (*string, bool)`

GetEnclosurePowerOk returns a tuple with the EnclosurePower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosurePower

`func (o *EnclosuresResourceInner) SetEnclosurePower(v string)`

SetEnclosurePower sets EnclosurePower field to given value.

### HasEnclosurePower

`func (o *EnclosuresResourceInner) HasEnclosurePower() bool`

HasEnclosurePower returns a boolean if a field has been set.

### GetEnclosureWwn

`func (o *EnclosuresResourceInner) GetEnclosureWwn() string`

GetEnclosureWwn returns the EnclosureWwn field if non-nil, zero value otherwise.

### GetEnclosureWwnOk

`func (o *EnclosuresResourceInner) GetEnclosureWwnOk() (*string, bool)`

GetEnclosureWwnOk returns a tuple with the EnclosureWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureWwn

`func (o *EnclosuresResourceInner) SetEnclosureWwn(v string)`

SetEnclosureWwn sets EnclosureWwn field to given value.

### HasEnclosureWwn

`func (o *EnclosuresResourceInner) HasEnclosureWwn() bool`

HasEnclosureWwn returns a boolean if a field has been set.

### GetExtendedStatus

`func (o *EnclosuresResourceInner) GetExtendedStatus() string`

GetExtendedStatus returns the ExtendedStatus field if non-nil, zero value otherwise.

### GetExtendedStatusOk

`func (o *EnclosuresResourceInner) GetExtendedStatusOk() (*string, bool)`

GetExtendedStatusOk returns a tuple with the ExtendedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedStatus

`func (o *EnclosuresResourceInner) SetExtendedStatus(v string)`

SetExtendedStatus sets ExtendedStatus field to given value.

### HasExtendedStatus

`func (o *EnclosuresResourceInner) HasExtendedStatus() bool`

HasExtendedStatus returns a boolean if a field has been set.

### GetFruLocation

`func (o *EnclosuresResourceInner) GetFruLocation() string`

GetFruLocation returns the FruLocation field if non-nil, zero value otherwise.

### GetFruLocationOk

`func (o *EnclosuresResourceInner) GetFruLocationOk() (*string, bool)`

GetFruLocationOk returns a tuple with the FruLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFruLocation

`func (o *EnclosuresResourceInner) SetFruLocation(v string)`

SetFruLocation sets FruLocation field to given value.

### HasFruLocation

`func (o *EnclosuresResourceInner) HasFruLocation() bool`

HasFruLocation returns a boolean if a field has been set.

### GetFruShortname

`func (o *EnclosuresResourceInner) GetFruShortname() string`

GetFruShortname returns the FruShortname field if non-nil, zero value otherwise.

### GetFruShortnameOk

`func (o *EnclosuresResourceInner) GetFruShortnameOk() (*string, bool)`

GetFruShortnameOk returns a tuple with the FruShortname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFruShortname

`func (o *EnclosuresResourceInner) SetFruShortname(v string)`

SetFruShortname sets FruShortname field to given value.

### HasFruShortname

`func (o *EnclosuresResourceInner) HasFruShortname() bool`

HasFruShortname returns a boolean if a field has been set.

### GetFruTlapn

`func (o *EnclosuresResourceInner) GetFruTlapn() string`

GetFruTlapn returns the FruTlapn field if non-nil, zero value otherwise.

### GetFruTlapnOk

`func (o *EnclosuresResourceInner) GetFruTlapnOk() (*string, bool)`

GetFruTlapnOk returns a tuple with the FruTlapn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFruTlapn

`func (o *EnclosuresResourceInner) SetFruTlapn(v string)`

SetFruTlapn sets FruTlapn field to given value.

### HasFruTlapn

`func (o *EnclosuresResourceInner) HasFruTlapn() bool`

HasFruTlapn returns a boolean if a field has been set.

### GetGemVersionA

`func (o *EnclosuresResourceInner) GetGemVersionA() string`

GetGemVersionA returns the GemVersionA field if non-nil, zero value otherwise.

### GetGemVersionAOk

`func (o *EnclosuresResourceInner) GetGemVersionAOk() (*string, bool)`

GetGemVersionAOk returns a tuple with the GemVersionA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemVersionA

`func (o *EnclosuresResourceInner) SetGemVersionA(v string)`

SetGemVersionA sets GemVersionA field to given value.

### HasGemVersionA

`func (o *EnclosuresResourceInner) HasGemVersionA() bool`

HasGemVersionA returns a boolean if a field has been set.

### GetGemVersionB

`func (o *EnclosuresResourceInner) GetGemVersionB() string`

GetGemVersionB returns the GemVersionB field if non-nil, zero value otherwise.

### GetGemVersionBOk

`func (o *EnclosuresResourceInner) GetGemVersionBOk() (*string, bool)`

GetGemVersionBOk returns a tuple with the GemVersionB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemVersionB

`func (o *EnclosuresResourceInner) SetGemVersionB(v string)`

SetGemVersionB sets GemVersionB field to given value.

### HasGemVersionB

`func (o *EnclosuresResourceInner) HasGemVersionB() bool`

HasGemVersionB returns a boolean if a field has been set.

### GetHealth

`func (o *EnclosuresResourceInner) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *EnclosuresResourceInner) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *EnclosuresResourceInner) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *EnclosuresResourceInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHealthNumeric

`func (o *EnclosuresResourceInner) GetHealthNumeric() int64`

GetHealthNumeric returns the HealthNumeric field if non-nil, zero value otherwise.

### GetHealthNumericOk

`func (o *EnclosuresResourceInner) GetHealthNumericOk() (*int64, bool)`

GetHealthNumericOk returns a tuple with the HealthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthNumeric

`func (o *EnclosuresResourceInner) SetHealthNumeric(v int64)`

SetHealthNumeric sets HealthNumeric field to given value.

### HasHealthNumeric

`func (o *EnclosuresResourceInner) HasHealthNumeric() bool`

HasHealthNumeric returns a boolean if a field has been set.

### GetHealthReason

`func (o *EnclosuresResourceInner) GetHealthReason() string`

GetHealthReason returns the HealthReason field if non-nil, zero value otherwise.

### GetHealthReasonOk

`func (o *EnclosuresResourceInner) GetHealthReasonOk() (*string, bool)`

GetHealthReasonOk returns a tuple with the HealthReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReason

`func (o *EnclosuresResourceInner) SetHealthReason(v string)`

SetHealthReason sets HealthReason field to given value.

### HasHealthReason

`func (o *EnclosuresResourceInner) HasHealthReason() bool`

HasHealthReason returns a boolean if a field has been set.

### GetHealthRecommendation

`func (o *EnclosuresResourceInner) GetHealthRecommendation() string`

GetHealthRecommendation returns the HealthRecommendation field if non-nil, zero value otherwise.

### GetHealthRecommendationOk

`func (o *EnclosuresResourceInner) GetHealthRecommendationOk() (*string, bool)`

GetHealthRecommendationOk returns a tuple with the HealthRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecommendation

`func (o *EnclosuresResourceInner) SetHealthRecommendation(v string)`

SetHealthRecommendation sets HealthRecommendation field to given value.

### HasHealthRecommendation

`func (o *EnclosuresResourceInner) HasHealthRecommendation() bool`

HasHealthRecommendation returns a boolean if a field has been set.

### GetIomType

`func (o *EnclosuresResourceInner) GetIomType() string`

GetIomType returns the IomType field if non-nil, zero value otherwise.

### GetIomTypeOk

`func (o *EnclosuresResourceInner) GetIomTypeOk() (*string, bool)`

GetIomTypeOk returns a tuple with the IomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIomType

`func (o *EnclosuresResourceInner) SetIomType(v string)`

SetIomType sets IomType field to given value.

### HasIomType

`func (o *EnclosuresResourceInner) HasIomType() bool`

HasIomType returns a boolean if a field has been set.

### GetIomTypeNumeric

`func (o *EnclosuresResourceInner) GetIomTypeNumeric() int64`

GetIomTypeNumeric returns the IomTypeNumeric field if non-nil, zero value otherwise.

### GetIomTypeNumericOk

`func (o *EnclosuresResourceInner) GetIomTypeNumericOk() (*int64, bool)`

GetIomTypeNumericOk returns a tuple with the IomTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIomTypeNumeric

`func (o *EnclosuresResourceInner) SetIomTypeNumeric(v int64)`

SetIomTypeNumeric sets IomTypeNumeric field to given value.

### HasIomTypeNumeric

`func (o *EnclosuresResourceInner) HasIomTypeNumeric() bool`

HasIomTypeNumeric returns a boolean if a field has been set.

### GetLocation

`func (o *EnclosuresResourceInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EnclosuresResourceInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EnclosuresResourceInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *EnclosuresResourceInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLocatorLed

`func (o *EnclosuresResourceInner) GetLocatorLed() string`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *EnclosuresResourceInner) GetLocatorLedOk() (*string, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *EnclosuresResourceInner) SetLocatorLed(v string)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *EnclosuresResourceInner) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### GetLocatorLedNumeric

`func (o *EnclosuresResourceInner) GetLocatorLedNumeric() int64`

GetLocatorLedNumeric returns the LocatorLedNumeric field if non-nil, zero value otherwise.

### GetLocatorLedNumericOk

`func (o *EnclosuresResourceInner) GetLocatorLedNumericOk() (*int64, bool)`

GetLocatorLedNumericOk returns a tuple with the LocatorLedNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLedNumeric

`func (o *EnclosuresResourceInner) SetLocatorLedNumeric(v int64)`

SetLocatorLedNumeric sets LocatorLedNumeric field to given value.

### HasLocatorLedNumeric

`func (o *EnclosuresResourceInner) HasLocatorLedNumeric() bool`

HasLocatorLedNumeric returns a boolean if a field has been set.

### GetMfgDate

`func (o *EnclosuresResourceInner) GetMfgDate() string`

GetMfgDate returns the MfgDate field if non-nil, zero value otherwise.

### GetMfgDateOk

`func (o *EnclosuresResourceInner) GetMfgDateOk() (*string, bool)`

GetMfgDateOk returns a tuple with the MfgDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgDate

`func (o *EnclosuresResourceInner) SetMfgDate(v string)`

SetMfgDate sets MfgDate field to given value.

### HasMfgDate

`func (o *EnclosuresResourceInner) HasMfgDate() bool`

HasMfgDate returns a boolean if a field has been set.

### GetMfgDateNumeric

`func (o *EnclosuresResourceInner) GetMfgDateNumeric() int64`

GetMfgDateNumeric returns the MfgDateNumeric field if non-nil, zero value otherwise.

### GetMfgDateNumericOk

`func (o *EnclosuresResourceInner) GetMfgDateNumericOk() (*int64, bool)`

GetMfgDateNumericOk returns a tuple with the MfgDateNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgDateNumeric

`func (o *EnclosuresResourceInner) SetMfgDateNumeric(v int64)`

SetMfgDateNumeric sets MfgDateNumeric field to given value.

### HasMfgDateNumeric

`func (o *EnclosuresResourceInner) HasMfgDateNumeric() bool`

HasMfgDateNumeric returns a boolean if a field has been set.

### GetMfgLocation

`func (o *EnclosuresResourceInner) GetMfgLocation() string`

GetMfgLocation returns the MfgLocation field if non-nil, zero value otherwise.

### GetMfgLocationOk

`func (o *EnclosuresResourceInner) GetMfgLocationOk() (*string, bool)`

GetMfgLocationOk returns a tuple with the MfgLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfgLocation

`func (o *EnclosuresResourceInner) SetMfgLocation(v string)`

SetMfgLocation sets MfgLocation field to given value.

### HasMfgLocation

`func (o *EnclosuresResourceInner) HasMfgLocation() bool`

HasMfgLocation returns a boolean if a field has been set.

### GetMidplaneRev

`func (o *EnclosuresResourceInner) GetMidplaneRev() int64`

GetMidplaneRev returns the MidplaneRev field if non-nil, zero value otherwise.

### GetMidplaneRevOk

`func (o *EnclosuresResourceInner) GetMidplaneRevOk() (*int64, bool)`

GetMidplaneRevOk returns a tuple with the MidplaneRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidplaneRev

`func (o *EnclosuresResourceInner) SetMidplaneRev(v int64)`

SetMidplaneRev sets MidplaneRev field to given value.

### HasMidplaneRev

`func (o *EnclosuresResourceInner) HasMidplaneRev() bool`

HasMidplaneRev returns a boolean if a field has been set.

### GetMidplaneSerialNumber

`func (o *EnclosuresResourceInner) GetMidplaneSerialNumber() string`

GetMidplaneSerialNumber returns the MidplaneSerialNumber field if non-nil, zero value otherwise.

### GetMidplaneSerialNumberOk

`func (o *EnclosuresResourceInner) GetMidplaneSerialNumberOk() (*string, bool)`

GetMidplaneSerialNumberOk returns a tuple with the MidplaneSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidplaneSerialNumber

`func (o *EnclosuresResourceInner) SetMidplaneSerialNumber(v string)`

SetMidplaneSerialNumber sets MidplaneSerialNumber field to given value.

### HasMidplaneSerialNumber

`func (o *EnclosuresResourceInner) HasMidplaneSerialNumber() bool`

HasMidplaneSerialNumber returns a boolean if a field has been set.

### GetMidplaneType

`func (o *EnclosuresResourceInner) GetMidplaneType() string`

GetMidplaneType returns the MidplaneType field if non-nil, zero value otherwise.

### GetMidplaneTypeOk

`func (o *EnclosuresResourceInner) GetMidplaneTypeOk() (*string, bool)`

GetMidplaneTypeOk returns a tuple with the MidplaneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidplaneType

`func (o *EnclosuresResourceInner) SetMidplaneType(v string)`

SetMidplaneType sets MidplaneType field to given value.

### HasMidplaneType

`func (o *EnclosuresResourceInner) HasMidplaneType() bool`

HasMidplaneType returns a boolean if a field has been set.

### GetMidplaneTypeNumeric

`func (o *EnclosuresResourceInner) GetMidplaneTypeNumeric() int64`

GetMidplaneTypeNumeric returns the MidplaneTypeNumeric field if non-nil, zero value otherwise.

### GetMidplaneTypeNumericOk

`func (o *EnclosuresResourceInner) GetMidplaneTypeNumericOk() (*int64, bool)`

GetMidplaneTypeNumericOk returns a tuple with the MidplaneTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidplaneTypeNumeric

`func (o *EnclosuresResourceInner) SetMidplaneTypeNumeric(v int64)`

SetMidplaneTypeNumeric sets MidplaneTypeNumeric field to given value.

### HasMidplaneTypeNumeric

`func (o *EnclosuresResourceInner) HasMidplaneTypeNumeric() bool`

HasMidplaneTypeNumeric returns a boolean if a field has been set.

### GetModel

`func (o *EnclosuresResourceInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EnclosuresResourceInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EnclosuresResourceInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EnclosuresResourceInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *EnclosuresResourceInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnclosuresResourceInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnclosuresResourceInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnclosuresResourceInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfCoolingsElements

`func (o *EnclosuresResourceInner) GetNumberOfCoolingsElements() int64`

GetNumberOfCoolingsElements returns the NumberOfCoolingsElements field if non-nil, zero value otherwise.

### GetNumberOfCoolingsElementsOk

`func (o *EnclosuresResourceInner) GetNumberOfCoolingsElementsOk() (*int64, bool)`

GetNumberOfCoolingsElementsOk returns a tuple with the NumberOfCoolingsElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCoolingsElements

`func (o *EnclosuresResourceInner) SetNumberOfCoolingsElements(v int64)`

SetNumberOfCoolingsElements sets NumberOfCoolingsElements field to given value.

### HasNumberOfCoolingsElements

`func (o *EnclosuresResourceInner) HasNumberOfCoolingsElements() bool`

HasNumberOfCoolingsElements returns a boolean if a field has been set.

### GetNumberOfDisks

`func (o *EnclosuresResourceInner) GetNumberOfDisks() int64`

GetNumberOfDisks returns the NumberOfDisks field if non-nil, zero value otherwise.

### GetNumberOfDisksOk

`func (o *EnclosuresResourceInner) GetNumberOfDisksOk() (*int64, bool)`

GetNumberOfDisksOk returns a tuple with the NumberOfDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDisks

`func (o *EnclosuresResourceInner) SetNumberOfDisks(v int64)`

SetNumberOfDisks sets NumberOfDisks field to given value.

### HasNumberOfDisks

`func (o *EnclosuresResourceInner) HasNumberOfDisks() bool`

HasNumberOfDisks returns a boolean if a field has been set.

### GetNumberOfPowerSupplies

`func (o *EnclosuresResourceInner) GetNumberOfPowerSupplies() int64`

GetNumberOfPowerSupplies returns the NumberOfPowerSupplies field if non-nil, zero value otherwise.

### GetNumberOfPowerSuppliesOk

`func (o *EnclosuresResourceInner) GetNumberOfPowerSuppliesOk() (*int64, bool)`

GetNumberOfPowerSuppliesOk returns a tuple with the NumberOfPowerSupplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPowerSupplies

`func (o *EnclosuresResourceInner) SetNumberOfPowerSupplies(v int64)`

SetNumberOfPowerSupplies sets NumberOfPowerSupplies field to given value.

### HasNumberOfPowerSupplies

`func (o *EnclosuresResourceInner) HasNumberOfPowerSupplies() bool`

HasNumberOfPowerSupplies returns a boolean if a field has been set.

### GetPartNumber

`func (o *EnclosuresResourceInner) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *EnclosuresResourceInner) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *EnclosuresResourceInner) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *EnclosuresResourceInner) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPcie2Capable

`func (o *EnclosuresResourceInner) GetPcie2Capable() string`

GetPcie2Capable returns the Pcie2Capable field if non-nil, zero value otherwise.

### GetPcie2CapableOk

`func (o *EnclosuresResourceInner) GetPcie2CapableOk() (*string, bool)`

GetPcie2CapableOk returns a tuple with the Pcie2Capable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcie2Capable

`func (o *EnclosuresResourceInner) SetPcie2Capable(v string)`

SetPcie2Capable sets Pcie2Capable field to given value.

### HasPcie2Capable

`func (o *EnclosuresResourceInner) HasPcie2Capable() bool`

HasPcie2Capable returns a boolean if a field has been set.

### GetPcie2CapableNumeric

`func (o *EnclosuresResourceInner) GetPcie2CapableNumeric() int64`

GetPcie2CapableNumeric returns the Pcie2CapableNumeric field if non-nil, zero value otherwise.

### GetPcie2CapableNumericOk

`func (o *EnclosuresResourceInner) GetPcie2CapableNumericOk() (*int64, bool)`

GetPcie2CapableNumericOk returns a tuple with the Pcie2CapableNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcie2CapableNumeric

`func (o *EnclosuresResourceInner) SetPcie2CapableNumeric(v int64)`

SetPcie2CapableNumeric sets Pcie2CapableNumeric field to given value.

### HasPcie2CapableNumeric

`func (o *EnclosuresResourceInner) HasPcie2CapableNumeric() bool`

HasPcie2CapableNumeric returns a boolean if a field has been set.

### GetPlatformType

`func (o *EnclosuresResourceInner) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *EnclosuresResourceInner) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *EnclosuresResourceInner) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *EnclosuresResourceInner) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetPlatformTypeNumeric

`func (o *EnclosuresResourceInner) GetPlatformTypeNumeric() int64`

GetPlatformTypeNumeric returns the PlatformTypeNumeric field if non-nil, zero value otherwise.

### GetPlatformTypeNumericOk

`func (o *EnclosuresResourceInner) GetPlatformTypeNumericOk() (*int64, bool)`

GetPlatformTypeNumericOk returns a tuple with the PlatformTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformTypeNumeric

`func (o *EnclosuresResourceInner) SetPlatformTypeNumeric(v int64)`

SetPlatformTypeNumeric sets PlatformTypeNumeric field to given value.

### HasPlatformTypeNumeric

`func (o *EnclosuresResourceInner) HasPlatformTypeNumeric() bool`

HasPlatformTypeNumeric returns a boolean if a field has been set.

### GetRackNumber

`func (o *EnclosuresResourceInner) GetRackNumber() int64`

GetRackNumber returns the RackNumber field if non-nil, zero value otherwise.

### GetRackNumberOk

`func (o *EnclosuresResourceInner) GetRackNumberOk() (*int64, bool)`

GetRackNumberOk returns a tuple with the RackNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackNumber

`func (o *EnclosuresResourceInner) SetRackNumber(v int64)`

SetRackNumber sets RackNumber field to given value.

### HasRackNumber

`func (o *EnclosuresResourceInner) HasRackNumber() bool`

HasRackNumber returns a boolean if a field has been set.

### GetRackPosition

`func (o *EnclosuresResourceInner) GetRackPosition() int64`

GetRackPosition returns the RackPosition field if non-nil, zero value otherwise.

### GetRackPositionOk

`func (o *EnclosuresResourceInner) GetRackPositionOk() (*int64, bool)`

GetRackPositionOk returns a tuple with the RackPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackPosition

`func (o *EnclosuresResourceInner) SetRackPosition(v int64)`

SetRackPosition sets RackPosition field to given value.

### HasRackPosition

`func (o *EnclosuresResourceInner) HasRackPosition() bool`

HasRackPosition returns a boolean if a field has been set.

### GetRevision

`func (o *EnclosuresResourceInner) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *EnclosuresResourceInner) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *EnclosuresResourceInner) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *EnclosuresResourceInner) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRows

`func (o *EnclosuresResourceInner) GetRows() int64`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *EnclosuresResourceInner) GetRowsOk() (*int64, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *EnclosuresResourceInner) SetRows(v int64)`

SetRows sets Rows field to given value.

### HasRows

`func (o *EnclosuresResourceInner) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetSlots

`func (o *EnclosuresResourceInner) GetSlots() int64`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *EnclosuresResourceInner) GetSlotsOk() (*int64, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *EnclosuresResourceInner) SetSlots(v int64)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *EnclosuresResourceInner) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### GetStatus

`func (o *EnclosuresResourceInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnclosuresResourceInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnclosuresResourceInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnclosuresResourceInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusNumeric

`func (o *EnclosuresResourceInner) GetStatusNumeric() int64`

GetStatusNumeric returns the StatusNumeric field if non-nil, zero value otherwise.

### GetStatusNumericOk

`func (o *EnclosuresResourceInner) GetStatusNumericOk() (*int64, bool)`

GetStatusNumericOk returns a tuple with the StatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusNumeric

`func (o *EnclosuresResourceInner) SetStatusNumeric(v int64)`

SetStatusNumeric sets StatusNumeric field to given value.

### HasStatusNumeric

`func (o *EnclosuresResourceInner) HasStatusNumeric() bool`

HasStatusNumeric returns a boolean if a field has been set.

### GetType

`func (o *EnclosuresResourceInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnclosuresResourceInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnclosuresResourceInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnclosuresResourceInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeNumeric

`func (o *EnclosuresResourceInner) GetTypeNumeric() int64`

GetTypeNumeric returns the TypeNumeric field if non-nil, zero value otherwise.

### GetTypeNumericOk

`func (o *EnclosuresResourceInner) GetTypeNumericOk() (*int64, bool)`

GetTypeNumericOk returns a tuple with the TypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeNumeric

`func (o *EnclosuresResourceInner) SetTypeNumeric(v int64)`

SetTypeNumeric sets TypeNumeric field to given value.

### HasTypeNumeric

`func (o *EnclosuresResourceInner) HasTypeNumeric() bool`

HasTypeNumeric returns a boolean if a field has been set.

### GetUrl

`func (o *EnclosuresResourceInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EnclosuresResourceInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EnclosuresResourceInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EnclosuresResourceInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVendor

`func (o *EnclosuresResourceInner) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *EnclosuresResourceInner) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *EnclosuresResourceInner) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *EnclosuresResourceInner) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetControllers

`func (o *EnclosuresResourceInner) GetControllers() []ControllersResourceInner`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *EnclosuresResourceInner) GetControllersOk() (*[]ControllersResourceInner, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *EnclosuresResourceInner) SetControllers(v []ControllersResourceInner)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *EnclosuresResourceInner) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetPowerSupplies

`func (o *EnclosuresResourceInner) GetPowerSupplies() []PowerSuppliesResourceInner`

GetPowerSupplies returns the PowerSupplies field if non-nil, zero value otherwise.

### GetPowerSuppliesOk

`func (o *EnclosuresResourceInner) GetPowerSuppliesOk() (*[]PowerSuppliesResourceInner, bool)`

GetPowerSuppliesOk returns a tuple with the PowerSupplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSupplies

`func (o *EnclosuresResourceInner) SetPowerSupplies(v []PowerSuppliesResourceInner)`

SetPowerSupplies sets PowerSupplies field to given value.

### HasPowerSupplies

`func (o *EnclosuresResourceInner) HasPowerSupplies() bool`

HasPowerSupplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


