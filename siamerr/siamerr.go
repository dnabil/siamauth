package siamerr

import "errors"

var (
	ErrorNotLoggedIn error = errors.New("please login first")
	ErrorLoggedIn    error = errors.New("already logged in")
)