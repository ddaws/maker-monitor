#!/bin/bash

script_dir=$(dirname $0)
commit=$(git rev-parse --short=12 HEAD)

docker build -t ddaws/maker-monitor:$commit $script_dir/../
docker tag ddaws/maker-monitor:$commit ddaws/maker-monitor:latest
