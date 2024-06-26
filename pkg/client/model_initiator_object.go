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

// checks if the InitiatorObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitiatorObject{}

// InitiatorObject struct for InitiatorObject
type InitiatorObject struct {
	Status    []StatusResourceInner    `json:"status,omitempty"`
	Initiator []InitiatorResourceInner `json:"initiator,omitempty"`
}

// NewInitiatorObject instantiates a new InitiatorObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitiatorObject() *InitiatorObject {
	this := InitiatorObject{}
	return &this
}

// NewInitiatorObjectWithDefaults instantiates a new InitiatorObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitiatorObjectWithDefaults() *InitiatorObject {
	this := InitiatorObject{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InitiatorObject) GetStatus() []StatusResourceInner {
	if o == nil || IsNil(o.Status) {
		var ret []StatusResourceInner
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitiatorObject) GetStatusOk() ([]StatusResourceInner, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InitiatorObject) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []StatusResourceInner and assigns it to the Status field.
func (o *InitiatorObject) SetStatus(v []StatusResourceInner) {
	o.Status = v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *InitiatorObject) GetInitiator() []InitiatorResourceInner {
	if o == nil || IsNil(o.Initiator) {
		var ret []InitiatorResourceInner
		return ret
	}
	return o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitiatorObject) GetInitiatorOk() ([]InitiatorResourceInner, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *InitiatorObject) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given []InitiatorResourceInner and assigns it to the Initiator field.
func (o *InitiatorObject) SetInitiator(v []InitiatorResourceInner) {
	o.Initiator = v
}

func (o InitiatorObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InitiatorObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	return toSerialize, nil
}

type NullableInitiatorObject struct {
	value *InitiatorObject
	isSet bool
}

func (v NullableInitiatorObject) Get() *InitiatorObject {
	return v.value
}

func (v *NullableInitiatorObject) Set(val *InitiatorObject) {
	v.value = val
	v.isSet = true
}

func (v NullableInitiatorObject) IsSet() bool {
	return v.isSet
}

func (v *NullableInitiatorObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitiatorObject(val *InitiatorObject) *NullableInitiatorObject {
	return &NullableInitiatorObject{value: val, isSet: true}
}

func (v NullableInitiatorObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInitiatorObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
