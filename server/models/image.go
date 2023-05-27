package models

import "math"

// import (
// 	"fmt"
// 	"strconv"
// )


type Image struct {
	URL         string    `bson:"url"`
	Width       int       `bson:"width"`
	Height      int       `bson:"height"`
	AspectRatio float64   `bson:"aspect_ratio"`
	Likes		int		  `bson:"likes"`
}


func NewImage(url string, width, height int)*Image{
	aspect_ratio := float64(width) / float64(height)
	aspect_ratio = math.Round(aspect_ratio * 100) / 100
	return &Image{
		URL: url,
		Width: width,
		Height: height,
		AspectRatio: aspect_ratio,
		Likes: 0,
	}

}
