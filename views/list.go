package views

import (
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"net/http"
)

func List(c echo.Context) error {
	sess, _ := session.Get("session", c)
	var username, _ = sess.Values["username"].(string)
	urls := bunny.UrlList(username)
	data := pongo2.Context{
		"title":    "Add New Host",
		"username": sess.Values["username"],
		"urls":     urls,
	}
	return c.Render(http.StatusFound, "templates/list.html", data)
}

func HostByID(c echo.Context) error {
	id := c.Param("id")
	sess, _ := session.Get("session", c)
	var username, _ = sess.Values["username"].(string)
	url, err := bunny.FindHostById(username, id)
	if err != nil {
		return c.String(http.StatusNotFound, "nai :/")
	} else {
		data := pongo2.Context{
			"title":    "Details of " + url.Name,
			"username": username,
			"url":      url,
		}
		return c.Render(http.StatusFound, "templates/host_by_id.html", data)
	}
}
