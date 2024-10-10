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

package responses

import "github.com/stacklok/trusty-sdk-go/pkg/types/v2/objects"

// Pkg captures the reply from the v2/pkg call. This API call returns the
// full data about a pacakge known to Trusty.
type Pkg struct {
	objects.Pkg
}
