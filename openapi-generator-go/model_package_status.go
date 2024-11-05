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

// PackageStatus is an enum.
type PackageStatus string

// Validate implements basic validation for this model
func (m PackageStatus) Validate() error {
	return InKnownPackageStatus.Validate(m)
}

var (
	PackageStatusComplete   PackageStatus = "complete"
	PackageStatusDeleted    PackageStatus = "deleted"
	PackageStatusFailed     PackageStatus = "failed"
	PackageStatusInitial    PackageStatus = "initial"
	PackageStatusNeighbours PackageStatus = "neighbours"
	PackageStatusPending    PackageStatus = "pending"
	PackageStatusPropagate  PackageStatus = "propagate"
	PackageStatusScoring    PackageStatus = "scoring"

	// KnownPackageStatus is the list of valid PackageStatus
	KnownPackageStatus = []PackageStatus{
		PackageStatusComplete,
		PackageStatusDeleted,
		PackageStatusFailed,
		PackageStatusInitial,
		PackageStatusNeighbours,
		PackageStatusPending,
		PackageStatusPropagate,
		PackageStatusScoring,
	}
	// KnownPackageStatusString is the list of valid PackageStatus as string
	KnownPackageStatusString = []string{
		string(PackageStatusComplete),
		string(PackageStatusDeleted),
		string(PackageStatusFailed),
		string(PackageStatusInitial),
		string(PackageStatusNeighbours),
		string(PackageStatusPending),
		string(PackageStatusPropagate),
		string(PackageStatusScoring),
	}

	// InKnownPackageStatus is an ozzo-validator for PackageStatus
	InKnownPackageStatus = validation.In(
		PackageStatusComplete,
		PackageStatusDeleted,
		PackageStatusFailed,
		PackageStatusInitial,
		PackageStatusNeighbours,
		PackageStatusPending,
		PackageStatusPropagate,
		PackageStatusScoring,
	)
)
