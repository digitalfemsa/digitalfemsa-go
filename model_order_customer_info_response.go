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
)

// checks if the OrderCustomerInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCustomerInfoResponse{}

// OrderCustomerInfoResponse struct for OrderCustomerInfoResponse
type OrderCustomerInfoResponse struct {
	// Custom reference
	CustomerCustomReference NullableString `json:"customer_custom_reference,omitempty"`
	Name                    *string        `json:"name,omitempty"`
	Email                   *string        `json:"email,omitempty"`
	Phone                   *string        `json:"phone,omitempty"`
	Corporate               *bool          `json:"corporate,omitempty"`
	Object                  *string        `json:"object,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _OrderCustomerInfoResponse OrderCustomerInfoResponse

// NewOrderCustomerInfoResponse instantiates a new OrderCustomerInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCustomerInfoResponse() *OrderCustomerInfoResponse {
	this := OrderCustomerInfoResponse{}
	var corporate bool = false
	this.Corporate = &corporate
	return &this
}

// NewOrderCustomerInfoResponseWithDefaults instantiates a new OrderCustomerInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCustomerInfoResponseWithDefaults() *OrderCustomerInfoResponse {
	this := OrderCustomerInfoResponse{}
	var corporate bool = false
	this.Corporate = &corporate
	return &this
}

// GetCustomerCustomReference returns the CustomerCustomReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderCustomerInfoResponse) GetCustomerCustomReference() string {
	if o == nil || IsNil(o.CustomerCustomReference.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerCustomReference.Get()
}

// GetCustomerCustomReferenceOk returns a tuple with the CustomerCustomReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderCustomerInfoResponse) GetCustomerCustomReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerCustomReference.Get(), o.CustomerCustomReference.IsSet()
}

// HasCustomerCustomReference returns a boolean if a field has been set.
func (o *OrderCustomerInfoResponse) HasCustomerCustomReference() bool {
	if o != nil && o.CustomerCustomReference.IsSet() {
		return true
	}

	return false
}

// SetCustomerCustomReference gets a reference to the given NullableString and assigns it to the CustomerCustomReference field.
func (o *OrderCustomerInfoResponse) SetCustomerCustomReference(v string) {
	o.CustomerCustomReference.Set(&v)
}

// SetCustomerCustomReferenceNil sets the value for CustomerCustomReference to be an explicit nil
func (o *OrderCustomerInfoResponse) SetCustomerCustomReferenceNil() {
	o.CustomerCustomReference.Set(nil)
}

// UnsetCustomerCustomReference ensures that no value is present for CustomerCustomReference, not even an explicit nil
func (o *OrderCustomerInfoResponse) UnsetCustomerCustomReference() {
	o.CustomerCustomReference.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrderCustomerInfoResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCustomerInfoResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrderCustomerInfoResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrderCustomerInfoResponse) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrderCustomerInfoResponse) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCustomerInfoResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrderCustomerInfoResponse) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrderCustomerInfoResponse) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *OrderCustomerInfoResponse) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCustomerInfoResponse) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *OrderCustomerInfoResponse) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *OrderCustomerInfoResponse) SetPhone(v string) {
	o.Phone = &v
}

// GetCorporate returns the Corporate field value if set, zero value otherwise.
func (o *OrderCustomerInfoResponse) GetCorporate() bool {
	if o == nil || IsNil(o.Corporate) {
		var ret bool
		return ret
	}
	return *o.Corporate
}

// GetCorporateOk returns a tuple with the Corporate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCustomerInfoResponse) GetCorporateOk() (*bool, bool) {
	if o == nil || IsNil(o.Corporate) {
		return nil, false
	}
	return o.Corporate, true
}

// HasCorporate returns a boolean if a field has been set.
func (o *OrderCustomerInfoResponse) HasCorporate() bool {
	if o != nil && !IsNil(o.Corporate) {
		return true
	}

	return false
}

// SetCorporate gets a reference to the given bool and assigns it to the Corporate field.
func (o *OrderCustomerInfoResponse) SetCorporate(v bool) {
	o.Corporate = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *OrderCustomerInfoResponse) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCustomerInfoResponse) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *OrderCustomerInfoResponse) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *OrderCustomerInfoResponse) SetObject(v string) {
	o.Object = &v
}

func (o OrderCustomerInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCustomerInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerCustomReference.IsSet() {
		toSerialize["customer_custom_reference"] = o.CustomerCustomReference.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Corporate) {
		toSerialize["corporate"] = o.Corporate
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderCustomerInfoResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderCustomerInfoResponse := _OrderCustomerInfoResponse{}

	err = json.Unmarshal(data, &varOrderCustomerInfoResponse)

	if err != nil {
		return err
	}

	*o = OrderCustomerInfoResponse(varOrderCustomerInfoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customer_custom_reference")
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "phone")
		delete(additionalProperties, "corporate")
		delete(additionalProperties, "object")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderCustomerInfoResponse struct {
	value *OrderCustomerInfoResponse
	isSet bool
}

func (v NullableOrderCustomerInfoResponse) Get() *OrderCustomerInfoResponse {
	return v.value
}

func (v *NullableOrderCustomerInfoResponse) Set(val *OrderCustomerInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCustomerInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCustomerInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCustomerInfoResponse(val *OrderCustomerInfoResponse) *NullableOrderCustomerInfoResponse {
	return &NullableOrderCustomerInfoResponse{value: val, isSet: true}
}

func (v NullableOrderCustomerInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCustomerInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
