#!/bin/bash

HOST_ADDR=$1

function checkResult() {
    result=$1
    if [ ${result} -ne 0 ]; then
        exit ${result}
    fi
}

# change host addr
if [ "X${HOST_ADDR}" = "X" ];then
  HOST_ADDR=$(cat cmd/helloworld/main.go | grep "@host" | awk '{print $3}')
  HOST_ADDR=$(echo  ${HOST_ADDR} | cut -d ':' -f 1)
else
    sed -i "s/@host .*:8080/@host ${HOST_ADDR}:8080/g" cmd/helloworld/main.go
fi

# generate api docs
swag init -g cmd/helloworld/main.go
checkResult $?

# modify duplicate numbers and error codes
sponge patch modify-dup-num --dir=internal/ecode
sponge patch modify-dup-err-code --dir=internal/ecode

colorCyan='\e[1;36m'
highBright='\e[1m'
markEnd='\e[0m'

echo ""
echo -e "${highBright}Tip:${markEnd} execute the command ${colorCyan}make run${markEnd} and then visit ${colorCyan}http://${HOST_ADDR}:8080/swagger/index.html${markEnd} in your browser."
echo ""
echo "generated api docs successfully."
echo ""
