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

//nolint:revive
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	v2client "github.com/stacklok/trusty-sdk-go/pkg/v2/client"
	v2types "github.com/stacklok/trusty-sdk-go/pkg/v2/types"
)

func main() {
	var endpoint, pname string
	flag.StringVar(&endpoint, "endpoint", "", "Trusty API endpoint to call")
	flag.StringVar(&pname, "pname", "", "Package name")
	flag.Parse()

	ctx := context.Background()
	client := v2client.New()

	switch endpoint {
	case "summary":
		if err := summary(ctx, client, pname); err != nil {
			fmt.Fprintf(os.Stderr, "error calling endpoint: %s\n", err)
			os.Exit(1)
		}
	case "pkg-meta":
		if err := pkg(ctx, client, pname); err != nil {
			fmt.Fprintf(os.Stderr, "error calling endpoint: %s\n", err)
			os.Exit(1)
		}
	case "alternatives":
		if err := alternatives(ctx, client, pname); err != nil {
			fmt.Fprintf(os.Stderr, "error calling endpoint: %s\n", err)
			os.Exit(1)
		}
	case "":
		fmt.Fprintf(os.Stderr, "endpoint is mandatory\n")
		os.Exit(1)
	default:
		fmt.Fprintf(os.Stderr, "invalid method: %s\n", endpoint)
		os.Exit(1)
	}
}

func summary(ctx context.Context, client v2client.Trusty, pname string) error {
	res, err := client.Summary(ctx, &v2types.Dependency{
		PackageName: pname,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)
	return nil
}

func pkg(ctx context.Context, client v2client.Trusty, pname string) error {
	res, err := client.PackageMetadata(ctx, &v2types.Dependency{
		PackageName: pname,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)
	fmt.Printf("STATUS: %+v\n", *res.Status)
	fmt.Printf("MALICIOUS: %+v\n", res.Malicious)
	for _, contributor := range res.Contributors {
		fmt.Printf("CONTRIBUTOR: %+v\n", contributor)
	}
	return nil
}

func alternatives(ctx context.Context, client v2client.Trusty, pname string) error {
	res, err := client.Alternatives(ctx, &v2types.Dependency{
		PackageName: pname,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)
	return nil
}
