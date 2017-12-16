package views

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"

	"github.com/pyprism/Hiren-UpBot/models"
	"github.com/spf13/viper"
	//"golang.org/x/crypto/bcrypt"
)

type LoginForm struct {
	User     string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

var db *pg.DB

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	db := pg.Connect(&pg.Options{
		Database: viper.GetString("db_name"),
		User:     viper.GetString("db_user"),
		Password: viper.GetString("db_password"),
	})
	fmt.Println(reflect.TypeOf(db))
	count, err := db.Model(&models.User{}).Count()
	if err != nil {
		panic(err)
	}

	fmt.Println(count)
	//defer db.Close()
}

func Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		fmt.Println("s")
		c.HTML(http.StatusOK, "login.tmpl", gin.H{})
	} else if c.Request.Method == "POST" {
		var form LoginForm
		if err := c.ShouldBind(&form); err == nil {
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
		c.HTML(http.StatusOK, "signup.tmpl", gin.H{})
	} else if c.Request.Method == "POST" {
		var form LoginForm
		if err := c.ShouldBind(&form); err == nil {
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
