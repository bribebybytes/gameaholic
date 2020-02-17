package actions

import "github.com/gobuffalo/buffalo"

// PlayerShow default implementation.
func PlayerShow(c buffalo.Context) error {
	return c.Render(200, r.HTML("player/show.html"))
}

// PlayerIndex default implementation.
func PlayerIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("player/index.html"))
}

// PlayerCreate default implementation.
func PlayerCreate(c buffalo.Context) error {
	return c.Render(200, r.HTML("player/create.html"))
}

