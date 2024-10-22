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

PLUGIN_LIST="action failure manual_control server_utility action_server follow_me shell
arm_authorizer_server ftp mission telemetry calibration ftp_server mission_raw telemetry_server
camera geofence mission_raw_server transponder camera_server gimbal mocap tune component_metadata
gripper offboard winch component_metadata_server info param core log_files param_server log_streaming rtk"

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
    camel_case_filename="$(snake_case_to_camel_case "$plugin").go"
    actual_file_path="${OUTPUT_DIR}/$plugin/$camel_case_filename"
    new_file_path="${OUTPUT_DIR}/$plugin/${plugin}.go"
    temp_file_path="${OUTPUT_DIR}/$plugin/temp.go"

    # Check if the plugin is in snake_case and if the corresponding file exists
    if [[ -f "$actual_file_path" ]]; then
        echo "Renaming $actual_file_path to $new_file_path"
        
        # Move the original CamelCase file to a temporary name
        mv "$actual_file_path" "$temp_file_path"
        
        # Rename the temporary file to the new name
        mv "$temp_file_path" "$new_file_path"
        
    fi
done

# Remove the temp directory.
rm -rf ${PROTO_DIR_TMP}
