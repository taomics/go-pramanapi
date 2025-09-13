#!/bin/bash

set -e

paths=${!#}
args=${@:1:$#-1}

protoc $args $paths
