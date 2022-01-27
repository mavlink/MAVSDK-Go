#!/usr/bin/env bash
docker run -it -v "$(dirname "$(pwd)")":/resources/ mavsdk

# To Build the docker
# docker build . -t mavsdk