# CreativeMcpPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**State** | **string** |  | 
**Url** | Pointer to **string** | Ready items only, delivered in tool-result _meta.creative_previews to the MCP App. Signed R2 derivative URL valid for at most 300 seconds; never an original or provider URL. | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**Width** | Pointer to **int32** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**ByteSize** | Pointer to **int32** | Stored derivative at most 1048576 bytes and 4000000 pixels. SVG refused; the existing decode cap applies. | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 

## Methods

### NewCreativeMcpPreview

`func NewCreativeMcpPreview(id string, state string, ) *CreativeMcpPreview`

NewCreativeMcpPreview instantiates a new CreativeMcpPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreativeMcpPreviewWithDefaults

`func NewCreativeMcpPreviewWithDefaults() *CreativeMcpPreview`

NewCreativeMcpPreviewWithDefaults instantiates a new CreativeMcpPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreativeMcpPreview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreativeMcpPreview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreativeMcpPreview) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *CreativeMcpPreview) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreativeMcpPreview) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreativeMcpPreview) SetState(v string)`

SetState sets State field to given value.


### GetUrl

`func (o *CreativeMcpPreview) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreativeMcpPreview) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreativeMcpPreview) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreativeMcpPreview) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMimeType

`func (o *CreativeMcpPreview) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *CreativeMcpPreview) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *CreativeMcpPreview) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *CreativeMcpPreview) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetWidth

`func (o *CreativeMcpPreview) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *CreativeMcpPreview) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *CreativeMcpPreview) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *CreativeMcpPreview) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *CreativeMcpPreview) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *CreativeMcpPreview) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *CreativeMcpPreview) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *CreativeMcpPreview) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetByteSize

`func (o *CreativeMcpPreview) GetByteSize() int32`

GetByteSize returns the ByteSize field if non-nil, zero value otherwise.

### GetByteSizeOk

`func (o *CreativeMcpPreview) GetByteSizeOk() (*int32, bool)`

GetByteSizeOk returns a tuple with the ByteSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteSize

`func (o *CreativeMcpPreview) SetByteSize(v int32)`

SetByteSize sets ByteSize field to given value.

### HasByteSize

`func (o *CreativeMcpPreview) HasByteSize() bool`

HasByteSize returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreativeMcpPreview) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreativeMcpPreview) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreativeMcpPreview) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreativeMcpPreview) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


