package utils

import "errors"

var ErrDuplicateUserNameOrPassword = errors.New("user name or email already exist")