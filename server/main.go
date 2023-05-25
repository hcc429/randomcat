package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/hcc429/randomcat/router"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	router.AddImageRoute(r)
	router.AddMetricRoute(r)
	r.Run("0.0.0.0:8080")
}