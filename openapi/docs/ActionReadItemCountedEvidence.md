# ActionReadItemCountedEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | [**ActionReadItemCountedEvidenceTarget**](ActionReadItemCountedEvidenceTarget.md) |  | 
**CapturedAt** | **string** |  | 
**Window** | [**NullableActionReadItemCountedEvidenceWindow**](ActionReadItemCountedEvidenceWindow.md) |  | 
**Currency** | **NullableString** |  | 
**CurrencyNote** | **string** |  | 
**MetaReported** | [**ActionReadItemCountedEvidenceMetaReported**](ActionReadItemCountedEvidenceMetaReported.md) |  | 
**Counted** | [**ActionReadItemCountedEvidenceCounted**](ActionReadItemCountedEvidenceCounted.md) |  | 
**Coverage** | [**ActionReadItemCountedEvidenceCoverage**](ActionReadItemCountedEvidenceCoverage.md) |  | 

## Methods

### NewActionReadItemCountedEvidence

`func NewActionReadItemCountedEvidence(target ActionReadItemCountedEvidenceTarget, capturedAt string, window NullableActionReadItemCountedEvidenceWindow, currency NullableString, currencyNote string, metaReported ActionReadItemCountedEvidenceMetaReported, counted ActionReadItemCountedEvidenceCounted, coverage ActionReadItemCountedEvidenceCoverage, ) *ActionReadItemCountedEvidence`

NewActionReadItemCountedEvidence instantiates a new ActionReadItemCountedEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionReadItemCountedEvidenceWithDefaults

`func NewActionReadItemCountedEvidenceWithDefaults() *ActionReadItemCountedEvidence`

NewActionReadItemCountedEvidenceWithDefaults instantiates a new ActionReadItemCountedEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *ActionReadItemCountedEvidence) GetTarget() ActionReadItemCountedEvidenceTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ActionReadItemCountedEvidence) GetTargetOk() (*ActionReadItemCountedEvidenceTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ActionReadItemCountedEvidence) SetTarget(v ActionReadItemCountedEvidenceTarget)`

SetTarget sets Target field to given value.


### GetCapturedAt

`func (o *ActionReadItemCountedEvidence) GetCapturedAt() string`

GetCapturedAt returns the CapturedAt field if non-nil, zero value otherwise.

### GetCapturedAtOk

`func (o *ActionReadItemCountedEvidence) GetCapturedAtOk() (*string, bool)`

GetCapturedAtOk returns a tuple with the CapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedAt

`func (o *ActionReadItemCountedEvidence) SetCapturedAt(v string)`

SetCapturedAt sets CapturedAt field to given value.


### GetWindow

`func (o *ActionReadItemCountedEvidence) GetWindow() ActionReadItemCountedEvidenceWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ActionReadItemCountedEvidence) GetWindowOk() (*ActionReadItemCountedEvidenceWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ActionReadItemCountedEvidence) SetWindow(v ActionReadItemCountedEvidenceWindow)`

SetWindow sets Window field to given value.


### SetWindowNil

`func (o *ActionReadItemCountedEvidence) SetWindowNil(b bool)`

 SetWindowNil sets the value for Window to be an explicit nil

### UnsetWindow
`func (o *ActionReadItemCountedEvidence) UnsetWindow()`

UnsetWindow ensures that no value is present for Window, not even an explicit nil
### GetCurrency

`func (o *ActionReadItemCountedEvidence) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ActionReadItemCountedEvidence) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ActionReadItemCountedEvidence) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *ActionReadItemCountedEvidence) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ActionReadItemCountedEvidence) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCurrencyNote

`func (o *ActionReadItemCountedEvidence) GetCurrencyNote() string`

GetCurrencyNote returns the CurrencyNote field if non-nil, zero value otherwise.

### GetCurrencyNoteOk

`func (o *ActionReadItemCountedEvidence) GetCurrencyNoteOk() (*string, bool)`

GetCurrencyNoteOk returns a tuple with the CurrencyNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyNote

`func (o *ActionReadItemCountedEvidence) SetCurrencyNote(v string)`

SetCurrencyNote sets CurrencyNote field to given value.


### GetMetaReported

`func (o *ActionReadItemCountedEvidence) GetMetaReported() ActionReadItemCountedEvidenceMetaReported`

GetMetaReported returns the MetaReported field if non-nil, zero value otherwise.

### GetMetaReportedOk

`func (o *ActionReadItemCountedEvidence) GetMetaReportedOk() (*ActionReadItemCountedEvidenceMetaReported, bool)`

GetMetaReportedOk returns a tuple with the MetaReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaReported

`func (o *ActionReadItemCountedEvidence) SetMetaReported(v ActionReadItemCountedEvidenceMetaReported)`

SetMetaReported sets MetaReported field to given value.


### GetCounted

`func (o *ActionReadItemCountedEvidence) GetCounted() ActionReadItemCountedEvidenceCounted`

GetCounted returns the Counted field if non-nil, zero value otherwise.

### GetCountedOk

`func (o *ActionReadItemCountedEvidence) GetCountedOk() (*ActionReadItemCountedEvidenceCounted, bool)`

GetCountedOk returns a tuple with the Counted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounted

`func (o *ActionReadItemCountedEvidence) SetCounted(v ActionReadItemCountedEvidenceCounted)`

SetCounted sets Counted field to given value.


### GetCoverage

`func (o *ActionReadItemCountedEvidence) GetCoverage() ActionReadItemCountedEvidenceCoverage`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *ActionReadItemCountedEvidence) GetCoverageOk() (*ActionReadItemCountedEvidenceCoverage, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *ActionReadItemCountedEvidence) SetCoverage(v ActionReadItemCountedEvidenceCoverage)`

SetCoverage sets Coverage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


