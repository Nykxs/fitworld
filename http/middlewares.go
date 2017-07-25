package http

import (
	"github.com/labstack/echo"
	"github.com/nykxs/fitworld"
)

type middlewaresHandler struct {
	sessionService fitworld.SessionService
	userService    fitworld.UserService
}

func RegisterMiddlewaresHandler(s *Server) {
	m := middlewaresHandler{
		sessionService: s.SessionService,
		userService:    s.UserService,
	}
	s.Middlewares = &m
}

func (m *middlewaresHandler) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get Cookie in this Middleware

		// @TODO : Handle error correctly here
		return next(c)
	}
}
