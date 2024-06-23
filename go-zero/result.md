## go-zero 创建的服务的压测结果

go-zero 版本 1.6.3

<br>

#### http 压测结果

使用压测工具k6，50个并发，总共100万次请求的结果：

```bash
$ K6_PROMETHEUS_RW_SERVER_URL="http://192.168.3.37:9090/api/v1/write" K6_PROMETHEUS_RW_TREND_STATS="min,max,avg,p(95),p(99)" K6_PROMETHEUS_RW_PUSH_INTERVAL=1s k6 run -u 50 -i 1000000 -o experimental-prometheus-rw http-load-test.js

  execution: local
     script: http-load-test.js
     output: Prometheus remote write (http://192.168.3.37:9090/api/v1/write)

  scenarios: (100.00%) 1 scenario, 50 max VUs, 10m30s max duration (incl. graceful stop):
           * default: 1000000 iterations shared among 50 VUs (maxDuration: 10m0s, gracefulStop: 30s)

     ✓ status is 200

     checks.........................: 100.00% ✓ 1000000      ✗ 0
     data_received..................: 222 MB  3.5 MB/s
     data_sent......................: 115 MB  1.8 MB/s
     http_req_blocked...............: avg=1.67µs  min=0s med=0s     max=8.64ms   p(90)=0s     p(95)=0s
     http_req_connecting............: avg=243ns   min=0s med=0s     max=8.64ms   p(90)=0s     p(95)=0s
     http_req_duration..............: avg=3.08ms  min=0s med=2.05ms max=124.92ms p(90)=6.78ms p(95)=9.13ms
       { expected_response:true }...: avg=3.08ms  min=0s med=2.05ms max=124.92ms p(90)=6.78ms p(95)=9.13ms
     http_req_failed................: 0.00%   ✓ 0            ✗ 1000000
     http_req_receiving.............: avg=21.31µs min=0s med=0s     max=5.89ms   p(90)=0s     p(95)=0s
     http_req_sending...............: avg=8.13µs  min=0s med=0s     max=5.2ms    p(90)=0s     p(95)=0s
     http_req_tls_handshaking.......: avg=0s      min=0s med=0s     max=0s       p(90)=0s     p(95)=0s
     http_req_waiting...............: avg=3.05ms  min=0s med=2.04ms max=124.92ms p(90)=6.75ms p(95)=9.09ms
     http_reqs......................: 1000000 15887.422209/s
     iteration_duration.............: avg=3.13ms  min=0s med=2.08ms max=124.92ms p(90)=6.85ms p(95)=9.21ms
     iterations.....................: 1000000 15887.422209/s
     vus............................: 50      min=50         max=50
     vus_max........................: 50      min=50         max=50


running (01m02.9s), 00/50 VUs, 1000000 complete and 0 interrupted iterations
default ✓ [======================================] 50 VUs  01m02.9s/10m0s  1000000/1000000 shared iters
```

<br>

压测http api指标的grafana界面：

![go-zero-http-k6](../test/assets/go-zero/go-zero-http-k6.png)

<br>

采集到的服务程序指标的grafana界面：

![go-zero-http-cpu](../test/assets/go-zero/go-zero-http-cpu.png)

![go-zero-http-go-process](../test/assets/go-zero/go-zero-http-go-process.png)

<br>

### grpc 压测结果数

使用压测工具ghz，50个并发，总共100万次，压测结果如下：

![go-zero-grpc-ghz](../test/assets/go-zero/go-zero-grpc-ghz.png)

<br>

采集到的服务程序指标的grafana界面：

![go-zero-grpc-cpu](../test/assets/go-zero/go-zero-grpc-cpu.png)

![go-zero-grpc-go-process](../test/assets/go-zero/go-zero-grpc-go-process.png)

<br>
