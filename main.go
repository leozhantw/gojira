package main

import "github.com/leozhantw/gojira/cmd"

var (
	// version and date are dynamically set by goreleaser.
	version = "dev"
	date    = ""
)

func main() {
	cmd.Version = version
	cmd.BuildDate = date
	cmd.Execute()
}
