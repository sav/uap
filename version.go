// From: github.com/MichaelMure/mdr

package main

import (
	"fmt"
	"runtime"
)

// These variables are initialized externally during the build. See the Makefile.
var GitCommit string
var GitLastTag string
var GitExactTag string

func printVersion() {
	if GitExactTag == "undefined" {
		GitExactTag = ""
	}

	version := GitLastTag

	if GitLastTag == "" {
		version = fmt.Sprintf("%s-dev-%.10s", version, GitCommit)
	}

	if GitCommit == "" {
		fmt.Println("UAP version: unknown (not compiled with the makefile)")
	} else {
		fmt.Printf("UAP version: %s\n", version)
	}

	fmt.Printf("System version: %s/%s\n", runtime.GOARCH, runtime.GOOS)
	fmt.Printf("Golang version: %s\n", runtime.Version())
}
