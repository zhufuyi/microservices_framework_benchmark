## kratos、go-zero、sponge三个微服务框架创建的http和gprc服务的性能测试

#### 主要压测指标

- **吞吐量(Throughput)**：单位时间内处理的请求数量，通常以每秒请求数(Requests per Second，RPS)表示。
- **响应时间(Response Time)**：从发出请求到收到响应的时间，包括p95、p99、avg、min、max。
- **错误率(Error Rate)**：请求处理失败或产生错误的比率。
- **资源利用率(Resource Utilization)**：包括cpu、内存、网络带宽等资源的使用情况。

<br>

#### 压测api

- **http**
    - 端口: 8080
    - 路由: /api/v1/helloworld/qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnm
    - 请求方法: GET
- **grpc**
    - 端口: 8282
    - path: /helloworld.v1.Greeter/SayHello
    - message: qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnm
    - 类型: unary

注： 8080端口的路由/metrics是采集go程序的指标。

<br>

#### 压测工具

- [**k6**](https://github.com/grafana/k6): 用来压测http服务。
- [**ghz**](https://github.com/bojand/ghz): 用来压测grpc服务。

<br>

#### 压测环境

因为不同的服务器硬件对性能测试结果不一样，本次是在宿主机和虚拟机之间进行负载测试：

- 宿主机
  - 硬件：R7 6800H CPU，16G内存
  - 用途：运行工具k6和ghz测试http api和grpc api
- VMware虚拟机
  - 系统：centos 8
  - 硬件：8核cpu、4G内存
  - 用途：用于单独运行kratos、go-zero、sponge创建的http和grpc服务

> 如果想要在自己的机器上进行负载测试，点击查看[压测说明文档](benchmark.md)。

<br>

#### 压测结果

50个并发，总共100万个请求，压测kratos、go-zero、sponge创建的`http`服务结果：

![http-server](test/assets/http-server.png)

<br>

50个并发，总共100万个请求，压测kratos、go-zero、sponge创建的`grpc`服务结果：

![grpc-server](test/assets/grpc-server.png)

<br>

- [**查看 kratos 压测的详细结果**](kratos/result.md)
- [**查看 go-zero 压测的详细结果**](go-zero/result.md)
- [**查看 sponge 压测详细结果**](sponge/result.md)

<br>

