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

<br>

### 性能测试

测试流程：

- 在服务器A运行kratos创建的http+grpc服务(1个程序文件)，然后在服务器B分别压测http和grpc的api，压测完成后停止服务。
- 在服务器A运行go-zero创建的http和grpc服务(2个程序文件，不要同时运行http和grpc服务，因为端口8080冲突)，然后在服务器B分别压测http和grpc的api，压测完成后停止服务。
- 在服务器A运行sponge创建的http+grpc服务(1个程序文件)，然后在服务器B分别压测http和grpc的api，压测完成后停止服务。

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

> 点击查看 [test/README.md](test/README.md) 文件里的 http api 压测说明。

3. 在服务器B执行命令压测grpc api

> 点击查看 [test/README.md](test/README.md) 文件里的 grpc api 压测说明。

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

> 点击查看 [test/README.md](test/README.md) 文件里的 http api 压测说明。

3. 在服务器A运行grpc服务

切换到目录 go-zero，在后台运行grpc服务：

```bash
make run-nohup-grpc

# 测试完毕后在服务器A执行命令停止grpc服务
# make run-nohup-grpc CMD=stop
```

4. 在服务器B执行命令压测grpc api

> 点击查看 [test/README.md](test/README.md) 文件里的 grpc api 压测说明。

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

> 点击查看 [test/README.md](test/README.md) 文件里的 http api 压测说明。

3. 在服务器B执行命令压测grpc api

> 点击查看 [test/README.md](test/README.md) 文件里的 grpc api 压测说明。
