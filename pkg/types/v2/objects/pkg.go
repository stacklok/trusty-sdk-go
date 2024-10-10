// Copyright 2024 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package objects

import (
	"time"

	"github.com/google/uuid"
)

// Pkg is the full package information as returned from the  v2/pkg API call.
// In contrast to PackageBrief this structure captures the complete data
// that trusty knows about the described package.
type Pkg struct {
	ID               *uuid.UUID
	Status           *string
	StatusCode       *string
	Name             *string
	Type             *string
	Version          *string
	VersionDate      *time.Time
	Author           *string
	AuthorEmail      *string `json:"author_email"`
	Description      *string `json:"packag_description"`
	RepoDescription  *string `json:"repo_description"`
	Origin           *string // What is this?
	StarGazersCount  *int    `json:"stargazers_count"`
	WatchersCount    *int    `json:"watchers_count"`
	HomePage         *string `json:"home_page"`
	HasIssues        *bool
	HasProjects      *bool
	HasDownloads     *bool
	ForksCount       *int
	Archived         *bool
	Deprecated       *bool `json:"is_deprecated"`
	Disabled         *bool `json:"disabled"`
	OpenIssuesCount  *int  `json:"open_issues_count"`
	Visibility       *string
	DefaultBranch    *string    `json:"default_branch"`
	RepositoryID     *uuid.UUID `json:"repository_id"`
	RepositoryName   *string    `json:"repository_name"`
	ContributorCount *int       `json:"contributor_count"`
	PublicRepos      *int       `json:"public_repos"`
	PublicGists      *int       `json:"public_gists"`
	Followers        *int
	Following        *int
	Owner            *Contributor
	Contributors     []*Contributor
	LastUpdate       *time.Time `json:"last_update"`
	// Scores
	// "scores": {},
	Malicious               *MaliciousData
	HasTriggeredReingestion *bool
}
