# CookieDeclaration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScanDate** | **NullableTime** | Date of the latest completed scan; null when no scan has completed. Blocked and failed scans do not replace it. | 
**Categories** | [**CookieDeclarationCategories**](CookieDeclarationCategories.md) |  | 

## Methods

### NewCookieDeclaration

`func NewCookieDeclaration(scanDate NullableTime, categories CookieDeclarationCategories, ) *CookieDeclaration`

NewCookieDeclaration instantiates a new CookieDeclaration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCookieDeclarationWithDefaults

`func NewCookieDeclarationWithDefaults() *CookieDeclaration`

NewCookieDeclarationWithDefaults instantiates a new CookieDeclaration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScanDate

`func (o *CookieDeclaration) GetScanDate() time.Time`

GetScanDate returns the ScanDate field if non-nil, zero value otherwise.

### GetScanDateOk

`func (o *CookieDeclaration) GetScanDateOk() (*time.Time, bool)`

GetScanDateOk returns a tuple with the ScanDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanDate

`func (o *CookieDeclaration) SetScanDate(v time.Time)`

SetScanDate sets ScanDate field to given value.


### SetScanDateNil

`func (o *CookieDeclaration) SetScanDateNil(b bool)`

 SetScanDateNil sets the value for ScanDate to be an explicit nil

### UnsetScanDate
`func (o *CookieDeclaration) UnsetScanDate()`

UnsetScanDate ensures that no value is present for ScanDate, not even an explicit nil
### GetCategories

`func (o *CookieDeclaration) GetCategories() CookieDeclarationCategories`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CookieDeclaration) GetCategoriesOk() (*CookieDeclarationCategories, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CookieDeclaration) SetCategories(v CookieDeclarationCategories)`

SetCategories sets Categories field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


