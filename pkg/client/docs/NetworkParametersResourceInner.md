# NetworkParametersResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
<<<<<<< HEAD
**ActiveVersion** | Pointer to **int64** | Currently assigned IP version | [optional] 
**AddressingMode** | Pointer to **string** | The way in which an IP address should be obtained | [optional] 
**AddressingModeNumeric** | Pointer to **int64** | The way in which an IP address should be obtained( In numeric form ) | [optional] 
**AutoNegotiation** | Pointer to **string** |  | [optional] 
**AutoNegotiationNumeric** | Pointer to **int64** |  | [optional] 
**DuplexMode** | Pointer to **string** | Type of duplex mode that is specified | [optional] 
**DuplexModeNumeric** | Pointer to **int64** | Type of duplex mode that is specified( In numeric form ) | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**HealthNumeric** | Pointer to **int64** |  | [optional] 
=======
**ActiveVersion** | Pointer to **int32** | Currently assigned IP version | [optional] 
**AddressingMode** | Pointer to **string** | The way in which an IP address should be obtained | [optional] 
**AddressingModeNumeric** | Pointer to **int32** | The way in which an IP address should be obtained( In numeric form ) | [optional] 
**AutoNegotiation** | Pointer to **string** |  | [optional] 
**AutoNegotiationNumeric** | Pointer to **int32** |  | [optional] 
**DuplexMode** | Pointer to **string** | Type of duplex mode that is specified | [optional] 
**DuplexModeNumeric** | Pointer to **int32** | Type of duplex mode that is specified( In numeric form ) | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**HealthNumeric** | Pointer to **int32** |  | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
**HealthReason** | Pointer to **string** |  | [optional] 
**HealthRecommendation** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**LinkSpeed** | Pointer to **string** | Link speed associated with this ethernet port | [optional] 
<<<<<<< HEAD
**LinkSpeedNumeric** | Pointer to **int64** | Link speed associated with this ethernet port( In numeric form ) | [optional] 
**MacAddress** | Pointer to **string** | MAC address for the network port | [optional] 
**PingBroadcast** | Pointer to **string** |  | [optional] 
**PingBroadcastNumeric** | Pointer to **int64** |  | [optional] 
=======
**LinkSpeedNumeric** | Pointer to **int32** | Link speed associated with this ethernet port( In numeric form ) | [optional] 
**MacAddress** | Pointer to **string** | MAC address for the network port | [optional] 
**PingBroadcast** | Pointer to **string** |  | [optional] 
**PingBroadcastNumeric** | Pointer to **int32** |  | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
**SubnetMask** | Pointer to **string** | Network port subnet mask | [optional] 

## Methods

### NewNetworkParametersResourceInner

`func NewNetworkParametersResourceInner() *NetworkParametersResourceInner`

NewNetworkParametersResourceInner instantiates a new NetworkParametersResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkParametersResourceInnerWithDefaults

`func NewNetworkParametersResourceInnerWithDefaults() *NetworkParametersResourceInner`

NewNetworkParametersResourceInnerWithDefaults instantiates a new NetworkParametersResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *NetworkParametersResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *NetworkParametersResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *NetworkParametersResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *NetworkParametersResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *NetworkParametersResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NetworkParametersResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NetworkParametersResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NetworkParametersResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetActiveVersion

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) GetActiveVersion() int64`
=======
`func (o *NetworkParametersResourceInner) GetActiveVersion() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetActiveVersion returns the ActiveVersion field if non-nil, zero value otherwise.

### GetActiveVersionOk

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) GetActiveVersionOk() (*int64, bool)`
=======
`func (o *NetworkParametersResourceInner) GetActiveVersionOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetActiveVersionOk returns a tuple with the ActiveVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveVersion

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) SetActiveVersion(v int64)`
=======
`func (o *NetworkParametersResourceInner) SetActiveVersion(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetActiveVersion sets ActiveVersion field to given value.

### HasActiveVersion

`func (o *NetworkParametersResourceInner) HasActiveVersion() bool`

HasActiveVersion returns a boolean if a field has been set.

### GetAddressingMode

`func (o *NetworkParametersResourceInner) GetAddressingMode() string`

GetAddressingMode returns the AddressingMode field if non-nil, zero value otherwise.

### GetAddressingModeOk

`func (o *NetworkParametersResourceInner) GetAddressingModeOk() (*string, bool)`

GetAddressingModeOk returns a tuple with the AddressingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressingMode

`func (o *NetworkParametersResourceInner) SetAddressingMode(v string)`

SetAddressingMode sets AddressingMode field to given value.

### HasAddressingMode

`func (o *NetworkParametersResourceInner) HasAddressingMode() bool`

HasAddressingMode returns a boolean if a field has been set.

### GetAddressingModeNumeric

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) GetAddressingModeNumeric() int64`
=======
`func (o *NetworkParametersResourceInner) GetAddressingModeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAddressingModeNumeric returns the AddressingModeNumeric field if non-nil, zero value otherwise.

### GetAddressingModeNumericOk

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) GetAddressingModeNumericOk() (*int64, bool)`
=======
`func (o *NetworkParametersResourceInner) GetAddressingModeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAddressingModeNumericOk returns a tuple with the AddressingModeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressingModeNumeric

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) SetAddressingModeNumeric(v int64)`
=======
`func (o *NetworkParametersResourceInner) SetAddressingModeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetAddressingModeNumeric sets AddressingModeNumeric field to given value.

### HasAddressingModeNumeric

`func (o *NetworkParametersResourceInner) HasAddressingModeNumeric() bool`

HasAddressingModeNumeric returns a boolean if a field has been set.

### GetAutoNegotiation

`func (o *NetworkParametersResourceInner) GetAutoNegotiation() string`

GetAutoNegotiation returns the AutoNegotiation field if non-nil, zero value otherwise.

### GetAutoNegotiationOk

`func (o *NetworkParametersResourceInner) GetAutoNegotiationOk() (*string, bool)`

GetAutoNegotiationOk returns a tuple with the AutoNegotiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoNegotiation

`func (o *NetworkParametersResourceInner) SetAutoNegotiation(v string)`

SetAutoNegotiation sets AutoNegotiation field to given value.

### HasAutoNegotiation

`func (o *NetworkParametersResourceInner) HasAutoNegotiation() bool`

HasAutoNegotiation returns a boolean if a field has been set.

### GetAutoNegotiationNumeric

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) GetAutoNegotiationNumeric() int64`
=======
`func (o *NetworkParametersResourceInner) GetAutoNegotiationNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAutoNegotiationNumeric returns the AutoNegotiationNumeric field if non-nil, zero value otherwise.

### GetAutoNegotiationNumericOk

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) GetAutoNegotiationNumericOk() (*int64, bool)`
=======
`func (o *NetworkParametersResourceInner) GetAutoNegotiationNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAutoNegotiationNumericOk returns a tuple with the AutoNegotiationNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoNegotiationNumeric

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) SetAutoNegotiationNumeric(v int64)`
=======
`func (o *NetworkParametersResourceInner) SetAutoNegotiationNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetAutoNegotiationNumeric sets AutoNegotiationNumeric field to given value.

### HasAutoNegotiationNumeric

`func (o *NetworkParametersResourceInner) HasAutoNegotiationNumeric() bool`

HasAutoNegotiationNumeric returns a boolean if a field has been set.

### GetDuplexMode

`func (o *NetworkParametersResourceInner) GetDuplexMode() string`

GetDuplexMode returns the DuplexMode field if non-nil, zero value otherwise.

### GetDuplexModeOk

`func (o *NetworkParametersResourceInner) GetDuplexModeOk() (*string, bool)`

GetDuplexModeOk returns a tuple with the DuplexMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplexMode

`func (o *NetworkParametersResourceInner) SetDuplexMode(v string)`

SetDuplexMode sets DuplexMode field to given value.

### HasDuplexMode

`func (o *NetworkParametersResourceInner) HasDuplexMode() bool`

HasDuplexMode returns a boolean if a field has been set.

### GetDuplexModeNumeric

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) GetDuplexModeNumeric() int64`
=======
`func (o *NetworkParametersResourceInner) GetDuplexModeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDuplexModeNumeric returns the DuplexModeNumeric field if non-nil, zero value otherwise.

### GetDuplexModeNumericOk

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) GetDuplexModeNumericOk() (*int64, bool)`
=======
`func (o *NetworkParametersResourceInner) GetDuplexModeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDuplexModeNumericOk returns a tuple with the DuplexModeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplexModeNumeric

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) SetDuplexModeNumeric(v int64)`
=======
`func (o *NetworkParametersResourceInner) SetDuplexModeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetDuplexModeNumeric sets DuplexModeNumeric field to given value.

### HasDuplexModeNumeric

`func (o *NetworkParametersResourceInner) HasDuplexModeNumeric() bool`

HasDuplexModeNumeric returns a boolean if a field has been set.

### GetDurableId

`func (o *NetworkParametersResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *NetworkParametersResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *NetworkParametersResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *NetworkParametersResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetGateway

`func (o *NetworkParametersResourceInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *NetworkParametersResourceInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *NetworkParametersResourceInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *NetworkParametersResourceInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetHealth

`func (o *NetworkParametersResourceInner) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *NetworkParametersResourceInner) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *NetworkParametersResourceInner) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *NetworkParametersResourceInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHealthNumeric

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) GetHealthNumeric() int64`
=======
`func (o *NetworkParametersResourceInner) GetHealthNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHealthNumeric returns the HealthNumeric field if non-nil, zero value otherwise.

### GetHealthNumericOk

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) GetHealthNumericOk() (*int64, bool)`
=======
`func (o *NetworkParametersResourceInner) GetHealthNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHealthNumericOk returns a tuple with the HealthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthNumeric

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) SetHealthNumeric(v int64)`
=======
`func (o *NetworkParametersResourceInner) SetHealthNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetHealthNumeric sets HealthNumeric field to given value.

### HasHealthNumeric

`func (o *NetworkParametersResourceInner) HasHealthNumeric() bool`

HasHealthNumeric returns a boolean if a field has been set.

### GetHealthReason

`func (o *NetworkParametersResourceInner) GetHealthReason() string`

GetHealthReason returns the HealthReason field if non-nil, zero value otherwise.

### GetHealthReasonOk

`func (o *NetworkParametersResourceInner) GetHealthReasonOk() (*string, bool)`

GetHealthReasonOk returns a tuple with the HealthReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReason

`func (o *NetworkParametersResourceInner) SetHealthReason(v string)`

SetHealthReason sets HealthReason field to given value.

### HasHealthReason

`func (o *NetworkParametersResourceInner) HasHealthReason() bool`

HasHealthReason returns a boolean if a field has been set.

### GetHealthRecommendation

`func (o *NetworkParametersResourceInner) GetHealthRecommendation() string`

GetHealthRecommendation returns the HealthRecommendation field if non-nil, zero value otherwise.

### GetHealthRecommendationOk

`func (o *NetworkParametersResourceInner) GetHealthRecommendationOk() (*string, bool)`

GetHealthRecommendationOk returns a tuple with the HealthRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecommendation

`func (o *NetworkParametersResourceInner) SetHealthRecommendation(v string)`

SetHealthRecommendation sets HealthRecommendation field to given value.

### HasHealthRecommendation

`func (o *NetworkParametersResourceInner) HasHealthRecommendation() bool`

HasHealthRecommendation returns a boolean if a field has been set.

### GetIpAddress

`func (o *NetworkParametersResourceInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkParametersResourceInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkParametersResourceInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NetworkParametersResourceInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *NetworkParametersResourceInner) GetLinkSpeed() string`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *NetworkParametersResourceInner) GetLinkSpeedOk() (*string, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *NetworkParametersResourceInner) SetLinkSpeed(v string)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *NetworkParametersResourceInner) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetLinkSpeedNumeric

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) GetLinkSpeedNumeric() int64`
=======
`func (o *NetworkParametersResourceInner) GetLinkSpeedNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLinkSpeedNumeric returns the LinkSpeedNumeric field if non-nil, zero value otherwise.

### GetLinkSpeedNumericOk

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) GetLinkSpeedNumericOk() (*int64, bool)`
=======
`func (o *NetworkParametersResourceInner) GetLinkSpeedNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLinkSpeedNumericOk returns a tuple with the LinkSpeedNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeedNumeric

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) SetLinkSpeedNumeric(v int64)`
=======
`func (o *NetworkParametersResourceInner) SetLinkSpeedNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetLinkSpeedNumeric sets LinkSpeedNumeric field to given value.

### HasLinkSpeedNumeric

`func (o *NetworkParametersResourceInner) HasLinkSpeedNumeric() bool`

HasLinkSpeedNumeric returns a boolean if a field has been set.

### GetMacAddress

`func (o *NetworkParametersResourceInner) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NetworkParametersResourceInner) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NetworkParametersResourceInner) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NetworkParametersResourceInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPingBroadcast

`func (o *NetworkParametersResourceInner) GetPingBroadcast() string`

GetPingBroadcast returns the PingBroadcast field if non-nil, zero value otherwise.

### GetPingBroadcastOk

`func (o *NetworkParametersResourceInner) GetPingBroadcastOk() (*string, bool)`

GetPingBroadcastOk returns a tuple with the PingBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingBroadcast

`func (o *NetworkParametersResourceInner) SetPingBroadcast(v string)`

SetPingBroadcast sets PingBroadcast field to given value.

### HasPingBroadcast

`func (o *NetworkParametersResourceInner) HasPingBroadcast() bool`

HasPingBroadcast returns a boolean if a field has been set.

### GetPingBroadcastNumeric

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) GetPingBroadcastNumeric() int64`
=======
`func (o *NetworkParametersResourceInner) GetPingBroadcastNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPingBroadcastNumeric returns the PingBroadcastNumeric field if non-nil, zero value otherwise.

### GetPingBroadcastNumericOk

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) GetPingBroadcastNumericOk() (*int64, bool)`
=======
`func (o *NetworkParametersResourceInner) GetPingBroadcastNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPingBroadcastNumericOk returns a tuple with the PingBroadcastNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingBroadcastNumeric

<<<<<<< HEAD
`func (o *NetworkParametersResourceInner) SetPingBroadcastNumeric(v int64)`
=======
`func (o *NetworkParametersResourceInner) SetPingBroadcastNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetPingBroadcastNumeric sets PingBroadcastNumeric field to given value.

### HasPingBroadcastNumeric

`func (o *NetworkParametersResourceInner) HasPingBroadcastNumeric() bool`

HasPingBroadcastNumeric returns a boolean if a field has been set.

### GetSubnetMask

`func (o *NetworkParametersResourceInner) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *NetworkParametersResourceInner) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *NetworkParametersResourceInner) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *NetworkParametersResourceInner) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


