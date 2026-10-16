# CookieDeclarationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ProviderDomain** | **string** |  | 
**Party** | **string** |  | 
**Expiry** | **string** | Observed UTC expiry date, Session, or Until removed. | 
**Purpose** | **NullableString** | Purpose from the Open Cookie Database, where available. | 

## Methods

### NewCookieDeclarationItem

`func NewCookieDeclarationItem(name string, providerDomain string, party string, expiry string, purpose NullableString, ) *CookieDeclarationItem`

NewCookieDeclarationItem instantiates a new CookieDeclarationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCookieDeclarationItemWithDefaults

`func NewCookieDeclarationItemWithDefaults() *CookieDeclarationItem`

NewCookieDeclarationItemWithDefaults instantiates a new CookieDeclarationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CookieDeclarationItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CookieDeclarationItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CookieDeclarationItem) SetName(v string)`

SetName sets Name field to given value.


### GetProviderDomain

`func (o *CookieDeclarationItem) GetProviderDomain() string`

GetProviderDomain returns the ProviderDomain field if non-nil, zero value otherwise.

### GetProviderDomainOk

`func (o *CookieDeclarationItem) GetProviderDomainOk() (*string, bool)`

GetProviderDomainOk returns a tuple with the ProviderDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderDomain

`func (o *CookieDeclarationItem) SetProviderDomain(v string)`

SetProviderDomain sets ProviderDomain field to given value.


### GetParty

`func (o *CookieDeclarationItem) GetParty() string`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *CookieDeclarationItem) GetPartyOk() (*string, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *CookieDeclarationItem) SetParty(v string)`

SetParty sets Party field to given value.


### GetExpiry

`func (o *CookieDeclarationItem) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *CookieDeclarationItem) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *CookieDeclarationItem) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.


### GetPurpose

`func (o *CookieDeclarationItem) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *CookieDeclarationItem) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *CookieDeclarationItem) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### SetPurposeNil

`func (o *CookieDeclarationItem) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *CookieDeclarationItem) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


