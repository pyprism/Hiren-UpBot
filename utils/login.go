package utils

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func AuthMiddleware(c *gin.Context) {
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
		c.Redirect(http.StatusPermanentRedirect, "/")
	} else {
		c.Next()
	}

}