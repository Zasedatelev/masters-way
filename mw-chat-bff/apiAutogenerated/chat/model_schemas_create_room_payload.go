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

// checks if the SchemasCreateRoomPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasCreateRoomPayload{}

// SchemasCreateRoomPayload struct for SchemasCreateRoomPayload
type SchemasCreateRoomPayload struct {
	Name NullableString `json:"name,omitempty"`
	RoomType string `json:"roomType"`
	UserId NullableString `json:"userId,omitempty"`
}

type _SchemasCreateRoomPayload SchemasCreateRoomPayload

// NewSchemasCreateRoomPayload instantiates a new SchemasCreateRoomPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasCreateRoomPayload(roomType string) *SchemasCreateRoomPayload {
	this := SchemasCreateRoomPayload{}
	this.RoomType = roomType
	return &this
}

// NewSchemasCreateRoomPayloadWithDefaults instantiates a new SchemasCreateRoomPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasCreateRoomPayloadWithDefaults() *SchemasCreateRoomPayload {
	this := SchemasCreateRoomPayload{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasCreateRoomPayload) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasCreateRoomPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SchemasCreateRoomPayload) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SchemasCreateRoomPayload) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SchemasCreateRoomPayload) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SchemasCreateRoomPayload) UnsetName() {
	o.Name.Unset()
}

// GetRoomType returns the RoomType field value
func (o *SchemasCreateRoomPayload) GetRoomType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateRoomPayload) GetRoomTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoomType, true
}

// SetRoomType sets field value
func (o *SchemasCreateRoomPayload) SetRoomType(v string) {
	o.RoomType = v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasCreateRoomPayload) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasCreateRoomPayload) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *SchemasCreateRoomPayload) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *SchemasCreateRoomPayload) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *SchemasCreateRoomPayload) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *SchemasCreateRoomPayload) UnsetUserId() {
	o.UserId.Unset()
}

func (o SchemasCreateRoomPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasCreateRoomPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["roomType"] = o.RoomType
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	return toSerialize, nil
}

func (o *SchemasCreateRoomPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roomType",
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

	varSchemasCreateRoomPayload := _SchemasCreateRoomPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasCreateRoomPayload)

	if err != nil {
		return err
	}

	*o = SchemasCreateRoomPayload(varSchemasCreateRoomPayload)

	return err
}

type NullableSchemasCreateRoomPayload struct {
	value *SchemasCreateRoomPayload
	isSet bool
}

func (v NullableSchemasCreateRoomPayload) Get() *SchemasCreateRoomPayload {
	return v.value
}

func (v *NullableSchemasCreateRoomPayload) Set(val *SchemasCreateRoomPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasCreateRoomPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasCreateRoomPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasCreateRoomPayload(val *SchemasCreateRoomPayload) *NullableSchemasCreateRoomPayload {
	return &NullableSchemasCreateRoomPayload{value: val, isSet: true}
}

func (v NullableSchemasCreateRoomPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasCreateRoomPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


