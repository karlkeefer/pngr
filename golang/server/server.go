package server

import (
	"net/http"
	"os"

	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/server/handlers/posts"
	"github.com/karlkeefer/pngr/golang/server/handlers/session"
	"github.com/karlkeefer/pngr/golang/server/handlers/user"
	"github.com/karlkeefer/pngr/golang/utils"
)

type server struct {
	env *env.Env
	fs  http.Handler
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
		// built front-end and static files get copied into the docker
		// container during the production build process
		fs: http.FileServer(http.Dir("/root/front")),
	}, nil
}

// ServeHTTP forks API traffic from static asset traffic
func (srv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	head, tail := utils.ShiftPath(r.URL.Path)
	if head == "api" {
		r.URL.Path = tail
		srv.ServeAPI(w, r)
	} else {
		srv.fs.ServeHTTP(w, r)
	}
}

// ServeAPI handles all of the API routes
func (srv *server) ServeAPI(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)

	var handler http.HandlerFunc
	var err error

	switch head {
	case "session":
		handler, err = session.Handler(srv.env, w, r)
	case "user":
		handler, err = user.Handler(srv.env, w, r)
	case "posts":
		handler, err = posts.Handler(srv.env, w, r)
	default:
		err = errors.RouteNotFound
	}

	if err != nil {
		errors.Write(w, err)
		return
	}

	// TODO: consider a middleware wrapper utility
	wrappedHandler := csrf(cors(handler))
	wrappedHandler(w, r)
}

// MIDDLEWARE

// csrf checks for the CSRF prevention header and compares the origin header
func csrf(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Origin") != os.Getenv("APP_ROOT") {
			errors.Write(w, errors.BadOrigin)
			return
		}
		if r.Header.Get("X-Requested-With") != "XMLHttpRequest" {
			errors.Write(w, errors.BadCSRF)
			return
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
