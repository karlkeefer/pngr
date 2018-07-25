package logout

import (
	"encoding/json"
	"log"
	"net/http"
)

type logoutRequest struct {
	Session string `json:"session"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t logoutRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	log.Println("Logging out", t.Session)
	w.WriteHeader(http.StatusOK)
}
