# CertificateStatusObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**[]StatusResourceInner**](StatusResourceInner.md) |  | [optional] 
**CertificateStatus** | Pointer to [**[]CertificateStatusResourceInner**](CertificateStatusResourceInner.md) |  | [optional] 

## Methods

### NewCertificateStatusObject

`func NewCertificateStatusObject() *CertificateStatusObject`

NewCertificateStatusObject instantiates a new CertificateStatusObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStatusObjectWithDefaults

`func NewCertificateStatusObjectWithDefaults() *CertificateStatusObject`

NewCertificateStatusObjectWithDefaults instantiates a new CertificateStatusObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CertificateStatusObject) GetStatus() []StatusResourceInner`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateStatusObject) GetStatusOk() (*[]StatusResourceInner, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateStatusObject) SetStatus(v []StatusResourceInner)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CertificateStatusObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCertificateStatus

`func (o *CertificateStatusObject) GetCertificateStatus() []CertificateStatusResourceInner`

GetCertificateStatus returns the CertificateStatus field if non-nil, zero value otherwise.

### GetCertificateStatusOk

`func (o *CertificateStatusObject) GetCertificateStatusOk() (*[]CertificateStatusResourceInner, bool)`

GetCertificateStatusOk returns a tuple with the CertificateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStatus

`func (o *CertificateStatusObject) SetCertificateStatus(v []CertificateStatusResourceInner)`

SetCertificateStatus sets CertificateStatus field to given value.

### HasCertificateStatus

`func (o *CertificateStatusObject) HasCertificateStatus() bool`

HasCertificateStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


