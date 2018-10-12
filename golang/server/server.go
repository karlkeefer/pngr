package server

import (
	"net/http"

	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/server/handlers/session"
	"github.com/karlkeefer/pngr/golang/server/handlers/user"
	"github.com/karlkeefer/pngr/golang/utils"
)

type server struct {
	env *env.Env
	fs  http.Handler
}

func (srv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	head, tail := utils.ShiftPath(r.URL.Path)
	if head == "api" {
		r.URL.Path = tail
		srv.ServeAPI(w, r)
	} else {
		srv.fs.ServeHTTP(w, r)
	}
}

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
	default:
		err = errors.RouteNotFound
	}

	if err != nil {
		errors.Write(w, err)
		return
	}

	// TODO: wrap with middleware for CORS, CSRF, etc.
	handler(w, r)
}

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
