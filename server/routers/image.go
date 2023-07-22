package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hcc429/randomcat/controllers"
	"github.com/hcc429/randomcat/middlewares"
)

func AddImageRoute(r *gin.Engine) {
	image := r.Group("/image")
	image.Use(middlewares.RateLimit)
	{
		image.GET("", controllers.GetRandImage)
	}
	images := r.Group("/images")
	images.Use(middlewares.RateLimit)
	{
		images.GET("", controllers.GetImages)
		images.POST("/like", controllers.LikeImages)
	}
}
