# PortResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**ActualSpeed** | Pointer to **string** | Actual speed for this port | [optional] 
**ActualSpeedNumeric** | Pointer to **int64** | Actual speed for this port( In numeric form ) | [optional] 
**ConfiguredSpeed** | Pointer to **string** |  | [optional] 
**ConfiguredSpeedNumeric** | Pointer to **int64** |  | [optional] 
**Controller** | Pointer to **string** |  | [optional] 
**ControllerNumeric** | Pointer to **int64** |  | [optional] 
**DurableId** | Pointer to **string** |  | [optional] 
**FanOut** | Pointer to **int64** |  | [optional] 
**Health** | Pointer to **string** |  | [optional] 
**HealthNumeric** | Pointer to **int64** |  | [optional] 
**HealthReason** | Pointer to **string** |  | [optional] 
**HealthRecommendation** | Pointer to **string** |  | [optional] 
**Media** | Pointer to **string** | Type of connection media used for this port | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**PortType** | Pointer to **string** | Port Type (FC, iSCSI, SAS) | [optional] 
**PortTypeNumeric** | Pointer to **int64** | Port Type (FC, iSCSI, SAS)( In numeric form ) | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusNumeric** | Pointer to **int64** |  | [optional] 
**TargetId** | Pointer to **string** | WWPN or IQN of the host port | [optional] 
**Url** | Pointer to **string** | The resource URL | [optional] 
**FcPort** | Pointer to [**[]FcPortResourceInner**](FcPortResourceInner.md) |  | [optional] 
**IscsiPort** | Pointer to [**[]IscsiPortResourceInner**](IscsiPortResourceInner.md) |  | [optional] 
**SasPort** | Pointer to [**[]SasPortResourceInner**](SasPortResourceInner.md) |  | [optional] 

## Methods

### NewPortResourceInner

`func NewPortResourceInner() *PortResourceInner`

NewPortResourceInner instantiates a new PortResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortResourceInnerWithDefaults

`func NewPortResourceInnerWithDefaults() *PortResourceInner`

NewPortResourceInnerWithDefaults instantiates a new PortResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *PortResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *PortResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *PortResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *PortResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *PortResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PortResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PortResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PortResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetActualSpeed

`func (o *PortResourceInner) GetActualSpeed() string`

GetActualSpeed returns the ActualSpeed field if non-nil, zero value otherwise.

### GetActualSpeedOk

`func (o *PortResourceInner) GetActualSpeedOk() (*string, bool)`

GetActualSpeedOk returns a tuple with the ActualSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSpeed

`func (o *PortResourceInner) SetActualSpeed(v string)`

SetActualSpeed sets ActualSpeed field to given value.

### HasActualSpeed

`func (o *PortResourceInner) HasActualSpeed() bool`

HasActualSpeed returns a boolean if a field has been set.

### GetActualSpeedNumeric

`func (o *PortResourceInner) GetActualSpeedNumeric() int64`

GetActualSpeedNumeric returns the ActualSpeedNumeric field if non-nil, zero value otherwise.

### GetActualSpeedNumericOk

`func (o *PortResourceInner) GetActualSpeedNumericOk() (*int64, bool)`

GetActualSpeedNumericOk returns a tuple with the ActualSpeedNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSpeedNumeric

`func (o *PortResourceInner) SetActualSpeedNumeric(v int64)`

SetActualSpeedNumeric sets ActualSpeedNumeric field to given value.

### HasActualSpeedNumeric

`func (o *PortResourceInner) HasActualSpeedNumeric() bool`

HasActualSpeedNumeric returns a boolean if a field has been set.

### GetConfiguredSpeed

`func (o *PortResourceInner) GetConfiguredSpeed() string`

GetConfiguredSpeed returns the ConfiguredSpeed field if non-nil, zero value otherwise.

### GetConfiguredSpeedOk

`func (o *PortResourceInner) GetConfiguredSpeedOk() (*string, bool)`

GetConfiguredSpeedOk returns a tuple with the ConfiguredSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredSpeed

`func (o *PortResourceInner) SetConfiguredSpeed(v string)`

SetConfiguredSpeed sets ConfiguredSpeed field to given value.

### HasConfiguredSpeed

`func (o *PortResourceInner) HasConfiguredSpeed() bool`

HasConfiguredSpeed returns a boolean if a field has been set.

### GetConfiguredSpeedNumeric

`func (o *PortResourceInner) GetConfiguredSpeedNumeric() int64`

GetConfiguredSpeedNumeric returns the ConfiguredSpeedNumeric field if non-nil, zero value otherwise.

### GetConfiguredSpeedNumericOk

`func (o *PortResourceInner) GetConfiguredSpeedNumericOk() (*int64, bool)`

GetConfiguredSpeedNumericOk returns a tuple with the ConfiguredSpeedNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredSpeedNumeric

`func (o *PortResourceInner) SetConfiguredSpeedNumeric(v int64)`

SetConfiguredSpeedNumeric sets ConfiguredSpeedNumeric field to given value.

### HasConfiguredSpeedNumeric

`func (o *PortResourceInner) HasConfiguredSpeedNumeric() bool`

HasConfiguredSpeedNumeric returns a boolean if a field has been set.

### GetController

`func (o *PortResourceInner) GetController() string`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *PortResourceInner) GetControllerOk() (*string, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *PortResourceInner) SetController(v string)`

SetController sets Controller field to given value.

### HasController

`func (o *PortResourceInner) HasController() bool`

HasController returns a boolean if a field has been set.

### GetControllerNumeric

`func (o *PortResourceInner) GetControllerNumeric() int64`

GetControllerNumeric returns the ControllerNumeric field if non-nil, zero value otherwise.

### GetControllerNumericOk

`func (o *PortResourceInner) GetControllerNumericOk() (*int64, bool)`

GetControllerNumericOk returns a tuple with the ControllerNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerNumeric

`func (o *PortResourceInner) SetControllerNumeric(v int64)`

SetControllerNumeric sets ControllerNumeric field to given value.

### HasControllerNumeric

`func (o *PortResourceInner) HasControllerNumeric() bool`

HasControllerNumeric returns a boolean if a field has been set.

### GetDurableId

`func (o *PortResourceInner) GetDurableId() string`

GetDurableId returns the DurableId field if non-nil, zero value otherwise.

### GetDurableIdOk

`func (o *PortResourceInner) GetDurableIdOk() (*string, bool)`

GetDurableIdOk returns a tuple with the DurableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableId

`func (o *PortResourceInner) SetDurableId(v string)`

SetDurableId sets DurableId field to given value.

### HasDurableId

`func (o *PortResourceInner) HasDurableId() bool`

HasDurableId returns a boolean if a field has been set.

### GetFanOut

`func (o *PortResourceInner) GetFanOut() int64`

GetFanOut returns the FanOut field if non-nil, zero value otherwise.

### GetFanOutOk

`func (o *PortResourceInner) GetFanOutOk() (*int64, bool)`

GetFanOutOk returns a tuple with the FanOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFanOut

`func (o *PortResourceInner) SetFanOut(v int64)`

SetFanOut sets FanOut field to given value.

### HasFanOut

`func (o *PortResourceInner) HasFanOut() bool`

HasFanOut returns a boolean if a field has been set.

### GetHealth

`func (o *PortResourceInner) GetHealth() string`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *PortResourceInner) GetHealthOk() (*string, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *PortResourceInner) SetHealth(v string)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *PortResourceInner) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHealthNumeric

`func (o *PortResourceInner) GetHealthNumeric() int64`

GetHealthNumeric returns the HealthNumeric field if non-nil, zero value otherwise.

### GetHealthNumericOk

`func (o *PortResourceInner) GetHealthNumericOk() (*int64, bool)`

GetHealthNumericOk returns a tuple with the HealthNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthNumeric

`func (o *PortResourceInner) SetHealthNumeric(v int64)`

SetHealthNumeric sets HealthNumeric field to given value.

### HasHealthNumeric

`func (o *PortResourceInner) HasHealthNumeric() bool`

HasHealthNumeric returns a boolean if a field has been set.

### GetHealthReason

`func (o *PortResourceInner) GetHealthReason() string`

GetHealthReason returns the HealthReason field if non-nil, zero value otherwise.

### GetHealthReasonOk

`func (o *PortResourceInner) GetHealthReasonOk() (*string, bool)`

GetHealthReasonOk returns a tuple with the HealthReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthReason

`func (o *PortResourceInner) SetHealthReason(v string)`

SetHealthReason sets HealthReason field to given value.

### HasHealthReason

`func (o *PortResourceInner) HasHealthReason() bool`

HasHealthReason returns a boolean if a field has been set.

### GetHealthRecommendation

`func (o *PortResourceInner) GetHealthRecommendation() string`

GetHealthRecommendation returns the HealthRecommendation field if non-nil, zero value otherwise.

### GetHealthRecommendationOk

`func (o *PortResourceInner) GetHealthRecommendationOk() (*string, bool)`

GetHealthRecommendationOk returns a tuple with the HealthRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthRecommendation

`func (o *PortResourceInner) SetHealthRecommendation(v string)`

SetHealthRecommendation sets HealthRecommendation field to given value.

### HasHealthRecommendation

`func (o *PortResourceInner) HasHealthRecommendation() bool`

HasHealthRecommendation returns a boolean if a field has been set.

### GetMedia

`func (o *PortResourceInner) GetMedia() string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *PortResourceInner) GetMediaOk() (*string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *PortResourceInner) SetMedia(v string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *PortResourceInner) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetPort

`func (o *PortResourceInner) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PortResourceInner) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PortResourceInner) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *PortResourceInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortType

`func (o *PortResourceInner) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *PortResourceInner) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *PortResourceInner) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *PortResourceInner) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetPortTypeNumeric

`func (o *PortResourceInner) GetPortTypeNumeric() int64`

GetPortTypeNumeric returns the PortTypeNumeric field if non-nil, zero value otherwise.

### GetPortTypeNumericOk

`func (o *PortResourceInner) GetPortTypeNumericOk() (*int64, bool)`

GetPortTypeNumericOk returns a tuple with the PortTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortTypeNumeric

`func (o *PortResourceInner) SetPortTypeNumeric(v int64)`

SetPortTypeNumeric sets PortTypeNumeric field to given value.

### HasPortTypeNumeric

`func (o *PortResourceInner) HasPortTypeNumeric() bool`

HasPortTypeNumeric returns a boolean if a field has been set.

### GetStatus

`func (o *PortResourceInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortResourceInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortResourceInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PortResourceInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusNumeric

`func (o *PortResourceInner) GetStatusNumeric() int64`

GetStatusNumeric returns the StatusNumeric field if non-nil, zero value otherwise.

### GetStatusNumericOk

`func (o *PortResourceInner) GetStatusNumericOk() (*int64, bool)`

GetStatusNumericOk returns a tuple with the StatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusNumeric

`func (o *PortResourceInner) SetStatusNumeric(v int64)`

SetStatusNumeric sets StatusNumeric field to given value.

### HasStatusNumeric

`func (o *PortResourceInner) HasStatusNumeric() bool`

HasStatusNumeric returns a boolean if a field has been set.

### GetTargetId

`func (o *PortResourceInner) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *PortResourceInner) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *PortResourceInner) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *PortResourceInner) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetUrl

`func (o *PortResourceInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PortResourceInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PortResourceInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PortResourceInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetFcPort

`func (o *PortResourceInner) GetFcPort() []FcPortResourceInner`

GetFcPort returns the FcPort field if non-nil, zero value otherwise.

### GetFcPortOk

`func (o *PortResourceInner) GetFcPortOk() (*[]FcPortResourceInner, bool)`

GetFcPortOk returns a tuple with the FcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcPort

`func (o *PortResourceInner) SetFcPort(v []FcPortResourceInner)`

SetFcPort sets FcPort field to given value.

### HasFcPort

`func (o *PortResourceInner) HasFcPort() bool`

HasFcPort returns a boolean if a field has been set.

### GetIscsiPort

`func (o *PortResourceInner) GetIscsiPort() []IscsiPortResourceInner`

GetIscsiPort returns the IscsiPort field if non-nil, zero value otherwise.

### GetIscsiPortOk

`func (o *PortResourceInner) GetIscsiPortOk() (*[]IscsiPortResourceInner, bool)`

GetIscsiPortOk returns a tuple with the IscsiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiPort

`func (o *PortResourceInner) SetIscsiPort(v []IscsiPortResourceInner)`

SetIscsiPort sets IscsiPort field to given value.

### HasIscsiPort

`func (o *PortResourceInner) HasIscsiPort() bool`

HasIscsiPort returns a boolean if a field has been set.

### GetSasPort

`func (o *PortResourceInner) GetSasPort() []SasPortResourceInner`

GetSasPort returns the SasPort field if non-nil, zero value otherwise.

### GetSasPortOk

`func (o *PortResourceInner) GetSasPortOk() (*[]SasPortResourceInner, bool)`

GetSasPortOk returns a tuple with the SasPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSasPort

`func (o *PortResourceInner) SetSasPort(v []SasPortResourceInner)`

SetSasPort sets SasPort field to given value.

### HasSasPort

`func (o *PortResourceInner) HasSasPort() bool`

HasSasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


