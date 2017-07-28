package http

import (
	"time"

	"github.com/labstack/echo"
	"github.com/nykxs/fitworld"
)

type middlewares struct {
	sessionService fitworld.SessionService
	userService    fitworld.UserService
}

func RegisterMiddlewares(s *Server) {
	m := middlewares{
		sessionService: s.SessionService,
		userService:    s.UserService,
	}
	s.Middlewares = &m
}

func (m *middlewares) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(CookieSession)
		if err != nil {
			return next(c)
		}

		session, err := m.sessionService.GetSession(cookie.Value)
		if err != nil {
			cookie.Expires = time.Now()
			c.SetCookie(cookie)
			return err
		}
		user, err := m.userService.GetByID(session.UserID)
		if err != nil {
			cookie.Expires = time.Now()
			c.SetCookie(cookie)
			return err
		}

		c.Set(ContextKeyCurrentUser, user)
		return next(c)
	}
}
