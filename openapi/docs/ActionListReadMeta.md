# ActionListReadMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | **int32** |  | 
**LastPage** | **int32** |  | 
**PerPage** | **int32** |  | 
**Total** | **int32** |  | 
**NextCursor** | Pointer to **NullableInt32** |  | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] 
**Continuation** | Pointer to **string** |  | [optional] 

## Methods

### NewActionListReadMeta

`func NewActionListReadMeta(currentPage int32, lastPage int32, perPage int32, total int32, ) *ActionListReadMeta`

NewActionListReadMeta instantiates a new ActionListReadMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionListReadMetaWithDefaults

`func NewActionListReadMetaWithDefaults() *ActionListReadMeta`

NewActionListReadMetaWithDefaults instantiates a new ActionListReadMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *ActionListReadMeta) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ActionListReadMeta) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ActionListReadMeta) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetLastPage

`func (o *ActionListReadMeta) GetLastPage() int32`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *ActionListReadMeta) GetLastPageOk() (*int32, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *ActionListReadMeta) SetLastPage(v int32)`

SetLastPage sets LastPage field to given value.


### GetPerPage

`func (o *ActionListReadMeta) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ActionListReadMeta) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ActionListReadMeta) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetTotal

`func (o *ActionListReadMeta) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ActionListReadMeta) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ActionListReadMeta) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetNextCursor

`func (o *ActionListReadMeta) GetNextCursor() int32`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ActionListReadMeta) GetNextCursorOk() (*int32, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ActionListReadMeta) SetNextCursor(v int32)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ActionListReadMeta) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *ActionListReadMeta) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ActionListReadMeta) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetTruncated

`func (o *ActionListReadMeta) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *ActionListReadMeta) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *ActionListReadMeta) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *ActionListReadMeta) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetContinuation

`func (o *ActionListReadMeta) GetContinuation() string`

GetContinuation returns the Continuation field if non-nil, zero value otherwise.

### GetContinuationOk

`func (o *ActionListReadMeta) GetContinuationOk() (*string, bool)`

GetContinuationOk returns a tuple with the Continuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuation

`func (o *ActionListReadMeta) SetContinuation(v string)`

SetContinuation sets Continuation field to given value.

### HasContinuation

`func (o *ActionListReadMeta) HasContinuation() bool`

HasContinuation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


