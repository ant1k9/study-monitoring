package models

import (
	"fmt"
	"time"
)

func (ms *ModelSuite) Test_User_Create() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &User{
		Email:                fmt.Sprintf("%d@test_user_create.com", time.Now().Nanosecond()),
		Password:             "password",
		PasswordConfirmation: "password",
	}

	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	defer ms.DB.Destroy(u)

	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u.PasswordHash)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Create_ValidationErrors() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &User{
		Password: "password",
	}

	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_User_Create_UserExists() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u1 := &User{
		Email:                fmt.Sprintf("%d@test_user_create_exists.com", time.Now().Nanosecond()),
		Password:             "password",
		PasswordConfirmation: "password",
	}

	ms.Zero(u1.PasswordHash)

	verrs, err := u1.Create(ms.DB)
	defer ms.DB.Destroy(u1)

	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u1.PasswordHash)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)

	u2 := &User{
		Email:    fmt.Sprintf("%d@test_user_create_exists.com", time.Now().Nanosecond()),
		Password: "password",
	}

	verrs, err = u2.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}
