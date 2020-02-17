package actions

import "github.com/gobuffalo/buffalo"

// PersonalAchievementShow default implementation.
func PersonalAchievementShow(c buffalo.Context) error {
	return c.Render(200, r.HTML("personal_achievement/show.html"))
}

// PersonalAchievementIndex default implementation.
func PersonalAchievementIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("personal_achievement/index.html"))
}

// PersonalAchievementCreate default implementation.
func PersonalAchievementCreate(c buffalo.Context) error {
	return c.Render(200, r.HTML("personal_achievement/create.html"))
}

