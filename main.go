package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/pyprism/Hiren-UpBot/views"
	"os"
)



func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", views.Login)
	router.POST("/", views.Login)
	router.GET("/signup/", views.SignUp)
	router.POST("/signup/", views.SignUp)

	PORT := os.Getenv("PORT")
	log.Fatal(router.Run(PORT))
}


