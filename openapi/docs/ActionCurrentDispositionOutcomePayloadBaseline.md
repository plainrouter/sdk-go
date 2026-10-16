# ActionCurrentDispositionOutcomePayloadBaseline

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

## Methods

### NewActionCurrentDispositionOutcomePayloadBaseline

`func NewActionCurrentDispositionOutcomePayloadBaseline() *ActionCurrentDispositionOutcomePayloadBaseline`

NewActionCurrentDispositionOutcomePayloadBaseline instantiates a new ActionCurrentDispositionOutcomePayloadBaseline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionCurrentDispositionOutcomePayloadBaselineWithDefaults

`func NewActionCurrentDispositionOutcomePayloadBaselineWithDefaults() *ActionCurrentDispositionOutcomePayloadBaseline`

NewActionCurrentDispositionOutcomePayloadBaselineWithDefaults instantiates a new ActionCurrentDispositionOutcomePayloadBaseline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDays

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) GetDays() []string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) GetDaysOk() (*[]string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) SetDays(v []string)`

SetDays sets Days field to given value.

### HasDays

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetCurrency

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTimezone

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSpendMinor

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) GetSpendMinor() string`

GetSpendMinor returns the SpendMinor field if non-nil, zero value otherwise.

### GetSpendMinorOk

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) GetSpendMinorOk() (*string, bool)`

GetSpendMinorOk returns a tuple with the SpendMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendMinor

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) SetSpendMinor(v string)`

SetSpendMinor sets SpendMinor field to given value.

### HasSpendMinor

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) HasSpendMinor() bool`

HasSpendMinor returns a boolean if a field has been set.

### GetCountedArrivals

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) GetCountedArrivals() int32`

GetCountedArrivals returns the CountedArrivals field if non-nil, zero value otherwise.

### GetCountedArrivalsOk

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) GetCountedArrivalsOk() (*int32, bool)`

GetCountedArrivalsOk returns a tuple with the CountedArrivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountedArrivals

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) SetCountedArrivals(v int32)`

SetCountedArrivals sets CountedArrivals field to given value.

### HasCountedArrivals

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) HasCountedArrivals() bool`

HasCountedArrivals returns a boolean if a field has been set.

### GetCostPerCountedArrivalMinor

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) GetCostPerCountedArrivalMinor() string`

GetCostPerCountedArrivalMinor returns the CostPerCountedArrivalMinor field if non-nil, zero value otherwise.

### GetCostPerCountedArrivalMinorOk

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) GetCostPerCountedArrivalMinorOk() (*string, bool)`

GetCostPerCountedArrivalMinorOk returns a tuple with the CostPerCountedArrivalMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerCountedArrivalMinor

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) SetCostPerCountedArrivalMinor(v string)`

SetCostPerCountedArrivalMinor sets CostPerCountedArrivalMinor field to given value.

### HasCostPerCountedArrivalMinor

`func (o *ActionCurrentDispositionOutcomePayloadBaseline) HasCostPerCountedArrivalMinor() bool`

HasCostPerCountedArrivalMinor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


