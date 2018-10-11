package handlers

import (
	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models/user"
	"github.com/karlkeefer/pngr/golang/utils"

	"encoding/json"
	"net/http"
)

func Session(env *env.Env, w http.ResponseWriter, r *http.Request) (http.HandlerFunc, error) {
	switch r.Method {
	case "POST":
		return login(env, w, r), nil
	case "DELETE":
		return logout(env, w, r), nil
	default:
		return nil, errors.BadRequestMethod
	}
}

func login(env *env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		u := &user.User{}
		err := decoder.Decode(u)
		if err != nil || &u == nil {
			errors.Write(w, errors.NoJSONBody)
			return
		}

		u, err = env.UserRepo().Authenticate(u)
		if err != nil {
			errors.Write(w, err)
			return
		}

		utils.SetCookieForUser(w, u)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
	}
}

func logout(env *env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := &user.User{}
		utils.SetCookieForUser(w, u)
		w.WriteHeader(http.StatusOK)
	}
}
