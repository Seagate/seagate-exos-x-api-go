/*
Seagate Management Controller (MC) OpenAPI

This API allows users to interact through the MC API to provision and manage storage.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the InitiatorViewResourceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitiatorViewResourceInner{}

// InitiatorViewResourceInner struct for InitiatorViewResourceInner
type InitiatorViewResourceInner struct {
	ObjectName         *string                   `json:"object-name,omitempty"`
	Meta               *string                   `json:"meta,omitempty"`
	HbaNickname        *string                   `json:"hba-nickname,omitempty"`
	HostProfile        *string                   `json:"host-profile,omitempty"`
	HostProfileNumeric *int32                    `json:"host-profile-numeric,omitempty"`
	Id                 *string                   `json:"id,omitempty"`
	VolumeView         []VolumeViewResourceInner `json:"volume-view,omitempty"`
}

// NewInitiatorViewResourceInner instantiates a new InitiatorViewResourceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitiatorViewResourceInner() *InitiatorViewResourceInner {
	this := InitiatorViewResourceInner{}
	return &this
}

// NewInitiatorViewResourceInnerWithDefaults instantiates a new InitiatorViewResourceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitiatorViewResourceInnerWithDefaults() *InitiatorViewResourceInner {
	this := InitiatorViewResourceInner{}
	return &this
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *InitiatorViewResourceInner) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitiatorViewResourceInner) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *InitiatorViewResourceInner) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *InitiatorViewResourceInner) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InitiatorViewResourceInner) GetMeta() string {
	if o == nil || IsNil(o.Meta) {
		var ret string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitiatorViewResourceInner) GetMetaOk() (*string, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InitiatorViewResourceInner) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given string and assigns it to the Meta field.
func (o *InitiatorViewResourceInner) SetMeta(v string) {
	o.Meta = &v
}

// GetHbaNickname returns the HbaNickname field value if set, zero value otherwise.
func (o *InitiatorViewResourceInner) GetHbaNickname() string {
	if o == nil || IsNil(o.HbaNickname) {
		var ret string
		return ret
	}
	return *o.HbaNickname
}

// GetHbaNicknameOk returns a tuple with the HbaNickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitiatorViewResourceInner) GetHbaNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.HbaNickname) {
		return nil, false
	}
	return o.HbaNickname, true
}

// HasHbaNickname returns a boolean if a field has been set.
func (o *InitiatorViewResourceInner) HasHbaNickname() bool {
	if o != nil && !IsNil(o.HbaNickname) {
		return true
	}

	return false
}

// SetHbaNickname gets a reference to the given string and assigns it to the HbaNickname field.
func (o *InitiatorViewResourceInner) SetHbaNickname(v string) {
	o.HbaNickname = &v
}

// GetHostProfile returns the HostProfile field value if set, zero value otherwise.
func (o *InitiatorViewResourceInner) GetHostProfile() string {
	if o == nil || IsNil(o.HostProfile) {
		var ret string
		return ret
	}
	return *o.HostProfile
}

// GetHostProfileOk returns a tuple with the HostProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitiatorViewResourceInner) GetHostProfileOk() (*string, bool) {
	if o == nil || IsNil(o.HostProfile) {
		return nil, false
	}
	return o.HostProfile, true
}

// HasHostProfile returns a boolean if a field has been set.
func (o *InitiatorViewResourceInner) HasHostProfile() bool {
	if o != nil && !IsNil(o.HostProfile) {
		return true
	}

	return false
}

// SetHostProfile gets a reference to the given string and assigns it to the HostProfile field.
func (o *InitiatorViewResourceInner) SetHostProfile(v string) {
	o.HostProfile = &v
}

// GetHostProfileNumeric returns the HostProfileNumeric field value if set, zero value otherwise.
func (o *InitiatorViewResourceInner) GetHostProfileNumeric() int32 {
	if o == nil || IsNil(o.HostProfileNumeric) {
		var ret int32
		return ret
	}
	return *o.HostProfileNumeric
}

// GetHostProfileNumericOk returns a tuple with the HostProfileNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitiatorViewResourceInner) GetHostProfileNumericOk() (*int32, bool) {
	if o == nil || IsNil(o.HostProfileNumeric) {
		return nil, false
	}
	return o.HostProfileNumeric, true
}

// HasHostProfileNumeric returns a boolean if a field has been set.
func (o *InitiatorViewResourceInner) HasHostProfileNumeric() bool {
	if o != nil && !IsNil(o.HostProfileNumeric) {
		return true
	}

	return false
}

// SetHostProfileNumeric gets a reference to the given int32 and assigns it to the HostProfileNumeric field.
func (o *InitiatorViewResourceInner) SetHostProfileNumeric(v int32) {
	o.HostProfileNumeric = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InitiatorViewResourceInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitiatorViewResourceInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InitiatorViewResourceInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InitiatorViewResourceInner) SetId(v string) {
	o.Id = &v
}

// GetVolumeView returns the VolumeView field value if set, zero value otherwise.
func (o *InitiatorViewResourceInner) GetVolumeView() []VolumeViewResourceInner {
	if o == nil || IsNil(o.VolumeView) {
		var ret []VolumeViewResourceInner
		return ret
	}
	return o.VolumeView
}

// GetVolumeViewOk returns a tuple with the VolumeView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitiatorViewResourceInner) GetVolumeViewOk() ([]VolumeViewResourceInner, bool) {
	if o == nil || IsNil(o.VolumeView) {
		return nil, false
	}
	return o.VolumeView, true
}

// HasVolumeView returns a boolean if a field has been set.
func (o *InitiatorViewResourceInner) HasVolumeView() bool {
	if o != nil && !IsNil(o.VolumeView) {
		return true
	}

	return false
}

// SetVolumeView gets a reference to the given []VolumeViewResourceInner and assigns it to the VolumeView field.
func (o *InitiatorViewResourceInner) SetVolumeView(v []VolumeViewResourceInner) {
	o.VolumeView = v
}

func (o InitiatorViewResourceInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InitiatorViewResourceInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectName) {
		toSerialize["object-name"] = o.ObjectName
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.HbaNickname) {
		toSerialize["hba-nickname"] = o.HbaNickname
	}
	if !IsNil(o.HostProfile) {
		toSerialize["host-profile"] = o.HostProfile
	}
	if !IsNil(o.HostProfileNumeric) {
		toSerialize["host-profile-numeric"] = o.HostProfileNumeric
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.VolumeView) {
		toSerialize["volume-view"] = o.VolumeView
	}
	return toSerialize, nil
}

type NullableInitiatorViewResourceInner struct {
	value *InitiatorViewResourceInner
	isSet bool
}

func (v NullableInitiatorViewResourceInner) Get() *InitiatorViewResourceInner {
	return v.value
}

func (v *NullableInitiatorViewResourceInner) Set(val *InitiatorViewResourceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInitiatorViewResourceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInitiatorViewResourceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitiatorViewResourceInner(val *InitiatorViewResourceInner) *NullableInitiatorViewResourceInner {
	return &NullableInitiatorViewResourceInner{value: val, isSet: true}
}

func (v NullableInitiatorViewResourceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitiatorViewResourceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}