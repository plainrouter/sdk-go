# ActionRuleDetailReadRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**WorkspaceId** | **int32** |  | 
**PlatformAdAccountId** | **int32** |  | 
**Name** | **string** |  | 
**Enabled** | **bool** |  | 
**SuspendedAt** | **NullableTime** |  | 
**SuspendedReason** | **NullableString** |  | 
**Window** | [**ActionRuleDetailReadRuleWindow**](ActionRuleDetailReadRuleWindow.md) |  | 
**CurrentRevision** | [**ActionRuleDetailReadRuleCurrentRevision**](ActionRuleDetailReadRuleCurrentRevision.md) |  | 
**Revisions** | [**[]ActionRuleDetailReadRuleCurrentRevision**](ActionRuleDetailReadRuleCurrentRevision.md) | Present on show and write responses. | 
**Runs** | [**[]ActionRuleDetailReadRuleRunsInner**](ActionRuleDetailReadRuleRunsInner.md) | Show only; at most five recent runs, each with at most 100 decisions. | 

## Methods

### NewActionRuleDetailReadRule

`func NewActionRuleDetailReadRule(id string, workspaceId int32, platformAdAccountId int32, name string, enabled bool, suspendedAt NullableTime, suspendedReason NullableString, window ActionRuleDetailReadRuleWindow, currentRevision ActionRuleDetailReadRuleCurrentRevision, revisions []ActionRuleDetailReadRuleCurrentRevision, runs []ActionRuleDetailReadRuleRunsInner, ) *ActionRuleDetailReadRule`

NewActionRuleDetailReadRule instantiates a new ActionRuleDetailReadRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRuleDetailReadRuleWithDefaults

`func NewActionRuleDetailReadRuleWithDefaults() *ActionRuleDetailReadRule`

NewActionRuleDetailReadRuleWithDefaults instantiates a new ActionRuleDetailReadRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionRuleDetailReadRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionRuleDetailReadRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionRuleDetailReadRule) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *ActionRuleDetailReadRule) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ActionRuleDetailReadRule) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ActionRuleDetailReadRule) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetPlatformAdAccountId

`func (o *ActionRuleDetailReadRule) GetPlatformAdAccountId() int32`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *ActionRuleDetailReadRule) GetPlatformAdAccountIdOk() (*int32, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *ActionRuleDetailReadRule) SetPlatformAdAccountId(v int32)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.


### GetName

`func (o *ActionRuleDetailReadRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionRuleDetailReadRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionRuleDetailReadRule) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *ActionRuleDetailReadRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ActionRuleDetailReadRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ActionRuleDetailReadRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSuspendedAt

`func (o *ActionRuleDetailReadRule) GetSuspendedAt() time.Time`

GetSuspendedAt returns the SuspendedAt field if non-nil, zero value otherwise.

### GetSuspendedAtOk

`func (o *ActionRuleDetailReadRule) GetSuspendedAtOk() (*time.Time, bool)`

GetSuspendedAtOk returns a tuple with the SuspendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedAt

`func (o *ActionRuleDetailReadRule) SetSuspendedAt(v time.Time)`

SetSuspendedAt sets SuspendedAt field to given value.


### SetSuspendedAtNil

`func (o *ActionRuleDetailReadRule) SetSuspendedAtNil(b bool)`

 SetSuspendedAtNil sets the value for SuspendedAt to be an explicit nil

### UnsetSuspendedAt
`func (o *ActionRuleDetailReadRule) UnsetSuspendedAt()`

UnsetSuspendedAt ensures that no value is present for SuspendedAt, not even an explicit nil
### GetSuspendedReason

`func (o *ActionRuleDetailReadRule) GetSuspendedReason() string`

GetSuspendedReason returns the SuspendedReason field if non-nil, zero value otherwise.

### GetSuspendedReasonOk

`func (o *ActionRuleDetailReadRule) GetSuspendedReasonOk() (*string, bool)`

GetSuspendedReasonOk returns a tuple with the SuspendedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedReason

`func (o *ActionRuleDetailReadRule) SetSuspendedReason(v string)`

SetSuspendedReason sets SuspendedReason field to given value.


### SetSuspendedReasonNil

`func (o *ActionRuleDetailReadRule) SetSuspendedReasonNil(b bool)`

 SetSuspendedReasonNil sets the value for SuspendedReason to be an explicit nil

### UnsetSuspendedReason
`func (o *ActionRuleDetailReadRule) UnsetSuspendedReason()`

UnsetSuspendedReason ensures that no value is present for SuspendedReason, not even an explicit nil
### GetWindow

`func (o *ActionRuleDetailReadRule) GetWindow() ActionRuleDetailReadRuleWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ActionRuleDetailReadRule) GetWindowOk() (*ActionRuleDetailReadRuleWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ActionRuleDetailReadRule) SetWindow(v ActionRuleDetailReadRuleWindow)`

SetWindow sets Window field to given value.


### GetCurrentRevision

`func (o *ActionRuleDetailReadRule) GetCurrentRevision() ActionRuleDetailReadRuleCurrentRevision`

GetCurrentRevision returns the CurrentRevision field if non-nil, zero value otherwise.

### GetCurrentRevisionOk

`func (o *ActionRuleDetailReadRule) GetCurrentRevisionOk() (*ActionRuleDetailReadRuleCurrentRevision, bool)`

GetCurrentRevisionOk returns a tuple with the CurrentRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevision

`func (o *ActionRuleDetailReadRule) SetCurrentRevision(v ActionRuleDetailReadRuleCurrentRevision)`

SetCurrentRevision sets CurrentRevision field to given value.


### GetRevisions

`func (o *ActionRuleDetailReadRule) GetRevisions() []ActionRuleDetailReadRuleCurrentRevision`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *ActionRuleDetailReadRule) GetRevisionsOk() (*[]ActionRuleDetailReadRuleCurrentRevision, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *ActionRuleDetailReadRule) SetRevisions(v []ActionRuleDetailReadRuleCurrentRevision)`

SetRevisions sets Revisions field to given value.


### GetRuns

`func (o *ActionRuleDetailReadRule) GetRuns() []ActionRuleDetailReadRuleRunsInner`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *ActionRuleDetailReadRule) GetRunsOk() (*[]ActionRuleDetailReadRuleRunsInner, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *ActionRuleDetailReadRule) SetRuns(v []ActionRuleDetailReadRuleRunsInner)`

SetRuns sets Runs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


