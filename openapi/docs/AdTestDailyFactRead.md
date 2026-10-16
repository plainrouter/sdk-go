# AdTestDailyFactRead

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

### NewAdTestDailyFactRead

`func NewAdTestDailyFactRead(date string, adTestMemberId string, adId string, spendMinor string, spend NullableString, impressions NullableInt32, countedArrivals NullableInt32, recordedConversions NullableInt32, admittedConversions NullableInt32, ) *AdTestDailyFactRead`

NewAdTestDailyFactRead instantiates a new AdTestDailyFactRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdTestDailyFactReadWithDefaults

`func NewAdTestDailyFactReadWithDefaults() *AdTestDailyFactRead`

NewAdTestDailyFactReadWithDefaults instantiates a new AdTestDailyFactRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *AdTestDailyFactRead) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AdTestDailyFactRead) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AdTestDailyFactRead) SetDate(v string)`

SetDate sets Date field to given value.


### GetAdTestMemberId

`func (o *AdTestDailyFactRead) GetAdTestMemberId() string`

GetAdTestMemberId returns the AdTestMemberId field if non-nil, zero value otherwise.

### GetAdTestMemberIdOk

`func (o *AdTestDailyFactRead) GetAdTestMemberIdOk() (*string, bool)`

GetAdTestMemberIdOk returns a tuple with the AdTestMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdTestMemberId

`func (o *AdTestDailyFactRead) SetAdTestMemberId(v string)`

SetAdTestMemberId sets AdTestMemberId field to given value.


### GetAdId

`func (o *AdTestDailyFactRead) GetAdId() string`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *AdTestDailyFactRead) GetAdIdOk() (*string, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *AdTestDailyFactRead) SetAdId(v string)`

SetAdId sets AdId field to given value.


### GetSpendMinor

`func (o *AdTestDailyFactRead) GetSpendMinor() string`

GetSpendMinor returns the SpendMinor field if non-nil, zero value otherwise.

### GetSpendMinorOk

`func (o *AdTestDailyFactRead) GetSpendMinorOk() (*string, bool)`

GetSpendMinorOk returns a tuple with the SpendMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendMinor

`func (o *AdTestDailyFactRead) SetSpendMinor(v string)`

SetSpendMinor sets SpendMinor field to given value.


### GetSpend

`func (o *AdTestDailyFactRead) GetSpend() string`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *AdTestDailyFactRead) GetSpendOk() (*string, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *AdTestDailyFactRead) SetSpend(v string)`

SetSpend sets Spend field to given value.


### SetSpendNil

`func (o *AdTestDailyFactRead) SetSpendNil(b bool)`

 SetSpendNil sets the value for Spend to be an explicit nil

### UnsetSpend
`func (o *AdTestDailyFactRead) UnsetSpend()`

UnsetSpend ensures that no value is present for Spend, not even an explicit nil
### GetImpressions

`func (o *AdTestDailyFactRead) GetImpressions() int32`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *AdTestDailyFactRead) GetImpressionsOk() (*int32, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *AdTestDailyFactRead) SetImpressions(v int32)`

SetImpressions sets Impressions field to given value.


### SetImpressionsNil

`func (o *AdTestDailyFactRead) SetImpressionsNil(b bool)`

 SetImpressionsNil sets the value for Impressions to be an explicit nil

### UnsetImpressions
`func (o *AdTestDailyFactRead) UnsetImpressions()`

UnsetImpressions ensures that no value is present for Impressions, not even an explicit nil
### GetInlineLinkClicks

`func (o *AdTestDailyFactRead) GetInlineLinkClicks() int32`

GetInlineLinkClicks returns the InlineLinkClicks field if non-nil, zero value otherwise.

### GetInlineLinkClicksOk

`func (o *AdTestDailyFactRead) GetInlineLinkClicksOk() (*int32, bool)`

GetInlineLinkClicksOk returns a tuple with the InlineLinkClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineLinkClicks

`func (o *AdTestDailyFactRead) SetInlineLinkClicks(v int32)`

SetInlineLinkClicks sets InlineLinkClicks field to given value.

### HasInlineLinkClicks

`func (o *AdTestDailyFactRead) HasInlineLinkClicks() bool`

HasInlineLinkClicks returns a boolean if a field has been set.

### SetInlineLinkClicksNil

`func (o *AdTestDailyFactRead) SetInlineLinkClicksNil(b bool)`

 SetInlineLinkClicksNil sets the value for InlineLinkClicks to be an explicit nil

### UnsetInlineLinkClicks
`func (o *AdTestDailyFactRead) UnsetInlineLinkClicks()`

UnsetInlineLinkClicks ensures that no value is present for InlineLinkClicks, not even an explicit nil
### GetCountedArrivals

`func (o *AdTestDailyFactRead) GetCountedArrivals() int32`

GetCountedArrivals returns the CountedArrivals field if non-nil, zero value otherwise.

### GetCountedArrivalsOk

`func (o *AdTestDailyFactRead) GetCountedArrivalsOk() (*int32, bool)`

GetCountedArrivalsOk returns a tuple with the CountedArrivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountedArrivals

`func (o *AdTestDailyFactRead) SetCountedArrivals(v int32)`

SetCountedArrivals sets CountedArrivals field to given value.


### SetCountedArrivalsNil

`func (o *AdTestDailyFactRead) SetCountedArrivalsNil(b bool)`

 SetCountedArrivalsNil sets the value for CountedArrivals to be an explicit nil

### UnsetCountedArrivals
`func (o *AdTestDailyFactRead) UnsetCountedArrivals()`

UnsetCountedArrivals ensures that no value is present for CountedArrivals, not even an explicit nil
### GetRecordedConversions

`func (o *AdTestDailyFactRead) GetRecordedConversions() int32`

GetRecordedConversions returns the RecordedConversions field if non-nil, zero value otherwise.

### GetRecordedConversionsOk

`func (o *AdTestDailyFactRead) GetRecordedConversionsOk() (*int32, bool)`

GetRecordedConversionsOk returns a tuple with the RecordedConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedConversions

`func (o *AdTestDailyFactRead) SetRecordedConversions(v int32)`

SetRecordedConversions sets RecordedConversions field to given value.


### SetRecordedConversionsNil

`func (o *AdTestDailyFactRead) SetRecordedConversionsNil(b bool)`

 SetRecordedConversionsNil sets the value for RecordedConversions to be an explicit nil

### UnsetRecordedConversions
`func (o *AdTestDailyFactRead) UnsetRecordedConversions()`

UnsetRecordedConversions ensures that no value is present for RecordedConversions, not even an explicit nil
### GetAdmittedConversions

`func (o *AdTestDailyFactRead) GetAdmittedConversions() int32`

GetAdmittedConversions returns the AdmittedConversions field if non-nil, zero value otherwise.

### GetAdmittedConversionsOk

`func (o *AdTestDailyFactRead) GetAdmittedConversionsOk() (*int32, bool)`

GetAdmittedConversionsOk returns a tuple with the AdmittedConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmittedConversions

`func (o *AdTestDailyFactRead) SetAdmittedConversions(v int32)`

SetAdmittedConversions sets AdmittedConversions field to given value.


### SetAdmittedConversionsNil

`func (o *AdTestDailyFactRead) SetAdmittedConversionsNil(b bool)`

 SetAdmittedConversionsNil sets the value for AdmittedConversions to be an explicit nil

### UnsetAdmittedConversions
`func (o *AdTestDailyFactRead) UnsetAdmittedConversions()`

UnsetAdmittedConversions ensures that no value is present for AdmittedConversions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


