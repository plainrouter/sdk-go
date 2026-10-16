# GetInventoryMetrics200ResponseObjectMetricsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** |  | 
**ExternalId** | **string** |  | 
**Date** | **string** |  | 
**Meta** | [**NullableGetInventoryMetrics200ResponseObjectMetricsInnerMeta**](GetInventoryMetrics200ResponseObjectMetricsInnerMeta.md) |  | 
**Arrivals** | **NullableInt32** |  | 
**CreativeId** | Pointer to **NullableString** |  | [optional] 
**Generator** | Pointer to **NullableString** |  | [optional] 
**GeneratorJob** | Pointer to **NullableString** |  | [optional] 
**ParentCreative** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetInventoryMetrics200ResponseObjectMetricsInner

`func NewGetInventoryMetrics200ResponseObjectMetricsInner(level string, externalId string, date string, meta NullableGetInventoryMetrics200ResponseObjectMetricsInnerMeta, arrivals NullableInt32, ) *GetInventoryMetrics200ResponseObjectMetricsInner`

NewGetInventoryMetrics200ResponseObjectMetricsInner instantiates a new GetInventoryMetrics200ResponseObjectMetricsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInventoryMetrics200ResponseObjectMetricsInnerWithDefaults

`func NewGetInventoryMetrics200ResponseObjectMetricsInnerWithDefaults() *GetInventoryMetrics200ResponseObjectMetricsInner`

NewGetInventoryMetrics200ResponseObjectMetricsInnerWithDefaults instantiates a new GetInventoryMetrics200ResponseObjectMetricsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetExternalId

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetDate

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) SetDate(v string)`

SetDate sets Date field to given value.


### GetMeta

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetMeta() GetInventoryMetrics200ResponseObjectMetricsInnerMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetMetaOk() (*GetInventoryMetrics200ResponseObjectMetricsInnerMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) SetMeta(v GetInventoryMetrics200ResponseObjectMetricsInnerMeta)`

SetMeta sets Meta field to given value.


### SetMetaNil

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetArrivals

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetArrivals() int32`

GetArrivals returns the Arrivals field if non-nil, zero value otherwise.

### GetArrivalsOk

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetArrivalsOk() (*int32, bool)`

GetArrivalsOk returns a tuple with the Arrivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivals

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) SetArrivals(v int32)`

SetArrivals sets Arrivals field to given value.


### SetArrivalsNil

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) SetArrivalsNil(b bool)`

 SetArrivalsNil sets the value for Arrivals to be an explicit nil

### UnsetArrivals
`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) UnsetArrivals()`

UnsetArrivals ensures that no value is present for Arrivals, not even an explicit nil
### GetCreativeId

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetCreativeId() string`

GetCreativeId returns the CreativeId field if non-nil, zero value otherwise.

### GetCreativeIdOk

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetCreativeIdOk() (*string, bool)`

GetCreativeIdOk returns a tuple with the CreativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeId

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) SetCreativeId(v string)`

SetCreativeId sets CreativeId field to given value.

### HasCreativeId

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) HasCreativeId() bool`

HasCreativeId returns a boolean if a field has been set.

### SetCreativeIdNil

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) SetCreativeIdNil(b bool)`

 SetCreativeIdNil sets the value for CreativeId to be an explicit nil

### UnsetCreativeId
`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) UnsetCreativeId()`

UnsetCreativeId ensures that no value is present for CreativeId, not even an explicit nil
### GetGenerator

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetGenerator() string`

GetGenerator returns the Generator field if non-nil, zero value otherwise.

### GetGeneratorOk

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetGeneratorOk() (*string, bool)`

GetGeneratorOk returns a tuple with the Generator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerator

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) SetGenerator(v string)`

SetGenerator sets Generator field to given value.

### HasGenerator

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) HasGenerator() bool`

HasGenerator returns a boolean if a field has been set.

### SetGeneratorNil

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) SetGeneratorNil(b bool)`

 SetGeneratorNil sets the value for Generator to be an explicit nil

### UnsetGenerator
`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) UnsetGenerator()`

UnsetGenerator ensures that no value is present for Generator, not even an explicit nil
### GetGeneratorJob

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetGeneratorJob() string`

GetGeneratorJob returns the GeneratorJob field if non-nil, zero value otherwise.

### GetGeneratorJobOk

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetGeneratorJobOk() (*string, bool)`

GetGeneratorJobOk returns a tuple with the GeneratorJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorJob

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) SetGeneratorJob(v string)`

SetGeneratorJob sets GeneratorJob field to given value.

### HasGeneratorJob

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) HasGeneratorJob() bool`

HasGeneratorJob returns a boolean if a field has been set.

### SetGeneratorJobNil

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) SetGeneratorJobNil(b bool)`

 SetGeneratorJobNil sets the value for GeneratorJob to be an explicit nil

### UnsetGeneratorJob
`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) UnsetGeneratorJob()`

UnsetGeneratorJob ensures that no value is present for GeneratorJob, not even an explicit nil
### GetParentCreative

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetParentCreative() string`

GetParentCreative returns the ParentCreative field if non-nil, zero value otherwise.

### GetParentCreativeOk

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) GetParentCreativeOk() (*string, bool)`

GetParentCreativeOk returns a tuple with the ParentCreative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCreative

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) SetParentCreative(v string)`

SetParentCreative sets ParentCreative field to given value.

### HasParentCreative

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) HasParentCreative() bool`

HasParentCreative returns a boolean if a field has been set.

### SetParentCreativeNil

`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) SetParentCreativeNil(b bool)`

 SetParentCreativeNil sets the value for ParentCreative to be an explicit nil

### UnsetParentCreative
`func (o *GetInventoryMetrics200ResponseObjectMetricsInner) UnsetParentCreative()`

UnsetParentCreative ensures that no value is present for ParentCreative, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


