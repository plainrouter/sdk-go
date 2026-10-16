# AdTestCreateReadTestVerdictMembersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdId** | **string** |  | 
**State** | **string** |  | 
**Rank** | **NullableInt32** |  | 
**Spend** | **NullableString** |  | 
**Impressions** | **NullableInt32** |  | 
**CountedArrivals** | **NullableInt32** |  | 
**RecordedConversions** | **NullableInt32** | PlainRouter recorded consented conversions of this Test event. | 
**AdmittedConversions** | **NullableInt32** | Admitted conversions are shown beside recorded conversions and are never substituted for them. | 
**SpendPerRecordedConversion** | **NullableString** | Major-unit spend divided by PlainRouter recorded conversions; used only to compare members within this Test. | 
**CoveredDays** | **[]string** |  | 

## Methods

### NewAdTestCreateReadTestVerdictMembersInner

`func NewAdTestCreateReadTestVerdictMembersInner(adId string, state string, rank NullableInt32, spend NullableString, impressions NullableInt32, countedArrivals NullableInt32, recordedConversions NullableInt32, admittedConversions NullableInt32, spendPerRecordedConversion NullableString, coveredDays []string, ) *AdTestCreateReadTestVerdictMembersInner`

NewAdTestCreateReadTestVerdictMembersInner instantiates a new AdTestCreateReadTestVerdictMembersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdTestCreateReadTestVerdictMembersInnerWithDefaults

`func NewAdTestCreateReadTestVerdictMembersInnerWithDefaults() *AdTestCreateReadTestVerdictMembersInner`

NewAdTestCreateReadTestVerdictMembersInnerWithDefaults instantiates a new AdTestCreateReadTestVerdictMembersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdId

`func (o *AdTestCreateReadTestVerdictMembersInner) GetAdId() string`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *AdTestCreateReadTestVerdictMembersInner) GetAdIdOk() (*string, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *AdTestCreateReadTestVerdictMembersInner) SetAdId(v string)`

SetAdId sets AdId field to given value.


### GetState

`func (o *AdTestCreateReadTestVerdictMembersInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AdTestCreateReadTestVerdictMembersInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AdTestCreateReadTestVerdictMembersInner) SetState(v string)`

SetState sets State field to given value.


### GetRank

`func (o *AdTestCreateReadTestVerdictMembersInner) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *AdTestCreateReadTestVerdictMembersInner) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *AdTestCreateReadTestVerdictMembersInner) SetRank(v int32)`

SetRank sets Rank field to given value.


### SetRankNil

`func (o *AdTestCreateReadTestVerdictMembersInner) SetRankNil(b bool)`

 SetRankNil sets the value for Rank to be an explicit nil

### UnsetRank
`func (o *AdTestCreateReadTestVerdictMembersInner) UnsetRank()`

UnsetRank ensures that no value is present for Rank, not even an explicit nil
### GetSpend

`func (o *AdTestCreateReadTestVerdictMembersInner) GetSpend() string`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *AdTestCreateReadTestVerdictMembersInner) GetSpendOk() (*string, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *AdTestCreateReadTestVerdictMembersInner) SetSpend(v string)`

SetSpend sets Spend field to given value.


### SetSpendNil

`func (o *AdTestCreateReadTestVerdictMembersInner) SetSpendNil(b bool)`

 SetSpendNil sets the value for Spend to be an explicit nil

### UnsetSpend
`func (o *AdTestCreateReadTestVerdictMembersInner) UnsetSpend()`

UnsetSpend ensures that no value is present for Spend, not even an explicit nil
### GetImpressions

`func (o *AdTestCreateReadTestVerdictMembersInner) GetImpressions() int32`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *AdTestCreateReadTestVerdictMembersInner) GetImpressionsOk() (*int32, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *AdTestCreateReadTestVerdictMembersInner) SetImpressions(v int32)`

SetImpressions sets Impressions field to given value.


### SetImpressionsNil

`func (o *AdTestCreateReadTestVerdictMembersInner) SetImpressionsNil(b bool)`

 SetImpressionsNil sets the value for Impressions to be an explicit nil

### UnsetImpressions
`func (o *AdTestCreateReadTestVerdictMembersInner) UnsetImpressions()`

UnsetImpressions ensures that no value is present for Impressions, not even an explicit nil
### GetCountedArrivals

`func (o *AdTestCreateReadTestVerdictMembersInner) GetCountedArrivals() int32`

GetCountedArrivals returns the CountedArrivals field if non-nil, zero value otherwise.

### GetCountedArrivalsOk

`func (o *AdTestCreateReadTestVerdictMembersInner) GetCountedArrivalsOk() (*int32, bool)`

GetCountedArrivalsOk returns a tuple with the CountedArrivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountedArrivals

`func (o *AdTestCreateReadTestVerdictMembersInner) SetCountedArrivals(v int32)`

SetCountedArrivals sets CountedArrivals field to given value.


### SetCountedArrivalsNil

`func (o *AdTestCreateReadTestVerdictMembersInner) SetCountedArrivalsNil(b bool)`

 SetCountedArrivalsNil sets the value for CountedArrivals to be an explicit nil

### UnsetCountedArrivals
`func (o *AdTestCreateReadTestVerdictMembersInner) UnsetCountedArrivals()`

UnsetCountedArrivals ensures that no value is present for CountedArrivals, not even an explicit nil
### GetRecordedConversions

`func (o *AdTestCreateReadTestVerdictMembersInner) GetRecordedConversions() int32`

GetRecordedConversions returns the RecordedConversions field if non-nil, zero value otherwise.

### GetRecordedConversionsOk

`func (o *AdTestCreateReadTestVerdictMembersInner) GetRecordedConversionsOk() (*int32, bool)`

GetRecordedConversionsOk returns a tuple with the RecordedConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedConversions

`func (o *AdTestCreateReadTestVerdictMembersInner) SetRecordedConversions(v int32)`

SetRecordedConversions sets RecordedConversions field to given value.


### SetRecordedConversionsNil

`func (o *AdTestCreateReadTestVerdictMembersInner) SetRecordedConversionsNil(b bool)`

 SetRecordedConversionsNil sets the value for RecordedConversions to be an explicit nil

### UnsetRecordedConversions
`func (o *AdTestCreateReadTestVerdictMembersInner) UnsetRecordedConversions()`

UnsetRecordedConversions ensures that no value is present for RecordedConversions, not even an explicit nil
### GetAdmittedConversions

`func (o *AdTestCreateReadTestVerdictMembersInner) GetAdmittedConversions() int32`

GetAdmittedConversions returns the AdmittedConversions field if non-nil, zero value otherwise.

### GetAdmittedConversionsOk

`func (o *AdTestCreateReadTestVerdictMembersInner) GetAdmittedConversionsOk() (*int32, bool)`

GetAdmittedConversionsOk returns a tuple with the AdmittedConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmittedConversions

`func (o *AdTestCreateReadTestVerdictMembersInner) SetAdmittedConversions(v int32)`

SetAdmittedConversions sets AdmittedConversions field to given value.


### SetAdmittedConversionsNil

`func (o *AdTestCreateReadTestVerdictMembersInner) SetAdmittedConversionsNil(b bool)`

 SetAdmittedConversionsNil sets the value for AdmittedConversions to be an explicit nil

### UnsetAdmittedConversions
`func (o *AdTestCreateReadTestVerdictMembersInner) UnsetAdmittedConversions()`

UnsetAdmittedConversions ensures that no value is present for AdmittedConversions, not even an explicit nil
### GetSpendPerRecordedConversion

`func (o *AdTestCreateReadTestVerdictMembersInner) GetSpendPerRecordedConversion() string`

GetSpendPerRecordedConversion returns the SpendPerRecordedConversion field if non-nil, zero value otherwise.

### GetSpendPerRecordedConversionOk

`func (o *AdTestCreateReadTestVerdictMembersInner) GetSpendPerRecordedConversionOk() (*string, bool)`

GetSpendPerRecordedConversionOk returns a tuple with the SpendPerRecordedConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendPerRecordedConversion

`func (o *AdTestCreateReadTestVerdictMembersInner) SetSpendPerRecordedConversion(v string)`

SetSpendPerRecordedConversion sets SpendPerRecordedConversion field to given value.


### SetSpendPerRecordedConversionNil

`func (o *AdTestCreateReadTestVerdictMembersInner) SetSpendPerRecordedConversionNil(b bool)`

 SetSpendPerRecordedConversionNil sets the value for SpendPerRecordedConversion to be an explicit nil

### UnsetSpendPerRecordedConversion
`func (o *AdTestCreateReadTestVerdictMembersInner) UnsetSpendPerRecordedConversion()`

UnsetSpendPerRecordedConversion ensures that no value is present for SpendPerRecordedConversion, not even an explicit nil
### GetCoveredDays

`func (o *AdTestCreateReadTestVerdictMembersInner) GetCoveredDays() []string`

GetCoveredDays returns the CoveredDays field if non-nil, zero value otherwise.

### GetCoveredDaysOk

`func (o *AdTestCreateReadTestVerdictMembersInner) GetCoveredDaysOk() (*[]string, bool)`

GetCoveredDaysOk returns a tuple with the CoveredDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoveredDays

`func (o *AdTestCreateReadTestVerdictMembersInner) SetCoveredDays(v []string)`

SetCoveredDays sets CoveredDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


