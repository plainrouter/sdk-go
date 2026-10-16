# AdTestCreateReadTest

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

### NewAdTestCreateReadTest

`func NewAdTestCreateReadTest(id string, workspaceId int32, platformAdAccountId int32, adSetId string, axis string, eventName string, arrivalsFloor int32, minimumConversions int32, maximumDays int32, maximumSpendMinor NullableString, anomalyThresholdPercent string, firstDay string, timezone string, currency string, createdAt NullableTime, members []AdTestCreateReadTestMembersInner, pauseProposals []AdTestCreateReadTestPauseProposalsInner, verdict NullableAdTestCreateReadTestVerdict, ) *AdTestCreateReadTest`

NewAdTestCreateReadTest instantiates a new AdTestCreateReadTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdTestCreateReadTestWithDefaults

`func NewAdTestCreateReadTestWithDefaults() *AdTestCreateReadTest`

NewAdTestCreateReadTestWithDefaults instantiates a new AdTestCreateReadTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdTestCreateReadTest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdTestCreateReadTest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdTestCreateReadTest) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *AdTestCreateReadTest) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *AdTestCreateReadTest) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *AdTestCreateReadTest) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetPlatformAdAccountId

`func (o *AdTestCreateReadTest) GetPlatformAdAccountId() int32`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *AdTestCreateReadTest) GetPlatformAdAccountIdOk() (*int32, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *AdTestCreateReadTest) SetPlatformAdAccountId(v int32)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.


### GetAdSetId

`func (o *AdTestCreateReadTest) GetAdSetId() string`

GetAdSetId returns the AdSetId field if non-nil, zero value otherwise.

### GetAdSetIdOk

`func (o *AdTestCreateReadTest) GetAdSetIdOk() (*string, bool)`

GetAdSetIdOk returns a tuple with the AdSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSetId

`func (o *AdTestCreateReadTest) SetAdSetId(v string)`

SetAdSetId sets AdSetId field to given value.


### GetAxis

`func (o *AdTestCreateReadTest) GetAxis() string`

GetAxis returns the Axis field if non-nil, zero value otherwise.

### GetAxisOk

`func (o *AdTestCreateReadTest) GetAxisOk() (*string, bool)`

GetAxisOk returns a tuple with the Axis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxis

`func (o *AdTestCreateReadTest) SetAxis(v string)`

SetAxis sets Axis field to given value.


### GetEventName

`func (o *AdTestCreateReadTest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *AdTestCreateReadTest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *AdTestCreateReadTest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetArrivalsFloor

`func (o *AdTestCreateReadTest) GetArrivalsFloor() int32`

GetArrivalsFloor returns the ArrivalsFloor field if non-nil, zero value otherwise.

### GetArrivalsFloorOk

`func (o *AdTestCreateReadTest) GetArrivalsFloorOk() (*int32, bool)`

GetArrivalsFloorOk returns a tuple with the ArrivalsFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalsFloor

`func (o *AdTestCreateReadTest) SetArrivalsFloor(v int32)`

SetArrivalsFloor sets ArrivalsFloor field to given value.


### GetCountHealthClickThreshold

`func (o *AdTestCreateReadTest) GetCountHealthClickThreshold() int32`

GetCountHealthClickThreshold returns the CountHealthClickThreshold field if non-nil, zero value otherwise.

### GetCountHealthClickThresholdOk

`func (o *AdTestCreateReadTest) GetCountHealthClickThresholdOk() (*int32, bool)`

GetCountHealthClickThresholdOk returns a tuple with the CountHealthClickThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountHealthClickThreshold

`func (o *AdTestCreateReadTest) SetCountHealthClickThreshold(v int32)`

SetCountHealthClickThreshold sets CountHealthClickThreshold field to given value.

### HasCountHealthClickThreshold

`func (o *AdTestCreateReadTest) HasCountHealthClickThreshold() bool`

HasCountHealthClickThreshold returns a boolean if a field has been set.

### SetCountHealthClickThresholdNil

`func (o *AdTestCreateReadTest) SetCountHealthClickThresholdNil(b bool)`

 SetCountHealthClickThresholdNil sets the value for CountHealthClickThreshold to be an explicit nil

### UnsetCountHealthClickThreshold
`func (o *AdTestCreateReadTest) UnsetCountHealthClickThreshold()`

UnsetCountHealthClickThreshold ensures that no value is present for CountHealthClickThreshold, not even an explicit nil
### GetMinimumConversions

`func (o *AdTestCreateReadTest) GetMinimumConversions() int32`

GetMinimumConversions returns the MinimumConversions field if non-nil, zero value otherwise.

### GetMinimumConversionsOk

`func (o *AdTestCreateReadTest) GetMinimumConversionsOk() (*int32, bool)`

GetMinimumConversionsOk returns a tuple with the MinimumConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumConversions

`func (o *AdTestCreateReadTest) SetMinimumConversions(v int32)`

SetMinimumConversions sets MinimumConversions field to given value.


### GetMaximumDays

`func (o *AdTestCreateReadTest) GetMaximumDays() int32`

GetMaximumDays returns the MaximumDays field if non-nil, zero value otherwise.

### GetMaximumDaysOk

`func (o *AdTestCreateReadTest) GetMaximumDaysOk() (*int32, bool)`

GetMaximumDaysOk returns a tuple with the MaximumDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDays

`func (o *AdTestCreateReadTest) SetMaximumDays(v int32)`

SetMaximumDays sets MaximumDays field to given value.


### GetMaximumSpendMinor

`func (o *AdTestCreateReadTest) GetMaximumSpendMinor() string`

GetMaximumSpendMinor returns the MaximumSpendMinor field if non-nil, zero value otherwise.

### GetMaximumSpendMinorOk

`func (o *AdTestCreateReadTest) GetMaximumSpendMinorOk() (*string, bool)`

GetMaximumSpendMinorOk returns a tuple with the MaximumSpendMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSpendMinor

`func (o *AdTestCreateReadTest) SetMaximumSpendMinor(v string)`

SetMaximumSpendMinor sets MaximumSpendMinor field to given value.


### SetMaximumSpendMinorNil

`func (o *AdTestCreateReadTest) SetMaximumSpendMinorNil(b bool)`

 SetMaximumSpendMinorNil sets the value for MaximumSpendMinor to be an explicit nil

### UnsetMaximumSpendMinor
`func (o *AdTestCreateReadTest) UnsetMaximumSpendMinor()`

UnsetMaximumSpendMinor ensures that no value is present for MaximumSpendMinor, not even an explicit nil
### GetAnomalyThresholdPercent

`func (o *AdTestCreateReadTest) GetAnomalyThresholdPercent() string`

GetAnomalyThresholdPercent returns the AnomalyThresholdPercent field if non-nil, zero value otherwise.

### GetAnomalyThresholdPercentOk

`func (o *AdTestCreateReadTest) GetAnomalyThresholdPercentOk() (*string, bool)`

GetAnomalyThresholdPercentOk returns a tuple with the AnomalyThresholdPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyThresholdPercent

`func (o *AdTestCreateReadTest) SetAnomalyThresholdPercent(v string)`

SetAnomalyThresholdPercent sets AnomalyThresholdPercent field to given value.


### GetFirstDay

`func (o *AdTestCreateReadTest) GetFirstDay() string`

GetFirstDay returns the FirstDay field if non-nil, zero value otherwise.

### GetFirstDayOk

`func (o *AdTestCreateReadTest) GetFirstDayOk() (*string, bool)`

GetFirstDayOk returns a tuple with the FirstDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDay

`func (o *AdTestCreateReadTest) SetFirstDay(v string)`

SetFirstDay sets FirstDay field to given value.


### GetTimezone

`func (o *AdTestCreateReadTest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AdTestCreateReadTest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AdTestCreateReadTest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetCurrency

`func (o *AdTestCreateReadTest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AdTestCreateReadTest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AdTestCreateReadTest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCreatedAt

`func (o *AdTestCreateReadTest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdTestCreateReadTest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdTestCreateReadTest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *AdTestCreateReadTest) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *AdTestCreateReadTest) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetMembers

`func (o *AdTestCreateReadTest) GetMembers() []AdTestCreateReadTestMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AdTestCreateReadTest) GetMembersOk() (*[]AdTestCreateReadTestMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AdTestCreateReadTest) SetMembers(v []AdTestCreateReadTestMembersInner)`

SetMembers sets Members field to given value.


### GetPauseProposals

`func (o *AdTestCreateReadTest) GetPauseProposals() []AdTestCreateReadTestPauseProposalsInner`

GetPauseProposals returns the PauseProposals field if non-nil, zero value otherwise.

### GetPauseProposalsOk

`func (o *AdTestCreateReadTest) GetPauseProposalsOk() (*[]AdTestCreateReadTestPauseProposalsInner, bool)`

GetPauseProposalsOk returns a tuple with the PauseProposals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseProposals

`func (o *AdTestCreateReadTest) SetPauseProposals(v []AdTestCreateReadTestPauseProposalsInner)`

SetPauseProposals sets PauseProposals field to given value.


### GetVerdict

`func (o *AdTestCreateReadTest) GetVerdict() AdTestCreateReadTestVerdict`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *AdTestCreateReadTest) GetVerdictOk() (*AdTestCreateReadTestVerdict, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *AdTestCreateReadTest) SetVerdict(v AdTestCreateReadTestVerdict)`

SetVerdict sets Verdict field to given value.


### SetVerdictNil

`func (o *AdTestCreateReadTest) SetVerdictNil(b bool)`

 SetVerdictNil sets the value for Verdict to be an explicit nil

### UnsetVerdict
`func (o *AdTestCreateReadTest) UnsetVerdict()`

UnsetVerdict ensures that no value is present for Verdict, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


