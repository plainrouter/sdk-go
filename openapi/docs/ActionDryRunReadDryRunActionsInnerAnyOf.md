# ActionDryRunReadDryRunActionsInnerAnyOf

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

## Methods

### NewActionDryRunReadDryRunActionsInnerAnyOf

`func NewActionDryRunReadDryRunActionsInnerAnyOf(type_ string, targetEntity ActionDryRunReadDryRunActionsInnerAnyOfTargetEntity, status string, policyDecision string, policyReasons []string, approvalRequired bool, wouldAutoExecute bool, diff ActionDryRunReadDryRunActionsInnerAnyOfDiff, ) *ActionDryRunReadDryRunActionsInnerAnyOf`

NewActionDryRunReadDryRunActionsInnerAnyOf instantiates a new ActionDryRunReadDryRunActionsInnerAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionDryRunReadDryRunActionsInnerAnyOfWithDefaults

`func NewActionDryRunReadDryRunActionsInnerAnyOfWithDefaults() *ActionDryRunReadDryRunActionsInnerAnyOf`

NewActionDryRunReadDryRunActionsInnerAnyOfWithDefaults instantiates a new ActionDryRunReadDryRunActionsInnerAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) SetType(v string)`

SetType sets Type field to given value.


### GetTargetEntity

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetTargetEntity() ActionDryRunReadDryRunActionsInnerAnyOfTargetEntity`

GetTargetEntity returns the TargetEntity field if non-nil, zero value otherwise.

### GetTargetEntityOk

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetTargetEntityOk() (*ActionDryRunReadDryRunActionsInnerAnyOfTargetEntity, bool)`

GetTargetEntityOk returns a tuple with the TargetEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntity

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) SetTargetEntity(v ActionDryRunReadDryRunActionsInnerAnyOfTargetEntity)`

SetTargetEntity sets TargetEntity field to given value.


### GetStatus

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPolicyDecision

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetPolicyDecision() string`

GetPolicyDecision returns the PolicyDecision field if non-nil, zero value otherwise.

### GetPolicyDecisionOk

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetPolicyDecisionOk() (*string, bool)`

GetPolicyDecisionOk returns a tuple with the PolicyDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDecision

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) SetPolicyDecision(v string)`

SetPolicyDecision sets PolicyDecision field to given value.


### GetPolicyReasons

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetPolicyReasons() []string`

GetPolicyReasons returns the PolicyReasons field if non-nil, zero value otherwise.

### GetPolicyReasonsOk

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetPolicyReasonsOk() (*[]string, bool)`

GetPolicyReasonsOk returns a tuple with the PolicyReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyReasons

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) SetPolicyReasons(v []string)`

SetPolicyReasons sets PolicyReasons field to given value.


### GetApprovalRequired

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetApprovalRequired() bool`

GetApprovalRequired returns the ApprovalRequired field if non-nil, zero value otherwise.

### GetApprovalRequiredOk

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetApprovalRequiredOk() (*bool, bool)`

GetApprovalRequiredOk returns a tuple with the ApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequired

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) SetApprovalRequired(v bool)`

SetApprovalRequired sets ApprovalRequired field to given value.


### GetWouldAutoExecute

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetWouldAutoExecute() bool`

GetWouldAutoExecute returns the WouldAutoExecute field if non-nil, zero value otherwise.

### GetWouldAutoExecuteOk

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetWouldAutoExecuteOk() (*bool, bool)`

GetWouldAutoExecuteOk returns a tuple with the WouldAutoExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWouldAutoExecute

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) SetWouldAutoExecute(v bool)`

SetWouldAutoExecute sets WouldAutoExecute field to given value.


### GetDiff

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetDiff() ActionDryRunReadDryRunActionsInnerAnyOfDiff`

GetDiff returns the Diff field if non-nil, zero value otherwise.

### GetDiffOk

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) GetDiffOk() (*ActionDryRunReadDryRunActionsInnerAnyOfDiff, bool)`

GetDiffOk returns a tuple with the Diff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiff

`func (o *ActionDryRunReadDryRunActionsInnerAnyOf) SetDiff(v ActionDryRunReadDryRunActionsInnerAnyOfDiff)`

SetDiff sets Diff field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


