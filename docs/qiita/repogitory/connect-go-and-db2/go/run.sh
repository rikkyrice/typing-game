#!/bin/bash
export PATH=../clidriver/bin:$PATH

go build -mod vendor ./main.go

./main

rm ./main