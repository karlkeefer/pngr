package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/karlkeefer/pngr/golang/db"
	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/handlers/jwt"
	"github.com/karlkeefer/pngr/golang/handlers/write"
)

func Login(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	u := db.User{}
	err := decoder.Decode(u)
	if err != nil || &u == nil {
		return write.Error(errors.NoJSONBody)
	}

	u, err = env.DB().FindUserByEmail(r.Context(), u.Email)
	if err != nil {
		return write.Error(err)
	}

	if !checkPasswordHash(u.Pass, u.Salt, u.Pass) {
		return write.Error(errors.FailedLogin)
	}

	jwt.WriteUserCookie(w, &u)
	return write.JSON(u)
}

type logoutResponse struct {
	success bool
}

func Logout(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	u := &db.User{}
	jwt.WriteUserCookie(w, u)
	return write.JSON(&logoutResponse{true})
}
