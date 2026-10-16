# AdTestCreateReadTestPauseProposalsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdId** | **string** |  | 
**Status** | **string** |  | 
**Reason** | **NullableString** | Stable reason code when the proposal was skipped or expired. | 
**DecisionDeadlineAt** | **time.Time** | Immutable deadline 24 hours after the Test&#39;s candidates are created on the database clock. | 

## Methods

### NewAdTestCreateReadTestPauseProposalsInner

`func NewAdTestCreateReadTestPauseProposalsInner(adId string, status string, reason NullableString, decisionDeadlineAt time.Time, ) *AdTestCreateReadTestPauseProposalsInner`

NewAdTestCreateReadTestPauseProposalsInner instantiates a new AdTestCreateReadTestPauseProposalsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdTestCreateReadTestPauseProposalsInnerWithDefaults

`func NewAdTestCreateReadTestPauseProposalsInnerWithDefaults() *AdTestCreateReadTestPauseProposalsInner`

NewAdTestCreateReadTestPauseProposalsInnerWithDefaults instantiates a new AdTestCreateReadTestPauseProposalsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdId

`func (o *AdTestCreateReadTestPauseProposalsInner) GetAdId() string`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *AdTestCreateReadTestPauseProposalsInner) GetAdIdOk() (*string, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *AdTestCreateReadTestPauseProposalsInner) SetAdId(v string)`

SetAdId sets AdId field to given value.


### GetStatus

`func (o *AdTestCreateReadTestPauseProposalsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdTestCreateReadTestPauseProposalsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdTestCreateReadTestPauseProposalsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *AdTestCreateReadTestPauseProposalsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AdTestCreateReadTestPauseProposalsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AdTestCreateReadTestPauseProposalsInner) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *AdTestCreateReadTestPauseProposalsInner) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *AdTestCreateReadTestPauseProposalsInner) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetDecisionDeadlineAt

`func (o *AdTestCreateReadTestPauseProposalsInner) GetDecisionDeadlineAt() time.Time`

GetDecisionDeadlineAt returns the DecisionDeadlineAt field if non-nil, zero value otherwise.

### GetDecisionDeadlineAtOk

`func (o *AdTestCreateReadTestPauseProposalsInner) GetDecisionDeadlineAtOk() (*time.Time, bool)`

GetDecisionDeadlineAtOk returns a tuple with the DecisionDeadlineAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDeadlineAt

`func (o *AdTestCreateReadTestPauseProposalsInner) SetDecisionDeadlineAt(v time.Time)`

SetDecisionDeadlineAt sets DecisionDeadlineAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


