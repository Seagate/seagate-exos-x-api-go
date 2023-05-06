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

// checks if the FanResourceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FanResourceInner{}

// FanResourceInner struct for FanResourceInner
type FanResourceInner struct {
	ObjectName *string `json:"object-name,omitempty"`
	Meta       *string `json:"meta,omitempty"`
	DurableId  *string `json:"durable-id,omitempty"`
	// Extended status (bits)
	ExtendedStatus *string `json:"extended-status,omitempty"`
	// Firmware version of the FRU
	FwRevision           *string `json:"fw-revision,omitempty"`
	Health               *string `json:"health,omitempty"`
	HealthNumeric        *int32  `json:"health-numeric,omitempty"`
	HealthReason         *string `json:"health-reason,omitempty"`
	HealthRecommendation *string `json:"health-recommendation,omitempty"`
	// Hardware version of the FRU
	HwRevision *string `json:"hw-revision,omitempty"`
	Location   *string `json:"location,omitempty"`
	// Indicates whether the locator LED is on
	LocatorLed *string `json:"locator-led,omitempty"`
	// Indicates whether the locator LED is on( In numeric form )
	LocatorLedNumeric *int32  `json:"locator-led-numeric,omitempty"`
	Name              *string `json:"name,omitempty"`
	PartNumber        *string `json:"part-number,omitempty"`
	// Position of the component in the enclosure
	Position *string `json:"position,omitempty"`
	// Position of the component in the enclosure( In numeric form )
	PositionNumeric *int32  `json:"position-numeric,omitempty"`
	SerialNumber    *string `json:"serial-number,omitempty"`
	Speed           *int32  `json:"speed,omitempty"`
	Status          *string `json:"status,omitempty"`
	StatusNumeric   *int32  `json:"status-numeric,omitempty"`
	// SES Common status
	StatusSes *string `json:"status-ses,omitempty"`
	// SES Common status( In numeric form )
	StatusSesNumeric *int32 `json:"status-ses-numeric,omitempty"`
	// The resource URL
	Url *string `json:"url,omitempty"`
}

// NewFanResourceInner instantiates a new FanResourceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFanResourceInner() *FanResourceInner {
	this := FanResourceInner{}
	return &this
}

// NewFanResourceInnerWithDefaults instantiates a new FanResourceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFanResourceInnerWithDefaults() *FanResourceInner {
	this := FanResourceInner{}
	return &this
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *FanResourceInner) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *FanResourceInner) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *FanResourceInner) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FanResourceInner) GetMeta() string {
	if o == nil || IsNil(o.Meta) {
		var ret string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetMetaOk() (*string, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FanResourceInner) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given string and assigns it to the Meta field.
func (o *FanResourceInner) SetMeta(v string) {
	o.Meta = &v
}

// GetDurableId returns the DurableId field value if set, zero value otherwise.
func (o *FanResourceInner) GetDurableId() string {
	if o == nil || IsNil(o.DurableId) {
		var ret string
		return ret
	}
	return *o.DurableId
}

// GetDurableIdOk returns a tuple with the DurableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetDurableIdOk() (*string, bool) {
	if o == nil || IsNil(o.DurableId) {
		return nil, false
	}
	return o.DurableId, true
}

// HasDurableId returns a boolean if a field has been set.
func (o *FanResourceInner) HasDurableId() bool {
	if o != nil && !IsNil(o.DurableId) {
		return true
	}

	return false
}

// SetDurableId gets a reference to the given string and assigns it to the DurableId field.
func (o *FanResourceInner) SetDurableId(v string) {
	o.DurableId = &v
}

// GetExtendedStatus returns the ExtendedStatus field value if set, zero value otherwise.
func (o *FanResourceInner) GetExtendedStatus() string {
	if o == nil || IsNil(o.ExtendedStatus) {
		var ret string
		return ret
	}
	return *o.ExtendedStatus
}

// GetExtendedStatusOk returns a tuple with the ExtendedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetExtendedStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ExtendedStatus) {
		return nil, false
	}
	return o.ExtendedStatus, true
}

// HasExtendedStatus returns a boolean if a field has been set.
func (o *FanResourceInner) HasExtendedStatus() bool {
	if o != nil && !IsNil(o.ExtendedStatus) {
		return true
	}

	return false
}

// SetExtendedStatus gets a reference to the given string and assigns it to the ExtendedStatus field.
func (o *FanResourceInner) SetExtendedStatus(v string) {
	o.ExtendedStatus = &v
}

// GetFwRevision returns the FwRevision field value if set, zero value otherwise.
func (o *FanResourceInner) GetFwRevision() string {
	if o == nil || IsNil(o.FwRevision) {
		var ret string
		return ret
	}
	return *o.FwRevision
}

// GetFwRevisionOk returns a tuple with the FwRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetFwRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.FwRevision) {
		return nil, false
	}
	return o.FwRevision, true
}

// HasFwRevision returns a boolean if a field has been set.
func (o *FanResourceInner) HasFwRevision() bool {
	if o != nil && !IsNil(o.FwRevision) {
		return true
	}

	return false
}

// SetFwRevision gets a reference to the given string and assigns it to the FwRevision field.
func (o *FanResourceInner) SetFwRevision(v string) {
	o.FwRevision = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *FanResourceInner) GetHealth() string {
	if o == nil || IsNil(o.Health) {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetHealthOk() (*string, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *FanResourceInner) HasHealth() bool {
	if o != nil && !IsNil(o.Health) {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *FanResourceInner) SetHealth(v string) {
	o.Health = &v
}

// GetHealthNumeric returns the HealthNumeric field value if set, zero value otherwise.
func (o *FanResourceInner) GetHealthNumeric() int32 {
	if o == nil || IsNil(o.HealthNumeric) {
		var ret int32
		return ret
	}
	return *o.HealthNumeric
}

// GetHealthNumericOk returns a tuple with the HealthNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetHealthNumericOk() (*int32, bool) {
	if o == nil || IsNil(o.HealthNumeric) {
		return nil, false
	}
	return o.HealthNumeric, true
}

// HasHealthNumeric returns a boolean if a field has been set.
func (o *FanResourceInner) HasHealthNumeric() bool {
	if o != nil && !IsNil(o.HealthNumeric) {
		return true
	}

	return false
}

// SetHealthNumeric gets a reference to the given int32 and assigns it to the HealthNumeric field.
func (o *FanResourceInner) SetHealthNumeric(v int32) {
	o.HealthNumeric = &v
}

// GetHealthReason returns the HealthReason field value if set, zero value otherwise.
func (o *FanResourceInner) GetHealthReason() string {
	if o == nil || IsNil(o.HealthReason) {
		var ret string
		return ret
	}
	return *o.HealthReason
}

// GetHealthReasonOk returns a tuple with the HealthReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetHealthReasonOk() (*string, bool) {
	if o == nil || IsNil(o.HealthReason) {
		return nil, false
	}
	return o.HealthReason, true
}

// HasHealthReason returns a boolean if a field has been set.
func (o *FanResourceInner) HasHealthReason() bool {
	if o != nil && !IsNil(o.HealthReason) {
		return true
	}

	return false
}

// SetHealthReason gets a reference to the given string and assigns it to the HealthReason field.
func (o *FanResourceInner) SetHealthReason(v string) {
	o.HealthReason = &v
}

// GetHealthRecommendation returns the HealthRecommendation field value if set, zero value otherwise.
func (o *FanResourceInner) GetHealthRecommendation() string {
	if o == nil || IsNil(o.HealthRecommendation) {
		var ret string
		return ret
	}
	return *o.HealthRecommendation
}

// GetHealthRecommendationOk returns a tuple with the HealthRecommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetHealthRecommendationOk() (*string, bool) {
	if o == nil || IsNil(o.HealthRecommendation) {
		return nil, false
	}
	return o.HealthRecommendation, true
}

// HasHealthRecommendation returns a boolean if a field has been set.
func (o *FanResourceInner) HasHealthRecommendation() bool {
	if o != nil && !IsNil(o.HealthRecommendation) {
		return true
	}

	return false
}

// SetHealthRecommendation gets a reference to the given string and assigns it to the HealthRecommendation field.
func (o *FanResourceInner) SetHealthRecommendation(v string) {
	o.HealthRecommendation = &v
}

// GetHwRevision returns the HwRevision field value if set, zero value otherwise.
func (o *FanResourceInner) GetHwRevision() string {
	if o == nil || IsNil(o.HwRevision) {
		var ret string
		return ret
	}
	return *o.HwRevision
}

// GetHwRevisionOk returns a tuple with the HwRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetHwRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.HwRevision) {
		return nil, false
	}
	return o.HwRevision, true
}

// HasHwRevision returns a boolean if a field has been set.
func (o *FanResourceInner) HasHwRevision() bool {
	if o != nil && !IsNil(o.HwRevision) {
		return true
	}

	return false
}

// SetHwRevision gets a reference to the given string and assigns it to the HwRevision field.
func (o *FanResourceInner) SetHwRevision(v string) {
	o.HwRevision = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *FanResourceInner) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *FanResourceInner) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *FanResourceInner) SetLocation(v string) {
	o.Location = &v
}

// GetLocatorLed returns the LocatorLed field value if set, zero value otherwise.
func (o *FanResourceInner) GetLocatorLed() string {
	if o == nil || IsNil(o.LocatorLed) {
		var ret string
		return ret
	}
	return *o.LocatorLed
}

// GetLocatorLedOk returns a tuple with the LocatorLed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetLocatorLedOk() (*string, bool) {
	if o == nil || IsNil(o.LocatorLed) {
		return nil, false
	}
	return o.LocatorLed, true
}

// HasLocatorLed returns a boolean if a field has been set.
func (o *FanResourceInner) HasLocatorLed() bool {
	if o != nil && !IsNil(o.LocatorLed) {
		return true
	}

	return false
}

// SetLocatorLed gets a reference to the given string and assigns it to the LocatorLed field.
func (o *FanResourceInner) SetLocatorLed(v string) {
	o.LocatorLed = &v
}

// GetLocatorLedNumeric returns the LocatorLedNumeric field value if set, zero value otherwise.
func (o *FanResourceInner) GetLocatorLedNumeric() int32 {
	if o == nil || IsNil(o.LocatorLedNumeric) {
		var ret int32
		return ret
	}
	return *o.LocatorLedNumeric
}

// GetLocatorLedNumericOk returns a tuple with the LocatorLedNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetLocatorLedNumericOk() (*int32, bool) {
	if o == nil || IsNil(o.LocatorLedNumeric) {
		return nil, false
	}
	return o.LocatorLedNumeric, true
}

// HasLocatorLedNumeric returns a boolean if a field has been set.
func (o *FanResourceInner) HasLocatorLedNumeric() bool {
	if o != nil && !IsNil(o.LocatorLedNumeric) {
		return true
	}

	return false
}

// SetLocatorLedNumeric gets a reference to the given int32 and assigns it to the LocatorLedNumeric field.
func (o *FanResourceInner) SetLocatorLedNumeric(v int32) {
	o.LocatorLedNumeric = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FanResourceInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FanResourceInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FanResourceInner) SetName(v string) {
	o.Name = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *FanResourceInner) GetPartNumber() string {
	if o == nil || IsNil(o.PartNumber) {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartNumber) {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *FanResourceInner) HasPartNumber() bool {
	if o != nil && !IsNil(o.PartNumber) {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *FanResourceInner) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *FanResourceInner) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *FanResourceInner) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *FanResourceInner) SetPosition(v string) {
	o.Position = &v
}

// GetPositionNumeric returns the PositionNumeric field value if set, zero value otherwise.
func (o *FanResourceInner) GetPositionNumeric() int32 {
	if o == nil || IsNil(o.PositionNumeric) {
		var ret int32
		return ret
	}
	return *o.PositionNumeric
}

// GetPositionNumericOk returns a tuple with the PositionNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetPositionNumericOk() (*int32, bool) {
	if o == nil || IsNil(o.PositionNumeric) {
		return nil, false
	}
	return o.PositionNumeric, true
}

// HasPositionNumeric returns a boolean if a field has been set.
func (o *FanResourceInner) HasPositionNumeric() bool {
	if o != nil && !IsNil(o.PositionNumeric) {
		return true
	}

	return false
}

// SetPositionNumeric gets a reference to the given int32 and assigns it to the PositionNumeric field.
func (o *FanResourceInner) SetPositionNumeric(v int32) {
	o.PositionNumeric = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *FanResourceInner) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *FanResourceInner) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *FanResourceInner) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *FanResourceInner) GetSpeed() int32 {
	if o == nil || IsNil(o.Speed) {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *FanResourceInner) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *FanResourceInner) SetSpeed(v int32) {
	o.Speed = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FanResourceInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FanResourceInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FanResourceInner) SetStatus(v string) {
	o.Status = &v
}

// GetStatusNumeric returns the StatusNumeric field value if set, zero value otherwise.
func (o *FanResourceInner) GetStatusNumeric() int32 {
	if o == nil || IsNil(o.StatusNumeric) {
		var ret int32
		return ret
	}
	return *o.StatusNumeric
}

// GetStatusNumericOk returns a tuple with the StatusNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetStatusNumericOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusNumeric) {
		return nil, false
	}
	return o.StatusNumeric, true
}

// HasStatusNumeric returns a boolean if a field has been set.
func (o *FanResourceInner) HasStatusNumeric() bool {
	if o != nil && !IsNil(o.StatusNumeric) {
		return true
	}

	return false
}

// SetStatusNumeric gets a reference to the given int32 and assigns it to the StatusNumeric field.
func (o *FanResourceInner) SetStatusNumeric(v int32) {
	o.StatusNumeric = &v
}

// GetStatusSes returns the StatusSes field value if set, zero value otherwise.
func (o *FanResourceInner) GetStatusSes() string {
	if o == nil || IsNil(o.StatusSes) {
		var ret string
		return ret
	}
	return *o.StatusSes
}

// GetStatusSesOk returns a tuple with the StatusSes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetStatusSesOk() (*string, bool) {
	if o == nil || IsNil(o.StatusSes) {
		return nil, false
	}
	return o.StatusSes, true
}

// HasStatusSes returns a boolean if a field has been set.
func (o *FanResourceInner) HasStatusSes() bool {
	if o != nil && !IsNil(o.StatusSes) {
		return true
	}

	return false
}

// SetStatusSes gets a reference to the given string and assigns it to the StatusSes field.
func (o *FanResourceInner) SetStatusSes(v string) {
	o.StatusSes = &v
}

// GetStatusSesNumeric returns the StatusSesNumeric field value if set, zero value otherwise.
func (o *FanResourceInner) GetStatusSesNumeric() int32 {
	if o == nil || IsNil(o.StatusSesNumeric) {
		var ret int32
		return ret
	}
	return *o.StatusSesNumeric
}

// GetStatusSesNumericOk returns a tuple with the StatusSesNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetStatusSesNumericOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusSesNumeric) {
		return nil, false
	}
	return o.StatusSesNumeric, true
}

// HasStatusSesNumeric returns a boolean if a field has been set.
func (o *FanResourceInner) HasStatusSesNumeric() bool {
	if o != nil && !IsNil(o.StatusSesNumeric) {
		return true
	}

	return false
}

// SetStatusSesNumeric gets a reference to the given int32 and assigns it to the StatusSesNumeric field.
func (o *FanResourceInner) SetStatusSesNumeric(v int32) {
	o.StatusSesNumeric = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *FanResourceInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanResourceInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *FanResourceInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *FanResourceInner) SetUrl(v string) {
	o.Url = &v
}

func (o FanResourceInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FanResourceInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ExtendedStatus) {
		toSerialize["extended-status"] = o.ExtendedStatus
	}
	if !IsNil(o.FwRevision) {
		toSerialize["fw-revision"] = o.FwRevision
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
	if !IsNil(o.HwRevision) {
		toSerialize["hw-revision"] = o.HwRevision
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.LocatorLed) {
		toSerialize["locator-led"] = o.LocatorLed
	}
	if !IsNil(o.LocatorLedNumeric) {
		toSerialize["locator-led-numeric"] = o.LocatorLedNumeric
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PartNumber) {
		toSerialize["part-number"] = o.PartNumber
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.PositionNumeric) {
		toSerialize["position-numeric"] = o.PositionNumeric
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serial-number"] = o.SerialNumber
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusNumeric) {
		toSerialize["status-numeric"] = o.StatusNumeric
	}
	if !IsNil(o.StatusSes) {
		toSerialize["status-ses"] = o.StatusSes
	}
	if !IsNil(o.StatusSesNumeric) {
		toSerialize["status-ses-numeric"] = o.StatusSesNumeric
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableFanResourceInner struct {
	value *FanResourceInner
	isSet bool
}

func (v NullableFanResourceInner) Get() *FanResourceInner {
	return v.value
}

func (v *NullableFanResourceInner) Set(val *FanResourceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFanResourceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFanResourceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFanResourceInner(val *FanResourceInner) *NullableFanResourceInner {
	return &NullableFanResourceInner{value: val, isSet: true}
}

func (v NullableFanResourceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFanResourceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
