#!/bin/bash

serverName="grpc_helloworld"
binaryDir="grpc/helloworld"

function checkResult() {
    result=$1
    if [ ${result} -ne 0 ]; then
        exit ${result}
    fi
}

cd ${binaryDir}
go build -o ${serverName} greeter.go
checkResult $?
cd - > /dev/null

# running server
./${binaryDir}/${serverName} -f grpc/helloworld/etc/greeter.yaml
checkResult $?
