# ActionReadItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**BatchId** | **string** |  | 
**WorkspaceId** | **int32** |  | 
**Type** | **string** |  | 
**TargetEntityType** | **string** |  | 
**TargetEntityId** | **string** |  | 
**TargetEntityName** | **NullableString** |  | 
**Params** | **interface{}** |  | 
**Rationale** | **string** |  | 
**Status** | **string** |  | 
**BatchStatus** | **string** |  | 
**Disposition** | [**ActionCurrentDisposition**](ActionCurrentDisposition.md) |  | 
**PolicyDecision** | **NullableString** |  | 
**PolicyReasons** | **[]string** |  | 

## Methods

### NewActionReadItem

`func NewActionReadItem(id string, batchId string, workspaceId int32, type_ string, targetEntityType string, targetEntityId string, targetEntityName NullableString, params interface{}, rationale string, status string, batchStatus string, disposition ActionCurrentDisposition, policyDecision NullableString, policyReasons []string, ) *ActionReadItem`

NewActionReadItem instantiates a new ActionReadItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionReadItemWithDefaults

`func NewActionReadItemWithDefaults() *ActionReadItem`

NewActionReadItemWithDefaults instantiates a new ActionReadItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionReadItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionReadItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionReadItem) SetId(v string)`

SetId sets Id field to given value.


### GetBatchId

`func (o *ActionReadItem) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *ActionReadItem) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *ActionReadItem) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.


### GetWorkspaceId

`func (o *ActionReadItem) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ActionReadItem) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ActionReadItem) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetType

`func (o *ActionReadItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionReadItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionReadItem) SetType(v string)`

SetType sets Type field to given value.


### GetTargetEntityType

`func (o *ActionReadItem) GetTargetEntityType() string`

GetTargetEntityType returns the TargetEntityType field if non-nil, zero value otherwise.

### GetTargetEntityTypeOk

`func (o *ActionReadItem) GetTargetEntityTypeOk() (*string, bool)`

GetTargetEntityTypeOk returns a tuple with the TargetEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntityType

`func (o *ActionReadItem) SetTargetEntityType(v string)`

SetTargetEntityType sets TargetEntityType field to given value.


### GetTargetEntityId

`func (o *ActionReadItem) GetTargetEntityId() string`

GetTargetEntityId returns the TargetEntityId field if non-nil, zero value otherwise.

### GetTargetEntityIdOk

`func (o *ActionReadItem) GetTargetEntityIdOk() (*string, bool)`

GetTargetEntityIdOk returns a tuple with the TargetEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntityId

`func (o *ActionReadItem) SetTargetEntityId(v string)`

SetTargetEntityId sets TargetEntityId field to given value.


### GetTargetEntityName

`func (o *ActionReadItem) GetTargetEntityName() string`

GetTargetEntityName returns the TargetEntityName field if non-nil, zero value otherwise.

### GetTargetEntityNameOk

`func (o *ActionReadItem) GetTargetEntityNameOk() (*string, bool)`

GetTargetEntityNameOk returns a tuple with the TargetEntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntityName

`func (o *ActionReadItem) SetTargetEntityName(v string)`

SetTargetEntityName sets TargetEntityName field to given value.


### SetTargetEntityNameNil

`func (o *ActionReadItem) SetTargetEntityNameNil(b bool)`

 SetTargetEntityNameNil sets the value for TargetEntityName to be an explicit nil

### UnsetTargetEntityName
`func (o *ActionReadItem) UnsetTargetEntityName()`

UnsetTargetEntityName ensures that no value is present for TargetEntityName, not even an explicit nil
### GetParams

`func (o *ActionReadItem) GetParams() interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ActionReadItem) GetParamsOk() (*interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ActionReadItem) SetParams(v interface{})`

SetParams sets Params field to given value.


### SetParamsNil

`func (o *ActionReadItem) SetParamsNil(b bool)`

 SetParamsNil sets the value for Params to be an explicit nil

### UnsetParams
`func (o *ActionReadItem) UnsetParams()`

UnsetParams ensures that no value is present for Params, not even an explicit nil
### GetRationale

`func (o *ActionReadItem) GetRationale() string`

GetRationale returns the Rationale field if non-nil, zero value otherwise.

### GetRationaleOk

`func (o *ActionReadItem) GetRationaleOk() (*string, bool)`

GetRationaleOk returns a tuple with the Rationale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRationale

`func (o *ActionReadItem) SetRationale(v string)`

SetRationale sets Rationale field to given value.


### GetStatus

`func (o *ActionReadItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionReadItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionReadItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBatchStatus

`func (o *ActionReadItem) GetBatchStatus() string`

GetBatchStatus returns the BatchStatus field if non-nil, zero value otherwise.

### GetBatchStatusOk

`func (o *ActionReadItem) GetBatchStatusOk() (*string, bool)`

GetBatchStatusOk returns a tuple with the BatchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchStatus

`func (o *ActionReadItem) SetBatchStatus(v string)`

SetBatchStatus sets BatchStatus field to given value.


### GetDisposition

`func (o *ActionReadItem) GetDisposition() ActionCurrentDisposition`

GetDisposition returns the Disposition field if non-nil, zero value otherwise.

### GetDispositionOk

`func (o *ActionReadItem) GetDispositionOk() (*ActionCurrentDisposition, bool)`

GetDispositionOk returns a tuple with the Disposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposition

`func (o *ActionReadItem) SetDisposition(v ActionCurrentDisposition)`

SetDisposition sets Disposition field to given value.


### GetPolicyDecision

`func (o *ActionReadItem) GetPolicyDecision() string`

GetPolicyDecision returns the PolicyDecision field if non-nil, zero value otherwise.

### GetPolicyDecisionOk

`func (o *ActionReadItem) GetPolicyDecisionOk() (*string, bool)`

GetPolicyDecisionOk returns a tuple with the PolicyDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDecision

`func (o *ActionReadItem) SetPolicyDecision(v string)`

SetPolicyDecision sets PolicyDecision field to given value.


### SetPolicyDecisionNil

`func (o *ActionReadItem) SetPolicyDecisionNil(b bool)`

 SetPolicyDecisionNil sets the value for PolicyDecision to be an explicit nil

### UnsetPolicyDecision
`func (o *ActionReadItem) UnsetPolicyDecision()`

UnsetPolicyDecision ensures that no value is present for PolicyDecision, not even an explicit nil
### GetPolicyReasons

`func (o *ActionReadItem) GetPolicyReasons() []string`

GetPolicyReasons returns the PolicyReasons field if non-nil, zero value otherwise.

### GetPolicyReasonsOk

`func (o *ActionReadItem) GetPolicyReasonsOk() (*[]string, bool)`

GetPolicyReasonsOk returns a tuple with the PolicyReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyReasons

`func (o *ActionReadItem) SetPolicyReasons(v []string)`

SetPolicyReasons sets PolicyReasons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


