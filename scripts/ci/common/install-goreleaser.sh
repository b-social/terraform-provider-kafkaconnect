#!/usr/bin/env sh

[ -n "$DEBUG" ] && set -x
set -e
set -o pipefail

go get -d github.com/goreleaser/goreleaser
cd "$GOPATH"/src/github.com/goreleaser/goreleaser
dep ensure -vendor-only
make setup build
go install