package domain

import (
	"fmt"
)

type User struct {
	Id           int64
	Name         string
	Email        string
	PasswordHash string
}

type UserUninitialized struct {
	Name     string
	Email    string
	Password string
}

func (u *User) Validate() error {
	nameLength := len([]rune(u.Name))

	if nameLength < 1 || nameLength > 32 {
		return fmt.Errorf("invalid `name` len: %d", nameLength)
	}

	return nil
}
