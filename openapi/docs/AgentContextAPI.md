# \AgentContextAPI

All URIs are relative to *https://plainrouter.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentContext**](AgentContextAPI.md#AgentContext) | **Get** /agent/context | Get agent context



## AgentContext

> AgentContext200Response AgentContext(ctx).AccountId(accountId).WorkspaceId(workspaceId).Execute()

Get agent context



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
	accountId := int32(56) // int32 | Optional positive account ID. A bound key may name only its own account; a different ID returns 422. An unbound key requires exactly one active available advertising account with an active connection; omission uses that account and a supplied ID must match it. Multiple active accounts return 422 even when account_id is supplied. With no active available account, a valid account_id is ignored and platform_ad_account is null. (optional)
	workspaceId := int32(56) // int32 | Optional workspace grant selector. Required for credentials with multiple workspace grants; it must identify an authorized workspace. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentContextAPI.AgentContext(context.Background()).AccountId(accountId).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentContextAPI.AgentContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AgentContext`: AgentContext200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentContextAPI.AgentContext`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAgentContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32** | Optional positive account ID. A bound key may name only its own account; a different ID returns 422. An unbound key requires exactly one active available advertising account with an active connection; omission uses that account and a supplied ID must match it. Multiple active accounts return 422 even when account_id is supplied. With no active available account, a valid account_id is ignored and platform_ad_account is null. | 
 **workspaceId** | **int32** | Optional workspace grant selector. Required for credentials with multiple workspace grants; it must identify an authorized workspace. | 

### Return type

[**AgentContext200Response**](AgentContext200Response.md)

### Authorization

[agentContextOAuth](../README.md#agentContextOAuth), [agentContextKey](../README.md#agentContextKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

