# ActionDryRunReadDryRunActionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**TargetEntity** | [**ActionDryRunReadDryRunActionsInnerAnyOfTargetEntity**](ActionDryRunReadDryRunActionsInnerAnyOfTargetEntity.md) |  | 
**Status** | **string** |  | 
**PolicyDecision** | **string** |  | 
**PolicyReasons** | **[]string** |  | 
**ApprovalRequired** | **bool** |  | 
**WouldAutoExecute** | **bool** |  | 
**Diff** | [**ActionDryRunReadDryRunActionsInnerAnyOfDiff**](ActionDryRunReadDryRunActionsInnerAnyOfDiff.md) |  | 
**Reason** | **string** |  | 

## Methods

### NewActionDryRunReadDryRunActionsInner

`func NewActionDryRunReadDryRunActionsInner(type_ string, targetEntity ActionDryRunReadDryRunActionsInnerAnyOfTargetEntity, status string, policyDecision string, policyReasons []string, approvalRequired bool, wouldAutoExecute bool, diff ActionDryRunReadDryRunActionsInnerAnyOfDiff, reason string, ) *ActionDryRunReadDryRunActionsInner`

NewActionDryRunReadDryRunActionsInner instantiates a new ActionDryRunReadDryRunActionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionDryRunReadDryRunActionsInnerWithDefaults

`func NewActionDryRunReadDryRunActionsInnerWithDefaults() *ActionDryRunReadDryRunActionsInner`

NewActionDryRunReadDryRunActionsInnerWithDefaults instantiates a new ActionDryRunReadDryRunActionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActionDryRunReadDryRunActionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionDryRunReadDryRunActionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionDryRunReadDryRunActionsInner) SetType(v string)`

SetType sets Type field to given value.


### GetTargetEntity

`func (o *ActionDryRunReadDryRunActionsInner) GetTargetEntity() ActionDryRunReadDryRunActionsInnerAnyOfTargetEntity`

GetTargetEntity returns the TargetEntity field if non-nil, zero value otherwise.

### GetTargetEntityOk

`func (o *ActionDryRunReadDryRunActionsInner) GetTargetEntityOk() (*ActionDryRunReadDryRunActionsInnerAnyOfTargetEntity, bool)`

GetTargetEntityOk returns a tuple with the TargetEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntity

`func (o *ActionDryRunReadDryRunActionsInner) SetTargetEntity(v ActionDryRunReadDryRunActionsInnerAnyOfTargetEntity)`

SetTargetEntity sets TargetEntity field to given value.


### GetStatus

`func (o *ActionDryRunReadDryRunActionsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionDryRunReadDryRunActionsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionDryRunReadDryRunActionsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPolicyDecision

`func (o *ActionDryRunReadDryRunActionsInner) GetPolicyDecision() string`

GetPolicyDecision returns the PolicyDecision field if non-nil, zero value otherwise.

### GetPolicyDecisionOk

`func (o *ActionDryRunReadDryRunActionsInner) GetPolicyDecisionOk() (*string, bool)`

GetPolicyDecisionOk returns a tuple with the PolicyDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDecision

`func (o *ActionDryRunReadDryRunActionsInner) SetPolicyDecision(v string)`

SetPolicyDecision sets PolicyDecision field to given value.


### GetPolicyReasons

`func (o *ActionDryRunReadDryRunActionsInner) GetPolicyReasons() []string`

GetPolicyReasons returns the PolicyReasons field if non-nil, zero value otherwise.

### GetPolicyReasonsOk

`func (o *ActionDryRunReadDryRunActionsInner) GetPolicyReasonsOk() (*[]string, bool)`

GetPolicyReasonsOk returns a tuple with the PolicyReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyReasons

`func (o *ActionDryRunReadDryRunActionsInner) SetPolicyReasons(v []string)`

SetPolicyReasons sets PolicyReasons field to given value.


### GetApprovalRequired

`func (o *ActionDryRunReadDryRunActionsInner) GetApprovalRequired() bool`

GetApprovalRequired returns the ApprovalRequired field if non-nil, zero value otherwise.

### GetApprovalRequiredOk

`func (o *ActionDryRunReadDryRunActionsInner) GetApprovalRequiredOk() (*bool, bool)`

GetApprovalRequiredOk returns a tuple with the ApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequired

`func (o *ActionDryRunReadDryRunActionsInner) SetApprovalRequired(v bool)`

SetApprovalRequired sets ApprovalRequired field to given value.


### GetWouldAutoExecute

`func (o *ActionDryRunReadDryRunActionsInner) GetWouldAutoExecute() bool`

GetWouldAutoExecute returns the WouldAutoExecute field if non-nil, zero value otherwise.

### GetWouldAutoExecuteOk

`func (o *ActionDryRunReadDryRunActionsInner) GetWouldAutoExecuteOk() (*bool, bool)`

GetWouldAutoExecuteOk returns a tuple with the WouldAutoExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWouldAutoExecute

`func (o *ActionDryRunReadDryRunActionsInner) SetWouldAutoExecute(v bool)`

SetWouldAutoExecute sets WouldAutoExecute field to given value.


### GetDiff

`func (o *ActionDryRunReadDryRunActionsInner) GetDiff() ActionDryRunReadDryRunActionsInnerAnyOfDiff`

GetDiff returns the Diff field if non-nil, zero value otherwise.

### GetDiffOk

`func (o *ActionDryRunReadDryRunActionsInner) GetDiffOk() (*ActionDryRunReadDryRunActionsInnerAnyOfDiff, bool)`

GetDiffOk returns a tuple with the Diff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiff

`func (o *ActionDryRunReadDryRunActionsInner) SetDiff(v ActionDryRunReadDryRunActionsInnerAnyOfDiff)`

SetDiff sets Diff field to given value.


### GetReason

`func (o *ActionDryRunReadDryRunActionsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ActionDryRunReadDryRunActionsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ActionDryRunReadDryRunActionsInner) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


