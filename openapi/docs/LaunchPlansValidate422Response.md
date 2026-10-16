# LaunchPlansValidate422Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Errors** | **map[string][]string** |  | 
**Error** | [**PlanValidateRejectedError**](PlanValidateRejectedError.md) |  | 

## Methods

### NewLaunchPlansValidate422Response

`func NewLaunchPlansValidate422Response(message string, errors map[string][]string, error_ PlanValidateRejectedError, ) *LaunchPlansValidate422Response`

NewLaunchPlansValidate422Response instantiates a new LaunchPlansValidate422Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchPlansValidate422ResponseWithDefaults

`func NewLaunchPlansValidate422ResponseWithDefaults() *LaunchPlansValidate422Response`

NewLaunchPlansValidate422ResponseWithDefaults instantiates a new LaunchPlansValidate422Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *LaunchPlansValidate422Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LaunchPlansValidate422Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LaunchPlansValidate422Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *LaunchPlansValidate422Response) GetErrors() map[string][]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LaunchPlansValidate422Response) GetErrorsOk() (*map[string][]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LaunchPlansValidate422Response) SetErrors(v map[string][]string)`

SetErrors sets Errors field to given value.


### GetError

`func (o *LaunchPlansValidate422Response) GetError() PlanValidateRejectedError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LaunchPlansValidate422Response) GetErrorOk() (*PlanValidateRejectedError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *LaunchPlansValidate422Response) SetError(v PlanValidateRejectedError)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


