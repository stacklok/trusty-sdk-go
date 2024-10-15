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

// Contributor is a reusable structure that captures the information about a
// user related to a package.
type Contributor struct {
	ID         *uuid.UUID `json:"ID"`
	Author     *string
	Email      *string `json:"author_email"`
	Login      *string `json:"login"`
	AvatarURL  *string `json:"avatar_url"`
	GravatarID *string `json:"gravatar_id"`
	URL        *string `json:"url"`
	HTMLURL    *string `json:"html_url"`
	Company    *string
	Blog       *string
	Location   *string
	// Email *string  // email
	Hireable        *bool
	TwitterUsername *string `json:"ywitter_username"`
	PublicRepos     *int
	PublicGists     *int
	Followers       *int
	Following       *int
	// Scores
}
