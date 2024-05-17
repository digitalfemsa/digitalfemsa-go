# CheckoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPaymentMethods** | **[]string** | Are the payment methods available for this link | 
**ExpiresAt** | Pointer to **int64** | Unix timestamp of checkout expiration | [optional] 
**FailureUrl** | Pointer to **string** | Redirection url back to the site in case of failed payment, applies only to HostedPayment. | [optional] 
**Name** | Pointer to **string** | Reason for payment | [optional] 
**OnDemandEnabled** | Pointer to **bool** |  | [optional] 
**SuccessUrl** | Pointer to **string** | Redirection url back to the site in case of successful payment, applies only to HostedPayment | [optional] 
**Type** | Pointer to **string** | This field represents the type of checkout | [optional] 

## Methods

### NewCheckoutRequest

`func NewCheckoutRequest(allowedPaymentMethods []string, ) *CheckoutRequest`

NewCheckoutRequest instantiates a new CheckoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutRequestWithDefaults

`func NewCheckoutRequestWithDefaults() *CheckoutRequest`

NewCheckoutRequestWithDefaults instantiates a new CheckoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPaymentMethods

`func (o *CheckoutRequest) GetAllowedPaymentMethods() []string`

GetAllowedPaymentMethods returns the AllowedPaymentMethods field if non-nil, zero value otherwise.

### GetAllowedPaymentMethodsOk

`func (o *CheckoutRequest) GetAllowedPaymentMethodsOk() (*[]string, bool)`

GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPaymentMethods

`func (o *CheckoutRequest) SetAllowedPaymentMethods(v []string)`

SetAllowedPaymentMethods sets AllowedPaymentMethods field to given value.


### GetExpiresAt

`func (o *CheckoutRequest) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CheckoutRequest) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CheckoutRequest) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CheckoutRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFailureUrl

`func (o *CheckoutRequest) GetFailureUrl() string`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *CheckoutRequest) GetFailureUrlOk() (*string, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *CheckoutRequest) SetFailureUrl(v string)`

SetFailureUrl sets FailureUrl field to given value.

### HasFailureUrl

`func (o *CheckoutRequest) HasFailureUrl() bool`

HasFailureUrl returns a boolean if a field has been set.

### GetName

`func (o *CheckoutRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckoutRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckoutRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CheckoutRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnDemandEnabled

`func (o *CheckoutRequest) GetOnDemandEnabled() bool`

GetOnDemandEnabled returns the OnDemandEnabled field if non-nil, zero value otherwise.

### GetOnDemandEnabledOk

`func (o *CheckoutRequest) GetOnDemandEnabledOk() (*bool, bool)`

GetOnDemandEnabledOk returns a tuple with the OnDemandEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemandEnabled

`func (o *CheckoutRequest) SetOnDemandEnabled(v bool)`

SetOnDemandEnabled sets OnDemandEnabled field to given value.

### HasOnDemandEnabled

`func (o *CheckoutRequest) HasOnDemandEnabled() bool`

HasOnDemandEnabled returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *CheckoutRequest) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *CheckoutRequest) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *CheckoutRequest) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *CheckoutRequest) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetType

`func (o *CheckoutRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CheckoutRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


