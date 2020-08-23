# To Run the docker 
docker run -it -v "$(dirname "$(pwd)")":/resources/ mavsdk

# To Build the docker
# docker build . -t mavsdk