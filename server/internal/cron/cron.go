package main 

import (
	"github.com/hcc429/randomcat/internal/cloudinary"
	"github.com/hcc429/randomcat/models"
	"github.com/hcc429/randomcat/db"
)

func main(){
	cld, ctx := cloudinary.GetCredentials()
	resp, err := cloudinary.GetAssetInfo(cld, ctx)
	if err != nil{
		panic("error when getting cloudinary info" + err.Error())
	}
	
	images := make([]models.Image, len(resp.Assets))
	for i, v := range resp.Assets{
		images[i] = *models.NewImage(v.URL, v.Width, v.Height)
	}
	db.InsertImages(images)
}