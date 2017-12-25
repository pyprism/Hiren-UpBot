package views

import (
	//"log"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/pyprism/Hiren-UpBot/db"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{"title": "home", "username": "hiren"})
}
