#!/bin/bash

serverName="helloworld"
binaryFile="cmd/${serverName}/${serverName}"

function checkResult() {
    result=$1
    if [ ${result} -ne 0 ]; then
        exit ${result}
    fi
}

cd cmd/${serverName}
go build -o ${serverName} ./...
checkResult $?
cd - > /dev/null

# running server
./${binaryFile} -conf configs/config.yaml
checkResult $?
