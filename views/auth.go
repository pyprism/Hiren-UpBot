package views

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pyprism/Hiren-UpBot/db"
	//"github.com/pyprism/Hiren-UpBot/models"
	"golang.org/x/crypto/bcrypt"
	//"log"
)

type LoginForm struct {
	User     string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var xxx = db.Hiren{}

func init() {
	xxx.Connect()
}

// login page
func Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		count := xxx.UserCount()
		log.Println(count)
		c.HTML(http.StatusOK, "login.tmpl", gin.H{})
	} else if c.Request.Method == "POST" {
		var form LoginForm
		if err := c.ShouldBind(&form); err == nil {
			//var hiren = db.GetDB()
			//var user models.User
			//hiren.Where(&models.User{UserName: form.User}).First(&user)
			//if user.ID > 0 { // if user found
			//	ok := CheckPasswordHash(form.Password, user.Password)
			//	if ok {
			//		session := sessions.Default(c)
			//		session.Set("username", form.User)
			//		session.Set("authenticated", true)
			//		session.Save()
			//		c.Redirect(http.StatusMovedPermanently, "/dashboard/")
			//	}
			//}
			c.HTML(http.StatusForbidden, "login.tmpl", gin.H{"status": "Username/Password is not valid!"})
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

			//var hiren = db.GetDB()
			//var count int64
			//hiren.Model(models.User{}).Count(&count)
			////count := xxx.Count()
			////log.Println(count)
			//
			//if count == 0 {
			//	hash, err := HashPassword(form.Password)
			//	if err != nil {
			//		log.Fatal(err)
			//	}
			//	user := models.User{UserName: form.User, Password: hash, Admin: true}
			//	hiren.Create(&user)
			//	success := hiren.NewRecord(user)
			//	if !success {
			//		c.HTML(http.StatusCreated, "signup.tmpl", gin.H{"status": "Signed up successfully!"})
			//	} else {
			//		c.HTML(http.StatusForbidden, "signup.tmpl", gin.H{"status": "Username already exists!"})
			//	}
			//} else {
			//	hash, err := HashPassword(form.Password)
			//	if err != nil {
			//		log.Fatal(err)
			//	}
			//	user := models.User{UserName: form.User, Password: hash, Admin: false}
			//	hiren.Create(&user)
			//	success := hiren.NewRecord(user)
			//	if !success {
			//		c.HTML(http.StatusCreated, "signup.tmpl", gin.H{"status": "Signed up successfully!"})
			//	} else {
			//		c.HTML(http.StatusForbidden, "signup.tmpl", gin.H{"status": "Username already exists!"})
			//	}
			//
			//}

		} else {
			c.HTML(http.StatusBadRequest, "signup.tmpl", gin.H{"status": err.Error()})
		}

	}
}

// logout and stfu
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("username", "")
	session.Set("authenticated", false)
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/")
}
