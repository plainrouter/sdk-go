# ExecuteDeploymentPlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReviewVersion** | **string** |  | 
**IntentKey** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExecuteDeploymentPlanRequest

`func NewExecuteDeploymentPlanRequest(reviewVersion string, ) *ExecuteDeploymentPlanRequest`

NewExecuteDeploymentPlanRequest instantiates a new ExecuteDeploymentPlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteDeploymentPlanRequestWithDefaults

`func NewExecuteDeploymentPlanRequestWithDefaults() *ExecuteDeploymentPlanRequest`

NewExecuteDeploymentPlanRequestWithDefaults instantiates a new ExecuteDeploymentPlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewVersion

`func (o *ExecuteDeploymentPlanRequest) GetReviewVersion() string`

GetReviewVersion returns the ReviewVersion field if non-nil, zero value otherwise.

### GetReviewVersionOk

`func (o *ExecuteDeploymentPlanRequest) GetReviewVersionOk() (*string, bool)`

GetReviewVersionOk returns a tuple with the ReviewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewVersion

`func (o *ExecuteDeploymentPlanRequest) SetReviewVersion(v string)`

SetReviewVersion sets ReviewVersion field to given value.


### GetIntentKey

`func (o *ExecuteDeploymentPlanRequest) GetIntentKey() string`

GetIntentKey returns the IntentKey field if non-nil, zero value otherwise.

### GetIntentKeyOk

`func (o *ExecuteDeploymentPlanRequest) GetIntentKeyOk() (*string, bool)`

GetIntentKeyOk returns a tuple with the IntentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentKey

`func (o *ExecuteDeploymentPlanRequest) SetIntentKey(v string)`

SetIntentKey sets IntentKey field to given value.

### HasIntentKey

`func (o *ExecuteDeploymentPlanRequest) HasIntentKey() bool`

HasIntentKey returns a boolean if a field has been set.

### SetIntentKeyNil

`func (o *ExecuteDeploymentPlanRequest) SetIntentKeyNil(b bool)`

 SetIntentKeyNil sets the value for IntentKey to be an explicit nil

### UnsetIntentKey
`func (o *ExecuteDeploymentPlanRequest) UnsetIntentKey()`

UnsetIntentKey ensures that no value is present for IntentKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


