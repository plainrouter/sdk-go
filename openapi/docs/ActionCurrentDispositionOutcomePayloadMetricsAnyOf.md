# ActionCurrentDispositionOutcomePayloadMetricsAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to **string** |  | [optional] 
**Days** | Pointer to **[]string** | Three whole days in the ad account timezone; excludes the change day. | [optional] 
**Currency** | Pointer to **NullableString** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**SpendMinor** | Pointer to **NullableString** | Decimal spend in ISO currency minor units; null when daily mirror data is missing. | [optional] 
**CountedArrivals** | Pointer to **NullableInt32** | Arrivals counted by PlainRouter through Signals, never provider conversions. | [optional] 
**CostPerCountedArrivalMinor** | Pointer to **NullableString** | Spend divided by counted arrivals, in ISO currency minor units, displayed to six decimal places. | [optional] 
**Verdict** | Pointer to **NullableString** |  | [optional] 
**LossThresholdPercent** | Pointer to **string** |  | [optional] 

## Methods

### NewActionCurrentDispositionOutcomePayloadMetricsAnyOf

`func NewActionCurrentDispositionOutcomePayloadMetricsAnyOf() *ActionCurrentDispositionOutcomePayloadMetricsAnyOf`

NewActionCurrentDispositionOutcomePayloadMetricsAnyOf instantiates a new ActionCurrentDispositionOutcomePayloadMetricsAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionCurrentDispositionOutcomePayloadMetricsAnyOfWithDefaults

`func NewActionCurrentDispositionOutcomePayloadMetricsAnyOfWithDefaults() *ActionCurrentDispositionOutcomePayloadMetricsAnyOf`

NewActionCurrentDispositionOutcomePayloadMetricsAnyOfWithDefaults instantiates a new ActionCurrentDispositionOutcomePayloadMetricsAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDays

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) SetDays(v []string)`

SetDays sets Days field to given value.

### HasDays

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetCurrency

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetTimezone

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSpendMinor

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetSpendMinor() string`

GetSpendMinor returns the SpendMinor field if non-nil, zero value otherwise.

### GetSpendMinorOk

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetSpendMinorOk() (*string, bool)`

GetSpendMinorOk returns a tuple with the SpendMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendMinor

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) SetSpendMinor(v string)`

SetSpendMinor sets SpendMinor field to given value.

### HasSpendMinor

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) HasSpendMinor() bool`

HasSpendMinor returns a boolean if a field has been set.

### SetSpendMinorNil

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) SetSpendMinorNil(b bool)`

 SetSpendMinorNil sets the value for SpendMinor to be an explicit nil

### UnsetSpendMinor
`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) UnsetSpendMinor()`

UnsetSpendMinor ensures that no value is present for SpendMinor, not even an explicit nil
### GetCountedArrivals

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetCountedArrivals() int32`

GetCountedArrivals returns the CountedArrivals field if non-nil, zero value otherwise.

### GetCountedArrivalsOk

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetCountedArrivalsOk() (*int32, bool)`

GetCountedArrivalsOk returns a tuple with the CountedArrivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountedArrivals

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) SetCountedArrivals(v int32)`

SetCountedArrivals sets CountedArrivals field to given value.

### HasCountedArrivals

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) HasCountedArrivals() bool`

HasCountedArrivals returns a boolean if a field has been set.

### SetCountedArrivalsNil

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) SetCountedArrivalsNil(b bool)`

 SetCountedArrivalsNil sets the value for CountedArrivals to be an explicit nil

### UnsetCountedArrivals
`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) UnsetCountedArrivals()`

UnsetCountedArrivals ensures that no value is present for CountedArrivals, not even an explicit nil
### GetCostPerCountedArrivalMinor

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetCostPerCountedArrivalMinor() string`

GetCostPerCountedArrivalMinor returns the CostPerCountedArrivalMinor field if non-nil, zero value otherwise.

### GetCostPerCountedArrivalMinorOk

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetCostPerCountedArrivalMinorOk() (*string, bool)`

GetCostPerCountedArrivalMinorOk returns a tuple with the CostPerCountedArrivalMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerCountedArrivalMinor

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) SetCostPerCountedArrivalMinor(v string)`

SetCostPerCountedArrivalMinor sets CostPerCountedArrivalMinor field to given value.

### HasCostPerCountedArrivalMinor

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) HasCostPerCountedArrivalMinor() bool`

HasCostPerCountedArrivalMinor returns a boolean if a field has been set.

### SetCostPerCountedArrivalMinorNil

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) SetCostPerCountedArrivalMinorNil(b bool)`

 SetCostPerCountedArrivalMinorNil sets the value for CostPerCountedArrivalMinor to be an explicit nil

### UnsetCostPerCountedArrivalMinor
`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) UnsetCostPerCountedArrivalMinor()`

UnsetCostPerCountedArrivalMinor ensures that no value is present for CostPerCountedArrivalMinor, not even an explicit nil
### GetVerdict

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.

### SetVerdictNil

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) SetVerdictNil(b bool)`

 SetVerdictNil sets the value for Verdict to be an explicit nil

### UnsetVerdict
`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) UnsetVerdict()`

UnsetVerdict ensures that no value is present for Verdict, not even an explicit nil
### GetLossThresholdPercent

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetLossThresholdPercent() string`

GetLossThresholdPercent returns the LossThresholdPercent field if non-nil, zero value otherwise.

### GetLossThresholdPercentOk

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) GetLossThresholdPercentOk() (*string, bool)`

GetLossThresholdPercentOk returns a tuple with the LossThresholdPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossThresholdPercent

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) SetLossThresholdPercent(v string)`

SetLossThresholdPercent sets LossThresholdPercent field to given value.

### HasLossThresholdPercent

`func (o *ActionCurrentDispositionOutcomePayloadMetricsAnyOf) HasLossThresholdPercent() bool`

HasLossThresholdPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


