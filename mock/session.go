package mock

import "github.com/nykxs/fitworld"

type SessionService struct {
	CreateSessionFn      func(string) (*fitworld.Session, error)
	CreateSessionInvoked bool

	LoginFn      func(string, string) (*fitworld.Session, error)
	LoginInvoked bool

	GetSessionFn      func(string) (*fitworld.Session, error)
	GetSessionInvoked bool

	DeleteSessionFn      func(string) error
	DeleteSessionInvoked bool
}

func NewSessionService() *SessionService {
	return &SessionService{}
}

func (s *SessionService) CreateSession(userID string) (*fitworld.Session, error) {
	s.CreateSessionInvoked = true
	return s.CreateSessionFn(userID)
}

func (s *SessionService) Login(email string, password string) (*fitworld.Session, error) {
	s.LoginInvoked = true
	return s.LoginFn(email, password)
}

func (s *SessionService) GetSession(id string) (*fitworld.Session, error) {
	s.GetSessionInvoked = true
	return s.GetSessionFn(id)
}

func (s *SessionService) DeleteSession(id string) error {
	s.DeleteSessionInvoked = true
	return s.DeleteSessionFn(id)
}
