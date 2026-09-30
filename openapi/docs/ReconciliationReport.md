# ReconciliationReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**WorkspaceId** | **int32** |  | 
**SignalTrackerId** | **string** | Deprecated alias of workspace_id; contains the workspace ID in decimal string form. | 
**DestinationId** | **string** |  | 
**ReportDate** | **time.Time** |  | 
**AcceptedCount** | **int32** |  | 
**MetaCount** | **int32** |  | 
**ObservedGap** | **int32** |  | 
**EventCounts** | **[]interface{}** |  | 
**Buckets** | **[]interface{}** |  | 
**UnexplainedResidual** | **int32** |  | 
**Status** | **string** |  | 
**CreatedAt** | **NullableTime** |  | 
**UpdatedAt** | **NullableTime** |  | 
**ClaimedClicks** | **NullableInt32** | Meta outbound clicks. Days before 2026-06-29, or not re-read by the daily sync since 2026-09-27, may still hold Meta link clicks or all clicks. | 

## Methods

### NewReconciliationReport

`func NewReconciliationReport(id int32, workspaceId int32, signalTrackerId string, destinationId string, reportDate time.Time, acceptedCount int32, metaCount int32, observedGap int32, eventCounts []interface{}, buckets []interface{}, unexplainedResidual int32, status string, createdAt NullableTime, updatedAt NullableTime, claimedClicks NullableInt32, ) *ReconciliationReport`

NewReconciliationReport instantiates a new ReconciliationReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReconciliationReportWithDefaults

`func NewReconciliationReportWithDefaults() *ReconciliationReport`

NewReconciliationReportWithDefaults instantiates a new ReconciliationReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReconciliationReport) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReconciliationReport) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReconciliationReport) SetId(v int32)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *ReconciliationReport) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ReconciliationReport) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ReconciliationReport) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetSignalTrackerId

`func (o *ReconciliationReport) GetSignalTrackerId() string`

GetSignalTrackerId returns the SignalTrackerId field if non-nil, zero value otherwise.

### GetSignalTrackerIdOk

`func (o *ReconciliationReport) GetSignalTrackerIdOk() (*string, bool)`

GetSignalTrackerIdOk returns a tuple with the SignalTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalTrackerId

`func (o *ReconciliationReport) SetSignalTrackerId(v string)`

SetSignalTrackerId sets SignalTrackerId field to given value.


### GetDestinationId

`func (o *ReconciliationReport) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *ReconciliationReport) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *ReconciliationReport) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetReportDate

`func (o *ReconciliationReport) GetReportDate() time.Time`

GetReportDate returns the ReportDate field if non-nil, zero value otherwise.

### GetReportDateOk

`func (o *ReconciliationReport) GetReportDateOk() (*time.Time, bool)`

GetReportDateOk returns a tuple with the ReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDate

`func (o *ReconciliationReport) SetReportDate(v time.Time)`

SetReportDate sets ReportDate field to given value.


### GetAcceptedCount

`func (o *ReconciliationReport) GetAcceptedCount() int32`

GetAcceptedCount returns the AcceptedCount field if non-nil, zero value otherwise.

### GetAcceptedCountOk

`func (o *ReconciliationReport) GetAcceptedCountOk() (*int32, bool)`

GetAcceptedCountOk returns a tuple with the AcceptedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCount

`func (o *ReconciliationReport) SetAcceptedCount(v int32)`

SetAcceptedCount sets AcceptedCount field to given value.


### GetMetaCount

`func (o *ReconciliationReport) GetMetaCount() int32`

GetMetaCount returns the MetaCount field if non-nil, zero value otherwise.

### GetMetaCountOk

`func (o *ReconciliationReport) GetMetaCountOk() (*int32, bool)`

GetMetaCountOk returns a tuple with the MetaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaCount

`func (o *ReconciliationReport) SetMetaCount(v int32)`

SetMetaCount sets MetaCount field to given value.


### GetObservedGap

`func (o *ReconciliationReport) GetObservedGap() int32`

GetObservedGap returns the ObservedGap field if non-nil, zero value otherwise.

### GetObservedGapOk

`func (o *ReconciliationReport) GetObservedGapOk() (*int32, bool)`

GetObservedGapOk returns a tuple with the ObservedGap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGap

`func (o *ReconciliationReport) SetObservedGap(v int32)`

SetObservedGap sets ObservedGap field to given value.


### GetEventCounts

`func (o *ReconciliationReport) GetEventCounts() []interface{}`

GetEventCounts returns the EventCounts field if non-nil, zero value otherwise.

### GetEventCountsOk

`func (o *ReconciliationReport) GetEventCountsOk() (*[]interface{}, bool)`

GetEventCountsOk returns a tuple with the EventCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCounts

`func (o *ReconciliationReport) SetEventCounts(v []interface{})`

SetEventCounts sets EventCounts field to given value.


### GetBuckets

`func (o *ReconciliationReport) GetBuckets() []interface{}`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *ReconciliationReport) GetBucketsOk() (*[]interface{}, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *ReconciliationReport) SetBuckets(v []interface{})`

SetBuckets sets Buckets field to given value.


### GetUnexplainedResidual

`func (o *ReconciliationReport) GetUnexplainedResidual() int32`

GetUnexplainedResidual returns the UnexplainedResidual field if non-nil, zero value otherwise.

### GetUnexplainedResidualOk

`func (o *ReconciliationReport) GetUnexplainedResidualOk() (*int32, bool)`

GetUnexplainedResidualOk returns a tuple with the UnexplainedResidual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnexplainedResidual

`func (o *ReconciliationReport) SetUnexplainedResidual(v int32)`

SetUnexplainedResidual sets UnexplainedResidual field to given value.


### GetStatus

`func (o *ReconciliationReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReconciliationReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReconciliationReport) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ReconciliationReport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReconciliationReport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReconciliationReport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *ReconciliationReport) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ReconciliationReport) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *ReconciliationReport) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReconciliationReport) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReconciliationReport) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *ReconciliationReport) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ReconciliationReport) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetClaimedClicks

`func (o *ReconciliationReport) GetClaimedClicks() int32`

GetClaimedClicks returns the ClaimedClicks field if non-nil, zero value otherwise.

### GetClaimedClicksOk

`func (o *ReconciliationReport) GetClaimedClicksOk() (*int32, bool)`

GetClaimedClicksOk returns a tuple with the ClaimedClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedClicks

`func (o *ReconciliationReport) SetClaimedClicks(v int32)`

SetClaimedClicks sets ClaimedClicks field to given value.


### SetClaimedClicksNil

`func (o *ReconciliationReport) SetClaimedClicksNil(b bool)`

 SetClaimedClicksNil sets the value for ClaimedClicks to be an explicit nil

### UnsetClaimedClicks
`func (o *ReconciliationReport) UnsetClaimedClicks()`

UnsetClaimedClicks ensures that no value is present for ClaimedClicks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


