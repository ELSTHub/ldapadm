#!/bin/bash

rm -rf dist
VERSION=v0.2.0

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