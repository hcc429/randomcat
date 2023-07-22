package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hcc429/randomcat/routers"
)

func main(){
	r := gin.Default()
	r.Use(cors.Default())
	router.InitRouter(r)
	r.Run("0.0.0.0:8080")
}