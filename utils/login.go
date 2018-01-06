package utils

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

//func AuthMiddleware() gin.HandlerFunc  {
//	return func(c *gin.Context) {
//		if strings.HasPrefix(c.Request.URL.Path, "/login") ||
//			strings.HasPrefix(c.Request.URL.Path, "/signup") {
//			return
//		}
//		if strings.HasPrefix(c.Request.URL.Path, "/static") {
//			return
//		}
//
//		session := sessions.Default(c)
//		bunny := session.Get("authenticated")
//		if bunny == nil || bunny == false {

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

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasPrefix(c.Request().URL.Path, "/static") {
			return next(c)
		}

		if c.Request().URL.Path == "/" || c.Request().URL.Path == "/signup" {
			return next(c)
		}

		sess, _ := session.Get("session", c)
		bunny := sess.Values["authenticated"]
		if bunny == nil || bunny != true {
			return c.Redirect(http.StatusPermanentRedirect, "/")
		} else {
			return next(c)
		}

	}

}
