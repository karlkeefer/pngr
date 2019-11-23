package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	BadRequestMethod = errors.New(http.StatusText(http.StatusMethodNotAllowed))
	InternalError    = errors.New(http.StatusText(http.StatusInternalServerError))

	NoJSONBody   = errors.New("Unable to decode JSON")
	InvalidEmail = errors.New("Invalid Email")

	FailedLogin          = errors.New("Invalid Email or Password")
	VerificationNotFound = errors.New("Invalid Verification Code")
	VerificationExpired  = errors.New("Verification Code Was Already Used")
	UserNotFound         = errors.New("User does not exist")

	BadCSRF           = errors.New("Missing CSRF Header")
	BadOrigin         = errors.New("Invalid Origin Header")
	RouteUnauthorized = errors.New("You don't have permission to view this resource")
	RouteNotFound     = errors.New("Route not found")
	ExpiredToken      = errors.New("Your access token expired")
)

// codeMap returns a map of errors to http status codes
func codeMap() map[error]int {
	return map[error]int{
		BadRequestMethod: http.StatusMethodNotAllowed,
		InternalError:    http.StatusInternalServerError,

		NoJSONBody:   http.StatusBadRequest,
		InvalidEmail: http.StatusBadRequest,

		FailedLogin:          http.StatusUnauthorized,
		VerificationNotFound: http.StatusNotFound,
		VerificationExpired:  http.StatusGone,
		UserNotFound:         http.StatusNotFound,

		BadCSRF:           http.StatusUnauthorized,
		BadOrigin:         http.StatusUnauthorized,
		RouteUnauthorized: http.StatusUnauthorized,
		RouteNotFound:     http.StatusNotFound,
		ExpiredToken:      http.StatusUnauthorized,
	}
}

// GetCode is a helper to get the relevant code for an error, or just return 500
func GetCode(e error) int {
	if code, ok := codeMap()[e]; ok {
		return code
	}
	return http.StatusInternalServerError
}

type errorResponse struct {
	Error string
}

// Write is a helper that writes an error out as JSON
func Write(w http.ResponseWriter, e error) {
	w.WriteHeader(GetCode(e))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&errorResponse{Error: e.Error()})
}
