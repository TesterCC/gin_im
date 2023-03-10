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
go get -u github.com/golang-jwt/jwt/v4    # token登录需要
go get -u github.com/jordan-wright/email
go get -u github.com/gin-contrib/cors     # Gin解决跨域问题
go get -u github.com/gin-contrib/sessions # Gin解决cookie session的方案session   
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

## import data

```
use navicat create mongodb - im db,
active im db, then right click "Excute Script File"
select im.sql and run it
```

## 数据库设计

见[MongoDB Collection Design](models/model.md)

## IM结构

- 用户模块：密码登录 Fin 、邮箱注册、用户详情 Fin
- 通讯模块（核心）：一对一通讯、多对多通讯、消息列表、聊天记录列表

## 进度

- [【项目实战】基于Gin、MongoDB的即时通讯系统](https://www.bilibili.com/video/BV1YL4y1c7ZX/?p=9)  0：38

注意启动时要使用root权限

## REF

- [Gitee项目地址](https://gitee.com/getcharzp/im)
- [Github项目地址](https://github.com/getcharzp/im)