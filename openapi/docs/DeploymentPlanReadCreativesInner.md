# DeploymentPlanReadCreativesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Position** | **int32** |  | 
**Copy** | [**NullableDeploymentPlanReadCreativesInnerCopy**](DeploymentPlanReadCreativesInnerCopy.md) |  | 
**LandingPage** | **NullableString** | Optional per-ad HTTPS landing URL override; null uses the plan landing URL. | 
**ResolvedCopy** | [**NullableDeploymentPlanReadCopy**](DeploymentPlanReadCopy.md) |  | 
**ResolvedLandingPage** | **string** |  | 
**CopyDiffers** | **bool** |  | 
**LandingPageDiffers** | **bool** |  | 

## Methods

### NewDeploymentPlanReadCreativesInner

`func NewDeploymentPlanReadCreativesInner(id string, position int32, copy NullableDeploymentPlanReadCreativesInnerCopy, landingPage NullableString, resolvedCopy NullableDeploymentPlanReadCopy, resolvedLandingPage string, copyDiffers bool, landingPageDiffers bool, ) *DeploymentPlanReadCreativesInner`

NewDeploymentPlanReadCreativesInner instantiates a new DeploymentPlanReadCreativesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentPlanReadCreativesInnerWithDefaults

`func NewDeploymentPlanReadCreativesInnerWithDefaults() *DeploymentPlanReadCreativesInner`

NewDeploymentPlanReadCreativesInnerWithDefaults instantiates a new DeploymentPlanReadCreativesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentPlanReadCreativesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentPlanReadCreativesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentPlanReadCreativesInner) SetId(v string)`

SetId sets Id field to given value.


### GetPosition

`func (o *DeploymentPlanReadCreativesInner) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *DeploymentPlanReadCreativesInner) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *DeploymentPlanReadCreativesInner) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetCopy

`func (o *DeploymentPlanReadCreativesInner) GetCopy() DeploymentPlanReadCreativesInnerCopy`

GetCopy returns the Copy field if non-nil, zero value otherwise.

### GetCopyOk

`func (o *DeploymentPlanReadCreativesInner) GetCopyOk() (*DeploymentPlanReadCreativesInnerCopy, bool)`

GetCopyOk returns a tuple with the Copy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopy

`func (o *DeploymentPlanReadCreativesInner) SetCopy(v DeploymentPlanReadCreativesInnerCopy)`

SetCopy sets Copy field to given value.


### SetCopyNil

`func (o *DeploymentPlanReadCreativesInner) SetCopyNil(b bool)`

 SetCopyNil sets the value for Copy to be an explicit nil

### UnsetCopy
`func (o *DeploymentPlanReadCreativesInner) UnsetCopy()`

UnsetCopy ensures that no value is present for Copy, not even an explicit nil
### GetLandingPage

`func (o *DeploymentPlanReadCreativesInner) GetLandingPage() string`

GetLandingPage returns the LandingPage field if non-nil, zero value otherwise.

### GetLandingPageOk

`func (o *DeploymentPlanReadCreativesInner) GetLandingPageOk() (*string, bool)`

GetLandingPageOk returns a tuple with the LandingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPage

`func (o *DeploymentPlanReadCreativesInner) SetLandingPage(v string)`

SetLandingPage sets LandingPage field to given value.


### SetLandingPageNil

`func (o *DeploymentPlanReadCreativesInner) SetLandingPageNil(b bool)`

 SetLandingPageNil sets the value for LandingPage to be an explicit nil

### UnsetLandingPage
`func (o *DeploymentPlanReadCreativesInner) UnsetLandingPage()`

UnsetLandingPage ensures that no value is present for LandingPage, not even an explicit nil
### GetResolvedCopy

`func (o *DeploymentPlanReadCreativesInner) GetResolvedCopy() DeploymentPlanReadCopy`

GetResolvedCopy returns the ResolvedCopy field if non-nil, zero value otherwise.

### GetResolvedCopyOk

`func (o *DeploymentPlanReadCreativesInner) GetResolvedCopyOk() (*DeploymentPlanReadCopy, bool)`

GetResolvedCopyOk returns a tuple with the ResolvedCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedCopy

`func (o *DeploymentPlanReadCreativesInner) SetResolvedCopy(v DeploymentPlanReadCopy)`

SetResolvedCopy sets ResolvedCopy field to given value.


### SetResolvedCopyNil

`func (o *DeploymentPlanReadCreativesInner) SetResolvedCopyNil(b bool)`

 SetResolvedCopyNil sets the value for ResolvedCopy to be an explicit nil

### UnsetResolvedCopy
`func (o *DeploymentPlanReadCreativesInner) UnsetResolvedCopy()`

UnsetResolvedCopy ensures that no value is present for ResolvedCopy, not even an explicit nil
### GetResolvedLandingPage

`func (o *DeploymentPlanReadCreativesInner) GetResolvedLandingPage() string`

GetResolvedLandingPage returns the ResolvedLandingPage field if non-nil, zero value otherwise.

### GetResolvedLandingPageOk

`func (o *DeploymentPlanReadCreativesInner) GetResolvedLandingPageOk() (*string, bool)`

GetResolvedLandingPageOk returns a tuple with the ResolvedLandingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedLandingPage

`func (o *DeploymentPlanReadCreativesInner) SetResolvedLandingPage(v string)`

SetResolvedLandingPage sets ResolvedLandingPage field to given value.


### GetCopyDiffers

`func (o *DeploymentPlanReadCreativesInner) GetCopyDiffers() bool`

GetCopyDiffers returns the CopyDiffers field if non-nil, zero value otherwise.

### GetCopyDiffersOk

`func (o *DeploymentPlanReadCreativesInner) GetCopyDiffersOk() (*bool, bool)`

GetCopyDiffersOk returns a tuple with the CopyDiffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyDiffers

`func (o *DeploymentPlanReadCreativesInner) SetCopyDiffers(v bool)`

SetCopyDiffers sets CopyDiffers field to given value.


### GetLandingPageDiffers

`func (o *DeploymentPlanReadCreativesInner) GetLandingPageDiffers() bool`

GetLandingPageDiffers returns the LandingPageDiffers field if non-nil, zero value otherwise.

### GetLandingPageDiffersOk

`func (o *DeploymentPlanReadCreativesInner) GetLandingPageDiffersOk() (*bool, bool)`

GetLandingPageDiffersOk returns a tuple with the LandingPageDiffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPageDiffers

`func (o *DeploymentPlanReadCreativesInner) SetLandingPageDiffers(v bool)`

SetLandingPageDiffers sets LandingPageDiffers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


