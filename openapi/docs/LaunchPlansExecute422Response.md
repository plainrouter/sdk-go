# LaunchPlansExecute422Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Errors** | **map[string][]string** |  | 
**Error** | [**PlanExecuteRejectedError**](PlanExecuteRejectedError.md) |  | 

## Methods

### NewLaunchPlansExecute422Response

`func NewLaunchPlansExecute422Response(message string, errors map[string][]string, error_ PlanExecuteRejectedError, ) *LaunchPlansExecute422Response`

NewLaunchPlansExecute422Response instantiates a new LaunchPlansExecute422Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchPlansExecute422ResponseWithDefaults

`func NewLaunchPlansExecute422ResponseWithDefaults() *LaunchPlansExecute422Response`

NewLaunchPlansExecute422ResponseWithDefaults instantiates a new LaunchPlansExecute422Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *LaunchPlansExecute422Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LaunchPlansExecute422Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LaunchPlansExecute422Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *LaunchPlansExecute422Response) GetErrors() map[string][]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LaunchPlansExecute422Response) GetErrorsOk() (*map[string][]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LaunchPlansExecute422Response) SetErrors(v map[string][]string)`

SetErrors sets Errors field to given value.


### GetError

`func (o *LaunchPlansExecute422Response) GetError() PlanExecuteRejectedError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LaunchPlansExecute422Response) GetErrorOk() (*PlanExecuteRejectedError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *LaunchPlansExecute422Response) SetError(v PlanExecuteRejectedError)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


