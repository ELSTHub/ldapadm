#!/bin/bash

# 检查是否提供了至少一个参数
if [ $# -eq 0 ]; then
    echo "请提供当前版本参数!"
    echo "使用方法: $0 <VERSION>"
    exit 1
fi


rm -rf dist
VERSION=$1

go env -w GOOS=linux
go env -w CGO_ENABLED=0
go build -ldflags "-s -w -X ldapadm/Version.VERSION=${VERSION}" -o dist/linux/bin/ldapadm ./main.go
echo "Linux build success!"

mkdir dist/linux/etc
echo "Create conf dir success!"

cp etc/ldapadm.yaml.example dist/linux/etc/ldapadm.yaml
echo "Copy config success!"

cp scripts/install.sh dist/linux/install.sh
echo "Copy install file success!"