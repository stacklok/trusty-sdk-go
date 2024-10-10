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

import "github.com/google/uuid"

// PackageBrief is the smaller representation of a package. It is mainly used
// in responses that ennumerate other packages.
type PackageBrief struct {
	ID              *uuid.UUID
	Name            *string `json:"package_name"`
	Type            *string `json:"package_type"`
	Version         *string `json:"package_version"`
	RepoDescription *string `json:"repo_description"`
	Score           *float64
	IsMalicious     *bool `json:"is_malicious"`
	Provenance      *Provenance
}
