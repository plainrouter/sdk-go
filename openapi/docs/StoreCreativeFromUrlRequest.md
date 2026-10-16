# StoreCreativeFromUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Vetted HTTPS URL to fetch; private addresses and unsafe redirects are refused. | 
**Tags** | Pointer to [**StoreCreativeFromUrlRequestTags**](StoreCreativeFromUrlRequestTags.md) |  | [optional] 
**SourceRef** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStoreCreativeFromUrlRequest

`func NewStoreCreativeFromUrlRequest(url string, ) *StoreCreativeFromUrlRequest`

NewStoreCreativeFromUrlRequest instantiates a new StoreCreativeFromUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreCreativeFromUrlRequestWithDefaults

`func NewStoreCreativeFromUrlRequestWithDefaults() *StoreCreativeFromUrlRequest`

NewStoreCreativeFromUrlRequestWithDefaults instantiates a new StoreCreativeFromUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *StoreCreativeFromUrlRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StoreCreativeFromUrlRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StoreCreativeFromUrlRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetTags

`func (o *StoreCreativeFromUrlRequest) GetTags() StoreCreativeFromUrlRequestTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StoreCreativeFromUrlRequest) GetTagsOk() (*StoreCreativeFromUrlRequestTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StoreCreativeFromUrlRequest) SetTags(v StoreCreativeFromUrlRequestTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StoreCreativeFromUrlRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSourceRef

`func (o *StoreCreativeFromUrlRequest) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *StoreCreativeFromUrlRequest) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *StoreCreativeFromUrlRequest) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *StoreCreativeFromUrlRequest) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### SetSourceRefNil

`func (o *StoreCreativeFromUrlRequest) SetSourceRefNil(b bool)`

 SetSourceRefNil sets the value for SourceRef to be an explicit nil

### UnsetSourceRef
`func (o *StoreCreativeFromUrlRequest) UnsetSourceRef()`

UnsetSourceRef ensures that no value is present for SourceRef, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


