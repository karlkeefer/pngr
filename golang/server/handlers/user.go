package handlers

import (
	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models/user"
	"github.com/karlkeefer/pngr/golang/utils"

	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type signupResponse struct {
	URL string
}

func User(env *env.Env, w http.ResponseWriter, r *http.Request) (http.HandlerFunc, error) {
	head, _ := utils.ShiftPath(r.URL.Path)

	switch r.Method {
	case "POST":
		if head == "verify" {
			return verify(env, w, r), nil
		} else {
			return signup(env, w, r), nil
		}
	case "GET":
		return whoami(env, w, r), nil
	default:
		return nil, errors.BadRequestMethod
	}
}

func signup(env *env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var u user.User
		err := decoder.Decode(&u)
		if err != nil || &u == nil {
			errors.Write(w, errors.NoJSONBody)
			return
		}

		err = env.UserRepo().Signup(&u)
		if err != nil {
			errors.Write(w, err)
			return
		}

		// TODO: this is where we should actually email the code with mailgun or something
		// for now we just pass verification code back in the response...
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&signupResponse{
			URL: fmt.Sprintf("%s/verify/%s", os.Getenv("APP_ROOT"), u.Verification),
		})
	}
}

func whoami(env *env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("jwt")
		var jwt string
		if cookie != nil {
			jwt = cookie.Value
		}

		u, err := utils.FromToken(jwt)
		if err != nil {
			errors.Write(w, err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
	}
}

type verifyRequest struct {
	Code string
}

func verify(env *env.Env, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var req verifyRequest

		err := decoder.Decode(&req)
		if err != nil || &req == nil || req.Code == "" {
			errors.Write(w, errors.NoJSONBody)
			return
		}

		u, err := env.UserRepo().Verify(req.Code)
		if err != nil {
			errors.Write(w, err)
		}

		utils.SetCookieForUser(w, u)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
	}
}
