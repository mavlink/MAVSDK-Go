# This is the docker file which is used to generate the code from proto buf files.
FROM ubuntu

RUN apt-get update
RUN apt-get install -y wget
RUN wget https://go.dev/dl/go1.17.6.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz
ENV PATH $PATH:/usr/local/go/bin
CMD source ~/.bashrc
RUN echo $PATH
ENV PATH="$PATH:$(go env GOPATH)/bin"
RUN echo $PATH
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Step 4: Install the python protoc-gendcsdk plugin to generate the class 
RUN apt-get install -y -q python3-pip git
RUN cd / && git clone https://github.com/mavlink/MAVSDK-Proto.git
RUN cd /MAVSDK-Proto/pb_plugins && pip3 install -r requirements.txt && pip3 install -e .
RUN pip3 install grpcio-tools
RUN apt-get install -y protobuf-compiler

ENTRYPOINT ["/bin/bash", "-l", "-c"]