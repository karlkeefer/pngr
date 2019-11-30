package server

import (
	"net/http"
	"os"

	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/server/handlers/posts"
	"github.com/karlkeefer/pngr/golang/server/handlers/session"
	"github.com/karlkeefer/pngr/golang/server/handlers/user"
	"github.com/karlkeefer/pngr/golang/server/write"
	"github.com/karlkeefer/pngr/golang/utils"
)

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
	handler := srv.getHandler(w, r)
	// TODO: consider a middleware wrapper utility
	wrappedHandler := lag(csrf(cors(handler)))
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

	switch head {
	case "session":
		return session.Handler(srv.env, w, r)
	case "user":
		return user.Handler(srv.env, w, r)
	case "posts":
		return posts.Handler(srv.env, w, r)
	default:
		return write.Error(errors.RouteNotFound)
	}
}

// MIDDLEWARE

// lag allows you to simiulate API lag locally to test "loading" states
// just uncomment the time.Sleep to see it in action
func lag(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// time.Sleep(time.Millisecond * 1000)
		fn(w, r)
	}
}

// csrf checks for the CSRF prevention header and compares the origin header
func csrf(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Origin") != "" && r.Header.Get("Origin") != os.Getenv("APP_ROOT") {
			fn = write.Error(errors.BadOrigin)
		} else if r.Header.Get("X-Requested-With") != "XMLHttpRequest" {
			fn = write.Error(errors.BadCSRF)
		}
		fn(w, r)
	}
}

// cors adds CORS headers to the response
// the current setup does not allow any cross-origin requests!
func cors(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("APP_ROOT"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")
		fn(w, r)
	}
}
