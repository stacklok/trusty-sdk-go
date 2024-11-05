/*
Trusty API

Trusty API

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the HistoricalProvenance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoricalProvenance{}

// HistoricalProvenance [Historical provenance](provenance.md#historical-provenance-hp) This contains the number of `tags` in the repo, the number of `versions` of the package, a count of the `common` tags and the ratio of tags to common as `overlap`.
type HistoricalProvenance struct {
	Overlap *float32 `json:"overlap,omitempty"`
	Common *float32 `json:"common,omitempty"`
	Tags *float32 `json:"tags,omitempty"`
	Versions *float32 `json:"versions,omitempty"`
	OverTime *OverTime `json:"over_time,omitempty"`
}

// NewHistoricalProvenance instantiates a new HistoricalProvenance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalProvenance() *HistoricalProvenance {
	this := HistoricalProvenance{}
	var overlap float32 = 0
	this.Overlap = &overlap
	var common float32 = 0
	this.Common = &common
	var tags float32 = 0
	this.Tags = &tags
	var versions float32 = 0
	this.Versions = &versions
	var overTime OverTime = {}
	this.OverTime = &overTime
	return &this
}

// NewHistoricalProvenanceWithDefaults instantiates a new HistoricalProvenance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalProvenanceWithDefaults() *HistoricalProvenance {
	this := HistoricalProvenance{}
	var overlap float32 = 0
	this.Overlap = &overlap
	var common float32 = 0
	this.Common = &common
	var tags float32 = 0
	this.Tags = &tags
	var versions float32 = 0
	this.Versions = &versions
	var overTime OverTime = {}
	this.OverTime = &overTime
	return &this
}

// GetOverlap returns the Overlap field value if set, zero value otherwise.
func (o *HistoricalProvenance) GetOverlap() float32 {
	if o == nil || IsNil(o.Overlap) {
		var ret float32
		return ret
	}
	return *o.Overlap
}

// GetOverlapOk returns a tuple with the Overlap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalProvenance) GetOverlapOk() (*float32, bool) {
	if o == nil || IsNil(o.Overlap) {
		return nil, false
	}
	return o.Overlap, true
}

// HasOverlap returns a boolean if a field has been set.
func (o *HistoricalProvenance) HasOverlap() bool {
	if o != nil && !IsNil(o.Overlap) {
		return true
	}

	return false
}

// SetOverlap gets a reference to the given float32 and assigns it to the Overlap field.
func (o *HistoricalProvenance) SetOverlap(v float32) {
	o.Overlap = &v
}

// GetCommon returns the Common field value if set, zero value otherwise.
func (o *HistoricalProvenance) GetCommon() float32 {
	if o == nil || IsNil(o.Common) {
		var ret float32
		return ret
	}
	return *o.Common
}

// GetCommonOk returns a tuple with the Common field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalProvenance) GetCommonOk() (*float32, bool) {
	if o == nil || IsNil(o.Common) {
		return nil, false
	}
	return o.Common, true
}

// HasCommon returns a boolean if a field has been set.
func (o *HistoricalProvenance) HasCommon() bool {
	if o != nil && !IsNil(o.Common) {
		return true
	}

	return false
}

// SetCommon gets a reference to the given float32 and assigns it to the Common field.
func (o *HistoricalProvenance) SetCommon(v float32) {
	o.Common = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *HistoricalProvenance) GetTags() float32 {
	if o == nil || IsNil(o.Tags) {
		var ret float32
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalProvenance) GetTagsOk() (*float32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *HistoricalProvenance) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given float32 and assigns it to the Tags field.
func (o *HistoricalProvenance) SetTags(v float32) {
	o.Tags = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *HistoricalProvenance) GetVersions() float32 {
	if o == nil || IsNil(o.Versions) {
		var ret float32
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalProvenance) GetVersionsOk() (*float32, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *HistoricalProvenance) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given float32 and assigns it to the Versions field.
func (o *HistoricalProvenance) SetVersions(v float32) {
	o.Versions = &v
}

// GetOverTime returns the OverTime field value if set, zero value otherwise.
func (o *HistoricalProvenance) GetOverTime() OverTime {
	if o == nil || IsNil(o.OverTime) {
		var ret OverTime
		return ret
	}
	return *o.OverTime
}

// GetOverTimeOk returns a tuple with the OverTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalProvenance) GetOverTimeOk() (*OverTime, bool) {
	if o == nil || IsNil(o.OverTime) {
		return nil, false
	}
	return o.OverTime, true
}

// HasOverTime returns a boolean if a field has been set.
func (o *HistoricalProvenance) HasOverTime() bool {
	if o != nil && !IsNil(o.OverTime) {
		return true
	}

	return false
}

// SetOverTime gets a reference to the given OverTime and assigns it to the OverTime field.
func (o *HistoricalProvenance) SetOverTime(v OverTime) {
	o.OverTime = &v
}

func (o HistoricalProvenance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricalProvenance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Overlap) {
		toSerialize["overlap"] = o.Overlap
	}
	if !IsNil(o.Common) {
		toSerialize["common"] = o.Common
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.OverTime) {
		toSerialize["over_time"] = o.OverTime
	}
	return toSerialize, nil
}

type NullableHistoricalProvenance struct {
	value *HistoricalProvenance
	isSet bool
}

func (v NullableHistoricalProvenance) Get() *HistoricalProvenance {
	return v.value
}

func (v *NullableHistoricalProvenance) Set(val *HistoricalProvenance) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricalProvenance) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricalProvenance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricalProvenance(val *HistoricalProvenance) *NullableHistoricalProvenance {
	return &NullableHistoricalProvenance{value: val, isSet: true}
}

func (v NullableHistoricalProvenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricalProvenance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


