# WAF

- 基于gin框架做反向代理实现waf
- 项目还在开发中

## 项目背景介绍

### 需求产生

在实习做业务开发的时候经常会花费一些时间写一些安全中间件，现在临近大学毕业，正好选题也在这个方向，尝试自己实现一个waf用于后续业务开发，顺便完成毕业设计。预计功能参考：[https://github.com/bukaleyang/zhongkui-waf](https://github.com/bukaleyang/zhongkui-waf)

### 功能列表：

1. 支持动态IP黑白名单功能，可配置过期时间，白名单访问绕过中间件，黑名单的IP访问返回403。
2. 支持CC防护，实现了4种限速算法，超过限速阀值返回403,计数器超过配置值，IP自动拉入永久黑名单。
3. 未完待续...

### 部署：

1. 克隆代码

```
git clone https://github.com/supercat0867/waf
```

2. 进入根目录，安装依赖，编译。

```
go mod download
go build
```

3. 编辑配置信息\
   配置文件为config目录下的config.yaml，可以根据需求更改，用途参考变量名。
```
server: "http://localhost"

port: 8080

proxyServer: "http://localhost:80"

jwtSetting:
  secretKey: "www.supercat.cc"

redis:
  address: "localhost:6379"
  password: ""
  database: 0

rateLimiterMode: 1

rateLimiter:
  maxCounter: 3
  tokenBucket:
    maxToken: 3
    tokenPerSecond: 1
  leakyBucket:
    capacity: 30
    leakyPerSecond: 10
  fixedWindow:
    windowSize: 1
    maxRequests: 20
  slideWindow:
    windowSize: 1
    maxRequests: 10
```

4. 生成API文档，在根目录输入
```azure
swag init
```
5. 运行

```
./waf
```

6. API文档路由\
   http://127.0.0.1:8080/swagger/index.html