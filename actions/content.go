package actions

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gofrs/uuid"

	"github.com/ant1k9/study-monitoring/models"
)

// ContentSave default implementation.
func ContentSave(c buffalo.Context) error {
	req := c.Request()
	if err := req.ParseForm(); err != nil {
		c.Logger().Errorf("content save: invalid form: %s", err)
		return c.Redirect(http.StatusSeeOther, "/")
	}

	t, err := strconv.ParseInt(req.Form.Get("time"), 10, 64)
	if err != nil {
		c.Logger().Errorf("content save: convert time: %s, %s", err, req.Form)
		return c.Redirect(http.StatusSeeOther, "/")
	}

	currentUserID := c.Session().Get("current_user_id")
	uid, ok := currentUserID.(uuid.UUID)
	if !ok {
		c.Logger().Errorf("content save: convert current_user_id to uuid: %s", currentUserID)
		return c.Redirect(http.StatusSeeOther, "/")
	}

	cont := &models.Content{
		Tag:    normalize(req.Form.Get("tag")),
		Type:   normalize(req.Form.Get("type")),
		Time:   t,
		UserID: uid,
	}
	err = models.DB.Create(cont)
	if err != nil {
		c.Logger().Errorf("content save: save to DB: %s", err)
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func normalize(s string) string {
	return strings.Trim(strings.ToLower(s), " ")
}
