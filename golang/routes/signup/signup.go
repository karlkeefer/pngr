package signup

import (
	"encoding/json"
	"log"
	"net/http"
)

type signupRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type signupResponse struct {
	Session string `json:"session"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t signupRequest
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	// TODO: check for dupes
	// TODO: don't actually reveal error in this case as people can fish for valid accounts
	if t.Id == "already@exists.com" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&errorResponse{
			Error: "Invalid Email",
		})
		return
	}

	log.Println("signing up:", t.Id, t.Password)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&signupResponse{
		Session: "placeholder",
	})
}
