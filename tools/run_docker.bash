#!/usr/bin/env bash
docker run -it -v "$(dirname "$(pwd)")":/resources/ mavsdk bash

# To Build the docker
# docker build . -t mavsdk