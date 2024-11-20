# HostGroupViewResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**GroupName** | Pointer to **string** | Name of a Management Group | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**HostViewMappings** | Pointer to [**[]HostViewMappingsResourceInner**](HostViewMappingsResourceInner.md) |  | [optional] 

## Methods

### NewHostGroupViewResourceInner

`func NewHostGroupViewResourceInner() *HostGroupViewResourceInner`

NewHostGroupViewResourceInner instantiates a new HostGroupViewResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostGroupViewResourceInnerWithDefaults

`func NewHostGroupViewResourceInnerWithDefaults() *HostGroupViewResourceInner`

NewHostGroupViewResourceInnerWithDefaults instantiates a new HostGroupViewResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *HostGroupViewResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *HostGroupViewResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *HostGroupViewResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *HostGroupViewResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *HostGroupViewResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HostGroupViewResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HostGroupViewResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HostGroupViewResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetDurableId

`func (o *HostGroupViewResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *HostGroupViewResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *HostGroupViewResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *HostGroupViewResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetGroupName

`func (o *HostGroupViewResourceInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *HostGroupViewResourceInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *HostGroupViewResourceInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *HostGroupViewResourceInner) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetSerialNumber

`func (o *HostGroupViewResourceInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *HostGroupViewResourceInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *HostGroupViewResourceInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *HostGroupViewResourceInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetStatus

`func (o *HostGroupViewResourceInner) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostGroupViewResourceInner) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostGroupViewResourceInner) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HostGroupViewResourceInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHostViewMappings

`func (o *HostGroupViewResourceInner) GetHostViewMappings() []HostViewMappingsResourceInner`

GetHostViewMappings returns the HostViewMappings field if non-nil, zero value otherwise.

### GetHostViewMappingsOk

`func (o *HostGroupViewResourceInner) GetHostViewMappingsOk() (*[]HostViewMappingsResourceInner, bool)`

GetHostViewMappingsOk returns a tuple with the HostViewMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostViewMappings

`func (o *HostGroupViewResourceInner) SetHostViewMappings(v []HostViewMappingsResourceInner)`

SetHostViewMappings sets HostViewMappings field to given value.

### HasHostViewMappings

`func (o *HostGroupViewResourceInner) HasHostViewMappings() bool`

HasHostViewMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


