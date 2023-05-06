# InitiatorResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**Discovered** | Pointer to **string** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**HostBusType** | Pointer to **string** |  | [optional] 
**HostBusTypeNumeric** | Pointer to **int32** |  | [optional] 
**HostId** | Pointer to **string** |  | [optional] 
**HostKey** | Pointer to **string** | Durable ID of a Host | [optional] 
**HostPortBitsA** | Pointer to **int32** |  | [optional] 
**HostPortBitsB** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Mapped** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** | User-defined alias for the mapped host | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**ProfileNumeric** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** | The resource URL | [optional] 

## Methods

### NewInitiatorResourceInner

`func NewInitiatorResourceInner() *InitiatorResourceInner`

NewInitiatorResourceInner instantiates a new InitiatorResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiatorResourceInnerWithDefaults

`func NewInitiatorResourceInnerWithDefaults() *InitiatorResourceInner`

NewInitiatorResourceInnerWithDefaults instantiates a new InitiatorResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *InitiatorResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *InitiatorResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *InitiatorResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *InitiatorResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *InitiatorResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InitiatorResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InitiatorResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InitiatorResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetDiscovered

`func (o *InitiatorResourceInner) GetDiscovered() string`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *InitiatorResourceInner) GetDiscoveredOk() (*string, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *InitiatorResourceInner) SetDiscovered(v string)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *InitiatorResourceInner) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetDurableId

`func (o *InitiatorResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *InitiatorResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *InitiatorResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *InitiatorResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetHostBusType

`func (o *InitiatorResourceInner) GetHostBusType() string`

GetHostBusType returns the HostBusType field if non-nil, zero value otherwise.

### GetHostBusTypeOk

`func (o *InitiatorResourceInner) GetHostBusTypeOk() (*string, bool)`

GetHostBusTypeOk returns a tuple with the HostBusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostBusType

`func (o *InitiatorResourceInner) SetHostBusType(v string)`

SetHostBusType sets HostBusType field to given value.

### HasHostBusType

`func (o *InitiatorResourceInner) HasHostBusType() bool`

HasHostBusType returns a boolean if a field has been set.

### GetHostBusTypeNumeric

`func (o *InitiatorResourceInner) GetHostBusTypeNumeric() int32`

GetHostBusTypeNumeric returns the HostBusTypeNumeric field if non-nil, zero value otherwise.

### GetHostBusTypeNumericOk

`func (o *InitiatorResourceInner) GetHostBusTypeNumericOk() (*int32, bool)`

GetHostBusTypeNumericOk returns a tuple with the HostBusTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostBusTypeNumeric

`func (o *InitiatorResourceInner) SetHostBusTypeNumeric(v int32)`

SetHostBusTypeNumeric sets HostBusTypeNumeric field to given value.

### HasHostBusTypeNumeric

`func (o *InitiatorResourceInner) HasHostBusTypeNumeric() bool`

HasHostBusTypeNumeric returns a boolean if a field has been set.

### GetHostId

`func (o *InitiatorResourceInner) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *InitiatorResourceInner) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *InitiatorResourceInner) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *InitiatorResourceInner) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetHostKey

`func (o *InitiatorResourceInner) GetHostKey() string`

GetHostKey returns the HostKey field if non-nil, zero value otherwise.

### GetHostKeyOk

`func (o *InitiatorResourceInner) GetHostKeyOk() (*string, bool)`

GetHostKeyOk returns a tuple with the HostKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostKey

`func (o *InitiatorResourceInner) SetHostKey(v string)`

SetHostKey sets HostKey field to given value.

### HasHostKey

`func (o *InitiatorResourceInner) HasHostKey() bool`

HasHostKey returns a boolean if a field has been set.

### GetHostPortBitsA

`func (o *InitiatorResourceInner) GetHostPortBitsA() int32`

GetHostPortBitsA returns the HostPortBitsA field if non-nil, zero value otherwise.

### GetHostPortBitsAOk

`func (o *InitiatorResourceInner) GetHostPortBitsAOk() (*int32, bool)`

GetHostPortBitsAOk returns a tuple with the HostPortBitsA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPortBitsA

`func (o *InitiatorResourceInner) SetHostPortBitsA(v int32)`

SetHostPortBitsA sets HostPortBitsA field to given value.

### HasHostPortBitsA

`func (o *InitiatorResourceInner) HasHostPortBitsA() bool`

HasHostPortBitsA returns a boolean if a field has been set.

### GetHostPortBitsB

`func (o *InitiatorResourceInner) GetHostPortBitsB() int32`

GetHostPortBitsB returns the HostPortBitsB field if non-nil, zero value otherwise.

### GetHostPortBitsBOk

`func (o *InitiatorResourceInner) GetHostPortBitsBOk() (*int32, bool)`

GetHostPortBitsBOk returns a tuple with the HostPortBitsB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPortBitsB

`func (o *InitiatorResourceInner) SetHostPortBitsB(v int32)`

SetHostPortBitsB sets HostPortBitsB field to given value.

### HasHostPortBitsB

`func (o *InitiatorResourceInner) HasHostPortBitsB() bool`

HasHostPortBitsB returns a boolean if a field has been set.

### GetId

`func (o *InitiatorResourceInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InitiatorResourceInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InitiatorResourceInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InitiatorResourceInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMapped

`func (o *InitiatorResourceInner) GetMapped() string`

GetMapped returns the Mapped field if non-nil, zero value otherwise.

### GetMappedOk

`func (o *InitiatorResourceInner) GetMappedOk() (*string, bool)`

GetMappedOk returns a tuple with the Mapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapped

`func (o *InitiatorResourceInner) SetMapped(v string)`

SetMapped sets Mapped field to given value.

### HasMapped

`func (o *InitiatorResourceInner) HasMapped() bool`

HasMapped returns a boolean if a field has been set.

### GetNickname

`func (o *InitiatorResourceInner) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *InitiatorResourceInner) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *InitiatorResourceInner) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *InitiatorResourceInner) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetProfile

`func (o *InitiatorResourceInner) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InitiatorResourceInner) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InitiatorResourceInner) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *InitiatorResourceInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetProfileNumeric

`func (o *InitiatorResourceInner) GetProfileNumeric() int32`

GetProfileNumeric returns the ProfileNumeric field if non-nil, zero value otherwise.

### GetProfileNumericOk

`func (o *InitiatorResourceInner) GetProfileNumericOk() (*int32, bool)`

GetProfileNumericOk returns a tuple with the ProfileNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileNumeric

`func (o *InitiatorResourceInner) SetProfileNumeric(v int32)`

SetProfileNumeric sets ProfileNumeric field to given value.

### HasProfileNumeric

`func (o *InitiatorResourceInner) HasProfileNumeric() bool`

HasProfileNumeric returns a boolean if a field has been set.

### GetUrl

`func (o *InitiatorResourceInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InitiatorResourceInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InitiatorResourceInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InitiatorResourceInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


