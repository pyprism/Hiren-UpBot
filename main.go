package main

import (
	"fmt"
	"github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
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
	connStr := "host=localhost" + " user=" + viper.GetString("db_user") +
		" dbname=" + viper.GetString("db_name") + " sslmode=disable" + " password=" + viper.GetString("db_password")
	db, dbErr := xorm.NewEngine("postgres", connStr)
	if dbErr != nil {
		panic(dbErr)
	}

	//dbErr := createSchema(db)
	//if dbErr != nil {
	//	panic(dbErr)
	//}
	db.ShowSQL(true)
	db.Logger().SetLevel(core.LOG_DEBUG)
	db.SetMapper(SameMapper{})
	defer db.Close()

	// routers
	router.GET("/", views.Login)
	router.POST("/", views.Login)
	router.GET("/signup/", views.SignUp)
	router.POST("/signup/", views.SignUp)

	log.Fatal(router.Run(viper.GetString("port")))
}
