package write

import (
	"encoding/json"
	"net/http"

	"github.com/karlkeefer/pngr/golang/errors"
)

type errorResponse struct {
	Error string
}

func Error(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(errors.GetCode(err))
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
