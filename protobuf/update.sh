#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

if [ ! -d "$PROTOBUF" ]; then
  echo "Please set the PROTOBUF environment variable to the include folder of protobuf (e.g. /protobuf/include/)"
  exit -1
fi

BASEDIR=$(dirname "$0")

mkdir -p ../app/lib/generated
protoc --go_out=plugins=grpc:../go-stuff/src/github.com/milvum/hummingbird/proto \
       -I$BASEDIR $BASEDIR/hummingbird.proto
protoc --dart_out=grpc:../app/lib/generated -I$BASEDIR $BASEDIR/hummingbird.proto
protoc --dart_out=grpc:../app/lib/generated -I$PROTOBUF $PROTOBUF/google/protobuf/timestamp.proto
