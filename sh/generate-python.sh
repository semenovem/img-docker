#!/bin/bash

DIR_SH=$(dirname $(realpath "$0"))
source "$DIR_SH/variables.sh"


PATH_OUT_PYTHON="$PATH_OUT/python"
PATH_PLUGIN_PYTHON=$(which grpc_python_plugin)

if ! [ -d "$PATH_OUT_PYTHON" ]; then
    mkdir -p "$PATH_OUT_PYTHON"
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
    protoc \
       --proto_path="$PATH_PROTO_FILE" \
       --python_out="$PATH_OUT_PYTHON" \
       --grpc_python_out="$PATH_OUT_PYTHON" \
       --plugin=protoc-gen-grpc_python="$PATH_PLUGIN_PYTHON" \
         registry-service.proto
fi


unset PATH_OUT_PYTHON PATH_PLUGIN_PYTHON ERROR_MSG
