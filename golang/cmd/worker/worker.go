package main

import (
	"log"

	"github.com/karlkeefer/pngr/golang/worker"
)

func main() {
	wrk, err := worker.New()
	if err != nil {
		log.Fatalln("Unable to initialize worker", err)
	}

	log.Println("Starting worker...")
	wrk.Run()
}
