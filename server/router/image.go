package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hcc429/randomcat/controllers"
	"github.com/hcc429/randomcat/middlewares"
)


func AddImageRoute(r *gin.Engine){
	image := r.Group("/image")
	image.Use(middleware.RateLimit)
	{
		image.GET("/", controllers.GetRandImage)
	}
}