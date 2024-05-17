# OrderResponseCheckout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPaymentMethods** | Pointer to **[]string** |  | [optional] 
**CanNotExpire** | Pointer to **bool** |  | [optional] 
**EmailsSent** | Pointer to **int32** |  | [optional] 
**ExpiresAt** | Pointer to **int64** |  | [optional] 
**FailureUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsRedirectOnFailure** | Pointer to **bool** |  | [optional] 
**Livemode** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NeedsShippingContact** | Pointer to **bool** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**OnDemandEnabled** | Pointer to **NullableBool** |  | [optional] 
**Recurrent** | Pointer to **bool** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**SmsSent** | Pointer to **int32** |  | [optional] 
**SuccessUrl** | Pointer to **string** |  | [optional] 
**StartsAt** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewOrderResponseCheckout

`func NewOrderResponseCheckout() *OrderResponseCheckout`

NewOrderResponseCheckout instantiates a new OrderResponseCheckout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResponseCheckoutWithDefaults

`func NewOrderResponseCheckoutWithDefaults() *OrderResponseCheckout`

NewOrderResponseCheckoutWithDefaults instantiates a new OrderResponseCheckout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPaymentMethods

`func (o *OrderResponseCheckout) GetAllowedPaymentMethods() []string`

GetAllowedPaymentMethods returns the AllowedPaymentMethods field if non-nil, zero value otherwise.

### GetAllowedPaymentMethodsOk

`func (o *OrderResponseCheckout) GetAllowedPaymentMethodsOk() (*[]string, bool)`

GetAllowedPaymentMethodsOk returns a tuple with the AllowedPaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPaymentMethods

`func (o *OrderResponseCheckout) SetAllowedPaymentMethods(v []string)`

SetAllowedPaymentMethods sets AllowedPaymentMethods field to given value.

### HasAllowedPaymentMethods

`func (o *OrderResponseCheckout) HasAllowedPaymentMethods() bool`

HasAllowedPaymentMethods returns a boolean if a field has been set.

### GetCanNotExpire

`func (o *OrderResponseCheckout) GetCanNotExpire() bool`

GetCanNotExpire returns the CanNotExpire field if non-nil, zero value otherwise.

### GetCanNotExpireOk

`func (o *OrderResponseCheckout) GetCanNotExpireOk() (*bool, bool)`

GetCanNotExpireOk returns a tuple with the CanNotExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanNotExpire

`func (o *OrderResponseCheckout) SetCanNotExpire(v bool)`

SetCanNotExpire sets CanNotExpire field to given value.

### HasCanNotExpire

`func (o *OrderResponseCheckout) HasCanNotExpire() bool`

HasCanNotExpire returns a boolean if a field has been set.

### GetEmailsSent

`func (o *OrderResponseCheckout) GetEmailsSent() int32`

GetEmailsSent returns the EmailsSent field if non-nil, zero value otherwise.

### GetEmailsSentOk

`func (o *OrderResponseCheckout) GetEmailsSentOk() (*int32, bool)`

GetEmailsSentOk returns a tuple with the EmailsSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsSent

`func (o *OrderResponseCheckout) SetEmailsSent(v int32)`

SetEmailsSent sets EmailsSent field to given value.

### HasEmailsSent

`func (o *OrderResponseCheckout) HasEmailsSent() bool`

HasEmailsSent returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OrderResponseCheckout) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrderResponseCheckout) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrderResponseCheckout) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OrderResponseCheckout) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFailureUrl

`func (o *OrderResponseCheckout) GetFailureUrl() string`

GetFailureUrl returns the FailureUrl field if non-nil, zero value otherwise.

### GetFailureUrlOk

`func (o *OrderResponseCheckout) GetFailureUrlOk() (*string, bool)`

GetFailureUrlOk returns a tuple with the FailureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureUrl

`func (o *OrderResponseCheckout) SetFailureUrl(v string)`

SetFailureUrl sets FailureUrl field to given value.

### HasFailureUrl

`func (o *OrderResponseCheckout) HasFailureUrl() bool`

HasFailureUrl returns a boolean if a field has been set.

### GetId

`func (o *OrderResponseCheckout) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderResponseCheckout) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderResponseCheckout) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderResponseCheckout) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRedirectOnFailure

`func (o *OrderResponseCheckout) GetIsRedirectOnFailure() bool`

GetIsRedirectOnFailure returns the IsRedirectOnFailure field if non-nil, zero value otherwise.

### GetIsRedirectOnFailureOk

`func (o *OrderResponseCheckout) GetIsRedirectOnFailureOk() (*bool, bool)`

GetIsRedirectOnFailureOk returns a tuple with the IsRedirectOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRedirectOnFailure

`func (o *OrderResponseCheckout) SetIsRedirectOnFailure(v bool)`

SetIsRedirectOnFailure sets IsRedirectOnFailure field to given value.

### HasIsRedirectOnFailure

`func (o *OrderResponseCheckout) HasIsRedirectOnFailure() bool`

HasIsRedirectOnFailure returns a boolean if a field has been set.

### GetLivemode

`func (o *OrderResponseCheckout) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *OrderResponseCheckout) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *OrderResponseCheckout) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *OrderResponseCheckout) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetMetadata

`func (o *OrderResponseCheckout) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderResponseCheckout) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderResponseCheckout) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderResponseCheckout) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *OrderResponseCheckout) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrderResponseCheckout) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrderResponseCheckout) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrderResponseCheckout) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedsShippingContact

`func (o *OrderResponseCheckout) GetNeedsShippingContact() bool`

GetNeedsShippingContact returns the NeedsShippingContact field if non-nil, zero value otherwise.

### GetNeedsShippingContactOk

`func (o *OrderResponseCheckout) GetNeedsShippingContactOk() (*bool, bool)`

GetNeedsShippingContactOk returns a tuple with the NeedsShippingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsShippingContact

`func (o *OrderResponseCheckout) SetNeedsShippingContact(v bool)`

SetNeedsShippingContact sets NeedsShippingContact field to given value.

### HasNeedsShippingContact

`func (o *OrderResponseCheckout) HasNeedsShippingContact() bool`

HasNeedsShippingContact returns a boolean if a field has been set.

### GetObject

`func (o *OrderResponseCheckout) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OrderResponseCheckout) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OrderResponseCheckout) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *OrderResponseCheckout) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOnDemandEnabled

`func (o *OrderResponseCheckout) GetOnDemandEnabled() bool`

GetOnDemandEnabled returns the OnDemandEnabled field if non-nil, zero value otherwise.

### GetOnDemandEnabledOk

`func (o *OrderResponseCheckout) GetOnDemandEnabledOk() (*bool, bool)`

GetOnDemandEnabledOk returns a tuple with the OnDemandEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemandEnabled

`func (o *OrderResponseCheckout) SetOnDemandEnabled(v bool)`

SetOnDemandEnabled sets OnDemandEnabled field to given value.

### HasOnDemandEnabled

`func (o *OrderResponseCheckout) HasOnDemandEnabled() bool`

HasOnDemandEnabled returns a boolean if a field has been set.

### SetOnDemandEnabledNil

`func (o *OrderResponseCheckout) SetOnDemandEnabledNil(b bool)`

 SetOnDemandEnabledNil sets the value for OnDemandEnabled to be an explicit nil

### UnsetOnDemandEnabled
`func (o *OrderResponseCheckout) UnsetOnDemandEnabled()`

UnsetOnDemandEnabled ensures that no value is present for OnDemandEnabled, not even an explicit nil
### GetRecurrent

`func (o *OrderResponseCheckout) GetRecurrent() bool`

GetRecurrent returns the Recurrent field if non-nil, zero value otherwise.

### GetRecurrentOk

`func (o *OrderResponseCheckout) GetRecurrentOk() (*bool, bool)`

GetRecurrentOk returns a tuple with the Recurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrent

`func (o *OrderResponseCheckout) SetRecurrent(v bool)`

SetRecurrent sets Recurrent field to given value.

### HasRecurrent

`func (o *OrderResponseCheckout) HasRecurrent() bool`

HasRecurrent returns a boolean if a field has been set.

### GetSlug

`func (o *OrderResponseCheckout) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OrderResponseCheckout) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OrderResponseCheckout) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *OrderResponseCheckout) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSmsSent

`func (o *OrderResponseCheckout) GetSmsSent() int32`

GetSmsSent returns the SmsSent field if non-nil, zero value otherwise.

### GetSmsSentOk

`func (o *OrderResponseCheckout) GetSmsSentOk() (*int32, bool)`

GetSmsSentOk returns a tuple with the SmsSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsSent

`func (o *OrderResponseCheckout) SetSmsSent(v int32)`

SetSmsSent sets SmsSent field to given value.

### HasSmsSent

`func (o *OrderResponseCheckout) HasSmsSent() bool`

HasSmsSent returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *OrderResponseCheckout) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *OrderResponseCheckout) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *OrderResponseCheckout) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *OrderResponseCheckout) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.

### GetStartsAt

`func (o *OrderResponseCheckout) GetStartsAt() int32`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *OrderResponseCheckout) GetStartsAtOk() (*int32, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *OrderResponseCheckout) SetStartsAt(v int32)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *OrderResponseCheckout) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetStatus

`func (o *OrderResponseCheckout) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderResponseCheckout) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderResponseCheckout) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderResponseCheckout) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *OrderResponseCheckout) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderResponseCheckout) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderResponseCheckout) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderResponseCheckout) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *OrderResponseCheckout) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrderResponseCheckout) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrderResponseCheckout) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OrderResponseCheckout) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


