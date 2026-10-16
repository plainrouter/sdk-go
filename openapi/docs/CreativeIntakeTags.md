# CreativeIntakeTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Generator** | Pointer to **string** | Lowercase slug naming the outside tool that made the creative. | [optional] 
**GeneratorJob** | Pointer to **string** | Printable characters naming the tool’s job or render; control and other non-printable characters are refused. | [optional] 
**ParentCreative** | Pointer to **string** | The id returned by the creative APIs: a lowercase ULID of an existing creative in the same workspace that this creative varies. | [optional] 

## Methods

### NewCreativeIntakeTags

`func NewCreativeIntakeTags() *CreativeIntakeTags`

NewCreativeIntakeTags instantiates a new CreativeIntakeTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreativeIntakeTagsWithDefaults

`func NewCreativeIntakeTagsWithDefaults() *CreativeIntakeTags`

NewCreativeIntakeTagsWithDefaults instantiates a new CreativeIntakeTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenerator

`func (o *CreativeIntakeTags) GetGenerator() string`

GetGenerator returns the Generator field if non-nil, zero value otherwise.

### GetGeneratorOk

`func (o *CreativeIntakeTags) GetGeneratorOk() (*string, bool)`

GetGeneratorOk returns a tuple with the Generator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerator

`func (o *CreativeIntakeTags) SetGenerator(v string)`

SetGenerator sets Generator field to given value.

### HasGenerator

`func (o *CreativeIntakeTags) HasGenerator() bool`

HasGenerator returns a boolean if a field has been set.

### GetGeneratorJob

`func (o *CreativeIntakeTags) GetGeneratorJob() string`

GetGeneratorJob returns the GeneratorJob field if non-nil, zero value otherwise.

### GetGeneratorJobOk

`func (o *CreativeIntakeTags) GetGeneratorJobOk() (*string, bool)`

GetGeneratorJobOk returns a tuple with the GeneratorJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorJob

`func (o *CreativeIntakeTags) SetGeneratorJob(v string)`

SetGeneratorJob sets GeneratorJob field to given value.

### HasGeneratorJob

`func (o *CreativeIntakeTags) HasGeneratorJob() bool`

HasGeneratorJob returns a boolean if a field has been set.

### GetParentCreative

`func (o *CreativeIntakeTags) GetParentCreative() string`

GetParentCreative returns the ParentCreative field if non-nil, zero value otherwise.

### GetParentCreativeOk

`func (o *CreativeIntakeTags) GetParentCreativeOk() (*string, bool)`

GetParentCreativeOk returns a tuple with the ParentCreative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCreative

`func (o *CreativeIntakeTags) SetParentCreative(v string)`

SetParentCreative sets ParentCreative field to given value.

### HasParentCreative

`func (o *CreativeIntakeTags) HasParentCreative() bool`

HasParentCreative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


