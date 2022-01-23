FROM ubuntu:18.04

RUN apt-get update && \
    apt-get install -y curl
    
RUN curl -L https://github.com/ykhedar/MAVSDK/releases/download/v0.50.2APM/mavsdk_server_manylinux2010-x64 -o /root/mavsdk_server

RUN chmod +x /root/mavsdk_server

ENTRYPOINT /root/mavsdk_server -p 50051
