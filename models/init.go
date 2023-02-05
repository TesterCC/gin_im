package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Mongo = InitMongo()

// encapsulated mongo query
func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 1 - local without password setting
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	//// 2 - dev server with password
	//client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
	//	Username: "admin",
	//	Password: "admin",
	//	//PasswordSet: false,
	//}).ApplyURI("mongodb://192.168.170.137:27017"))

	if err != nil {
		log.Println("Connection MongoDB Error: ", err)
		return nil
	}

	// if normal, return *mongo.Database
	return client.Database("im")
}
