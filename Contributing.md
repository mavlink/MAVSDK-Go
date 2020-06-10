More advanced user can also generate the files from the protobuf definition. In order to do this the go packages proto-gen-go and proto-gen-gorpc are required.
This is already installed in the docker file available in tools/Dockerfile.

Build this dockerfile using
```bash
docker build . -t mavsdk
```

Run the image using the following command with the volume with MAVSDK-GO mounted in the docker.

```bash
bash run_docker.bash
```

Then run the script ```generate_from_protos.bash``` to generate the plugins. The script is more or less self explanatory.


Now the proto-gen-dcsdk can be used to generate the MAVSDK language related files using the jinja2 templates.
In order to generate the plugin classes run the following command from the root of the MAVSDK-Go

```bash
export TEMPLATE_PATH="$(pwd)/templates/"
protoc --plugin=protoc-gen-custom=$(which protoc-gen-dcsdk) -I./proto/protos/action --custom_out=. --custom_opt=file-ext=go ./proto/protos/action/action.proto

```

## Note
There was a problem generating the plugins in docker. If you get similar problems, please generate the plugins on a normal machine (ubuntu)