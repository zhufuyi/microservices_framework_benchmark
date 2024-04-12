## 性能测试

执行命令测试http api，

```bash
# 50个虚拟用户，总共100万个请求
bash test.sh http 50 1000000

# 50个虚拟用户，总共100万个请求，把压测指标推送到prometheus
#bash test.sh http 50 1000000 http://192.168.3.37:9090

```

<br>

执行命令测试grpc api：

```bash
# 50个虚拟用户，总共100万个请求
bash test.sh grpc 50 1000

# 50个虚拟用户，总共100万个请求，把压测指标推送到prometheus
#bash test.sh grpc 50 1000  http://192.168.3.37:9090
```
