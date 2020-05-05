#!/bin/bash

DIR_SH=$(dirname $(realpath "$0"))

source sh/cmd.sh &&

bash "$DIR_SH/generate-python.sh"
bash "$DIR_SH/generate-golang.sh"
