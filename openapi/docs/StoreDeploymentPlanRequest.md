# StoreDeploymentPlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**PlatformAdAccountId** | **int32** |  | 
**Platform** | **string** |  | 
**CampaignRef** | Pointer to **NullableString** |  | [optional] 
**CampaignSpec** | Pointer to **[]string** |  | [optional] 
**AdSetRef** | Pointer to **NullableString** |  | [optional] 
**AdSetSpec** | Pointer to [**NullableStoreDeploymentPlanRequestAdSetSpec**](StoreDeploymentPlanRequestAdSetSpec.md) |  | [optional] 
**BudgetType** | **string** |  | 
**BudgetAmountMinor** | **int32** |  | 
**Currency** | **string** |  | 
**Copy** | [**StoreDeploymentPlanRequestCopy**](StoreDeploymentPlanRequestCopy.md) |  | 
**LandingPage** | **string** |  | 
**UtmPolicyName** | Pointer to **NullableString** |  | [optional] 
**NamingPolicyName** | Pointer to **NullableString** |  | [optional] 
**CreativeIds** | Pointer to **[]string** |  | [optional] 
**CreativeOverrides** | Pointer to [**[]StoreDeploymentPlanRequestCreativeOverridesInner**](StoreDeploymentPlanRequestCreativeOverridesInner.md) |  | [optional] 

## Methods

### NewStoreDeploymentPlanRequest

`func NewStoreDeploymentPlanRequest(platformAdAccountId int32, platform string, budgetType string, budgetAmountMinor int32, currency string, copy StoreDeploymentPlanRequestCopy, landingPage string, ) *StoreDeploymentPlanRequest`

NewStoreDeploymentPlanRequest instantiates a new StoreDeploymentPlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreDeploymentPlanRequestWithDefaults

`func NewStoreDeploymentPlanRequestWithDefaults() *StoreDeploymentPlanRequest`

NewStoreDeploymentPlanRequestWithDefaults instantiates a new StoreDeploymentPlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StoreDeploymentPlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoreDeploymentPlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoreDeploymentPlanRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoreDeploymentPlanRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StoreDeploymentPlanRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StoreDeploymentPlanRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPlatformAdAccountId

`func (o *StoreDeploymentPlanRequest) GetPlatformAdAccountId() int32`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *StoreDeploymentPlanRequest) GetPlatformAdAccountIdOk() (*int32, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *StoreDeploymentPlanRequest) SetPlatformAdAccountId(v int32)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.


### GetPlatform

`func (o *StoreDeploymentPlanRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *StoreDeploymentPlanRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *StoreDeploymentPlanRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetCampaignRef

`func (o *StoreDeploymentPlanRequest) GetCampaignRef() string`

GetCampaignRef returns the CampaignRef field if non-nil, zero value otherwise.

### GetCampaignRefOk

`func (o *StoreDeploymentPlanRequest) GetCampaignRefOk() (*string, bool)`

GetCampaignRefOk returns a tuple with the CampaignRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignRef

`func (o *StoreDeploymentPlanRequest) SetCampaignRef(v string)`

SetCampaignRef sets CampaignRef field to given value.

### HasCampaignRef

`func (o *StoreDeploymentPlanRequest) HasCampaignRef() bool`

HasCampaignRef returns a boolean if a field has been set.

### SetCampaignRefNil

`func (o *StoreDeploymentPlanRequest) SetCampaignRefNil(b bool)`

 SetCampaignRefNil sets the value for CampaignRef to be an explicit nil

### UnsetCampaignRef
`func (o *StoreDeploymentPlanRequest) UnsetCampaignRef()`

UnsetCampaignRef ensures that no value is present for CampaignRef, not even an explicit nil
### GetCampaignSpec

`func (o *StoreDeploymentPlanRequest) GetCampaignSpec() []string`

GetCampaignSpec returns the CampaignSpec field if non-nil, zero value otherwise.

### GetCampaignSpecOk

`func (o *StoreDeploymentPlanRequest) GetCampaignSpecOk() (*[]string, bool)`

GetCampaignSpecOk returns a tuple with the CampaignSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignSpec

`func (o *StoreDeploymentPlanRequest) SetCampaignSpec(v []string)`

SetCampaignSpec sets CampaignSpec field to given value.

### HasCampaignSpec

`func (o *StoreDeploymentPlanRequest) HasCampaignSpec() bool`

HasCampaignSpec returns a boolean if a field has been set.

### SetCampaignSpecNil

`func (o *StoreDeploymentPlanRequest) SetCampaignSpecNil(b bool)`

 SetCampaignSpecNil sets the value for CampaignSpec to be an explicit nil

### UnsetCampaignSpec
`func (o *StoreDeploymentPlanRequest) UnsetCampaignSpec()`

UnsetCampaignSpec ensures that no value is present for CampaignSpec, not even an explicit nil
### GetAdSetRef

`func (o *StoreDeploymentPlanRequest) GetAdSetRef() string`

GetAdSetRef returns the AdSetRef field if non-nil, zero value otherwise.

### GetAdSetRefOk

`func (o *StoreDeploymentPlanRequest) GetAdSetRefOk() (*string, bool)`

GetAdSetRefOk returns a tuple with the AdSetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSetRef

`func (o *StoreDeploymentPlanRequest) SetAdSetRef(v string)`

SetAdSetRef sets AdSetRef field to given value.

### HasAdSetRef

`func (o *StoreDeploymentPlanRequest) HasAdSetRef() bool`

HasAdSetRef returns a boolean if a field has been set.

### SetAdSetRefNil

`func (o *StoreDeploymentPlanRequest) SetAdSetRefNil(b bool)`

 SetAdSetRefNil sets the value for AdSetRef to be an explicit nil

### UnsetAdSetRef
`func (o *StoreDeploymentPlanRequest) UnsetAdSetRef()`

UnsetAdSetRef ensures that no value is present for AdSetRef, not even an explicit nil
### GetAdSetSpec

`func (o *StoreDeploymentPlanRequest) GetAdSetSpec() StoreDeploymentPlanRequestAdSetSpec`

GetAdSetSpec returns the AdSetSpec field if non-nil, zero value otherwise.

### GetAdSetSpecOk

`func (o *StoreDeploymentPlanRequest) GetAdSetSpecOk() (*StoreDeploymentPlanRequestAdSetSpec, bool)`

GetAdSetSpecOk returns a tuple with the AdSetSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSetSpec

`func (o *StoreDeploymentPlanRequest) SetAdSetSpec(v StoreDeploymentPlanRequestAdSetSpec)`

SetAdSetSpec sets AdSetSpec field to given value.

### HasAdSetSpec

`func (o *StoreDeploymentPlanRequest) HasAdSetSpec() bool`

HasAdSetSpec returns a boolean if a field has been set.

### SetAdSetSpecNil

`func (o *StoreDeploymentPlanRequest) SetAdSetSpecNil(b bool)`

 SetAdSetSpecNil sets the value for AdSetSpec to be an explicit nil

### UnsetAdSetSpec
`func (o *StoreDeploymentPlanRequest) UnsetAdSetSpec()`

UnsetAdSetSpec ensures that no value is present for AdSetSpec, not even an explicit nil
### GetBudgetType

`func (o *StoreDeploymentPlanRequest) GetBudgetType() string`

GetBudgetType returns the BudgetType field if non-nil, zero value otherwise.

### GetBudgetTypeOk

`func (o *StoreDeploymentPlanRequest) GetBudgetTypeOk() (*string, bool)`

GetBudgetTypeOk returns a tuple with the BudgetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetType

`func (o *StoreDeploymentPlanRequest) SetBudgetType(v string)`

SetBudgetType sets BudgetType field to given value.


### GetBudgetAmountMinor

`func (o *StoreDeploymentPlanRequest) GetBudgetAmountMinor() int32`

GetBudgetAmountMinor returns the BudgetAmountMinor field if non-nil, zero value otherwise.

### GetBudgetAmountMinorOk

`func (o *StoreDeploymentPlanRequest) GetBudgetAmountMinorOk() (*int32, bool)`

GetBudgetAmountMinorOk returns a tuple with the BudgetAmountMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetAmountMinor

`func (o *StoreDeploymentPlanRequest) SetBudgetAmountMinor(v int32)`

SetBudgetAmountMinor sets BudgetAmountMinor field to given value.


### GetCurrency

`func (o *StoreDeploymentPlanRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *StoreDeploymentPlanRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *StoreDeploymentPlanRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCopy

`func (o *StoreDeploymentPlanRequest) GetCopy() StoreDeploymentPlanRequestCopy`

GetCopy returns the Copy field if non-nil, zero value otherwise.

### GetCopyOk

`func (o *StoreDeploymentPlanRequest) GetCopyOk() (*StoreDeploymentPlanRequestCopy, bool)`

GetCopyOk returns a tuple with the Copy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopy

`func (o *StoreDeploymentPlanRequest) SetCopy(v StoreDeploymentPlanRequestCopy)`

SetCopy sets Copy field to given value.


### GetLandingPage

`func (o *StoreDeploymentPlanRequest) GetLandingPage() string`

GetLandingPage returns the LandingPage field if non-nil, zero value otherwise.

### GetLandingPageOk

`func (o *StoreDeploymentPlanRequest) GetLandingPageOk() (*string, bool)`

GetLandingPageOk returns a tuple with the LandingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPage

`func (o *StoreDeploymentPlanRequest) SetLandingPage(v string)`

SetLandingPage sets LandingPage field to given value.


### GetUtmPolicyName

`func (o *StoreDeploymentPlanRequest) GetUtmPolicyName() string`

GetUtmPolicyName returns the UtmPolicyName field if non-nil, zero value otherwise.

### GetUtmPolicyNameOk

`func (o *StoreDeploymentPlanRequest) GetUtmPolicyNameOk() (*string, bool)`

GetUtmPolicyNameOk returns a tuple with the UtmPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtmPolicyName

`func (o *StoreDeploymentPlanRequest) SetUtmPolicyName(v string)`

SetUtmPolicyName sets UtmPolicyName field to given value.

### HasUtmPolicyName

`func (o *StoreDeploymentPlanRequest) HasUtmPolicyName() bool`

HasUtmPolicyName returns a boolean if a field has been set.

### SetUtmPolicyNameNil

`func (o *StoreDeploymentPlanRequest) SetUtmPolicyNameNil(b bool)`

 SetUtmPolicyNameNil sets the value for UtmPolicyName to be an explicit nil

### UnsetUtmPolicyName
`func (o *StoreDeploymentPlanRequest) UnsetUtmPolicyName()`

UnsetUtmPolicyName ensures that no value is present for UtmPolicyName, not even an explicit nil
### GetNamingPolicyName

`func (o *StoreDeploymentPlanRequest) GetNamingPolicyName() string`

GetNamingPolicyName returns the NamingPolicyName field if non-nil, zero value otherwise.

### GetNamingPolicyNameOk

`func (o *StoreDeploymentPlanRequest) GetNamingPolicyNameOk() (*string, bool)`

GetNamingPolicyNameOk returns a tuple with the NamingPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingPolicyName

`func (o *StoreDeploymentPlanRequest) SetNamingPolicyName(v string)`

SetNamingPolicyName sets NamingPolicyName field to given value.

### HasNamingPolicyName

`func (o *StoreDeploymentPlanRequest) HasNamingPolicyName() bool`

HasNamingPolicyName returns a boolean if a field has been set.

### SetNamingPolicyNameNil

`func (o *StoreDeploymentPlanRequest) SetNamingPolicyNameNil(b bool)`

 SetNamingPolicyNameNil sets the value for NamingPolicyName to be an explicit nil

### UnsetNamingPolicyName
`func (o *StoreDeploymentPlanRequest) UnsetNamingPolicyName()`

UnsetNamingPolicyName ensures that no value is present for NamingPolicyName, not even an explicit nil
### GetCreativeIds

`func (o *StoreDeploymentPlanRequest) GetCreativeIds() []string`

GetCreativeIds returns the CreativeIds field if non-nil, zero value otherwise.

### GetCreativeIdsOk

`func (o *StoreDeploymentPlanRequest) GetCreativeIdsOk() (*[]string, bool)`

GetCreativeIdsOk returns a tuple with the CreativeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeIds

`func (o *StoreDeploymentPlanRequest) SetCreativeIds(v []string)`

SetCreativeIds sets CreativeIds field to given value.

### HasCreativeIds

`func (o *StoreDeploymentPlanRequest) HasCreativeIds() bool`

HasCreativeIds returns a boolean if a field has been set.

### GetCreativeOverrides

`func (o *StoreDeploymentPlanRequest) GetCreativeOverrides() []StoreDeploymentPlanRequestCreativeOverridesInner`

GetCreativeOverrides returns the CreativeOverrides field if non-nil, zero value otherwise.

### GetCreativeOverridesOk

`func (o *StoreDeploymentPlanRequest) GetCreativeOverridesOk() (*[]StoreDeploymentPlanRequestCreativeOverridesInner, bool)`

GetCreativeOverridesOk returns a tuple with the CreativeOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeOverrides

`func (o *StoreDeploymentPlanRequest) SetCreativeOverrides(v []StoreDeploymentPlanRequestCreativeOverridesInner)`

SetCreativeOverrides sets CreativeOverrides field to given value.

### HasCreativeOverrides

`func (o *StoreDeploymentPlanRequest) HasCreativeOverrides() bool`

HasCreativeOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


