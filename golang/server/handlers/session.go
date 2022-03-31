package handlers

import (
	"encoding/json"
	"net/http"

	mdb "github.com/QuinnMain/infograph/golang/db"
	"github.com/QuinnMain/infograph/golang/env"
	"github.com/QuinnMain/infograph/golang/errors"
	"github.com/QuinnMain/infograph/golang/server/jwt"
	"github.com/QuinnMain/infograph/golang/server/write"
)

type loginRequest struct {
	Email string
	Pass  string
}

func Login(env env.Env, user *mdb.MUser, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	req := loginRequest{}
	err := decoder.Decode(&req)
	if err != nil || &req == nil {
		return write.Error(errors.NoJSONBody)
	}

	u, err := env.DB().FindUserByEmail(r.Context(), req.Email)
	if err != nil {
		if isNotFound(err) {
			return write.Error(errors.FailedLogin)
		}
		return write.Error(err)
	}

	if !checkPasswordHash(req.Pass, u.Salt, u.Pass) {
		return write.Error(errors.FailedLogin)
	}

	jwt.WriteUserCookie(w, &u)
	return write.JSON(&u)
}

func Logout(env env.Env, user *mdb.MUser, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	u := &mdb.MUser{}
	jwt.WriteUserCookie(w, u)
	return write.Success()
}

func LoginMon()
