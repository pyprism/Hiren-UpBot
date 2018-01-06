package views

import (
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"net/http"
)

type URL struct {
	Name string `form:"name" validate:"required"`
	Url  string `form:"url" validate:"required,url"`
	Poll int64  `form:"poll" validate:"required,min=1"`
	Slow int64  `form:"slow" validate:"required,min=1"`
}

func AddDomain(c echo.Context) (err error) {
	sess, _ := session.Get("session", c)
	if c.Request().Method == "GET" {
		data := pongo2.Context{
			"title":    "Add New Host",
			"username": sess.Values["username"],
		}
		return c.Render(http.StatusOK, "templates/add_domain.html", data)
	} else if c.Request().Method == "POST" {
		u := new(URL)

		errData := pongo2.Context{
			"title":    "Add New Host",
			"username": sess.Values["username"],
			"error":    "Data is not valid!",
		}

		if err = c.Bind(u); err != nil {
			return c.Render(http.StatusConflict, "templates/add_domain.html", errData)
		}
		if err = c.Validate(u); err != nil {
			return c.Render(http.StatusConflict, "templates/add_domain.html", errData)
		}
		sess, _ := session.Get("session", c)
		data := pongo2.Context{
			"title":    "Add New Host",
			"username": sess.Values["username"],
			"success":  "New host added.",
		}
		var username, _ = sess.Values["username"].(string)
		ok := bunny.UrlCreate(u.Name, u.Url, username, u.Poll, u.Slow)
		if !ok {
			return c.Render(http.StatusOK, "templates/add_domain.html", data)
		} else {
			return c.String(http.StatusConflict, "WTF ! Something bad happened :O !") // lol :P
		}
	} else {
		return c.String(http.StatusBadRequest, "f@ck off")
	}
}

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
