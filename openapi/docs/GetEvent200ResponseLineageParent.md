# GetEvent200ResponseLineageParent

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

### NewGetEvent200ResponseLineageParent

`func NewGetEvent200ResponseLineageParent(id string, workspaceId int32, signalTrackerId string, parentEventId NullableString, eventName string, eventTime string, actionSource string, eventClass string, orderId NullableString, valueAmount NullableInt32, valueCurrency NullableString, createdAt string, consentBasis string, measurementClass string, attributionJoin string, enforcementScope string, consentNormalizationVersion string, consent interface{}, userDataHashed interface{}, clickIds interface{}, session interface{}, valueData interface{}, eventSource NullableString, payloadExpired bool, ) *GetEvent200ResponseLineageParent`

NewGetEvent200ResponseLineageParent instantiates a new GetEvent200ResponseLineageParent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvent200ResponseLineageParentWithDefaults

`func NewGetEvent200ResponseLineageParentWithDefaults() *GetEvent200ResponseLineageParent`

NewGetEvent200ResponseLineageParentWithDefaults instantiates a new GetEvent200ResponseLineageParent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEvent200ResponseLineageParent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEvent200ResponseLineageParent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEvent200ResponseLineageParent) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *GetEvent200ResponseLineageParent) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GetEvent200ResponseLineageParent) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GetEvent200ResponseLineageParent) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetSignalTrackerId

`func (o *GetEvent200ResponseLineageParent) GetSignalTrackerId() string`

GetSignalTrackerId returns the SignalTrackerId field if non-nil, zero value otherwise.

### GetSignalTrackerIdOk

`func (o *GetEvent200ResponseLineageParent) GetSignalTrackerIdOk() (*string, bool)`

GetSignalTrackerIdOk returns a tuple with the SignalTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalTrackerId

`func (o *GetEvent200ResponseLineageParent) SetSignalTrackerId(v string)`

SetSignalTrackerId sets SignalTrackerId field to given value.


### GetParentEventId

`func (o *GetEvent200ResponseLineageParent) GetParentEventId() string`

GetParentEventId returns the ParentEventId field if non-nil, zero value otherwise.

### GetParentEventIdOk

`func (o *GetEvent200ResponseLineageParent) GetParentEventIdOk() (*string, bool)`

GetParentEventIdOk returns a tuple with the ParentEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEventId

`func (o *GetEvent200ResponseLineageParent) SetParentEventId(v string)`

SetParentEventId sets ParentEventId field to given value.


### SetParentEventIdNil

`func (o *GetEvent200ResponseLineageParent) SetParentEventIdNil(b bool)`

 SetParentEventIdNil sets the value for ParentEventId to be an explicit nil

### UnsetParentEventId
`func (o *GetEvent200ResponseLineageParent) UnsetParentEventId()`

UnsetParentEventId ensures that no value is present for ParentEventId, not even an explicit nil
### GetEventName

`func (o *GetEvent200ResponseLineageParent) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *GetEvent200ResponseLineageParent) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *GetEvent200ResponseLineageParent) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetEventTime

`func (o *GetEvent200ResponseLineageParent) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *GetEvent200ResponseLineageParent) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *GetEvent200ResponseLineageParent) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.


### GetActionSource

`func (o *GetEvent200ResponseLineageParent) GetActionSource() string`

GetActionSource returns the ActionSource field if non-nil, zero value otherwise.

### GetActionSourceOk

`func (o *GetEvent200ResponseLineageParent) GetActionSourceOk() (*string, bool)`

GetActionSourceOk returns a tuple with the ActionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionSource

`func (o *GetEvent200ResponseLineageParent) SetActionSource(v string)`

SetActionSource sets ActionSource field to given value.


### GetEventClass

`func (o *GetEvent200ResponseLineageParent) GetEventClass() string`

GetEventClass returns the EventClass field if non-nil, zero value otherwise.

### GetEventClassOk

`func (o *GetEvent200ResponseLineageParent) GetEventClassOk() (*string, bool)`

GetEventClassOk returns a tuple with the EventClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClass

`func (o *GetEvent200ResponseLineageParent) SetEventClass(v string)`

SetEventClass sets EventClass field to given value.


### GetOrderId

`func (o *GetEvent200ResponseLineageParent) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetEvent200ResponseLineageParent) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetEvent200ResponseLineageParent) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### SetOrderIdNil

`func (o *GetEvent200ResponseLineageParent) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *GetEvent200ResponseLineageParent) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetValueAmount

`func (o *GetEvent200ResponseLineageParent) GetValueAmount() int32`

GetValueAmount returns the ValueAmount field if non-nil, zero value otherwise.

### GetValueAmountOk

`func (o *GetEvent200ResponseLineageParent) GetValueAmountOk() (*int32, bool)`

GetValueAmountOk returns a tuple with the ValueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAmount

`func (o *GetEvent200ResponseLineageParent) SetValueAmount(v int32)`

SetValueAmount sets ValueAmount field to given value.


### SetValueAmountNil

`func (o *GetEvent200ResponseLineageParent) SetValueAmountNil(b bool)`

 SetValueAmountNil sets the value for ValueAmount to be an explicit nil

### UnsetValueAmount
`func (o *GetEvent200ResponseLineageParent) UnsetValueAmount()`

UnsetValueAmount ensures that no value is present for ValueAmount, not even an explicit nil
### GetValueCurrency

`func (o *GetEvent200ResponseLineageParent) GetValueCurrency() string`

GetValueCurrency returns the ValueCurrency field if non-nil, zero value otherwise.

### GetValueCurrencyOk

`func (o *GetEvent200ResponseLineageParent) GetValueCurrencyOk() (*string, bool)`

GetValueCurrencyOk returns a tuple with the ValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCurrency

`func (o *GetEvent200ResponseLineageParent) SetValueCurrency(v string)`

SetValueCurrency sets ValueCurrency field to given value.


### SetValueCurrencyNil

`func (o *GetEvent200ResponseLineageParent) SetValueCurrencyNil(b bool)`

 SetValueCurrencyNil sets the value for ValueCurrency to be an explicit nil

### UnsetValueCurrency
`func (o *GetEvent200ResponseLineageParent) UnsetValueCurrency()`

UnsetValueCurrency ensures that no value is present for ValueCurrency, not even an explicit nil
### GetCreatedAt

`func (o *GetEvent200ResponseLineageParent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetEvent200ResponseLineageParent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetEvent200ResponseLineageParent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetConsentBasis

`func (o *GetEvent200ResponseLineageParent) GetConsentBasis() string`

GetConsentBasis returns the ConsentBasis field if non-nil, zero value otherwise.

### GetConsentBasisOk

`func (o *GetEvent200ResponseLineageParent) GetConsentBasisOk() (*string, bool)`

GetConsentBasisOk returns a tuple with the ConsentBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentBasis

`func (o *GetEvent200ResponseLineageParent) SetConsentBasis(v string)`

SetConsentBasis sets ConsentBasis field to given value.


### GetMeasurementClass

`func (o *GetEvent200ResponseLineageParent) GetMeasurementClass() string`

GetMeasurementClass returns the MeasurementClass field if non-nil, zero value otherwise.

### GetMeasurementClassOk

`func (o *GetEvent200ResponseLineageParent) GetMeasurementClassOk() (*string, bool)`

GetMeasurementClassOk returns a tuple with the MeasurementClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementClass

`func (o *GetEvent200ResponseLineageParent) SetMeasurementClass(v string)`

SetMeasurementClass sets MeasurementClass field to given value.


### GetAttributionJoin

`func (o *GetEvent200ResponseLineageParent) GetAttributionJoin() string`

GetAttributionJoin returns the AttributionJoin field if non-nil, zero value otherwise.

### GetAttributionJoinOk

`func (o *GetEvent200ResponseLineageParent) GetAttributionJoinOk() (*string, bool)`

GetAttributionJoinOk returns a tuple with the AttributionJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionJoin

`func (o *GetEvent200ResponseLineageParent) SetAttributionJoin(v string)`

SetAttributionJoin sets AttributionJoin field to given value.


### GetEnforcementScope

`func (o *GetEvent200ResponseLineageParent) GetEnforcementScope() string`

GetEnforcementScope returns the EnforcementScope field if non-nil, zero value otherwise.

### GetEnforcementScopeOk

`func (o *GetEvent200ResponseLineageParent) GetEnforcementScopeOk() (*string, bool)`

GetEnforcementScopeOk returns a tuple with the EnforcementScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementScope

`func (o *GetEvent200ResponseLineageParent) SetEnforcementScope(v string)`

SetEnforcementScope sets EnforcementScope field to given value.


### GetConsentNormalizationVersion

`func (o *GetEvent200ResponseLineageParent) GetConsentNormalizationVersion() string`

GetConsentNormalizationVersion returns the ConsentNormalizationVersion field if non-nil, zero value otherwise.

### GetConsentNormalizationVersionOk

`func (o *GetEvent200ResponseLineageParent) GetConsentNormalizationVersionOk() (*string, bool)`

GetConsentNormalizationVersionOk returns a tuple with the ConsentNormalizationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentNormalizationVersion

`func (o *GetEvent200ResponseLineageParent) SetConsentNormalizationVersion(v string)`

SetConsentNormalizationVersion sets ConsentNormalizationVersion field to given value.


### GetConsent

`func (o *GetEvent200ResponseLineageParent) GetConsent() interface{}`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *GetEvent200ResponseLineageParent) GetConsentOk() (*interface{}, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *GetEvent200ResponseLineageParent) SetConsent(v interface{})`

SetConsent sets Consent field to given value.


### SetConsentNil

`func (o *GetEvent200ResponseLineageParent) SetConsentNil(b bool)`

 SetConsentNil sets the value for Consent to be an explicit nil

### UnsetConsent
`func (o *GetEvent200ResponseLineageParent) UnsetConsent()`

UnsetConsent ensures that no value is present for Consent, not even an explicit nil
### GetUserDataHashed

`func (o *GetEvent200ResponseLineageParent) GetUserDataHashed() interface{}`

GetUserDataHashed returns the UserDataHashed field if non-nil, zero value otherwise.

### GetUserDataHashedOk

`func (o *GetEvent200ResponseLineageParent) GetUserDataHashedOk() (*interface{}, bool)`

GetUserDataHashedOk returns a tuple with the UserDataHashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataHashed

`func (o *GetEvent200ResponseLineageParent) SetUserDataHashed(v interface{})`

SetUserDataHashed sets UserDataHashed field to given value.


### SetUserDataHashedNil

`func (o *GetEvent200ResponseLineageParent) SetUserDataHashedNil(b bool)`

 SetUserDataHashedNil sets the value for UserDataHashed to be an explicit nil

### UnsetUserDataHashed
`func (o *GetEvent200ResponseLineageParent) UnsetUserDataHashed()`

UnsetUserDataHashed ensures that no value is present for UserDataHashed, not even an explicit nil
### GetClickIds

`func (o *GetEvent200ResponseLineageParent) GetClickIds() interface{}`

GetClickIds returns the ClickIds field if non-nil, zero value otherwise.

### GetClickIdsOk

`func (o *GetEvent200ResponseLineageParent) GetClickIdsOk() (*interface{}, bool)`

GetClickIdsOk returns a tuple with the ClickIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickIds

`func (o *GetEvent200ResponseLineageParent) SetClickIds(v interface{})`

SetClickIds sets ClickIds field to given value.


### SetClickIdsNil

`func (o *GetEvent200ResponseLineageParent) SetClickIdsNil(b bool)`

 SetClickIdsNil sets the value for ClickIds to be an explicit nil

### UnsetClickIds
`func (o *GetEvent200ResponseLineageParent) UnsetClickIds()`

UnsetClickIds ensures that no value is present for ClickIds, not even an explicit nil
### GetSession

`func (o *GetEvent200ResponseLineageParent) GetSession() interface{}`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *GetEvent200ResponseLineageParent) GetSessionOk() (*interface{}, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *GetEvent200ResponseLineageParent) SetSession(v interface{})`

SetSession sets Session field to given value.


### SetSessionNil

`func (o *GetEvent200ResponseLineageParent) SetSessionNil(b bool)`

 SetSessionNil sets the value for Session to be an explicit nil

### UnsetSession
`func (o *GetEvent200ResponseLineageParent) UnsetSession()`

UnsetSession ensures that no value is present for Session, not even an explicit nil
### GetValueData

`func (o *GetEvent200ResponseLineageParent) GetValueData() interface{}`

GetValueData returns the ValueData field if non-nil, zero value otherwise.

### GetValueDataOk

`func (o *GetEvent200ResponseLineageParent) GetValueDataOk() (*interface{}, bool)`

GetValueDataOk returns a tuple with the ValueData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueData

`func (o *GetEvent200ResponseLineageParent) SetValueData(v interface{})`

SetValueData sets ValueData field to given value.


### SetValueDataNil

`func (o *GetEvent200ResponseLineageParent) SetValueDataNil(b bool)`

 SetValueDataNil sets the value for ValueData to be an explicit nil

### UnsetValueData
`func (o *GetEvent200ResponseLineageParent) UnsetValueData()`

UnsetValueData ensures that no value is present for ValueData, not even an explicit nil
### GetEventSource

`func (o *GetEvent200ResponseLineageParent) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *GetEvent200ResponseLineageParent) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *GetEvent200ResponseLineageParent) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.


### SetEventSourceNil

`func (o *GetEvent200ResponseLineageParent) SetEventSourceNil(b bool)`

 SetEventSourceNil sets the value for EventSource to be an explicit nil

### UnsetEventSource
`func (o *GetEvent200ResponseLineageParent) UnsetEventSource()`

UnsetEventSource ensures that no value is present for EventSource, not even an explicit nil
### GetPayloadExpired

`func (o *GetEvent200ResponseLineageParent) GetPayloadExpired() bool`

GetPayloadExpired returns the PayloadExpired field if non-nil, zero value otherwise.

### GetPayloadExpiredOk

`func (o *GetEvent200ResponseLineageParent) GetPayloadExpiredOk() (*bool, bool)`

GetPayloadExpiredOk returns a tuple with the PayloadExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadExpired

`func (o *GetEvent200ResponseLineageParent) SetPayloadExpired(v bool)`

SetPayloadExpired sets PayloadExpired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


