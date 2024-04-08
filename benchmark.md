## 压测说明

准备两台服务器A和B，并在服务A和B上安装一些压测时需要的程序和工具。

### 环境搭建

#### 服务器A的环境搭建

服务器A用于单独运行kratos、go-zero、sponge创建的服务。

1. 安装go环境。

2. 下载代码 git clone https://github.com/zhufuyi/microservices_framework_benchmark.git

<br>

#### 服务器B的环境搭建

服务器B运行http和grpc客户端进行测试。

1. 安装go环境。

2. 安装负载压测工具k6，下载地址 https://github.com/grafana/k6/releases/tag/v0.50.0 ，把二进制文件移动到系统环境path下，如果想把压测指标实时推送给Prometheus，请按照[xk6-output-prometheus-remote](https://github.com/grafana/xk6-output-prometheus-remote)重新编译新的k6。

3. 下载代码  git clone https://github.com/zhufuyi/microservices_framework_benchmark.git

4. 配置服务器A的ip地址

- 切换到目录`test`，打开 http-load-test.js 脚本文件，把 `192.168.3.37` 改为服务器A的ip地址，用于测试http服务。
- 切换到目录`sponge/helloworld/configs`，打开 `helloworld.yml` 配置文件，把配置 `grpcClient` 下的host地址`192.168.3.37` 改为服务器A的ip地址，用于测试grpc服务。

<br>

### 性能压测

压测流程：

- 在服务器A启动kratos创建的服务(http+grpc)，然后在服务器B分别压测http和grpc的api，压测完成后停止服务。
- 在服务器A启动go-zero创建的http和grpc服务，然后在服务器B分别压测http和grpc的api，压测完成后停止服务。
- 在服务器A启动sponge创建的服务(http+grpc)，然后在服务器B分别压测http和grpc的api，压测完成后停止服务。

<br>

#### 压测 kratos 创建的服务

1. 在服务器A运行http+grpc服务，点击查看[kratos创建的服务说明](kratos/README.md)。

切换到目录`kratos/helloworld`，在后台运行服务：

```bash
make run-nohup

# 测试完毕后在服务器A执行命令停止http+grpc服务
# make run-nohup CMD=stop
```

2. 在服务器B执行命令压测http api

使用50个虚拟用户，100万次请求，执行命令压测http api：

```bash
bash test.sh http 50 1000000
```

3. 在服务器B执行命令压测grpc api

切换到目录`sponge/helloworld/internal/service`。

并发50个协程，100万次请求，执行命令压测grpc api：

```bash
go test -run Test_service_greeter_benchmark/SayHello
```

<br>

#### 压测 go-zero 创建的服务

因为go-zero不支持在一个服务上同时创建http和grpc协议，因此分别创建了http和grpc两个单独的服务，点击查看[goctl创建的服务说明](go-zero/README.md)。

1. 在服务器A运行http服务

切换到目录 go-zero，在后台运行http服务：

```bash
make run-nohup-http

# 测试完毕后在服务器A执行命令停止http服务
# make run-nohup-http CMD=stop
```

2. 在服务器B执行命令压测http api

使用50个虚拟用户，100万次请求，执行命令压测http api：

```bash
bash test.sh http 50 1000000
```

3. 在服务器A运行grpc服务

切换到目录 go-zero，在后台运行grpc服务：

```bash
make run-nohup-grpc

# 测试完毕后在服务器A执行命令停止grpc服务
# make run-nohup-grpc CMD=stop
```

4. 在服务器B执行命令压测grpc api

切换到目录`sponge/helloworld/internal/service`。

并发50个协程，100万次请求，执行命令压测grpc api：

```bash
go test -run Test_service_greeter_benchmark/SayHello
```

<br>

#### 压测 sponge 创建的服务

1. 在服务器A运行http+grpc服务，点击查看[sponge创建的服务说明](sponge/README.md)。

切换到目录`sponge/helloworld`，在后台运行服务：

```bash
make run-nohup

# 测试完毕后在服务器A执行命令停止http+grpc服务
# make run-nohup CMD=stop
```

2. 在服务器B执行命令压测http api

使用50个虚拟用户，100万次请求，执行命令压测http api：

```bash
bash test.sh http 50 1000000
```

3. 在服务器B执行命令压测grpc api

切换到目录`sponge/helloworld/internal/service`。

并发50个协程，100万次请求，执行命令压测grpc api：

```bash
go test -run Test_service_greeter_benchmark/SayHello
```
