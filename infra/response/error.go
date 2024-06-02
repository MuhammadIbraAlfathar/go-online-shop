package response

import "errors"

// error general
var (
	ErrNotFound = errors.New("not found")
)

var (
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password must have minimum 6 character")
	ErrAuthIsNotExists       = errors.New("auth is not exists")
	ErrEmailAlreadyUse       = errors.New("email already use")
)
