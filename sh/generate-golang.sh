#!/bin/bash

export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN

DIR_SH=$(dirname $(realpath "$0"))
source sh/cmd.sh &&
source "$DIR_SH/variables.sh"


PATH_OUT_GOLANG="$PATH_OUT/golang"

if ! [ -d "$PATH_OUT_GOLANG" ]; then
    mkdir -p "$PATH_OUT_GOLANG"
    chmod -R 777 "$PATH_OUT_GOLANG"
fi


# проверки
ERROR_MSG=""
if ! [ -d "$PATH_OUT_GOLANG" ]; then
    ERROR_MSG+="Нет директории для golang: '$PATH_OUT_GOLANG'"
fi



if [ "$ERROR_MSG" ]; then
    echo "Ошибка! Нельзя выполнить генерацию кода для golang"
    echo -e "$ERROR_MSG"
else
    echo "Генерация кода для golang"
    protoc \
       --proto_path="$PATH_PROTO_FILE" \
       --go_out=plugins=grpc:"$PATH_OUT_GOLANG" \
        registry-service.proto

    chmod -R 777 "$PATH_OUT_GOLANG"
fi

unset DIR_SH PATH_OUT_GOLANG ERROR_MSG
