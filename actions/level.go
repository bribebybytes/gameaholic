package actions

import "github.com/gobuffalo/buffalo"

// LevelShow default implementation.
func LevelShow(c buffalo.Context) error {
	return c.Render(200, r.HTML("level/show.html"))
}

// LevelIndex default implementation.
func LevelIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("level/index.html"))
}

// LevelCreate default implementation.
func LevelCreate(c buffalo.Context) error {
	return c.Render(200, r.HTML("level/create.html"))
}





