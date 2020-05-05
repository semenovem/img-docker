#!/bin/bash

source ./sh/constants.sh

#docker run --rm \
#-v "${PWD}/proto-files:/usr/app/$NAME_DIR_PROTO_FILES" \
#-v "${PWD}/pkg:/usr/app/$NAME_DIR_PKG_HOST_PC" \
# gen_proto
#
#exit



docker run --rm \
-it \
-u $(id -u "${USER}"):$(id -g "${USER}") \
-v "${PWD}/proto-files:/usr/app/$NAME_DIR_PROTO_FILES" \
-v "${PWD}/pkg:/usr/app/$NAME_DIR_PKG_HOST_PC" \
 gen_proto \
 bash



exit

-u $(id -u "${USER}"):$(id -g "${USER}") \
