/*
Masters way API

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

// checks if the SchemasGenerateMetricsPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasGenerateMetricsPayload{}

// SchemasGenerateMetricsPayload struct for SchemasGenerateMetricsPayload
type SchemasGenerateMetricsPayload struct {
	GoalDescription string `json:"goalDescription"`
	Metrics []string `json:"metrics"`
	WayName string `json:"wayName"`
}

type _SchemasGenerateMetricsPayload SchemasGenerateMetricsPayload

// NewSchemasGenerateMetricsPayload instantiates a new SchemasGenerateMetricsPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasGenerateMetricsPayload(goalDescription string, metrics []string, wayName string) *SchemasGenerateMetricsPayload {
	this := SchemasGenerateMetricsPayload{}
	this.GoalDescription = goalDescription
	this.Metrics = metrics
	this.WayName = wayName
	return &this
}

// NewSchemasGenerateMetricsPayloadWithDefaults instantiates a new SchemasGenerateMetricsPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasGenerateMetricsPayloadWithDefaults() *SchemasGenerateMetricsPayload {
	this := SchemasGenerateMetricsPayload{}
	return &this
}

// GetGoalDescription returns the GoalDescription field value
func (o *SchemasGenerateMetricsPayload) GetGoalDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GoalDescription
}

// GetGoalDescriptionOk returns a tuple with the GoalDescription field value
// and a boolean to check if the value has been set.
func (o *SchemasGenerateMetricsPayload) GetGoalDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoalDescription, true
}

// SetGoalDescription sets field value
func (o *SchemasGenerateMetricsPayload) SetGoalDescription(v string) {
	o.GoalDescription = v
}

// GetMetrics returns the Metrics field value
func (o *SchemasGenerateMetricsPayload) GetMetrics() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *SchemasGenerateMetricsPayload) GetMetricsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metrics, true
}

// SetMetrics sets field value
func (o *SchemasGenerateMetricsPayload) SetMetrics(v []string) {
	o.Metrics = v
}

// GetWayName returns the WayName field value
func (o *SchemasGenerateMetricsPayload) GetWayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WayName
}

// GetWayNameOk returns a tuple with the WayName field value
// and a boolean to check if the value has been set.
func (o *SchemasGenerateMetricsPayload) GetWayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WayName, true
}

// SetWayName sets field value
func (o *SchemasGenerateMetricsPayload) SetWayName(v string) {
	o.WayName = v
}

func (o SchemasGenerateMetricsPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasGenerateMetricsPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["goalDescription"] = o.GoalDescription
	toSerialize["metrics"] = o.Metrics
	toSerialize["wayName"] = o.WayName
	return toSerialize, nil
}

func (o *SchemasGenerateMetricsPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"goalDescription",
		"metrics",
		"wayName",
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

	varSchemasGenerateMetricsPayload := _SchemasGenerateMetricsPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasGenerateMetricsPayload)

	if err != nil {
		return err
	}

	*o = SchemasGenerateMetricsPayload(varSchemasGenerateMetricsPayload)

	return err
}

type NullableSchemasGenerateMetricsPayload struct {
	value *SchemasGenerateMetricsPayload
	isSet bool
}

func (v NullableSchemasGenerateMetricsPayload) Get() *SchemasGenerateMetricsPayload {
	return v.value
}

func (v *NullableSchemasGenerateMetricsPayload) Set(val *SchemasGenerateMetricsPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasGenerateMetricsPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasGenerateMetricsPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasGenerateMetricsPayload(val *SchemasGenerateMetricsPayload) *NullableSchemasGenerateMetricsPayload {
	return &NullableSchemasGenerateMetricsPayload{value: val, isSet: true}
}

func (v NullableSchemasGenerateMetricsPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasGenerateMetricsPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

