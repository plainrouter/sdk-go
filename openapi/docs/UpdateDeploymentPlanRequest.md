# UpdateDeploymentPlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**PlatformAdAccountId** | Pointer to **int32** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**CampaignRef** | Pointer to **NullableString** |  | [optional] 
**CampaignSpec** | Pointer to **[]string** |  | [optional] 
**AdSetRef** | Pointer to **NullableString** |  | [optional] 
**AdSetSpec** | Pointer to [**NullableStoreDeploymentPlanRequestAdSetSpec**](StoreDeploymentPlanRequestAdSetSpec.md) |  | [optional] 
**BudgetType** | Pointer to **string** |  | [optional] 
**BudgetAmountMinor** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Copy** | Pointer to [**StoreDeploymentPlanRequestCopy**](StoreDeploymentPlanRequestCopy.md) |  | [optional] 
**LandingPage** | Pointer to **string** |  | [optional] 
**UtmPolicyName** | Pointer to **NullableString** |  | [optional] 
**NamingPolicyName** | Pointer to **NullableString** |  | [optional] 
**CreativeIds** | Pointer to **[]string** |  | [optional] 
**CreativeOverrides** | Pointer to [**[]StoreDeploymentPlanRequestCreativeOverridesInner**](StoreDeploymentPlanRequestCreativeOverridesInner.md) |  | [optional] 

## Methods

### NewUpdateDeploymentPlanRequest

`func NewUpdateDeploymentPlanRequest() *UpdateDeploymentPlanRequest`

NewUpdateDeploymentPlanRequest instantiates a new UpdateDeploymentPlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeploymentPlanRequestWithDefaults

`func NewUpdateDeploymentPlanRequestWithDefaults() *UpdateDeploymentPlanRequest`

NewUpdateDeploymentPlanRequestWithDefaults instantiates a new UpdateDeploymentPlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDeploymentPlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDeploymentPlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDeploymentPlanRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDeploymentPlanRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateDeploymentPlanRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateDeploymentPlanRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPlatformAdAccountId

`func (o *UpdateDeploymentPlanRequest) GetPlatformAdAccountId() int32`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *UpdateDeploymentPlanRequest) GetPlatformAdAccountIdOk() (*int32, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *UpdateDeploymentPlanRequest) SetPlatformAdAccountId(v int32)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.

### HasPlatformAdAccountId

`func (o *UpdateDeploymentPlanRequest) HasPlatformAdAccountId() bool`

HasPlatformAdAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *UpdateDeploymentPlanRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdateDeploymentPlanRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdateDeploymentPlanRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *UpdateDeploymentPlanRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetCampaignRef

`func (o *UpdateDeploymentPlanRequest) GetCampaignRef() string`

GetCampaignRef returns the CampaignRef field if non-nil, zero value otherwise.

### GetCampaignRefOk

`func (o *UpdateDeploymentPlanRequest) GetCampaignRefOk() (*string, bool)`

GetCampaignRefOk returns a tuple with the CampaignRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignRef

`func (o *UpdateDeploymentPlanRequest) SetCampaignRef(v string)`

SetCampaignRef sets CampaignRef field to given value.

### HasCampaignRef

`func (o *UpdateDeploymentPlanRequest) HasCampaignRef() bool`

HasCampaignRef returns a boolean if a field has been set.

### SetCampaignRefNil

`func (o *UpdateDeploymentPlanRequest) SetCampaignRefNil(b bool)`

 SetCampaignRefNil sets the value for CampaignRef to be an explicit nil

### UnsetCampaignRef
`func (o *UpdateDeploymentPlanRequest) UnsetCampaignRef()`

UnsetCampaignRef ensures that no value is present for CampaignRef, not even an explicit nil
### GetCampaignSpec

`func (o *UpdateDeploymentPlanRequest) GetCampaignSpec() []string`

GetCampaignSpec returns the CampaignSpec field if non-nil, zero value otherwise.

### GetCampaignSpecOk

`func (o *UpdateDeploymentPlanRequest) GetCampaignSpecOk() (*[]string, bool)`

GetCampaignSpecOk returns a tuple with the CampaignSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignSpec

`func (o *UpdateDeploymentPlanRequest) SetCampaignSpec(v []string)`

SetCampaignSpec sets CampaignSpec field to given value.

### HasCampaignSpec

`func (o *UpdateDeploymentPlanRequest) HasCampaignSpec() bool`

HasCampaignSpec returns a boolean if a field has been set.

### SetCampaignSpecNil

`func (o *UpdateDeploymentPlanRequest) SetCampaignSpecNil(b bool)`

 SetCampaignSpecNil sets the value for CampaignSpec to be an explicit nil

### UnsetCampaignSpec
`func (o *UpdateDeploymentPlanRequest) UnsetCampaignSpec()`

UnsetCampaignSpec ensures that no value is present for CampaignSpec, not even an explicit nil
### GetAdSetRef

`func (o *UpdateDeploymentPlanRequest) GetAdSetRef() string`

GetAdSetRef returns the AdSetRef field if non-nil, zero value otherwise.

### GetAdSetRefOk

`func (o *UpdateDeploymentPlanRequest) GetAdSetRefOk() (*string, bool)`

GetAdSetRefOk returns a tuple with the AdSetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSetRef

`func (o *UpdateDeploymentPlanRequest) SetAdSetRef(v string)`

SetAdSetRef sets AdSetRef field to given value.

### HasAdSetRef

`func (o *UpdateDeploymentPlanRequest) HasAdSetRef() bool`

HasAdSetRef returns a boolean if a field has been set.

### SetAdSetRefNil

`func (o *UpdateDeploymentPlanRequest) SetAdSetRefNil(b bool)`

 SetAdSetRefNil sets the value for AdSetRef to be an explicit nil

### UnsetAdSetRef
`func (o *UpdateDeploymentPlanRequest) UnsetAdSetRef()`

UnsetAdSetRef ensures that no value is present for AdSetRef, not even an explicit nil
### GetAdSetSpec

`func (o *UpdateDeploymentPlanRequest) GetAdSetSpec() StoreDeploymentPlanRequestAdSetSpec`

GetAdSetSpec returns the AdSetSpec field if non-nil, zero value otherwise.

### GetAdSetSpecOk

`func (o *UpdateDeploymentPlanRequest) GetAdSetSpecOk() (*StoreDeploymentPlanRequestAdSetSpec, bool)`

GetAdSetSpecOk returns a tuple with the AdSetSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSetSpec

`func (o *UpdateDeploymentPlanRequest) SetAdSetSpec(v StoreDeploymentPlanRequestAdSetSpec)`

SetAdSetSpec sets AdSetSpec field to given value.

### HasAdSetSpec

`func (o *UpdateDeploymentPlanRequest) HasAdSetSpec() bool`

HasAdSetSpec returns a boolean if a field has been set.

### SetAdSetSpecNil

`func (o *UpdateDeploymentPlanRequest) SetAdSetSpecNil(b bool)`

 SetAdSetSpecNil sets the value for AdSetSpec to be an explicit nil

### UnsetAdSetSpec
`func (o *UpdateDeploymentPlanRequest) UnsetAdSetSpec()`

UnsetAdSetSpec ensures that no value is present for AdSetSpec, not even an explicit nil
### GetBudgetType

`func (o *UpdateDeploymentPlanRequest) GetBudgetType() string`

GetBudgetType returns the BudgetType field if non-nil, zero value otherwise.

### GetBudgetTypeOk

`func (o *UpdateDeploymentPlanRequest) GetBudgetTypeOk() (*string, bool)`

GetBudgetTypeOk returns a tuple with the BudgetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetType

`func (o *UpdateDeploymentPlanRequest) SetBudgetType(v string)`

SetBudgetType sets BudgetType field to given value.

### HasBudgetType

`func (o *UpdateDeploymentPlanRequest) HasBudgetType() bool`

HasBudgetType returns a boolean if a field has been set.

### GetBudgetAmountMinor

`func (o *UpdateDeploymentPlanRequest) GetBudgetAmountMinor() int32`

GetBudgetAmountMinor returns the BudgetAmountMinor field if non-nil, zero value otherwise.

### GetBudgetAmountMinorOk

`func (o *UpdateDeploymentPlanRequest) GetBudgetAmountMinorOk() (*int32, bool)`

GetBudgetAmountMinorOk returns a tuple with the BudgetAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetAmountMinor

`func (o *UpdateDeploymentPlanRequest) SetBudgetAmountMinor(v int32)`

SetBudgetAmountMinor sets BudgetAmountMinor field to given value.

### HasBudgetAmountMinor

`func (o *UpdateDeploymentPlanRequest) HasBudgetAmountMinor() bool`

HasBudgetAmountMinor returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdateDeploymentPlanRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateDeploymentPlanRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateDeploymentPlanRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdateDeploymentPlanRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCopy

`func (o *UpdateDeploymentPlanRequest) GetCopy() StoreDeploymentPlanRequestCopy`

GetCopy returns the Copy field if non-nil, zero value otherwise.

### GetCopyOk

`func (o *UpdateDeploymentPlanRequest) GetCopyOk() (*StoreDeploymentPlanRequestCopy, bool)`

GetCopyOk returns a tuple with the Copy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopy

`func (o *UpdateDeploymentPlanRequest) SetCopy(v StoreDeploymentPlanRequestCopy)`

SetCopy sets Copy field to given value.

### HasCopy

`func (o *UpdateDeploymentPlanRequest) HasCopy() bool`

HasCopy returns a boolean if a field has been set.

### GetLandingPage

`func (o *UpdateDeploymentPlanRequest) GetLandingPage() string`

GetLandingPage returns the LandingPage field if non-nil, zero value otherwise.

### GetLandingPageOk

`func (o *UpdateDeploymentPlanRequest) GetLandingPageOk() (*string, bool)`

GetLandingPageOk returns a tuple with the LandingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPage

`func (o *UpdateDeploymentPlanRequest) SetLandingPage(v string)`

SetLandingPage sets LandingPage field to given value.

### HasLandingPage

`func (o *UpdateDeploymentPlanRequest) HasLandingPage() bool`

HasLandingPage returns a boolean if a field has been set.

### GetUtmPolicyName

`func (o *UpdateDeploymentPlanRequest) GetUtmPolicyName() string`

GetUtmPolicyName returns the UtmPolicyName field if non-nil, zero value otherwise.

### GetUtmPolicyNameOk

`func (o *UpdateDeploymentPlanRequest) GetUtmPolicyNameOk() (*string, bool)`

GetUtmPolicyNameOk returns a tuple with the UtmPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmPolicyName

`func (o *UpdateDeploymentPlanRequest) SetUtmPolicyName(v string)`

SetUtmPolicyName sets UtmPolicyName field to given value.

### HasUtmPolicyName

`func (o *UpdateDeploymentPlanRequest) HasUtmPolicyName() bool`

HasUtmPolicyName returns a boolean if a field has been set.

### SetUtmPolicyNameNil

`func (o *UpdateDeploymentPlanRequest) SetUtmPolicyNameNil(b bool)`

 SetUtmPolicyNameNil sets the value for UtmPolicyName to be an explicit nil

### UnsetUtmPolicyName
`func (o *UpdateDeploymentPlanRequest) UnsetUtmPolicyName()`

UnsetUtmPolicyName ensures that no value is present for UtmPolicyName, not even an explicit nil
### GetNamingPolicyName

`func (o *UpdateDeploymentPlanRequest) GetNamingPolicyName() string`

GetNamingPolicyName returns the NamingPolicyName field if non-nil, zero value otherwise.

### GetNamingPolicyNameOk

`func (o *UpdateDeploymentPlanRequest) GetNamingPolicyNameOk() (*string, bool)`

GetNamingPolicyNameOk returns a tuple with the NamingPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingPolicyName

`func (o *UpdateDeploymentPlanRequest) SetNamingPolicyName(v string)`

SetNamingPolicyName sets NamingPolicyName field to given value.

### HasNamingPolicyName

`func (o *UpdateDeploymentPlanRequest) HasNamingPolicyName() bool`

HasNamingPolicyName returns a boolean if a field has been set.

### SetNamingPolicyNameNil

`func (o *UpdateDeploymentPlanRequest) SetNamingPolicyNameNil(b bool)`

 SetNamingPolicyNameNil sets the value for NamingPolicyName to be an explicit nil

### UnsetNamingPolicyName
`func (o *UpdateDeploymentPlanRequest) UnsetNamingPolicyName()`

UnsetNamingPolicyName ensures that no value is present for NamingPolicyName, not even an explicit nil
### GetCreativeIds

`func (o *UpdateDeploymentPlanRequest) GetCreativeIds() []string`

GetCreativeIds returns the CreativeIds field if non-nil, zero value otherwise.

### GetCreativeIdsOk

`func (o *UpdateDeploymentPlanRequest) GetCreativeIdsOk() (*[]string, bool)`

GetCreativeIdsOk returns a tuple with the CreativeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeIds

`func (o *UpdateDeploymentPlanRequest) SetCreativeIds(v []string)`

SetCreativeIds sets CreativeIds field to given value.

### HasCreativeIds

`func (o *UpdateDeploymentPlanRequest) HasCreativeIds() bool`

HasCreativeIds returns a boolean if a field has been set.

### GetCreativeOverrides

`func (o *UpdateDeploymentPlanRequest) GetCreativeOverrides() []StoreDeploymentPlanRequestCreativeOverridesInner`

GetCreativeOverrides returns the CreativeOverrides field if non-nil, zero value otherwise.

### GetCreativeOverridesOk

`func (o *UpdateDeploymentPlanRequest) GetCreativeOverridesOk() (*[]StoreDeploymentPlanRequestCreativeOverridesInner, bool)`

GetCreativeOverridesOk returns a tuple with the CreativeOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeOverrides

`func (o *UpdateDeploymentPlanRequest) SetCreativeOverrides(v []StoreDeploymentPlanRequestCreativeOverridesInner)`

SetCreativeOverrides sets CreativeOverrides field to given value.

### HasCreativeOverrides

`func (o *UpdateDeploymentPlanRequest) HasCreativeOverrides() bool`

HasCreativeOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


