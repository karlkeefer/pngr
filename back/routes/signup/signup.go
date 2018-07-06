package signup

import (
    "../jsonapi"
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

func Handler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var t signupRequest
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }

    if t.Id == "already@exists.com" {
        jsonapi.ReturnError(w, jsonapi.ErrorAccountAlreadyExists)
        return
    }

    log.Println("signing up:", t.Id, t.Password)

    jsonapi.ReturnSuccess(w, &signupResponse{
        Session: "placeholder",
    })
}
