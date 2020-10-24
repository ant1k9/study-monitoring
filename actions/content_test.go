package actions

import (
	"net/http"

	_ "github.com/ant1k9/study-monitoring/models"
)

func (as *ActionSuite) Test_Content_Save() {
	u, err := as.createUser()
	as.NoError(err)
	defer as.DB.Destroy(u)

	res := as.HTML("/auth/new/").Post(u)
	as.Equal(http.StatusMovedPermanently, res.Code)

	res = as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)
	as.NotContains(res.Body.String(), "hobby-bobby")

	res = as.HTML("/content/save").Post(
		map[string]string{
			"type": "reading",
			"tag":  "hobby-bobby ",
			"time": "100",
		},
	)
	as.Equal(303, res.Code)

	res = as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "hobby-bobby")
	as.NotContains(res.Body.String(), "hobby-bobby ")
}

func (as *ActionSuite) Test_Saved_Content_Only_For_Creator() {
	u1, err := as.createUser()
	as.NoError(err)
	defer as.DB.Destroy(u1)

	res := as.HTML("/auth/new/").Post(u1)
	as.Equal(http.StatusMovedPermanently, res.Code)

	res = as.HTML("/content/save").Post(
		map[string]string{
			"type": "reading ",
			"tag":  "hobby-bobby",
			"time": "100",
		},
	)
	as.Equal(303, res.Code)

	res = as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "hobby-bobby")

	as.HTML("/auth").Delete()
	as.Equal(http.StatusOK, res.Code)

	u2, err := as.createUser()
	as.NoError(err)
	defer as.DB.Destroy(u2)

	res = as.HTML("/auth/new/").Post(u2)
	as.Equal(http.StatusMovedPermanently, res.Code)

	res = as.HTML("/").Get()
	as.Equal(http.StatusOK, res.Code)
	as.NotContains(res.Body.String(), "hobby-bobby")
}
