# \ActionDryRunApiAPI

All URIs are relative to *https://plainrouter.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionsApiDryRun**](ActionDryRunApiAPI.md#ActionsApiDryRun) | **Post** /agent/workspaces/{workspace}/actions/dry-run | Preview actions



## ActionsApiDryRun

> ActionDryRunRead ActionsApiDryRun(ctx, workspace).ActionProposalInput(actionProposalInput).Execute()

Preview actions



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
	actionProposalInput := *openapiclient.NewActionProposalInput([]openapiclient.ActionProposalInputActionsInner{*openapiclient.NewActionProposalInputActionsInner("Type_example", *openapiclient.NewActionProposalInputActionsInnerAnyOf4TargetEntity("Type_example", "Id_example"), *openapiclient.NewActionProposalInputActionsInnerAnyOf13Params("AssetId_example", "Status_example"), "Rationale_example")}, "Rationale_example", "IdempotencyKey_example", []openapiclient.ActionProposalInputEvidenceInner{*openapiclient.NewActionProposalInputEvidenceInner("SourceTool_example", []string{"FieldsUsed_example"})}, "TargetSource_example") // ActionProposalInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionDryRunApiAPI.ActionsApiDryRun(context.Background(), workspace).ActionProposalInput(actionProposalInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionDryRunApiAPI.ActionsApiDryRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsApiDryRun`: ActionDryRunRead
	fmt.Fprintf(os.Stdout, "Response from `ActionDryRunApiAPI.ActionsApiDryRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionsApiDryRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actionProposalInput** | [**ActionProposalInput**](ActionProposalInput.md) |  | 

### Return type

[**ActionDryRunRead**](ActionDryRunRead.md)

### Authorization

[workspaceActionKey](../README.md#workspaceActionKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

