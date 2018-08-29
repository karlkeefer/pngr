package verify

import (
	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/utils"

	"encoding/json"
	"net/http"
)

type verifyResponse struct {
	Msg string
}

func ServeHTTP(env *env.Env, w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)
	if head == "" {
		errors.Write(w, errors.RouteNotFound)
		return
	}

	switch r.Method {
	case "GET":
		handleGet(head, env, w, r)
	default:
		errors.Write(w, errors.BadRequestMethod)
	}
}

func handleGet(code string, env *env.Env, w http.ResponseWriter, r *http.Request) {
	err, auth := env.UserRepo().Verify(code)
	if err != nil {
		errors.Write(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auth)
}
