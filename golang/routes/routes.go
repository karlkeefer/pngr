package routes

import (
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/utils"

	"github.com/karlkeefer/pngr/golang/routes/login"
	"github.com/karlkeefer/pngr/golang/routes/signup"
	"github.com/karlkeefer/pngr/golang/routes/verify"

	"net/http"
)

type Server struct {
	fs  http.Handler
	api *apiHandler
}

func (h *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	head, tail := utils.ShiftPath(r.URL.Path)
	if head == "api" {
		r.URL.Path = tail
		h.api.ServeHTTP(w, r)
	} else {
		h.fs.ServeHTTP(w, r)
	}
}

type apiHandler struct {
	loginHandler  *login.Handler
	signupHandler *signup.Handler
	verifyHandler *verify.Handler
}

func (h *apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)
	switch head {
	case "login":
		h.loginHandler.ServeHTTP(w, r)
	case "signup":
		h.signupHandler.ServeHTTP(w, r)
	case "verify":
		h.verifyHandler.ServeHTTP(w, r)
	default:
		errors.Write(w, errors.RouteNotFound)
	}
}

func Configure() *Server {
	return &Server{
		// built front-end and static files get copied into the docker
		// container during the production build process
		fs: http.FileServer(http.Dir("/root/front")),
		api: &apiHandler{
			loginHandler:  &login.Handler{},
			signupHandler: &signup.Handler{},
			verifyHandler: &verify.Handler{},
		},
	}
}
