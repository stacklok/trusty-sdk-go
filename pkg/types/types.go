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

// Package types is the collection of main data types used by the Trusty libraries
//
// Deprecated: moved to pkg/v1/types
package types

import (
	v1 "github.com/stacklok/trusty-sdk-go/pkg/v1/types"
)

// Ecosystem is an identifier of a packaging system supported by Trusty
//
// Deprecated: moved to pkg/v1/types
type Ecosystem = v1.Ecosystem

// Dependency represents a generic dependency structure
//
// Deprecated: moved to pkg/v1/types
type Dependency = v1.Dependency

const (
	// ECOSYSTEM_NPM identifies the NPM ecosystem
	//
	// Deprecated: moved to pkg/v1/types
	ECOSYSTEM_NPM Ecosystem = v1.ECOSYSTEM_NPM

	// ECOSYSTEM_GO identifies the Go language
	//
	// Deprecated: moved to pkg/v1/types
	ECOSYSTEM_GO Ecosystem = v1.ECOSYSTEM_GO

	// ECOSYSTEM_PYPI identifies the Python Package Index
	//
	// Deprecated: moved to pkg/v1/types
	ECOSYSTEM_PYPI Ecosystem = v1.ECOSYSTEM_PYPI

	// IngestStatusFailed ingestion failed permanently
	//
	// Deprecated: moved to pkg/v1/types
	IngestStatusFailed = v1.IngestStatusFailed

	// IngestStatusComplete means ingestion is done, data available
	//
	// Deprecated: moved to pkg/v1/types
	IngestStatusComplete = v1.IngestStatusComplete

	// IngestStatusPending means that the ingestion process is waiting to start
	//
	// Deprecated: moved to pkg/v1/types
	IngestStatusPending = v1.IngestStatusPending

	// IngestStatusScoring means the scoring process is underway
	//
	// Deprecated: moved to pkg/v1/types
	IngestStatusScoring = v1.IngestStatusScoring
)

// Ecosystems enumerates the supported ecosystems
//
// Deprecated: moved to pkg/v1/types
var Ecosystems = v1.Ecosystems

// ConvertDepsToMap converts a slice of Dependency structs to a map for easier comparison
//
// Deprecated: moved to pkg/v1/types
var ConvertDepsToMap = v1.ConvertDepsToMap

// DiffDependencies compares two sets of dependencies (represented as maps) and finds what's added in newDeps.
//
// Deprecated: moved to pkg/v1/types
var DiffDependencies = v1.DiffDependencies
