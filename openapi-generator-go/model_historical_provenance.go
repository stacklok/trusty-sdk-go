// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Trusty API
//	Version: v2
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// HistoricalProvenance is an object. [Historical provenance](provenance.md#historical-provenance-hp)
// This contains the number of `tags` in the repo, the number of `versions`
// of the package, a count of the `common` tags and the ratio of tags to common
// as `overlap`.
type HistoricalProvenance struct {
	// Common:
	Common float32 `json:"common,omitempty" mapstructure:"common,omitempty"`
	// OverTime:
	OverTime interface{} `json:"over_time,omitempty" mapstructure:"over_time,omitempty"`
	// Overlap:
	Overlap float32 `json:"overlap,omitempty" mapstructure:"overlap,omitempty"`
	// Tags:
	Tags float32 `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Versions:
	Versions float32 `json:"versions,omitempty" mapstructure:"versions,omitempty"`
}

// Validate implements basic validation for this model
func (m HistoricalProvenance) Validate() error {
	return validation.Errors{}.Filter()
}

// GetCommon returns the Common property
func (m HistoricalProvenance) GetCommon() float32 {
	return m.Common
}

// SetCommon sets the Common property
func (m *HistoricalProvenance) SetCommon(val float32) {
	m.Common = val
}

// GetOverTime returns the OverTime property
func (m HistoricalProvenance) GetOverTime() interface{} {
	return m.OverTime
}

// SetOverTime sets the OverTime property
func (m *HistoricalProvenance) SetOverTime(val interface{}) {
	m.OverTime = val
}

// GetOverlap returns the Overlap property
func (m HistoricalProvenance) GetOverlap() float32 {
	return m.Overlap
}

// SetOverlap sets the Overlap property
func (m *HistoricalProvenance) SetOverlap(val float32) {
	m.Overlap = val
}

// GetTags returns the Tags property
func (m HistoricalProvenance) GetTags() float32 {
	return m.Tags
}

// SetTags sets the Tags property
func (m *HistoricalProvenance) SetTags(val float32) {
	m.Tags = val
}

// GetVersions returns the Versions property
func (m HistoricalProvenance) GetVersions() float32 {
	return m.Versions
}

// SetVersions sets the Versions property
func (m *HistoricalProvenance) SetVersions(val float32) {
	m.Versions = val
}
