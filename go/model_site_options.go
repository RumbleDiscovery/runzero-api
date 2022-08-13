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

// SiteOptions struct for SiteOptions
type SiteOptions struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Scope       *string `json:"scope,omitempty"`
	Excludes    *string `json:"excludes,omitempty"`
}

// NewSiteOptions instantiates a new SiteOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteOptions(name string) *SiteOptions {
	this := SiteOptions{}
	this.Name = name
	return &this
}

// NewSiteOptionsWithDefaults instantiates a new SiteOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteOptionsWithDefaults() *SiteOptions {
	this := SiteOptions{}
	return &this
}

// GetName returns the Name field value
func (o *SiteOptions) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteOptions) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteOptions) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SiteOptions) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteOptions) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SiteOptions) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SiteOptions) SetDescription(v string) {
	o.Description = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *SiteOptions) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteOptions) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *SiteOptions) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *SiteOptions) SetScope(v string) {
	o.Scope = &v
}

// GetExcludes returns the Excludes field value if set, zero value otherwise.
func (o *SiteOptions) GetExcludes() string {
	if o == nil || o.Excludes == nil {
		var ret string
		return ret
	}
	return *o.Excludes
}

// GetExcludesOk returns a tuple with the Excludes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteOptions) GetExcludesOk() (*string, bool) {
	if o == nil || o.Excludes == nil {
		return nil, false
	}
	return o.Excludes, true
}

// HasExcludes returns a boolean if a field has been set.
func (o *SiteOptions) HasExcludes() bool {
	if o != nil && o.Excludes != nil {
		return true
	}

	return false
}

// SetExcludes gets a reference to the given string and assigns it to the Excludes field.
func (o *SiteOptions) SetExcludes(v string) {
	o.Excludes = &v
}

func (o SiteOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Excludes != nil {
		toSerialize["excludes"] = o.Excludes
	}
	return json.Marshal(toSerialize)
}

type NullableSiteOptions struct {
	value *SiteOptions
	isSet bool
}

func (v NullableSiteOptions) Get() *SiteOptions {
	return v.value
}

func (v *NullableSiteOptions) Set(val *SiteOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteOptions(val *SiteOptions) *NullableSiteOptions {
	return &NullableSiteOptions{value: val, isSet: true}
}

func (v NullableSiteOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
