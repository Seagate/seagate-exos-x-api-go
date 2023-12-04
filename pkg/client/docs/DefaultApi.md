# \DefaultApi

All URIs are relative to *http://localhost:8080/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyVolumeDestinationPoolNameSourceGet**](DefaultApi.md#CopyVolumeDestinationPoolNameSourceGet) | **Get** /copy/volume/destination-pool/{destination-poolOption}/name/{nameOption}/{sourceOption} | 
[**CreateSnapshotsVolumesNamesGet**](DefaultApi.md#CreateSnapshotsVolumesNamesGet) | **Get** /create/snapshots/volumes/{volumesOption}/{namesOption} | 
[**CreateVolumePoolSizeTierAffinityNameGet**](DefaultApi.md#CreateVolumePoolSizeTierAffinityNameGet) | **Get** /create/volume/pool/{poolOption}/size/{sizeOption}/tier-affinity/{tier-affinityOption}/{nameOption} | 
[**DeleteHostsNamesGet**](DefaultApi.md#DeleteHostsNamesGet) | **Get** /delete/hosts/{namesOption} | 
[**DeleteInitiatorNicknameNameGet**](DefaultApi.md#DeleteInitiatorNicknameNameGet) | **Get** /delete/initiator-nickname/{nameOption} | 
[**DeleteSnapshotNamesGet**](DefaultApi.md#DeleteSnapshotNamesGet) | **Get** /delete/snapshot/{namesOption} | 
[**DeleteVolumesNamesGet**](DefaultApi.md#DeleteVolumesNamesGet) | **Get** /delete/volumes/{namesOption} | 
[**ExpandVolumeSizeNameGet**](DefaultApi.md#ExpandVolumeSizeNameGet) | **Get** /expand/volume/size/{sizeOption}/{nameOption} | 
[**LoginGet**](DefaultApi.md#LoginGet) | **Get** /login | 
[**LogoutGet**](DefaultApi.md#LogoutGet) | **Get** /logout | 
[**MapVolumeAccessLunInitiatorNamesGet**](DefaultApi.md#MapVolumeAccessLunInitiatorNamesGet) | **Get** /map/volume/access/{accessOption}/lun/{lunOption}/initiator/{initiatorOption}/{namesOption} | 
[**SchemaGet**](DefaultApi.md#SchemaGet) | **Get** /meta/{schemaId} | 
[**SetInitiatorIdNicknameGet**](DefaultApi.md#SetInitiatorIdNicknameGet) | **Get** /set/initiator/id/{idOption}/nickname/{nicknameOption} | 
[**ShowAdvancedSettingsGet**](DefaultApi.md#ShowAdvancedSettingsGet) | **Get** /show/advanced-settings | 
[**ShowCacheParametersGet**](DefaultApi.md#ShowCacheParametersGet) | **Get** /show/cache-parameters | 
[**ShowCertificateGet**](DefaultApi.md#ShowCertificateGet) | **Get** /show/certificate | 
[**ShowControllersGet**](DefaultApi.md#ShowControllersGet) | **Get** /show/controllers | 
[**ShowDiskGroupsGet**](DefaultApi.md#ShowDiskGroupsGet) | **Get** /show/disk-groups | 
[**ShowDisksGet**](DefaultApi.md#ShowDisksGet) | **Get** /show/disks | 
[**ShowEnclosuresGet**](DefaultApi.md#ShowEnclosuresGet) | **Get** /show/enclosures | 
[**ShowFansGet**](DefaultApi.md#ShowFansGet) | **Get** /show/fans | 
[**ShowHostGroupsGet**](DefaultApi.md#ShowHostGroupsGet) | **Get** /show/host-groups | 
[**ShowInitiatorNamesGet**](DefaultApi.md#ShowInitiatorNamesGet) | **Get** /show/initiator/{namesOption} | 
[**ShowInitiatorsGet**](DefaultApi.md#ShowInitiatorsGet) | **Get** /show/initiators | 
[**ShowMapsAllGet**](DefaultApi.md#ShowMapsAllGet) | **Get** /show/maps/all | 
[**ShowMapsGet**](DefaultApi.md#ShowMapsGet) | **Get** /show/maps | 
[**ShowMapsInitiatorNamesGet**](DefaultApi.md#ShowMapsInitiatorNamesGet) | **Get** /show/maps/initiator/{namesOption} | 
[**ShowMapsNamesGet**](DefaultApi.md#ShowMapsNamesGet) | **Get** /show/maps/{namesOption} | 
[**ShowPoolsGet**](DefaultApi.md#ShowPoolsGet) | **Get** /show/pools | 
[**ShowPowerSuppliesGet**](DefaultApi.md#ShowPowerSuppliesGet) | **Get** /show/power-supplies | 
[**ShowSnapshotsGet**](DefaultApi.md#ShowSnapshotsGet) | **Get** /show/snapshots | 
[**ShowSnapshotsPatternGet**](DefaultApi.md#ShowSnapshotsPatternGet) | **Get** /show/snapshots/pattern/{patternOption} | 
[**ShowSnapshotsVolumeGet**](DefaultApi.md#ShowSnapshotsVolumeGet) | **Get** /show/snapshots/volume/{volumeOption} | 
[**ShowSystemGet**](DefaultApi.md#ShowSystemGet) | **Get** /show/system | 
[**ShowVersionsDetailGet**](DefaultApi.md#ShowVersionsDetailGet) | **Get** /show/versions/detail | 
[**ShowVersionsGet**](DefaultApi.md#ShowVersionsGet) | **Get** /show/versions | 
[**ShowVolumesNamesGet**](DefaultApi.md#ShowVolumesNamesGet) | **Get** /show/volumes/{namesOption} | 
[**UnmapVolumeInitiatorNamesGet**](DefaultApi.md#UnmapVolumeInitiatorNamesGet) | **Get** /unmap/volume/initiator/{initiatorOption}/{namesOption} | 
[**UnmapVolumeNamesGet**](DefaultApi.md#UnmapVolumeNamesGet) | **Get** /unmap/volume/{namesOption} | 



## CopyVolumeDestinationPoolNameSourceGet

> StatusObject CopyVolumeDestinationPoolNameSourceGet(ctx, destinationPoolOption, nameOption, sourceOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    destinationPoolOption := "destinationPoolOption_example" // string | 
    nameOption := "nameOption_example" // string | 
    sourceOption := "sourceOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CopyVolumeDestinationPoolNameSourceGet(context.Background(), destinationPoolOption, nameOption, sourceOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CopyVolumeDestinationPoolNameSourceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyVolumeDestinationPoolNameSourceGet`: StatusObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CopyVolumeDestinationPoolNameSourceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationPoolOption** | **string** |  | 
**nameOption** | **string** |  | 
**sourceOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyVolumeDestinationPoolNameSourceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**StatusObject**](StatusObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSnapshotsVolumesNamesGet

> StatusObject CreateSnapshotsVolumesNamesGet(ctx, volumesOption, namesOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    volumesOption := "volumesOption_example" // string | 
    namesOption := "namesOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSnapshotsVolumesNamesGet(context.Background(), volumesOption, namesOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSnapshotsVolumesNamesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSnapshotsVolumesNamesGet`: StatusObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSnapshotsVolumesNamesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumesOption** | **string** |  | 
**namesOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotsVolumesNamesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StatusObject**](StatusObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVolumePoolSizeTierAffinityNameGet

> VolumesObject CreateVolumePoolSizeTierAffinityNameGet(ctx, poolOption, sizeOption, tierAffinityOption, nameOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    poolOption := "poolOption_example" // string | 
    sizeOption := "sizeOption_example" // string | 
    tierAffinityOption := "tierAffinityOption_example" // string | 
    nameOption := "nameOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateVolumePoolSizeTierAffinityNameGet(context.Background(), poolOption, sizeOption, tierAffinityOption, nameOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateVolumePoolSizeTierAffinityNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVolumePoolSizeTierAffinityNameGet`: VolumesObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateVolumePoolSizeTierAffinityNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolOption** | **string** |  | 
**sizeOption** | **string** |  | 
**tierAffinityOption** | **string** |  | 
**nameOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumePoolSizeTierAffinityNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**VolumesObject**](VolumesObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHostsNamesGet

> StatusObject DeleteHostsNamesGet(ctx, namesOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    namesOption := "namesOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteHostsNamesGet(context.Background(), namesOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteHostsNamesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHostsNamesGet`: StatusObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteHostsNamesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namesOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHostsNamesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusObject**](StatusObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInitiatorNicknameNameGet

> StatusObject DeleteInitiatorNicknameNameGet(ctx, nameOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    nameOption := "nameOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteInitiatorNicknameNameGet(context.Background(), nameOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteInitiatorNicknameNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInitiatorNicknameNameGet`: StatusObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteInitiatorNicknameNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInitiatorNicknameNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusObject**](StatusObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSnapshotNamesGet

> StatusObject DeleteSnapshotNamesGet(ctx, namesOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    namesOption := "namesOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSnapshotNamesGet(context.Background(), namesOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSnapshotNamesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSnapshotNamesGet`: StatusObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteSnapshotNamesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namesOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotNamesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusObject**](StatusObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolumesNamesGet

> StatusObject DeleteVolumesNamesGet(ctx, namesOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    namesOption := "namesOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteVolumesNamesGet(context.Background(), namesOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteVolumesNamesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVolumesNamesGet`: StatusObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteVolumesNamesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namesOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumesNamesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusObject**](StatusObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpandVolumeSizeNameGet

> StatusObject ExpandVolumeSizeNameGet(ctx, sizeOption, nameOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    sizeOption := "sizeOption_example" // string | 
    nameOption := "nameOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ExpandVolumeSizeNameGet(context.Background(), sizeOption, nameOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ExpandVolumeSizeNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpandVolumeSizeNameGet`: StatusObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ExpandVolumeSizeNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sizeOption** | **string** |  | 
**nameOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpandVolumeSizeNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StatusObject**](StatusObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginGet

> StatusObject LoginGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.LoginGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.LoginGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoginGet`: StatusObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.LoginGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLoginGetRequest struct via the builder pattern


### Return type

[**StatusObject**](StatusObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogoutGet

> StatusObject LogoutGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.LogoutGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.LogoutGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogoutGet`: StatusObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.LogoutGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutGetRequest struct via the builder pattern


### Return type

[**StatusObject**](StatusObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MapVolumeAccessLunInitiatorNamesGet

> StatusObject MapVolumeAccessLunInitiatorNamesGet(ctx, accessOption, lunOption, initiatorOption, namesOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    accessOption := "accessOption_example" // string | 
    lunOption := "lunOption_example" // string | 
    initiatorOption := "initiatorOption_example" // string | 
    namesOption := "namesOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MapVolumeAccessLunInitiatorNamesGet(context.Background(), accessOption, lunOption, initiatorOption, namesOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MapVolumeAccessLunInitiatorNamesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MapVolumeAccessLunInitiatorNamesGet`: StatusObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MapVolumeAccessLunInitiatorNamesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessOption** | **string** |  | 
**lunOption** | **string** |  | 
**initiatorOption** | **string** |  | 
**namesOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMapVolumeAccessLunInitiatorNamesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**StatusObject**](StatusObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchemaGet

> interface{} SchemaGet(ctx, schemaId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    schemaId := "volumes" // string | A string label for a resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SchemaGet(context.Background(), schemaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SchemaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SchemaGet`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SchemaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** | A string label for a resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetInitiatorIdNicknameGet

> StatusObject SetInitiatorIdNicknameGet(ctx, idOption, nicknameOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    idOption := "idOption_example" // string | 
    nicknameOption := "nicknameOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SetInitiatorIdNicknameGet(context.Background(), idOption, nicknameOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetInitiatorIdNicknameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetInitiatorIdNicknameGet`: StatusObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SetInitiatorIdNicknameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOption** | **string** |  | 
**nicknameOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetInitiatorIdNicknameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StatusObject**](StatusObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowAdvancedSettingsGet

> AdvancedSettingsTableObject ShowAdvancedSettingsGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowAdvancedSettingsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowAdvancedSettingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowAdvancedSettingsGet`: AdvancedSettingsTableObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowAdvancedSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowAdvancedSettingsGetRequest struct via the builder pattern


### Return type

[**AdvancedSettingsTableObject**](AdvancedSettingsTableObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowCacheParametersGet

> CacheSettingsObject ShowCacheParametersGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowCacheParametersGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowCacheParametersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowCacheParametersGet`: CacheSettingsObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowCacheParametersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowCacheParametersGetRequest struct via the builder pattern


### Return type

[**CacheSettingsObject**](CacheSettingsObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowCertificateGet

> CertificateStatusObject ShowCertificateGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowCertificateGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowCertificateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowCertificateGet`: CertificateStatusObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowCertificateGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowCertificateGetRequest struct via the builder pattern


### Return type

[**CertificateStatusObject**](CertificateStatusObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowControllersGet

> ControllersObject ShowControllersGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowControllersGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowControllersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowControllersGet`: ControllersObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowControllersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowControllersGetRequest struct via the builder pattern


### Return type

[**ControllersObject**](ControllersObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowDiskGroupsGet

> DiskGroupsObject ShowDiskGroupsGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowDiskGroupsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowDiskGroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowDiskGroupsGet`: DiskGroupsObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowDiskGroupsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowDiskGroupsGetRequest struct via the builder pattern


### Return type

[**DiskGroupsObject**](DiskGroupsObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowDisksGet

> DrivesObject ShowDisksGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowDisksGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowDisksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowDisksGet`: DrivesObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowDisksGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowDisksGetRequest struct via the builder pattern


### Return type

[**DrivesObject**](DrivesObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowEnclosuresGet

> EnclosuresObject ShowEnclosuresGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowEnclosuresGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowEnclosuresGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowEnclosuresGet`: EnclosuresObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowEnclosuresGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowEnclosuresGetRequest struct via the builder pattern


### Return type

[**EnclosuresObject**](EnclosuresObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowFansGet

> FanObject ShowFansGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowFansGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowFansGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowFansGet`: FanObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowFansGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowFansGetRequest struct via the builder pattern


### Return type

[**FanObject**](FanObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowHostGroupsGet

> HostGroupObject ShowHostGroupsGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowHostGroupsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowHostGroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowHostGroupsGet`: HostGroupObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowHostGroupsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowHostGroupsGetRequest struct via the builder pattern


### Return type

[**HostGroupObject**](HostGroupObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowInitiatorNamesGet

> InitiatorObject ShowInitiatorNamesGet(ctx, namesOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    namesOption := "namesOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowInitiatorNamesGet(context.Background(), namesOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowInitiatorNamesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowInitiatorNamesGet`: InitiatorObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowInitiatorNamesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namesOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowInitiatorNamesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InitiatorObject**](InitiatorObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowInitiatorsGet

> InitiatorObject ShowInitiatorsGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowInitiatorsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowInitiatorsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowInitiatorsGet`: InitiatorObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowInitiatorsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowInitiatorsGetRequest struct via the builder pattern


### Return type

[**InitiatorObject**](InitiatorObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowMapsAllGet

> VolumeViewObject ShowMapsAllGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowMapsAllGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowMapsAllGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowMapsAllGet`: VolumeViewObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowMapsAllGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowMapsAllGetRequest struct via the builder pattern


### Return type

[**VolumeViewObject**](VolumeViewObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowMapsGet

> VolumeViewObject ShowMapsGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowMapsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowMapsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowMapsGet`: VolumeViewObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowMapsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowMapsGetRequest struct via the builder pattern


### Return type

[**VolumeViewObject**](VolumeViewObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowMapsInitiatorNamesGet

> InitiatorViewObject ShowMapsInitiatorNamesGet(ctx, namesOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    namesOption := "namesOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowMapsInitiatorNamesGet(context.Background(), namesOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowMapsInitiatorNamesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowMapsInitiatorNamesGet`: InitiatorViewObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowMapsInitiatorNamesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namesOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowMapsInitiatorNamesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InitiatorViewObject**](InitiatorViewObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowMapsNamesGet

> VolumeViewObject ShowMapsNamesGet(ctx, namesOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    namesOption := "namesOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowMapsNamesGet(context.Background(), namesOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowMapsNamesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowMapsNamesGet`: VolumeViewObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowMapsNamesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namesOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowMapsNamesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeViewObject**](VolumeViewObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPoolsGet

> PoolsObject ShowPoolsGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowPoolsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowPoolsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowPoolsGet`: PoolsObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowPoolsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowPoolsGetRequest struct via the builder pattern


### Return type

[**PoolsObject**](PoolsObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPowerSuppliesGet

> PowerSuppliesObject ShowPowerSuppliesGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowPowerSuppliesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowPowerSuppliesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowPowerSuppliesGet`: PowerSuppliesObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowPowerSuppliesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowPowerSuppliesGetRequest struct via the builder pattern


### Return type

[**PowerSuppliesObject**](PowerSuppliesObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowSnapshotsGet

> SnapshotsObject ShowSnapshotsGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowSnapshotsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowSnapshotsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowSnapshotsGet`: SnapshotsObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowSnapshotsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowSnapshotsGetRequest struct via the builder pattern


### Return type

[**SnapshotsObject**](SnapshotsObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowSnapshotsPatternGet

> SnapshotsObject ShowSnapshotsPatternGet(ctx, patternOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    patternOption := "patternOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowSnapshotsPatternGet(context.Background(), patternOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowSnapshotsPatternGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowSnapshotsPatternGet`: SnapshotsObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowSnapshotsPatternGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patternOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowSnapshotsPatternGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SnapshotsObject**](SnapshotsObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowSnapshotsVolumeGet

> SnapshotsObject ShowSnapshotsVolumeGet(ctx, volumeOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    volumeOption := "volumeOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowSnapshotsVolumeGet(context.Background(), volumeOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowSnapshotsVolumeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowSnapshotsVolumeGet`: SnapshotsObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowSnapshotsVolumeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowSnapshotsVolumeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SnapshotsObject**](SnapshotsObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowSystemGet

> SystemObject ShowSystemGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowSystemGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowSystemGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowSystemGet`: SystemObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowSystemGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowSystemGetRequest struct via the builder pattern


### Return type

[**SystemObject**](SystemObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVersionsDetailGet

> VersionsObject ShowVersionsDetailGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowVersionsDetailGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowVersionsDetailGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowVersionsDetailGet`: VersionsObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowVersionsDetailGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowVersionsDetailGetRequest struct via the builder pattern


### Return type

[**VersionsObject**](VersionsObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVersionsGet

> VersionsObject ShowVersionsGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowVersionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowVersionsGet`: VersionsObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowVersionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowVersionsGetRequest struct via the builder pattern


### Return type

[**VersionsObject**](VersionsObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVolumesNamesGet

> VolumesObject ShowVolumesNamesGet(ctx, namesOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    namesOption := "namesOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ShowVolumesNamesGet(context.Background(), namesOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ShowVolumesNamesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowVolumesNamesGet`: VolumesObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ShowVolumesNamesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namesOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowVolumesNamesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumesObject**](VolumesObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnmapVolumeInitiatorNamesGet

> StatusObject UnmapVolumeInitiatorNamesGet(ctx, initiatorOption, namesOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    initiatorOption := "initiatorOption_example" // string | 
    namesOption := "namesOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UnmapVolumeInitiatorNamesGet(context.Background(), initiatorOption, namesOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UnmapVolumeInitiatorNamesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnmapVolumeInitiatorNamesGet`: StatusObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UnmapVolumeInitiatorNamesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**initiatorOption** | **string** |  | 
**namesOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmapVolumeInitiatorNamesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StatusObject**](StatusObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnmapVolumeNamesGet

> StatusObject UnmapVolumeNamesGet(ctx, namesOption).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    namesOption := "namesOption_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UnmapVolumeNamesGet(context.Background(), namesOption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UnmapVolumeNamesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnmapVolumeNamesGet`: StatusObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UnmapVolumeNamesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namesOption** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmapVolumeNamesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusObject**](StatusObject.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

