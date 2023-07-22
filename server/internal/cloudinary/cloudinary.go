package cloudinary

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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

func GetAssetInfo(cld *cloudinary.Cloudinary, ctx context.Context)(*admin.AssetsResult, error){
	resp, err := cld.Admin.Assets(ctx, admin.AssetsParams{NextCursor: ctx.Value("NextCursor").(string), MaxResults: 50})
	return resp, err 
}

type Transform struct{
	Width int
	Height int
	Blur int
}

func GetTransformUrl(cld *cloudinary.Cloudinary, publicID string, transform *Transform)(string, error){

	image, err := cld.Image(publicID)
	if err != nil{
		return "", err 
	}
	filter := make([]string, 0)
	if transform.Width == 0{ // default
		filter = append(filter, "w_300")
	} else {
		filter = append(filter, "w_" + strconv.Itoa(transform.Width))
	}
	
	if transform.Height != 0{
		filter = append(filter, "h_" + strconv.Itoa(transform.Height))
	}
	if transform.Blur != 0{
		filter = append(filter, "e_blur:" + strconv.Itoa(transform.Blur))
	}

	image.Transformation = strings.Join(filter, ",")
	return image.String()
}
