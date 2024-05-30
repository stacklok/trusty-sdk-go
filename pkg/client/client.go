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
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/stacklok/trusty-sdk-go/pkg/types"
)

const (
	defaultEndpoint = "https://api.trustypkg.dev"
	endpointEnvVar  = "TRUSTY_ENDPOINT"
	reportPath      = "v1/report"
)

// Options configures the Trusty API client
type Options struct {
	HttpClient netClient
	BaseURL    string
}

// DefaultOptions is the default Trusty client options set
var DefaultOptions = Options{
	HttpClient: &http.Client{},
	BaseURL:    defaultEndpoint,
}

type netClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// New returns a new Trusty REST client
func New() *Trusty {
	opts := DefaultOptions
	if ep := os.Getenv(endpointEnvVar); ep != "" {
		opts.BaseURL = ep
	}
	return NewWithOptions(opts)
}

// NewWithOptions returns a new client with the dspecified options set
func NewWithOptions(opts Options) *Trusty {
	return &Trusty{
		Options: opts,
	}
}

func urlFromEndpointAndPaths(
	baseUrl, endpoint string, params map[string]string,
) (*url.URL, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse endpoint: %w", err)
	}
	u = u.JoinPath(endpoint)

	// Add query parameters for package_name and package_type
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	return u, nil
}

// Trusty is the main trusty client
type Trusty struct {
	Options Options
}

// newRequest buids a new http GET request using the preconfigured trusty API uri
func (t *Trusty) newRequest(ctx context.Context, path string, params map[string]string) (*http.Request, error) {
	u, err := urlFromEndpointAndPaths(t.Options.BaseURL, path, params)
	if err != nil {
		return nil, fmt.Errorf("could not parse endpoint: %w", err)
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}
	req = req.WithContext(ctx)
	return req, nil
}

// Report returns a dependency report with all the data that Trust has
// available for a package
func (t *Trusty) Report(ctx context.Context, dep *types.Dependency) (*types.Reply, error) {
	// Check dependency:
	errs := []error{}
	if dep.Name == "" {
		errs = append(errs, fmt.Errorf("dependency has no name defined"))
	}
	if dep.Ecosystem.AsString() == "" {
		errs = append(errs, fmt.Errorf("dependency has no ecosystem set"))
	}

	preErr := errors.Join(errs...)
	if preErr != nil {
		return nil, preErr
	}

	params := map[string]string{
		"package_name": dep.Name,
		"package_type": strings.ToLower(dep.Ecosystem.AsString()),
	}
	req, err := t.newRequest(ctx, reportPath, params)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	resp, err := t.Options.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	var r types.Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %w", err)
	}

	return &r, nil
}
