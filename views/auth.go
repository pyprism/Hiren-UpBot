package views

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pyprism/Hiren-UpBot/db"
	"golang.org/x/crypto/bcrypt"
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

var bunny = db.Hiren{}

func init() {
	bunny.Connect()
}

// login page
func Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{})
	} else if c.Request.Method == "POST" {
		var form LoginForm
		if err := c.ShouldBind(&form); err == nil {
			user, er := bunny.FindUserByUsername(form.User)
			if er == nil { // if user was found
				ok := CheckPasswordHash(form.Password, user.Password)
				if ok {
					session := sessions.Default(c)
					session.Set("username", form.User)
					session.Set("authenticated", true)
					session.Save()
					c.Redirect(http.StatusMovedPermanently, "/dashboard/")
				} else {
					c.HTML(http.StatusForbidden, "login.tmpl", gin.H{"status": "Username/Password is not valid!"})
				}
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

			count := bunny.UserCount()
			hash, err := HashPassword(form.Password)
			if err != nil {
				log.Fatal(err)
			}

			if count == 0 {
				success := bunny.UserCreate(form.User, hash, true)
				if !success {
					c.HTML(http.StatusCreated, "signup.tmpl", gin.H{"status": "Signed up successfully!"})
				} else {
					c.HTML(http.StatusForbidden, "signup.tmpl", gin.H{"status": "Username already exists!"})
				}
			} else {
				success := bunny.UserCreate(form.User, hash, false)
				if !success {
					c.HTML(http.StatusCreated, "signup.tmpl", gin.H{"status": "Signed up successfully!"})
				} else {
					c.HTML(http.StatusForbidden, "signup.tmpl", gin.H{"status": "Username already exists!"})
				}

			}

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
