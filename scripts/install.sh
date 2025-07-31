#!/bin/bash

PREFIX=/usr/local/ldapadm
SYS_CONF_PATH=/etc/ldapadm

mkdir -p $SYS_CONF_PATH

cp ./etc/ldapadm.yaml $SYS_CONF_PATH/ldapadm.yaml

chmod 600 $SYS_CONF_PATH/*

mkdir -p $PREFIX/bin

cp -r ./bin/* $PREFIX/bin

chmod +x $PREFIX/bin/*

chmod 4755 $PREFIX/bin/*

ln -s $PREFIX/bin/* /usr/local/bin/
