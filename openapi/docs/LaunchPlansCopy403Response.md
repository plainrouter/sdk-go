# LaunchPlansCopy403Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**PlanCopyForbiddenError**](PlanCopyForbiddenError.md) |  | 
**Message** | **string** |  | 

## Methods

### NewLaunchPlansCopy403Response

`func NewLaunchPlansCopy403Response(error_ PlanCopyForbiddenError, message string, ) *LaunchPlansCopy403Response`

NewLaunchPlansCopy403Response instantiates a new LaunchPlansCopy403Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchPlansCopy403ResponseWithDefaults

`func NewLaunchPlansCopy403ResponseWithDefaults() *LaunchPlansCopy403Response`

NewLaunchPlansCopy403ResponseWithDefaults instantiates a new LaunchPlansCopy403Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *LaunchPlansCopy403Response) GetError() PlanCopyForbiddenError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LaunchPlansCopy403Response) GetErrorOk() (*PlanCopyForbiddenError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *LaunchPlansCopy403Response) SetError(v PlanCopyForbiddenError)`

SetError sets Error field to given value.


### GetMessage

`func (o *LaunchPlansCopy403Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LaunchPlansCopy403Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LaunchPlansCopy403Response) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


