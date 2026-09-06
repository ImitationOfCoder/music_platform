package domain

import (
	"errors"
)

var (
	ErrUserNotFound         = errors.New("User not found.")
	ErrUserNameIsTooLong    = errors.New("User's name is too long.")
	ErrUserNameIsRequired   = errors.New("User's name is required.")
	ErrUserPasswordTooShort = errors.New("User's password is too short.")
	ErrUserAlreadyExists    = errors.New("User already exists.")

	ErrInvalidCredentials = errors.New("Invalid credentials.")
)
