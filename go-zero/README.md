
### 创建http服务

创建http服务参考 https://go-zero.dev/docs/tasks/cli/api-demo

1. 创建服务

> goctl api new helloworld

2. 打开代码文件 internal/logic/helloworldlogic.go， 添加下面两行代码

```go
resp = new(types.Response)
resp.Message = req.Name
```

3. 打开配置文件 etc/helloworld-api.yaml，修改端口为8080

4. 启动服务

```bash
go mod tidy

go run helloworld.go
```

5. 请求api

> curl http://127.0.0.1:8080/helloworld/foobar

<br>

### 创建grpc服务

创建grpc服务参考 https://go-zero.dev/docs/tasks/cli/grpc-demo

1. 创建服务

> goctl rpc new helloworld

把准备好的greeter.proto文件替换helloworld.proto，执行命令生成代码

> goctl rpc protoc helloworld.proto --go_out=. --go-grpc_out=. --zrpc_out=.

2. 打开代码文件 internal/logic/sayhellologic.go， 修改代码为：

```go
return &helloworldV1.HelloReply{Message: in.Name}, nil
```

3. 打开etc/greeter.yaml，修改端口为8282，并删除etcd配置

```yaml
Name: helloworld.rpc
ListenOn: 0.0.0.0:8282
#Etcd:
#  Hosts:
#  - 127.0.0.1:2379
#  Key: helloworld.rpc
```

4. 启动服务

```bash
go mod tidy

go run helloworld.go
```

5. 使用grpc客户端请求api

- port:  8282
- path: /helloworld.v1.Greeter/SayHello
- massage: {name: foobar}

<br>

---

### 启动和停止服务

注： 不要同时启动http和grpc服务，因为grpc服务也使用了8080端口(采集go metrics)，会与http服务的8080端口冲突。

#### 启动和停止http服务

```bash
# 后台启动http服务，日志输出到文件 http/helloworld/helloworld.log
make run-http-nohup

# 停止服务
make run-http-nohup CMD=stop
```

<br>

#### 启动和停止grpc服务

```bash
# 后台启动grpc服务，日志输出到文件 grpc/helloworld/helloworld.log
make run-grpc-nohup

# 停止服务
make run-grpc-nohup CMD=stop
```
