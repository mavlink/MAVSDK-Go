## generating the protobuf files.
Steps to generate:


Generate the .pb.go and .pb.rpc.go files using the plugins: proto-gen-go and proto-gen-gorpc respectively.
These require the use of the protoc compiler and also the protobuf files.

Now the proto-gen-dcsdk can be used to generate the MAVSDK language related files using the jinja2 templates.



## List of Protobuf Message Support
1. Action: action.proto
2. Telemetry: telemetry.proto
3. Geofence: geofence.proto
4. Mission: mission.proto
5. Log Files: log_files.proto 
6. Core: core.proto
 
