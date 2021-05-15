## A package for loading config data from AWS SecretsManager

![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/gofor-little/cfg?include_prereleases)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gofor-little/cfg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://raw.githubusercontent.com/gofor-little/cfg/main/LICENSE)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/gofor-little/cfg/CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofor-little/cfg)](https://goreportcard.com/report/github.com/gofor-little/cfg)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gofor-little/cfg)](https://pkg.go.dev/github.com/gofor-little/cfg)

### Introduction
* Secure config loading.
* Support AWS SecretsManager.
* Easy JSON unmarshaling into a struct.

### Example
```go
package main

import (
	"context"

	"github.com/gofor-little/cfg"
)

type Config struct {
	SomeStringValue string `json:"someStringValue"`
	SomeIntValue    int    `json:"someIntValue"`
}

func main() {
	// Initialize the cfg package.
	if err := cfg.Initialize("AWS_PROFILE", "AWS_REGION"); err != nil {
		panic(err)
	}

	// Load and parse the config data into the passed Config struct.
	cfg := &Config{}
	if err := cfg.Load(context.Background(), "AWS_SECRET_ARN", cfg); err != nil {
		panic(err)
	}
}
```

### Testing
Ensure the following environment variables are set, usually with a .env file.
* ```AWS_PROFILE``` (an AWS CLI profile name)
* ```AWS_REGION``` (a valid AWS region)

Run ```go test -v ./...``` in the root directory.