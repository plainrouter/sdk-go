# GetEvent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**GetEvent200ResponseEvent**](GetEvent200ResponseEvent.md) |  | 
**Lineage** | [**GetEvent200ResponseLineage**](GetEvent200ResponseLineage.md) |  | 
**Deliveries** | [**[]GetEvent200ResponseDeliveriesInner**](GetEvent200ResponseDeliveriesInner.md) |  | 

## Methods

### NewGetEvent200Response

`func NewGetEvent200Response(event GetEvent200ResponseEvent, lineage GetEvent200ResponseLineage, deliveries []GetEvent200ResponseDeliveriesInner, ) *GetEvent200Response`

NewGetEvent200Response instantiates a new GetEvent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvent200ResponseWithDefaults

`func NewGetEvent200ResponseWithDefaults() *GetEvent200Response`

NewGetEvent200ResponseWithDefaults instantiates a new GetEvent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *GetEvent200Response) GetEvent() GetEvent200ResponseEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *GetEvent200Response) GetEventOk() (*GetEvent200ResponseEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *GetEvent200Response) SetEvent(v GetEvent200ResponseEvent)`

SetEvent sets Event field to given value.


### GetLineage

`func (o *GetEvent200Response) GetLineage() GetEvent200ResponseLineage`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *GetEvent200Response) GetLineageOk() (*GetEvent200ResponseLineage, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *GetEvent200Response) SetLineage(v GetEvent200ResponseLineage)`

SetLineage sets Lineage field to given value.


### GetDeliveries

`func (o *GetEvent200Response) GetDeliveries() []GetEvent200ResponseDeliveriesInner`

GetDeliveries returns the Deliveries field if non-nil, zero value otherwise.

### GetDeliveriesOk

`func (o *GetEvent200Response) GetDeliveriesOk() (*[]GetEvent200ResponseDeliveriesInner, bool)`

GetDeliveriesOk returns a tuple with the Deliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveries

`func (o *GetEvent200Response) SetDeliveries(v []GetEvent200ResponseDeliveriesInner)`

SetDeliveries sets Deliveries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


