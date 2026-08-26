# ReplayDeliveries202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queued** | **int32** |  | 
**Expired** | **int32** |  | 
**PayloadExpired** | **int32** |  | 
**Capped** | **bool** |  | 

## Methods

### NewReplayDeliveries202Response

`func NewReplayDeliveries202Response(queued int32, expired int32, payloadExpired int32, capped bool, ) *ReplayDeliveries202Response`

NewReplayDeliveries202Response instantiates a new ReplayDeliveries202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplayDeliveries202ResponseWithDefaults

`func NewReplayDeliveries202ResponseWithDefaults() *ReplayDeliveries202Response`

NewReplayDeliveries202ResponseWithDefaults instantiates a new ReplayDeliveries202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueued

`func (o *ReplayDeliveries202Response) GetQueued() int32`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *ReplayDeliveries202Response) GetQueuedOk() (*int32, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *ReplayDeliveries202Response) SetQueued(v int32)`

SetQueued sets Queued field to given value.


### GetExpired

`func (o *ReplayDeliveries202Response) GetExpired() int32`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *ReplayDeliveries202Response) GetExpiredOk() (*int32, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *ReplayDeliveries202Response) SetExpired(v int32)`

SetExpired sets Expired field to given value.


### GetPayloadExpired

`func (o *ReplayDeliveries202Response) GetPayloadExpired() int32`

GetPayloadExpired returns the PayloadExpired field if non-nil, zero value otherwise.

### GetPayloadExpiredOk

`func (o *ReplayDeliveries202Response) GetPayloadExpiredOk() (*int32, bool)`

GetPayloadExpiredOk returns a tuple with the PayloadExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadExpired

`func (o *ReplayDeliveries202Response) SetPayloadExpired(v int32)`

SetPayloadExpired sets PayloadExpired field to given value.


### GetCapped

`func (o *ReplayDeliveries202Response) GetCapped() bool`

GetCapped returns the Capped field if non-nil, zero value otherwise.

### GetCappedOk

`func (o *ReplayDeliveries202Response) GetCappedOk() (*bool, bool)`

GetCappedOk returns a tuple with the Capped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapped

`func (o *ReplayDeliveries202Response) SetCapped(v bool)`

SetCapped sets Capped field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


