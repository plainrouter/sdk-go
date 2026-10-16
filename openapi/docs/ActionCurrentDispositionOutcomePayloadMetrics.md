# ActionCurrentDispositionOutcomePayloadMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to **string** |  | [optional] 
**Days** | Pointer to **[]string** | Three whole days in the ad account timezone; excludes the change day. | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**SpendMinor** | Pointer to **string** | Decimal spend in ISO currency minor units; null when daily mirror data is missing. | [optional] 
**CountedArrivals** | Pointer to **int32** | Arrivals counted by PlainRouter through Signals, never provider conversions. | [optional] 
**CostPerCountedArrivalMinor** | Pointer to **string** | Spend divided by counted arrivals, in ISO currency minor units, displayed to six decimal places. | [optional] 
**Verdict** | Pointer to **NullableString** |  | [optional] 
**LossThresholdPercent** | Pointer to **string** |  | [optional] 

## Methods

### NewActionCurrentDispositionOutcomePayloadMetrics

`func NewActionCurrentDispositionOutcomePayloadMetrics() *ActionCurrentDispositionOutcomePayloadMetrics`

NewActionCurrentDispositionOutcomePayloadMetrics instantiates a new ActionCurrentDispositionOutcomePayloadMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionCurrentDispositionOutcomePayloadMetricsWithDefaults

`func NewActionCurrentDispositionOutcomePayloadMetricsWithDefaults() *ActionCurrentDispositionOutcomePayloadMetrics`

NewActionCurrentDispositionOutcomePayloadMetricsWithDefaults instantiates a new ActionCurrentDispositionOutcomePayloadMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDays

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) SetDays(v []string)`

SetDays sets Days field to given value.

### HasDays

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetCurrency

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTimezone

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSpendMinor

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetSpendMinor() string`

GetSpendMinor returns the SpendMinor field if non-nil, zero value otherwise.

### GetSpendMinorOk

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetSpendMinorOk() (*string, bool)`

GetSpendMinorOk returns a tuple with the SpendMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendMinor

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) SetSpendMinor(v string)`

SetSpendMinor sets SpendMinor field to given value.

### HasSpendMinor

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) HasSpendMinor() bool`

HasSpendMinor returns a boolean if a field has been set.

### GetCountedArrivals

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetCountedArrivals() int32`

GetCountedArrivals returns the CountedArrivals field if non-nil, zero value otherwise.

### GetCountedArrivalsOk

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetCountedArrivalsOk() (*int32, bool)`

GetCountedArrivalsOk returns a tuple with the CountedArrivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountedArrivals

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) SetCountedArrivals(v int32)`

SetCountedArrivals sets CountedArrivals field to given value.

### HasCountedArrivals

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) HasCountedArrivals() bool`

HasCountedArrivals returns a boolean if a field has been set.

### GetCostPerCountedArrivalMinor

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetCostPerCountedArrivalMinor() string`

GetCostPerCountedArrivalMinor returns the CostPerCountedArrivalMinor field if non-nil, zero value otherwise.

### GetCostPerCountedArrivalMinorOk

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetCostPerCountedArrivalMinorOk() (*string, bool)`

GetCostPerCountedArrivalMinorOk returns a tuple with the CostPerCountedArrivalMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerCountedArrivalMinor

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) SetCostPerCountedArrivalMinor(v string)`

SetCostPerCountedArrivalMinor sets CostPerCountedArrivalMinor field to given value.

### HasCostPerCountedArrivalMinor

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) HasCostPerCountedArrivalMinor() bool`

HasCostPerCountedArrivalMinor returns a boolean if a field has been set.

### GetVerdict

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.

### SetVerdictNil

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) SetVerdictNil(b bool)`

 SetVerdictNil sets the value for Verdict to be an explicit nil

### UnsetVerdict
`func (o *ActionCurrentDispositionOutcomePayloadMetrics) UnsetVerdict()`

UnsetVerdict ensures that no value is present for Verdict, not even an explicit nil
### GetLossThresholdPercent

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetLossThresholdPercent() string`

GetLossThresholdPercent returns the LossThresholdPercent field if non-nil, zero value otherwise.

### GetLossThresholdPercentOk

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) GetLossThresholdPercentOk() (*string, bool)`

GetLossThresholdPercentOk returns a tuple with the LossThresholdPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossThresholdPercent

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) SetLossThresholdPercent(v string)`

SetLossThresholdPercent sets LossThresholdPercent field to given value.

### HasLossThresholdPercent

`func (o *ActionCurrentDispositionOutcomePayloadMetrics) HasLossThresholdPercent() bool`

HasLossThresholdPercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


