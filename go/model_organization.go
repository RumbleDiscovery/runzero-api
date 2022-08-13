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

// Organization struct for Organization
type Organization struct {
	Id                     string  `json:"id"`
	CreatedAt              *int64  `json:"created_at,omitempty"`
	UpdatedAt              *int64  `json:"updated_at,omitempty"`
	ClientId               *string `json:"client_id,omitempty"`
	DownloadToken          *string `json:"download_token,omitempty"`
	DownloadTokenCreatedAt *int64  `json:"download_token_created_at,omitempty"`
	Permanent              *bool   `json:"permanent,omitempty"`
	Name                   string  `json:"name"`
	Description            *string `json:"description,omitempty"`
	Inactive               *bool   `json:"inactive,omitempty"`
	DeactivatedAt          *int64  `json:"deactivated_at,omitempty"`
	ServiceCount           *int64  `json:"service_count,omitempty"`
	ServiceCountTcp        *int64  `json:"service_count_tcp,omitempty"`
	ServiceCountUdp        *int64  `json:"service_count_udp,omitempty"`
	ServiceCountArp        *int64  `json:"service_count_arp,omitempty"`
	ServiceCountIcmp       *int64  `json:"service_count_icmp,omitempty"`
	AssetCount             *int64  `json:"asset_count,omitempty"`
	ExportToken            *string `json:"export_token,omitempty"`
	ExportTokenCreatedAt   *int64  `json:"export_token_created_at,omitempty"`
	ExportTokenLastUsedAt  *int64  `json:"export_token_last_used_at,omitempty"`
	ExportTokenLastUsedBy  *string `json:"export_token_last_used_by,omitempty"`
	ExportTokenCounter     *int64  `json:"export_token_counter,omitempty"`
}

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization(id string, name string) *Organization {
	this := Organization{}
	this.Id = id
	this.Name = name
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	return &this
}

// GetId returns the Id field value
func (o *Organization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Organization) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Organization) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Organization) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Organization) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *Organization) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Organization) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Organization) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *Organization) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *Organization) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *Organization) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *Organization) SetClientId(v string) {
	o.ClientId = &v
}

// GetDownloadToken returns the DownloadToken field value if set, zero value otherwise.
func (o *Organization) GetDownloadToken() string {
	if o == nil || o.DownloadToken == nil {
		var ret string
		return ret
	}
	return *o.DownloadToken
}

// GetDownloadTokenOk returns a tuple with the DownloadToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetDownloadTokenOk() (*string, bool) {
	if o == nil || o.DownloadToken == nil {
		return nil, false
	}
	return o.DownloadToken, true
}

// HasDownloadToken returns a boolean if a field has been set.
func (o *Organization) HasDownloadToken() bool {
	if o != nil && o.DownloadToken != nil {
		return true
	}

	return false
}

// SetDownloadToken gets a reference to the given string and assigns it to the DownloadToken field.
func (o *Organization) SetDownloadToken(v string) {
	o.DownloadToken = &v
}

// GetDownloadTokenCreatedAt returns the DownloadTokenCreatedAt field value if set, zero value otherwise.
func (o *Organization) GetDownloadTokenCreatedAt() int64 {
	if o == nil || o.DownloadTokenCreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.DownloadTokenCreatedAt
}

// GetDownloadTokenCreatedAtOk returns a tuple with the DownloadTokenCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetDownloadTokenCreatedAtOk() (*int64, bool) {
	if o == nil || o.DownloadTokenCreatedAt == nil {
		return nil, false
	}
	return o.DownloadTokenCreatedAt, true
}

// HasDownloadTokenCreatedAt returns a boolean if a field has been set.
func (o *Organization) HasDownloadTokenCreatedAt() bool {
	if o != nil && o.DownloadTokenCreatedAt != nil {
		return true
	}

	return false
}

// SetDownloadTokenCreatedAt gets a reference to the given int64 and assigns it to the DownloadTokenCreatedAt field.
func (o *Organization) SetDownloadTokenCreatedAt(v int64) {
	o.DownloadTokenCreatedAt = &v
}

// GetPermanent returns the Permanent field value if set, zero value otherwise.
func (o *Organization) GetPermanent() bool {
	if o == nil || o.Permanent == nil {
		var ret bool
		return ret
	}
	return *o.Permanent
}

// GetPermanentOk returns a tuple with the Permanent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetPermanentOk() (*bool, bool) {
	if o == nil || o.Permanent == nil {
		return nil, false
	}
	return o.Permanent, true
}

// HasPermanent returns a boolean if a field has been set.
func (o *Organization) HasPermanent() bool {
	if o != nil && o.Permanent != nil {
		return true
	}

	return false
}

// SetPermanent gets a reference to the given bool and assigns it to the Permanent field.
func (o *Organization) SetPermanent(v bool) {
	o.Permanent = &v
}

// GetName returns the Name field value
func (o *Organization) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Organization) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Organization) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Organization) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Organization) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Organization) SetDescription(v string) {
	o.Description = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *Organization) GetInactive() bool {
	if o == nil || o.Inactive == nil {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetInactiveOk() (*bool, bool) {
	if o == nil || o.Inactive == nil {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *Organization) HasInactive() bool {
	if o != nil && o.Inactive != nil {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *Organization) SetInactive(v bool) {
	o.Inactive = &v
}

// GetDeactivatedAt returns the DeactivatedAt field value if set, zero value otherwise.
func (o *Organization) GetDeactivatedAt() int64 {
	if o == nil || o.DeactivatedAt == nil {
		var ret int64
		return ret
	}
	return *o.DeactivatedAt
}

// GetDeactivatedAtOk returns a tuple with the DeactivatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetDeactivatedAtOk() (*int64, bool) {
	if o == nil || o.DeactivatedAt == nil {
		return nil, false
	}
	return o.DeactivatedAt, true
}

// HasDeactivatedAt returns a boolean if a field has been set.
func (o *Organization) HasDeactivatedAt() bool {
	if o != nil && o.DeactivatedAt != nil {
		return true
	}

	return false
}

// SetDeactivatedAt gets a reference to the given int64 and assigns it to the DeactivatedAt field.
func (o *Organization) SetDeactivatedAt(v int64) {
	o.DeactivatedAt = &v
}

// GetServiceCount returns the ServiceCount field value if set, zero value otherwise.
func (o *Organization) GetServiceCount() int64 {
	if o == nil || o.ServiceCount == nil {
		var ret int64
		return ret
	}
	return *o.ServiceCount
}

// GetServiceCountOk returns a tuple with the ServiceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetServiceCountOk() (*int64, bool) {
	if o == nil || o.ServiceCount == nil {
		return nil, false
	}
	return o.ServiceCount, true
}

// HasServiceCount returns a boolean if a field has been set.
func (o *Organization) HasServiceCount() bool {
	if o != nil && o.ServiceCount != nil {
		return true
	}

	return false
}

// SetServiceCount gets a reference to the given int64 and assigns it to the ServiceCount field.
func (o *Organization) SetServiceCount(v int64) {
	o.ServiceCount = &v
}

// GetServiceCountTcp returns the ServiceCountTcp field value if set, zero value otherwise.
func (o *Organization) GetServiceCountTcp() int64 {
	if o == nil || o.ServiceCountTcp == nil {
		var ret int64
		return ret
	}
	return *o.ServiceCountTcp
}

// GetServiceCountTcpOk returns a tuple with the ServiceCountTcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetServiceCountTcpOk() (*int64, bool) {
	if o == nil || o.ServiceCountTcp == nil {
		return nil, false
	}
	return o.ServiceCountTcp, true
}

// HasServiceCountTcp returns a boolean if a field has been set.
func (o *Organization) HasServiceCountTcp() bool {
	if o != nil && o.ServiceCountTcp != nil {
		return true
	}

	return false
}

// SetServiceCountTcp gets a reference to the given int64 and assigns it to the ServiceCountTcp field.
func (o *Organization) SetServiceCountTcp(v int64) {
	o.ServiceCountTcp = &v
}

// GetServiceCountUdp returns the ServiceCountUdp field value if set, zero value otherwise.
func (o *Organization) GetServiceCountUdp() int64 {
	if o == nil || o.ServiceCountUdp == nil {
		var ret int64
		return ret
	}
	return *o.ServiceCountUdp
}

// GetServiceCountUdpOk returns a tuple with the ServiceCountUdp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetServiceCountUdpOk() (*int64, bool) {
	if o == nil || o.ServiceCountUdp == nil {
		return nil, false
	}
	return o.ServiceCountUdp, true
}

// HasServiceCountUdp returns a boolean if a field has been set.
func (o *Organization) HasServiceCountUdp() bool {
	if o != nil && o.ServiceCountUdp != nil {
		return true
	}

	return false
}

// SetServiceCountUdp gets a reference to the given int64 and assigns it to the ServiceCountUdp field.
func (o *Organization) SetServiceCountUdp(v int64) {
	o.ServiceCountUdp = &v
}

// GetServiceCountArp returns the ServiceCountArp field value if set, zero value otherwise.
func (o *Organization) GetServiceCountArp() int64 {
	if o == nil || o.ServiceCountArp == nil {
		var ret int64
		return ret
	}
	return *o.ServiceCountArp
}

// GetServiceCountArpOk returns a tuple with the ServiceCountArp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetServiceCountArpOk() (*int64, bool) {
	if o == nil || o.ServiceCountArp == nil {
		return nil, false
	}
	return o.ServiceCountArp, true
}

// HasServiceCountArp returns a boolean if a field has been set.
func (o *Organization) HasServiceCountArp() bool {
	if o != nil && o.ServiceCountArp != nil {
		return true
	}

	return false
}

// SetServiceCountArp gets a reference to the given int64 and assigns it to the ServiceCountArp field.
func (o *Organization) SetServiceCountArp(v int64) {
	o.ServiceCountArp = &v
}

// GetServiceCountIcmp returns the ServiceCountIcmp field value if set, zero value otherwise.
func (o *Organization) GetServiceCountIcmp() int64 {
	if o == nil || o.ServiceCountIcmp == nil {
		var ret int64
		return ret
	}
	return *o.ServiceCountIcmp
}

// GetServiceCountIcmpOk returns a tuple with the ServiceCountIcmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetServiceCountIcmpOk() (*int64, bool) {
	if o == nil || o.ServiceCountIcmp == nil {
		return nil, false
	}
	return o.ServiceCountIcmp, true
}

// HasServiceCountIcmp returns a boolean if a field has been set.
func (o *Organization) HasServiceCountIcmp() bool {
	if o != nil && o.ServiceCountIcmp != nil {
		return true
	}

	return false
}

// SetServiceCountIcmp gets a reference to the given int64 and assigns it to the ServiceCountIcmp field.
func (o *Organization) SetServiceCountIcmp(v int64) {
	o.ServiceCountIcmp = &v
}

// GetAssetCount returns the AssetCount field value if set, zero value otherwise.
func (o *Organization) GetAssetCount() int64 {
	if o == nil || o.AssetCount == nil {
		var ret int64
		return ret
	}
	return *o.AssetCount
}

// GetAssetCountOk returns a tuple with the AssetCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetAssetCountOk() (*int64, bool) {
	if o == nil || o.AssetCount == nil {
		return nil, false
	}
	return o.AssetCount, true
}

// HasAssetCount returns a boolean if a field has been set.
func (o *Organization) HasAssetCount() bool {
	if o != nil && o.AssetCount != nil {
		return true
	}

	return false
}

// SetAssetCount gets a reference to the given int64 and assigns it to the AssetCount field.
func (o *Organization) SetAssetCount(v int64) {
	o.AssetCount = &v
}

// GetExportToken returns the ExportToken field value if set, zero value otherwise.
func (o *Organization) GetExportToken() string {
	if o == nil || o.ExportToken == nil {
		var ret string
		return ret
	}
	return *o.ExportToken
}

// GetExportTokenOk returns a tuple with the ExportToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetExportTokenOk() (*string, bool) {
	if o == nil || o.ExportToken == nil {
		return nil, false
	}
	return o.ExportToken, true
}

// HasExportToken returns a boolean if a field has been set.
func (o *Organization) HasExportToken() bool {
	if o != nil && o.ExportToken != nil {
		return true
	}

	return false
}

// SetExportToken gets a reference to the given string and assigns it to the ExportToken field.
func (o *Organization) SetExportToken(v string) {
	o.ExportToken = &v
}

// GetExportTokenCreatedAt returns the ExportTokenCreatedAt field value if set, zero value otherwise.
func (o *Organization) GetExportTokenCreatedAt() int64 {
	if o == nil || o.ExportTokenCreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.ExportTokenCreatedAt
}

// GetExportTokenCreatedAtOk returns a tuple with the ExportTokenCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetExportTokenCreatedAtOk() (*int64, bool) {
	if o == nil || o.ExportTokenCreatedAt == nil {
		return nil, false
	}
	return o.ExportTokenCreatedAt, true
}

// HasExportTokenCreatedAt returns a boolean if a field has been set.
func (o *Organization) HasExportTokenCreatedAt() bool {
	if o != nil && o.ExportTokenCreatedAt != nil {
		return true
	}

	return false
}

// SetExportTokenCreatedAt gets a reference to the given int64 and assigns it to the ExportTokenCreatedAt field.
func (o *Organization) SetExportTokenCreatedAt(v int64) {
	o.ExportTokenCreatedAt = &v
}

// GetExportTokenLastUsedAt returns the ExportTokenLastUsedAt field value if set, zero value otherwise.
func (o *Organization) GetExportTokenLastUsedAt() int64 {
	if o == nil || o.ExportTokenLastUsedAt == nil {
		var ret int64
		return ret
	}
	return *o.ExportTokenLastUsedAt
}

// GetExportTokenLastUsedAtOk returns a tuple with the ExportTokenLastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetExportTokenLastUsedAtOk() (*int64, bool) {
	if o == nil || o.ExportTokenLastUsedAt == nil {
		return nil, false
	}
	return o.ExportTokenLastUsedAt, true
}

// HasExportTokenLastUsedAt returns a boolean if a field has been set.
func (o *Organization) HasExportTokenLastUsedAt() bool {
	if o != nil && o.ExportTokenLastUsedAt != nil {
		return true
	}

	return false
}

// SetExportTokenLastUsedAt gets a reference to the given int64 and assigns it to the ExportTokenLastUsedAt field.
func (o *Organization) SetExportTokenLastUsedAt(v int64) {
	o.ExportTokenLastUsedAt = &v
}

// GetExportTokenLastUsedBy returns the ExportTokenLastUsedBy field value if set, zero value otherwise.
func (o *Organization) GetExportTokenLastUsedBy() string {
	if o == nil || o.ExportTokenLastUsedBy == nil {
		var ret string
		return ret
	}
	return *o.ExportTokenLastUsedBy
}

// GetExportTokenLastUsedByOk returns a tuple with the ExportTokenLastUsedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetExportTokenLastUsedByOk() (*string, bool) {
	if o == nil || o.ExportTokenLastUsedBy == nil {
		return nil, false
	}
	return o.ExportTokenLastUsedBy, true
}

// HasExportTokenLastUsedBy returns a boolean if a field has been set.
func (o *Organization) HasExportTokenLastUsedBy() bool {
	if o != nil && o.ExportTokenLastUsedBy != nil {
		return true
	}

	return false
}

// SetExportTokenLastUsedBy gets a reference to the given string and assigns it to the ExportTokenLastUsedBy field.
func (o *Organization) SetExportTokenLastUsedBy(v string) {
	o.ExportTokenLastUsedBy = &v
}

// GetExportTokenCounter returns the ExportTokenCounter field value if set, zero value otherwise.
func (o *Organization) GetExportTokenCounter() int64 {
	if o == nil || o.ExportTokenCounter == nil {
		var ret int64
		return ret
	}
	return *o.ExportTokenCounter
}

// GetExportTokenCounterOk returns a tuple with the ExportTokenCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetExportTokenCounterOk() (*int64, bool) {
	if o == nil || o.ExportTokenCounter == nil {
		return nil, false
	}
	return o.ExportTokenCounter, true
}

// HasExportTokenCounter returns a boolean if a field has been set.
func (o *Organization) HasExportTokenCounter() bool {
	if o != nil && o.ExportTokenCounter != nil {
		return true
	}

	return false
}

// SetExportTokenCounter gets a reference to the given int64 and assigns it to the ExportTokenCounter field.
func (o *Organization) SetExportTokenCounter(v int64) {
	o.ExportTokenCounter = &v
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.DownloadToken != nil {
		toSerialize["download_token"] = o.DownloadToken
	}
	if o.DownloadTokenCreatedAt != nil {
		toSerialize["download_token_created_at"] = o.DownloadTokenCreatedAt
	}
	if o.Permanent != nil {
		toSerialize["permanent"] = o.Permanent
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Inactive != nil {
		toSerialize["inactive"] = o.Inactive
	}
	if o.DeactivatedAt != nil {
		toSerialize["deactivated_at"] = o.DeactivatedAt
	}
	if o.ServiceCount != nil {
		toSerialize["service_count"] = o.ServiceCount
	}
	if o.ServiceCountTcp != nil {
		toSerialize["service_count_tcp"] = o.ServiceCountTcp
	}
	if o.ServiceCountUdp != nil {
		toSerialize["service_count_udp"] = o.ServiceCountUdp
	}
	if o.ServiceCountArp != nil {
		toSerialize["service_count_arp"] = o.ServiceCountArp
	}
	if o.ServiceCountIcmp != nil {
		toSerialize["service_count_icmp"] = o.ServiceCountIcmp
	}
	if o.AssetCount != nil {
		toSerialize["asset_count"] = o.AssetCount
	}
	if o.ExportToken != nil {
		toSerialize["export_token"] = o.ExportToken
	}
	if o.ExportTokenCreatedAt != nil {
		toSerialize["export_token_created_at"] = o.ExportTokenCreatedAt
	}
	if o.ExportTokenLastUsedAt != nil {
		toSerialize["export_token_last_used_at"] = o.ExportTokenLastUsedAt
	}
	if o.ExportTokenLastUsedBy != nil {
		toSerialize["export_token_last_used_by"] = o.ExportTokenLastUsedBy
	}
	if o.ExportTokenCounter != nil {
		toSerialize["export_token_counter"] = o.ExportTokenCounter
	}
	return json.Marshal(toSerialize)
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
