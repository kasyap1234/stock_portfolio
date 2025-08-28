package appErrors

import "errors"

var ErrEmptyPassword = errors.New("empty password")

var ErrInvalidCredentials = errors.New("invalid credentials")
