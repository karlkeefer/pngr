package main

import (
    "./routes"
    "log"
    "net/http"
)

const port = ":5000"

func main() {
    routes.Configure()

    log.Println("Listening on port", port)
    err := http.ListenAndServe(port, nil)

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
