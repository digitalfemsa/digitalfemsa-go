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

// checks if the WebhookKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookKeyRequest{}

// WebhookKeyRequest struct for WebhookKeyRequest
type WebhookKeyRequest struct {
	// Indicates if the webhook key is active
	Active               *bool `json:"active,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebhookKeyRequest WebhookKeyRequest

// NewWebhookKeyRequest instantiates a new WebhookKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookKeyRequest() *WebhookKeyRequest {
	this := WebhookKeyRequest{}
	var active bool = true
	this.Active = &active
	return &this
}

// NewWebhookKeyRequestWithDefaults instantiates a new WebhookKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookKeyRequestWithDefaults() *WebhookKeyRequest {
	this := WebhookKeyRequest{}
	var active bool = true
	this.Active = &active
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *WebhookKeyRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookKeyRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *WebhookKeyRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *WebhookKeyRequest) SetActive(v bool) {
	o.Active = &v
}

func (o WebhookKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebhookKeyRequest) UnmarshalJSON(data []byte) (err error) {
	varWebhookKeyRequest := _WebhookKeyRequest{}

	err = json.Unmarshal(data, &varWebhookKeyRequest)

	if err != nil {
		return err
	}

	*o = WebhookKeyRequest(varWebhookKeyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhookKeyRequest struct {
	value *WebhookKeyRequest
	isSet bool
}

func (v NullableWebhookKeyRequest) Get() *WebhookKeyRequest {
	return v.value
}

func (v *NullableWebhookKeyRequest) Set(val *WebhookKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookKeyRequest(val *WebhookKeyRequest) *NullableWebhookKeyRequest {
	return &NullableWebhookKeyRequest{value: val, isSet: true}
}

func (v NullableWebhookKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}