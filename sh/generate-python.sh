#!/bin/bash

DIR_SH=$(dirname $(realpath "$0"))
source "$DIR_SH/current-version.sh"
source "$DIR_SH/variables.sh"


PATH_OUT_PYTHON="$PATH_OUT/python"
PATH_PLUGIN_PYTHON=$(which grpc_python_plugin)

if ! [ -d "$PATH_OUT_PYTHON" ]; then
    mkdir -p "$PATH_OUT_PYTHON"
    chmod -R 777 "$PATH_OUT_PYTHON"
fi


# проверки
ERROR_MSG=""
if ! [ -d "$PATH_OUT_PYTHON" ]; then
    ERROR_MSG+="Нет директории для python: '$PATH_OUT_PYTHON'"
fi
if ! [ -f "$PATH_PLUGIN_PYTHON" ]; then
    ERROR_MSG+="Нет плагина для python: '$PATH_PLUGIN_PYTHON'"
fi


if [ "$ERROR_MSG" ]; then
    echo "Ошибка! Нельзя выполнить генерацию кода для python"
    echo -e "$ERROR_MSG"
else
    echo "Генерация кода для python"

    PATH_OUT_TMP=$(pwd)/tmp
    mkdir "$PATH_OUT_TMP"

    protoc \
       --proto_path="$PATH_PROTO_FILE" \
       --python_out="$PATH_OUT_TMP" \
       --grpc_python_out="$PATH_OUT_TMP" \
       --plugin=protoc-gen-grpc_python="$PATH_PLUGIN_PYTHON" \
         registry-service.proto

    cd "$PATH_OUT_TMP" > /dev/null
    tar -czf "$PATH_OUT_PYTHON/python-$VERSION.tar.gz" *
    cd - > /dev/null

    cp -R "$PATH_OUT_TMP"/* "$PATH_OUT_PYTHON"/

    chmod -R 777 "$PATH_OUT_PYTHON"

    rm -rf "$PATH_OUT_TMP"
fi


unset PATH_OUT_PYTHON PATH_PLUGIN_PYTHON ERROR_MSG
