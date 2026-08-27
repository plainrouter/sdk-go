# SetDestinationTestMode200ResponseDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SignalTrackerId** | **string** |  | 
**PlatformAdAccountId** | **NullableInt32** |  | 
**Type** | [**DestinationType**](DestinationType.md) |  | 
**CredentialSource** | [**DestinationCredentialSource**](DestinationCredentialSource.md) |  | 
**Config** | **interface{}** |  | 
**Status** | [**DestinationStatus**](DestinationStatus.md) |  | 
**CreatedAt** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 

## Methods

### NewSetDestinationTestMode200ResponseDestination

`func NewSetDestinationTestMode200ResponseDestination(id string, signalTrackerId string, platformAdAccountId NullableInt32, type_ DestinationType, credentialSource DestinationCredentialSource, config interface{}, status DestinationStatus, createdAt NullableString, updatedAt NullableString, ) *SetDestinationTestMode200ResponseDestination`

NewSetDestinationTestMode200ResponseDestination instantiates a new SetDestinationTestMode200ResponseDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetDestinationTestMode200ResponseDestinationWithDefaults

`func NewSetDestinationTestMode200ResponseDestinationWithDefaults() *SetDestinationTestMode200ResponseDestination`

NewSetDestinationTestMode200ResponseDestinationWithDefaults instantiates a new SetDestinationTestMode200ResponseDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SetDestinationTestMode200ResponseDestination) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SetDestinationTestMode200ResponseDestination) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SetDestinationTestMode200ResponseDestination) SetId(v string)`

SetId sets Id field to given value.


### GetSignalTrackerId

`func (o *SetDestinationTestMode200ResponseDestination) GetSignalTrackerId() string`

GetSignalTrackerId returns the SignalTrackerId field if non-nil, zero value otherwise.

### GetSignalTrackerIdOk

`func (o *SetDestinationTestMode200ResponseDestination) GetSignalTrackerIdOk() (*string, bool)`

GetSignalTrackerIdOk returns a tuple with the SignalTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalTrackerId

`func (o *SetDestinationTestMode200ResponseDestination) SetSignalTrackerId(v string)`

SetSignalTrackerId sets SignalTrackerId field to given value.


### GetPlatformAdAccountId

`func (o *SetDestinationTestMode200ResponseDestination) GetPlatformAdAccountId() int32`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *SetDestinationTestMode200ResponseDestination) GetPlatformAdAccountIdOk() (*int32, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *SetDestinationTestMode200ResponseDestination) SetPlatformAdAccountId(v int32)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.


### SetPlatformAdAccountIdNil

`func (o *SetDestinationTestMode200ResponseDestination) SetPlatformAdAccountIdNil(b bool)`

 SetPlatformAdAccountIdNil sets the value for PlatformAdAccountId to be an explicit nil

### UnsetPlatformAdAccountId
`func (o *SetDestinationTestMode200ResponseDestination) UnsetPlatformAdAccountId()`

UnsetPlatformAdAccountId ensures that no value is present for PlatformAdAccountId, not even an explicit nil
### GetType

`func (o *SetDestinationTestMode200ResponseDestination) GetType() DestinationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SetDestinationTestMode200ResponseDestination) GetTypeOk() (*DestinationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SetDestinationTestMode200ResponseDestination) SetType(v DestinationType)`

SetType sets Type field to given value.


### GetCredentialSource

`func (o *SetDestinationTestMode200ResponseDestination) GetCredentialSource() DestinationCredentialSource`

GetCredentialSource returns the CredentialSource field if non-nil, zero value otherwise.

### GetCredentialSourceOk

`func (o *SetDestinationTestMode200ResponseDestination) GetCredentialSourceOk() (*DestinationCredentialSource, bool)`

GetCredentialSourceOk returns a tuple with the CredentialSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialSource

`func (o *SetDestinationTestMode200ResponseDestination) SetCredentialSource(v DestinationCredentialSource)`

SetCredentialSource sets CredentialSource field to given value.


### GetConfig

`func (o *SetDestinationTestMode200ResponseDestination) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SetDestinationTestMode200ResponseDestination) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SetDestinationTestMode200ResponseDestination) SetConfig(v interface{})`

SetConfig sets Config field to given value.


### SetConfigNil

`func (o *SetDestinationTestMode200ResponseDestination) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *SetDestinationTestMode200ResponseDestination) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetStatus

`func (o *SetDestinationTestMode200ResponseDestination) GetStatus() DestinationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SetDestinationTestMode200ResponseDestination) GetStatusOk() (*DestinationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SetDestinationTestMode200ResponseDestination) SetStatus(v DestinationStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *SetDestinationTestMode200ResponseDestination) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SetDestinationTestMode200ResponseDestination) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SetDestinationTestMode200ResponseDestination) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *SetDestinationTestMode200ResponseDestination) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *SetDestinationTestMode200ResponseDestination) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *SetDestinationTestMode200ResponseDestination) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SetDestinationTestMode200ResponseDestination) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SetDestinationTestMode200ResponseDestination) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *SetDestinationTestMode200ResponseDestination) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *SetDestinationTestMode200ResponseDestination) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


