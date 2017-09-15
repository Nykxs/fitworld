package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	"github.com/nykxs/fitworld"
	"github.com/nykxs/fitworld/http/validator"
)

type sessionHandler struct {
	sessionService fitworld.SessionService
	userService    fitworld.UserService
}

// RegisterSessionHandler will handle all reaquired actions to expose http endpoints to interact with sessions.
func RegisterSessionHandler(s *Server) {
	sessionHandler := &sessionHandler{
		sessionService: s.SessionService,
		userService:    s.UserService,
	}

	sessionGroup := s.Router.Group("/session")

	sessionGroup.POST("/", sessionHandler.Login)
	sessionGroup.GET("/delete", sessionHandler.Logout)
}

// SessionLoginPayload defines fields that should be sent to the Login endpoint to use it.
type SessionLoginPayload struct {
	Email    string `json:"email" form:"email" query:"email" valid:"required,email"`
	Password string `json:"password" form:"password" query:"password" valid:"required"`
}

// Login endpoint is used to create a session for a user without already existing session.
// It takes as parameter an Email/Password pair and will check info using userService.
// If everything's ok, it will call the sessionService.
func (h *sessionHandler) Login(c echo.Context) error {
	payload := new(SessionLoginPayload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	validator := httpvalidator.NewValidator()
	if err := validator.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// @TODO : Check that currentuser is nil
	//	-> else : check cookie, session and returns

	user, err := h.userService.GetByEmail(payload.Email)
	if err != nil {
		if err == fitworld.ErrUserNotFound {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	session, err := h.sessionService.CreateSession(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Set Cookie
	cookie := http.Cookie{
		Name:    CookieSession,
		Value:   session.ID,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	}
	c.SetCookie(&cookie)

	return nil
}

// Logout endpoint is used to delete the session associated to a user.
// If the given user does not have a session cookie set,the Logout funtion will not do anything.
func (h *sessionHandler) Logout(c echo.Context) error {
	cookie, err := c.Cookie(CookieSession)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	session, err := h.sessionService.GetSession(cookie.Value)
	if err != nil {
		if err == fitworld.ErrSessionNotFound {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	// @TODO : Check currentuser.ID and compare with session.userID

	if err := h.sessionService.DeleteSession(session.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	cookie.Expires = time.Now()
	c.SetCookie(cookie)
	return nil
}
