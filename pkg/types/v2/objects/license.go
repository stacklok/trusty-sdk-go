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

// LicenseIdentifier overrides string and captures an SPDX license identifier
type LicenseIdentifier string

// LicenseClaim is a struct that captures the licensing data from a package.
type LicenseClaim struct {
	ID          *uuid.UUID
	OwnerID     *uuid.UUID           `json:"owner_id"`
	Licenses    []*LicenseIdentifier `json:"licenses"`
	Claim       *LicenseClaim
	Content     *string // ?
	URL         *string // ?
	Source      *string
	Description *string
}