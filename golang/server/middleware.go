package server

import (
	"net/http"
	"os"
	"time"

	"github.com/karlkeefer/pngr/golang/db"
	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/server/jwt"
	"github.com/karlkeefer/pngr/golang/server/write"
)

// withUserAndEnv populates our custom srvHandler args for our route handlers
func withUserAndEnv(env env.Env, h srvHandler, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	head, _ := shiftPath(r.URL.Path)

	var user *db.User

	// don't parse user cookie on session routes!
	if head != "session" {
		var err error
		user, err = jwt.HandleUserCookie(env, w, r)
		if err != nil {
			write.Error(err)
		}
	}

	return h(env, user, w, r)
}

// lag allows you to simiulate API lag locally to test "loading" states
func lag(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isDev {
			time.Sleep(time.Millisecond * 500)
		}
		fn(w, r)
	}
}

const localDev = "http://localhost:4500"

// only returns an origin if it matches our list
func validateOrigin(r *http.Request) string {
	appRoot := os.Getenv("APP_ROOT")
	switch r.Header.Get("Origin") {
	case appRoot:
		return appRoot
	case localDev:
		return localDev
	default:
		return ""
	}
}

// csrf checks for the CSRF prevention header and compares the origin header
func csrf(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if skipCorsAndCSRF(r.URL.Path) {
			fn(w, r)
			return
		}

		if r.Method != http.MethodOptions {
			if r.Header.Get("Origin") != "" && validateOrigin(r) == "" {
				// if an origin is provided, but didn't match our list
				fn = write.Error(errors.BadOrigin)
			} else if r.Header.Get("X-Requested-With") != "XMLHttpRequest" {
				fn = write.Error(errors.BadCSRF)
			}
		}
		fn(w, r)
	}
}

// cors adds CORS headers to the response
func cors(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if skipCorsAndCSRF(r.URL.Path) {
			fn(w, r)
			return
		}

		if origin := validateOrigin(r); origin != "" {
			// if we were given an origin that matches our list
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")

		if r.Method == http.MethodOptions {
			// simple response for the preflight check
			fn = write.Success()
		}
		fn(w, r)
	}
}

// a list of paths to bypass cors checks - this is useful for webhooks and stuff
var bypassPaths = []string{}

func skipCorsAndCSRF(path string) bool {
	for _, c := range bypassPaths {
		if path == c {
			return true
		}
	}

	return false
}
