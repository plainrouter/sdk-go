# DeleteUserData200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletionRequestId** | **string** |  | 
**Duplicate** | **bool** |  | 
**EventsUpdated** | **int32** |  | 
**SessionsUpdated** | **int32** |  | 
**CompletedAt** | **string** |  | 

## Methods

### NewDeleteUserData200Response

`func NewDeleteUserData200Response(deletionRequestId string, duplicate bool, eventsUpdated int32, sessionsUpdated int32, completedAt string, ) *DeleteUserData200Response`

NewDeleteUserData200Response instantiates a new DeleteUserData200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteUserData200ResponseWithDefaults

`func NewDeleteUserData200ResponseWithDefaults() *DeleteUserData200Response`

NewDeleteUserData200ResponseWithDefaults instantiates a new DeleteUserData200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletionRequestId

`func (o *DeleteUserData200Response) GetDeletionRequestId() string`

GetDeletionRequestId returns the DeletionRequestId field if non-nil, zero value otherwise.

### GetDeletionRequestIdOk

`func (o *DeleteUserData200Response) GetDeletionRequestIdOk() (*string, bool)`

GetDeletionRequestIdOk returns a tuple with the DeletionRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionRequestId

`func (o *DeleteUserData200Response) SetDeletionRequestId(v string)`

SetDeletionRequestId sets DeletionRequestId field to given value.


### GetDuplicate

`func (o *DeleteUserData200Response) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *DeleteUserData200Response) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *DeleteUserData200Response) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.


### GetEventsUpdated

`func (o *DeleteUserData200Response) GetEventsUpdated() int32`

GetEventsUpdated returns the EventsUpdated field if non-nil, zero value otherwise.

### GetEventsUpdatedOk

`func (o *DeleteUserData200Response) GetEventsUpdatedOk() (*int32, bool)`

GetEventsUpdatedOk returns a tuple with the EventsUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsUpdated

`func (o *DeleteUserData200Response) SetEventsUpdated(v int32)`

SetEventsUpdated sets EventsUpdated field to given value.


### GetSessionsUpdated

`func (o *DeleteUserData200Response) GetSessionsUpdated() int32`

GetSessionsUpdated returns the SessionsUpdated field if non-nil, zero value otherwise.

### GetSessionsUpdatedOk

`func (o *DeleteUserData200Response) GetSessionsUpdatedOk() (*int32, bool)`

GetSessionsUpdatedOk returns a tuple with the SessionsUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionsUpdated

`func (o *DeleteUserData200Response) SetSessionsUpdated(v int32)`

SetSessionsUpdated sets SessionsUpdated field to given value.


### GetCompletedAt

`func (o *DeleteUserData200Response) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *DeleteUserData200Response) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *DeleteUserData200Response) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


