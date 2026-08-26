# ValidateSandboxEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** | Optional synthetic idempotency key; maximum 128 characters. | [optional] 
**EventName** | **string** | Synthetic event name; maximum 100 characters. | 
**ActionSource** | Pointer to **string** | Synthetic action source. | [optional] 
**ValueData** | Pointer to [**ValidateSandboxEventRequestValueData**](ValidateSandboxEventRequestValueData.md) |  | [optional] 

## Methods

### NewValidateSandboxEventRequest

`func NewValidateSandboxEventRequest(eventName string, ) *ValidateSandboxEventRequest`

NewValidateSandboxEventRequest instantiates a new ValidateSandboxEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateSandboxEventRequestWithDefaults

`func NewValidateSandboxEventRequestWithDefaults() *ValidateSandboxEventRequest`

NewValidateSandboxEventRequestWithDefaults instantiates a new ValidateSandboxEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *ValidateSandboxEventRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *ValidateSandboxEventRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *ValidateSandboxEventRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *ValidateSandboxEventRequest) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventName

`func (o *ValidateSandboxEventRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *ValidateSandboxEventRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *ValidateSandboxEventRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetActionSource

`func (o *ValidateSandboxEventRequest) GetActionSource() string`

GetActionSource returns the ActionSource field if non-nil, zero value otherwise.

### GetActionSourceOk

`func (o *ValidateSandboxEventRequest) GetActionSourceOk() (*string, bool)`

GetActionSourceOk returns a tuple with the ActionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionSource

`func (o *ValidateSandboxEventRequest) SetActionSource(v string)`

SetActionSource sets ActionSource field to given value.

### HasActionSource

`func (o *ValidateSandboxEventRequest) HasActionSource() bool`

HasActionSource returns a boolean if a field has been set.

### GetValueData

`func (o *ValidateSandboxEventRequest) GetValueData() ValidateSandboxEventRequestValueData`

GetValueData returns the ValueData field if non-nil, zero value otherwise.

### GetValueDataOk

`func (o *ValidateSandboxEventRequest) GetValueDataOk() (*ValidateSandboxEventRequestValueData, bool)`

GetValueDataOk returns a tuple with the ValueData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueData

`func (o *ValidateSandboxEventRequest) SetValueData(v ValidateSandboxEventRequestValueData)`

SetValueData sets ValueData field to given value.

### HasValueData

`func (o *ValidateSandboxEventRequest) HasValueData() bool`

HasValueData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


