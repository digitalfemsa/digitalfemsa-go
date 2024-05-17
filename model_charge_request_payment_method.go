/*
Femsa API

Femsa sdk

API version: 2.1.0
Contact: engineering@femsa.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package digitalfemsa

import (
	"encoding/json"
	"fmt"
)

// checks if the ChargeRequestPaymentMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargeRequestPaymentMethod{}

// ChargeRequestPaymentMethod Payment method used in the charge. Go to the [payment methods](https://developers.femsa.com/reference/m%C3%A9todos-de-pago) section for more details
type ChargeRequestPaymentMethod struct {
	// Method expiration date as unix timestamp
	ExpiresAt            *int64  `json:"expires_at,omitempty"`
	Type                 string  `json:"type"`
	PaymentSourceId      *string `json:"payment_source_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChargeRequestPaymentMethod ChargeRequestPaymentMethod

// NewChargeRequestPaymentMethod instantiates a new ChargeRequestPaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargeRequestPaymentMethod(type_ string) *ChargeRequestPaymentMethod {
	this := ChargeRequestPaymentMethod{}
	this.Type = type_
	return &this
}

// NewChargeRequestPaymentMethodWithDefaults instantiates a new ChargeRequestPaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargeRequestPaymentMethodWithDefaults() *ChargeRequestPaymentMethod {
	this := ChargeRequestPaymentMethod{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ChargeRequestPaymentMethod) GetExpiresAt() int64 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeRequestPaymentMethod) GetExpiresAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ChargeRequestPaymentMethod) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *ChargeRequestPaymentMethod) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetType returns the Type field value
func (o *ChargeRequestPaymentMethod) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChargeRequestPaymentMethod) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ChargeRequestPaymentMethod) SetType(v string) {
	o.Type = v
}

// GetPaymentSourceId returns the PaymentSourceId field value if set, zero value otherwise.
func (o *ChargeRequestPaymentMethod) GetPaymentSourceId() string {
	if o == nil || IsNil(o.PaymentSourceId) {
		var ret string
		return ret
	}
	return *o.PaymentSourceId
}

// GetPaymentSourceIdOk returns a tuple with the PaymentSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargeRequestPaymentMethod) GetPaymentSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentSourceId) {
		return nil, false
	}
	return o.PaymentSourceId, true
}

// HasPaymentSourceId returns a boolean if a field has been set.
func (o *ChargeRequestPaymentMethod) HasPaymentSourceId() bool {
	if o != nil && !IsNil(o.PaymentSourceId) {
		return true
	}

	return false
}

// SetPaymentSourceId gets a reference to the given string and assigns it to the PaymentSourceId field.
func (o *ChargeRequestPaymentMethod) SetPaymentSourceId(v string) {
	o.PaymentSourceId = &v
}

func (o ChargeRequestPaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargeRequestPaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.PaymentSourceId) {
		toSerialize["payment_source_id"] = o.PaymentSourceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChargeRequestPaymentMethod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varChargeRequestPaymentMethod := _ChargeRequestPaymentMethod{}

	err = json.Unmarshal(data, &varChargeRequestPaymentMethod)

	if err != nil {
		return err
	}

	*o = ChargeRequestPaymentMethod(varChargeRequestPaymentMethod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "type")
		delete(additionalProperties, "payment_source_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChargeRequestPaymentMethod struct {
	value *ChargeRequestPaymentMethod
	isSet bool
}

func (v NullableChargeRequestPaymentMethod) Get() *ChargeRequestPaymentMethod {
	return v.value
}

func (v *NullableChargeRequestPaymentMethod) Set(val *ChargeRequestPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableChargeRequestPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableChargeRequestPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargeRequestPaymentMethod(val *ChargeRequestPaymentMethod) *NullableChargeRequestPaymentMethod {
	return &NullableChargeRequestPaymentMethod{value: val, isSet: true}
}

func (v NullableChargeRequestPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargeRequestPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
