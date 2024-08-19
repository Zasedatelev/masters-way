/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SchemasCreateMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasCreateMessageResponse{}

// SchemasCreateMessageResponse struct for SchemasCreateMessageResponse
type SchemasCreateMessageResponse struct {
	Message SchemasMessageResponse `json:"message"`
	Users []string `json:"users"`
}

type _SchemasCreateMessageResponse SchemasCreateMessageResponse

// NewSchemasCreateMessageResponse instantiates a new SchemasCreateMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasCreateMessageResponse(message SchemasMessageResponse, users []string) *SchemasCreateMessageResponse {
	this := SchemasCreateMessageResponse{}
	this.Message = message
	this.Users = users
	return &this
}

// NewSchemasCreateMessageResponseWithDefaults instantiates a new SchemasCreateMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasCreateMessageResponseWithDefaults() *SchemasCreateMessageResponse {
	this := SchemasCreateMessageResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *SchemasCreateMessageResponse) GetMessage() SchemasMessageResponse {
	if o == nil {
		var ret SchemasMessageResponse
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateMessageResponse) GetMessageOk() (*SchemasMessageResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SchemasCreateMessageResponse) SetMessage(v SchemasMessageResponse) {
	o.Message = v
}

// GetUsers returns the Users field value
func (o *SchemasCreateMessageResponse) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateMessageResponse) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *SchemasCreateMessageResponse) SetUsers(v []string) {
	o.Users = v
}

func (o SchemasCreateMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasCreateMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *SchemasCreateMessageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"users",
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

	varSchemasCreateMessageResponse := _SchemasCreateMessageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasCreateMessageResponse)

	if err != nil {
		return err
	}

	*o = SchemasCreateMessageResponse(varSchemasCreateMessageResponse)

	return err
}

type NullableSchemasCreateMessageResponse struct {
	value *SchemasCreateMessageResponse
	isSet bool
}

func (v NullableSchemasCreateMessageResponse) Get() *SchemasCreateMessageResponse {
	return v.value
}

func (v *NullableSchemasCreateMessageResponse) Set(val *SchemasCreateMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasCreateMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasCreateMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasCreateMessageResponse(val *SchemasCreateMessageResponse) *NullableSchemasCreateMessageResponse {
	return &NullableSchemasCreateMessageResponse{value: val, isSet: true}
}

func (v NullableSchemasCreateMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasCreateMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

