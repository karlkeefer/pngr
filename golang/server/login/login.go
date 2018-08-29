package login

import (
	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models/user"

	"encoding/json"
	"net/http"
)

func ServeHTTP(env *env.Env, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handlePost(env, w, r)
	default:
		errors.Write(w, errors.BadRequestMethod)
	}
}

func handlePost(env *env.Env, w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u user.User
	err := decoder.Decode(&u)
	if err != nil || &u == nil {
		errors.Write(w, errors.NoJSONBody)
		return
	}

	err, auth := env.UserRepo().Authenticate(&u)
	if err != nil {
		errors.Write(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auth)
}
