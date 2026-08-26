# ListEventsByCursor200ResponseEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetEvent200ResponseEvent**](GetEvent200ResponseEvent.md) |  | 
**Path** | **NullableString** |  | 
**PerPage** | **int32** |  | 
**NextCursor** | **NullableString** |  | 
**NextPageUrl** | **NullableString** |  | 
**PrevCursor** | **NullableString** |  | 
**PrevPageUrl** | **NullableString** |  | 

## Methods

### NewListEventsByCursor200ResponseEvents

`func NewListEventsByCursor200ResponseEvents(data []GetEvent200ResponseEvent, path NullableString, perPage int32, nextCursor NullableString, nextPageUrl NullableString, prevCursor NullableString, prevPageUrl NullableString, ) *ListEventsByCursor200ResponseEvents`

NewListEventsByCursor200ResponseEvents instantiates a new ListEventsByCursor200ResponseEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEventsByCursor200ResponseEventsWithDefaults

`func NewListEventsByCursor200ResponseEventsWithDefaults() *ListEventsByCursor200ResponseEvents`

NewListEventsByCursor200ResponseEventsWithDefaults instantiates a new ListEventsByCursor200ResponseEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListEventsByCursor200ResponseEvents) GetData() []GetEvent200ResponseEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListEventsByCursor200ResponseEvents) GetDataOk() (*[]GetEvent200ResponseEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListEventsByCursor200ResponseEvents) SetData(v []GetEvent200ResponseEvent)`

SetData sets Data field to given value.


### GetPath

`func (o *ListEventsByCursor200ResponseEvents) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ListEventsByCursor200ResponseEvents) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ListEventsByCursor200ResponseEvents) SetPath(v string)`

SetPath sets Path field to given value.


### SetPathNil

`func (o *ListEventsByCursor200ResponseEvents) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ListEventsByCursor200ResponseEvents) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPerPage

`func (o *ListEventsByCursor200ResponseEvents) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListEventsByCursor200ResponseEvents) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListEventsByCursor200ResponseEvents) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetNextCursor

`func (o *ListEventsByCursor200ResponseEvents) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListEventsByCursor200ResponseEvents) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListEventsByCursor200ResponseEvents) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *ListEventsByCursor200ResponseEvents) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ListEventsByCursor200ResponseEvents) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetNextPageUrl

`func (o *ListEventsByCursor200ResponseEvents) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *ListEventsByCursor200ResponseEvents) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *ListEventsByCursor200ResponseEvents) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.


### SetNextPageUrlNil

`func (o *ListEventsByCursor200ResponseEvents) SetNextPageUrlNil(b bool)`

 SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil

### UnsetNextPageUrl
`func (o *ListEventsByCursor200ResponseEvents) UnsetNextPageUrl()`

UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
### GetPrevCursor

`func (o *ListEventsByCursor200ResponseEvents) GetPrevCursor() string`

GetPrevCursor returns the PrevCursor field if non-nil, zero value otherwise.

### GetPrevCursorOk

`func (o *ListEventsByCursor200ResponseEvents) GetPrevCursorOk() (*string, bool)`

GetPrevCursorOk returns a tuple with the PrevCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevCursor

`func (o *ListEventsByCursor200ResponseEvents) SetPrevCursor(v string)`

SetPrevCursor sets PrevCursor field to given value.


### SetPrevCursorNil

`func (o *ListEventsByCursor200ResponseEvents) SetPrevCursorNil(b bool)`

 SetPrevCursorNil sets the value for PrevCursor to be an explicit nil

### UnsetPrevCursor
`func (o *ListEventsByCursor200ResponseEvents) UnsetPrevCursor()`

UnsetPrevCursor ensures that no value is present for PrevCursor, not even an explicit nil
### GetPrevPageUrl

`func (o *ListEventsByCursor200ResponseEvents) GetPrevPageUrl() string`

GetPrevPageUrl returns the PrevPageUrl field if non-nil, zero value otherwise.

### GetPrevPageUrlOk

`func (o *ListEventsByCursor200ResponseEvents) GetPrevPageUrlOk() (*string, bool)`

GetPrevPageUrlOk returns a tuple with the PrevPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPageUrl

`func (o *ListEventsByCursor200ResponseEvents) SetPrevPageUrl(v string)`

SetPrevPageUrl sets PrevPageUrl field to given value.


### SetPrevPageUrlNil

`func (o *ListEventsByCursor200ResponseEvents) SetPrevPageUrlNil(b bool)`

 SetPrevPageUrlNil sets the value for PrevPageUrl to be an explicit nil

### UnsetPrevPageUrl
`func (o *ListEventsByCursor200ResponseEvents) UnsetPrevPageUrl()`

UnsetPrevPageUrl ensures that no value is present for PrevPageUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


