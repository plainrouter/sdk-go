# ActionsTestsApiStoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformAdAccountId** | **int32** |  | 
**AdSetId** | **string** |  | 
**MemberAdIds** | **[]string** | Two to four distinct Meta ad IDs shown in the selected ad set inventory. | 
**Axis** | **string** |  | 
**EventName** | **string** | An observed workspace conversion event other than PageView. | 
**ArrivalsFloor** | Pointer to **int32** | Optional override; default 50 counted arrivals per member. | [optional] 
**MinimumConversions** | Pointer to **int32** | Optional override; default 10 recorded conversions for a winner. | [optional] 
**MaximumDays** | Pointer to **int32** | Optional override; default 14 complete account-local days. | [optional] 
**MaximumSpendMinor** | Pointer to [**NullableActionsTestsApiStoreRequestMaximumSpendMinor**](ActionsTestsApiStoreRequestMaximumSpendMinor.md) |  | [optional] 

## Methods

### NewActionsTestsApiStoreRequest

`func NewActionsTestsApiStoreRequest(platformAdAccountId int32, adSetId string, memberAdIds []string, axis string, eventName string, ) *ActionsTestsApiStoreRequest`

NewActionsTestsApiStoreRequest instantiates a new ActionsTestsApiStoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsTestsApiStoreRequestWithDefaults

`func NewActionsTestsApiStoreRequestWithDefaults() *ActionsTestsApiStoreRequest`

NewActionsTestsApiStoreRequestWithDefaults instantiates a new ActionsTestsApiStoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformAdAccountId

`func (o *ActionsTestsApiStoreRequest) GetPlatformAdAccountId() int32`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *ActionsTestsApiStoreRequest) GetPlatformAdAccountIdOk() (*int32, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *ActionsTestsApiStoreRequest) SetPlatformAdAccountId(v int32)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.


### GetAdSetId

`func (o *ActionsTestsApiStoreRequest) GetAdSetId() string`

GetAdSetId returns the AdSetId field if non-nil, zero value otherwise.

### GetAdSetIdOk

`func (o *ActionsTestsApiStoreRequest) GetAdSetIdOk() (*string, bool)`

GetAdSetIdOk returns a tuple with the AdSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSetId

`func (o *ActionsTestsApiStoreRequest) SetAdSetId(v string)`

SetAdSetId sets AdSetId field to given value.


### GetMemberAdIds

`func (o *ActionsTestsApiStoreRequest) GetMemberAdIds() []string`

GetMemberAdIds returns the MemberAdIds field if non-nil, zero value otherwise.

### GetMemberAdIdsOk

`func (o *ActionsTestsApiStoreRequest) GetMemberAdIdsOk() (*[]string, bool)`

GetMemberAdIdsOk returns a tuple with the MemberAdIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberAdIds

`func (o *ActionsTestsApiStoreRequest) SetMemberAdIds(v []string)`

SetMemberAdIds sets MemberAdIds field to given value.


### GetAxis

`func (o *ActionsTestsApiStoreRequest) GetAxis() string`

GetAxis returns the Axis field if non-nil, zero value otherwise.

### GetAxisOk

`func (o *ActionsTestsApiStoreRequest) GetAxisOk() (*string, bool)`

GetAxisOk returns a tuple with the Axis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxis

`func (o *ActionsTestsApiStoreRequest) SetAxis(v string)`

SetAxis sets Axis field to given value.


### GetEventName

`func (o *ActionsTestsApiStoreRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *ActionsTestsApiStoreRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *ActionsTestsApiStoreRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetArrivalsFloor

`func (o *ActionsTestsApiStoreRequest) GetArrivalsFloor() int32`

GetArrivalsFloor returns the ArrivalsFloor field if non-nil, zero value otherwise.

### GetArrivalsFloorOk

`func (o *ActionsTestsApiStoreRequest) GetArrivalsFloorOk() (*int32, bool)`

GetArrivalsFloorOk returns a tuple with the ArrivalsFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalsFloor

`func (o *ActionsTestsApiStoreRequest) SetArrivalsFloor(v int32)`

SetArrivalsFloor sets ArrivalsFloor field to given value.

### HasArrivalsFloor

`func (o *ActionsTestsApiStoreRequest) HasArrivalsFloor() bool`

HasArrivalsFloor returns a boolean if a field has been set.

### GetMinimumConversions

`func (o *ActionsTestsApiStoreRequest) GetMinimumConversions() int32`

GetMinimumConversions returns the MinimumConversions field if non-nil, zero value otherwise.

### GetMinimumConversionsOk

`func (o *ActionsTestsApiStoreRequest) GetMinimumConversionsOk() (*int32, bool)`

GetMinimumConversionsOk returns a tuple with the MinimumConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumConversions

`func (o *ActionsTestsApiStoreRequest) SetMinimumConversions(v int32)`

SetMinimumConversions sets MinimumConversions field to given value.

### HasMinimumConversions

`func (o *ActionsTestsApiStoreRequest) HasMinimumConversions() bool`

HasMinimumConversions returns a boolean if a field has been set.

### GetMaximumDays

`func (o *ActionsTestsApiStoreRequest) GetMaximumDays() int32`

GetMaximumDays returns the MaximumDays field if non-nil, zero value otherwise.

### GetMaximumDaysOk

`func (o *ActionsTestsApiStoreRequest) GetMaximumDaysOk() (*int32, bool)`

GetMaximumDaysOk returns a tuple with the MaximumDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDays

`func (o *ActionsTestsApiStoreRequest) SetMaximumDays(v int32)`

SetMaximumDays sets MaximumDays field to given value.

### HasMaximumDays

`func (o *ActionsTestsApiStoreRequest) HasMaximumDays() bool`

HasMaximumDays returns a boolean if a field has been set.

### GetMaximumSpendMinor

`func (o *ActionsTestsApiStoreRequest) GetMaximumSpendMinor() ActionsTestsApiStoreRequestMaximumSpendMinor`

GetMaximumSpendMinor returns the MaximumSpendMinor field if non-nil, zero value otherwise.

### GetMaximumSpendMinorOk

`func (o *ActionsTestsApiStoreRequest) GetMaximumSpendMinorOk() (*ActionsTestsApiStoreRequestMaximumSpendMinor, bool)`

GetMaximumSpendMinorOk returns a tuple with the MaximumSpendMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSpendMinor

`func (o *ActionsTestsApiStoreRequest) SetMaximumSpendMinor(v ActionsTestsApiStoreRequestMaximumSpendMinor)`

SetMaximumSpendMinor sets MaximumSpendMinor field to given value.

### HasMaximumSpendMinor

`func (o *ActionsTestsApiStoreRequest) HasMaximumSpendMinor() bool`

HasMaximumSpendMinor returns a boolean if a field has been set.

### SetMaximumSpendMinorNil

`func (o *ActionsTestsApiStoreRequest) SetMaximumSpendMinorNil(b bool)`

 SetMaximumSpendMinorNil sets the value for MaximumSpendMinor to be an explicit nil

### UnsetMaximumSpendMinor
`func (o *ActionsTestsApiStoreRequest) UnsetMaximumSpendMinor()`

UnsetMaximumSpendMinor ensures that no value is present for MaximumSpendMinor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


