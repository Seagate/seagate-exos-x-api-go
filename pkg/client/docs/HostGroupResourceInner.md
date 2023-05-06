# HostGroupResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
<<<<<<< HEAD
**MemberCount** | Pointer to **int64** | Number of members | [optional] 
=======
**MemberCount** | Pointer to **int32** | Number of members | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
**Name** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | The resource URL | [optional] 
**Host** | Pointer to [**[]HostResourceInner**](HostResourceInner.md) |  | [optional] 

## Methods

### NewHostGroupResourceInner

`func NewHostGroupResourceInner() *HostGroupResourceInner`

NewHostGroupResourceInner instantiates a new HostGroupResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostGroupResourceInnerWithDefaults

`func NewHostGroupResourceInnerWithDefaults() *HostGroupResourceInner`

NewHostGroupResourceInnerWithDefaults instantiates a new HostGroupResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *HostGroupResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *HostGroupResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *HostGroupResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *HostGroupResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *HostGroupResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HostGroupResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HostGroupResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HostGroupResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetDurableId

`func (o *HostGroupResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *HostGroupResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *HostGroupResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *HostGroupResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetMemberCount

<<<<<<< HEAD
`func (o *HostGroupResourceInner) GetMemberCount() int64`
=======
`func (o *HostGroupResourceInner) GetMemberCount() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

<<<<<<< HEAD
`func (o *HostGroupResourceInner) GetMemberCountOk() (*int64, bool)`
=======
`func (o *HostGroupResourceInner) GetMemberCountOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

<<<<<<< HEAD
`func (o *HostGroupResourceInner) SetMemberCount(v int64)`
=======
`func (o *HostGroupResourceInner) SetMemberCount(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *HostGroupResourceInner) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetName

`func (o *HostGroupResourceInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostGroupResourceInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostGroupResourceInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HostGroupResourceInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerialNumber

`func (o *HostGroupResourceInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *HostGroupResourceInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *HostGroupResourceInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *HostGroupResourceInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetUrl

`func (o *HostGroupResourceInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HostGroupResourceInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HostGroupResourceInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HostGroupResourceInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHost

`func (o *HostGroupResourceInner) GetHost() []HostResourceInner`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HostGroupResourceInner) GetHostOk() (*[]HostResourceInner, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HostGroupResourceInner) SetHost(v []HostResourceInner)`

SetHost sets Host field to given value.

### HasHost

`func (o *HostGroupResourceInner) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


