# \LaunchPolicyAPI

All URIs are relative to *https://plainrouter.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LaunchPoliciesIndex**](LaunchPolicyAPI.md#LaunchPoliciesIndex) | **Get** /workspaces/{workspace}/admin/policies | List Launch policies
[**LaunchPoliciesRetire**](LaunchPolicyAPI.md#LaunchPoliciesRetire) | **Post** /workspaces/{workspace}/admin/policies/{launch_policy}/retire | Retire a Launch policy
[**LaunchPoliciesStore**](LaunchPolicyAPI.md#LaunchPoliciesStore) | **Post** /workspaces/{workspace}/admin/policies | Create a Launch policy



## LaunchPoliciesIndex

> LaunchPoliciesIndex200Response LaunchPoliciesIndex(ctx, workspace).PerPage(perPage).Page(page).Execute()

List Launch policies



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
	resp, r, err := apiClient.LaunchPolicyAPI.LaunchPoliciesIndex(context.Background(), workspace).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LaunchPolicyAPI.LaunchPoliciesIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchPoliciesIndex`: LaunchPoliciesIndex200Response
	fmt.Fprintf(os.Stdout, "Response from `LaunchPolicyAPI.LaunchPoliciesIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchPoliciesIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** |  | 
 **page** | **int64** | Integer page from 1 through 10000000000000000; integer strings are accepted. Zero, negatives, fractions, non-numeric strings and booleans return 422 with a page field validation error. | 

### Return type

[**LaunchPoliciesIndex200Response**](LaunchPoliciesIndex200Response.md)

### Authorization

[workspaceSecret](../README.md#workspaceSecret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchPoliciesRetire

> LaunchPoliciesStore201Response LaunchPoliciesRetire(ctx, workspace, launchPolicy).Execute()

Retire a Launch policy



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
	launchPolicy := "launchPolicy_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LaunchPolicyAPI.LaunchPoliciesRetire(context.Background(), workspace, launchPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LaunchPolicyAPI.LaunchPoliciesRetire``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchPoliciesRetire`: LaunchPoliciesStore201Response
	fmt.Fprintf(os.Stdout, "Response from `LaunchPolicyAPI.LaunchPoliciesRetire`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**launchPolicy** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchPoliciesRetireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LaunchPoliciesStore201Response**](LaunchPoliciesStore201Response.md)

### Authorization

[workspaceSecret](../README.md#workspaceSecret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchPoliciesStore

> LaunchPoliciesStore201Response LaunchPoliciesStore(ctx, workspace).StoreLaunchPolicyRequest(storeLaunchPolicyRequest).Execute()

Create a Launch policy



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
	storeLaunchPolicyRequest := *openapiclient.NewStoreLaunchPolicyRequest("Kind_example", "Name_example", []string{"Body_example"}) // StoreLaunchPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LaunchPolicyAPI.LaunchPoliciesStore(context.Background(), workspace).StoreLaunchPolicyRequest(storeLaunchPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LaunchPolicyAPI.LaunchPoliciesStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchPoliciesStore`: LaunchPoliciesStore201Response
	fmt.Fprintf(os.Stdout, "Response from `LaunchPolicyAPI.LaunchPoliciesStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchPoliciesStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storeLaunchPolicyRequest** | [**StoreLaunchPolicyRequest**](StoreLaunchPolicyRequest.md) |  | 

### Return type

[**LaunchPoliciesStore201Response**](LaunchPoliciesStore201Response.md)

### Authorization

[workspaceSecret](../README.md#workspaceSecret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

