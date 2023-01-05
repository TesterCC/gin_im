## Gin + MongoDB 即时通讯系统

## 技术栈

- 语言：Golang
- 数据库：MongoDB
- 协议：Websocket

## 核心包安装

```
go get -u github.com/gin-gonic/gin
go get -u go.mongodb.org/mongo-driver/mongo
go get -u github.com/gorilla/websocket
go get -u github.com/golang-jwt/jwt/v4
```

## Docker install MongoDB

```shell
docker pull mongo
# docker pull mongo-express  # web console

docker run -d --name some-mongo \
-e MONGO_INITDB_ROOT_USERNAME=admin \
-e MONGO_INITDB_ROOT_PASSWORD=admin \
-p 27017:27017 \
mongo
```

## 数据库设计

见[MongoDB Collection Design](models/model.md)

## IM结构

- 用户模块：密码登录、邮箱注册、用户详情
- 通讯模块（核心）：一对一通讯、多对多通讯、消息列表、聊天记录列表

## 进度

- [【项目实战】基于Gin、MongoDB的即时通讯系统](https://www.bilibili.com/video/BV1YL4y1c7ZX/?p=3)