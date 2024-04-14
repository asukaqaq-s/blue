# fim_server

即时通信项目后端工程

## 项目介绍

fim项目（枫枫知道即时通讯项目Fengfeng Instant Messaging）已完成后端部分课程录制

该项目选题新颖，技术栈相对丰富，是面试或go语言开发进阶的不二选择

该项目有8大服务

用户服务、对话服务、群聊服务、认证服务、系统服务、文件服务、日志服务、网关组成



相关功能包括但不限于：

- 好友搜索，好友申请，好友验证功能
- 安全的文件服务
- 用户实时聊天功能，支持文字、表情、图片、文件，支持好友上线提醒
- 高级用户对话功能，如撤回消息，回复消息，引用消息，用户视频、语音通话功能
- 实时群聊功能，支持用户建群、邀请用户入群、管理员、群禁言、群搜索
- 用户会话列表，好友置顶，群会话列表，群置顶
- 统一日志服务
- 管理员视角，直观了解系统用户使用情况



相关技术栈（后端）：

go-zero，gorm,  redis，mysql，kafka， docker-compose



在这个项目中你将会学到

1. 系统架构设计，表结构相关设计
2. 全链路整合，将会学习到中大型项目的开发流程及开发思路
3. 独立开发微服务项目以及独立排查错误的能力
4. 一款不输于市面上的即时通讯项目

## 系统架构图

![](https://image.fengfengzhidao.com/pic/20240329002540.png)


## 项目运行

在项目根目录下载依赖

```bash
go mod tidy
```

然后先运行rpc服务

先修改每个服务的配置文件，主要是修改你的etcd、mysql、redis的地址

运行服务需要保证根目录是各个服务的根目录

例如运行user的rpc服务，需要cd到 fim_server/user_rpc
```bash
go run userrpc.go
```

logs服务需要kafka的依赖

本地运行kafka
```bash

docker network create app-tier --driver bridge


docker run -d --restart=always --name zookeeper-server --network app-tier -e ALLOW_ANONYMOUS_LOGIN=yes bitnami/zookeeper:latest


docker run -d --restart=always  --name kafka-server  --network app-tier  -p 9092:9092 -e ALLOW_PLAINTEXT_LISTENER=yes  -e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper-server:2181 -e KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://192.168.0.108:9092  bitnami/kafka:latest


docker run -d --restart=always  --name kafka-map --network app-tier  -p 9001:8080 -v /opt/kafka-map/data:/usr/local/kafka-map/data  -e DEFAULT_USERNAME=admin  -e DEFAULT_PASSWORD=admin --restart always dushixiang/kafka-map:latest
# 用户名密码都是 admin
```

## 项目部署

参考课程部署视频及文档