package main

import (
    "./routes"
    "log"
    "net/http"
)

const port = ":3000"

func main() {
    routes.Configure()

    // static files get copied into the docker container during the build process
    fs := http.FileServer(http.Dir("/root/front"))
  	http.Handle("/", fs)

  	// TODO: add special auth handler for admin js bundle after code-splitting

    log.Println("Listening on port", port)
    err := http.ListenAndServe(port, nil)

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
