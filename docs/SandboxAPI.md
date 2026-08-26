# \SandboxAPI

All URIs are relative to *https://plainrouter.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSandboxKey**](SandboxAPI.md#CreateSandboxKey) | **Post** /sandbox/keys | Create a sandbox API key
[**GetSandbox**](SandboxAPI.md#GetSandbox) | **Get** /sandbox | Discover the zero-auth sandbox
[**ValidateSandboxEvent**](SandboxAPI.md#ValidateSandboxEvent) | **Post** /sandbox/events | Validate a synthetic event
[**ValidateSandboxEventWithKey**](SandboxAPI.md#ValidateSandboxEventWithKey) | **Post** /sandbox/keyed-events | Validate a synthetic event with a sandbox key



## CreateSandboxKey

> CreateSandboxKey201Response CreateSandboxKey(ctx).Execute()

Create a sandbox API key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wudaku/plainrouter-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.CreateSandboxKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.CreateSandboxKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSandboxKey`: CreateSandboxKey201Response
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.CreateSandboxKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSandboxKeyRequest struct via the builder pattern


### Return type

[**CreateSandboxKey201Response**](CreateSandboxKey201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSandbox

> GetSandbox200Response GetSandbox(ctx).Execute()

Discover the zero-auth sandbox



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wudaku/plainrouter-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.GetSandbox(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.GetSandbox``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSandbox`: GetSandbox200Response
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.GetSandbox`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxRequest struct via the builder pattern


### Return type

[**GetSandbox200Response**](GetSandbox200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateSandboxEvent

> ValidateSandboxEvent200Response ValidateSandboxEvent(ctx).ValidateSandboxEventRequest(validateSandboxEventRequest).Execute()

Validate a synthetic event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wudaku/plainrouter-go"
)

func main() {
	validateSandboxEventRequest := *openapiclient.NewValidateSandboxEventRequest("EventName_example") // ValidateSandboxEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.ValidateSandboxEvent(context.Background()).ValidateSandboxEventRequest(validateSandboxEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.ValidateSandboxEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateSandboxEvent`: ValidateSandboxEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.ValidateSandboxEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateSandboxEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateSandboxEventRequest** | [**ValidateSandboxEventRequest**](ValidateSandboxEventRequest.md) |  | 

### Return type

[**ValidateSandboxEvent200Response**](ValidateSandboxEvent200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateSandboxEventWithKey

> ValidateSandboxEvent200Response ValidateSandboxEventWithKey(ctx).ValidateSandboxEventRequest(validateSandboxEventRequest).Execute()

Validate a synthetic event with a sandbox key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wudaku/plainrouter-go"
)

func main() {
	validateSandboxEventRequest := *openapiclient.NewValidateSandboxEventRequest("EventName_example") // ValidateSandboxEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SandboxAPI.ValidateSandboxEventWithKey(context.Background()).ValidateSandboxEventRequest(validateSandboxEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SandboxAPI.ValidateSandboxEventWithKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateSandboxEventWithKey`: ValidateSandboxEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `SandboxAPI.ValidateSandboxEventWithKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateSandboxEventWithKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateSandboxEventRequest** | [**ValidateSandboxEventRequest**](ValidateSandboxEventRequest.md) |  | 

### Return type

[**ValidateSandboxEvent200Response**](ValidateSandboxEvent200Response.md)

### Authorization

[sandboxCredential](../README.md#sandboxCredential)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

