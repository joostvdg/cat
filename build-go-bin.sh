#!/usr/bin/env bash

rm -rf bin/
env GOOS=linux GOARCH=amd64 go build -v -tags netgo -o bin/cat
