# CertificateStatusResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectName** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **string** |  | [optional] 
**CertificateSignature** | Pointer to **string** |  | [optional] 
**CertificateStatus** | Pointer to **string** |  | [optional] 
<<<<<<< HEAD
**CertificateStatusNumeric** | Pointer to **int64** |  | [optional] 
**CertificateText** | Pointer to **string** |  | [optional] 
**CertificateTime** | Pointer to **string** |  | [optional] 
**Controller** | Pointer to **string** |  | [optional] 
**ControllerNumeric** | Pointer to **int64** |  | [optional] 
=======
**CertificateStatusNumeric** | Pointer to **int32** |  | [optional] 
**CertificateText** | Pointer to **string** |  | [optional] 
**CertificateTime** | Pointer to **string** |  | [optional] 
**Controller** | Pointer to **string** |  | [optional] 
**ControllerNumeric** | Pointer to **int32** |  | [optional] 
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

## Methods

### NewCertificateStatusResourceInner

`func NewCertificateStatusResourceInner() *CertificateStatusResourceInner`

NewCertificateStatusResourceInner instantiates a new CertificateStatusResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStatusResourceInnerWithDefaults

`func NewCertificateStatusResourceInnerWithDefaults() *CertificateStatusResourceInner`

NewCertificateStatusResourceInnerWithDefaults instantiates a new CertificateStatusResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectName

`func (o *CertificateStatusResourceInner) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *CertificateStatusResourceInner) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *CertificateStatusResourceInner) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *CertificateStatusResourceInner) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetMeta

`func (o *CertificateStatusResourceInner) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CertificateStatusResourceInner) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CertificateStatusResourceInner) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CertificateStatusResourceInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetCertificateSignature

`func (o *CertificateStatusResourceInner) GetCertificateSignature() string`

GetCertificateSignature returns the CertificateSignature field if non-nil, zero value otherwise.

### GetCertificateSignatureOk

`func (o *CertificateStatusResourceInner) GetCertificateSignatureOk() (*string, bool)`

GetCertificateSignatureOk returns a tuple with the CertificateSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateSignature

`func (o *CertificateStatusResourceInner) SetCertificateSignature(v string)`

SetCertificateSignature sets CertificateSignature field to given value.

### HasCertificateSignature

`func (o *CertificateStatusResourceInner) HasCertificateSignature() bool`

HasCertificateSignature returns a boolean if a field has been set.

### GetCertificateStatus

`func (o *CertificateStatusResourceInner) GetCertificateStatus() string`

GetCertificateStatus returns the CertificateStatus field if non-nil, zero value otherwise.

### GetCertificateStatusOk

`func (o *CertificateStatusResourceInner) GetCertificateStatusOk() (*string, bool)`

GetCertificateStatusOk returns a tuple with the CertificateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStatus

`func (o *CertificateStatusResourceInner) SetCertificateStatus(v string)`

SetCertificateStatus sets CertificateStatus field to given value.

### HasCertificateStatus

`func (o *CertificateStatusResourceInner) HasCertificateStatus() bool`

HasCertificateStatus returns a boolean if a field has been set.

### GetCertificateStatusNumeric

<<<<<<< HEAD
`func (o *CertificateStatusResourceInner) GetCertificateStatusNumeric() int64`
=======
`func (o *CertificateStatusResourceInner) GetCertificateStatusNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCertificateStatusNumeric returns the CertificateStatusNumeric field if non-nil, zero value otherwise.

### GetCertificateStatusNumericOk

<<<<<<< HEAD
`func (o *CertificateStatusResourceInner) GetCertificateStatusNumericOk() (*int64, bool)`
=======
`func (o *CertificateStatusResourceInner) GetCertificateStatusNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetCertificateStatusNumericOk returns a tuple with the CertificateStatusNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStatusNumeric

<<<<<<< HEAD
`func (o *CertificateStatusResourceInner) SetCertificateStatusNumeric(v int64)`
=======
`func (o *CertificateStatusResourceInner) SetCertificateStatusNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetCertificateStatusNumeric sets CertificateStatusNumeric field to given value.

### HasCertificateStatusNumeric

`func (o *CertificateStatusResourceInner) HasCertificateStatusNumeric() bool`

HasCertificateStatusNumeric returns a boolean if a field has been set.

### GetCertificateText

`func (o *CertificateStatusResourceInner) GetCertificateText() string`

GetCertificateText returns the CertificateText field if non-nil, zero value otherwise.

### GetCertificateTextOk

`func (o *CertificateStatusResourceInner) GetCertificateTextOk() (*string, bool)`

GetCertificateTextOk returns a tuple with the CertificateText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateText

`func (o *CertificateStatusResourceInner) SetCertificateText(v string)`

SetCertificateText sets CertificateText field to given value.

### HasCertificateText

`func (o *CertificateStatusResourceInner) HasCertificateText() bool`

HasCertificateText returns a boolean if a field has been set.

### GetCertificateTime

`func (o *CertificateStatusResourceInner) GetCertificateTime() string`

GetCertificateTime returns the CertificateTime field if non-nil, zero value otherwise.

### GetCertificateTimeOk

`func (o *CertificateStatusResourceInner) GetCertificateTimeOk() (*string, bool)`

GetCertificateTimeOk returns a tuple with the CertificateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTime

`func (o *CertificateStatusResourceInner) SetCertificateTime(v string)`

SetCertificateTime sets CertificateTime field to given value.

### HasCertificateTime

`func (o *CertificateStatusResourceInner) HasCertificateTime() bool`

HasCertificateTime returns a boolean if a field has been set.

### GetController

`func (o *CertificateStatusResourceInner) GetController() string`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *CertificateStatusResourceInner) GetControllerOk() (*string, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *CertificateStatusResourceInner) SetController(v string)`

SetController sets Controller field to given value.

### HasController

`func (o *CertificateStatusResourceInner) HasController() bool`

HasController returns a boolean if a field has been set.

### GetControllerNumeric

<<<<<<< HEAD
`func (o *CertificateStatusResourceInner) GetControllerNumeric() int64`
=======
`func (o *CertificateStatusResourceInner) GetControllerNumeric() int32`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetControllerNumeric returns the ControllerNumeric field if non-nil, zero value otherwise.

### GetControllerNumericOk

<<<<<<< HEAD
`func (o *CertificateStatusResourceInner) GetControllerNumericOk() (*int64, bool)`
=======
`func (o *CertificateStatusResourceInner) GetControllerNumericOk() (*int32, bool)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

GetControllerNumericOk returns a tuple with the ControllerNumeric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerNumeric

<<<<<<< HEAD
`func (o *CertificateStatusResourceInner) SetControllerNumeric(v int64)`
=======
`func (o *CertificateStatusResourceInner) SetControllerNumeric(v int32)`
>>>>>>> aac8175 (feat(apiv2): openapi generator and validator and spec)

SetControllerNumeric sets ControllerNumeric field to given value.

### HasControllerNumeric

`func (o *CertificateStatusResourceInner) HasControllerNumeric() bool`

HasControllerNumeric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


