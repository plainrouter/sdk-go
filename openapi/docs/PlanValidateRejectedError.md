# PlanValidateRejectedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Launch plan refusal code, including launch_video_upload_unavailable when a new Meta video upload would be required. | 
**Message** | **string** |  | 

## Methods

### NewPlanValidateRejectedError

`func NewPlanValidateRejectedError(code string, message string, ) *PlanValidateRejectedError`

NewPlanValidateRejectedError instantiates a new PlanValidateRejectedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanValidateRejectedErrorWithDefaults

`func NewPlanValidateRejectedErrorWithDefaults() *PlanValidateRejectedError`

NewPlanValidateRejectedErrorWithDefaults instantiates a new PlanValidateRejectedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *PlanValidateRejectedError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PlanValidateRejectedError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PlanValidateRejectedError) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *PlanValidateRejectedError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PlanValidateRejectedError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PlanValidateRejectedError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


