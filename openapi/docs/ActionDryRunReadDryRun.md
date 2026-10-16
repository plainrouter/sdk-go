# ActionDryRunReadDryRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**[]ActionDryRunReadDryRunActionsInner**](ActionDryRunReadDryRunActionsInner.md) |  | 
**NextCursor** | Pointer to **NullableInt32** |  | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] 
**Continuation** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewActionDryRunReadDryRun

`func NewActionDryRunReadDryRun(actions []ActionDryRunReadDryRunActionsInner, ) *ActionDryRunReadDryRun`

NewActionDryRunReadDryRun instantiates a new ActionDryRunReadDryRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionDryRunReadDryRunWithDefaults

`func NewActionDryRunReadDryRunWithDefaults() *ActionDryRunReadDryRun`

NewActionDryRunReadDryRunWithDefaults instantiates a new ActionDryRunReadDryRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ActionDryRunReadDryRun) GetActions() []ActionDryRunReadDryRunActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ActionDryRunReadDryRun) GetActionsOk() (*[]ActionDryRunReadDryRunActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ActionDryRunReadDryRun) SetActions(v []ActionDryRunReadDryRunActionsInner)`

SetActions sets Actions field to given value.


### GetNextCursor

`func (o *ActionDryRunReadDryRun) GetNextCursor() int32`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ActionDryRunReadDryRun) GetNextCursorOk() (*int32, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ActionDryRunReadDryRun) SetNextCursor(v int32)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ActionDryRunReadDryRun) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *ActionDryRunReadDryRun) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ActionDryRunReadDryRun) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetTruncated

`func (o *ActionDryRunReadDryRun) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *ActionDryRunReadDryRun) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *ActionDryRunReadDryRun) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *ActionDryRunReadDryRun) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetContinuation

`func (o *ActionDryRunReadDryRun) GetContinuation() string`

GetContinuation returns the Continuation field if non-nil, zero value otherwise.

### GetContinuationOk

`func (o *ActionDryRunReadDryRun) GetContinuationOk() (*string, bool)`

GetContinuationOk returns a tuple with the Continuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuation

`func (o *ActionDryRunReadDryRun) SetContinuation(v string)`

SetContinuation sets Continuation field to given value.

### HasContinuation

`func (o *ActionDryRunReadDryRun) HasContinuation() bool`

HasContinuation returns a boolean if a field has been set.

### GetTotal

`func (o *ActionDryRunReadDryRun) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ActionDryRunReadDryRun) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ActionDryRunReadDryRun) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ActionDryRunReadDryRun) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


