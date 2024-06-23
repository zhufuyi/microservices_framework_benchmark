## kratos 创建的服务的压测结果

kratos 版本 2.7.2

<br>

### http 压测结果

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
     data_received..................: 137 MB  2.3 MB/s
     data_sent......................: 115 MB  2.0 MB/s
     http_req_blocked...............: avg=1.68µs  min=0s med=0s     max=10.12ms p(90)=0s     p(95)=0s
     http_req_connecting............: avg=262ns   min=0s med=0s     max=10.12ms p(90)=0s     p(95)=0s
     http_req_duration..............: avg=2.86ms  min=0s med=2.01ms max=47.44ms p(90)=6.62ms p(95)=8.68ms
       { expected_response:true }...: avg=2.86ms  min=0s med=2.01ms max=47.44ms p(90)=6.62ms p(95)=8.68ms
     http_req_failed................: 0.00%   ✓ 0            ✗ 1000000
     http_req_receiving.............: avg=20.56µs min=0s med=0s     max=4.68ms  p(90)=0s     p(95)=0s
     http_req_sending...............: avg=8.01µs  min=0s med=0s     max=3.77ms  p(90)=0s     p(95)=0s
     http_req_tls_handshaking.......: avg=0s      min=0s med=0s     max=0s      p(90)=0s     p(95)=0s
     http_req_waiting...............: avg=2.83ms  min=0s med=2ms    max=47.44ms p(90)=6.58ms p(95)=8.64ms
     http_reqs......................: 1000000 17085.130811/s
     iteration_duration.............: avg=2.91ms  min=0s med=2.02ms max=47.44ms p(90)=6.68ms p(95)=8.76ms
     iterations.....................: 1000000 17085.130811/s
     vus............................: 50      min=50         max=50
     vus_max........................: 50      min=50         max=50


running (00m58.5s), 00/50 VUs, 1000000 complete and 0 interrupted iterations
default ✓ [======================================] 50 VUs  00m58.5s/10m0s  1000000/1000000 shared iters
```

<br>

压测http api指标的grafana界面：

![kratos-http-k6](../test/assets/kratos/kratos-http-k6.png)

<br>

采集到的服务程序指标的grafana界面：

![kratos-http-cpu](../test/assets/kratos/kratos-http-cpu.png)

![kratos-http-go-process](../test/assets/kratos/kratos-http-go-process.png)

<br>

### grpc 压测结果

使用压测工具ghz，50个并发，总共100万次请求， 压测结果如下：

![kratos-grpc-ghz](../test/assets/kratos/kratos-grpc-ghz.png)

<br>

采集到的服务程序指标的grafana界面：

![kratos-grpc-cpu](../test/assets/kratos/kratos-grpc-cpu.png)

![kratos-grpc-go-process](../test/assets/kratos/kratos-grpc-go-process.png)

<br>
