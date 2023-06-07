# IscsiPortResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**DefaultRouter** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**IpVersion** | Pointer to **string** |  | [optional] 
**LinkLocalAddress** | Pointer to **string** | IPv6 Link-Local Address | [optional] 
**MacAddress** | Pointer to **string** | Primary MAC address for the port | [optional] 
**Netmask** | Pointer to **string** | Netmask of the iSCSI IP address for this port | [optional] 
**Sfp10GCompliance** | Pointer to **string** |  | [optional] 
**Sfp10GComplianceNumeric** | Pointer to **int64** |  | [optional] 
**SfpCableLength** | Pointer to **int64** |  | [optional] 
**SfpCableTechnology** | Pointer to **string** |  | [optional] 
**SfpCableTechnologyNumeric** | Pointer to **int64** |  | [optional] 
**SfpEthernetCompliance** | Pointer to **string** |  | [optional] 
**SfpEthernetComplianceNumeric** | Pointer to **int64** |  | [optional] 
**SfpPartNumber** | Pointer to **string** |  | [optional] 
**SfpPresent** | Pointer to **string** |  | [optional] 
**SfpPresentNumeric** | Pointer to **int64** |  | [optional] 
**SfpRevision** | Pointer to **string** |  | [optional] 
**SfpStatus** | Pointer to **string** |  | [optional] 
**SfpStatusNumeric** | Pointer to **int64** |  | [optional] 
**SfpVendor** | Pointer to **string** |  | [optional] 

## Methods

### NewIscsiPortResourceInner

`func NewIscsiPortResourceInner() *IscsiPortResourceInner`

NewIscsiPortResourceInner instantiates a new IscsiPortResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIscsiPortResourceInnerWithDefaults

`func NewIscsiPortResourceInnerWithDefaults() *IscsiPortResourceInner`

NewIscsiPortResourceInnerWithDefaults instantiates a new IscsiPortResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *IscsiPortResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *IscsiPortResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *IscsiPortResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *IscsiPortResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *IscsiPortResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *IscsiPortResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *IscsiPortResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *IscsiPortResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetDefaultRouter

`func (o *IscsiPortResourceInner) GetDefaultRouter() string`

GetDefaultRouter returns the DefaultRouter field if non-nil, zero value otherwise.

### GetDefaultRouterOk

`func (o *IscsiPortResourceInner) GetDefaultRouterOk() (*string, bool)`

GetDefaultRouterOk returns a tuple with the DefaultRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRouter

`func (o *IscsiPortResourceInner) SetDefaultRouter(v string)`

SetDefaultRouter sets DefaultRouter field to given value.

### HasDefaultRouter

`func (o *IscsiPortResourceInner) HasDefaultRouter() bool`

HasDefaultRouter returns a boolean if a field has been set.

### GetGateway

`func (o *IscsiPortResourceInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IscsiPortResourceInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IscsiPortResourceInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *IscsiPortResourceInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpAddress

`func (o *IscsiPortResourceInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *IscsiPortResourceInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *IscsiPortResourceInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *IscsiPortResourceInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpVersion

`func (o *IscsiPortResourceInner) GetIpVersion() string`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *IscsiPortResourceInner) GetIpVersionOk() (*string, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *IscsiPortResourceInner) SetIpVersion(v string)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *IscsiPortResourceInner) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetLinkLocalAddress

`func (o *IscsiPortResourceInner) GetLinkLocalAddress() string`

GetLinkLocalAddress returns the LinkLocalAddress field if non-nil, zero value otherwise.

### GetLinkLocalAddressOk

`func (o *IscsiPortResourceInner) GetLinkLocalAddressOk() (*string, bool)`

GetLinkLocalAddressOk returns a tuple with the LinkLocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkLocalAddress

`func (o *IscsiPortResourceInner) SetLinkLocalAddress(v string)`

SetLinkLocalAddress sets LinkLocalAddress field to given value.

### HasLinkLocalAddress

`func (o *IscsiPortResourceInner) HasLinkLocalAddress() bool`

HasLinkLocalAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *IscsiPortResourceInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *IscsiPortResourceInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *IscsiPortResourceInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *IscsiPortResourceInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetNetmask

`func (o *IscsiPortResourceInner) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *IscsiPortResourceInner) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *IscsiPortResourceInner) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *IscsiPortResourceInner) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetSfp10GCompliance

`func (o *IscsiPortResourceInner) GetSfp10GCompliance() string`

GetSfp10GCompliance returns the Sfp10GCompliance field if non-nil, zero value otherwise.

### GetSfp10GComplianceOk

`func (o *IscsiPortResourceInner) GetSfp10GComplianceOk() (*string, bool)`

GetSfp10GComplianceOk returns a tuple with the Sfp10GCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfp10GCompliance

`func (o *IscsiPortResourceInner) SetSfp10GCompliance(v string)`

SetSfp10GCompliance sets Sfp10GCompliance field to given value.

### HasSfp10GCompliance

`func (o *IscsiPortResourceInner) HasSfp10GCompliance() bool`

HasSfp10GCompliance returns a boolean if a field has been set.

### GetSfp10GComplianceNumeric

`func (o *IscsiPortResourceInner) GetSfp10GComplianceNumeric() int64`

GetSfp10GComplianceNumeric returns the Sfp10GComplianceNumeric field if non-nil, zero value otherwise.

### GetSfp10GComplianceNumericOk

`func (o *IscsiPortResourceInner) GetSfp10GComplianceNumericOk() (*int64, bool)`

GetSfp10GComplianceNumericOk returns a tuple with the Sfp10GComplianceNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfp10GComplianceNumeric

`func (o *IscsiPortResourceInner) SetSfp10GComplianceNumeric(v int64)`

SetSfp10GComplianceNumeric sets Sfp10GComplianceNumeric field to given value.

### HasSfp10GComplianceNumeric

`func (o *IscsiPortResourceInner) HasSfp10GComplianceNumeric() bool`

HasSfp10GComplianceNumeric returns a boolean if a field has been set.

### GetSfpCableLength

`func (o *IscsiPortResourceInner) GetSfpCableLength() int64`

GetSfpCableLength returns the SfpCableLength field if non-nil, zero value otherwise.

### GetSfpCableLengthOk

`func (o *IscsiPortResourceInner) GetSfpCableLengthOk() (*int64, bool)`

GetSfpCableLengthOk returns a tuple with the SfpCableLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpCableLength

`func (o *IscsiPortResourceInner) SetSfpCableLength(v int64)`

SetSfpCableLength sets SfpCableLength field to given value.

### HasSfpCableLength

`func (o *IscsiPortResourceInner) HasSfpCableLength() bool`

HasSfpCableLength returns a boolean if a field has been set.

### GetSfpCableTechnology

`func (o *IscsiPortResourceInner) GetSfpCableTechnology() string`

GetSfpCableTechnology returns the SfpCableTechnology field if non-nil, zero value otherwise.

### GetSfpCableTechnologyOk

`func (o *IscsiPortResourceInner) GetSfpCableTechnologyOk() (*string, bool)`

GetSfpCableTechnologyOk returns a tuple with the SfpCableTechnology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpCableTechnology

`func (o *IscsiPortResourceInner) SetSfpCableTechnology(v string)`

SetSfpCableTechnology sets SfpCableTechnology field to given value.

### HasSfpCableTechnology

`func (o *IscsiPortResourceInner) HasSfpCableTechnology() bool`

HasSfpCableTechnology returns a boolean if a field has been set.

### GetSfpCableTechnologyNumeric

`func (o *IscsiPortResourceInner) GetSfpCableTechnologyNumeric() int64`

GetSfpCableTechnologyNumeric returns the SfpCableTechnologyNumeric field if non-nil, zero value otherwise.

### GetSfpCableTechnologyNumericOk

`func (o *IscsiPortResourceInner) GetSfpCableTechnologyNumericOk() (*int64, bool)`

GetSfpCableTechnologyNumericOk returns a tuple with the SfpCableTechnologyNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpCableTechnologyNumeric

`func (o *IscsiPortResourceInner) SetSfpCableTechnologyNumeric(v int64)`

SetSfpCableTechnologyNumeric sets SfpCableTechnologyNumeric field to given value.

### HasSfpCableTechnologyNumeric

`func (o *IscsiPortResourceInner) HasSfpCableTechnologyNumeric() bool`

HasSfpCableTechnologyNumeric returns a boolean if a field has been set.

### GetSfpEthernetCompliance

`func (o *IscsiPortResourceInner) GetSfpEthernetCompliance() string`

GetSfpEthernetCompliance returns the SfpEthernetCompliance field if non-nil, zero value otherwise.

### GetSfpEthernetComplianceOk

`func (o *IscsiPortResourceInner) GetSfpEthernetComplianceOk() (*string, bool)`

GetSfpEthernetComplianceOk returns a tuple with the SfpEthernetCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpEthernetCompliance

`func (o *IscsiPortResourceInner) SetSfpEthernetCompliance(v string)`

SetSfpEthernetCompliance sets SfpEthernetCompliance field to given value.

### HasSfpEthernetCompliance

`func (o *IscsiPortResourceInner) HasSfpEthernetCompliance() bool`

HasSfpEthernetCompliance returns a boolean if a field has been set.

### GetSfpEthernetComplianceNumeric

`func (o *IscsiPortResourceInner) GetSfpEthernetComplianceNumeric() int64`

GetSfpEthernetComplianceNumeric returns the SfpEthernetComplianceNumeric field if non-nil, zero value otherwise.

### GetSfpEthernetComplianceNumericOk

`func (o *IscsiPortResourceInner) GetSfpEthernetComplianceNumericOk() (*int64, bool)`

GetSfpEthernetComplianceNumericOk returns a tuple with the SfpEthernetComplianceNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpEthernetComplianceNumeric

`func (o *IscsiPortResourceInner) SetSfpEthernetComplianceNumeric(v int64)`

SetSfpEthernetComplianceNumeric sets SfpEthernetComplianceNumeric field to given value.

### HasSfpEthernetComplianceNumeric

`func (o *IscsiPortResourceInner) HasSfpEthernetComplianceNumeric() bool`

HasSfpEthernetComplianceNumeric returns a boolean if a field has been set.

### GetSfpPartNumber

`func (o *IscsiPortResourceInner) GetSfpPartNumber() string`

GetSfpPartNumber returns the SfpPartNumber field if non-nil, zero value otherwise.

### GetSfpPartNumberOk

`func (o *IscsiPortResourceInner) GetSfpPartNumberOk() (*string, bool)`

GetSfpPartNumberOk returns a tuple with the SfpPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpPartNumber

`func (o *IscsiPortResourceInner) SetSfpPartNumber(v string)`

SetSfpPartNumber sets SfpPartNumber field to given value.

### HasSfpPartNumber

`func (o *IscsiPortResourceInner) HasSfpPartNumber() bool`

HasSfpPartNumber returns a boolean if a field has been set.

### GetSfpPresent

`func (o *IscsiPortResourceInner) GetSfpPresent() string`

GetSfpPresent returns the SfpPresent field if non-nil, zero value otherwise.

### GetSfpPresentOk

`func (o *IscsiPortResourceInner) GetSfpPresentOk() (*string, bool)`

GetSfpPresentOk returns a tuple with the SfpPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpPresent

`func (o *IscsiPortResourceInner) SetSfpPresent(v string)`

SetSfpPresent sets SfpPresent field to given value.

### HasSfpPresent

`func (o *IscsiPortResourceInner) HasSfpPresent() bool`

HasSfpPresent returns a boolean if a field has been set.

### GetSfpPresentNumeric

`func (o *IscsiPortResourceInner) GetSfpPresentNumeric() int64`

GetSfpPresentNumeric returns the SfpPresentNumeric field if non-nil, zero value otherwise.

### GetSfpPresentNumericOk

`func (o *IscsiPortResourceInner) GetSfpPresentNumericOk() (*int64, bool)`

GetSfpPresentNumericOk returns a tuple with the SfpPresentNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpPresentNumeric

`func (o *IscsiPortResourceInner) SetSfpPresentNumeric(v int64)`

SetSfpPresentNumeric sets SfpPresentNumeric field to given value.

### HasSfpPresentNumeric

`func (o *IscsiPortResourceInner) HasSfpPresentNumeric() bool`

HasSfpPresentNumeric returns a boolean if a field has been set.

### GetSfpRevision

`func (o *IscsiPortResourceInner) GetSfpRevision() string`

GetSfpRevision returns the SfpRevision field if non-nil, zero value otherwise.

### GetSfpRevisionOk

`func (o *IscsiPortResourceInner) GetSfpRevisionOk() (*string, bool)`

GetSfpRevisionOk returns a tuple with the SfpRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpRevision

`func (o *IscsiPortResourceInner) SetSfpRevision(v string)`

SetSfpRevision sets SfpRevision field to given value.

### HasSfpRevision

`func (o *IscsiPortResourceInner) HasSfpRevision() bool`

HasSfpRevision returns a boolean if a field has been set.

### GetSfpStatus

`func (o *IscsiPortResourceInner) GetSfpStatus() string`

GetSfpStatus returns the SfpStatus field if non-nil, zero value otherwise.

### GetSfpStatusOk

`func (o *IscsiPortResourceInner) GetSfpStatusOk() (*string, bool)`

GetSfpStatusOk returns a tuple with the SfpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpStatus

`func (o *IscsiPortResourceInner) SetSfpStatus(v string)`

SetSfpStatus sets SfpStatus field to given value.

### HasSfpStatus

`func (o *IscsiPortResourceInner) HasSfpStatus() bool`

HasSfpStatus returns a boolean if a field has been set.

### GetSfpStatusNumeric

`func (o *IscsiPortResourceInner) GetSfpStatusNumeric() int64`

GetSfpStatusNumeric returns the SfpStatusNumeric field if non-nil, zero value otherwise.

### GetSfpStatusNumericOk

`func (o *IscsiPortResourceInner) GetSfpStatusNumericOk() (*int64, bool)`

GetSfpStatusNumericOk returns a tuple with the SfpStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpStatusNumeric

`func (o *IscsiPortResourceInner) SetSfpStatusNumeric(v int64)`

SetSfpStatusNumeric sets SfpStatusNumeric field to given value.

### HasSfpStatusNumeric

`func (o *IscsiPortResourceInner) HasSfpStatusNumeric() bool`

HasSfpStatusNumeric returns a boolean if a field has been set.

### GetSfpVendor

`func (o *IscsiPortResourceInner) GetSfpVendor() string`

GetSfpVendor returns the SfpVendor field if non-nil, zero value otherwise.

### GetSfpVendorOk

`func (o *IscsiPortResourceInner) GetSfpVendorOk() (*string, bool)`

GetSfpVendorOk returns a tuple with the SfpVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfpVendor

`func (o *IscsiPortResourceInner) SetSfpVendor(v string)`

SetSfpVendor sets SfpVendor field to given value.

### HasSfpVendor

`func (o *IscsiPortResourceInner) HasSfpVendor() bool`

HasSfpVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


