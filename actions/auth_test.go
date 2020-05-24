package actions

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ant1k9/study-monitoring/models"
)

func (as *ActionSuite) createUser() (*models.User, error) {
	u := &models.User{
		Email:                fmt.Sprintf("%d@create_user.com", time.Now().Nanosecond()),
		Password:             "password",
		PasswordConfirmation: "password",
	}

	verrs, err := u.Create(as.DB)
	as.False(verrs.HasAny())

	return u, err
}

func (as *ActionSuite) Test_Auth_Signin() {
	res := as.HTML("/auth/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), `<a href="/auth/new/">Sign In</a>`)
}

func (as *ActionSuite) Test_Auth_New() {
	res := as.HTML("/auth/new").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign In")
}

func (as *ActionSuite) Test_Auth_Create() {
	u, err := as.createUser()
	as.NoError(err)
	defer as.DB.Destroy(u)

	tcases := []struct {
		Email       string
		Password    string
		Status      int
		RedirectURL string

		Identifier string
	}{
		{u.Email, u.Password, http.StatusFound, "/", "Valid"},
		{"noexist@example.com", "password", http.StatusUnauthorized, "", "Email Invalid"},
		{u.Email, "invalidPassword", http.StatusUnauthorized, "", "Password Invalid"},
	}

	for _, tcase := range tcases {
		as.Run(tcase.Identifier, func() {
			res := as.HTML("/auth/new/").Post(&models.User{
				Email:    tcase.Email,
				Password: tcase.Password,
			})

			as.Equal(tcase.Status, res.Code)
			as.Equal(tcase.RedirectURL, res.Location())
		})
	}
}

func (as *ActionSuite) Test_Auth_Redirect() {
	u, err := as.createUser()
	as.NoError(err)
	defer as.DB.Destroy(u)

	tcases := []struct {
		redirectURL    interface{}
		resultLocation string

		identifier string
	}{
		{"/some/url", "/some/url", "RedirectURL defined"},
		{nil, "/", "RedirectURL nil"},
		{"", "/", "RedirectURL empty"},
	}

	for _, tcase := range tcases {
		as.Run(tcase.identifier, func() {
			as.Session.Set("redirectURL", tcase.redirectURL)

			res := as.HTML("/auth/new").Post(u)

			as.Equal(302, res.Code)
			as.Equal(res.Location(), tcase.resultLocation)
		})
	}
}
