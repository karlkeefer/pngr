package server

import (
	"encoding/json"
	"net/http"

	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models"
	"github.com/karlkeefer/pngr/golang/server/jwt"
	"github.com/karlkeefer/pngr/golang/server/write"
)

func Login(env env.Env, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
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

func Logout(env env.Env, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	u := &models.User{}
	jwt.WriteUserCookie(w, u)
	return write.JSON(&logoutResponse{true})
}
