# GetInventoryMetrics200ResponseMetricsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**PlatformAdAccountId** | **int32** |  | 
**AdExternalId** | **string** |  | 
**CampaignExternalId** | **NullableString** |  | 
**AdsetExternalId** | **NullableString** |  | 
**Date** | **string** |  | 
**AccountCurrency** | **string** |  | 
**Spend** | **string** |  | 
**Impressions** | **int32** |  | 
**Clicks** | **int32** |  | 
**InlineLinkClicks** | **int32** |  | 
**Actions** | **interface{}** |  | 
**ActionValues** | **interface{}** |  | 
**FetchedAt** | **NullableString** |  | 
**Arrivals** | **NullableInt32** |  | 
**CreativeId** | **NullableString** |  | 
**Generator** | **NullableString** |  | 
**GeneratorJob** | **NullableString** |  | 
**ParentCreative** | **NullableString** |  | 

## Methods

### NewGetInventoryMetrics200ResponseMetricsInner

`func NewGetInventoryMetrics200ResponseMetricsInner(id int32, platformAdAccountId int32, adExternalId string, campaignExternalId NullableString, adsetExternalId NullableString, date string, accountCurrency string, spend string, impressions int32, clicks int32, inlineLinkClicks int32, actions interface{}, actionValues interface{}, fetchedAt NullableString, arrivals NullableInt32, creativeId NullableString, generator NullableString, generatorJob NullableString, parentCreative NullableString, ) *GetInventoryMetrics200ResponseMetricsInner`

NewGetInventoryMetrics200ResponseMetricsInner instantiates a new GetInventoryMetrics200ResponseMetricsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInventoryMetrics200ResponseMetricsInnerWithDefaults

`func NewGetInventoryMetrics200ResponseMetricsInnerWithDefaults() *GetInventoryMetrics200ResponseMetricsInner`

NewGetInventoryMetrics200ResponseMetricsInnerWithDefaults instantiates a new GetInventoryMetrics200ResponseMetricsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetPlatformAdAccountId

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetPlatformAdAccountId() int32`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetPlatformAdAccountIdOk() (*int32, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetPlatformAdAccountId(v int32)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.


### GetAdExternalId

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetAdExternalId() string`

GetAdExternalId returns the AdExternalId field if non-nil, zero value otherwise.

### GetAdExternalIdOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetAdExternalIdOk() (*string, bool)`

GetAdExternalIdOk returns a tuple with the AdExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdExternalId

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetAdExternalId(v string)`

SetAdExternalId sets AdExternalId field to given value.


### GetCampaignExternalId

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetCampaignExternalId() string`

GetCampaignExternalId returns the CampaignExternalId field if non-nil, zero value otherwise.

### GetCampaignExternalIdOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetCampaignExternalIdOk() (*string, bool)`

GetCampaignExternalIdOk returns a tuple with the CampaignExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignExternalId

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetCampaignExternalId(v string)`

SetCampaignExternalId sets CampaignExternalId field to given value.


### SetCampaignExternalIdNil

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetCampaignExternalIdNil(b bool)`

 SetCampaignExternalIdNil sets the value for CampaignExternalId to be an explicit nil

### UnsetCampaignExternalId
`func (o *GetInventoryMetrics200ResponseMetricsInner) UnsetCampaignExternalId()`

UnsetCampaignExternalId ensures that no value is present for CampaignExternalId, not even an explicit nil
### GetAdsetExternalId

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetAdsetExternalId() string`

GetAdsetExternalId returns the AdsetExternalId field if non-nil, zero value otherwise.

### GetAdsetExternalIdOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetAdsetExternalIdOk() (*string, bool)`

GetAdsetExternalIdOk returns a tuple with the AdsetExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdsetExternalId

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetAdsetExternalId(v string)`

SetAdsetExternalId sets AdsetExternalId field to given value.


### SetAdsetExternalIdNil

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetAdsetExternalIdNil(b bool)`

 SetAdsetExternalIdNil sets the value for AdsetExternalId to be an explicit nil

### UnsetAdsetExternalId
`func (o *GetInventoryMetrics200ResponseMetricsInner) UnsetAdsetExternalId()`

UnsetAdsetExternalId ensures that no value is present for AdsetExternalId, not even an explicit nil
### GetDate

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetAccountCurrency

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetAccountCurrency() string`

GetAccountCurrency returns the AccountCurrency field if non-nil, zero value otherwise.

### GetAccountCurrencyOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetAccountCurrencyOk() (*string, bool)`

GetAccountCurrencyOk returns a tuple with the AccountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCurrency

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetAccountCurrency(v string)`

SetAccountCurrency sets AccountCurrency field to given value.


### GetSpend

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetSpend() string`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetSpendOk() (*string, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetSpend(v string)`

SetSpend sets Spend field to given value.


### GetImpressions

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetImpressions() int32`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetImpressionsOk() (*int32, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetImpressions(v int32)`

SetImpressions sets Impressions field to given value.


### GetClicks

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetClicks() int32`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetClicksOk() (*int32, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetClicks(v int32)`

SetClicks sets Clicks field to given value.


### GetInlineLinkClicks

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetInlineLinkClicks() int32`

GetInlineLinkClicks returns the InlineLinkClicks field if non-nil, zero value otherwise.

### GetInlineLinkClicksOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetInlineLinkClicksOk() (*int32, bool)`

GetInlineLinkClicksOk returns a tuple with the InlineLinkClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineLinkClicks

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetInlineLinkClicks(v int32)`

SetInlineLinkClicks sets InlineLinkClicks field to given value.


### GetActions

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetActions() interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetActionsOk() (*interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetActions(v interface{})`

SetActions sets Actions field to given value.


### SetActionsNil

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *GetInventoryMetrics200ResponseMetricsInner) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetActionValues

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetActionValues() interface{}`

GetActionValues returns the ActionValues field if non-nil, zero value otherwise.

### GetActionValuesOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetActionValuesOk() (*interface{}, bool)`

GetActionValuesOk returns a tuple with the ActionValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValues

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetActionValues(v interface{})`

SetActionValues sets ActionValues field to given value.


### SetActionValuesNil

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetActionValuesNil(b bool)`

 SetActionValuesNil sets the value for ActionValues to be an explicit nil

### UnsetActionValues
`func (o *GetInventoryMetrics200ResponseMetricsInner) UnsetActionValues()`

UnsetActionValues ensures that no value is present for ActionValues, not even an explicit nil
### GetFetchedAt

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetFetchedAt() string`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetFetchedAtOk() (*string, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetFetchedAt(v string)`

SetFetchedAt sets FetchedAt field to given value.


### SetFetchedAtNil

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetFetchedAtNil(b bool)`

 SetFetchedAtNil sets the value for FetchedAt to be an explicit nil

### UnsetFetchedAt
`func (o *GetInventoryMetrics200ResponseMetricsInner) UnsetFetchedAt()`

UnsetFetchedAt ensures that no value is present for FetchedAt, not even an explicit nil
### GetArrivals

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetArrivals() int32`

GetArrivals returns the Arrivals field if non-nil, zero value otherwise.

### GetArrivalsOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetArrivalsOk() (*int32, bool)`

GetArrivalsOk returns a tuple with the Arrivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivals

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetArrivals(v int32)`

SetArrivals sets Arrivals field to given value.


### SetArrivalsNil

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetArrivalsNil(b bool)`

 SetArrivalsNil sets the value for Arrivals to be an explicit nil

### UnsetArrivals
`func (o *GetInventoryMetrics200ResponseMetricsInner) UnsetArrivals()`

UnsetArrivals ensures that no value is present for Arrivals, not even an explicit nil
### GetCreativeId

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetCreativeId() string`

GetCreativeId returns the CreativeId field if non-nil, zero value otherwise.

### GetCreativeIdOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetCreativeIdOk() (*string, bool)`

GetCreativeIdOk returns a tuple with the CreativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeId

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetCreativeId(v string)`

SetCreativeId sets CreativeId field to given value.


### SetCreativeIdNil

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetCreativeIdNil(b bool)`

 SetCreativeIdNil sets the value for CreativeId to be an explicit nil

### UnsetCreativeId
`func (o *GetInventoryMetrics200ResponseMetricsInner) UnsetCreativeId()`

UnsetCreativeId ensures that no value is present for CreativeId, not even an explicit nil
### GetGenerator

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetGenerator() string`

GetGenerator returns the Generator field if non-nil, zero value otherwise.

### GetGeneratorOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetGeneratorOk() (*string, bool)`

GetGeneratorOk returns a tuple with the Generator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerator

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetGenerator(v string)`

SetGenerator sets Generator field to given value.


### SetGeneratorNil

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetGeneratorNil(b bool)`

 SetGeneratorNil sets the value for Generator to be an explicit nil

### UnsetGenerator
`func (o *GetInventoryMetrics200ResponseMetricsInner) UnsetGenerator()`

UnsetGenerator ensures that no value is present for Generator, not even an explicit nil
### GetGeneratorJob

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetGeneratorJob() string`

GetGeneratorJob returns the GeneratorJob field if non-nil, zero value otherwise.

### GetGeneratorJobOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetGeneratorJobOk() (*string, bool)`

GetGeneratorJobOk returns a tuple with the GeneratorJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorJob

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetGeneratorJob(v string)`

SetGeneratorJob sets GeneratorJob field to given value.


### SetGeneratorJobNil

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetGeneratorJobNil(b bool)`

 SetGeneratorJobNil sets the value for GeneratorJob to be an explicit nil

### UnsetGeneratorJob
`func (o *GetInventoryMetrics200ResponseMetricsInner) UnsetGeneratorJob()`

UnsetGeneratorJob ensures that no value is present for GeneratorJob, not even an explicit nil
### GetParentCreative

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetParentCreative() string`

GetParentCreative returns the ParentCreative field if non-nil, zero value otherwise.

### GetParentCreativeOk

`func (o *GetInventoryMetrics200ResponseMetricsInner) GetParentCreativeOk() (*string, bool)`

GetParentCreativeOk returns a tuple with the ParentCreative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCreative

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetParentCreative(v string)`

SetParentCreative sets ParentCreative field to given value.


### SetParentCreativeNil

`func (o *GetInventoryMetrics200ResponseMetricsInner) SetParentCreativeNil(b bool)`

 SetParentCreativeNil sets the value for ParentCreative to be an explicit nil

### UnsetParentCreative
`func (o *GetInventoryMetrics200ResponseMetricsInner) UnsetParentCreative()`

UnsetParentCreative ensures that no value is present for ParentCreative, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


