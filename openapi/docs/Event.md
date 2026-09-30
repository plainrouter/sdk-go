# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ParentEventId** | **NullableString** |  | 
**EventName** | **string** |  | 
**EventTime** | **time.Time** |  | 
**ActionSource** | **string** |  | 
**EventClass** | **string** |  | 
**OrderId** | **NullableString** |  | 
**ValueAmount** | **NullableInt32** |  | 
**ValueCurrency** | **NullableString** |  | 
**CreatedAt** | **time.Time** |  | 
**ConsentBasis** | **string** |  | 
**MeasurementClass** | **string** |  | 
**AttributionJoin** | **string** |  | 
**EnforcementScope** | **string** |  | 
**ConsentNormalizationVersion** | **string** |  | 
**PolicyClass** | [**JurisdictionPolicyClass**](JurisdictionPolicyClass.md) |  | 
**TrafficClass** | [**TrafficClass**](TrafficClass.md) |  | 
**WorkspaceId** | **int32** |  | 
**Consent** | **string** |  | 
**UserDataHashed** | **interface{}** | Deprecated compatibility field. The value is always null; delivery identity is never returned. | 
**ClickIds** | **string** |  | 
**Session** | **string** |  | 
**ValueData** | **string** |  | 
**EventSource** | **string** |  | 
**PayloadExpired** | **bool** |  | 
**ConsentSource** | Pointer to **string** |  | [optional] 
**ConsentUiVersion** | Pointer to **int32** |  | [optional] 
**Deliveries** | **[]interface{}** |  | 

## Methods

### NewEvent

`func NewEvent(id string, parentEventId NullableString, eventName string, eventTime time.Time, actionSource string, eventClass string, orderId NullableString, valueAmount NullableInt32, valueCurrency NullableString, createdAt time.Time, consentBasis string, measurementClass string, attributionJoin string, enforcementScope string, consentNormalizationVersion string, policyClass JurisdictionPolicyClass, trafficClass TrafficClass, workspaceId int32, consent string, userDataHashed interface{}, clickIds string, session string, valueData string, eventSource string, payloadExpired bool, deliveries []interface{}, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Event) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v string)`

SetId sets Id field to given value.


### GetParentEventId

`func (o *Event) GetParentEventId() string`

GetParentEventId returns the ParentEventId field if non-nil, zero value otherwise.

### GetParentEventIdOk

`func (o *Event) GetParentEventIdOk() (*string, bool)`

GetParentEventIdOk returns a tuple with the ParentEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEventId

`func (o *Event) SetParentEventId(v string)`

SetParentEventId sets ParentEventId field to given value.


### SetParentEventIdNil

`func (o *Event) SetParentEventIdNil(b bool)`

 SetParentEventIdNil sets the value for ParentEventId to be an explicit nil

### UnsetParentEventId
`func (o *Event) UnsetParentEventId()`

UnsetParentEventId ensures that no value is present for ParentEventId, not even an explicit nil
### GetEventName

`func (o *Event) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *Event) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *Event) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetEventTime

`func (o *Event) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *Event) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *Event) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetActionSource

`func (o *Event) GetActionSource() string`

GetActionSource returns the ActionSource field if non-nil, zero value otherwise.

### GetActionSourceOk

`func (o *Event) GetActionSourceOk() (*string, bool)`

GetActionSourceOk returns a tuple with the ActionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionSource

`func (o *Event) SetActionSource(v string)`

SetActionSource sets ActionSource field to given value.


### GetEventClass

`func (o *Event) GetEventClass() string`

GetEventClass returns the EventClass field if non-nil, zero value otherwise.

### GetEventClassOk

`func (o *Event) GetEventClassOk() (*string, bool)`

GetEventClassOk returns a tuple with the EventClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClass

`func (o *Event) SetEventClass(v string)`

SetEventClass sets EventClass field to given value.


### GetOrderId

`func (o *Event) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Event) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Event) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### SetOrderIdNil

`func (o *Event) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *Event) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetValueAmount

`func (o *Event) GetValueAmount() int32`

GetValueAmount returns the ValueAmount field if non-nil, zero value otherwise.

### GetValueAmountOk

`func (o *Event) GetValueAmountOk() (*int32, bool)`

GetValueAmountOk returns a tuple with the ValueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAmount

`func (o *Event) SetValueAmount(v int32)`

SetValueAmount sets ValueAmount field to given value.


### SetValueAmountNil

`func (o *Event) SetValueAmountNil(b bool)`

 SetValueAmountNil sets the value for ValueAmount to be an explicit nil

### UnsetValueAmount
`func (o *Event) UnsetValueAmount()`

UnsetValueAmount ensures that no value is present for ValueAmount, not even an explicit nil
### GetValueCurrency

`func (o *Event) GetValueCurrency() string`

GetValueCurrency returns the ValueCurrency field if non-nil, zero value otherwise.

### GetValueCurrencyOk

`func (o *Event) GetValueCurrencyOk() (*string, bool)`

GetValueCurrencyOk returns a tuple with the ValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCurrency

`func (o *Event) SetValueCurrency(v string)`

SetValueCurrency sets ValueCurrency field to given value.


### SetValueCurrencyNil

`func (o *Event) SetValueCurrencyNil(b bool)`

 SetValueCurrencyNil sets the value for ValueCurrency to be an explicit nil

### UnsetValueCurrency
`func (o *Event) UnsetValueCurrency()`

UnsetValueCurrency ensures that no value is present for ValueCurrency, not even an explicit nil
### GetCreatedAt

`func (o *Event) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Event) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Event) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetConsentBasis

`func (o *Event) GetConsentBasis() string`

GetConsentBasis returns the ConsentBasis field if non-nil, zero value otherwise.

### GetConsentBasisOk

`func (o *Event) GetConsentBasisOk() (*string, bool)`

GetConsentBasisOk returns a tuple with the ConsentBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentBasis

`func (o *Event) SetConsentBasis(v string)`

SetConsentBasis sets ConsentBasis field to given value.


### GetMeasurementClass

`func (o *Event) GetMeasurementClass() string`

GetMeasurementClass returns the MeasurementClass field if non-nil, zero value otherwise.

### GetMeasurementClassOk

`func (o *Event) GetMeasurementClassOk() (*string, bool)`

GetMeasurementClassOk returns a tuple with the MeasurementClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementClass

`func (o *Event) SetMeasurementClass(v string)`

SetMeasurementClass sets MeasurementClass field to given value.


### GetAttributionJoin

`func (o *Event) GetAttributionJoin() string`

GetAttributionJoin returns the AttributionJoin field if non-nil, zero value otherwise.

### GetAttributionJoinOk

`func (o *Event) GetAttributionJoinOk() (*string, bool)`

GetAttributionJoinOk returns a tuple with the AttributionJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionJoin

`func (o *Event) SetAttributionJoin(v string)`

SetAttributionJoin sets AttributionJoin field to given value.


### GetEnforcementScope

`func (o *Event) GetEnforcementScope() string`

GetEnforcementScope returns the EnforcementScope field if non-nil, zero value otherwise.

### GetEnforcementScopeOk

`func (o *Event) GetEnforcementScopeOk() (*string, bool)`

GetEnforcementScopeOk returns a tuple with the EnforcementScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementScope

`func (o *Event) SetEnforcementScope(v string)`

SetEnforcementScope sets EnforcementScope field to given value.


### GetConsentNormalizationVersion

`func (o *Event) GetConsentNormalizationVersion() string`

GetConsentNormalizationVersion returns the ConsentNormalizationVersion field if non-nil, zero value otherwise.

### GetConsentNormalizationVersionOk

`func (o *Event) GetConsentNormalizationVersionOk() (*string, bool)`

GetConsentNormalizationVersionOk returns a tuple with the ConsentNormalizationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentNormalizationVersion

`func (o *Event) SetConsentNormalizationVersion(v string)`

SetConsentNormalizationVersion sets ConsentNormalizationVersion field to given value.


### GetPolicyClass

`func (o *Event) GetPolicyClass() JurisdictionPolicyClass`

GetPolicyClass returns the PolicyClass field if non-nil, zero value otherwise.

### GetPolicyClassOk

`func (o *Event) GetPolicyClassOk() (*JurisdictionPolicyClass, bool)`

GetPolicyClassOk returns a tuple with the PolicyClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyClass

`func (o *Event) SetPolicyClass(v JurisdictionPolicyClass)`

SetPolicyClass sets PolicyClass field to given value.


### GetTrafficClass

`func (o *Event) GetTrafficClass() TrafficClass`

GetTrafficClass returns the TrafficClass field if non-nil, zero value otherwise.

### GetTrafficClassOk

`func (o *Event) GetTrafficClassOk() (*TrafficClass, bool)`

GetTrafficClassOk returns a tuple with the TrafficClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficClass

`func (o *Event) SetTrafficClass(v TrafficClass)`

SetTrafficClass sets TrafficClass field to given value.


### GetWorkspaceId

`func (o *Event) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *Event) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *Event) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetConsent

`func (o *Event) GetConsent() string`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *Event) GetConsentOk() (*string, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *Event) SetConsent(v string)`

SetConsent sets Consent field to given value.


### GetUserDataHashed

`func (o *Event) GetUserDataHashed() interface{}`

GetUserDataHashed returns the UserDataHashed field if non-nil, zero value otherwise.

### GetUserDataHashedOk

`func (o *Event) GetUserDataHashedOk() (*interface{}, bool)`

GetUserDataHashedOk returns a tuple with the UserDataHashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataHashed

`func (o *Event) SetUserDataHashed(v interface{})`

SetUserDataHashed sets UserDataHashed field to given value.


### SetUserDataHashedNil

`func (o *Event) SetUserDataHashedNil(b bool)`

 SetUserDataHashedNil sets the value for UserDataHashed to be an explicit nil

### UnsetUserDataHashed
`func (o *Event) UnsetUserDataHashed()`

UnsetUserDataHashed ensures that no value is present for UserDataHashed, not even an explicit nil
### GetClickIds

`func (o *Event) GetClickIds() string`

GetClickIds returns the ClickIds field if non-nil, zero value otherwise.

### GetClickIdsOk

`func (o *Event) GetClickIdsOk() (*string, bool)`

GetClickIdsOk returns a tuple with the ClickIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickIds

`func (o *Event) SetClickIds(v string)`

SetClickIds sets ClickIds field to given value.


### GetSession

`func (o *Event) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *Event) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *Event) SetSession(v string)`

SetSession sets Session field to given value.


### GetValueData

`func (o *Event) GetValueData() string`

GetValueData returns the ValueData field if non-nil, zero value otherwise.

### GetValueDataOk

`func (o *Event) GetValueDataOk() (*string, bool)`

GetValueDataOk returns a tuple with the ValueData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueData

`func (o *Event) SetValueData(v string)`

SetValueData sets ValueData field to given value.


### GetEventSource

`func (o *Event) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *Event) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *Event) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.


### GetPayloadExpired

`func (o *Event) GetPayloadExpired() bool`

GetPayloadExpired returns the PayloadExpired field if non-nil, zero value otherwise.

### GetPayloadExpiredOk

`func (o *Event) GetPayloadExpiredOk() (*bool, bool)`

GetPayloadExpiredOk returns a tuple with the PayloadExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadExpired

`func (o *Event) SetPayloadExpired(v bool)`

SetPayloadExpired sets PayloadExpired field to given value.


### GetConsentSource

`func (o *Event) GetConsentSource() string`

GetConsentSource returns the ConsentSource field if non-nil, zero value otherwise.

### GetConsentSourceOk

`func (o *Event) GetConsentSourceOk() (*string, bool)`

GetConsentSourceOk returns a tuple with the ConsentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentSource

`func (o *Event) SetConsentSource(v string)`

SetConsentSource sets ConsentSource field to given value.

### HasConsentSource

`func (o *Event) HasConsentSource() bool`

HasConsentSource returns a boolean if a field has been set.

### GetConsentUiVersion

`func (o *Event) GetConsentUiVersion() int32`

GetConsentUiVersion returns the ConsentUiVersion field if non-nil, zero value otherwise.

### GetConsentUiVersionOk

`func (o *Event) GetConsentUiVersionOk() (*int32, bool)`

GetConsentUiVersionOk returns a tuple with the ConsentUiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentUiVersion

`func (o *Event) SetConsentUiVersion(v int32)`

SetConsentUiVersion sets ConsentUiVersion field to given value.

### HasConsentUiVersion

`func (o *Event) HasConsentUiVersion() bool`

HasConsentUiVersion returns a boolean if a field has been set.

### GetDeliveries

`func (o *Event) GetDeliveries() []interface{}`

GetDeliveries returns the Deliveries field if non-nil, zero value otherwise.

### GetDeliveriesOk

`func (o *Event) GetDeliveriesOk() (*[]interface{}, bool)`

GetDeliveriesOk returns a tuple with the Deliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveries

`func (o *Event) SetDeliveries(v []interface{})`

SetDeliveries sets Deliveries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


