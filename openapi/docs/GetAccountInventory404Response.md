# GetAccountInventory404Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | [**ApiRouteNotFoundError**](ApiRouteNotFoundError.md) |  | 
**Request** | [**ApiRouteNotFoundRequest**](ApiRouteNotFoundRequest.md) |  | 
**Resources** | [**ApiRouteNotFoundResources**](ApiRouteNotFoundResources.md) |  | 
**Message** | **string** |  | 

## Methods

### NewGetAccountInventory404Response

`func NewGetAccountInventory404Response(error_ ApiRouteNotFoundError, request ApiRouteNotFoundRequest, resources ApiRouteNotFoundResources, message string, ) *GetAccountInventory404Response`

NewGetAccountInventory404Response instantiates a new GetAccountInventory404Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountInventory404ResponseWithDefaults

`func NewGetAccountInventory404ResponseWithDefaults() *GetAccountInventory404Response`

NewGetAccountInventory404ResponseWithDefaults instantiates a new GetAccountInventory404Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GetAccountInventory404Response) GetError() ApiRouteNotFoundError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetAccountInventory404Response) GetErrorOk() (*ApiRouteNotFoundError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetAccountInventory404Response) SetError(v ApiRouteNotFoundError)`

SetError sets Error field to given value.


### GetRequest

`func (o *GetAccountInventory404Response) GetRequest() ApiRouteNotFoundRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *GetAccountInventory404Response) GetRequestOk() (*ApiRouteNotFoundRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *GetAccountInventory404Response) SetRequest(v ApiRouteNotFoundRequest)`

SetRequest sets Request field to given value.


### GetResources

`func (o *GetAccountInventory404Response) GetResources() ApiRouteNotFoundResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GetAccountInventory404Response) GetResourcesOk() (*ApiRouteNotFoundResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GetAccountInventory404Response) SetResources(v ApiRouteNotFoundResources)`

SetResources sets Resources field to given value.


### GetMessage

`func (o *GetAccountInventory404Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetAccountInventory404Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetAccountInventory404Response) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


