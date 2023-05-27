package main 

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/hcc429/randomcat/internal/cloudinary"
	"github.com/hcc429/randomcat/db"
)

func main(){
	cld, ctx := cloudinary.GetCredentials()
	resp, err := cloudinary.GetAssetInfo(cld, ctx)
	if err != nil{
		panic("error when getting cloudinary info" + err.Error())
	}
	
	image_table := db.SyncImages(resp.Assets)
	db.ClearUnusedImages(*image_table)
}