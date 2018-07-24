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
    // TODO: guard bad method
    decoder := json.NewDecoder(r.Body)
    var t loginRequest
    err := decoder.Decode(&t)
    if err != nil {
        log.Println("Invalid JSON", err)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    log.Println("Logging in:", t.Id, t.Password)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(&loginResponse{
        Session: "placeholder",
    })
}
