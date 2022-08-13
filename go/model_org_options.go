/*
 * runZero API
 *
 * runZero Network Discovery API
 *
 * API version: 1.0.4
 * Contact: support@runzero.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// OrgOptions struct for OrgOptions
type OrgOptions struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewOrgOptions instantiates a new OrgOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgOptions() *OrgOptions {
	this := OrgOptions{}
	return &this
}

// NewOrgOptionsWithDefaults instantiates a new OrgOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgOptionsWithDefaults() *OrgOptions {
	this := OrgOptions{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrgOptions) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOptions) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrgOptions) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrgOptions) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrgOptions) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOptions) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrgOptions) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrgOptions) SetDescription(v string) {
	o.Description = &v
}

func (o OrgOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableOrgOptions struct {
	value *OrgOptions
	isSet bool
}

func (v NullableOrgOptions) Get() *OrgOptions {
	return v.value
}

func (v *NullableOrgOptions) Set(val *OrgOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgOptions(val *OrgOptions) *NullableOrgOptions {
	return &NullableOrgOptions{value: val, isSet: true}
}

func (v NullableOrgOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
