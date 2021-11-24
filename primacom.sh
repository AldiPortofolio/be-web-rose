#!/usr/bin/env bash

env GOOS=linux GOARCH=amd64 go build
scp rose-be-go nc_ketut@10.10.43.49:/home/nc_ketut/