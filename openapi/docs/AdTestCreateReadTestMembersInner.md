# AdTestCreateReadTestMembersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdId** | **string** |  | 
**Position** | **int32** |  | 
**BaselineHost** | Pointer to **NullableString** | Immutable normalized landing host once the member has a baseline. | [optional] 
**BaselinePath** | Pointer to **NullableString** | Immutable normalized landing path once the member has a baseline. | [optional] 

## Methods

### NewAdTestCreateReadTestMembersInner

`func NewAdTestCreateReadTestMembersInner(adId string, position int32, ) *AdTestCreateReadTestMembersInner`

NewAdTestCreateReadTestMembersInner instantiates a new AdTestCreateReadTestMembersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdTestCreateReadTestMembersInnerWithDefaults

`func NewAdTestCreateReadTestMembersInnerWithDefaults() *AdTestCreateReadTestMembersInner`

NewAdTestCreateReadTestMembersInnerWithDefaults instantiates a new AdTestCreateReadTestMembersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdId

`func (o *AdTestCreateReadTestMembersInner) GetAdId() string`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *AdTestCreateReadTestMembersInner) GetAdIdOk() (*string, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *AdTestCreateReadTestMembersInner) SetAdId(v string)`

SetAdId sets AdId field to given value.


### GetPosition

`func (o *AdTestCreateReadTestMembersInner) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *AdTestCreateReadTestMembersInner) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *AdTestCreateReadTestMembersInner) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetBaselineHost

`func (o *AdTestCreateReadTestMembersInner) GetBaselineHost() string`

GetBaselineHost returns the BaselineHost field if non-nil, zero value otherwise.

### GetBaselineHostOk

`func (o *AdTestCreateReadTestMembersInner) GetBaselineHostOk() (*string, bool)`

GetBaselineHostOk returns a tuple with the BaselineHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineHost

`func (o *AdTestCreateReadTestMembersInner) SetBaselineHost(v string)`

SetBaselineHost sets BaselineHost field to given value.

### HasBaselineHost

`func (o *AdTestCreateReadTestMembersInner) HasBaselineHost() bool`

HasBaselineHost returns a boolean if a field has been set.

### SetBaselineHostNil

`func (o *AdTestCreateReadTestMembersInner) SetBaselineHostNil(b bool)`

 SetBaselineHostNil sets the value for BaselineHost to be an explicit nil

### UnsetBaselineHost
`func (o *AdTestCreateReadTestMembersInner) UnsetBaselineHost()`

UnsetBaselineHost ensures that no value is present for BaselineHost, not even an explicit nil
### GetBaselinePath

`func (o *AdTestCreateReadTestMembersInner) GetBaselinePath() string`

GetBaselinePath returns the BaselinePath field if non-nil, zero value otherwise.

### GetBaselinePathOk

`func (o *AdTestCreateReadTestMembersInner) GetBaselinePathOk() (*string, bool)`

GetBaselinePathOk returns a tuple with the BaselinePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselinePath

`func (o *AdTestCreateReadTestMembersInner) SetBaselinePath(v string)`

SetBaselinePath sets BaselinePath field to given value.

### HasBaselinePath

`func (o *AdTestCreateReadTestMembersInner) HasBaselinePath() bool`

HasBaselinePath returns a boolean if a field has been set.

### SetBaselinePathNil

`func (o *AdTestCreateReadTestMembersInner) SetBaselinePathNil(b bool)`

 SetBaselinePathNil sets the value for BaselinePath to be an explicit nil

### UnsetBaselinePath
`func (o *AdTestCreateReadTestMembersInner) UnsetBaselinePath()`

UnsetBaselinePath ensures that no value is present for BaselinePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


