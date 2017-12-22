package views

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pyprism/Hiren-UpBot/db"
	"github.com/pyprism/Hiren-UpBot/models"
	"github.com/pyprism/Hiren-UpBot/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type LoginForm struct {
	User     string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

//var db *gorm.DB
//
//func init() {
//	viper.SetConfigName("config")
//	viper.AddConfigPath(".")
//	err := viper.ReadInConfig()
//	if err != nil { // Handle errors reading the config file
//		panic(fmt.Errorf("Fatal error config file: %s \n", err))
//	}
//
//	db, err := gorm.Open("postgres", "host=localhost"+" user="+viper.GetString("db_user")+
//		" dbname="+viper.GetString("db_name")+" sslmode=disable"+" password="+viper.GetString("db_password"))
//	if err != nil {
//		panic("failed to connect database")
//	}
//	//defer db.Close()
//}

//func x() {
//	var dbx = db.GetDB()
//	user := new(models.User)
//	total, err := dbx.Count(user)
//	fmt.Println(total, err)
//}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// login page
func Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{})
	} else if c.Request.Method == "POST" {
		var form LoginForm
		if err := c.ShouldBind(&form); err == nil {
			var hiren = db.GetDB()
			var user = models.User{UserName: form.User}
			has, err := hiren.Get(&user)
			if err != nil {
				log.Fatal(err)
			}
			var ok bool
			if has {
				ok = CheckPasswordHash(form.Password, user.Password)
			}
			if has && ok {
				session := sessions.Default(c)
				session.Set("username", form.User)
				session.Set("authenticated", true)
				session.Save()
				c.HTML(http.StatusAccepted, "login.tmpl", gin.H{"status": "connected"})
			} else {
				c.HTML(http.StatusForbidden, "login.tmpl", gin.H{"status": "Username/Password is not valid!"})
			}
		} else {
			c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{"status": err.Error()})
		}

	}
}

// sign up page
func SignUp(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "signup.tmpl", gin.H{})
	} else if c.Request.Method == "POST" {
		var form LoginForm
		if err := c.ShouldBind(&form); err == nil {

			var hiren = db.GetDB()
			user := new(models.User)
			count, err := hiren.Count(user)
			if err != nil {
				log.Fatal(err)
			}

			if count == 0 {
				hash, err := HashPassword(form.Password)
				if err != nil {
					log.Fatal(err)
				}
				user := new(models.User)
				user.UserName = form.User
				user.Password = hash
				user.Admin = true
				affected, err := hiren.Insert(user)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(affected)
				fmt.Println(user.Id)
				c.HTML(http.StatusAccepted, "signup.tmpl", gin.H{"status": "Signed up successfully!"})
			} else {
				hash, err := HashPassword(form.Password)
				if err != nil {
					log.Fatal(err)
				}
				user := new(models.User)
				user.UserName = form.User
				user.Password = hash
				user.Admin = false
				affected, err := hiren.Insert(user)
				if err != nil {
					c.HTML(http.StatusNotAcceptable, "signup.tmpl", gin.H{"status": "Username already exists!"})
				}
				if affected == 1 {
					c.HTML(http.StatusAccepted, "signup.tmpl", gin.H{"status": "Signed up successfully!"})
				}
			}

		} else {
			c.HTML(http.StatusBadRequest, "signup.tmpl", gin.H{"status": err.Error()})
		}

	}
}
