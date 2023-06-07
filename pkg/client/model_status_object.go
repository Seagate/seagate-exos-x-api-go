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

// checks if the StatusObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusObject{}

// StatusObject struct for StatusObject
type StatusObject struct {
	Status []StatusResourceInner `json:"status,omitempty"`
}

// NewStatusObject instantiates a new StatusObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusObject() *StatusObject {
	this := StatusObject{}
	return &this
}

// NewStatusObjectWithDefaults instantiates a new StatusObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusObjectWithDefaults() *StatusObject {
	this := StatusObject{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StatusObject) GetStatus() []StatusResourceInner {
	if o == nil || IsNil(o.Status) {
		var ret []StatusResourceInner
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusObject) GetStatusOk() ([]StatusResourceInner, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StatusObject) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []StatusResourceInner and assigns it to the Status field.
func (o *StatusObject) SetStatus(v []StatusResourceInner) {
	o.Status = v
}

func (o StatusObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableStatusObject struct {
	value *StatusObject
	isSet bool
}

func (v NullableStatusObject) Get() *StatusObject {
	return v.value
}

func (v *NullableStatusObject) Set(val *StatusObject) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusObject) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusObject(val *StatusObject) *NullableStatusObject {
	return &NullableStatusObject{value: val, isSet: true}
}

func (v NullableStatusObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
