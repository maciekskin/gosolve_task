package main

import "os"

func main() {
	err := runApp()
	if err != nil {
		// TODO: log error
		os.Exit(-1)
	}
}

func runApp() error {
	return nil
}

// TODO:
// - create http server and define a GET endpoint
// - add go.mod
// - use separate request/response structs and internal data model
// - load data from input.txt to slice in a repository during app start
// - implement lookup service as separate module (see description in recruitment_task.md)
// - add logging
// - add configuration file with service port and log level (Debug, Info, Error)
// - write unit tests
// - add README.md with service description
// - write Makefile
// - return result nested in nice struct that can optionally return an error
// - save periodically and commit to show your iterational code changes in final repository
// - upload to GitHub
