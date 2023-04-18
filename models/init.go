package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Mongo = InitMongo()

var DBHost = "192.168.80.129"
var DBPort = "27017"
//var DBURI = fmt.Sprintf("mongodb://%s:%d", DBHost, DBPort)   // can cat different type, but performance low
var DBURI = fmt.Sprintf("mongodb://"+DBHost+":"+DBPort)

// encapsulated mongo query
func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1.local without password setting
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	fmt.Println("[C] DBURI: ",DBURI)

	// 2.dev server with password
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "admin",
		//PasswordSet: false,
	}).ApplyURI(DBURI))
	//}).ApplyURI("mongodb://192.168.80.129:27017"))
	//}).ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Println("Connection MongoDB Error: ", err)
		return nil
	}

	// if normal, return *mongo.Database
	return client.Database("im")
}
