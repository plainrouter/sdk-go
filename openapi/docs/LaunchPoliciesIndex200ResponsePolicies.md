# LaunchPoliciesIndex200ResponsePolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | **int32** |  | 
**Data** | [**[]LaunchPolicy**](LaunchPolicy.md) |  | 
**FirstPageUrl** | **NullableString** |  | 
**From** | **NullableInt32** |  | 
**LastPageUrl** | **NullableString** |  | 
**LastPage** | **int32** |  | 
**Links** | [**[]LaunchPoliciesIndex200ResponsePoliciesLinksInner**](LaunchPoliciesIndex200ResponsePoliciesLinksInner.md) | Generated paginator links. | 
**NextPageUrl** | **NullableString** |  | 
**Path** | **NullableString** | Base path for paginator generated URLs. | 
**PerPage** | **int32** | Number of items shown per page. | 
**PrevPageUrl** | **NullableString** |  | 
**To** | **NullableInt32** | Number of the last item in the slice. | 
**Total** | **int32** | Total number of items being paginated. | 

## Methods

### NewLaunchPoliciesIndex200ResponsePolicies

`func NewLaunchPoliciesIndex200ResponsePolicies(currentPage int32, data []LaunchPolicy, firstPageUrl NullableString, from NullableInt32, lastPageUrl NullableString, lastPage int32, links []LaunchPoliciesIndex200ResponsePoliciesLinksInner, nextPageUrl NullableString, path NullableString, perPage int32, prevPageUrl NullableString, to NullableInt32, total int32, ) *LaunchPoliciesIndex200ResponsePolicies`

NewLaunchPoliciesIndex200ResponsePolicies instantiates a new LaunchPoliciesIndex200ResponsePolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchPoliciesIndex200ResponsePoliciesWithDefaults

`func NewLaunchPoliciesIndex200ResponsePoliciesWithDefaults() *LaunchPoliciesIndex200ResponsePolicies`

NewLaunchPoliciesIndex200ResponsePoliciesWithDefaults instantiates a new LaunchPoliciesIndex200ResponsePolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetData

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetData() []LaunchPolicy`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetDataOk() (*[]LaunchPolicy, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetData(v []LaunchPolicy)`

SetData sets Data field to given value.


### GetFirstPageUrl

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetFirstPageUrl() string`

GetFirstPageUrl returns the FirstPageUrl field if non-nil, zero value otherwise.

### GetFirstPageUrlOk

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetFirstPageUrlOk() (*string, bool)`

GetFirstPageUrlOk returns a tuple with the FirstPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPageUrl

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetFirstPageUrl(v string)`

SetFirstPageUrl sets FirstPageUrl field to given value.


### SetFirstPageUrlNil

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetFirstPageUrlNil(b bool)`

 SetFirstPageUrlNil sets the value for FirstPageUrl to be an explicit nil

### UnsetFirstPageUrl
`func (o *LaunchPoliciesIndex200ResponsePolicies) UnsetFirstPageUrl()`

UnsetFirstPageUrl ensures that no value is present for FirstPageUrl, not even an explicit nil
### GetFrom

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetFrom(v int32)`

SetFrom sets From field to given value.


### SetFromNil

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *LaunchPoliciesIndex200ResponsePolicies) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetLastPageUrl

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetLastPageUrl() string`

GetLastPageUrl returns the LastPageUrl field if non-nil, zero value otherwise.

### GetLastPageUrlOk

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetLastPageUrlOk() (*string, bool)`

GetLastPageUrlOk returns a tuple with the LastPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPageUrl

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetLastPageUrl(v string)`

SetLastPageUrl sets LastPageUrl field to given value.


### SetLastPageUrlNil

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetLastPageUrlNil(b bool)`

 SetLastPageUrlNil sets the value for LastPageUrl to be an explicit nil

### UnsetLastPageUrl
`func (o *LaunchPoliciesIndex200ResponsePolicies) UnsetLastPageUrl()`

UnsetLastPageUrl ensures that no value is present for LastPageUrl, not even an explicit nil
### GetLastPage

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetLastPage() int32`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetLastPageOk() (*int32, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetLastPage(v int32)`

SetLastPage sets LastPage field to given value.


### GetLinks

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetLinks() []LaunchPoliciesIndex200ResponsePoliciesLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetLinksOk() (*[]LaunchPoliciesIndex200ResponsePoliciesLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetLinks(v []LaunchPoliciesIndex200ResponsePoliciesLinksInner)`

SetLinks sets Links field to given value.


### GetNextPageUrl

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetNextPageUrl() string`

GetNextPageUrl returns the NextPageUrl field if non-nil, zero value otherwise.

### GetNextPageUrlOk

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetNextPageUrlOk() (*string, bool)`

GetNextPageUrlOk returns a tuple with the NextPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageUrl

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetNextPageUrl(v string)`

SetNextPageUrl sets NextPageUrl field to given value.


### SetNextPageUrlNil

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetNextPageUrlNil(b bool)`

 SetNextPageUrlNil sets the value for NextPageUrl to be an explicit nil

### UnsetNextPageUrl
`func (o *LaunchPoliciesIndex200ResponsePolicies) UnsetNextPageUrl()`

UnsetNextPageUrl ensures that no value is present for NextPageUrl, not even an explicit nil
### GetPath

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetPath(v string)`

SetPath sets Path field to given value.


### SetPathNil

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *LaunchPoliciesIndex200ResponsePolicies) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPerPage

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetPrevPageUrl

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetPrevPageUrl() string`

GetPrevPageUrl returns the PrevPageUrl field if non-nil, zero value otherwise.

### GetPrevPageUrlOk

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetPrevPageUrlOk() (*string, bool)`

GetPrevPageUrlOk returns a tuple with the PrevPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevPageUrl

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetPrevPageUrl(v string)`

SetPrevPageUrl sets PrevPageUrl field to given value.


### SetPrevPageUrlNil

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetPrevPageUrlNil(b bool)`

 SetPrevPageUrlNil sets the value for PrevPageUrl to be an explicit nil

### UnsetPrevPageUrl
`func (o *LaunchPoliciesIndex200ResponsePolicies) UnsetPrevPageUrl()`

UnsetPrevPageUrl ensures that no value is present for PrevPageUrl, not even an explicit nil
### GetTo

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetTo(v int32)`

SetTo sets To field to given value.


### SetToNil

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *LaunchPoliciesIndex200ResponsePolicies) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetTotal

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LaunchPoliciesIndex200ResponsePolicies) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LaunchPoliciesIndex200ResponsePolicies) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


