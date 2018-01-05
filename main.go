package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/pyprism/Hiren-UpBot/db"
	"github.com/pyprism/Hiren-UpBot/utils"
	"gopkg.in/go-playground/validator.v9"
	"github.com/pyprism/Hiren-UpBot/views"
	"github.com/spf13/viper"
)

func main() {
	router := echo.New()

	// config file
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	//middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.Secure())
	router.Use(session.Middleware(sessions.NewCookieStore([]byte(viper.GetString("secret_key")))))
	router.Use(utils.AuthMiddleware)
	//router.Use(helmet.Default())
	//router.Use(gzip.Gzip(gzip.BestCompression))
	//
	router.Static("/static", "static")
	//router.LoadHTMLGlob("templates/*")

	renderer := utils.Renderer{
		Debug: true,
	}

	router.Renderer = renderer
	router.Validator = &utils.CustomValidator{Validator: validator.New()}

	// database
	bunny := db.Hiren{}
	bunny.Connect()

	// cookie based session
	//store := sessions.NewCookieStore([]byte(viper.GetString("secret_key")))
	//router.Use(sessions.Sessions("bunny", store))
	//router.Use(utils.AuthMiddleware())

	// routers
	router.GET("/", views.Login)
	router.POST("/", views.Login)
	//router.GET("/signup/", views.SignUp)
	//router.POST("/signup/", views.SignUp)
	router.GET("/logout/", views.Logout)
	router.GET("/dashboard/", views.Dashboard)
	router.Any("/add/", views.AddDomain)
	router.GET("/list/:id/", views.HostByID)
	router.GET("/list/", views.List)

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
