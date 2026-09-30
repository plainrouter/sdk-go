# DeleteUserDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentifierType** | **string** |  | 
**IdentifierHash** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteUserDataRequest

`func NewDeleteUserDataRequest(identifierType string, ) *DeleteUserDataRequest`

NewDeleteUserDataRequest instantiates a new DeleteUserDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteUserDataRequestWithDefaults

`func NewDeleteUserDataRequestWithDefaults() *DeleteUserDataRequest`

NewDeleteUserDataRequestWithDefaults instantiates a new DeleteUserDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifierType

`func (o *DeleteUserDataRequest) GetIdentifierType() string`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *DeleteUserDataRequest) GetIdentifierTypeOk() (*string, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *DeleteUserDataRequest) SetIdentifierType(v string)`

SetIdentifierType sets IdentifierType field to given value.


### GetIdentifierHash

`func (o *DeleteUserDataRequest) GetIdentifierHash() string`

GetIdentifierHash returns the IdentifierHash field if non-nil, zero value otherwise.

### GetIdentifierHashOk

`func (o *DeleteUserDataRequest) GetIdentifierHashOk() (*string, bool)`

GetIdentifierHashOk returns a tuple with the IdentifierHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierHash

`func (o *DeleteUserDataRequest) SetIdentifierHash(v string)`

SetIdentifierHash sets IdentifierHash field to given value.

### HasIdentifierHash

`func (o *DeleteUserDataRequest) HasIdentifierHash() bool`

HasIdentifierHash returns a boolean if a field has been set.

### GetIdentifier

`func (o *DeleteUserDataRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DeleteUserDataRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DeleteUserDataRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *DeleteUserDataRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


