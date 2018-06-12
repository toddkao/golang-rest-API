package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/toddkao/ecomm2/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
