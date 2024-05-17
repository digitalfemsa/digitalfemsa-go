# ChargeOrderResponsePaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Object** | **string** |  | 
**AuthCode** | Pointer to **NullableInt32** |  | [optional] 
**CashierId** | Pointer to **NullableString** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**BarcodeUrl** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **int64** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**Store** | Pointer to **NullableString** |  | [optional] 
**StoreName** | Pointer to **string** |  | [optional] 

## Methods

### NewChargeOrderResponsePaymentMethod

`func NewChargeOrderResponsePaymentMethod(object string, ) *ChargeOrderResponsePaymentMethod`

NewChargeOrderResponsePaymentMethod instantiates a new ChargeOrderResponsePaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeOrderResponsePaymentMethodWithDefaults

`func NewChargeOrderResponsePaymentMethodWithDefaults() *ChargeOrderResponsePaymentMethod`

NewChargeOrderResponsePaymentMethodWithDefaults instantiates a new ChargeOrderResponsePaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChargeOrderResponsePaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChargeOrderResponsePaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChargeOrderResponsePaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChargeOrderResponsePaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetObject

`func (o *ChargeOrderResponsePaymentMethod) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChargeOrderResponsePaymentMethod) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChargeOrderResponsePaymentMethod) SetObject(v string)`

SetObject sets Object field to given value.


### GetAuthCode

`func (o *ChargeOrderResponsePaymentMethod) GetAuthCode() int32`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *ChargeOrderResponsePaymentMethod) GetAuthCodeOk() (*int32, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *ChargeOrderResponsePaymentMethod) SetAuthCode(v int32)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *ChargeOrderResponsePaymentMethod) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### SetAuthCodeNil

`func (o *ChargeOrderResponsePaymentMethod) SetAuthCodeNil(b bool)`

 SetAuthCodeNil sets the value for AuthCode to be an explicit nil

### UnsetAuthCode
`func (o *ChargeOrderResponsePaymentMethod) UnsetAuthCode()`

UnsetAuthCode ensures that no value is present for AuthCode, not even an explicit nil
### GetCashierId

`func (o *ChargeOrderResponsePaymentMethod) GetCashierId() string`

GetCashierId returns the CashierId field if non-nil, zero value otherwise.

### GetCashierIdOk

`func (o *ChargeOrderResponsePaymentMethod) GetCashierIdOk() (*string, bool)`

GetCashierIdOk returns a tuple with the CashierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashierId

`func (o *ChargeOrderResponsePaymentMethod) SetCashierId(v string)`

SetCashierId sets CashierId field to given value.

### HasCashierId

`func (o *ChargeOrderResponsePaymentMethod) HasCashierId() bool`

HasCashierId returns a boolean if a field has been set.

### SetCashierIdNil

`func (o *ChargeOrderResponsePaymentMethod) SetCashierIdNil(b bool)`

 SetCashierIdNil sets the value for CashierId to be an explicit nil

### UnsetCashierId
`func (o *ChargeOrderResponsePaymentMethod) UnsetCashierId()`

UnsetCashierId ensures that no value is present for CashierId, not even an explicit nil
### GetReference

`func (o *ChargeOrderResponsePaymentMethod) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ChargeOrderResponsePaymentMethod) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ChargeOrderResponsePaymentMethod) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ChargeOrderResponsePaymentMethod) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetBarcodeUrl

`func (o *ChargeOrderResponsePaymentMethod) GetBarcodeUrl() string`

GetBarcodeUrl returns the BarcodeUrl field if non-nil, zero value otherwise.

### GetBarcodeUrlOk

`func (o *ChargeOrderResponsePaymentMethod) GetBarcodeUrlOk() (*string, bool)`

GetBarcodeUrlOk returns a tuple with the BarcodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeUrl

`func (o *ChargeOrderResponsePaymentMethod) SetBarcodeUrl(v string)`

SetBarcodeUrl sets BarcodeUrl field to given value.

### HasBarcodeUrl

`func (o *ChargeOrderResponsePaymentMethod) HasBarcodeUrl() bool`

HasBarcodeUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ChargeOrderResponsePaymentMethod) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ChargeOrderResponsePaymentMethod) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ChargeOrderResponsePaymentMethod) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ChargeOrderResponsePaymentMethod) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetServiceName

`func (o *ChargeOrderResponsePaymentMethod) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ChargeOrderResponsePaymentMethod) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ChargeOrderResponsePaymentMethod) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ChargeOrderResponsePaymentMethod) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStore

`func (o *ChargeOrderResponsePaymentMethod) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ChargeOrderResponsePaymentMethod) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ChargeOrderResponsePaymentMethod) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *ChargeOrderResponsePaymentMethod) HasStore() bool`

HasStore returns a boolean if a field has been set.

### SetStoreNil

`func (o *ChargeOrderResponsePaymentMethod) SetStoreNil(b bool)`

 SetStoreNil sets the value for Store to be an explicit nil

### UnsetStore
`func (o *ChargeOrderResponsePaymentMethod) UnsetStore()`

UnsetStore ensures that no value is present for Store, not even an explicit nil
### GetStoreName

`func (o *ChargeOrderResponsePaymentMethod) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *ChargeOrderResponsePaymentMethod) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *ChargeOrderResponsePaymentMethod) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *ChargeOrderResponsePaymentMethod) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


