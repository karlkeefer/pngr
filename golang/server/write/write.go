package write

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/karlkeefer/pngr/golang/errors"
)

type errorResponse struct {
	Error string
}

func Error(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		found, code := errors.GetCode(err)
		if !found {
			// unexpected error - we should clean this up to avoid showing sql errors in the browser
			log.Println("Unexpected Error: ", err)
			err = errors.InternalError
		}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&errorResponse{Error: err.Error()})
	}
}

func JSON(obj interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(obj)
	}
}

func JSONorErr(obj interface{}, err error) http.HandlerFunc {
	if err != nil {
		return Error(err)
	}

	return JSON(obj)
}
