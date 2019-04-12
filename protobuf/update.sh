#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

BASEDIR=$(dirname "$0")
mkdir -p ../app/lib/generated
protoc --dart_out=grpc:../app/lib/generated -I$BASEDIR $BASEDIR/hummingbird.proto
