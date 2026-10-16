# StoreCreative422Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**CreativeIntakeRejectedError**](CreativeIntakeRejectedError.md) |  | 
**Message** | **string** |  | 
**Errors** | **map[string][]string** |  | 

## Methods

### NewStoreCreative422Response

`func NewStoreCreative422Response(error_ CreativeIntakeRejectedError, message string, errors map[string][]string, ) *StoreCreative422Response`

NewStoreCreative422Response instantiates a new StoreCreative422Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreCreative422ResponseWithDefaults

`func NewStoreCreative422ResponseWithDefaults() *StoreCreative422Response`

NewStoreCreative422ResponseWithDefaults instantiates a new StoreCreative422Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *StoreCreative422Response) GetError() CreativeIntakeRejectedError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StoreCreative422Response) GetErrorOk() (*CreativeIntakeRejectedError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StoreCreative422Response) SetError(v CreativeIntakeRejectedError)`

SetError sets Error field to given value.


### GetMessage

`func (o *StoreCreative422Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StoreCreative422Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StoreCreative422Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *StoreCreative422Response) GetErrors() map[string][]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *StoreCreative422Response) GetErrorsOk() (*map[string][]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *StoreCreative422Response) SetErrors(v map[string][]string)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


