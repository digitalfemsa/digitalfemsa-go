# CustomerPaymentMethodsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of payment method | 
**ExpiresAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewCustomerPaymentMethodsRequest

`func NewCustomerPaymentMethodsRequest(type_ string, ) *CustomerPaymentMethodsRequest`

NewCustomerPaymentMethodsRequest instantiates a new CustomerPaymentMethodsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPaymentMethodsRequestWithDefaults

`func NewCustomerPaymentMethodsRequestWithDefaults() *CustomerPaymentMethodsRequest`

NewCustomerPaymentMethodsRequestWithDefaults instantiates a new CustomerPaymentMethodsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerPaymentMethodsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerPaymentMethodsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerPaymentMethodsRequest) SetType(v string)`

SetType sets Type field to given value.


### GetExpiresAt

`func (o *CustomerPaymentMethodsRequest) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CustomerPaymentMethodsRequest) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CustomerPaymentMethodsRequest) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CustomerPaymentMethodsRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


