# ActionBatchReadData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**WorkspaceId** | **int32** |  | 
**PlatformAdAccountId** | **int32** |  | 
**Status** | **string** |  | 
**BatchStatus** | **string** |  | 
**RestorationSummary** | **NullableString** |  | 
**PolicyDecision** | **NullableString** |  | 
**PolicyReasons** | **[]string** |  | 
**Rationale** | **string** |  | 
**IdempotencyKey** | **string** |  | 
**Actions** | [**[]ActionReadItem**](ActionReadItem.md) |  | 

## Methods

### NewActionBatchReadData

`func NewActionBatchReadData(id string, workspaceId int32, platformAdAccountId int32, status string, batchStatus string, restorationSummary NullableString, policyDecision NullableString, policyReasons []string, rationale string, idempotencyKey string, actions []ActionReadItem, ) *ActionBatchReadData`

NewActionBatchReadData instantiates a new ActionBatchReadData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionBatchReadDataWithDefaults

`func NewActionBatchReadDataWithDefaults() *ActionBatchReadData`

NewActionBatchReadDataWithDefaults instantiates a new ActionBatchReadData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionBatchReadData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionBatchReadData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionBatchReadData) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *ActionBatchReadData) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ActionBatchReadData) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ActionBatchReadData) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetPlatformAdAccountId

`func (o *ActionBatchReadData) GetPlatformAdAccountId() int32`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *ActionBatchReadData) GetPlatformAdAccountIdOk() (*int32, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *ActionBatchReadData) SetPlatformAdAccountId(v int32)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.


### GetStatus

`func (o *ActionBatchReadData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionBatchReadData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionBatchReadData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBatchStatus

`func (o *ActionBatchReadData) GetBatchStatus() string`

GetBatchStatus returns the BatchStatus field if non-nil, zero value otherwise.

### GetBatchStatusOk

`func (o *ActionBatchReadData) GetBatchStatusOk() (*string, bool)`

GetBatchStatusOk returns a tuple with the BatchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchStatus

`func (o *ActionBatchReadData) SetBatchStatus(v string)`

SetBatchStatus sets BatchStatus field to given value.


### GetRestorationSummary

`func (o *ActionBatchReadData) GetRestorationSummary() string`

GetRestorationSummary returns the RestorationSummary field if non-nil, zero value otherwise.

### GetRestorationSummaryOk

`func (o *ActionBatchReadData) GetRestorationSummaryOk() (*string, bool)`

GetRestorationSummaryOk returns a tuple with the RestorationSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorationSummary

`func (o *ActionBatchReadData) SetRestorationSummary(v string)`

SetRestorationSummary sets RestorationSummary field to given value.


### SetRestorationSummaryNil

`func (o *ActionBatchReadData) SetRestorationSummaryNil(b bool)`

 SetRestorationSummaryNil sets the value for RestorationSummary to be an explicit nil

### UnsetRestorationSummary
`func (o *ActionBatchReadData) UnsetRestorationSummary()`

UnsetRestorationSummary ensures that no value is present for RestorationSummary, not even an explicit nil
### GetPolicyDecision

`func (o *ActionBatchReadData) GetPolicyDecision() string`

GetPolicyDecision returns the PolicyDecision field if non-nil, zero value otherwise.

### GetPolicyDecisionOk

`func (o *ActionBatchReadData) GetPolicyDecisionOk() (*string, bool)`

GetPolicyDecisionOk returns a tuple with the PolicyDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDecision

`func (o *ActionBatchReadData) SetPolicyDecision(v string)`

SetPolicyDecision sets PolicyDecision field to given value.


### SetPolicyDecisionNil

`func (o *ActionBatchReadData) SetPolicyDecisionNil(b bool)`

 SetPolicyDecisionNil sets the value for PolicyDecision to be an explicit nil

### UnsetPolicyDecision
`func (o *ActionBatchReadData) UnsetPolicyDecision()`

UnsetPolicyDecision ensures that no value is present for PolicyDecision, not even an explicit nil
### GetPolicyReasons

`func (o *ActionBatchReadData) GetPolicyReasons() []string`

GetPolicyReasons returns the PolicyReasons field if non-nil, zero value otherwise.

### GetPolicyReasonsOk

`func (o *ActionBatchReadData) GetPolicyReasonsOk() (*[]string, bool)`

GetPolicyReasonsOk returns a tuple with the PolicyReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyReasons

`func (o *ActionBatchReadData) SetPolicyReasons(v []string)`

SetPolicyReasons sets PolicyReasons field to given value.


### GetRationale

`func (o *ActionBatchReadData) GetRationale() string`

GetRationale returns the Rationale field if non-nil, zero value otherwise.

### GetRationaleOk

`func (o *ActionBatchReadData) GetRationaleOk() (*string, bool)`

GetRationaleOk returns a tuple with the Rationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationale

`func (o *ActionBatchReadData) SetRationale(v string)`

SetRationale sets Rationale field to given value.


### GetIdempotencyKey

`func (o *ActionBatchReadData) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *ActionBatchReadData) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *ActionBatchReadData) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### GetActions

`func (o *ActionBatchReadData) GetActions() []ActionReadItem`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ActionBatchReadData) GetActionsOk() (*[]ActionReadItem, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ActionBatchReadData) SetActions(v []ActionReadItem)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


