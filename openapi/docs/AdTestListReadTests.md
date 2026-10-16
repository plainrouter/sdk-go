# AdTestListReadTests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | **int32** |  | 
**Data** | [**[]AdTestCreateReadTest**](AdTestCreateReadTest.md) |  | 
**FirstPageUrl** | **NullableString** |  | 
**From** | **NullableInt32** |  | 
**LastPage** | **int32** |  | 
**LastPageUrl** | **NullableString** |  | 
**Links** | [**[]LaunchPoliciesIndex200ResponsePoliciesLinksInner**](LaunchPoliciesIndex200ResponsePoliciesLinksInner.md) |  | 
**NextPageUrl** | **NullableString** |  | 
**Path** | **string** |  | 
**PerPage** | **int32** |  | 
**PrevPageUrl** | **NullableString** |  | 
**To** | **NullableInt32** |  | 
**Total** | **int32** |  | 

## Methods

### NewAdTestListReadTests

`func NewAdTestListReadTests(currentPage int32, data []AdTestCreateReadTest, firstPageUrl NullableString, from NullableInt32, lastPage int32, lastPageUrl NullableString, links []LaunchPoliciesIndex200ResponsePoliciesLinksInner, nextPageUrl NullableString, path string, perPage int32, prevPageUrl NullableString, to NullableInt32, total int32, ) *AdTestListReadTests`

NewAdTestListReadTests instantiates a new AdTestListReadTests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdTestListReadTestsWithDefaults

`func NewAdTestListReadTestsWithDefaults() *AdTestListReadTests`

NewAdTestListReadTestsWithDefaults instantiates a new AdTestListReadTests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *AdTestListReadTests) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *AdTestListReadTests) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *AdTestListReadTests) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetData

`func (o *AdTestListReadTests) GetData() []AdTestCreateReadTest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AdTestListReadTests) GetDataOk() (*[]AdTestCreateReadTest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AdTestListReadTests) SetData(v []AdTestCreateReadTest)`

SetData sets Data field to given value.


### GetFirstPageUrl

`func (o *AdTestListReadTests) GetFirstPageUrl() string`

GetFirstPageUrl returns the FirstPageUrl field if non-nil, zero value otherwise.

### GetFirstPageUrlOk

`func (o *AdTestListReadTests) GetFirstPageUrlOk() (*string, bool)`

GetFirstPageUrlOk returns a tuple with the FirstPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUrl

`func (o *AdTestListReadTests) SetFirstPageUrl(v string)`

SetFirstPageUrl sets FirstPageUrl field to given value.


### SetFirstPageUrlNil

`func (o *AdTestListReadTests) SetFirstPageUrlNil(b bool)`

 SetFirstPageUrlNil sets the value for FirstPageUrl to be an explicit nil

### UnsetFirstPageUrl
`func (o *AdTestListReadTests) UnsetFirstPageUrl()`

UnsetFirstPageUrl ensures that no value is present for FirstPageUrl, not even an explicit nil
### GetFrom

`func (o *AdTestListReadTests) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *AdTestListReadTests) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *AdTestListReadTests) SetFrom(v int32)`

SetFrom sets From field to given value.


### SetFromNil

`func (o *AdTestListReadTests) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *AdTestListReadTests) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetLastPage

`func (o *AdTestListReadTests) GetLastPage() int32`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *AdTestListReadTests) GetLastPageOk() (*int32, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *AdTestListReadTests) SetLastPage(v int32)`

SetLastPage sets LastPage field to given value.


### GetLastPageUrl

`func (o *AdTestListReadTests) GetLastPageUrl() string`

GetLastPageUrl returns the LastPageUrl field if non-nil, zero value otherwise.

### GetLastPageUrlOk

`func (o *AdTestListReadTests) GetLastPageUrlOk() (*string, bool)`

GetLastPageUrlOk returns a tuple with the LastPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPageUrl

`func (o *AdTestListReadTests) SetLastPageUrl(v string)`

SetLastPageUrl sets LastPageUrl field to given value.


### SetLastPageUrlNil

`func (o *AdTestListReadTests) SetLastPageUrlNil(b bool)`

 SetLastPageUrlNil sets the value for LastPageUrl to be an explicit nil

### UnsetLastPageUrl
`func (o *AdTestListReadTests) UnsetLastPageUrl()`

UnsetLastPageUrl ensures that no value is present for LastPageUrl, not even an explicit nil
### GetLinks

`func (o *AdTestListReadTests) GetLinks() []LaunchPoliciesIndex200ResponsePoliciesLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AdTestListReadTests) GetLinksOk() (*[]LaunchPoliciesIndex200ResponsePoliciesLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AdTestListReadTests) SetLinks(v []LaunchPoliciesIndex200ResponsePoliciesLinksInner)`

SetLinks sets Links field to given value.


### GetNextPageUrl

`func (o *AdTestListReadTests) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *AdTestListReadTests) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *AdTestListReadTests) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.


### SetNextPageUrlNil

`func (o *AdTestListReadTests) SetNextPageUrlNil(b bool)`

 SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil

### UnsetNextPageUrl
`func (o *AdTestListReadTests) UnsetNextPageUrl()`

UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
### GetPath

`func (o *AdTestListReadTests) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AdTestListReadTests) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AdTestListReadTests) SetPath(v string)`

SetPath sets Path field to given value.


### GetPerPage

`func (o *AdTestListReadTests) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *AdTestListReadTests) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *AdTestListReadTests) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetPrevPageUrl

`func (o *AdTestListReadTests) GetPrevPageUrl() string`

GetPrevPageUrl returns the PrevPageUrl field if non-nil, zero value otherwise.

### GetPrevPageUrlOk

`func (o *AdTestListReadTests) GetPrevPageUrlOk() (*string, bool)`

GetPrevPageUrlOk returns a tuple with the PrevPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPageUrl

`func (o *AdTestListReadTests) SetPrevPageUrl(v string)`

SetPrevPageUrl sets PrevPageUrl field to given value.


### SetPrevPageUrlNil

`func (o *AdTestListReadTests) SetPrevPageUrlNil(b bool)`

 SetPrevPageUrlNil sets the value for PrevPageUrl to be an explicit nil

### UnsetPrevPageUrl
`func (o *AdTestListReadTests) UnsetPrevPageUrl()`

UnsetPrevPageUrl ensures that no value is present for PrevPageUrl, not even an explicit nil
### GetTo

`func (o *AdTestListReadTests) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *AdTestListReadTests) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *AdTestListReadTests) SetTo(v int32)`

SetTo sets To field to given value.


### SetToNil

`func (o *AdTestListReadTests) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *AdTestListReadTests) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetTotal

`func (o *AdTestListReadTests) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AdTestListReadTests) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AdTestListReadTests) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


