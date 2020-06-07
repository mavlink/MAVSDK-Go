# MAVSDK-Go
This package contains the GoLang Plugin for the MAVSDK. This is still work in progress so use at your own descretion.

## generating the protobuf files.
Steps to generate:


Generate the .pb.go and .pb.rpc.go files using the plugins: proto-gen-go and proto-gen-gorpc respectively.
These require the use of the protoc compiler and also the protobuf files.


## List of Protobuf Message Support
1. Action: action.proto
2. Telemetry: telemetry.proto
3. Geofence: geofence.proto
4. Mission: mission.proto
5. Log Files: log_files.proto 
6. Core: core.proto
 

## generate plugin classes from Jinja

Now the proto-gen-dcsdk can be used to generate the MAVSDK language related files using the jinja2 templates.
In order to generate the plugin classes run the following command from the root of the MAVSDK-Go

export TEMPLATE_PATH="$(pwd)/templates/"
protoc --plugin=protoc-gen-custom=$(which protoc-gen-dcsdk) -I./proto/protos/action --custom_out=. --custom_opt=file-ext=go ./proto/protos/action/action.proto