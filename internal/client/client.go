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
	"time"

	packageurl "github.com/package-url/packageurl-go"
	khttp "sigs.k8s.io/release-utils/http"

	v1types "github.com/stacklok/trusty-sdk-go/pkg/v1/types"
	v2types "github.com/stacklok/trusty-sdk-go/pkg/v2/types"
)

const (
	defaultEndpoint = "https://api.trustypkg.dev"
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

	// WaitForIngestion causes the http client to wait and retry if Trusty
	// responds with a successful request but with a "pending" or "scoring" status
	WaitForIngestion bool

	// ErrOnFailedIngestion makes the client return an error on a Report call
	// when the ingestion failed internally withing trusty. If false, the
	// report data willbe returned but the application needs to check the
	// ingestion status and handle it.
	ErrOnFailedIngestion bool

	// IngestionRetryWait is the number of seconds that the client will wait for
	// package ingestion before retrying.
	IngestionRetryWait int

	// IngestionMaxRetries is the maximum number of requests the client will
	// send while waiting for ingestion to finish
	IngestionMaxRetries int
}

// DefaultOptions is the default Trusty client options set
var DefaultOptions = Options{
	Workers:            2,
	BaseURL:            defaultEndpoint,
	WaitForIngestion:   true,
	IngestionRetryWait: 5,
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
func (t *Trusty) GroupReport(_ context.Context, deps []*v1types.Dependency) ([]*v1types.Reply, error) {
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
	resps := make([]*v1types.Reply, len(responses))
	for i := range responses {
		defer responses[i].Body.Close()
		dec := json.NewDecoder(responses[i].Body)
		resps[i] = &v1types.Reply{}
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
func (t *Trusty) PackageEndpoint(dep *v1types.Dependency) (string, error) {
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
func (_ *Trusty) PurlToEcosystem(purl string) v1types.Ecosystem {
	switch {
	case strings.HasPrefix(purl, "pkg:golang"):
		return v1types.ECOSYSTEM_GO
	case strings.HasPrefix(purl, "pkg:npm"):
		return v1types.ECOSYSTEM_NPM
	case strings.HasPrefix(purl, "pkg:pypi"):
		return v1types.ECOSYSTEM_PYPI
	default:
		return v1types.Ecosystem(0)
	}
}

// PurlToDependency takes a string with a package url
func (t *Trusty) PurlToDependency(purlString string) (*v1types.Dependency, error) {
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
	return &v1types.Dependency{
		Ecosystem: e,
		Name:      name,
		Version:   purl.Version,
	}, nil
}

// Report returns a dependency report with all the data that Trusty has
// available for a package.
func (t *Trusty) Report(_ context.Context, dep *v1types.Dependency) (*v1types.Reply, error) {
	u, err := t.PackageEndpoint(dep)
	if err != nil {
		return nil, fmt.Errorf("computing package endpoint: %w", err)
	}

	var r v1types.Reply
	tries := 0
	for {
		resp, err := t.Options.HttpClient.GetRequest(u)
		if err != nil {
			return nil, fmt.Errorf("could not send request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("received non-200 response: %d", resp.StatusCode)
		}

		dec := json.NewDecoder(resp.Body)
		if err := dec.Decode(&r); err != nil {
			return nil, fmt.Errorf("could not unmarshal response: %w", err)
		}
		fmt.Printf("Attempt #%d to fetch package, status: %s", tries, r.PackageData.Status)

		shouldRetry, err := evalRetry(r.PackageData.Status, t.Options)
		if err != nil {
			return nil, err
		}

		if !shouldRetry {
			break
		}

		tries++
		if tries > t.Options.IngestionMaxRetries {
			return nil, fmt.Errorf("time out reached waiting for package ingestion")
		}
		time.Sleep(time.Duration(t.Options.IngestionRetryWait) * time.Second)
	}

	return &r, err
}

func evalRetry(status string, opts Options) (shouldRetry bool, err error) {
	// First, error if the ingestion status is invalid
	if status != v1types.IngestStatusFailed && status != v1types.IngestStatusComplete &&
		status != v1types.IngestStatusPending && status != v1types.IngestStatusScoring {

		return false, fmt.Errorf("unexpected ingestion status when querying package")
	}

	if status == v1types.IngestStatusFailed && opts.ErrOnFailedIngestion {
		return false, fmt.Errorf("upstream error ingesting package data")
	}

	// Package ingestion is ready
	if status == v1types.IngestStatusComplete {
		return false, nil
	}

	// Client configured to return raw response (even when package is not ready)
	if !opts.WaitForIngestion || status == v1types.IngestStatusFailed {
		return false, nil
	}

	return true, nil
}

const (
	// v2 paths
	v2SummaryPath  = "v2/summary"
	v2PkgPath      = "v2/pkg"
	v2Alternatives = "v2/alternatives"
	v2Provenance   = "v2/provenance"
)

// Summary fetches a summary of Security Signal information
// for the package.
func (t *Trusty) Summary(
	_ context.Context,
	dep *v2types.Dependency,
) (*v2types.PackageSummaryAnnotation, error) {
	if dep.PackageName == "" {
		return nil, fmt.Errorf("dependency has no name defined")
	}
	if dep.PackageType == "" {
		return nil, fmt.Errorf("dependency has no ecosystem defined")
	}

	u, err := urlFor(t.Options.BaseURL, v2SummaryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse endpoint: %w", err)
	}

	// Add query parameters for package_name, package_type, and
	// package_version.
	q := u.Query()
	q.Set("package_name", dep.PackageName)
	q.Set("package_type", strings.ToLower(dep.PackageType))
	if dep.PackageVersion != nil && *dep.PackageVersion != "" {
		q.Set("package_version", *dep.PackageVersion)
	}
	u.RawQuery = q.Encode()

	return doRequest[v2types.PackageSummaryAnnotation](t.Options.HttpClient, u.String())
}

// PackageMetadata fetched the metadata for a package.
//
// This includes the package's name, version, description, and
// other metadata about contributors.
func (t *Trusty) PackageMetadata(
	_ context.Context,
	dep *v2types.Dependency,
) (*v2types.TrustyPackageData, error) {
	if dep.PackageName == "" {
		return nil, fmt.Errorf("dependency has no name defined")
	}
	if dep.PackageType == "" {
		return nil, fmt.Errorf("dependency has no ecosystem defined")
	}

	u, err := urlFor(t.Options.BaseURL, v2PkgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse endpoint: %w", err)
	}

	// Add query parameters for package_name, package_type, and
	// package_version.
	q := u.Query()
	q.Set("package_name", dep.PackageName)
	q.Set("package_type", strings.ToLower(dep.PackageType))
	if dep.PackageVersion != nil && *dep.PackageVersion != "" {
		q.Set("package_version", *dep.PackageVersion)
	}
	u.RawQuery = q.Encode()

	return doRequest[v2types.TrustyPackageData](t.Options.HttpClient, u.String())
}

// Alternatives fetches packages that can be used in place of the
// given one.
func (t *Trusty) Alternatives(
	_ context.Context,
	dep *v2types.Dependency,
) (*v2types.PackageAlternatives, error) {
	if dep.PackageName == "" {
		return nil, fmt.Errorf("dependency has no name defined")
	}
	if dep.PackageType == "" {
		return nil, fmt.Errorf("dependency has no ecosystem defined")
	}

	u, err := urlFor(t.Options.BaseURL, v2Alternatives)
	if err != nil {
		return nil, fmt.Errorf("failed to parse endpoint: %w", err)
	}

	// Add query parameters for package_name, package_type, and
	// package_version.
	q := u.Query()
	q.Set("package_name", dep.PackageName)
	q.Set("package_type", strings.ToLower(dep.PackageType))
	if dep.PackageVersion != nil && *dep.PackageVersion != "" {
		q.Set("package_version", *dep.PackageVersion)
	}
	u.RawQuery = q.Encode()

	return doRequest[v2types.PackageAlternatives](t.Options.HttpClient, u.String())
}

// Provenance fetches detailed provenance information of a given
// package.
func (t *Trusty) Provenance(
	_ context.Context,
	dep *v2types.Dependency,
) (*v2types.Provenance, error) {
	if dep.PackageName == "" {
		return nil, fmt.Errorf("dependency has no name defined")
	}
	if dep.PackageType == "" {
		return nil, fmt.Errorf("dependency has no ecosystem defined")
	}

	u, err := urlFor(t.Options.BaseURL, v2Provenance)
	if err != nil {
		return nil, fmt.Errorf("failed to parse endpoint: %w", err)
	}

	// Add query parameters for package_name, package_type, and
	// package_version.
	q := u.Query()
	q.Set("package_name", dep.PackageName)
	q.Set("package_type", strings.ToLower(dep.PackageType))
	if dep.PackageVersion != nil && *dep.PackageVersion != "" {
		q.Set("package_version", *dep.PackageVersion)
	}
	u.RawQuery = q.Encode()

	return doRequest[v2types.Provenance](t.Options.HttpClient, u.String())
}

// doRequest only wraps (1) an HTTP GET issued to the given URL using
// the given client, and (2) result deserialization.
func doRequest[T any](client netClient, fullurl string) (*T, error) {
	resp, err := client.GetRequest(fullurl)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	var res T
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&res); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %w", err)
	}

	return &res, nil
}

func urlFor(baseURL, path string) (*url.URL, error) {
	ustr, err := url.JoinPath(baseURL, path)
	if err != nil {
		return nil, err
	}
	return url.Parse(ustr)
}
