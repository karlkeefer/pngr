package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/karlkeefer/pngr/golang/db"
	"github.com/karlkeefer/pngr/golang/routes"
)

const port = ":5000"

func main() {
	go handleSignals()

	// connect to postgres
	_ = db.DB()
	log.Println("Successfully connected to postgres")

	log.Println("Listening on", port)
	err := http.ListenAndServe(port, routes.Configure())

	if err != nil {
		log.Fatalln("ListenAndServe error:", err)
	}
}

func handleSignals() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	exitChan := make(chan int)
	go func() {
		for {
			s := <-signalChan
			switch s {
			case syscall.SIGHUP,
				syscall.SIGINT,
				syscall.SIGTERM,
				syscall.SIGQUIT:
				log.Println("Received shutdown signal:", s)
				exitChan <- 0

			default:
				log.Println("Unknown signal!?", s)
				exitChan <- 1
			}
		}
	}()

	code := <-exitChan
	os.Exit(code)
}
