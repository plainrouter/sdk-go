# CreateEventRequestAnyOf1

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

### NewCreateEventRequestAnyOf1

`func NewCreateEventRequestAnyOf1(eventName string, consentBasis string, ) *CreateEventRequestAnyOf1`

NewCreateEventRequestAnyOf1 instantiates a new CreateEventRequestAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEventRequestAnyOf1WithDefaults

`func NewCreateEventRequestAnyOf1WithDefaults() *CreateEventRequestAnyOf1`

NewCreateEventRequestAnyOf1WithDefaults instantiates a new CreateEventRequestAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *CreateEventRequestAnyOf1) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *CreateEventRequestAnyOf1) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *CreateEventRequestAnyOf1) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *CreateEventRequestAnyOf1) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventName

`func (o *CreateEventRequestAnyOf1) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *CreateEventRequestAnyOf1) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *CreateEventRequestAnyOf1) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetParentEventId

`func (o *CreateEventRequestAnyOf1) GetParentEventId() string`

GetParentEventId returns the ParentEventId field if non-nil, zero value otherwise.

### GetParentEventIdOk

`func (o *CreateEventRequestAnyOf1) GetParentEventIdOk() (*string, bool)`

GetParentEventIdOk returns a tuple with the ParentEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEventId

`func (o *CreateEventRequestAnyOf1) SetParentEventId(v string)`

SetParentEventId sets ParentEventId field to given value.

### HasParentEventId

`func (o *CreateEventRequestAnyOf1) HasParentEventId() bool`

HasParentEventId returns a boolean if a field has been set.

### GetEventTime

`func (o *CreateEventRequestAnyOf1) GetEventTime() CreateEventRequestAnyOfEventTime`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *CreateEventRequestAnyOf1) GetEventTimeOk() (*CreateEventRequestAnyOfEventTime, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *CreateEventRequestAnyOf1) SetEventTime(v CreateEventRequestAnyOfEventTime)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *CreateEventRequestAnyOf1) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventSource

`func (o *CreateEventRequestAnyOf1) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *CreateEventRequestAnyOf1) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *CreateEventRequestAnyOf1) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *CreateEventRequestAnyOf1) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### GetActionSource

`func (o *CreateEventRequestAnyOf1) GetActionSource() string`

GetActionSource returns the ActionSource field if non-nil, zero value otherwise.

### GetActionSourceOk

`func (o *CreateEventRequestAnyOf1) GetActionSourceOk() (*string, bool)`

GetActionSourceOk returns a tuple with the ActionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionSource

`func (o *CreateEventRequestAnyOf1) SetActionSource(v string)`

SetActionSource sets ActionSource field to given value.

### HasActionSource

`func (o *CreateEventRequestAnyOf1) HasActionSource() bool`

HasActionSource returns a boolean if a field has been set.

### GetVisitorId

`func (o *CreateEventRequestAnyOf1) GetVisitorId() string`

GetVisitorId returns the VisitorId field if non-nil, zero value otherwise.

### GetVisitorIdOk

`func (o *CreateEventRequestAnyOf1) GetVisitorIdOk() (*string, bool)`

GetVisitorIdOk returns a tuple with the VisitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitorId

`func (o *CreateEventRequestAnyOf1) SetVisitorId(v string)`

SetVisitorId sets VisitorId field to given value.

### HasVisitorId

`func (o *CreateEventRequestAnyOf1) HasVisitorId() bool`

HasVisitorId returns a boolean if a field has been set.

### GetConsentBasis

`func (o *CreateEventRequestAnyOf1) GetConsentBasis() string`

GetConsentBasis returns the ConsentBasis field if non-nil, zero value otherwise.

### GetConsentBasisOk

`func (o *CreateEventRequestAnyOf1) GetConsentBasisOk() (*string, bool)`

GetConsentBasisOk returns a tuple with the ConsentBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentBasis

`func (o *CreateEventRequestAnyOf1) SetConsentBasis(v string)`

SetConsentBasis sets ConsentBasis field to given value.


### GetConsent

`func (o *CreateEventRequestAnyOf1) GetConsent() map[string]*interface{}`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *CreateEventRequestAnyOf1) GetConsentOk() (*map[string]*interface{}, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *CreateEventRequestAnyOf1) SetConsent(v map[string]*interface{})`

SetConsent sets Consent field to given value.

### HasConsent

`func (o *CreateEventRequestAnyOf1) HasConsent() bool`

HasConsent returns a boolean if a field has been set.

### GetConsentMode

`func (o *CreateEventRequestAnyOf1) GetConsentMode() map[string]*interface{}`

GetConsentMode returns the ConsentMode field if non-nil, zero value otherwise.

### GetConsentModeOk

`func (o *CreateEventRequestAnyOf1) GetConsentModeOk() (*map[string]*interface{}, bool)`

GetConsentModeOk returns a tuple with the ConsentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentMode

`func (o *CreateEventRequestAnyOf1) SetConsentMode(v map[string]*interface{})`

SetConsentMode sets ConsentMode field to given value.

### HasConsentMode

`func (o *CreateEventRequestAnyOf1) HasConsentMode() bool`

HasConsentMode returns a boolean if a field has been set.

### GetTcf

`func (o *CreateEventRequestAnyOf1) GetTcf() map[string]*interface{}`

GetTcf returns the Tcf field if non-nil, zero value otherwise.

### GetTcfOk

`func (o *CreateEventRequestAnyOf1) GetTcfOk() (*map[string]*interface{}, bool)`

GetTcfOk returns a tuple with the Tcf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcf

`func (o *CreateEventRequestAnyOf1) SetTcf(v map[string]*interface{})`

SetTcf sets Tcf field to given value.

### HasTcf

`func (o *CreateEventRequestAnyOf1) HasTcf() bool`

HasTcf returns a boolean if a field has been set.

### GetUserData

`func (o *CreateEventRequestAnyOf1) GetUserData() map[string]*interface{}`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreateEventRequestAnyOf1) GetUserDataOk() (*map[string]*interface{}, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreateEventRequestAnyOf1) SetUserData(v map[string]*interface{})`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *CreateEventRequestAnyOf1) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetClickIds

`func (o *CreateEventRequestAnyOf1) GetClickIds() map[string]*interface{}`

GetClickIds returns the ClickIds field if non-nil, zero value otherwise.

### GetClickIdsOk

`func (o *CreateEventRequestAnyOf1) GetClickIdsOk() (*map[string]*interface{}, bool)`

GetClickIdsOk returns a tuple with the ClickIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickIds

`func (o *CreateEventRequestAnyOf1) SetClickIds(v map[string]*interface{})`

SetClickIds sets ClickIds field to given value.

### HasClickIds

`func (o *CreateEventRequestAnyOf1) HasClickIds() bool`

HasClickIds returns a boolean if a field has been set.

### GetValueData

`func (o *CreateEventRequestAnyOf1) GetValueData() CreateEventRequestAnyOfValueData`

GetValueData returns the ValueData field if non-nil, zero value otherwise.

### GetValueDataOk

`func (o *CreateEventRequestAnyOf1) GetValueDataOk() (*CreateEventRequestAnyOfValueData, bool)`

GetValueDataOk returns a tuple with the ValueData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueData

`func (o *CreateEventRequestAnyOf1) SetValueData(v CreateEventRequestAnyOfValueData)`

SetValueData sets ValueData field to given value.

### HasValueData

`func (o *CreateEventRequestAnyOf1) HasValueData() bool`

HasValueData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


