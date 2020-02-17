package actions

import "github.com/gobuffalo/buffalo"

// AchievementShow default implementation.
func AchievementShow(c buffalo.Context) error {
	return c.Render(200, r.HTML("achievement/show.html"))
}

// AchievementIndex default implementation.
func AchievementIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("achievement/index.html"))
}

// AchievementCreate default implementation.
func AchievementCreate(c buffalo.Context) error {
	return c.Render(200, r.HTML("achievement/create.html"))
}


