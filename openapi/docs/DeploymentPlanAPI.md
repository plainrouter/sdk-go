# \DeploymentPlanAPI

All URIs are relative to *https://plainrouter.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LaunchPlansCopy**](DeploymentPlanAPI.md#LaunchPlansCopy) | **Post** /workspaces/{workspace}/admin/plans/{deployment_plan}/copy | Copy a failed plan to a new draft
[**LaunchPlansExecute**](DeploymentPlanAPI.md#LaunchPlansExecute) | **Post** /workspaces/{workspace}/admin/plans/{deployment_plan}/execute | Execute a deployment plan
[**LaunchPlansIndex**](DeploymentPlanAPI.md#LaunchPlansIndex) | **Get** /workspaces/{workspace}/admin/plans | List deployment plans
[**LaunchPlansShow**](DeploymentPlanAPI.md#LaunchPlansShow) | **Get** /workspaces/{workspace}/admin/plans/{deployment_plan} | Read a deployment plan
[**LaunchPlansStore**](DeploymentPlanAPI.md#LaunchPlansStore) | **Post** /workspaces/{workspace}/admin/plans | Create a deployment plan
[**LaunchPlansUpdate**](DeploymentPlanAPI.md#LaunchPlansUpdate) | **Patch** /workspaces/{workspace}/admin/plans/{deployment_plan} | Update a deployment plan
[**LaunchPlansValidate**](DeploymentPlanAPI.md#LaunchPlansValidate) | **Post** /workspaces/{workspace}/admin/plans/{deployment_plan}/validate | Validate a deployment plan



## LaunchPlansCopy

> PlanCopyRead LaunchPlansCopy(ctx, workspace, deploymentPlan).Execute()

Copy a failed plan to a new draft



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
	deploymentPlan := "deploymentPlan_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentPlanAPI.LaunchPlansCopy(context.Background(), workspace, deploymentPlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentPlanAPI.LaunchPlansCopy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchPlansCopy`: PlanCopyRead
	fmt.Fprintf(os.Stdout, "Response from `DeploymentPlanAPI.LaunchPlansCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**deploymentPlan** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchPlansCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PlanCopyRead**](PlanCopyRead.md)

### Authorization

[planWriter](../README.md#planWriter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchPlansExecute

> LaunchIntentRead LaunchPlansExecute(ctx, workspace, deploymentPlan).ExecuteDeploymentPlanRequest(executeDeploymentPlanRequest).Execute()

Execute a deployment plan



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
	deploymentPlan := "deploymentPlan_example" // string | 
	executeDeploymentPlanRequest := *openapiclient.NewExecuteDeploymentPlanRequest("ReviewVersion_example") // ExecuteDeploymentPlanRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentPlanAPI.LaunchPlansExecute(context.Background(), workspace, deploymentPlan).ExecuteDeploymentPlanRequest(executeDeploymentPlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentPlanAPI.LaunchPlansExecute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchPlansExecute`: LaunchIntentRead
	fmt.Fprintf(os.Stdout, "Response from `DeploymentPlanAPI.LaunchPlansExecute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**deploymentPlan** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchPlansExecuteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **executeDeploymentPlanRequest** | [**ExecuteDeploymentPlanRequest**](ExecuteDeploymentPlanRequest.md) |  | 

### Return type

[**LaunchIntentRead**](LaunchIntentRead.md)

### Authorization

[planWriter](../README.md#planWriter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchPlansIndex

> LaunchPlansIndex200Response LaunchPlansIndex(ctx, workspace).Cursor(cursor).PerPage(perPage).Page(page).Execute()

List deployment plans



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
	resp, r, err := apiClient.DeploymentPlanAPI.LaunchPlansIndex(context.Background(), workspace).Cursor(cursor).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentPlanAPI.LaunchPlansIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchPlansIndex`: LaunchPlansIndex200Response
	fmt.Fprintf(os.Stdout, "Response from `DeploymentPlanAPI.LaunchPlansIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchPlansIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **int32** | Continue using next_cursor with the same filters; rule lists use an immutable rule ID, other lists an offset or object ID. | 
 **perPage** | **int32** | Maximum complete rows; the byte budget may return fewer. Defaults to 25. | [default to 25]
 **page** | **int64** | Integer page from 1 through 10000000000000000; integer strings are accepted. Zero, negatives, fractions, non-numeric strings and booleans return 422 with a page field validation error. | 

### Return type

[**LaunchPlansIndex200Response**](LaunchPlansIndex200Response.md)

### Authorization

[planReader](../README.md#planReader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchPlansShow

> LaunchPlansShow200Response LaunchPlansShow(ctx, workspace, deploymentPlan).Cursor(cursor).Limit(limit).Execute()

Read a deployment plan



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
	deploymentPlan := "deploymentPlan_example" // string | 
	cursor := int32(56) // int32 | Continue using next_cursor with the same filters; rule lists use an immutable rule ID, other lists an offset or object ID. (optional)
	limit := int32(56) // int32 | Maximum complete rows; the byte budget may return fewer. Defaults to 5. (optional) (default to 5)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentPlanAPI.LaunchPlansShow(context.Background(), workspace, deploymentPlan).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentPlanAPI.LaunchPlansShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchPlansShow`: LaunchPlansShow200Response
	fmt.Fprintf(os.Stdout, "Response from `DeploymentPlanAPI.LaunchPlansShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**deploymentPlan** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchPlansShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **int32** | Continue using next_cursor with the same filters; rule lists use an immutable rule ID, other lists an offset or object ID. | 
 **limit** | **int32** | Maximum complete rows; the byte budget may return fewer. Defaults to 5. | [default to 5]

### Return type

[**LaunchPlansShow200Response**](LaunchPlansShow200Response.md)

### Authorization

[planReader](../README.md#planReader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchPlansStore

> LaunchPlansStore201Response LaunchPlansStore(ctx, workspace).StoreDeploymentPlanRequest(storeDeploymentPlanRequest).Execute()

Create a deployment plan



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
	storeDeploymentPlanRequest := *openapiclient.NewStoreDeploymentPlanRequest(int32(123), "Platform_example", "BudgetType_example", int32(123), "Currency_example", *openapiclient.NewStoreDeploymentPlanRequestCopy("PrimaryText_example", "Headline_example", "Description_example", "Cta_example"), "LandingPage_example") // StoreDeploymentPlanRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentPlanAPI.LaunchPlansStore(context.Background(), workspace).StoreDeploymentPlanRequest(storeDeploymentPlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentPlanAPI.LaunchPlansStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchPlansStore`: LaunchPlansStore201Response
	fmt.Fprintf(os.Stdout, "Response from `DeploymentPlanAPI.LaunchPlansStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchPlansStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storeDeploymentPlanRequest** | [**StoreDeploymentPlanRequest**](StoreDeploymentPlanRequest.md) |  | 

### Return type

[**LaunchPlansStore201Response**](LaunchPlansStore201Response.md)

### Authorization

[workspaceSecret](../README.md#workspaceSecret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchPlansUpdate

> LaunchPlansStore201Response LaunchPlansUpdate(ctx, workspace, deploymentPlan).UpdateDeploymentPlanRequest(updateDeploymentPlanRequest).Execute()

Update a deployment plan



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
	deploymentPlan := "deploymentPlan_example" // string | 
	updateDeploymentPlanRequest := *openapiclient.NewUpdateDeploymentPlanRequest() // UpdateDeploymentPlanRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentPlanAPI.LaunchPlansUpdate(context.Background(), workspace, deploymentPlan).UpdateDeploymentPlanRequest(updateDeploymentPlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentPlanAPI.LaunchPlansUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchPlansUpdate`: LaunchPlansStore201Response
	fmt.Fprintf(os.Stdout, "Response from `DeploymentPlanAPI.LaunchPlansUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**deploymentPlan** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchPlansUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDeploymentPlanRequest** | [**UpdateDeploymentPlanRequest**](UpdateDeploymentPlanRequest.md) |  | 

### Return type

[**LaunchPlansStore201Response**](LaunchPlansStore201Response.md)

### Authorization

[workspaceSecret](../README.md#workspaceSecret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchPlansValidate

> LaunchPlansShow200Response LaunchPlansValidate(ctx, workspace, deploymentPlan).Execute()

Validate a deployment plan



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
	deploymentPlan := "deploymentPlan_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentPlanAPI.LaunchPlansValidate(context.Background(), workspace, deploymentPlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentPlanAPI.LaunchPlansValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LaunchPlansValidate`: LaunchPlansShow200Response
	fmt.Fprintf(os.Stdout, "Response from `DeploymentPlanAPI.LaunchPlansValidate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**deploymentPlan** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchPlansValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LaunchPlansShow200Response**](LaunchPlansShow200Response.md)

### Authorization

[planWriter](../README.md#planWriter)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

