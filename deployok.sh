#!/usr/bin/env bash

env GOOS=linux GOARCH=amd64 go build
scp rose-be-go syandi@34.101.247.136:/home/syandi/rose-be-go