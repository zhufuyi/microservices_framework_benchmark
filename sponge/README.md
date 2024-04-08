
### 创建http+grpc服务

准备好greeter.proto文件

1. 创建服务

> sponge micro grpc-http-pb --module-name=helloworld --server-name=helloworld --project-name=helloworld --protobuf-file=./greeter.proto

2. 打开代码文件 internal/service/greeter.go， 把去掉代码 panic("implement me")，添加代码：

```go
return &helloworldV1.HelloReply{Message: req.Name}, nil
```

3. 启动服务

```bash
# 生成代码
make proto

# 运行服务
make run
```

4. 测试api

- 在浏览器访问 http://localhost:8080/apis/swagger/index.html ，测试http api
- 使用goland或vs code 打开代码 helloworld/internal/service/greeter._client_test.go，填写参数测试grpc api

<br>

---

### 启动和停止服务

切换到目录 kratos/helloworld

```bash
# 后台启动http服务，日志输出到文件 helloworld.log
make run-nohup

# 停止服务
make run-http-nohup CMD=stop
```
