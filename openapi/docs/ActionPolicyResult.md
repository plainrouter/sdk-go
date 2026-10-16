# ActionPolicyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PolicyId** | **string** |  | 
**PolicyVersion** | **int32** |  | 
**Phase** | **string** |  | 
**Outcome** | **string** |  | 
**Reasons** | **[]string** |  | 
**ReasonDetails** | **[]string** |  | 
**RequestedMinor** | **NullableInt32** |  | 
**EvaluatedAt** | **time.Time** |  | 

## Methods

### NewActionPolicyResult

`func NewActionPolicyResult(id string, policyId string, policyVersion int32, phase string, outcome string, reasons []string, reasonDetails []string, requestedMinor NullableInt32, evaluatedAt time.Time, ) *ActionPolicyResult`

NewActionPolicyResult instantiates a new ActionPolicyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionPolicyResultWithDefaults

`func NewActionPolicyResultWithDefaults() *ActionPolicyResult`

NewActionPolicyResultWithDefaults instantiates a new ActionPolicyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionPolicyResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionPolicyResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionPolicyResult) SetId(v string)`

SetId sets Id field to given value.


### GetPolicyId

`func (o *ActionPolicyResult) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ActionPolicyResult) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ActionPolicyResult) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.


### GetPolicyVersion

`func (o *ActionPolicyResult) GetPolicyVersion() int32`

GetPolicyVersion returns the PolicyVersion field if non-nil, zero value otherwise.

### GetPolicyVersionOk

`func (o *ActionPolicyResult) GetPolicyVersionOk() (*int32, bool)`

GetPolicyVersionOk returns a tuple with the PolicyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyVersion

`func (o *ActionPolicyResult) SetPolicyVersion(v int32)`

SetPolicyVersion sets PolicyVersion field to given value.


### GetPhase

`func (o *ActionPolicyResult) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ActionPolicyResult) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ActionPolicyResult) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetOutcome

`func (o *ActionPolicyResult) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *ActionPolicyResult) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *ActionPolicyResult) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetReasons

`func (o *ActionPolicyResult) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *ActionPolicyResult) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *ActionPolicyResult) SetReasons(v []string)`

SetReasons sets Reasons field to given value.


### GetReasonDetails

`func (o *ActionPolicyResult) GetReasonDetails() []string`

GetReasonDetails returns the ReasonDetails field if non-nil, zero value otherwise.

### GetReasonDetailsOk

`func (o *ActionPolicyResult) GetReasonDetailsOk() (*[]string, bool)`

GetReasonDetailsOk returns a tuple with the ReasonDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonDetails

`func (o *ActionPolicyResult) SetReasonDetails(v []string)`

SetReasonDetails sets ReasonDetails field to given value.


### SetReasonDetailsNil

`func (o *ActionPolicyResult) SetReasonDetailsNil(b bool)`

 SetReasonDetailsNil sets the value for ReasonDetails to be an explicit nil

### UnsetReasonDetails
`func (o *ActionPolicyResult) UnsetReasonDetails()`

UnsetReasonDetails ensures that no value is present for ReasonDetails, not even an explicit nil
### GetRequestedMinor

`func (o *ActionPolicyResult) GetRequestedMinor() int32`

GetRequestedMinor returns the RequestedMinor field if non-nil, zero value otherwise.

### GetRequestedMinorOk

`func (o *ActionPolicyResult) GetRequestedMinorOk() (*int32, bool)`

GetRequestedMinorOk returns a tuple with the RequestedMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedMinor

`func (o *ActionPolicyResult) SetRequestedMinor(v int32)`

SetRequestedMinor sets RequestedMinor field to given value.


### SetRequestedMinorNil

`func (o *ActionPolicyResult) SetRequestedMinorNil(b bool)`

 SetRequestedMinorNil sets the value for RequestedMinor to be an explicit nil

### UnsetRequestedMinor
`func (o *ActionPolicyResult) UnsetRequestedMinor()`

UnsetRequestedMinor ensures that no value is present for RequestedMinor, not even an explicit nil
### GetEvaluatedAt

`func (o *ActionPolicyResult) GetEvaluatedAt() time.Time`

GetEvaluatedAt returns the EvaluatedAt field if non-nil, zero value otherwise.

### GetEvaluatedAtOk

`func (o *ActionPolicyResult) GetEvaluatedAtOk() (*time.Time, bool)`

GetEvaluatedAtOk returns a tuple with the EvaluatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedAt

`func (o *ActionPolicyResult) SetEvaluatedAt(v time.Time)`

SetEvaluatedAt sets EvaluatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


