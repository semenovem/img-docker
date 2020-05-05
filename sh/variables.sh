#!/bin/bash

DIR_SH=$(dirname $(realpath "$0"))
source "$DIR_SH/constants.sh"
source "$DIR_SH/current-version.sh"

PATH_PROTO_FILE="/usr/app/$NAME_DIR_PROTO_FILES"
PATH_OUT="/usr/app/$NAME_DIR_PKG_HOST_PC/$VERSION"

if ! [ -d "$PATH_OUT" ]; then
    mkdir -p "$PATH_OUT"
    chmod -R 777 "$PATH_OUT"
fi

# проверки
ERROR_MSG=""

if ! [ -d "$PATH_PROTO_FILE" ]; then
    ERROR_MSG+="Нет директории с proto файлами\n"
fi

if ! [ -d "$PATH_OUT" ]; then
    ERROR_MSG+="Нет директории для выхода компиляции"
fi

if [ "$ERROR_MSG" ]; then
    echo -e ERROR_MSG
    exit 1
fi

unset ERROR_MSG
