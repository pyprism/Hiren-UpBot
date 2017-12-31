package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pyprism/Hiren-UpBot/utils"

	//"github.com/pyprism/Hiren-UpBot/db"
	"github.com/pyprism/Hiren-UpBot/views"
	"github.com/spf13/viper"
)

func main() {
	router := echo.New()

	//middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	//router.Use(helmet.Default())
	//router.Use(gzip.Gzip(gzip.BestCompression))
	//
	router.Static("/static", "static")
	//router.LoadHTMLGlob("templates/*")

	renderer := utils.Renderer{
		Debug: true,
	}

	router.Renderer = renderer

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
	//store := sessions.NewCookieStore([]byte(viper.GetString("secret_key")))
	//router.Use(sessions.Sessions("bunny", store))
	//router.Use(utils.AuthMiddleware())

	// routers
	router.GET("/", views.Login)
	router.POST("/", views.Login)
	//router.GET("/signup/", views.SignUp)
	//router.POST("/signup/", views.SignUp)
	//router.GET("/logout/", views.Logout)
	//router.GET("/dashboard/", views.Home)

	router.Logger.Fatal(router.Start(viper.GetString("PORT")))
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
