/*
Masters way general API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SchemasAIChatResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasAIChatResponse{}

// SchemasAIChatResponse struct for SchemasAIChatResponse
type SchemasAIChatResponse struct {
	Message string `json:"message"`
}

type _SchemasAIChatResponse SchemasAIChatResponse

// NewSchemasAIChatResponse instantiates a new SchemasAIChatResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasAIChatResponse(message string) *SchemasAIChatResponse {
	this := SchemasAIChatResponse{}
	this.Message = message
	return &this
}

// NewSchemasAIChatResponseWithDefaults instantiates a new SchemasAIChatResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasAIChatResponseWithDefaults() *SchemasAIChatResponse {
	this := SchemasAIChatResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *SchemasAIChatResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SchemasAIChatResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SchemasAIChatResponse) SetMessage(v string) {
	o.Message = v
}

func (o SchemasAIChatResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasAIChatResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *SchemasAIChatResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSchemasAIChatResponse := _SchemasAIChatResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasAIChatResponse)

	if err != nil {
		return err
	}

	*o = SchemasAIChatResponse(varSchemasAIChatResponse)

	return err
}

type NullableSchemasAIChatResponse struct {
	value *SchemasAIChatResponse
	isSet bool
}

func (v NullableSchemasAIChatResponse) Get() *SchemasAIChatResponse {
	return v.value
}

func (v *NullableSchemasAIChatResponse) Set(val *SchemasAIChatResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasAIChatResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasAIChatResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasAIChatResponse(val *SchemasAIChatResponse) *NullableSchemasAIChatResponse {
	return &NullableSchemasAIChatResponse{value: val, isSet: true}
}

func (v NullableSchemasAIChatResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasAIChatResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


