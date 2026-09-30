# GetEvent200ResponseLineageChildrenInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**WorkspaceId** | **int32** |  | 
**SignalTrackerId** | **string** | Deprecated alias of workspace_id; contains the workspace ID in decimal string form. | 
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
**UserDataHashed** | **interface{}** | Deprecated compatibility field. The value is always null; delivery identity is never returned. | 
**ClickIds** | **interface{}** |  | 
**Session** | **interface{}** |  | 
**ValueData** | **interface{}** |  | 
**EventSource** | **NullableString** |  | 
**PayloadExpired** | **bool** |  | 

## Methods

### NewGetEvent200ResponseLineageChildrenInner

`func NewGetEvent200ResponseLineageChildrenInner(id string, workspaceId int32, signalTrackerId string, parentEventId NullableString, eventName string, eventTime string, actionSource string, eventClass string, orderId NullableString, valueAmount NullableInt32, valueCurrency NullableString, createdAt string, consentBasis string, measurementClass string, attributionJoin string, enforcementScope string, consentNormalizationVersion string, consent interface{}, userDataHashed interface{}, clickIds interface{}, session interface{}, valueData interface{}, eventSource NullableString, payloadExpired bool, ) *GetEvent200ResponseLineageChildrenInner`

NewGetEvent200ResponseLineageChildrenInner instantiates a new GetEvent200ResponseLineageChildrenInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvent200ResponseLineageChildrenInnerWithDefaults

`func NewGetEvent200ResponseLineageChildrenInnerWithDefaults() *GetEvent200ResponseLineageChildrenInner`

NewGetEvent200ResponseLineageChildrenInnerWithDefaults instantiates a new GetEvent200ResponseLineageChildrenInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEvent200ResponseLineageChildrenInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEvent200ResponseLineageChildrenInner) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *GetEvent200ResponseLineageChildrenInner) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GetEvent200ResponseLineageChildrenInner) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetSignalTrackerId

`func (o *GetEvent200ResponseLineageChildrenInner) GetSignalTrackerId() string`

GetSignalTrackerId returns the SignalTrackerId field if non-nil, zero value otherwise.

### GetSignalTrackerIdOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetSignalTrackerIdOk() (*string, bool)`

GetSignalTrackerIdOk returns a tuple with the SignalTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalTrackerId

`func (o *GetEvent200ResponseLineageChildrenInner) SetSignalTrackerId(v string)`

SetSignalTrackerId sets SignalTrackerId field to given value.


### GetParentEventId

`func (o *GetEvent200ResponseLineageChildrenInner) GetParentEventId() string`

GetParentEventId returns the ParentEventId field if non-nil, zero value otherwise.

### GetParentEventIdOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetParentEventIdOk() (*string, bool)`

GetParentEventIdOk returns a tuple with the ParentEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEventId

`func (o *GetEvent200ResponseLineageChildrenInner) SetParentEventId(v string)`

SetParentEventId sets ParentEventId field to given value.


### SetParentEventIdNil

`func (o *GetEvent200ResponseLineageChildrenInner) SetParentEventIdNil(b bool)`

 SetParentEventIdNil sets the value for ParentEventId to be an explicit nil

### UnsetParentEventId
`func (o *GetEvent200ResponseLineageChildrenInner) UnsetParentEventId()`

UnsetParentEventId ensures that no value is present for ParentEventId, not even an explicit nil
### GetEventName

`func (o *GetEvent200ResponseLineageChildrenInner) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *GetEvent200ResponseLineageChildrenInner) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetEventTime

`func (o *GetEvent200ResponseLineageChildrenInner) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *GetEvent200ResponseLineageChildrenInner) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.


### GetActionSource

`func (o *GetEvent200ResponseLineageChildrenInner) GetActionSource() string`

GetActionSource returns the ActionSource field if non-nil, zero value otherwise.

### GetActionSourceOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetActionSourceOk() (*string, bool)`

GetActionSourceOk returns a tuple with the ActionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionSource

`func (o *GetEvent200ResponseLineageChildrenInner) SetActionSource(v string)`

SetActionSource sets ActionSource field to given value.


### GetEventClass

`func (o *GetEvent200ResponseLineageChildrenInner) GetEventClass() string`

GetEventClass returns the EventClass field if non-nil, zero value otherwise.

### GetEventClassOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetEventClassOk() (*string, bool)`

GetEventClassOk returns a tuple with the EventClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClass

`func (o *GetEvent200ResponseLineageChildrenInner) SetEventClass(v string)`

SetEventClass sets EventClass field to given value.


### GetOrderId

`func (o *GetEvent200ResponseLineageChildrenInner) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetEvent200ResponseLineageChildrenInner) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### SetOrderIdNil

`func (o *GetEvent200ResponseLineageChildrenInner) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *GetEvent200ResponseLineageChildrenInner) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetValueAmount

`func (o *GetEvent200ResponseLineageChildrenInner) GetValueAmount() int32`

GetValueAmount returns the ValueAmount field if non-nil, zero value otherwise.

### GetValueAmountOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetValueAmountOk() (*int32, bool)`

GetValueAmountOk returns a tuple with the ValueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAmount

`func (o *GetEvent200ResponseLineageChildrenInner) SetValueAmount(v int32)`

SetValueAmount sets ValueAmount field to given value.


### SetValueAmountNil

`func (o *GetEvent200ResponseLineageChildrenInner) SetValueAmountNil(b bool)`

 SetValueAmountNil sets the value for ValueAmount to be an explicit nil

### UnsetValueAmount
`func (o *GetEvent200ResponseLineageChildrenInner) UnsetValueAmount()`

UnsetValueAmount ensures that no value is present for ValueAmount, not even an explicit nil
### GetValueCurrency

`func (o *GetEvent200ResponseLineageChildrenInner) GetValueCurrency() string`

GetValueCurrency returns the ValueCurrency field if non-nil, zero value otherwise.

### GetValueCurrencyOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetValueCurrencyOk() (*string, bool)`

GetValueCurrencyOk returns a tuple with the ValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCurrency

`func (o *GetEvent200ResponseLineageChildrenInner) SetValueCurrency(v string)`

SetValueCurrency sets ValueCurrency field to given value.


### SetValueCurrencyNil

`func (o *GetEvent200ResponseLineageChildrenInner) SetValueCurrencyNil(b bool)`

 SetValueCurrencyNil sets the value for ValueCurrency to be an explicit nil

### UnsetValueCurrency
`func (o *GetEvent200ResponseLineageChildrenInner) UnsetValueCurrency()`

UnsetValueCurrency ensures that no value is present for ValueCurrency, not even an explicit nil
### GetCreatedAt

`func (o *GetEvent200ResponseLineageChildrenInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetEvent200ResponseLineageChildrenInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetConsentBasis

`func (o *GetEvent200ResponseLineageChildrenInner) GetConsentBasis() string`

GetConsentBasis returns the ConsentBasis field if non-nil, zero value otherwise.

### GetConsentBasisOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetConsentBasisOk() (*string, bool)`

GetConsentBasisOk returns a tuple with the ConsentBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentBasis

`func (o *GetEvent200ResponseLineageChildrenInner) SetConsentBasis(v string)`

SetConsentBasis sets ConsentBasis field to given value.


### GetMeasurementClass

`func (o *GetEvent200ResponseLineageChildrenInner) GetMeasurementClass() string`

GetMeasurementClass returns the MeasurementClass field if non-nil, zero value otherwise.

### GetMeasurementClassOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetMeasurementClassOk() (*string, bool)`

GetMeasurementClassOk returns a tuple with the MeasurementClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementClass

`func (o *GetEvent200ResponseLineageChildrenInner) SetMeasurementClass(v string)`

SetMeasurementClass sets MeasurementClass field to given value.


### GetAttributionJoin

`func (o *GetEvent200ResponseLineageChildrenInner) GetAttributionJoin() string`

GetAttributionJoin returns the AttributionJoin field if non-nil, zero value otherwise.

### GetAttributionJoinOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetAttributionJoinOk() (*string, bool)`

GetAttributionJoinOk returns a tuple with the AttributionJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionJoin

`func (o *GetEvent200ResponseLineageChildrenInner) SetAttributionJoin(v string)`

SetAttributionJoin sets AttributionJoin field to given value.


### GetEnforcementScope

`func (o *GetEvent200ResponseLineageChildrenInner) GetEnforcementScope() string`

GetEnforcementScope returns the EnforcementScope field if non-nil, zero value otherwise.

### GetEnforcementScopeOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetEnforcementScopeOk() (*string, bool)`

GetEnforcementScopeOk returns a tuple with the EnforcementScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementScope

`func (o *GetEvent200ResponseLineageChildrenInner) SetEnforcementScope(v string)`

SetEnforcementScope sets EnforcementScope field to given value.


### GetConsentNormalizationVersion

`func (o *GetEvent200ResponseLineageChildrenInner) GetConsentNormalizationVersion() string`

GetConsentNormalizationVersion returns the ConsentNormalizationVersion field if non-nil, zero value otherwise.

### GetConsentNormalizationVersionOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetConsentNormalizationVersionOk() (*string, bool)`

GetConsentNormalizationVersionOk returns a tuple with the ConsentNormalizationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentNormalizationVersion

`func (o *GetEvent200ResponseLineageChildrenInner) SetConsentNormalizationVersion(v string)`

SetConsentNormalizationVersion sets ConsentNormalizationVersion field to given value.


### GetConsent

`func (o *GetEvent200ResponseLineageChildrenInner) GetConsent() interface{}`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetConsentOk() (*interface{}, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *GetEvent200ResponseLineageChildrenInner) SetConsent(v interface{})`

SetConsent sets Consent field to given value.


### SetConsentNil

`func (o *GetEvent200ResponseLineageChildrenInner) SetConsentNil(b bool)`

 SetConsentNil sets the value for Consent to be an explicit nil

### UnsetConsent
`func (o *GetEvent200ResponseLineageChildrenInner) UnsetConsent()`

UnsetConsent ensures that no value is present for Consent, not even an explicit nil
### GetUserDataHashed

`func (o *GetEvent200ResponseLineageChildrenInner) GetUserDataHashed() interface{}`

GetUserDataHashed returns the UserDataHashed field if non-nil, zero value otherwise.

### GetUserDataHashedOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetUserDataHashedOk() (*interface{}, bool)`

GetUserDataHashedOk returns a tuple with the UserDataHashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataHashed

`func (o *GetEvent200ResponseLineageChildrenInner) SetUserDataHashed(v interface{})`

SetUserDataHashed sets UserDataHashed field to given value.


### SetUserDataHashedNil

`func (o *GetEvent200ResponseLineageChildrenInner) SetUserDataHashedNil(b bool)`

 SetUserDataHashedNil sets the value for UserDataHashed to be an explicit nil

### UnsetUserDataHashed
`func (o *GetEvent200ResponseLineageChildrenInner) UnsetUserDataHashed()`

UnsetUserDataHashed ensures that no value is present for UserDataHashed, not even an explicit nil
### GetClickIds

`func (o *GetEvent200ResponseLineageChildrenInner) GetClickIds() interface{}`

GetClickIds returns the ClickIds field if non-nil, zero value otherwise.

### GetClickIdsOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetClickIdsOk() (*interface{}, bool)`

GetClickIdsOk returns a tuple with the ClickIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickIds

`func (o *GetEvent200ResponseLineageChildrenInner) SetClickIds(v interface{})`

SetClickIds sets ClickIds field to given value.


### SetClickIdsNil

`func (o *GetEvent200ResponseLineageChildrenInner) SetClickIdsNil(b bool)`

 SetClickIdsNil sets the value for ClickIds to be an explicit nil

### UnsetClickIds
`func (o *GetEvent200ResponseLineageChildrenInner) UnsetClickIds()`

UnsetClickIds ensures that no value is present for ClickIds, not even an explicit nil
### GetSession

`func (o *GetEvent200ResponseLineageChildrenInner) GetSession() interface{}`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetSessionOk() (*interface{}, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *GetEvent200ResponseLineageChildrenInner) SetSession(v interface{})`

SetSession sets Session field to given value.


### SetSessionNil

`func (o *GetEvent200ResponseLineageChildrenInner) SetSessionNil(b bool)`

 SetSessionNil sets the value for Session to be an explicit nil

### UnsetSession
`func (o *GetEvent200ResponseLineageChildrenInner) UnsetSession()`

UnsetSession ensures that no value is present for Session, not even an explicit nil
### GetValueData

`func (o *GetEvent200ResponseLineageChildrenInner) GetValueData() interface{}`

GetValueData returns the ValueData field if non-nil, zero value otherwise.

### GetValueDataOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetValueDataOk() (*interface{}, bool)`

GetValueDataOk returns a tuple with the ValueData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueData

`func (o *GetEvent200ResponseLineageChildrenInner) SetValueData(v interface{})`

SetValueData sets ValueData field to given value.


### SetValueDataNil

`func (o *GetEvent200ResponseLineageChildrenInner) SetValueDataNil(b bool)`

 SetValueDataNil sets the value for ValueData to be an explicit nil

### UnsetValueData
`func (o *GetEvent200ResponseLineageChildrenInner) UnsetValueData()`

UnsetValueData ensures that no value is present for ValueData, not even an explicit nil
### GetEventSource

`func (o *GetEvent200ResponseLineageChildrenInner) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *GetEvent200ResponseLineageChildrenInner) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.


### SetEventSourceNil

`func (o *GetEvent200ResponseLineageChildrenInner) SetEventSourceNil(b bool)`

 SetEventSourceNil sets the value for EventSource to be an explicit nil

### UnsetEventSource
`func (o *GetEvent200ResponseLineageChildrenInner) UnsetEventSource()`

UnsetEventSource ensures that no value is present for EventSource, not even an explicit nil
### GetPayloadExpired

`func (o *GetEvent200ResponseLineageChildrenInner) GetPayloadExpired() bool`

GetPayloadExpired returns the PayloadExpired field if non-nil, zero value otherwise.

### GetPayloadExpiredOk

`func (o *GetEvent200ResponseLineageChildrenInner) GetPayloadExpiredOk() (*bool, bool)`

GetPayloadExpiredOk returns a tuple with the PayloadExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadExpired

`func (o *GetEvent200ResponseLineageChildrenInner) SetPayloadExpired(v bool)`

SetPayloadExpired sets PayloadExpired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


