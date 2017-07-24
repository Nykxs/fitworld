package mock

import "github.com/nykxs/fitworld"

type UserService struct {
	DB map[string]*fitworld.User

	RegisterFn      func(*fitworld.UserRegister) (*fitworld.User, error)
	RegisterInvoked bool

	MatchPasswordFn      func(email string, password string) (bool, error)
	MatchPasswordInvoked bool

	GetByIDFn      func(string) (*fitworld.User, error)
	GetByIDInvoked bool

	GetByEmailFn      func(string) (*fitworld.User, error)
	GetByEmailInvoked bool

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
func (s *UserService) MatchPassword(email string, password string) (bool, error) {
	s.MatchPasswordInvoked = true
	return s.MatchPasswordFn(email, password)
}

func (s *UserService) GetByID(id string) (*fitworld.User, error) {
	s.GetByIDInvoked = true
	return s.GetByIDFn(id)
}

func (s *UserService) GetByEmail(email string) (*fitworld.User, error) {
	s.GetByEmailInvoked = true
	return s.GetByEmailFn(email)
}

func (s *UserService) Delete(id string) error {
	s.DeleteInvoked = true
	return s.DeleteFn(id)
}
