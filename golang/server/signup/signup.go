package signup

import (
	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models/user"

	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type signupResponse struct {
	URL string
}

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

	err = env.UserRepo().Signup(&u)
	if err != nil {
		errors.Write(w, err)
		return
	}

	// TODO: this is where we should actually email the code with mailgun or something
	// for now we just pass verification code back in the response...
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&signupResponse{
		URL: fmt.Sprintf("%s/verify/%s", os.Getenv("APP_ROOT"), u.Verification),
	})
}
