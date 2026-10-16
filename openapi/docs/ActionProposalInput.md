# ActionProposalInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Maximum complete result actions; the shared JSON byte budget can return fewer. | [optional] [default to 25]
**Cursor** | Pointer to **int32** | Continue using next_cursor. Read saved proposal actions with get_action_batch to avoid another write; dry runs repeat the same inputs. | [optional] 
**Actions** | [**[]ActionProposalInputActionsInner**](ActionProposalInputActionsInner.md) |  | 
**Rationale** | **string** |  | 
**IdempotencyKey** | **string** |  | 
**Evidence** | [**[]ActionProposalInputEvidenceInner**](ActionProposalInputEvidenceInner.md) |  | 
**TargetSource** | **string** |  | 
**AccountId** | Pointer to **int32** |  | [optional] 
**WorkspaceId** | Pointer to **int32** |  | [optional] 

## Methods

### NewActionProposalInput

`func NewActionProposalInput(actions []ActionProposalInputActionsInner, rationale string, idempotencyKey string, evidence []ActionProposalInputEvidenceInner, targetSource string, ) *ActionProposalInput`

NewActionProposalInput instantiates a new ActionProposalInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionProposalInputWithDefaults

`func NewActionProposalInputWithDefaults() *ActionProposalInput`

NewActionProposalInputWithDefaults instantiates a new ActionProposalInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ActionProposalInput) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ActionProposalInput) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ActionProposalInput) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ActionProposalInput) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetCursor

`func (o *ActionProposalInput) GetCursor() int32`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ActionProposalInput) GetCursorOk() (*int32, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ActionProposalInput) SetCursor(v int32)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ActionProposalInput) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetActions

`func (o *ActionProposalInput) GetActions() []ActionProposalInputActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ActionProposalInput) GetActionsOk() (*[]ActionProposalInputActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ActionProposalInput) SetActions(v []ActionProposalInputActionsInner)`

SetActions sets Actions field to given value.


### GetRationale

`func (o *ActionProposalInput) GetRationale() string`

GetRationale returns the Rationale field if non-nil, zero value otherwise.

### GetRationaleOk

`func (o *ActionProposalInput) GetRationaleOk() (*string, bool)`

GetRationaleOk returns a tuple with the Rationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationale

`func (o *ActionProposalInput) SetRationale(v string)`

SetRationale sets Rationale field to given value.


### GetIdempotencyKey

`func (o *ActionProposalInput) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *ActionProposalInput) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *ActionProposalInput) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### GetEvidence

`func (o *ActionProposalInput) GetEvidence() []ActionProposalInputEvidenceInner`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *ActionProposalInput) GetEvidenceOk() (*[]ActionProposalInputEvidenceInner, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *ActionProposalInput) SetEvidence(v []ActionProposalInputEvidenceInner)`

SetEvidence sets Evidence field to given value.


### GetTargetSource

`func (o *ActionProposalInput) GetTargetSource() string`

GetTargetSource returns the TargetSource field if non-nil, zero value otherwise.

### GetTargetSourceOk

`func (o *ActionProposalInput) GetTargetSourceOk() (*string, bool)`

GetTargetSourceOk returns a tuple with the TargetSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSource

`func (o *ActionProposalInput) SetTargetSource(v string)`

SetTargetSource sets TargetSource field to given value.


### GetAccountId

`func (o *ActionProposalInput) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ActionProposalInput) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ActionProposalInput) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ActionProposalInput) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *ActionProposalInput) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ActionProposalInput) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ActionProposalInput) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *ActionProposalInput) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


