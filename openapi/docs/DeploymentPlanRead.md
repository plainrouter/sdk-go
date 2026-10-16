# DeploymentPlanRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**PlatformAdAccountId** | **int32** |  | 
**ReviewVersion** | **string** | Current review version. Pass this value when executing the plan you reviewed. | 
**Copy** | Pointer to [**NullableDeploymentPlanReadCopy**](DeploymentPlanReadCopy.md) |  | [optional] 
**LandingPage** | Pointer to **string** |  | [optional] 
**Creatives** | Pointer to [**[]DeploymentPlanReadCreativesInner**](DeploymentPlanReadCreativesInner.md) |  | [optional] 
**NextCursor** | Pointer to **NullableInt32** |  | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] 
**Continuation** | Pointer to **string** |  | [optional] 
**CreativeCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewDeploymentPlanRead

`func NewDeploymentPlanRead(id string, platformAdAccountId int32, reviewVersion string, ) *DeploymentPlanRead`

NewDeploymentPlanRead instantiates a new DeploymentPlanRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentPlanReadWithDefaults

`func NewDeploymentPlanReadWithDefaults() *DeploymentPlanRead`

NewDeploymentPlanReadWithDefaults instantiates a new DeploymentPlanRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentPlanRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentPlanRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentPlanRead) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentPlanRead) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeploymentPlanRead) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeploymentPlanRead) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *DeploymentPlanRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentPlanRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentPlanRead) SetId(v string)`

SetId sets Id field to given value.


### GetPlatformAdAccountId

`func (o *DeploymentPlanRead) GetPlatformAdAccountId() int32`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *DeploymentPlanRead) GetPlatformAdAccountIdOk() (*int32, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *DeploymentPlanRead) SetPlatformAdAccountId(v int32)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.


### GetReviewVersion

`func (o *DeploymentPlanRead) GetReviewVersion() string`

GetReviewVersion returns the ReviewVersion field if non-nil, zero value otherwise.

### GetReviewVersionOk

`func (o *DeploymentPlanRead) GetReviewVersionOk() (*string, bool)`

GetReviewVersionOk returns a tuple with the ReviewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewVersion

`func (o *DeploymentPlanRead) SetReviewVersion(v string)`

SetReviewVersion sets ReviewVersion field to given value.


### GetCopy

`func (o *DeploymentPlanRead) GetCopy() DeploymentPlanReadCopy`

GetCopy returns the Copy field if non-nil, zero value otherwise.

### GetCopyOk

`func (o *DeploymentPlanRead) GetCopyOk() (*DeploymentPlanReadCopy, bool)`

GetCopyOk returns a tuple with the Copy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopy

`func (o *DeploymentPlanRead) SetCopy(v DeploymentPlanReadCopy)`

SetCopy sets Copy field to given value.

### HasCopy

`func (o *DeploymentPlanRead) HasCopy() bool`

HasCopy returns a boolean if a field has been set.

### SetCopyNil

`func (o *DeploymentPlanRead) SetCopyNil(b bool)`

 SetCopyNil sets the value for Copy to be an explicit nil

### UnsetCopy
`func (o *DeploymentPlanRead) UnsetCopy()`

UnsetCopy ensures that no value is present for Copy, not even an explicit nil
### GetLandingPage

`func (o *DeploymentPlanRead) GetLandingPage() string`

GetLandingPage returns the LandingPage field if non-nil, zero value otherwise.

### GetLandingPageOk

`func (o *DeploymentPlanRead) GetLandingPageOk() (*string, bool)`

GetLandingPageOk returns a tuple with the LandingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPage

`func (o *DeploymentPlanRead) SetLandingPage(v string)`

SetLandingPage sets LandingPage field to given value.

### HasLandingPage

`func (o *DeploymentPlanRead) HasLandingPage() bool`

HasLandingPage returns a boolean if a field has been set.

### GetCreatives

`func (o *DeploymentPlanRead) GetCreatives() []DeploymentPlanReadCreativesInner`

GetCreatives returns the Creatives field if non-nil, zero value otherwise.

### GetCreativesOk

`func (o *DeploymentPlanRead) GetCreativesOk() (*[]DeploymentPlanReadCreativesInner, bool)`

GetCreativesOk returns a tuple with the Creatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatives

`func (o *DeploymentPlanRead) SetCreatives(v []DeploymentPlanReadCreativesInner)`

SetCreatives sets Creatives field to given value.

### HasCreatives

`func (o *DeploymentPlanRead) HasCreatives() bool`

HasCreatives returns a boolean if a field has been set.

### GetNextCursor

`func (o *DeploymentPlanRead) GetNextCursor() int32`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *DeploymentPlanRead) GetNextCursorOk() (*int32, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *DeploymentPlanRead) SetNextCursor(v int32)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *DeploymentPlanRead) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *DeploymentPlanRead) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *DeploymentPlanRead) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetTruncated

`func (o *DeploymentPlanRead) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *DeploymentPlanRead) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *DeploymentPlanRead) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *DeploymentPlanRead) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetContinuation

`func (o *DeploymentPlanRead) GetContinuation() string`

GetContinuation returns the Continuation field if non-nil, zero value otherwise.

### GetContinuationOk

`func (o *DeploymentPlanRead) GetContinuationOk() (*string, bool)`

GetContinuationOk returns a tuple with the Continuation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuation

`func (o *DeploymentPlanRead) SetContinuation(v string)`

SetContinuation sets Continuation field to given value.

### HasContinuation

`func (o *DeploymentPlanRead) HasContinuation() bool`

HasContinuation returns a boolean if a field has been set.

### GetCreativeCount

`func (o *DeploymentPlanRead) GetCreativeCount() int32`

GetCreativeCount returns the CreativeCount field if non-nil, zero value otherwise.

### GetCreativeCountOk

`func (o *DeploymentPlanRead) GetCreativeCountOk() (*int32, bool)`

GetCreativeCountOk returns a tuple with the CreativeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeCount

`func (o *DeploymentPlanRead) SetCreativeCount(v int32)`

SetCreativeCount sets CreativeCount field to given value.

### HasCreativeCount

`func (o *DeploymentPlanRead) HasCreativeCount() bool`

HasCreativeCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


