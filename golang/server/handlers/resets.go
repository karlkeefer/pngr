package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/karlkeefer/pngr/golang/db"
	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/server/jwt"
	"github.com/karlkeefer/pngr/golang/server/write"
)

type createResetRequest struct {
	Email string `json:"email"`
}

func CreateReset(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	req := &createResetRequest{}
	err := decoder.Decode(req)
	if err != nil || req == nil {
		return write.Error(errors.NoJSONBody)
	}

	u, err := env.DB().FindUserByEmail(r.Context(), req.Email)
	if err != nil {
		if isNotFound(err) {
			return write.Error(errors.InvalidEmail)
		}
		return write.Error(err)
	}

	reset, err := env.DB().CreateReset(r.Context(), db.CreateResetParams{
		UserID: u.ID,
		Code:   generateRandomString(32),
	})
	if err != nil {
		return write.Error(err)
	}

	err = env.Mailer().ResetPassword(u.Email, reset.Code)
	if err != nil {
		return write.Error(err)
	}

	return write.Success()
}

func DoReset(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	code := getString("code", r)

	reset, err := env.DB().FindResetByCode(r.Context(), code)
	if err != nil {
		if isNotFound(err) {
			return write.Error(errors.ResetNotFound)
		}
		return write.Error(err)
	}

	u, err := env.DB().FindUserByID(r.Context(), reset.UserID)
	if err != nil {
		return write.Error(err)
	}

	// clean up resets from DB
	err = env.DB().DeleteResetsForUser(r.Context(), reset.UserID)
	if err != nil {
		return write.Error(err)
	}

	jwt.WriteUserCookie(w, &u)
	return write.JSON(&u)
}
