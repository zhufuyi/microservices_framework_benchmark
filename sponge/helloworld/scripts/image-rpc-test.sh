#!/bin/bash

# build rpc service test image

serverName="helloworld"
# image name of the service, prohibit uppercase letters in names.
IMAGE_NAME="helloworld/helloworld.rpc-test"
# Dockerfile file directory
DOCKERFILE_PATH="scripts/build"
DOCKERFILE="${DOCKERFILE_PATH}/Dockerfile_test"

# image repo address, REPO_HOST="ip or domain", passed in via the first parameter
REPO_HOST=$1
if [ "X${REPO_HOST}" = "X" ];then
        echo "param 'repo host' cannot be empty, example: ./image-rpc-test.sh hub.docker.com v1.0.0"
        exit 1
fi
# the version tag, which defaults to latest if empty, is passed in via the second parameter
TAG=$2
if [ "X${TAG}" = "X" ];then
        TAG="latest"
fi
# image name and tag
IMAGE_NAME_TAG="${REPO_HOST}/${IMAGE_NAME}:${TAG}"

PROJECT_FILES=$(ls)
tar zcf ${serverName}.tar.gz ${PROJECT_FILES}
mv -f ${serverName}.tar.gz ${DOCKERFILE_PATH}

echo "docker build -f ${DOCKERFILE} -t ${IMAGE_NAME_TAG} ${DOCKERFILE_PATH}"
docker build --force-rm -f ${DOCKERFILE} -t ${IMAGE_NAME_TAG} ${DOCKERFILE_PATH}

rm -rf ${DOCKERFILE_PATH}/${serverName}.tar.gz
