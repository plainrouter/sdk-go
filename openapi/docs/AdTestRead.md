# AdTestRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**WorkspaceId** | **int32** |  | 
**PlatformAdAccountId** | **int32** |  | 
**AdSetId** | **string** |  | 
**Axis** | **string** |  | 
**EventName** | **string** |  | 
**ArrivalsFloor** | **int32** |  | 
**CountHealthClickThreshold** | Pointer to **NullableInt32** | Link-click threshold frozen at conclusion; null until concluded or for legacy Tests. | [optional] 
**MinimumConversions** | **int32** |  | 
**MaximumDays** | **int32** |  | 
**MaximumSpendMinor** | **NullableString** | Optional maximum in positive whole Meta minor units, using the same unit scale as Launch plan budget_amount_minor. | 
**AnomalyThresholdPercent** | **string** |  | 
**FirstDay** | **string** |  | 
**Timezone** | **string** |  | 
**Currency** | **string** |  | 
**CreatedAt** | **NullableTime** |  | 
**Members** | [**[]AdTestCreateReadTestMembersInner**](AdTestCreateReadTestMembersInner.md) |  | 
**PauseProposals** | [**[]AdTestCreateReadTestPauseProposalsInner**](AdTestCreateReadTestPauseProposalsInner.md) |  | 
**Verdict** | [**NullableAdTestCreateReadTestVerdict**](AdTestCreateReadTestVerdict.md) |  | 

## Methods

### NewAdTestRead

`func NewAdTestRead(id string, workspaceId int32, platformAdAccountId int32, adSetId string, axis string, eventName string, arrivalsFloor int32, minimumConversions int32, maximumDays int32, maximumSpendMinor NullableString, anomalyThresholdPercent string, firstDay string, timezone string, currency string, createdAt NullableTime, members []AdTestCreateReadTestMembersInner, pauseProposals []AdTestCreateReadTestPauseProposalsInner, verdict NullableAdTestCreateReadTestVerdict, ) *AdTestRead`

NewAdTestRead instantiates a new AdTestRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdTestReadWithDefaults

`func NewAdTestReadWithDefaults() *AdTestRead`

NewAdTestReadWithDefaults instantiates a new AdTestRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdTestRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdTestRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdTestRead) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *AdTestRead) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *AdTestRead) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *AdTestRead) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetPlatformAdAccountId

`func (o *AdTestRead) GetPlatformAdAccountId() int32`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *AdTestRead) GetPlatformAdAccountIdOk() (*int32, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *AdTestRead) SetPlatformAdAccountId(v int32)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.


### GetAdSetId

`func (o *AdTestRead) GetAdSetId() string`

GetAdSetId returns the AdSetId field if non-nil, zero value otherwise.

### GetAdSetIdOk

`func (o *AdTestRead) GetAdSetIdOk() (*string, bool)`

GetAdSetIdOk returns a tuple with the AdSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSetId

`func (o *AdTestRead) SetAdSetId(v string)`

SetAdSetId sets AdSetId field to given value.


### GetAxis

`func (o *AdTestRead) GetAxis() string`

GetAxis returns the Axis field if non-nil, zero value otherwise.

### GetAxisOk

`func (o *AdTestRead) GetAxisOk() (*string, bool)`

GetAxisOk returns a tuple with the Axis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxis

`func (o *AdTestRead) SetAxis(v string)`

SetAxis sets Axis field to given value.


### GetEventName

`func (o *AdTestRead) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *AdTestRead) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *AdTestRead) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetArrivalsFloor

`func (o *AdTestRead) GetArrivalsFloor() int32`

GetArrivalsFloor returns the ArrivalsFloor field if non-nil, zero value otherwise.

### GetArrivalsFloorOk

`func (o *AdTestRead) GetArrivalsFloorOk() (*int32, bool)`

GetArrivalsFloorOk returns a tuple with the ArrivalsFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalsFloor

`func (o *AdTestRead) SetArrivalsFloor(v int32)`

SetArrivalsFloor sets ArrivalsFloor field to given value.


### GetCountHealthClickThreshold

`func (o *AdTestRead) GetCountHealthClickThreshold() int32`

GetCountHealthClickThreshold returns the CountHealthClickThreshold field if non-nil, zero value otherwise.

### GetCountHealthClickThresholdOk

`func (o *AdTestRead) GetCountHealthClickThresholdOk() (*int32, bool)`

GetCountHealthClickThresholdOk returns a tuple with the CountHealthClickThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountHealthClickThreshold

`func (o *AdTestRead) SetCountHealthClickThreshold(v int32)`

SetCountHealthClickThreshold sets CountHealthClickThreshold field to given value.

### HasCountHealthClickThreshold

`func (o *AdTestRead) HasCountHealthClickThreshold() bool`

HasCountHealthClickThreshold returns a boolean if a field has been set.

### SetCountHealthClickThresholdNil

`func (o *AdTestRead) SetCountHealthClickThresholdNil(b bool)`

 SetCountHealthClickThresholdNil sets the value for CountHealthClickThreshold to be an explicit nil

### UnsetCountHealthClickThreshold
`func (o *AdTestRead) UnsetCountHealthClickThreshold()`

UnsetCountHealthClickThreshold ensures that no value is present for CountHealthClickThreshold, not even an explicit nil
### GetMinimumConversions

`func (o *AdTestRead) GetMinimumConversions() int32`

GetMinimumConversions returns the MinimumConversions field if non-nil, zero value otherwise.

### GetMinimumConversionsOk

`func (o *AdTestRead) GetMinimumConversionsOk() (*int32, bool)`

GetMinimumConversionsOk returns a tuple with the MinimumConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumConversions

`func (o *AdTestRead) SetMinimumConversions(v int32)`

SetMinimumConversions sets MinimumConversions field to given value.


### GetMaximumDays

`func (o *AdTestRead) GetMaximumDays() int32`

GetMaximumDays returns the MaximumDays field if non-nil, zero value otherwise.

### GetMaximumDaysOk

`func (o *AdTestRead) GetMaximumDaysOk() (*int32, bool)`

GetMaximumDaysOk returns a tuple with the MaximumDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDays

`func (o *AdTestRead) SetMaximumDays(v int32)`

SetMaximumDays sets MaximumDays field to given value.


### GetMaximumSpendMinor

`func (o *AdTestRead) GetMaximumSpendMinor() string`

GetMaximumSpendMinor returns the MaximumSpendMinor field if non-nil, zero value otherwise.

### GetMaximumSpendMinorOk

`func (o *AdTestRead) GetMaximumSpendMinorOk() (*string, bool)`

GetMaximumSpendMinorOk returns a tuple with the MaximumSpendMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSpendMinor

`func (o *AdTestRead) SetMaximumSpendMinor(v string)`

SetMaximumSpendMinor sets MaximumSpendMinor field to given value.


### SetMaximumSpendMinorNil

`func (o *AdTestRead) SetMaximumSpendMinorNil(b bool)`

 SetMaximumSpendMinorNil sets the value for MaximumSpendMinor to be an explicit nil

### UnsetMaximumSpendMinor
`func (o *AdTestRead) UnsetMaximumSpendMinor()`

UnsetMaximumSpendMinor ensures that no value is present for MaximumSpendMinor, not even an explicit nil
### GetAnomalyThresholdPercent

`func (o *AdTestRead) GetAnomalyThresholdPercent() string`

GetAnomalyThresholdPercent returns the AnomalyThresholdPercent field if non-nil, zero value otherwise.

### GetAnomalyThresholdPercentOk

`func (o *AdTestRead) GetAnomalyThresholdPercentOk() (*string, bool)`

GetAnomalyThresholdPercentOk returns a tuple with the AnomalyThresholdPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyThresholdPercent

`func (o *AdTestRead) SetAnomalyThresholdPercent(v string)`

SetAnomalyThresholdPercent sets AnomalyThresholdPercent field to given value.


### GetFirstDay

`func (o *AdTestRead) GetFirstDay() string`

GetFirstDay returns the FirstDay field if non-nil, zero value otherwise.

### GetFirstDayOk

`func (o *AdTestRead) GetFirstDayOk() (*string, bool)`

GetFirstDayOk returns a tuple with the FirstDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDay

`func (o *AdTestRead) SetFirstDay(v string)`

SetFirstDay sets FirstDay field to given value.


### GetTimezone

`func (o *AdTestRead) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AdTestRead) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AdTestRead) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetCurrency

`func (o *AdTestRead) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AdTestRead) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AdTestRead) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCreatedAt

`func (o *AdTestRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdTestRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdTestRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *AdTestRead) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *AdTestRead) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetMembers

`func (o *AdTestRead) GetMembers() []AdTestCreateReadTestMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AdTestRead) GetMembersOk() (*[]AdTestCreateReadTestMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AdTestRead) SetMembers(v []AdTestCreateReadTestMembersInner)`

SetMembers sets Members field to given value.


### GetPauseProposals

`func (o *AdTestRead) GetPauseProposals() []AdTestCreateReadTestPauseProposalsInner`

GetPauseProposals returns the PauseProposals field if non-nil, zero value otherwise.

### GetPauseProposalsOk

`func (o *AdTestRead) GetPauseProposalsOk() (*[]AdTestCreateReadTestPauseProposalsInner, bool)`

GetPauseProposalsOk returns a tuple with the PauseProposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseProposals

`func (o *AdTestRead) SetPauseProposals(v []AdTestCreateReadTestPauseProposalsInner)`

SetPauseProposals sets PauseProposals field to given value.


### GetVerdict

`func (o *AdTestRead) GetVerdict() AdTestCreateReadTestVerdict`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *AdTestRead) GetVerdictOk() (*AdTestCreateReadTestVerdict, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *AdTestRead) SetVerdict(v AdTestCreateReadTestVerdict)`

SetVerdict sets Verdict field to given value.


### SetVerdictNil

`func (o *AdTestRead) SetVerdictNil(b bool)`

 SetVerdictNil sets the value for Verdict to be an explicit nil

### UnsetVerdict
`func (o *AdTestRead) UnsetVerdict()`

UnsetVerdict ensures that no value is present for Verdict, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


