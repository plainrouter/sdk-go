# CreateEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **string** | Caller-supplied idempotency key; maximum 128 characters. | [optional] 
**EventName** | **string** | Signal event name; maximum 100 characters. | 
**ParentEventId** | Pointer to **string** | Optional parent event id; maximum 128 characters. | [optional] 
**EventTime** | Pointer to [**CreateEventRequestAnyOfEventTime**](CreateEventRequestAnyOfEventTime.md) |  | [optional] 
**EventSource** | Pointer to **string** | Optional absolute source URL. | [optional] 
**ActionSource** | Pointer to **string** | Optional action source; defaults to website and is limited to 50 characters. | [optional] 
**VisitorId** | Pointer to **string** | Optional visitor identifier; maximum 255 characters. | [optional] 
**ConsentBasis** | **string** | Legal basis for processing. Legitimate-interest revenue lifecycle events are rejected; use an authenticated server adapter. | 
**Consent** | Pointer to **map[string]interface{}** | Consent state supplied with the event. | [optional] 
**ConsentMode** | Pointer to **map[string]interface{}** | Consent Mode v2 signal values supplied with the event. | [optional] 
**Tcf** | Pointer to **map[string]interface{}** | TCF v2 data containing string and optional captured_at. | [optional] 
**UserData** | Pointer to **map[string]interface{}** | Identity fields accepted by the tracker. | [optional] 
**ClickIds** | Pointer to **map[string]interface{}** | Advertising click identifiers. | [optional] 
**ValueData** | Pointer to [**CreateEventRequestAnyOfValueData**](CreateEventRequestAnyOfValueData.md) |  | [optional] 

## Methods

### NewCreateEventRequest

`func NewCreateEventRequest(eventName string, consentBasis string, ) *CreateEventRequest`

NewCreateEventRequest instantiates a new CreateEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEventRequestWithDefaults

`func NewCreateEventRequestWithDefaults() *CreateEventRequest`

NewCreateEventRequestWithDefaults instantiates a new CreateEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *CreateEventRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *CreateEventRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *CreateEventRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *CreateEventRequest) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventName

`func (o *CreateEventRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *CreateEventRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *CreateEventRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetParentEventId

`func (o *CreateEventRequest) GetParentEventId() string`

GetParentEventId returns the ParentEventId field if non-nil, zero value otherwise.

### GetParentEventIdOk

`func (o *CreateEventRequest) GetParentEventIdOk() (*string, bool)`

GetParentEventIdOk returns a tuple with the ParentEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEventId

`func (o *CreateEventRequest) SetParentEventId(v string)`

SetParentEventId sets ParentEventId field to given value.

### HasParentEventId

`func (o *CreateEventRequest) HasParentEventId() bool`

HasParentEventId returns a boolean if a field has been set.

### GetEventTime

`func (o *CreateEventRequest) GetEventTime() CreateEventRequestAnyOfEventTime`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *CreateEventRequest) GetEventTimeOk() (*CreateEventRequestAnyOfEventTime, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *CreateEventRequest) SetEventTime(v CreateEventRequestAnyOfEventTime)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *CreateEventRequest) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventSource

`func (o *CreateEventRequest) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *CreateEventRequest) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *CreateEventRequest) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *CreateEventRequest) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### GetActionSource

`func (o *CreateEventRequest) GetActionSource() string`

GetActionSource returns the ActionSource field if non-nil, zero value otherwise.

### GetActionSourceOk

`func (o *CreateEventRequest) GetActionSourceOk() (*string, bool)`

GetActionSourceOk returns a tuple with the ActionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionSource

`func (o *CreateEventRequest) SetActionSource(v string)`

SetActionSource sets ActionSource field to given value.

### HasActionSource

`func (o *CreateEventRequest) HasActionSource() bool`

HasActionSource returns a boolean if a field has been set.

### GetVisitorId

`func (o *CreateEventRequest) GetVisitorId() string`

GetVisitorId returns the VisitorId field if non-nil, zero value otherwise.

### GetVisitorIdOk

`func (o *CreateEventRequest) GetVisitorIdOk() (*string, bool)`

GetVisitorIdOk returns a tuple with the VisitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitorId

`func (o *CreateEventRequest) SetVisitorId(v string)`

SetVisitorId sets VisitorId field to given value.

### HasVisitorId

`func (o *CreateEventRequest) HasVisitorId() bool`

HasVisitorId returns a boolean if a field has been set.

### GetConsentBasis

`func (o *CreateEventRequest) GetConsentBasis() string`

GetConsentBasis returns the ConsentBasis field if non-nil, zero value otherwise.

### GetConsentBasisOk

`func (o *CreateEventRequest) GetConsentBasisOk() (*string, bool)`

GetConsentBasisOk returns a tuple with the ConsentBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentBasis

`func (o *CreateEventRequest) SetConsentBasis(v string)`

SetConsentBasis sets ConsentBasis field to given value.


### GetConsent

`func (o *CreateEventRequest) GetConsent() map[string]interface{}`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *CreateEventRequest) GetConsentOk() (*map[string]interface{}, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *CreateEventRequest) SetConsent(v map[string]interface{})`

SetConsent sets Consent field to given value.

### HasConsent

`func (o *CreateEventRequest) HasConsent() bool`

HasConsent returns a boolean if a field has been set.

### GetConsentMode

`func (o *CreateEventRequest) GetConsentMode() map[string]interface{}`

GetConsentMode returns the ConsentMode field if non-nil, zero value otherwise.

### GetConsentModeOk

`func (o *CreateEventRequest) GetConsentModeOk() (*map[string]interface{}, bool)`

GetConsentModeOk returns a tuple with the ConsentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentMode

`func (o *CreateEventRequest) SetConsentMode(v map[string]interface{})`

SetConsentMode sets ConsentMode field to given value.

### HasConsentMode

`func (o *CreateEventRequest) HasConsentMode() bool`

HasConsentMode returns a boolean if a field has been set.

### GetTcf

`func (o *CreateEventRequest) GetTcf() map[string]interface{}`

GetTcf returns the Tcf field if non-nil, zero value otherwise.

### GetTcfOk

`func (o *CreateEventRequest) GetTcfOk() (*map[string]interface{}, bool)`

GetTcfOk returns a tuple with the Tcf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcf

`func (o *CreateEventRequest) SetTcf(v map[string]interface{})`

SetTcf sets Tcf field to given value.

### HasTcf

`func (o *CreateEventRequest) HasTcf() bool`

HasTcf returns a boolean if a field has been set.

### GetUserData

`func (o *CreateEventRequest) GetUserData() map[string]interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreateEventRequest) GetUserDataOk() (*map[string]interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreateEventRequest) SetUserData(v map[string]interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *CreateEventRequest) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetClickIds

`func (o *CreateEventRequest) GetClickIds() map[string]interface{}`

GetClickIds returns the ClickIds field if non-nil, zero value otherwise.

### GetClickIdsOk

`func (o *CreateEventRequest) GetClickIdsOk() (*map[string]interface{}, bool)`

GetClickIdsOk returns a tuple with the ClickIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickIds

`func (o *CreateEventRequest) SetClickIds(v map[string]interface{})`

SetClickIds sets ClickIds field to given value.

### HasClickIds

`func (o *CreateEventRequest) HasClickIds() bool`

HasClickIds returns a boolean if a field has been set.

### GetValueData

`func (o *CreateEventRequest) GetValueData() CreateEventRequestAnyOfValueData`

GetValueData returns the ValueData field if non-nil, zero value otherwise.

### GetValueDataOk

`func (o *CreateEventRequest) GetValueDataOk() (*CreateEventRequestAnyOfValueData, bool)`

GetValueDataOk returns a tuple with the ValueData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueData

`func (o *CreateEventRequest) SetValueData(v CreateEventRequestAnyOfValueData)`

SetValueData sets ValueData field to given value.

### HasValueData

`func (o *CreateEventRequest) HasValueData() bool`

HasValueData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


