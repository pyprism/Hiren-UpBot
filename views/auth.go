package views

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pyprism/Hiren-UpBot/models"
	"github.com/spf13/viper"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//"golang.org/x/crypto/bcrypt"
	//"github.com/pyprism/Hiren-UpBot/models"
)

type LoginForm struct {
	User string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

var db *gorm.DB

func init()  {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	db, err := gorm.Open("postgres", "host=localhost"+" user="+viper.GetString("db_user")+
		" dbname="+viper.GetString("db_name")+" sslmode=disable"+" password="+viper.GetString("db_password"))
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
}


func Login(c *gin.Context)  {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{})
	} else if c.Request.Method == "POST" {
		var form LoginForm
		if err:=c.ShouldBind(&form); err == nil {
			if form.User == "demo" && form.Password == "demo" {
				c.HTML(http.StatusAccepted, "login.tmpl", gin.H{"status": "connected"})
			} else {
				c.HTML(http.StatusForbidden, "login.tmpl", gin.H{"status": "Username/Password is not valid!"})
			}
		} else {
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{"status": err.Error()})
		}

	}
}

func SignUp(c *gin.Context) {
	if c.Request.Method == "GET" {
		var user []models.User
		db.Count(&user)
		fmt.Println(user)

		c.HTML(http.StatusOK, "signup.tmpl", gin.H{})
	} else if c.Request.Method == "POST" {
		var form LoginForm
		if err:=c.ShouldBind(&form); err == nil {
			if form.User == "demo" && form.Password == "demo" {
				c.HTML(http.StatusAccepted, "signup.tmpl", gin.H{"status": "Signed up successfully!"})
			} else {
				c.HTML(http.StatusForbidden, "signup.tmpl", gin.H{"status": "Username/Password is not valid!"})
			}
		} else {
			c.HTML(http.StatusBadRequest, "signup.tmpl", gin.H{"status": err.Error()})
		}

	}
}