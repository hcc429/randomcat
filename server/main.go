package main

import (
	_ "github.com/joho/godotenv/autoload"
	"randomcat/db"


	"github.com/gin-gonic/gin"
)

func main(){

	db.InitClient()
	db.GetDB()
	r := gin.Default()


	r.Run()
}