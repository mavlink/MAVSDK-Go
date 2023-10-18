#!/usr/bin/env bash

set -e

function snake_case_to_camel_case {
    echo $1 | awk -v FS="_" -v OFS="" '{for (i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) substr($i,2)} 1'
}

command -v protoc || { echo >&2 "Protobuf needs to be installed (e.g. '$ apt install protobuf-compiler') for this script to run!"; exit 1; }

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROTO_DIR=${PROTO_DIR:-"${SCRIPT_DIR}/../proto/protos"}
OUTPUT_DIR=${OUTPUT_DIR:-"${SCRIPT_DIR}/../Sources/"}
PROTO_DIR_TMP=${PROTO_DIR_TMP:-"${SCRIPT_DIR}/tmp/protos"}
export TEMPLATE_PATH="$(pwd)/../templates/"

PLUGIN_LIST="action calibration camera core failure follow_me ftp geofence gimbal info log_files manual_control mission mission_raw
             mocap offboard param shell telemetry transponder tune
             action_server mission_raw_server param_server telemetry_server tracking_server"

echo "Plugin List consist of: " ${PLUGIN_LIST}
rm -rf ${PROTO_DIR_TMP}
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

GO_GEN_CMD="/root/go/bin/protoc-gen-go"
GO_GEN_RPC_CMD="/root/go/bin/protoc-gen-go-grpc"

echo "Generating proto definitions."
# Generate the message and service definitions using grpc plugins.
for plugin in ${PLUGIN_LIST}; do
    mkdir -p ${OUTPUT_DIR}/$plugin 
    protoc ${plugin}.proto -I${PROTO_DIR}/$plugin --go_out=${OUTPUT_DIR}/$plugin --gogrpc_out=${OUTPUT_DIR}/$plugin --plugin=protoc-gen-go=${GO_GEN_CMD} --plugin=protoc-gen-gogrpc=${GO_GEN_RPC_CMD}
done

echo "Generating final plugins."
# Generate the final plugins
for plugin in ${PLUGIN_LIST}; do
	echo "+=> Doing $plugin"
	python3 -m grpc_tools.protoc --plugin=protoc-gen-custom=$(which protoc-gen-mavsdk) -I${PROTO_DIR}/$plugin --custom_out=${OUTPUT_DIR}/$plugin --custom_opt=file_ext=go ${plugin}.proto
    # Again move generated file to its place.
    mv ${OUTPUT_DIR}/$plugin/$(snake_case_to_camel_case ${plugin}).go ${OUTPUT_DIR}/$plugin/temp.go
    mv ${OUTPUT_DIR}/$plugin/temp.go ${OUTPUT_DIR}/$plugin/$plugin.go
done

# Remove the temp directory.
rm -rf ${PROTO_DIR_TMP}
