package pg

import (
	"testing"

	"github.com/nykxs/fitworld"
	"github.com/stretchr/testify/require"
)

func Setup(t *testing.T) *UserService {
	store, err := NewStore("dbname=fitworld user=postgres sslmode=disable host=localhost")
	require.NoError(t, err)

	userService := NewUserService(store)
	return userService
}

func TestUserRegister(t *testing.T) {
	userService := Setup(t)
	user := &fitworld.UserRegister{
		FirstName: "Vincent",
		LastName:  "Vielle",
		Email:     "vincent.vielle@gmail.com",
		Password:  "123abc",
	}

	u, err := userService.Register(user)
	require.NoError(t, err)
	require.NotNil(t, u)
}

func TestUserGetByEmail(t *testing.T) {
	userService := Setup(t)
	user := &fitworld.UserRegister{
		FirstName: "Vincent",
		LastName:  "Vielle",
		Email:     "vincent.vielle@gmail.com",
		Password:  "123abc",
	}

	u, err := userService.Register(user)
	require.NoError(t, err)
	require.NotNil(t, u)

	resp, err := userService.GetByEmail("vincent.vielle@gmail.com")
	require.NoError(t, err)
	require.Equal(t, "Vincent", resp.FirstName)
}

func TestUserGetByID(t *testing.T) {
	userService := Setup(t)
	user := &fitworld.UserRegister{
		FirstName: "Vincent",
		LastName:  "Vielle",
		Email:     "vincent.vielle@gmail.com",
		Password:  "123abc",
	}

	u, err := userService.Register(user)
	require.NoError(t, err)
	require.NotNil(t, u)

	resp, err := userService.GetByID("1")
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestUserDelete(t *testing.T) {
	userService := Setup(t)
	user := &fitworld.UserRegister{
		FirstName: "Vincent",
		LastName:  "Vielle",
		Email:     "vincent.vielle@gmail.com",
		Password:  "123abc",
	}

	u, err := userService.Register(user)
	require.NoError(t, err)
	require.NotNil(t, u)

	err = userService.Delete("1")
	require.NoError(t, err)
}
