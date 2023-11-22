# WAF

- 基于gin框架做反向代理实现waf
- 项目还在开发中

## 项目背景介绍

### 需求产生

在实习做业务开发的时候经常会花费一些时间写一些安全中间件，现在临近大学毕业，正好选题也在这个方向，尝试自己实现一个waf用于后续业务开发，顺便完成毕业设计。预计功能参考：[https://github.com/bukaleyang/zhongkui-waf](https://github.com/bukaleyang/zhongkui-waf)

### 目前功能列表：

1. 支持IP黑白名单功能，白名单访问绕过中间件，黑名单的IP访问返回403。
2. 支持CC攻击防护，实现了4种限速算法，超过限速阀值返回403。
3. 未完待续...

### 部署：

1. 克隆代码

```
git clone https://github.com/supercat0867/waf
```

2. 进入根目录，安装依赖，编译

```
go mod download
go build
```

3. 编辑配置信息\
   配置文件为config目录下的config.yaml，可以根据需求更改，用途参考变量名

4. 运行

```
./waf
```