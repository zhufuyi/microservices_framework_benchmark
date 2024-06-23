## 性能测试

### 执行命令测试http api

测试前打开 [test/http-load-test.js](http-load-test.js) 文件，修改`baseUrl`地址(被测web服务地址)，然后执行命令测试：

```bash
# 50个虚拟用户，总共100万个请求
bash test.sh http 50 1000000

# 50个虚拟用户，总共100万个请求，把压测指标推送到prometheus
#bash test.sh http 50 1000000 http://192.168.3.37:9090
```

<br>

### 执行命令测试grpc api

#### 方式一：使用工具k6

测试前打开 [test/grpc-load-test.js](grpc-load-test.js) 文件，修改`grpcAddr`地址(被测grpc服务地址)，然后执行命令：

```bash
# 50个虚拟用户，总共100万个请求
bash test.sh grpc 50 1000

# 50个虚拟用户，总共100万个请求，把压测指标推送到prometheus
#bash test.sh grpc 50 1000  http://192.168.3.37:9090
```

<br>

#### 方式二：使用工具ghz

测试前打开配置文件 [sponge/helloworld/configs/helloworld.yml](../sponge/helloworld/configs/helloworld.yml)，修改 `grpcClient` 下的 **host** 和 **port** 值(被测grpc服务地址)

然后在终端切换到目录 `sponge/helloworld/internal/service`，执行命令：

```bash
# 50个并发，总共100万次请求
go test -run Test_service_greeter_benchmark/SayHello
```

<br>

## 监控

监控使用Prometheus，把测试结果推送到Prometheus，然后通过Grafana可视化展示。搭建监控环境请参考 [test/monitor/grafana-prometheus/README.md](monitor/grafana-prometheus/README.md)

为了监控进程指标，可以使用process-exporter，请参考 [test/monitor/process-exporter/README.md](monitor/process-exporter/README.md)
