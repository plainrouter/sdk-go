# \CookieDeclarationAPI

All URIs are relative to *https://plainrouter.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCookieDeclaration**](CookieDeclarationAPI.md#GetCookieDeclaration) | **Get** /cookie-declarations/{publishableKey} | Read the public cookie declaration



## GetCookieDeclaration

> CookieDeclaration GetCookieDeclaration(ctx, publishableKey).Execute()

Read the public cookie declaration



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
	publishableKey := "publishableKey_example" // string | Workspace publishable key from the customer page source; pk_live_ followed by 32 alphanumeric characters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CookieDeclarationAPI.GetCookieDeclaration(context.Background(), publishableKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CookieDeclarationAPI.GetCookieDeclaration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCookieDeclaration`: CookieDeclaration
	fmt.Fprintf(os.Stdout, "Response from `CookieDeclarationAPI.GetCookieDeclaration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publishableKey** | **string** | Workspace publishable key from the customer page source; pk_live_ followed by 32 alphanumeric characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCookieDeclarationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CookieDeclaration**](CookieDeclaration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

