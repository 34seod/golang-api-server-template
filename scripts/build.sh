#!/bin/bash

last_path=${PWD##*/} # get last folder
if [ $last_path = "scripts" ]; then
    cd ..
fi
build_version=$(date "+%Y%m%d%H%M%S")
docker build -t golang-api-server-template:${build_version} .
