package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/pyprism/Hiren-UpBot/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
