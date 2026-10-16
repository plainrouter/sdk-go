# LaunchPlansCopy422Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**PlanCopyRejectedError**](PlanCopyRejectedError.md) |  | 
**Message** | **string** |  | 
**Errors** | **map[string][]string** |  | 

## Methods

### NewLaunchPlansCopy422Response

`func NewLaunchPlansCopy422Response(error_ PlanCopyRejectedError, message string, errors map[string][]string, ) *LaunchPlansCopy422Response`

NewLaunchPlansCopy422Response instantiates a new LaunchPlansCopy422Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchPlansCopy422ResponseWithDefaults

`func NewLaunchPlansCopy422ResponseWithDefaults() *LaunchPlansCopy422Response`

NewLaunchPlansCopy422ResponseWithDefaults instantiates a new LaunchPlansCopy422Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *LaunchPlansCopy422Response) GetError() PlanCopyRejectedError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LaunchPlansCopy422Response) GetErrorOk() (*PlanCopyRejectedError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *LaunchPlansCopy422Response) SetError(v PlanCopyRejectedError)`

SetError sets Error field to given value.


### GetMessage

`func (o *LaunchPlansCopy422Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LaunchPlansCopy422Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LaunchPlansCopy422Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *LaunchPlansCopy422Response) GetErrors() map[string][]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LaunchPlansCopy422Response) GetErrorsOk() (*map[string][]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LaunchPlansCopy422Response) SetErrors(v map[string][]string)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


