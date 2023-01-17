/*
 Navicat Premium Data Transfer

 Source Server         : CentOSdocker-mongo
 Source Server Type    : MongoDB
 Source Server Version : 50005
 Source Host           : 192.168.170.133:27017
 Source Schema         : im

 Target Server Type    : MongoDB
 Target Server Version : 50005
 File Encoding         : 65001

 Date: 17/01/2023 09:23:30
*/


// ----------------------------
// Collection structure for message_basic
// ----------------------------
db.getCollection("message_basic").drop();
db.createCollection("message_basic");

// ----------------------------
// Documents of message_basic
// ----------------------------
db.getCollection("message_basic").insert([ {
    _id: ObjectId("63b62ed5276d000033007a05"),
    "user_id": "用户唯一标识",
    "room_id": "房间唯一标识",
    data: "发送的数据",
    "created_at": 1,
    "updated_at": 1
} ]);

// ----------------------------
// Collection structure for room_basic
// ----------------------------
db.getCollection("room_basic").drop();
db.createCollection("room_basic");

// ----------------------------
// Documents of room_basic
// ----------------------------
db.getCollection("room_basic").insert([ {
    _id: ObjectId("63b63073276d000033007a06"),
    number: "房间号",
    name: "房间名称",
    info: "房间简介",
    "user_id": "房间创建者的唯一标识",
    "created_at": 1,
    "updated_at": 1
} ]);

// ----------------------------
// Collection structure for user_basic
// ----------------------------
db.getCollection("user_basic").drop();
db.createCollection("user_basic");

// ----------------------------
// Documents of user_basic
// ----------------------------
db.getCollection("user_basic").insert([ {
    _id: ObjectId("63b62cfd276d000033007a04"),
    account: "admin",
    password: "25d55ad283aa400af464c76d713c07ad",
    nickname: "Admin",
    gender: 1,
    email: "admin@admin",
    avatar: "头像",
    "created_at": 1672887500,
    "updated_at": 1672887601
} ]);

// ----------------------------
// Collection structure for user_room
// ----------------------------
db.getCollection("user_room").drop();
db.createCollection("user_room");

// ----------------------------
// Documents of user_room
// ----------------------------
db.getCollection("user_room").insert([ {
    _id: ObjectId("63b63196276d000033007a07"),
    "user_id": "用户唯一标识",
    "room_id": "房间唯一标识",
    "message_id": "消息唯一标识",
    "created_at": 1,
    "updated_at": 1
} ]);
