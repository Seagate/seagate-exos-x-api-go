# AdvancedSettingsTableResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**AutoMap** | Pointer to **string** |  | [optional] 
<<<<<<< HEAD
**AutoMapNumeric** | Pointer to **int64** |  | [optional] 
**AutoStallRecovery** | Pointer to **string** | Provides detection of when hung controllers are preventing IO from completing | [optional] 
**AutoStallRecoveryNumeric** | Pointer to **int64** | Provides detection of when hung controllers are preventing IO from completing( In numeric form ) | [optional] 
**AutoUnmap** | Pointer to **string** |  | [optional] 
**AutoUnmapNumeric** | Pointer to **int64** |  | [optional] 
**AutoWriteBack** | Pointer to **string** | Data is cached to memory before being written to disk | [optional] 
**AutoWriteBackNumeric** | Pointer to **int64** | Data is cached to memory before being written to disk( In numeric form ) | [optional] 
**BackgroundScrub** | Pointer to **string** | Low priority disk group scrub | [optional] 
**BackgroundScrubInterval** | Pointer to **int64** | Disk group scrub start interval | [optional] 
**BackgroundScrubNumeric** | Pointer to **int64** | Low priority disk group scrub( In numeric form ) | [optional] 
**CacheFlushTimeout** | Pointer to **string** |  | [optional] 
**CacheFlushTimeoutNumeric** | Pointer to **int64** |  | [optional] 
**ControllerFailure** | Pointer to **string** |  | [optional] 
**ControllerFailureNumeric** | Pointer to **int64** |  | [optional] 
**DefaultMapping** | Pointer to **string** |  | [optional] 
**DefaultMappingNumeric** | Pointer to **int64** |  | [optional] 
**DeleteOverride** | Pointer to **string** | Bypass checks when deleting a pool. | [optional] 
**DeleteOverrideNumeric** | Pointer to **int64** | Bypass checks when deleting a pool.( In numeric form ) | [optional] 
**DiskDsdDelay** | Pointer to **int64** |  | [optional] 
**DiskDsdEnable** | Pointer to **string** |  | [optional] 
**DiskDsdEnableNumeric** | Pointer to **int64** |  | [optional] 
**DiskFirmwareUpdate** | Pointer to **string** |  | [optional] 
**DiskFirmwareUpdateNumeric** | Pointer to **int64** |  | [optional] 
**DynamicSpares** | Pointer to **string** | If enabled, replaces failed disks in a degraded vdisk with properly sized available disks | [optional] 
**EmpPollRate** | Pointer to **string** |  | [optional] 
**FanFailure** | Pointer to **string** |  | [optional] 
**FanFailureNumeric** | Pointer to **int64** |  | [optional] 
**HedgedReadsTimeout** | Pointer to **string** |  | [optional] 
**HedgedReadsTimeoutNumeric** | Pointer to **int64** |  | [optional] 
**HostCacheControl** | Pointer to **string** | Whether hosts can change the system&#39;s write-back cache setting | [optional] 
**HostCacheControlNumeric** | Pointer to **int64** | Whether hosts can change the system&#39;s write-back cache setting( In numeric form ) | [optional] 
**IndependentCache** | Pointer to **string** |  | [optional] 
**IndependentCacheNumeric** | Pointer to **int64** |  | [optional] 
**LargePools** | Pointer to **string** |  | [optional] 
**LargePoolsNumeric** | Pointer to **int64** |  | [optional] 
**ManagedLogs** | Pointer to **string** |  | [optional] 
**ManagedLogsNumeric** | Pointer to **int64** |  | [optional] 
**MemoryCardFailure** | Pointer to **string** |  | [optional] 
**MemoryCardFailureNumeric** | Pointer to **int64** |  | [optional] 
**MissingLunResponse** | Pointer to **string** |  | [optional] 
**MissingLunResponseNumeric** | Pointer to **int64** |  | [optional] 
**PartnerFirmwareUpgrade** | Pointer to **string** | Indicates whether automated update of firmware of other controller is enabled | [optional] 
**PartnerFirmwareUpgradeNumeric** | Pointer to **int64** | Indicates whether automated update of firmware of other controller is enabled( In numeric form ) | [optional] 
**PartnerNotify** | Pointer to **string** | Indicates whether the write-through trigger is enabled | [optional] 
**PartnerNotifyNumeric** | Pointer to **int64** | Indicates whether the write-through trigger is enabled( In numeric form ) | [optional] 
**PcieHotplug** | Pointer to **string** |  | [optional] 
**PcieHotplugNumeric** | Pointer to **int64** |  | [optional] 
**PowerSupplyFailure** | Pointer to **string** | Indicates whether the write-through trigger is enabled | [optional] 
**PowerSupplyFailureNumeric** | Pointer to **int64** | Indicates whether the write-through trigger is enabled( In numeric form ) | [optional] 
**RandomIoPerformanceOptimization** | Pointer to **string** |  | [optional] 
**RandomIoPerformanceOptimizationNumeric** | Pointer to **int64** |  | [optional] 
**Remanufacture** | Pointer to **string** |  | [optional] 
**RemanufactureNumeric** | Pointer to **int64** |  | [optional] 
**RestartOnCapiFail** | Pointer to **string** |  | [optional] 
**RestartOnCapiFailNumeric** | Pointer to **int64** |  | [optional] 
**ScrubSchedule** | Pointer to **string** |  | [optional] 
**ScrubScheduleNumeric** | Pointer to **int64** |  | [optional] 
**SingleController** | Pointer to **string** |  | [optional] 
**SingleControllerNumeric** | Pointer to **int64** |  | [optional] 
**SlotAffinity** | Pointer to **string** |  | [optional] 
**SlotAffinityNumeric** | Pointer to **int64** |  | [optional] 
**SlowDiskDetection** | Pointer to **string** |  | [optional] 
**SlowDiskDetectionNumeric** | Pointer to **int64** |  | [optional] 
**Smart** | Pointer to **string** |  | [optional] 
**SmartNumeric** | Pointer to **int64** |  | [optional] 
**SsdConcurrentAccess** | Pointer to **string** |  | [optional] 
**SsdConcurrentAccessNumeric** | Pointer to **int64** |  | [optional] 
**SuperCapFailure** | Pointer to **string** | Indicates whether the write-through trigger is enabled | [optional] 
**SuperCapFailureNumeric** | Pointer to **int64** | Indicates whether the write-through trigger is enabled( In numeric form ) | [optional] 
**SyncCacheMode** | Pointer to **string** |  | [optional] 
**SyncCacheModeNumeric** | Pointer to **int64** |  | [optional] 
**TemperatureExceeded** | Pointer to **string** | Indicates whether the write-through trigger is enabled | [optional] 
**TemperatureExceededNumeric** | Pointer to **int64** | Indicates whether the write-through trigger is enabled( In numeric form ) | [optional] 
**UtilityPriority** | Pointer to **string** | Configured processing priority for utilities | [optional] 
**UtilityPriorityNumeric** | Pointer to **int64** | Configured processing priority for utilities( In numeric form ) | [optional] 
=======
**AutoMapNumeric** | Pointer to **int32** |  | [optional] 
**AutoStallRecovery** | Pointer to **string** | Provides detection of when hung controllers are preventing IO from completing | [optional] 
**AutoStallRecoveryNumeric** | Pointer to **int32** | Provides detection of when hung controllers are preventing IO from completing( In numeric form ) | [optional] 
**AutoUnmap** | Pointer to **string** |  | [optional] 
**AutoUnmapNumeric** | Pointer to **int32** |  | [optional] 
**AutoWriteBack** | Pointer to **string** | Data is cached to memory before being written to disk | [optional] 
**AutoWriteBackNumeric** | Pointer to **int32** | Data is cached to memory before being written to disk( In numeric form ) | [optional] 
**BackgroundScrub** | Pointer to **string** | Low priority disk group scrub | [optional] 
**BackgroundScrubInterval** | Pointer to **int32** | Disk group scrub start interval | [optional] 
**BackgroundScrubNumeric** | Pointer to **int32** | Low priority disk group scrub( In numeric form ) | [optional] 
**CacheFlushTimeout** | Pointer to **string** |  | [optional] 
**CacheFlushTimeoutNumeric** | Pointer to **int32** |  | [optional] 
**ControllerFailure** | Pointer to **string** |  | [optional] 
**ControllerFailureNumeric** | Pointer to **int32** |  | [optional] 
**DefaultMapping** | Pointer to **string** |  | [optional] 
**DefaultMappingNumeric** | Pointer to **int32** |  | [optional] 
**DeleteOverride** | Pointer to **string** | Bypass checks when deleting a pool. | [optional] 
**DeleteOverrideNumeric** | Pointer to **int32** | Bypass checks when deleting a pool.( In numeric form ) | [optional] 
**DiskDsdDelay** | Pointer to **int32** |  | [optional] 
**DiskDsdEnable** | Pointer to **string** |  | [optional] 
**DiskDsdEnableNumeric** | Pointer to **int32** |  | [optional] 
**DiskFirmwareUpdate** | Pointer to **string** |  | [optional] 
**DiskFirmwareUpdateNumeric** | Pointer to **int32** |  | [optional] 
**DynamicSpares** | Pointer to **string** | If enabled, replaces failed disks in a degraded vdisk with properly sized available disks | [optional] 
**EmpPollRate** | Pointer to **string** |  | [optional] 
**FanFailure** | Pointer to **string** |  | [optional] 
**FanFailureNumeric** | Pointer to **int32** |  | [optional] 
**HedgedReadsTimeout** | Pointer to **string** |  | [optional] 
**HedgedReadsTimeoutNumeric** | Pointer to **int32** |  | [optional] 
**HostCacheControl** | Pointer to **string** | Whether hosts can change the system&#39;s write-back cache setting | [optional] 
**HostCacheControlNumeric** | Pointer to **int32** | Whether hosts can change the system&#39;s write-back cache setting( In numeric form ) | [optional] 
**IndependentCache** | Pointer to **string** |  | [optional] 
**IndependentCacheNumeric** | Pointer to **int32** |  | [optional] 
**LargePools** | Pointer to **string** |  | [optional] 
**LargePoolsNumeric** | Pointer to **int32** |  | [optional] 
**ManagedLogs** | Pointer to **string** |  | [optional] 
**ManagedLogsNumeric** | Pointer to **int32** |  | [optional] 
**MemoryCardFailure** | Pointer to **string** |  | [optional] 
**MemoryCardFailureNumeric** | Pointer to **int32** |  | [optional] 
**MissingLunResponse** | Pointer to **string** |  | [optional] 
**MissingLunResponseNumeric** | Pointer to **int32** |  | [optional] 
**PartnerFirmwareUpgrade** | Pointer to **string** | Indicates whether automated update of firmware of other controller is enabled | [optional] 
**PartnerFirmwareUpgradeNumeric** | Pointer to **int32** | Indicates whether automated update of firmware of other controller is enabled( In numeric form ) | [optional] 
**PartnerNotify** | Pointer to **string** | Indicates whether the write-through trigger is enabled | [optional] 
**PartnerNotifyNumeric** | Pointer to **int32** | Indicates whether the write-through trigger is enabled( In numeric form ) | [optional] 
**PcieHotplug** | Pointer to **string** |  | [optional] 
**PcieHotplugNumeric** | Pointer to **int32** |  | [optional] 
**PowerSupplyFailure** | Pointer to **string** | Indicates whether the write-through trigger is enabled | [optional] 
**PowerSupplyFailureNumeric** | Pointer to **int32** | Indicates whether the write-through trigger is enabled( In numeric form ) | [optional] 
**RandomIoPerformanceOptimization** | Pointer to **string** |  | [optional] 
**RandomIoPerformanceOptimizationNumeric** | Pointer to **int32** |  | [optional] 
**Remanufacture** | Pointer to **string** |  | [optional] 
**RemanufactureNumeric** | Pointer to **int32** |  | [optional] 
**RestartOnCapiFail** | Pointer to **string** |  | [optional] 
**RestartOnCapiFailNumeric** | Pointer to **int32** |  | [optional] 
**ScrubSchedule** | Pointer to **string** |  | [optional] 
**ScrubScheduleNumeric** | Pointer to **int32** |  | [optional] 
**SingleController** | Pointer to **string** |  | [optional] 
**SingleControllerNumeric** | Pointer to **int32** |  | [optional] 
**SlotAffinity** | Pointer to **string** |  | [optional] 
**SlotAffinityNumeric** | Pointer to **int32** |  | [optional] 
**SlowDiskDetection** | Pointer to **string** |  | [optional] 
**SlowDiskDetectionNumeric** | Pointer to **int32** |  | [optional] 
**Smart** | Pointer to **string** |  | [optional] 
**SmartNumeric** | Pointer to **int32** |  | [optional] 
**SsdConcurrentAccess** | Pointer to **string** |  | [optional] 
**SsdConcurrentAccessNumeric** | Pointer to **int32** |  | [optional] 
**SuperCapFailure** | Pointer to **string** | Indicates whether the write-through trigger is enabled | [optional] 
**SuperCapFailureNumeric** | Pointer to **int32** | Indicates whether the write-through trigger is enabled( In numeric form ) | [optional] 
**SyncCacheMode** | Pointer to **string** |  | [optional] 
**SyncCacheModeNumeric** | Pointer to **int32** |  | [optional] 
**TemperatureExceeded** | Pointer to **string** | Indicates whether the write-through trigger is enabled | [optional] 
**TemperatureExceededNumeric** | Pointer to **int32** | Indicates whether the write-through trigger is enabled( In numeric form ) | [optional] 
**UtilityPriority** | Pointer to **string** | Configured processing priority for utilities | [optional] 
**UtilityPriorityNumeric** | Pointer to **int32** | Configured processing priority for utilities( In numeric form ) | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

## Methods

### NewAdvancedSettingsTableResourceInner

`func NewAdvancedSettingsTableResourceInner() *AdvancedSettingsTableResourceInner`

NewAdvancedSettingsTableResourceInner instantiates a new AdvancedSettingsTableResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedSettingsTableResourceInnerWithDefaults

`func NewAdvancedSettingsTableResourceInnerWithDefaults() *AdvancedSettingsTableResourceInner`

NewAdvancedSettingsTableResourceInnerWithDefaults instantiates a new AdvancedSettingsTableResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *AdvancedSettingsTableResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *AdvancedSettingsTableResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *AdvancedSettingsTableResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *AdvancedSettingsTableResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *AdvancedSettingsTableResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AdvancedSettingsTableResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AdvancedSettingsTableResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AdvancedSettingsTableResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetAutoMap

`func (o *AdvancedSettingsTableResourceInner) GetAutoMap() string`

GetAutoMap returns the AutoMap field if non-nil, zero value otherwise.

### GetAutoMapOk

`func (o *AdvancedSettingsTableResourceInner) GetAutoMapOk() (*string, bool)`

GetAutoMapOk returns a tuple with the AutoMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMap

`func (o *AdvancedSettingsTableResourceInner) SetAutoMap(v string)`

SetAutoMap sets AutoMap field to given value.

### HasAutoMap

`func (o *AdvancedSettingsTableResourceInner) HasAutoMap() bool`

HasAutoMap returns a boolean if a field has been set.

### GetAutoMapNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetAutoMapNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetAutoMapNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAutoMapNumeric returns the AutoMapNumeric field if non-nil, zero value otherwise.

### GetAutoMapNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetAutoMapNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetAutoMapNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAutoMapNumericOk returns a tuple with the AutoMapNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMapNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetAutoMapNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetAutoMapNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetAutoMapNumeric sets AutoMapNumeric field to given value.

### HasAutoMapNumeric

`func (o *AdvancedSettingsTableResourceInner) HasAutoMapNumeric() bool`

HasAutoMapNumeric returns a boolean if a field has been set.

### GetAutoStallRecovery

`func (o *AdvancedSettingsTableResourceInner) GetAutoStallRecovery() string`

GetAutoStallRecovery returns the AutoStallRecovery field if non-nil, zero value otherwise.

### GetAutoStallRecoveryOk

`func (o *AdvancedSettingsTableResourceInner) GetAutoStallRecoveryOk() (*string, bool)`

GetAutoStallRecoveryOk returns a tuple with the AutoStallRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStallRecovery

`func (o *AdvancedSettingsTableResourceInner) SetAutoStallRecovery(v string)`

SetAutoStallRecovery sets AutoStallRecovery field to given value.

### HasAutoStallRecovery

`func (o *AdvancedSettingsTableResourceInner) HasAutoStallRecovery() bool`

HasAutoStallRecovery returns a boolean if a field has been set.

### GetAutoStallRecoveryNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetAutoStallRecoveryNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetAutoStallRecoveryNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAutoStallRecoveryNumeric returns the AutoStallRecoveryNumeric field if non-nil, zero value otherwise.

### GetAutoStallRecoveryNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetAutoStallRecoveryNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetAutoStallRecoveryNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAutoStallRecoveryNumericOk returns a tuple with the AutoStallRecoveryNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStallRecoveryNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetAutoStallRecoveryNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetAutoStallRecoveryNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetAutoStallRecoveryNumeric sets AutoStallRecoveryNumeric field to given value.

### HasAutoStallRecoveryNumeric

`func (o *AdvancedSettingsTableResourceInner) HasAutoStallRecoveryNumeric() bool`

HasAutoStallRecoveryNumeric returns a boolean if a field has been set.

### GetAutoUnmap

`func (o *AdvancedSettingsTableResourceInner) GetAutoUnmap() string`

GetAutoUnmap returns the AutoUnmap field if non-nil, zero value otherwise.

### GetAutoUnmapOk

`func (o *AdvancedSettingsTableResourceInner) GetAutoUnmapOk() (*string, bool)`

GetAutoUnmapOk returns a tuple with the AutoUnmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUnmap

`func (o *AdvancedSettingsTableResourceInner) SetAutoUnmap(v string)`

SetAutoUnmap sets AutoUnmap field to given value.

### HasAutoUnmap

`func (o *AdvancedSettingsTableResourceInner) HasAutoUnmap() bool`

HasAutoUnmap returns a boolean if a field has been set.

### GetAutoUnmapNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetAutoUnmapNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetAutoUnmapNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAutoUnmapNumeric returns the AutoUnmapNumeric field if non-nil, zero value otherwise.

### GetAutoUnmapNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetAutoUnmapNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetAutoUnmapNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAutoUnmapNumericOk returns a tuple with the AutoUnmapNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUnmapNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetAutoUnmapNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetAutoUnmapNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetAutoUnmapNumeric sets AutoUnmapNumeric field to given value.

### HasAutoUnmapNumeric

`func (o *AdvancedSettingsTableResourceInner) HasAutoUnmapNumeric() bool`

HasAutoUnmapNumeric returns a boolean if a field has been set.

### GetAutoWriteBack

`func (o *AdvancedSettingsTableResourceInner) GetAutoWriteBack() string`

GetAutoWriteBack returns the AutoWriteBack field if non-nil, zero value otherwise.

### GetAutoWriteBackOk

`func (o *AdvancedSettingsTableResourceInner) GetAutoWriteBackOk() (*string, bool)`

GetAutoWriteBackOk returns a tuple with the AutoWriteBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoWriteBack

`func (o *AdvancedSettingsTableResourceInner) SetAutoWriteBack(v string)`

SetAutoWriteBack sets AutoWriteBack field to given value.

### HasAutoWriteBack

`func (o *AdvancedSettingsTableResourceInner) HasAutoWriteBack() bool`

HasAutoWriteBack returns a boolean if a field has been set.

### GetAutoWriteBackNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetAutoWriteBackNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetAutoWriteBackNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAutoWriteBackNumeric returns the AutoWriteBackNumeric field if non-nil, zero value otherwise.

### GetAutoWriteBackNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetAutoWriteBackNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetAutoWriteBackNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetAutoWriteBackNumericOk returns a tuple with the AutoWriteBackNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoWriteBackNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetAutoWriteBackNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetAutoWriteBackNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetAutoWriteBackNumeric sets AutoWriteBackNumeric field to given value.

### HasAutoWriteBackNumeric

`func (o *AdvancedSettingsTableResourceInner) HasAutoWriteBackNumeric() bool`

HasAutoWriteBackNumeric returns a boolean if a field has been set.

### GetBackgroundScrub

`func (o *AdvancedSettingsTableResourceInner) GetBackgroundScrub() string`

GetBackgroundScrub returns the BackgroundScrub field if non-nil, zero value otherwise.

### GetBackgroundScrubOk

`func (o *AdvancedSettingsTableResourceInner) GetBackgroundScrubOk() (*string, bool)`

GetBackgroundScrubOk returns a tuple with the BackgroundScrub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundScrub

`func (o *AdvancedSettingsTableResourceInner) SetBackgroundScrub(v string)`

SetBackgroundScrub sets BackgroundScrub field to given value.

### HasBackgroundScrub

`func (o *AdvancedSettingsTableResourceInner) HasBackgroundScrub() bool`

HasBackgroundScrub returns a boolean if a field has been set.

### GetBackgroundScrubInterval

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetBackgroundScrubInterval() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetBackgroundScrubInterval() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetBackgroundScrubInterval returns the BackgroundScrubInterval field if non-nil, zero value otherwise.

### GetBackgroundScrubIntervalOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetBackgroundScrubIntervalOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetBackgroundScrubIntervalOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetBackgroundScrubIntervalOk returns a tuple with the BackgroundScrubInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundScrubInterval

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetBackgroundScrubInterval(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetBackgroundScrubInterval(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetBackgroundScrubInterval sets BackgroundScrubInterval field to given value.

### HasBackgroundScrubInterval

`func (o *AdvancedSettingsTableResourceInner) HasBackgroundScrubInterval() bool`

HasBackgroundScrubInterval returns a boolean if a field has been set.

### GetBackgroundScrubNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetBackgroundScrubNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetBackgroundScrubNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetBackgroundScrubNumeric returns the BackgroundScrubNumeric field if non-nil, zero value otherwise.

### GetBackgroundScrubNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetBackgroundScrubNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetBackgroundScrubNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetBackgroundScrubNumericOk returns a tuple with the BackgroundScrubNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundScrubNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetBackgroundScrubNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetBackgroundScrubNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetBackgroundScrubNumeric sets BackgroundScrubNumeric field to given value.

### HasBackgroundScrubNumeric

`func (o *AdvancedSettingsTableResourceInner) HasBackgroundScrubNumeric() bool`

HasBackgroundScrubNumeric returns a boolean if a field has been set.

### GetCacheFlushTimeout

`func (o *AdvancedSettingsTableResourceInner) GetCacheFlushTimeout() string`

GetCacheFlushTimeout returns the CacheFlushTimeout field if non-nil, zero value otherwise.

### GetCacheFlushTimeoutOk

`func (o *AdvancedSettingsTableResourceInner) GetCacheFlushTimeoutOk() (*string, bool)`

GetCacheFlushTimeoutOk returns a tuple with the CacheFlushTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheFlushTimeout

`func (o *AdvancedSettingsTableResourceInner) SetCacheFlushTimeout(v string)`

SetCacheFlushTimeout sets CacheFlushTimeout field to given value.

### HasCacheFlushTimeout

`func (o *AdvancedSettingsTableResourceInner) HasCacheFlushTimeout() bool`

HasCacheFlushTimeout returns a boolean if a field has been set.

### GetCacheFlushTimeoutNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetCacheFlushTimeoutNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetCacheFlushTimeoutNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCacheFlushTimeoutNumeric returns the CacheFlushTimeoutNumeric field if non-nil, zero value otherwise.

### GetCacheFlushTimeoutNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetCacheFlushTimeoutNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetCacheFlushTimeoutNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCacheFlushTimeoutNumericOk returns a tuple with the CacheFlushTimeoutNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheFlushTimeoutNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetCacheFlushTimeoutNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetCacheFlushTimeoutNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetCacheFlushTimeoutNumeric sets CacheFlushTimeoutNumeric field to given value.

### HasCacheFlushTimeoutNumeric

`func (o *AdvancedSettingsTableResourceInner) HasCacheFlushTimeoutNumeric() bool`

HasCacheFlushTimeoutNumeric returns a boolean if a field has been set.

### GetControllerFailure

`func (o *AdvancedSettingsTableResourceInner) GetControllerFailure() string`

GetControllerFailure returns the ControllerFailure field if non-nil, zero value otherwise.

### GetControllerFailureOk

`func (o *AdvancedSettingsTableResourceInner) GetControllerFailureOk() (*string, bool)`

GetControllerFailureOk returns a tuple with the ControllerFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerFailure

`func (o *AdvancedSettingsTableResourceInner) SetControllerFailure(v string)`

SetControllerFailure sets ControllerFailure field to given value.

### HasControllerFailure

`func (o *AdvancedSettingsTableResourceInner) HasControllerFailure() bool`

HasControllerFailure returns a boolean if a field has been set.

### GetControllerFailureNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetControllerFailureNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetControllerFailureNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetControllerFailureNumeric returns the ControllerFailureNumeric field if non-nil, zero value otherwise.

### GetControllerFailureNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetControllerFailureNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetControllerFailureNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetControllerFailureNumericOk returns a tuple with the ControllerFailureNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerFailureNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetControllerFailureNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetControllerFailureNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetControllerFailureNumeric sets ControllerFailureNumeric field to given value.

### HasControllerFailureNumeric

`func (o *AdvancedSettingsTableResourceInner) HasControllerFailureNumeric() bool`

HasControllerFailureNumeric returns a boolean if a field has been set.

### GetDefaultMapping

`func (o *AdvancedSettingsTableResourceInner) GetDefaultMapping() string`

GetDefaultMapping returns the DefaultMapping field if non-nil, zero value otherwise.

### GetDefaultMappingOk

`func (o *AdvancedSettingsTableResourceInner) GetDefaultMappingOk() (*string, bool)`

GetDefaultMappingOk returns a tuple with the DefaultMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMapping

`func (o *AdvancedSettingsTableResourceInner) SetDefaultMapping(v string)`

SetDefaultMapping sets DefaultMapping field to given value.

### HasDefaultMapping

`func (o *AdvancedSettingsTableResourceInner) HasDefaultMapping() bool`

HasDefaultMapping returns a boolean if a field has been set.

### GetDefaultMappingNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetDefaultMappingNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetDefaultMappingNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDefaultMappingNumeric returns the DefaultMappingNumeric field if non-nil, zero value otherwise.

### GetDefaultMappingNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetDefaultMappingNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetDefaultMappingNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDefaultMappingNumericOk returns a tuple with the DefaultMappingNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMappingNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetDefaultMappingNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetDefaultMappingNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetDefaultMappingNumeric sets DefaultMappingNumeric field to given value.

### HasDefaultMappingNumeric

`func (o *AdvancedSettingsTableResourceInner) HasDefaultMappingNumeric() bool`

HasDefaultMappingNumeric returns a boolean if a field has been set.

### GetDeleteOverride

`func (o *AdvancedSettingsTableResourceInner) GetDeleteOverride() string`

GetDeleteOverride returns the DeleteOverride field if non-nil, zero value otherwise.

### GetDeleteOverrideOk

`func (o *AdvancedSettingsTableResourceInner) GetDeleteOverrideOk() (*string, bool)`

GetDeleteOverrideOk returns a tuple with the DeleteOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOverride

`func (o *AdvancedSettingsTableResourceInner) SetDeleteOverride(v string)`

SetDeleteOverride sets DeleteOverride field to given value.

### HasDeleteOverride

`func (o *AdvancedSettingsTableResourceInner) HasDeleteOverride() bool`

HasDeleteOverride returns a boolean if a field has been set.

### GetDeleteOverrideNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetDeleteOverrideNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetDeleteOverrideNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDeleteOverrideNumeric returns the DeleteOverrideNumeric field if non-nil, zero value otherwise.

### GetDeleteOverrideNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetDeleteOverrideNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetDeleteOverrideNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDeleteOverrideNumericOk returns a tuple with the DeleteOverrideNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOverrideNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetDeleteOverrideNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetDeleteOverrideNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetDeleteOverrideNumeric sets DeleteOverrideNumeric field to given value.

### HasDeleteOverrideNumeric

`func (o *AdvancedSettingsTableResourceInner) HasDeleteOverrideNumeric() bool`

HasDeleteOverrideNumeric returns a boolean if a field has been set.

### GetDiskDsdDelay

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetDiskDsdDelay() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetDiskDsdDelay() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDiskDsdDelay returns the DiskDsdDelay field if non-nil, zero value otherwise.

### GetDiskDsdDelayOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetDiskDsdDelayOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetDiskDsdDelayOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDiskDsdDelayOk returns a tuple with the DiskDsdDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDsdDelay

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetDiskDsdDelay(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetDiskDsdDelay(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetDiskDsdDelay sets DiskDsdDelay field to given value.

### HasDiskDsdDelay

`func (o *AdvancedSettingsTableResourceInner) HasDiskDsdDelay() bool`

HasDiskDsdDelay returns a boolean if a field has been set.

### GetDiskDsdEnable

`func (o *AdvancedSettingsTableResourceInner) GetDiskDsdEnable() string`

GetDiskDsdEnable returns the DiskDsdEnable field if non-nil, zero value otherwise.

### GetDiskDsdEnableOk

`func (o *AdvancedSettingsTableResourceInner) GetDiskDsdEnableOk() (*string, bool)`

GetDiskDsdEnableOk returns a tuple with the DiskDsdEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDsdEnable

`func (o *AdvancedSettingsTableResourceInner) SetDiskDsdEnable(v string)`

SetDiskDsdEnable sets DiskDsdEnable field to given value.

### HasDiskDsdEnable

`func (o *AdvancedSettingsTableResourceInner) HasDiskDsdEnable() bool`

HasDiskDsdEnable returns a boolean if a field has been set.

### GetDiskDsdEnableNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetDiskDsdEnableNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetDiskDsdEnableNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDiskDsdEnableNumeric returns the DiskDsdEnableNumeric field if non-nil, zero value otherwise.

### GetDiskDsdEnableNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetDiskDsdEnableNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetDiskDsdEnableNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDiskDsdEnableNumericOk returns a tuple with the DiskDsdEnableNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDsdEnableNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetDiskDsdEnableNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetDiskDsdEnableNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetDiskDsdEnableNumeric sets DiskDsdEnableNumeric field to given value.

### HasDiskDsdEnableNumeric

`func (o *AdvancedSettingsTableResourceInner) HasDiskDsdEnableNumeric() bool`

HasDiskDsdEnableNumeric returns a boolean if a field has been set.

### GetDiskFirmwareUpdate

`func (o *AdvancedSettingsTableResourceInner) GetDiskFirmwareUpdate() string`

GetDiskFirmwareUpdate returns the DiskFirmwareUpdate field if non-nil, zero value otherwise.

### GetDiskFirmwareUpdateOk

`func (o *AdvancedSettingsTableResourceInner) GetDiskFirmwareUpdateOk() (*string, bool)`

GetDiskFirmwareUpdateOk returns a tuple with the DiskFirmwareUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFirmwareUpdate

`func (o *AdvancedSettingsTableResourceInner) SetDiskFirmwareUpdate(v string)`

SetDiskFirmwareUpdate sets DiskFirmwareUpdate field to given value.

### HasDiskFirmwareUpdate

`func (o *AdvancedSettingsTableResourceInner) HasDiskFirmwareUpdate() bool`

HasDiskFirmwareUpdate returns a boolean if a field has been set.

### GetDiskFirmwareUpdateNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetDiskFirmwareUpdateNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetDiskFirmwareUpdateNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDiskFirmwareUpdateNumeric returns the DiskFirmwareUpdateNumeric field if non-nil, zero value otherwise.

### GetDiskFirmwareUpdateNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetDiskFirmwareUpdateNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetDiskFirmwareUpdateNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetDiskFirmwareUpdateNumericOk returns a tuple with the DiskFirmwareUpdateNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFirmwareUpdateNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetDiskFirmwareUpdateNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetDiskFirmwareUpdateNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetDiskFirmwareUpdateNumeric sets DiskFirmwareUpdateNumeric field to given value.

### HasDiskFirmwareUpdateNumeric

`func (o *AdvancedSettingsTableResourceInner) HasDiskFirmwareUpdateNumeric() bool`

HasDiskFirmwareUpdateNumeric returns a boolean if a field has been set.

### GetDynamicSpares

`func (o *AdvancedSettingsTableResourceInner) GetDynamicSpares() string`

GetDynamicSpares returns the DynamicSpares field if non-nil, zero value otherwise.

### GetDynamicSparesOk

`func (o *AdvancedSettingsTableResourceInner) GetDynamicSparesOk() (*string, bool)`

GetDynamicSparesOk returns a tuple with the DynamicSpares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSpares

`func (o *AdvancedSettingsTableResourceInner) SetDynamicSpares(v string)`

SetDynamicSpares sets DynamicSpares field to given value.

### HasDynamicSpares

`func (o *AdvancedSettingsTableResourceInner) HasDynamicSpares() bool`

HasDynamicSpares returns a boolean if a field has been set.

### GetEmpPollRate

`func (o *AdvancedSettingsTableResourceInner) GetEmpPollRate() string`

GetEmpPollRate returns the EmpPollRate field if non-nil, zero value otherwise.

### GetEmpPollRateOk

`func (o *AdvancedSettingsTableResourceInner) GetEmpPollRateOk() (*string, bool)`

GetEmpPollRateOk returns a tuple with the EmpPollRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpPollRate

`func (o *AdvancedSettingsTableResourceInner) SetEmpPollRate(v string)`

SetEmpPollRate sets EmpPollRate field to given value.

### HasEmpPollRate

`func (o *AdvancedSettingsTableResourceInner) HasEmpPollRate() bool`

HasEmpPollRate returns a boolean if a field has been set.

### GetFanFailure

`func (o *AdvancedSettingsTableResourceInner) GetFanFailure() string`

GetFanFailure returns the FanFailure field if non-nil, zero value otherwise.

### GetFanFailureOk

`func (o *AdvancedSettingsTableResourceInner) GetFanFailureOk() (*string, bool)`

GetFanFailureOk returns a tuple with the FanFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanFailure

`func (o *AdvancedSettingsTableResourceInner) SetFanFailure(v string)`

SetFanFailure sets FanFailure field to given value.

### HasFanFailure

`func (o *AdvancedSettingsTableResourceInner) HasFanFailure() bool`

HasFanFailure returns a boolean if a field has been set.

### GetFanFailureNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetFanFailureNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetFanFailureNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetFanFailureNumeric returns the FanFailureNumeric field if non-nil, zero value otherwise.

### GetFanFailureNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetFanFailureNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetFanFailureNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetFanFailureNumericOk returns a tuple with the FanFailureNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanFailureNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetFanFailureNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetFanFailureNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetFanFailureNumeric sets FanFailureNumeric field to given value.

### HasFanFailureNumeric

`func (o *AdvancedSettingsTableResourceInner) HasFanFailureNumeric() bool`

HasFanFailureNumeric returns a boolean if a field has been set.

### GetHedgedReadsTimeout

`func (o *AdvancedSettingsTableResourceInner) GetHedgedReadsTimeout() string`

GetHedgedReadsTimeout returns the HedgedReadsTimeout field if non-nil, zero value otherwise.

### GetHedgedReadsTimeoutOk

`func (o *AdvancedSettingsTableResourceInner) GetHedgedReadsTimeoutOk() (*string, bool)`

GetHedgedReadsTimeoutOk returns a tuple with the HedgedReadsTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHedgedReadsTimeout

`func (o *AdvancedSettingsTableResourceInner) SetHedgedReadsTimeout(v string)`

SetHedgedReadsTimeout sets HedgedReadsTimeout field to given value.

### HasHedgedReadsTimeout

`func (o *AdvancedSettingsTableResourceInner) HasHedgedReadsTimeout() bool`

HasHedgedReadsTimeout returns a boolean if a field has been set.

### GetHedgedReadsTimeoutNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetHedgedReadsTimeoutNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetHedgedReadsTimeoutNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHedgedReadsTimeoutNumeric returns the HedgedReadsTimeoutNumeric field if non-nil, zero value otherwise.

### GetHedgedReadsTimeoutNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetHedgedReadsTimeoutNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetHedgedReadsTimeoutNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHedgedReadsTimeoutNumericOk returns a tuple with the HedgedReadsTimeoutNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHedgedReadsTimeoutNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetHedgedReadsTimeoutNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetHedgedReadsTimeoutNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetHedgedReadsTimeoutNumeric sets HedgedReadsTimeoutNumeric field to given value.

### HasHedgedReadsTimeoutNumeric

`func (o *AdvancedSettingsTableResourceInner) HasHedgedReadsTimeoutNumeric() bool`

HasHedgedReadsTimeoutNumeric returns a boolean if a field has been set.

### GetHostCacheControl

`func (o *AdvancedSettingsTableResourceInner) GetHostCacheControl() string`

GetHostCacheControl returns the HostCacheControl field if non-nil, zero value otherwise.

### GetHostCacheControlOk

`func (o *AdvancedSettingsTableResourceInner) GetHostCacheControlOk() (*string, bool)`

GetHostCacheControlOk returns a tuple with the HostCacheControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCacheControl

`func (o *AdvancedSettingsTableResourceInner) SetHostCacheControl(v string)`

SetHostCacheControl sets HostCacheControl field to given value.

### HasHostCacheControl

`func (o *AdvancedSettingsTableResourceInner) HasHostCacheControl() bool`

HasHostCacheControl returns a boolean if a field has been set.

### GetHostCacheControlNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetHostCacheControlNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetHostCacheControlNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHostCacheControlNumeric returns the HostCacheControlNumeric field if non-nil, zero value otherwise.

### GetHostCacheControlNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetHostCacheControlNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetHostCacheControlNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetHostCacheControlNumericOk returns a tuple with the HostCacheControlNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCacheControlNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetHostCacheControlNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetHostCacheControlNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetHostCacheControlNumeric sets HostCacheControlNumeric field to given value.

### HasHostCacheControlNumeric

`func (o *AdvancedSettingsTableResourceInner) HasHostCacheControlNumeric() bool`

HasHostCacheControlNumeric returns a boolean if a field has been set.

### GetIndependentCache

`func (o *AdvancedSettingsTableResourceInner) GetIndependentCache() string`

GetIndependentCache returns the IndependentCache field if non-nil, zero value otherwise.

### GetIndependentCacheOk

`func (o *AdvancedSettingsTableResourceInner) GetIndependentCacheOk() (*string, bool)`

GetIndependentCacheOk returns a tuple with the IndependentCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndependentCache

`func (o *AdvancedSettingsTableResourceInner) SetIndependentCache(v string)`

SetIndependentCache sets IndependentCache field to given value.

### HasIndependentCache

`func (o *AdvancedSettingsTableResourceInner) HasIndependentCache() bool`

HasIndependentCache returns a boolean if a field has been set.

### GetIndependentCacheNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetIndependentCacheNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetIndependentCacheNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetIndependentCacheNumeric returns the IndependentCacheNumeric field if non-nil, zero value otherwise.

### GetIndependentCacheNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetIndependentCacheNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetIndependentCacheNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetIndependentCacheNumericOk returns a tuple with the IndependentCacheNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndependentCacheNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetIndependentCacheNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetIndependentCacheNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetIndependentCacheNumeric sets IndependentCacheNumeric field to given value.

### HasIndependentCacheNumeric

`func (o *AdvancedSettingsTableResourceInner) HasIndependentCacheNumeric() bool`

HasIndependentCacheNumeric returns a boolean if a field has been set.

### GetLargePools

`func (o *AdvancedSettingsTableResourceInner) GetLargePools() string`

GetLargePools returns the LargePools field if non-nil, zero value otherwise.

### GetLargePoolsOk

`func (o *AdvancedSettingsTableResourceInner) GetLargePoolsOk() (*string, bool)`

GetLargePoolsOk returns a tuple with the LargePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargePools

`func (o *AdvancedSettingsTableResourceInner) SetLargePools(v string)`

SetLargePools sets LargePools field to given value.

### HasLargePools

`func (o *AdvancedSettingsTableResourceInner) HasLargePools() bool`

HasLargePools returns a boolean if a field has been set.

### GetLargePoolsNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetLargePoolsNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetLargePoolsNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLargePoolsNumeric returns the LargePoolsNumeric field if non-nil, zero value otherwise.

### GetLargePoolsNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetLargePoolsNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetLargePoolsNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetLargePoolsNumericOk returns a tuple with the LargePoolsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargePoolsNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetLargePoolsNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetLargePoolsNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetLargePoolsNumeric sets LargePoolsNumeric field to given value.

### HasLargePoolsNumeric

`func (o *AdvancedSettingsTableResourceInner) HasLargePoolsNumeric() bool`

HasLargePoolsNumeric returns a boolean if a field has been set.

### GetManagedLogs

`func (o *AdvancedSettingsTableResourceInner) GetManagedLogs() string`

GetManagedLogs returns the ManagedLogs field if non-nil, zero value otherwise.

### GetManagedLogsOk

`func (o *AdvancedSettingsTableResourceInner) GetManagedLogsOk() (*string, bool)`

GetManagedLogsOk returns a tuple with the ManagedLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedLogs

`func (o *AdvancedSettingsTableResourceInner) SetManagedLogs(v string)`

SetManagedLogs sets ManagedLogs field to given value.

### HasManagedLogs

`func (o *AdvancedSettingsTableResourceInner) HasManagedLogs() bool`

HasManagedLogs returns a boolean if a field has been set.

### GetManagedLogsNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetManagedLogsNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetManagedLogsNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetManagedLogsNumeric returns the ManagedLogsNumeric field if non-nil, zero value otherwise.

### GetManagedLogsNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetManagedLogsNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetManagedLogsNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetManagedLogsNumericOk returns a tuple with the ManagedLogsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedLogsNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetManagedLogsNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetManagedLogsNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetManagedLogsNumeric sets ManagedLogsNumeric field to given value.

### HasManagedLogsNumeric

`func (o *AdvancedSettingsTableResourceInner) HasManagedLogsNumeric() bool`

HasManagedLogsNumeric returns a boolean if a field has been set.

### GetMemoryCardFailure

`func (o *AdvancedSettingsTableResourceInner) GetMemoryCardFailure() string`

GetMemoryCardFailure returns the MemoryCardFailure field if non-nil, zero value otherwise.

### GetMemoryCardFailureOk

`func (o *AdvancedSettingsTableResourceInner) GetMemoryCardFailureOk() (*string, bool)`

GetMemoryCardFailureOk returns a tuple with the MemoryCardFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCardFailure

`func (o *AdvancedSettingsTableResourceInner) SetMemoryCardFailure(v string)`

SetMemoryCardFailure sets MemoryCardFailure field to given value.

### HasMemoryCardFailure

`func (o *AdvancedSettingsTableResourceInner) HasMemoryCardFailure() bool`

HasMemoryCardFailure returns a boolean if a field has been set.

### GetMemoryCardFailureNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetMemoryCardFailureNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetMemoryCardFailureNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMemoryCardFailureNumeric returns the MemoryCardFailureNumeric field if non-nil, zero value otherwise.

### GetMemoryCardFailureNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetMemoryCardFailureNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetMemoryCardFailureNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMemoryCardFailureNumericOk returns a tuple with the MemoryCardFailureNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCardFailureNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetMemoryCardFailureNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetMemoryCardFailureNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetMemoryCardFailureNumeric sets MemoryCardFailureNumeric field to given value.

### HasMemoryCardFailureNumeric

`func (o *AdvancedSettingsTableResourceInner) HasMemoryCardFailureNumeric() bool`

HasMemoryCardFailureNumeric returns a boolean if a field has been set.

### GetMissingLunResponse

`func (o *AdvancedSettingsTableResourceInner) GetMissingLunResponse() string`

GetMissingLunResponse returns the MissingLunResponse field if non-nil, zero value otherwise.

### GetMissingLunResponseOk

`func (o *AdvancedSettingsTableResourceInner) GetMissingLunResponseOk() (*string, bool)`

GetMissingLunResponseOk returns a tuple with the MissingLunResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingLunResponse

`func (o *AdvancedSettingsTableResourceInner) SetMissingLunResponse(v string)`

SetMissingLunResponse sets MissingLunResponse field to given value.

### HasMissingLunResponse

`func (o *AdvancedSettingsTableResourceInner) HasMissingLunResponse() bool`

HasMissingLunResponse returns a boolean if a field has been set.

### GetMissingLunResponseNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetMissingLunResponseNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetMissingLunResponseNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMissingLunResponseNumeric returns the MissingLunResponseNumeric field if non-nil, zero value otherwise.

### GetMissingLunResponseNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetMissingLunResponseNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetMissingLunResponseNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetMissingLunResponseNumericOk returns a tuple with the MissingLunResponseNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingLunResponseNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetMissingLunResponseNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetMissingLunResponseNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetMissingLunResponseNumeric sets MissingLunResponseNumeric field to given value.

### HasMissingLunResponseNumeric

`func (o *AdvancedSettingsTableResourceInner) HasMissingLunResponseNumeric() bool`

HasMissingLunResponseNumeric returns a boolean if a field has been set.

### GetPartnerFirmwareUpgrade

`func (o *AdvancedSettingsTableResourceInner) GetPartnerFirmwareUpgrade() string`

GetPartnerFirmwareUpgrade returns the PartnerFirmwareUpgrade field if non-nil, zero value otherwise.

### GetPartnerFirmwareUpgradeOk

`func (o *AdvancedSettingsTableResourceInner) GetPartnerFirmwareUpgradeOk() (*string, bool)`

GetPartnerFirmwareUpgradeOk returns a tuple with the PartnerFirmwareUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerFirmwareUpgrade

`func (o *AdvancedSettingsTableResourceInner) SetPartnerFirmwareUpgrade(v string)`

SetPartnerFirmwareUpgrade sets PartnerFirmwareUpgrade field to given value.

### HasPartnerFirmwareUpgrade

`func (o *AdvancedSettingsTableResourceInner) HasPartnerFirmwareUpgrade() bool`

HasPartnerFirmwareUpgrade returns a boolean if a field has been set.

### GetPartnerFirmwareUpgradeNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetPartnerFirmwareUpgradeNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetPartnerFirmwareUpgradeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPartnerFirmwareUpgradeNumeric returns the PartnerFirmwareUpgradeNumeric field if non-nil, zero value otherwise.

### GetPartnerFirmwareUpgradeNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetPartnerFirmwareUpgradeNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetPartnerFirmwareUpgradeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPartnerFirmwareUpgradeNumericOk returns a tuple with the PartnerFirmwareUpgradeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerFirmwareUpgradeNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetPartnerFirmwareUpgradeNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetPartnerFirmwareUpgradeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetPartnerFirmwareUpgradeNumeric sets PartnerFirmwareUpgradeNumeric field to given value.

### HasPartnerFirmwareUpgradeNumeric

`func (o *AdvancedSettingsTableResourceInner) HasPartnerFirmwareUpgradeNumeric() bool`

HasPartnerFirmwareUpgradeNumeric returns a boolean if a field has been set.

### GetPartnerNotify

`func (o *AdvancedSettingsTableResourceInner) GetPartnerNotify() string`

GetPartnerNotify returns the PartnerNotify field if non-nil, zero value otherwise.

### GetPartnerNotifyOk

`func (o *AdvancedSettingsTableResourceInner) GetPartnerNotifyOk() (*string, bool)`

GetPartnerNotifyOk returns a tuple with the PartnerNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerNotify

`func (o *AdvancedSettingsTableResourceInner) SetPartnerNotify(v string)`

SetPartnerNotify sets PartnerNotify field to given value.

### HasPartnerNotify

`func (o *AdvancedSettingsTableResourceInner) HasPartnerNotify() bool`

HasPartnerNotify returns a boolean if a field has been set.

### GetPartnerNotifyNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetPartnerNotifyNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetPartnerNotifyNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPartnerNotifyNumeric returns the PartnerNotifyNumeric field if non-nil, zero value otherwise.

### GetPartnerNotifyNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetPartnerNotifyNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetPartnerNotifyNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPartnerNotifyNumericOk returns a tuple with the PartnerNotifyNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerNotifyNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetPartnerNotifyNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetPartnerNotifyNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetPartnerNotifyNumeric sets PartnerNotifyNumeric field to given value.

### HasPartnerNotifyNumeric

`func (o *AdvancedSettingsTableResourceInner) HasPartnerNotifyNumeric() bool`

HasPartnerNotifyNumeric returns a boolean if a field has been set.

### GetPcieHotplug

`func (o *AdvancedSettingsTableResourceInner) GetPcieHotplug() string`

GetPcieHotplug returns the PcieHotplug field if non-nil, zero value otherwise.

### GetPcieHotplugOk

`func (o *AdvancedSettingsTableResourceInner) GetPcieHotplugOk() (*string, bool)`

GetPcieHotplugOk returns a tuple with the PcieHotplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieHotplug

`func (o *AdvancedSettingsTableResourceInner) SetPcieHotplug(v string)`

SetPcieHotplug sets PcieHotplug field to given value.

### HasPcieHotplug

`func (o *AdvancedSettingsTableResourceInner) HasPcieHotplug() bool`

HasPcieHotplug returns a boolean if a field has been set.

### GetPcieHotplugNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetPcieHotplugNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetPcieHotplugNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPcieHotplugNumeric returns the PcieHotplugNumeric field if non-nil, zero value otherwise.

### GetPcieHotplugNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetPcieHotplugNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetPcieHotplugNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPcieHotplugNumericOk returns a tuple with the PcieHotplugNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieHotplugNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetPcieHotplugNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetPcieHotplugNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetPcieHotplugNumeric sets PcieHotplugNumeric field to given value.

### HasPcieHotplugNumeric

`func (o *AdvancedSettingsTableResourceInner) HasPcieHotplugNumeric() bool`

HasPcieHotplugNumeric returns a boolean if a field has been set.

### GetPowerSupplyFailure

`func (o *AdvancedSettingsTableResourceInner) GetPowerSupplyFailure() string`

GetPowerSupplyFailure returns the PowerSupplyFailure field if non-nil, zero value otherwise.

### GetPowerSupplyFailureOk

`func (o *AdvancedSettingsTableResourceInner) GetPowerSupplyFailureOk() (*string, bool)`

GetPowerSupplyFailureOk returns a tuple with the PowerSupplyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSupplyFailure

`func (o *AdvancedSettingsTableResourceInner) SetPowerSupplyFailure(v string)`

SetPowerSupplyFailure sets PowerSupplyFailure field to given value.

### HasPowerSupplyFailure

`func (o *AdvancedSettingsTableResourceInner) HasPowerSupplyFailure() bool`

HasPowerSupplyFailure returns a boolean if a field has been set.

### GetPowerSupplyFailureNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetPowerSupplyFailureNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetPowerSupplyFailureNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPowerSupplyFailureNumeric returns the PowerSupplyFailureNumeric field if non-nil, zero value otherwise.

### GetPowerSupplyFailureNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetPowerSupplyFailureNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetPowerSupplyFailureNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetPowerSupplyFailureNumericOk returns a tuple with the PowerSupplyFailureNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSupplyFailureNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetPowerSupplyFailureNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetPowerSupplyFailureNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetPowerSupplyFailureNumeric sets PowerSupplyFailureNumeric field to given value.

### HasPowerSupplyFailureNumeric

`func (o *AdvancedSettingsTableResourceInner) HasPowerSupplyFailureNumeric() bool`

HasPowerSupplyFailureNumeric returns a boolean if a field has been set.

### GetRandomIoPerformanceOptimization

`func (o *AdvancedSettingsTableResourceInner) GetRandomIoPerformanceOptimization() string`

GetRandomIoPerformanceOptimization returns the RandomIoPerformanceOptimization field if non-nil, zero value otherwise.

### GetRandomIoPerformanceOptimizationOk

`func (o *AdvancedSettingsTableResourceInner) GetRandomIoPerformanceOptimizationOk() (*string, bool)`

GetRandomIoPerformanceOptimizationOk returns a tuple with the RandomIoPerformanceOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomIoPerformanceOptimization

`func (o *AdvancedSettingsTableResourceInner) SetRandomIoPerformanceOptimization(v string)`

SetRandomIoPerformanceOptimization sets RandomIoPerformanceOptimization field to given value.

### HasRandomIoPerformanceOptimization

`func (o *AdvancedSettingsTableResourceInner) HasRandomIoPerformanceOptimization() bool`

HasRandomIoPerformanceOptimization returns a boolean if a field has been set.

### GetRandomIoPerformanceOptimizationNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetRandomIoPerformanceOptimizationNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetRandomIoPerformanceOptimizationNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRandomIoPerformanceOptimizationNumeric returns the RandomIoPerformanceOptimizationNumeric field if non-nil, zero value otherwise.

### GetRandomIoPerformanceOptimizationNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetRandomIoPerformanceOptimizationNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetRandomIoPerformanceOptimizationNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRandomIoPerformanceOptimizationNumericOk returns a tuple with the RandomIoPerformanceOptimizationNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomIoPerformanceOptimizationNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetRandomIoPerformanceOptimizationNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetRandomIoPerformanceOptimizationNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetRandomIoPerformanceOptimizationNumeric sets RandomIoPerformanceOptimizationNumeric field to given value.

### HasRandomIoPerformanceOptimizationNumeric

`func (o *AdvancedSettingsTableResourceInner) HasRandomIoPerformanceOptimizationNumeric() bool`

HasRandomIoPerformanceOptimizationNumeric returns a boolean if a field has been set.

### GetRemanufacture

`func (o *AdvancedSettingsTableResourceInner) GetRemanufacture() string`

GetRemanufacture returns the Remanufacture field if non-nil, zero value otherwise.

### GetRemanufactureOk

`func (o *AdvancedSettingsTableResourceInner) GetRemanufactureOk() (*string, bool)`

GetRemanufactureOk returns a tuple with the Remanufacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemanufacture

`func (o *AdvancedSettingsTableResourceInner) SetRemanufacture(v string)`

SetRemanufacture sets Remanufacture field to given value.

### HasRemanufacture

`func (o *AdvancedSettingsTableResourceInner) HasRemanufacture() bool`

HasRemanufacture returns a boolean if a field has been set.

### GetRemanufactureNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetRemanufactureNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetRemanufactureNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRemanufactureNumeric returns the RemanufactureNumeric field if non-nil, zero value otherwise.

### GetRemanufactureNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetRemanufactureNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetRemanufactureNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRemanufactureNumericOk returns a tuple with the RemanufactureNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemanufactureNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetRemanufactureNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetRemanufactureNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetRemanufactureNumeric sets RemanufactureNumeric field to given value.

### HasRemanufactureNumeric

`func (o *AdvancedSettingsTableResourceInner) HasRemanufactureNumeric() bool`

HasRemanufactureNumeric returns a boolean if a field has been set.

### GetRestartOnCapiFail

`func (o *AdvancedSettingsTableResourceInner) GetRestartOnCapiFail() string`

GetRestartOnCapiFail returns the RestartOnCapiFail field if non-nil, zero value otherwise.

### GetRestartOnCapiFailOk

`func (o *AdvancedSettingsTableResourceInner) GetRestartOnCapiFailOk() (*string, bool)`

GetRestartOnCapiFailOk returns a tuple with the RestartOnCapiFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartOnCapiFail

`func (o *AdvancedSettingsTableResourceInner) SetRestartOnCapiFail(v string)`

SetRestartOnCapiFail sets RestartOnCapiFail field to given value.

### HasRestartOnCapiFail

`func (o *AdvancedSettingsTableResourceInner) HasRestartOnCapiFail() bool`

HasRestartOnCapiFail returns a boolean if a field has been set.

### GetRestartOnCapiFailNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetRestartOnCapiFailNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetRestartOnCapiFailNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRestartOnCapiFailNumeric returns the RestartOnCapiFailNumeric field if non-nil, zero value otherwise.

### GetRestartOnCapiFailNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetRestartOnCapiFailNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetRestartOnCapiFailNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetRestartOnCapiFailNumericOk returns a tuple with the RestartOnCapiFailNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartOnCapiFailNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetRestartOnCapiFailNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetRestartOnCapiFailNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetRestartOnCapiFailNumeric sets RestartOnCapiFailNumeric field to given value.

### HasRestartOnCapiFailNumeric

`func (o *AdvancedSettingsTableResourceInner) HasRestartOnCapiFailNumeric() bool`

HasRestartOnCapiFailNumeric returns a boolean if a field has been set.

### GetScrubSchedule

`func (o *AdvancedSettingsTableResourceInner) GetScrubSchedule() string`

GetScrubSchedule returns the ScrubSchedule field if non-nil, zero value otherwise.

### GetScrubScheduleOk

`func (o *AdvancedSettingsTableResourceInner) GetScrubScheduleOk() (*string, bool)`

GetScrubScheduleOk returns a tuple with the ScrubSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrubSchedule

`func (o *AdvancedSettingsTableResourceInner) SetScrubSchedule(v string)`

SetScrubSchedule sets ScrubSchedule field to given value.

### HasScrubSchedule

`func (o *AdvancedSettingsTableResourceInner) HasScrubSchedule() bool`

HasScrubSchedule returns a boolean if a field has been set.

### GetScrubScheduleNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetScrubScheduleNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetScrubScheduleNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetScrubScheduleNumeric returns the ScrubScheduleNumeric field if non-nil, zero value otherwise.

### GetScrubScheduleNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetScrubScheduleNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetScrubScheduleNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetScrubScheduleNumericOk returns a tuple with the ScrubScheduleNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrubScheduleNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetScrubScheduleNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetScrubScheduleNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetScrubScheduleNumeric sets ScrubScheduleNumeric field to given value.

### HasScrubScheduleNumeric

`func (o *AdvancedSettingsTableResourceInner) HasScrubScheduleNumeric() bool`

HasScrubScheduleNumeric returns a boolean if a field has been set.

### GetSingleController

`func (o *AdvancedSettingsTableResourceInner) GetSingleController() string`

GetSingleController returns the SingleController field if non-nil, zero value otherwise.

### GetSingleControllerOk

`func (o *AdvancedSettingsTableResourceInner) GetSingleControllerOk() (*string, bool)`

GetSingleControllerOk returns a tuple with the SingleController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleController

`func (o *AdvancedSettingsTableResourceInner) SetSingleController(v string)`

SetSingleController sets SingleController field to given value.

### HasSingleController

`func (o *AdvancedSettingsTableResourceInner) HasSingleController() bool`

HasSingleController returns a boolean if a field has been set.

### GetSingleControllerNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetSingleControllerNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetSingleControllerNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSingleControllerNumeric returns the SingleControllerNumeric field if non-nil, zero value otherwise.

### GetSingleControllerNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetSingleControllerNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetSingleControllerNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSingleControllerNumericOk returns a tuple with the SingleControllerNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleControllerNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetSingleControllerNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetSingleControllerNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSingleControllerNumeric sets SingleControllerNumeric field to given value.

### HasSingleControllerNumeric

`func (o *AdvancedSettingsTableResourceInner) HasSingleControllerNumeric() bool`

HasSingleControllerNumeric returns a boolean if a field has been set.

### GetSlotAffinity

`func (o *AdvancedSettingsTableResourceInner) GetSlotAffinity() string`

GetSlotAffinity returns the SlotAffinity field if non-nil, zero value otherwise.

### GetSlotAffinityOk

`func (o *AdvancedSettingsTableResourceInner) GetSlotAffinityOk() (*string, bool)`

GetSlotAffinityOk returns a tuple with the SlotAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotAffinity

`func (o *AdvancedSettingsTableResourceInner) SetSlotAffinity(v string)`

SetSlotAffinity sets SlotAffinity field to given value.

### HasSlotAffinity

`func (o *AdvancedSettingsTableResourceInner) HasSlotAffinity() bool`

HasSlotAffinity returns a boolean if a field has been set.

### GetSlotAffinityNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetSlotAffinityNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetSlotAffinityNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSlotAffinityNumeric returns the SlotAffinityNumeric field if non-nil, zero value otherwise.

### GetSlotAffinityNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetSlotAffinityNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetSlotAffinityNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSlotAffinityNumericOk returns a tuple with the SlotAffinityNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotAffinityNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetSlotAffinityNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetSlotAffinityNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSlotAffinityNumeric sets SlotAffinityNumeric field to given value.

### HasSlotAffinityNumeric

`func (o *AdvancedSettingsTableResourceInner) HasSlotAffinityNumeric() bool`

HasSlotAffinityNumeric returns a boolean if a field has been set.

### GetSlowDiskDetection

`func (o *AdvancedSettingsTableResourceInner) GetSlowDiskDetection() string`

GetSlowDiskDetection returns the SlowDiskDetection field if non-nil, zero value otherwise.

### GetSlowDiskDetectionOk

`func (o *AdvancedSettingsTableResourceInner) GetSlowDiskDetectionOk() (*string, bool)`

GetSlowDiskDetectionOk returns a tuple with the SlowDiskDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowDiskDetection

`func (o *AdvancedSettingsTableResourceInner) SetSlowDiskDetection(v string)`

SetSlowDiskDetection sets SlowDiskDetection field to given value.

### HasSlowDiskDetection

`func (o *AdvancedSettingsTableResourceInner) HasSlowDiskDetection() bool`

HasSlowDiskDetection returns a boolean if a field has been set.

### GetSlowDiskDetectionNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetSlowDiskDetectionNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetSlowDiskDetectionNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSlowDiskDetectionNumeric returns the SlowDiskDetectionNumeric field if non-nil, zero value otherwise.

### GetSlowDiskDetectionNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetSlowDiskDetectionNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetSlowDiskDetectionNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSlowDiskDetectionNumericOk returns a tuple with the SlowDiskDetectionNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowDiskDetectionNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetSlowDiskDetectionNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetSlowDiskDetectionNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSlowDiskDetectionNumeric sets SlowDiskDetectionNumeric field to given value.

### HasSlowDiskDetectionNumeric

`func (o *AdvancedSettingsTableResourceInner) HasSlowDiskDetectionNumeric() bool`

HasSlowDiskDetectionNumeric returns a boolean if a field has been set.

### GetSmart

`func (o *AdvancedSettingsTableResourceInner) GetSmart() string`

GetSmart returns the Smart field if non-nil, zero value otherwise.

### GetSmartOk

`func (o *AdvancedSettingsTableResourceInner) GetSmartOk() (*string, bool)`

GetSmartOk returns a tuple with the Smart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmart

`func (o *AdvancedSettingsTableResourceInner) SetSmart(v string)`

SetSmart sets Smart field to given value.

### HasSmart

`func (o *AdvancedSettingsTableResourceInner) HasSmart() bool`

HasSmart returns a boolean if a field has been set.

### GetSmartNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetSmartNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetSmartNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSmartNumeric returns the SmartNumeric field if non-nil, zero value otherwise.

### GetSmartNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetSmartNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetSmartNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSmartNumericOk returns a tuple with the SmartNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetSmartNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetSmartNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSmartNumeric sets SmartNumeric field to given value.

### HasSmartNumeric

`func (o *AdvancedSettingsTableResourceInner) HasSmartNumeric() bool`

HasSmartNumeric returns a boolean if a field has been set.

### GetSsdConcurrentAccess

`func (o *AdvancedSettingsTableResourceInner) GetSsdConcurrentAccess() string`

GetSsdConcurrentAccess returns the SsdConcurrentAccess field if non-nil, zero value otherwise.

### GetSsdConcurrentAccessOk

`func (o *AdvancedSettingsTableResourceInner) GetSsdConcurrentAccessOk() (*string, bool)`

GetSsdConcurrentAccessOk returns a tuple with the SsdConcurrentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdConcurrentAccess

`func (o *AdvancedSettingsTableResourceInner) SetSsdConcurrentAccess(v string)`

SetSsdConcurrentAccess sets SsdConcurrentAccess field to given value.

### HasSsdConcurrentAccess

`func (o *AdvancedSettingsTableResourceInner) HasSsdConcurrentAccess() bool`

HasSsdConcurrentAccess returns a boolean if a field has been set.

### GetSsdConcurrentAccessNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetSsdConcurrentAccessNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetSsdConcurrentAccessNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSsdConcurrentAccessNumeric returns the SsdConcurrentAccessNumeric field if non-nil, zero value otherwise.

### GetSsdConcurrentAccessNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetSsdConcurrentAccessNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetSsdConcurrentAccessNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSsdConcurrentAccessNumericOk returns a tuple with the SsdConcurrentAccessNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdConcurrentAccessNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetSsdConcurrentAccessNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetSsdConcurrentAccessNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSsdConcurrentAccessNumeric sets SsdConcurrentAccessNumeric field to given value.

### HasSsdConcurrentAccessNumeric

`func (o *AdvancedSettingsTableResourceInner) HasSsdConcurrentAccessNumeric() bool`

HasSsdConcurrentAccessNumeric returns a boolean if a field has been set.

### GetSuperCapFailure

`func (o *AdvancedSettingsTableResourceInner) GetSuperCapFailure() string`

GetSuperCapFailure returns the SuperCapFailure field if non-nil, zero value otherwise.

### GetSuperCapFailureOk

`func (o *AdvancedSettingsTableResourceInner) GetSuperCapFailureOk() (*string, bool)`

GetSuperCapFailureOk returns a tuple with the SuperCapFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperCapFailure

`func (o *AdvancedSettingsTableResourceInner) SetSuperCapFailure(v string)`

SetSuperCapFailure sets SuperCapFailure field to given value.

### HasSuperCapFailure

`func (o *AdvancedSettingsTableResourceInner) HasSuperCapFailure() bool`

HasSuperCapFailure returns a boolean if a field has been set.

### GetSuperCapFailureNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetSuperCapFailureNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetSuperCapFailureNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSuperCapFailureNumeric returns the SuperCapFailureNumeric field if non-nil, zero value otherwise.

### GetSuperCapFailureNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetSuperCapFailureNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetSuperCapFailureNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSuperCapFailureNumericOk returns a tuple with the SuperCapFailureNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperCapFailureNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetSuperCapFailureNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetSuperCapFailureNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSuperCapFailureNumeric sets SuperCapFailureNumeric field to given value.

### HasSuperCapFailureNumeric

`func (o *AdvancedSettingsTableResourceInner) HasSuperCapFailureNumeric() bool`

HasSuperCapFailureNumeric returns a boolean if a field has been set.

### GetSyncCacheMode

`func (o *AdvancedSettingsTableResourceInner) GetSyncCacheMode() string`

GetSyncCacheMode returns the SyncCacheMode field if non-nil, zero value otherwise.

### GetSyncCacheModeOk

`func (o *AdvancedSettingsTableResourceInner) GetSyncCacheModeOk() (*string, bool)`

GetSyncCacheModeOk returns a tuple with the SyncCacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCacheMode

`func (o *AdvancedSettingsTableResourceInner) SetSyncCacheMode(v string)`

SetSyncCacheMode sets SyncCacheMode field to given value.

### HasSyncCacheMode

`func (o *AdvancedSettingsTableResourceInner) HasSyncCacheMode() bool`

HasSyncCacheMode returns a boolean if a field has been set.

### GetSyncCacheModeNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetSyncCacheModeNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetSyncCacheModeNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSyncCacheModeNumeric returns the SyncCacheModeNumeric field if non-nil, zero value otherwise.

### GetSyncCacheModeNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetSyncCacheModeNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetSyncCacheModeNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetSyncCacheModeNumericOk returns a tuple with the SyncCacheModeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCacheModeNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetSyncCacheModeNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetSyncCacheModeNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetSyncCacheModeNumeric sets SyncCacheModeNumeric field to given value.

### HasSyncCacheModeNumeric

`func (o *AdvancedSettingsTableResourceInner) HasSyncCacheModeNumeric() bool`

HasSyncCacheModeNumeric returns a boolean if a field has been set.

### GetTemperatureExceeded

`func (o *AdvancedSettingsTableResourceInner) GetTemperatureExceeded() string`

GetTemperatureExceeded returns the TemperatureExceeded field if non-nil, zero value otherwise.

### GetTemperatureExceededOk

`func (o *AdvancedSettingsTableResourceInner) GetTemperatureExceededOk() (*string, bool)`

GetTemperatureExceededOk returns a tuple with the TemperatureExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureExceeded

`func (o *AdvancedSettingsTableResourceInner) SetTemperatureExceeded(v string)`

SetTemperatureExceeded sets TemperatureExceeded field to given value.

### HasTemperatureExceeded

`func (o *AdvancedSettingsTableResourceInner) HasTemperatureExceeded() bool`

HasTemperatureExceeded returns a boolean if a field has been set.

### GetTemperatureExceededNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetTemperatureExceededNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetTemperatureExceededNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetTemperatureExceededNumeric returns the TemperatureExceededNumeric field if non-nil, zero value otherwise.

### GetTemperatureExceededNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetTemperatureExceededNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetTemperatureExceededNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetTemperatureExceededNumericOk returns a tuple with the TemperatureExceededNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureExceededNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetTemperatureExceededNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetTemperatureExceededNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetTemperatureExceededNumeric sets TemperatureExceededNumeric field to given value.

### HasTemperatureExceededNumeric

`func (o *AdvancedSettingsTableResourceInner) HasTemperatureExceededNumeric() bool`

HasTemperatureExceededNumeric returns a boolean if a field has been set.

### GetUtilityPriority

`func (o *AdvancedSettingsTableResourceInner) GetUtilityPriority() string`

GetUtilityPriority returns the UtilityPriority field if non-nil, zero value otherwise.

### GetUtilityPriorityOk

`func (o *AdvancedSettingsTableResourceInner) GetUtilityPriorityOk() (*string, bool)`

GetUtilityPriorityOk returns a tuple with the UtilityPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilityPriority

`func (o *AdvancedSettingsTableResourceInner) SetUtilityPriority(v string)`

SetUtilityPriority sets UtilityPriority field to given value.

### HasUtilityPriority

`func (o *AdvancedSettingsTableResourceInner) HasUtilityPriority() bool`

HasUtilityPriority returns a boolean if a field has been set.

### GetUtilityPriorityNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetUtilityPriorityNumeric() int64`
=======
`func (o *AdvancedSettingsTableResourceInner) GetUtilityPriorityNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetUtilityPriorityNumeric returns the UtilityPriorityNumeric field if non-nil, zero value otherwise.

### GetUtilityPriorityNumericOk

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) GetUtilityPriorityNumericOk() (*int64, bool)`
=======
`func (o *AdvancedSettingsTableResourceInner) GetUtilityPriorityNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetUtilityPriorityNumericOk returns a tuple with the UtilityPriorityNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilityPriorityNumeric

<<<<<<< HEAD
`func (o *AdvancedSettingsTableResourceInner) SetUtilityPriorityNumeric(v int64)`
=======
`func (o *AdvancedSettingsTableResourceInner) SetUtilityPriorityNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetUtilityPriorityNumeric sets UtilityPriorityNumeric field to given value.

### HasUtilityPriorityNumeric

`func (o *AdvancedSettingsTableResourceInner) HasUtilityPriorityNumeric() bool`

HasUtilityPriorityNumeric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


