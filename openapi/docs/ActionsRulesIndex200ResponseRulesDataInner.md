# ActionsRulesIndex200ResponseRulesDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Status** | **string** |  | 
**Account** | [**ActionsRulesIndex200ResponseRulesDataInnerAccount**](ActionsRulesIndex200ResponseRulesDataInnerAccount.md) |  | 
**CurrentRevision** | **map[string]interface{}** |  | 
**LastRunStatus** | **NullableString** |  | 
**LastRunAt** | **NullableString** |  | 
**Counts** | [**ActionsRulesIndex200ResponseRulesDataInnerCounts**](ActionsRulesIndex200ResponseRulesDataInnerCounts.md) |  | 

## Methods

### NewActionsRulesIndex200ResponseRulesDataInner

`func NewActionsRulesIndex200ResponseRulesDataInner(id string, name string, status string, account ActionsRulesIndex200ResponseRulesDataInnerAccount, currentRevision map[string]*interface{}, lastRunStatus NullableString, lastRunAt NullableString, counts ActionsRulesIndex200ResponseRulesDataInnerCounts, ) *ActionsRulesIndex200ResponseRulesDataInner`

NewActionsRulesIndex200ResponseRulesDataInner instantiates a new ActionsRulesIndex200ResponseRulesDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsRulesIndex200ResponseRulesDataInnerWithDefaults

`func NewActionsRulesIndex200ResponseRulesDataInnerWithDefaults() *ActionsRulesIndex200ResponseRulesDataInner`

NewActionsRulesIndex200ResponseRulesDataInnerWithDefaults instantiates a new ActionsRulesIndex200ResponseRulesDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionsRulesIndex200ResponseRulesDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionsRulesIndex200ResponseRulesDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionsRulesIndex200ResponseRulesDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAccount

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetAccount() ActionsRulesIndex200ResponseRulesDataInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetAccountOk() (*ActionsRulesIndex200ResponseRulesDataInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ActionsRulesIndex200ResponseRulesDataInner) SetAccount(v ActionsRulesIndex200ResponseRulesDataInnerAccount)`

SetAccount sets Account field to given value.


### GetCurrentRevision

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetCurrentRevision() map[string]*interface{}`

GetCurrentRevision returns the CurrentRevision field if non-nil, zero value otherwise.

### GetCurrentRevisionOk

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetCurrentRevisionOk() (*map[string]*interface{}, bool)`

GetCurrentRevisionOk returns a tuple with the CurrentRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevision

`func (o *ActionsRulesIndex200ResponseRulesDataInner) SetCurrentRevision(v map[string]*interface{})`

SetCurrentRevision sets CurrentRevision field to given value.


### GetLastRunStatus

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetLastRunStatus() string`

GetLastRunStatus returns the LastRunStatus field if non-nil, zero value otherwise.

### GetLastRunStatusOk

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetLastRunStatusOk() (*string, bool)`

GetLastRunStatusOk returns a tuple with the LastRunStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunStatus

`func (o *ActionsRulesIndex200ResponseRulesDataInner) SetLastRunStatus(v string)`

SetLastRunStatus sets LastRunStatus field to given value.


### SetLastRunStatusNil

`func (o *ActionsRulesIndex200ResponseRulesDataInner) SetLastRunStatusNil(b bool)`

 SetLastRunStatusNil sets the value for LastRunStatus to be an explicit nil

### UnsetLastRunStatus
`func (o *ActionsRulesIndex200ResponseRulesDataInner) UnsetLastRunStatus()`

UnsetLastRunStatus ensures that no value is present for LastRunStatus, not even an explicit nil
### GetLastRunAt

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetLastRunAt() string`

GetLastRunAt returns the LastRunAt field if non-nil, zero value otherwise.

### GetLastRunAtOk

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetLastRunAtOk() (*string, bool)`

GetLastRunAtOk returns a tuple with the LastRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAt

`func (o *ActionsRulesIndex200ResponseRulesDataInner) SetLastRunAt(v string)`

SetLastRunAt sets LastRunAt field to given value.


### SetLastRunAtNil

`func (o *ActionsRulesIndex200ResponseRulesDataInner) SetLastRunAtNil(b bool)`

 SetLastRunAtNil sets the value for LastRunAt to be an explicit nil

### UnsetLastRunAt
`func (o *ActionsRulesIndex200ResponseRulesDataInner) UnsetLastRunAt()`

UnsetLastRunAt ensures that no value is present for LastRunAt, not even an explicit nil
### GetCounts

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetCounts() ActionsRulesIndex200ResponseRulesDataInnerCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *ActionsRulesIndex200ResponseRulesDataInner) GetCountsOk() (*ActionsRulesIndex200ResponseRulesDataInnerCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *ActionsRulesIndex200ResponseRulesDataInner) SetCounts(v ActionsRulesIndex200ResponseRulesDataInnerCounts)`

SetCounts sets Counts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


