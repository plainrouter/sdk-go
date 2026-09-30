# ActionPolicyReadData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt32** |  | 
**WorkspaceId** | **int32** |  | 
**ExecutionMode** | **string** |  | 
**MaxSpendDeltaPercent** | **string** |  | 
**HardAccountDailyCapMinor** | **NullableInt32** |  | 
**ProtectedEntities** | **[]interface{}** |  | 
**QuietHoursStart** | **NullableString** |  | 
**QuietHoursEnd** | **NullableString** |  | 
**ProtectLearningPhase** | **bool** |  | 
**OutcomeCheckAfterHours** | **int32** |  | 
**AnomalyThresholdPercent** | **string** |  | 

## Methods

### NewActionPolicyReadData

`func NewActionPolicyReadData(id NullableInt32, workspaceId int32, executionMode string, maxSpendDeltaPercent string, hardAccountDailyCapMinor NullableInt32, protectedEntities []interface{}, quietHoursStart NullableString, quietHoursEnd NullableString, protectLearningPhase bool, outcomeCheckAfterHours int32, anomalyThresholdPercent string, ) *ActionPolicyReadData`

NewActionPolicyReadData instantiates a new ActionPolicyReadData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionPolicyReadDataWithDefaults

`func NewActionPolicyReadDataWithDefaults() *ActionPolicyReadData`

NewActionPolicyReadDataWithDefaults instantiates a new ActionPolicyReadData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionPolicyReadData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionPolicyReadData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionPolicyReadData) SetId(v int32)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ActionPolicyReadData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ActionPolicyReadData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetWorkspaceId

`func (o *ActionPolicyReadData) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ActionPolicyReadData) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ActionPolicyReadData) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetExecutionMode

`func (o *ActionPolicyReadData) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *ActionPolicyReadData) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *ActionPolicyReadData) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.


### GetMaxSpendDeltaPercent

`func (o *ActionPolicyReadData) GetMaxSpendDeltaPercent() string`

GetMaxSpendDeltaPercent returns the MaxSpendDeltaPercent field if non-nil, zero value otherwise.

### GetMaxSpendDeltaPercentOk

`func (o *ActionPolicyReadData) GetMaxSpendDeltaPercentOk() (*string, bool)`

GetMaxSpendDeltaPercentOk returns a tuple with the MaxSpendDeltaPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpendDeltaPercent

`func (o *ActionPolicyReadData) SetMaxSpendDeltaPercent(v string)`

SetMaxSpendDeltaPercent sets MaxSpendDeltaPercent field to given value.


### GetHardAccountDailyCapMinor

`func (o *ActionPolicyReadData) GetHardAccountDailyCapMinor() int32`

GetHardAccountDailyCapMinor returns the HardAccountDailyCapMinor field if non-nil, zero value otherwise.

### GetHardAccountDailyCapMinorOk

`func (o *ActionPolicyReadData) GetHardAccountDailyCapMinorOk() (*int32, bool)`

GetHardAccountDailyCapMinorOk returns a tuple with the HardAccountDailyCapMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardAccountDailyCapMinor

`func (o *ActionPolicyReadData) SetHardAccountDailyCapMinor(v int32)`

SetHardAccountDailyCapMinor sets HardAccountDailyCapMinor field to given value.


### SetHardAccountDailyCapMinorNil

`func (o *ActionPolicyReadData) SetHardAccountDailyCapMinorNil(b bool)`

 SetHardAccountDailyCapMinorNil sets the value for HardAccountDailyCapMinor to be an explicit nil

### UnsetHardAccountDailyCapMinor
`func (o *ActionPolicyReadData) UnsetHardAccountDailyCapMinor()`

UnsetHardAccountDailyCapMinor ensures that no value is present for HardAccountDailyCapMinor, not even an explicit nil
### GetProtectedEntities

`func (o *ActionPolicyReadData) GetProtectedEntities() []interface{}`

GetProtectedEntities returns the ProtectedEntities field if non-nil, zero value otherwise.

### GetProtectedEntitiesOk

`func (o *ActionPolicyReadData) GetProtectedEntitiesOk() (*[]interface{}, bool)`

GetProtectedEntitiesOk returns a tuple with the ProtectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedEntities

`func (o *ActionPolicyReadData) SetProtectedEntities(v []interface{})`

SetProtectedEntities sets ProtectedEntities field to given value.


### SetProtectedEntitiesNil

`func (o *ActionPolicyReadData) SetProtectedEntitiesNil(b bool)`

 SetProtectedEntitiesNil sets the value for ProtectedEntities to be an explicit nil

### UnsetProtectedEntities
`func (o *ActionPolicyReadData) UnsetProtectedEntities()`

UnsetProtectedEntities ensures that no value is present for ProtectedEntities, not even an explicit nil
### GetQuietHoursStart

`func (o *ActionPolicyReadData) GetQuietHoursStart() string`

GetQuietHoursStart returns the QuietHoursStart field if non-nil, zero value otherwise.

### GetQuietHoursStartOk

`func (o *ActionPolicyReadData) GetQuietHoursStartOk() (*string, bool)`

GetQuietHoursStartOk returns a tuple with the QuietHoursStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietHoursStart

`func (o *ActionPolicyReadData) SetQuietHoursStart(v string)`

SetQuietHoursStart sets QuietHoursStart field to given value.


### SetQuietHoursStartNil

`func (o *ActionPolicyReadData) SetQuietHoursStartNil(b bool)`

 SetQuietHoursStartNil sets the value for QuietHoursStart to be an explicit nil

### UnsetQuietHoursStart
`func (o *ActionPolicyReadData) UnsetQuietHoursStart()`

UnsetQuietHoursStart ensures that no value is present for QuietHoursStart, not even an explicit nil
### GetQuietHoursEnd

`func (o *ActionPolicyReadData) GetQuietHoursEnd() string`

GetQuietHoursEnd returns the QuietHoursEnd field if non-nil, zero value otherwise.

### GetQuietHoursEndOk

`func (o *ActionPolicyReadData) GetQuietHoursEndOk() (*string, bool)`

GetQuietHoursEndOk returns a tuple with the QuietHoursEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietHoursEnd

`func (o *ActionPolicyReadData) SetQuietHoursEnd(v string)`

SetQuietHoursEnd sets QuietHoursEnd field to given value.


### SetQuietHoursEndNil

`func (o *ActionPolicyReadData) SetQuietHoursEndNil(b bool)`

 SetQuietHoursEndNil sets the value for QuietHoursEnd to be an explicit nil

### UnsetQuietHoursEnd
`func (o *ActionPolicyReadData) UnsetQuietHoursEnd()`

UnsetQuietHoursEnd ensures that no value is present for QuietHoursEnd, not even an explicit nil
### GetProtectLearningPhase

`func (o *ActionPolicyReadData) GetProtectLearningPhase() bool`

GetProtectLearningPhase returns the ProtectLearningPhase field if non-nil, zero value otherwise.

### GetProtectLearningPhaseOk

`func (o *ActionPolicyReadData) GetProtectLearningPhaseOk() (*bool, bool)`

GetProtectLearningPhaseOk returns a tuple with the ProtectLearningPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectLearningPhase

`func (o *ActionPolicyReadData) SetProtectLearningPhase(v bool)`

SetProtectLearningPhase sets ProtectLearningPhase field to given value.


### GetOutcomeCheckAfterHours

`func (o *ActionPolicyReadData) GetOutcomeCheckAfterHours() int32`

GetOutcomeCheckAfterHours returns the OutcomeCheckAfterHours field if non-nil, zero value otherwise.

### GetOutcomeCheckAfterHoursOk

`func (o *ActionPolicyReadData) GetOutcomeCheckAfterHoursOk() (*int32, bool)`

GetOutcomeCheckAfterHoursOk returns a tuple with the OutcomeCheckAfterHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomeCheckAfterHours

`func (o *ActionPolicyReadData) SetOutcomeCheckAfterHours(v int32)`

SetOutcomeCheckAfterHours sets OutcomeCheckAfterHours field to given value.


### GetAnomalyThresholdPercent

`func (o *ActionPolicyReadData) GetAnomalyThresholdPercent() string`

GetAnomalyThresholdPercent returns the AnomalyThresholdPercent field if non-nil, zero value otherwise.

### GetAnomalyThresholdPercentOk

`func (o *ActionPolicyReadData) GetAnomalyThresholdPercentOk() (*string, bool)`

GetAnomalyThresholdPercentOk returns a tuple with the AnomalyThresholdPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyThresholdPercent

`func (o *ActionPolicyReadData) SetAnomalyThresholdPercent(v string)`

SetAnomalyThresholdPercent sets AnomalyThresholdPercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


