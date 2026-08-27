# GetEvent200ResponseEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SignalTrackerId** | **string** |  | 
**ParentEventId** | **NullableString** |  | 
**EventName** | **string** |  | 
**EventTime** | **string** |  | 
**ActionSource** | **string** |  | 
**EventClass** | **string** |  | 
**OrderId** | **NullableString** |  | 
**ValueAmount** | **NullableInt32** |  | 
**ValueCurrency** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**ConsentBasis** | **string** |  | 
**MeasurementClass** | **string** |  | 
**AttributionJoin** | **string** |  | 
**EnforcementScope** | **string** |  | 
**ConsentNormalizationVersion** | **string** |  | 
**Consent** | **interface{}** |  | 
**UserDataHashed** | **interface{}** |  | 
**ClickIds** | **interface{}** |  | 
**Session** | **interface{}** |  | 
**ValueData** | **interface{}** |  | 
**EventSource** | **NullableString** |  | 
**PayloadExpired** | **bool** |  | 
**Deliveries** | [**[]GetEvent200ResponseEventDeliveriesInner**](GetEvent200ResponseEventDeliveriesInner.md) |  | 

## Methods

### NewGetEvent200ResponseEvent

`func NewGetEvent200ResponseEvent(id string, signalTrackerId string, parentEventId NullableString, eventName string, eventTime string, actionSource string, eventClass string, orderId NullableString, valueAmount NullableInt32, valueCurrency NullableString, createdAt string, consentBasis string, measurementClass string, attributionJoin string, enforcementScope string, consentNormalizationVersion string, consent interface{}, userDataHashed interface{}, clickIds interface{}, session interface{}, valueData interface{}, eventSource NullableString, payloadExpired bool, deliveries []GetEvent200ResponseEventDeliveriesInner, ) *GetEvent200ResponseEvent`

NewGetEvent200ResponseEvent instantiates a new GetEvent200ResponseEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvent200ResponseEventWithDefaults

`func NewGetEvent200ResponseEventWithDefaults() *GetEvent200ResponseEvent`

NewGetEvent200ResponseEventWithDefaults instantiates a new GetEvent200ResponseEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEvent200ResponseEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEvent200ResponseEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEvent200ResponseEvent) SetId(v string)`

SetId sets Id field to given value.


### GetSignalTrackerId

`func (o *GetEvent200ResponseEvent) GetSignalTrackerId() string`

GetSignalTrackerId returns the SignalTrackerId field if non-nil, zero value otherwise.

### GetSignalTrackerIdOk

`func (o *GetEvent200ResponseEvent) GetSignalTrackerIdOk() (*string, bool)`

GetSignalTrackerIdOk returns a tuple with the SignalTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalTrackerId

`func (o *GetEvent200ResponseEvent) SetSignalTrackerId(v string)`

SetSignalTrackerId sets SignalTrackerId field to given value.


### GetParentEventId

`func (o *GetEvent200ResponseEvent) GetParentEventId() string`

GetParentEventId returns the ParentEventId field if non-nil, zero value otherwise.

### GetParentEventIdOk

`func (o *GetEvent200ResponseEvent) GetParentEventIdOk() (*string, bool)`

GetParentEventIdOk returns a tuple with the ParentEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEventId

`func (o *GetEvent200ResponseEvent) SetParentEventId(v string)`

SetParentEventId sets ParentEventId field to given value.


### SetParentEventIdNil

`func (o *GetEvent200ResponseEvent) SetParentEventIdNil(b bool)`

 SetParentEventIdNil sets the value for ParentEventId to be an explicit nil

### UnsetParentEventId
`func (o *GetEvent200ResponseEvent) UnsetParentEventId()`

UnsetParentEventId ensures that no value is present for ParentEventId, not even an explicit nil
### GetEventName

`func (o *GetEvent200ResponseEvent) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *GetEvent200ResponseEvent) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *GetEvent200ResponseEvent) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetEventTime

`func (o *GetEvent200ResponseEvent) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *GetEvent200ResponseEvent) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *GetEvent200ResponseEvent) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.


### GetActionSource

`func (o *GetEvent200ResponseEvent) GetActionSource() string`

GetActionSource returns the ActionSource field if non-nil, zero value otherwise.

### GetActionSourceOk

`func (o *GetEvent200ResponseEvent) GetActionSourceOk() (*string, bool)`

GetActionSourceOk returns a tuple with the ActionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionSource

`func (o *GetEvent200ResponseEvent) SetActionSource(v string)`

SetActionSource sets ActionSource field to given value.


### GetEventClass

`func (o *GetEvent200ResponseEvent) GetEventClass() string`

GetEventClass returns the EventClass field if non-nil, zero value otherwise.

### GetEventClassOk

`func (o *GetEvent200ResponseEvent) GetEventClassOk() (*string, bool)`

GetEventClassOk returns a tuple with the EventClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClass

`func (o *GetEvent200ResponseEvent) SetEventClass(v string)`

SetEventClass sets EventClass field to given value.


### GetOrderId

`func (o *GetEvent200ResponseEvent) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetEvent200ResponseEvent) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetEvent200ResponseEvent) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### SetOrderIdNil

`func (o *GetEvent200ResponseEvent) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *GetEvent200ResponseEvent) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetValueAmount

`func (o *GetEvent200ResponseEvent) GetValueAmount() int32`

GetValueAmount returns the ValueAmount field if non-nil, zero value otherwise.

### GetValueAmountOk

`func (o *GetEvent200ResponseEvent) GetValueAmountOk() (*int32, bool)`

GetValueAmountOk returns a tuple with the ValueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAmount

`func (o *GetEvent200ResponseEvent) SetValueAmount(v int32)`

SetValueAmount sets ValueAmount field to given value.


### SetValueAmountNil

`func (o *GetEvent200ResponseEvent) SetValueAmountNil(b bool)`

 SetValueAmountNil sets the value for ValueAmount to be an explicit nil

### UnsetValueAmount
`func (o *GetEvent200ResponseEvent) UnsetValueAmount()`

UnsetValueAmount ensures that no value is present for ValueAmount, not even an explicit nil
### GetValueCurrency

`func (o *GetEvent200ResponseEvent) GetValueCurrency() string`

GetValueCurrency returns the ValueCurrency field if non-nil, zero value otherwise.

### GetValueCurrencyOk

`func (o *GetEvent200ResponseEvent) GetValueCurrencyOk() (*string, bool)`

GetValueCurrencyOk returns a tuple with the ValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCurrency

`func (o *GetEvent200ResponseEvent) SetValueCurrency(v string)`

SetValueCurrency sets ValueCurrency field to given value.


### SetValueCurrencyNil

`func (o *GetEvent200ResponseEvent) SetValueCurrencyNil(b bool)`

 SetValueCurrencyNil sets the value for ValueCurrency to be an explicit nil

### UnsetValueCurrency
`func (o *GetEvent200ResponseEvent) UnsetValueCurrency()`

UnsetValueCurrency ensures that no value is present for ValueCurrency, not even an explicit nil
### GetCreatedAt

`func (o *GetEvent200ResponseEvent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetEvent200ResponseEvent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetEvent200ResponseEvent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetConsentBasis

`func (o *GetEvent200ResponseEvent) GetConsentBasis() string`

GetConsentBasis returns the ConsentBasis field if non-nil, zero value otherwise.

### GetConsentBasisOk

`func (o *GetEvent200ResponseEvent) GetConsentBasisOk() (*string, bool)`

GetConsentBasisOk returns a tuple with the ConsentBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentBasis

`func (o *GetEvent200ResponseEvent) SetConsentBasis(v string)`

SetConsentBasis sets ConsentBasis field to given value.


### GetMeasurementClass

`func (o *GetEvent200ResponseEvent) GetMeasurementClass() string`

GetMeasurementClass returns the MeasurementClass field if non-nil, zero value otherwise.

### GetMeasurementClassOk

`func (o *GetEvent200ResponseEvent) GetMeasurementClassOk() (*string, bool)`

GetMeasurementClassOk returns a tuple with the MeasurementClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementClass

`func (o *GetEvent200ResponseEvent) SetMeasurementClass(v string)`

SetMeasurementClass sets MeasurementClass field to given value.


### GetAttributionJoin

`func (o *GetEvent200ResponseEvent) GetAttributionJoin() string`

GetAttributionJoin returns the AttributionJoin field if non-nil, zero value otherwise.

### GetAttributionJoinOk

`func (o *GetEvent200ResponseEvent) GetAttributionJoinOk() (*string, bool)`

GetAttributionJoinOk returns a tuple with the AttributionJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionJoin

`func (o *GetEvent200ResponseEvent) SetAttributionJoin(v string)`

SetAttributionJoin sets AttributionJoin field to given value.


### GetEnforcementScope

`func (o *GetEvent200ResponseEvent) GetEnforcementScope() string`

GetEnforcementScope returns the EnforcementScope field if non-nil, zero value otherwise.

### GetEnforcementScopeOk

`func (o *GetEvent200ResponseEvent) GetEnforcementScopeOk() (*string, bool)`

GetEnforcementScopeOk returns a tuple with the EnforcementScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementScope

`func (o *GetEvent200ResponseEvent) SetEnforcementScope(v string)`

SetEnforcementScope sets EnforcementScope field to given value.


### GetConsentNormalizationVersion

`func (o *GetEvent200ResponseEvent) GetConsentNormalizationVersion() string`

GetConsentNormalizationVersion returns the ConsentNormalizationVersion field if non-nil, zero value otherwise.

### GetConsentNormalizationVersionOk

`func (o *GetEvent200ResponseEvent) GetConsentNormalizationVersionOk() (*string, bool)`

GetConsentNormalizationVersionOk returns a tuple with the ConsentNormalizationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentNormalizationVersion

`func (o *GetEvent200ResponseEvent) SetConsentNormalizationVersion(v string)`

SetConsentNormalizationVersion sets ConsentNormalizationVersion field to given value.


### GetConsent

`func (o *GetEvent200ResponseEvent) GetConsent() interface{}`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *GetEvent200ResponseEvent) GetConsentOk() (*interface{}, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *GetEvent200ResponseEvent) SetConsent(v interface{})`

SetConsent sets Consent field to given value.


### SetConsentNil

`func (o *GetEvent200ResponseEvent) SetConsentNil(b bool)`

 SetConsentNil sets the value for Consent to be an explicit nil

### UnsetConsent
`func (o *GetEvent200ResponseEvent) UnsetConsent()`

UnsetConsent ensures that no value is present for Consent, not even an explicit nil
### GetUserDataHashed

`func (o *GetEvent200ResponseEvent) GetUserDataHashed() interface{}`

GetUserDataHashed returns the UserDataHashed field if non-nil, zero value otherwise.

### GetUserDataHashedOk

`func (o *GetEvent200ResponseEvent) GetUserDataHashedOk() (*interface{}, bool)`

GetUserDataHashedOk returns a tuple with the UserDataHashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataHashed

`func (o *GetEvent200ResponseEvent) SetUserDataHashed(v interface{})`

SetUserDataHashed sets UserDataHashed field to given value.


### SetUserDataHashedNil

`func (o *GetEvent200ResponseEvent) SetUserDataHashedNil(b bool)`

 SetUserDataHashedNil sets the value for UserDataHashed to be an explicit nil

### UnsetUserDataHashed
`func (o *GetEvent200ResponseEvent) UnsetUserDataHashed()`

UnsetUserDataHashed ensures that no value is present for UserDataHashed, not even an explicit nil
### GetClickIds

`func (o *GetEvent200ResponseEvent) GetClickIds() interface{}`

GetClickIds returns the ClickIds field if non-nil, zero value otherwise.

### GetClickIdsOk

`func (o *GetEvent200ResponseEvent) GetClickIdsOk() (*interface{}, bool)`

GetClickIdsOk returns a tuple with the ClickIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickIds

`func (o *GetEvent200ResponseEvent) SetClickIds(v interface{})`

SetClickIds sets ClickIds field to given value.


### SetClickIdsNil

`func (o *GetEvent200ResponseEvent) SetClickIdsNil(b bool)`

 SetClickIdsNil sets the value for ClickIds to be an explicit nil

### UnsetClickIds
`func (o *GetEvent200ResponseEvent) UnsetClickIds()`

UnsetClickIds ensures that no value is present for ClickIds, not even an explicit nil
### GetSession

`func (o *GetEvent200ResponseEvent) GetSession() interface{}`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *GetEvent200ResponseEvent) GetSessionOk() (*interface{}, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *GetEvent200ResponseEvent) SetSession(v interface{})`

SetSession sets Session field to given value.


### SetSessionNil

`func (o *GetEvent200ResponseEvent) SetSessionNil(b bool)`

 SetSessionNil sets the value for Session to be an explicit nil

### UnsetSession
`func (o *GetEvent200ResponseEvent) UnsetSession()`

UnsetSession ensures that no value is present for Session, not even an explicit nil
### GetValueData

`func (o *GetEvent200ResponseEvent) GetValueData() interface{}`

GetValueData returns the ValueData field if non-nil, zero value otherwise.

### GetValueDataOk

`func (o *GetEvent200ResponseEvent) GetValueDataOk() (*interface{}, bool)`

GetValueDataOk returns a tuple with the ValueData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueData

`func (o *GetEvent200ResponseEvent) SetValueData(v interface{})`

SetValueData sets ValueData field to given value.


### SetValueDataNil

`func (o *GetEvent200ResponseEvent) SetValueDataNil(b bool)`

 SetValueDataNil sets the value for ValueData to be an explicit nil

### UnsetValueData
`func (o *GetEvent200ResponseEvent) UnsetValueData()`

UnsetValueData ensures that no value is present for ValueData, not even an explicit nil
### GetEventSource

`func (o *GetEvent200ResponseEvent) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *GetEvent200ResponseEvent) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *GetEvent200ResponseEvent) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.


### SetEventSourceNil

`func (o *GetEvent200ResponseEvent) SetEventSourceNil(b bool)`

 SetEventSourceNil sets the value for EventSource to be an explicit nil

### UnsetEventSource
`func (o *GetEvent200ResponseEvent) UnsetEventSource()`

UnsetEventSource ensures that no value is present for EventSource, not even an explicit nil
### GetPayloadExpired

`func (o *GetEvent200ResponseEvent) GetPayloadExpired() bool`

GetPayloadExpired returns the PayloadExpired field if non-nil, zero value otherwise.

### GetPayloadExpiredOk

`func (o *GetEvent200ResponseEvent) GetPayloadExpiredOk() (*bool, bool)`

GetPayloadExpiredOk returns a tuple with the PayloadExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadExpired

`func (o *GetEvent200ResponseEvent) SetPayloadExpired(v bool)`

SetPayloadExpired sets PayloadExpired field to given value.


### GetDeliveries

`func (o *GetEvent200ResponseEvent) GetDeliveries() []GetEvent200ResponseEventDeliveriesInner`

GetDeliveries returns the Deliveries field if non-nil, zero value otherwise.

### GetDeliveriesOk

`func (o *GetEvent200ResponseEvent) GetDeliveriesOk() (*[]GetEvent200ResponseEventDeliveriesInner, bool)`

GetDeliveriesOk returns a tuple with the Deliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveries

`func (o *GetEvent200ResponseEvent) SetDeliveries(v []GetEvent200ResponseEventDeliveriesInner)`

SetDeliveries sets Deliveries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


