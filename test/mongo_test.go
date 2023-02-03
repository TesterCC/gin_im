package test

import (
	"context"
	"fmt"
	"gin_im/models"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestFindOne(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// local without password setting
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	// dev server with password
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username:    "admin",
		Password:    "admin",
		PasswordSet: false,
	}).ApplyURI("mongodb://192.168.170.137:27017"))

	if err != nil {
		t.Fatal(err)
	}

	db := client.Database("im")
	ub := new(models.UserBasic)
	// 将查询到的数据解析到ub上
	db.Collection("user_basic").FindOne(context.Background(), bson.D{}).Decode(ub)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("ub --> ", ub)
	fmt.Println("avatar --> ", ub.Avatar)
	fmt.Println("gender --> ", ub.Gender)
}
