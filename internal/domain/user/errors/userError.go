package errors

import "errors"

var (
	UserListEmptyError     = errors.New("user list empty")
	UserNotFoundError      = errors.New("user not found")
	UserAlreadyExistsError = errors.New("user already exists")
	UserNotAuthError       = errors.New("user not authorized")
)
