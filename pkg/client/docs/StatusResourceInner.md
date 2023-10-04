# StatusResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**ComponentId** | Pointer to **string** |  | [optional] 
**Response** | Pointer to **string** |  | [optional] 
**ResponseType** | Pointer to **string** | Indicates whether the command returned a Success, Error, Info, or Warning message | [optional] 
**ResponseTypeNumeric** | Pointer to **int32** | Indicates whether the command returned a Success (0), Error (1), Info (2), or Warning (3) message (numeric form) | [optional] 
**ReturnCode** | Pointer to **int32** | Numeric return code for the command | [optional] 
**TimeStamp** | Pointer to **string** | Time at which this event was detected | [optional] 
**TimeStampNumeric** | Pointer to **int32** | Time at which this event was detected( In numeric form ) | [optional] 

## Methods

### NewStatusResourceInner

`func NewStatusResourceInner() *StatusResourceInner`

NewStatusResourceInner instantiates a new StatusResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusResourceInnerWithDefaults

`func NewStatusResourceInnerWithDefaults() *StatusResourceInner`

NewStatusResourceInnerWithDefaults instantiates a new StatusResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *StatusResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *StatusResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *StatusResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *StatusResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *StatusResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *StatusResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *StatusResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *StatusResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetComponentId

`func (o *StatusResourceInner) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *StatusResourceInner) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *StatusResourceInner) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *StatusResourceInner) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetResponse

`func (o *StatusResourceInner) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *StatusResourceInner) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *StatusResourceInner) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *StatusResourceInner) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetResponseType

`func (o *StatusResourceInner) GetResponseType() string`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *StatusResourceInner) GetResponseTypeOk() (*string, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *StatusResourceInner) SetResponseType(v string)`

SetResponseType sets ResponseType field to given value.

### HasResponseType

`func (o *StatusResourceInner) HasResponseType() bool`

HasResponseType returns a boolean if a field has been set.

### GetResponseTypeNumeric

`func (o *StatusResourceInner) GetResponseTypeNumeric() int32`

GetResponseTypeNumeric returns the ResponseTypeNumeric field if non-nil, zero value otherwise.

### GetResponseTypeNumericOk

`func (o *StatusResourceInner) GetResponseTypeNumericOk() (*int32, bool)`

GetResponseTypeNumericOk returns a tuple with the ResponseTypeNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypeNumeric

`func (o *StatusResourceInner) SetResponseTypeNumeric(v int32)`

SetResponseTypeNumeric sets ResponseTypeNumeric field to given value.

### HasResponseTypeNumeric

`func (o *StatusResourceInner) HasResponseTypeNumeric() bool`

HasResponseTypeNumeric returns a boolean if a field has been set.

### GetReturnCode

`func (o *StatusResourceInner) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *StatusResourceInner) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *StatusResourceInner) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *StatusResourceInner) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetTimeStamp

`func (o *StatusResourceInner) GetTimeStamp() string`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *StatusResourceInner) GetTimeStampOk() (*string, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *StatusResourceInner) SetTimeStamp(v string)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *StatusResourceInner) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetTimeStampNumeric

`func (o *StatusResourceInner) GetTimeStampNumeric() int32`

GetTimeStampNumeric returns the TimeStampNumeric field if non-nil, zero value otherwise.

### GetTimeStampNumericOk

`func (o *StatusResourceInner) GetTimeStampNumericOk() (*int32, bool)`

GetTimeStampNumericOk returns a tuple with the TimeStampNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStampNumeric

`func (o *StatusResourceInner) SetTimeStampNumeric(v int32)`

SetTimeStampNumeric sets TimeStampNumeric field to given value.

### HasTimeStampNumeric

`func (o *StatusResourceInner) HasTimeStampNumeric() bool`

HasTimeStampNumeric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


