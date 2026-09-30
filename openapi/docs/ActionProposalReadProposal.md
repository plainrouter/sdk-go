# ActionProposalReadProposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Scope** | **interface{}** |  | 
**EvidenceProvenance** | **interface{}** |  | 
**Status** | **string** |  | 
**PolicyDecision** | **string** |  | 
**PolicyReasons** | **[]string** |  | 
**ApprovalRequired** | **bool** |  | 
**InboxUrl** | **string** |  | 
**ApprovalQueueUrl** | **string** |  | 
**Rationale** | **string** |  | 
**ProposedBy** | [**ActionProposalReadProposalProposedBy**](ActionProposalReadProposalProposedBy.md) |  | 
**Actions** | [**[]ActionProposalReadProposalActionsInner**](ActionProposalReadProposalActionsInner.md) |  | 

## Methods

### NewActionProposalReadProposal

`func NewActionProposalReadProposal(id string, scope interface{}, evidenceProvenance interface{}, status string, policyDecision string, policyReasons []string, approvalRequired bool, inboxUrl string, approvalQueueUrl string, rationale string, proposedBy ActionProposalReadProposalProposedBy, actions []ActionProposalReadProposalActionsInner, ) *ActionProposalReadProposal`

NewActionProposalReadProposal instantiates a new ActionProposalReadProposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionProposalReadProposalWithDefaults

`func NewActionProposalReadProposalWithDefaults() *ActionProposalReadProposal`

NewActionProposalReadProposalWithDefaults instantiates a new ActionProposalReadProposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionProposalReadProposal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionProposalReadProposal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionProposalReadProposal) SetId(v string)`

SetId sets Id field to given value.


### GetScope

`func (o *ActionProposalReadProposal) GetScope() interface{}`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ActionProposalReadProposal) GetScopeOk() (*interface{}, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ActionProposalReadProposal) SetScope(v interface{})`

SetScope sets Scope field to given value.


### SetScopeNil

`func (o *ActionProposalReadProposal) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *ActionProposalReadProposal) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetEvidenceProvenance

`func (o *ActionProposalReadProposal) GetEvidenceProvenance() interface{}`

GetEvidenceProvenance returns the EvidenceProvenance field if non-nil, zero value otherwise.

### GetEvidenceProvenanceOk

`func (o *ActionProposalReadProposal) GetEvidenceProvenanceOk() (*interface{}, bool)`

GetEvidenceProvenanceOk returns a tuple with the EvidenceProvenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceProvenance

`func (o *ActionProposalReadProposal) SetEvidenceProvenance(v interface{})`

SetEvidenceProvenance sets EvidenceProvenance field to given value.


### SetEvidenceProvenanceNil

`func (o *ActionProposalReadProposal) SetEvidenceProvenanceNil(b bool)`

 SetEvidenceProvenanceNil sets the value for EvidenceProvenance to be an explicit nil

### UnsetEvidenceProvenance
`func (o *ActionProposalReadProposal) UnsetEvidenceProvenance()`

UnsetEvidenceProvenance ensures that no value is present for EvidenceProvenance, not even an explicit nil
### GetStatus

`func (o *ActionProposalReadProposal) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionProposalReadProposal) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionProposalReadProposal) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPolicyDecision

`func (o *ActionProposalReadProposal) GetPolicyDecision() string`

GetPolicyDecision returns the PolicyDecision field if non-nil, zero value otherwise.

### GetPolicyDecisionOk

`func (o *ActionProposalReadProposal) GetPolicyDecisionOk() (*string, bool)`

GetPolicyDecisionOk returns a tuple with the PolicyDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDecision

`func (o *ActionProposalReadProposal) SetPolicyDecision(v string)`

SetPolicyDecision sets PolicyDecision field to given value.


### GetPolicyReasons

`func (o *ActionProposalReadProposal) GetPolicyReasons() []string`

GetPolicyReasons returns the PolicyReasons field if non-nil, zero value otherwise.

### GetPolicyReasonsOk

`func (o *ActionProposalReadProposal) GetPolicyReasonsOk() (*[]string, bool)`

GetPolicyReasonsOk returns a tuple with the PolicyReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyReasons

`func (o *ActionProposalReadProposal) SetPolicyReasons(v []string)`

SetPolicyReasons sets PolicyReasons field to given value.


### GetApprovalRequired

`func (o *ActionProposalReadProposal) GetApprovalRequired() bool`

GetApprovalRequired returns the ApprovalRequired field if non-nil, zero value otherwise.

### GetApprovalRequiredOk

`func (o *ActionProposalReadProposal) GetApprovalRequiredOk() (*bool, bool)`

GetApprovalRequiredOk returns a tuple with the ApprovalRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequired

`func (o *ActionProposalReadProposal) SetApprovalRequired(v bool)`

SetApprovalRequired sets ApprovalRequired field to given value.


### GetInboxUrl

`func (o *ActionProposalReadProposal) GetInboxUrl() string`

GetInboxUrl returns the InboxUrl field if non-nil, zero value otherwise.

### GetInboxUrlOk

`func (o *ActionProposalReadProposal) GetInboxUrlOk() (*string, bool)`

GetInboxUrlOk returns a tuple with the InboxUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboxUrl

`func (o *ActionProposalReadProposal) SetInboxUrl(v string)`

SetInboxUrl sets InboxUrl field to given value.


### GetApprovalQueueUrl

`func (o *ActionProposalReadProposal) GetApprovalQueueUrl() string`

GetApprovalQueueUrl returns the ApprovalQueueUrl field if non-nil, zero value otherwise.

### GetApprovalQueueUrlOk

`func (o *ActionProposalReadProposal) GetApprovalQueueUrlOk() (*string, bool)`

GetApprovalQueueUrlOk returns a tuple with the ApprovalQueueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalQueueUrl

`func (o *ActionProposalReadProposal) SetApprovalQueueUrl(v string)`

SetApprovalQueueUrl sets ApprovalQueueUrl field to given value.


### GetRationale

`func (o *ActionProposalReadProposal) GetRationale() string`

GetRationale returns the Rationale field if non-nil, zero value otherwise.

### GetRationaleOk

`func (o *ActionProposalReadProposal) GetRationaleOk() (*string, bool)`

GetRationaleOk returns a tuple with the Rationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationale

`func (o *ActionProposalReadProposal) SetRationale(v string)`

SetRationale sets Rationale field to given value.


### GetProposedBy

`func (o *ActionProposalReadProposal) GetProposedBy() ActionProposalReadProposalProposedBy`

GetProposedBy returns the ProposedBy field if non-nil, zero value otherwise.

### GetProposedByOk

`func (o *ActionProposalReadProposal) GetProposedByOk() (*ActionProposalReadProposalProposedBy, bool)`

GetProposedByOk returns a tuple with the ProposedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedBy

`func (o *ActionProposalReadProposal) SetProposedBy(v ActionProposalReadProposalProposedBy)`

SetProposedBy sets ProposedBy field to given value.


### GetActions

`func (o *ActionProposalReadProposal) GetActions() []ActionProposalReadProposalActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ActionProposalReadProposal) GetActionsOk() (*[]ActionProposalReadProposalActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ActionProposalReadProposal) SetActions(v []ActionProposalReadProposalActionsInner)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


