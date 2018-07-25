package routes

import (
	"github.com/karlkeefer/pngr/golang/routes/login"
	"github.com/karlkeefer/pngr/golang/routes/logout"
	"github.com/karlkeefer/pngr/golang/routes/signup"
	"net/http"
)

func Configure() {
	http.HandleFunc("/api/login", login.Handler)
	http.HandleFunc("/api/logout", logout.Handler)
	http.HandleFunc("/api/signup", signup.Handler)
}
