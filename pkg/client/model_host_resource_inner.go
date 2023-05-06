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

// checks if the HostResourceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostResourceInner{}

// HostResourceInner struct for HostResourceInner
type HostResourceInner struct {
	ObjectName *string `json:"object-name,omitempty"`
	Meta       *string `json:"meta,omitempty"`
	DurableId  *string `json:"durable-id,omitempty"`
	// Durable ID of a Management Group
	GroupKey  *string `json:"group-key,omitempty"`
	HostGroup *string `json:"host-group,omitempty"`
	// Number of members
	MemberCount  *int32                   `json:"member-count,omitempty"`
	Name         *string                  `json:"name,omitempty"`
	SerialNumber *string                  `json:"serial-number,omitempty"`
	Initiator    []InitiatorResourceInner `json:"initiator,omitempty"`
}

// NewHostResourceInner instantiates a new HostResourceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostResourceInner() *HostResourceInner {
	this := HostResourceInner{}
	return &this
}

// NewHostResourceInnerWithDefaults instantiates a new HostResourceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostResourceInnerWithDefaults() *HostResourceInner {
	this := HostResourceInner{}
	return &this
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *HostResourceInner) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResourceInner) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *HostResourceInner) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *HostResourceInner) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *HostResourceInner) GetMeta() string {
	if o == nil || IsNil(o.Meta) {
		var ret string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResourceInner) GetMetaOk() (*string, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *HostResourceInner) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given string and assigns it to the Meta field.
func (o *HostResourceInner) SetMeta(v string) {
	o.Meta = &v
}

// GetDurableId returns the DurableId field value if set, zero value otherwise.
func (o *HostResourceInner) GetDurableId() string {
	if o == nil || IsNil(o.DurableId) {
		var ret string
		return ret
	}
	return *o.DurableId
}

// GetDurableIdOk returns a tuple with the DurableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResourceInner) GetDurableIdOk() (*string, bool) {
	if o == nil || IsNil(o.DurableId) {
		return nil, false
	}
	return o.DurableId, true
}

// HasDurableId returns a boolean if a field has been set.
func (o *HostResourceInner) HasDurableId() bool {
	if o != nil && !IsNil(o.DurableId) {
		return true
	}

	return false
}

// SetDurableId gets a reference to the given string and assigns it to the DurableId field.
func (o *HostResourceInner) SetDurableId(v string) {
	o.DurableId = &v
}

// GetGroupKey returns the GroupKey field value if set, zero value otherwise.
func (o *HostResourceInner) GetGroupKey() string {
	if o == nil || IsNil(o.GroupKey) {
		var ret string
		return ret
	}
	return *o.GroupKey
}

// GetGroupKeyOk returns a tuple with the GroupKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResourceInner) GetGroupKeyOk() (*string, bool) {
	if o == nil || IsNil(o.GroupKey) {
		return nil, false
	}
	return o.GroupKey, true
}

// HasGroupKey returns a boolean if a field has been set.
func (o *HostResourceInner) HasGroupKey() bool {
	if o != nil && !IsNil(o.GroupKey) {
		return true
	}

	return false
}

// SetGroupKey gets a reference to the given string and assigns it to the GroupKey field.
func (o *HostResourceInner) SetGroupKey(v string) {
	o.GroupKey = &v
}

// GetHostGroup returns the HostGroup field value if set, zero value otherwise.
func (o *HostResourceInner) GetHostGroup() string {
	if o == nil || IsNil(o.HostGroup) {
		var ret string
		return ret
	}
	return *o.HostGroup
}

// GetHostGroupOk returns a tuple with the HostGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResourceInner) GetHostGroupOk() (*string, bool) {
	if o == nil || IsNil(o.HostGroup) {
		return nil, false
	}
	return o.HostGroup, true
}

// HasHostGroup returns a boolean if a field has been set.
func (o *HostResourceInner) HasHostGroup() bool {
	if o != nil && !IsNil(o.HostGroup) {
		return true
	}

	return false
}

// SetHostGroup gets a reference to the given string and assigns it to the HostGroup field.
func (o *HostResourceInner) SetHostGroup(v string) {
	o.HostGroup = &v
}

// GetMemberCount returns the MemberCount field value if set, zero value otherwise.
func (o *HostResourceInner) GetMemberCount() int32 {
	if o == nil || IsNil(o.MemberCount) {
		var ret int32
		return ret
	}
	return *o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResourceInner) GetMemberCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MemberCount) {
		return nil, false
	}
	return o.MemberCount, true
}

// HasMemberCount returns a boolean if a field has been set.
func (o *HostResourceInner) HasMemberCount() bool {
	if o != nil && !IsNil(o.MemberCount) {
		return true
	}

	return false
}

// SetMemberCount gets a reference to the given int32 and assigns it to the MemberCount field.
func (o *HostResourceInner) SetMemberCount(v int32) {
	o.MemberCount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HostResourceInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResourceInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HostResourceInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HostResourceInner) SetName(v string) {
	o.Name = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *HostResourceInner) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResourceInner) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *HostResourceInner) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *HostResourceInner) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetInitiator returns the Initiator field value if set, zero value otherwise.
func (o *HostResourceInner) GetInitiator() []InitiatorResourceInner {
	if o == nil || IsNil(o.Initiator) {
		var ret []InitiatorResourceInner
		return ret
	}
	return o.Initiator
}

// GetInitiatorOk returns a tuple with the Initiator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostResourceInner) GetInitiatorOk() ([]InitiatorResourceInner, bool) {
	if o == nil || IsNil(o.Initiator) {
		return nil, false
	}
	return o.Initiator, true
}

// HasInitiator returns a boolean if a field has been set.
func (o *HostResourceInner) HasInitiator() bool {
	if o != nil && !IsNil(o.Initiator) {
		return true
	}

	return false
}

// SetInitiator gets a reference to the given []InitiatorResourceInner and assigns it to the Initiator field.
func (o *HostResourceInner) SetInitiator(v []InitiatorResourceInner) {
	o.Initiator = v
}

func (o HostResourceInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostResourceInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectName) {
		toSerialize["object-name"] = o.ObjectName
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.DurableId) {
		toSerialize["durable-id"] = o.DurableId
	}
	if !IsNil(o.GroupKey) {
		toSerialize["group-key"] = o.GroupKey
	}
	if !IsNil(o.HostGroup) {
		toSerialize["host-group"] = o.HostGroup
	}
	if !IsNil(o.MemberCount) {
		toSerialize["member-count"] = o.MemberCount
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serial-number"] = o.SerialNumber
	}
	if !IsNil(o.Initiator) {
		toSerialize["initiator"] = o.Initiator
	}
	return toSerialize, nil
}

type NullableHostResourceInner struct {
	value *HostResourceInner
	isSet bool
}

func (v NullableHostResourceInner) Get() *HostResourceInner {
	return v.value
}

func (v *NullableHostResourceInner) Set(val *HostResourceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableHostResourceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableHostResourceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostResourceInner(val *HostResourceInner) *NullableHostResourceInner {
	return &NullableHostResourceInner{value: val, isSet: true}
}

func (v NullableHostResourceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostResourceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
