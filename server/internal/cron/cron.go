package main

import (
	_ "github.com/joho/godotenv/autoload"
	"context"

	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/hcc429/randomcat/db"
	"github.com/hcc429/randomcat/internal/cloudinary"
)

func main(){
	cld, ctx := cloudinary.GetCredentials()

	ctx = context.WithValue(ctx, "NextCursor", "")
	assets := admin.AssetsResult{}
	for{
		resp, err := cloudinary.GetAssetInfo(cld, ctx)
		
		if err != nil{
			panic("error when getting cloudinary info" + err.Error())
		}
		assets.Assets = append(assets.Assets, resp.Assets...)
		if resp.NextCursor == ""{
			break
		}
		ctx = context.WithValue(ctx, "NextCursor", resp.NextCursor)
	}
	image_table := db.SyncImages(assets.Assets)
	db.ClearUnusedImages(*image_table)
}