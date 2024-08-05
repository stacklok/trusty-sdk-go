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

	packageurl "github.com/package-url/packageurl-go"
	khttp "sigs.k8s.io/release-utils/http"

	"github.com/stacklok/trusty-sdk-go/pkg/types"
)

const (
	defaultEndpoint = "https://gh.trustypkg.dev"
	endpointEnvVar  = "TRUSTY_ENDPOINT"
	reportPath      = "v1/report"
)

// Options configures the Trusty API client
type Options struct {
	HttpClient netClient
	// Workers is the number of parallel request the client makes to the API
	Workers int

	// BaseURL of the Trusty API
	BaseURL string
}

// DefaultOptions is the default Trusty client options set
var DefaultOptions = Options{
	Workers: 2,
	BaseURL: defaultEndpoint,
}

type netClient interface {
	GetRequestGroup([]string) ([]*http.Response, []error)
	GetRequest(string) (*http.Response, error)
}

// New returns a new Trusty REST client
func New() *Trusty {
	opts := DefaultOptions
	opts.HttpClient = khttp.NewAgent().WithMaxParallel(opts.Workers).WithFailOnHTTPError(true)
	if ep := os.Getenv(endpointEnvVar); ep != "" {
		opts.BaseURL = ep
	}
	return NewWithOptions(opts)
}

// NewWithOptions returns a new client with the specified options set
func NewWithOptions(opts Options) *Trusty {
	if opts.BaseURL == "" {
		opts.BaseURL = DefaultOptions.BaseURL
	}

	if opts.Workers == 0 {
		opts.Workers = DefaultOptions.Workers
	}

	if opts.HttpClient == nil {
		opts.HttpClient = khttp.NewAgent().WithMaxParallel(opts.Workers).WithFailOnHTTPError(true)
	}

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

// GroupReport queries the Trusty API in parallel for a group of dependencies.
func (t *Trusty) GroupReport(_ context.Context, deps []*types.Dependency) ([]*types.Reply, error) {
	urls := []string{}
	for _, dep := range deps {
		u, err := t.PackageEndpoint(dep)
		if err != nil {
			return nil, fmt.Errorf("unable to get endpoint for: %q: %w", dep.Name, err)
		}
		urls = append(urls, u)
	}

	responses, errs := t.Options.HttpClient.GetRequestGroup(urls)
	if err := errors.Join(errs...); err != nil {
		return nil, fmt.Errorf("fetching data from Trusty: %w", err)
	}

	// Parse the replies
	resps := make([]*types.Reply, len(responses))
	for i := range responses {
		defer responses[i].Body.Close()
		dec := json.NewDecoder(responses[i].Body)
		resps[i] = &types.Reply{}
		if err := dec.Decode(resps[i]); err != nil {
			return nil, fmt.Errorf("could not unmarshal response #%d: %w", i, err)
		}
	}
	return resps, nil
}

// PurlEndpoint returns the API endpoint url to query for data about a purl
func (t *Trusty) PurlEndpoint(purl string) (string, error) {
	dep, err := t.PurlToDependency(purl)
	if err != nil {
		return "", fmt.Errorf("getting dependency from %q", purl)
	}
	ep, err := t.PackageEndpoint(dep)
	if err != nil {
		return "", fmt.Errorf("getting package endpoint: %w", err)
	}
	return ep, nil
}

// PackageEndpoint takes a dependency and returns the Trusty endpoint to
// query data about it.
func (t *Trusty) PackageEndpoint(dep *types.Dependency) (string, error) {
	// Check dependency data:
	errs := []error{}
	if dep.Name == "" {
		errs = append(errs, fmt.Errorf("dependency has no name defined"))
	}
	if dep.Ecosystem.AsString() == "" {
		errs = append(errs, fmt.Errorf("dependency has no ecosystem set"))
	}

	if err := errors.Join(errs...); err != nil {
		return "", err
	}

	u, err := url.Parse(t.Options.BaseURL + "/" + reportPath)
	if err != nil {
		return "", fmt.Errorf("failed to parse endpoint: %w", err)
	}

	params := map[string]string{
		"package_name": dep.Name,
		"package_type": strings.ToLower(dep.Ecosystem.AsString()),
	}

	// Add query parameters for package_name and package_type
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// PurlToEcosystem returns a trusty ecosystem constant from a Package URL's type
func (_ *Trusty) PurlToEcosystem(purl string) types.Ecosystem {
	switch {
	case strings.HasPrefix(purl, "pkg:golang"):
		return types.ECOSYSTEM_GO
	case strings.HasPrefix(purl, "pkg:npm"):
		return types.ECOSYSTEM_NPM
	case strings.HasPrefix(purl, "pkg:pypi"):
		return types.ECOSYSTEM_PYPI
	default:
		return types.Ecosystem(0)
	}
}

// PurlToDependency takes a string with a package url
func (t *Trusty) PurlToDependency(purlString string) (*types.Dependency, error) {
	e := t.PurlToEcosystem(purlString)
	if e == 0 {
		// Ecosystem nil or not supported
		return nil, fmt.Errorf("ecosystem not supported")
	}

	purl, err := packageurl.FromString(purlString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse package url: %w", err)
	}
	name := purl.Name
	if purl.Namespace != "" {
		name = purl.Namespace + "/" + purl.Name
	}
	return &types.Dependency{
		Ecosystem: e,
		Name:      name,
		Version:   purl.Version,
	}, nil
}

// Report returns a dependency report with all the data that Trust has
// available for a package
func (t *Trusty) Report(_ context.Context, dep *types.Dependency) (*types.Reply, error) {
	u, err := t.PackageEndpoint(dep)
	if err != nil {
		return nil, fmt.Errorf("computing package endpoint: %w", err)
	}

	resp, err := t.Options.HttpClient.GetRequest(u)
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
