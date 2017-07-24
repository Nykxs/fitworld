package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/nykxs/fitworld"
	"github.com/nykxs/fitworld/mock"
	"github.com/stretchr/testify/assert"
)

var (
	userRegisterJSON        = `{"email":"John Doe","email":"johndoe@mail.com","first_name":"John","last_name":"Doe","password":"123abc"}`
	userRegisterJSONInvalid = `{"email":"John Doe","first_name":"John","last_name":"Doe"}`
)

func TestRegisterUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(userRegisterJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUser := mock.NewUserService()
	mockUser.RegisterFn = func(u *fitworld.UserRegister) (*fitworld.User, error) {
		stored := &fitworld.User{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Password:  u.Password,
		}

		mockUser.DB[stored.Email] = stored
		return stored, nil
	}

	handler := &userHandler{
		userService: mockUser,
	}

	if assert.NoError(t, handler.Register(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Len(t, mockUser.DB, 1)
	}
}

func TestRegisterUserInvalidPayload(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(userRegisterJSONInvalid))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUser := mock.NewUserService()
	mockUser.RegisterFn = func(u *fitworld.UserRegister) (*fitworld.User, error) {
		stored := &fitworld.User{
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Password:  u.Password,
		}

		mockUser.DB[stored.Email] = stored
		return stored, nil
	}

	handler := &userHandler{
		userService: mockUser,
	}

	if assert.NoError(t, handler.Register(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Empty(t, mockUser.DB)
	}
}
