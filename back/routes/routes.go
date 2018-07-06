package routes

import (
  "./login"
  "./logout"
  "./signup"
  "net/http"
)

func Configure() {
  http.HandleFunc("/api/login", login.Handler)
  http.HandleFunc("/api/logout", logout.Handler)
  http.HandleFunc("/api/signup", signup.Handler)
}
