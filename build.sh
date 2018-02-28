#!/usr/bin/env bash
set -e

echo goimports
goimports -l -w $(go list -f {{.Dir}} ./...)

echo go test
go test -cover ./...

echo go vet
go vet ./...

echo dep status
if ! ds=$(dep status 2> /dev/null); then
    echo "${ds}"
    echo "Run 'dep ensure' to fix"
    exit 1
fi
