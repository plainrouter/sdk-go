# ActionListRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ActionReadItem**](ActionReadItem.md) |  | 
**Meta** | [**ActionListReadMeta**](ActionListReadMeta.md) |  | 

## Methods

### NewActionListRead

`func NewActionListRead(data []ActionReadItem, meta ActionListReadMeta, ) *ActionListRead`

NewActionListRead instantiates a new ActionListRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionListReadWithDefaults

`func NewActionListReadWithDefaults() *ActionListRead`

NewActionListReadWithDefaults instantiates a new ActionListRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ActionListRead) GetData() []ActionReadItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ActionListRead) GetDataOk() (*[]ActionReadItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ActionListRead) SetData(v []ActionReadItem)`

SetData sets Data field to given value.


### GetMeta

`func (o *ActionListRead) GetMeta() ActionListReadMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ActionListRead) GetMetaOk() (*ActionListReadMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ActionListRead) SetMeta(v ActionListReadMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


