# EmqSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**SignalTrackerId** | **string** |  | 
**DestinationId** | **string** |  | 
**Score** | **float32** |  | 
**WeekOverWeekChange** | **NullableFloat32** |  | 
**Alerted** | **bool** |  | 
**PlatformResponse** | **[]interface{}** |  | 
**MeasuredAt** | **time.Time** |  | 
**CreatedAt** | **NullableTime** |  | 
**UpdatedAt** | **NullableTime** |  | 

## Methods

### NewEmqSnapshot

`func NewEmqSnapshot(id int32, signalTrackerId string, destinationId string, score float32, weekOverWeekChange NullableFloat32, alerted bool, platformResponse []interface{}, measuredAt time.Time, createdAt NullableTime, updatedAt NullableTime, ) *EmqSnapshot`

NewEmqSnapshot instantiates a new EmqSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmqSnapshotWithDefaults

`func NewEmqSnapshotWithDefaults() *EmqSnapshot`

NewEmqSnapshotWithDefaults instantiates a new EmqSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmqSnapshot) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmqSnapshot) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmqSnapshot) SetId(v int32)`

SetId sets Id field to given value.


### GetSignalTrackerId

`func (o *EmqSnapshot) GetSignalTrackerId() string`

GetSignalTrackerId returns the SignalTrackerId field if non-nil, zero value otherwise.

### GetSignalTrackerIdOk

`func (o *EmqSnapshot) GetSignalTrackerIdOk() (*string, bool)`

GetSignalTrackerIdOk returns a tuple with the SignalTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalTrackerId

`func (o *EmqSnapshot) SetSignalTrackerId(v string)`

SetSignalTrackerId sets SignalTrackerId field to given value.


### GetDestinationId

`func (o *EmqSnapshot) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *EmqSnapshot) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *EmqSnapshot) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetScore

`func (o *EmqSnapshot) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *EmqSnapshot) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *EmqSnapshot) SetScore(v float32)`

SetScore sets Score field to given value.


### GetWeekOverWeekChange

`func (o *EmqSnapshot) GetWeekOverWeekChange() float32`

GetWeekOverWeekChange returns the WeekOverWeekChange field if non-nil, zero value otherwise.

### GetWeekOverWeekChangeOk

`func (o *EmqSnapshot) GetWeekOverWeekChangeOk() (*float32, bool)`

GetWeekOverWeekChangeOk returns a tuple with the WeekOverWeekChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekOverWeekChange

`func (o *EmqSnapshot) SetWeekOverWeekChange(v float32)`

SetWeekOverWeekChange sets WeekOverWeekChange field to given value.


### SetWeekOverWeekChangeNil

`func (o *EmqSnapshot) SetWeekOverWeekChangeNil(b bool)`

 SetWeekOverWeekChangeNil sets the value for WeekOverWeekChange to be an explicit nil

### UnsetWeekOverWeekChange
`func (o *EmqSnapshot) UnsetWeekOverWeekChange()`

UnsetWeekOverWeekChange ensures that no value is present for WeekOverWeekChange, not even an explicit nil
### GetAlerted

`func (o *EmqSnapshot) GetAlerted() bool`

GetAlerted returns the Alerted field if non-nil, zero value otherwise.

### GetAlertedOk

`func (o *EmqSnapshot) GetAlertedOk() (*bool, bool)`

GetAlertedOk returns a tuple with the Alerted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerted

`func (o *EmqSnapshot) SetAlerted(v bool)`

SetAlerted sets Alerted field to given value.


### GetPlatformResponse

`func (o *EmqSnapshot) GetPlatformResponse() []interface{}`

GetPlatformResponse returns the PlatformResponse field if non-nil, zero value otherwise.

### GetPlatformResponseOk

`func (o *EmqSnapshot) GetPlatformResponseOk() (*[]interface{}, bool)`

GetPlatformResponseOk returns a tuple with the PlatformResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformResponse

`func (o *EmqSnapshot) SetPlatformResponse(v []interface{})`

SetPlatformResponse sets PlatformResponse field to given value.


### SetPlatformResponseNil

`func (o *EmqSnapshot) SetPlatformResponseNil(b bool)`

 SetPlatformResponseNil sets the value for PlatformResponse to be an explicit nil

### UnsetPlatformResponse
`func (o *EmqSnapshot) UnsetPlatformResponse()`

UnsetPlatformResponse ensures that no value is present for PlatformResponse, not even an explicit nil
### GetMeasuredAt

`func (o *EmqSnapshot) GetMeasuredAt() time.Time`

GetMeasuredAt returns the MeasuredAt field if non-nil, zero value otherwise.

### GetMeasuredAtOk

`func (o *EmqSnapshot) GetMeasuredAtOk() (*time.Time, bool)`

GetMeasuredAtOk returns a tuple with the MeasuredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasuredAt

`func (o *EmqSnapshot) SetMeasuredAt(v time.Time)`

SetMeasuredAt sets MeasuredAt field to given value.


### GetCreatedAt

`func (o *EmqSnapshot) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EmqSnapshot) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EmqSnapshot) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *EmqSnapshot) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *EmqSnapshot) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *EmqSnapshot) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EmqSnapshot) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EmqSnapshot) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *EmqSnapshot) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *EmqSnapshot) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


