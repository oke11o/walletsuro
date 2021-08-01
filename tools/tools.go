// +build tools

//go:generate go build -o ../bin/mockgen github.com/golang/mock/mockgen
//go:generate go build -o ../bin/swagger github.com/go-swagger/go-swagger/cmd/swagger
//go:generate go build -o ../bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate
//go:generate bash -c "go build -ldflags \"-X 'main.version=$(go list -m -f '{{.Version}}' github.com/golangci/golangci-lint)' -X 'main.commit=test' -X 'main.date=test'\" -o ../bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint"

// Package tools contains go:generate commands for all project tools with versions stored in local go.mod file
// See https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
