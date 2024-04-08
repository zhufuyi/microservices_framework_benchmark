#!/bin/bash

serverName="http_helloworld"
binaryDir="http/helloworld"

function checkResult() {
    result=$1
    if [ ${result} -ne 0 ]; then
        exit ${result}
    fi
}

cd ${binaryDir}
go build -o ${serverName} helloworld.go
checkResult $?
cd - > /dev/null

# running server
./${binaryDir}/${serverName} -f http/helloworld/etc/helloworld-api.yaml
checkResult $?
