# ActionsRulesUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Definition** | Pointer to [**ActionsRulesStoreRequestDefinition**](ActionsRulesStoreRequestDefinition.md) |  | [optional] 
**AccountId** | Pointer to **int32** | Optional integer account ID. JSON booleans return 422 with message plus field errors. A valid supplied ID is ignored by an account-bound key. An unbound key requires exactly one active available advertising account with an active connection; omission uses that account and a supplied ID must match it. Multiple active accounts return 422 even when account_id is supplied. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Window** | Pointer to [**NullableActionsRulesStoreRequestWindow**](ActionsRulesStoreRequestWindow.md) |  | [optional] 
**WindowDays** | Pointer to [**NullableActionsRulesStoreRequestWindow**](ActionsRulesStoreRequestWindow.md) |  | [optional] 
**AuthorType** | Pointer to [**NullableActionsRulesStoreRequestWindow**](ActionsRulesStoreRequestWindow.md) |  | [optional] 
**CurrentRevisionId** | Pointer to [**NullableActionsRulesStoreRequestWindow**](ActionsRulesStoreRequestWindow.md) |  | [optional] 

## Methods

### NewActionsRulesUpdateRequest

`func NewActionsRulesUpdateRequest() *ActionsRulesUpdateRequest`

NewActionsRulesUpdateRequest instantiates a new ActionsRulesUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsRulesUpdateRequestWithDefaults

`func NewActionsRulesUpdateRequestWithDefaults() *ActionsRulesUpdateRequest`

NewActionsRulesUpdateRequestWithDefaults instantiates a new ActionsRulesUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ActionsRulesUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionsRulesUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionsRulesUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActionsRulesUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefinition

`func (o *ActionsRulesUpdateRequest) GetDefinition() ActionsRulesStoreRequestDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *ActionsRulesUpdateRequest) GetDefinitionOk() (*ActionsRulesStoreRequestDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *ActionsRulesUpdateRequest) SetDefinition(v ActionsRulesStoreRequestDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *ActionsRulesUpdateRequest) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetAccountId

`func (o *ActionsRulesUpdateRequest) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ActionsRulesUpdateRequest) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ActionsRulesUpdateRequest) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ActionsRulesUpdateRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetEnabled

`func (o *ActionsRulesUpdateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ActionsRulesUpdateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ActionsRulesUpdateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ActionsRulesUpdateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetWindow

`func (o *ActionsRulesUpdateRequest) GetWindow() ActionsRulesStoreRequestWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ActionsRulesUpdateRequest) GetWindowOk() (*ActionsRulesStoreRequestWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ActionsRulesUpdateRequest) SetWindow(v ActionsRulesStoreRequestWindow)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *ActionsRulesUpdateRequest) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### SetWindowNil

`func (o *ActionsRulesUpdateRequest) SetWindowNil(b bool)`

 SetWindowNil sets the value for Window to be an explicit nil

### UnsetWindow
`func (o *ActionsRulesUpdateRequest) UnsetWindow()`

UnsetWindow ensures that no value is present for Window, not even an explicit nil
### GetWindowDays

`func (o *ActionsRulesUpdateRequest) GetWindowDays() ActionsRulesStoreRequestWindow`

GetWindowDays returns the WindowDays field if non-nil, zero value otherwise.

### GetWindowDaysOk

`func (o *ActionsRulesUpdateRequest) GetWindowDaysOk() (*ActionsRulesStoreRequestWindow, bool)`

GetWindowDaysOk returns a tuple with the WindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowDays

`func (o *ActionsRulesUpdateRequest) SetWindowDays(v ActionsRulesStoreRequestWindow)`

SetWindowDays sets WindowDays field to given value.

### HasWindowDays

`func (o *ActionsRulesUpdateRequest) HasWindowDays() bool`

HasWindowDays returns a boolean if a field has been set.

### SetWindowDaysNil

`func (o *ActionsRulesUpdateRequest) SetWindowDaysNil(b bool)`

 SetWindowDaysNil sets the value for WindowDays to be an explicit nil

### UnsetWindowDays
`func (o *ActionsRulesUpdateRequest) UnsetWindowDays()`

UnsetWindowDays ensures that no value is present for WindowDays, not even an explicit nil
### GetAuthorType

`func (o *ActionsRulesUpdateRequest) GetAuthorType() ActionsRulesStoreRequestWindow`

GetAuthorType returns the AuthorType field if non-nil, zero value otherwise.

### GetAuthorTypeOk

`func (o *ActionsRulesUpdateRequest) GetAuthorTypeOk() (*ActionsRulesStoreRequestWindow, bool)`

GetAuthorTypeOk returns a tuple with the AuthorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorType

`func (o *ActionsRulesUpdateRequest) SetAuthorType(v ActionsRulesStoreRequestWindow)`

SetAuthorType sets AuthorType field to given value.

### HasAuthorType

`func (o *ActionsRulesUpdateRequest) HasAuthorType() bool`

HasAuthorType returns a boolean if a field has been set.

### SetAuthorTypeNil

`func (o *ActionsRulesUpdateRequest) SetAuthorTypeNil(b bool)`

 SetAuthorTypeNil sets the value for AuthorType to be an explicit nil

### UnsetAuthorType
`func (o *ActionsRulesUpdateRequest) UnsetAuthorType()`

UnsetAuthorType ensures that no value is present for AuthorType, not even an explicit nil
### GetCurrentRevisionId

`func (o *ActionsRulesUpdateRequest) GetCurrentRevisionId() ActionsRulesStoreRequestWindow`

GetCurrentRevisionId returns the CurrentRevisionId field if non-nil, zero value otherwise.

### GetCurrentRevisionIdOk

`func (o *ActionsRulesUpdateRequest) GetCurrentRevisionIdOk() (*ActionsRulesStoreRequestWindow, bool)`

GetCurrentRevisionIdOk returns a tuple with the CurrentRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevisionId

`func (o *ActionsRulesUpdateRequest) SetCurrentRevisionId(v ActionsRulesStoreRequestWindow)`

SetCurrentRevisionId sets CurrentRevisionId field to given value.

### HasCurrentRevisionId

`func (o *ActionsRulesUpdateRequest) HasCurrentRevisionId() bool`

HasCurrentRevisionId returns a boolean if a field has been set.

### SetCurrentRevisionIdNil

`func (o *ActionsRulesUpdateRequest) SetCurrentRevisionIdNil(b bool)`

 SetCurrentRevisionIdNil sets the value for CurrentRevisionId to be an explicit nil

### UnsetCurrentRevisionId
`func (o *ActionsRulesUpdateRequest) UnsetCurrentRevisionId()`

UnsetCurrentRevisionId ensures that no value is present for CurrentRevisionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


