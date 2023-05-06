# VolumeViewMappingsResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**Access** | Pointer to **string** | Access rights for this volume | [optional] 
<<<<<<< HEAD
**AccessNumeric** | Pointer to **int64** | Access rights for this volume( In numeric form ) | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**HostProfile** | Pointer to **string** |  | [optional] 
**HostProfileNumeric** | Pointer to **int64** |  | [optional] 
=======
**AccessNumeric** | Pointer to **int32** | Access rights for this volume( In numeric form ) | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**HostProfile** | Pointer to **string** |  | [optional] 
**HostProfileNumeric** | Pointer to **int32** |  | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
**Identifier** | Pointer to **string** | WWPN or IQN or Host Serial Number or Host Group Serial Number | [optional] 
**InitiatorsUrl** | Pointer to **string** |  | [optional] 
**Lun** | Pointer to **string** | Logical Unit Number | [optional] 
**MappedId** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** | User-defined alias for the mapped host | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to **string** |  | [optional] 

## Methods

### NewVolumeViewMappingsResourceInner

`func NewVolumeViewMappingsResourceInner() *VolumeViewMappingsResourceInner`

NewVolumeViewMappingsResourceInner instantiates a new VolumeViewMappingsResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeViewMappingsResourceInnerWithDefaults

`func NewVolumeViewMappingsResourceInnerWithDefaults() *VolumeViewMappingsResourceInner`

NewVolumeViewMappingsResourceInnerWithDefaults instantiates a new VolumeViewMappingsResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *VolumeViewMappingsResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *VolumeViewMappingsResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *VolumeViewMappingsResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *VolumeViewMappingsResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *VolumeViewMappingsResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VolumeViewMappingsResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VolumeViewMappingsResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *VolumeViewMappingsResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetAccess

`func (o *VolumeViewMappingsResourceInner) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *VolumeViewMappingsResourceInner) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *VolumeViewMappingsResourceInner) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *VolumeViewMappingsResourceInner) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetAccessNumeric

<<<<<<< HEAD
`func (o *VolumeViewMappingsResourceInner) GetAccessNumeric() int64`
=======
`func (o *VolumeViewMappingsResourceInner) GetAccessNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAccessNumeric returns the AccessNumeric field if non-nil, zero value otherwise.

### GetAccessNumericOk

<<<<<<< HEAD
`func (o *VolumeViewMappingsResourceInner) GetAccessNumericOk() (*int64, bool)`
=======
`func (o *VolumeViewMappingsResourceInner) GetAccessNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAccessNumericOk returns a tuple with the AccessNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessNumeric

<<<<<<< HEAD
`func (o *VolumeViewMappingsResourceInner) SetAccessNumeric(v int64)`
=======
`func (o *VolumeViewMappingsResourceInner) SetAccessNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetAccessNumeric sets AccessNumeric field to given value.

### HasAccessNumeric

`func (o *VolumeViewMappingsResourceInner) HasAccessNumeric() bool`

HasAccessNumeric returns a boolean if a field has been set.

### GetDurableId

`func (o *VolumeViewMappingsResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *VolumeViewMappingsResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *VolumeViewMappingsResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *VolumeViewMappingsResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetHostProfile

`func (o *VolumeViewMappingsResourceInner) GetHostProfile() string`

GetHostProfile returns the HostProfile field if non-nil, zero value otherwise.

### GetHostProfileOk

`func (o *VolumeViewMappingsResourceInner) GetHostProfileOk() (*string, bool)`

GetHostProfileOk returns a tuple with the HostProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostProfile

`func (o *VolumeViewMappingsResourceInner) SetHostProfile(v string)`

SetHostProfile sets HostProfile field to given value.

### HasHostProfile

`func (o *VolumeViewMappingsResourceInner) HasHostProfile() bool`

HasHostProfile returns a boolean if a field has been set.

### GetHostProfileNumeric

<<<<<<< HEAD
`func (o *VolumeViewMappingsResourceInner) GetHostProfileNumeric() int64`
=======
`func (o *VolumeViewMappingsResourceInner) GetHostProfileNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHostProfileNumeric returns the HostProfileNumeric field if non-nil, zero value otherwise.

### GetHostProfileNumericOk

<<<<<<< HEAD
`func (o *VolumeViewMappingsResourceInner) GetHostProfileNumericOk() (*int64, bool)`
=======
`func (o *VolumeViewMappingsResourceInner) GetHostProfileNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHostProfileNumericOk returns a tuple with the HostProfileNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostProfileNumeric

<<<<<<< HEAD
`func (o *VolumeViewMappingsResourceInner) SetHostProfileNumeric(v int64)`
=======
`func (o *VolumeViewMappingsResourceInner) SetHostProfileNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetHostProfileNumeric sets HostProfileNumeric field to given value.

### HasHostProfileNumeric

`func (o *VolumeViewMappingsResourceInner) HasHostProfileNumeric() bool`

HasHostProfileNumeric returns a boolean if a field has been set.

### GetIdentifier

`func (o *VolumeViewMappingsResourceInner) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VolumeViewMappingsResourceInner) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VolumeViewMappingsResourceInner) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *VolumeViewMappingsResourceInner) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetInitiatorsUrl

`func (o *VolumeViewMappingsResourceInner) GetInitiatorsUrl() string`

GetInitiatorsUrl returns the InitiatorsUrl field if non-nil, zero value otherwise.

### GetInitiatorsUrlOk

`func (o *VolumeViewMappingsResourceInner) GetInitiatorsUrlOk() (*string, bool)`

GetInitiatorsUrlOk returns a tuple with the InitiatorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorsUrl

`func (o *VolumeViewMappingsResourceInner) SetInitiatorsUrl(v string)`

SetInitiatorsUrl sets InitiatorsUrl field to given value.

### HasInitiatorsUrl

`func (o *VolumeViewMappingsResourceInner) HasInitiatorsUrl() bool`

HasInitiatorsUrl returns a boolean if a field has been set.

### GetLun

`func (o *VolumeViewMappingsResourceInner) GetLun() string`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *VolumeViewMappingsResourceInner) GetLunOk() (*string, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *VolumeViewMappingsResourceInner) SetLun(v string)`

SetLun sets Lun field to given value.

### HasLun

`func (o *VolumeViewMappingsResourceInner) HasLun() bool`

HasLun returns a boolean if a field has been set.

### GetMappedId

`func (o *VolumeViewMappingsResourceInner) GetMappedId() string`

GetMappedId returns the MappedId field if non-nil, zero value otherwise.

### GetMappedIdOk

`func (o *VolumeViewMappingsResourceInner) GetMappedIdOk() (*string, bool)`

GetMappedIdOk returns a tuple with the MappedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedId

`func (o *VolumeViewMappingsResourceInner) SetMappedId(v string)`

SetMappedId sets MappedId field to given value.

### HasMappedId

`func (o *VolumeViewMappingsResourceInner) HasMappedId() bool`

HasMappedId returns a boolean if a field has been set.

### GetNickname

`func (o *VolumeViewMappingsResourceInner) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *VolumeViewMappingsResourceInner) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *VolumeViewMappingsResourceInner) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *VolumeViewMappingsResourceInner) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetParentId

`func (o *VolumeViewMappingsResourceInner) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *VolumeViewMappingsResourceInner) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *VolumeViewMappingsResourceInner) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *VolumeViewMappingsResourceInner) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPorts

`func (o *VolumeViewMappingsResourceInner) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *VolumeViewMappingsResourceInner) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *VolumeViewMappingsResourceInner) SetPorts(v string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *VolumeViewMappingsResourceInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


