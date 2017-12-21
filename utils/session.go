package utils

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)


func SetSession(userName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("username", userName)

		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			fmt.Println(count)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.Next()
	}
}