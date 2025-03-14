#!/bin/bash

cd app 
export GOOS="linux"
export GOARCH="amd64"
export CGO_ENABLED="0"
go build  -o bin/go-practice  ./cmd/main.go
