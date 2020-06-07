# This is the docker file which is used to generate the code from proto buf files.
FROM shifty21/golang_mav

RUN apt install -y -q python3-pip

RUN cd /resources/pb_plugins && pip3 install -r requirements.txt && pip3 -e .

ENTRYPOINT [ "/bin/bash"]