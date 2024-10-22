# MAVSDK-Go
This package contains the Go wrapper based on gRPC for the MAVSDK. This is still work in progress so use at your own descretion.
MAVSDK server (also called backend) needs to be running on the same system as this plugin. The Go Wrapper communicates using gRPC to the server.

MAVSDK uses Google Protobuf based message definitions in the protobuf files. [MAVSDK-ProtoBuf](https://github.com/mavlink/MAVSDK-Proto)

A example is also provided which contains example usage of the mavsdk-go package.


## Contributing
For general users, it is enough to use the generated files as such. But in order to keep the sources updated based on the latest proto definitions it is required to generate the source file again.
This can be done using the instructions from ```Contributing.md```

If you have any issues or questions, please feel free to open an issue in the repository. We will try to address them as soon as possible.

Happy flying :)