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

// checks if the SchemasLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasLabel{}

// SchemasLabel struct for SchemasLabel
type SchemasLabel struct {
	Color string `json:"color"`
	Description string `json:"description"`
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

type _SchemasLabel SchemasLabel

// NewSchemasLabel instantiates a new SchemasLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasLabel(color string, description string, name string, uuid string) *SchemasLabel {
	this := SchemasLabel{}
	this.Color = color
	this.Description = description
	this.Name = name
	this.Uuid = uuid
	return &this
}

// NewSchemasLabelWithDefaults instantiates a new SchemasLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasLabelWithDefaults() *SchemasLabel {
	this := SchemasLabel{}
	return &this
}

// GetColor returns the Color field value
func (o *SchemasLabel) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *SchemasLabel) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *SchemasLabel) SetColor(v string) {
	o.Color = v
}

// GetDescription returns the Description field value
func (o *SchemasLabel) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SchemasLabel) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SchemasLabel) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *SchemasLabel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemasLabel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemasLabel) SetName(v string) {
	o.Name = v
}

// GetUuid returns the Uuid field value
func (o *SchemasLabel) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *SchemasLabel) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *SchemasLabel) SetUuid(v string) {
	o.Uuid = v
}

func (o SchemasLabel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["color"] = o.Color
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	toSerialize["uuid"] = o.Uuid
	return toSerialize, nil
}

func (o *SchemasLabel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"color",
		"description",
		"name",
		"uuid",
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

	varSchemasLabel := _SchemasLabel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasLabel)

	if err != nil {
		return err
	}

	*o = SchemasLabel(varSchemasLabel)

	return err
}

type NullableSchemasLabel struct {
	value *SchemasLabel
	isSet bool
}

func (v NullableSchemasLabel) Get() *SchemasLabel {
	return v.value
}

func (v *NullableSchemasLabel) Set(val *SchemasLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasLabel(val *SchemasLabel) *NullableSchemasLabel {
	return &NullableSchemasLabel{value: val, isSet: true}
}

func (v NullableSchemasLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


