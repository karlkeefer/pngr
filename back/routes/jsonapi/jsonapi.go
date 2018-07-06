package jsonapi

import (
  "encoding/json"
  "errors"
  "net/http"
)

type Error struct {
  Error      error
  StatusCode int
}

type errorResponse struct {
  Error string `json:"error"`
}

// TODO: reformulate golang errors as consts
var (
  ErrorAccountAlreadyExists = Error{errors.New("An account with this identifier already exists"), http.StatusBadRequest}
)

func ReturnError(w http.ResponseWriter, err Error) {
  w.WriteHeader(err.StatusCode)
  w.Header().Set("Content-Type", "application/json")
  encoder := json.NewEncoder(w)
  encoder.Encode(&errorResponse{
    Error: err.Error.Error(),
  })
}

func ReturnSuccess(w http.ResponseWriter, resp interface{}) {
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")
  encoder := json.NewEncoder(w)
  encoder.Encode(resp)
}
