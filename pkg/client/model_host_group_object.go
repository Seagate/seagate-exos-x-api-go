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

// checks if the HostGroupObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostGroupObject{}

// HostGroupObject struct for HostGroupObject
type HostGroupObject struct {
	Status    []StatusResourceInner    `json:"status,omitempty"`
	HostGroup []HostGroupResourceInner `json:"host-group,omitempty"`
}

// NewHostGroupObject instantiates a new HostGroupObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostGroupObject() *HostGroupObject {
	this := HostGroupObject{}
	return &this
}

// NewHostGroupObjectWithDefaults instantiates a new HostGroupObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostGroupObjectWithDefaults() *HostGroupObject {
	this := HostGroupObject{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HostGroupObject) GetStatus() []StatusResourceInner {
	if o == nil || IsNil(o.Status) {
		var ret []StatusResourceInner
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostGroupObject) GetStatusOk() ([]StatusResourceInner, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HostGroupObject) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []StatusResourceInner and assigns it to the Status field.
func (o *HostGroupObject) SetStatus(v []StatusResourceInner) {
	o.Status = v
}

// GetHostGroup returns the HostGroup field value if set, zero value otherwise.
func (o *HostGroupObject) GetHostGroup() []HostGroupResourceInner {
	if o == nil || IsNil(o.HostGroup) {
		var ret []HostGroupResourceInner
		return ret
	}
	return o.HostGroup
}

// GetHostGroupOk returns a tuple with the HostGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostGroupObject) GetHostGroupOk() ([]HostGroupResourceInner, bool) {
	if o == nil || IsNil(o.HostGroup) {
		return nil, false
	}
	return o.HostGroup, true
}

// HasHostGroup returns a boolean if a field has been set.
func (o *HostGroupObject) HasHostGroup() bool {
	if o != nil && !IsNil(o.HostGroup) {
		return true
	}

	return false
}

// SetHostGroup gets a reference to the given []HostGroupResourceInner and assigns it to the HostGroup field.
func (o *HostGroupObject) SetHostGroup(v []HostGroupResourceInner) {
	o.HostGroup = v
}

func (o HostGroupObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostGroupObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.HostGroup) {
		toSerialize["host-group"] = o.HostGroup
	}
	return toSerialize, nil
}

type NullableHostGroupObject struct {
	value *HostGroupObject
	isSet bool
}

func (v NullableHostGroupObject) Get() *HostGroupObject {
	return v.value
}

func (v *NullableHostGroupObject) Set(val *HostGroupObject) {
	v.value = val
	v.isSet = true
}

func (v NullableHostGroupObject) IsSet() bool {
	return v.isSet
}

func (v *NullableHostGroupObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostGroupObject(val *HostGroupObject) *NullableHostGroupObject {
	return &NullableHostGroupObject{value: val, isSet: true}
}

func (v NullableHostGroupObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostGroupObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
