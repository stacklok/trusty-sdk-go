/*
Trusty API

Trusty API

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SameOriginPackagesResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SameOriginPackagesResult{}

// SameOriginPackagesResult struct for SameOriginPackagesResult
type SameOriginPackagesResult struct {
	NextToken string `json:"next_token"`
	SameOriginPackages []PackageBasicInfo `json:"same_origin_packages,omitempty"`
}

type _SameOriginPackagesResult SameOriginPackagesResult

// NewSameOriginPackagesResult instantiates a new SameOriginPackagesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSameOriginPackagesResult(nextToken string) *SameOriginPackagesResult {
	this := SameOriginPackagesResult{}
	this.NextToken = nextToken
	return &this
}

// NewSameOriginPackagesResultWithDefaults instantiates a new SameOriginPackagesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSameOriginPackagesResultWithDefaults() *SameOriginPackagesResult {
	this := SameOriginPackagesResult{}
	return &this
}

// GetNextToken returns the NextToken field value
func (o *SameOriginPackagesResult) GetNextToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value
// and a boolean to check if the value has been set.
func (o *SameOriginPackagesResult) GetNextTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextToken, true
}

// SetNextToken sets field value
func (o *SameOriginPackagesResult) SetNextToken(v string) {
	o.NextToken = v
}

// GetSameOriginPackages returns the SameOriginPackages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SameOriginPackagesResult) GetSameOriginPackages() []PackageBasicInfo {
	if o == nil {
		var ret []PackageBasicInfo
		return ret
	}
	return o.SameOriginPackages
}

// GetSameOriginPackagesOk returns a tuple with the SameOriginPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SameOriginPackagesResult) GetSameOriginPackagesOk() ([]PackageBasicInfo, bool) {
	if o == nil || IsNil(o.SameOriginPackages) {
		return nil, false
	}
	return o.SameOriginPackages, true
}

// HasSameOriginPackages returns a boolean if a field has been set.
func (o *SameOriginPackagesResult) HasSameOriginPackages() bool {
	if o != nil && !IsNil(o.SameOriginPackages) {
		return true
	}

	return false
}

// SetSameOriginPackages gets a reference to the given []PackageBasicInfo and assigns it to the SameOriginPackages field.
func (o *SameOriginPackagesResult) SetSameOriginPackages(v []PackageBasicInfo) {
	o.SameOriginPackages = v
}

func (o SameOriginPackagesResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SameOriginPackagesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["next_token"] = o.NextToken
	if o.SameOriginPackages != nil {
		toSerialize["same_origin_packages"] = o.SameOriginPackages
	}
	return toSerialize, nil
}

func (o *SameOriginPackagesResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"next_token",
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

	varSameOriginPackagesResult := _SameOriginPackagesResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSameOriginPackagesResult)

	if err != nil {
		return err
	}

	*o = SameOriginPackagesResult(varSameOriginPackagesResult)

	return err
}

type NullableSameOriginPackagesResult struct {
	value *SameOriginPackagesResult
	isSet bool
}

func (v NullableSameOriginPackagesResult) Get() *SameOriginPackagesResult {
	return v.value
}

func (v *NullableSameOriginPackagesResult) Set(val *SameOriginPackagesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSameOriginPackagesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSameOriginPackagesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSameOriginPackagesResult(val *SameOriginPackagesResult) *NullableSameOriginPackagesResult {
	return &NullableSameOriginPackagesResult{value: val, isSet: true}
}

func (v NullableSameOriginPackagesResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSameOriginPackagesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


