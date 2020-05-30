#!/usr/bin/env bash

set -e

command -v protoc || { echo >&2 "Protobuf needs to be installed (e.g. '$ apt install protobuf-compiler') for this script to run!"; exit 1; }

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PB_PLUGINS_DIR=${PB_PLUGINS_DIR:-"${SCRIPT_DIR}/../proto/pb_plugins"}
PROTO_DIR=${PROTO_DIR:-"${SCRIPT_DIR}/../proto/protos"}
OUTPUT_DIR=${OUTPUT_DIR:-"${SCRIPT_DIR}/../Sources/MAVSDK-Go/Generated"}

PLUGIN_LIST="action core mission geofence telemetry log_files"

if [ ! -d ${PROTO_DIR} ]; then
    echo "Script is not in the right location! It will look for the proto files in '${PROTO_DIR}', which doesn't exist!"

    exit 1
fi

if [ ! -d ${OUTPUT_DIR} ]; then
    echo "Script is not in the right location! It is made to generate the files in '${OUTPUT_DIR}', which doesn't exist!"

    exit 1
fi

echo "-------------------------------"
echo "Generating pb and grpc.pb files"
echo "-------------------------------"

GO_GEN_CMD=""
GO_GEN_RPC_CMD=""

for plugin in ${PLUGIN_LIST}; do
    protoc ${plugin}.proto -I${PROTO_DIR}/${plugin} --go_out=${OUTPUT_DIR} --gogrpc_out=${OUTPUT_DIR} --plugin=protoc-gen-go=${GO_GEN_CMD} --plugin=protoc-gen-gogrpc=${GO_GEN_RPC_CMD}
done
