// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Trusty API
//	Version: v2
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

// PersonAnnotation is an object. This represents a package annotation.
type PersonAnnotation struct {
	// Description:
	Description interface{} `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Score:
	Score interface{} `json:"score,omitempty" mapstructure:"score,omitempty"`
	// UpdatedAt:
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"`
}

// Validate implements basic validation for this model
func (m PersonAnnotation) Validate() error {
	return validation.Errors{
		"updatedAt": validation.Validate(
			m.UpdatedAt, validation.Required, validation.Date(time.RFC3339),
		),
	}.Filter()
}

// GetDescription returns the Description property
func (m PersonAnnotation) GetDescription() interface{} {
	return m.Description
}

// SetDescription sets the Description property
func (m *PersonAnnotation) SetDescription(val interface{}) {
	m.Description = val
}

// GetScore returns the Score property
func (m PersonAnnotation) GetScore() interface{} {
	return m.Score
}

// SetScore sets the Score property
func (m *PersonAnnotation) SetScore(val interface{}) {
	m.Score = val
}

// GetUpdatedAt returns the UpdatedAt property
func (m PersonAnnotation) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// SetUpdatedAt sets the UpdatedAt property
func (m *PersonAnnotation) SetUpdatedAt(val time.Time) {
	m.UpdatedAt = val
}
