/*
Technitium API

Go SDK for interacting with Technitium APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 12.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package technitium

import (
	"encoding/json"
)

// checks if the CommonResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonResponse{}

// CommonResponse Describes the result of processing a request
type CommonResponse struct {
	Status *string `json:"status,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonResponse CommonResponse

// NewCommonResponse instantiates a new CommonResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonResponse() *CommonResponse {
	this := CommonResponse{}
	return &this
}

// NewCommonResponseWithDefaults instantiates a new CommonResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonResponseWithDefaults() *CommonResponse {
	this := CommonResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CommonResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CommonResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CommonResponse) SetStatus(v string) {
	o.Status = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *CommonResponse) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *CommonResponse) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *CommonResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o CommonResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonResponse) UnmarshalJSON(data []byte) (err error) {
	varCommonResponse := _CommonResponse{}

	err = json.Unmarshal(data, &varCommonResponse)

	if err != nil {
		return err
	}

	*o = CommonResponse(varCommonResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "errorMessage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonResponse struct {
	value *CommonResponse
	isSet bool
}

func (v NullableCommonResponse) Get() *CommonResponse {
	return v.value
}

func (v *NullableCommonResponse) Set(val *CommonResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonResponse(val *CommonResponse) *NullableCommonResponse {
	return &NullableCommonResponse{value: val, isSet: true}
}

func (v NullableCommonResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


