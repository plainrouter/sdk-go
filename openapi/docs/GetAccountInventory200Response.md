# GetAccountInventory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **map[string]interface{}** |  | 
**Sync** | **map[string]interface{}** |  | 
**Staleness** | **map[string]interface{}** |  | 
**Summary** | [**GetAccountInventory200ResponseSummary**](GetAccountInventory200ResponseSummary.md) |  | 
**Objects** | **[]map[string]*interface{}** |  | 
**NextCursor** | **NullableInt32** |  | 
**Truncated** | **bool** |  | 
**Continuation** | **string** |  | 

## Methods

### NewGetAccountInventory200Response

`func NewGetAccountInventory200Response(account map[string]*interface{}, sync map[string]*interface{}, staleness map[string]*interface{}, summary GetAccountInventory200ResponseSummary, objects []map[string]*interface{}, nextCursor NullableInt32, truncated bool, continuation string, ) *GetAccountInventory200Response`

NewGetAccountInventory200Response instantiates a new GetAccountInventory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountInventory200ResponseWithDefaults

`func NewGetAccountInventory200ResponseWithDefaults() *GetAccountInventory200Response`

NewGetAccountInventory200ResponseWithDefaults instantiates a new GetAccountInventory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetAccountInventory200Response) GetAccount() map[string]*interface{}`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetAccountInventory200Response) GetAccountOk() (*map[string]*interface{}, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetAccountInventory200Response) SetAccount(v map[string]*interface{})`

SetAccount sets Account field to given value.


### GetSync

`func (o *GetAccountInventory200Response) GetSync() map[string]*interface{}`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *GetAccountInventory200Response) GetSyncOk() (*map[string]*interface{}, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *GetAccountInventory200Response) SetSync(v map[string]*interface{})`

SetSync sets Sync field to given value.


### GetStaleness

`func (o *GetAccountInventory200Response) GetStaleness() map[string]*interface{}`

GetStaleness returns the Staleness field if non-nil, zero value otherwise.

### GetStalenessOk

`func (o *GetAccountInventory200Response) GetStalenessOk() (*map[string]*interface{}, bool)`

GetStalenessOk returns a tuple with the Staleness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleness

`func (o *GetAccountInventory200Response) SetStaleness(v map[string]*interface{})`

SetStaleness sets Staleness field to given value.


### GetSummary

`func (o *GetAccountInventory200Response) GetSummary() GetAccountInventory200ResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetAccountInventory200Response) GetSummaryOk() (*GetAccountInventory200ResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetAccountInventory200Response) SetSummary(v GetAccountInventory200ResponseSummary)`

SetSummary sets Summary field to given value.


### GetObjects

`func (o *GetAccountInventory200Response) GetObjects() []map[string]*interface{}`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetAccountInventory200Response) GetObjectsOk() (*[]map[string]*interface{}, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetAccountInventory200Response) SetObjects(v []map[string]*interface{})`

SetObjects sets Objects field to given value.


### GetNextCursor

`func (o *GetAccountInventory200Response) GetNextCursor() int32`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *GetAccountInventory200Response) GetNextCursorOk() (*int32, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *GetAccountInventory200Response) SetNextCursor(v int32)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *GetAccountInventory200Response) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *GetAccountInventory200Response) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetTruncated

`func (o *GetAccountInventory200Response) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GetAccountInventory200Response) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GetAccountInventory200Response) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.


### GetContinuation

`func (o *GetAccountInventory200Response) GetContinuation() string`

GetContinuation returns the Continuation field if non-nil, zero value otherwise.

### GetContinuationOk

`func (o *GetAccountInventory200Response) GetContinuationOk() (*string, bool)`

GetContinuationOk returns a tuple with the Continuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuation

`func (o *GetAccountInventory200Response) SetContinuation(v string)`

SetContinuation sets Continuation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


