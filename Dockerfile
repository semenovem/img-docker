FROM ubuntu:18.04

#RUN apt-get install python

RUN apt-get update
RUN apt-get upgrade
RUN yes | apt-get install python3.7

    #    apt-get install software-properties-common \
    #    add-apt-repository ppa:deadsnakes/ppa \


#RUN groupadd --gid 2000 node \
#  && useradd --uid 2000 --gid node --shell /bin/bash --create-home node

WORKDIR /usr/app

#COPY 'protoc-3.11.4-linux-x86_64' ./util

#RUN mkdir build

#CMD /usr/src/app/util/bin/protoc \
#   --proto_path=source-proto-files \
#   --python_out=build \
#   --grpc_python_out=build
#     registry-service.proto
#
#~/_dev/docker/gen-grpc-python/protoc-3.11.4-linux-x86_64/bin
