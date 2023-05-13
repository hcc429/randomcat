package db

import (
	"context"
	"log"
	"os"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


var (
	Client *mongo.Client
	DB_URL = os.Getenv("DATABASE_URL")
	DB_NAME = os.Getenv("DATABASE_NAME")
)
func InitClient(){
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(DB_URL))
	if err != nil{
		log.Fatal("Error when connect to mongodb\n")
		panic(err)
	}
	err = Client.Ping(ctx, readpref.Primary())
	if err != nil{
		log.Fatal("Couldn't connect to database!\n", err)
	} else{
		log.Println("Connected!")
	}
}

func GetDB() *mongo.Database{
	return Client.Database(DB_NAME)
}