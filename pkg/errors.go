package pkg

import "errors"

var (
	// Login Error
	ErrUnauthorized = errors.New("unauthorized")
	// Hash Error
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)
