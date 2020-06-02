FROM ubuntu
RUN apt-get update && apt-get install -y git curl
RUN git clone https://github.com/PX4/Firmware.git --recursive
RUN apt-get install -y wget
RUN cd ~/ && wget https://github.com/mavlink/MAVSDK/releases/download/v0.26.0/mavsdk_server_manylinux2010-x64
RUN chmod +x ~/mavsdk_server_manylinux2010-x64
RUN cd Firmware && bash ./Tools/setup/ubuntu.sh --no-nuttx

# Run the MAVSDK server
RUN cd ~ && ./mavsdk_server_manylinux2010-x64 &

# Run the Simulator
RUN cd /Firmware
CMD ["HEADLESS=1 make px4_sitl jmavsim"]
#4560 - simulator port
