
### 创建http+grpc服务

创建http+grpc服务参考 https://go-kratos.dev/docs/getting-started/start

1. 创建服务

> kratos new helloworld

2. 打开 helloworld/configs/config.yaml 配置文件，修改http服务端口为8080，修改grpc服务端口为8282

3. 启动服务

```bash
go mod tidy

cd cmd\helloworld\
go run ./...
```

4. 测试api

- curl http://127.0.0.1:8080/helloworld/foobar
- 使用grpc客户端测试api
  - port:  8282
  - path: /helloworld.v1.Greeter/SayHello
  - massage: {name: foobar}
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
