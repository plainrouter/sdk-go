# StoreCreativeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | ***os.File** |  | 
**SourceRef** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewStoreCreativeRequest

`func NewStoreCreativeRequest(file *os.File, ) *StoreCreativeRequest`

NewStoreCreativeRequest instantiates a new StoreCreativeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreCreativeRequestWithDefaults

`func NewStoreCreativeRequestWithDefaults() *StoreCreativeRequest`

NewStoreCreativeRequestWithDefaults instantiates a new StoreCreativeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *StoreCreativeRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *StoreCreativeRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *StoreCreativeRequest) SetFile(v *os.File)`

SetFile sets File field to given value.


### GetSourceRef

`func (o *StoreCreativeRequest) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *StoreCreativeRequest) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *StoreCreativeRequest) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *StoreCreativeRequest) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### SetSourceRefNil

`func (o *StoreCreativeRequest) SetSourceRefNil(b bool)`

 SetSourceRefNil sets the value for SourceRef to be an explicit nil

### UnsetSourceRef
`func (o *StoreCreativeRequest) UnsetSourceRef()`

UnsetSourceRef ensures that no value is present for SourceRef, not even an explicit nil
### GetTags

`func (o *StoreCreativeRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StoreCreativeRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StoreCreativeRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StoreCreativeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


