# CreateEvent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**Duplicate** | **bool** |  | 

## Methods

### NewCreateEvent200Response

`func NewCreateEvent200Response(eventId string, duplicate bool, ) *CreateEvent200Response`

NewCreateEvent200Response instantiates a new CreateEvent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEvent200ResponseWithDefaults

`func NewCreateEvent200ResponseWithDefaults() *CreateEvent200Response`

NewCreateEvent200ResponseWithDefaults instantiates a new CreateEvent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *CreateEvent200Response) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *CreateEvent200Response) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *CreateEvent200Response) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetDuplicate

`func (o *CreateEvent200Response) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *CreateEvent200Response) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *CreateEvent200Response) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


