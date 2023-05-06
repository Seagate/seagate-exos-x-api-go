# HostResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**GroupKey** | Pointer to **string** | Durable ID of a Management Group | [optional] 
**HostGroup** | Pointer to **string** |  | [optional] 
<<<<<<< HEAD
**MemberCount** | Pointer to **int64** | Number of members | [optional] 
=======
**MemberCount** | Pointer to **int32** | Number of members | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
**Name** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Initiator** | Pointer to [**[]InitiatorResourceInner**](InitiatorResourceInner.md) |  | [optional] 

## Methods

### NewHostResourceInner

`func NewHostResourceInner() *HostResourceInner`

NewHostResourceInner instantiates a new HostResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostResourceInnerWithDefaults

`func NewHostResourceInnerWithDefaults() *HostResourceInner`

NewHostResourceInnerWithDefaults instantiates a new HostResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *HostResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *HostResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *HostResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *HostResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *HostResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HostResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HostResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HostResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetDurableId

`func (o *HostResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *HostResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *HostResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *HostResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetGroupKey

`func (o *HostResourceInner) GetGroupKey() string`

GetGroupKey returns the GroupKey field if non-nil, zero value otherwise.

### GetGroupKeyOk

`func (o *HostResourceInner) GetGroupKeyOk() (*string, bool)`

GetGroupKeyOk returns a tuple with the GroupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKey

`func (o *HostResourceInner) SetGroupKey(v string)`

SetGroupKey sets GroupKey field to given value.

### HasGroupKey

`func (o *HostResourceInner) HasGroupKey() bool`

HasGroupKey returns a boolean if a field has been set.

### GetHostGroup

`func (o *HostResourceInner) GetHostGroup() string`

GetHostGroup returns the HostGroup field if non-nil, zero value otherwise.

### GetHostGroupOk

`func (o *HostResourceInner) GetHostGroupOk() (*string, bool)`

GetHostGroupOk returns a tuple with the HostGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroup

`func (o *HostResourceInner) SetHostGroup(v string)`

SetHostGroup sets HostGroup field to given value.

### HasHostGroup

`func (o *HostResourceInner) HasHostGroup() bool`

HasHostGroup returns a boolean if a field has been set.

### GetMemberCount

<<<<<<< HEAD
`func (o *HostResourceInner) GetMemberCount() int64`
=======
`func (o *HostResourceInner) GetMemberCount() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

<<<<<<< HEAD
`func (o *HostResourceInner) GetMemberCountOk() (*int64, bool)`
=======
`func (o *HostResourceInner) GetMemberCountOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

<<<<<<< HEAD
`func (o *HostResourceInner) SetMemberCount(v int64)`
=======
`func (o *HostResourceInner) SetMemberCount(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *HostResourceInner) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetName

`func (o *HostResourceInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostResourceInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostResourceInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HostResourceInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerialNumber

`func (o *HostResourceInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *HostResourceInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *HostResourceInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *HostResourceInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetInitiator

`func (o *HostResourceInner) GetInitiator() []InitiatorResourceInner`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *HostResourceInner) GetInitiatorOk() (*[]InitiatorResourceInner, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *HostResourceInner) SetInitiator(v []InitiatorResourceInner)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *HostResourceInner) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


