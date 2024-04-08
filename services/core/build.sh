#!/usr/bin/env bash

cd cmd/scrumerization
go build -tags netgo -ldflags '-s -w' -o ../../app