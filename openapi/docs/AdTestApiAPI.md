# \AdTestApiAPI

All URIs are relative to *https://plainrouter.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionsTestsApiIndex**](AdTestApiAPI.md#ActionsTestsApiIndex) | **Get** /agent/workspaces/{workspace}/tests | List Tests and observed conversion event choices
[**ActionsTestsApiShow**](AdTestApiAPI.md#ActionsTestsApiShow) | **Get** /agent/workspaces/{workspace}/tests/{ad_test} | Read a Test, verdict, and pause candidates
[**ActionsTestsApiStore**](AdTestApiAPI.md#ActionsTestsApiStore) | **Post** /agent/workspaces/{workspace}/tests | Create a judging Test



## ActionsTestsApiIndex

> AdTestListRead ActionsTestsApiIndex(ctx, workspace).PerPage(perPage).Page(page).Execute()

List Tests and observed conversion event choices



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
	perPage := int32(56) // int32 |  (optional)
	page := int64(789) // int64 | Integer page from 1 through 10000000000000000; integer strings are accepted. Zero, negatives, fractions, non-numeric strings and booleans return 422 with a page field validation error. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdTestApiAPI.ActionsTestsApiIndex(context.Background(), workspace).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdTestApiAPI.ActionsTestsApiIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsTestsApiIndex`: AdTestListRead
	fmt.Fprintf(os.Stdout, "Response from `AdTestApiAPI.ActionsTestsApiIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsTestsApiIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** |  | 
 **page** | **int64** | Integer page from 1 through 10000000000000000; integer strings are accepted. Zero, negatives, fractions, non-numeric strings and booleans return 422 with a page field validation error. | 

### Return type

[**AdTestListRead**](AdTestListRead.md)

### Authorization

[testReader](../README.md#testReader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsTestsApiShow

> AdTestShowRead ActionsTestsApiShow(ctx, workspace, adTest).Execute()

Read a Test, verdict, and pause candidates



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
	adTest := "adTest_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdTestApiAPI.ActionsTestsApiShow(context.Background(), workspace, adTest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdTestApiAPI.ActionsTestsApiShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsTestsApiShow`: AdTestShowRead
	fmt.Fprintf(os.Stdout, "Response from `AdTestApiAPI.ActionsTestsApiShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**adTest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsTestsApiShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AdTestShowRead**](AdTestShowRead.md)

### Authorization

[testReader](../README.md#testReader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsTestsApiStore

> AdTestCreateRead ActionsTestsApiStore(ctx, workspace).ActionsTestsApiStoreRequest(actionsTestsApiStoreRequest).Execute()

Create a judging Test



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
	actionsTestsApiStoreRequest := *openapiclient.NewActionsTestsApiStoreRequest(int32(123), "AdSetId_example", []string{"MemberAdIds_example"}, "Axis_example", "EventName_example") // ActionsTestsApiStoreRequest | A Test using ads present in the selected mirrored Meta ad set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdTestApiAPI.ActionsTestsApiStore(context.Background(), workspace).ActionsTestsApiStoreRequest(actionsTestsApiStoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdTestApiAPI.ActionsTestsApiStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsTestsApiStore`: AdTestCreateRead
	fmt.Fprintf(os.Stdout, "Response from `AdTestApiAPI.ActionsTestsApiStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsTestsApiStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actionsTestsApiStoreRequest** | [**ActionsTestsApiStoreRequest**](ActionsTestsApiStoreRequest.md) | A Test using ads present in the selected mirrored Meta ad set. | 

### Return type

[**AdTestCreateRead**](AdTestCreateRead.md)

### Authorization

[testWriter](../README.md#testWriter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

