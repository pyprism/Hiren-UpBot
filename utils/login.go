package utils

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	//"log"
)

<<<<<<< HEAD
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
=======
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
>>>>>>> edea98f5a84f2f88fee6a11b19e6688f7d64c1a2
			c.Redirect(http.StatusPermanentRedirect, "/")
		} else {
			c.Next()
		}

	}
<<<<<<< HEAD
}
=======
}
>>>>>>> edea98f5a84f2f88fee6a11b19e6688f7d64c1a2
