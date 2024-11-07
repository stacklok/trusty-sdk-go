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

// Package client provides a rest client to talk to the Trusty API.
//
// Deprecated: moved to pkg/v1/client
package client

import (
	v1 "github.com/stacklok/trusty-sdk-go/pkg/v1/client"
)

// Options configures the Trusty API client
//
// Deprecated: moved to pkg/v1/client
type Options = v1.Options

// DefaultOptions is the default Trusty client options set
//
// Deprecated: moved to pkg/v1/client
var DefaultOptions = v1.DefaultOptions

// New returns a new Trusty REST client
//
// Deprecated: moved to pkg/v1/client
var New = v1.New

// NewWithOptions returns a new client with the specified options set
//
// Deprecated: moved to pkg/v1/client
var NewWithOptions = v1.NewWithOptions

// Trusty is the main trusty client
//
// Deprecated: moved to pkg/v1/client
type Trusty = v1.Trusty
