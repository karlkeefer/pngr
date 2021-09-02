package errors

import (
	"errors"
	"net/http"
)

var (
	BadRequestMethod = errors.New(http.StatusText(http.StatusMethodNotAllowed))
	InternalError    = errors.New(http.StatusText(http.StatusInternalServerError))

	NoJSONBody   = errors.New("Unable to decode JSON")
	InvalidEmail = errors.New("Invalid Email")

	FailedLogin          = errors.New("Invalid Email or Password")
	AlreadyRegistered    = errors.New("An account already exists for this email")
	VerificationNotFound = errors.New("Invalid Verification Code")
	VerificationExpired  = errors.New("Verification Code Was Already Used")

	UserNotFound  = errors.New("User does not exist")
	PostNotFound  = errors.New("Post does not exist")
	ResetNotFound = errors.New("Invalid password reset code")

	BadCSRF           = errors.New("Missing CSRF Header")
	BadOrigin         = errors.New("Invalid Origin Header")
	RouteUnauthorized = errors.New("You don't have permission to view this resource")
	RouteNotFound     = errors.New("Route not found")
	ExpiredToken      = errors.New("Your access token expired")
	InvalidToken      = errors.New("Your access token is invalid")
)

// codeMap returns a map of errors to http status codes
func codeMap() map[error]int {
	return map[error]int{
		BadRequestMethod: http.StatusMethodNotAllowed,
		InternalError:    http.StatusInternalServerError,

		NoJSONBody:        http.StatusBadRequest,
		InvalidEmail:      http.StatusBadRequest,
		AlreadyRegistered: http.StatusBadRequest,

		FailedLogin:          http.StatusUnauthorized,
		VerificationNotFound: http.StatusNotFound,
		VerificationExpired:  http.StatusGone,
		UserNotFound:         http.StatusNotFound,
		PostNotFound:         http.StatusNotFound,
		ResetNotFound:        http.StatusNotFound,

		BadCSRF:           http.StatusUnauthorized,
		BadOrigin:         http.StatusUnauthorized,
		RouteUnauthorized: http.StatusUnauthorized,
		RouteNotFound:     http.StatusNotFound,
		ExpiredToken:      http.StatusUnauthorized,
	}
}

// GetCode is a helper to get the relevant code for an error, or just return 500
func GetCode(e error) (bool, int) {
	if code, ok := codeMap()[e]; ok {
		return true, code
	}
	return false, http.StatusInternalServerError
}
