package domain

import (
	"fmt"
)

type User struct {
	Id    int64
	Name  string
	Email string
}

type UserUninitialized struct {
	Name     string
	Email    string
	Password string
}

func NewUserUninitialized(name string, email string, password string) (UserUninitialized, error) {
	nameLength := len([]rune(name))

	if nameLength < 1 {
		return UserUninitialized{}, ErrUserNameIsRequired
	} else if nameLength > 32 {
		return UserUninitialized{}, ErrUserNameIsTooLong
	}

	passwordLength := len([]rune(password))

	if passwordLength < 8 || passwordLength > 64 {
		return UserUninitialized{}, ErrUserPasswordTooShort
	}

	return UserUninitialized{
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}

func (u *User) Validate() error {
	nameLength := len([]rune(u.Name))

	if nameLength < 1 || nameLength > 32 {
		return fmt.Errorf("invalid `name` len: %d", nameLength)
	}

	return nil
}
