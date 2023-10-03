package siamauth

import "errors"

var (
	// not logged in err
	ErrNotLoggedIn 	error 	= errors.New("please login first")
	// logged in err, y u try to log back in
	ErrLoggedIn    	error 	= errors.New("already logged in")
	// caused by siam UI changes/passed in the wrong page
	ErrNoElement 	error	= errors.New("element not found")
	// login fail
	ErrLoginFail	error	= errors.New("login fail, wrong credentials?")
)