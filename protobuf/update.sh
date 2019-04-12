#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

BASEDIR=$(dirname "$0")

protoc --dart_out=grpc:../app/lib/generated -I$BASEDIR $BASEDIR/hummingbird.proto
