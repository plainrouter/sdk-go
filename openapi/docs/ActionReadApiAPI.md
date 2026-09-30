# \ActionReadApiAPI

All URIs are relative to *https://plainrouter.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionsApiBatch**](ActionReadApiAPI.md#ActionsApiBatch) | **Get** /agent/workspaces/{workspace}/action-batches/{actionBatch} | Get action batch
[**ActionsApiDecisionReceipt**](ActionReadApiAPI.md#ActionsApiDecisionReceipt) | **Get** /agent/workspaces/{workspace}/actions/{action}/decision-receipt | Get decision receipt
[**ActionsApiIndex**](ActionReadApiAPI.md#ActionsApiIndex) | **Get** /agent/workspaces/{workspace}/actions | List actions
[**ActionsApiPolicy**](ActionReadApiAPI.md#ActionsApiPolicy) | **Get** /agent/workspaces/{workspace}/actions/policy | Get action policy
[**ActionsApiShow**](ActionReadApiAPI.md#ActionsApiShow) | **Get** /agent/workspaces/{workspace}/actions/{action} | Get action



## ActionsApiBatch

> ActionBatchRead ActionsApiBatch(ctx, workspace, actionBatch).AccountId(accountId).Execute()

Get action batch



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
	actionBatch := "actionBatch_example" // string | 
	accountId := int32(56) // int32 | Optional positive account ID. A bound key may name only its own account; any other value returns 422. An unbound key may name only an active account with an active connection in its workspace, and must name one when the workspace has more than one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionReadApiAPI.ActionsApiBatch(context.Background(), workspace, actionBatch).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionReadApiAPI.ActionsApiBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsApiBatch`: ActionBatchRead
	fmt.Fprintf(os.Stdout, "Response from `ActionReadApiAPI.ActionsApiBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**actionBatch** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsApiBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **int32** | Optional positive account ID. A bound key may name only its own account; any other value returns 422. An unbound key may name only an active account with an active connection in its workspace, and must name one when the workspace has more than one. | 

### Return type

[**ActionBatchRead**](ActionBatchRead.md)

### Authorization

[workspaceActionKey](../README.md#workspaceActionKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsApiDecisionReceipt

> ActionDecisionReceiptRead ActionsApiDecisionReceipt(ctx, workspace, action).AccountId(accountId).Execute()

Get decision receipt



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
	action := "action_example" // string | 
	accountId := int32(56) // int32 | Optional positive account ID. A bound key may name only its own account; any other value returns 422. An unbound key may name only an active account with an active connection in its workspace, and must name one when the workspace has more than one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionReadApiAPI.ActionsApiDecisionReceipt(context.Background(), workspace, action).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionReadApiAPI.ActionsApiDecisionReceipt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsApiDecisionReceipt`: ActionDecisionReceiptRead
	fmt.Fprintf(os.Stdout, "Response from `ActionReadApiAPI.ActionsApiDecisionReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**action** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsApiDecisionReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **int32** | Optional positive account ID. A bound key may name only its own account; any other value returns 422. An unbound key may name only an active account with an active connection in its workspace, and must name one when the workspace has more than one. | 

### Return type

[**ActionDecisionReceiptRead**](ActionDecisionReceiptRead.md)

### Authorization

[workspaceActionKey](../README.md#workspaceActionKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsApiIndex

> ActionListRead ActionsApiIndex(ctx, workspace).Status(status).Page(page).AccountId(accountId).Execute()

List actions



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
	status := "status_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	accountId := int32(56) // int32 | Optional positive account ID. A bound key may name only its own account; any other value returns 422. An unbound key may name only an active account with an active connection in its workspace, and must name one when the workspace has more than one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionReadApiAPI.ActionsApiIndex(context.Background(), workspace).Status(status).Page(page).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionReadApiAPI.ActionsApiIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsApiIndex`: ActionListRead
	fmt.Fprintf(os.Stdout, "Response from `ActionReadApiAPI.ActionsApiIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsApiIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** |  | 
 **page** | **int32** |  | 
 **accountId** | **int32** | Optional positive account ID. A bound key may name only its own account; any other value returns 422. An unbound key may name only an active account with an active connection in its workspace, and must name one when the workspace has more than one. | 

### Return type

[**ActionListRead**](ActionListRead.md)

### Authorization

[workspaceActionKey](../README.md#workspaceActionKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsApiPolicy

> ActionPolicyRead ActionsApiPolicy(ctx, workspace).AccountId(accountId).Execute()

Get action policy



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
	accountId := int32(56) // int32 | Optional positive account ID. A bound key may name only its own account; any other value returns 422. An unbound key may name only an active account with an active connection in its workspace, and must name one when the workspace has more than one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionReadApiAPI.ActionsApiPolicy(context.Background(), workspace).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionReadApiAPI.ActionsApiPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsApiPolicy`: ActionPolicyRead
	fmt.Fprintf(os.Stdout, "Response from `ActionReadApiAPI.ActionsApiPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsApiPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **int32** | Optional positive account ID. A bound key may name only its own account; any other value returns 422. An unbound key may name only an active account with an active connection in its workspace, and must name one when the workspace has more than one. | 

### Return type

[**ActionPolicyRead**](ActionPolicyRead.md)

### Authorization

[workspaceActionKey](../README.md#workspaceActionKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsApiShow

> ActionDetailRead ActionsApiShow(ctx, workspace, action).AccountId(accountId).Execute()

Get action



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
	action := "action_example" // string | 
	accountId := int32(56) // int32 | Optional positive account ID. A bound key may name only its own account; any other value returns 422. An unbound key may name only an active account with an active connection in its workspace, and must name one when the workspace has more than one. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionReadApiAPI.ActionsApiShow(context.Background(), workspace, action).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionReadApiAPI.ActionsApiShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsApiShow`: ActionDetailRead
	fmt.Fprintf(os.Stdout, "Response from `ActionReadApiAPI.ActionsApiShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**action** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsApiShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **int32** | Optional positive account ID. A bound key may name only its own account; any other value returns 422. An unbound key may name only an active account with an active connection in its workspace, and must name one when the workspace has more than one. | 

### Return type

[**ActionDetailRead**](ActionDetailRead.md)

### Authorization

[workspaceActionKey](../README.md#workspaceActionKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

