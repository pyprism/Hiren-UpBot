package utils

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	//"log"
)


func AuthMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/login") ||
			strings.HasPrefix(c.Request.URL.Path, "/signup") {
			return
		}
		if strings.HasPrefix(c.Request.URL.Path, "/static") {
			return
		}

		session := sessions.Default(c)
		bunny := session.Get("authenticated")
		if bunny == nil || bunny == false {

//func AuthMiddleware(c *gin.Context) {
//	if strings.HasPrefix(c.Request.URL.Path, "/") ||
//		strings.HasPrefix(c.Request.URL.Path, "/signup") {
//		return
//	}
//	if strings.HasPrefix(c.Request.URL.Path, "/static") {
//		return
//	}
//
//	session := sessions.Default(c)
//	bunny := session.Get("authenticated")
//	log.Println(bunny)
//	if bunny == nil || bunny != true {
//		c.Redirect(http.StatusPermanentRedirect, "/")
//	} else {
//		c.Next()
//	}
//
//}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//	if strings.HasPrefix(c.Request.URL.Path, "/xx") ||
		//		strings.HasPrefix(c.Request.URL.Path, "/signup") {
		//		return
		//	}
		if strings.HasPrefix(c.Request.URL.Path, "/static") {
			return
		}

		if c.Request.URL.Path == "/" || c.Request.URL.Path == "/signup" {
			return
		}

		session := sessions.Default(c)
		bunny := session.Get("authenticated")
		if bunny == nil || bunny != true {
			c.Redirect(http.StatusPermanentRedirect, "/")
		} else {
			c.Next()
		}

	}

}

