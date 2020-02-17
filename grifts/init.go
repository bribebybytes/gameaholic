package grifts

import (
	"github.com/gobuffalo/buffalo"
	"bribebybytes.in/gameaholic/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
