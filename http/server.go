package http

import (
	"context"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/middleware"
	"github.com/nykxs/fitworld"
)

const (
	// ContextKeyCurrentUser defines the name of the key in our echo.Context to get CurrentUser if any
	ContextKeyCurrentUser = "currentUser"
	// CookieSession defines the name of the cookie that is created for each user's session
	CookieSession = "session"
)

// Server defines the object that will manage all the HTTP stuff (Endpoints, router...)
type Server struct {
	Router         *echo.Echo
	Middlewares    *middlewares
	UserService    fitworld.UserService
	SessionService fitworld.SessionService
}

// NewServer returns a Server up to run with fields initialized.
func NewServer(u fitworld.UserService) *Server {
	return &Server{
		Router:      echo.New(),
		UserService: u,
	}
}

// Setup function register both middlewares and http endpoints into the server in order to be exposed will Start function is called.
func (s *Server) Setup() error {
	s.Router.Logger.SetLevel(log.INFO)

	// Register middlewares.
	RegisterMiddlewares(s)

	// Defines middlewares we'll use in all our endpoints
	// First middleware should be a recover one
	// Second middleware should be a error handler one that , following the returned error, will format it
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.RequestID())
	s.Router.Use(s.Middlewares.Auth)

	// Register endpoints in this function.
	RegisterUserHandler(s)
	RegisterSessionHandler(s)

	return nil
}

// Start runs the server and start to expose http endpoints.
func (s *Server) Start() error {
	if err := s.Router.Start(":8000"); err != nil {
		s.Router.Logger.Warn(err)
		return err
	}
	return nil
}

// Stop tries to stop the server as gracefully as possible.
func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Router.Shutdown(ctx); err != nil {
		s.Router.Logger.Warn(err)
		return err
	}
	return nil
}
