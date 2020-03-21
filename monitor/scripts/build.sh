#!/bin/bash

script_dir=$(dirname $0)
commit=$(git rev-parse --short=12 HEAD)

docker build -t monitor:$commit $script_dir/../
