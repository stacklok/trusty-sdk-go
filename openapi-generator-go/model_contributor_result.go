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

// ContributorResult is an object.
type ContributorResult struct {
	// ContributorData:
	ContributorData *PackageContributorData `json:"contributor_data,omitempty" mapstructure:"contributor_data,omitempty"`
	// ContributorRepos:
	ContributorRepos interface{} `json:"contributor_repos,omitempty" mapstructure:"contributor_repos,omitempty"`
}

// Validate implements basic validation for this model
func (m ContributorResult) Validate() error {
	return validation.Errors{
		"contributorData": validation.Validate(
			m.ContributorData,
		),
	}.Filter()
}

// GetContributorData returns the ContributorData property
func (m ContributorResult) GetContributorData() *PackageContributorData {
	return m.ContributorData
}

// SetContributorData sets the ContributorData property
func (m *ContributorResult) SetContributorData(val *PackageContributorData) {
	m.ContributorData = val
}

// GetContributorRepos returns the ContributorRepos property
func (m ContributorResult) GetContributorRepos() interface{} {
	return m.ContributorRepos
}

// SetContributorRepos sets the ContributorRepos property
func (m *ContributorResult) SetContributorRepos(val interface{}) {
	m.ContributorRepos = val
}
