# GetInventoryMetrics200ResponseChangesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**PlatformAdAccountId** | **int32** |  | 
**Level** | **string** |  | 
**ExternalId** | **string** |  | 
**Field** | **string** |  | 
**OldValue** | **NullableString** |  | 
**NewValue** | **NullableString** |  | 
**ObservedAt** | **NullableString** |  | 

## Methods

### NewGetInventoryMetrics200ResponseChangesInner

`func NewGetInventoryMetrics200ResponseChangesInner(id int32, platformAdAccountId int32, level string, externalId string, field string, oldValue NullableString, newValue NullableString, observedAt NullableString, ) *GetInventoryMetrics200ResponseChangesInner`

NewGetInventoryMetrics200ResponseChangesInner instantiates a new GetInventoryMetrics200ResponseChangesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInventoryMetrics200ResponseChangesInnerWithDefaults

`func NewGetInventoryMetrics200ResponseChangesInnerWithDefaults() *GetInventoryMetrics200ResponseChangesInner`

NewGetInventoryMetrics200ResponseChangesInnerWithDefaults instantiates a new GetInventoryMetrics200ResponseChangesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInventoryMetrics200ResponseChangesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInventoryMetrics200ResponseChangesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInventoryMetrics200ResponseChangesInner) SetId(v int32)`

SetId sets Id field to given value.


### GetPlatformAdAccountId

`func (o *GetInventoryMetrics200ResponseChangesInner) GetPlatformAdAccountId() int32`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *GetInventoryMetrics200ResponseChangesInner) GetPlatformAdAccountIdOk() (*int32, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *GetInventoryMetrics200ResponseChangesInner) SetPlatformAdAccountId(v int32)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.


### GetLevel

`func (o *GetInventoryMetrics200ResponseChangesInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GetInventoryMetrics200ResponseChangesInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GetInventoryMetrics200ResponseChangesInner) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetExternalId

`func (o *GetInventoryMetrics200ResponseChangesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetInventoryMetrics200ResponseChangesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetInventoryMetrics200ResponseChangesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetField

`func (o *GetInventoryMetrics200ResponseChangesInner) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *GetInventoryMetrics200ResponseChangesInner) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *GetInventoryMetrics200ResponseChangesInner) SetField(v string)`

SetField sets Field field to given value.


### GetOldValue

`func (o *GetInventoryMetrics200ResponseChangesInner) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *GetInventoryMetrics200ResponseChangesInner) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *GetInventoryMetrics200ResponseChangesInner) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.


### SetOldValueNil

`func (o *GetInventoryMetrics200ResponseChangesInner) SetOldValueNil(b bool)`

 SetOldValueNil sets the value for OldValue to be an explicit nil

### UnsetOldValue
`func (o *GetInventoryMetrics200ResponseChangesInner) UnsetOldValue()`

UnsetOldValue ensures that no value is present for OldValue, not even an explicit nil
### GetNewValue

`func (o *GetInventoryMetrics200ResponseChangesInner) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *GetInventoryMetrics200ResponseChangesInner) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *GetInventoryMetrics200ResponseChangesInner) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.


### SetNewValueNil

`func (o *GetInventoryMetrics200ResponseChangesInner) SetNewValueNil(b bool)`

 SetNewValueNil sets the value for NewValue to be an explicit nil

### UnsetNewValue
`func (o *GetInventoryMetrics200ResponseChangesInner) UnsetNewValue()`

UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil
### GetObservedAt

`func (o *GetInventoryMetrics200ResponseChangesInner) GetObservedAt() string`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *GetInventoryMetrics200ResponseChangesInner) GetObservedAtOk() (*string, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *GetInventoryMetrics200ResponseChangesInner) SetObservedAt(v string)`

SetObservedAt sets ObservedAt field to given value.


### SetObservedAtNil

`func (o *GetInventoryMetrics200ResponseChangesInner) SetObservedAtNil(b bool)`

 SetObservedAtNil sets the value for ObservedAt to be an explicit nil

### UnsetObservedAt
`func (o *GetInventoryMetrics200ResponseChangesInner) UnsetObservedAt()`

UnsetObservedAt ensures that no value is present for ObservedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


