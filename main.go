package main

import (
	"fmt"
	"github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	//"github.com/pyprism/Hiren-UpBot/db"
	"github.com/pyprism/Hiren-UpBot/views"
	"github.com/spf13/viper"
	"log"
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
	//db.Init()

	// cookie based session
	store := sessions.NewCookieStore([]byte(viper.GetString("secret_key")))
	router.Use(sessions.Sessions("bunny", store))

	// routers
	router.GET("/", views.Login)
	router.POST("/", views.Login)
	router.GET("/signup/", views.SignUp)
	router.POST("/signup/", views.SignUp)
	router.GET("/logout/", views.Logout)

	log.Fatal(router.Run(viper.GetString("port")))
}

//func main() {
//	conString := ""
//	db, err := gorm.Open("postgres", conString)
//	if err != nil {
//		panic("failed to connect database")
//	}
//	defer db.Close()
//	//db.DropTable(&models.User{})
//	db.AutoMigrate(&models.User{})
//	var count int64
//	db.Debug().Model(&models.User{}).Count(&count)
//	fmt.Println(count)
//}
