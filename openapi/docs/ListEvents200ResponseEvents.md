# ListEvents200ResponseEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | **int32** |  | 
**Data** | [**[]GetEvent200ResponseEvent**](GetEvent200ResponseEvent.md) |  | 
**FirstPageUrl** | **NullableString** |  | 
**From** | **NullableInt32** |  | 
**LastPage** | **int32** |  | 
**LastPageUrl** | **NullableString** |  | 
**Links** | [**[]ListEvents200ResponseEventsLinksInner**](ListEvents200ResponseEventsLinksInner.md) |  | 
**NextPageUrl** | **NullableString** |  | 
**Path** | **NullableString** |  | 
**PerPage** | **int32** |  | 
**PrevPageUrl** | **NullableString** |  | 
**To** | **NullableInt32** |  | 
**Total** | **int32** |  | 

## Methods

### NewListEvents200ResponseEvents

`func NewListEvents200ResponseEvents(currentPage int32, data []GetEvent200ResponseEvent, firstPageUrl NullableString, from NullableInt32, lastPage int32, lastPageUrl NullableString, links []ListEvents200ResponseEventsLinksInner, nextPageUrl NullableString, path NullableString, perPage int32, prevPageUrl NullableString, to NullableInt32, total int32, ) *ListEvents200ResponseEvents`

NewListEvents200ResponseEvents instantiates a new ListEvents200ResponseEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEvents200ResponseEventsWithDefaults

`func NewListEvents200ResponseEventsWithDefaults() *ListEvents200ResponseEvents`

NewListEvents200ResponseEventsWithDefaults instantiates a new ListEvents200ResponseEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *ListEvents200ResponseEvents) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ListEvents200ResponseEvents) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ListEvents200ResponseEvents) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetData

`func (o *ListEvents200ResponseEvents) GetData() []GetEvent200ResponseEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListEvents200ResponseEvents) GetDataOk() (*[]GetEvent200ResponseEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListEvents200ResponseEvents) SetData(v []GetEvent200ResponseEvent)`

SetData sets Data field to given value.


### GetFirstPageUrl

`func (o *ListEvents200ResponseEvents) GetFirstPageUrl() string`

GetFirstPageUrl returns the FirstPageUrl field if non-nil, zero value otherwise.

### GetFirstPageUrlOk

`func (o *ListEvents200ResponseEvents) GetFirstPageUrlOk() (*string, bool)`

GetFirstPageUrlOk returns a tuple with the FirstPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUrl

`func (o *ListEvents200ResponseEvents) SetFirstPageUrl(v string)`

SetFirstPageUrl sets FirstPageUrl field to given value.


### SetFirstPageUrlNil

`func (o *ListEvents200ResponseEvents) SetFirstPageUrlNil(b bool)`

 SetFirstPageUrlNil sets the value for FirstPageUrl to be an explicit nil

### UnsetFirstPageUrl
`func (o *ListEvents200ResponseEvents) UnsetFirstPageUrl()`

UnsetFirstPageUrl ensures that no value is present for FirstPageUrl, not even an explicit nil
### GetFrom

`func (o *ListEvents200ResponseEvents) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ListEvents200ResponseEvents) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ListEvents200ResponseEvents) SetFrom(v int32)`

SetFrom sets From field to given value.


### SetFromNil

`func (o *ListEvents200ResponseEvents) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *ListEvents200ResponseEvents) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetLastPage

`func (o *ListEvents200ResponseEvents) GetLastPage() int32`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *ListEvents200ResponseEvents) GetLastPageOk() (*int32, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *ListEvents200ResponseEvents) SetLastPage(v int32)`

SetLastPage sets LastPage field to given value.


### GetLastPageUrl

`func (o *ListEvents200ResponseEvents) GetLastPageUrl() string`

GetLastPageUrl returns the LastPageUrl field if non-nil, zero value otherwise.

### GetLastPageUrlOk

`func (o *ListEvents200ResponseEvents) GetLastPageUrlOk() (*string, bool)`

GetLastPageUrlOk returns a tuple with the LastPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPageUrl

`func (o *ListEvents200ResponseEvents) SetLastPageUrl(v string)`

SetLastPageUrl sets LastPageUrl field to given value.


### SetLastPageUrlNil

`func (o *ListEvents200ResponseEvents) SetLastPageUrlNil(b bool)`

 SetLastPageUrlNil sets the value for LastPageUrl to be an explicit nil

### UnsetLastPageUrl
`func (o *ListEvents200ResponseEvents) UnsetLastPageUrl()`

UnsetLastPageUrl ensures that no value is present for LastPageUrl, not even an explicit nil
### GetLinks

`func (o *ListEvents200ResponseEvents) GetLinks() []ListEvents200ResponseEventsLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListEvents200ResponseEvents) GetLinksOk() (*[]ListEvents200ResponseEventsLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListEvents200ResponseEvents) SetLinks(v []ListEvents200ResponseEventsLinksInner)`

SetLinks sets Links field to given value.


### GetNextPageUrl

`func (o *ListEvents200ResponseEvents) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *ListEvents200ResponseEvents) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *ListEvents200ResponseEvents) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.


### SetNextPageUrlNil

`func (o *ListEvents200ResponseEvents) SetNextPageUrlNil(b bool)`

 SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil

### UnsetNextPageUrl
`func (o *ListEvents200ResponseEvents) UnsetNextPageUrl()`

UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
### GetPath

`func (o *ListEvents200ResponseEvents) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ListEvents200ResponseEvents) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ListEvents200ResponseEvents) SetPath(v string)`

SetPath sets Path field to given value.


### SetPathNil

`func (o *ListEvents200ResponseEvents) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ListEvents200ResponseEvents) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPerPage

`func (o *ListEvents200ResponseEvents) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ListEvents200ResponseEvents) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ListEvents200ResponseEvents) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetPrevPageUrl

`func (o *ListEvents200ResponseEvents) GetPrevPageUrl() string`

GetPrevPageUrl returns the PrevPageUrl field if non-nil, zero value otherwise.

### GetPrevPageUrlOk

`func (o *ListEvents200ResponseEvents) GetPrevPageUrlOk() (*string, bool)`

GetPrevPageUrlOk returns a tuple with the PrevPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPageUrl

`func (o *ListEvents200ResponseEvents) SetPrevPageUrl(v string)`

SetPrevPageUrl sets PrevPageUrl field to given value.


### SetPrevPageUrlNil

`func (o *ListEvents200ResponseEvents) SetPrevPageUrlNil(b bool)`

 SetPrevPageUrlNil sets the value for PrevPageUrl to be an explicit nil

### UnsetPrevPageUrl
`func (o *ListEvents200ResponseEvents) UnsetPrevPageUrl()`

UnsetPrevPageUrl ensures that no value is present for PrevPageUrl, not even an explicit nil
### GetTo

`func (o *ListEvents200ResponseEvents) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ListEvents200ResponseEvents) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ListEvents200ResponseEvents) SetTo(v int32)`

SetTo sets To field to given value.


### SetToNil

`func (o *ListEvents200ResponseEvents) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ListEvents200ResponseEvents) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetTotal

`func (o *ListEvents200ResponseEvents) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListEvents200ResponseEvents) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListEvents200ResponseEvents) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


