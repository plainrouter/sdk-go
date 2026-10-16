# LaunchPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**WorkspaceId** | **int32** |  | 
**Kind** | **string** |  | 
**Name** | **string** |  | 
**Version** | **int32** |  | 
**Body** | **[]interface{}** |  | 
**Status** | **string** |  | 
**CreatedAt** | **NullableTime** |  | 
**UpdatedAt** | **NullableTime** |  | 

## Methods

### NewLaunchPolicy

`func NewLaunchPolicy(id string, workspaceId int32, kind string, name string, version int32, body []interface{}, status string, createdAt NullableTime, updatedAt NullableTime, ) *LaunchPolicy`

NewLaunchPolicy instantiates a new LaunchPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchPolicyWithDefaults

`func NewLaunchPolicyWithDefaults() *LaunchPolicy`

NewLaunchPolicyWithDefaults instantiates a new LaunchPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LaunchPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LaunchPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LaunchPolicy) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *LaunchPolicy) GetWorkspaceId() int32`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *LaunchPolicy) GetWorkspaceIdOk() (*int32, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *LaunchPolicy) SetWorkspaceId(v int32)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetKind

`func (o *LaunchPolicy) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LaunchPolicy) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LaunchPolicy) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *LaunchPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LaunchPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LaunchPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *LaunchPolicy) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LaunchPolicy) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LaunchPolicy) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetBody

`func (o *LaunchPolicy) GetBody() []interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *LaunchPolicy) GetBodyOk() (*[]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *LaunchPolicy) SetBody(v []interface{})`

SetBody sets Body field to given value.


### GetStatus

`func (o *LaunchPolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LaunchPolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LaunchPolicy) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *LaunchPolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LaunchPolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LaunchPolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *LaunchPolicy) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LaunchPolicy) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LaunchPolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LaunchPolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LaunchPolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *LaunchPolicy) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LaunchPolicy) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


