package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pyprism/Hiren-UpBot/views"
	"github.com/spf13/viper"
	"github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/gzip"
	"github.com/pyprism/Hiren-UpBot/models"
	"log"
	//"os"
)

func main() {
	router := gin.Default()

	//middleware
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	// config file
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// database
	db, err := gorm.Open("postgres", "host=localhost"+" user="+viper.GetString("db_user")+
		" dbname="+viper.GetString("db_name")+" sslmode=disable"+" password="+viper.GetString("db_password"))
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&models.User{})

	// routers
	router.GET("/", views.Login)
	router.POST("/", views.Login)
	router.GET("/signup/", views.SignUp)
	router.POST("/signup/", views.SignUp)

	log.Fatal(router.Run(viper.GetString("port")))
}
