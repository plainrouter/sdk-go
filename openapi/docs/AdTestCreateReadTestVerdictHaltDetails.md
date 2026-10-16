# AdTestCreateReadTestVerdictHaltDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdId** | Pointer to **string** |  | [optional] 
**EffectiveStatus** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**LinkClicks** | Pointer to **int32** |  | [optional] 
**CountedArrivals** | Pointer to **int32** |  | [optional] 
**AdTestMemberId** | Pointer to **string** |  | [optional] 
**Baseline** | Pointer to [**AdTestCreateReadTestVerdictHaltDetailsBaseline**](AdTestCreateReadTestVerdictHaltDetailsBaseline.md) |  | [optional] 
**ObservedDestinations** | Pointer to [**[]AdTestCreateReadTestVerdictHaltDetailsObservedDestinationsInner**](AdTestCreateReadTestVerdictHaltDetailsObservedDestinationsInner.md) |  | [optional] 

## Methods

### NewAdTestCreateReadTestVerdictHaltDetails

`func NewAdTestCreateReadTestVerdictHaltDetails() *AdTestCreateReadTestVerdictHaltDetails`

NewAdTestCreateReadTestVerdictHaltDetails instantiates a new AdTestCreateReadTestVerdictHaltDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdTestCreateReadTestVerdictHaltDetailsWithDefaults

`func NewAdTestCreateReadTestVerdictHaltDetailsWithDefaults() *AdTestCreateReadTestVerdictHaltDetails`

NewAdTestCreateReadTestVerdictHaltDetailsWithDefaults instantiates a new AdTestCreateReadTestVerdictHaltDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdId

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetAdId() string`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetAdIdOk() (*string, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *AdTestCreateReadTestVerdictHaltDetails) SetAdId(v string)`

SetAdId sets AdId field to given value.

### HasAdId

`func (o *AdTestCreateReadTestVerdictHaltDetails) HasAdId() bool`

HasAdId returns a boolean if a field has been set.

### GetEffectiveStatus

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetEffectiveStatus() string`

GetEffectiveStatus returns the EffectiveStatus field if non-nil, zero value otherwise.

### GetEffectiveStatusOk

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetEffectiveStatusOk() (*string, bool)`

GetEffectiveStatusOk returns a tuple with the EffectiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveStatus

`func (o *AdTestCreateReadTestVerdictHaltDetails) SetEffectiveStatus(v string)`

SetEffectiveStatus sets EffectiveStatus field to given value.

### HasEffectiveStatus

`func (o *AdTestCreateReadTestVerdictHaltDetails) HasEffectiveStatus() bool`

HasEffectiveStatus returns a boolean if a field has been set.

### GetDate

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AdTestCreateReadTestVerdictHaltDetails) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AdTestCreateReadTestVerdictHaltDetails) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetLinkClicks

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetLinkClicks() int32`

GetLinkClicks returns the LinkClicks field if non-nil, zero value otherwise.

### GetLinkClicksOk

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetLinkClicksOk() (*int32, bool)`

GetLinkClicksOk returns a tuple with the LinkClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkClicks

`func (o *AdTestCreateReadTestVerdictHaltDetails) SetLinkClicks(v int32)`

SetLinkClicks sets LinkClicks field to given value.

### HasLinkClicks

`func (o *AdTestCreateReadTestVerdictHaltDetails) HasLinkClicks() bool`

HasLinkClicks returns a boolean if a field has been set.

### GetCountedArrivals

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetCountedArrivals() int32`

GetCountedArrivals returns the CountedArrivals field if non-nil, zero value otherwise.

### GetCountedArrivalsOk

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetCountedArrivalsOk() (*int32, bool)`

GetCountedArrivalsOk returns a tuple with the CountedArrivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountedArrivals

`func (o *AdTestCreateReadTestVerdictHaltDetails) SetCountedArrivals(v int32)`

SetCountedArrivals sets CountedArrivals field to given value.

### HasCountedArrivals

`func (o *AdTestCreateReadTestVerdictHaltDetails) HasCountedArrivals() bool`

HasCountedArrivals returns a boolean if a field has been set.

### GetAdTestMemberId

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetAdTestMemberId() string`

GetAdTestMemberId returns the AdTestMemberId field if non-nil, zero value otherwise.

### GetAdTestMemberIdOk

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetAdTestMemberIdOk() (*string, bool)`

GetAdTestMemberIdOk returns a tuple with the AdTestMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdTestMemberId

`func (o *AdTestCreateReadTestVerdictHaltDetails) SetAdTestMemberId(v string)`

SetAdTestMemberId sets AdTestMemberId field to given value.

### HasAdTestMemberId

`func (o *AdTestCreateReadTestVerdictHaltDetails) HasAdTestMemberId() bool`

HasAdTestMemberId returns a boolean if a field has been set.

### GetBaseline

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetBaseline() AdTestCreateReadTestVerdictHaltDetailsBaseline`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetBaselineOk() (*AdTestCreateReadTestVerdictHaltDetailsBaseline, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *AdTestCreateReadTestVerdictHaltDetails) SetBaseline(v AdTestCreateReadTestVerdictHaltDetailsBaseline)`

SetBaseline sets Baseline field to given value.

### HasBaseline

`func (o *AdTestCreateReadTestVerdictHaltDetails) HasBaseline() bool`

HasBaseline returns a boolean if a field has been set.

### GetObservedDestinations

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetObservedDestinations() []AdTestCreateReadTestVerdictHaltDetailsObservedDestinationsInner`

GetObservedDestinations returns the ObservedDestinations field if non-nil, zero value otherwise.

### GetObservedDestinationsOk

`func (o *AdTestCreateReadTestVerdictHaltDetails) GetObservedDestinationsOk() (*[]AdTestCreateReadTestVerdictHaltDetailsObservedDestinationsInner, bool)`

GetObservedDestinationsOk returns a tuple with the ObservedDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedDestinations

`func (o *AdTestCreateReadTestVerdictHaltDetails) SetObservedDestinations(v []AdTestCreateReadTestVerdictHaltDetailsObservedDestinationsInner)`

SetObservedDestinations sets ObservedDestinations field to given value.

### HasObservedDestinations

`func (o *AdTestCreateReadTestVerdictHaltDetails) HasObservedDestinations() bool`

HasObservedDestinations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


