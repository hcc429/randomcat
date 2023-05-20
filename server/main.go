package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/hcc429/randomcat/controllers"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.GET("/", controllers.GetRandImage)
	r.Run()
}