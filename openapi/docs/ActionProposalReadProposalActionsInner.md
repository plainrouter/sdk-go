# ActionProposalReadProposalActionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **string** |  | 
**TargetEntity** | [**ActionProposalReadProposalProposedBy**](ActionProposalReadProposalProposedBy.md) |  | 
**Params** | **interface{}** |  | 
**Rationale** | **string** |  | 
**ProposedBy** | [**ActionProposalReadProposalProposedBy**](ActionProposalReadProposalProposedBy.md) |  | 
**PolicyDecision** | **string** |  | 
**PolicyReasons** | **[]string** |  | 
**PolicyEvidence** | **interface{}** |  | 

## Methods

### NewActionProposalReadProposalActionsInner

`func NewActionProposalReadProposalActionsInner(id string, type_ string, targetEntity ActionProposalReadProposalProposedBy, params interface{}, rationale string, proposedBy ActionProposalReadProposalProposedBy, policyDecision string, policyReasons []string, policyEvidence interface{}, ) *ActionProposalReadProposalActionsInner`

NewActionProposalReadProposalActionsInner instantiates a new ActionProposalReadProposalActionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionProposalReadProposalActionsInnerWithDefaults

`func NewActionProposalReadProposalActionsInnerWithDefaults() *ActionProposalReadProposalActionsInner`

NewActionProposalReadProposalActionsInnerWithDefaults instantiates a new ActionProposalReadProposalActionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionProposalReadProposalActionsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionProposalReadProposalActionsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionProposalReadProposalActionsInner) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ActionProposalReadProposalActionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionProposalReadProposalActionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionProposalReadProposalActionsInner) SetType(v string)`

SetType sets Type field to given value.


### GetTargetEntity

`func (o *ActionProposalReadProposalActionsInner) GetTargetEntity() ActionProposalReadProposalProposedBy`

GetTargetEntity returns the TargetEntity field if non-nil, zero value otherwise.

### GetTargetEntityOk

`func (o *ActionProposalReadProposalActionsInner) GetTargetEntityOk() (*ActionProposalReadProposalProposedBy, bool)`

GetTargetEntityOk returns a tuple with the TargetEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntity

`func (o *ActionProposalReadProposalActionsInner) SetTargetEntity(v ActionProposalReadProposalProposedBy)`

SetTargetEntity sets TargetEntity field to given value.


### GetParams

`func (o *ActionProposalReadProposalActionsInner) GetParams() interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ActionProposalReadProposalActionsInner) GetParamsOk() (*interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ActionProposalReadProposalActionsInner) SetParams(v interface{})`

SetParams sets Params field to given value.


### SetParamsNil

`func (o *ActionProposalReadProposalActionsInner) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *ActionProposalReadProposalActionsInner) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil
### GetRationale

`func (o *ActionProposalReadProposalActionsInner) GetRationale() string`

GetRationale returns the Rationale field if non-nil, zero value otherwise.

### GetRationaleOk

`func (o *ActionProposalReadProposalActionsInner) GetRationaleOk() (*string, bool)`

GetRationaleOk returns a tuple with the Rationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationale

`func (o *ActionProposalReadProposalActionsInner) SetRationale(v string)`

SetRationale sets Rationale field to given value.


### GetProposedBy

`func (o *ActionProposalReadProposalActionsInner) GetProposedBy() ActionProposalReadProposalProposedBy`

GetProposedBy returns the ProposedBy field if non-nil, zero value otherwise.

### GetProposedByOk

`func (o *ActionProposalReadProposalActionsInner) GetProposedByOk() (*ActionProposalReadProposalProposedBy, bool)`

GetProposedByOk returns a tuple with the ProposedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedBy

`func (o *ActionProposalReadProposalActionsInner) SetProposedBy(v ActionProposalReadProposalProposedBy)`

SetProposedBy sets ProposedBy field to given value.


### GetPolicyDecision

`func (o *ActionProposalReadProposalActionsInner) GetPolicyDecision() string`

GetPolicyDecision returns the PolicyDecision field if non-nil, zero value otherwise.

### GetPolicyDecisionOk

`func (o *ActionProposalReadProposalActionsInner) GetPolicyDecisionOk() (*string, bool)`

GetPolicyDecisionOk returns a tuple with the PolicyDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDecision

`func (o *ActionProposalReadProposalActionsInner) SetPolicyDecision(v string)`

SetPolicyDecision sets PolicyDecision field to given value.


### GetPolicyReasons

`func (o *ActionProposalReadProposalActionsInner) GetPolicyReasons() []string`

GetPolicyReasons returns the PolicyReasons field if non-nil, zero value otherwise.

### GetPolicyReasonsOk

`func (o *ActionProposalReadProposalActionsInner) GetPolicyReasonsOk() (*[]string, bool)`

GetPolicyReasonsOk returns a tuple with the PolicyReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyReasons

`func (o *ActionProposalReadProposalActionsInner) SetPolicyReasons(v []string)`

SetPolicyReasons sets PolicyReasons field to given value.


### GetPolicyEvidence

`func (o *ActionProposalReadProposalActionsInner) GetPolicyEvidence() interface{}`

GetPolicyEvidence returns the PolicyEvidence field if non-nil, zero value otherwise.

### GetPolicyEvidenceOk

`func (o *ActionProposalReadProposalActionsInner) GetPolicyEvidenceOk() (*interface{}, bool)`

GetPolicyEvidenceOk returns a tuple with the PolicyEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEvidence

`func (o *ActionProposalReadProposalActionsInner) SetPolicyEvidence(v interface{})`

SetPolicyEvidence sets PolicyEvidence field to given value.


### SetPolicyEvidenceNil

`func (o *ActionProposalReadProposalActionsInner) SetPolicyEvidenceNil(b bool)`

 SetPolicyEvidenceNil sets the value for PolicyEvidence to be an explicit nil

### UnsetPolicyEvidence
`func (o *ActionProposalReadProposalActionsInner) UnsetPolicyEvidence()`

UnsetPolicyEvidence ensures that no value is present for PolicyEvidence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


