# AdTestCreateReadTestVerdictDailyFactsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**AdTestMemberId** | **string** |  | 
**AdId** | **string** |  | 
**SpendMinor** | **string** | Exact spend in Meta minor units for the account currency, including any stored fractional component. | 
**Spend** | **NullableString** | Account-currency spend in major units. | 
**Impressions** | **NullableInt32** |  | 
**InlineLinkClicks** | Pointer to **NullableInt32** |  | [optional] 
**CountedArrivals** | **NullableInt32** |  | 
**RecordedConversions** | **NullableInt32** |  | 
**AdmittedConversions** | **NullableInt32** |  | 

## Methods

### NewAdTestCreateReadTestVerdictDailyFactsInner

`func NewAdTestCreateReadTestVerdictDailyFactsInner(date string, adTestMemberId string, adId string, spendMinor string, spend NullableString, impressions NullableInt32, countedArrivals NullableInt32, recordedConversions NullableInt32, admittedConversions NullableInt32, ) *AdTestCreateReadTestVerdictDailyFactsInner`

NewAdTestCreateReadTestVerdictDailyFactsInner instantiates a new AdTestCreateReadTestVerdictDailyFactsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdTestCreateReadTestVerdictDailyFactsInnerWithDefaults

`func NewAdTestCreateReadTestVerdictDailyFactsInnerWithDefaults() *AdTestCreateReadTestVerdictDailyFactsInner`

NewAdTestCreateReadTestVerdictDailyFactsInnerWithDefaults instantiates a new AdTestCreateReadTestVerdictDailyFactsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetAdTestMemberId

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetAdTestMemberId() string`

GetAdTestMemberId returns the AdTestMemberId field if non-nil, zero value otherwise.

### GetAdTestMemberIdOk

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetAdTestMemberIdOk() (*string, bool)`

GetAdTestMemberIdOk returns a tuple with the AdTestMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdTestMemberId

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetAdTestMemberId(v string)`

SetAdTestMemberId sets AdTestMemberId field to given value.


### GetAdId

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetAdId() string`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetAdIdOk() (*string, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetAdId(v string)`

SetAdId sets AdId field to given value.


### GetSpendMinor

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetSpendMinor() string`

GetSpendMinor returns the SpendMinor field if non-nil, zero value otherwise.

### GetSpendMinorOk

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetSpendMinorOk() (*string, bool)`

GetSpendMinorOk returns a tuple with the SpendMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendMinor

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetSpendMinor(v string)`

SetSpendMinor sets SpendMinor field to given value.


### GetSpend

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetSpend() string`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetSpendOk() (*string, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetSpend(v string)`

SetSpend sets Spend field to given value.


### SetSpendNil

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetSpendNil(b bool)`

 SetSpendNil sets the value for Spend to be an explicit nil

### UnsetSpend
`func (o *AdTestCreateReadTestVerdictDailyFactsInner) UnsetSpend()`

UnsetSpend ensures that no value is present for Spend, not even an explicit nil
### GetImpressions

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetImpressions() int32`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetImpressionsOk() (*int32, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetImpressions(v int32)`

SetImpressions sets Impressions field to given value.


### SetImpressionsNil

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetImpressionsNil(b bool)`

 SetImpressionsNil sets the value for Impressions to be an explicit nil

### UnsetImpressions
`func (o *AdTestCreateReadTestVerdictDailyFactsInner) UnsetImpressions()`

UnsetImpressions ensures that no value is present for Impressions, not even an explicit nil
### GetInlineLinkClicks

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetInlineLinkClicks() int32`

GetInlineLinkClicks returns the InlineLinkClicks field if non-nil, zero value otherwise.

### GetInlineLinkClicksOk

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetInlineLinkClicksOk() (*int32, bool)`

GetInlineLinkClicksOk returns a tuple with the InlineLinkClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineLinkClicks

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetInlineLinkClicks(v int32)`

SetInlineLinkClicks sets InlineLinkClicks field to given value.

### HasInlineLinkClicks

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) HasInlineLinkClicks() bool`

HasInlineLinkClicks returns a boolean if a field has been set.

### SetInlineLinkClicksNil

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetInlineLinkClicksNil(b bool)`

 SetInlineLinkClicksNil sets the value for InlineLinkClicks to be an explicit nil

### UnsetInlineLinkClicks
`func (o *AdTestCreateReadTestVerdictDailyFactsInner) UnsetInlineLinkClicks()`

UnsetInlineLinkClicks ensures that no value is present for InlineLinkClicks, not even an explicit nil
### GetCountedArrivals

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetCountedArrivals() int32`

GetCountedArrivals returns the CountedArrivals field if non-nil, zero value otherwise.

### GetCountedArrivalsOk

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetCountedArrivalsOk() (*int32, bool)`

GetCountedArrivalsOk returns a tuple with the CountedArrivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountedArrivals

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetCountedArrivals(v int32)`

SetCountedArrivals sets CountedArrivals field to given value.


### SetCountedArrivalsNil

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetCountedArrivalsNil(b bool)`

 SetCountedArrivalsNil sets the value for CountedArrivals to be an explicit nil

### UnsetCountedArrivals
`func (o *AdTestCreateReadTestVerdictDailyFactsInner) UnsetCountedArrivals()`

UnsetCountedArrivals ensures that no value is present for CountedArrivals, not even an explicit nil
### GetRecordedConversions

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetRecordedConversions() int32`

GetRecordedConversions returns the RecordedConversions field if non-nil, zero value otherwise.

### GetRecordedConversionsOk

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetRecordedConversionsOk() (*int32, bool)`

GetRecordedConversionsOk returns a tuple with the RecordedConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedConversions

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetRecordedConversions(v int32)`

SetRecordedConversions sets RecordedConversions field to given value.


### SetRecordedConversionsNil

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetRecordedConversionsNil(b bool)`

 SetRecordedConversionsNil sets the value for RecordedConversions to be an explicit nil

### UnsetRecordedConversions
`func (o *AdTestCreateReadTestVerdictDailyFactsInner) UnsetRecordedConversions()`

UnsetRecordedConversions ensures that no value is present for RecordedConversions, not even an explicit nil
### GetAdmittedConversions

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetAdmittedConversions() int32`

GetAdmittedConversions returns the AdmittedConversions field if non-nil, zero value otherwise.

### GetAdmittedConversionsOk

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) GetAdmittedConversionsOk() (*int32, bool)`

GetAdmittedConversionsOk returns a tuple with the AdmittedConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmittedConversions

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetAdmittedConversions(v int32)`

SetAdmittedConversions sets AdmittedConversions field to given value.


### SetAdmittedConversionsNil

`func (o *AdTestCreateReadTestVerdictDailyFactsInner) SetAdmittedConversionsNil(b bool)`

 SetAdmittedConversionsNil sets the value for AdmittedConversions to be an explicit nil

### UnsetAdmittedConversions
`func (o *AdTestCreateReadTestVerdictDailyFactsInner) UnsetAdmittedConversions()`

UnsetAdmittedConversions ensures that no value is present for AdmittedConversions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


