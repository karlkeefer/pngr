package verify

import (
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models/user"
	"github.com/karlkeefer/pngr/golang/utils"

	"encoding/json"
	"net/http"
)

type verifyResponse struct {
	Msg string
}

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)
	if head == "" {
		errors.Write(w, errors.RouteNotFound)
		return
	}

	switch r.Method {
	case "GET":
		h.handleGet(head)(w, r)
	default:
		errors.Write(w, errors.BadRequestMethod)
	}
}

func (h *Handler) handleGet(code string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err, auth := user.Verify(code)
		if err != nil {
			errors.Write(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(auth)
	}
}
