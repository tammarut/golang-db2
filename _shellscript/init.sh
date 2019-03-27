#!/bin/sh
export GO111MODULE=on
export DB2HOME=$PWD/_clidriver/linux
export CGO_CFLAGS=-I$DB2HOME/include
export CGO_LDFLAGS=-L$DB2HOME/lib
export LD_LIBRARY_PATH=$DB2HOME/lib