package controllers

import (
	"context"

	// "encoding/json"
	"fmt"
	_ "image/png"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hcc429/randomcat/cache"
	"github.com/hcc429/randomcat/controllers/dto"
	"github.com/hcc429/randomcat/db"
	"github.com/hcc429/randomcat/internal/cloudinary"
	"github.com/hcc429/randomcat/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Tolerance     = 0.15
	IntervalSize  = 0.25
	URLExpireTime = 5
)

var (
	imgCollection *mongo.Collection = db.GetDB().Collection("image")
	validate                        = validator.New()
	cld, _                          = cloudinary.GetCredentials()
)

func GetRandImage(c *gin.Context) {
	width, errW := strconv.Atoi(c.Query("w"))
	height, errH := strconv.Atoi(c.Query("h"))

	var filter interface{}
	var err error
	var publicID string
	var results []struct {
		PublicID string `bson:"public_id"`
	}

	if errW != nil || errH != nil {
		
		publicID, err = cache.GetValue("Rand")
		//publicID, err = db.GetValue("Rand")
		if err != nil {
			//c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query: width, height"})
			filter = bson.D{{}}
			cur, err := db.FindImagesByFilter(filter)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			ctx := context.Background()
			if err = cur.All(ctx, &results); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			selected := results[rand.Intn(len(results))]

			cache.AddKeyValuePair("Rand", selected.PublicID, URLExpireTime)
			publicID = selected.PublicID
		}
	} else {
		aspect_ratio := float64(width) / float64(height)
		// Check if redis has cache
		interval := strconv.Itoa(int(math.Floor(aspect_ratio / IntervalSize)))
		publicID, err = cache.GetValue(interval)

		if err != nil {
			fmt.Println("cache miss")
			filter = bson.D{
				{Key: "$and", Value: bson.A{
					bson.D{{Key: "aspect_ratio", Value: bson.D{{Key: "$gte", Value: aspect_ratio - Tolerance}}}},
					bson.D{{Key: "aspect_ratio", Value: bson.D{{Key: "$lte", Value: aspect_ratio + Tolerance}}}},
				},
				},
			}

			cur, err := db.FindImagesByFilter(filter)
			//fmt.Print(cur)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			ctx := context.Background()
			if err = cur.All(ctx, &results); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			if len(results) == 0{
				// TODO: error handling
			}
			selected := results[rand.Intn(len(results))]
			cache.AddKeyValuePair(interval, selected.PublicID, URLExpireTime)
			publicID = selected.PublicID
		}
	}

	// Resize the Image to required size
	t := cloudinary.Transform{
		Width:  width,
		Height: height,
		Blur:   0,
	}
	//fmt.Println(publicID)
	URL, err := cloudinary.GetTransformUrl(cld, publicID, &t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cloudinary Transform failed"})
		return
	}
	//fmt.Println(URL)
	c.Redirect(http.StatusTemporaryRedirect, URL)
}
func GetImages(c *gin.Context) {
	page, errP := strconv.Atoi(c.Query("p"))
	limit, errL := strconv.Atoi(c.Query("limit"))
	sortMethod := c.Query("sort")
	if errP != nil || errL != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query: page, limit"})
		return
	}

	filter := bson.D{}
	skip := int64((page - 1) * limit)
	sortBy := bson.D{}
	switch sortMethod {
	case "most_like":
		sortBy = bson.D{{Key: "likes", Value:-1}}
	case "newest":
		sortBy = bson.D{{Key: "_id", Value: -1}}
	case "oldest":
		sortBy = bson.D{{Key: "_id", Value: 1}}
	default:
		sortBy = bson.D{{Key: "likes", Value: -1}, {Key: "_id", Value:  -1}}
	}
	opts := options.Find().SetSort(sortBy).SetLimit(int64(limit)).SetSkip(skip)

	cur, err := db.FindImagesByFilterAndOption(filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	results := make([]models.Image, 0)
	ctx := context.Background()
	err = cur.All(ctx, &results)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"images": results})
	return
}

func LikeImages(c *gin.Context) {
	likes := new(dto.Likes)
	if err := c.ShouldBindJSON(&likes); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}


	writeModels := make([]mongo.WriteModel, 0)
	for _, like := range likes.Likes{
		writeModels = append(writeModels, mongo.NewUpdateOneModel().SetFilter(
			bson.D{{Key: "url", Value: like.URL}}).SetUpdate(
				bson.D{{Key: "$inc", Value:  bson.D{{Key: "likes", Value: like.Amount}}}}))
	}

	opts := options.BulkWrite().SetOrdered(true)
	results, err := imgCollection.BulkWrite(context.Background(), writeModels, opts)
	if err != nil{
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"status": results})
	return
}
