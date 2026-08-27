# \EventAPI

All URIs are relative to *https://plainrouter.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEvent**](EventAPI.md#CreateEvent) | **Post** /events | Submit a conversion event
[**GetEvent**](EventAPI.md#GetEvent) | **Get** /events/{event} | Get an event and delivery trace
[**VerifySignalIngestion**](EventAPI.md#VerifySignalIngestion) | **Post** /verification-events | Verify server-side Signal ingestion



## CreateEvent

> CreateEvent200Response CreateEvent(ctx).CreateEventRequest(createEventRequest).IdempotencyKey(idempotencyKey).Execute()

Submit a conversion event



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
	createEventRequest := *openapiclient.NewCreateEventRequest("EventName_example", "ConsentBasis_example") // CreateEventRequest | 
	idempotencyKey := "idempotencyKey_example" // string | Optional idempotency key. When event_id is omitted, PlainRouter uses this value as event_id. If both are supplied, they must match. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventAPI.CreateEvent(context.Background()).CreateEventRequest(createEventRequest).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.CreateEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEvent`: CreateEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `EventAPI.CreateEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEventRequest** | [**CreateEventRequest**](CreateEventRequest.md) |  | 
 **idempotencyKey** | **string** | Optional idempotency key. When event_id is omitted, PlainRouter uses this value as event_id. If both are supplied, they must match. | 

### Return type

[**CreateEvent200Response**](CreateEvent200Response.md)

### Authorization

[signalTrackerSecret](../README.md#signalTrackerSecret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvent

> GetEvent200Response GetEvent(ctx, event).Execute()

Get an event and delivery trace



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
	event := "event_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventAPI.GetEvent(context.Background(), event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.GetEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvent`: GetEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `EventAPI.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**event** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEvent200Response**](GetEvent200Response.md)

### Authorization

[signalTrackerSecret](../README.md#signalTrackerSecret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifySignalIngestion

> CreateEvent200Response VerifySignalIngestion(ctx).Execute()

Verify server-side Signal ingestion



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventAPI.VerifySignalIngestion(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.VerifySignalIngestion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifySignalIngestion`: CreateEvent200Response
	fmt.Fprintf(os.Stdout, "Response from `EventAPI.VerifySignalIngestion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVerifySignalIngestionRequest struct via the builder pattern


### Return type

[**CreateEvent200Response**](CreateEvent200Response.md)

### Authorization

[signalTrackerSecret](../README.md#signalTrackerSecret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

