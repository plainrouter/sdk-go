# LaunchCreativesIndex200ResponseCreativesDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Fingerprint** | **string** |  | 
**Type** | **string** |  | 
**MimeType** | **string** |  | 
**ByteSize** | **int32** |  | 
**Width** | **NullableInt32** |  | 
**Height** | **NullableInt32** |  | 
**AspectRatio** | **NullableFloat32** |  | 
**DurationMs** | **NullableInt32** |  | 
**OriginalFilename** | Pointer to **NullableString** |  | [optional] 
**SourceRef** | **NullableString** |  | 
**Tags** | [**LaunchCreativesIndex200ResponseCreativesDataInnerTags**](LaunchCreativesIndex200ResponseCreativesDataInnerTags.md) |  | 
**Status** | **string** |  | 
**CreatedAt** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 

## Methods

### NewLaunchCreativesIndex200ResponseCreativesDataInner

`func NewLaunchCreativesIndex200ResponseCreativesDataInner(id string, fingerprint string, type_ string, mimeType string, byteSize int32, width NullableInt32, height NullableInt32, aspectRatio NullableFloat32, durationMs NullableInt32, sourceRef NullableString, tags LaunchCreativesIndex200ResponseCreativesDataInnerTags, status string, createdAt NullableString, updatedAt NullableString, ) *LaunchCreativesIndex200ResponseCreativesDataInner`

NewLaunchCreativesIndex200ResponseCreativesDataInner instantiates a new LaunchCreativesIndex200ResponseCreativesDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchCreativesIndex200ResponseCreativesDataInnerWithDefaults

`func NewLaunchCreativesIndex200ResponseCreativesDataInnerWithDefaults() *LaunchCreativesIndex200ResponseCreativesDataInner`

NewLaunchCreativesIndex200ResponseCreativesDataInnerWithDefaults instantiates a new LaunchCreativesIndex200ResponseCreativesDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetFingerprint

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetType

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetType(v string)`

SetType sets Type field to given value.


### GetMimeType

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetByteSize

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetByteSize() int32`

GetByteSize returns the ByteSize field if non-nil, zero value otherwise.

### GetByteSizeOk

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetByteSizeOk() (*int32, bool)`

GetByteSizeOk returns a tuple with the ByteSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteSize

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetByteSize(v int32)`

SetByteSize sets ByteSize field to given value.


### GetWidth

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetWidth(v int32)`

SetWidth sets Width field to given value.


### SetWidthNil

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetHeight(v int32)`

SetHeight sets Height field to given value.


### SetHeightNil

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetAspectRatio

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetAspectRatio() float32`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetAspectRatioOk() (*float32, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetAspectRatio(v float32)`

SetAspectRatio sets AspectRatio field to given value.


### SetAspectRatioNil

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetAspectRatioNil(b bool)`

 SetAspectRatioNil sets the value for AspectRatio to be an explicit nil

### UnsetAspectRatio
`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) UnsetAspectRatio()`

UnsetAspectRatio ensures that no value is present for AspectRatio, not even an explicit nil
### GetDurationMs

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.


### SetDurationMsNil

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetDurationMsNil(b bool)`

 SetDurationMsNil sets the value for DurationMs to be an explicit nil

### UnsetDurationMs
`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) UnsetDurationMs()`

UnsetDurationMs ensures that no value is present for DurationMs, not even an explicit nil
### GetOriginalFilename

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetOriginalFilename() string`

GetOriginalFilename returns the OriginalFilename field if non-nil, zero value otherwise.

### GetOriginalFilenameOk

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetOriginalFilenameOk() (*string, bool)`

GetOriginalFilenameOk returns a tuple with the OriginalFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFilename

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetOriginalFilename(v string)`

SetOriginalFilename sets OriginalFilename field to given value.

### HasOriginalFilename

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) HasOriginalFilename() bool`

HasOriginalFilename returns a boolean if a field has been set.

### SetOriginalFilenameNil

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetOriginalFilenameNil(b bool)`

 SetOriginalFilenameNil sets the value for OriginalFilename to be an explicit nil

### UnsetOriginalFilename
`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) UnsetOriginalFilename()`

UnsetOriginalFilename ensures that no value is present for OriginalFilename, not even an explicit nil
### GetSourceRef

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetSourceRef() string`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetSourceRefOk() (*string, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetSourceRef(v string)`

SetSourceRef sets SourceRef field to given value.


### SetSourceRefNil

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetSourceRefNil(b bool)`

 SetSourceRefNil sets the value for SourceRef to be an explicit nil

### UnsetSourceRef
`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) UnsetSourceRef()`

UnsetSourceRef ensures that no value is present for SourceRef, not even an explicit nil
### GetTags

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetTags() LaunchCreativesIndex200ResponseCreativesDataInnerTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetTagsOk() (*LaunchCreativesIndex200ResponseCreativesDataInnerTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetTags(v LaunchCreativesIndex200ResponseCreativesDataInnerTags)`

SetTags sets Tags field to given value.


### GetStatus

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LaunchCreativesIndex200ResponseCreativesDataInner) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


