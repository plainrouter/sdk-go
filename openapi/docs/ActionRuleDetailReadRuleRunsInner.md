# ActionRuleDetailReadRuleRunsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**RevisionId** | **string** |  | 
**WindowStart** | **time.Time** |  | 
**WindowEnd** | **time.Time** |  | 
**Timezone** | **string** |  | 
**Status** | **string** |  | 
**Reason** | **NullableString** |  | 
**AccountSpendMinor** | **NullableString** |  | 
**AccountCountedArrivals** | **NullableInt32** |  | 
**EvaluatedAt** | **NullableTime** |  | 
**Decisions** | [**[]ActionRuleDetailReadRuleRunsInnerDecisionsInner**](ActionRuleDetailReadRuleRunsInnerDecisionsInner.md) |  | 
**DecisionsTruncated** | **bool** |  | 

## Methods

### NewActionRuleDetailReadRuleRunsInner

`func NewActionRuleDetailReadRuleRunsInner(id string, revisionId string, windowStart time.Time, windowEnd time.Time, timezone string, status string, reason NullableString, accountSpendMinor NullableString, accountCountedArrivals NullableInt32, evaluatedAt NullableTime, decisions []ActionRuleDetailReadRuleRunsInnerDecisionsInner, decisionsTruncated bool, ) *ActionRuleDetailReadRuleRunsInner`

NewActionRuleDetailReadRuleRunsInner instantiates a new ActionRuleDetailReadRuleRunsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRuleDetailReadRuleRunsInnerWithDefaults

`func NewActionRuleDetailReadRuleRunsInnerWithDefaults() *ActionRuleDetailReadRuleRunsInner`

NewActionRuleDetailReadRuleRunsInnerWithDefaults instantiates a new ActionRuleDetailReadRuleRunsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionRuleDetailReadRuleRunsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionRuleDetailReadRuleRunsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionRuleDetailReadRuleRunsInner) SetId(v string)`

SetId sets Id field to given value.


### GetRevisionId

`func (o *ActionRuleDetailReadRuleRunsInner) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *ActionRuleDetailReadRuleRunsInner) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *ActionRuleDetailReadRuleRunsInner) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.


### GetWindowStart

`func (o *ActionRuleDetailReadRuleRunsInner) GetWindowStart() time.Time`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *ActionRuleDetailReadRuleRunsInner) GetWindowStartOk() (*time.Time, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *ActionRuleDetailReadRuleRunsInner) SetWindowStart(v time.Time)`

SetWindowStart sets WindowStart field to given value.


### GetWindowEnd

`func (o *ActionRuleDetailReadRuleRunsInner) GetWindowEnd() time.Time`

GetWindowEnd returns the WindowEnd field if non-nil, zero value otherwise.

### GetWindowEndOk

`func (o *ActionRuleDetailReadRuleRunsInner) GetWindowEndOk() (*time.Time, bool)`

GetWindowEndOk returns a tuple with the WindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowEnd

`func (o *ActionRuleDetailReadRuleRunsInner) SetWindowEnd(v time.Time)`

SetWindowEnd sets WindowEnd field to given value.


### GetTimezone

`func (o *ActionRuleDetailReadRuleRunsInner) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ActionRuleDetailReadRuleRunsInner) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ActionRuleDetailReadRuleRunsInner) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStatus

`func (o *ActionRuleDetailReadRuleRunsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionRuleDetailReadRuleRunsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionRuleDetailReadRuleRunsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *ActionRuleDetailReadRuleRunsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ActionRuleDetailReadRuleRunsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ActionRuleDetailReadRuleRunsInner) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *ActionRuleDetailReadRuleRunsInner) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ActionRuleDetailReadRuleRunsInner) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetAccountSpendMinor

`func (o *ActionRuleDetailReadRuleRunsInner) GetAccountSpendMinor() string`

GetAccountSpendMinor returns the AccountSpendMinor field if non-nil, zero value otherwise.

### GetAccountSpendMinorOk

`func (o *ActionRuleDetailReadRuleRunsInner) GetAccountSpendMinorOk() (*string, bool)`

GetAccountSpendMinorOk returns a tuple with the AccountSpendMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSpendMinor

`func (o *ActionRuleDetailReadRuleRunsInner) SetAccountSpendMinor(v string)`

SetAccountSpendMinor sets AccountSpendMinor field to given value.


### SetAccountSpendMinorNil

`func (o *ActionRuleDetailReadRuleRunsInner) SetAccountSpendMinorNil(b bool)`

 SetAccountSpendMinorNil sets the value for AccountSpendMinor to be an explicit nil

### UnsetAccountSpendMinor
`func (o *ActionRuleDetailReadRuleRunsInner) UnsetAccountSpendMinor()`

UnsetAccountSpendMinor ensures that no value is present for AccountSpendMinor, not even an explicit nil
### GetAccountCountedArrivals

`func (o *ActionRuleDetailReadRuleRunsInner) GetAccountCountedArrivals() int32`

GetAccountCountedArrivals returns the AccountCountedArrivals field if non-nil, zero value otherwise.

### GetAccountCountedArrivalsOk

`func (o *ActionRuleDetailReadRuleRunsInner) GetAccountCountedArrivalsOk() (*int32, bool)`

GetAccountCountedArrivalsOk returns a tuple with the AccountCountedArrivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCountedArrivals

`func (o *ActionRuleDetailReadRuleRunsInner) SetAccountCountedArrivals(v int32)`

SetAccountCountedArrivals sets AccountCountedArrivals field to given value.


### SetAccountCountedArrivalsNil

`func (o *ActionRuleDetailReadRuleRunsInner) SetAccountCountedArrivalsNil(b bool)`

 SetAccountCountedArrivalsNil sets the value for AccountCountedArrivals to be an explicit nil

### UnsetAccountCountedArrivals
`func (o *ActionRuleDetailReadRuleRunsInner) UnsetAccountCountedArrivals()`

UnsetAccountCountedArrivals ensures that no value is present for AccountCountedArrivals, not even an explicit nil
### GetEvaluatedAt

`func (o *ActionRuleDetailReadRuleRunsInner) GetEvaluatedAt() time.Time`

GetEvaluatedAt returns the EvaluatedAt field if non-nil, zero value otherwise.

### GetEvaluatedAtOk

`func (o *ActionRuleDetailReadRuleRunsInner) GetEvaluatedAtOk() (*time.Time, bool)`

GetEvaluatedAtOk returns a tuple with the EvaluatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedAt

`func (o *ActionRuleDetailReadRuleRunsInner) SetEvaluatedAt(v time.Time)`

SetEvaluatedAt sets EvaluatedAt field to given value.


### SetEvaluatedAtNil

`func (o *ActionRuleDetailReadRuleRunsInner) SetEvaluatedAtNil(b bool)`

 SetEvaluatedAtNil sets the value for EvaluatedAt to be an explicit nil

### UnsetEvaluatedAt
`func (o *ActionRuleDetailReadRuleRunsInner) UnsetEvaluatedAt()`

UnsetEvaluatedAt ensures that no value is present for EvaluatedAt, not even an explicit nil
### GetDecisions

`func (o *ActionRuleDetailReadRuleRunsInner) GetDecisions() []ActionRuleDetailReadRuleRunsInnerDecisionsInner`

GetDecisions returns the Decisions field if non-nil, zero value otherwise.

### GetDecisionsOk

`func (o *ActionRuleDetailReadRuleRunsInner) GetDecisionsOk() (*[]ActionRuleDetailReadRuleRunsInnerDecisionsInner, bool)`

GetDecisionsOk returns a tuple with the Decisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisions

`func (o *ActionRuleDetailReadRuleRunsInner) SetDecisions(v []ActionRuleDetailReadRuleRunsInnerDecisionsInner)`

SetDecisions sets Decisions field to given value.


### GetDecisionsTruncated

`func (o *ActionRuleDetailReadRuleRunsInner) GetDecisionsTruncated() bool`

GetDecisionsTruncated returns the DecisionsTruncated field if non-nil, zero value otherwise.

### GetDecisionsTruncatedOk

`func (o *ActionRuleDetailReadRuleRunsInner) GetDecisionsTruncatedOk() (*bool, bool)`

GetDecisionsTruncatedOk returns a tuple with the DecisionsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionsTruncated

`func (o *ActionRuleDetailReadRuleRunsInner) SetDecisionsTruncated(v bool)`

SetDecisionsTruncated sets DecisionsTruncated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


