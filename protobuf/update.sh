#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

BASEDIR=$(dirname "$0")
mkdir -p ../app/lib/generated
protoc --go_out=plugins=grpc:../go-stuff/src/github.com/milvum/hummingbird/proto \
       --dart_out=grpc:../app/lib/generated \
       -I$BASEDIR $BASEDIR/hummingbird.proto
