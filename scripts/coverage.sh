#!/bin/sh
set -ue

coverprofile="/tmp/go-codegen-coverage.out"

pkg="./pkg/..."

covermode=${COVER_MODE:-"set"}

go test -coverpkg="${pkg}" -coverprofile="${coverprofile}" -covermode="${covermode}" ./...
go tool cover -html="$coverprofile"
