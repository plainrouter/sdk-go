# ActionRuleWriteReadRule

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

## Methods

### NewActionRuleWriteReadRule

`func NewActionRuleWriteReadRule(id string, workspaceId int32, platformAdAccountId int32, name string, enabled bool, suspendedAt NullableTime, suspendedReason NullableString, window ActionRuleDetailReadRuleWindow, currentRevision ActionRuleDetailReadRuleCurrentRevision, revisions []ActionRuleDetailReadRuleCurrentRevision, ) *ActionRuleWriteReadRule`

NewActionRuleWriteReadRule instantiates a new ActionRuleWriteReadRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRuleWriteReadRuleWithDefaults

`func NewActionRuleWriteReadRuleWithDefaults() *ActionRuleWriteReadRule`

NewActionRuleWriteReadRuleWithDefaults instantiates a new ActionRuleWriteReadRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionRuleWriteReadRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionRuleWriteReadRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionRuleWriteReadRule) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *ActionRuleWriteReadRule) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ActionRuleWriteReadRule) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ActionRuleWriteReadRule) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetPlatformAdAccountId

`func (o *ActionRuleWriteReadRule) GetPlatformAdAccountId() int32`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *ActionRuleWriteReadRule) GetPlatformAdAccountIdOk() (*int32, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *ActionRuleWriteReadRule) SetPlatformAdAccountId(v int32)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.


### GetName

`func (o *ActionRuleWriteReadRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionRuleWriteReadRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionRuleWriteReadRule) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *ActionRuleWriteReadRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ActionRuleWriteReadRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ActionRuleWriteReadRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSuspendedAt

`func (o *ActionRuleWriteReadRule) GetSuspendedAt() time.Time`

GetSuspendedAt returns the SuspendedAt field if non-nil, zero value otherwise.

### GetSuspendedAtOk

`func (o *ActionRuleWriteReadRule) GetSuspendedAtOk() (*time.Time, bool)`

GetSuspendedAtOk returns a tuple with the SuspendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedAt

`func (o *ActionRuleWriteReadRule) SetSuspendedAt(v time.Time)`

SetSuspendedAt sets SuspendedAt field to given value.


### SetSuspendedAtNil

`func (o *ActionRuleWriteReadRule) SetSuspendedAtNil(b bool)`

 SetSuspendedAtNil sets the value for SuspendedAt to be an explicit nil

### UnsetSuspendedAt
`func (o *ActionRuleWriteReadRule) UnsetSuspendedAt()`

UnsetSuspendedAt ensures that no value is present for SuspendedAt, not even an explicit nil
### GetSuspendedReason

`func (o *ActionRuleWriteReadRule) GetSuspendedReason() string`

GetSuspendedReason returns the SuspendedReason field if non-nil, zero value otherwise.

### GetSuspendedReasonOk

`func (o *ActionRuleWriteReadRule) GetSuspendedReasonOk() (*string, bool)`

GetSuspendedReasonOk returns a tuple with the SuspendedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedReason

`func (o *ActionRuleWriteReadRule) SetSuspendedReason(v string)`

SetSuspendedReason sets SuspendedReason field to given value.


### SetSuspendedReasonNil

`func (o *ActionRuleWriteReadRule) SetSuspendedReasonNil(b bool)`

 SetSuspendedReasonNil sets the value for SuspendedReason to be an explicit nil

### UnsetSuspendedReason
`func (o *ActionRuleWriteReadRule) UnsetSuspendedReason()`

UnsetSuspendedReason ensures that no value is present for SuspendedReason, not even an explicit nil
### GetWindow

`func (o *ActionRuleWriteReadRule) GetWindow() ActionRuleDetailReadRuleWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ActionRuleWriteReadRule) GetWindowOk() (*ActionRuleDetailReadRuleWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ActionRuleWriteReadRule) SetWindow(v ActionRuleDetailReadRuleWindow)`

SetWindow sets Window field to given value.


### GetCurrentRevision

`func (o *ActionRuleWriteReadRule) GetCurrentRevision() ActionRuleDetailReadRuleCurrentRevision`

GetCurrentRevision returns the CurrentRevision field if non-nil, zero value otherwise.

### GetCurrentRevisionOk

`func (o *ActionRuleWriteReadRule) GetCurrentRevisionOk() (*ActionRuleDetailReadRuleCurrentRevision, bool)`

GetCurrentRevisionOk returns a tuple with the CurrentRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevision

`func (o *ActionRuleWriteReadRule) SetCurrentRevision(v ActionRuleDetailReadRuleCurrentRevision)`

SetCurrentRevision sets CurrentRevision field to given value.


### GetRevisions

`func (o *ActionRuleWriteReadRule) GetRevisions() []ActionRuleDetailReadRuleCurrentRevision`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *ActionRuleWriteReadRule) GetRevisionsOk() (*[]ActionRuleDetailReadRuleCurrentRevision, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *ActionRuleWriteReadRule) SetRevisions(v []ActionRuleDetailReadRuleCurrentRevision)`

SetRevisions sets Revisions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


