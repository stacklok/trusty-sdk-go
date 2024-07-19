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

package types

import "time"

// Reply is the response from the package report API
type Reply struct {
	PackageName    string           `json:"package_name"`
	PackageType    string           `json:"package_type"`
	PackageVersion string           `json:"package_version"`
	Status         string           `json:"status"`
	Summary        ScoreSummary     `json:"summary"`
	Provenance     *Provenance      `json:"provenance"`
	Activity       *Activity        `json:"activity"`
	Typosquatting  *Typosquatting   `json:"typosquatting"`
	Alternatives   AlternativesList `json:"alternatives"`
	PackageData    PackageData      `json:"package_data"`
}

// Activity captures a package's activity score
type Activity struct {
	Score       float64 `json:"score"`
	Description string  `json:"description"`
}

// Typosquatting score for the package's name
type Typosquatting struct {
	Score       float64 `json:"score"`
	Description string  `json:"description"`
}

// Alternative is an alternative package returned from the package intelligence API
type Alternative struct {
	PackageName    string  `json:"package_name"`
	Score          float64 `json:"score"`
	PackageNameURL string
}

// AlternativesList is the alternatives block in the trusty API response
type AlternativesList struct {
	Status   string        `json:"status"`
	Packages []Alternative `json:"packages"`
}

// ScoreSummary is the summary score returned from the package intelligence API
type ScoreSummary struct {
	Score       *float64       `json:"score"`
	Description map[string]any `json:"description"`
}

// PackageData contains the data about the queried package
type PackageData struct {
	Archived   bool           `json:"archived"`
	Deprecated bool           `json:"is_deprecated"`
	Malicious  *MaliciousData `json:"malicious"`
}

// MaliciousData contains the security details when a dependency is malicious
type MaliciousData struct {
	Summary   string     `json:"summary"`
	Details   string     `json:"details"`
	Published *time.Time `json:"published"`
	Modified  *time.Time `json:"modified"`
	Source    string     `json:"source"`
}

// Provenance has the package's provenance score and provenance type components
type Provenance struct {
	Score       float64               `json:"score"`
	Description ProvenanceDescription `json:"description"`
	// UpdatedAt   time.Time             `json:"updated_at"`
}

// ProvenanceDescription contians the provenance types
type ProvenanceDescription struct {
	Historical HistoricalProvenance `json:"hp"`
	Sigstore   SigstoreProvenance   `json:"provenance"`
}

// HistoricalProvenance has the historical provenance components from a package
type HistoricalProvenance struct {
	Tags     float64 `json:"tags"`
	Common   float64 `json:"common"`
	Overlap  float64 `json:"overlap"`
	Versions float64 `json:"versions"`
}

// SigstoreProvenance has the sigstore certificate data when a package was signed
// using a github actions workflow
type SigstoreProvenance struct {
	Issuer           string `json:"issuer"`
	Workflow         string `json:"workflow"`
	SourceRepository string `json:"source_repo"`
	TokenIssuer      string `json:"token_issuer"`
	Transparency     string `json:"transparency"`
}
