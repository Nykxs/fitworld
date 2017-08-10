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
	userRegisterJSON                = `{"email":"John Doe","email":"johndoe@mail.com","first_name":"John","last_name":"Doe","password":"123abc"}`
	userRegisterJSONInvalidEmail    = `{"email":"John Doe","first_name":"John","last_name":"Doe","password":"123abc"}`
	userRegisterJSONInvalidPassword = `{"first_name":"John","last_name":"Doe","email":"johndoe@mail.com"}`

	userJSON = `{"id":"user-01","email":"John Doe","email":"johndoe@mail.com","first_name":"John","last_name":"Doe","password":"123abc"}`
)

// ****
// TESTS RELATED TO REGISTER NEW USER
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
		return nil, nil
	}

	handler := &userHandler{
		userService: mockUser,
	}

	if assert.NoError(t, handler.Register(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Len(t, mockUser.DB, 1)
	}
}

func TestRegisterUserInvalidEmailPayload(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(userRegisterJSONInvalidEmail))
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
		return nil, nil
	}

	handler := &userHandler{
		userService: mockUser,
	}

	if assert.NoError(t, handler.Register(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Empty(t, mockUser.DB)
	}
}

func TestRegisterUserInvalidPasswordPayload(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(userRegisterJSONInvalidPassword))
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
		return nil, nil
	}

	handler := &userHandler{
		userService: mockUser,
	}

	if assert.NoError(t, handler.Register(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Empty(t, mockUser.DB)
	}
}

// ****
// TESTS RELATED TO GET USER
func TestGetUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("123abc")

	mockUser := mock.NewUserService()
	mockUser.DB["123abc"] = &fitworld.User{
		ID:        "user-01",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@mail.com",
		Password:  "123abc",
	}
	mockUser.GetByIDFn = func(id string) (*fitworld.User, error) {
		u, ok := mockUser.DB[id]
		if !ok {
			return nil, fitworld.ErrUserNotFound
		}

		return u, nil
	}

	handler := &userHandler{
		userService: mockUser,
	}

	if assert.NoError(t, handler.Get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, userJSON, rec.Body.String())
	}
}
func TestGetUserNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("123abc")

	mockUser := mock.NewUserService()
	mockUser.GetByIDFn = func(id string) (*fitworld.User, error) {
		u, ok := mockUser.DB[id]
		if !ok {
			return nil, fitworld.ErrUserNotFound
		}

		return u, nil
	}

	handler := &userHandler{
		userService: mockUser,
	}

	if assert.NoError(t, handler.Get(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
