#!/bin/bash

source ./sh/constants.sh

docker build -t "$NAME_DOCKER_IMAGE" .
