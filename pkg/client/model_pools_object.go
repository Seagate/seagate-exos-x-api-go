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

// checks if the PoolsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolsObject{}

// PoolsObject struct for PoolsObject
type PoolsObject struct {
	Status []StatusResourceInner `json:"status,omitempty"`
	Pools  []PoolsResourceInner  `json:"pools,omitempty"`
}

// NewPoolsObject instantiates a new PoolsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolsObject() *PoolsObject {
	this := PoolsObject{}
	return &this
}

// NewPoolsObjectWithDefaults instantiates a new PoolsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolsObjectWithDefaults() *PoolsObject {
	this := PoolsObject{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PoolsObject) GetStatus() []StatusResourceInner {
	if o == nil || IsNil(o.Status) {
		var ret []StatusResourceInner
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsObject) GetStatusOk() ([]StatusResourceInner, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PoolsObject) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []StatusResourceInner and assigns it to the Status field.
func (o *PoolsObject) SetStatus(v []StatusResourceInner) {
	o.Status = v
}

// GetPools returns the Pools field value if set, zero value otherwise.
func (o *PoolsObject) GetPools() []PoolsResourceInner {
	if o == nil || IsNil(o.Pools) {
		var ret []PoolsResourceInner
		return ret
	}
	return o.Pools
}

// GetPoolsOk returns a tuple with the Pools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolsObject) GetPoolsOk() ([]PoolsResourceInner, bool) {
	if o == nil || IsNil(o.Pools) {
		return nil, false
	}
	return o.Pools, true
}

// HasPools returns a boolean if a field has been set.
func (o *PoolsObject) HasPools() bool {
	if o != nil && !IsNil(o.Pools) {
		return true
	}

	return false
}

// SetPools gets a reference to the given []PoolsResourceInner and assigns it to the Pools field.
func (o *PoolsObject) SetPools(v []PoolsResourceInner) {
	o.Pools = v
}

func (o PoolsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Pools) {
		toSerialize["pools"] = o.Pools
	}
	return toSerialize, nil
}

type NullablePoolsObject struct {
	value *PoolsObject
	isSet bool
}

func (v NullablePoolsObject) Get() *PoolsObject {
	return v.value
}

func (v *NullablePoolsObject) Set(val *PoolsObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolsObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolsObject(val *PoolsObject) *NullablePoolsObject {
	return &NullablePoolsObject{value: val, isSet: true}
}

func (v NullablePoolsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
