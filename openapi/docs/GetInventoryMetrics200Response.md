# GetInventoryMetrics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**GetInventoryMetrics200ResponseAccount**](GetInventoryMetrics200ResponseAccount.md) |  | 
**Sync** | [**GetInventoryMetrics200ResponseSync**](GetInventoryMetrics200ResponseSync.md) |  | 
**Staleness** | [**GetInventoryMetrics200ResponseStaleness**](GetInventoryMetrics200ResponseStaleness.md) |  | 
**Metrics** | [**[]GetInventoryMetrics200ResponseMetricsInner**](GetInventoryMetrics200ResponseMetricsInner.md) |  | 
**Changes** | [**[]GetInventoryMetrics200ResponseChangesInner**](GetInventoryMetrics200ResponseChangesInner.md) |  | 
**EffectiveRange** | [**NullableGetInventoryMetrics200ResponseEffectiveRange**](GetInventoryMetrics200ResponseEffectiveRange.md) |  | 
**ObjectMetrics** | [**[]GetInventoryMetrics200ResponseObjectMetricsInner**](GetInventoryMetrics200ResponseObjectMetricsInner.md) |  | 
**UnmatchedArrivals** | **NullableInt32** |  | 
**Pagination** | Pointer to [**map[string]GetInventoryMetrics200ResponsePaginationValue**](GetInventoryMetrics200ResponsePaginationValue.md) |  | [optional] 
**NextCursor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetInventoryMetrics200Response

`func NewGetInventoryMetrics200Response(account GetInventoryMetrics200ResponseAccount, sync GetInventoryMetrics200ResponseSync, staleness GetInventoryMetrics200ResponseStaleness, metrics []GetInventoryMetrics200ResponseMetricsInner, changes []GetInventoryMetrics200ResponseChangesInner, effectiveRange NullableGetInventoryMetrics200ResponseEffectiveRange, objectMetrics []GetInventoryMetrics200ResponseObjectMetricsInner, unmatchedArrivals NullableInt32, ) *GetInventoryMetrics200Response`

NewGetInventoryMetrics200Response instantiates a new GetInventoryMetrics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInventoryMetrics200ResponseWithDefaults

`func NewGetInventoryMetrics200ResponseWithDefaults() *GetInventoryMetrics200Response`

NewGetInventoryMetrics200ResponseWithDefaults instantiates a new GetInventoryMetrics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetInventoryMetrics200Response) GetAccount() GetInventoryMetrics200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetInventoryMetrics200Response) GetAccountOk() (*GetInventoryMetrics200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetInventoryMetrics200Response) SetAccount(v GetInventoryMetrics200ResponseAccount)`

SetAccount sets Account field to given value.


### GetSync

`func (o *GetInventoryMetrics200Response) GetSync() GetInventoryMetrics200ResponseSync`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *GetInventoryMetrics200Response) GetSyncOk() (*GetInventoryMetrics200ResponseSync, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *GetInventoryMetrics200Response) SetSync(v GetInventoryMetrics200ResponseSync)`

SetSync sets Sync field to given value.


### GetStaleness

`func (o *GetInventoryMetrics200Response) GetStaleness() GetInventoryMetrics200ResponseStaleness`

GetStaleness returns the Staleness field if non-nil, zero value otherwise.

### GetStalenessOk

`func (o *GetInventoryMetrics200Response) GetStalenessOk() (*GetInventoryMetrics200ResponseStaleness, bool)`

GetStalenessOk returns a tuple with the Staleness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleness

`func (o *GetInventoryMetrics200Response) SetStaleness(v GetInventoryMetrics200ResponseStaleness)`

SetStaleness sets Staleness field to given value.


### GetMetrics

`func (o *GetInventoryMetrics200Response) GetMetrics() []GetInventoryMetrics200ResponseMetricsInner`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetInventoryMetrics200Response) GetMetricsOk() (*[]GetInventoryMetrics200ResponseMetricsInner, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetInventoryMetrics200Response) SetMetrics(v []GetInventoryMetrics200ResponseMetricsInner)`

SetMetrics sets Metrics field to given value.


### GetChanges

`func (o *GetInventoryMetrics200Response) GetChanges() []GetInventoryMetrics200ResponseChangesInner`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *GetInventoryMetrics200Response) GetChangesOk() (*[]GetInventoryMetrics200ResponseChangesInner, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *GetInventoryMetrics200Response) SetChanges(v []GetInventoryMetrics200ResponseChangesInner)`

SetChanges sets Changes field to given value.


### GetEffectiveRange

`func (o *GetInventoryMetrics200Response) GetEffectiveRange() GetInventoryMetrics200ResponseEffectiveRange`

GetEffectiveRange returns the EffectiveRange field if non-nil, zero value otherwise.

### GetEffectiveRangeOk

`func (o *GetInventoryMetrics200Response) GetEffectiveRangeOk() (*GetInventoryMetrics200ResponseEffectiveRange, bool)`

GetEffectiveRangeOk returns a tuple with the EffectiveRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveRange

`func (o *GetInventoryMetrics200Response) SetEffectiveRange(v GetInventoryMetrics200ResponseEffectiveRange)`

SetEffectiveRange sets EffectiveRange field to given value.


### SetEffectiveRangeNil

`func (o *GetInventoryMetrics200Response) SetEffectiveRangeNil(b bool)`

 SetEffectiveRangeNil sets the value for EffectiveRange to be an explicit nil

### UnsetEffectiveRange
`func (o *GetInventoryMetrics200Response) UnsetEffectiveRange()`

UnsetEffectiveRange ensures that no value is present for EffectiveRange, not even an explicit nil
### GetObjectMetrics

`func (o *GetInventoryMetrics200Response) GetObjectMetrics() []GetInventoryMetrics200ResponseObjectMetricsInner`

GetObjectMetrics returns the ObjectMetrics field if non-nil, zero value otherwise.

### GetObjectMetricsOk

`func (o *GetInventoryMetrics200Response) GetObjectMetricsOk() (*[]GetInventoryMetrics200ResponseObjectMetricsInner, bool)`

GetObjectMetricsOk returns a tuple with the ObjectMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectMetrics

`func (o *GetInventoryMetrics200Response) SetObjectMetrics(v []GetInventoryMetrics200ResponseObjectMetricsInner)`

SetObjectMetrics sets ObjectMetrics field to given value.


### GetUnmatchedArrivals

`func (o *GetInventoryMetrics200Response) GetUnmatchedArrivals() int32`

GetUnmatchedArrivals returns the UnmatchedArrivals field if non-nil, zero value otherwise.

### GetUnmatchedArrivalsOk

`func (o *GetInventoryMetrics200Response) GetUnmatchedArrivalsOk() (*int32, bool)`

GetUnmatchedArrivalsOk returns a tuple with the UnmatchedArrivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatchedArrivals

`func (o *GetInventoryMetrics200Response) SetUnmatchedArrivals(v int32)`

SetUnmatchedArrivals sets UnmatchedArrivals field to given value.


### SetUnmatchedArrivalsNil

`func (o *GetInventoryMetrics200Response) SetUnmatchedArrivalsNil(b bool)`

 SetUnmatchedArrivalsNil sets the value for UnmatchedArrivals to be an explicit nil

### UnsetUnmatchedArrivals
`func (o *GetInventoryMetrics200Response) UnsetUnmatchedArrivals()`

UnsetUnmatchedArrivals ensures that no value is present for UnmatchedArrivals, not even an explicit nil
### GetPagination

`func (o *GetInventoryMetrics200Response) GetPagination() map[string]GetInventoryMetrics200ResponsePaginationValue`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetInventoryMetrics200Response) GetPaginationOk() (*map[string]GetInventoryMetrics200ResponsePaginationValue, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetInventoryMetrics200Response) SetPagination(v map[string]GetInventoryMetrics200ResponsePaginationValue)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetInventoryMetrics200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetNextCursor

`func (o *GetInventoryMetrics200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *GetInventoryMetrics200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *GetInventoryMetrics200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *GetInventoryMetrics200Response) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *GetInventoryMetrics200Response) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *GetInventoryMetrics200Response) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


