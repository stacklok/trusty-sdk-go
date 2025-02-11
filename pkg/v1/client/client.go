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
package client

import (
	"context"

	internalclient "github.com/stacklok/trusty-sdk-go/internal/client"
	"github.com/stacklok/trusty-sdk-go/pkg/v1/types"
)

// Options configures the Trusty API client
type Options = internalclient.Options

// DefaultOptions is the default Trusty client options set
var DefaultOptions = internalclient.DefaultOptions

// Trusty is a client on v1 Trusty APIs.
type Trusty interface {
	// Report returns a dependency report with all the data that
	// Trusty has available for a package.
	Report(context.Context, *types.Dependency) (*types.Reply, error)
	// GroupReport queries the Trusty API in parallel for a group
	// of dependencies.
	GroupReport(context.Context, []*types.Dependency) ([]*types.Reply, error)

	// PurlEndpoint returns the API endpoint url to query for data
	// about a purl.
	PurlEndpoint(string) (string, error)
	// PackageEndpoint takes a dependency and returns the Trusty
	// endpoint to query data about it.
	PackageEndpoint(*types.Dependency) (string, error)
	// PurlToEcosystem returns a trusty ecosystem constant from a
	// Package URL's type.
	PurlToEcosystem(string) types.Ecosystem
	// PurlToDependency takes a string with a package url.
	PurlToDependency(string) (*types.Dependency, error)
}

// New returns a new Trusty REST client
func New() Trusty {
	return internalclient.New()
}

// NewWithOptions returns a new client with the specified options set
func NewWithOptions(opts Options) Trusty {
	return internalclient.NewWithOptions(opts)
}
