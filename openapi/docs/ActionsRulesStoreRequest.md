# ActionsRulesStoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Definition** | [**ActionsRulesStoreRequestDefinition**](ActionsRulesStoreRequestDefinition.md) |  | 
**AccountId** | Pointer to **int32** | Optional integer account ID. JSON booleans return 422 with message plus field errors. A valid supplied ID is ignored by an account-bound key. An unbound key requires exactly one active available advertising account with an active connection; omission uses that account and a supplied ID must match it. Multiple active accounts return 422 even when account_id is supplied. | [optional] 
**PlatformAdAccountId** | **int32** |  | 
**Window** | Pointer to [**NullableActionsRulesStoreRequestWindow**](ActionsRulesStoreRequestWindow.md) |  | [optional] 
**WindowDays** | Pointer to [**NullableActionsRulesStoreRequestWindow**](ActionsRulesStoreRequestWindow.md) |  | [optional] 
**AuthorType** | Pointer to [**NullableActionsRulesStoreRequestWindow**](ActionsRulesStoreRequestWindow.md) |  | [optional] 
**CurrentRevisionId** | Pointer to [**NullableActionsRulesStoreRequestWindow**](ActionsRulesStoreRequestWindow.md) |  | [optional] 

## Methods

### NewActionsRulesStoreRequest

`func NewActionsRulesStoreRequest(name string, definition ActionsRulesStoreRequestDefinition, platformAdAccountId int32, ) *ActionsRulesStoreRequest`

NewActionsRulesStoreRequest instantiates a new ActionsRulesStoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsRulesStoreRequestWithDefaults

`func NewActionsRulesStoreRequestWithDefaults() *ActionsRulesStoreRequest`

NewActionsRulesStoreRequestWithDefaults instantiates a new ActionsRulesStoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ActionsRulesStoreRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionsRulesStoreRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionsRulesStoreRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDefinition

`func (o *ActionsRulesStoreRequest) GetDefinition() ActionsRulesStoreRequestDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *ActionsRulesStoreRequest) GetDefinitionOk() (*ActionsRulesStoreRequestDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *ActionsRulesStoreRequest) SetDefinition(v ActionsRulesStoreRequestDefinition)`

SetDefinition sets Definition field to given value.


### GetAccountId

`func (o *ActionsRulesStoreRequest) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ActionsRulesStoreRequest) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ActionsRulesStoreRequest) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ActionsRulesStoreRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatformAdAccountId

`func (o *ActionsRulesStoreRequest) GetPlatformAdAccountId() int32`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *ActionsRulesStoreRequest) GetPlatformAdAccountIdOk() (*int32, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *ActionsRulesStoreRequest) SetPlatformAdAccountId(v int32)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.


### GetWindow

`func (o *ActionsRulesStoreRequest) GetWindow() ActionsRulesStoreRequestWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ActionsRulesStoreRequest) GetWindowOk() (*ActionsRulesStoreRequestWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ActionsRulesStoreRequest) SetWindow(v ActionsRulesStoreRequestWindow)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *ActionsRulesStoreRequest) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### SetWindowNil

`func (o *ActionsRulesStoreRequest) SetWindowNil(b bool)`

 SetWindowNil sets the value for Window to be an explicit nil

### UnsetWindow
`func (o *ActionsRulesStoreRequest) UnsetWindow()`

UnsetWindow ensures that no value is present for Window, not even an explicit nil
### GetWindowDays

`func (o *ActionsRulesStoreRequest) GetWindowDays() ActionsRulesStoreRequestWindow`

GetWindowDays returns the WindowDays field if non-nil, zero value otherwise.

### GetWindowDaysOk

`func (o *ActionsRulesStoreRequest) GetWindowDaysOk() (*ActionsRulesStoreRequestWindow, bool)`

GetWindowDaysOk returns a tuple with the WindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowDays

`func (o *ActionsRulesStoreRequest) SetWindowDays(v ActionsRulesStoreRequestWindow)`

SetWindowDays sets WindowDays field to given value.

### HasWindowDays

`func (o *ActionsRulesStoreRequest) HasWindowDays() bool`

HasWindowDays returns a boolean if a field has been set.

### SetWindowDaysNil

`func (o *ActionsRulesStoreRequest) SetWindowDaysNil(b bool)`

 SetWindowDaysNil sets the value for WindowDays to be an explicit nil

### UnsetWindowDays
`func (o *ActionsRulesStoreRequest) UnsetWindowDays()`

UnsetWindowDays ensures that no value is present for WindowDays, not even an explicit nil
### GetAuthorType

`func (o *ActionsRulesStoreRequest) GetAuthorType() ActionsRulesStoreRequestWindow`

GetAuthorType returns the AuthorType field if non-nil, zero value otherwise.

### GetAuthorTypeOk

`func (o *ActionsRulesStoreRequest) GetAuthorTypeOk() (*ActionsRulesStoreRequestWindow, bool)`

GetAuthorTypeOk returns a tuple with the AuthorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorType

`func (o *ActionsRulesStoreRequest) SetAuthorType(v ActionsRulesStoreRequestWindow)`

SetAuthorType sets AuthorType field to given value.

### HasAuthorType

`func (o *ActionsRulesStoreRequest) HasAuthorType() bool`

HasAuthorType returns a boolean if a field has been set.

### SetAuthorTypeNil

`func (o *ActionsRulesStoreRequest) SetAuthorTypeNil(b bool)`

 SetAuthorTypeNil sets the value for AuthorType to be an explicit nil

### UnsetAuthorType
`func (o *ActionsRulesStoreRequest) UnsetAuthorType()`

UnsetAuthorType ensures that no value is present for AuthorType, not even an explicit nil
### GetCurrentRevisionId

`func (o *ActionsRulesStoreRequest) GetCurrentRevisionId() ActionsRulesStoreRequestWindow`

GetCurrentRevisionId returns the CurrentRevisionId field if non-nil, zero value otherwise.

### GetCurrentRevisionIdOk

`func (o *ActionsRulesStoreRequest) GetCurrentRevisionIdOk() (*ActionsRulesStoreRequestWindow, bool)`

GetCurrentRevisionIdOk returns a tuple with the CurrentRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevisionId

`func (o *ActionsRulesStoreRequest) SetCurrentRevisionId(v ActionsRulesStoreRequestWindow)`

SetCurrentRevisionId sets CurrentRevisionId field to given value.

### HasCurrentRevisionId

`func (o *ActionsRulesStoreRequest) HasCurrentRevisionId() bool`

HasCurrentRevisionId returns a boolean if a field has been set.

### SetCurrentRevisionIdNil

`func (o *ActionsRulesStoreRequest) SetCurrentRevisionIdNil(b bool)`

 SetCurrentRevisionIdNil sets the value for CurrentRevisionId to be an explicit nil

### UnsetCurrentRevisionId
`func (o *ActionsRulesStoreRequest) UnsetCurrentRevisionId()`

UnsetCurrentRevisionId ensures that no value is present for CurrentRevisionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


