# ActionRuleDetailReadRuleCurrentRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TargetLevel** | **string** |  | 
**Action** | **string** |  | 
**BudgetChangePercent** | **NullableInt32** |  | 
**MinimumCountedArrivals** | **int32** |  | 
**CostComparison** | **string** |  | 
**CostRatio** | **string** |  | 
**AuthorType** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewActionRuleDetailReadRuleCurrentRevision

`func NewActionRuleDetailReadRuleCurrentRevision(id string, targetLevel string, action string, budgetChangePercent NullableInt32, minimumCountedArrivals int32, costComparison string, costRatio string, authorType string, createdAt time.Time, ) *ActionRuleDetailReadRuleCurrentRevision`

NewActionRuleDetailReadRuleCurrentRevision instantiates a new ActionRuleDetailReadRuleCurrentRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionRuleDetailReadRuleCurrentRevisionWithDefaults

`func NewActionRuleDetailReadRuleCurrentRevisionWithDefaults() *ActionRuleDetailReadRuleCurrentRevision`

NewActionRuleDetailReadRuleCurrentRevisionWithDefaults instantiates a new ActionRuleDetailReadRuleCurrentRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionRuleDetailReadRuleCurrentRevision) SetId(v string)`

SetId sets Id field to given value.


### GetTargetLevel

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetTargetLevel() string`

GetTargetLevel returns the TargetLevel field if non-nil, zero value otherwise.

### GetTargetLevelOk

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetTargetLevelOk() (*string, bool)`

GetTargetLevelOk returns a tuple with the TargetLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLevel

`func (o *ActionRuleDetailReadRuleCurrentRevision) SetTargetLevel(v string)`

SetTargetLevel sets TargetLevel field to given value.


### GetAction

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ActionRuleDetailReadRuleCurrentRevision) SetAction(v string)`

SetAction sets Action field to given value.


### GetBudgetChangePercent

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetBudgetChangePercent() int32`

GetBudgetChangePercent returns the BudgetChangePercent field if non-nil, zero value otherwise.

### GetBudgetChangePercentOk

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetBudgetChangePercentOk() (*int32, bool)`

GetBudgetChangePercentOk returns a tuple with the BudgetChangePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetChangePercent

`func (o *ActionRuleDetailReadRuleCurrentRevision) SetBudgetChangePercent(v int32)`

SetBudgetChangePercent sets BudgetChangePercent field to given value.


### SetBudgetChangePercentNil

`func (o *ActionRuleDetailReadRuleCurrentRevision) SetBudgetChangePercentNil(b bool)`

 SetBudgetChangePercentNil sets the value for BudgetChangePercent to be an explicit nil

### UnsetBudgetChangePercent
`func (o *ActionRuleDetailReadRuleCurrentRevision) UnsetBudgetChangePercent()`

UnsetBudgetChangePercent ensures that no value is present for BudgetChangePercent, not even an explicit nil
### GetMinimumCountedArrivals

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetMinimumCountedArrivals() int32`

GetMinimumCountedArrivals returns the MinimumCountedArrivals field if non-nil, zero value otherwise.

### GetMinimumCountedArrivalsOk

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetMinimumCountedArrivalsOk() (*int32, bool)`

GetMinimumCountedArrivalsOk returns a tuple with the MinimumCountedArrivals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCountedArrivals

`func (o *ActionRuleDetailReadRuleCurrentRevision) SetMinimumCountedArrivals(v int32)`

SetMinimumCountedArrivals sets MinimumCountedArrivals field to given value.


### GetCostComparison

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetCostComparison() string`

GetCostComparison returns the CostComparison field if non-nil, zero value otherwise.

### GetCostComparisonOk

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetCostComparisonOk() (*string, bool)`

GetCostComparisonOk returns a tuple with the CostComparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostComparison

`func (o *ActionRuleDetailReadRuleCurrentRevision) SetCostComparison(v string)`

SetCostComparison sets CostComparison field to given value.


### GetCostRatio

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetCostRatio() string`

GetCostRatio returns the CostRatio field if non-nil, zero value otherwise.

### GetCostRatioOk

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetCostRatioOk() (*string, bool)`

GetCostRatioOk returns a tuple with the CostRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostRatio

`func (o *ActionRuleDetailReadRuleCurrentRevision) SetCostRatio(v string)`

SetCostRatio sets CostRatio field to given value.


### GetAuthorType

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetAuthorType() string`

GetAuthorType returns the AuthorType field if non-nil, zero value otherwise.

### GetAuthorTypeOk

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetAuthorTypeOk() (*string, bool)`

GetAuthorTypeOk returns a tuple with the AuthorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorType

`func (o *ActionRuleDetailReadRuleCurrentRevision) SetAuthorType(v string)`

SetAuthorType sets AuthorType field to given value.


### GetCreatedAt

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ActionRuleDetailReadRuleCurrentRevision) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ActionRuleDetailReadRuleCurrentRevision) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


