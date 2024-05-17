# ChargeRequestPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **int64** | Method expiration date as unix timestamp | [optional] 
**Type** | **string** |  | 
**PaymentSourceId** | Pointer to **string** |  | [optional] 

## Methods

### NewChargeRequestPaymentMethod

`func NewChargeRequestPaymentMethod(type_ string, ) *ChargeRequestPaymentMethod`

NewChargeRequestPaymentMethod instantiates a new ChargeRequestPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeRequestPaymentMethodWithDefaults

`func NewChargeRequestPaymentMethodWithDefaults() *ChargeRequestPaymentMethod`

NewChargeRequestPaymentMethodWithDefaults instantiates a new ChargeRequestPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *ChargeRequestPaymentMethod) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ChargeRequestPaymentMethod) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ChargeRequestPaymentMethod) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ChargeRequestPaymentMethod) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetType

`func (o *ChargeRequestPaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChargeRequestPaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChargeRequestPaymentMethod) SetType(v string)`

SetType sets Type field to given value.


### GetPaymentSourceId

`func (o *ChargeRequestPaymentMethod) GetPaymentSourceId() string`

GetPaymentSourceId returns the PaymentSourceId field if non-nil, zero value otherwise.

### GetPaymentSourceIdOk

`func (o *ChargeRequestPaymentMethod) GetPaymentSourceIdOk() (*string, bool)`

GetPaymentSourceIdOk returns a tuple with the PaymentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSourceId

`func (o *ChargeRequestPaymentMethod) SetPaymentSourceId(v string)`

SetPaymentSourceId sets PaymentSourceId field to given value.

### HasPaymentSourceId

`func (o *ChargeRequestPaymentMethod) HasPaymentSourceId() bool`

HasPaymentSourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


