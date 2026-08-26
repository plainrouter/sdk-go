# SetDestinationTestModeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**TestEventCode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSetDestinationTestModeRequest

`func NewSetDestinationTestModeRequest(enabled bool, ) *SetDestinationTestModeRequest`

NewSetDestinationTestModeRequest instantiates a new SetDestinationTestModeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetDestinationTestModeRequestWithDefaults

`func NewSetDestinationTestModeRequestWithDefaults() *SetDestinationTestModeRequest`

NewSetDestinationTestModeRequestWithDefaults instantiates a new SetDestinationTestModeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SetDestinationTestModeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SetDestinationTestModeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SetDestinationTestModeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetTestEventCode

`func (o *SetDestinationTestModeRequest) GetTestEventCode() string`

GetTestEventCode returns the TestEventCode field if non-nil, zero value otherwise.

### GetTestEventCodeOk

`func (o *SetDestinationTestModeRequest) GetTestEventCodeOk() (*string, bool)`

GetTestEventCodeOk returns a tuple with the TestEventCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestEventCode

`func (o *SetDestinationTestModeRequest) SetTestEventCode(v string)`

SetTestEventCode sets TestEventCode field to given value.

### HasTestEventCode

`func (o *SetDestinationTestModeRequest) HasTestEventCode() bool`

HasTestEventCode returns a boolean if a field has been set.

### SetTestEventCodeNil

`func (o *SetDestinationTestModeRequest) SetTestEventCodeNil(b bool)`

 SetTestEventCodeNil sets the value for TestEventCode to be an explicit nil

### UnsetTestEventCode
`func (o *SetDestinationTestModeRequest) UnsetTestEventCode()`

UnsetTestEventCode ensures that no value is present for TestEventCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


