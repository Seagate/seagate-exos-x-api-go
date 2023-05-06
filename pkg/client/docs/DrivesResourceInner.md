# DrivesResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**Architecture** | Pointer to **string** |  | [optional] 
**ArchitectureNumeric** | Pointer to **int32** |  | [optional] 
**AssuranceLevel** | Pointer to **string** |  | [optional] 
**AssuranceLevelNumeric** | Pointer to **int32** |  | [optional] 
**Attributes** | Pointer to **string** | Indicates if the disk is single-pathed | [optional] 
**AttributesNumeric** | Pointer to **int32** | Indicates if the disk is single-pathed( In numeric form ) | [optional] 
**AvgRspTime** | Pointer to **int32** |  | [optional] 
**Blink** | Pointer to **int32** | Indicates whether the locator LED is on | [optional] 
**Blocks** | Pointer to **int32** | The size in blocks | [optional] 
**Blocksize** | Pointer to **int32** |  | [optional] 
**ContainerIndex** | Pointer to **int32** |  | [optional] 
**CopybackState** | Pointer to **string** | Copyback State | [optional] 
**CopybackStateNumeric** | Pointer to **int32** | Copyback State( In numeric form ) | [optional] 
**CurrentJobCompletion** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DescriptionNumeric** | Pointer to **int32** |  | [optional] 
**DiskDsdCount** | Pointer to **int32** |  | [optional] 
**DiskGroup** | Pointer to **string** | Disk Group details | [optional] 
**DrawerId** | Pointer to **int32** |  | [optional] 
**DriveDownCode** | Pointer to **int32** |  | [optional] 
**DualPort** | Pointer to **int32** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**EnclosureId** | Pointer to **int32** |  | [optional] 
**EnclosureWwn** | Pointer to **string** | Enclosure World Wide Name | [optional] 
**EnclosuresUrl** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **int32** |  | [optional] 
**ExtendedStatus** | Pointer to **int32** | Extended status (bits) | [optional] 
**FcP1Channel** | Pointer to **int32** |  | [optional] 
**FcP1DeviceId** | Pointer to **int32** |  | [optional] 
**FcP1NodeWwn** | Pointer to **string** |  | [optional] 
**FcP1PortWwn** | Pointer to **string** |  | [optional] 
**FcP1UnitNumber** | Pointer to **int32** |  | [optional] 
**FcP2Channel** | Pointer to **int32** |  | [optional] 
**FcP2DeviceId** | Pointer to **int32** |  | [optional] 
**FcP2NodeWwn** | Pointer to **string** |  | [optional] 
**FcP2PortWwn** | Pointer to **string** |  | [optional] 
**FcP2UnitNumber** | Pointer to **int32** |  | [optional] 
**FdeConfigTime** | Pointer to **string** |  | [optional] 
**FdeConfigTimeNumeric** | Pointer to **int32** |  | [optional] 
**FdeState** | Pointer to **string** |  | [optional] 
**FdeStateNumeric** | Pointer to **int32** |  | [optional] 
**FipsCapable** | Pointer to **string** |  | [optional] 
**FipsCapableNumeric** | Pointer to **int32** |  | [optional] 
**FirmwareUpdateStatus** | Pointer to **string** |  | [optional] 
**FirmwareUpdateStatusNumeric** | Pointer to **int32** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**HealthNumeric** | Pointer to **int32** |  | [optional] 
**HealthReason** | Pointer to **string** |  | [optional] 
**HealthReasonNumeric** | Pointer to **int32** |  | [optional] 
**HealthRecommendation** | Pointer to **string** |  | [optional] 
**HealthRecommendationNumeric** | Pointer to **int32** |  | [optional] 
**ImportLockKeyId** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Interface** | Pointer to **string** |  | [optional] 
**InterfaceNumeric** | Pointer to **int32** |  | [optional] 
**JobRunning** | Pointer to **string** |  | [optional] 
**JobRunningNumeric** | Pointer to **int32** |  | [optional] 
**KmipState** | Pointer to **string** |  | [optional] 
**KmipStateNumeric** | Pointer to **int32** |  | [optional] 
**LedStatus** | Pointer to **string** |  | [optional] 
**LedStatusNumeric** | Pointer to **int32** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**LocatorLed** | Pointer to **string** | Indicates whether the locator LED is on | [optional] 
**LocatorLedNumeric** | Pointer to **int32** | Indicates whether the locator LED is on( In numeric form ) | [optional] 
**LockKeyId** | Pointer to **string** |  | [optional] 
**MemberIndex** | Pointer to **int32** | Index for this disk in the vdisk list | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**NumDepoppedHeads** | Pointer to **int32** | Number of depopulated heads | [optional] 
**NumberOfIos** | Pointer to **int32** |  | [optional] 
**Owner** | Pointer to **string** | Controller owning the component | [optional] 
**OwnerNumeric** | Pointer to **int32** | Controller owning the component( In numeric form ) | [optional] 
**PiFormatted** | Pointer to **string** | Used to describe the Disk Protection Information | [optional] 
**PiFormattedNumeric** | Pointer to **int32** | Used to describe the Disk Protection Information( In numeric form ) | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**PowerOnHours** | Pointer to **int32** |  | [optional] 
**ReconState** | Pointer to **string** | Reconstruction State | [optional] 
**ReconStateNumeric** | Pointer to **int32** | Reconstruction State( In numeric form ) | [optional] 
**Remanufacture** | Pointer to **string** |  | [optional] 
**RemanufactureNumeric** | Pointer to **int32** |  | [optional] 
**Revision** | Pointer to **string** | Current revision for this FRU | [optional] 
**Rpm** | Pointer to **int32** | Vendor-specified disk speed in thousands of revolutions per minute | [optional] 
**ScsiId** | Pointer to **int32** | SCSI ID assigned to this disk for the primary channel | [optional] 
**SecondaryChannel** | Pointer to **int32** | SCSI ID assigned to this disk for the secondary channel | [optional] 
**SectorFormat** | Pointer to **string** |  | [optional] 
**SectorFormatNumeric** | Pointer to **int32** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**SinglePorted** | Pointer to **string** |  | [optional] 
**SinglePortedNumeric** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **string** | The size or capacity formatted with the current session base, precision, and units | [optional] 
**SizeNumeric** | Pointer to **int32** | The size or capacity formatted with the current session base, precision, and units( In numeric form ) | [optional] 
**Slot** | Pointer to **int32** | The slot number where the disk is located | [optional] 
**Smart** | Pointer to **string** |  | [optional] 
**SmartNumeric** | Pointer to **int32** |  | [optional] 
**Speed** | Pointer to **int32** |  | [optional] 
**SpunDown** | Pointer to **int32** | Indicates a drive is spun down | [optional] 
**SsdLifeLeft** | Pointer to **string** |  | [optional] 
**SsdLifeLeftNumeric** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StoragePoolName** | Pointer to **string** | User-defined name for the pool | [optional] 
**StorageTier** | Pointer to **string** | Disk group tier assignment for tiered migration | [optional] 
**StorageTierNumeric** | Pointer to **int32** | Disk group tier assignment for tiered migration( In numeric form ) | [optional] 
**SupportsUnmap** | Pointer to **string** |  | [optional] 
**SupportsUnmapNumeric** | Pointer to **int32** |  | [optional] 
**Temperature** | Pointer to **string** |  | [optional] 
**TemperatureNumeric** | Pointer to **int32** |  | [optional] 
**TemperatureStatus** | Pointer to **string** |  | [optional] 
**TemperatureStatusNumeric** | Pointer to **int32** |  | [optional] 
**TotalDataTransferred** | Pointer to **string** |  | [optional] 
**TotalDataTransferredNumeric** | Pointer to **int32** |  | [optional] 
**TransferRate** | Pointer to **string** | Transfer rate of the disk | [optional] 
**TransferRateNumeric** | Pointer to **int32** | Transfer rate of the disk( In numeric form ) | [optional] 
**Type** | Pointer to **string** | The type of disk | [optional] 
**TypeNumeric** | Pointer to **int32** | The type of disk( In numeric form ) | [optional] 
**Url** | Pointer to **string** | The resource URL | [optional] 
**Usage** | Pointer to **string** | Disk usage | [optional] 
**UsageNumeric** | Pointer to **int32** | Disk usage | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**VirtualDiskSerial** | Pointer to **string** | Unique serial number for the disk group | [optional] 

## Methods

### NewDrivesResourceInner

`func NewDrivesResourceInner() *DrivesResourceInner`

NewDrivesResourceInner instantiates a new DrivesResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrivesResourceInnerWithDefaults

`func NewDrivesResourceInnerWithDefaults() *DrivesResourceInner`

NewDrivesResourceInnerWithDefaults instantiates a new DrivesResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *DrivesResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *DrivesResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *DrivesResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *DrivesResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *DrivesResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DrivesResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DrivesResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DrivesResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetArchitecture

`func (o *DrivesResourceInner) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *DrivesResourceInner) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *DrivesResourceInner) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *DrivesResourceInner) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetArchitectureNumeric

`func (o *DrivesResourceInner) GetArchitectureNumeric() int32`

GetArchitectureNumeric returns the ArchitectureNumeric field if non-nil, zero value otherwise.

### GetArchitectureNumericOk

`func (o *DrivesResourceInner) GetArchitectureNumericOk() (*int32, bool)`

GetArchitectureNumericOk returns a tuple with the ArchitectureNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectureNumeric

`func (o *DrivesResourceInner) SetArchitectureNumeric(v int32)`

SetArchitectureNumeric sets ArchitectureNumeric field to given value.

### HasArchitectureNumeric

`func (o *DrivesResourceInner) HasArchitectureNumeric() bool`

HasArchitectureNumeric returns a boolean if a field has been set.

### GetAssuranceLevel

`func (o *DrivesResourceInner) GetAssuranceLevel() string`

GetAssuranceLevel returns the AssuranceLevel field if non-nil, zero value otherwise.

### GetAssuranceLevelOk

`func (o *DrivesResourceInner) GetAssuranceLevelOk() (*string, bool)`

GetAssuranceLevelOk returns a tuple with the AssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceLevel

`func (o *DrivesResourceInner) SetAssuranceLevel(v string)`

SetAssuranceLevel sets AssuranceLevel field to given value.

### HasAssuranceLevel

`func (o *DrivesResourceInner) HasAssuranceLevel() bool`

HasAssuranceLevel returns a boolean if a field has been set.

### GetAssuranceLevelNumeric

`func (o *DrivesResourceInner) GetAssuranceLevelNumeric() int32`

GetAssuranceLevelNumeric returns the AssuranceLevelNumeric field if non-nil, zero value otherwise.

### GetAssuranceLevelNumericOk

`func (o *DrivesResourceInner) GetAssuranceLevelNumericOk() (*int32, bool)`

GetAssuranceLevelNumericOk returns a tuple with the AssuranceLevelNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceLevelNumeric

`func (o *DrivesResourceInner) SetAssuranceLevelNumeric(v int32)`

SetAssuranceLevelNumeric sets AssuranceLevelNumeric field to given value.

### HasAssuranceLevelNumeric

`func (o *DrivesResourceInner) HasAssuranceLevelNumeric() bool`

HasAssuranceLevelNumeric returns a boolean if a field has been set.

### GetAttributes

`func (o *DrivesResourceInner) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DrivesResourceInner) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DrivesResourceInner) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DrivesResourceInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetAttributesNumeric

`func (o *DrivesResourceInner) GetAttributesNumeric() int32`

GetAttributesNumeric returns the AttributesNumeric field if non-nil, zero value otherwise.

### GetAttributesNumericOk

`func (o *DrivesResourceInner) GetAttributesNumericOk() (*int32, bool)`

GetAttributesNumericOk returns a tuple with the AttributesNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesNumeric

`func (o *DrivesResourceInner) SetAttributesNumeric(v int32)`

SetAttributesNumeric sets AttributesNumeric field to given value.

### HasAttributesNumeric

`func (o *DrivesResourceInner) HasAttributesNumeric() bool`

HasAttributesNumeric returns a boolean if a field has been set.

### GetAvgRspTime

`func (o *DrivesResourceInner) GetAvgRspTime() int32`

GetAvgRspTime returns the AvgRspTime field if non-nil, zero value otherwise.

### GetAvgRspTimeOk

`func (o *DrivesResourceInner) GetAvgRspTimeOk() (*int32, bool)`

GetAvgRspTimeOk returns a tuple with the AvgRspTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgRspTime

`func (o *DrivesResourceInner) SetAvgRspTime(v int32)`

SetAvgRspTime sets AvgRspTime field to given value.

### HasAvgRspTime

`func (o *DrivesResourceInner) HasAvgRspTime() bool`

HasAvgRspTime returns a boolean if a field has been set.

### GetBlink

`func (o *DrivesResourceInner) GetBlink() int32`

GetBlink returns the Blink field if non-nil, zero value otherwise.

### GetBlinkOk

`func (o *DrivesResourceInner) GetBlinkOk() (*int32, bool)`

GetBlinkOk returns a tuple with the Blink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlink

`func (o *DrivesResourceInner) SetBlink(v int32)`

SetBlink sets Blink field to given value.

### HasBlink

`func (o *DrivesResourceInner) HasBlink() bool`

HasBlink returns a boolean if a field has been set.

### GetBlocks

`func (o *DrivesResourceInner) GetBlocks() int32`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *DrivesResourceInner) GetBlocksOk() (*int32, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *DrivesResourceInner) SetBlocks(v int32)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *DrivesResourceInner) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetBlocksize

`func (o *DrivesResourceInner) GetBlocksize() int32`

GetBlocksize returns the Blocksize field if non-nil, zero value otherwise.

### GetBlocksizeOk

`func (o *DrivesResourceInner) GetBlocksizeOk() (*int32, bool)`

GetBlocksizeOk returns a tuple with the Blocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocksize

`func (o *DrivesResourceInner) SetBlocksize(v int32)`

SetBlocksize sets Blocksize field to given value.

### HasBlocksize

`func (o *DrivesResourceInner) HasBlocksize() bool`

HasBlocksize returns a boolean if a field has been set.

### GetContainerIndex

`func (o *DrivesResourceInner) GetContainerIndex() int32`

GetContainerIndex returns the ContainerIndex field if non-nil, zero value otherwise.

### GetContainerIndexOk

`func (o *DrivesResourceInner) GetContainerIndexOk() (*int32, bool)`

GetContainerIndexOk returns a tuple with the ContainerIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerIndex

`func (o *DrivesResourceInner) SetContainerIndex(v int32)`

SetContainerIndex sets ContainerIndex field to given value.

### HasContainerIndex

`func (o *DrivesResourceInner) HasContainerIndex() bool`

HasContainerIndex returns a boolean if a field has been set.

### GetCopybackState

`func (o *DrivesResourceInner) GetCopybackState() string`

GetCopybackState returns the CopybackState field if non-nil, zero value otherwise.

### GetCopybackStateOk

`func (o *DrivesResourceInner) GetCopybackStateOk() (*string, bool)`

GetCopybackStateOk returns a tuple with the CopybackState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopybackState

`func (o *DrivesResourceInner) SetCopybackState(v string)`

SetCopybackState sets CopybackState field to given value.

### HasCopybackState

`func (o *DrivesResourceInner) HasCopybackState() bool`

HasCopybackState returns a boolean if a field has been set.

### GetCopybackStateNumeric

`func (o *DrivesResourceInner) GetCopybackStateNumeric() int32`

GetCopybackStateNumeric returns the CopybackStateNumeric field if non-nil, zero value otherwise.

### GetCopybackStateNumericOk

`func (o *DrivesResourceInner) GetCopybackStateNumericOk() (*int32, bool)`

GetCopybackStateNumericOk returns a tuple with the CopybackStateNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopybackStateNumeric

`func (o *DrivesResourceInner) SetCopybackStateNumeric(v int32)`

SetCopybackStateNumeric sets CopybackStateNumeric field to given value.

### HasCopybackStateNumeric

`func (o *DrivesResourceInner) HasCopybackStateNumeric() bool`

HasCopybackStateNumeric returns a boolean if a field has been set.

### GetCurrentJobCompletion

`func (o *DrivesResourceInner) GetCurrentJobCompletion() string`

GetCurrentJobCompletion returns the CurrentJobCompletion field if non-nil, zero value otherwise.

### GetCurrentJobCompletionOk

`func (o *DrivesResourceInner) GetCurrentJobCompletionOk() (*string, bool)`

GetCurrentJobCompletionOk returns a tuple with the CurrentJobCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentJobCompletion

`func (o *DrivesResourceInner) SetCurrentJobCompletion(v string)`

SetCurrentJobCompletion sets CurrentJobCompletion field to given value.

### HasCurrentJobCompletion

`func (o *DrivesResourceInner) HasCurrentJobCompletion() bool`

HasCurrentJobCompletion returns a boolean if a field has been set.

### GetDescription

`func (o *DrivesResourceInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DrivesResourceInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DrivesResourceInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DrivesResourceInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionNumeric

`func (o *DrivesResourceInner) GetDescriptionNumeric() int32`

GetDescriptionNumeric returns the DescriptionNumeric field if non-nil, zero value otherwise.

### GetDescriptionNumericOk

`func (o *DrivesResourceInner) GetDescriptionNumericOk() (*int32, bool)`

GetDescriptionNumericOk returns a tuple with the DescriptionNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionNumeric

`func (o *DrivesResourceInner) SetDescriptionNumeric(v int32)`

SetDescriptionNumeric sets DescriptionNumeric field to given value.

### HasDescriptionNumeric

`func (o *DrivesResourceInner) HasDescriptionNumeric() bool`

HasDescriptionNumeric returns a boolean if a field has been set.

### GetDiskDsdCount

`func (o *DrivesResourceInner) GetDiskDsdCount() int32`

GetDiskDsdCount returns the DiskDsdCount field if non-nil, zero value otherwise.

### GetDiskDsdCountOk

`func (o *DrivesResourceInner) GetDiskDsdCountOk() (*int32, bool)`

GetDiskDsdCountOk returns a tuple with the DiskDsdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDsdCount

`func (o *DrivesResourceInner) SetDiskDsdCount(v int32)`

SetDiskDsdCount sets DiskDsdCount field to given value.

### HasDiskDsdCount

`func (o *DrivesResourceInner) HasDiskDsdCount() bool`

HasDiskDsdCount returns a boolean if a field has been set.

### GetDiskGroup

`func (o *DrivesResourceInner) GetDiskGroup() string`

GetDiskGroup returns the DiskGroup field if non-nil, zero value otherwise.

### GetDiskGroupOk

`func (o *DrivesResourceInner) GetDiskGroupOk() (*string, bool)`

GetDiskGroupOk returns a tuple with the DiskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGroup

`func (o *DrivesResourceInner) SetDiskGroup(v string)`

SetDiskGroup sets DiskGroup field to given value.

### HasDiskGroup

`func (o *DrivesResourceInner) HasDiskGroup() bool`

HasDiskGroup returns a boolean if a field has been set.

### GetDrawerId

`func (o *DrivesResourceInner) GetDrawerId() int32`

GetDrawerId returns the DrawerId field if non-nil, zero value otherwise.

### GetDrawerIdOk

`func (o *DrivesResourceInner) GetDrawerIdOk() (*int32, bool)`

GetDrawerIdOk returns a tuple with the DrawerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawerId

`func (o *DrivesResourceInner) SetDrawerId(v int32)`

SetDrawerId sets DrawerId field to given value.

### HasDrawerId

`func (o *DrivesResourceInner) HasDrawerId() bool`

HasDrawerId returns a boolean if a field has been set.

### GetDriveDownCode

`func (o *DrivesResourceInner) GetDriveDownCode() int32`

GetDriveDownCode returns the DriveDownCode field if non-nil, zero value otherwise.

### GetDriveDownCodeOk

`func (o *DrivesResourceInner) GetDriveDownCodeOk() (*int32, bool)`

GetDriveDownCodeOk returns a tuple with the DriveDownCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveDownCode

`func (o *DrivesResourceInner) SetDriveDownCode(v int32)`

SetDriveDownCode sets DriveDownCode field to given value.

### HasDriveDownCode

`func (o *DrivesResourceInner) HasDriveDownCode() bool`

HasDriveDownCode returns a boolean if a field has been set.

### GetDualPort

`func (o *DrivesResourceInner) GetDualPort() int32`

GetDualPort returns the DualPort field if non-nil, zero value otherwise.

### GetDualPortOk

`func (o *DrivesResourceInner) GetDualPortOk() (*int32, bool)`

GetDualPortOk returns a tuple with the DualPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualPort

`func (o *DrivesResourceInner) SetDualPort(v int32)`

SetDualPort sets DualPort field to given value.

### HasDualPort

`func (o *DrivesResourceInner) HasDualPort() bool`

HasDualPort returns a boolean if a field has been set.

### GetDurableId

`func (o *DrivesResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *DrivesResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *DrivesResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *DrivesResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetEnclosureId

`func (o *DrivesResourceInner) GetEnclosureId() int32`

GetEnclosureId returns the EnclosureId field if non-nil, zero value otherwise.

### GetEnclosureIdOk

`func (o *DrivesResourceInner) GetEnclosureIdOk() (*int32, bool)`

GetEnclosureIdOk returns a tuple with the EnclosureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureId

`func (o *DrivesResourceInner) SetEnclosureId(v int32)`

SetEnclosureId sets EnclosureId field to given value.

### HasEnclosureId

`func (o *DrivesResourceInner) HasEnclosureId() bool`

HasEnclosureId returns a boolean if a field has been set.

### GetEnclosureWwn

`func (o *DrivesResourceInner) GetEnclosureWwn() string`

GetEnclosureWwn returns the EnclosureWwn field if non-nil, zero value otherwise.

### GetEnclosureWwnOk

`func (o *DrivesResourceInner) GetEnclosureWwnOk() (*string, bool)`

GetEnclosureWwnOk returns a tuple with the EnclosureWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosureWwn

`func (o *DrivesResourceInner) SetEnclosureWwn(v string)`

SetEnclosureWwn sets EnclosureWwn field to given value.

### HasEnclosureWwn

`func (o *DrivesResourceInner) HasEnclosureWwn() bool`

HasEnclosureWwn returns a boolean if a field has been set.

### GetEnclosuresUrl

`func (o *DrivesResourceInner) GetEnclosuresUrl() string`

GetEnclosuresUrl returns the EnclosuresUrl field if non-nil, zero value otherwise.

### GetEnclosuresUrlOk

`func (o *DrivesResourceInner) GetEnclosuresUrlOk() (*string, bool)`

GetEnclosuresUrlOk returns a tuple with the EnclosuresUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclosuresUrl

`func (o *DrivesResourceInner) SetEnclosuresUrl(v string)`

SetEnclosuresUrl sets EnclosuresUrl field to given value.

### HasEnclosuresUrl

`func (o *DrivesResourceInner) HasEnclosuresUrl() bool`

HasEnclosuresUrl returns a boolean if a field has been set.

### GetError

`func (o *DrivesResourceInner) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DrivesResourceInner) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DrivesResourceInner) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *DrivesResourceInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExtendedStatus

`func (o *DrivesResourceInner) GetExtendedStatus() int32`

GetExtendedStatus returns the ExtendedStatus field if non-nil, zero value otherwise.

### GetExtendedStatusOk

`func (o *DrivesResourceInner) GetExtendedStatusOk() (*int32, bool)`

GetExtendedStatusOk returns a tuple with the ExtendedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedStatus

`func (o *DrivesResourceInner) SetExtendedStatus(v int32)`

SetExtendedStatus sets ExtendedStatus field to given value.

### HasExtendedStatus

`func (o *DrivesResourceInner) HasExtendedStatus() bool`

HasExtendedStatus returns a boolean if a field has been set.

### GetFcP1Channel

`func (o *DrivesResourceInner) GetFcP1Channel() int32`

GetFcP1Channel returns the FcP1Channel field if non-nil, zero value otherwise.

### GetFcP1ChannelOk

`func (o *DrivesResourceInner) GetFcP1ChannelOk() (*int32, bool)`

GetFcP1ChannelOk returns a tuple with the FcP1Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcP1Channel

`func (o *DrivesResourceInner) SetFcP1Channel(v int32)`

SetFcP1Channel sets FcP1Channel field to given value.

### HasFcP1Channel

`func (o *DrivesResourceInner) HasFcP1Channel() bool`

HasFcP1Channel returns a boolean if a field has been set.

### GetFcP1DeviceId

`func (o *DrivesResourceInner) GetFcP1DeviceId() int32`

GetFcP1DeviceId returns the FcP1DeviceId field if non-nil, zero value otherwise.

### GetFcP1DeviceIdOk

`func (o *DrivesResourceInner) GetFcP1DeviceIdOk() (*int32, bool)`

GetFcP1DeviceIdOk returns a tuple with the FcP1DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcP1DeviceId

`func (o *DrivesResourceInner) SetFcP1DeviceId(v int32)`

SetFcP1DeviceId sets FcP1DeviceId field to given value.

### HasFcP1DeviceId

`func (o *DrivesResourceInner) HasFcP1DeviceId() bool`

HasFcP1DeviceId returns a boolean if a field has been set.

### GetFcP1NodeWwn

`func (o *DrivesResourceInner) GetFcP1NodeWwn() string`

GetFcP1NodeWwn returns the FcP1NodeWwn field if non-nil, zero value otherwise.

### GetFcP1NodeWwnOk

`func (o *DrivesResourceInner) GetFcP1NodeWwnOk() (*string, bool)`

GetFcP1NodeWwnOk returns a tuple with the FcP1NodeWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcP1NodeWwn

`func (o *DrivesResourceInner) SetFcP1NodeWwn(v string)`

SetFcP1NodeWwn sets FcP1NodeWwn field to given value.

### HasFcP1NodeWwn

`func (o *DrivesResourceInner) HasFcP1NodeWwn() bool`

HasFcP1NodeWwn returns a boolean if a field has been set.

### GetFcP1PortWwn

`func (o *DrivesResourceInner) GetFcP1PortWwn() string`

GetFcP1PortWwn returns the FcP1PortWwn field if non-nil, zero value otherwise.

### GetFcP1PortWwnOk

`func (o *DrivesResourceInner) GetFcP1PortWwnOk() (*string, bool)`

GetFcP1PortWwnOk returns a tuple with the FcP1PortWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcP1PortWwn

`func (o *DrivesResourceInner) SetFcP1PortWwn(v string)`

SetFcP1PortWwn sets FcP1PortWwn field to given value.

### HasFcP1PortWwn

`func (o *DrivesResourceInner) HasFcP1PortWwn() bool`

HasFcP1PortWwn returns a boolean if a field has been set.

### GetFcP1UnitNumber

`func (o *DrivesResourceInner) GetFcP1UnitNumber() int32`

GetFcP1UnitNumber returns the FcP1UnitNumber field if non-nil, zero value otherwise.

### GetFcP1UnitNumberOk

`func (o *DrivesResourceInner) GetFcP1UnitNumberOk() (*int32, bool)`

GetFcP1UnitNumberOk returns a tuple with the FcP1UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcP1UnitNumber

`func (o *DrivesResourceInner) SetFcP1UnitNumber(v int32)`

SetFcP1UnitNumber sets FcP1UnitNumber field to given value.

### HasFcP1UnitNumber

`func (o *DrivesResourceInner) HasFcP1UnitNumber() bool`

HasFcP1UnitNumber returns a boolean if a field has been set.

### GetFcP2Channel

`func (o *DrivesResourceInner) GetFcP2Channel() int32`

GetFcP2Channel returns the FcP2Channel field if non-nil, zero value otherwise.

### GetFcP2ChannelOk

`func (o *DrivesResourceInner) GetFcP2ChannelOk() (*int32, bool)`

GetFcP2ChannelOk returns a tuple with the FcP2Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcP2Channel

`func (o *DrivesResourceInner) SetFcP2Channel(v int32)`

SetFcP2Channel sets FcP2Channel field to given value.

### HasFcP2Channel

`func (o *DrivesResourceInner) HasFcP2Channel() bool`

HasFcP2Channel returns a boolean if a field has been set.

### GetFcP2DeviceId

`func (o *DrivesResourceInner) GetFcP2DeviceId() int32`

GetFcP2DeviceId returns the FcP2DeviceId field if non-nil, zero value otherwise.

### GetFcP2DeviceIdOk

`func (o *DrivesResourceInner) GetFcP2DeviceIdOk() (*int32, bool)`

GetFcP2DeviceIdOk returns a tuple with the FcP2DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcP2DeviceId

`func (o *DrivesResourceInner) SetFcP2DeviceId(v int32)`

SetFcP2DeviceId sets FcP2DeviceId field to given value.

### HasFcP2DeviceId

`func (o *DrivesResourceInner) HasFcP2DeviceId() bool`

HasFcP2DeviceId returns a boolean if a field has been set.

### GetFcP2NodeWwn

`func (o *DrivesResourceInner) GetFcP2NodeWwn() string`

GetFcP2NodeWwn returns the FcP2NodeWwn field if non-nil, zero value otherwise.

### GetFcP2NodeWwnOk

`func (o *DrivesResourceInner) GetFcP2NodeWwnOk() (*string, bool)`

GetFcP2NodeWwnOk returns a tuple with the FcP2NodeWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcP2NodeWwn

`func (o *DrivesResourceInner) SetFcP2NodeWwn(v string)`

SetFcP2NodeWwn sets FcP2NodeWwn field to given value.

### HasFcP2NodeWwn

`func (o *DrivesResourceInner) HasFcP2NodeWwn() bool`

HasFcP2NodeWwn returns a boolean if a field has been set.

### GetFcP2PortWwn

`func (o *DrivesResourceInner) GetFcP2PortWwn() string`

GetFcP2PortWwn returns the FcP2PortWwn field if non-nil, zero value otherwise.

### GetFcP2PortWwnOk

`func (o *DrivesResourceInner) GetFcP2PortWwnOk() (*string, bool)`

GetFcP2PortWwnOk returns a tuple with the FcP2PortWwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcP2PortWwn

`func (o *DrivesResourceInner) SetFcP2PortWwn(v string)`

SetFcP2PortWwn sets FcP2PortWwn field to given value.

### HasFcP2PortWwn

`func (o *DrivesResourceInner) HasFcP2PortWwn() bool`

HasFcP2PortWwn returns a boolean if a field has been set.

### GetFcP2UnitNumber

`func (o *DrivesResourceInner) GetFcP2UnitNumber() int32`

GetFcP2UnitNumber returns the FcP2UnitNumber field if non-nil, zero value otherwise.

### GetFcP2UnitNumberOk

`func (o *DrivesResourceInner) GetFcP2UnitNumberOk() (*int32, bool)`

GetFcP2UnitNumberOk returns a tuple with the FcP2UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcP2UnitNumber

`func (o *DrivesResourceInner) SetFcP2UnitNumber(v int32)`

SetFcP2UnitNumber sets FcP2UnitNumber field to given value.

### HasFcP2UnitNumber

`func (o *DrivesResourceInner) HasFcP2UnitNumber() bool`

HasFcP2UnitNumber returns a boolean if a field has been set.

### GetFdeConfigTime

`func (o *DrivesResourceInner) GetFdeConfigTime() string`

GetFdeConfigTime returns the FdeConfigTime field if non-nil, zero value otherwise.

### GetFdeConfigTimeOk

`func (o *DrivesResourceInner) GetFdeConfigTimeOk() (*string, bool)`

GetFdeConfigTimeOk returns a tuple with the FdeConfigTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdeConfigTime

`func (o *DrivesResourceInner) SetFdeConfigTime(v string)`

SetFdeConfigTime sets FdeConfigTime field to given value.

### HasFdeConfigTime

`func (o *DrivesResourceInner) HasFdeConfigTime() bool`

HasFdeConfigTime returns a boolean if a field has been set.

### GetFdeConfigTimeNumeric

`func (o *DrivesResourceInner) GetFdeConfigTimeNumeric() int32`

GetFdeConfigTimeNumeric returns the FdeConfigTimeNumeric field if non-nil, zero value otherwise.

### GetFdeConfigTimeNumericOk

`func (o *DrivesResourceInner) GetFdeConfigTimeNumericOk() (*int32, bool)`

GetFdeConfigTimeNumericOk returns a tuple with the FdeConfigTimeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdeConfigTimeNumeric

`func (o *DrivesResourceInner) SetFdeConfigTimeNumeric(v int32)`

SetFdeConfigTimeNumeric sets FdeConfigTimeNumeric field to given value.

### HasFdeConfigTimeNumeric

`func (o *DrivesResourceInner) HasFdeConfigTimeNumeric() bool`

HasFdeConfigTimeNumeric returns a boolean if a field has been set.

### GetFdeState

`func (o *DrivesResourceInner) GetFdeState() string`

GetFdeState returns the FdeState field if non-nil, zero value otherwise.

### GetFdeStateOk

`func (o *DrivesResourceInner) GetFdeStateOk() (*string, bool)`

GetFdeStateOk returns a tuple with the FdeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdeState

`func (o *DrivesResourceInner) SetFdeState(v string)`

SetFdeState sets FdeState field to given value.

### HasFdeState

`func (o *DrivesResourceInner) HasFdeState() bool`

HasFdeState returns a boolean if a field has been set.

### GetFdeStateNumeric

`func (o *DrivesResourceInner) GetFdeStateNumeric() int32`

GetFdeStateNumeric returns the FdeStateNumeric field if non-nil, zero value otherwise.

### GetFdeStateNumericOk

`func (o *DrivesResourceInner) GetFdeStateNumericOk() (*int32, bool)`

GetFdeStateNumericOk returns a tuple with the FdeStateNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdeStateNumeric

`func (o *DrivesResourceInner) SetFdeStateNumeric(v int32)`

SetFdeStateNumeric sets FdeStateNumeric field to given value.

### HasFdeStateNumeric

`func (o *DrivesResourceInner) HasFdeStateNumeric() bool`

HasFdeStateNumeric returns a boolean if a field has been set.

### GetFipsCapable

`func (o *DrivesResourceInner) GetFipsCapable() string`

GetFipsCapable returns the FipsCapable field if non-nil, zero value otherwise.

### GetFipsCapableOk

`func (o *DrivesResourceInner) GetFipsCapableOk() (*string, bool)`

GetFipsCapableOk returns a tuple with the FipsCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsCapable

`func (o *DrivesResourceInner) SetFipsCapable(v string)`

SetFipsCapable sets FipsCapable field to given value.

### HasFipsCapable

`func (o *DrivesResourceInner) HasFipsCapable() bool`

HasFipsCapable returns a boolean if a field has been set.

### GetFipsCapableNumeric

`func (o *DrivesResourceInner) GetFipsCapableNumeric() int32`

GetFipsCapableNumeric returns the FipsCapableNumeric field if non-nil, zero value otherwise.

### GetFipsCapableNumericOk

`func (o *DrivesResourceInner) GetFipsCapableNumericOk() (*int32, bool)`

GetFipsCapableNumericOk returns a tuple with the FipsCapableNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsCapableNumeric

`func (o *DrivesResourceInner) SetFipsCapableNumeric(v int32)`

SetFipsCapableNumeric sets FipsCapableNumeric field to given value.

### HasFipsCapableNumeric

`func (o *DrivesResourceInner) HasFipsCapableNumeric() bool`

HasFipsCapableNumeric returns a boolean if a field has been set.

### GetFirmwareUpdateStatus

`func (o *DrivesResourceInner) GetFirmwareUpdateStatus() string`

GetFirmwareUpdateStatus returns the FirmwareUpdateStatus field if non-nil, zero value otherwise.

### GetFirmwareUpdateStatusOk

`func (o *DrivesResourceInner) GetFirmwareUpdateStatusOk() (*string, bool)`

GetFirmwareUpdateStatusOk returns a tuple with the FirmwareUpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareUpdateStatus

`func (o *DrivesResourceInner) SetFirmwareUpdateStatus(v string)`

SetFirmwareUpdateStatus sets FirmwareUpdateStatus field to given value.

### HasFirmwareUpdateStatus

`func (o *DrivesResourceInner) HasFirmwareUpdateStatus() bool`

HasFirmwareUpdateStatus returns a boolean if a field has been set.

### GetFirmwareUpdateStatusNumeric

`func (o *DrivesResourceInner) GetFirmwareUpdateStatusNumeric() int32`

GetFirmwareUpdateStatusNumeric returns the FirmwareUpdateStatusNumeric field if non-nil, zero value otherwise.

### GetFirmwareUpdateStatusNumericOk

`func (o *DrivesResourceInner) GetFirmwareUpdateStatusNumericOk() (*int32, bool)`

GetFirmwareUpdateStatusNumericOk returns a tuple with the FirmwareUpdateStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareUpdateStatusNumeric

`func (o *DrivesResourceInner) SetFirmwareUpdateStatusNumeric(v int32)`

SetFirmwareUpdateStatusNumeric sets FirmwareUpdateStatusNumeric field to given value.

### HasFirmwareUpdateStatusNumeric

`func (o *DrivesResourceInner) HasFirmwareUpdateStatusNumeric() bool`

HasFirmwareUpdateStatusNumeric returns a boolean if a field has been set.

### GetHealth

`func (o *DrivesResourceInner) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *DrivesResourceInner) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *DrivesResourceInner) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *DrivesResourceInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHealthNumeric

`func (o *DrivesResourceInner) GetHealthNumeric() int32`

GetHealthNumeric returns the HealthNumeric field if non-nil, zero value otherwise.

### GetHealthNumericOk

`func (o *DrivesResourceInner) GetHealthNumericOk() (*int32, bool)`

GetHealthNumericOk returns a tuple with the HealthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthNumeric

`func (o *DrivesResourceInner) SetHealthNumeric(v int32)`

SetHealthNumeric sets HealthNumeric field to given value.

### HasHealthNumeric

`func (o *DrivesResourceInner) HasHealthNumeric() bool`

HasHealthNumeric returns a boolean if a field has been set.

### GetHealthReason

`func (o *DrivesResourceInner) GetHealthReason() string`

GetHealthReason returns the HealthReason field if non-nil, zero value otherwise.

### GetHealthReasonOk

`func (o *DrivesResourceInner) GetHealthReasonOk() (*string, bool)`

GetHealthReasonOk returns a tuple with the HealthReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReason

`func (o *DrivesResourceInner) SetHealthReason(v string)`

SetHealthReason sets HealthReason field to given value.

### HasHealthReason

`func (o *DrivesResourceInner) HasHealthReason() bool`

HasHealthReason returns a boolean if a field has been set.

### GetHealthReasonNumeric

`func (o *DrivesResourceInner) GetHealthReasonNumeric() int32`

GetHealthReasonNumeric returns the HealthReasonNumeric field if non-nil, zero value otherwise.

### GetHealthReasonNumericOk

`func (o *DrivesResourceInner) GetHealthReasonNumericOk() (*int32, bool)`

GetHealthReasonNumericOk returns a tuple with the HealthReasonNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReasonNumeric

`func (o *DrivesResourceInner) SetHealthReasonNumeric(v int32)`

SetHealthReasonNumeric sets HealthReasonNumeric field to given value.

### HasHealthReasonNumeric

`func (o *DrivesResourceInner) HasHealthReasonNumeric() bool`

HasHealthReasonNumeric returns a boolean if a field has been set.

### GetHealthRecommendation

`func (o *DrivesResourceInner) GetHealthRecommendation() string`

GetHealthRecommendation returns the HealthRecommendation field if non-nil, zero value otherwise.

### GetHealthRecommendationOk

`func (o *DrivesResourceInner) GetHealthRecommendationOk() (*string, bool)`

GetHealthRecommendationOk returns a tuple with the HealthRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecommendation

`func (o *DrivesResourceInner) SetHealthRecommendation(v string)`

SetHealthRecommendation sets HealthRecommendation field to given value.

### HasHealthRecommendation

`func (o *DrivesResourceInner) HasHealthRecommendation() bool`

HasHealthRecommendation returns a boolean if a field has been set.

### GetHealthRecommendationNumeric

`func (o *DrivesResourceInner) GetHealthRecommendationNumeric() int32`

GetHealthRecommendationNumeric returns the HealthRecommendationNumeric field if non-nil, zero value otherwise.

### GetHealthRecommendationNumericOk

`func (o *DrivesResourceInner) GetHealthRecommendationNumericOk() (*int32, bool)`

GetHealthRecommendationNumericOk returns a tuple with the HealthRecommendationNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecommendationNumeric

`func (o *DrivesResourceInner) SetHealthRecommendationNumeric(v int32)`

SetHealthRecommendationNumeric sets HealthRecommendationNumeric field to given value.

### HasHealthRecommendationNumeric

`func (o *DrivesResourceInner) HasHealthRecommendationNumeric() bool`

HasHealthRecommendationNumeric returns a boolean if a field has been set.

### GetImportLockKeyId

`func (o *DrivesResourceInner) GetImportLockKeyId() string`

GetImportLockKeyId returns the ImportLockKeyId field if non-nil, zero value otherwise.

### GetImportLockKeyIdOk

`func (o *DrivesResourceInner) GetImportLockKeyIdOk() (*string, bool)`

GetImportLockKeyIdOk returns a tuple with the ImportLockKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportLockKeyId

`func (o *DrivesResourceInner) SetImportLockKeyId(v string)`

SetImportLockKeyId sets ImportLockKeyId field to given value.

### HasImportLockKeyId

`func (o *DrivesResourceInner) HasImportLockKeyId() bool`

HasImportLockKeyId returns a boolean if a field has been set.

### GetIndex

`func (o *DrivesResourceInner) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *DrivesResourceInner) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *DrivesResourceInner) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *DrivesResourceInner) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetInterface

`func (o *DrivesResourceInner) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *DrivesResourceInner) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *DrivesResourceInner) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *DrivesResourceInner) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetInterfaceNumeric

`func (o *DrivesResourceInner) GetInterfaceNumeric() int32`

GetInterfaceNumeric returns the InterfaceNumeric field if non-nil, zero value otherwise.

### GetInterfaceNumericOk

`func (o *DrivesResourceInner) GetInterfaceNumericOk() (*int32, bool)`

GetInterfaceNumericOk returns a tuple with the InterfaceNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceNumeric

`func (o *DrivesResourceInner) SetInterfaceNumeric(v int32)`

SetInterfaceNumeric sets InterfaceNumeric field to given value.

### HasInterfaceNumeric

`func (o *DrivesResourceInner) HasInterfaceNumeric() bool`

HasInterfaceNumeric returns a boolean if a field has been set.

### GetJobRunning

`func (o *DrivesResourceInner) GetJobRunning() string`

GetJobRunning returns the JobRunning field if non-nil, zero value otherwise.

### GetJobRunningOk

`func (o *DrivesResourceInner) GetJobRunningOk() (*string, bool)`

GetJobRunningOk returns a tuple with the JobRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunning

`func (o *DrivesResourceInner) SetJobRunning(v string)`

SetJobRunning sets JobRunning field to given value.

### HasJobRunning

`func (o *DrivesResourceInner) HasJobRunning() bool`

HasJobRunning returns a boolean if a field has been set.

### GetJobRunningNumeric

`func (o *DrivesResourceInner) GetJobRunningNumeric() int32`

GetJobRunningNumeric returns the JobRunningNumeric field if non-nil, zero value otherwise.

### GetJobRunningNumericOk

`func (o *DrivesResourceInner) GetJobRunningNumericOk() (*int32, bool)`

GetJobRunningNumericOk returns a tuple with the JobRunningNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRunningNumeric

`func (o *DrivesResourceInner) SetJobRunningNumeric(v int32)`

SetJobRunningNumeric sets JobRunningNumeric field to given value.

### HasJobRunningNumeric

`func (o *DrivesResourceInner) HasJobRunningNumeric() bool`

HasJobRunningNumeric returns a boolean if a field has been set.

### GetKmipState

`func (o *DrivesResourceInner) GetKmipState() string`

GetKmipState returns the KmipState field if non-nil, zero value otherwise.

### GetKmipStateOk

`func (o *DrivesResourceInner) GetKmipStateOk() (*string, bool)`

GetKmipStateOk returns a tuple with the KmipState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmipState

`func (o *DrivesResourceInner) SetKmipState(v string)`

SetKmipState sets KmipState field to given value.

### HasKmipState

`func (o *DrivesResourceInner) HasKmipState() bool`

HasKmipState returns a boolean if a field has been set.

### GetKmipStateNumeric

`func (o *DrivesResourceInner) GetKmipStateNumeric() int32`

GetKmipStateNumeric returns the KmipStateNumeric field if non-nil, zero value otherwise.

### GetKmipStateNumericOk

`func (o *DrivesResourceInner) GetKmipStateNumericOk() (*int32, bool)`

GetKmipStateNumericOk returns a tuple with the KmipStateNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmipStateNumeric

`func (o *DrivesResourceInner) SetKmipStateNumeric(v int32)`

SetKmipStateNumeric sets KmipStateNumeric field to given value.

### HasKmipStateNumeric

`func (o *DrivesResourceInner) HasKmipStateNumeric() bool`

HasKmipStateNumeric returns a boolean if a field has been set.

### GetLedStatus

`func (o *DrivesResourceInner) GetLedStatus() string`

GetLedStatus returns the LedStatus field if non-nil, zero value otherwise.

### GetLedStatusOk

`func (o *DrivesResourceInner) GetLedStatusOk() (*string, bool)`

GetLedStatusOk returns a tuple with the LedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedStatus

`func (o *DrivesResourceInner) SetLedStatus(v string)`

SetLedStatus sets LedStatus field to given value.

### HasLedStatus

`func (o *DrivesResourceInner) HasLedStatus() bool`

HasLedStatus returns a boolean if a field has been set.

### GetLedStatusNumeric

`func (o *DrivesResourceInner) GetLedStatusNumeric() int32`

GetLedStatusNumeric returns the LedStatusNumeric field if non-nil, zero value otherwise.

### GetLedStatusNumericOk

`func (o *DrivesResourceInner) GetLedStatusNumericOk() (*int32, bool)`

GetLedStatusNumericOk returns a tuple with the LedStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedStatusNumeric

`func (o *DrivesResourceInner) SetLedStatusNumeric(v int32)`

SetLedStatusNumeric sets LedStatusNumeric field to given value.

### HasLedStatusNumeric

`func (o *DrivesResourceInner) HasLedStatusNumeric() bool`

HasLedStatusNumeric returns a boolean if a field has been set.

### GetLocation

`func (o *DrivesResourceInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DrivesResourceInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DrivesResourceInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DrivesResourceInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLocatorLed

`func (o *DrivesResourceInner) GetLocatorLed() string`

GetLocatorLed returns the LocatorLed field if non-nil, zero value otherwise.

### GetLocatorLedOk

`func (o *DrivesResourceInner) GetLocatorLedOk() (*string, bool)`

GetLocatorLedOk returns a tuple with the LocatorLed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLed

`func (o *DrivesResourceInner) SetLocatorLed(v string)`

SetLocatorLed sets LocatorLed field to given value.

### HasLocatorLed

`func (o *DrivesResourceInner) HasLocatorLed() bool`

HasLocatorLed returns a boolean if a field has been set.

### GetLocatorLedNumeric

`func (o *DrivesResourceInner) GetLocatorLedNumeric() int32`

GetLocatorLedNumeric returns the LocatorLedNumeric field if non-nil, zero value otherwise.

### GetLocatorLedNumericOk

`func (o *DrivesResourceInner) GetLocatorLedNumericOk() (*int32, bool)`

GetLocatorLedNumericOk returns a tuple with the LocatorLedNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocatorLedNumeric

`func (o *DrivesResourceInner) SetLocatorLedNumeric(v int32)`

SetLocatorLedNumeric sets LocatorLedNumeric field to given value.

### HasLocatorLedNumeric

`func (o *DrivesResourceInner) HasLocatorLedNumeric() bool`

HasLocatorLedNumeric returns a boolean if a field has been set.

### GetLockKeyId

`func (o *DrivesResourceInner) GetLockKeyId() string`

GetLockKeyId returns the LockKeyId field if non-nil, zero value otherwise.

### GetLockKeyIdOk

`func (o *DrivesResourceInner) GetLockKeyIdOk() (*string, bool)`

GetLockKeyIdOk returns a tuple with the LockKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockKeyId

`func (o *DrivesResourceInner) SetLockKeyId(v string)`

SetLockKeyId sets LockKeyId field to given value.

### HasLockKeyId

`func (o *DrivesResourceInner) HasLockKeyId() bool`

HasLockKeyId returns a boolean if a field has been set.

### GetMemberIndex

`func (o *DrivesResourceInner) GetMemberIndex() int32`

GetMemberIndex returns the MemberIndex field if non-nil, zero value otherwise.

### GetMemberIndexOk

`func (o *DrivesResourceInner) GetMemberIndexOk() (*int32, bool)`

GetMemberIndexOk returns a tuple with the MemberIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIndex

`func (o *DrivesResourceInner) SetMemberIndex(v int32)`

SetMemberIndex sets MemberIndex field to given value.

### HasMemberIndex

`func (o *DrivesResourceInner) HasMemberIndex() bool`

HasMemberIndex returns a boolean if a field has been set.

### GetModel

`func (o *DrivesResourceInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DrivesResourceInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DrivesResourceInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DrivesResourceInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNumDepoppedHeads

`func (o *DrivesResourceInner) GetNumDepoppedHeads() int32`

GetNumDepoppedHeads returns the NumDepoppedHeads field if non-nil, zero value otherwise.

### GetNumDepoppedHeadsOk

`func (o *DrivesResourceInner) GetNumDepoppedHeadsOk() (*int32, bool)`

GetNumDepoppedHeadsOk returns a tuple with the NumDepoppedHeads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDepoppedHeads

`func (o *DrivesResourceInner) SetNumDepoppedHeads(v int32)`

SetNumDepoppedHeads sets NumDepoppedHeads field to given value.

### HasNumDepoppedHeads

`func (o *DrivesResourceInner) HasNumDepoppedHeads() bool`

HasNumDepoppedHeads returns a boolean if a field has been set.

### GetNumberOfIos

`func (o *DrivesResourceInner) GetNumberOfIos() int32`

GetNumberOfIos returns the NumberOfIos field if non-nil, zero value otherwise.

### GetNumberOfIosOk

`func (o *DrivesResourceInner) GetNumberOfIosOk() (*int32, bool)`

GetNumberOfIosOk returns a tuple with the NumberOfIos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfIos

`func (o *DrivesResourceInner) SetNumberOfIos(v int32)`

SetNumberOfIos sets NumberOfIos field to given value.

### HasNumberOfIos

`func (o *DrivesResourceInner) HasNumberOfIos() bool`

HasNumberOfIos returns a boolean if a field has been set.

### GetOwner

`func (o *DrivesResourceInner) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DrivesResourceInner) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DrivesResourceInner) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DrivesResourceInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnerNumeric

`func (o *DrivesResourceInner) GetOwnerNumeric() int32`

GetOwnerNumeric returns the OwnerNumeric field if non-nil, zero value otherwise.

### GetOwnerNumericOk

`func (o *DrivesResourceInner) GetOwnerNumericOk() (*int32, bool)`

GetOwnerNumericOk returns a tuple with the OwnerNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerNumeric

`func (o *DrivesResourceInner) SetOwnerNumeric(v int32)`

SetOwnerNumeric sets OwnerNumeric field to given value.

### HasOwnerNumeric

`func (o *DrivesResourceInner) HasOwnerNumeric() bool`

HasOwnerNumeric returns a boolean if a field has been set.

### GetPiFormatted

`func (o *DrivesResourceInner) GetPiFormatted() string`

GetPiFormatted returns the PiFormatted field if non-nil, zero value otherwise.

### GetPiFormattedOk

`func (o *DrivesResourceInner) GetPiFormattedOk() (*string, bool)`

GetPiFormattedOk returns a tuple with the PiFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiFormatted

`func (o *DrivesResourceInner) SetPiFormatted(v string)`

SetPiFormatted sets PiFormatted field to given value.

### HasPiFormatted

`func (o *DrivesResourceInner) HasPiFormatted() bool`

HasPiFormatted returns a boolean if a field has been set.

### GetPiFormattedNumeric

`func (o *DrivesResourceInner) GetPiFormattedNumeric() int32`

GetPiFormattedNumeric returns the PiFormattedNumeric field if non-nil, zero value otherwise.

### GetPiFormattedNumericOk

`func (o *DrivesResourceInner) GetPiFormattedNumericOk() (*int32, bool)`

GetPiFormattedNumericOk returns a tuple with the PiFormattedNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiFormattedNumeric

`func (o *DrivesResourceInner) SetPiFormattedNumeric(v int32)`

SetPiFormattedNumeric sets PiFormattedNumeric field to given value.

### HasPiFormattedNumeric

`func (o *DrivesResourceInner) HasPiFormattedNumeric() bool`

HasPiFormattedNumeric returns a boolean if a field has been set.

### GetPort

`func (o *DrivesResourceInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DrivesResourceInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DrivesResourceInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DrivesResourceInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPowerOnHours

`func (o *DrivesResourceInner) GetPowerOnHours() int32`

GetPowerOnHours returns the PowerOnHours field if non-nil, zero value otherwise.

### GetPowerOnHoursOk

`func (o *DrivesResourceInner) GetPowerOnHoursOk() (*int32, bool)`

GetPowerOnHoursOk returns a tuple with the PowerOnHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOnHours

`func (o *DrivesResourceInner) SetPowerOnHours(v int32)`

SetPowerOnHours sets PowerOnHours field to given value.

### HasPowerOnHours

`func (o *DrivesResourceInner) HasPowerOnHours() bool`

HasPowerOnHours returns a boolean if a field has been set.

### GetReconState

`func (o *DrivesResourceInner) GetReconState() string`

GetReconState returns the ReconState field if non-nil, zero value otherwise.

### GetReconStateOk

`func (o *DrivesResourceInner) GetReconStateOk() (*string, bool)`

GetReconStateOk returns a tuple with the ReconState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconState

`func (o *DrivesResourceInner) SetReconState(v string)`

SetReconState sets ReconState field to given value.

### HasReconState

`func (o *DrivesResourceInner) HasReconState() bool`

HasReconState returns a boolean if a field has been set.

### GetReconStateNumeric

`func (o *DrivesResourceInner) GetReconStateNumeric() int32`

GetReconStateNumeric returns the ReconStateNumeric field if non-nil, zero value otherwise.

### GetReconStateNumericOk

`func (o *DrivesResourceInner) GetReconStateNumericOk() (*int32, bool)`

GetReconStateNumericOk returns a tuple with the ReconStateNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconStateNumeric

`func (o *DrivesResourceInner) SetReconStateNumeric(v int32)`

SetReconStateNumeric sets ReconStateNumeric field to given value.

### HasReconStateNumeric

`func (o *DrivesResourceInner) HasReconStateNumeric() bool`

HasReconStateNumeric returns a boolean if a field has been set.

### GetRemanufacture

`func (o *DrivesResourceInner) GetRemanufacture() string`

GetRemanufacture returns the Remanufacture field if non-nil, zero value otherwise.

### GetRemanufactureOk

`func (o *DrivesResourceInner) GetRemanufactureOk() (*string, bool)`

GetRemanufactureOk returns a tuple with the Remanufacture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemanufacture

`func (o *DrivesResourceInner) SetRemanufacture(v string)`

SetRemanufacture sets Remanufacture field to given value.

### HasRemanufacture

`func (o *DrivesResourceInner) HasRemanufacture() bool`

HasRemanufacture returns a boolean if a field has been set.

### GetRemanufactureNumeric

`func (o *DrivesResourceInner) GetRemanufactureNumeric() int32`

GetRemanufactureNumeric returns the RemanufactureNumeric field if non-nil, zero value otherwise.

### GetRemanufactureNumericOk

`func (o *DrivesResourceInner) GetRemanufactureNumericOk() (*int32, bool)`

GetRemanufactureNumericOk returns a tuple with the RemanufactureNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemanufactureNumeric

`func (o *DrivesResourceInner) SetRemanufactureNumeric(v int32)`

SetRemanufactureNumeric sets RemanufactureNumeric field to given value.

### HasRemanufactureNumeric

`func (o *DrivesResourceInner) HasRemanufactureNumeric() bool`

HasRemanufactureNumeric returns a boolean if a field has been set.

### GetRevision

`func (o *DrivesResourceInner) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DrivesResourceInner) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DrivesResourceInner) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *DrivesResourceInner) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRpm

`func (o *DrivesResourceInner) GetRpm() int32`

GetRpm returns the Rpm field if non-nil, zero value otherwise.

### GetRpmOk

`func (o *DrivesResourceInner) GetRpmOk() (*int32, bool)`

GetRpmOk returns a tuple with the Rpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpm

`func (o *DrivesResourceInner) SetRpm(v int32)`

SetRpm sets Rpm field to given value.

### HasRpm

`func (o *DrivesResourceInner) HasRpm() bool`

HasRpm returns a boolean if a field has been set.

### GetScsiId

`func (o *DrivesResourceInner) GetScsiId() int32`

GetScsiId returns the ScsiId field if non-nil, zero value otherwise.

### GetScsiIdOk

`func (o *DrivesResourceInner) GetScsiIdOk() (*int32, bool)`

GetScsiIdOk returns a tuple with the ScsiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsiId

`func (o *DrivesResourceInner) SetScsiId(v int32)`

SetScsiId sets ScsiId field to given value.

### HasScsiId

`func (o *DrivesResourceInner) HasScsiId() bool`

HasScsiId returns a boolean if a field has been set.

### GetSecondaryChannel

`func (o *DrivesResourceInner) GetSecondaryChannel() int32`

GetSecondaryChannel returns the SecondaryChannel field if non-nil, zero value otherwise.

### GetSecondaryChannelOk

`func (o *DrivesResourceInner) GetSecondaryChannelOk() (*int32, bool)`

GetSecondaryChannelOk returns a tuple with the SecondaryChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryChannel

`func (o *DrivesResourceInner) SetSecondaryChannel(v int32)`

SetSecondaryChannel sets SecondaryChannel field to given value.

### HasSecondaryChannel

`func (o *DrivesResourceInner) HasSecondaryChannel() bool`

HasSecondaryChannel returns a boolean if a field has been set.

### GetSectorFormat

`func (o *DrivesResourceInner) GetSectorFormat() string`

GetSectorFormat returns the SectorFormat field if non-nil, zero value otherwise.

### GetSectorFormatOk

`func (o *DrivesResourceInner) GetSectorFormatOk() (*string, bool)`

GetSectorFormatOk returns a tuple with the SectorFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorFormat

`func (o *DrivesResourceInner) SetSectorFormat(v string)`

SetSectorFormat sets SectorFormat field to given value.

### HasSectorFormat

`func (o *DrivesResourceInner) HasSectorFormat() bool`

HasSectorFormat returns a boolean if a field has been set.

### GetSectorFormatNumeric

`func (o *DrivesResourceInner) GetSectorFormatNumeric() int32`

GetSectorFormatNumeric returns the SectorFormatNumeric field if non-nil, zero value otherwise.

### GetSectorFormatNumericOk

`func (o *DrivesResourceInner) GetSectorFormatNumericOk() (*int32, bool)`

GetSectorFormatNumericOk returns a tuple with the SectorFormatNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorFormatNumeric

`func (o *DrivesResourceInner) SetSectorFormatNumeric(v int32)`

SetSectorFormatNumeric sets SectorFormatNumeric field to given value.

### HasSectorFormatNumeric

`func (o *DrivesResourceInner) HasSectorFormatNumeric() bool`

HasSectorFormatNumeric returns a boolean if a field has been set.

### GetSerialNumber

`func (o *DrivesResourceInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DrivesResourceInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DrivesResourceInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DrivesResourceInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSinglePorted

`func (o *DrivesResourceInner) GetSinglePorted() string`

GetSinglePorted returns the SinglePorted field if non-nil, zero value otherwise.

### GetSinglePortedOk

`func (o *DrivesResourceInner) GetSinglePortedOk() (*string, bool)`

GetSinglePortedOk returns a tuple with the SinglePorted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinglePorted

`func (o *DrivesResourceInner) SetSinglePorted(v string)`

SetSinglePorted sets SinglePorted field to given value.

### HasSinglePorted

`func (o *DrivesResourceInner) HasSinglePorted() bool`

HasSinglePorted returns a boolean if a field has been set.

### GetSinglePortedNumeric

`func (o *DrivesResourceInner) GetSinglePortedNumeric() int32`

GetSinglePortedNumeric returns the SinglePortedNumeric field if non-nil, zero value otherwise.

### GetSinglePortedNumericOk

`func (o *DrivesResourceInner) GetSinglePortedNumericOk() (*int32, bool)`

GetSinglePortedNumericOk returns a tuple with the SinglePortedNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinglePortedNumeric

`func (o *DrivesResourceInner) SetSinglePortedNumeric(v int32)`

SetSinglePortedNumeric sets SinglePortedNumeric field to given value.

### HasSinglePortedNumeric

`func (o *DrivesResourceInner) HasSinglePortedNumeric() bool`

HasSinglePortedNumeric returns a boolean if a field has been set.

### GetSize

`func (o *DrivesResourceInner) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DrivesResourceInner) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DrivesResourceInner) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *DrivesResourceInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeNumeric

`func (o *DrivesResourceInner) GetSizeNumeric() int32`

GetSizeNumeric returns the SizeNumeric field if non-nil, zero value otherwise.

### GetSizeNumericOk

`func (o *DrivesResourceInner) GetSizeNumericOk() (*int32, bool)`

GetSizeNumericOk returns a tuple with the SizeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeNumeric

`func (o *DrivesResourceInner) SetSizeNumeric(v int32)`

SetSizeNumeric sets SizeNumeric field to given value.

### HasSizeNumeric

`func (o *DrivesResourceInner) HasSizeNumeric() bool`

HasSizeNumeric returns a boolean if a field has been set.

### GetSlot

`func (o *DrivesResourceInner) GetSlot() int32`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *DrivesResourceInner) GetSlotOk() (*int32, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *DrivesResourceInner) SetSlot(v int32)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *DrivesResourceInner) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetSmart

`func (o *DrivesResourceInner) GetSmart() string`

GetSmart returns the Smart field if non-nil, zero value otherwise.

### GetSmartOk

`func (o *DrivesResourceInner) GetSmartOk() (*string, bool)`

GetSmartOk returns a tuple with the Smart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmart

`func (o *DrivesResourceInner) SetSmart(v string)`

SetSmart sets Smart field to given value.

### HasSmart

`func (o *DrivesResourceInner) HasSmart() bool`

HasSmart returns a boolean if a field has been set.

### GetSmartNumeric

`func (o *DrivesResourceInner) GetSmartNumeric() int32`

GetSmartNumeric returns the SmartNumeric field if non-nil, zero value otherwise.

### GetSmartNumericOk

`func (o *DrivesResourceInner) GetSmartNumericOk() (*int32, bool)`

GetSmartNumericOk returns a tuple with the SmartNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartNumeric

`func (o *DrivesResourceInner) SetSmartNumeric(v int32)`

SetSmartNumeric sets SmartNumeric field to given value.

### HasSmartNumeric

`func (o *DrivesResourceInner) HasSmartNumeric() bool`

HasSmartNumeric returns a boolean if a field has been set.

### GetSpeed

`func (o *DrivesResourceInner) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *DrivesResourceInner) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *DrivesResourceInner) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *DrivesResourceInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetSpunDown

`func (o *DrivesResourceInner) GetSpunDown() int32`

GetSpunDown returns the SpunDown field if non-nil, zero value otherwise.

### GetSpunDownOk

`func (o *DrivesResourceInner) GetSpunDownOk() (*int32, bool)`

GetSpunDownOk returns a tuple with the SpunDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpunDown

`func (o *DrivesResourceInner) SetSpunDown(v int32)`

SetSpunDown sets SpunDown field to given value.

### HasSpunDown

`func (o *DrivesResourceInner) HasSpunDown() bool`

HasSpunDown returns a boolean if a field has been set.

### GetSsdLifeLeft

`func (o *DrivesResourceInner) GetSsdLifeLeft() string`

GetSsdLifeLeft returns the SsdLifeLeft field if non-nil, zero value otherwise.

### GetSsdLifeLeftOk

`func (o *DrivesResourceInner) GetSsdLifeLeftOk() (*string, bool)`

GetSsdLifeLeftOk returns a tuple with the SsdLifeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdLifeLeft

`func (o *DrivesResourceInner) SetSsdLifeLeft(v string)`

SetSsdLifeLeft sets SsdLifeLeft field to given value.

### HasSsdLifeLeft

`func (o *DrivesResourceInner) HasSsdLifeLeft() bool`

HasSsdLifeLeft returns a boolean if a field has been set.

### GetSsdLifeLeftNumeric

`func (o *DrivesResourceInner) GetSsdLifeLeftNumeric() int32`

GetSsdLifeLeftNumeric returns the SsdLifeLeftNumeric field if non-nil, zero value otherwise.

### GetSsdLifeLeftNumericOk

`func (o *DrivesResourceInner) GetSsdLifeLeftNumericOk() (*int32, bool)`

GetSsdLifeLeftNumericOk returns a tuple with the SsdLifeLeftNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdLifeLeftNumeric

`func (o *DrivesResourceInner) SetSsdLifeLeftNumeric(v int32)`

SetSsdLifeLeftNumeric sets SsdLifeLeftNumeric field to given value.

### HasSsdLifeLeftNumeric

`func (o *DrivesResourceInner) HasSsdLifeLeftNumeric() bool`

HasSsdLifeLeftNumeric returns a boolean if a field has been set.

### GetState

`func (o *DrivesResourceInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DrivesResourceInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DrivesResourceInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DrivesResourceInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *DrivesResourceInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DrivesResourceInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DrivesResourceInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DrivesResourceInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStoragePoolName

`func (o *DrivesResourceInner) GetStoragePoolName() string`

GetStoragePoolName returns the StoragePoolName field if non-nil, zero value otherwise.

### GetStoragePoolNameOk

`func (o *DrivesResourceInner) GetStoragePoolNameOk() (*string, bool)`

GetStoragePoolNameOk returns a tuple with the StoragePoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolName

`func (o *DrivesResourceInner) SetStoragePoolName(v string)`

SetStoragePoolName sets StoragePoolName field to given value.

### HasStoragePoolName

`func (o *DrivesResourceInner) HasStoragePoolName() bool`

HasStoragePoolName returns a boolean if a field has been set.

### GetStorageTier

`func (o *DrivesResourceInner) GetStorageTier() string`

GetStorageTier returns the StorageTier field if non-nil, zero value otherwise.

### GetStorageTierOk

`func (o *DrivesResourceInner) GetStorageTierOk() (*string, bool)`

GetStorageTierOk returns a tuple with the StorageTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTier

`func (o *DrivesResourceInner) SetStorageTier(v string)`

SetStorageTier sets StorageTier field to given value.

### HasStorageTier

`func (o *DrivesResourceInner) HasStorageTier() bool`

HasStorageTier returns a boolean if a field has been set.

### GetStorageTierNumeric

`func (o *DrivesResourceInner) GetStorageTierNumeric() int32`

GetStorageTierNumeric returns the StorageTierNumeric field if non-nil, zero value otherwise.

### GetStorageTierNumericOk

`func (o *DrivesResourceInner) GetStorageTierNumericOk() (*int32, bool)`

GetStorageTierNumericOk returns a tuple with the StorageTierNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageTierNumeric

`func (o *DrivesResourceInner) SetStorageTierNumeric(v int32)`

SetStorageTierNumeric sets StorageTierNumeric field to given value.

### HasStorageTierNumeric

`func (o *DrivesResourceInner) HasStorageTierNumeric() bool`

HasStorageTierNumeric returns a boolean if a field has been set.

### GetSupportsUnmap

`func (o *DrivesResourceInner) GetSupportsUnmap() string`

GetSupportsUnmap returns the SupportsUnmap field if non-nil, zero value otherwise.

### GetSupportsUnmapOk

`func (o *DrivesResourceInner) GetSupportsUnmapOk() (*string, bool)`

GetSupportsUnmapOk returns a tuple with the SupportsUnmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsUnmap

`func (o *DrivesResourceInner) SetSupportsUnmap(v string)`

SetSupportsUnmap sets SupportsUnmap field to given value.

### HasSupportsUnmap

`func (o *DrivesResourceInner) HasSupportsUnmap() bool`

HasSupportsUnmap returns a boolean if a field has been set.

### GetSupportsUnmapNumeric

`func (o *DrivesResourceInner) GetSupportsUnmapNumeric() int32`

GetSupportsUnmapNumeric returns the SupportsUnmapNumeric field if non-nil, zero value otherwise.

### GetSupportsUnmapNumericOk

`func (o *DrivesResourceInner) GetSupportsUnmapNumericOk() (*int32, bool)`

GetSupportsUnmapNumericOk returns a tuple with the SupportsUnmapNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsUnmapNumeric

`func (o *DrivesResourceInner) SetSupportsUnmapNumeric(v int32)`

SetSupportsUnmapNumeric sets SupportsUnmapNumeric field to given value.

### HasSupportsUnmapNumeric

`func (o *DrivesResourceInner) HasSupportsUnmapNumeric() bool`

HasSupportsUnmapNumeric returns a boolean if a field has been set.

### GetTemperature

`func (o *DrivesResourceInner) GetTemperature() string`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *DrivesResourceInner) GetTemperatureOk() (*string, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *DrivesResourceInner) SetTemperature(v string)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *DrivesResourceInner) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTemperatureNumeric

`func (o *DrivesResourceInner) GetTemperatureNumeric() int32`

GetTemperatureNumeric returns the TemperatureNumeric field if non-nil, zero value otherwise.

### GetTemperatureNumericOk

`func (o *DrivesResourceInner) GetTemperatureNumericOk() (*int32, bool)`

GetTemperatureNumericOk returns a tuple with the TemperatureNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureNumeric

`func (o *DrivesResourceInner) SetTemperatureNumeric(v int32)`

SetTemperatureNumeric sets TemperatureNumeric field to given value.

### HasTemperatureNumeric

`func (o *DrivesResourceInner) HasTemperatureNumeric() bool`

HasTemperatureNumeric returns a boolean if a field has been set.

### GetTemperatureStatus

`func (o *DrivesResourceInner) GetTemperatureStatus() string`

GetTemperatureStatus returns the TemperatureStatus field if non-nil, zero value otherwise.

### GetTemperatureStatusOk

`func (o *DrivesResourceInner) GetTemperatureStatusOk() (*string, bool)`

GetTemperatureStatusOk returns a tuple with the TemperatureStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureStatus

`func (o *DrivesResourceInner) SetTemperatureStatus(v string)`

SetTemperatureStatus sets TemperatureStatus field to given value.

### HasTemperatureStatus

`func (o *DrivesResourceInner) HasTemperatureStatus() bool`

HasTemperatureStatus returns a boolean if a field has been set.

### GetTemperatureStatusNumeric

`func (o *DrivesResourceInner) GetTemperatureStatusNumeric() int32`

GetTemperatureStatusNumeric returns the TemperatureStatusNumeric field if non-nil, zero value otherwise.

### GetTemperatureStatusNumericOk

`func (o *DrivesResourceInner) GetTemperatureStatusNumericOk() (*int32, bool)`

GetTemperatureStatusNumericOk returns a tuple with the TemperatureStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureStatusNumeric

`func (o *DrivesResourceInner) SetTemperatureStatusNumeric(v int32)`

SetTemperatureStatusNumeric sets TemperatureStatusNumeric field to given value.

### HasTemperatureStatusNumeric

`func (o *DrivesResourceInner) HasTemperatureStatusNumeric() bool`

HasTemperatureStatusNumeric returns a boolean if a field has been set.

### GetTotalDataTransferred

`func (o *DrivesResourceInner) GetTotalDataTransferred() string`

GetTotalDataTransferred returns the TotalDataTransferred field if non-nil, zero value otherwise.

### GetTotalDataTransferredOk

`func (o *DrivesResourceInner) GetTotalDataTransferredOk() (*string, bool)`

GetTotalDataTransferredOk returns a tuple with the TotalDataTransferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDataTransferred

`func (o *DrivesResourceInner) SetTotalDataTransferred(v string)`

SetTotalDataTransferred sets TotalDataTransferred field to given value.

### HasTotalDataTransferred

`func (o *DrivesResourceInner) HasTotalDataTransferred() bool`

HasTotalDataTransferred returns a boolean if a field has been set.

### GetTotalDataTransferredNumeric

`func (o *DrivesResourceInner) GetTotalDataTransferredNumeric() int32`

GetTotalDataTransferredNumeric returns the TotalDataTransferredNumeric field if non-nil, zero value otherwise.

### GetTotalDataTransferredNumericOk

`func (o *DrivesResourceInner) GetTotalDataTransferredNumericOk() (*int32, bool)`

GetTotalDataTransferredNumericOk returns a tuple with the TotalDataTransferredNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDataTransferredNumeric

`func (o *DrivesResourceInner) SetTotalDataTransferredNumeric(v int32)`

SetTotalDataTransferredNumeric sets TotalDataTransferredNumeric field to given value.

### HasTotalDataTransferredNumeric

`func (o *DrivesResourceInner) HasTotalDataTransferredNumeric() bool`

HasTotalDataTransferredNumeric returns a boolean if a field has been set.

### GetTransferRate

`func (o *DrivesResourceInner) GetTransferRate() string`

GetTransferRate returns the TransferRate field if non-nil, zero value otherwise.

### GetTransferRateOk

`func (o *DrivesResourceInner) GetTransferRateOk() (*string, bool)`

GetTransferRateOk returns a tuple with the TransferRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferRate

`func (o *DrivesResourceInner) SetTransferRate(v string)`

SetTransferRate sets TransferRate field to given value.

### HasTransferRate

`func (o *DrivesResourceInner) HasTransferRate() bool`

HasTransferRate returns a boolean if a field has been set.

### GetTransferRateNumeric

`func (o *DrivesResourceInner) GetTransferRateNumeric() int32`

GetTransferRateNumeric returns the TransferRateNumeric field if non-nil, zero value otherwise.

### GetTransferRateNumericOk

`func (o *DrivesResourceInner) GetTransferRateNumericOk() (*int32, bool)`

GetTransferRateNumericOk returns a tuple with the TransferRateNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferRateNumeric

`func (o *DrivesResourceInner) SetTransferRateNumeric(v int32)`

SetTransferRateNumeric sets TransferRateNumeric field to given value.

### HasTransferRateNumeric

`func (o *DrivesResourceInner) HasTransferRateNumeric() bool`

HasTransferRateNumeric returns a boolean if a field has been set.

### GetType

`func (o *DrivesResourceInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DrivesResourceInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DrivesResourceInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DrivesResourceInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeNumeric

`func (o *DrivesResourceInner) GetTypeNumeric() int32`

GetTypeNumeric returns the TypeNumeric field if non-nil, zero value otherwise.

### GetTypeNumericOk

`func (o *DrivesResourceInner) GetTypeNumericOk() (*int32, bool)`

GetTypeNumericOk returns a tuple with the TypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeNumeric

`func (o *DrivesResourceInner) SetTypeNumeric(v int32)`

SetTypeNumeric sets TypeNumeric field to given value.

### HasTypeNumeric

`func (o *DrivesResourceInner) HasTypeNumeric() bool`

HasTypeNumeric returns a boolean if a field has been set.

### GetUrl

`func (o *DrivesResourceInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DrivesResourceInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DrivesResourceInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DrivesResourceInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsage

`func (o *DrivesResourceInner) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DrivesResourceInner) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DrivesResourceInner) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *DrivesResourceInner) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUsageNumeric

`func (o *DrivesResourceInner) GetUsageNumeric() int32`

GetUsageNumeric returns the UsageNumeric field if non-nil, zero value otherwise.

### GetUsageNumericOk

`func (o *DrivesResourceInner) GetUsageNumericOk() (*int32, bool)`

GetUsageNumericOk returns a tuple with the UsageNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageNumeric

`func (o *DrivesResourceInner) SetUsageNumeric(v int32)`

SetUsageNumeric sets UsageNumeric field to given value.

### HasUsageNumeric

`func (o *DrivesResourceInner) HasUsageNumeric() bool`

HasUsageNumeric returns a boolean if a field has been set.

### GetVendor

`func (o *DrivesResourceInner) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DrivesResourceInner) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DrivesResourceInner) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *DrivesResourceInner) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVirtualDiskSerial

`func (o *DrivesResourceInner) GetVirtualDiskSerial() string`

GetVirtualDiskSerial returns the VirtualDiskSerial field if non-nil, zero value otherwise.

### GetVirtualDiskSerialOk

`func (o *DrivesResourceInner) GetVirtualDiskSerialOk() (*string, bool)`

GetVirtualDiskSerialOk returns a tuple with the VirtualDiskSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskSerial

`func (o *DrivesResourceInner) SetVirtualDiskSerial(v string)`

SetVirtualDiskSerial sets VirtualDiskSerial field to given value.

### HasVirtualDiskSerial

`func (o *DrivesResourceInner) HasVirtualDiskSerial() bool`

HasVirtualDiskSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


