# \MetaInventoryExposureAPI

All URIs are relative to *https://plainrouter.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountInventory**](MetaInventoryExposureAPI.md#GetAccountInventory) | **Get** /workspaces/{workspace}/admin/ad-accounts/{account_id}/inventory | Read account inventory
[**GetInventoryMetrics**](MetaInventoryExposureAPI.md#GetInventoryMetrics) | **Get** /workspaces/{workspace}/admin/ad-accounts/{account_id}/inventory/metrics | Compare Meta metrics with counted arrivals



## GetAccountInventory

> GetAccountInventory200Response GetAccountInventory(ctx, workspace, accountId).Level(level).EffectiveStatus(effectiveStatus).CampaignId(campaignId).AdsetId(adsetId).AfterId(afterId).Limit(limit).Execute()

Read account inventory



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
	accountId := "accountId_example" // string | 
	level := "level_example" // string |  (optional)
	effectiveStatus := "effectiveStatus_example" // string |  (optional)
	campaignId := "campaignId_example" // string |  (optional)
	adsetId := "adsetId_example" // string |  (optional)
	afterId := int32(56) // int32 | Continue using next_cursor with the same filters; rule lists use an immutable rule ID, other lists an offset or object ID. (optional)
	limit := int32(56) // int32 | Maximum complete rows; the byte budget may return fewer. Defaults to 25. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaInventoryExposureAPI.GetAccountInventory(context.Background(), workspace, accountId).Level(level).EffectiveStatus(effectiveStatus).CampaignId(campaignId).AdsetId(adsetId).AfterId(afterId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaInventoryExposureAPI.GetAccountInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInventory`: GetAccountInventory200Response
	fmt.Fprintf(os.Stdout, "Response from `MetaInventoryExposureAPI.GetAccountInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **level** | **string** |  | 
 **effectiveStatus** | **string** |  | 
 **campaignId** | **string** |  | 
 **adsetId** | **string** |  | 
 **afterId** | **int32** | Continue using next_cursor with the same filters; rule lists use an immutable rule ID, other lists an offset or object ID. | 
 **limit** | **int32** | Maximum complete rows; the byte budget may return fewer. Defaults to 25. | [default to 25]

### Return type

[**GetAccountInventory200Response**](GetAccountInventory200Response.md)

### Authorization

[workspaceSecret](../README.md#workspaceSecret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInventoryMetrics

> GetInventoryMetrics200Response GetInventoryMetrics(ctx, workspace, accountId).StartDate(startDate).EndDate(endDate).CampaignId(campaignId).AdsetId(adsetId).AdId(adId).IncludeChanges(includeChanges).Limit(limit).Cursor(cursor).Execute()

Compare Meta metrics with counted arrivals



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/plainrouter/sdk-go/openapi"
)

func main() {
	workspace := int32(56) // int32 | The workspace ID
	accountId := "accountId_example" // string | 
	startDate := time.Now() // string |  (optional)
	endDate := time.Now() // string |  (optional)
	campaignId := "campaignId_example" // string |  (optional)
	adsetId := "adsetId_example" // string |  (optional)
	adId := "adId_example" // string |  (optional)
	includeChanges := true // bool |  (optional)
	limit := int32(56) // int32 |  (optional)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetaInventoryExposureAPI.GetInventoryMetrics(context.Background(), workspace, accountId).StartDate(startDate).EndDate(endDate).CampaignId(campaignId).AdsetId(adsetId).AdId(adId).IncludeChanges(includeChanges).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetaInventoryExposureAPI.GetInventoryMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInventoryMetrics`: GetInventoryMetrics200Response
	fmt.Fprintf(os.Stdout, "Response from `MetaInventoryExposureAPI.GetInventoryMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **int32** | The workspace ID | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInventoryMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **campaignId** | **string** |  | 
 **adsetId** | **string** |  | 
 **adId** | **string** |  | 
 **includeChanges** | **bool** |  | 
 **limit** | **int32** |  | 
 **cursor** | **string** |  | 

### Return type

[**GetInventoryMetrics200Response**](GetInventoryMetrics200Response.md)

### Authorization

[workspaceInventoryKey](../README.md#workspaceInventoryKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

