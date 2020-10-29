#!/bin/bash

case $(uname | tr '[:upper:]' '[:lower:]') in
    darwin*) export OS='macos';;
    mingw*) export OS='win'
esac

if [ OS = 'macos' ]; then
    export DB2HOME=../clidriver
    export CGO_CFLAGS=-I${DB2HOME}/include
    export CGO_LDFLAGS=-L${DB2HOME}/lib
    export DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:${DB2HOME}/lib
else
    export PATH=../clidriver/bin:$PATH
fi

go mod vendor

go build -mod vendor ./main.go

./main

rm ./main