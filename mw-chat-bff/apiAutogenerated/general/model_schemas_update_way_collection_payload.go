/*
Masters way API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SchemasUpdateWayCollectionPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasUpdateWayCollectionPayload{}

// SchemasUpdateWayCollectionPayload struct for SchemasUpdateWayCollectionPayload
type SchemasUpdateWayCollectionPayload struct {
	Name *string `json:"name,omitempty"`
}

// NewSchemasUpdateWayCollectionPayload instantiates a new SchemasUpdateWayCollectionPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasUpdateWayCollectionPayload() *SchemasUpdateWayCollectionPayload {
	this := SchemasUpdateWayCollectionPayload{}
	return &this
}

// NewSchemasUpdateWayCollectionPayloadWithDefaults instantiates a new SchemasUpdateWayCollectionPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasUpdateWayCollectionPayloadWithDefaults() *SchemasUpdateWayCollectionPayload {
	this := SchemasUpdateWayCollectionPayload{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SchemasUpdateWayCollectionPayload) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdateWayCollectionPayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SchemasUpdateWayCollectionPayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SchemasUpdateWayCollectionPayload) SetName(v string) {
	o.Name = &v
}

func (o SchemasUpdateWayCollectionPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasUpdateWayCollectionPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableSchemasUpdateWayCollectionPayload struct {
	value *SchemasUpdateWayCollectionPayload
	isSet bool
}

func (v NullableSchemasUpdateWayCollectionPayload) Get() *SchemasUpdateWayCollectionPayload {
	return v.value
}

func (v *NullableSchemasUpdateWayCollectionPayload) Set(val *SchemasUpdateWayCollectionPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasUpdateWayCollectionPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasUpdateWayCollectionPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasUpdateWayCollectionPayload(val *SchemasUpdateWayCollectionPayload) *NullableSchemasUpdateWayCollectionPayload {
	return &NullableSchemasUpdateWayCollectionPayload{value: val, isSet: true}
}

func (v NullableSchemasUpdateWayCollectionPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasUpdateWayCollectionPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


