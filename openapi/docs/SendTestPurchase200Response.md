# SendTestPurchase200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**Status** | [**DeliveryStatus**](DeliveryStatus.md) |  | 
**TraceId** | **NullableString** |  | 

## Methods

### NewSendTestPurchase200Response

`func NewSendTestPurchase200Response(eventId string, status DeliveryStatus, traceId NullableString, ) *SendTestPurchase200Response`

NewSendTestPurchase200Response instantiates a new SendTestPurchase200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendTestPurchase200ResponseWithDefaults

`func NewSendTestPurchase200ResponseWithDefaults() *SendTestPurchase200Response`

NewSendTestPurchase200ResponseWithDefaults instantiates a new SendTestPurchase200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *SendTestPurchase200Response) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SendTestPurchase200Response) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SendTestPurchase200Response) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetStatus

`func (o *SendTestPurchase200Response) GetStatus() DeliveryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SendTestPurchase200Response) GetStatusOk() (*DeliveryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SendTestPurchase200Response) SetStatus(v DeliveryStatus)`

SetStatus sets Status field to given value.


### GetTraceId

`func (o *SendTestPurchase200Response) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *SendTestPurchase200Response) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *SendTestPurchase200Response) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### SetTraceIdNil

`func (o *SendTestPurchase200Response) SetTraceIdNil(b bool)`

 SetTraceIdNil sets the value for TraceId to be an explicit nil

### UnsetTraceId
`func (o *SendTestPurchase200Response) UnsetTraceId()`

UnsetTraceId ensures that no value is present for TraceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


