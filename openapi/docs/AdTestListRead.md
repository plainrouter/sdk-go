# AdTestListRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tests** | [**AdTestListReadTests**](AdTestListReadTests.md) |  | 
**ObservedConversionEventsAvailable** | **bool** |  | 
**ObservedConversionEvents** | [**[]AdTestListReadObservedConversionEventsInner**](AdTestListReadObservedConversionEventsInner.md) |  | 

## Methods

### NewAdTestListRead

`func NewAdTestListRead(tests AdTestListReadTests, observedConversionEventsAvailable bool, observedConversionEvents []AdTestListReadObservedConversionEventsInner, ) *AdTestListRead`

NewAdTestListRead instantiates a new AdTestListRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdTestListReadWithDefaults

`func NewAdTestListReadWithDefaults() *AdTestListRead`

NewAdTestListReadWithDefaults instantiates a new AdTestListRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTests

`func (o *AdTestListRead) GetTests() AdTestListReadTests`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *AdTestListRead) GetTestsOk() (*AdTestListReadTests, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *AdTestListRead) SetTests(v AdTestListReadTests)`

SetTests sets Tests field to given value.


### GetObservedConversionEventsAvailable

`func (o *AdTestListRead) GetObservedConversionEventsAvailable() bool`

GetObservedConversionEventsAvailable returns the ObservedConversionEventsAvailable field if non-nil, zero value otherwise.

### GetObservedConversionEventsAvailableOk

`func (o *AdTestListRead) GetObservedConversionEventsAvailableOk() (*bool, bool)`

GetObservedConversionEventsAvailableOk returns a tuple with the ObservedConversionEventsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedConversionEventsAvailable

`func (o *AdTestListRead) SetObservedConversionEventsAvailable(v bool)`

SetObservedConversionEventsAvailable sets ObservedConversionEventsAvailable field to given value.


### GetObservedConversionEvents

`func (o *AdTestListRead) GetObservedConversionEvents() []AdTestListReadObservedConversionEventsInner`

GetObservedConversionEvents returns the ObservedConversionEvents field if non-nil, zero value otherwise.

### GetObservedConversionEventsOk

`func (o *AdTestListRead) GetObservedConversionEventsOk() (*[]AdTestListReadObservedConversionEventsInner, bool)`

GetObservedConversionEventsOk returns a tuple with the ObservedConversionEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedConversionEvents

`func (o *AdTestListRead) SetObservedConversionEvents(v []AdTestListReadObservedConversionEventsInner)`

SetObservedConversionEvents sets ObservedConversionEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


