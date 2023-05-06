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

// checks if the TiersResourceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TiersResourceInner{}

// TiersResourceInner struct for TiersResourceInner
type TiersResourceInner struct {
	ObjectName           *string `json:"object-name,omitempty"`
	Meta                 *string `json:"meta,omitempty"`
	AffinitySize         *string `json:"affinity-size,omitempty"`
<<<<<<< HEAD
	AffinitySizeNumeric  *int64  `json:"affinity-size-numeric,omitempty"`
	AllocatedSize        *string `json:"allocated-size,omitempty"`
	AllocatedSizeNumeric *int64  `json:"allocated-size-numeric,omitempty"`
	AvailableSize        *string `json:"available-size,omitempty"`
	AvailableSizeNumeric *int64  `json:"available-size-numeric,omitempty"`
	// Number of disks
	Diskcount *int64 `json:"diskcount,omitempty"`
	// Pool
	Pool *string `json:"pool,omitempty"`
	// Portion of the virtual pool used by this disk group
	PoolPercentage *int64  `json:"pool-percentage,omitempty"`
	RawSize        *string `json:"raw-size,omitempty"`
	RawSizeNumeric *int64  `json:"raw-size-numeric,omitempty"`
	SerialNumber   *string `json:"serial-number,omitempty"`
	Tier           *string `json:"tier,omitempty"`
	TierNumeric    *int64  `json:"tier-numeric,omitempty"`
	// The total size formatted using the session settings for base, precision, and units
	TotalSize *string `json:"total-size,omitempty"`
	// The total size formatted using the session settings for base, precision, and units( In numeric form )
	TotalSizeNumeric *int64 `json:"total-size-numeric,omitempty"`
=======
	AffinitySizeNumeric  *int32  `json:"affinity-size-numeric,omitempty"`
	AllocatedSize        *string `json:"allocated-size,omitempty"`
	AllocatedSizeNumeric *int32  `json:"allocated-size-numeric,omitempty"`
	AvailableSize        *string `json:"available-size,omitempty"`
	AvailableSizeNumeric *int32  `json:"available-size-numeric,omitempty"`
	// Number of disks
	Diskcount *int32 `json:"diskcount,omitempty"`
	// Pool
	Pool *string `json:"pool,omitempty"`
	// Portion of the virtual pool used by this disk group
	PoolPercentage *int32  `json:"pool-percentage,omitempty"`
	RawSize        *string `json:"raw-size,omitempty"`
	RawSizeNumeric *int32  `json:"raw-size-numeric,omitempty"`
	SerialNumber   *string `json:"serial-number,omitempty"`
	Tier           *string `json:"tier,omitempty"`
	TierNumeric    *int32  `json:"tier-numeric,omitempty"`
	// The total size formatted using the session settings for base, precision, and units
	TotalSize *string `json:"total-size,omitempty"`
	// The total size formatted using the session settings for base, precision, and units( In numeric form )
	TotalSizeNumeric *int32 `json:"total-size-numeric,omitempty"`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
}

// NewTiersResourceInner instantiates a new TiersResourceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTiersResourceInner() *TiersResourceInner {
	this := TiersResourceInner{}
	return &this
}

// NewTiersResourceInnerWithDefaults instantiates a new TiersResourceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTiersResourceInnerWithDefaults() *TiersResourceInner {
	this := TiersResourceInner{}
	return &this
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *TiersResourceInner) GetObjectName() string {
	if o == nil || IsNil(o.ObjectName) {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TiersResourceInner) GetObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectName) {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *TiersResourceInner) HasObjectName() bool {
	if o != nil && !IsNil(o.ObjectName) {
		return true
	}

	return false
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *TiersResourceInner) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TiersResourceInner) GetMeta() string {
	if o == nil || IsNil(o.Meta) {
		var ret string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TiersResourceInner) GetMetaOk() (*string, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TiersResourceInner) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given string and assigns it to the Meta field.
func (o *TiersResourceInner) SetMeta(v string) {
	o.Meta = &v
}

// GetAffinitySize returns the AffinitySize field value if set, zero value otherwise.
func (o *TiersResourceInner) GetAffinitySize() string {
	if o == nil || IsNil(o.AffinitySize) {
		var ret string
		return ret
	}
	return *o.AffinitySize
}

// GetAffinitySizeOk returns a tuple with the AffinitySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TiersResourceInner) GetAffinitySizeOk() (*string, bool) {
	if o == nil || IsNil(o.AffinitySize) {
		return nil, false
	}
	return o.AffinitySize, true
}

// HasAffinitySize returns a boolean if a field has been set.
func (o *TiersResourceInner) HasAffinitySize() bool {
	if o != nil && !IsNil(o.AffinitySize) {
		return true
	}

	return false
}

// SetAffinitySize gets a reference to the given string and assigns it to the AffinitySize field.
func (o *TiersResourceInner) SetAffinitySize(v string) {
	o.AffinitySize = &v
}

// GetAffinitySizeNumeric returns the AffinitySizeNumeric field value if set, zero value otherwise.
<<<<<<< HEAD
func (o *TiersResourceInner) GetAffinitySizeNumeric() int64 {
	if o == nil || IsNil(o.AffinitySizeNumeric) {
		var ret int64
=======
func (o *TiersResourceInner) GetAffinitySizeNumeric() int32 {
	if o == nil || IsNil(o.AffinitySizeNumeric) {
		var ret int32
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
		return ret
	}
	return *o.AffinitySizeNumeric
}

// GetAffinitySizeNumericOk returns a tuple with the AffinitySizeNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
<<<<<<< HEAD
func (o *TiersResourceInner) GetAffinitySizeNumericOk() (*int64, bool) {
=======
func (o *TiersResourceInner) GetAffinitySizeNumericOk() (*int32, bool) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	if o == nil || IsNil(o.AffinitySizeNumeric) {
		return nil, false
	}
	return o.AffinitySizeNumeric, true
}

// HasAffinitySizeNumeric returns a boolean if a field has been set.
func (o *TiersResourceInner) HasAffinitySizeNumeric() bool {
	if o != nil && !IsNil(o.AffinitySizeNumeric) {
		return true
	}

	return false
}

<<<<<<< HEAD
// SetAffinitySizeNumeric gets a reference to the given int64 and assigns it to the AffinitySizeNumeric field.
func (o *TiersResourceInner) SetAffinitySizeNumeric(v int64) {
=======
// SetAffinitySizeNumeric gets a reference to the given int32 and assigns it to the AffinitySizeNumeric field.
func (o *TiersResourceInner) SetAffinitySizeNumeric(v int32) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	o.AffinitySizeNumeric = &v
}

// GetAllocatedSize returns the AllocatedSize field value if set, zero value otherwise.
func (o *TiersResourceInner) GetAllocatedSize() string {
	if o == nil || IsNil(o.AllocatedSize) {
		var ret string
		return ret
	}
	return *o.AllocatedSize
}

// GetAllocatedSizeOk returns a tuple with the AllocatedSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TiersResourceInner) GetAllocatedSizeOk() (*string, bool) {
	if o == nil || IsNil(o.AllocatedSize) {
		return nil, false
	}
	return o.AllocatedSize, true
}

// HasAllocatedSize returns a boolean if a field has been set.
func (o *TiersResourceInner) HasAllocatedSize() bool {
	if o != nil && !IsNil(o.AllocatedSize) {
		return true
	}

	return false
}

// SetAllocatedSize gets a reference to the given string and assigns it to the AllocatedSize field.
func (o *TiersResourceInner) SetAllocatedSize(v string) {
	o.AllocatedSize = &v
}

// GetAllocatedSizeNumeric returns the AllocatedSizeNumeric field value if set, zero value otherwise.
<<<<<<< HEAD
func (o *TiersResourceInner) GetAllocatedSizeNumeric() int64 {
	if o == nil || IsNil(o.AllocatedSizeNumeric) {
		var ret int64
=======
func (o *TiersResourceInner) GetAllocatedSizeNumeric() int32 {
	if o == nil || IsNil(o.AllocatedSizeNumeric) {
		var ret int32
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
		return ret
	}
	return *o.AllocatedSizeNumeric
}

// GetAllocatedSizeNumericOk returns a tuple with the AllocatedSizeNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
<<<<<<< HEAD
func (o *TiersResourceInner) GetAllocatedSizeNumericOk() (*int64, bool) {
=======
func (o *TiersResourceInner) GetAllocatedSizeNumericOk() (*int32, bool) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	if o == nil || IsNil(o.AllocatedSizeNumeric) {
		return nil, false
	}
	return o.AllocatedSizeNumeric, true
}

// HasAllocatedSizeNumeric returns a boolean if a field has been set.
func (o *TiersResourceInner) HasAllocatedSizeNumeric() bool {
	if o != nil && !IsNil(o.AllocatedSizeNumeric) {
		return true
	}

	return false
}

<<<<<<< HEAD
// SetAllocatedSizeNumeric gets a reference to the given int64 and assigns it to the AllocatedSizeNumeric field.
func (o *TiersResourceInner) SetAllocatedSizeNumeric(v int64) {
=======
// SetAllocatedSizeNumeric gets a reference to the given int32 and assigns it to the AllocatedSizeNumeric field.
func (o *TiersResourceInner) SetAllocatedSizeNumeric(v int32) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	o.AllocatedSizeNumeric = &v
}

// GetAvailableSize returns the AvailableSize field value if set, zero value otherwise.
func (o *TiersResourceInner) GetAvailableSize() string {
	if o == nil || IsNil(o.AvailableSize) {
		var ret string
		return ret
	}
	return *o.AvailableSize
}

// GetAvailableSizeOk returns a tuple with the AvailableSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TiersResourceInner) GetAvailableSizeOk() (*string, bool) {
	if o == nil || IsNil(o.AvailableSize) {
		return nil, false
	}
	return o.AvailableSize, true
}

// HasAvailableSize returns a boolean if a field has been set.
func (o *TiersResourceInner) HasAvailableSize() bool {
	if o != nil && !IsNil(o.AvailableSize) {
		return true
	}

	return false
}

// SetAvailableSize gets a reference to the given string and assigns it to the AvailableSize field.
func (o *TiersResourceInner) SetAvailableSize(v string) {
	o.AvailableSize = &v
}

// GetAvailableSizeNumeric returns the AvailableSizeNumeric field value if set, zero value otherwise.
<<<<<<< HEAD
func (o *TiersResourceInner) GetAvailableSizeNumeric() int64 {
	if o == nil || IsNil(o.AvailableSizeNumeric) {
		var ret int64
=======
func (o *TiersResourceInner) GetAvailableSizeNumeric() int32 {
	if o == nil || IsNil(o.AvailableSizeNumeric) {
		var ret int32
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
		return ret
	}
	return *o.AvailableSizeNumeric
}

// GetAvailableSizeNumericOk returns a tuple with the AvailableSizeNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
<<<<<<< HEAD
func (o *TiersResourceInner) GetAvailableSizeNumericOk() (*int64, bool) {
=======
func (o *TiersResourceInner) GetAvailableSizeNumericOk() (*int32, bool) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	if o == nil || IsNil(o.AvailableSizeNumeric) {
		return nil, false
	}
	return o.AvailableSizeNumeric, true
}

// HasAvailableSizeNumeric returns a boolean if a field has been set.
func (o *TiersResourceInner) HasAvailableSizeNumeric() bool {
	if o != nil && !IsNil(o.AvailableSizeNumeric) {
		return true
	}

	return false
}

<<<<<<< HEAD
// SetAvailableSizeNumeric gets a reference to the given int64 and assigns it to the AvailableSizeNumeric field.
func (o *TiersResourceInner) SetAvailableSizeNumeric(v int64) {
=======
// SetAvailableSizeNumeric gets a reference to the given int32 and assigns it to the AvailableSizeNumeric field.
func (o *TiersResourceInner) SetAvailableSizeNumeric(v int32) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	o.AvailableSizeNumeric = &v
}

// GetDiskcount returns the Diskcount field value if set, zero value otherwise.
<<<<<<< HEAD
func (o *TiersResourceInner) GetDiskcount() int64 {
	if o == nil || IsNil(o.Diskcount) {
		var ret int64
=======
func (o *TiersResourceInner) GetDiskcount() int32 {
	if o == nil || IsNil(o.Diskcount) {
		var ret int32
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
		return ret
	}
	return *o.Diskcount
}

// GetDiskcountOk returns a tuple with the Diskcount field value if set, nil otherwise
// and a boolean to check if the value has been set.
<<<<<<< HEAD
func (o *TiersResourceInner) GetDiskcountOk() (*int64, bool) {
=======
func (o *TiersResourceInner) GetDiskcountOk() (*int32, bool) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	if o == nil || IsNil(o.Diskcount) {
		return nil, false
	}
	return o.Diskcount, true
}

// HasDiskcount returns a boolean if a field has been set.
func (o *TiersResourceInner) HasDiskcount() bool {
	if o != nil && !IsNil(o.Diskcount) {
		return true
	}

	return false
}

<<<<<<< HEAD
// SetDiskcount gets a reference to the given int64 and assigns it to the Diskcount field.
func (o *TiersResourceInner) SetDiskcount(v int64) {
=======
// SetDiskcount gets a reference to the given int32 and assigns it to the Diskcount field.
func (o *TiersResourceInner) SetDiskcount(v int32) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	o.Diskcount = &v
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *TiersResourceInner) GetPool() string {
	if o == nil || IsNil(o.Pool) {
		var ret string
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TiersResourceInner) GetPoolOk() (*string, bool) {
	if o == nil || IsNil(o.Pool) {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *TiersResourceInner) HasPool() bool {
	if o != nil && !IsNil(o.Pool) {
		return true
	}

	return false
}

// SetPool gets a reference to the given string and assigns it to the Pool field.
func (o *TiersResourceInner) SetPool(v string) {
	o.Pool = &v
}

// GetPoolPercentage returns the PoolPercentage field value if set, zero value otherwise.
<<<<<<< HEAD
func (o *TiersResourceInner) GetPoolPercentage() int64 {
	if o == nil || IsNil(o.PoolPercentage) {
		var ret int64
=======
func (o *TiersResourceInner) GetPoolPercentage() int32 {
	if o == nil || IsNil(o.PoolPercentage) {
		var ret int32
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
		return ret
	}
	return *o.PoolPercentage
}

// GetPoolPercentageOk returns a tuple with the PoolPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
<<<<<<< HEAD
func (o *TiersResourceInner) GetPoolPercentageOk() (*int64, bool) {
=======
func (o *TiersResourceInner) GetPoolPercentageOk() (*int32, bool) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	if o == nil || IsNil(o.PoolPercentage) {
		return nil, false
	}
	return o.PoolPercentage, true
}

// HasPoolPercentage returns a boolean if a field has been set.
func (o *TiersResourceInner) HasPoolPercentage() bool {
	if o != nil && !IsNil(o.PoolPercentage) {
		return true
	}

	return false
}

<<<<<<< HEAD
// SetPoolPercentage gets a reference to the given int64 and assigns it to the PoolPercentage field.
func (o *TiersResourceInner) SetPoolPercentage(v int64) {
=======
// SetPoolPercentage gets a reference to the given int32 and assigns it to the PoolPercentage field.
func (o *TiersResourceInner) SetPoolPercentage(v int32) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	o.PoolPercentage = &v
}

// GetRawSize returns the RawSize field value if set, zero value otherwise.
func (o *TiersResourceInner) GetRawSize() string {
	if o == nil || IsNil(o.RawSize) {
		var ret string
		return ret
	}
	return *o.RawSize
}

// GetRawSizeOk returns a tuple with the RawSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TiersResourceInner) GetRawSizeOk() (*string, bool) {
	if o == nil || IsNil(o.RawSize) {
		return nil, false
	}
	return o.RawSize, true
}

// HasRawSize returns a boolean if a field has been set.
func (o *TiersResourceInner) HasRawSize() bool {
	if o != nil && !IsNil(o.RawSize) {
		return true
	}

	return false
}

// SetRawSize gets a reference to the given string and assigns it to the RawSize field.
func (o *TiersResourceInner) SetRawSize(v string) {
	o.RawSize = &v
}

// GetRawSizeNumeric returns the RawSizeNumeric field value if set, zero value otherwise.
<<<<<<< HEAD
func (o *TiersResourceInner) GetRawSizeNumeric() int64 {
	if o == nil || IsNil(o.RawSizeNumeric) {
		var ret int64
=======
func (o *TiersResourceInner) GetRawSizeNumeric() int32 {
	if o == nil || IsNil(o.RawSizeNumeric) {
		var ret int32
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
		return ret
	}
	return *o.RawSizeNumeric
}

// GetRawSizeNumericOk returns a tuple with the RawSizeNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
<<<<<<< HEAD
func (o *TiersResourceInner) GetRawSizeNumericOk() (*int64, bool) {
=======
func (o *TiersResourceInner) GetRawSizeNumericOk() (*int32, bool) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	if o == nil || IsNil(o.RawSizeNumeric) {
		return nil, false
	}
	return o.RawSizeNumeric, true
}

// HasRawSizeNumeric returns a boolean if a field has been set.
func (o *TiersResourceInner) HasRawSizeNumeric() bool {
	if o != nil && !IsNil(o.RawSizeNumeric) {
		return true
	}

	return false
}

<<<<<<< HEAD
// SetRawSizeNumeric gets a reference to the given int64 and assigns it to the RawSizeNumeric field.
func (o *TiersResourceInner) SetRawSizeNumeric(v int64) {
=======
// SetRawSizeNumeric gets a reference to the given int32 and assigns it to the RawSizeNumeric field.
func (o *TiersResourceInner) SetRawSizeNumeric(v int32) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	o.RawSizeNumeric = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *TiersResourceInner) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TiersResourceInner) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *TiersResourceInner) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *TiersResourceInner) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *TiersResourceInner) GetTier() string {
	if o == nil || IsNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TiersResourceInner) GetTierOk() (*string, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *TiersResourceInner) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *TiersResourceInner) SetTier(v string) {
	o.Tier = &v
}

// GetTierNumeric returns the TierNumeric field value if set, zero value otherwise.
<<<<<<< HEAD
func (o *TiersResourceInner) GetTierNumeric() int64 {
	if o == nil || IsNil(o.TierNumeric) {
		var ret int64
=======
func (o *TiersResourceInner) GetTierNumeric() int32 {
	if o == nil || IsNil(o.TierNumeric) {
		var ret int32
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
		return ret
	}
	return *o.TierNumeric
}

// GetTierNumericOk returns a tuple with the TierNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
<<<<<<< HEAD
func (o *TiersResourceInner) GetTierNumericOk() (*int64, bool) {
=======
func (o *TiersResourceInner) GetTierNumericOk() (*int32, bool) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	if o == nil || IsNil(o.TierNumeric) {
		return nil, false
	}
	return o.TierNumeric, true
}

// HasTierNumeric returns a boolean if a field has been set.
func (o *TiersResourceInner) HasTierNumeric() bool {
	if o != nil && !IsNil(o.TierNumeric) {
		return true
	}

	return false
}

<<<<<<< HEAD
// SetTierNumeric gets a reference to the given int64 and assigns it to the TierNumeric field.
func (o *TiersResourceInner) SetTierNumeric(v int64) {
=======
// SetTierNumeric gets a reference to the given int32 and assigns it to the TierNumeric field.
func (o *TiersResourceInner) SetTierNumeric(v int32) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	o.TierNumeric = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *TiersResourceInner) GetTotalSize() string {
	if o == nil || IsNil(o.TotalSize) {
		var ret string
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TiersResourceInner) GetTotalSizeOk() (*string, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *TiersResourceInner) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given string and assigns it to the TotalSize field.
func (o *TiersResourceInner) SetTotalSize(v string) {
	o.TotalSize = &v
}

// GetTotalSizeNumeric returns the TotalSizeNumeric field value if set, zero value otherwise.
<<<<<<< HEAD
func (o *TiersResourceInner) GetTotalSizeNumeric() int64 {
	if o == nil || IsNil(o.TotalSizeNumeric) {
		var ret int64
=======
func (o *TiersResourceInner) GetTotalSizeNumeric() int32 {
	if o == nil || IsNil(o.TotalSizeNumeric) {
		var ret int32
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
		return ret
	}
	return *o.TotalSizeNumeric
}

// GetTotalSizeNumericOk returns a tuple with the TotalSizeNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
<<<<<<< HEAD
func (o *TiersResourceInner) GetTotalSizeNumericOk() (*int64, bool) {
=======
func (o *TiersResourceInner) GetTotalSizeNumericOk() (*int32, bool) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	if o == nil || IsNil(o.TotalSizeNumeric) {
		return nil, false
	}
	return o.TotalSizeNumeric, true
}

// HasTotalSizeNumeric returns a boolean if a field has been set.
func (o *TiersResourceInner) HasTotalSizeNumeric() bool {
	if o != nil && !IsNil(o.TotalSizeNumeric) {
		return true
	}

	return false
}

<<<<<<< HEAD
// SetTotalSizeNumeric gets a reference to the given int64 and assigns it to the TotalSizeNumeric field.
func (o *TiersResourceInner) SetTotalSizeNumeric(v int64) {
=======
// SetTotalSizeNumeric gets a reference to the given int32 and assigns it to the TotalSizeNumeric field.
func (o *TiersResourceInner) SetTotalSizeNumeric(v int32) {
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)
	o.TotalSizeNumeric = &v
}

func (o TiersResourceInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TiersResourceInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectName) {
		toSerialize["object-name"] = o.ObjectName
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.AffinitySize) {
		toSerialize["affinity-size"] = o.AffinitySize
	}
	if !IsNil(o.AffinitySizeNumeric) {
		toSerialize["affinity-size-numeric"] = o.AffinitySizeNumeric
	}
	if !IsNil(o.AllocatedSize) {
		toSerialize["allocated-size"] = o.AllocatedSize
	}
	if !IsNil(o.AllocatedSizeNumeric) {
		toSerialize["allocated-size-numeric"] = o.AllocatedSizeNumeric
	}
	if !IsNil(o.AvailableSize) {
		toSerialize["available-size"] = o.AvailableSize
	}
	if !IsNil(o.AvailableSizeNumeric) {
		toSerialize["available-size-numeric"] = o.AvailableSizeNumeric
	}
	if !IsNil(o.Diskcount) {
		toSerialize["diskcount"] = o.Diskcount
	}
	if !IsNil(o.Pool) {
		toSerialize["pool"] = o.Pool
	}
	if !IsNil(o.PoolPercentage) {
		toSerialize["pool-percentage"] = o.PoolPercentage
	}
	if !IsNil(o.RawSize) {
		toSerialize["raw-size"] = o.RawSize
	}
	if !IsNil(o.RawSizeNumeric) {
		toSerialize["raw-size-numeric"] = o.RawSizeNumeric
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serial-number"] = o.SerialNumber
	}
	if !IsNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}
	if !IsNil(o.TierNumeric) {
		toSerialize["tier-numeric"] = o.TierNumeric
	}
	if !IsNil(o.TotalSize) {
		toSerialize["total-size"] = o.TotalSize
	}
	if !IsNil(o.TotalSizeNumeric) {
		toSerialize["total-size-numeric"] = o.TotalSizeNumeric
	}
	return toSerialize, nil
}

type NullableTiersResourceInner struct {
	value *TiersResourceInner
	isSet bool
}

func (v NullableTiersResourceInner) Get() *TiersResourceInner {
	return v.value
}

func (v *NullableTiersResourceInner) Set(val *TiersResourceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTiersResourceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTiersResourceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTiersResourceInner(val *TiersResourceInner) *NullableTiersResourceInner {
	return &NullableTiersResourceInner{value: val, isSet: true}
}

func (v NullableTiersResourceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTiersResourceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
