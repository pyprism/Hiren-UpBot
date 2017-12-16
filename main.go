package main

import (
	"fmt"
	"github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/pyprism/Hiren-UpBot/models"
	"github.com/pyprism/Hiren-UpBot/views"
	"github.com/spf13/viper"
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
	db := pg.Connect(&pg.Options{
		Database: viper.GetString("db_name"),
		User:     viper.GetString("db_user"),
		Password: viper.GetString("db_password"),
	})

	//dbErr := createSchema(db)
	//if dbErr != nil {
	//	panic(dbErr)
	//}
	defer db.Close()

	// routers
	router.GET("/", views.Login)
	router.POST("/", views.Login)
	router.GET("/signup/", views.SignUp)
	router.POST("/signup/", views.SignUp)

	log.Fatal(router.Run(viper.GetString("port")))
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{&models.User{}} {
		err := db.CreateTable(model, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
