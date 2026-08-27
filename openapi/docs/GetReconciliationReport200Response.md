# GetReconciliationReport200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**Reports** | [**[]GetReconciliationReport200ResponseReportsInner**](GetReconciliationReport200ResponseReportsInner.md) |  | 

## Methods

### NewGetReconciliationReport200Response

`func NewGetReconciliationReport200Response(date string, reports []GetReconciliationReport200ResponseReportsInner, ) *GetReconciliationReport200Response`

NewGetReconciliationReport200Response instantiates a new GetReconciliationReport200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReconciliationReport200ResponseWithDefaults

`func NewGetReconciliationReport200ResponseWithDefaults() *GetReconciliationReport200Response`

NewGetReconciliationReport200ResponseWithDefaults instantiates a new GetReconciliationReport200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetReconciliationReport200Response) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetReconciliationReport200Response) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetReconciliationReport200Response) SetDate(v string)`

SetDate sets Date field to given value.


### GetReports

`func (o *GetReconciliationReport200Response) GetReports() []GetReconciliationReport200ResponseReportsInner`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *GetReconciliationReport200Response) GetReportsOk() (*[]GetReconciliationReport200ResponseReportsInner, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *GetReconciliationReport200Response) SetReports(v []GetReconciliationReport200ResponseReportsInner)`

SetReports sets Reports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


