# VolumeViewResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | The resource URL | [optional] 
**VolumeName** | Pointer to **string** | User-defined name for the volume | [optional] 
**VolumeSerial** | Pointer to **string** | Unique serial number for the volume | [optional] 
**VolumesUrl** | Pointer to **string** | Volumes URL | [optional] 
**VolumeViewMappings** | Pointer to [**[]VolumeViewMappingsResourceInner**](VolumeViewMappingsResourceInner.md) |  | [optional] 

## Methods

### NewVolumeViewResourceInner

`func NewVolumeViewResourceInner() *VolumeViewResourceInner`

NewVolumeViewResourceInner instantiates a new VolumeViewResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeViewResourceInnerWithDefaults

`func NewVolumeViewResourceInnerWithDefaults() *VolumeViewResourceInner`

NewVolumeViewResourceInnerWithDefaults instantiates a new VolumeViewResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *VolumeViewResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *VolumeViewResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *VolumeViewResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *VolumeViewResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *VolumeViewResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VolumeViewResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VolumeViewResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *VolumeViewResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetDurableId

`func (o *VolumeViewResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *VolumeViewResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *VolumeViewResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *VolumeViewResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetUrl

`func (o *VolumeViewResourceInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VolumeViewResourceInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VolumeViewResourceInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VolumeViewResourceInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVolumeName

`func (o *VolumeViewResourceInner) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *VolumeViewResourceInner) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *VolumeViewResourceInner) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *VolumeViewResourceInner) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumeSerial

`func (o *VolumeViewResourceInner) GetVolumeSerial() string`

GetVolumeSerial returns the VolumeSerial field if non-nil, zero value otherwise.

### GetVolumeSerialOk

`func (o *VolumeViewResourceInner) GetVolumeSerialOk() (*string, bool)`

GetVolumeSerialOk returns a tuple with the VolumeSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSerial

`func (o *VolumeViewResourceInner) SetVolumeSerial(v string)`

SetVolumeSerial sets VolumeSerial field to given value.

### HasVolumeSerial

`func (o *VolumeViewResourceInner) HasVolumeSerial() bool`

HasVolumeSerial returns a boolean if a field has been set.

### GetVolumesUrl

`func (o *VolumeViewResourceInner) GetVolumesUrl() string`

GetVolumesUrl returns the VolumesUrl field if non-nil, zero value otherwise.

### GetVolumesUrlOk

`func (o *VolumeViewResourceInner) GetVolumesUrlOk() (*string, bool)`

GetVolumesUrlOk returns a tuple with the VolumesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesUrl

`func (o *VolumeViewResourceInner) SetVolumesUrl(v string)`

SetVolumesUrl sets VolumesUrl field to given value.

### HasVolumesUrl

`func (o *VolumeViewResourceInner) HasVolumesUrl() bool`

HasVolumesUrl returns a boolean if a field has been set.

### GetVolumeViewMappings

`func (o *VolumeViewResourceInner) GetVolumeViewMappings() []VolumeViewMappingsResourceInner`

GetVolumeViewMappings returns the VolumeViewMappings field if non-nil, zero value otherwise.

### GetVolumeViewMappingsOk

`func (o *VolumeViewResourceInner) GetVolumeViewMappingsOk() (*[]VolumeViewMappingsResourceInner, bool)`

GetVolumeViewMappingsOk returns a tuple with the VolumeViewMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeViewMappings

`func (o *VolumeViewResourceInner) SetVolumeViewMappings(v []VolumeViewMappingsResourceInner)`

SetVolumeViewMappings sets VolumeViewMappings field to given value.

### HasVolumeViewMappings

`func (o *VolumeViewResourceInner) HasVolumeViewMappings() bool`

HasVolumeViewMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


