#!/bin/bash
export PATH=../clidriver/bin:$PATH

go mod vendor

go build -mod vendor ./main.go

./main

rm ./main