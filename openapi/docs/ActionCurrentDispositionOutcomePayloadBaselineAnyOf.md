# ActionCurrentDispositionOutcomePayloadBaselineAnyOf

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

## Methods

### NewActionCurrentDispositionOutcomePayloadBaselineAnyOf

`func NewActionCurrentDispositionOutcomePayloadBaselineAnyOf() *ActionCurrentDispositionOutcomePayloadBaselineAnyOf`

NewActionCurrentDispositionOutcomePayloadBaselineAnyOf instantiates a new ActionCurrentDispositionOutcomePayloadBaselineAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionCurrentDispositionOutcomePayloadBaselineAnyOfWithDefaults

`func NewActionCurrentDispositionOutcomePayloadBaselineAnyOfWithDefaults() *ActionCurrentDispositionOutcomePayloadBaselineAnyOf`

NewActionCurrentDispositionOutcomePayloadBaselineAnyOfWithDefaults instantiates a new ActionCurrentDispositionOutcomePayloadBaselineAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDays

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) SetDays(v []string)`

SetDays sets Days field to given value.

### HasDays

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetCurrency

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetTimezone

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSpendMinor

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) GetSpendMinor() string`

GetSpendMinor returns the SpendMinor field if non-nil, zero value otherwise.

### GetSpendMinorOk

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) GetSpendMinorOk() (*string, bool)`

GetSpendMinorOk returns a tuple with the SpendMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendMinor

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) SetSpendMinor(v string)`

SetSpendMinor sets SpendMinor field to given value.

### HasSpendMinor

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) HasSpendMinor() bool`

HasSpendMinor returns a boolean if a field has been set.

### SetSpendMinorNil

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) SetSpendMinorNil(b bool)`

 SetSpendMinorNil sets the value for SpendMinor to be an explicit nil

### UnsetSpendMinor
`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) UnsetSpendMinor()`

UnsetSpendMinor ensures that no value is present for SpendMinor, not even an explicit nil
### GetCountedArrivals

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) GetCountedArrivals() int32`

GetCountedArrivals returns the CountedArrivals field if non-nil, zero value otherwise.

### GetCountedArrivalsOk

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) GetCountedArrivalsOk() (*int32, bool)`

GetCountedArrivalsOk returns a tuple with the CountedArrivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountedArrivals

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) SetCountedArrivals(v int32)`

SetCountedArrivals sets CountedArrivals field to given value.

### HasCountedArrivals

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) HasCountedArrivals() bool`

HasCountedArrivals returns a boolean if a field has been set.

### SetCountedArrivalsNil

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) SetCountedArrivalsNil(b bool)`

 SetCountedArrivalsNil sets the value for CountedArrivals to be an explicit nil

### UnsetCountedArrivals
`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) UnsetCountedArrivals()`

UnsetCountedArrivals ensures that no value is present for CountedArrivals, not even an explicit nil
### GetCostPerCountedArrivalMinor

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) GetCostPerCountedArrivalMinor() string`

GetCostPerCountedArrivalMinor returns the CostPerCountedArrivalMinor field if non-nil, zero value otherwise.

### GetCostPerCountedArrivalMinorOk

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) GetCostPerCountedArrivalMinorOk() (*string, bool)`

GetCostPerCountedArrivalMinorOk returns a tuple with the CostPerCountedArrivalMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerCountedArrivalMinor

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) SetCostPerCountedArrivalMinor(v string)`

SetCostPerCountedArrivalMinor sets CostPerCountedArrivalMinor field to given value.

### HasCostPerCountedArrivalMinor

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) HasCostPerCountedArrivalMinor() bool`

HasCostPerCountedArrivalMinor returns a boolean if a field has been set.

### SetCostPerCountedArrivalMinorNil

`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) SetCostPerCountedArrivalMinorNil(b bool)`

 SetCostPerCountedArrivalMinorNil sets the value for CostPerCountedArrivalMinor to be an explicit nil

### UnsetCostPerCountedArrivalMinor
`func (o *ActionCurrentDispositionOutcomePayloadBaselineAnyOf) UnsetCostPerCountedArrivalMinor()`

UnsetCostPerCountedArrivalMinor ensures that no value is present for CostPerCountedArrivalMinor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


