package server

import (
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/julienschmidt/httprouter"

	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models"
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
	env    env.Env
	router *httprouter.Router
}

// New initializes env (database connections and whatnot)
// and creates a server that implements ServeHTTP
func New() (*server, error) {
	env, err := env.New()
	if err != nil {
		return nil, err
	}

	router := httprouter.New()

	// setup error handlers for our router
	router.MethodNotAllowed = write.Error(errors.BadRequestMethod)
	router.NotFound = write.Error(errors.RouteNotFound)
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		log.Println("Panic on", r.URL.Path)
		debug.PrintStack()
		write.Error(errors.InternalError)(w, r)
	}

	srv := &server{
		env:    env,
		router: router,
	}

	// SESSION
	srv.POST("/session", Login)
	srv.DELETE("/session", Logout)

	// USER
	srv.POST("/user", Signup)
	srv.GET("/user", Whoami)
	srv.POST("/user/verify", Verify)

	// POSTS
	srv.GET("/posts", GetPosts)
	srv.GET("/posts/:id", GetPost)
	srv.POST("/posts", CreatePost)
	srv.PUT("/posts", UpdatePost)
	srv.DELETE("/posts/:id", DeletePost)

	return srv, nil
}

// ServeHTTP handles routes that start with api
func (srv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	// shift head and tail to get below "api/" part of the path
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)
	if head != "api" {
		write.Error(errors.RouteNotFound)
	}

	srv.router.ServeHTTP(w, r)
}

// srvHandler is the extended handler function that our API routes use
type srvHandler func(env env.Env, user *models.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc

// helpers for easily adding routes
func (srv *server) GET(path string, handler srvHandler) {
	srv.router.HandlerFunc(http.MethodGet, path, srv.wrap(handler))
}
func (srv *server) PUT(path string, handler srvHandler) {
	srv.router.HandlerFunc(http.MethodPut, path, srv.wrap(handler))
}
func (srv *server) POST(path string, handler srvHandler) {
	srv.router.HandlerFunc(http.MethodPost, path, srv.wrap(handler))
}
func (srv *server) DELETE(path string, handler srvHandler) {
	srv.router.HandlerFunc(http.MethodDelete, path, srv.wrap(handler))
}
