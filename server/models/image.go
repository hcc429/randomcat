package models


type Image struct {
	URL         string    `bson:"url"`
	Width       int       `bson:"width"`
	Height      int       `bson:"height"`
	AspectRatio float64   `bson:"aspect_ratio"`
	Likes		int		  `bson:"likes"`
	PublicID    string    `bson:"public_id"`
}


func NewImage(url string, width, height int, public_id string)*Image{
	aspect_ratio := float64(width) / float64(height)
	return &Image{
		URL: url,
		Width: width,
		Height: height,
		AspectRatio: aspect_ratio,
		Likes: 0,
		PublicID: public_id,
	}

}
