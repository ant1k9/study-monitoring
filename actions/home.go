package actions

import (
	"math"
	"net/http"

	"github.com/gobuffalo/buffalo"

	"github.com/ant1k9/study-monitoring/models"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	uid := c.Session().Get("current_user_id")

	var content models.Contents
	err := models.DB.Select(`tag, type, SUM(time)::INT "time"`).
		Where("user_id = ? AND TO_CHAR(created_at, 'iyyy-iw') = TO_CHAR(NOW(), 'iyyy-iw')", uid).
		GroupBy("tag", "type").
		All(&content)
	if err != nil {
		c.Logger().Errorf("get home page: %s", err)
	}

	currentWeekTime, avgWeekTime := &models.Content{}, &models.Content{}
	err = models.DB.Select(`SUM(time) AS time`).
		Where(`TO_CHAR(created_at, 'iyyy-iw') = TO_CHAR(NOW(), 'iyyy-iw')`).
		First(currentWeekTime)
	if err != nil {
		c.Logger().Errorf("get home page: %s", err)
	}

	err = models.DB.RawQuery(`
		SELECT ROUND(AVG(t)) as time FROM (
			SELECT SUM(time) t, TO_CHAR(created_at, 'iyyy-iw') w
				FROM contents
				GROUP BY w
		) _ WHERE w != TO_CHAR(NOW(), 'iyyy-iw')
		`).First(avgWeekTime)
	if err != nil {
		c.Logger().Errorf("get home page: %s", err)
	}

	c.Set("contents", content)
	c.Set("types", content.GetUniqueTypes())
	c.Set("progress", asPercents(currentWeekTime, avgWeekTime))
	return c.Render(http.StatusOK, r.HTML("index.html"))
}

func asPercents(currentWeekTime, avgWeekTime *models.Content) float64 {
	if avgWeekTime != nil && currentWeekTime != nil && currentWeekTime.Time < avgWeekTime.Time {
		return math.Round(float64(currentWeekTime.Time) / float64(avgWeekTime.Time) * 100)
	}
	return 100.0
}
