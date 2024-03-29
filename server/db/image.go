package db

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/hcc429/randomcat/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func removeImageById(id primitive.ObjectID) {
	imgCollection := GetDB().Collection("image")
	ctx := context.Background()
	result, err := imgCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		panic(err)
	}
	fmt.Println(result.DeletedCount)
}

func FindImagesByFilter(filter interface{}) (*mongo.Cursor, error) {
	imgCollection := GetDB().Collection("image")
	ctx := context.Background()
	cursor, err := imgCollection.Find(ctx, filter)
	return cursor, err
}

func FindImagesByFilterAndOption(filter interface{}, opts *options.FindOptions) (*mongo.Cursor, error) {
	imgCollection := GetDB().Collection("image")
	ctx := context.Background()
	cursor, err := imgCollection.Find(ctx, filter, opts)
	return cursor, err
}

func UrlExist(url string) bool {
	imgCollection := GetDB().Collection("image")
	ctx := context.Background()
	var result models.Image
	err := imgCollection.FindOne(ctx, bson.D{{Key: "url", Value: url}}).Decode(&result)
	return err != mongo.ErrNoDocuments
}

func UpdateImage(filter interface{}, update interface{}, opts *options.FindOneAndUpdateOptions) error {
	imgCollection := GetDB().Collection("image")
	ctx := context.Background()
	var result models.Image
	err := imgCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	return err
}

func SyncImages(images []api.BriefAssetResult) *map[string]bool {
	image_table := map[string]bool{}
	imgCollection := GetDB().Collection("image")

	for _, image := range images {
		image_table[image.SecureURL] = true
		if !UrlExist(image.SecureURL) {
			ctx := context.Background()
			img := models.NewImage(image.SecureURL, image.Width, image.Height, image.PublicID)
			_, err := imgCollection.InsertOne(ctx, img)
			if err != nil {
				panic(err)
			}
			fmt.Println("Insert New Image to Database, url: ", image.SecureURL)
		}
	}
	return &image_table
}

func ClearUnusedImages(image_table map[string]bool) {
	imgCollection := GetDB().Collection("image")
	ctx := context.Background()
	cursor, err := imgCollection.Find(ctx, bson.D{{}})
	if err != nil {
		panic(err)
	}
	var results []models.Image
	if err = cursor.All(context.Background(), &results); err != nil {
		panic(err)
	}

	totalDelete := 0
	for _, image := range results {
		if !image_table[image.URL] { // image no longer exist
			result, err := imgCollection.DeleteOne(ctx, bson.M{"url": image.URL})
			if err != nil {
				panic(err)
			}
			totalDelete += int(result.DeletedCount)
		}
	}
	fmt.Println("Delete ", totalDelete, "image(s) from Database")
}
