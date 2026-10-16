# \ActionRuleApiAPI

All URIs are relative to *https://plainrouter.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionsRulesIndex**](ActionRuleApiAPI.md#ActionsRulesIndex) | **Get** /agent/workspaces/{workspace}/action-rules | List rule summaries
[**ActionsRulesShow**](ActionRuleApiAPI.md#ActionsRulesShow) | **Get** /agent/workspaces/{workspace}/action-rules/{action_rule} | Get an action rule
[**ActionsRulesStore**](ActionRuleApiAPI.md#ActionsRulesStore) | **Post** /agent/workspaces/{workspace}/action-rules | Create an action rule
[**ActionsRulesUpdate**](ActionRuleApiAPI.md#ActionsRulesUpdate) | **Patch** /agent/workspaces/{workspace}/action-rules/{action_rule} | Update an action rule



## ActionsRulesIndex

> ActionsRulesIndex200Response ActionsRulesIndex(ctx, workspace).AccountId(accountId).PlatformAdAccountId(platformAdAccountId).Cursor(cursor).PerPage(perPage).Page(page).Execute()

List rule summaries



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
	accountId := int32(56) // int32 | Optional integer account ID. JSON booleans return 422 with message plus field errors. A valid supplied ID is ignored by an account-bound key. An unbound key requires exactly one active available advertising account with an active connection; omission uses that account and a supplied ID must match it. Multiple active accounts return 422 even when account_id is supplied. (optional)
	platformAdAccountId := int32(56) // int32 | Optional integer account ID. JSON booleans return 422 with message plus field errors. A valid supplied ID is ignored by an account-bound key. An unbound key requires exactly one active available advertising account with an active connection; omission uses that account and a supplied ID must match it. Multiple active accounts return 422 even when account_id is supplied. Alias of account_id; account_id takes precedence when both are supplied. (optional)
	cursor := "cursor_example" // string | Continue using next_cursor with the same filters; rule lists use an immutable rule ID, other lists an offset or object ID. (optional)
	perPage := int32(56) // int32 | Maximum complete rows; the byte budget may return fewer. Defaults to 25. (optional) (default to 25)
	page := int64(789) // int64 | Integer page from 1 through 10000000000000000; integer strings are accepted. Zero, negatives, fractions, non-numeric strings and booleans return 422 with a page field validation error. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionRuleApiAPI.ActionsRulesIndex(context.Background(), workspace).AccountId(accountId).PlatformAdAccountId(platformAdAccountId).Cursor(cursor).PerPage(perPage).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionRuleApiAPI.ActionsRulesIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsRulesIndex`: ActionsRulesIndex200Response
	fmt.Fprintf(os.Stdout, "Response from `ActionRuleApiAPI.ActionsRulesIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsRulesIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **int32** | Optional integer account ID. JSON booleans return 422 with message plus field errors. A valid supplied ID is ignored by an account-bound key. An unbound key requires exactly one active available advertising account with an active connection; omission uses that account and a supplied ID must match it. Multiple active accounts return 422 even when account_id is supplied. | 
 **platformAdAccountId** | **int32** | Optional integer account ID. JSON booleans return 422 with message plus field errors. A valid supplied ID is ignored by an account-bound key. An unbound key requires exactly one active available advertising account with an active connection; omission uses that account and a supplied ID must match it. Multiple active accounts return 422 even when account_id is supplied. Alias of account_id; account_id takes precedence when both are supplied. | 
 **cursor** | **string** | Continue using next_cursor with the same filters; rule lists use an immutable rule ID, other lists an offset or object ID. | 
 **perPage** | **int32** | Maximum complete rows; the byte budget may return fewer. Defaults to 25. | [default to 25]
 **page** | **int64** | Integer page from 1 through 10000000000000000; integer strings are accepted. Zero, negatives, fractions, non-numeric strings and booleans return 422 with a page field validation error. | 

### Return type

[**ActionsRulesIndex200Response**](ActionsRulesIndex200Response.md)

### Authorization

[ruleReader](../README.md#ruleReader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsRulesShow

> ActionRuleDetailRead ActionsRulesShow(ctx, workspace, actionRule).AccountId(accountId).Execute()

Get an action rule



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
	actionRule := "actionRule_example" // string | 
	accountId := int32(56) // int32 | Optional integer account ID. JSON booleans return 422 with message plus field errors. A valid supplied ID is ignored by an account-bound key. An unbound key requires exactly one active available advertising account with an active connection; omission uses that account and a supplied ID must match it. Multiple active accounts return 422 even when account_id is supplied. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionRuleApiAPI.ActionsRulesShow(context.Background(), workspace, actionRule).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionRuleApiAPI.ActionsRulesShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsRulesShow`: ActionRuleDetailRead
	fmt.Fprintf(os.Stdout, "Response from `ActionRuleApiAPI.ActionsRulesShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**actionRule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsRulesShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **int32** | Optional integer account ID. JSON booleans return 422 with message plus field errors. A valid supplied ID is ignored by an account-bound key. An unbound key requires exactly one active available advertising account with an active connection; omission uses that account and a supplied ID must match it. Multiple active accounts return 422 even when account_id is supplied. | 

### Return type

[**ActionRuleDetailRead**](ActionRuleDetailRead.md)

### Authorization

[ruleReader](../README.md#ruleReader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsRulesStore

> ActionRuleWriteRead ActionsRulesStore(ctx, workspace).ActionsRulesStoreRequest(actionsRulesStoreRequest).Execute()

Create an action rule



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
	actionsRulesStoreRequest := *openapiclient.NewActionsRulesStoreRequest("Name_example", *openapiclient.NewActionsRulesStoreRequestDefinition("TargetLevel_example", "Action_example", int32(123), "CostComparison_example", *openapiclient.NewActionsRulesStoreRequestDefinitionAnyOfCostRatio(), int32(123)), int32(123)) // ActionsRulesStoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionRuleApiAPI.ActionsRulesStore(context.Background(), workspace).ActionsRulesStoreRequest(actionsRulesStoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionRuleApiAPI.ActionsRulesStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsRulesStore`: ActionRuleWriteRead
	fmt.Fprintf(os.Stdout, "Response from `ActionRuleApiAPI.ActionsRulesStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsRulesStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actionsRulesStoreRequest** | [**ActionsRulesStoreRequest**](ActionsRulesStoreRequest.md) |  | 

### Return type

[**ActionRuleWriteRead**](ActionRuleWriteRead.md)

### Authorization

[ruleWriter](../README.md#ruleWriter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionsRulesUpdate

> ActionRuleWriteRead ActionsRulesUpdate(ctx, workspace, actionRule).ActionsRulesUpdateRequest(actionsRulesUpdateRequest).Execute()

Update an action rule



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
	actionRule := "actionRule_example" // string | 
	actionsRulesUpdateRequest := *openapiclient.NewActionsRulesUpdateRequest() // ActionsRulesUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionRuleApiAPI.ActionsRulesUpdate(context.Background(), workspace, actionRule).ActionsRulesUpdateRequest(actionsRulesUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionRuleApiAPI.ActionsRulesUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsRulesUpdate`: ActionRuleWriteRead
	fmt.Fprintf(os.Stdout, "Response from `ActionRuleApiAPI.ActionsRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**actionRule** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **actionsRulesUpdateRequest** | [**ActionsRulesUpdateRequest**](ActionsRulesUpdateRequest.md) |  | 

### Return type

[**ActionRuleWriteRead**](ActionRuleWriteRead.md)

### Authorization

[ruleWriter](../README.md#ruleWriter)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

