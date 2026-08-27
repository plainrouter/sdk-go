# \OperationsAPI

All URIs are relative to *https://plainrouter.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserData**](OperationsAPI.md#DeleteUserData) | **Delete** /user-data | Delete user data by hashed identifier
[**GetEmqReport**](OperationsAPI.md#GetEmqReport) | **Get** /reports/emq | Get Event Match Quality history
[**GetReconciliationReport**](OperationsAPI.md#GetReconciliationReport) | **Get** /reports/reconciliation | Get a reconciliation report
[**ListEvents**](OperationsAPI.md#ListEvents) | **Get** /dashboard/events | List recent events
[**ListEventsByCursor**](OperationsAPI.md#ListEventsByCursor) | **Get** /dashboard/events/cursor | List recent events by cursor
[**ReplayDeliveries**](OperationsAPI.md#ReplayDeliveries) | **Post** /deliveries/replay | Replay eligible deliveries
[**SendTestPurchase**](OperationsAPI.md#SendTestPurchase) | **Post** /destinations/{destination}/test-purchase | Send a controlled test purchase
[**SetDestinationTestMode**](OperationsAPI.md#SetDestinationTestMode) | **Patch** /destinations/{destination}/test-mode | Configure destination test mode



## DeleteUserData

> DeleteUserData200Response DeleteUserData(ctx).DeleteUserDataRequest(deleteUserDataRequest).Execute()

Delete user data by hashed identifier



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
	deleteUserDataRequest := *openapiclient.NewDeleteUserDataRequest("IdentifierType_example", "IdentifierHash_example") // DeleteUserDataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.DeleteUserData(context.Background()).DeleteUserDataRequest(deleteUserDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.DeleteUserData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserData`: DeleteUserData200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.DeleteUserData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteUserDataRequest** | [**DeleteUserDataRequest**](DeleteUserDataRequest.md) |  | 

### Return type

[**DeleteUserData200Response**](DeleteUserData200Response.md)

### Authorization

[signalTrackerSecret](../README.md#signalTrackerSecret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmqReport

> GetEmqReport200Response GetEmqReport(ctx).Execute()

Get Event Match Quality history



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
	resp, r, err := apiClient.OperationsAPI.GetEmqReport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.GetEmqReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmqReport`: GetEmqReport200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.GetEmqReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmqReportRequest struct via the builder pattern


### Return type

[**GetEmqReport200Response**](GetEmqReport200Response.md)

### Authorization

[signalTrackerSecret](../README.md#signalTrackerSecret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReconciliationReport

> GetReconciliationReport200Response GetReconciliationReport(ctx).Date(date).Execute()

Get a reconciliation report



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
	date := time.Now() // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.GetReconciliationReport(context.Background()).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.GetReconciliationReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReconciliationReport`: GetReconciliationReport200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.GetReconciliationReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReconciliationReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **string** |  | 

### Return type

[**GetReconciliationReport200Response**](GetReconciliationReport200Response.md)

### Authorization

[signalTrackerSecret](../README.md#signalTrackerSecret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvents

> ListEvents200Response ListEvents(ctx).PerPage(perPage).Execute()

List recent events



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
	perPage := int32(56) // int32 | Events per page, capped at 100. (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.ListEvents(context.Background()).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ListEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEvents`: ListEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ListEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Events per page, capped at 100. | [default to 25]

### Return type

[**ListEvents200Response**](ListEvents200Response.md)

### Authorization

[signalTrackerSecret](../README.md#signalTrackerSecret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventsByCursor

> ListEventsByCursor200Response ListEventsByCursor(ctx).PerPage(perPage).Cursor(cursor).Execute()

List recent events by cursor



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
	perPage := int32(56) // int32 | Events per cursor page, capped at 100. (optional) (default to 25)
	cursor := "cursor_example" // string | Opaque cursor returned by a previous response. Omit it on the first request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.ListEventsByCursor(context.Background()).PerPage(perPage).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ListEventsByCursor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEventsByCursor`: ListEventsByCursor200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ListEventsByCursor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventsByCursorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Events per cursor page, capped at 100. | [default to 25]
 **cursor** | **string** | Opaque cursor returned by a previous response. Omit it on the first request. | 

### Return type

[**ListEventsByCursor200Response**](ListEventsByCursor200Response.md)

### Authorization

[signalTrackerSecret](../README.md#signalTrackerSecret)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplayDeliveries

> ReplayDeliveries202Response ReplayDeliveries(ctx).ReplayDeliveriesRequest(replayDeliveriesRequest).Execute()

Replay eligible deliveries



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
	replayDeliveriesRequest := *openapiclient.NewReplayDeliveriesRequest() // ReplayDeliveriesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.ReplayDeliveries(context.Background()).ReplayDeliveriesRequest(replayDeliveriesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.ReplayDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplayDeliveries`: ReplayDeliveries202Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.ReplayDeliveries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplayDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replayDeliveriesRequest** | [**ReplayDeliveriesRequest**](ReplayDeliveriesRequest.md) |  | 

### Return type

[**ReplayDeliveries202Response**](ReplayDeliveries202Response.md)

### Authorization

[signalTrackerSecret](../README.md#signalTrackerSecret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTestPurchase

> SendTestPurchase200Response SendTestPurchase(ctx, destination).SendTestPurchaseRequest(sendTestPurchaseRequest).Execute()

Send a controlled test purchase



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
	destination := "destination_example" // string | The destination ID
	sendTestPurchaseRequest := *openapiclient.NewSendTestPurchaseRequest() // SendTestPurchaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.SendTestPurchase(context.Background(), destination).SendTestPurchaseRequest(sendTestPurchaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.SendTestPurchase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendTestPurchase`: SendTestPurchase200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.SendTestPurchase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destination** | **string** | The destination ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendTestPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendTestPurchaseRequest** | [**SendTestPurchaseRequest**](SendTestPurchaseRequest.md) |  | 

### Return type

[**SendTestPurchase200Response**](SendTestPurchase200Response.md)

### Authorization

[signalTrackerSecret](../README.md#signalTrackerSecret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDestinationTestMode

> SetDestinationTestMode200Response SetDestinationTestMode(ctx, destination).SetDestinationTestModeRequest(setDestinationTestModeRequest).Execute()

Configure destination test mode



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
	destination := "destination_example" // string | The destination ID
	setDestinationTestModeRequest := *openapiclient.NewSetDestinationTestModeRequest(false) // SetDestinationTestModeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperationsAPI.SetDestinationTestMode(context.Background(), destination).SetDestinationTestModeRequest(setDestinationTestModeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperationsAPI.SetDestinationTestMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDestinationTestMode`: SetDestinationTestMode200Response
	fmt.Fprintf(os.Stdout, "Response from `OperationsAPI.SetDestinationTestMode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destination** | **string** | The destination ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDestinationTestModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setDestinationTestModeRequest** | [**SetDestinationTestModeRequest**](SetDestinationTestModeRequest.md) |  | 

### Return type

[**SetDestinationTestMode200Response**](SetDestinationTestMode200Response.md)

### Authorization

[signalTrackerSecret](../README.md#signalTrackerSecret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

