package views

import (
	"log"
	//"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	//"github.com/pyprism/Hiren-UpBot/db"
)

func Home(c *gin.Context) {
	session := sessions.Default(c)
	bunny := session.Get("authenticated")
	log.Println(bunny)
	c.HTML(http.StatusOK, "home.tmpl", gin.H{"title": "home", "username": "hiren"})
}
