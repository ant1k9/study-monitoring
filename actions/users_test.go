package actions

import (
	"fmt"
	"time"

	"github.com/ant1k9/study-monitoring/models"
)

func (as *ActionSuite) Test_Users_New() {
	res := as.HTML("/users/new").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_Users_Create() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)

	u := &models.User{
		Email:                fmt.Sprintf("%d@test_users_create.com", time.Now().Nanosecond()),
		Password:             "password",
		PasswordConfirmation: "password",
	}

	res := as.HTML("/users").Post(u)
	as.Equal(302, res.Code)
	defer as.DB.Destroy(u)

	count, err = as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)
}
