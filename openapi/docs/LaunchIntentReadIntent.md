# LaunchIntentReadIntent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Status** | **string** |  | 
**ActionBatchId** | **NullableString** |  | 
**PolicyReasons** | **[]string** |  | 
**ProviderObjectIds** | **interface{}** |  | 
**CampaignId** | **NullableString** |  | 
**AdSetId** | **NullableString** |  | 
**Currency** | **NullableString** |  | 
**Actions** | [**[]LaunchIntentReadIntentActionsInner**](LaunchIntentReadIntentActionsInner.md) |  | 

## Methods

### NewLaunchIntentReadIntent

`func NewLaunchIntentReadIntent(id string, status string, actionBatchId NullableString, policyReasons []string, providerObjectIds interface{}, campaignId NullableString, adSetId NullableString, currency NullableString, actions []LaunchIntentReadIntentActionsInner, ) *LaunchIntentReadIntent`

NewLaunchIntentReadIntent instantiates a new LaunchIntentReadIntent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchIntentReadIntentWithDefaults

`func NewLaunchIntentReadIntentWithDefaults() *LaunchIntentReadIntent`

NewLaunchIntentReadIntentWithDefaults instantiates a new LaunchIntentReadIntent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LaunchIntentReadIntent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LaunchIntentReadIntent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LaunchIntentReadIntent) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *LaunchIntentReadIntent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LaunchIntentReadIntent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LaunchIntentReadIntent) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetActionBatchId

`func (o *LaunchIntentReadIntent) GetActionBatchId() string`

GetActionBatchId returns the ActionBatchId field if non-nil, zero value otherwise.

### GetActionBatchIdOk

`func (o *LaunchIntentReadIntent) GetActionBatchIdOk() (*string, bool)`

GetActionBatchIdOk returns a tuple with the ActionBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionBatchId

`func (o *LaunchIntentReadIntent) SetActionBatchId(v string)`

SetActionBatchId sets ActionBatchId field to given value.


### SetActionBatchIdNil

`func (o *LaunchIntentReadIntent) SetActionBatchIdNil(b bool)`

 SetActionBatchIdNil sets the value for ActionBatchId to be an explicit nil

### UnsetActionBatchId
`func (o *LaunchIntentReadIntent) UnsetActionBatchId()`

UnsetActionBatchId ensures that no value is present for ActionBatchId, not even an explicit nil
### GetPolicyReasons

`func (o *LaunchIntentReadIntent) GetPolicyReasons() []string`

GetPolicyReasons returns the PolicyReasons field if non-nil, zero value otherwise.

### GetPolicyReasonsOk

`func (o *LaunchIntentReadIntent) GetPolicyReasonsOk() (*[]string, bool)`

GetPolicyReasonsOk returns a tuple with the PolicyReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyReasons

`func (o *LaunchIntentReadIntent) SetPolicyReasons(v []string)`

SetPolicyReasons sets PolicyReasons field to given value.


### GetProviderObjectIds

`func (o *LaunchIntentReadIntent) GetProviderObjectIds() interface{}`

GetProviderObjectIds returns the ProviderObjectIds field if non-nil, zero value otherwise.

### GetProviderObjectIdsOk

`func (o *LaunchIntentReadIntent) GetProviderObjectIdsOk() (*interface{}, bool)`

GetProviderObjectIdsOk returns a tuple with the ProviderObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderObjectIds

`func (o *LaunchIntentReadIntent) SetProviderObjectIds(v interface{})`

SetProviderObjectIds sets ProviderObjectIds field to given value.


### SetProviderObjectIdsNil

`func (o *LaunchIntentReadIntent) SetProviderObjectIdsNil(b bool)`

 SetProviderObjectIdsNil sets the value for ProviderObjectIds to be an explicit nil

### UnsetProviderObjectIds
`func (o *LaunchIntentReadIntent) UnsetProviderObjectIds()`

UnsetProviderObjectIds ensures that no value is present for ProviderObjectIds, not even an explicit nil
### GetCampaignId

`func (o *LaunchIntentReadIntent) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *LaunchIntentReadIntent) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *LaunchIntentReadIntent) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.


### SetCampaignIdNil

`func (o *LaunchIntentReadIntent) SetCampaignIdNil(b bool)`

 SetCampaignIdNil sets the value for CampaignId to be an explicit nil

### UnsetCampaignId
`func (o *LaunchIntentReadIntent) UnsetCampaignId()`

UnsetCampaignId ensures that no value is present for CampaignId, not even an explicit nil
### GetAdSetId

`func (o *LaunchIntentReadIntent) GetAdSetId() string`

GetAdSetId returns the AdSetId field if non-nil, zero value otherwise.

### GetAdSetIdOk

`func (o *LaunchIntentReadIntent) GetAdSetIdOk() (*string, bool)`

GetAdSetIdOk returns a tuple with the AdSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSetId

`func (o *LaunchIntentReadIntent) SetAdSetId(v string)`

SetAdSetId sets AdSetId field to given value.


### SetAdSetIdNil

`func (o *LaunchIntentReadIntent) SetAdSetIdNil(b bool)`

 SetAdSetIdNil sets the value for AdSetId to be an explicit nil

### UnsetAdSetId
`func (o *LaunchIntentReadIntent) UnsetAdSetId()`

UnsetAdSetId ensures that no value is present for AdSetId, not even an explicit nil
### GetCurrency

`func (o *LaunchIntentReadIntent) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LaunchIntentReadIntent) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LaunchIntentReadIntent) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *LaunchIntentReadIntent) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *LaunchIntentReadIntent) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetActions

`func (o *LaunchIntentReadIntent) GetActions() []LaunchIntentReadIntentActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *LaunchIntentReadIntent) GetActionsOk() (*[]LaunchIntentReadIntentActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *LaunchIntentReadIntent) SetActions(v []LaunchIntentReadIntentActionsInner)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


