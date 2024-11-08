//
// Copyright 2024 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package types is the collection of main data types used by the
// Trusty libraries
package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Dependency represents request arguments for various endpoints.
type Dependency struct {
	PackageName    string
	PackageType    *string
	PackageVersion *string
}

// PackageSummaryAnnotation represents a package annotation.
type PackageSummaryAnnotation struct {
	Score       *float64           `json:"score"`
	Description SummaryDescription `json:"description"`
	Status      *Status            `json:"status"`
	// The following field is currently not straightroward to
	// parse because it lacks timezone information,
	// i.e. 2024-11-14T11:24:09.119788.
	//
	// It is not a huge gap at the moment, so I'd rather add it
	// back once we agree on the format.
	//
	// UpdatedAt   *time.Time              `json:"updated_at"`
}

// Status represents that processing status of a package. It might be
// `"in_progress"` if the package was never seen previously, and
// changes to `"complete"` once processed.
type Status string

var (
	// StatusInProgress represents a package being processed.
	StatusInProgress Status = "in_progress"
	// StatusComplete represents an already processed package.
	StatusComplete Status = "complete"
)

//nolint:revive
func (t *Status) UnmarshalJSON(data []byte) error {
	var tmp string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	switch tmp {
	case "in_progress":
		*t = StatusInProgress
	case "complete":
		*t = StatusComplete
	default:
		return fmt.Errorf("invalid status type: %s", tmp)
	}

	return nil
}

// SummaryDescription is the response body of a `GET /v2/summary`
type SummaryDescription struct {
	From           string          `json:"from"`
	Provenance     float64         `json:"provenance"`
	TrustSummary   float64         `json:"trust-summary"`
	TypoSquatting  float64         `json:"typosquatting"`
	ActivityUser   float64         `json:"activity_user"`
	ActivityRepo   float64         `json:"activity_repo"`
	Activity       float64         `json:"activity"`
	TrustActivity  float64         `json:"trust-activity"`
	Malicious      bool            `json:"malicious"`
	ProvenanceType *ProvenanceType `json:"provenance_type"`
}

// ProvenanceType is the type of provenance information that Trusty
// was able to gather.
type ProvenanceType string

var (
	// ProvenanceTypeVerifiedProvenance represents a fully
	// verified provenance information.
	ProvenanceTypeVerifiedProvenance ProvenanceType = "verified_provenance"
	// ProvenanceTypeHistoricalProvenance represents a verified
	// historical provenance information.
	ProvenanceTypeHistoricalProvenance ProvenanceType = "historical_provenance_match"
	// ProvenanceTypeUnknown represents no provenance information.
	ProvenanceTypeUnknown ProvenanceType = "unknown"
	// ProvenanceTypeMismatched represents conflicting provenance
	// information.
	ProvenanceTypeMismatched ProvenanceType = "mismatched"
)

//nolint:revive
func (t *ProvenanceType) UnmarshalJSON(data []byte) error {
	var tmp string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	switch tmp {
	case "verified_provenance":
		*t = ProvenanceTypeVerifiedProvenance
	case "historical_provenance_match":
		*t = ProvenanceTypeHistoricalProvenance
	case "unknown":
		*t = ProvenanceTypeUnknown
	case "mismatched":
		*t = ProvenanceTypeMismatched
	default:
		return fmt.Errorf("invalid provenance type: %s", tmp)
	}

	return nil
}

// PackageType represents a package's ecosystem.
type PackageType string

// Implementation note: we do not implement `UnmarshalJSON` for
// `PackageType` because we want to get new package types seamlessly
// as they're added to Trusty. The downside of this is that sdk users
// must match new types manually until we add the case to the list.

var (
	// PackageTypePypi is the ecosystem of Python packages.
	PackageTypePypi PackageType = "pypi"
	// PackageTypeNpm is the ecosystem of JavaScript packages.
	PackageTypeNpm PackageType = "npm"
	// PackageTypeCrates is the ecosystem of Rust packages.
	PackageTypeCrates PackageType = "crates"
	// PackageTypeMaven is the ecosystem of Java packages.
	PackageTypeMaven PackageType = "maven"
	// PackageTypeGo is the ecosystem of Go packages.
	PackageTypeGo PackageType = "go"
)

// PackageStatus represents a package's status in the package
// repository.
type PackageStatus string

var (
	// PackageStatusPending represents status pending
	PackageStatusPending PackageStatus = "pending"
	// PackageStatusInitial represents status initial
	PackageStatusInitial PackageStatus = "initial"
	// PackageStatusNeighbours represents status neoghbours
	PackageStatusNeighbours PackageStatus = "neighbours"
	// PackageStatusComplete represents status complete
	PackageStatusComplete PackageStatus = "complete"
	// PackageStatusFailed represents status failed
	PackageStatusFailed PackageStatus = "failed"
	// PackageStatusScoring represents status scoring
	PackageStatusScoring PackageStatus = "scoring"
	// PackageStatusPropagate represents status propagate
	PackageStatusPropagate PackageStatus = "propagate"
	// PackageStatusDeleted represents status deleted
	PackageStatusDeleted PackageStatus = "deleted"
)

//nolint:revive
func (t *PackageStatus) UnmarshalJSON(data []byte) error {
	var tmp string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	switch tmp {
	case "pending":
		*t = PackageStatusPending
	case "initial":
		*t = PackageStatusInitial
	case "neighbours":
		*t = PackageStatusNeighbours
	case "complete":
		*t = PackageStatusComplete
	case "failed":
		*t = PackageStatusFailed
	case "scoring":
		*t = PackageStatusScoring
	case "propagate":
		*t = PackageStatusPropagate
	case "deleted":
		*t = PackageStatusDeleted
	default:
		return fmt.Errorf("invalid package status type: %s", tmp)
	}

	return nil
}

// TrustyPackageData is the full package information as returned from
// the `GET /v2/pkg` API call. In contrast to `PackageBrief` this
// structure captures the complete data that trusty knows about the
// described package.
type TrustyPackageData struct {
	ID         *uuid.UUID     `json:"id"`
	Status     *PackageStatus `json:"status"`
	StatusCode *string        `json:"status_code"`
	Name       string         `json:"name"`
	Type       PackageType    `json:"type"`
	Version    *string        `json:"version"`
	// The following field is currently not straightroward to
	// parse because it lacks timezone information,
	// i.e. 2024-11-14T11:24:09.119788.
	//
	// It is not a huge gap at the moment, so I'd rather add it
	// back once we agree on the format.
	//
	// VersionDate      *time.Time     `json:"version_date"`
	Author           *string    `json:"author"`
	AuthorEmail      *string    `json:"author_email"`
	Description      *string    `json:"packag_description"`
	RepoDescription  *string    `json:"repo_description"`
	Origin           *string    `json:"origin"`
	StarGazersCount  *int       `json:"stargazers_count"`
	WatchersCount    *int       `json:"watchers_count"`
	HomePage         *string    `json:"home_page"`
	HasIssues        *bool      `json:"has_issues"`
	HasProjects      *bool      `json:"has_projects"`
	HasDownloads     *bool      `json:"has_downloads"`
	ForksCount       *int       `json:"forks_count"`
	Archived         *bool      `json:"archived"`
	IsDeprecated     *bool      `json:"is_deprecated"`
	Disabled         *bool      `json:"disabled"`
	OpenIssuesCount  *int       `json:"open_issues_count"`
	Visibility       *string    `json:"visibility"`
	DefaultBranch    *string    `json:"default_branch"`
	RepositoryID     *uuid.UUID `json:"repository_id"`
	RepositoryName   *string    `json:"repository_name"`
	ContributorCount *int       `json:"contributor_count"`
	PublicRepos      *int       `json:"public_repos"`
	PublicGists      *int       `json:"public_gists"`
	Followers        *int       `json:"followers"`
	Following        *int       `json:"following"`
	Owner            *User      `json:"owner"`
	Contributors     []*User    `json:"contributors"`
	// The following field is currently not straightroward to
	// parse because it lacks timezone information,
	// i.e. 2024-11-14T11:24:09.119788.
	//
	// It is not a huge gap at the moment, so I'd rather add it
	// back once we agree on the format.
	//
	// LastUpdate *time.Time `json:"last_update"`
	Scores                  interface{}              `json:"scores"`
	Malicious               *PackageMaliciousPayload `json:"malicious"`
	HasTriggeredReingestion *bool                    `json:"has_triggered_reingestion"`
}

// User represents an individual or bot that acts on repositories in
// some way.
type User struct {
	Id               uuid.UUID   `json:"id"`
	Author           *string     `json:"author"`
	Author_email     *string     `json:"author_email"`
	Login            *string     `json:"login"`
	Avatar_url       *string     `json:"avatar_url"`
	Gravatar_id      *string     `json:"gravatar_id"`
	Url              *string     `json:"url"`
	Html_url         *string     `json:"html_url"`
	Company          *string     `json:"company"`
	Blog             *string     `json:"blog"`
	Location         *string     `json:"location"`
	Email            *string     `json:"email"`
	Hireable         bool        `json:"hireable"`
	Twitter_username *string     `json:"twitter_username"`
	Public_repos     *int        `json:"public_repos"`
	Public_gists     *int        `json:"public_gists"`
	Followers        *int        `json:"followers"`
	Following        *int        `json:"following"`
	Scores           interface{} `json:"scores"`
}

// PackageMaliciousPayload represents the payload details for a
// malicious package.
type PackageMaliciousPayload struct {
	Summary   string     `json:"summary"`
	Details   *string    `json:"details"`
	Published *time.Time `json:"published"`
	Modified  *time.Time `json:"modified"`
	Source    *string    `json:"source"`
}

// PackageAlternatives is a list of alternative packages to the one
// requested.
type PackageAlternatives struct {
	Status   Status              `json:"status"` // in_progress or complete
	Packages []*PackageBasicInfo `json:"packages"`
}

// PackageBasicInfo contains basic information about a package.
type PackageBasicInfo struct {
	ID              *uuid.UUID   `json:"id"`
	PackageName     string       `json:"package_name"`
	PackageType     *PackageType `json:"package_type"`
	PackageVersion  *string      `json:"package_version"`
	RepoDescription *string      `json:"repo_description"`
	Score           *float64     `json:"score"`
	IsMalicious     bool         `json:"is_malicious"`
}
