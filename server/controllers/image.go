package controllers

import (
	"context"
	"image"
	_ "image/png"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hcc429/randomcat/db"
	"github.com/hcc429/randomcat/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const Tolerance = 0.1

var imgCollection *mongo.Collection = db.GetDB().Collection("image")
var validate = validator.New()

func AddImage(c *gin.Context) {
	// upload single file
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"form error": err.Error()})
	}
	img_file, err := file.Open()
	defer img_file.Close()
	img, _, err := image.DecodeConfig(img_file)

	c.SaveUploadedFile(file, "./cdn-path")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var image models.Image
	defer cancel()

	//validate the request body
	if err := c.BindJSON(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//use the validator library to validate required fields
	if validationErr := validate.Struct(&image); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	newImage := models.image{
		//////// REPLACE THIS PART ////////
		URL:         "",
		W:           img.Width,
		H:           img.Height,
		AspectRatio: float64(img.Width) / float64(img.Height),
		SizeKB:      100,
		UploadTime:  time.Now(),
		UploadBy:    "",
		//////// REPLACE THIS PART ////////
	}

	result, err := imgCollection.InsertOne(ctx, newImage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": result}) // result include acknowledged, insertedId
}

func GetRandImage(c *gin.Context) {
	width, errW := strconv.Atoi(c.Query("width"))
	height, errH := strconv.Atoi(c.Query("height"))
	if errW != nil || errH != nil {
		c.JSON(400, gin.H{"error": "invalid query: width, height"})
		return
	}
	aspect_ratio := float64(width) / float64(height)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{
		"$and", bson.A{
			bson.D{"aspect_ratio", bson.D{"$gte", aspect_ratio - Tolerance}},
			bson.D{"aspect_ratio", bson.D{"$lte", aspect_ratio + Tolerance}},
		},
	}
	opts := mongo.options.Find().SetProjection(bson.D{"image_url", 1})
	cur, err := imgCollection.Find(ctx, filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var results []struct {
		URL string
	}
	if err = cur.All(ctx, &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rand.Seed(time.Now().Unix())
	selected := results[rand.Intn(len(results))]

	c.Redirect(http.StatusMovedPermanently, selected.URL)
}

func GetImageById(c *gin.Context) {
	//
}

func DeleteImageById(c *gin.Context) {
	//
}
