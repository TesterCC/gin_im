package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBasic struct {
	Id        string `bson:"_id"` // mongodb uuid
	Account   string `bson:"account"`
	Password  string `bson:"password"`
	Nickname  string `bson:"nickname"`
	Gender    int    `bson:"gender"` // 0-unknown, 1-male, 2-female
	Email     string `bson:"email"`
	Avatar    string `bson:"avatar"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

// 定义函数，返回结构体对应的集合名称，define a func return struct UserBasic collection name
func (UserBasic) CollectionName() string {
	return "user_basic"
}

func GerUserBasicByAccountPassword(account, password string) (*UserBasic, error) {
	ub := new(UserBasic)
	// 查询结果映射到user basic下
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"account", account}, {"password", password}}).Decode(ub)
	return ub, err
}

//func GetUserBasicById(id string) (*UserBasic, error) {
//	ub := new(UserBasic)
//	// 查询结果映射到user basic下
//	err := Mongo.Collection(UserBasic{}.CollectionName()).
//		FindOne(context.Background(), bson.D{{"_id", id}}).Decode(ub)
//	return ub, err
//}

func GetUserBasicById(_id primitive.ObjectID) (*UserBasic, error) {
	ub := new(UserBasic)
	// 查询结果映射到user basic下
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"_id", _id}}).Decode(ub)
	return ub, err
}


func GetUserBasicByEmail(email string) (*UserBasic, error) {
	ub := new(UserBasic)
	// 查询结果映射到user basic下
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"email", email}}).Decode(ub)
	return ub, err
}