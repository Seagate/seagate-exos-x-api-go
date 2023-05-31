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

// checks if the VolumeViewObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeViewObject{}

// VolumeViewObject struct for VolumeViewObject
type VolumeViewObject struct {
	Status          []StatusResourceInner          `json:"status,omitempty"`
	VolumeGroupView []VolumeGroupViewResourceInner `json:"volume-group-view,omitempty"`
	VolumeView      []VolumeViewResourceInner      `json:"volume-view,omitempty"`
}

// NewVolumeViewObject instantiates a new VolumeViewObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeViewObject() *VolumeViewObject {
	this := VolumeViewObject{}
	return &this
}

// NewVolumeViewObjectWithDefaults instantiates a new VolumeViewObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeViewObjectWithDefaults() *VolumeViewObject {
	this := VolumeViewObject{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VolumeViewObject) GetStatus() []StatusResourceInner {
	if o == nil || IsNil(o.Status) {
		var ret []StatusResourceInner
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewObject) GetStatusOk() ([]StatusResourceInner, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VolumeViewObject) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []StatusResourceInner and assigns it to the Status field.
func (o *VolumeViewObject) SetStatus(v []StatusResourceInner) {
	o.Status = v
}

// GetVolumeGroupView returns the VolumeGroupView field value if set, zero value otherwise.
func (o *VolumeViewObject) GetVolumeGroupView() []VolumeGroupViewResourceInner {
	if o == nil || IsNil(o.VolumeGroupView) {
		var ret []VolumeGroupViewResourceInner
		return ret
	}
	return o.VolumeGroupView
}

// GetVolumeGroupViewOk returns a tuple with the VolumeGroupView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewObject) GetVolumeGroupViewOk() ([]VolumeGroupViewResourceInner, bool) {
	if o == nil || IsNil(o.VolumeGroupView) {
		return nil, false
	}
	return o.VolumeGroupView, true
}

// HasVolumeGroupView returns a boolean if a field has been set.
func (o *VolumeViewObject) HasVolumeGroupView() bool {
	if o != nil && !IsNil(o.VolumeGroupView) {
		return true
	}

	return false
}

// SetVolumeGroupView gets a reference to the given []VolumeGroupViewResourceInner and assigns it to the VolumeGroupView field.
func (o *VolumeViewObject) SetVolumeGroupView(v []VolumeGroupViewResourceInner) {
	o.VolumeGroupView = v
}

// GetVolumeView returns the VolumeView field value if set, zero value otherwise.
func (o *VolumeViewObject) GetVolumeView() []VolumeViewResourceInner {
	if o == nil || IsNil(o.VolumeView) {
		var ret []VolumeViewResourceInner
		return ret
	}
	return o.VolumeView
}

// GetVolumeViewOk returns a tuple with the VolumeView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeViewObject) GetVolumeViewOk() ([]VolumeViewResourceInner, bool) {
	if o == nil || IsNil(o.VolumeView) {
		return nil, false
	}
	return o.VolumeView, true
}

// HasVolumeView returns a boolean if a field has been set.
func (o *VolumeViewObject) HasVolumeView() bool {
	if o != nil && !IsNil(o.VolumeView) {
		return true
	}

	return false
}

// SetVolumeView gets a reference to the given []VolumeViewResourceInner and assigns it to the VolumeView field.
func (o *VolumeViewObject) SetVolumeView(v []VolumeViewResourceInner) {
	o.VolumeView = v
}

func (o VolumeViewObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeViewObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.VolumeGroupView) {
		toSerialize["volume-group-view"] = o.VolumeGroupView
	}
	if !IsNil(o.VolumeView) {
		toSerialize["volume-view"] = o.VolumeView
	}
	return toSerialize, nil
}

type NullableVolumeViewObject struct {
	value *VolumeViewObject
	isSet bool
}

func (v NullableVolumeViewObject) Get() *VolumeViewObject {
	return v.value
}

func (v *NullableVolumeViewObject) Set(val *VolumeViewObject) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeViewObject) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeViewObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeViewObject(val *VolumeViewObject) *NullableVolumeViewObject {
	return &NullableVolumeViewObject{value: val, isSet: true}
}

func (v NullableVolumeViewObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeViewObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
