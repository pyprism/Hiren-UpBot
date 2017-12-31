package views

import (
	"net/http"
	"log"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo-contrib/session"

	"github.com/labstack/echo"

	//"github.com/pyprism/Hiren-UpBot/db"
)

func Dashboard(c echo.Context) error{

	sess, _ := session.Get("session", c)
	log.Println(sess.Values["authenticated"])
	data := pongo2.Context{
		"title": "Dashboard",
		"username": sess.Values["username"],
	}
	return c.Render(http.StatusOK, "templates/dashboard.html", data)
}
