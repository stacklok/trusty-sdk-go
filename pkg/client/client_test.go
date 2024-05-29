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
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stacklok/trusty-sdk-go/pkg/types"
)

// fakeClient mocks the http client used by the trusty client
type fakeClient struct {
	resp *http.Response
	err  error
}

func (fc *fakeClient) Do(_ *http.Request) (*http.Response, error) {
	return fc.resp, fc.err
}

func buildReader(s string) io.ReadCloser {
	stringReader := strings.NewReader(s)
	return io.NopCloser(stringReader)
}

func TestReport(t *testing.T) {
	t.Parallel()
	respBody := `{"package_name":"requestts","package_type":"pypi"}`

	testdep := &types.Dependency{
		Name:      "requestts",
		Ecosystem: 1,
	}

	for _, tc := range []struct {
		name     string
		dep      *types.Dependency
		prepare  func(*fakeClient)
		expected *types.Reply
		mustErr  bool
	}{
		{
			name: "normal",
			dep:  testdep,
			prepare: func(fc *fakeClient) {
				fc.resp = &http.Response{
					StatusCode: http.StatusOK,
					Body:       buildReader(respBody),
				}
			},
			expected: &types.Reply{
				PackageName: "requestts",
				PackageType: "pypi",
			},
		},
		{
			name: "no-dep-name",
			dep: &types.Dependency{
				Ecosystem: 1,
			},
			prepare: func(_ *fakeClient) {},
			mustErr: true,
		},
		{
			name: "no-dep-ecosystem",
			dep: &types.Dependency{
				Name: "test",
			},
			prepare: func(_ *fakeClient) {},
			mustErr: true,
		},
		{
			name: "http-fails",
			dep:  testdep,
			prepare: func(fc *fakeClient) {
				fc.err = fmt.Errorf("fake error")
			},
			mustErr: true,
		},
		{
			name: "http-non-200",
			dep:  testdep,
			prepare: func(fc *fakeClient) {
				fc.resp = &http.Response{
					Body:       buildReader(respBody),
					Status:     "Not found",
					StatusCode: 404,
				}
			},
			mustErr: true,
		},
		{
			name: "bad-response-json",
			dep:  testdep,
			prepare: func(fc *fakeClient) {
				fc.resp = &http.Response{
					Body:       buildReader("HEy Fr1end!"),
					Status:     "OK",
					StatusCode: 200,
				}
			},
			mustErr: true,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fake := &fakeClient{}
			tc.prepare(fake)
			client := &Trusty{
				Options: Options{
					HttpClient: fake,
					BaseURL:    defaultEndpoint,
				},
			}

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
