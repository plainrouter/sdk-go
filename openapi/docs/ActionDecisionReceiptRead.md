# ActionDecisionReceiptRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | **map[string]interface{}** |  | 
**DocumentSha256** | **string** |  | 
**ChainEntry** | [**ActionDecisionReceiptReadChainEntry**](ActionDecisionReceiptReadChainEntry.md) |  | 
**Disposition** | [**ActionCurrentDisposition**](ActionCurrentDisposition.md) |  | 

## Methods

### NewActionDecisionReceiptRead

`func NewActionDecisionReceiptRead(document map[string]*interface{}, documentSha256 string, chainEntry ActionDecisionReceiptReadChainEntry, disposition ActionCurrentDisposition, ) *ActionDecisionReceiptRead`

NewActionDecisionReceiptRead instantiates a new ActionDecisionReceiptRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionDecisionReceiptReadWithDefaults

`func NewActionDecisionReceiptReadWithDefaults() *ActionDecisionReceiptRead`

NewActionDecisionReceiptReadWithDefaults instantiates a new ActionDecisionReceiptRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *ActionDecisionReceiptRead) GetDocument() map[string]*interface{}`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *ActionDecisionReceiptRead) GetDocumentOk() (*map[string]*interface{}, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *ActionDecisionReceiptRead) SetDocument(v map[string]*interface{})`

SetDocument sets Document field to given value.


### GetDocumentSha256

`func (o *ActionDecisionReceiptRead) GetDocumentSha256() string`

GetDocumentSha256 returns the DocumentSha256 field if non-nil, zero value otherwise.

### GetDocumentSha256Ok

`func (o *ActionDecisionReceiptRead) GetDocumentSha256Ok() (*string, bool)`

GetDocumentSha256Ok returns a tuple with the DocumentSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSha256

`func (o *ActionDecisionReceiptRead) SetDocumentSha256(v string)`

SetDocumentSha256 sets DocumentSha256 field to given value.


### GetChainEntry

`func (o *ActionDecisionReceiptRead) GetChainEntry() ActionDecisionReceiptReadChainEntry`

GetChainEntry returns the ChainEntry field if non-nil, zero value otherwise.

### GetChainEntryOk

`func (o *ActionDecisionReceiptRead) GetChainEntryOk() (*ActionDecisionReceiptReadChainEntry, bool)`

GetChainEntryOk returns a tuple with the ChainEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainEntry

`func (o *ActionDecisionReceiptRead) SetChainEntry(v ActionDecisionReceiptReadChainEntry)`

SetChainEntry sets ChainEntry field to given value.


### GetDisposition

`func (o *ActionDecisionReceiptRead) GetDisposition() ActionCurrentDisposition`

GetDisposition returns the Disposition field if non-nil, zero value otherwise.

### GetDispositionOk

`func (o *ActionDecisionReceiptRead) GetDispositionOk() (*ActionCurrentDisposition, bool)`

GetDispositionOk returns a tuple with the Disposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposition

`func (o *ActionDecisionReceiptRead) SetDisposition(v ActionCurrentDisposition)`

SetDisposition sets Disposition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


