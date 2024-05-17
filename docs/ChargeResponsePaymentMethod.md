# ChargeResponsePaymentMethod

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

### NewChargeResponsePaymentMethod

`func NewChargeResponsePaymentMethod(object string, ) *ChargeResponsePaymentMethod`

NewChargeResponsePaymentMethod instantiates a new ChargeResponsePaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeResponsePaymentMethodWithDefaults

`func NewChargeResponsePaymentMethodWithDefaults() *ChargeResponsePaymentMethod`

NewChargeResponsePaymentMethodWithDefaults instantiates a new ChargeResponsePaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChargeResponsePaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChargeResponsePaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChargeResponsePaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChargeResponsePaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetObject

`func (o *ChargeResponsePaymentMethod) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChargeResponsePaymentMethod) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChargeResponsePaymentMethod) SetObject(v string)`

SetObject sets Object field to given value.


### GetAuthCode

`func (o *ChargeResponsePaymentMethod) GetAuthCode() int32`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *ChargeResponsePaymentMethod) GetAuthCodeOk() (*int32, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *ChargeResponsePaymentMethod) SetAuthCode(v int32)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *ChargeResponsePaymentMethod) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### SetAuthCodeNil

`func (o *ChargeResponsePaymentMethod) SetAuthCodeNil(b bool)`

 SetAuthCodeNil sets the value for AuthCode to be an explicit nil

### UnsetAuthCode
`func (o *ChargeResponsePaymentMethod) UnsetAuthCode()`

UnsetAuthCode ensures that no value is present for AuthCode, not even an explicit nil
### GetCashierId

`func (o *ChargeResponsePaymentMethod) GetCashierId() string`

GetCashierId returns the CashierId field if non-nil, zero value otherwise.

### GetCashierIdOk

`func (o *ChargeResponsePaymentMethod) GetCashierIdOk() (*string, bool)`

GetCashierIdOk returns a tuple with the CashierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashierId

`func (o *ChargeResponsePaymentMethod) SetCashierId(v string)`

SetCashierId sets CashierId field to given value.

### HasCashierId

`func (o *ChargeResponsePaymentMethod) HasCashierId() bool`

HasCashierId returns a boolean if a field has been set.

### SetCashierIdNil

`func (o *ChargeResponsePaymentMethod) SetCashierIdNil(b bool)`

 SetCashierIdNil sets the value for CashierId to be an explicit nil

### UnsetCashierId
`func (o *ChargeResponsePaymentMethod) UnsetCashierId()`

UnsetCashierId ensures that no value is present for CashierId, not even an explicit nil
### GetReference

`func (o *ChargeResponsePaymentMethod) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ChargeResponsePaymentMethod) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ChargeResponsePaymentMethod) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ChargeResponsePaymentMethod) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetBarcodeUrl

`func (o *ChargeResponsePaymentMethod) GetBarcodeUrl() string`

GetBarcodeUrl returns the BarcodeUrl field if non-nil, zero value otherwise.

### GetBarcodeUrlOk

`func (o *ChargeResponsePaymentMethod) GetBarcodeUrlOk() (*string, bool)`

GetBarcodeUrlOk returns a tuple with the BarcodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeUrl

`func (o *ChargeResponsePaymentMethod) SetBarcodeUrl(v string)`

SetBarcodeUrl sets BarcodeUrl field to given value.

### HasBarcodeUrl

`func (o *ChargeResponsePaymentMethod) HasBarcodeUrl() bool`

HasBarcodeUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ChargeResponsePaymentMethod) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ChargeResponsePaymentMethod) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ChargeResponsePaymentMethod) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ChargeResponsePaymentMethod) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetServiceName

`func (o *ChargeResponsePaymentMethod) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ChargeResponsePaymentMethod) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ChargeResponsePaymentMethod) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ChargeResponsePaymentMethod) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStore

`func (o *ChargeResponsePaymentMethod) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ChargeResponsePaymentMethod) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ChargeResponsePaymentMethod) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *ChargeResponsePaymentMethod) HasStore() bool`

HasStore returns a boolean if a field has been set.

### SetStoreNil

`func (o *ChargeResponsePaymentMethod) SetStoreNil(b bool)`

 SetStoreNil sets the value for Store to be an explicit nil

### UnsetStore
`func (o *ChargeResponsePaymentMethod) UnsetStore()`

UnsetStore ensures that no value is present for Store, not even an explicit nil
### GetStoreName

`func (o *ChargeResponsePaymentMethod) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *ChargeResponsePaymentMethod) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *ChargeResponsePaymentMethod) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *ChargeResponsePaymentMethod) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


