package models


type Image struct {
	URL         string    `bson:"url"`
	Width       int       `bson:"width"`
	Height      int       `bson:"height"`
	AspectRatio float64   `bson:"aspect_ratio"`
	Likes		int		  `bson:"likes"`
}


func NewImage(url string, width, height int)*Image{

	return &Image{
		URL: url,
		Width: width,
		Height: height,
		AspectRatio: float64(width) / float64(height),
		Likes: 0,
	}

}
