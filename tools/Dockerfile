# This is the docker file which is used to generate the code from proto buf files.
FROM shifty21/golang_mav
# This docker image already contains the protobuf and grpc plugins for go languages. The Dockerfile 
# will be shared soon.
# Step 1: Install the protoc-gen-go plugin
# Step 2: Install the protoc-gen-gogrpc
# Step 3: Setup the go paths.

# Step 4: Install the python protoc-gendcsdk plugin to generate the class 
RUN apt install -y -q python3-pip git
RUN cd / && git clone https://github.com/mavlink/MAVSDK-Proto.git
RUN cd /MAVSDK-Proto/pb_plugins && pip3 install -r requirements.txt && pip3 -e .

ENTRYPOINT [ "/bin/bash"]