package auth

import (
	"errors"
)

var (
	// errors that are specific to the auth domain
	ErrInvalidCredentials = errors.New("invalid credentials")
)
