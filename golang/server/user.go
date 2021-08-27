package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models"
	"github.com/karlkeefer/pngr/golang/server/jwt"
	"github.com/karlkeefer/pngr/golang/server/write"
)

type signupResponse struct {
	URL string
}

func Signup(env env.Env, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	var u models.User
	err := decoder.Decode(&u)
	if err != nil || &u == nil {
		return write.Error(errors.NoJSONBody)
	}

	dbUser, err := env.UserRepo().Signup(&u)
	if err != nil {
		return write.Error(err)
	}

	// TODO: this is where we should actually email the code with mailgun or something
	// for now we just pass verification code back in the response...
	return write.JSON(&signupResponse{
		URL: fmt.Sprintf("%s/verify/%s", os.Getenv("APP_ROOT"), dbUser.Verification),
	})
}

func Whoami(env env.Env, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return write.JSON(user)
}

type verifyRequest struct {
	Code string
}

func Verify(env env.Env, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
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
