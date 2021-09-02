package handlers

import (
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"

	"github.com/karlkeefer/pngr/golang/db"
	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/handlers/write"
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
func NewServer() (*server, error) {
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

	// RESETS
	srv.POST("/reset", CreateReset)
	srv.GET("/reset/:code", DoReset)

	// USER
	srv.POST("/user", Signup)
	srv.GET("/user", Whoami)
	srv.POST("/user/verify", Verify)
	srv.PUT("/user/password", UpdatePassword)

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
	head, r.URL.Path = shiftPath(r.URL.Path)
	if head != "api" {
		write.Error(errors.RouteNotFound)
	}

	srv.router.ServeHTTP(w, r)
}

func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

// srvHandler is the extended handler function that our API routes use
type srvHandler func(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc

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

// helpers for easily parsing params
func getInt64(name string, r *http.Request) (out int64, err error) {
	params := httprouter.ParamsFromContext(r.Context())
	arg := params.ByName(name)
	out, err = strconv.ParseInt(arg, 10, 64)
	return
}

func getID(r *http.Request) (out int64, err error) {
	return getInt64("id", r)
}

func getString(name string, r *http.Request) (param string) {
	params := httprouter.ParamsFromContext(r.Context())
	return params.ByName(name)
}
