package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/pyprism/Hiren-UpBot/views"
	"fmt"
)



func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", views.Login)
	router.POST("/", views.Login)

	PORT := ":8000"
	fmt.Println("Server running on:", PORT)
	log.Fatal(router.Run(PORT))
}


