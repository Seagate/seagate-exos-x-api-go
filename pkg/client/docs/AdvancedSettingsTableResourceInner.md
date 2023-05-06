# AdvancedSettingsTableResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**AutoMap** | Pointer to **string** |  | [optional] 
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

`func (o *AdvancedSettingsTableResourceInner) GetAutoMapNumeric() int32`

GetAutoMapNumeric returns the AutoMapNumeric field if non-nil, zero value otherwise.

### GetAutoMapNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetAutoMapNumericOk() (*int32, bool)`

GetAutoMapNumericOk returns a tuple with the AutoMapNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoMapNumeric

`func (o *AdvancedSettingsTableResourceInner) SetAutoMapNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetAutoStallRecoveryNumeric() int32`

GetAutoStallRecoveryNumeric returns the AutoStallRecoveryNumeric field if non-nil, zero value otherwise.

### GetAutoStallRecoveryNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetAutoStallRecoveryNumericOk() (*int32, bool)`

GetAutoStallRecoveryNumericOk returns a tuple with the AutoStallRecoveryNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStallRecoveryNumeric

`func (o *AdvancedSettingsTableResourceInner) SetAutoStallRecoveryNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetAutoUnmapNumeric() int32`

GetAutoUnmapNumeric returns the AutoUnmapNumeric field if non-nil, zero value otherwise.

### GetAutoUnmapNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetAutoUnmapNumericOk() (*int32, bool)`

GetAutoUnmapNumericOk returns a tuple with the AutoUnmapNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUnmapNumeric

`func (o *AdvancedSettingsTableResourceInner) SetAutoUnmapNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetAutoWriteBackNumeric() int32`

GetAutoWriteBackNumeric returns the AutoWriteBackNumeric field if non-nil, zero value otherwise.

### GetAutoWriteBackNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetAutoWriteBackNumericOk() (*int32, bool)`

GetAutoWriteBackNumericOk returns a tuple with the AutoWriteBackNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoWriteBackNumeric

`func (o *AdvancedSettingsTableResourceInner) SetAutoWriteBackNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetBackgroundScrubInterval() int32`

GetBackgroundScrubInterval returns the BackgroundScrubInterval field if non-nil, zero value otherwise.

### GetBackgroundScrubIntervalOk

`func (o *AdvancedSettingsTableResourceInner) GetBackgroundScrubIntervalOk() (*int32, bool)`

GetBackgroundScrubIntervalOk returns a tuple with the BackgroundScrubInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundScrubInterval

`func (o *AdvancedSettingsTableResourceInner) SetBackgroundScrubInterval(v int32)`

SetBackgroundScrubInterval sets BackgroundScrubInterval field to given value.

### HasBackgroundScrubInterval

`func (o *AdvancedSettingsTableResourceInner) HasBackgroundScrubInterval() bool`

HasBackgroundScrubInterval returns a boolean if a field has been set.

### GetBackgroundScrubNumeric

`func (o *AdvancedSettingsTableResourceInner) GetBackgroundScrubNumeric() int32`

GetBackgroundScrubNumeric returns the BackgroundScrubNumeric field if non-nil, zero value otherwise.

### GetBackgroundScrubNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetBackgroundScrubNumericOk() (*int32, bool)`

GetBackgroundScrubNumericOk returns a tuple with the BackgroundScrubNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundScrubNumeric

`func (o *AdvancedSettingsTableResourceInner) SetBackgroundScrubNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetCacheFlushTimeoutNumeric() int32`

GetCacheFlushTimeoutNumeric returns the CacheFlushTimeoutNumeric field if non-nil, zero value otherwise.

### GetCacheFlushTimeoutNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetCacheFlushTimeoutNumericOk() (*int32, bool)`

GetCacheFlushTimeoutNumericOk returns a tuple with the CacheFlushTimeoutNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheFlushTimeoutNumeric

`func (o *AdvancedSettingsTableResourceInner) SetCacheFlushTimeoutNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetControllerFailureNumeric() int32`

GetControllerFailureNumeric returns the ControllerFailureNumeric field if non-nil, zero value otherwise.

### GetControllerFailureNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetControllerFailureNumericOk() (*int32, bool)`

GetControllerFailureNumericOk returns a tuple with the ControllerFailureNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerFailureNumeric

`func (o *AdvancedSettingsTableResourceInner) SetControllerFailureNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetDefaultMappingNumeric() int32`

GetDefaultMappingNumeric returns the DefaultMappingNumeric field if non-nil, zero value otherwise.

### GetDefaultMappingNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetDefaultMappingNumericOk() (*int32, bool)`

GetDefaultMappingNumericOk returns a tuple with the DefaultMappingNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMappingNumeric

`func (o *AdvancedSettingsTableResourceInner) SetDefaultMappingNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetDeleteOverrideNumeric() int32`

GetDeleteOverrideNumeric returns the DeleteOverrideNumeric field if non-nil, zero value otherwise.

### GetDeleteOverrideNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetDeleteOverrideNumericOk() (*int32, bool)`

GetDeleteOverrideNumericOk returns a tuple with the DeleteOverrideNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOverrideNumeric

`func (o *AdvancedSettingsTableResourceInner) SetDeleteOverrideNumeric(v int32)`

SetDeleteOverrideNumeric sets DeleteOverrideNumeric field to given value.

### HasDeleteOverrideNumeric

`func (o *AdvancedSettingsTableResourceInner) HasDeleteOverrideNumeric() bool`

HasDeleteOverrideNumeric returns a boolean if a field has been set.

### GetDiskDsdDelay

`func (o *AdvancedSettingsTableResourceInner) GetDiskDsdDelay() int32`

GetDiskDsdDelay returns the DiskDsdDelay field if non-nil, zero value otherwise.

### GetDiskDsdDelayOk

`func (o *AdvancedSettingsTableResourceInner) GetDiskDsdDelayOk() (*int32, bool)`

GetDiskDsdDelayOk returns a tuple with the DiskDsdDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDsdDelay

`func (o *AdvancedSettingsTableResourceInner) SetDiskDsdDelay(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetDiskDsdEnableNumeric() int32`

GetDiskDsdEnableNumeric returns the DiskDsdEnableNumeric field if non-nil, zero value otherwise.

### GetDiskDsdEnableNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetDiskDsdEnableNumericOk() (*int32, bool)`

GetDiskDsdEnableNumericOk returns a tuple with the DiskDsdEnableNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDsdEnableNumeric

`func (o *AdvancedSettingsTableResourceInner) SetDiskDsdEnableNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetDiskFirmwareUpdateNumeric() int32`

GetDiskFirmwareUpdateNumeric returns the DiskFirmwareUpdateNumeric field if non-nil, zero value otherwise.

### GetDiskFirmwareUpdateNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetDiskFirmwareUpdateNumericOk() (*int32, bool)`

GetDiskFirmwareUpdateNumericOk returns a tuple with the DiskFirmwareUpdateNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFirmwareUpdateNumeric

`func (o *AdvancedSettingsTableResourceInner) SetDiskFirmwareUpdateNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetFanFailureNumeric() int32`

GetFanFailureNumeric returns the FanFailureNumeric field if non-nil, zero value otherwise.

### GetFanFailureNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetFanFailureNumericOk() (*int32, bool)`

GetFanFailureNumericOk returns a tuple with the FanFailureNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanFailureNumeric

`func (o *AdvancedSettingsTableResourceInner) SetFanFailureNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetHedgedReadsTimeoutNumeric() int32`

GetHedgedReadsTimeoutNumeric returns the HedgedReadsTimeoutNumeric field if non-nil, zero value otherwise.

### GetHedgedReadsTimeoutNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetHedgedReadsTimeoutNumericOk() (*int32, bool)`

GetHedgedReadsTimeoutNumericOk returns a tuple with the HedgedReadsTimeoutNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHedgedReadsTimeoutNumeric

`func (o *AdvancedSettingsTableResourceInner) SetHedgedReadsTimeoutNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetHostCacheControlNumeric() int32`

GetHostCacheControlNumeric returns the HostCacheControlNumeric field if non-nil, zero value otherwise.

### GetHostCacheControlNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetHostCacheControlNumericOk() (*int32, bool)`

GetHostCacheControlNumericOk returns a tuple with the HostCacheControlNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostCacheControlNumeric

`func (o *AdvancedSettingsTableResourceInner) SetHostCacheControlNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetIndependentCacheNumeric() int32`

GetIndependentCacheNumeric returns the IndependentCacheNumeric field if non-nil, zero value otherwise.

### GetIndependentCacheNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetIndependentCacheNumericOk() (*int32, bool)`

GetIndependentCacheNumericOk returns a tuple with the IndependentCacheNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndependentCacheNumeric

`func (o *AdvancedSettingsTableResourceInner) SetIndependentCacheNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetLargePoolsNumeric() int32`

GetLargePoolsNumeric returns the LargePoolsNumeric field if non-nil, zero value otherwise.

### GetLargePoolsNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetLargePoolsNumericOk() (*int32, bool)`

GetLargePoolsNumericOk returns a tuple with the LargePoolsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargePoolsNumeric

`func (o *AdvancedSettingsTableResourceInner) SetLargePoolsNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetManagedLogsNumeric() int32`

GetManagedLogsNumeric returns the ManagedLogsNumeric field if non-nil, zero value otherwise.

### GetManagedLogsNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetManagedLogsNumericOk() (*int32, bool)`

GetManagedLogsNumericOk returns a tuple with the ManagedLogsNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedLogsNumeric

`func (o *AdvancedSettingsTableResourceInner) SetManagedLogsNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetMemoryCardFailureNumeric() int32`

GetMemoryCardFailureNumeric returns the MemoryCardFailureNumeric field if non-nil, zero value otherwise.

### GetMemoryCardFailureNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetMemoryCardFailureNumericOk() (*int32, bool)`

GetMemoryCardFailureNumericOk returns a tuple with the MemoryCardFailureNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCardFailureNumeric

`func (o *AdvancedSettingsTableResourceInner) SetMemoryCardFailureNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetMissingLunResponseNumeric() int32`

GetMissingLunResponseNumeric returns the MissingLunResponseNumeric field if non-nil, zero value otherwise.

### GetMissingLunResponseNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetMissingLunResponseNumericOk() (*int32, bool)`

GetMissingLunResponseNumericOk returns a tuple with the MissingLunResponseNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingLunResponseNumeric

`func (o *AdvancedSettingsTableResourceInner) SetMissingLunResponseNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetPartnerFirmwareUpgradeNumeric() int32`

GetPartnerFirmwareUpgradeNumeric returns the PartnerFirmwareUpgradeNumeric field if non-nil, zero value otherwise.

### GetPartnerFirmwareUpgradeNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetPartnerFirmwareUpgradeNumericOk() (*int32, bool)`

GetPartnerFirmwareUpgradeNumericOk returns a tuple with the PartnerFirmwareUpgradeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerFirmwareUpgradeNumeric

`func (o *AdvancedSettingsTableResourceInner) SetPartnerFirmwareUpgradeNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetPartnerNotifyNumeric() int32`

GetPartnerNotifyNumeric returns the PartnerNotifyNumeric field if non-nil, zero value otherwise.

### GetPartnerNotifyNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetPartnerNotifyNumericOk() (*int32, bool)`

GetPartnerNotifyNumericOk returns a tuple with the PartnerNotifyNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerNotifyNumeric

`func (o *AdvancedSettingsTableResourceInner) SetPartnerNotifyNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetPcieHotplugNumeric() int32`

GetPcieHotplugNumeric returns the PcieHotplugNumeric field if non-nil, zero value otherwise.

### GetPcieHotplugNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetPcieHotplugNumericOk() (*int32, bool)`

GetPcieHotplugNumericOk returns a tuple with the PcieHotplugNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieHotplugNumeric

`func (o *AdvancedSettingsTableResourceInner) SetPcieHotplugNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetPowerSupplyFailureNumeric() int32`

GetPowerSupplyFailureNumeric returns the PowerSupplyFailureNumeric field if non-nil, zero value otherwise.

### GetPowerSupplyFailureNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetPowerSupplyFailureNumericOk() (*int32, bool)`

GetPowerSupplyFailureNumericOk returns a tuple with the PowerSupplyFailureNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerSupplyFailureNumeric

`func (o *AdvancedSettingsTableResourceInner) SetPowerSupplyFailureNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetRandomIoPerformanceOptimizationNumeric() int32`

GetRandomIoPerformanceOptimizationNumeric returns the RandomIoPerformanceOptimizationNumeric field if non-nil, zero value otherwise.

### GetRandomIoPerformanceOptimizationNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetRandomIoPerformanceOptimizationNumericOk() (*int32, bool)`

GetRandomIoPerformanceOptimizationNumericOk returns a tuple with the RandomIoPerformanceOptimizationNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomIoPerformanceOptimizationNumeric

`func (o *AdvancedSettingsTableResourceInner) SetRandomIoPerformanceOptimizationNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetRemanufactureNumeric() int32`

GetRemanufactureNumeric returns the RemanufactureNumeric field if non-nil, zero value otherwise.

### GetRemanufactureNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetRemanufactureNumericOk() (*int32, bool)`

GetRemanufactureNumericOk returns a tuple with the RemanufactureNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemanufactureNumeric

`func (o *AdvancedSettingsTableResourceInner) SetRemanufactureNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetRestartOnCapiFailNumeric() int32`

GetRestartOnCapiFailNumeric returns the RestartOnCapiFailNumeric field if non-nil, zero value otherwise.

### GetRestartOnCapiFailNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetRestartOnCapiFailNumericOk() (*int32, bool)`

GetRestartOnCapiFailNumericOk returns a tuple with the RestartOnCapiFailNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartOnCapiFailNumeric

`func (o *AdvancedSettingsTableResourceInner) SetRestartOnCapiFailNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetScrubScheduleNumeric() int32`

GetScrubScheduleNumeric returns the ScrubScheduleNumeric field if non-nil, zero value otherwise.

### GetScrubScheduleNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetScrubScheduleNumericOk() (*int32, bool)`

GetScrubScheduleNumericOk returns a tuple with the ScrubScheduleNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrubScheduleNumeric

`func (o *AdvancedSettingsTableResourceInner) SetScrubScheduleNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetSingleControllerNumeric() int32`

GetSingleControllerNumeric returns the SingleControllerNumeric field if non-nil, zero value otherwise.

### GetSingleControllerNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetSingleControllerNumericOk() (*int32, bool)`

GetSingleControllerNumericOk returns a tuple with the SingleControllerNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleControllerNumeric

`func (o *AdvancedSettingsTableResourceInner) SetSingleControllerNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetSlotAffinityNumeric() int32`

GetSlotAffinityNumeric returns the SlotAffinityNumeric field if non-nil, zero value otherwise.

### GetSlotAffinityNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetSlotAffinityNumericOk() (*int32, bool)`

GetSlotAffinityNumericOk returns a tuple with the SlotAffinityNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotAffinityNumeric

`func (o *AdvancedSettingsTableResourceInner) SetSlotAffinityNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetSlowDiskDetectionNumeric() int32`

GetSlowDiskDetectionNumeric returns the SlowDiskDetectionNumeric field if non-nil, zero value otherwise.

### GetSlowDiskDetectionNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetSlowDiskDetectionNumericOk() (*int32, bool)`

GetSlowDiskDetectionNumericOk returns a tuple with the SlowDiskDetectionNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowDiskDetectionNumeric

`func (o *AdvancedSettingsTableResourceInner) SetSlowDiskDetectionNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetSmartNumeric() int32`

GetSmartNumeric returns the SmartNumeric field if non-nil, zero value otherwise.

### GetSmartNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetSmartNumericOk() (*int32, bool)`

GetSmartNumericOk returns a tuple with the SmartNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartNumeric

`func (o *AdvancedSettingsTableResourceInner) SetSmartNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetSsdConcurrentAccessNumeric() int32`

GetSsdConcurrentAccessNumeric returns the SsdConcurrentAccessNumeric field if non-nil, zero value otherwise.

### GetSsdConcurrentAccessNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetSsdConcurrentAccessNumericOk() (*int32, bool)`

GetSsdConcurrentAccessNumericOk returns a tuple with the SsdConcurrentAccessNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdConcurrentAccessNumeric

`func (o *AdvancedSettingsTableResourceInner) SetSsdConcurrentAccessNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetSuperCapFailureNumeric() int32`

GetSuperCapFailureNumeric returns the SuperCapFailureNumeric field if non-nil, zero value otherwise.

### GetSuperCapFailureNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetSuperCapFailureNumericOk() (*int32, bool)`

GetSuperCapFailureNumericOk returns a tuple with the SuperCapFailureNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperCapFailureNumeric

`func (o *AdvancedSettingsTableResourceInner) SetSuperCapFailureNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetSyncCacheModeNumeric() int32`

GetSyncCacheModeNumeric returns the SyncCacheModeNumeric field if non-nil, zero value otherwise.

### GetSyncCacheModeNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetSyncCacheModeNumericOk() (*int32, bool)`

GetSyncCacheModeNumericOk returns a tuple with the SyncCacheModeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCacheModeNumeric

`func (o *AdvancedSettingsTableResourceInner) SetSyncCacheModeNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetTemperatureExceededNumeric() int32`

GetTemperatureExceededNumeric returns the TemperatureExceededNumeric field if non-nil, zero value otherwise.

### GetTemperatureExceededNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetTemperatureExceededNumericOk() (*int32, bool)`

GetTemperatureExceededNumericOk returns a tuple with the TemperatureExceededNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureExceededNumeric

`func (o *AdvancedSettingsTableResourceInner) SetTemperatureExceededNumeric(v int32)`

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

`func (o *AdvancedSettingsTableResourceInner) GetUtilityPriorityNumeric() int32`

GetUtilityPriorityNumeric returns the UtilityPriorityNumeric field if non-nil, zero value otherwise.

### GetUtilityPriorityNumericOk

`func (o *AdvancedSettingsTableResourceInner) GetUtilityPriorityNumericOk() (*int32, bool)`

GetUtilityPriorityNumericOk returns a tuple with the UtilityPriorityNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilityPriorityNumeric

`func (o *AdvancedSettingsTableResourceInner) SetUtilityPriorityNumeric(v int32)`

SetUtilityPriorityNumeric sets UtilityPriorityNumeric field to given value.

### HasUtilityPriorityNumeric

`func (o *AdvancedSettingsTableResourceInner) HasUtilityPriorityNumeric() bool`

HasUtilityPriorityNumeric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


