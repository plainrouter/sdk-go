# GetInventoryMetrics200ResponsePaginationValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** |  | 
**Truncated** | **bool** |  | 
**NextCursor** | **NullableInt32** |  | 

## Methods

### NewGetInventoryMetrics200ResponsePaginationValue

`func NewGetInventoryMetrics200ResponsePaginationValue(total int32, truncated bool, nextCursor NullableInt32, ) *GetInventoryMetrics200ResponsePaginationValue`

NewGetInventoryMetrics200ResponsePaginationValue instantiates a new GetInventoryMetrics200ResponsePaginationValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInventoryMetrics200ResponsePaginationValueWithDefaults

`func NewGetInventoryMetrics200ResponsePaginationValueWithDefaults() *GetInventoryMetrics200ResponsePaginationValue`

NewGetInventoryMetrics200ResponsePaginationValueWithDefaults instantiates a new GetInventoryMetrics200ResponsePaginationValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetInventoryMetrics200ResponsePaginationValue) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetInventoryMetrics200ResponsePaginationValue) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetInventoryMetrics200ResponsePaginationValue) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetTruncated

`func (o *GetInventoryMetrics200ResponsePaginationValue) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GetInventoryMetrics200ResponsePaginationValue) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GetInventoryMetrics200ResponsePaginationValue) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.


### GetNextCursor

`func (o *GetInventoryMetrics200ResponsePaginationValue) GetNextCursor() int32`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *GetInventoryMetrics200ResponsePaginationValue) GetNextCursorOk() (*int32, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *GetInventoryMetrics200ResponsePaginationValue) SetNextCursor(v int32)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *GetInventoryMetrics200ResponsePaginationValue) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *GetInventoryMetrics200ResponsePaginationValue) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


