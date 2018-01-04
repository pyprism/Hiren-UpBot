package views

import (
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type (
	URL struct {
		Name string `form:"name" validate:"required"`
		Url  string `form:"url" validate:"required,url"`
		Poll int    `form:"poll" validate:"required,min=1"`
		Slow int    `form:"slow" validate:"required,min=1"`
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

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
		return c.Render(http.StatusOK, "templates/add_domain.html", data)
	} else {
		return c.String(http.StatusBadRequest, "f@ck off")
	}
}
