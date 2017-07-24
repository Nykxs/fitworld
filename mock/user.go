package mock

import "github.com/nykxs/fitworld"

type UserService struct {
	DB map[string]*fitworld.User

	RegisterFn      func(*fitworld.UserRegister) (*fitworld.User, error)
	RegisterInvoked bool

	GetFn      func(string) (*fitworld.User, error)
	GetInvoked bool

	DeleteFn      func(string) error
	DeleteInvoked bool
}

func NewUserService() *UserService {
	return &UserService{
		DB: make(map[string]*fitworld.User),
	}
}

func (s *UserService) Register(u *fitworld.UserRegister) (*fitworld.User, error) {
	s.RegisterInvoked = true
	return s.RegisterFn(u)
}

func (s *UserService) Get(id string) (*fitworld.User, error) {
	s.GetInvoked = true
	return s.GetFn(id)
}

func (s *UserService) Delete(id string) error {
	s.DeleteInvoked = true
	return s.DeleteFn(id)
}
