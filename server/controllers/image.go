package controllers

import (
	"context"
	_ "image/png"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hcc429/randomcat/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const Tolerance = 2
const binWidth = 2.0

var imgCollection *mongo.Collection = db.GetDB().Collection("image")
var validate = validator.New()

func GetRandImage(c *gin.Context) {
	width, errW := strconv.Atoi(c.Query("w"))
	height, errH := strconv.Atoi(c.Query("h"))
	if errW != nil || errH != nil {
		c.JSON(400, gin.H{"error": "invalid query: width, height"})
		return
	}
	aspect_ratio := float64(width) / float64(height)
	// Check if redis has cache
	binNum := strconv.Itoa(int(math.Floor(aspect_ratio / binWidth)))
	url, err := db.GetValue(binNum)

	db.GetDB()

	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, url)
		return
	}
	// Query db
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{
		{Key: "$and", Value: bson.A{
			bson.D{{Key: "aspect_ratio", Value: bson.D{{Key: "$gte", Value: aspect_ratio - Tolerance}}}},
			bson.D{{Key: "aspect_ratio", Value: bson.D{{Key: "$lte", Value: aspect_ratio + Tolerance}}}},
		},
		},
	}

	cur, err := imgCollection.Find(ctx, filter)
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

	c.Redirect(http.StatusTemporaryRedirect, selected.URL)

	db.AddKeyValuePair(binNum, selected.URL)

}

func GetImageById(c *gin.Context) {
	//
}

func DeleteImageById(c *gin.Context) {
	//
}
