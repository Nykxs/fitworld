package pg

import (
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/nykxs/fitworld"
)

// UserService is the implementation of the UserService interface defined in our domain that uses PG.
type UserService struct {
	Store *Store
}

// NewUserService returns a new UserService implementation working with PG and implementing the interface from domain.
func NewUserService(s *Store) *UserService {
	return &UserService{
		Store: s,
	}
}

func (u *UserService) Register(user *fitworld.UserRegister) (*fitworld.User, error) {
	record := ToUserRegisterRecord(user)

	var userid int
	query := fmt.Sprintf(`INSERT INTO users(first_name, last_name, email, password)
	VALUES('%v', '%v', '%v', '%v') RETURNING id`, record.FirstName, record.LastName, record.Email, record.Password)

	err := u.Store.QueryRow(query).Scan(&userid)
	if err != nil {
		return nil, err
	}

	return &fitworld.User{
		ID:        strconv.Itoa(userid),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}

func (u *UserService) MatchPassword(email string, password string) (bool, error) {
	query := fmt.Sprintf(`SELECT password FROM users WHERE email = '%v'`, email)

	q, err := u.Store.Query(query)
	if err != nil {
		return false, err
	}

	var actualPassword string
	for q.Next() {
		err = q.Scan(&actualPassword)
		if err != nil {
			return false, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(actualPassword), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (u *UserService) GetByID(id string) (*fitworld.User, error) {
	query := fmt.Sprintf(`SELECT * FROM users WHERE id = '%v'`, id)

	q, err := u.Store.Query(query)
	if err != nil {
		return nil, err
	}

	var user UserRecord

	for q.Next() {
		err = q.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
	}

	resp := FromUserRecord(&user)
	return resp, nil
}

func (u *UserService) GetByEmail(email string) (*fitworld.User, error) {
	query := fmt.Sprintf(`SELECT * FROM users WHERE email = '%v'`, email)

	q, err := u.Store.Query(query)
	if err != nil {
		return nil, err
	}

	var user UserRecord

	for q.Next() {
		err = q.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
	}

	resp := FromUserRecord(&user)
	return resp, nil
}

func (u *UserService) Delete(id string) error {
	_, err := u.Store.Exec(`DELETE FROM users WHERE id=$1`, id)
	return err
}
