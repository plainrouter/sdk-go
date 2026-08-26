# ListEvents200ResponseMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepted** | **int32** |  | 
**TotalDeliveries** | **int32** |  | 
**AcceptanceRate** | **NullableFloat32** |  | 
**AcceptanceRateWindow** | **string** |  | 

## Methods

### NewListEvents200ResponseMetrics

`func NewListEvents200ResponseMetrics(accepted int32, totalDeliveries int32, acceptanceRate NullableFloat32, acceptanceRateWindow string, ) *ListEvents200ResponseMetrics`

NewListEvents200ResponseMetrics instantiates a new ListEvents200ResponseMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEvents200ResponseMetricsWithDefaults

`func NewListEvents200ResponseMetricsWithDefaults() *ListEvents200ResponseMetrics`

NewListEvents200ResponseMetricsWithDefaults instantiates a new ListEvents200ResponseMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepted

`func (o *ListEvents200ResponseMetrics) GetAccepted() int32`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *ListEvents200ResponseMetrics) GetAcceptedOk() (*int32, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *ListEvents200ResponseMetrics) SetAccepted(v int32)`

SetAccepted sets Accepted field to given value.


### GetTotalDeliveries

`func (o *ListEvents200ResponseMetrics) GetTotalDeliveries() int32`

GetTotalDeliveries returns the TotalDeliveries field if non-nil, zero value otherwise.

### GetTotalDeliveriesOk

`func (o *ListEvents200ResponseMetrics) GetTotalDeliveriesOk() (*int32, bool)`

GetTotalDeliveriesOk returns a tuple with the TotalDeliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDeliveries

`func (o *ListEvents200ResponseMetrics) SetTotalDeliveries(v int32)`

SetTotalDeliveries sets TotalDeliveries field to given value.


### GetAcceptanceRate

`func (o *ListEvents200ResponseMetrics) GetAcceptanceRate() float32`

GetAcceptanceRate returns the AcceptanceRate field if non-nil, zero value otherwise.

### GetAcceptanceRateOk

`func (o *ListEvents200ResponseMetrics) GetAcceptanceRateOk() (*float32, bool)`

GetAcceptanceRateOk returns a tuple with the AcceptanceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceRate

`func (o *ListEvents200ResponseMetrics) SetAcceptanceRate(v float32)`

SetAcceptanceRate sets AcceptanceRate field to given value.


### SetAcceptanceRateNil

`func (o *ListEvents200ResponseMetrics) SetAcceptanceRateNil(b bool)`

 SetAcceptanceRateNil sets the value for AcceptanceRate to be an explicit nil

### UnsetAcceptanceRate
`func (o *ListEvents200ResponseMetrics) UnsetAcceptanceRate()`

UnsetAcceptanceRate ensures that no value is present for AcceptanceRate, not even an explicit nil
### GetAcceptanceRateWindow

`func (o *ListEvents200ResponseMetrics) GetAcceptanceRateWindow() string`

GetAcceptanceRateWindow returns the AcceptanceRateWindow field if non-nil, zero value otherwise.

### GetAcceptanceRateWindowOk

`func (o *ListEvents200ResponseMetrics) GetAcceptanceRateWindowOk() (*string, bool)`

GetAcceptanceRateWindowOk returns a tuple with the AcceptanceRateWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceRateWindow

`func (o *ListEvents200ResponseMetrics) SetAcceptanceRateWindow(v string)`

SetAcceptanceRateWindow sets AcceptanceRateWindow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


