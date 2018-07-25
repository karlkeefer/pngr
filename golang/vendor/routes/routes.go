package routes

import (
	"fs"
	"net/http"
	"routes/login"
	"routes/logout"
	"routes/signup"
)

func Configure() {
	http.HandleFunc("/api/login", login.Handler)
	http.HandleFunc("/api/logout", logout.Handler)
	http.HandleFunc("/api/signup", signup.Handler)

	// built front-end and static files get copied into the docker
	// container during the production build process
	fs := http.FileServer(http.Dir("/root/front"))
	http.Handle("/", fs)
}
