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

// CustomerPaymentMethodsData - struct for CustomerPaymentMethodsData
type CustomerPaymentMethodsData struct {
	PaymentMethodCashResponse *PaymentMethodCashResponse
}

// PaymentMethodCashResponseAsCustomerPaymentMethodsData is a convenience function that returns PaymentMethodCashResponse wrapped in CustomerPaymentMethodsData
func PaymentMethodCashResponseAsCustomerPaymentMethodsData(v *PaymentMethodCashResponse) CustomerPaymentMethodsData {
	return CustomerPaymentMethodsData{
		PaymentMethodCashResponse: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CustomerPaymentMethodsData) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'cash'
	if jsonDict["type"] == "cash" {
		// try to unmarshal JSON data into PaymentMethodCashResponse
		err = json.Unmarshal(data, &dst.PaymentMethodCashResponse)
		if err == nil {
			return nil // data stored in dst.PaymentMethodCashResponse, return on the first match
		} else {
			dst.PaymentMethodCashResponse = nil
			return fmt.Errorf("failed to unmarshal CustomerPaymentMethodsData as PaymentMethodCashResponse: %s", err.Error())
		}
	}

	// check if the discriminator value is 'oxxo_recurrent'
	if jsonDict["type"] == "oxxo_recurrent" {
		// try to unmarshal JSON data into PaymentMethodCashResponse
		err = json.Unmarshal(data, &dst.PaymentMethodCashResponse)
		if err == nil {
			return nil // data stored in dst.PaymentMethodCashResponse, return on the first match
		} else {
			dst.PaymentMethodCashResponse = nil
			return fmt.Errorf("failed to unmarshal CustomerPaymentMethodsData as PaymentMethodCashResponse: %s", err.Error())
		}
	}

	// check if the discriminator value is 'payment_method_cash_response'
	if jsonDict["type"] == "payment_method_cash_response" {
		// try to unmarshal JSON data into PaymentMethodCashResponse
		err = json.Unmarshal(data, &dst.PaymentMethodCashResponse)
		if err == nil {
			return nil // data stored in dst.PaymentMethodCashResponse, return on the first match
		} else {
			dst.PaymentMethodCashResponse = nil
			return fmt.Errorf("failed to unmarshal CustomerPaymentMethodsData as PaymentMethodCashResponse: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CustomerPaymentMethodsData) MarshalJSON() ([]byte, error) {
	if src.PaymentMethodCashResponse != nil {
		return json.Marshal(&src.PaymentMethodCashResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CustomerPaymentMethodsData) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.PaymentMethodCashResponse != nil {
		return obj.PaymentMethodCashResponse
	}

	// all schemas are nil
	return nil
}

type NullableCustomerPaymentMethodsData struct {
	value *CustomerPaymentMethodsData
	isSet bool
}

func (v NullableCustomerPaymentMethodsData) Get() *CustomerPaymentMethodsData {
	return v.value
}

func (v *NullableCustomerPaymentMethodsData) Set(val *CustomerPaymentMethodsData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentMethodsData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentMethodsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentMethodsData(val *CustomerPaymentMethodsData) *NullableCustomerPaymentMethodsData {
	return &NullableCustomerPaymentMethodsData{value: val, isSet: true}
}

func (v NullableCustomerPaymentMethodsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentMethodsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
