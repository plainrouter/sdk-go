# \CreativeAPI

All URIs are relative to *https://plainrouter.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LaunchCreativesIndex**](CreativeAPI.md#LaunchCreativesIndex) | **Get** /workspaces/{workspace}/admin/creatives | List workspace creatives
[**LaunchCreativesShow**](CreativeAPI.md#LaunchCreativesShow) | **Get** /workspaces/{workspace}/admin/creatives/{creative} | Read a workspace creative
[**LaunchCreativesStatus**](CreativeAPI.md#LaunchCreativesStatus) | **Patch** /workspaces/{workspace}/admin/creatives/{creative}/status | Update a workspace creative status
[**StoreCreative**](CreativeAPI.md#StoreCreative) | **Post** /workspaces/{workspace}/admin/creatives | Upload a creative
[**StoreCreativeFromUrl**](CreativeAPI.md#StoreCreativeFromUrl) | **Post** /workspaces/{workspace}/admin/creatives/from-url | Store a creative from a URL



## LaunchCreativesIndex

> LaunchCreativesIndex200Response LaunchCreativesIndex(ctx, workspace).Cursor(cursor).PerPage(perPage).Page(page).Execute()

List workspace creatives



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/plainrouter/sdk-go/openapi"
)

func main() {
	workspace := int32(56) // int32 | The workspace ID
	cursor := int32(56) // int32 | Continue using next_cursor with the same filters; rule lists use an immutable rule ID, other lists an offset or object ID. (optional)
	perPage := int32(56) // int32 | Maximum complete rows; the byte budget may return fewer. Defaults to 25. (optional) (default to 25)
	page := int64(789) // int64 | Integer page from 1 through 10000000000000000; integer strings are accepted. Zero, negatives, fractions, non-numeric strings and booleans return 422 with a page field validation error. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeAPI.LaunchCreativesIndex(context.Background(), workspace).Cursor(cursor).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeAPI.LaunchCreativesIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchCreativesIndex`: LaunchCreativesIndex200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeAPI.LaunchCreativesIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchCreativesIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **int32** | Continue using next_cursor with the same filters; rule lists use an immutable rule ID, other lists an offset or object ID. | 
 **perPage** | **int32** | Maximum complete rows; the byte budget may return fewer. Defaults to 25. | [default to 25]
 **page** | **int64** | Integer page from 1 through 10000000000000000; integer strings are accepted. Zero, negatives, fractions, non-numeric strings and booleans return 422 with a page field validation error. | 

### Return type

[**LaunchCreativesIndex200Response**](LaunchCreativesIndex200Response.md)

### Authorization

[workspaceSecret](../README.md#workspaceSecret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchCreativesShow

> LaunchCreativesStatus200Response LaunchCreativesShow(ctx, workspace, creative).Execute()

Read a workspace creative



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/plainrouter/sdk-go/openapi"
)

func main() {
	workspace := int32(56) // int32 | The workspace ID
	creative := "creative_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeAPI.LaunchCreativesShow(context.Background(), workspace, creative).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeAPI.LaunchCreativesShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchCreativesShow`: LaunchCreativesStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeAPI.LaunchCreativesShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**creative** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchCreativesShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LaunchCreativesStatus200Response**](LaunchCreativesStatus200Response.md)

### Authorization

[workspaceSecret](../README.md#workspaceSecret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchCreativesStatus

> LaunchCreativesStatus200Response LaunchCreativesStatus(ctx, workspace, creative).UpdateCreativeStatusRequest(updateCreativeStatusRequest).Execute()

Update a workspace creative status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/plainrouter/sdk-go/openapi"
)

func main() {
	workspace := int32(56) // int32 | The workspace ID
	creative := "creative_example" // string | 
	updateCreativeStatusRequest := *openapiclient.NewUpdateCreativeStatusRequest("Status_example") // UpdateCreativeStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeAPI.LaunchCreativesStatus(context.Background(), workspace, creative).UpdateCreativeStatusRequest(updateCreativeStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeAPI.LaunchCreativesStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchCreativesStatus`: LaunchCreativesStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `CreativeAPI.LaunchCreativesStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**creative** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchCreativesStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCreativeStatusRequest** | [**UpdateCreativeStatusRequest**](UpdateCreativeStatusRequest.md) |  | 

### Return type

[**LaunchCreativesStatus200Response**](LaunchCreativesStatus200Response.md)

### Authorization

[workspaceSecret](../README.md#workspaceSecret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreCreative

> CreativeIntakeRead StoreCreative(ctx, workspace).File(file).Tags(tags).SourceRef(sourceRef).Execute()

Upload a creative



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/plainrouter/sdk-go/openapi"
)

func main() {
	workspace := int32(56) // int32 | The workspace ID
	file := os.NewFile(1234, "some_file") // *os.File | JPEG, PNG, MP4 or QuickTime bytes, at most 4 GiB.
	tags := *openapiclient.NewStoreCreativeRequestTags() // StoreCreativeRequestTags |  (optional)
	sourceRef := "sourceRef_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeAPI.StoreCreative(context.Background(), workspace).File(file).Tags(tags).SourceRef(sourceRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeAPI.StoreCreative``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreCreative`: CreativeIntakeRead
	fmt.Fprintf(os.Stdout, "Response from `CreativeAPI.StoreCreative`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreCreativeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | JPEG, PNG, MP4 or QuickTime bytes, at most 4 GiB. | 
 **tags** | [**StoreCreativeRequestTags**](StoreCreativeRequestTags.md) |  | 
 **sourceRef** | **string** |  | 

### Return type

[**CreativeIntakeRead**](CreativeIntakeRead.md)

### Authorization

[creativeWriter](../README.md#creativeWriter)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreCreativeFromUrl

> CreativeIntakeRead StoreCreativeFromUrl(ctx, workspace).StoreCreativeFromUrlRequest(storeCreativeFromUrlRequest).Execute()

Store a creative from a URL



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/plainrouter/sdk-go/openapi"
)

func main() {
	workspace := int32(56) // int32 | The workspace ID
	storeCreativeFromUrlRequest := *openapiclient.NewStoreCreativeFromUrlRequest("Url_example") // StoreCreativeFromUrlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreativeAPI.StoreCreativeFromUrl(context.Background(), workspace).StoreCreativeFromUrlRequest(storeCreativeFromUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreativeAPI.StoreCreativeFromUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreCreativeFromUrl`: CreativeIntakeRead
	fmt.Fprintf(os.Stdout, "Response from `CreativeAPI.StoreCreativeFromUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreCreativeFromUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storeCreativeFromUrlRequest** | [**StoreCreativeFromUrlRequest**](StoreCreativeFromUrlRequest.md) |  | 

### Return type

[**CreativeIntakeRead**](CreativeIntakeRead.md)

### Authorization

[creativeWriter](../README.md#creativeWriter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

