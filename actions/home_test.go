package actions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ant1k9/study-monitoring/models"
)

func (as *ActionSuite) Test_HomeHandler() {
	res := as.HTML("/").Get()
	as.Equal(http.StatusMovedPermanently, res.Code)
	as.Equal(res.Location(), "/auth/new")
}

func (as *ActionSuite) Test_HomeHandler_LoggedIn() {
	u := &models.User{
		Email:                fmt.Sprintf("%d@test_logged_in.com", time.Now().Nanosecond()),
		Password:             "password",
		PasswordConfirmation: "password",
	}

	verrs, err := u.Create(as.DB)
	defer as.DB.Destroy(u)

	as.NoError(err)
	as.False(verrs.HasAny())
	as.Session.Set("current_user_id", u.ID)

	res := as.HTML("/auth").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Sign Out")

	as.Session.Clear()
	res = as.HTML("/auth").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Sign In")
}
