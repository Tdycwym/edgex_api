#!/bin/bash

RUN_NAME="edgex_api"

export GO111MODULE=on
export GOPROXY=https://goproxy.cn
export GOSUMDB=off
go mod download
mkdir -p output/bin output/log
cp script/bootstrap.sh output 2>/dev/null
chmod +x output/bootstrap.sh

go build -o output/bin/${RUN_NAME}
