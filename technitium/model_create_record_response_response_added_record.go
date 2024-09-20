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

// checks if the CreateRecordResponseResponseAddedRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRecordResponseResponseAddedRecord{}

// CreateRecordResponseResponseAddedRecord struct for CreateRecordResponseResponseAddedRecord
type CreateRecordResponseResponseAddedRecord struct {
	Disabled *bool `json:"disabled,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Ttl *int32 `json:"ttl,omitempty"`
	RData *CreateRecordResponseResponseAddedRecordRData `json:"rData,omitempty"`
	DnssecStatus *string `json:"dnssecStatus,omitempty"`
	LastUsedOn *string `json:"lastUsedOn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateRecordResponseResponseAddedRecord CreateRecordResponseResponseAddedRecord

// NewCreateRecordResponseResponseAddedRecord instantiates a new CreateRecordResponseResponseAddedRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRecordResponseResponseAddedRecord() *CreateRecordResponseResponseAddedRecord {
	this := CreateRecordResponseResponseAddedRecord{}
	return &this
}

// NewCreateRecordResponseResponseAddedRecordWithDefaults instantiates a new CreateRecordResponseResponseAddedRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRecordResponseResponseAddedRecordWithDefaults() *CreateRecordResponseResponseAddedRecord {
	this := CreateRecordResponseResponseAddedRecord{}
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *CreateRecordResponseResponseAddedRecord) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordResponseResponseAddedRecord) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *CreateRecordResponseResponseAddedRecord) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *CreateRecordResponseResponseAddedRecord) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateRecordResponseResponseAddedRecord) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordResponseResponseAddedRecord) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateRecordResponseResponseAddedRecord) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateRecordResponseResponseAddedRecord) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateRecordResponseResponseAddedRecord) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordResponseResponseAddedRecord) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateRecordResponseResponseAddedRecord) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateRecordResponseResponseAddedRecord) SetType(v string) {
	o.Type = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *CreateRecordResponseResponseAddedRecord) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordResponseResponseAddedRecord) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *CreateRecordResponseResponseAddedRecord) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *CreateRecordResponseResponseAddedRecord) SetTtl(v int32) {
	o.Ttl = &v
}

// GetRData returns the RData field value if set, zero value otherwise.
func (o *CreateRecordResponseResponseAddedRecord) GetRData() CreateRecordResponseResponseAddedRecordRData {
	if o == nil || IsNil(o.RData) {
		var ret CreateRecordResponseResponseAddedRecordRData
		return ret
	}
	return *o.RData
}

// GetRDataOk returns a tuple with the RData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordResponseResponseAddedRecord) GetRDataOk() (*CreateRecordResponseResponseAddedRecordRData, bool) {
	if o == nil || IsNil(o.RData) {
		return nil, false
	}
	return o.RData, true
}

// HasRData returns a boolean if a field has been set.
func (o *CreateRecordResponseResponseAddedRecord) HasRData() bool {
	if o != nil && !IsNil(o.RData) {
		return true
	}

	return false
}

// SetRData gets a reference to the given CreateRecordResponseResponseAddedRecordRData and assigns it to the RData field.
func (o *CreateRecordResponseResponseAddedRecord) SetRData(v CreateRecordResponseResponseAddedRecordRData) {
	o.RData = &v
}

// GetDnssecStatus returns the DnssecStatus field value if set, zero value otherwise.
func (o *CreateRecordResponseResponseAddedRecord) GetDnssecStatus() string {
	if o == nil || IsNil(o.DnssecStatus) {
		var ret string
		return ret
	}
	return *o.DnssecStatus
}

// GetDnssecStatusOk returns a tuple with the DnssecStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordResponseResponseAddedRecord) GetDnssecStatusOk() (*string, bool) {
	if o == nil || IsNil(o.DnssecStatus) {
		return nil, false
	}
	return o.DnssecStatus, true
}

// HasDnssecStatus returns a boolean if a field has been set.
func (o *CreateRecordResponseResponseAddedRecord) HasDnssecStatus() bool {
	if o != nil && !IsNil(o.DnssecStatus) {
		return true
	}

	return false
}

// SetDnssecStatus gets a reference to the given string and assigns it to the DnssecStatus field.
func (o *CreateRecordResponseResponseAddedRecord) SetDnssecStatus(v string) {
	o.DnssecStatus = &v
}

// GetLastUsedOn returns the LastUsedOn field value if set, zero value otherwise.
func (o *CreateRecordResponseResponseAddedRecord) GetLastUsedOn() string {
	if o == nil || IsNil(o.LastUsedOn) {
		var ret string
		return ret
	}
	return *o.LastUsedOn
}

// GetLastUsedOnOk returns a tuple with the LastUsedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordResponseResponseAddedRecord) GetLastUsedOnOk() (*string, bool) {
	if o == nil || IsNil(o.LastUsedOn) {
		return nil, false
	}
	return o.LastUsedOn, true
}

// HasLastUsedOn returns a boolean if a field has been set.
func (o *CreateRecordResponseResponseAddedRecord) HasLastUsedOn() bool {
	if o != nil && !IsNil(o.LastUsedOn) {
		return true
	}

	return false
}

// SetLastUsedOn gets a reference to the given string and assigns it to the LastUsedOn field.
func (o *CreateRecordResponseResponseAddedRecord) SetLastUsedOn(v string) {
	o.LastUsedOn = &v
}

func (o CreateRecordResponseResponseAddedRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRecordResponseResponseAddedRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.RData) {
		toSerialize["rData"] = o.RData
	}
	if !IsNil(o.DnssecStatus) {
		toSerialize["dnssecStatus"] = o.DnssecStatus
	}
	if !IsNil(o.LastUsedOn) {
		toSerialize["lastUsedOn"] = o.LastUsedOn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateRecordResponseResponseAddedRecord) UnmarshalJSON(data []byte) (err error) {
	varCreateRecordResponseResponseAddedRecord := _CreateRecordResponseResponseAddedRecord{}

	err = json.Unmarshal(data, &varCreateRecordResponseResponseAddedRecord)

	if err != nil {
		return err
	}

	*o = CreateRecordResponseResponseAddedRecord(varCreateRecordResponseResponseAddedRecord)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "ttl")
		delete(additionalProperties, "rData")
		delete(additionalProperties, "dnssecStatus")
		delete(additionalProperties, "lastUsedOn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateRecordResponseResponseAddedRecord struct {
	value *CreateRecordResponseResponseAddedRecord
	isSet bool
}

func (v NullableCreateRecordResponseResponseAddedRecord) Get() *CreateRecordResponseResponseAddedRecord {
	return v.value
}

func (v *NullableCreateRecordResponseResponseAddedRecord) Set(val *CreateRecordResponseResponseAddedRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecordResponseResponseAddedRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecordResponseResponseAddedRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecordResponseResponseAddedRecord(val *CreateRecordResponseResponseAddedRecord) *NullableCreateRecordResponseResponseAddedRecord {
	return &NullableCreateRecordResponseResponseAddedRecord{value: val, isSet: true}
}

func (v NullableCreateRecordResponseResponseAddedRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecordResponseResponseAddedRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


