package fitworld

import (
	"time"
)

// Session defines fields that are used to manipulate a session.
type Session struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	ExpiresAt time.Time
}

// SessionService defines the behaviour that should be implemented by any object that want to handle sessions for a user.
type SessionService interface {
	CreateSession(userID string) (*Session, error)
	Login(email string, password string) (*Session, error)
	GetSession(id string) (*Session, error)
	DeleteSession(id string) error
}
