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

import (
	v1 "github.com/stacklok/trusty-sdk-go/pkg/v1/types"
)

// Reply is the response from the package report API
//
// Deprecated: moved to pkg/v1/types
type Reply = v1.Reply

// Activity captures a package's activity score
//
// Deprecated: moved to pkg/v1/types
type Activity = v1.Activity

// ActivityDescription captures the fields of the activuty score
//
// Deprecated: moved to pkg/v1/types
type ActivityDescription = v1.ActivityDescription

// Typosquatting score for the package's name
//
// Deprecated: moved to pkg/v1/types
type Typosquatting = v1.Typosquatting

// TyposquattingDescription captures the dat details of the typosquatting score
//
// Deprecated: moved to pkg/v1/types
type TyposquattingDescription = v1.TyposquattingDescription

// Alternative is an alternative package returned from the package intelligence API
//
// Deprecated: moved to pkg/v1/types
type Alternative = v1.Alternative

// AlternativesList is the alternatives block in the trusty API response
//
// Deprecated: moved to pkg/v1/types
type AlternativesList = v1.AlternativesList

// ScoreSummary is the summary score returned from the package intelligence API
//
// Deprecated: moved to pkg/v1/types
type ScoreSummary = v1.ScoreSummary

// PackageData contains the data about the queried package
//
// Deprecated: moved to pkg/v1/types
type PackageData = v1.PackageData

// MaliciousData contains the security details when a dependency is malicious
//
// Deprecated: moved to pkg/v1/types
type MaliciousData = v1.MaliciousData

// Provenance has the package's provenance score and provenance type components
//
// Deprecated: moved to pkg/v1/types
type Provenance = v1.Provenance

// ProvenanceDescription contians the provenance types
//
// Deprecated: moved to pkg/v1/types
type ProvenanceDescription = v1.ProvenanceDescription

// HistoricalProvenance has the historical provenance components from a package
//
// Deprecated: moved to pkg/v1/types
type HistoricalProvenance = v1.HistoricalProvenance

// SigstoreProvenance has the sigstore certificate data when a package was signed
// using a github actions workflow
//
// Deprecated: moved to pkg/v1/types
type SigstoreProvenance = v1.SigstoreProvenance
