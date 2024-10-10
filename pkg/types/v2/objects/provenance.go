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

import "time"

// Provenance has the package's provenance score and provenance type components
type Provenance struct {
	Score       float64               `json:"score"`
	Description ProvenanceDescription `json:"description"`
	UpdatedAt   *time.Time            `json:"updated_at"`
}

// ProvenanceDescription contians the provenance types
type ProvenanceDescription struct {
	Historical HistoricalProvenance `json:"hp"`
	Sigstore   SigstoreProvenance   `json:"sigstore"`
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
