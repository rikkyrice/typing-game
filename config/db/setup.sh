#!/bin/sh
set -e
. ./env.list

# Build Docker image
docker pull ibmcom/db2:11.5.4.0
docker build -t ${REPOSITORY}:${TAG} .