package http

import (
	"context"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"

	"github.com/nykxs/fitworld"
)

// Server defines the object that will manage all the HTTP stuff (Endpoints, router...)
type Server struct {
	Router      *echo.Echo
	UserService fitworld.UserService
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

	// Register middleware here too.
	// Register endpoints in this function.
	RegisterUserHandler(s)

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
