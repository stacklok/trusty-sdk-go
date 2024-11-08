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
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	v1types "github.com/stacklok/trusty-sdk-go/pkg/v1/types"
)

func newFakeClient() *fakeClient {
	return &fakeClient{
		resps: []*http.Response{},
		errs:  []error{},
	}
}

// fakeClient mocks the http client used by the trusty client
type fakeClient struct {
	resps []*http.Response
	errs  []error
}

func (fc *fakeClient) GetRequest(_ string) (response *http.Response, err error) {
	if len(fc.resps) != 0 {
		response = fc.resps[0]
		if _, err := response.Body.(fakeCloser).Seek(0, 0); err != nil {
			return nil, fmt.Errorf("seeking fake response: %w", err)
		}
	}

	if len(fc.errs) != 0 {
		err = fc.errs[0]
	}
	return response, err
}

func (fc *fakeClient) GetRequestGroup(_ []string) (response []*http.Response, errs []error) {
	return fc.resps, fc.errs
}

type fakeCloser struct {
	*strings.Reader
}

func (_ fakeCloser) Close() error {
	return nil
}

func buildReader(s string) fakeCloser {
	stringReader := strings.NewReader(s)
	f := fakeCloser{
		Reader: stringReader,
	}
	return f
}

func TestNewWithOptions(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name               string
		constructorOptions Options
		usedOptions        Options
	}{
		{
			name: "defaults workers",
			constructorOptions: Options{
				BaseURL: "https://test.com",
			},
			usedOptions: Options{
				BaseURL: "https://test.com",
				Workers: DefaultOptions.Workers,
			},
		},
		{
			name: "defaults base URL",
			constructorOptions: Options{
				Workers: 1,
			},
			usedOptions: Options{
				BaseURL: DefaultOptions.BaseURL,
				Workers: 1,
			},
		},
		{
			name: "defaults http client",
			constructorOptions: Options{
				Workers: 1,
				BaseURL: "https://test.com",
			},
			usedOptions: Options{
				Workers: 1,
				BaseURL: "https://test.com",
			},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			client := NewWithOptions(tc.constructorOptions)

			require.Equal(t, tc.usedOptions.BaseURL, client.Options.BaseURL)
			require.Equal(t, tc.usedOptions.Workers, client.Options.Workers)
			require.NotNil(t, client.Options.HttpClient)
		})
	}
}

func TestReport(t *testing.T) {
	t.Parallel()
	respBody := `{"package_name":"requestts","package_type":"pypi", "package_data": { "status":"complete"} }`

	testdep := &v1types.Dependency{
		Name:      "requestts",
		Ecosystem: 1,
	}

	defaultOpts := Options{
		BaseURL: defaultEndpoint,
	}

	for _, tc := range []struct {
		name     string
		dep      *v1types.Dependency
		prepare  func(*fakeClient)
		expected *v1types.Reply
		mustErr  bool
		options  *Options
	}{
		{
			name: "normal",
			dep:  testdep,
			prepare: func(fc *fakeClient) {
				fc.resps = append(fc.resps, &http.Response{
					StatusCode: http.StatusOK,
					Body:       buildReader(respBody),
				})
			},
			expected: &v1types.Reply{
				PackageName: "requestts",
				PackageType: "pypi",
			},
		},
		{
			name: "no-dep-name",
			dep: &v1types.Dependency{
				Ecosystem: 1,
			},
			prepare: func(_ *fakeClient) {},
			mustErr: true,
		},
		{
			name: "no-dep-ecosystem",
			dep: &v1types.Dependency{
				Name: "test",
			},
			prepare: func(_ *fakeClient) {},
			mustErr: true,
		},
		{
			name: "http-fails",
			dep:  testdep,
			prepare: func(fc *fakeClient) {
				fc.errs = append(fc.errs, fmt.Errorf("fake error"))
			},
			mustErr: true,
		},
		{
			name: "http-non-200",
			dep:  testdep,
			prepare: func(fc *fakeClient) {
				fc.resps = append(fc.resps, &http.Response{
					Body:       buildReader(respBody),
					Status:     "Not found",
					StatusCode: 404,
				})
			},
			mustErr: true,
		},
		{
			name: "bad-response-json",
			dep:  testdep,
			prepare: func(fc *fakeClient) {
				fc.resps = append(fc.resps, &http.Response{
					Body:       buildReader("HEy Fr1end!"),
					Status:     "OK",
					StatusCode: 200,
				})
			},
			mustErr: true,
		},
		{
			name: "bad-ingestion-status",
			dep:  testdep,
			prepare: func(fc *fakeClient) {
				fc.resps = append(fc.resps, &http.Response{
					Body:       buildReader(`{"package_name":"requestts","package_type":"pypi", "package_data": { "status":"bad status"} }`),
					StatusCode: http.StatusOK,
				})
			},
			options: &Options{
				BaseURL:              defaultEndpoint,
				ErrOnFailedIngestion: true,
			},
			mustErr: true,
		},
		{
			name: "normal-err-non-ingested",
			dep:  testdep,
			prepare: func(fc *fakeClient) {
				fc.resps = append(fc.resps, &http.Response{
					StatusCode: http.StatusOK,
					Body:       buildReader(`{"package_name":"requestts","package_type":"pypi", "package_data": { "status":"failed"} }`),
				})
			},
			options: &Options{
				BaseURL:              defaultEndpoint,
				ErrOnFailedIngestion: true,
			},
			mustErr: true,
		},
		{
			name: "fail-retrying-timeout",
			dep:  testdep,
			prepare: func(fc *fakeClient) {
				fc.resps = append(fc.resps, &http.Response{
					StatusCode: http.StatusOK,
					Body:       buildReader(`{"package_name":"requestts","package_type":"pypi", "package_data": { "status":"pending"} }`),
				})
			},
			options: &Options{
				BaseURL:             defaultEndpoint,
				IngestionMaxRetries: 2,
				WaitForIngestion:    true,
			},
			mustErr: true,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fake := newFakeClient()
			tc.prepare(fake)
			if tc.options == nil {
				tc.options = &defaultOpts
			}
			client := &Trusty{
				Options: *tc.options,
			}
			client.Options.HttpClient = fake

			res, err := client.Report(context.Background(), tc.dep)
			if tc.mustErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, res)
			require.Equal(t, tc.expected.PackageName, res.PackageName)
		})
	}
}

func TestGroupReport(t *testing.T) {
	t.Parallel()
	respBody1 := `{"package_name":"requestts","package_type":"pypi"}`
	respBody2 := `{"package_name":"tensorflow","package_type":"pypi"}`

	testdep1 := &v1types.Dependency{
		Name:      "requestts",
		Ecosystem: 1,
	}
	testdep2 := &v1types.Dependency{
		Name:      "tensorflow",
		Ecosystem: 1,
	}

	for _, tc := range []struct {
		name     string
		deps     []*v1types.Dependency
		prepare  func(*fakeClient)
		expected []*v1types.Reply
		mustErr  bool
	}{
		{
			name: "normal",
			deps: []*v1types.Dependency{testdep1, testdep2},
			prepare: func(fc *fakeClient) {
				fc.resps = append(fc.resps,
					&http.Response{
						StatusCode: http.StatusOK,
						Body:       buildReader(respBody1),
					},
					&http.Response{
						StatusCode: http.StatusOK,
						Body:       buildReader(respBody2),
					},
				)
			},
			expected: []*v1types.Reply{
				{
					PackageName: "requestts",
					PackageType: "pypi",
				},
				{
					PackageName: "tensorflow",
					PackageType: "pypi",
				},
			},
		},

		{
			name: "no-dep-name",
			deps: []*v1types.Dependency{
				{Ecosystem: 1}, testdep1,
			},
			prepare: func(_ *fakeClient) {},
			mustErr: true,
		},
		{
			name: "no-dep-ecosystem",
			deps: []*v1types.Dependency{
				{Name: "test"}, testdep1,
			},
			prepare: func(_ *fakeClient) {},
			mustErr: true,
		},
		{
			name: "http-fails",
			deps: []*v1types.Dependency{testdep1},
			prepare: func(fc *fakeClient) {
				fc.errs = append(fc.errs, fmt.Errorf("fake error"))
			},
			mustErr: true,
		},
		{
			name: "http-non-200",
			deps: []*v1types.Dependency{testdep1, testdep2},
			prepare: func(fc *fakeClient) {
				fc.resps = append(
					fc.resps, &http.Response{
						Body:       buildReader(respBody1),
						Status:     "Not found",
						StatusCode: 404,
					},
					&http.Response{
						StatusCode: http.StatusOK,
						Body:       buildReader(respBody2),
					},
				)
				fc.errs = append(
					fc.errs, errors.New("HTTP Error"), nil,
				)
			},
			mustErr: true,
		},
		{
			name: "bad-response-json",
			deps: []*v1types.Dependency{testdep1, testdep2},
			prepare: func(fc *fakeClient) {
				fc.resps = append(fc.resps,
					&http.Response{
						Body:       buildReader("HEy Fr1end!"),
						Status:     "OK",
						StatusCode: 200,
					},
					&http.Response{
						StatusCode: http.StatusOK,
						Body:       buildReader(respBody2),
					},
				)
			},
			mustErr: true,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fake := newFakeClient()
			tc.prepare(fake)
			client := &Trusty{
				Options: Options{
					HttpClient: fake,
					BaseURL:    defaultEndpoint,
				},
			}

			res, err := client.GroupReport(context.Background(), tc.deps)
			if tc.mustErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, res)
			require.Equal(t, tc.expected, res)
		})
	}
}

func TestUrlFromEndpointAndPaths(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		baseUrl  string
		endpoint string
		params   map[string]string
		expected string
		mustErr  bool
	}{
		{
			name:     "no-query",
			endpoint: "/test/",
			baseUrl:  defaultEndpoint,
			expected: "https://api.trustypkg.dev/test/",
		},
		{
			name:     "query-string",
			endpoint: "/test/",
			baseUrl:  defaultEndpoint,
			params:   map[string]string{"key": "value"},
			expected: "https://api.trustypkg.dev/test/?key=value",
		},
		{
			name:     "invalid-base",
			endpoint: "/test/",
			baseUrl:  "Even!\nFlow!",
			mustErr:  true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			res, err := urlFromEndpointAndPaths(tc.baseUrl, tc.endpoint, tc.params)
			if tc.mustErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expected, res.String())

		})
	}
}

func TestPurlToDependency(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name     string
		purl     string
		expected *v1types.Dependency
		mustErr  bool
	}{
		{
			name:     "golang",
			purl:     "pkg:golang/github.com/k8s.io/release@v1.0.8",
			expected: &v1types.Dependency{Name: "github.com/k8s.io/release", Version: "v1.0.8", Ecosystem: v1types.ECOSYSTEM_GO},
			mustErr:  false,
		},
		{
			name:     "pypi",
			purl:     "pkg:pypi/requests@v1.2.3",
			expected: &v1types.Dependency{Name: "requests", Version: "v1.2.3", Ecosystem: v1types.ECOSYSTEM_PYPI},
			mustErr:  false,
		},
		{
			name:     "npm",
			purl:     "pkg:npm/%40react-stately/color@3.7.0",
			expected: &v1types.Dependency{Name: "@react-stately/color", Version: "3.7.0", Ecosystem: v1types.ECOSYSTEM_NPM},
			mustErr:  false,
		},
		{
			name:     "no-version",
			purl:     "pkg:npm/%40react-stately/color",
			expected: &v1types.Dependency{Name: "@react-stately/color", Version: "", Ecosystem: v1types.ECOSYSTEM_NPM},
			mustErr:  false,
		},
		{
			name:    "unsupported-ecosystem",
			purl:    "pkg:bugget/hello/there@1234",
			mustErr: true,
		},
		{
			name:    "invalid-purl",
			purl:    "http:npm/hello/there@1234",
			mustErr: true,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			dep, err := New().PurlToDependency(tc.purl)
			if tc.mustErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expected.Ecosystem, dep.Ecosystem)
			require.Equal(t, tc.expected.Name, dep.Name)
			require.Equal(t, tc.expected.Version, dep.Version)
		})
	}
}
