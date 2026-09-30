# GetEvent200ResponseEventDeliveriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**WorkspaceId** | **int32** |  | 
**SignalTrackerId** | **string** | Deprecated alias of workspace_id; contains the workspace ID in decimal string form. | 
**EventId** | **string** |  | 
**DestinationId** | **NullableString** |  | 
**Status** | [**DeliveryStatus**](DeliveryStatus.md) |  | 
**IsTest** | **bool** |  | 
**AttemptCount** | **int32** |  | 
**LastError** | **interface{}** |  | 
**PlatformResponse** | **interface{}** |  | 
**PlatformTraceId** | **NullableString** |  | 
**NextAttemptAt** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **NullableString** |  | 

## Methods

### NewGetEvent200ResponseEventDeliveriesInner

`func NewGetEvent200ResponseEventDeliveriesInner(id int32, workspaceId int32, signalTrackerId string, eventId string, destinationId NullableString, status DeliveryStatus, isTest bool, attemptCount int32, lastError interface{}, platformResponse interface{}, platformTraceId NullableString, nextAttemptAt NullableString, createdAt string, updatedAt NullableString, ) *GetEvent200ResponseEventDeliveriesInner`

NewGetEvent200ResponseEventDeliveriesInner instantiates a new GetEvent200ResponseEventDeliveriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvent200ResponseEventDeliveriesInnerWithDefaults

`func NewGetEvent200ResponseEventDeliveriesInnerWithDefaults() *GetEvent200ResponseEventDeliveriesInner`

NewGetEvent200ResponseEventDeliveriesInnerWithDefaults instantiates a new GetEvent200ResponseEventDeliveriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEvent200ResponseEventDeliveriesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEvent200ResponseEventDeliveriesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEvent200ResponseEventDeliveriesInner) SetId(v int32)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *GetEvent200ResponseEventDeliveriesInner) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GetEvent200ResponseEventDeliveriesInner) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GetEvent200ResponseEventDeliveriesInner) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetSignalTrackerId

`func (o *GetEvent200ResponseEventDeliveriesInner) GetSignalTrackerId() string`

GetSignalTrackerId returns the SignalTrackerId field if non-nil, zero value otherwise.

### GetSignalTrackerIdOk

`func (o *GetEvent200ResponseEventDeliveriesInner) GetSignalTrackerIdOk() (*string, bool)`

GetSignalTrackerIdOk returns a tuple with the SignalTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalTrackerId

`func (o *GetEvent200ResponseEventDeliveriesInner) SetSignalTrackerId(v string)`

SetSignalTrackerId sets SignalTrackerId field to given value.


### GetEventId

`func (o *GetEvent200ResponseEventDeliveriesInner) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *GetEvent200ResponseEventDeliveriesInner) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *GetEvent200ResponseEventDeliveriesInner) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetDestinationId

`func (o *GetEvent200ResponseEventDeliveriesInner) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *GetEvent200ResponseEventDeliveriesInner) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *GetEvent200ResponseEventDeliveriesInner) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### SetDestinationIdNil

`func (o *GetEvent200ResponseEventDeliveriesInner) SetDestinationIdNil(b bool)`

 SetDestinationIdNil sets the value for DestinationId to be an explicit nil

### UnsetDestinationId
`func (o *GetEvent200ResponseEventDeliveriesInner) UnsetDestinationId()`

UnsetDestinationId ensures that no value is present for DestinationId, not even an explicit nil
### GetStatus

`func (o *GetEvent200ResponseEventDeliveriesInner) GetStatus() DeliveryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEvent200ResponseEventDeliveriesInner) GetStatusOk() (*DeliveryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEvent200ResponseEventDeliveriesInner) SetStatus(v DeliveryStatus)`

SetStatus sets Status field to given value.


### GetIsTest

`func (o *GetEvent200ResponseEventDeliveriesInner) GetIsTest() bool`

GetIsTest returns the IsTest field if non-nil, zero value otherwise.

### GetIsTestOk

`func (o *GetEvent200ResponseEventDeliveriesInner) GetIsTestOk() (*bool, bool)`

GetIsTestOk returns a tuple with the IsTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTest

`func (o *GetEvent200ResponseEventDeliveriesInner) SetIsTest(v bool)`

SetIsTest sets IsTest field to given value.


### GetAttemptCount

`func (o *GetEvent200ResponseEventDeliveriesInner) GetAttemptCount() int32`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *GetEvent200ResponseEventDeliveriesInner) GetAttemptCountOk() (*int32, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *GetEvent200ResponseEventDeliveriesInner) SetAttemptCount(v int32)`

SetAttemptCount sets AttemptCount field to given value.


### GetLastError

`func (o *GetEvent200ResponseEventDeliveriesInner) GetLastError() interface{}`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *GetEvent200ResponseEventDeliveriesInner) GetLastErrorOk() (*interface{}, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *GetEvent200ResponseEventDeliveriesInner) SetLastError(v interface{})`

SetLastError sets LastError field to given value.


### SetLastErrorNil

`func (o *GetEvent200ResponseEventDeliveriesInner) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *GetEvent200ResponseEventDeliveriesInner) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetPlatformResponse

`func (o *GetEvent200ResponseEventDeliveriesInner) GetPlatformResponse() interface{}`

GetPlatformResponse returns the PlatformResponse field if non-nil, zero value otherwise.

### GetPlatformResponseOk

`func (o *GetEvent200ResponseEventDeliveriesInner) GetPlatformResponseOk() (*interface{}, bool)`

GetPlatformResponseOk returns a tuple with the PlatformResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformResponse

`func (o *GetEvent200ResponseEventDeliveriesInner) SetPlatformResponse(v interface{})`

SetPlatformResponse sets PlatformResponse field to given value.


### SetPlatformResponseNil

`func (o *GetEvent200ResponseEventDeliveriesInner) SetPlatformResponseNil(b bool)`

 SetPlatformResponseNil sets the value for PlatformResponse to be an explicit nil

### UnsetPlatformResponse
`func (o *GetEvent200ResponseEventDeliveriesInner) UnsetPlatformResponse()`

UnsetPlatformResponse ensures that no value is present for PlatformResponse, not even an explicit nil
### GetPlatformTraceId

`func (o *GetEvent200ResponseEventDeliveriesInner) GetPlatformTraceId() string`

GetPlatformTraceId returns the PlatformTraceId field if non-nil, zero value otherwise.

### GetPlatformTraceIdOk

`func (o *GetEvent200ResponseEventDeliveriesInner) GetPlatformTraceIdOk() (*string, bool)`

GetPlatformTraceIdOk returns a tuple with the PlatformTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformTraceId

`func (o *GetEvent200ResponseEventDeliveriesInner) SetPlatformTraceId(v string)`

SetPlatformTraceId sets PlatformTraceId field to given value.


### SetPlatformTraceIdNil

`func (o *GetEvent200ResponseEventDeliveriesInner) SetPlatformTraceIdNil(b bool)`

 SetPlatformTraceIdNil sets the value for PlatformTraceId to be an explicit nil

### UnsetPlatformTraceId
`func (o *GetEvent200ResponseEventDeliveriesInner) UnsetPlatformTraceId()`

UnsetPlatformTraceId ensures that no value is present for PlatformTraceId, not even an explicit nil
### GetNextAttemptAt

`func (o *GetEvent200ResponseEventDeliveriesInner) GetNextAttemptAt() string`

GetNextAttemptAt returns the NextAttemptAt field if non-nil, zero value otherwise.

### GetNextAttemptAtOk

`func (o *GetEvent200ResponseEventDeliveriesInner) GetNextAttemptAtOk() (*string, bool)`

GetNextAttemptAtOk returns a tuple with the NextAttemptAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAttemptAt

`func (o *GetEvent200ResponseEventDeliveriesInner) SetNextAttemptAt(v string)`

SetNextAttemptAt sets NextAttemptAt field to given value.


### SetNextAttemptAtNil

`func (o *GetEvent200ResponseEventDeliveriesInner) SetNextAttemptAtNil(b bool)`

 SetNextAttemptAtNil sets the value for NextAttemptAt to be an explicit nil

### UnsetNextAttemptAt
`func (o *GetEvent200ResponseEventDeliveriesInner) UnsetNextAttemptAt()`

UnsetNextAttemptAt ensures that no value is present for NextAttemptAt, not even an explicit nil
### GetCreatedAt

`func (o *GetEvent200ResponseEventDeliveriesInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetEvent200ResponseEventDeliveriesInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetEvent200ResponseEventDeliveriesInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetEvent200ResponseEventDeliveriesInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetEvent200ResponseEventDeliveriesInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetEvent200ResponseEventDeliveriesInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *GetEvent200ResponseEventDeliveriesInner) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GetEvent200ResponseEventDeliveriesInner) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


