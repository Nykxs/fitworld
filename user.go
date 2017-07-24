package fitworld

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
)

// User defines fields that can be filled for a user.
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// UserRegister defines mandatory fields that should be filled to register a new user using a UserService.
type UserRegister struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// UserService defines the behaviour that should be implemented by each object that want to manage users.
type UserService interface {
	Register(*UserRegister) (*User, error)
	MatchPassword(email string, password string) (bool, error)
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	Delete(id string) error
}
