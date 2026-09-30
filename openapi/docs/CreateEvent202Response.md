# CreateEvent202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**Duplicate** | **bool** |  | 
**Warnings** | [**[]CreateEvent202ResponseWarningsInner**](CreateEvent202ResponseWarningsInner.md) |  | 

## Methods

### NewCreateEvent202Response

`func NewCreateEvent202Response(eventId string, duplicate bool, warnings []CreateEvent202ResponseWarningsInner, ) *CreateEvent202Response`

NewCreateEvent202Response instantiates a new CreateEvent202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEvent202ResponseWithDefaults

`func NewCreateEvent202ResponseWithDefaults() *CreateEvent202Response`

NewCreateEvent202ResponseWithDefaults instantiates a new CreateEvent202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *CreateEvent202Response) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *CreateEvent202Response) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *CreateEvent202Response) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetDuplicate

`func (o *CreateEvent202Response) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *CreateEvent202Response) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *CreateEvent202Response) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.


### GetWarnings

`func (o *CreateEvent202Response) GetWarnings() []CreateEvent202ResponseWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *CreateEvent202Response) GetWarningsOk() (*[]CreateEvent202ResponseWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *CreateEvent202Response) SetWarnings(v []CreateEvent202ResponseWarningsInner)`

SetWarnings sets Warnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


