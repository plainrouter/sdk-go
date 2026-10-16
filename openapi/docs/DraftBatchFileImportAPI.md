# \DraftBatchFileImportAPI

All URIs are relative to *https://plainrouter.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LauncherFileImport**](DraftBatchFileImportAPI.md#LauncherFileImport) | **Post** /mcp/launcher/import | Import a Launcher draft batch



## LauncherFileImport

> map[string]*interface{} LauncherFileImport(ctx).File(file).SourceAdsetId(sourceAdsetId).SourceAdId(sourceAdId).Mode(mode).FolderToAdset(folderToAdset).GroupingEnabled(groupingEnabled).AccountId(accountId).WorkspaceId(workspaceId).Execute()

Import a Launcher draft batch



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
	file := os.NewFile(1234, "some_file") // *os.File | 
	sourceAdsetId := "sourceAdsetId_example" // string | 
	sourceAdId := "sourceAdId_example" // string | 
	mode := openapiclient.DraftBatchMode("duplicate_adset") // DraftBatchMode | 
	folderToAdset := true // bool |  (optional)
	groupingEnabled := true // bool |  (optional)
	accountId := int32(56) // int32 |  (optional)
	workspaceId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DraftBatchFileImportAPI.LauncherFileImport(context.Background()).File(file).SourceAdsetId(sourceAdsetId).SourceAdId(sourceAdId).Mode(mode).FolderToAdset(folderToAdset).GroupingEnabled(groupingEnabled).AccountId(accountId).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DraftBatchFileImportAPI.LauncherFileImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LauncherFileImport`: map[string]*interface{}
	fmt.Fprintf(os.Stdout, "Response from `DraftBatchFileImportAPI.LauncherFileImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLauncherFileImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 
 **sourceAdsetId** | **string** |  | 
 **sourceAdId** | **string** |  | 
 **mode** | [**DraftBatchMode**](DraftBatchMode.md) |  | 
 **folderToAdset** | **bool** |  | 
 **groupingEnabled** | **bool** |  | 
 **accountId** | **int32** |  | 
 **workspaceId** | **int32** |  | 

### Return type

**map[string]*interface{}**

### Authorization

[workspaceSecret](../README.md#workspaceSecret)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

