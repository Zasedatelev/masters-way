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

// checks if the SchemasCreateCommentPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasCreateCommentPayload{}

// SchemasCreateCommentPayload struct for SchemasCreateCommentPayload
type SchemasCreateCommentPayload struct {
	DayReportUuid string
	Description string
	OwnerUuid string
}

type _SchemasCreateCommentPayload SchemasCreateCommentPayload

// NewSchemasCreateCommentPayload instantiates a new SchemasCreateCommentPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasCreateCommentPayload(dayReportUuid string, description string, ownerUuid string) *SchemasCreateCommentPayload {
	this := SchemasCreateCommentPayload{}
	this.DayReportUuid = dayReportUuid
	this.Description = description
	this.OwnerUuid = ownerUuid
	return &this
}

// NewSchemasCreateCommentPayloadWithDefaults instantiates a new SchemasCreateCommentPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasCreateCommentPayloadWithDefaults() *SchemasCreateCommentPayload {
	this := SchemasCreateCommentPayload{}
	return &this
}

// GetDayReportUuid returns the DayReportUuid field value
func (o *SchemasCreateCommentPayload) GetDayReportUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DayReportUuid
}

// GetDayReportUuidOk returns a tuple with the DayReportUuid field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateCommentPayload) GetDayReportUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayReportUuid, true
}

// SetDayReportUuid sets field value
func (o *SchemasCreateCommentPayload) SetDayReportUuid(v string) {
	o.DayReportUuid = v
}

// GetDescription returns the Description field value
func (o *SchemasCreateCommentPayload) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateCommentPayload) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SchemasCreateCommentPayload) SetDescription(v string) {
	o.Description = v
}

// GetOwnerUuid returns the OwnerUuid field value
func (o *SchemasCreateCommentPayload) GetOwnerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerUuid
}

// GetOwnerUuidOk returns a tuple with the OwnerUuid field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateCommentPayload) GetOwnerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerUuid, true
}

// SetOwnerUuid sets field value
func (o *SchemasCreateCommentPayload) SetOwnerUuid(v string) {
	o.OwnerUuid = v
}

func (o SchemasCreateCommentPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasCreateCommentPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dayReportUuid"] = o.DayReportUuid
	toSerialize["description"] = o.Description
	toSerialize["ownerUuid"] = o.OwnerUuid
	return toSerialize, nil
}

func (o *SchemasCreateCommentPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dayReportUuid",
		"description",
		"ownerUuid",
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

	varSchemasCreateCommentPayload := _SchemasCreateCommentPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasCreateCommentPayload)

	if err != nil {
		return err
	}

	*o = SchemasCreateCommentPayload(varSchemasCreateCommentPayload)

	return err
}

type NullableSchemasCreateCommentPayload struct {
	value *SchemasCreateCommentPayload
	isSet bool
}

func (v NullableSchemasCreateCommentPayload) Get() *SchemasCreateCommentPayload {
	return v.value
}

func (v *NullableSchemasCreateCommentPayload) Set(val *SchemasCreateCommentPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasCreateCommentPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasCreateCommentPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasCreateCommentPayload(val *SchemasCreateCommentPayload) *NullableSchemasCreateCommentPayload {
	return &NullableSchemasCreateCommentPayload{value: val, isSet: true}
}

func (v NullableSchemasCreateCommentPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasCreateCommentPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


