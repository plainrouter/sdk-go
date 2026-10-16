# LauncherFileImport422Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Errors** | **map[string][]string** |  | 
**Error** | **string** |  | 
**ReasonCode** | **string** |  | 
**SyntacticRowFailures** | [**[]LauncherFileImport422ResponseAnyOfSyntacticRowFailuresInner**](LauncherFileImport422ResponseAnyOfSyntacticRowFailuresInner.md) |  | 

## Methods

### NewLauncherFileImport422Response

`func NewLauncherFileImport422Response(message string, errors map[string][]string, error_ string, reasonCode string, syntacticRowFailures []LauncherFileImport422ResponseAnyOfSyntacticRowFailuresInner, ) *LauncherFileImport422Response`

NewLauncherFileImport422Response instantiates a new LauncherFileImport422Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLauncherFileImport422ResponseWithDefaults

`func NewLauncherFileImport422ResponseWithDefaults() *LauncherFileImport422Response`

NewLauncherFileImport422ResponseWithDefaults instantiates a new LauncherFileImport422Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *LauncherFileImport422Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LauncherFileImport422Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LauncherFileImport422Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *LauncherFileImport422Response) GetErrors() map[string][]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *LauncherFileImport422Response) GetErrorsOk() (*map[string][]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *LauncherFileImport422Response) SetErrors(v map[string][]string)`

SetErrors sets Errors field to given value.


### GetError

`func (o *LauncherFileImport422Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *LauncherFileImport422Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *LauncherFileImport422Response) SetError(v string)`

SetError sets Error field to given value.


### GetReasonCode

`func (o *LauncherFileImport422Response) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *LauncherFileImport422Response) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *LauncherFileImport422Response) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.


### GetSyntacticRowFailures

`func (o *LauncherFileImport422Response) GetSyntacticRowFailures() []LauncherFileImport422ResponseAnyOfSyntacticRowFailuresInner`

GetSyntacticRowFailures returns the SyntacticRowFailures field if non-nil, zero value otherwise.

### GetSyntacticRowFailuresOk

`func (o *LauncherFileImport422Response) GetSyntacticRowFailuresOk() (*[]LauncherFileImport422ResponseAnyOfSyntacticRowFailuresInner, bool)`

GetSyntacticRowFailuresOk returns a tuple with the SyntacticRowFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntacticRowFailures

`func (o *LauncherFileImport422Response) SetSyntacticRowFailures(v []LauncherFileImport422ResponseAnyOfSyntacticRowFailuresInner)`

SetSyntacticRowFailures sets SyntacticRowFailures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


