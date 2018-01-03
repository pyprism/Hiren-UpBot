package views

import (
	"log"
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

type (
	URL struct {
		Name string `form:"name" validate:"required"`
		Url  string `form:"url" validate:"required,url"`
		Poll string `form:"poll" validate:"required,min=1"`
		Slow string `form:"slow" validate:"required,min=1"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func AddDomain(c echo.Context) (err error) {
	if c.Request().Method == "GET" {
		sess, _ := session.Get("session", c)
		data := pongo2.Context{
			"title":    "Add New Host",
			"username": sess.Values["username"],
		}
		return c.Render(http.StatusOK, "templates/add_domain.html", data)
	} else if c.Request().Method == "POST" {
		u := new(URL)
		if err = c.Bind(u); err != nil {
			log.Println(err)
			return
		}
		if err = c.Validate(u); err != nil {
			log.Println(err)
			return
		}
		sess, _ := session.Get("session", c)
		data := pongo2.Context{
			"title":    "Add New Host",
			"username": sess.Values["username"],
		}
		return c.Render(http.StatusOK, "templates/add_domain.html", data)
	} else {
		return c.String(http.StatusBadRequest, "f@ck off")
	}
}
