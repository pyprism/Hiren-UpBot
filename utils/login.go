package utils

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func AuthMiddleware(c *gin.Context) {
	if strings.HasPrefix(c.Request.URL.Path, "/") ||
		strings.HasPrefix(c.Request.URL.Path, "/signup") {
		return
	}
	if strings.HasPrefix(c.Request.URL.Path, "/static") {
		return
	}

	session := sessions.Default(c)
	bunny := session.Get("authenticated")
	log.Println(bunny)
	if bunny == nil || bunny == "" {
		c.Redirect(http.StatusPermanentRedirect, "/")
	} else {
		log.Println("sasa")
		c.Next()
	}
	log.Println(bunny)

}
