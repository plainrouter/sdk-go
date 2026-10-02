# \DeploymentPlanAPI

All URIs are relative to *https://plainrouter.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LaunchPlansCopy**](DeploymentPlanAPI.md#LaunchPlansCopy) | **Post** /workspaces/{workspace}/admin/plans/{deployment_plan}/copy | Copy a failed plan to a new draft



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

