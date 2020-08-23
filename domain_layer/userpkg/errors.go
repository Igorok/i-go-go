package userpkg

import "errors"

var (
	ErrUserNotFound       = errors.New("user_not_found")
	ErrInvalidAccessToken = errors.New("invalid_access_token")
)
