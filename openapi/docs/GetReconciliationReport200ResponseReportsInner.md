# GetReconciliationReport200ResponseReportsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**WorkspaceId** | **int32** |  | 
**SignalTrackerId** | **string** | Deprecated alias of workspace_id; contains the workspace ID in decimal string form. | 
**DestinationId** | **string** |  | 
**ReportDate** | **string** |  | 
**AcceptedCount** | **int32** |  | 
**MetaCount** | **int32** |  | 
**ObservedGap** | **int32** |  | 
**EventCounts** | [**GetReconciliationReport200ResponseReportsInnerEventCounts**](GetReconciliationReport200ResponseReportsInnerEventCounts.md) |  | 
**Buckets** | [**map[string]GetReconciliationReport200ResponseReportsInnerBucketsValue**](GetReconciliationReport200ResponseReportsInnerBucketsValue.md) |  | 
**UnexplainedResidual** | **int32** |  | 
**Status** | **string** |  | 
**CreatedAt** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 
**Destination** | [**SetDestinationTestMode200ResponseDestination**](SetDestinationTestMode200ResponseDestination.md) |  | 

## Methods

### NewGetReconciliationReport200ResponseReportsInner

`func NewGetReconciliationReport200ResponseReportsInner(id int32, workspaceId int32, signalTrackerId string, destinationId string, reportDate string, acceptedCount int32, metaCount int32, observedGap int32, eventCounts GetReconciliationReport200ResponseReportsInnerEventCounts, buckets map[string]GetReconciliationReport200ResponseReportsInnerBucketsValue, unexplainedResidual int32, status string, createdAt NullableString, updatedAt NullableString, destination SetDestinationTestMode200ResponseDestination, ) *GetReconciliationReport200ResponseReportsInner`

NewGetReconciliationReport200ResponseReportsInner instantiates a new GetReconciliationReport200ResponseReportsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReconciliationReport200ResponseReportsInnerWithDefaults

`func NewGetReconciliationReport200ResponseReportsInnerWithDefaults() *GetReconciliationReport200ResponseReportsInner`

NewGetReconciliationReport200ResponseReportsInnerWithDefaults instantiates a new GetReconciliationReport200ResponseReportsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetReconciliationReport200ResponseReportsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReconciliationReport200ResponseReportsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReconciliationReport200ResponseReportsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *GetReconciliationReport200ResponseReportsInner) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *GetReconciliationReport200ResponseReportsInner) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *GetReconciliationReport200ResponseReportsInner) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetSignalTrackerId

`func (o *GetReconciliationReport200ResponseReportsInner) GetSignalTrackerId() string`

GetSignalTrackerId returns the SignalTrackerId field if non-nil, zero value otherwise.

### GetSignalTrackerIdOk

`func (o *GetReconciliationReport200ResponseReportsInner) GetSignalTrackerIdOk() (*string, bool)`

GetSignalTrackerIdOk returns a tuple with the SignalTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalTrackerId

`func (o *GetReconciliationReport200ResponseReportsInner) SetSignalTrackerId(v string)`

SetSignalTrackerId sets SignalTrackerId field to given value.


### GetDestinationId

`func (o *GetReconciliationReport200ResponseReportsInner) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *GetReconciliationReport200ResponseReportsInner) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *GetReconciliationReport200ResponseReportsInner) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetReportDate

`func (o *GetReconciliationReport200ResponseReportsInner) GetReportDate() string`

GetReportDate returns the ReportDate field if non-nil, zero value otherwise.

### GetReportDateOk

`func (o *GetReconciliationReport200ResponseReportsInner) GetReportDateOk() (*string, bool)`

GetReportDateOk returns a tuple with the ReportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDate

`func (o *GetReconciliationReport200ResponseReportsInner) SetReportDate(v string)`

SetReportDate sets ReportDate field to given value.


### GetAcceptedCount

`func (o *GetReconciliationReport200ResponseReportsInner) GetAcceptedCount() int32`

GetAcceptedCount returns the AcceptedCount field if non-nil, zero value otherwise.

### GetAcceptedCountOk

`func (o *GetReconciliationReport200ResponseReportsInner) GetAcceptedCountOk() (*int32, bool)`

GetAcceptedCountOk returns a tuple with the AcceptedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedCount

`func (o *GetReconciliationReport200ResponseReportsInner) SetAcceptedCount(v int32)`

SetAcceptedCount sets AcceptedCount field to given value.


### GetMetaCount

`func (o *GetReconciliationReport200ResponseReportsInner) GetMetaCount() int32`

GetMetaCount returns the MetaCount field if non-nil, zero value otherwise.

### GetMetaCountOk

`func (o *GetReconciliationReport200ResponseReportsInner) GetMetaCountOk() (*int32, bool)`

GetMetaCountOk returns a tuple with the MetaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaCount

`func (o *GetReconciliationReport200ResponseReportsInner) SetMetaCount(v int32)`

SetMetaCount sets MetaCount field to given value.


### GetObservedGap

`func (o *GetReconciliationReport200ResponseReportsInner) GetObservedGap() int32`

GetObservedGap returns the ObservedGap field if non-nil, zero value otherwise.

### GetObservedGapOk

`func (o *GetReconciliationReport200ResponseReportsInner) GetObservedGapOk() (*int32, bool)`

GetObservedGapOk returns a tuple with the ObservedGap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGap

`func (o *GetReconciliationReport200ResponseReportsInner) SetObservedGap(v int32)`

SetObservedGap sets ObservedGap field to given value.


### GetEventCounts

`func (o *GetReconciliationReport200ResponseReportsInner) GetEventCounts() GetReconciliationReport200ResponseReportsInnerEventCounts`

GetEventCounts returns the EventCounts field if non-nil, zero value otherwise.

### GetEventCountsOk

`func (o *GetReconciliationReport200ResponseReportsInner) GetEventCountsOk() (*GetReconciliationReport200ResponseReportsInnerEventCounts, bool)`

GetEventCountsOk returns a tuple with the EventCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCounts

`func (o *GetReconciliationReport200ResponseReportsInner) SetEventCounts(v GetReconciliationReport200ResponseReportsInnerEventCounts)`

SetEventCounts sets EventCounts field to given value.


### GetBuckets

`func (o *GetReconciliationReport200ResponseReportsInner) GetBuckets() map[string]GetReconciliationReport200ResponseReportsInnerBucketsValue`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *GetReconciliationReport200ResponseReportsInner) GetBucketsOk() (*map[string]GetReconciliationReport200ResponseReportsInnerBucketsValue, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *GetReconciliationReport200ResponseReportsInner) SetBuckets(v map[string]GetReconciliationReport200ResponseReportsInnerBucketsValue)`

SetBuckets sets Buckets field to given value.


### GetUnexplainedResidual

`func (o *GetReconciliationReport200ResponseReportsInner) GetUnexplainedResidual() int32`

GetUnexplainedResidual returns the UnexplainedResidual field if non-nil, zero value otherwise.

### GetUnexplainedResidualOk

`func (o *GetReconciliationReport200ResponseReportsInner) GetUnexplainedResidualOk() (*int32, bool)`

GetUnexplainedResidualOk returns a tuple with the UnexplainedResidual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnexplainedResidual

`func (o *GetReconciliationReport200ResponseReportsInner) SetUnexplainedResidual(v int32)`

SetUnexplainedResidual sets UnexplainedResidual field to given value.


### GetStatus

`func (o *GetReconciliationReport200ResponseReportsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetReconciliationReport200ResponseReportsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetReconciliationReport200ResponseReportsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *GetReconciliationReport200ResponseReportsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetReconciliationReport200ResponseReportsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetReconciliationReport200ResponseReportsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *GetReconciliationReport200ResponseReportsInner) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GetReconciliationReport200ResponseReportsInner) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GetReconciliationReport200ResponseReportsInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetReconciliationReport200ResponseReportsInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetReconciliationReport200ResponseReportsInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *GetReconciliationReport200ResponseReportsInner) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GetReconciliationReport200ResponseReportsInner) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetDestination

`func (o *GetReconciliationReport200ResponseReportsInner) GetDestination() SetDestinationTestMode200ResponseDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetReconciliationReport200ResponseReportsInner) GetDestinationOk() (*SetDestinationTestMode200ResponseDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetReconciliationReport200ResponseReportsInner) SetDestination(v SetDestinationTestMode200ResponseDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


