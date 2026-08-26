# ReplayDeliveriesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryIds** | Pointer to **[]int32** |  | [optional] 
**EventName** | Pointer to **NullableString** |  | [optional] 
**Limit** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewReplayDeliveriesRequest

`func NewReplayDeliveriesRequest() *ReplayDeliveriesRequest`

NewReplayDeliveriesRequest instantiates a new ReplayDeliveriesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplayDeliveriesRequestWithDefaults

`func NewReplayDeliveriesRequestWithDefaults() *ReplayDeliveriesRequest`

NewReplayDeliveriesRequestWithDefaults instantiates a new ReplayDeliveriesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryIds

`func (o *ReplayDeliveriesRequest) GetDeliveryIds() []int32`

GetDeliveryIds returns the DeliveryIds field if non-nil, zero value otherwise.

### GetDeliveryIdsOk

`func (o *ReplayDeliveriesRequest) GetDeliveryIdsOk() (*[]int32, bool)`

GetDeliveryIdsOk returns a tuple with the DeliveryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryIds

`func (o *ReplayDeliveriesRequest) SetDeliveryIds(v []int32)`

SetDeliveryIds sets DeliveryIds field to given value.

### HasDeliveryIds

`func (o *ReplayDeliveriesRequest) HasDeliveryIds() bool`

HasDeliveryIds returns a boolean if a field has been set.

### GetEventName

`func (o *ReplayDeliveriesRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *ReplayDeliveriesRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *ReplayDeliveriesRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *ReplayDeliveriesRequest) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### SetEventNameNil

`func (o *ReplayDeliveriesRequest) SetEventNameNil(b bool)`

 SetEventNameNil sets the value for EventName to be an explicit nil

### UnsetEventName
`func (o *ReplayDeliveriesRequest) UnsetEventName()`

UnsetEventName ensures that no value is present for EventName, not even an explicit nil
### GetLimit

`func (o *ReplayDeliveriesRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ReplayDeliveriesRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ReplayDeliveriesRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ReplayDeliveriesRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *ReplayDeliveriesRequest) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *ReplayDeliveriesRequest) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


