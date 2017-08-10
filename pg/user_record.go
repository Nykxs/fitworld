package pg

import "github.com/nykxs/fitworld"

type UserRegisterRecord struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func ToUserRegisterRecord(user *fitworld.UserRegister) *UserRegisterRecord {
	return &UserRegisterRecord{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}
}
func FromUserRegisterRecord(user *UserRegisterRecord) *fitworld.UserRegister {
	return &fitworld.UserRegister{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}
}

type UserRecord struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func ToUserRecord(user *fitworld.User) *UserRecord {
	return &UserRecord{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func FromUserRecord(user *UserRecord) *fitworld.User {
	return &fitworld.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}
}
