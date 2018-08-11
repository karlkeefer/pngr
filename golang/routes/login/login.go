package login

import (
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models/user"

	"encoding/json"
	"net/http"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.handlePost(w, r)
	default:
		errors.Write(w, errors.BadRequestMethod)
	}
}

func (h *Handler) handlePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u user.User
	err := decoder.Decode(&u)
	if err != nil || &u == nil {
		errors.Write(w, errors.NoJSONBody)
		return
	}

	err, auth := user.Authenticate(&u)
	if err != nil {
		errors.Write(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auth)
}
