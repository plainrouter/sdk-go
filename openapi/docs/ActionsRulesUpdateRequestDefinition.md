# ActionsRulesUpdateRequestDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetLevel** | **string** |  | 
**Action** | **string** |  | 
**MinimumCountedArrivals** | **int32** |  | 
**CostComparison** | **string** |  | 
**CostRatio** | [**ActionsRulesStoreRequestDefinitionAnyOfCostRatio**](ActionsRulesStoreRequestDefinitionAnyOfCostRatio.md) |  | 
**BudgetChangePercent** | **int32** |  | 

## Methods

### NewActionsRulesUpdateRequestDefinition

`func NewActionsRulesUpdateRequestDefinition(targetLevel string, action string, minimumCountedArrivals int32, costComparison string, costRatio ActionsRulesStoreRequestDefinitionAnyOfCostRatio, budgetChangePercent int32, ) *ActionsRulesUpdateRequestDefinition`

NewActionsRulesUpdateRequestDefinition instantiates a new ActionsRulesUpdateRequestDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsRulesUpdateRequestDefinitionWithDefaults

`func NewActionsRulesUpdateRequestDefinitionWithDefaults() *ActionsRulesUpdateRequestDefinition`

NewActionsRulesUpdateRequestDefinitionWithDefaults instantiates a new ActionsRulesUpdateRequestDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetLevel

`func (o *ActionsRulesUpdateRequestDefinition) GetTargetLevel() string`

GetTargetLevel returns the TargetLevel field if non-nil, zero value otherwise.

### GetTargetLevelOk

`func (o *ActionsRulesUpdateRequestDefinition) GetTargetLevelOk() (*string, bool)`

GetTargetLevelOk returns a tuple with the TargetLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLevel

`func (o *ActionsRulesUpdateRequestDefinition) SetTargetLevel(v string)`

SetTargetLevel sets TargetLevel field to given value.


### GetAction

`func (o *ActionsRulesUpdateRequestDefinition) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ActionsRulesUpdateRequestDefinition) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ActionsRulesUpdateRequestDefinition) SetAction(v string)`

SetAction sets Action field to given value.


### GetMinimumCountedArrivals

`func (o *ActionsRulesUpdateRequestDefinition) GetMinimumCountedArrivals() int32`

GetMinimumCountedArrivals returns the MinimumCountedArrivals field if non-nil, zero value otherwise.

### GetMinimumCountedArrivalsOk

`func (o *ActionsRulesUpdateRequestDefinition) GetMinimumCountedArrivalsOk() (*int32, bool)`

GetMinimumCountedArrivalsOk returns a tuple with the MinimumCountedArrivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCountedArrivals

`func (o *ActionsRulesUpdateRequestDefinition) SetMinimumCountedArrivals(v int32)`

SetMinimumCountedArrivals sets MinimumCountedArrivals field to given value.


### GetCostComparison

`func (o *ActionsRulesUpdateRequestDefinition) GetCostComparison() string`

GetCostComparison returns the CostComparison field if non-nil, zero value otherwise.

### GetCostComparisonOk

`func (o *ActionsRulesUpdateRequestDefinition) GetCostComparisonOk() (*string, bool)`

GetCostComparisonOk returns a tuple with the CostComparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostComparison

`func (o *ActionsRulesUpdateRequestDefinition) SetCostComparison(v string)`

SetCostComparison sets CostComparison field to given value.


### GetCostRatio

`func (o *ActionsRulesUpdateRequestDefinition) GetCostRatio() ActionsRulesStoreRequestDefinitionAnyOfCostRatio`

GetCostRatio returns the CostRatio field if non-nil, zero value otherwise.

### GetCostRatioOk

`func (o *ActionsRulesUpdateRequestDefinition) GetCostRatioOk() (*ActionsRulesStoreRequestDefinitionAnyOfCostRatio, bool)`

GetCostRatioOk returns a tuple with the CostRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostRatio

`func (o *ActionsRulesUpdateRequestDefinition) SetCostRatio(v ActionsRulesStoreRequestDefinitionAnyOfCostRatio)`

SetCostRatio sets CostRatio field to given value.


### GetBudgetChangePercent

`func (o *ActionsRulesUpdateRequestDefinition) GetBudgetChangePercent() int32`

GetBudgetChangePercent returns the BudgetChangePercent field if non-nil, zero value otherwise.

### GetBudgetChangePercentOk

`func (o *ActionsRulesUpdateRequestDefinition) GetBudgetChangePercentOk() (*int32, bool)`

GetBudgetChangePercentOk returns a tuple with the BudgetChangePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetChangePercent

`func (o *ActionsRulesUpdateRequestDefinition) SetBudgetChangePercent(v int32)`

SetBudgetChangePercent sets BudgetChangePercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


