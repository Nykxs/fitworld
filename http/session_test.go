package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"strings"

	"github.com/nykxs/fitworld"
	"github.com/nykxs/fitworld/mock"
)

var (
	sessionLoginJSON = `{"email":"johndoe@mail.com", "password":"123abc"}`
)

func TestLogin(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(sessionLoginJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	fakeUser := fitworld.User{
		Email:    "johndoe@mail.com",
		Password: "123abc",
	}

	mockUser := mock.NewUserService()
	mockUser.DB[fakeUser.Email] = &fakeUser
	mockUser.GetByEmailFn = func(email string) (*fitworld.User, error) {
		user, ok := mockUser.DB[email]
		if !ok {
			return nil, errors.New("empty")
		}
		return user, nil
	}

	mockSession := mock.NewSessionService()
	mockSession.CreateSessionFn = func(userID string) (*fitworld.Session, error) {
		session := fitworld.Session{
			ID:     "session-id",
			UserID: "user-id",
		}
		return &session, nil
	}

	handler := &sessionHandler{
		userService:    mockUser,
		sessionService: mockSession,
	}

	if assert.NoError(t, handler.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
func TestLoginNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(sessionLoginJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUser := mock.NewUserService()
	mockUser.GetByEmailFn = func(email string) (*fitworld.User, error) {
		user, ok := mockUser.DB[email]
		if !ok {
			return nil, fitworld.ErrUserNotFound
		}
		return user, nil
	}

	mockSession := mock.NewSessionService()
	mockSession.CreateSessionFn = func(userID string) (*fitworld.Session, error) {
		session := fitworld.Session{
			ID:     "session-id",
			UserID: "user-id",
		}
		return &session, nil
	}

	handler := &sessionHandler{
		userService:    mockUser,
		sessionService: mockSession,
	}

	if assert.NoError(t, handler.Login(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestLogout(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(ContextKeyCurrentUser, &fitworld.User{
		Email:    "johndoe@mail.com",
		Password: "123abc",
	})
	c.Request().AddCookie(&http.Cookie{
		Name:  CookieSession,
		Value: "session-id",
		Path:  "/",
	})

	mockUser := mock.NewUserService()
	mockUser.GetByEmailFn = func(email string) (*fitworld.User, error) {
		user, ok := mockUser.DB[email]
		if !ok {
			return nil, fitworld.ErrUserNotFound
		}
		return user, nil
	}

	mockSession := mock.NewSessionService()

	mockSession.GetSessionFn = func(userID string) (*fitworld.Session, error) {
		return &fitworld.Session{
			ID:     "session-id",
			UserID: "user-id",
		}, nil
	}

	mockSession.DeleteSessionFn = func(userID string) error {
		return nil
	}

	handler := &sessionHandler{
		userService:    mockUser,
		sessionService: mockSession,
	}

	if assert.NoError(t, handler.Logout(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, mockSession.GetSessionInvoked)
		assert.True(t, mockSession.DeleteSessionInvoked)
	}
}
