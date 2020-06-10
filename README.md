# MAVSDK-Go
This package contains the Go wrapper based on gRPC for the MAVSDK. This is still work in progress so use at your own descretion.
MAVSDK server (also called backend) needs to be running on the same system as this plugin. The Go Wrapper communicates using gRPC to the server.

MAVSDK uses Google Protobuf based message definitions in the protobuf files. [MAVSDK-ProtoBuf](https://github.com/mavlink/MAVSDK-Proto)

The wrapper can be used directly using the generated files in Source folder.
A example file is provided in the example folder which contains example for take-off, land, geofence upload and telemetry data.


## Contributing
For general users, it is enough to use the generated files as such. But in order to keep the sources updated based on the latest proto definitions it is required to generate the source file again.
This can be done using the instructions from ```Contributing.md```


## Supported Plugins

Currently only the following plugins are supported.
1. Action: action.proto
2. Telemetry: telemetry.proto
3. Geofence: geofence.proto
4. Mission: mission.proto
5. Log Files: log_files.proto 
6. Core: core.proto


## Issues/Questions
For any problems with the wrapper please create an issue.

Happy flying :)