package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	// "text/template"

	"github.com/stacklok/trusty-sdk-go/internal/gen"
)

func main() {
	fd, err := os.Open("v2.json")
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
	defer fd.Close()

	bs, err := io.ReadAll(fd)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	var result gen.OapiSpec
	if err := json.Unmarshal(bs, &result); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", result)
}
