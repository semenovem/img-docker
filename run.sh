#!/bin/bash

source ./sh/constants.sh

if ! [ -d "pkg" ]; then
    make -p "pkg"
fi



docker run --rm \
-v "${PWD}/proto-files:/usr/app/$NAME_DIR_PROTO_FILES" \
-v "${PWD}/pkg:/usr/app/$NAME_DIR_PKG_HOST_PC" \
 "$NAME_DOCKER_IMAGE"

exit



docker run --rm \
-it \
-v "${PWD}/proto-files:/usr/app/$NAME_DIR_PROTO_FILES" \
-v "${PWD}/pkg:/usr/app/$NAME_DIR_PKG_HOST_PC" \
 "$NAME_DOCKER_IMAGE" \
 bash



exit

-u $(id -u "${USER}"):$(id -g "${USER}") \
