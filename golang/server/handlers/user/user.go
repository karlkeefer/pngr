package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models/user"
	"github.com/karlkeefer/pngr/golang/server/jwt"
	"github.com/karlkeefer/pngr/golang/server/write"
	"github.com/karlkeefer/pngr/golang/utils"
)

func Handler(env *env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	var head string
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)
	switch r.Method {
	case "POST":
		if head == "verify" {
			return verify(env, w, r)
		} else {
			return signup(env, w, r)
		}
	case "GET":
		return whoami(env, w, r)
	default:
		return write.Error(errors.BadRequestMethod)
	}
}

type signupResponse struct {
	URL string
}

func signup(env *env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	var u user.User
	err := decoder.Decode(&u)
	if err != nil || &u == nil {
		return write.Error(errors.NoJSONBody)
	}

	fromDB, err := env.UserRepo().Signup(&u)
	if err != nil {
		return write.Error(err)
	}

	// TODO: this is where we should actually email the code with mailgun or something
	// for now we just pass verification code back in the response...
	return write.JSON(&signupResponse{
		URL: fmt.Sprintf("%s/verify/%s", os.Getenv("APP_ROOT"), fromDB.Verification),
	})
}

func whoami(env *env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return write.JSONorErr(jwt.HandleUserCookie(env, w, r))
}

type verifyRequest struct {
	Code string
}

func verify(env *env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	var req verifyRequest

	err := decoder.Decode(&req)
	if err != nil || &req == nil || req.Code == "" {
		return write.Error(errors.NoJSONBody)
	}

	u, err := env.UserRepo().Verify(req.Code)
	if err != nil {
		return write.Error(err)
	}

	jwt.WriteUserCookie(w, u)
	return write.JSON(u)
}
