# 蔓茵科技



> create by P-F on 2022/10/21



## 1. 项目介绍

本项目是一个大创结项作品，基本实现了一个商城，订单业务也基本完善了。



## 2. 技术栈

**后端**

- go 语言

- beego 框架

- jwt 登录 鉴权 

- Redis 缓存 下单

- MySQL 存储数据

**前端**

- Vue

- Element-ui

- axios

**部署**

- Nginx  反向代理，静态资源压缩



## 3. 项目启动



### 3.1 前端

安装依赖包：

```shell
npm install
```

修改配置信息：public/env.json。model 指应用哪套配置，每套配置中 server 是后端服务地址，用于请求后台。host 是前端所在的域名，用于生成付款二维码。

运行项目：

```shell
npm run serve
```



### 3.2 后端启动

进入 go-service 目录，执行命令：

```shell
go mod tidy
```

将依赖包封装进项目内：

```shell
go mod vendor
```

修改配置文件：conf/app.conf。主要修改 MySQL 的地址和 Redis 的地址和密码。

导入数据库：进入数据库创建 manyin 数据库，然后将 sql 文件导入：`mysql -u xxx -p manyin < manyin.sql`.

启动项目：

```shell
go ./main.go
```



### 3.3 配置 Nginx

这一步不是必须的，如果没配置 Nginx ，在前端里面的 env.json 里后端地址就得加上端口 9090，同时前端访问的时候得访问 localhost:8080.

将 nginx 目录中的 manyin-dev.conf 复制到 Nginx 目录的conf中。配置 Nginx 的 nginx.conf，加上一行`include ./manyin-dev.conf;` 即可。



## 4. 项目部署

- Vue 打包 dist 给 Nginx 部署

- 配置 nginx 目录中的manyin-prod.conf 到 nginx.conf

- 打包 go-service 并后台启动


