/*
Masters way general API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SchemasUpdatePlanPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasUpdatePlanPayload{}

// SchemasUpdatePlanPayload struct for SchemasUpdatePlanPayload
type SchemasUpdatePlanPayload struct {
	Description *string `json:"description,omitempty"`
	IsDone *bool `json:"isDone,omitempty"`
	Time *int32 `json:"time,omitempty"`
}

// NewSchemasUpdatePlanPayload instantiates a new SchemasUpdatePlanPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasUpdatePlanPayload() *SchemasUpdatePlanPayload {
	this := SchemasUpdatePlanPayload{}
	return &this
}

// NewSchemasUpdatePlanPayloadWithDefaults instantiates a new SchemasUpdatePlanPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasUpdatePlanPayloadWithDefaults() *SchemasUpdatePlanPayload {
	this := SchemasUpdatePlanPayload{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SchemasUpdatePlanPayload) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdatePlanPayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SchemasUpdatePlanPayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SchemasUpdatePlanPayload) SetDescription(v string) {
	o.Description = &v
}

// GetIsDone returns the IsDone field value if set, zero value otherwise.
func (o *SchemasUpdatePlanPayload) GetIsDone() bool {
	if o == nil || IsNil(o.IsDone) {
		var ret bool
		return ret
	}
	return *o.IsDone
}

// GetIsDoneOk returns a tuple with the IsDone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdatePlanPayload) GetIsDoneOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDone) {
		return nil, false
	}
	return o.IsDone, true
}

// HasIsDone returns a boolean if a field has been set.
func (o *SchemasUpdatePlanPayload) HasIsDone() bool {
	if o != nil && !IsNil(o.IsDone) {
		return true
	}

	return false
}

// SetIsDone gets a reference to the given bool and assigns it to the IsDone field.
func (o *SchemasUpdatePlanPayload) SetIsDone(v bool) {
	o.IsDone = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *SchemasUpdatePlanPayload) GetTime() int32 {
	if o == nil || IsNil(o.Time) {
		var ret int32
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdatePlanPayload) GetTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *SchemasUpdatePlanPayload) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int32 and assigns it to the Time field.
func (o *SchemasUpdatePlanPayload) SetTime(v int32) {
	o.Time = &v
}

func (o SchemasUpdatePlanPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasUpdatePlanPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsDone) {
		toSerialize["isDone"] = o.IsDone
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableSchemasUpdatePlanPayload struct {
	value *SchemasUpdatePlanPayload
	isSet bool
}

func (v NullableSchemasUpdatePlanPayload) Get() *SchemasUpdatePlanPayload {
	return v.value
}

func (v *NullableSchemasUpdatePlanPayload) Set(val *SchemasUpdatePlanPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasUpdatePlanPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasUpdatePlanPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasUpdatePlanPayload(val *SchemasUpdatePlanPayload) *NullableSchemasUpdatePlanPayload {
	return &NullableSchemasUpdatePlanPayload{value: val, isSet: true}
}

func (v NullableSchemasUpdatePlanPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasUpdatePlanPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


