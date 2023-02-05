# 苏宁易购
领先的综合网上购物商城，正品行货，全国联保，让您尽享购物乐趣
## 功能
实现了以下功能

 注册、登录、修改密码、注销登录

 设置个人信息

 充值、查看钱包余额
 
 查看、添加、删除收货地址

 收藏商品、取消收藏

 查看订单、取消订单、确认订单

 查看所有商品
 
 购买商品

 添加、删除、查看、支付购物车内商品

 评论、查看商品评论区
## 项目框架

采用gin框架进行项目构建

### 什么是Gin框架

Go世界里最流行的Web框架，Github上有32K+star。 基于httprouter开发的Web框架。 中文文档齐全，简 单易用的轻量级框架。Gin是一个用Go语言编写的web框架。它是一个类似于martini但拥有更好性能的API框架 , 由于使用了httprouter，速度提高了近40倍。 如果你是性能和高效的追求者, 你会爱上Gin。

### 框架结构

README.md：项目的说明文档

api：接口层，在里面是详细的逻辑实现以及路由。

dao：全名为 data access object，操作数据库。

service:调用dao层的一些函数从而实现api层的一些功能

model：模型层，主要放数据库实例的结构体。

utils：一些常用的工具函数，封装在这里减少代码的重复使用。

middleware:存放一些中间件

cmd:项目的入口，存放main函数

go.mod：依赖管理

## Mysql——数据库

### MySQL介绍

该项目运用MySQL作为数据库，进行数据储存和数据反馈

MySQL是一个关系型数据库管理系统，由瑞典MySQL AB 公司开发，属于 Oracle 旗下产品。MySQL 是最流行的关系型数据库管理系统关系型数据库管理系统之一，在 WEB 应用方面，MySQL是最好的 RDBMS (Relational Database Management System，关系数据库管理系统) 应用软件之一

## 中间件

运用了cookie和cors

### cookie

cookie 存储在客户端： cookie 是服务器发送到用户浏览器并保存在本地的一小块数据，它会在浏览器下次向同一服务器再发起请求时被携带并发送到服务器上。因此，服务端脚本就可以读、写存储在客户端的cookie的值。

cookie用于获取登录状态，确保用户在进行其他操作的时候，必须处于登录状态。避免资源的浪费。



### cors

CORS 是一种基于 [HTTP Header](https://link.juejin.cn?target=https%3A%2F%2Fdeveloper.mozilla.org%2Fen-US%2Fdocs%2FGlossary%2FHeader) 的机制，该机制通过允许服务器标示除了它自己以外的其它域。服务器端配合浏览器实现 `CORS` 机制，可以突破浏览器对跨域资源访问的限制，实现跨域资源请求。

采用cors中间件，解决跨域问题

### 不足

未考虑并发处理，仅接受单线程访问

防止sql注入方面并未完善
