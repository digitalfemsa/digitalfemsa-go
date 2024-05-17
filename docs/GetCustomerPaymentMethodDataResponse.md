# GetCustomerPaymentMethodDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Object** | **string** |  | 
**CreatedAt** | **int64** |  | 
**ParentId** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**Barcode** | Pointer to **string** |  | [optional] 
**BarcodeUrl** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **int64** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 

## Methods

### NewGetCustomerPaymentMethodDataResponse

`func NewGetCustomerPaymentMethodDataResponse(type_ string, id string, object string, createdAt int64, ) *GetCustomerPaymentMethodDataResponse`

NewGetCustomerPaymentMethodDataResponse instantiates a new GetCustomerPaymentMethodDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomerPaymentMethodDataResponseWithDefaults

`func NewGetCustomerPaymentMethodDataResponseWithDefaults() *GetCustomerPaymentMethodDataResponse`

NewGetCustomerPaymentMethodDataResponseWithDefaults instantiates a new GetCustomerPaymentMethodDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetCustomerPaymentMethodDataResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCustomerPaymentMethodDataResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCustomerPaymentMethodDataResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GetCustomerPaymentMethodDataResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCustomerPaymentMethodDataResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCustomerPaymentMethodDataResponse) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *GetCustomerPaymentMethodDataResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GetCustomerPaymentMethodDataResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GetCustomerPaymentMethodDataResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetCreatedAt

`func (o *GetCustomerPaymentMethodDataResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetCustomerPaymentMethodDataResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetCustomerPaymentMethodDataResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetParentId

`func (o *GetCustomerPaymentMethodDataResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *GetCustomerPaymentMethodDataResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *GetCustomerPaymentMethodDataResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *GetCustomerPaymentMethodDataResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetReference

`func (o *GetCustomerPaymentMethodDataResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GetCustomerPaymentMethodDataResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GetCustomerPaymentMethodDataResponse) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *GetCustomerPaymentMethodDataResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetBarcode

`func (o *GetCustomerPaymentMethodDataResponse) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *GetCustomerPaymentMethodDataResponse) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *GetCustomerPaymentMethodDataResponse) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *GetCustomerPaymentMethodDataResponse) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetBarcodeUrl

`func (o *GetCustomerPaymentMethodDataResponse) GetBarcodeUrl() string`

GetBarcodeUrl returns the BarcodeUrl field if non-nil, zero value otherwise.

### GetBarcodeUrlOk

`func (o *GetCustomerPaymentMethodDataResponse) GetBarcodeUrlOk() (*string, bool)`

GetBarcodeUrlOk returns a tuple with the BarcodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeUrl

`func (o *GetCustomerPaymentMethodDataResponse) SetBarcodeUrl(v string)`

SetBarcodeUrl sets BarcodeUrl field to given value.

### HasBarcodeUrl

`func (o *GetCustomerPaymentMethodDataResponse) HasBarcodeUrl() bool`

HasBarcodeUrl returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GetCustomerPaymentMethodDataResponse) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GetCustomerPaymentMethodDataResponse) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GetCustomerPaymentMethodDataResponse) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GetCustomerPaymentMethodDataResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetProvider

`func (o *GetCustomerPaymentMethodDataResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GetCustomerPaymentMethodDataResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GetCustomerPaymentMethodDataResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GetCustomerPaymentMethodDataResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


