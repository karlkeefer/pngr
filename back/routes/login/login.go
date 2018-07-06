package login

import (
    "encoding/json"
    "log"
    "net/http"
)

type loginRequest struct {
    Id       string `json:"id"`
    Password string `json:"password"`
}

type loginResponse struct {
    Session string `json:"session"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var t loginRequest
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }

    log.Println("Logging in:", t.Id, t.Password)

    resp := &loginResponse{
        Session: "placeholder",
    }

    w.Header().Set("Content-Type", "application/json")
    encoder := json.NewEncoder(w)
    encoder.Encode(resp)
}
