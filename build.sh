#!/bin/bash

TAG_NAME="v0.1.1"
IMAGE_NAME="argocd-sync-forcer"

set -e

# go to current shell file path 
SCRIPT_DIR=$(dirname "$(readlink -f "$0")")
echo "Change directory to $SCRIPT_DIR"
cd $SCRIPT_DIR

source .env

# build
#GOOS=linux GOARCH=amd64 go build -o ./bin/app .

# make a docker image
docker buildx build --platform linux/amd64  -t ${DOCKER_REPO}${IMAGE_NAME}:${TAG_NAME} .

# if DOCKER_REPO exists, push the image to repo
if [ -n "$DOCKER_REPO" ]; then
    docker push "${DOCKER_REPO}${IMAGE_NAME}:${TAG_NAME}"
fi

# run
