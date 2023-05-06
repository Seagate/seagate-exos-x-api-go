# VersionsResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**BackplaneType** | Pointer to **int32** |  | [optional] 
**BuildDate** | Pointer to **string** |  | [optional] 
**BundleBaseVersion** | Pointer to **string** |  | [optional] 
**BundleStatus** | Pointer to **string** |  | [optional] 
**BundleStatusNumeric** | Pointer to **int32** |  | [optional] 
**BundleVersion** | Pointer to **string** |  | [optional] 
**BundleVersionOnly** | Pointer to **string** |  | [optional] 
**CapiVersion** | Pointer to **string** |  | [optional] 
**CtkVersion** | Pointer to **string** | Customization Toolkit (CTK) Version | [optional] 
**DiskChannelRevision** | Pointer to **int32** |  | [optional] 
**EcFw** | Pointer to **string** | The Expander Controller firmware version | [optional] 
**FwDefaultPlatformBrand** | Pointer to **string** | Default hardware platform brand of the firmware | [optional] 
**FwDefaultPlatformBrandNumeric** | Pointer to **int32** | Default hardware platform brand of the firmware( In numeric form ) | [optional] 
**GemVersion** | Pointer to **string** |  | [optional] 
**HimModel** | Pointer to **string** |  | [optional] 
**HimRev** | Pointer to **string** |  | [optional] 
**HostChannelRevision** | Pointer to **int32** |  | [optional] 
**HwRev** | Pointer to **string** | The main hardware version of the controller | [optional] 
**McBaseFw** | Pointer to **string** | The Management Controller main (&#39;app&#39;) baselevel firmware version | [optional] 
**McFw** | Pointer to **string** | The Management Controller main (&#39;app&#39;) firmware version | [optional] 
**McLoader** | Pointer to **string** | The Management Controller loader firmware version | [optional] 
**McosVersion** | Pointer to **string** | MC Operating System (OS) Version | [optional] 
**MrcVersion** | Pointer to **string** | The SC Boot Memory Reference Code (MRC) version | [optional] 
**PcieSwitchBackendConfigurationVersion** | Pointer to **string** |  | [optional] 
**PcieSwitchBackendFirmwareVersion** | Pointer to **string** |  | [optional] 
**PcieSwitchFrontendConfigurationVersion** | Pointer to **string** |  | [optional] 
**PcieSwitchFrontendFirmwareVersion** | Pointer to **string** |  | [optional] 
**PldRev** | Pointer to **string** | The CPLD code version | [optional] 
**PmCpldVersion** | Pointer to **string** |  | [optional] 
**PrmVersion** | Pointer to **string** | PRM CPLD code version | [optional] 
**PubsVersion** | Pointer to **string** | CLI Help version | [optional] 
**ScBaselevel** | Pointer to **string** |  | [optional] 
**ScCpuType** | Pointer to **string** | The Storage Controller processor type | [optional] 
**ScFuVersion** | Pointer to **string** | ASIC Controller Version | [optional] 
**ScFw** | Pointer to **string** | The Storage Controller main (&#39;app&#39;) firmware version | [optional] 
**ScLoader** | Pointer to **string** | The Storage Controller loader firmware version | [optional] 
**ScMemory** | Pointer to **string** | The Storage Controller memory-controller FPGA code version | [optional] 
**TranslationVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewVersionsResourceInner

`func NewVersionsResourceInner() *VersionsResourceInner`

NewVersionsResourceInner instantiates a new VersionsResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionsResourceInnerWithDefaults

`func NewVersionsResourceInnerWithDefaults() *VersionsResourceInner`

NewVersionsResourceInnerWithDefaults instantiates a new VersionsResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *VersionsResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *VersionsResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *VersionsResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *VersionsResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *VersionsResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VersionsResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VersionsResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *VersionsResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetBackplaneType

`func (o *VersionsResourceInner) GetBackplaneType() int32`

GetBackplaneType returns the BackplaneType field if non-nil, zero value otherwise.

### GetBackplaneTypeOk

`func (o *VersionsResourceInner) GetBackplaneTypeOk() (*int32, bool)`

GetBackplaneTypeOk returns a tuple with the BackplaneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackplaneType

`func (o *VersionsResourceInner) SetBackplaneType(v int32)`

SetBackplaneType sets BackplaneType field to given value.

### HasBackplaneType

`func (o *VersionsResourceInner) HasBackplaneType() bool`

HasBackplaneType returns a boolean if a field has been set.

### GetBuildDate

`func (o *VersionsResourceInner) GetBuildDate() string`

GetBuildDate returns the BuildDate field if non-nil, zero value otherwise.

### GetBuildDateOk

`func (o *VersionsResourceInner) GetBuildDateOk() (*string, bool)`

GetBuildDateOk returns a tuple with the BuildDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDate

`func (o *VersionsResourceInner) SetBuildDate(v string)`

SetBuildDate sets BuildDate field to given value.

### HasBuildDate

`func (o *VersionsResourceInner) HasBuildDate() bool`

HasBuildDate returns a boolean if a field has been set.

### GetBundleBaseVersion

`func (o *VersionsResourceInner) GetBundleBaseVersion() string`

GetBundleBaseVersion returns the BundleBaseVersion field if non-nil, zero value otherwise.

### GetBundleBaseVersionOk

`func (o *VersionsResourceInner) GetBundleBaseVersionOk() (*string, bool)`

GetBundleBaseVersionOk returns a tuple with the BundleBaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleBaseVersion

`func (o *VersionsResourceInner) SetBundleBaseVersion(v string)`

SetBundleBaseVersion sets BundleBaseVersion field to given value.

### HasBundleBaseVersion

`func (o *VersionsResourceInner) HasBundleBaseVersion() bool`

HasBundleBaseVersion returns a boolean if a field has been set.

### GetBundleStatus

`func (o *VersionsResourceInner) GetBundleStatus() string`

GetBundleStatus returns the BundleStatus field if non-nil, zero value otherwise.

### GetBundleStatusOk

`func (o *VersionsResourceInner) GetBundleStatusOk() (*string, bool)`

GetBundleStatusOk returns a tuple with the BundleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleStatus

`func (o *VersionsResourceInner) SetBundleStatus(v string)`

SetBundleStatus sets BundleStatus field to given value.

### HasBundleStatus

`func (o *VersionsResourceInner) HasBundleStatus() bool`

HasBundleStatus returns a boolean if a field has been set.

### GetBundleStatusNumeric

`func (o *VersionsResourceInner) GetBundleStatusNumeric() int32`

GetBundleStatusNumeric returns the BundleStatusNumeric field if non-nil, zero value otherwise.

### GetBundleStatusNumericOk

`func (o *VersionsResourceInner) GetBundleStatusNumericOk() (*int32, bool)`

GetBundleStatusNumericOk returns a tuple with the BundleStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleStatusNumeric

`func (o *VersionsResourceInner) SetBundleStatusNumeric(v int32)`

SetBundleStatusNumeric sets BundleStatusNumeric field to given value.

### HasBundleStatusNumeric

`func (o *VersionsResourceInner) HasBundleStatusNumeric() bool`

HasBundleStatusNumeric returns a boolean if a field has been set.

### GetBundleVersion

`func (o *VersionsResourceInner) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *VersionsResourceInner) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *VersionsResourceInner) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.

### HasBundleVersion

`func (o *VersionsResourceInner) HasBundleVersion() bool`

HasBundleVersion returns a boolean if a field has been set.

### GetBundleVersionOnly

`func (o *VersionsResourceInner) GetBundleVersionOnly() string`

GetBundleVersionOnly returns the BundleVersionOnly field if non-nil, zero value otherwise.

### GetBundleVersionOnlyOk

`func (o *VersionsResourceInner) GetBundleVersionOnlyOk() (*string, bool)`

GetBundleVersionOnlyOk returns a tuple with the BundleVersionOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersionOnly

`func (o *VersionsResourceInner) SetBundleVersionOnly(v string)`

SetBundleVersionOnly sets BundleVersionOnly field to given value.

### HasBundleVersionOnly

`func (o *VersionsResourceInner) HasBundleVersionOnly() bool`

HasBundleVersionOnly returns a boolean if a field has been set.

### GetCapiVersion

`func (o *VersionsResourceInner) GetCapiVersion() string`

GetCapiVersion returns the CapiVersion field if non-nil, zero value otherwise.

### GetCapiVersionOk

`func (o *VersionsResourceInner) GetCapiVersionOk() (*string, bool)`

GetCapiVersionOk returns a tuple with the CapiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapiVersion

`func (o *VersionsResourceInner) SetCapiVersion(v string)`

SetCapiVersion sets CapiVersion field to given value.

### HasCapiVersion

`func (o *VersionsResourceInner) HasCapiVersion() bool`

HasCapiVersion returns a boolean if a field has been set.

### GetCtkVersion

`func (o *VersionsResourceInner) GetCtkVersion() string`

GetCtkVersion returns the CtkVersion field if non-nil, zero value otherwise.

### GetCtkVersionOk

`func (o *VersionsResourceInner) GetCtkVersionOk() (*string, bool)`

GetCtkVersionOk returns a tuple with the CtkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtkVersion

`func (o *VersionsResourceInner) SetCtkVersion(v string)`

SetCtkVersion sets CtkVersion field to given value.

### HasCtkVersion

`func (o *VersionsResourceInner) HasCtkVersion() bool`

HasCtkVersion returns a boolean if a field has been set.

### GetDiskChannelRevision

`func (o *VersionsResourceInner) GetDiskChannelRevision() int32`

GetDiskChannelRevision returns the DiskChannelRevision field if non-nil, zero value otherwise.

### GetDiskChannelRevisionOk

`func (o *VersionsResourceInner) GetDiskChannelRevisionOk() (*int32, bool)`

GetDiskChannelRevisionOk returns a tuple with the DiskChannelRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskChannelRevision

`func (o *VersionsResourceInner) SetDiskChannelRevision(v int32)`

SetDiskChannelRevision sets DiskChannelRevision field to given value.

### HasDiskChannelRevision

`func (o *VersionsResourceInner) HasDiskChannelRevision() bool`

HasDiskChannelRevision returns a boolean if a field has been set.

### GetEcFw

`func (o *VersionsResourceInner) GetEcFw() string`

GetEcFw returns the EcFw field if non-nil, zero value otherwise.

### GetEcFwOk

`func (o *VersionsResourceInner) GetEcFwOk() (*string, bool)`

GetEcFwOk returns a tuple with the EcFw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcFw

`func (o *VersionsResourceInner) SetEcFw(v string)`

SetEcFw sets EcFw field to given value.

### HasEcFw

`func (o *VersionsResourceInner) HasEcFw() bool`

HasEcFw returns a boolean if a field has been set.

### GetFwDefaultPlatformBrand

`func (o *VersionsResourceInner) GetFwDefaultPlatformBrand() string`

GetFwDefaultPlatformBrand returns the FwDefaultPlatformBrand field if non-nil, zero value otherwise.

### GetFwDefaultPlatformBrandOk

`func (o *VersionsResourceInner) GetFwDefaultPlatformBrandOk() (*string, bool)`

GetFwDefaultPlatformBrandOk returns a tuple with the FwDefaultPlatformBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwDefaultPlatformBrand

`func (o *VersionsResourceInner) SetFwDefaultPlatformBrand(v string)`

SetFwDefaultPlatformBrand sets FwDefaultPlatformBrand field to given value.

### HasFwDefaultPlatformBrand

`func (o *VersionsResourceInner) HasFwDefaultPlatformBrand() bool`

HasFwDefaultPlatformBrand returns a boolean if a field has been set.

### GetFwDefaultPlatformBrandNumeric

`func (o *VersionsResourceInner) GetFwDefaultPlatformBrandNumeric() int32`

GetFwDefaultPlatformBrandNumeric returns the FwDefaultPlatformBrandNumeric field if non-nil, zero value otherwise.

### GetFwDefaultPlatformBrandNumericOk

`func (o *VersionsResourceInner) GetFwDefaultPlatformBrandNumericOk() (*int32, bool)`

GetFwDefaultPlatformBrandNumericOk returns a tuple with the FwDefaultPlatformBrandNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwDefaultPlatformBrandNumeric

`func (o *VersionsResourceInner) SetFwDefaultPlatformBrandNumeric(v int32)`

SetFwDefaultPlatformBrandNumeric sets FwDefaultPlatformBrandNumeric field to given value.

### HasFwDefaultPlatformBrandNumeric

`func (o *VersionsResourceInner) HasFwDefaultPlatformBrandNumeric() bool`

HasFwDefaultPlatformBrandNumeric returns a boolean if a field has been set.

### GetGemVersion

`func (o *VersionsResourceInner) GetGemVersion() string`

GetGemVersion returns the GemVersion field if non-nil, zero value otherwise.

### GetGemVersionOk

`func (o *VersionsResourceInner) GetGemVersionOk() (*string, bool)`

GetGemVersionOk returns a tuple with the GemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGemVersion

`func (o *VersionsResourceInner) SetGemVersion(v string)`

SetGemVersion sets GemVersion field to given value.

### HasGemVersion

`func (o *VersionsResourceInner) HasGemVersion() bool`

HasGemVersion returns a boolean if a field has been set.

### GetHimModel

`func (o *VersionsResourceInner) GetHimModel() string`

GetHimModel returns the HimModel field if non-nil, zero value otherwise.

### GetHimModelOk

`func (o *VersionsResourceInner) GetHimModelOk() (*string, bool)`

GetHimModelOk returns a tuple with the HimModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHimModel

`func (o *VersionsResourceInner) SetHimModel(v string)`

SetHimModel sets HimModel field to given value.

### HasHimModel

`func (o *VersionsResourceInner) HasHimModel() bool`

HasHimModel returns a boolean if a field has been set.

### GetHimRev

`func (o *VersionsResourceInner) GetHimRev() string`

GetHimRev returns the HimRev field if non-nil, zero value otherwise.

### GetHimRevOk

`func (o *VersionsResourceInner) GetHimRevOk() (*string, bool)`

GetHimRevOk returns a tuple with the HimRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHimRev

`func (o *VersionsResourceInner) SetHimRev(v string)`

SetHimRev sets HimRev field to given value.

### HasHimRev

`func (o *VersionsResourceInner) HasHimRev() bool`

HasHimRev returns a boolean if a field has been set.

### GetHostChannelRevision

`func (o *VersionsResourceInner) GetHostChannelRevision() int32`

GetHostChannelRevision returns the HostChannelRevision field if non-nil, zero value otherwise.

### GetHostChannelRevisionOk

`func (o *VersionsResourceInner) GetHostChannelRevisionOk() (*int32, bool)`

GetHostChannelRevisionOk returns a tuple with the HostChannelRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostChannelRevision

`func (o *VersionsResourceInner) SetHostChannelRevision(v int32)`

SetHostChannelRevision sets HostChannelRevision field to given value.

### HasHostChannelRevision

`func (o *VersionsResourceInner) HasHostChannelRevision() bool`

HasHostChannelRevision returns a boolean if a field has been set.

### GetHwRev

`func (o *VersionsResourceInner) GetHwRev() string`

GetHwRev returns the HwRev field if non-nil, zero value otherwise.

### GetHwRevOk

`func (o *VersionsResourceInner) GetHwRevOk() (*string, bool)`

GetHwRevOk returns a tuple with the HwRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwRev

`func (o *VersionsResourceInner) SetHwRev(v string)`

SetHwRev sets HwRev field to given value.

### HasHwRev

`func (o *VersionsResourceInner) HasHwRev() bool`

HasHwRev returns a boolean if a field has been set.

### GetMcBaseFw

`func (o *VersionsResourceInner) GetMcBaseFw() string`

GetMcBaseFw returns the McBaseFw field if non-nil, zero value otherwise.

### GetMcBaseFwOk

`func (o *VersionsResourceInner) GetMcBaseFwOk() (*string, bool)`

GetMcBaseFwOk returns a tuple with the McBaseFw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcBaseFw

`func (o *VersionsResourceInner) SetMcBaseFw(v string)`

SetMcBaseFw sets McBaseFw field to given value.

### HasMcBaseFw

`func (o *VersionsResourceInner) HasMcBaseFw() bool`

HasMcBaseFw returns a boolean if a field has been set.

### GetMcFw

`func (o *VersionsResourceInner) GetMcFw() string`

GetMcFw returns the McFw field if non-nil, zero value otherwise.

### GetMcFwOk

`func (o *VersionsResourceInner) GetMcFwOk() (*string, bool)`

GetMcFwOk returns a tuple with the McFw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcFw

`func (o *VersionsResourceInner) SetMcFw(v string)`

SetMcFw sets McFw field to given value.

### HasMcFw

`func (o *VersionsResourceInner) HasMcFw() bool`

HasMcFw returns a boolean if a field has been set.

### GetMcLoader

`func (o *VersionsResourceInner) GetMcLoader() string`

GetMcLoader returns the McLoader field if non-nil, zero value otherwise.

### GetMcLoaderOk

`func (o *VersionsResourceInner) GetMcLoaderOk() (*string, bool)`

GetMcLoaderOk returns a tuple with the McLoader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcLoader

`func (o *VersionsResourceInner) SetMcLoader(v string)`

SetMcLoader sets McLoader field to given value.

### HasMcLoader

`func (o *VersionsResourceInner) HasMcLoader() bool`

HasMcLoader returns a boolean if a field has been set.

### GetMcosVersion

`func (o *VersionsResourceInner) GetMcosVersion() string`

GetMcosVersion returns the McosVersion field if non-nil, zero value otherwise.

### GetMcosVersionOk

`func (o *VersionsResourceInner) GetMcosVersionOk() (*string, bool)`

GetMcosVersionOk returns a tuple with the McosVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcosVersion

`func (o *VersionsResourceInner) SetMcosVersion(v string)`

SetMcosVersion sets McosVersion field to given value.

### HasMcosVersion

`func (o *VersionsResourceInner) HasMcosVersion() bool`

HasMcosVersion returns a boolean if a field has been set.

### GetMrcVersion

`func (o *VersionsResourceInner) GetMrcVersion() string`

GetMrcVersion returns the MrcVersion field if non-nil, zero value otherwise.

### GetMrcVersionOk

`func (o *VersionsResourceInner) GetMrcVersionOk() (*string, bool)`

GetMrcVersionOk returns a tuple with the MrcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrcVersion

`func (o *VersionsResourceInner) SetMrcVersion(v string)`

SetMrcVersion sets MrcVersion field to given value.

### HasMrcVersion

`func (o *VersionsResourceInner) HasMrcVersion() bool`

HasMrcVersion returns a boolean if a field has been set.

### GetPcieSwitchBackendConfigurationVersion

`func (o *VersionsResourceInner) GetPcieSwitchBackendConfigurationVersion() string`

GetPcieSwitchBackendConfigurationVersion returns the PcieSwitchBackendConfigurationVersion field if non-nil, zero value otherwise.

### GetPcieSwitchBackendConfigurationVersionOk

`func (o *VersionsResourceInner) GetPcieSwitchBackendConfigurationVersionOk() (*string, bool)`

GetPcieSwitchBackendConfigurationVersionOk returns a tuple with the PcieSwitchBackendConfigurationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSwitchBackendConfigurationVersion

`func (o *VersionsResourceInner) SetPcieSwitchBackendConfigurationVersion(v string)`

SetPcieSwitchBackendConfigurationVersion sets PcieSwitchBackendConfigurationVersion field to given value.

### HasPcieSwitchBackendConfigurationVersion

`func (o *VersionsResourceInner) HasPcieSwitchBackendConfigurationVersion() bool`

HasPcieSwitchBackendConfigurationVersion returns a boolean if a field has been set.

### GetPcieSwitchBackendFirmwareVersion

`func (o *VersionsResourceInner) GetPcieSwitchBackendFirmwareVersion() string`

GetPcieSwitchBackendFirmwareVersion returns the PcieSwitchBackendFirmwareVersion field if non-nil, zero value otherwise.

### GetPcieSwitchBackendFirmwareVersionOk

`func (o *VersionsResourceInner) GetPcieSwitchBackendFirmwareVersionOk() (*string, bool)`

GetPcieSwitchBackendFirmwareVersionOk returns a tuple with the PcieSwitchBackendFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSwitchBackendFirmwareVersion

`func (o *VersionsResourceInner) SetPcieSwitchBackendFirmwareVersion(v string)`

SetPcieSwitchBackendFirmwareVersion sets PcieSwitchBackendFirmwareVersion field to given value.

### HasPcieSwitchBackendFirmwareVersion

`func (o *VersionsResourceInner) HasPcieSwitchBackendFirmwareVersion() bool`

HasPcieSwitchBackendFirmwareVersion returns a boolean if a field has been set.

### GetPcieSwitchFrontendConfigurationVersion

`func (o *VersionsResourceInner) GetPcieSwitchFrontendConfigurationVersion() string`

GetPcieSwitchFrontendConfigurationVersion returns the PcieSwitchFrontendConfigurationVersion field if non-nil, zero value otherwise.

### GetPcieSwitchFrontendConfigurationVersionOk

`func (o *VersionsResourceInner) GetPcieSwitchFrontendConfigurationVersionOk() (*string, bool)`

GetPcieSwitchFrontendConfigurationVersionOk returns a tuple with the PcieSwitchFrontendConfigurationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSwitchFrontendConfigurationVersion

`func (o *VersionsResourceInner) SetPcieSwitchFrontendConfigurationVersion(v string)`

SetPcieSwitchFrontendConfigurationVersion sets PcieSwitchFrontendConfigurationVersion field to given value.

### HasPcieSwitchFrontendConfigurationVersion

`func (o *VersionsResourceInner) HasPcieSwitchFrontendConfigurationVersion() bool`

HasPcieSwitchFrontendConfigurationVersion returns a boolean if a field has been set.

### GetPcieSwitchFrontendFirmwareVersion

`func (o *VersionsResourceInner) GetPcieSwitchFrontendFirmwareVersion() string`

GetPcieSwitchFrontendFirmwareVersion returns the PcieSwitchFrontendFirmwareVersion field if non-nil, zero value otherwise.

### GetPcieSwitchFrontendFirmwareVersionOk

`func (o *VersionsResourceInner) GetPcieSwitchFrontendFirmwareVersionOk() (*string, bool)`

GetPcieSwitchFrontendFirmwareVersionOk returns a tuple with the PcieSwitchFrontendFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieSwitchFrontendFirmwareVersion

`func (o *VersionsResourceInner) SetPcieSwitchFrontendFirmwareVersion(v string)`

SetPcieSwitchFrontendFirmwareVersion sets PcieSwitchFrontendFirmwareVersion field to given value.

### HasPcieSwitchFrontendFirmwareVersion

`func (o *VersionsResourceInner) HasPcieSwitchFrontendFirmwareVersion() bool`

HasPcieSwitchFrontendFirmwareVersion returns a boolean if a field has been set.

### GetPldRev

`func (o *VersionsResourceInner) GetPldRev() string`

GetPldRev returns the PldRev field if non-nil, zero value otherwise.

### GetPldRevOk

`func (o *VersionsResourceInner) GetPldRevOk() (*string, bool)`

GetPldRevOk returns a tuple with the PldRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPldRev

`func (o *VersionsResourceInner) SetPldRev(v string)`

SetPldRev sets PldRev field to given value.

### HasPldRev

`func (o *VersionsResourceInner) HasPldRev() bool`

HasPldRev returns a boolean if a field has been set.

### GetPmCpldVersion

`func (o *VersionsResourceInner) GetPmCpldVersion() string`

GetPmCpldVersion returns the PmCpldVersion field if non-nil, zero value otherwise.

### GetPmCpldVersionOk

`func (o *VersionsResourceInner) GetPmCpldVersionOk() (*string, bool)`

GetPmCpldVersionOk returns a tuple with the PmCpldVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmCpldVersion

`func (o *VersionsResourceInner) SetPmCpldVersion(v string)`

SetPmCpldVersion sets PmCpldVersion field to given value.

### HasPmCpldVersion

`func (o *VersionsResourceInner) HasPmCpldVersion() bool`

HasPmCpldVersion returns a boolean if a field has been set.

### GetPrmVersion

`func (o *VersionsResourceInner) GetPrmVersion() string`

GetPrmVersion returns the PrmVersion field if non-nil, zero value otherwise.

### GetPrmVersionOk

`func (o *VersionsResourceInner) GetPrmVersionOk() (*string, bool)`

GetPrmVersionOk returns a tuple with the PrmVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrmVersion

`func (o *VersionsResourceInner) SetPrmVersion(v string)`

SetPrmVersion sets PrmVersion field to given value.

### HasPrmVersion

`func (o *VersionsResourceInner) HasPrmVersion() bool`

HasPrmVersion returns a boolean if a field has been set.

### GetPubsVersion

`func (o *VersionsResourceInner) GetPubsVersion() string`

GetPubsVersion returns the PubsVersion field if non-nil, zero value otherwise.

### GetPubsVersionOk

`func (o *VersionsResourceInner) GetPubsVersionOk() (*string, bool)`

GetPubsVersionOk returns a tuple with the PubsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubsVersion

`func (o *VersionsResourceInner) SetPubsVersion(v string)`

SetPubsVersion sets PubsVersion field to given value.

### HasPubsVersion

`func (o *VersionsResourceInner) HasPubsVersion() bool`

HasPubsVersion returns a boolean if a field has been set.

### GetScBaselevel

`func (o *VersionsResourceInner) GetScBaselevel() string`

GetScBaselevel returns the ScBaselevel field if non-nil, zero value otherwise.

### GetScBaselevelOk

`func (o *VersionsResourceInner) GetScBaselevelOk() (*string, bool)`

GetScBaselevelOk returns a tuple with the ScBaselevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScBaselevel

`func (o *VersionsResourceInner) SetScBaselevel(v string)`

SetScBaselevel sets ScBaselevel field to given value.

### HasScBaselevel

`func (o *VersionsResourceInner) HasScBaselevel() bool`

HasScBaselevel returns a boolean if a field has been set.

### GetScCpuType

`func (o *VersionsResourceInner) GetScCpuType() string`

GetScCpuType returns the ScCpuType field if non-nil, zero value otherwise.

### GetScCpuTypeOk

`func (o *VersionsResourceInner) GetScCpuTypeOk() (*string, bool)`

GetScCpuTypeOk returns a tuple with the ScCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScCpuType

`func (o *VersionsResourceInner) SetScCpuType(v string)`

SetScCpuType sets ScCpuType field to given value.

### HasScCpuType

`func (o *VersionsResourceInner) HasScCpuType() bool`

HasScCpuType returns a boolean if a field has been set.

### GetScFuVersion

`func (o *VersionsResourceInner) GetScFuVersion() string`

GetScFuVersion returns the ScFuVersion field if non-nil, zero value otherwise.

### GetScFuVersionOk

`func (o *VersionsResourceInner) GetScFuVersionOk() (*string, bool)`

GetScFuVersionOk returns a tuple with the ScFuVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScFuVersion

`func (o *VersionsResourceInner) SetScFuVersion(v string)`

SetScFuVersion sets ScFuVersion field to given value.

### HasScFuVersion

`func (o *VersionsResourceInner) HasScFuVersion() bool`

HasScFuVersion returns a boolean if a field has been set.

### GetScFw

`func (o *VersionsResourceInner) GetScFw() string`

GetScFw returns the ScFw field if non-nil, zero value otherwise.

### GetScFwOk

`func (o *VersionsResourceInner) GetScFwOk() (*string, bool)`

GetScFwOk returns a tuple with the ScFw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScFw

`func (o *VersionsResourceInner) SetScFw(v string)`

SetScFw sets ScFw field to given value.

### HasScFw

`func (o *VersionsResourceInner) HasScFw() bool`

HasScFw returns a boolean if a field has been set.

### GetScLoader

`func (o *VersionsResourceInner) GetScLoader() string`

GetScLoader returns the ScLoader field if non-nil, zero value otherwise.

### GetScLoaderOk

`func (o *VersionsResourceInner) GetScLoaderOk() (*string, bool)`

GetScLoaderOk returns a tuple with the ScLoader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScLoader

`func (o *VersionsResourceInner) SetScLoader(v string)`

SetScLoader sets ScLoader field to given value.

### HasScLoader

`func (o *VersionsResourceInner) HasScLoader() bool`

HasScLoader returns a boolean if a field has been set.

### GetScMemory

`func (o *VersionsResourceInner) GetScMemory() string`

GetScMemory returns the ScMemory field if non-nil, zero value otherwise.

### GetScMemoryOk

`func (o *VersionsResourceInner) GetScMemoryOk() (*string, bool)`

GetScMemoryOk returns a tuple with the ScMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScMemory

`func (o *VersionsResourceInner) SetScMemory(v string)`

SetScMemory sets ScMemory field to given value.

### HasScMemory

`func (o *VersionsResourceInner) HasScMemory() bool`

HasScMemory returns a boolean if a field has been set.

### GetTranslationVersion

`func (o *VersionsResourceInner) GetTranslationVersion() string`

GetTranslationVersion returns the TranslationVersion field if non-nil, zero value otherwise.

### GetTranslationVersionOk

`func (o *VersionsResourceInner) GetTranslationVersionOk() (*string, bool)`

GetTranslationVersionOk returns a tuple with the TranslationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationVersion

`func (o *VersionsResourceInner) SetTranslationVersion(v string)`

SetTranslationVersion sets TranslationVersion field to given value.

### HasTranslationVersion

`func (o *VersionsResourceInner) HasTranslationVersion() bool`

HasTranslationVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


