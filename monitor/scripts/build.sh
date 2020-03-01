#!/bin/bash

script_dir=$(dirname $0)
GOOS=linux GOARCH=386 go build -o $script_dir/../build/monitor $script_dir/../main.go
