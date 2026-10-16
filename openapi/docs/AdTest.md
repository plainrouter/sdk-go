# AdTest

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
**MinimumConversions** | **int32** |  | 
**MaximumDays** | **int32** |  | 
**MaximumSpendMinor** | **NullableString** |  | 
**AnomalyThresholdPercent** | **string** |  | 
**Currency** | **string** |  | 
**Timezone** | **string** |  | 
**FirstDay** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**CountHealthClickThreshold** | **NullableInt32** |  | 

## Methods

### NewAdTest

`func NewAdTest(id string, workspaceId int32, platformAdAccountId int32, adSetId string, axis string, eventName string, arrivalsFloor int32, minimumConversions int32, maximumDays int32, maximumSpendMinor NullableString, anomalyThresholdPercent string, currency string, timezone string, firstDay time.Time, createdAt time.Time, updatedAt time.Time, countHealthClickThreshold NullableInt32, ) *AdTest`

NewAdTest instantiates a new AdTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdTestWithDefaults

`func NewAdTestWithDefaults() *AdTest`

NewAdTestWithDefaults instantiates a new AdTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdTest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdTest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdTest) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *AdTest) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *AdTest) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *AdTest) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetPlatformAdAccountId

`func (o *AdTest) GetPlatformAdAccountId() int32`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *AdTest) GetPlatformAdAccountIdOk() (*int32, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *AdTest) SetPlatformAdAccountId(v int32)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.


### GetAdSetId

`func (o *AdTest) GetAdSetId() string`

GetAdSetId returns the AdSetId field if non-nil, zero value otherwise.

### GetAdSetIdOk

`func (o *AdTest) GetAdSetIdOk() (*string, bool)`

GetAdSetIdOk returns a tuple with the AdSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSetId

`func (o *AdTest) SetAdSetId(v string)`

SetAdSetId sets AdSetId field to given value.


### GetAxis

`func (o *AdTest) GetAxis() string`

GetAxis returns the Axis field if non-nil, zero value otherwise.

### GetAxisOk

`func (o *AdTest) GetAxisOk() (*string, bool)`

GetAxisOk returns a tuple with the Axis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxis

`func (o *AdTest) SetAxis(v string)`

SetAxis sets Axis field to given value.


### GetEventName

`func (o *AdTest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *AdTest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *AdTest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetArrivalsFloor

`func (o *AdTest) GetArrivalsFloor() int32`

GetArrivalsFloor returns the ArrivalsFloor field if non-nil, zero value otherwise.

### GetArrivalsFloorOk

`func (o *AdTest) GetArrivalsFloorOk() (*int32, bool)`

GetArrivalsFloorOk returns a tuple with the ArrivalsFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalsFloor

`func (o *AdTest) SetArrivalsFloor(v int32)`

SetArrivalsFloor sets ArrivalsFloor field to given value.


### GetMinimumConversions

`func (o *AdTest) GetMinimumConversions() int32`

GetMinimumConversions returns the MinimumConversions field if non-nil, zero value otherwise.

### GetMinimumConversionsOk

`func (o *AdTest) GetMinimumConversionsOk() (*int32, bool)`

GetMinimumConversionsOk returns a tuple with the MinimumConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumConversions

`func (o *AdTest) SetMinimumConversions(v int32)`

SetMinimumConversions sets MinimumConversions field to given value.


### GetMaximumDays

`func (o *AdTest) GetMaximumDays() int32`

GetMaximumDays returns the MaximumDays field if non-nil, zero value otherwise.

### GetMaximumDaysOk

`func (o *AdTest) GetMaximumDaysOk() (*int32, bool)`

GetMaximumDaysOk returns a tuple with the MaximumDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDays

`func (o *AdTest) SetMaximumDays(v int32)`

SetMaximumDays sets MaximumDays field to given value.


### GetMaximumSpendMinor

`func (o *AdTest) GetMaximumSpendMinor() string`

GetMaximumSpendMinor returns the MaximumSpendMinor field if non-nil, zero value otherwise.

### GetMaximumSpendMinorOk

`func (o *AdTest) GetMaximumSpendMinorOk() (*string, bool)`

GetMaximumSpendMinorOk returns a tuple with the MaximumSpendMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSpendMinor

`func (o *AdTest) SetMaximumSpendMinor(v string)`

SetMaximumSpendMinor sets MaximumSpendMinor field to given value.


### SetMaximumSpendMinorNil

`func (o *AdTest) SetMaximumSpendMinorNil(b bool)`

 SetMaximumSpendMinorNil sets the value for MaximumSpendMinor to be an explicit nil

### UnsetMaximumSpendMinor
`func (o *AdTest) UnsetMaximumSpendMinor()`

UnsetMaximumSpendMinor ensures that no value is present for MaximumSpendMinor, not even an explicit nil
### GetAnomalyThresholdPercent

`func (o *AdTest) GetAnomalyThresholdPercent() string`

GetAnomalyThresholdPercent returns the AnomalyThresholdPercent field if non-nil, zero value otherwise.

### GetAnomalyThresholdPercentOk

`func (o *AdTest) GetAnomalyThresholdPercentOk() (*string, bool)`

GetAnomalyThresholdPercentOk returns a tuple with the AnomalyThresholdPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyThresholdPercent

`func (o *AdTest) SetAnomalyThresholdPercent(v string)`

SetAnomalyThresholdPercent sets AnomalyThresholdPercent field to given value.


### GetCurrency

`func (o *AdTest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AdTest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AdTest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetTimezone

`func (o *AdTest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AdTest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AdTest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetFirstDay

`func (o *AdTest) GetFirstDay() time.Time`

GetFirstDay returns the FirstDay field if non-nil, zero value otherwise.

### GetFirstDayOk

`func (o *AdTest) GetFirstDayOk() (*time.Time, bool)`

GetFirstDayOk returns a tuple with the FirstDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDay

`func (o *AdTest) SetFirstDay(v time.Time)`

SetFirstDay sets FirstDay field to given value.


### GetCreatedAt

`func (o *AdTest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdTest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdTest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AdTest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdTest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdTest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCountHealthClickThreshold

`func (o *AdTest) GetCountHealthClickThreshold() int32`

GetCountHealthClickThreshold returns the CountHealthClickThreshold field if non-nil, zero value otherwise.

### GetCountHealthClickThresholdOk

`func (o *AdTest) GetCountHealthClickThresholdOk() (*int32, bool)`

GetCountHealthClickThresholdOk returns a tuple with the CountHealthClickThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountHealthClickThreshold

`func (o *AdTest) SetCountHealthClickThreshold(v int32)`

SetCountHealthClickThreshold sets CountHealthClickThreshold field to given value.


### SetCountHealthClickThresholdNil

`func (o *AdTest) SetCountHealthClickThresholdNil(b bool)`

 SetCountHealthClickThresholdNil sets the value for CountHealthClickThreshold to be an explicit nil

### UnsetCountHealthClickThreshold
`func (o *AdTest) UnsetCountHealthClickThreshold()`

UnsetCountHealthClickThreshold ensures that no value is present for CountHealthClickThreshold, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


