package session

import (
	"encoding/json"
	"net/http"

	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models"
	"github.com/karlkeefer/pngr/golang/server/jwt"
	"github.com/karlkeefer/pngr/golang/server/write"
)

func Handler(env env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	switch r.Method {
	case http.MethodPost:
		return login(env, w, r)
	case http.MethodDelete:
		return logout(env, w)
	default:
		return write.Error(errors.BadRequestMethod)
	}
}

func login(env env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	u := &models.User{}
	err := decoder.Decode(u)
	if err != nil || &u == nil {
		return write.Error(errors.NoJSONBody)
	}

	u, err = env.UserRepo().Authenticate(u)
	if err != nil {
		return write.Error(err)
	}

	jwt.WriteUserCookie(w, u)
	return write.JSON(u)
}

type logoutResponse struct {
	success bool
}

func logout(env env.Env, w http.ResponseWriter) http.HandlerFunc {
	u := &models.User{}
	jwt.WriteUserCookie(w, u)
	return write.JSON(&logoutResponse{true})
}
