package views

import (
	"github.com/flosch/pongo2"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/pyprism/Hiren-UpBot/db"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var bunny = db.Hiren{}

func init() {
	bunny.Connect()
}

// login page
func Login(c echo.Context) error {
	if c.Request().Method == "GET" {
		data := pongo2.Context{
			"title": "Sign In",
		}
		return c.Render(http.StatusOK, "templates/login.html", data)
	} else if c.Request().Method == "POST" {
		user, er := bunny.FindUserByUsername(c.FormValue("username"))
		if er == nil { // if user was found
			ok := CheckPasswordHash(c.FormValue("password"), user.Password)
			if ok {
				sess, _ := session.Get("session", c)
				sess.Options = &sessions.Options{
					Path:     "/",
					MaxAge:   86400 * 7,
					HttpOnly: true,
				}
				sess.Values["username"] = c.FormValue("username")
				sess.Values["authenticated"] = true
				sess.Save(c.Request(), c.Response())
				return c.Redirect(http.StatusPermanentRedirect, "/dashboard/")
			} else {
				data := pongo2.Context{
					"title":  "Sign In",
					"status": "Username/Password is not valid!",
				}
				return c.Render(http.StatusOK, "templates/login.html", data)
			}
		} else {
			data := pongo2.Context{
				"title":  "Sign In",
				"status": "Username/Password is not valid!",
			}
			return c.Render(http.StatusOK, "templates/login.html", data)
		}
	} else {
		return c.String(http.StatusBadRequest, "f@ck off")
	}
}

// sign up page
//func SignUp(c *gin.Context) {
//	if c.Request.Method == "GET" {
//		c.HTML(http.StatusOK, "signup.tmpl", gin.H{})
//	} else if c.Request.Method == "POST" {
//		var form LoginForm
//		if err := c.ShouldBind(&form); err == nil {
//
//			count := bunny.UserCount()
//			hash, err := HashPassword(form.Password)
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			if count == 0 {
//				success := bunny.UserCreate(form.User, hash, true)
//				if !success {
//					c.HTML(http.StatusCreated, "signup.tmpl", gin.H{"status": "Signed up successfully!"})
//				} else {
//					c.HTML(http.StatusForbidden, "signup.tmpl", gin.H{"status": "Username already exists!"})
//				}
//			} else {
//				success := bunny.UserCreate(form.User, hash, false)
//				if !success {
//					c.HTML(http.StatusCreated, "signup.tmpl", gin.H{"status": "Signed up successfully!"})
//				} else {
//					c.HTML(http.StatusForbidden, "signup.tmpl", gin.H{"status": "Username already exists!"})
//				}
//
//			}
//
//		} else {
//			c.HTML(http.StatusBadRequest, "signup.tmpl", gin.H{"status": err.Error()})
//		}
//
//	}
//}

// logout and stfo
func Logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Values["username"] = ""
	sess.Values["authenticated"] = ""
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, "/")
}
