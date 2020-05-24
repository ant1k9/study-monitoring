package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"

	"github.com/ant1k9/study-monitoring/models"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	uid := c.Session().Get("current_user_id")

	var content models.Contents
	err := models.DB.Select(`tag, type, sum(time)::int "time"`).
		Where("user_id = ?", uid).
		GroupBy("tag", "type").
		All(&content)

	if err != nil {
		c.Logger().Errorf("get home page: %s", err)
	}

	c.Set("contents", content)
	c.Set("types", content.GetUniqueTypes())
	return c.Render(http.StatusOK, r.HTML("index.html"))
}
