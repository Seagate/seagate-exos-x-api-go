# SnapshotsResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**BaseSerialNumber** | Pointer to **string** | The root of the snapshot tree serial number | [optional] 
**BaseVolume** | Pointer to **string** | The root of the snapshot tree | [optional] 
**BaseVolumesUrl** | Pointer to **string** | The resource URL of the root of the snapshot tree | [optional] 
**CreationDateTime** | Pointer to **string** |  | [optional] 
**CreationDateTimeNumeric** | Pointer to **int32** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**MasterVolumeName** | Pointer to **string** | Source volume for snapshots | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumChildren** | Pointer to **int32** |  | [optional] 
**NumSnapsTree** | Pointer to **int32** |  | [optional] 
**PriorityValue** | Pointer to **string** | Attribute priority | [optional] 
**RetentionPriority** | Pointer to **string** | Retention priority of the snapshot or children | [optional] 
**RetentionPriorityNumeric** | Pointer to **int32** | Retention priority of the snapshot or children( In numeric form ) | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Shareddata** | Pointer to **string** |  | [optional] 
**ShareddataNumeric** | Pointer to **int32** |  | [optional] 
**SnapData** | Pointer to **string** | Amount of snapshot data associated with this volume | [optional] 
**SnapDataNumeric** | Pointer to **int32** | Amount of snapshot data associated with this volume( In numeric form ) | [optional] 
**SnapPoolName** | Pointer to **string** | User-defined name of the snap pool | [optional] 
**SnapshotType** | Pointer to **string** | Associated snapshot type | [optional] 
**SnapshotTypeNumeric** | Pointer to **int32** | Associated snapshot type( In numeric form ) | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusNumeric** | Pointer to **int32** |  | [optional] 
**StatusReason** | Pointer to **string** | Probable cause for the current status | [optional] 
**StatusReasonNumeric** | Pointer to **int32** | Probable cause for the current status( In numeric form ) | [optional] 
**StoragePoolName** | Pointer to **string** | User-defined name for the pool | [optional] 
**StoragePoolsUrl** | Pointer to **string** |  | [optional] 
**StorageType** | Pointer to **string** | Storage type | [optional] 
**StorageTypeNumeric** | Pointer to **int32** | Storage type( In numeric form ) | [optional] 
**TotalSize** | Pointer to **string** | The total size formatted using the session settings for base, precision, and units | [optional] 
**TotalSizeNumeric** | Pointer to **int32** | The total size formatted using the session settings for base, precision, and units( In numeric form ) | [optional] 
**Uniquedata** | Pointer to **string** | Amount of preserved and write data that is unique to the snapshot | [optional] 
**UniquedataNumeric** | Pointer to **int32** | Amount of preserved and write data that is unique to the snapshot( In numeric form ) | [optional] 
**Url** | Pointer to **string** | The resource URL | [optional] 
**UserPriorityValue** | Pointer to **string** | User-assigned priority | [optional] 
**VirtualDiskName** | Pointer to **string** | The name of the pool that contains the volume | [optional] 
**VolumeParent** | Pointer to **string** | Serial number of the associated volume | [optional] 

## Methods

### NewSnapshotsResourceInner

`func NewSnapshotsResourceInner() *SnapshotsResourceInner`

NewSnapshotsResourceInner instantiates a new SnapshotsResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotsResourceInnerWithDefaults

`func NewSnapshotsResourceInnerWithDefaults() *SnapshotsResourceInner`

NewSnapshotsResourceInnerWithDefaults instantiates a new SnapshotsResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *SnapshotsResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *SnapshotsResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *SnapshotsResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *SnapshotsResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *SnapshotsResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SnapshotsResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SnapshotsResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SnapshotsResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetBaseSerialNumber

`func (o *SnapshotsResourceInner) GetBaseSerialNumber() string`

GetBaseSerialNumber returns the BaseSerialNumber field if non-nil, zero value otherwise.

### GetBaseSerialNumberOk

`func (o *SnapshotsResourceInner) GetBaseSerialNumberOk() (*string, bool)`

GetBaseSerialNumberOk returns a tuple with the BaseSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSerialNumber

`func (o *SnapshotsResourceInner) SetBaseSerialNumber(v string)`

SetBaseSerialNumber sets BaseSerialNumber field to given value.

### HasBaseSerialNumber

`func (o *SnapshotsResourceInner) HasBaseSerialNumber() bool`

HasBaseSerialNumber returns a boolean if a field has been set.

### GetBaseVolume

`func (o *SnapshotsResourceInner) GetBaseVolume() string`

GetBaseVolume returns the BaseVolume field if non-nil, zero value otherwise.

### GetBaseVolumeOk

`func (o *SnapshotsResourceInner) GetBaseVolumeOk() (*string, bool)`

GetBaseVolumeOk returns a tuple with the BaseVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVolume

`func (o *SnapshotsResourceInner) SetBaseVolume(v string)`

SetBaseVolume sets BaseVolume field to given value.

### HasBaseVolume

`func (o *SnapshotsResourceInner) HasBaseVolume() bool`

HasBaseVolume returns a boolean if a field has been set.

### GetBaseVolumesUrl

`func (o *SnapshotsResourceInner) GetBaseVolumesUrl() string`

GetBaseVolumesUrl returns the BaseVolumesUrl field if non-nil, zero value otherwise.

### GetBaseVolumesUrlOk

`func (o *SnapshotsResourceInner) GetBaseVolumesUrlOk() (*string, bool)`

GetBaseVolumesUrlOk returns a tuple with the BaseVolumesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVolumesUrl

`func (o *SnapshotsResourceInner) SetBaseVolumesUrl(v string)`

SetBaseVolumesUrl sets BaseVolumesUrl field to given value.

### HasBaseVolumesUrl

`func (o *SnapshotsResourceInner) HasBaseVolumesUrl() bool`

HasBaseVolumesUrl returns a boolean if a field has been set.

### GetCreationDateTime

`func (o *SnapshotsResourceInner) GetCreationDateTime() string`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *SnapshotsResourceInner) GetCreationDateTimeOk() (*string, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *SnapshotsResourceInner) SetCreationDateTime(v string)`

SetCreationDateTime sets CreationDateTime field to given value.

### HasCreationDateTime

`func (o *SnapshotsResourceInner) HasCreationDateTime() bool`

HasCreationDateTime returns a boolean if a field has been set.

### GetCreationDateTimeNumeric

`func (o *SnapshotsResourceInner) GetCreationDateTimeNumeric() int32`

GetCreationDateTimeNumeric returns the CreationDateTimeNumeric field if non-nil, zero value otherwise.

### GetCreationDateTimeNumericOk

`func (o *SnapshotsResourceInner) GetCreationDateTimeNumericOk() (*int32, bool)`

GetCreationDateTimeNumericOk returns a tuple with the CreationDateTimeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTimeNumeric

`func (o *SnapshotsResourceInner) SetCreationDateTimeNumeric(v int32)`

SetCreationDateTimeNumeric sets CreationDateTimeNumeric field to given value.

### HasCreationDateTimeNumeric

`func (o *SnapshotsResourceInner) HasCreationDateTimeNumeric() bool`

HasCreationDateTimeNumeric returns a boolean if a field has been set.

### GetDurableId

`func (o *SnapshotsResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *SnapshotsResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *SnapshotsResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *SnapshotsResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetMasterVolumeName

`func (o *SnapshotsResourceInner) GetMasterVolumeName() string`

GetMasterVolumeName returns the MasterVolumeName field if non-nil, zero value otherwise.

### GetMasterVolumeNameOk

`func (o *SnapshotsResourceInner) GetMasterVolumeNameOk() (*string, bool)`

GetMasterVolumeNameOk returns a tuple with the MasterVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterVolumeName

`func (o *SnapshotsResourceInner) SetMasterVolumeName(v string)`

SetMasterVolumeName sets MasterVolumeName field to given value.

### HasMasterVolumeName

`func (o *SnapshotsResourceInner) HasMasterVolumeName() bool`

HasMasterVolumeName returns a boolean if a field has been set.

### GetName

`func (o *SnapshotsResourceInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnapshotsResourceInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnapshotsResourceInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnapshotsResourceInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumChildren

`func (o *SnapshotsResourceInner) GetNumChildren() int32`

GetNumChildren returns the NumChildren field if non-nil, zero value otherwise.

### GetNumChildrenOk

`func (o *SnapshotsResourceInner) GetNumChildrenOk() (*int32, bool)`

GetNumChildrenOk returns a tuple with the NumChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumChildren

`func (o *SnapshotsResourceInner) SetNumChildren(v int32)`

SetNumChildren sets NumChildren field to given value.

### HasNumChildren

`func (o *SnapshotsResourceInner) HasNumChildren() bool`

HasNumChildren returns a boolean if a field has been set.

### GetNumSnapsTree

`func (o *SnapshotsResourceInner) GetNumSnapsTree() int32`

GetNumSnapsTree returns the NumSnapsTree field if non-nil, zero value otherwise.

### GetNumSnapsTreeOk

`func (o *SnapshotsResourceInner) GetNumSnapsTreeOk() (*int32, bool)`

GetNumSnapsTreeOk returns a tuple with the NumSnapsTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSnapsTree

`func (o *SnapshotsResourceInner) SetNumSnapsTree(v int32)`

SetNumSnapsTree sets NumSnapsTree field to given value.

### HasNumSnapsTree

`func (o *SnapshotsResourceInner) HasNumSnapsTree() bool`

HasNumSnapsTree returns a boolean if a field has been set.

### GetPriorityValue

`func (o *SnapshotsResourceInner) GetPriorityValue() string`

GetPriorityValue returns the PriorityValue field if non-nil, zero value otherwise.

### GetPriorityValueOk

`func (o *SnapshotsResourceInner) GetPriorityValueOk() (*string, bool)`

GetPriorityValueOk returns a tuple with the PriorityValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityValue

`func (o *SnapshotsResourceInner) SetPriorityValue(v string)`

SetPriorityValue sets PriorityValue field to given value.

### HasPriorityValue

`func (o *SnapshotsResourceInner) HasPriorityValue() bool`

HasPriorityValue returns a boolean if a field has been set.

### GetRetentionPriority

`func (o *SnapshotsResourceInner) GetRetentionPriority() string`

GetRetentionPriority returns the RetentionPriority field if non-nil, zero value otherwise.

### GetRetentionPriorityOk

`func (o *SnapshotsResourceInner) GetRetentionPriorityOk() (*string, bool)`

GetRetentionPriorityOk returns a tuple with the RetentionPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPriority

`func (o *SnapshotsResourceInner) SetRetentionPriority(v string)`

SetRetentionPriority sets RetentionPriority field to given value.

### HasRetentionPriority

`func (o *SnapshotsResourceInner) HasRetentionPriority() bool`

HasRetentionPriority returns a boolean if a field has been set.

### GetRetentionPriorityNumeric

`func (o *SnapshotsResourceInner) GetRetentionPriorityNumeric() int32`

GetRetentionPriorityNumeric returns the RetentionPriorityNumeric field if non-nil, zero value otherwise.

### GetRetentionPriorityNumericOk

`func (o *SnapshotsResourceInner) GetRetentionPriorityNumericOk() (*int32, bool)`

GetRetentionPriorityNumericOk returns a tuple with the RetentionPriorityNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPriorityNumeric

`func (o *SnapshotsResourceInner) SetRetentionPriorityNumeric(v int32)`

SetRetentionPriorityNumeric sets RetentionPriorityNumeric field to given value.

### HasRetentionPriorityNumeric

`func (o *SnapshotsResourceInner) HasRetentionPriorityNumeric() bool`

HasRetentionPriorityNumeric returns a boolean if a field has been set.

### GetSerialNumber

`func (o *SnapshotsResourceInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SnapshotsResourceInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SnapshotsResourceInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *SnapshotsResourceInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetShareddata

`func (o *SnapshotsResourceInner) GetShareddata() string`

GetShareddata returns the Shareddata field if non-nil, zero value otherwise.

### GetShareddataOk

`func (o *SnapshotsResourceInner) GetShareddataOk() (*string, bool)`

GetShareddataOk returns a tuple with the Shareddata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareddata

`func (o *SnapshotsResourceInner) SetShareddata(v string)`

SetShareddata sets Shareddata field to given value.

### HasShareddata

`func (o *SnapshotsResourceInner) HasShareddata() bool`

HasShareddata returns a boolean if a field has been set.

### GetShareddataNumeric

`func (o *SnapshotsResourceInner) GetShareddataNumeric() int32`

GetShareddataNumeric returns the ShareddataNumeric field if non-nil, zero value otherwise.

### GetShareddataNumericOk

`func (o *SnapshotsResourceInner) GetShareddataNumericOk() (*int32, bool)`

GetShareddataNumericOk returns a tuple with the ShareddataNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareddataNumeric

`func (o *SnapshotsResourceInner) SetShareddataNumeric(v int32)`

SetShareddataNumeric sets ShareddataNumeric field to given value.

### HasShareddataNumeric

`func (o *SnapshotsResourceInner) HasShareddataNumeric() bool`

HasShareddataNumeric returns a boolean if a field has been set.

### GetSnapData

`func (o *SnapshotsResourceInner) GetSnapData() string`

GetSnapData returns the SnapData field if non-nil, zero value otherwise.

### GetSnapDataOk

`func (o *SnapshotsResourceInner) GetSnapDataOk() (*string, bool)`

GetSnapDataOk returns a tuple with the SnapData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapData

`func (o *SnapshotsResourceInner) SetSnapData(v string)`

SetSnapData sets SnapData field to given value.

### HasSnapData

`func (o *SnapshotsResourceInner) HasSnapData() bool`

HasSnapData returns a boolean if a field has been set.

### GetSnapDataNumeric

`func (o *SnapshotsResourceInner) GetSnapDataNumeric() int32`

GetSnapDataNumeric returns the SnapDataNumeric field if non-nil, zero value otherwise.

### GetSnapDataNumericOk

`func (o *SnapshotsResourceInner) GetSnapDataNumericOk() (*int32, bool)`

GetSnapDataNumericOk returns a tuple with the SnapDataNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapDataNumeric

`func (o *SnapshotsResourceInner) SetSnapDataNumeric(v int32)`

SetSnapDataNumeric sets SnapDataNumeric field to given value.

### HasSnapDataNumeric

`func (o *SnapshotsResourceInner) HasSnapDataNumeric() bool`

HasSnapDataNumeric returns a boolean if a field has been set.

### GetSnapPoolName

`func (o *SnapshotsResourceInner) GetSnapPoolName() string`

GetSnapPoolName returns the SnapPoolName field if non-nil, zero value otherwise.

### GetSnapPoolNameOk

`func (o *SnapshotsResourceInner) GetSnapPoolNameOk() (*string, bool)`

GetSnapPoolNameOk returns a tuple with the SnapPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapPoolName

`func (o *SnapshotsResourceInner) SetSnapPoolName(v string)`

SetSnapPoolName sets SnapPoolName field to given value.

### HasSnapPoolName

`func (o *SnapshotsResourceInner) HasSnapPoolName() bool`

HasSnapPoolName returns a boolean if a field has been set.

### GetSnapshotType

`func (o *SnapshotsResourceInner) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *SnapshotsResourceInner) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *SnapshotsResourceInner) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *SnapshotsResourceInner) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetSnapshotTypeNumeric

`func (o *SnapshotsResourceInner) GetSnapshotTypeNumeric() int32`

GetSnapshotTypeNumeric returns the SnapshotTypeNumeric field if non-nil, zero value otherwise.

### GetSnapshotTypeNumericOk

`func (o *SnapshotsResourceInner) GetSnapshotTypeNumericOk() (*int32, bool)`

GetSnapshotTypeNumericOk returns a tuple with the SnapshotTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTypeNumeric

`func (o *SnapshotsResourceInner) SetSnapshotTypeNumeric(v int32)`

SetSnapshotTypeNumeric sets SnapshotTypeNumeric field to given value.

### HasSnapshotTypeNumeric

`func (o *SnapshotsResourceInner) HasSnapshotTypeNumeric() bool`

HasSnapshotTypeNumeric returns a boolean if a field has been set.

### GetStatus

`func (o *SnapshotsResourceInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SnapshotsResourceInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SnapshotsResourceInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SnapshotsResourceInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusNumeric

`func (o *SnapshotsResourceInner) GetStatusNumeric() int32`

GetStatusNumeric returns the StatusNumeric field if non-nil, zero value otherwise.

### GetStatusNumericOk

`func (o *SnapshotsResourceInner) GetStatusNumericOk() (*int32, bool)`

GetStatusNumericOk returns a tuple with the StatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusNumeric

`func (o *SnapshotsResourceInner) SetStatusNumeric(v int32)`

SetStatusNumeric sets StatusNumeric field to given value.

### HasStatusNumeric

`func (o *SnapshotsResourceInner) HasStatusNumeric() bool`

HasStatusNumeric returns a boolean if a field has been set.

### GetStatusReason

`func (o *SnapshotsResourceInner) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *SnapshotsResourceInner) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *SnapshotsResourceInner) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *SnapshotsResourceInner) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetStatusReasonNumeric

`func (o *SnapshotsResourceInner) GetStatusReasonNumeric() int32`

GetStatusReasonNumeric returns the StatusReasonNumeric field if non-nil, zero value otherwise.

### GetStatusReasonNumericOk

`func (o *SnapshotsResourceInner) GetStatusReasonNumericOk() (*int32, bool)`

GetStatusReasonNumericOk returns a tuple with the StatusReasonNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReasonNumeric

`func (o *SnapshotsResourceInner) SetStatusReasonNumeric(v int32)`

SetStatusReasonNumeric sets StatusReasonNumeric field to given value.

### HasStatusReasonNumeric

`func (o *SnapshotsResourceInner) HasStatusReasonNumeric() bool`

HasStatusReasonNumeric returns a boolean if a field has been set.

### GetStoragePoolName

`func (o *SnapshotsResourceInner) GetStoragePoolName() string`

GetStoragePoolName returns the StoragePoolName field if non-nil, zero value otherwise.

### GetStoragePoolNameOk

`func (o *SnapshotsResourceInner) GetStoragePoolNameOk() (*string, bool)`

GetStoragePoolNameOk returns a tuple with the StoragePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolName

`func (o *SnapshotsResourceInner) SetStoragePoolName(v string)`

SetStoragePoolName sets StoragePoolName field to given value.

### HasStoragePoolName

`func (o *SnapshotsResourceInner) HasStoragePoolName() bool`

HasStoragePoolName returns a boolean if a field has been set.

### GetStoragePoolsUrl

`func (o *SnapshotsResourceInner) GetStoragePoolsUrl() string`

GetStoragePoolsUrl returns the StoragePoolsUrl field if non-nil, zero value otherwise.

### GetStoragePoolsUrlOk

`func (o *SnapshotsResourceInner) GetStoragePoolsUrlOk() (*string, bool)`

GetStoragePoolsUrlOk returns a tuple with the StoragePoolsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolsUrl

`func (o *SnapshotsResourceInner) SetStoragePoolsUrl(v string)`

SetStoragePoolsUrl sets StoragePoolsUrl field to given value.

### HasStoragePoolsUrl

`func (o *SnapshotsResourceInner) HasStoragePoolsUrl() bool`

HasStoragePoolsUrl returns a boolean if a field has been set.

### GetStorageType

`func (o *SnapshotsResourceInner) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *SnapshotsResourceInner) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *SnapshotsResourceInner) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *SnapshotsResourceInner) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetStorageTypeNumeric

`func (o *SnapshotsResourceInner) GetStorageTypeNumeric() int32`

GetStorageTypeNumeric returns the StorageTypeNumeric field if non-nil, zero value otherwise.

### GetStorageTypeNumericOk

`func (o *SnapshotsResourceInner) GetStorageTypeNumericOk() (*int32, bool)`

GetStorageTypeNumericOk returns a tuple with the StorageTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTypeNumeric

`func (o *SnapshotsResourceInner) SetStorageTypeNumeric(v int32)`

SetStorageTypeNumeric sets StorageTypeNumeric field to given value.

### HasStorageTypeNumeric

`func (o *SnapshotsResourceInner) HasStorageTypeNumeric() bool`

HasStorageTypeNumeric returns a boolean if a field has been set.

### GetTotalSize

`func (o *SnapshotsResourceInner) GetTotalSize() string`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *SnapshotsResourceInner) GetTotalSizeOk() (*string, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *SnapshotsResourceInner) SetTotalSize(v string)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *SnapshotsResourceInner) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetTotalSizeNumeric

`func (o *SnapshotsResourceInner) GetTotalSizeNumeric() int32`

GetTotalSizeNumeric returns the TotalSizeNumeric field if non-nil, zero value otherwise.

### GetTotalSizeNumericOk

`func (o *SnapshotsResourceInner) GetTotalSizeNumericOk() (*int32, bool)`

GetTotalSizeNumericOk returns a tuple with the TotalSizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSizeNumeric

`func (o *SnapshotsResourceInner) SetTotalSizeNumeric(v int32)`

SetTotalSizeNumeric sets TotalSizeNumeric field to given value.

### HasTotalSizeNumeric

`func (o *SnapshotsResourceInner) HasTotalSizeNumeric() bool`

HasTotalSizeNumeric returns a boolean if a field has been set.

### GetUniquedata

`func (o *SnapshotsResourceInner) GetUniquedata() string`

GetUniquedata returns the Uniquedata field if non-nil, zero value otherwise.

### GetUniquedataOk

`func (o *SnapshotsResourceInner) GetUniquedataOk() (*string, bool)`

GetUniquedataOk returns a tuple with the Uniquedata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniquedata

`func (o *SnapshotsResourceInner) SetUniquedata(v string)`

SetUniquedata sets Uniquedata field to given value.

### HasUniquedata

`func (o *SnapshotsResourceInner) HasUniquedata() bool`

HasUniquedata returns a boolean if a field has been set.

### GetUniquedataNumeric

`func (o *SnapshotsResourceInner) GetUniquedataNumeric() int32`

GetUniquedataNumeric returns the UniquedataNumeric field if non-nil, zero value otherwise.

### GetUniquedataNumericOk

`func (o *SnapshotsResourceInner) GetUniquedataNumericOk() (*int32, bool)`

GetUniquedataNumericOk returns a tuple with the UniquedataNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniquedataNumeric

`func (o *SnapshotsResourceInner) SetUniquedataNumeric(v int32)`

SetUniquedataNumeric sets UniquedataNumeric field to given value.

### HasUniquedataNumeric

`func (o *SnapshotsResourceInner) HasUniquedataNumeric() bool`

HasUniquedataNumeric returns a boolean if a field has been set.

### GetUrl

`func (o *SnapshotsResourceInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SnapshotsResourceInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SnapshotsResourceInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SnapshotsResourceInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserPriorityValue

`func (o *SnapshotsResourceInner) GetUserPriorityValue() string`

GetUserPriorityValue returns the UserPriorityValue field if non-nil, zero value otherwise.

### GetUserPriorityValueOk

`func (o *SnapshotsResourceInner) GetUserPriorityValueOk() (*string, bool)`

GetUserPriorityValueOk returns a tuple with the UserPriorityValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPriorityValue

`func (o *SnapshotsResourceInner) SetUserPriorityValue(v string)`

SetUserPriorityValue sets UserPriorityValue field to given value.

### HasUserPriorityValue

`func (o *SnapshotsResourceInner) HasUserPriorityValue() bool`

HasUserPriorityValue returns a boolean if a field has been set.

### GetVirtualDiskName

`func (o *SnapshotsResourceInner) GetVirtualDiskName() string`

GetVirtualDiskName returns the VirtualDiskName field if non-nil, zero value otherwise.

### GetVirtualDiskNameOk

`func (o *SnapshotsResourceInner) GetVirtualDiskNameOk() (*string, bool)`

GetVirtualDiskNameOk returns a tuple with the VirtualDiskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskName

`func (o *SnapshotsResourceInner) SetVirtualDiskName(v string)`

SetVirtualDiskName sets VirtualDiskName field to given value.

### HasVirtualDiskName

`func (o *SnapshotsResourceInner) HasVirtualDiskName() bool`

HasVirtualDiskName returns a boolean if a field has been set.

### GetVolumeParent

`func (o *SnapshotsResourceInner) GetVolumeParent() string`

GetVolumeParent returns the VolumeParent field if non-nil, zero value otherwise.

### GetVolumeParentOk

`func (o *SnapshotsResourceInner) GetVolumeParentOk() (*string, bool)`

GetVolumeParentOk returns a tuple with the VolumeParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeParent

`func (o *SnapshotsResourceInner) SetVolumeParent(v string)`

SetVolumeParent sets VolumeParent field to given value.

### HasVolumeParent

`func (o *SnapshotsResourceInner) HasVolumeParent() bool`

HasVolumeParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


