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

// checks if the OrderDiscountLinesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDiscountLinesRequest{}

// OrderDiscountLinesRequest List of discounts that apply to the order.
type OrderDiscountLinesRequest struct {
	// The amount to be deducted from the total sum of all payments, in cents.
	Amount int64 `json:"amount"`
	// Discount code.
	Code string `json:"code"`
	// It can be 'loyalty', 'campaign', 'coupon' o 'sign'
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _OrderDiscountLinesRequest OrderDiscountLinesRequest

// NewOrderDiscountLinesRequest instantiates a new OrderDiscountLinesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDiscountLinesRequest(amount int64, code string, type_ string) *OrderDiscountLinesRequest {
	this := OrderDiscountLinesRequest{}
	this.Amount = amount
	this.Code = code
	this.Type = type_
	return &this
}

// NewOrderDiscountLinesRequestWithDefaults instantiates a new OrderDiscountLinesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDiscountLinesRequestWithDefaults() *OrderDiscountLinesRequest {
	this := OrderDiscountLinesRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *OrderDiscountLinesRequest) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *OrderDiscountLinesRequest) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *OrderDiscountLinesRequest) SetAmount(v int64) {
	o.Amount = v
}

// GetCode returns the Code field value
func (o *OrderDiscountLinesRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OrderDiscountLinesRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OrderDiscountLinesRequest) SetCode(v string) {
	o.Code = v
}

// GetType returns the Type field value
func (o *OrderDiscountLinesRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrderDiscountLinesRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderDiscountLinesRequest) SetType(v string) {
	o.Type = v
}

func (o OrderDiscountLinesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDiscountLinesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["code"] = o.Code
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderDiscountLinesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"code",
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

	varOrderDiscountLinesRequest := _OrderDiscountLinesRequest{}

	err = json.Unmarshal(data, &varOrderDiscountLinesRequest)

	if err != nil {
		return err
	}

	*o = OrderDiscountLinesRequest(varOrderDiscountLinesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "code")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderDiscountLinesRequest struct {
	value *OrderDiscountLinesRequest
	isSet bool
}

func (v NullableOrderDiscountLinesRequest) Get() *OrderDiscountLinesRequest {
	return v.value
}

func (v *NullableOrderDiscountLinesRequest) Set(val *OrderDiscountLinesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDiscountLinesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDiscountLinesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDiscountLinesRequest(val *OrderDiscountLinesRequest) *NullableOrderDiscountLinesRequest {
	return &NullableOrderDiscountLinesRequest{value: val, isSet: true}
}

func (v NullableOrderDiscountLinesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDiscountLinesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}