package db

import (
	"context"
	"fmt"
	"github.com/hcc429/randomcat/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)



func InsertImages(images []models.Image){
	if images == nil{
		return 
	}
	imgCollection := GetDB().Collection("image")
	ctx := context.Background()
	for _, img := range images{
		var result models.Image
		err := imgCollection.FindOne(ctx, bson.D{{Key: "url",Value: img.URL}}).Decode(&result)
		if err == mongo.ErrNoDocuments{// didn't find image with same url
			_, err := imgCollection.InsertOne(ctx, img)
			if err != nil{
				panic(err);
			}
			fmt.Println(img)
		}
	}
}
