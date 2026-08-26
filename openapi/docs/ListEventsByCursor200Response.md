# ListEventsByCursor200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**ListEventsByCursor200ResponseEvents**](ListEventsByCursor200ResponseEvents.md) |  | 
**Metrics** | [**ListEvents200ResponseMetrics**](ListEvents200ResponseMetrics.md) |  | 

## Methods

### NewListEventsByCursor200Response

`func NewListEventsByCursor200Response(events ListEventsByCursor200ResponseEvents, metrics ListEvents200ResponseMetrics, ) *ListEventsByCursor200Response`

NewListEventsByCursor200Response instantiates a new ListEventsByCursor200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEventsByCursor200ResponseWithDefaults

`func NewListEventsByCursor200ResponseWithDefaults() *ListEventsByCursor200Response`

NewListEventsByCursor200ResponseWithDefaults instantiates a new ListEventsByCursor200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *ListEventsByCursor200Response) GetEvents() ListEventsByCursor200ResponseEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListEventsByCursor200Response) GetEventsOk() (*ListEventsByCursor200ResponseEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListEventsByCursor200Response) SetEvents(v ListEventsByCursor200ResponseEvents)`

SetEvents sets Events field to given value.


### GetMetrics

`func (o *ListEventsByCursor200Response) GetMetrics() ListEvents200ResponseMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ListEventsByCursor200Response) GetMetricsOk() (*ListEvents200ResponseMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ListEventsByCursor200Response) SetMetrics(v ListEvents200ResponseMetrics)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


