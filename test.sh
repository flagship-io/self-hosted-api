#!/bin/bash

go test -coverpkg=./... ./... -coverprofile cover.out.tmp -covermode=atomic
cat cover.out.tmp | grep -v "_mock.go" > cover.out
go tool cover -html=cover.out