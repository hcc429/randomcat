package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	_ "image/png"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hcc429/randomcat/db"
	"github.com/hcc429/randomcat/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Tolerance     = 2
	binWidth      = 0.25
	URLExpireTime = 5
)

var imgCollection *mongo.Collection = db.GetDB().Collection("image")
var validate = validator.New()

func GetRandImage(c *gin.Context) {
	width, errW := strconv.Atoi(c.Query("w"))
	height, errH := strconv.Atoi(c.Query("h"))
	if errW != nil || errH != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query: width, height"})
		return
	}
	aspect_ratio := float64(width) / float64(height)
	// Check if redis has cache
	binNum := strconv.Itoa(int(math.Floor(aspect_ratio / binWidth)))
	url, err := db.GetValue(binNum)

	if err != nil {
		fmt.Println("cache miss")
		filter := bson.D{
			{Key: "$and", Value: bson.A{
				bson.D{{Key: "aspect_ratio", Value: bson.D{{Key: "$gte", Value: aspect_ratio - Tolerance}}}},
				bson.D{{Key: "aspect_ratio", Value: bson.D{{Key: "$lte", Value: aspect_ratio + Tolerance}}}},
			},
			},
		}

		cur, err := db.FindImagesByFilter(filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var results []struct {
			URL string
		}
		ctx := context.Background()
		if err = cur.All(ctx, &results); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		selected := results[rand.Intn(len(results))]
		db.AddKeyValuePair(binNum, selected.URL, URLExpireTime)
		url = selected.URL
	}
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func validSort(s string) bool {
	validSortMethod := []string{"most_like", "newest", "oldest"}
	for _, method := range validSortMethod {
		if s == method {
			return true
		}
	}
	return false
}

func GetImagePage(c *gin.Context) {
	page, errP := strconv.Atoi(c.Query("p"))
	limit, errL := strconv.Atoi(c.Query("limit"))
	sortMethod := c.Query("sort")
	if errP != nil || errL != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query: page, limit"})
		return
	}
	if !validSort(sortMethod) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query: sort should be one of the following: most_like, newest, oldest"})
	}

	filter := bson.D{}
	skip := int64((page - 1) * limit)
	sortBy := bson.D{}
	switch sortMethod {
	case "most_like":
		sortBy = bson.D{{"likes", -1}}
	case "newest":
		sortBy = bson.D{{"_id", -1}}
	case "oldest":
		sortBy = bson.D{{"_id", 1}}
	default:
		sortBy = bson.D{{"likes", -1}}
	}
	opts := options.Find().SetSort(sortBy).SetLimit(int64(limit)).SetSkip(skip)

	cur, err := db.FindImagesByFilterAndOption(filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var results []models.Image
	ctx := context.Background()
	err = cur.All(ctx, &results)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var returnVal []models.Image
	for _, r := range results {
		returnVal = append(returnVal, r)
	}
	jsonVal, _ := json.Marshal(returnVal)
	c.JSON(http.StatusOK, gin.H{"images": string(jsonVal)})
	return
}

func LikeImage(c *gin.Context) {
	imageUrl := c.PostForm("URL")
	filter := bson.D{{"url", imageUrl}}
	update := bson.D{{"$inc", bson.D{{"likes", 1}}}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err := db.UpdateImage(filter, update, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return
}
