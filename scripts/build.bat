@echo off

rmdir /s /q .\dist
mkdir dist
mkdir dist\linux
mkdir dist\linux\bin

go env -w GOOS=linux
go env -w CGO_ENABLED=0
go build -ldflags "-s -w" -o dist/linux/bin/ldapadm .\main.go
echo "Linux build success!"

mkdir dist\linux\etc
echo "Create conf dir success!"

copy etc\ldapadm.yaml.example dist\linux\etc\ldapadm.yaml
echo "Copy config success!"

copy scripts\install.sh dist\linux\install.sh
echo "Copy install file success!"