package views

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func List(c echo.Context) error {
	sess, _ := session.Get("session", c)
	var username , _ = sess.Values["username"].(string)
	urls := bunny.UrlList(username)
	data := pongo2.Context{
		"title":    "Add New Host",
		"username": sess.Values["username"],
		"urls": urls,
	}
	return c.Render(http.StatusOK, "templates/list.html", data)
}
