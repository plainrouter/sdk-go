# GetSandbox200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **string** |  | 
**AuthenticationRequired** | **bool** |  | 
**AccountRequired** | **bool** |  | 
**ProductionData** | **bool** |  | 
**PersistsData** | **bool** |  | 
**ProviderDelivery** | **bool** |  | 
**Description** | **string** |  | 
**SelfServeKey** | [**GetSandbox200ResponseSelfServeKey**](GetSandbox200ResponseSelfServeKey.md) |  | 
**Try** | [**GetSandbox200ResponseTry**](GetSandbox200ResponseTry.md) |  | 

## Methods

### NewGetSandbox200Response

`func NewGetSandbox200Response(environment string, authenticationRequired bool, accountRequired bool, productionData bool, persistsData bool, providerDelivery bool, description string, selfServeKey GetSandbox200ResponseSelfServeKey, try GetSandbox200ResponseTry, ) *GetSandbox200Response`

NewGetSandbox200Response instantiates a new GetSandbox200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSandbox200ResponseWithDefaults

`func NewGetSandbox200ResponseWithDefaults() *GetSandbox200Response`

NewGetSandbox200ResponseWithDefaults instantiates a new GetSandbox200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *GetSandbox200Response) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GetSandbox200Response) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GetSandbox200Response) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetAuthenticationRequired

`func (o *GetSandbox200Response) GetAuthenticationRequired() bool`

GetAuthenticationRequired returns the AuthenticationRequired field if non-nil, zero value otherwise.

### GetAuthenticationRequiredOk

`func (o *GetSandbox200Response) GetAuthenticationRequiredOk() (*bool, bool)`

GetAuthenticationRequiredOk returns a tuple with the AuthenticationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationRequired

`func (o *GetSandbox200Response) SetAuthenticationRequired(v bool)`

SetAuthenticationRequired sets AuthenticationRequired field to given value.


### GetAccountRequired

`func (o *GetSandbox200Response) GetAccountRequired() bool`

GetAccountRequired returns the AccountRequired field if non-nil, zero value otherwise.

### GetAccountRequiredOk

`func (o *GetSandbox200Response) GetAccountRequiredOk() (*bool, bool)`

GetAccountRequiredOk returns a tuple with the AccountRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountRequired

`func (o *GetSandbox200Response) SetAccountRequired(v bool)`

SetAccountRequired sets AccountRequired field to given value.


### GetProductionData

`func (o *GetSandbox200Response) GetProductionData() bool`

GetProductionData returns the ProductionData field if non-nil, zero value otherwise.

### GetProductionDataOk

`func (o *GetSandbox200Response) GetProductionDataOk() (*bool, bool)`

GetProductionDataOk returns a tuple with the ProductionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionData

`func (o *GetSandbox200Response) SetProductionData(v bool)`

SetProductionData sets ProductionData field to given value.


### GetPersistsData

`func (o *GetSandbox200Response) GetPersistsData() bool`

GetPersistsData returns the PersistsData field if non-nil, zero value otherwise.

### GetPersistsDataOk

`func (o *GetSandbox200Response) GetPersistsDataOk() (*bool, bool)`

GetPersistsDataOk returns a tuple with the PersistsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistsData

`func (o *GetSandbox200Response) SetPersistsData(v bool)`

SetPersistsData sets PersistsData field to given value.


### GetProviderDelivery

`func (o *GetSandbox200Response) GetProviderDelivery() bool`

GetProviderDelivery returns the ProviderDelivery field if non-nil, zero value otherwise.

### GetProviderDeliveryOk

`func (o *GetSandbox200Response) GetProviderDeliveryOk() (*bool, bool)`

GetProviderDeliveryOk returns a tuple with the ProviderDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderDelivery

`func (o *GetSandbox200Response) SetProviderDelivery(v bool)`

SetProviderDelivery sets ProviderDelivery field to given value.


### GetDescription

`func (o *GetSandbox200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetSandbox200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetSandbox200Response) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSelfServeKey

`func (o *GetSandbox200Response) GetSelfServeKey() GetSandbox200ResponseSelfServeKey`

GetSelfServeKey returns the SelfServeKey field if non-nil, zero value otherwise.

### GetSelfServeKeyOk

`func (o *GetSandbox200Response) GetSelfServeKeyOk() (*GetSandbox200ResponseSelfServeKey, bool)`

GetSelfServeKeyOk returns a tuple with the SelfServeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServeKey

`func (o *GetSandbox200Response) SetSelfServeKey(v GetSandbox200ResponseSelfServeKey)`

SetSelfServeKey sets SelfServeKey field to given value.


### GetTry

`func (o *GetSandbox200Response) GetTry() GetSandbox200ResponseTry`

GetTry returns the Try field if non-nil, zero value otherwise.

### GetTryOk

`func (o *GetSandbox200Response) GetTryOk() (*GetSandbox200ResponseTry, bool)`

GetTryOk returns a tuple with the Try field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTry

`func (o *GetSandbox200Response) SetTry(v GetSandbox200ResponseTry)`

SetTry sets Try field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


