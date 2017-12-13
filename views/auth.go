package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
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