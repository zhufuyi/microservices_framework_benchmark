## monitor

启动服务前先修改prometheus配置文件prometheus.yml访问权限，方便修改配置后prometheus可以热重载。

```bash
chmod 0777 prometheus.yml
```

<br>

### 无持久化数据启动服务

打开`docker-compose.yml`，把下面两行注释掉：

- $PWD/prometheus_data/data/:/prometheus/data/
- $PWD/grafana_data/grafana/:/var/lib/grafana/

然后启动服务：

```bash
docker-compose up -d
```

<br>

### 持久化数据启动服务

在无持久化运行服务的基础上， 把运行中的容器文件复制到本地：

```bash
# 复制prometheus数据到本地
docker cp prometheus:/prometheus/data/ prometheus_data/

# 复制grafana数据到本地
docker cp grafana:/var/lib/grafana/ grafana_data/
```

修改数据访问权限。

```bash
chomod -R 0777  prometheus_data/data
chomod -R 0777  grafana_data/grafana
```

打开docker-compose.yml，取消之前注释的映射，重启服务：

```bash
docker-compose down
docker-compose up -d
```

<br>

### 导入grafana面板

#### 获取数据源的uid

因为 [grafana-dashboards](grafana-dashboards) 的目录下grafana面板依赖于Prometheus数据源uid，在导入到grafana中要修改这个数据源的uid，否则会报数据源不存在的错误。

打开grafana添加prometheus数据源：

1. 打开浏览器访问http://localhost:3000。
2. 点击左边菜单栏的Dashboards，点击右上角的+号，选择Import，选择导入的面板文件。
3. 选择导入的面板文件，点击右上角的导入按钮。
4. 选择数据源，选择Prometheus，输入Prometheus的地址 http://<ip_address>:9090，点击Save & Test按钮。

获取uid方式：

1. 在 https://grafana.com/grafana/dashboards/ 随便找一个prometheus类型的面板复制id出来(例如6671)。
2. 在grafana导入并进入面板，在面板的右上角，找到并点击齿轮图标(设置)，点击左边菜单栏的JSON Model查看json文件，搜索"uid"就可以看到数据源的uid值。
3. 复制uid值，修改 [grafana-dashboards](grafana-dashboards) 目录下对应的面板文件，把所有json文件的uid值(d408553f-a33a-474f-ba16-912057230c24)改成刚才复制的uid值。
4. 把 [grafana-dashboards](grafana-dashboards) 目录下面的所有面板文件导入到grafana中。

<br>
