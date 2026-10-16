# AgentContext200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaConnected** | **bool** |  | 
**Agent** | [**AgentContext200ResponseAgent**](AgentContext200ResponseAgent.md) |  | 
**PlatformAdAccount** | [**NullableAgentContext200ResponsePlatformAdAccount**](AgentContext200ResponsePlatformAdAccount.md) |  | 
**NextSteps** | Pointer to **[]string** | Present only when meta_connected is false. | [optional] 

## Methods

### NewAgentContext200Response

`func NewAgentContext200Response(metaConnected bool, agent AgentContext200ResponseAgent, platformAdAccount NullableAgentContext200ResponsePlatformAdAccount, ) *AgentContext200Response`

NewAgentContext200Response instantiates a new AgentContext200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentContext200ResponseWithDefaults

`func NewAgentContext200ResponseWithDefaults() *AgentContext200Response`

NewAgentContext200ResponseWithDefaults instantiates a new AgentContext200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaConnected

`func (o *AgentContext200Response) GetMetaConnected() bool`

GetMetaConnected returns the MetaConnected field if non-nil, zero value otherwise.

### GetMetaConnectedOk

`func (o *AgentContext200Response) GetMetaConnectedOk() (*bool, bool)`

GetMetaConnectedOk returns a tuple with the MetaConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaConnected

`func (o *AgentContext200Response) SetMetaConnected(v bool)`

SetMetaConnected sets MetaConnected field to given value.


### GetAgent

`func (o *AgentContext200Response) GetAgent() AgentContext200ResponseAgent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentContext200Response) GetAgentOk() (*AgentContext200ResponseAgent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentContext200Response) SetAgent(v AgentContext200ResponseAgent)`

SetAgent sets Agent field to given value.


### GetPlatformAdAccount

`func (o *AgentContext200Response) GetPlatformAdAccount() AgentContext200ResponsePlatformAdAccount`

GetPlatformAdAccount returns the PlatformAdAccount field if non-nil, zero value otherwise.

### GetPlatformAdAccountOk

`func (o *AgentContext200Response) GetPlatformAdAccountOk() (*AgentContext200ResponsePlatformAdAccount, bool)`

GetPlatformAdAccountOk returns a tuple with the PlatformAdAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccount

`func (o *AgentContext200Response) SetPlatformAdAccount(v AgentContext200ResponsePlatformAdAccount)`

SetPlatformAdAccount sets PlatformAdAccount field to given value.


### SetPlatformAdAccountNil

`func (o *AgentContext200Response) SetPlatformAdAccountNil(b bool)`

 SetPlatformAdAccountNil sets the value for PlatformAdAccount to be an explicit nil

### UnsetPlatformAdAccount
`func (o *AgentContext200Response) UnsetPlatformAdAccount()`

UnsetPlatformAdAccount ensures that no value is present for PlatformAdAccount, not even an explicit nil
### GetNextSteps

`func (o *AgentContext200Response) GetNextSteps() []string`

GetNextSteps returns the NextSteps field if non-nil, zero value otherwise.

### GetNextStepsOk

`func (o *AgentContext200Response) GetNextStepsOk() (*[]string, bool)`

GetNextStepsOk returns a tuple with the NextSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSteps

`func (o *AgentContext200Response) SetNextSteps(v []string)`

SetNextSteps sets NextSteps field to given value.

### HasNextSteps

`func (o *AgentContext200Response) HasNextSteps() bool`

HasNextSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


