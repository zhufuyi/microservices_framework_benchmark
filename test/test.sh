#!/bin/bash

testType=$1   # 第一个参数，测试类型，http或grpc
VUs=$2          # 第二个参数，并发数，默认为1
iterations=$3  # 第三个参数，循环次数，默认1次
prometheusAddr=$4  # 第四个参数，prometheus地址(例如http://192.168.3.37:9090)，如果为空，忽略把压测指标推送给prometheus

if [ -z "$VUs" ]; then
  VUs="1"
fi

if [ -z "$iterations" ]; then
  iterations="1"
fi

function httpLoadTest() {
  if [  -z "$prometheusAddr" ]; then
    echo -e "k6 run -u $VUs -i $iterations  http-load-test.js\n\n"
    k6 run -u $VUs -i $iterations  http-load-test.js
  else
    echo "K6_PROMETHEUS_RW_SERVER_URL=\"${prometheusAddr}/api/v1/write\" K6_PROMETHEUS_RW_TREND_STATS=\"min,max,avg,p(95),p(99)\" K6_PROMETHEUS_RW_PUSH_INTERVAL=1s k6 run -u $VUs -i $iterations -o experimental-prometheus-rw http-load-test.js"
    K6_PROMETHEUS_RW_SERVER_URL="${prometheusAddr}/api/v1/write" K6_PROMETHEUS_RW_TREND_STATS="min,max,avg,p(95),p(99)" K6_PROMETHEUS_RW_PUSH_INTERVAL=1s k6 run -u $VUs -i $iterations -o experimental-prometheus-rw http-load-test.js
  fi
}

function grpcLoadTest() {
  if [  -z "$prometheusAddr" ]; then
    echo -e "k6 run -u $VUs -i $iterations grpc-load-test.js\n\n"
    k6 run -u $VUs -i $iterations grpc-load-test.js
  else
    echo "K6_PROMETHEUS_RW_SERVER_URL=\"${prometheusAddr}/api/v1/write\" K6_PROMETHEUS_RW_TREND_STATS=\"min,max,avg,p(95),p(99)\" K6_PROMETHEUS_RW_PUSH_INTERVAL=1s k6 run -u $VUs -i $iterations -o experimental-prometheus-rw grpc-load-test.js"
    K6_PROMETHEUS_RW_SERVER_URL="${prometheusAddr}/api/v1/write" K6_PROMETHEUS_RW_TREND_STATS="min,max,avg,p(95),p(99)" K6_PROMETHEUS_RW_PUSH_INTERVAL=1s k6 run -u $VUs -i $iterations -o experimental-prometheus-rw grpc-load-test.js
  fi
}

if [ "$testType" == "http" ]; then
  httpLoadTest
elif [ "$testType" == "grpc" ]; then
  echo -e "因为k6工具压测grpc服务的本身的性能测试能力有限，建议使用ghz工具来压测获得更准确的结果。\n\n"
  grpcLoadTest
else
  echo "Usage: $0 <http|grpc> <VUs> <iterations>"
fi
