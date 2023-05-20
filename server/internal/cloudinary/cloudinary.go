package cloudinary

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
)



func GetCredentials() (*cloudinary.Cloudinary, context.Context){

	cld, err := cloudinary.New()// New will init from env variable CLOUDINARY_URL
	if err != nil{
		fmt.Println("error when creating cloudinary")
	}
	cld.Config.URL.Secure = true 
	ctx := context.Background()
	return cld, ctx
}

func GetAssetInfo(cld * cloudinary.Cloudinary, ctx context.Context)(*admin.AssetsResult, error){
	resp, err := cld.Admin.Assets(ctx, admin.AssetsParams{})
	return resp, err 
}


