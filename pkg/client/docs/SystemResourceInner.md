# SystemResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**ConfigurationSelector** | Pointer to **string** |  | [optional] 
**CurrentNodeWwn** | Pointer to **string** |  | [optional] 
**EnclosureCount** | Pointer to **int32** |  | [optional] 
**FdeSecurityStatus** | Pointer to **string** |  | [optional] 
**FdeSecurityStatusNumeric** | Pointer to **int32** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**HealthNumeric** | Pointer to **int32** |  | [optional] 
**HealthReason** | Pointer to **string** |  | [optional] 
**MidplaneSerialNumber** | Pointer to **string** | Serial number of the enclosure | [optional] 
**OtherMCStatus** | Pointer to **string** | Identifies the availability of the partner MC | [optional] 
**OtherMCStatusNumeric** | Pointer to **int32** | Identifies the availability of the partner MC( In numeric form ) | [optional] 
**PfuStatus** | Pointer to **string** |  | [optional] 
**PfuStatusNumeric** | Pointer to **int32** |  | [optional] 
**PlatformBrand** | Pointer to **string** | HW Platform Brand | [optional] 
**PlatformBrandNumeric** | Pointer to **int32** | HW Platform Brand( In numeric form ) | [optional] 
**PlatformType** | Pointer to **string** | HW Platform Type | [optional] 
**PlatformTypeNumeric** | Pointer to **int32** | HW Platform Type( In numeric form ) | [optional] 
**ProductBrand** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**ScsiProductId** | Pointer to **string** | SCSI Product ID | [optional] 
**ScsiVendorId** | Pointer to **string** | SCSI vendor name | [optional] 
**SecuritySystemManagement** | Pointer to **string** |  | [optional] 
**SecuritySystemManagementNumeric** | Pointer to **int32** |  | [optional] 
**SupportedLocales** | Pointer to **string** |  | [optional] 
**SystemContact** | Pointer to **string** | User-defined contact for this system | [optional] 
**SystemInformation** | Pointer to **string** |  | [optional] 
**SystemLocation** | Pointer to **string** | User-defined location of the system | [optional] 
**SystemName** | Pointer to **string** | User-defined name for the system | [optional] 
**Url** | Pointer to **string** | The resource URL | [optional] 
**VendorName** | Pointer to **string** |  | [optional] 
**Redundancy** | Pointer to [**[]RedundancyResourceInner**](RedundancyResourceInner.md) |  | [optional] 

## Methods

### NewSystemResourceInner

`func NewSystemResourceInner() *SystemResourceInner`

NewSystemResourceInner instantiates a new SystemResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemResourceInnerWithDefaults

`func NewSystemResourceInnerWithDefaults() *SystemResourceInner`

NewSystemResourceInnerWithDefaults instantiates a new SystemResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *SystemResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *SystemResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *SystemResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *SystemResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *SystemResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SystemResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SystemResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SystemResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetConfigurationSelector

`func (o *SystemResourceInner) GetConfigurationSelector() string`

GetConfigurationSelector returns the ConfigurationSelector field if non-nil, zero value otherwise.

### GetConfigurationSelectorOk

`func (o *SystemResourceInner) GetConfigurationSelectorOk() (*string, bool)`

GetConfigurationSelectorOk returns a tuple with the ConfigurationSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationSelector

`func (o *SystemResourceInner) SetConfigurationSelector(v string)`

SetConfigurationSelector sets ConfigurationSelector field to given value.

### HasConfigurationSelector

`func (o *SystemResourceInner) HasConfigurationSelector() bool`

HasConfigurationSelector returns a boolean if a field has been set.

### GetCurrentNodeWwn

`func (o *SystemResourceInner) GetCurrentNodeWwn() string`

GetCurrentNodeWwn returns the CurrentNodeWwn field if non-nil, zero value otherwise.

### GetCurrentNodeWwnOk

`func (o *SystemResourceInner) GetCurrentNodeWwnOk() (*string, bool)`

GetCurrentNodeWwnOk returns a tuple with the CurrentNodeWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNodeWwn

`func (o *SystemResourceInner) SetCurrentNodeWwn(v string)`

SetCurrentNodeWwn sets CurrentNodeWwn field to given value.

### HasCurrentNodeWwn

`func (o *SystemResourceInner) HasCurrentNodeWwn() bool`

HasCurrentNodeWwn returns a boolean if a field has been set.

### GetEnclosureCount

`func (o *SystemResourceInner) GetEnclosureCount() int32`

GetEnclosureCount returns the EnclosureCount field if non-nil, zero value otherwise.

### GetEnclosureCountOk

`func (o *SystemResourceInner) GetEnclosureCountOk() (*int32, bool)`

GetEnclosureCountOk returns a tuple with the EnclosureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureCount

`func (o *SystemResourceInner) SetEnclosureCount(v int32)`

SetEnclosureCount sets EnclosureCount field to given value.

### HasEnclosureCount

`func (o *SystemResourceInner) HasEnclosureCount() bool`

HasEnclosureCount returns a boolean if a field has been set.

### GetFdeSecurityStatus

`func (o *SystemResourceInner) GetFdeSecurityStatus() string`

GetFdeSecurityStatus returns the FdeSecurityStatus field if non-nil, zero value otherwise.

### GetFdeSecurityStatusOk

`func (o *SystemResourceInner) GetFdeSecurityStatusOk() (*string, bool)`

GetFdeSecurityStatusOk returns a tuple with the FdeSecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdeSecurityStatus

`func (o *SystemResourceInner) SetFdeSecurityStatus(v string)`

SetFdeSecurityStatus sets FdeSecurityStatus field to given value.

### HasFdeSecurityStatus

`func (o *SystemResourceInner) HasFdeSecurityStatus() bool`

HasFdeSecurityStatus returns a boolean if a field has been set.

### GetFdeSecurityStatusNumeric

`func (o *SystemResourceInner) GetFdeSecurityStatusNumeric() int32`

GetFdeSecurityStatusNumeric returns the FdeSecurityStatusNumeric field if non-nil, zero value otherwise.

### GetFdeSecurityStatusNumericOk

`func (o *SystemResourceInner) GetFdeSecurityStatusNumericOk() (*int32, bool)`

GetFdeSecurityStatusNumericOk returns a tuple with the FdeSecurityStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdeSecurityStatusNumeric

`func (o *SystemResourceInner) SetFdeSecurityStatusNumeric(v int32)`

SetFdeSecurityStatusNumeric sets FdeSecurityStatusNumeric field to given value.

### HasFdeSecurityStatusNumeric

`func (o *SystemResourceInner) HasFdeSecurityStatusNumeric() bool`

HasFdeSecurityStatusNumeric returns a boolean if a field has been set.

### GetHealth

`func (o *SystemResourceInner) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *SystemResourceInner) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *SystemResourceInner) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *SystemResourceInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHealthNumeric

`func (o *SystemResourceInner) GetHealthNumeric() int32`

GetHealthNumeric returns the HealthNumeric field if non-nil, zero value otherwise.

### GetHealthNumericOk

`func (o *SystemResourceInner) GetHealthNumericOk() (*int32, bool)`

GetHealthNumericOk returns a tuple with the HealthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthNumeric

`func (o *SystemResourceInner) SetHealthNumeric(v int32)`

SetHealthNumeric sets HealthNumeric field to given value.

### HasHealthNumeric

`func (o *SystemResourceInner) HasHealthNumeric() bool`

HasHealthNumeric returns a boolean if a field has been set.

### GetHealthReason

`func (o *SystemResourceInner) GetHealthReason() string`

GetHealthReason returns the HealthReason field if non-nil, zero value otherwise.

### GetHealthReasonOk

`func (o *SystemResourceInner) GetHealthReasonOk() (*string, bool)`

GetHealthReasonOk returns a tuple with the HealthReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReason

`func (o *SystemResourceInner) SetHealthReason(v string)`

SetHealthReason sets HealthReason field to given value.

### HasHealthReason

`func (o *SystemResourceInner) HasHealthReason() bool`

HasHealthReason returns a boolean if a field has been set.

### GetMidplaneSerialNumber

`func (o *SystemResourceInner) GetMidplaneSerialNumber() string`

GetMidplaneSerialNumber returns the MidplaneSerialNumber field if non-nil, zero value otherwise.

### GetMidplaneSerialNumberOk

`func (o *SystemResourceInner) GetMidplaneSerialNumberOk() (*string, bool)`

GetMidplaneSerialNumberOk returns a tuple with the MidplaneSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidplaneSerialNumber

`func (o *SystemResourceInner) SetMidplaneSerialNumber(v string)`

SetMidplaneSerialNumber sets MidplaneSerialNumber field to given value.

### HasMidplaneSerialNumber

`func (o *SystemResourceInner) HasMidplaneSerialNumber() bool`

HasMidplaneSerialNumber returns a boolean if a field has been set.

### GetOtherMCStatus

`func (o *SystemResourceInner) GetOtherMCStatus() string`

GetOtherMCStatus returns the OtherMCStatus field if non-nil, zero value otherwise.

### GetOtherMCStatusOk

`func (o *SystemResourceInner) GetOtherMCStatusOk() (*string, bool)`

GetOtherMCStatusOk returns a tuple with the OtherMCStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherMCStatus

`func (o *SystemResourceInner) SetOtherMCStatus(v string)`

SetOtherMCStatus sets OtherMCStatus field to given value.

### HasOtherMCStatus

`func (o *SystemResourceInner) HasOtherMCStatus() bool`

HasOtherMCStatus returns a boolean if a field has been set.

### GetOtherMCStatusNumeric

`func (o *SystemResourceInner) GetOtherMCStatusNumeric() int32`

GetOtherMCStatusNumeric returns the OtherMCStatusNumeric field if non-nil, zero value otherwise.

### GetOtherMCStatusNumericOk

`func (o *SystemResourceInner) GetOtherMCStatusNumericOk() (*int32, bool)`

GetOtherMCStatusNumericOk returns a tuple with the OtherMCStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherMCStatusNumeric

`func (o *SystemResourceInner) SetOtherMCStatusNumeric(v int32)`

SetOtherMCStatusNumeric sets OtherMCStatusNumeric field to given value.

### HasOtherMCStatusNumeric

`func (o *SystemResourceInner) HasOtherMCStatusNumeric() bool`

HasOtherMCStatusNumeric returns a boolean if a field has been set.

### GetPfuStatus

`func (o *SystemResourceInner) GetPfuStatus() string`

GetPfuStatus returns the PfuStatus field if non-nil, zero value otherwise.

### GetPfuStatusOk

`func (o *SystemResourceInner) GetPfuStatusOk() (*string, bool)`

GetPfuStatusOk returns a tuple with the PfuStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfuStatus

`func (o *SystemResourceInner) SetPfuStatus(v string)`

SetPfuStatus sets PfuStatus field to given value.

### HasPfuStatus

`func (o *SystemResourceInner) HasPfuStatus() bool`

HasPfuStatus returns a boolean if a field has been set.

### GetPfuStatusNumeric

`func (o *SystemResourceInner) GetPfuStatusNumeric() int32`

GetPfuStatusNumeric returns the PfuStatusNumeric field if non-nil, zero value otherwise.

### GetPfuStatusNumericOk

`func (o *SystemResourceInner) GetPfuStatusNumericOk() (*int32, bool)`

GetPfuStatusNumericOk returns a tuple with the PfuStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfuStatusNumeric

`func (o *SystemResourceInner) SetPfuStatusNumeric(v int32)`

SetPfuStatusNumeric sets PfuStatusNumeric field to given value.

### HasPfuStatusNumeric

`func (o *SystemResourceInner) HasPfuStatusNumeric() bool`

HasPfuStatusNumeric returns a boolean if a field has been set.

### GetPlatformBrand

`func (o *SystemResourceInner) GetPlatformBrand() string`

GetPlatformBrand returns the PlatformBrand field if non-nil, zero value otherwise.

### GetPlatformBrandOk

`func (o *SystemResourceInner) GetPlatformBrandOk() (*string, bool)`

GetPlatformBrandOk returns a tuple with the PlatformBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformBrand

`func (o *SystemResourceInner) SetPlatformBrand(v string)`

SetPlatformBrand sets PlatformBrand field to given value.

### HasPlatformBrand

`func (o *SystemResourceInner) HasPlatformBrand() bool`

HasPlatformBrand returns a boolean if a field has been set.

### GetPlatformBrandNumeric

`func (o *SystemResourceInner) GetPlatformBrandNumeric() int32`

GetPlatformBrandNumeric returns the PlatformBrandNumeric field if non-nil, zero value otherwise.

### GetPlatformBrandNumericOk

`func (o *SystemResourceInner) GetPlatformBrandNumericOk() (*int32, bool)`

GetPlatformBrandNumericOk returns a tuple with the PlatformBrandNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformBrandNumeric

`func (o *SystemResourceInner) SetPlatformBrandNumeric(v int32)`

SetPlatformBrandNumeric sets PlatformBrandNumeric field to given value.

### HasPlatformBrandNumeric

`func (o *SystemResourceInner) HasPlatformBrandNumeric() bool`

HasPlatformBrandNumeric returns a boolean if a field has been set.

### GetPlatformType

`func (o *SystemResourceInner) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *SystemResourceInner) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *SystemResourceInner) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *SystemResourceInner) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetPlatformTypeNumeric

`func (o *SystemResourceInner) GetPlatformTypeNumeric() int32`

GetPlatformTypeNumeric returns the PlatformTypeNumeric field if non-nil, zero value otherwise.

### GetPlatformTypeNumericOk

`func (o *SystemResourceInner) GetPlatformTypeNumericOk() (*int32, bool)`

GetPlatformTypeNumericOk returns a tuple with the PlatformTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformTypeNumeric

`func (o *SystemResourceInner) SetPlatformTypeNumeric(v int32)`

SetPlatformTypeNumeric sets PlatformTypeNumeric field to given value.

### HasPlatformTypeNumeric

`func (o *SystemResourceInner) HasPlatformTypeNumeric() bool`

HasPlatformTypeNumeric returns a boolean if a field has been set.

### GetProductBrand

`func (o *SystemResourceInner) GetProductBrand() string`

GetProductBrand returns the ProductBrand field if non-nil, zero value otherwise.

### GetProductBrandOk

`func (o *SystemResourceInner) GetProductBrandOk() (*string, bool)`

GetProductBrandOk returns a tuple with the ProductBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBrand

`func (o *SystemResourceInner) SetProductBrand(v string)`

SetProductBrand sets ProductBrand field to given value.

### HasProductBrand

`func (o *SystemResourceInner) HasProductBrand() bool`

HasProductBrand returns a boolean if a field has been set.

### GetProductId

`func (o *SystemResourceInner) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SystemResourceInner) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SystemResourceInner) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *SystemResourceInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetScsiProductId

`func (o *SystemResourceInner) GetScsiProductId() string`

GetScsiProductId returns the ScsiProductId field if non-nil, zero value otherwise.

### GetScsiProductIdOk

`func (o *SystemResourceInner) GetScsiProductIdOk() (*string, bool)`

GetScsiProductIdOk returns a tuple with the ScsiProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsiProductId

`func (o *SystemResourceInner) SetScsiProductId(v string)`

SetScsiProductId sets ScsiProductId field to given value.

### HasScsiProductId

`func (o *SystemResourceInner) HasScsiProductId() bool`

HasScsiProductId returns a boolean if a field has been set.

### GetScsiVendorId

`func (o *SystemResourceInner) GetScsiVendorId() string`

GetScsiVendorId returns the ScsiVendorId field if non-nil, zero value otherwise.

### GetScsiVendorIdOk

`func (o *SystemResourceInner) GetScsiVendorIdOk() (*string, bool)`

GetScsiVendorIdOk returns a tuple with the ScsiVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsiVendorId

`func (o *SystemResourceInner) SetScsiVendorId(v string)`

SetScsiVendorId sets ScsiVendorId field to given value.

### HasScsiVendorId

`func (o *SystemResourceInner) HasScsiVendorId() bool`

HasScsiVendorId returns a boolean if a field has been set.

### GetSecuritySystemManagement

`func (o *SystemResourceInner) GetSecuritySystemManagement() string`

GetSecuritySystemManagement returns the SecuritySystemManagement field if non-nil, zero value otherwise.

### GetSecuritySystemManagementOk

`func (o *SystemResourceInner) GetSecuritySystemManagementOk() (*string, bool)`

GetSecuritySystemManagementOk returns a tuple with the SecuritySystemManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySystemManagement

`func (o *SystemResourceInner) SetSecuritySystemManagement(v string)`

SetSecuritySystemManagement sets SecuritySystemManagement field to given value.

### HasSecuritySystemManagement

`func (o *SystemResourceInner) HasSecuritySystemManagement() bool`

HasSecuritySystemManagement returns a boolean if a field has been set.

### GetSecuritySystemManagementNumeric

`func (o *SystemResourceInner) GetSecuritySystemManagementNumeric() int32`

GetSecuritySystemManagementNumeric returns the SecuritySystemManagementNumeric field if non-nil, zero value otherwise.

### GetSecuritySystemManagementNumericOk

`func (o *SystemResourceInner) GetSecuritySystemManagementNumericOk() (*int32, bool)`

GetSecuritySystemManagementNumericOk returns a tuple with the SecuritySystemManagementNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySystemManagementNumeric

`func (o *SystemResourceInner) SetSecuritySystemManagementNumeric(v int32)`

SetSecuritySystemManagementNumeric sets SecuritySystemManagementNumeric field to given value.

### HasSecuritySystemManagementNumeric

`func (o *SystemResourceInner) HasSecuritySystemManagementNumeric() bool`

HasSecuritySystemManagementNumeric returns a boolean if a field has been set.

### GetSupportedLocales

`func (o *SystemResourceInner) GetSupportedLocales() string`

GetSupportedLocales returns the SupportedLocales field if non-nil, zero value otherwise.

### GetSupportedLocalesOk

`func (o *SystemResourceInner) GetSupportedLocalesOk() (*string, bool)`

GetSupportedLocalesOk returns a tuple with the SupportedLocales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedLocales

`func (o *SystemResourceInner) SetSupportedLocales(v string)`

SetSupportedLocales sets SupportedLocales field to given value.

### HasSupportedLocales

`func (o *SystemResourceInner) HasSupportedLocales() bool`

HasSupportedLocales returns a boolean if a field has been set.

### GetSystemContact

`func (o *SystemResourceInner) GetSystemContact() string`

GetSystemContact returns the SystemContact field if non-nil, zero value otherwise.

### GetSystemContactOk

`func (o *SystemResourceInner) GetSystemContactOk() (*string, bool)`

GetSystemContactOk returns a tuple with the SystemContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemContact

`func (o *SystemResourceInner) SetSystemContact(v string)`

SetSystemContact sets SystemContact field to given value.

### HasSystemContact

`func (o *SystemResourceInner) HasSystemContact() bool`

HasSystemContact returns a boolean if a field has been set.

### GetSystemInformation

`func (o *SystemResourceInner) GetSystemInformation() string`

GetSystemInformation returns the SystemInformation field if non-nil, zero value otherwise.

### GetSystemInformationOk

`func (o *SystemResourceInner) GetSystemInformationOk() (*string, bool)`

GetSystemInformationOk returns a tuple with the SystemInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInformation

`func (o *SystemResourceInner) SetSystemInformation(v string)`

SetSystemInformation sets SystemInformation field to given value.

### HasSystemInformation

`func (o *SystemResourceInner) HasSystemInformation() bool`

HasSystemInformation returns a boolean if a field has been set.

### GetSystemLocation

`func (o *SystemResourceInner) GetSystemLocation() string`

GetSystemLocation returns the SystemLocation field if non-nil, zero value otherwise.

### GetSystemLocationOk

`func (o *SystemResourceInner) GetSystemLocationOk() (*string, bool)`

GetSystemLocationOk returns a tuple with the SystemLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLocation

`func (o *SystemResourceInner) SetSystemLocation(v string)`

SetSystemLocation sets SystemLocation field to given value.

### HasSystemLocation

`func (o *SystemResourceInner) HasSystemLocation() bool`

HasSystemLocation returns a boolean if a field has been set.

### GetSystemName

`func (o *SystemResourceInner) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *SystemResourceInner) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *SystemResourceInner) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *SystemResourceInner) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetUrl

`func (o *SystemResourceInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SystemResourceInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SystemResourceInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SystemResourceInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVendorName

`func (o *SystemResourceInner) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *SystemResourceInner) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *SystemResourceInner) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *SystemResourceInner) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetRedundancy

`func (o *SystemResourceInner) GetRedundancy() []RedundancyResourceInner`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *SystemResourceInner) GetRedundancyOk() (*[]RedundancyResourceInner, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *SystemResourceInner) SetRedundancy(v []RedundancyResourceInner)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *SystemResourceInner) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


