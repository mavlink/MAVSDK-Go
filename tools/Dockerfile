# This is the docker file which is used to generate the code from proto buf files.
FROM --platform=linux/amd64 ubuntu:22.04

RUN apt-get update
RUN apt-get install -y wget python3 python3-distutils python3-dev python3-pip
RUN wget https://go.dev/dl/go1.22.3.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz
ENV PATH $PATH:/usr/local/go/bin:/root/go/bin
ENV PATH="$PATH:$(go env GOPATH)/bin"
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install golang.org/x/tools/cmd/goimports@latest

# Step 4: Install the python protoc-gendcsdk plugin to generate the class 
RUN apt-get install -y -q python3-pip git
RUN cd / && git clone https://github.com/mavlink/MAVSDK-Proto.git
RUN cd /MAVSDK-Proto/pb_plugins && pip3 install -r requirements.txt && pip3 install -e .
RUN pip3 install grpcio-tools==1.47.0
RUN apt install protobuf-compiler -y 
ENV PROTOCOL_BUFFERS_PYTHON_IMPLEMENTATION=python

ENTRYPOINT ["/bin/bash", "-l", "-c"]