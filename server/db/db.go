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
	client  *mongo.Client
	db_url  = os.Getenv("DATABASE_URL")
	db_name = os.Getenv("DATABASE_NAME")
)

func init() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(db_url))
	if err != nil {
		panic("Error when connect to mongodb\n" + err.Error())
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to database!\n", err)
	} else {
		log.Println("Connected to MongoDB!")
	}
}

func GetDB() *mongo.Database {
	return client.Database(db_name)
}
