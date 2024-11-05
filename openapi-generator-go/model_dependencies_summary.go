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

// DependenciesSummary is an object.
type DependenciesSummary struct {
	// DeclaredLicenses:
	DeclaredLicenses map[string]int32 `json:"declared_licenses,omitempty" mapstructure:"declared_licenses,omitempty"`
	// Depths:
	Depths map[string]int32 `json:"depths,omitempty" mapstructure:"depths,omitempty"`
	// MaxScore:
	MaxScore float32 `json:"max_score,omitempty" mapstructure:"max_score,omitempty"`
	// MeanScore:
	MeanScore float32 `json:"mean_score,omitempty" mapstructure:"mean_score,omitempty"`
	// MinScore:
	MinScore float32 `json:"min_score,omitempty" mapstructure:"min_score,omitempty"`
	// Total:
	Total int32 `json:"total,omitempty" mapstructure:"total,omitempty"`
	// VulnSeverity:
	VulnSeverity map[string]int32 `json:"vuln_severity,omitempty" mapstructure:"vuln_severity,omitempty"`
}

// Validate implements basic validation for this model
func (m DependenciesSummary) Validate() error {
	return validation.Errors{
		"declaredLicenses": validation.Validate(
			m.DeclaredLicenses,
		),
		"depths": validation.Validate(
			m.Depths,
		),
		"vulnSeverity": validation.Validate(
			m.VulnSeverity,
		),
	}.Filter()
}

// GetDeclaredLicenses returns the DeclaredLicenses property
func (m DependenciesSummary) GetDeclaredLicenses() map[string]int32 {
	return m.DeclaredLicenses
}

// SetDeclaredLicenses sets the DeclaredLicenses property
func (m *DependenciesSummary) SetDeclaredLicenses(val map[string]int32) {
	m.DeclaredLicenses = val
}

// GetDepths returns the Depths property
func (m DependenciesSummary) GetDepths() map[string]int32 {
	return m.Depths
}

// SetDepths sets the Depths property
func (m *DependenciesSummary) SetDepths(val map[string]int32) {
	m.Depths = val
}

// GetMaxScore returns the MaxScore property
func (m DependenciesSummary) GetMaxScore() float32 {
	return m.MaxScore
}

// SetMaxScore sets the MaxScore property
func (m *DependenciesSummary) SetMaxScore(val float32) {
	m.MaxScore = val
}

// GetMeanScore returns the MeanScore property
func (m DependenciesSummary) GetMeanScore() float32 {
	return m.MeanScore
}

// SetMeanScore sets the MeanScore property
func (m *DependenciesSummary) SetMeanScore(val float32) {
	m.MeanScore = val
}

// GetMinScore returns the MinScore property
func (m DependenciesSummary) GetMinScore() float32 {
	return m.MinScore
}

// SetMinScore sets the MinScore property
func (m *DependenciesSummary) SetMinScore(val float32) {
	m.MinScore = val
}

// GetTotal returns the Total property
func (m DependenciesSummary) GetTotal() int32 {
	return m.Total
}

// SetTotal sets the Total property
func (m *DependenciesSummary) SetTotal(val int32) {
	m.Total = val
}

// GetVulnSeverity returns the VulnSeverity property
func (m DependenciesSummary) GetVulnSeverity() map[string]int32 {
	return m.VulnSeverity
}

// SetVulnSeverity sets the VulnSeverity property
func (m *DependenciesSummary) SetVulnSeverity(val map[string]int32) {
	m.VulnSeverity = val
}
