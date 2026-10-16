# AdTestVerdictRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outcome** | **NullableString** |  | 
**WinnerAdId** | **NullableString** |  | 
**Reason** | **NullableString** |  | 
**HaltDetails** | Pointer to [**NullableAdTestCreateReadTestVerdictHaltDetails**](AdTestCreateReadTestVerdictHaltDetails.md) |  | [optional] 
**WindowStart** | **NullableString** |  | 
**WindowEnd** | **NullableString** |  | 
**ConcludedAt** | **NullableTime** |  | 
**Members** | [**[]AdTestCreateReadTestVerdictMembersInner**](AdTestCreateReadTestVerdictMembersInner.md) |  | 
**DailyFacts** | [**[]AdTestCreateReadTestVerdictDailyFactsInner**](AdTestCreateReadTestVerdictDailyFactsInner.md) |  | 

## Methods

### NewAdTestVerdictRead

`func NewAdTestVerdictRead(outcome NullableString, winnerAdId NullableString, reason NullableString, windowStart NullableString, windowEnd NullableString, concludedAt NullableTime, members []AdTestCreateReadTestVerdictMembersInner, dailyFacts []AdTestCreateReadTestVerdictDailyFactsInner, ) *AdTestVerdictRead`

NewAdTestVerdictRead instantiates a new AdTestVerdictRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdTestVerdictReadWithDefaults

`func NewAdTestVerdictReadWithDefaults() *AdTestVerdictRead`

NewAdTestVerdictReadWithDefaults instantiates a new AdTestVerdictRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutcome

`func (o *AdTestVerdictRead) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *AdTestVerdictRead) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *AdTestVerdictRead) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### SetOutcomeNil

`func (o *AdTestVerdictRead) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *AdTestVerdictRead) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetWinnerAdId

`func (o *AdTestVerdictRead) GetWinnerAdId() string`

GetWinnerAdId returns the WinnerAdId field if non-nil, zero value otherwise.

### GetWinnerAdIdOk

`func (o *AdTestVerdictRead) GetWinnerAdIdOk() (*string, bool)`

GetWinnerAdIdOk returns a tuple with the WinnerAdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinnerAdId

`func (o *AdTestVerdictRead) SetWinnerAdId(v string)`

SetWinnerAdId sets WinnerAdId field to given value.


### SetWinnerAdIdNil

`func (o *AdTestVerdictRead) SetWinnerAdIdNil(b bool)`

 SetWinnerAdIdNil sets the value for WinnerAdId to be an explicit nil

### UnsetWinnerAdId
`func (o *AdTestVerdictRead) UnsetWinnerAdId()`

UnsetWinnerAdId ensures that no value is present for WinnerAdId, not even an explicit nil
### GetReason

`func (o *AdTestVerdictRead) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AdTestVerdictRead) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AdTestVerdictRead) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *AdTestVerdictRead) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *AdTestVerdictRead) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetHaltDetails

`func (o *AdTestVerdictRead) GetHaltDetails() AdTestCreateReadTestVerdictHaltDetails`

GetHaltDetails returns the HaltDetails field if non-nil, zero value otherwise.

### GetHaltDetailsOk

`func (o *AdTestVerdictRead) GetHaltDetailsOk() (*AdTestCreateReadTestVerdictHaltDetails, bool)`

GetHaltDetailsOk returns a tuple with the HaltDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaltDetails

`func (o *AdTestVerdictRead) SetHaltDetails(v AdTestCreateReadTestVerdictHaltDetails)`

SetHaltDetails sets HaltDetails field to given value.

### HasHaltDetails

`func (o *AdTestVerdictRead) HasHaltDetails() bool`

HasHaltDetails returns a boolean if a field has been set.

### SetHaltDetailsNil

`func (o *AdTestVerdictRead) SetHaltDetailsNil(b bool)`

 SetHaltDetailsNil sets the value for HaltDetails to be an explicit nil

### UnsetHaltDetails
`func (o *AdTestVerdictRead) UnsetHaltDetails()`

UnsetHaltDetails ensures that no value is present for HaltDetails, not even an explicit nil
### GetWindowStart

`func (o *AdTestVerdictRead) GetWindowStart() string`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *AdTestVerdictRead) GetWindowStartOk() (*string, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *AdTestVerdictRead) SetWindowStart(v string)`

SetWindowStart sets WindowStart field to given value.


### SetWindowStartNil

`func (o *AdTestVerdictRead) SetWindowStartNil(b bool)`

 SetWindowStartNil sets the value for WindowStart to be an explicit nil

### UnsetWindowStart
`func (o *AdTestVerdictRead) UnsetWindowStart()`

UnsetWindowStart ensures that no value is present for WindowStart, not even an explicit nil
### GetWindowEnd

`func (o *AdTestVerdictRead) GetWindowEnd() string`

GetWindowEnd returns the WindowEnd field if non-nil, zero value otherwise.

### GetWindowEndOk

`func (o *AdTestVerdictRead) GetWindowEndOk() (*string, bool)`

GetWindowEndOk returns a tuple with the WindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowEnd

`func (o *AdTestVerdictRead) SetWindowEnd(v string)`

SetWindowEnd sets WindowEnd field to given value.


### SetWindowEndNil

`func (o *AdTestVerdictRead) SetWindowEndNil(b bool)`

 SetWindowEndNil sets the value for WindowEnd to be an explicit nil

### UnsetWindowEnd
`func (o *AdTestVerdictRead) UnsetWindowEnd()`

UnsetWindowEnd ensures that no value is present for WindowEnd, not even an explicit nil
### GetConcludedAt

`func (o *AdTestVerdictRead) GetConcludedAt() time.Time`

GetConcludedAt returns the ConcludedAt field if non-nil, zero value otherwise.

### GetConcludedAtOk

`func (o *AdTestVerdictRead) GetConcludedAtOk() (*time.Time, bool)`

GetConcludedAtOk returns a tuple with the ConcludedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcludedAt

`func (o *AdTestVerdictRead) SetConcludedAt(v time.Time)`

SetConcludedAt sets ConcludedAt field to given value.


### SetConcludedAtNil

`func (o *AdTestVerdictRead) SetConcludedAtNil(b bool)`

 SetConcludedAtNil sets the value for ConcludedAt to be an explicit nil

### UnsetConcludedAt
`func (o *AdTestVerdictRead) UnsetConcludedAt()`

UnsetConcludedAt ensures that no value is present for ConcludedAt, not even an explicit nil
### GetMembers

`func (o *AdTestVerdictRead) GetMembers() []AdTestCreateReadTestVerdictMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AdTestVerdictRead) GetMembersOk() (*[]AdTestCreateReadTestVerdictMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AdTestVerdictRead) SetMembers(v []AdTestCreateReadTestVerdictMembersInner)`

SetMembers sets Members field to given value.


### GetDailyFacts

`func (o *AdTestVerdictRead) GetDailyFacts() []AdTestCreateReadTestVerdictDailyFactsInner`

GetDailyFacts returns the DailyFacts field if non-nil, zero value otherwise.

### GetDailyFactsOk

`func (o *AdTestVerdictRead) GetDailyFactsOk() (*[]AdTestCreateReadTestVerdictDailyFactsInner, bool)`

GetDailyFactsOk returns a tuple with the DailyFacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyFacts

`func (o *AdTestVerdictRead) SetDailyFacts(v []AdTestCreateReadTestVerdictDailyFactsInner)`

SetDailyFacts sets DailyFacts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


