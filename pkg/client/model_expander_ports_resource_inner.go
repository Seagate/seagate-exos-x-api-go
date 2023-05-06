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

// checks if the ExpanderPortsResourceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpanderPortsResourceInner{}

// ExpanderPortsResourceInner struct for ExpanderPortsResourceInner
type ExpanderPortsResourceInner struct {
	ObjectName           *string `json:"object-name,omitempty"`
	Meta                 *string `json:"meta,omitempty"`
	Controller           *string `json:"controller,omitempty"`
	ControllerNumeric    *int32  `json:"controller-numeric,omitempty"`
	DurableId            *string `json:"durable-id,omitempty"`
	EnclosureId          *int32  `json:"enclosure-id,omitempty"`
	Health               *string `json:"health,omitempty"`
	HealthNumeric        *int32  `json:"health-numeric,omitempty"`
	HealthReason         *string `json:"health-reason,omitempty"`
	HealthRecommendation *string `json:"health-recommendation,omitempty"`
	Name                 *string `json:"name,omitempty"`
	// Port number for this SAS port
	SasPortIndex *int32 `json:"sas-port-index,omitempty"`
	// The type of SAS port
	SasPortType *string `json:"sas-port-type,omitempty"`
	// The type of SAS port( In numeric form )
	SasPortTypeNumeric *int32  `json:"sas-port-type-numeric,omitempty"`
	Status             *string `json:"status,omitempty"`
	StatusNumeric      *int32  `json:"status-numeric,omitempty"`
}

// NewExpanderPortsResourceInner instantiates a new ExpanderPortsResourceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpanderPortsResourceInner() *ExpanderPortsResourceInner {
	this := ExpanderPortsResourceInner{}
	return &this
}

// NewExpanderPortsResourceInnerWithDefaults instantiates a new ExpanderPortsResourceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpanderPortsResourceInnerWithDefaults() *ExpanderPortsResourceInner {
	this := ExpanderPortsResourceInner{}
	return &this
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *ExpanderPortsResourceInner) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetMeta() string {
	if o == nil || IsNil(o.Meta) {
		var ret string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetMetaOk() (*string, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given string and assigns it to the Meta field.
func (o *ExpanderPortsResourceInner) SetMeta(v string) {
	o.Meta = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetController() string {
	if o == nil || IsNil(o.Controller) {
		var ret string
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetControllerOk() (*string, bool) {
	if o == nil || IsNil(o.Controller) {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasController() bool {
	if o != nil && !IsNil(o.Controller) {
		return true
	}

	return false
}

// SetController gets a reference to the given string and assigns it to the Controller field.
func (o *ExpanderPortsResourceInner) SetController(v string) {
	o.Controller = &v
}

// GetControllerNumeric returns the ControllerNumeric field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetControllerNumeric() int32 {
	if o == nil || IsNil(o.ControllerNumeric) {
		var ret int32
		return ret
	}
	return *o.ControllerNumeric
}

// GetControllerNumericOk returns a tuple with the ControllerNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetControllerNumericOk() (*int32, bool) {
	if o == nil || IsNil(o.ControllerNumeric) {
		return nil, false
	}
	return o.ControllerNumeric, true
}

// HasControllerNumeric returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasControllerNumeric() bool {
	if o != nil && !IsNil(o.ControllerNumeric) {
		return true
	}

	return false
}

// SetControllerNumeric gets a reference to the given int32 and assigns it to the ControllerNumeric field.
func (o *ExpanderPortsResourceInner) SetControllerNumeric(v int32) {
	o.ControllerNumeric = &v
}

// GetDurableId returns the DurableId field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetDurableId() string {
	if o == nil || IsNil(o.DurableId) {
		var ret string
		return ret
	}
	return *o.DurableId
}

// GetDurableIdOk returns a tuple with the DurableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetDurableIdOk() (*string, bool) {
	if o == nil || IsNil(o.DurableId) {
		return nil, false
	}
	return o.DurableId, true
}

// HasDurableId returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasDurableId() bool {
	if o != nil && !IsNil(o.DurableId) {
		return true
	}

	return false
}

// SetDurableId gets a reference to the given string and assigns it to the DurableId field.
func (o *ExpanderPortsResourceInner) SetDurableId(v string) {
	o.DurableId = &v
}

// GetEnclosureId returns the EnclosureId field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetEnclosureId() int32 {
	if o == nil || IsNil(o.EnclosureId) {
		var ret int32
		return ret
	}
	return *o.EnclosureId
}

// GetEnclosureIdOk returns a tuple with the EnclosureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetEnclosureIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EnclosureId) {
		return nil, false
	}
	return o.EnclosureId, true
}

// HasEnclosureId returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasEnclosureId() bool {
	if o != nil && !IsNil(o.EnclosureId) {
		return true
	}

	return false
}

// SetEnclosureId gets a reference to the given int32 and assigns it to the EnclosureId field.
func (o *ExpanderPortsResourceInner) SetEnclosureId(v int32) {
	o.EnclosureId = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetHealth() string {
	if o == nil || IsNil(o.Health) {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetHealthOk() (*string, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasHealth() bool {
	if o != nil && !IsNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *ExpanderPortsResourceInner) SetHealth(v string) {
	o.Health = &v
}

// GetHealthNumeric returns the HealthNumeric field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetHealthNumeric() int32 {
	if o == nil || IsNil(o.HealthNumeric) {
		var ret int32
		return ret
	}
	return *o.HealthNumeric
}

// GetHealthNumericOk returns a tuple with the HealthNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetHealthNumericOk() (*int32, bool) {
	if o == nil || IsNil(o.HealthNumeric) {
		return nil, false
	}
	return o.HealthNumeric, true
}

// HasHealthNumeric returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasHealthNumeric() bool {
	if o != nil && !IsNil(o.HealthNumeric) {
		return true
	}

	return false
}

// SetHealthNumeric gets a reference to the given int32 and assigns it to the HealthNumeric field.
func (o *ExpanderPortsResourceInner) SetHealthNumeric(v int32) {
	o.HealthNumeric = &v
}

// GetHealthReason returns the HealthReason field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetHealthReason() string {
	if o == nil || IsNil(o.HealthReason) {
		var ret string
		return ret
	}
	return *o.HealthReason
}

// GetHealthReasonOk returns a tuple with the HealthReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetHealthReasonOk() (*string, bool) {
	if o == nil || IsNil(o.HealthReason) {
		return nil, false
	}
	return o.HealthReason, true
}

// HasHealthReason returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasHealthReason() bool {
	if o != nil && !IsNil(o.HealthReason) {
		return true
	}

	return false
}

// SetHealthReason gets a reference to the given string and assigns it to the HealthReason field.
func (o *ExpanderPortsResourceInner) SetHealthReason(v string) {
	o.HealthReason = &v
}

// GetHealthRecommendation returns the HealthRecommendation field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetHealthRecommendation() string {
	if o == nil || IsNil(o.HealthRecommendation) {
		var ret string
		return ret
	}
	return *o.HealthRecommendation
}

// GetHealthRecommendationOk returns a tuple with the HealthRecommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetHealthRecommendationOk() (*string, bool) {
	if o == nil || IsNil(o.HealthRecommendation) {
		return nil, false
	}
	return o.HealthRecommendation, true
}

// HasHealthRecommendation returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasHealthRecommendation() bool {
	if o != nil && !IsNil(o.HealthRecommendation) {
		return true
	}

	return false
}

// SetHealthRecommendation gets a reference to the given string and assigns it to the HealthRecommendation field.
func (o *ExpanderPortsResourceInner) SetHealthRecommendation(v string) {
	o.HealthRecommendation = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExpanderPortsResourceInner) SetName(v string) {
	o.Name = &v
}

// GetSasPortIndex returns the SasPortIndex field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetSasPortIndex() int32 {
	if o == nil || IsNil(o.SasPortIndex) {
		var ret int32
		return ret
	}
	return *o.SasPortIndex
}

// GetSasPortIndexOk returns a tuple with the SasPortIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetSasPortIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.SasPortIndex) {
		return nil, false
	}
	return o.SasPortIndex, true
}

// HasSasPortIndex returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasSasPortIndex() bool {
	if o != nil && !IsNil(o.SasPortIndex) {
		return true
	}

	return false
}

// SetSasPortIndex gets a reference to the given int32 and assigns it to the SasPortIndex field.
func (o *ExpanderPortsResourceInner) SetSasPortIndex(v int32) {
	o.SasPortIndex = &v
}

// GetSasPortType returns the SasPortType field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetSasPortType() string {
	if o == nil || IsNil(o.SasPortType) {
		var ret string
		return ret
	}
	return *o.SasPortType
}

// GetSasPortTypeOk returns a tuple with the SasPortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetSasPortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SasPortType) {
		return nil, false
	}
	return o.SasPortType, true
}

// HasSasPortType returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasSasPortType() bool {
	if o != nil && !IsNil(o.SasPortType) {
		return true
	}

	return false
}

// SetSasPortType gets a reference to the given string and assigns it to the SasPortType field.
func (o *ExpanderPortsResourceInner) SetSasPortType(v string) {
	o.SasPortType = &v
}

// GetSasPortTypeNumeric returns the SasPortTypeNumeric field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetSasPortTypeNumeric() int32 {
	if o == nil || IsNil(o.SasPortTypeNumeric) {
		var ret int32
		return ret
	}
	return *o.SasPortTypeNumeric
}

// GetSasPortTypeNumericOk returns a tuple with the SasPortTypeNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetSasPortTypeNumericOk() (*int32, bool) {
	if o == nil || IsNil(o.SasPortTypeNumeric) {
		return nil, false
	}
	return o.SasPortTypeNumeric, true
}

// HasSasPortTypeNumeric returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasSasPortTypeNumeric() bool {
	if o != nil && !IsNil(o.SasPortTypeNumeric) {
		return true
	}

	return false
}

// SetSasPortTypeNumeric gets a reference to the given int32 and assigns it to the SasPortTypeNumeric field.
func (o *ExpanderPortsResourceInner) SetSasPortTypeNumeric(v int32) {
	o.SasPortTypeNumeric = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ExpanderPortsResourceInner) SetStatus(v string) {
	o.Status = &v
}

// GetStatusNumeric returns the StatusNumeric field value if set, zero value otherwise.
func (o *ExpanderPortsResourceInner) GetStatusNumeric() int32 {
	if o == nil || IsNil(o.StatusNumeric) {
		var ret int32
		return ret
	}
	return *o.StatusNumeric
}

// GetStatusNumericOk returns a tuple with the StatusNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpanderPortsResourceInner) GetStatusNumericOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusNumeric) {
		return nil, false
	}
	return o.StatusNumeric, true
}

// HasStatusNumeric returns a boolean if a field has been set.
func (o *ExpanderPortsResourceInner) HasStatusNumeric() bool {
	if o != nil && !IsNil(o.StatusNumeric) {
		return true
	}

	return false
}

// SetStatusNumeric gets a reference to the given int32 and assigns it to the StatusNumeric field.
func (o *ExpanderPortsResourceInner) SetStatusNumeric(v int32) {
	o.StatusNumeric = &v
}

func (o ExpanderPortsResourceInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpanderPortsResourceInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectName) {
		toSerialize["object-name"] = o.ObjectName
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Controller) {
		toSerialize["controller"] = o.Controller
	}
	if !IsNil(o.ControllerNumeric) {
		toSerialize["controller-numeric"] = o.ControllerNumeric
	}
	if !IsNil(o.DurableId) {
		toSerialize["durable-id"] = o.DurableId
	}
	if !IsNil(o.EnclosureId) {
		toSerialize["enclosure-id"] = o.EnclosureId
	}
	if !IsNil(o.Health) {
		toSerialize["health"] = o.Health
	}
	if !IsNil(o.HealthNumeric) {
		toSerialize["health-numeric"] = o.HealthNumeric
	}
	if !IsNil(o.HealthReason) {
		toSerialize["health-reason"] = o.HealthReason
	}
	if !IsNil(o.HealthRecommendation) {
		toSerialize["health-recommendation"] = o.HealthRecommendation
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SasPortIndex) {
		toSerialize["sas-port-index"] = o.SasPortIndex
	}
	if !IsNil(o.SasPortType) {
		toSerialize["sas-port-type"] = o.SasPortType
	}
	if !IsNil(o.SasPortTypeNumeric) {
		toSerialize["sas-port-type-numeric"] = o.SasPortTypeNumeric
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusNumeric) {
		toSerialize["status-numeric"] = o.StatusNumeric
	}
	return toSerialize, nil
}

type NullableExpanderPortsResourceInner struct {
	value *ExpanderPortsResourceInner
	isSet bool
}

func (v NullableExpanderPortsResourceInner) Get() *ExpanderPortsResourceInner {
	return v.value
}

func (v *NullableExpanderPortsResourceInner) Set(val *ExpanderPortsResourceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExpanderPortsResourceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExpanderPortsResourceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpanderPortsResourceInner(val *ExpanderPortsResourceInner) *NullableExpanderPortsResourceInner {
	return &NullableExpanderPortsResourceInner{value: val, isSet: true}
}

func (v NullableExpanderPortsResourceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpanderPortsResourceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}