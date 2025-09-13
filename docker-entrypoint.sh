#!/bin/bash

set -e

# Extract the last argument (expected to be the path(s) for protoc)
paths=${!#}
# Extract all arguments except the last (expected to be protoc flags)
args=${@:1:$#-1}

protoc $args $paths
