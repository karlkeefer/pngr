package server

import (
	"net/http"
	"os"
	"time"

	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/server/handlers/posts"
	"github.com/karlkeefer/pngr/golang/server/handlers/session"
	"github.com/karlkeefer/pngr/golang/server/handlers/user"
	"github.com/karlkeefer/pngr/golang/server/jwt"
	"github.com/karlkeefer/pngr/golang/server/write"
	"github.com/karlkeefer/pngr/golang/utils"
)

var isDev = false
var appRoot = ""

const localDev = "https://localhost"

func init() {
	isDev = os.Getenv("ENV") == "dev"
	appRoot = os.Getenv("APP_ROOT")
}

type server struct {
	env env.Env
}

// New initializes env (database connections and whatnot)
// and creates a server that implements ServeHTTP
func New() (*server, error) {
	env, err := env.New()
	if err != nil {
		return nil, err
	}

	return &server{
		env: env,
	}, nil
}

// ServeHTTP forks API traffic from static asset traffic
func (srv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	originalPath := r.URL.Path
	handler := srv.getHandler(w, r)
	// TODO: consider a middleware wrapper utility
	wrappedHandler := lag(csrf(originalPath, cors(originalPath, handler)))
	wrappedHandler(w, r)
}

// getHandler parses the request path to return a relevant handler
// sub-sections of the API have their own Handler() func which handles route parsing for that piece of the API
func (srv *server) getHandler(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	var head string
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)
	if head != "api" {
		return write.Error(errors.RouteNotFound)
	}

	// shift head and tail to get below "api/" part of the path
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)

	// handle session routes before cookie parsing
	switch head {
	case "session":
		return session.Handler(srv.env, w, r)
	}

	// read user off cookie
	u, err := jwt.HandleUserCookie(srv.env, w, r)
	if err != nil {
		return write.Error(err)
	}

	switch head {
	case "user":
		return user.Handler(srv.env, w, r, u)
	case "posts":
		return posts.Handler(srv.env, w, r, u)
	default:
		return write.Error(errors.RouteNotFound)
	}
}

// MIDDLEWARE

// lag allows you to simiulate API lag locally to test "loading" states
func lag(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isDev {
			// minimal lag so that we can see loading states at least briefly
			time.Sleep(time.Millisecond * 500)
		}
		fn(w, r)
	}
}

// only returns an origin if it matches our list
func validateOrigin(r *http.Request) string {
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
func csrf(path string, fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if skipCorsAndCSRF(path) {
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
func cors(path string, fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if skipCorsAndCSRF(path) {
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
