# LaunchPlansIndex200ResponsePlans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DeploymentPlanRead**](DeploymentPlanRead.md) |  | 
**CurrentPage** | **int32** |  | 
**PerPage** | **int32** |  | 
**Total** | **int32** |  | 
**LastPage** | **int32** |  | 
**NextCursor** | **NullableInt32** | Always present. The next row offset, or null when the list is complete. | 
**Truncated** | **bool** |  | 
**NextPageArguments** | [**NullableLaunchPlansIndex200ResponsePlansNextPageArguments**](LaunchPlansIndex200ResponsePlansNextPageArguments.md) |  | 
**Continuation** | **string** |  | 

## Methods

### NewLaunchPlansIndex200ResponsePlans

`func NewLaunchPlansIndex200ResponsePlans(data []DeploymentPlanRead, currentPage int32, perPage int32, total int32, lastPage int32, nextCursor NullableInt32, truncated bool, nextPageArguments NullableLaunchPlansIndex200ResponsePlansNextPageArguments, continuation string, ) *LaunchPlansIndex200ResponsePlans`

NewLaunchPlansIndex200ResponsePlans instantiates a new LaunchPlansIndex200ResponsePlans object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchPlansIndex200ResponsePlansWithDefaults

`func NewLaunchPlansIndex200ResponsePlansWithDefaults() *LaunchPlansIndex200ResponsePlans`

NewLaunchPlansIndex200ResponsePlansWithDefaults instantiates a new LaunchPlansIndex200ResponsePlans object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LaunchPlansIndex200ResponsePlans) GetData() []DeploymentPlanRead`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LaunchPlansIndex200ResponsePlans) GetDataOk() (*[]DeploymentPlanRead, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LaunchPlansIndex200ResponsePlans) SetData(v []DeploymentPlanRead)`

SetData sets Data field to given value.


### GetCurrentPage

`func (o *LaunchPlansIndex200ResponsePlans) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *LaunchPlansIndex200ResponsePlans) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *LaunchPlansIndex200ResponsePlans) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetPerPage

`func (o *LaunchPlansIndex200ResponsePlans) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *LaunchPlansIndex200ResponsePlans) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *LaunchPlansIndex200ResponsePlans) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetTotal

`func (o *LaunchPlansIndex200ResponsePlans) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LaunchPlansIndex200ResponsePlans) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LaunchPlansIndex200ResponsePlans) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetLastPage

`func (o *LaunchPlansIndex200ResponsePlans) GetLastPage() int32`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *LaunchPlansIndex200ResponsePlans) GetLastPageOk() (*int32, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *LaunchPlansIndex200ResponsePlans) SetLastPage(v int32)`

SetLastPage sets LastPage field to given value.


### GetNextCursor

`func (o *LaunchPlansIndex200ResponsePlans) GetNextCursor() int32`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *LaunchPlansIndex200ResponsePlans) GetNextCursorOk() (*int32, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *LaunchPlansIndex200ResponsePlans) SetNextCursor(v int32)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *LaunchPlansIndex200ResponsePlans) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *LaunchPlansIndex200ResponsePlans) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetTruncated

`func (o *LaunchPlansIndex200ResponsePlans) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *LaunchPlansIndex200ResponsePlans) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *LaunchPlansIndex200ResponsePlans) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.


### GetNextPageArguments

`func (o *LaunchPlansIndex200ResponsePlans) GetNextPageArguments() LaunchPlansIndex200ResponsePlansNextPageArguments`

GetNextPageArguments returns the NextPageArguments field if non-nil, zero value otherwise.

### GetNextPageArgumentsOk

`func (o *LaunchPlansIndex200ResponsePlans) GetNextPageArgumentsOk() (*LaunchPlansIndex200ResponsePlansNextPageArguments, bool)`

GetNextPageArgumentsOk returns a tuple with the NextPageArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageArguments

`func (o *LaunchPlansIndex200ResponsePlans) SetNextPageArguments(v LaunchPlansIndex200ResponsePlansNextPageArguments)`

SetNextPageArguments sets NextPageArguments field to given value.


### SetNextPageArgumentsNil

`func (o *LaunchPlansIndex200ResponsePlans) SetNextPageArgumentsNil(b bool)`

 SetNextPageArgumentsNil sets the value for NextPageArguments to be an explicit nil

### UnsetNextPageArguments
`func (o *LaunchPlansIndex200ResponsePlans) UnsetNextPageArguments()`

UnsetNextPageArguments ensures that no value is present for NextPageArguments, not even an explicit nil
### GetContinuation

`func (o *LaunchPlansIndex200ResponsePlans) GetContinuation() string`

GetContinuation returns the Continuation field if non-nil, zero value otherwise.

### GetContinuationOk

`func (o *LaunchPlansIndex200ResponsePlans) GetContinuationOk() (*string, bool)`

GetContinuationOk returns a tuple with the Continuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuation

`func (o *LaunchPlansIndex200ResponsePlans) SetContinuation(v string)`

SetContinuation sets Continuation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


