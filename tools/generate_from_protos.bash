#!/usr/bin/env bash

set -e

command -v protoc || { echo >&2 "Protobuf needs to be installed (e.g. '$ apt install protobuf-compiler') for this script to run!"; exit 1; }

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROTO_DIR=${PROTO_DIR:-"${SCRIPT_DIR}/../proto/protos"}
OUTPUT_DIR=${OUTPUT_DIR:-"${SCRIPT_DIR}/../Sources/"}
PROTO_DIR_TMP=${PROTO_DIR_TMP:-"${SCRIPT_DIR}/tmp/protos"}
export TEMPLATE_PATH="$(pwd)/../templates/"



PLUGIN_LIST="action core mission geofence telemetry log_files"

mkdir -p ${PROTO_DIR_TMP}
cp -r ${PROTO_DIR}/* ${PROTO_DIR_TMP}

for plugin in ${PLUGIN_LIST}; do
    sed -i "/java_package.*/c option go_package = \".;${plugin}\";" ${PROTO_DIR_TMP}/$plugin/$plugin.proto
    cp ${PROTO_DIR_TMP}/mavsdk_options.proto ${PROTO_DIR_TMP}/$plugin/mavsdk_options.proto 
    sed -i "/java_package.*/c option go_package = \".;${plugin}\";" ${PROTO_DIR_TMP}/$plugin/mavsdk_options.proto
done

PROTO_DIR=${PROTO_DIR_TMP}

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

GO_GEN_CMD="/go/bin/protoc-gen-go"
GO_GEN_RPC_CMD="/go/bin/protoc-gen-gogrpc"

# Generate the message and service definitions using grpc plugins.
for plugin in ${PLUGIN_LIST}; do
    mkdir -p ${OUTPUT_DIR}/$plugin 
    protoc ${plugin}.proto -I${PROTO_DIR}/$plugin --go_out=${OUTPUT_DIR}/$plugin --gogrpc_out=${OUTPUT_DIR}/$plugin --plugin=protoc-gen-go=${GO_GEN_CMD} --plugin=protoc-gen-gogrpc=${GO_GEN_RPC_CMD}
done


# Generate the final plugins
for plugin in ${PLUGIN_LIST}; do
	echo "+=> Doing $plugin"
	python3 -m grpc_tools.protoc --plugin=protoc-gen-custom=$(which protoc-gen-mavsdk) -I${PROTO_DIR}/$plugin --custom_out=${OUTPUT_DIR}/$plugin --custom_opt=file_ext=go ${plugin}.proto
done

# Remove the temp directory.
rm ${PROTO_DIR_TMP}
