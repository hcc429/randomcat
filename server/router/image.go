package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hcc429/randomcat/controllers"
	"github.com/hcc429/randomcat/middleware"
)

func AddImageRoute(r *gin.Engine) {
	image := r.Group("/image")
	image.Use(middleware.RateLimit)
	{
		image.GET("", controllers.GetRandImage)
	}
	images := r.Group("/images")
	images.Use(middleware.RateLimit)
	{
		images.GET("", controllers.GetImages)
		images.POST("/like", controllers.LikeImage)
	}
}
