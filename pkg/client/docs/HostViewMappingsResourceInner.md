# HostViewMappingsResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**Access** | Pointer to **string** | Access rights for this host | [optional] 
**AccessNumeric** | Pointer to **int64** | Access rights for this host( In numeric form ) | [optional] 
**Lun** | Pointer to **string** | Logical Unit Number | [optional] 
**Ports** | Pointer to **string** |  | [optional] 
**VolumeName** | Pointer to **string** | User-defined name for the volume | [optional] 
**VolumeSerial** | Pointer to **string** | Unique serial number for the volume | [optional] 

## Methods

### NewHostViewMappingsResourceInner

`func NewHostViewMappingsResourceInner() *HostViewMappingsResourceInner`

NewHostViewMappingsResourceInner instantiates a new HostViewMappingsResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostViewMappingsResourceInnerWithDefaults

`func NewHostViewMappingsResourceInnerWithDefaults() *HostViewMappingsResourceInner`

NewHostViewMappingsResourceInnerWithDefaults instantiates a new HostViewMappingsResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *HostViewMappingsResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *HostViewMappingsResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *HostViewMappingsResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *HostViewMappingsResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *HostViewMappingsResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HostViewMappingsResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HostViewMappingsResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HostViewMappingsResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetAccess

`func (o *HostViewMappingsResourceInner) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *HostViewMappingsResourceInner) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *HostViewMappingsResourceInner) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *HostViewMappingsResourceInner) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetAccessNumeric

`func (o *HostViewMappingsResourceInner) GetAccessNumeric() int64`

GetAccessNumeric returns the AccessNumeric field if non-nil, zero value otherwise.

### GetAccessNumericOk

`func (o *HostViewMappingsResourceInner) GetAccessNumericOk() (*int64, bool)`

GetAccessNumericOk returns a tuple with the AccessNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessNumeric

`func (o *HostViewMappingsResourceInner) SetAccessNumeric(v int64)`

SetAccessNumeric sets AccessNumeric field to given value.

### HasAccessNumeric

`func (o *HostViewMappingsResourceInner) HasAccessNumeric() bool`

HasAccessNumeric returns a boolean if a field has been set.

### GetLun

`func (o *HostViewMappingsResourceInner) GetLun() string`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *HostViewMappingsResourceInner) GetLunOk() (*string, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *HostViewMappingsResourceInner) SetLun(v string)`

SetLun sets Lun field to given value.

### HasLun

`func (o *HostViewMappingsResourceInner) HasLun() bool`

HasLun returns a boolean if a field has been set.

### GetPorts

`func (o *HostViewMappingsResourceInner) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *HostViewMappingsResourceInner) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *HostViewMappingsResourceInner) SetPorts(v string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *HostViewMappingsResourceInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetVolumeName

`func (o *HostViewMappingsResourceInner) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *HostViewMappingsResourceInner) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *HostViewMappingsResourceInner) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *HostViewMappingsResourceInner) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumeSerial

`func (o *HostViewMappingsResourceInner) GetVolumeSerial() string`

GetVolumeSerial returns the VolumeSerial field if non-nil, zero value otherwise.

### GetVolumeSerialOk

`func (o *HostViewMappingsResourceInner) GetVolumeSerialOk() (*string, bool)`

GetVolumeSerialOk returns a tuple with the VolumeSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSerial

`func (o *HostViewMappingsResourceInner) SetVolumeSerial(v string)`

SetVolumeSerial sets VolumeSerial field to given value.

### HasVolumeSerial

`func (o *HostViewMappingsResourceInner) HasVolumeSerial() bool`

HasVolumeSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


