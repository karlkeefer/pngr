package write

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/karlkeefer/pngr/golang/errors"
)

type errorResponse struct {
	Error string `json:"error"`
}

// Error is a shortcut for converting custom error types into appropriate response codes and JSON
func Error(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, code := errors.GetCode(err)
		if !found {
			// unexpected error - we should clean this up to avoid showing sql errors in the browser
			log.Printf("%s %s: Unexpected Error: %s", r.Method, r.URL.Path, err)
			err = errors.InternalError
		}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&errorResponse{Error: err.Error()})
	}
}

// JSON is a shortcut for writing JSON api responses
func JSON(obj interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(obj)
	}
}

// JSONorErr is a simple helper to handle the last `if err != nil` check in a handler
func JSONorErr(obj interface{}, err error) http.HandlerFunc {
	if err != nil {
		return Error(err)
	}

	return JSON(obj)
}

// SuccessOrErr is a simple helper to handle the last `if err != nil` check in a handler
func SuccessOrErr(err error) http.HandlerFunc {
	if err != nil {
		return Error(err)
	}

	return Success()
}

// Success writes a generic { success: true } response
func Success() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&map[string]bool{"success": true})
	}
}
