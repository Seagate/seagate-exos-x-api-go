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

// checks if the CacheSettingsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CacheSettingsObject{}

// CacheSettingsObject struct for CacheSettingsObject
type CacheSettingsObject struct {
	Status        []StatusResourceInner        `json:"status,omitempty"`
	CacheSettings []CacheSettingsResourceInner `json:"cache-settings,omitempty"`
}

// NewCacheSettingsObject instantiates a new CacheSettingsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCacheSettingsObject() *CacheSettingsObject {
	this := CacheSettingsObject{}
	return &this
}

// NewCacheSettingsObjectWithDefaults instantiates a new CacheSettingsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCacheSettingsObjectWithDefaults() *CacheSettingsObject {
	this := CacheSettingsObject{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CacheSettingsObject) GetStatus() []StatusResourceInner {
	if o == nil || IsNil(o.Status) {
		var ret []StatusResourceInner
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheSettingsObject) GetStatusOk() ([]StatusResourceInner, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CacheSettingsObject) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []StatusResourceInner and assigns it to the Status field.
func (o *CacheSettingsObject) SetStatus(v []StatusResourceInner) {
	o.Status = v
}

// GetCacheSettings returns the CacheSettings field value if set, zero value otherwise.
func (o *CacheSettingsObject) GetCacheSettings() []CacheSettingsResourceInner {
	if o == nil || IsNil(o.CacheSettings) {
		var ret []CacheSettingsResourceInner
		return ret
	}
	return o.CacheSettings
}

// GetCacheSettingsOk returns a tuple with the CacheSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheSettingsObject) GetCacheSettingsOk() ([]CacheSettingsResourceInner, bool) {
	if o == nil || IsNil(o.CacheSettings) {
		return nil, false
	}
	return o.CacheSettings, true
}

// HasCacheSettings returns a boolean if a field has been set.
func (o *CacheSettingsObject) HasCacheSettings() bool {
	if o != nil && !IsNil(o.CacheSettings) {
		return true
	}

	return false
}

// SetCacheSettings gets a reference to the given []CacheSettingsResourceInner and assigns it to the CacheSettings field.
func (o *CacheSettingsObject) SetCacheSettings(v []CacheSettingsResourceInner) {
	o.CacheSettings = v
}

func (o CacheSettingsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CacheSettingsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CacheSettings) {
		toSerialize["cache-settings"] = o.CacheSettings
	}
	return toSerialize, nil
}

type NullableCacheSettingsObject struct {
	value *CacheSettingsObject
	isSet bool
}

func (v NullableCacheSettingsObject) Get() *CacheSettingsObject {
	return v.value
}

func (v *NullableCacheSettingsObject) Set(val *CacheSettingsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCacheSettingsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCacheSettingsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCacheSettingsObject(val *CacheSettingsObject) *NullableCacheSettingsObject {
	return &NullableCacheSettingsObject{value: val, isSet: true}
}

func (v NullableCacheSettingsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCacheSettingsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
