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

// checks if the WebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookRequest{}

// WebhookRequest a webhook
type WebhookRequest struct {
	// Here you must place the URL of your Webhook remember that you must program what you will do with the events received. Also do not forget to handle the HTTPS protocol for greater security.
	Url string `json:"url"`
	// It is a value that allows to decide if the events will be synchronous or asynchronous. We recommend asynchronous = false
	Synchronous          bool `json:"synchronous"`
	AdditionalProperties map[string]interface{}
}

type _WebhookRequest WebhookRequest

// NewWebhookRequest instantiates a new WebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookRequest(url string, synchronous bool) *WebhookRequest {
	this := WebhookRequest{}
	this.Url = url
	this.Synchronous = synchronous
	return &this
}

// NewWebhookRequestWithDefaults instantiates a new WebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookRequestWithDefaults() *WebhookRequest {
	this := WebhookRequest{}
	var synchronous bool = false
	this.Synchronous = synchronous
	return &this
}

// GetUrl returns the Url field value
func (o *WebhookRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WebhookRequest) SetUrl(v string) {
	o.Url = v
}

// GetSynchronous returns the Synchronous field value
func (o *WebhookRequest) GetSynchronous() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Synchronous
}

// GetSynchronousOk returns a tuple with the Synchronous field value
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetSynchronousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Synchronous, true
}

// SetSynchronous sets field value
func (o *WebhookRequest) SetSynchronous(v bool) {
	o.Synchronous = v
}

func (o WebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["synchronous"] = o.Synchronous

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebhookRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"synchronous",
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

	varWebhookRequest := _WebhookRequest{}

	err = json.Unmarshal(data, &varWebhookRequest)

	if err != nil {
		return err
	}

	*o = WebhookRequest(varWebhookRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "synchronous")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhookRequest struct {
	value *WebhookRequest
	isSet bool
}

func (v NullableWebhookRequest) Get() *WebhookRequest {
	return v.value
}

func (v *NullableWebhookRequest) Set(val *WebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookRequest(val *WebhookRequest) *NullableWebhookRequest {
	return &NullableWebhookRequest{value: val, isSet: true}
}

func (v NullableWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
