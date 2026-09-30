# GetEmqReport200ResponseSnapshotsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**WorkspaceId** | **int32** |  | 
**SignalTrackerId** | **string** | Deprecated alias of workspace_id; contains the workspace ID in decimal string form. | 
**DestinationId** | **string** |  | 
**Score** | **float32** |  | 
**WeekOverWeekChange** | **NullableFloat32** |  | 
**Alerted** | **bool** |  | 
**PlatformResponse** | **interface{}** |  | 
**MeasuredAt** | **string** |  | 
**CreatedAt** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 

## Methods

### NewGetEmqReport200ResponseSnapshotsInner

`func NewGetEmqReport200ResponseSnapshotsInner(id int32, workspaceId int32, signalTrackerId string, destinationId string, score float32, weekOverWeekChange NullableFloat32, alerted bool, platformResponse interface{}, measuredAt string, createdAt NullableString, updatedAt NullableString, ) *GetEmqReport200ResponseSnapshotsInner`

NewGetEmqReport200ResponseSnapshotsInner instantiates a new GetEmqReport200ResponseSnapshotsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEmqReport200ResponseSnapshotsInnerWithDefaults

`func NewGetEmqReport200ResponseSnapshotsInnerWithDefaults() *GetEmqReport200ResponseSnapshotsInner`

NewGetEmqReport200ResponseSnapshotsInnerWithDefaults instantiates a new GetEmqReport200ResponseSnapshotsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEmqReport200ResponseSnapshotsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEmqReport200ResponseSnapshotsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEmqReport200ResponseSnapshotsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *GetEmqReport200ResponseSnapshotsInner) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GetEmqReport200ResponseSnapshotsInner) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GetEmqReport200ResponseSnapshotsInner) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetSignalTrackerId

`func (o *GetEmqReport200ResponseSnapshotsInner) GetSignalTrackerId() string`

GetSignalTrackerId returns the SignalTrackerId field if non-nil, zero value otherwise.

### GetSignalTrackerIdOk

`func (o *GetEmqReport200ResponseSnapshotsInner) GetSignalTrackerIdOk() (*string, bool)`

GetSignalTrackerIdOk returns a tuple with the SignalTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalTrackerId

`func (o *GetEmqReport200ResponseSnapshotsInner) SetSignalTrackerId(v string)`

SetSignalTrackerId sets SignalTrackerId field to given value.


### GetDestinationId

`func (o *GetEmqReport200ResponseSnapshotsInner) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *GetEmqReport200ResponseSnapshotsInner) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *GetEmqReport200ResponseSnapshotsInner) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetScore

`func (o *GetEmqReport200ResponseSnapshotsInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *GetEmqReport200ResponseSnapshotsInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *GetEmqReport200ResponseSnapshotsInner) SetScore(v float32)`

SetScore sets Score field to given value.


### GetWeekOverWeekChange

`func (o *GetEmqReport200ResponseSnapshotsInner) GetWeekOverWeekChange() float32`

GetWeekOverWeekChange returns the WeekOverWeekChange field if non-nil, zero value otherwise.

### GetWeekOverWeekChangeOk

`func (o *GetEmqReport200ResponseSnapshotsInner) GetWeekOverWeekChangeOk() (*float32, bool)`

GetWeekOverWeekChangeOk returns a tuple with the WeekOverWeekChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekOverWeekChange

`func (o *GetEmqReport200ResponseSnapshotsInner) SetWeekOverWeekChange(v float32)`

SetWeekOverWeekChange sets WeekOverWeekChange field to given value.


### SetWeekOverWeekChangeNil

`func (o *GetEmqReport200ResponseSnapshotsInner) SetWeekOverWeekChangeNil(b bool)`

 SetWeekOverWeekChangeNil sets the value for WeekOverWeekChange to be an explicit nil

### UnsetWeekOverWeekChange
`func (o *GetEmqReport200ResponseSnapshotsInner) UnsetWeekOverWeekChange()`

UnsetWeekOverWeekChange ensures that no value is present for WeekOverWeekChange, not even an explicit nil
### GetAlerted

`func (o *GetEmqReport200ResponseSnapshotsInner) GetAlerted() bool`

GetAlerted returns the Alerted field if non-nil, zero value otherwise.

### GetAlertedOk

`func (o *GetEmqReport200ResponseSnapshotsInner) GetAlertedOk() (*bool, bool)`

GetAlertedOk returns a tuple with the Alerted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerted

`func (o *GetEmqReport200ResponseSnapshotsInner) SetAlerted(v bool)`

SetAlerted sets Alerted field to given value.


### GetPlatformResponse

`func (o *GetEmqReport200ResponseSnapshotsInner) GetPlatformResponse() interface{}`

GetPlatformResponse returns the PlatformResponse field if non-nil, zero value otherwise.

### GetPlatformResponseOk

`func (o *GetEmqReport200ResponseSnapshotsInner) GetPlatformResponseOk() (*interface{}, bool)`

GetPlatformResponseOk returns a tuple with the PlatformResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformResponse

`func (o *GetEmqReport200ResponseSnapshotsInner) SetPlatformResponse(v interface{})`

SetPlatformResponse sets PlatformResponse field to given value.


### SetPlatformResponseNil

`func (o *GetEmqReport200ResponseSnapshotsInner) SetPlatformResponseNil(b bool)`

 SetPlatformResponseNil sets the value for PlatformResponse to be an explicit nil

### UnsetPlatformResponse
`func (o *GetEmqReport200ResponseSnapshotsInner) UnsetPlatformResponse()`

UnsetPlatformResponse ensures that no value is present for PlatformResponse, not even an explicit nil
### GetMeasuredAt

`func (o *GetEmqReport200ResponseSnapshotsInner) GetMeasuredAt() string`

GetMeasuredAt returns the MeasuredAt field if non-nil, zero value otherwise.

### GetMeasuredAtOk

`func (o *GetEmqReport200ResponseSnapshotsInner) GetMeasuredAtOk() (*string, bool)`

GetMeasuredAtOk returns a tuple with the MeasuredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredAt

`func (o *GetEmqReport200ResponseSnapshotsInner) SetMeasuredAt(v string)`

SetMeasuredAt sets MeasuredAt field to given value.


### GetCreatedAt

`func (o *GetEmqReport200ResponseSnapshotsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetEmqReport200ResponseSnapshotsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetEmqReport200ResponseSnapshotsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *GetEmqReport200ResponseSnapshotsInner) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GetEmqReport200ResponseSnapshotsInner) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GetEmqReport200ResponseSnapshotsInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetEmqReport200ResponseSnapshotsInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetEmqReport200ResponseSnapshotsInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *GetEmqReport200ResponseSnapshotsInner) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GetEmqReport200ResponseSnapshotsInner) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


