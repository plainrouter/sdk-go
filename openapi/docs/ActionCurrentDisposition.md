# ActionCurrentDisposition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReceiptStatus** | **NullableString** |  | 
**OutcomeStatus** | **NullableString** |  | 
**OutcomeReasonCode** | **NullableString** |  | 
**OutcomeCheckedAt** | **NullableTime** |  | 
**CompensationReasonCode** | **NullableString** |  | 
**RecoveryDisposition** | **NullableString** |  | 
**LateRestored** | **bool** |  | 

## Methods

### NewActionCurrentDisposition

`func NewActionCurrentDisposition(receiptStatus NullableString, outcomeStatus NullableString, outcomeReasonCode NullableString, outcomeCheckedAt NullableTime, compensationReasonCode NullableString, recoveryDisposition NullableString, lateRestored bool, ) *ActionCurrentDisposition`

NewActionCurrentDisposition instantiates a new ActionCurrentDisposition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionCurrentDispositionWithDefaults

`func NewActionCurrentDispositionWithDefaults() *ActionCurrentDisposition`

NewActionCurrentDispositionWithDefaults instantiates a new ActionCurrentDisposition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiptStatus

`func (o *ActionCurrentDisposition) GetReceiptStatus() string`

GetReceiptStatus returns the ReceiptStatus field if non-nil, zero value otherwise.

### GetReceiptStatusOk

`func (o *ActionCurrentDisposition) GetReceiptStatusOk() (*string, bool)`

GetReceiptStatusOk returns a tuple with the ReceiptStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptStatus

`func (o *ActionCurrentDisposition) SetReceiptStatus(v string)`

SetReceiptStatus sets ReceiptStatus field to given value.


### SetReceiptStatusNil

`func (o *ActionCurrentDisposition) SetReceiptStatusNil(b bool)`

 SetReceiptStatusNil sets the value for ReceiptStatus to be an explicit nil

### UnsetReceiptStatus
`func (o *ActionCurrentDisposition) UnsetReceiptStatus()`

UnsetReceiptStatus ensures that no value is present for ReceiptStatus, not even an explicit nil
### GetOutcomeStatus

`func (o *ActionCurrentDisposition) GetOutcomeStatus() string`

GetOutcomeStatus returns the OutcomeStatus field if non-nil, zero value otherwise.

### GetOutcomeStatusOk

`func (o *ActionCurrentDisposition) GetOutcomeStatusOk() (*string, bool)`

GetOutcomeStatusOk returns a tuple with the OutcomeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeStatus

`func (o *ActionCurrentDisposition) SetOutcomeStatus(v string)`

SetOutcomeStatus sets OutcomeStatus field to given value.


### SetOutcomeStatusNil

`func (o *ActionCurrentDisposition) SetOutcomeStatusNil(b bool)`

 SetOutcomeStatusNil sets the value for OutcomeStatus to be an explicit nil

### UnsetOutcomeStatus
`func (o *ActionCurrentDisposition) UnsetOutcomeStatus()`

UnsetOutcomeStatus ensures that no value is present for OutcomeStatus, not even an explicit nil
### GetOutcomeReasonCode

`func (o *ActionCurrentDisposition) GetOutcomeReasonCode() string`

GetOutcomeReasonCode returns the OutcomeReasonCode field if non-nil, zero value otherwise.

### GetOutcomeReasonCodeOk

`func (o *ActionCurrentDisposition) GetOutcomeReasonCodeOk() (*string, bool)`

GetOutcomeReasonCodeOk returns a tuple with the OutcomeReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeReasonCode

`func (o *ActionCurrentDisposition) SetOutcomeReasonCode(v string)`

SetOutcomeReasonCode sets OutcomeReasonCode field to given value.


### SetOutcomeReasonCodeNil

`func (o *ActionCurrentDisposition) SetOutcomeReasonCodeNil(b bool)`

 SetOutcomeReasonCodeNil sets the value for OutcomeReasonCode to be an explicit nil

### UnsetOutcomeReasonCode
`func (o *ActionCurrentDisposition) UnsetOutcomeReasonCode()`

UnsetOutcomeReasonCode ensures that no value is present for OutcomeReasonCode, not even an explicit nil
### GetOutcomeCheckedAt

`func (o *ActionCurrentDisposition) GetOutcomeCheckedAt() time.Time`

GetOutcomeCheckedAt returns the OutcomeCheckedAt field if non-nil, zero value otherwise.

### GetOutcomeCheckedAtOk

`func (o *ActionCurrentDisposition) GetOutcomeCheckedAtOk() (*time.Time, bool)`

GetOutcomeCheckedAtOk returns a tuple with the OutcomeCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeCheckedAt

`func (o *ActionCurrentDisposition) SetOutcomeCheckedAt(v time.Time)`

SetOutcomeCheckedAt sets OutcomeCheckedAt field to given value.


### SetOutcomeCheckedAtNil

`func (o *ActionCurrentDisposition) SetOutcomeCheckedAtNil(b bool)`

 SetOutcomeCheckedAtNil sets the value for OutcomeCheckedAt to be an explicit nil

### UnsetOutcomeCheckedAt
`func (o *ActionCurrentDisposition) UnsetOutcomeCheckedAt()`

UnsetOutcomeCheckedAt ensures that no value is present for OutcomeCheckedAt, not even an explicit nil
### GetCompensationReasonCode

`func (o *ActionCurrentDisposition) GetCompensationReasonCode() string`

GetCompensationReasonCode returns the CompensationReasonCode field if non-nil, zero value otherwise.

### GetCompensationReasonCodeOk

`func (o *ActionCurrentDisposition) GetCompensationReasonCodeOk() (*string, bool)`

GetCompensationReasonCodeOk returns a tuple with the CompensationReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompensationReasonCode

`func (o *ActionCurrentDisposition) SetCompensationReasonCode(v string)`

SetCompensationReasonCode sets CompensationReasonCode field to given value.


### SetCompensationReasonCodeNil

`func (o *ActionCurrentDisposition) SetCompensationReasonCodeNil(b bool)`

 SetCompensationReasonCodeNil sets the value for CompensationReasonCode to be an explicit nil

### UnsetCompensationReasonCode
`func (o *ActionCurrentDisposition) UnsetCompensationReasonCode()`

UnsetCompensationReasonCode ensures that no value is present for CompensationReasonCode, not even an explicit nil
### GetRecoveryDisposition

`func (o *ActionCurrentDisposition) GetRecoveryDisposition() string`

GetRecoveryDisposition returns the RecoveryDisposition field if non-nil, zero value otherwise.

### GetRecoveryDispositionOk

`func (o *ActionCurrentDisposition) GetRecoveryDispositionOk() (*string, bool)`

GetRecoveryDispositionOk returns a tuple with the RecoveryDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryDisposition

`func (o *ActionCurrentDisposition) SetRecoveryDisposition(v string)`

SetRecoveryDisposition sets RecoveryDisposition field to given value.


### SetRecoveryDispositionNil

`func (o *ActionCurrentDisposition) SetRecoveryDispositionNil(b bool)`

 SetRecoveryDispositionNil sets the value for RecoveryDisposition to be an explicit nil

### UnsetRecoveryDisposition
`func (o *ActionCurrentDisposition) UnsetRecoveryDisposition()`

UnsetRecoveryDisposition ensures that no value is present for RecoveryDisposition, not even an explicit nil
### GetLateRestored

`func (o *ActionCurrentDisposition) GetLateRestored() bool`

GetLateRestored returns the LateRestored field if non-nil, zero value otherwise.

### GetLateRestoredOk

`func (o *ActionCurrentDisposition) GetLateRestoredOk() (*bool, bool)`

GetLateRestoredOk returns a tuple with the LateRestored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLateRestored

`func (o *ActionCurrentDisposition) SetLateRestored(v bool)`

SetLateRestored sets LateRestored field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


