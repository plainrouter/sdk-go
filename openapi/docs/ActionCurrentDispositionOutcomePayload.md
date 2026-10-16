# ActionCurrentDispositionOutcomePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**ObservedAt** | Pointer to **time.Time** |  | [optional] 
**ReasonCode** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Baseline** | Pointer to [**ActionCurrentDispositionOutcomePayloadBaseline**](ActionCurrentDispositionOutcomePayloadBaseline.md) |  | [optional] 
**Metrics** | Pointer to [**ActionCurrentDispositionOutcomePayloadMetrics**](ActionCurrentDispositionOutcomePayloadMetrics.md) |  | [optional] 
**Inverse** | Pointer to [**ActionCurrentDispositionOutcomePayloadInverse**](ActionCurrentDispositionOutcomePayloadInverse.md) |  | [optional] 

## Methods

### NewActionCurrentDispositionOutcomePayload

`func NewActionCurrentDispositionOutcomePayload() *ActionCurrentDispositionOutcomePayload`

NewActionCurrentDispositionOutcomePayload instantiates a new ActionCurrentDispositionOutcomePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionCurrentDispositionOutcomePayloadWithDefaults

`func NewActionCurrentDispositionOutcomePayloadWithDefaults() *ActionCurrentDispositionOutcomePayload`

NewActionCurrentDispositionOutcomePayloadWithDefaults instantiates a new ActionCurrentDispositionOutcomePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ActionCurrentDispositionOutcomePayload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionCurrentDispositionOutcomePayload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionCurrentDispositionOutcomePayload) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActionCurrentDispositionOutcomePayload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetObservedAt

`func (o *ActionCurrentDispositionOutcomePayload) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *ActionCurrentDispositionOutcomePayload) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *ActionCurrentDispositionOutcomePayload) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.

### HasObservedAt

`func (o *ActionCurrentDispositionOutcomePayload) HasObservedAt() bool`

HasObservedAt returns a boolean if a field has been set.

### GetReasonCode

`func (o *ActionCurrentDispositionOutcomePayload) GetReasonCode() string`

GetReasonCode returns the ReasonCode field if non-nil, zero value otherwise.

### GetReasonCodeOk

`func (o *ActionCurrentDispositionOutcomePayload) GetReasonCodeOk() (*string, bool)`

GetReasonCodeOk returns a tuple with the ReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonCode

`func (o *ActionCurrentDispositionOutcomePayload) SetReasonCode(v string)`

SetReasonCode sets ReasonCode field to given value.

### HasReasonCode

`func (o *ActionCurrentDispositionOutcomePayload) HasReasonCode() bool`

HasReasonCode returns a boolean if a field has been set.

### SetReasonCodeNil

`func (o *ActionCurrentDispositionOutcomePayload) SetReasonCodeNil(b bool)`

 SetReasonCodeNil sets the value for ReasonCode to be an explicit nil

### UnsetReasonCode
`func (o *ActionCurrentDispositionOutcomePayload) UnsetReasonCode()`

UnsetReasonCode ensures that no value is present for ReasonCode, not even an explicit nil
### GetReason

`func (o *ActionCurrentDispositionOutcomePayload) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ActionCurrentDispositionOutcomePayload) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ActionCurrentDispositionOutcomePayload) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ActionCurrentDispositionOutcomePayload) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *ActionCurrentDispositionOutcomePayload) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ActionCurrentDispositionOutcomePayload) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetBaseline

`func (o *ActionCurrentDispositionOutcomePayload) GetBaseline() ActionCurrentDispositionOutcomePayloadBaseline`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *ActionCurrentDispositionOutcomePayload) GetBaselineOk() (*ActionCurrentDispositionOutcomePayloadBaseline, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *ActionCurrentDispositionOutcomePayload) SetBaseline(v ActionCurrentDispositionOutcomePayloadBaseline)`

SetBaseline sets Baseline field to given value.

### HasBaseline

`func (o *ActionCurrentDispositionOutcomePayload) HasBaseline() bool`

HasBaseline returns a boolean if a field has been set.

### GetMetrics

`func (o *ActionCurrentDispositionOutcomePayload) GetMetrics() ActionCurrentDispositionOutcomePayloadMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ActionCurrentDispositionOutcomePayload) GetMetricsOk() (*ActionCurrentDispositionOutcomePayloadMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ActionCurrentDispositionOutcomePayload) SetMetrics(v ActionCurrentDispositionOutcomePayloadMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ActionCurrentDispositionOutcomePayload) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetInverse

`func (o *ActionCurrentDispositionOutcomePayload) GetInverse() ActionCurrentDispositionOutcomePayloadInverse`

GetInverse returns the Inverse field if non-nil, zero value otherwise.

### GetInverseOk

`func (o *ActionCurrentDispositionOutcomePayload) GetInverseOk() (*ActionCurrentDispositionOutcomePayloadInverse, bool)`

GetInverseOk returns a tuple with the Inverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverse

`func (o *ActionCurrentDispositionOutcomePayload) SetInverse(v ActionCurrentDispositionOutcomePayloadInverse)`

SetInverse sets Inverse field to given value.

### HasInverse

`func (o *ActionCurrentDispositionOutcomePayload) HasInverse() bool`

HasInverse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


