package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	// "github.com/jmoiron/sqlx"
	"github.com/karlkeefer/pngr/golang/routes"
)

const port = ":5050"

func getConnectionString() string {
	return fmt.Sprintf("postgresql://%s@%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_DB"))
}

func main() {
	go handleSignals()

	routes.Configure()

	// connectionString := getConnectionString()
	// log.Println("Connecting to", connectionString)
	// db, err := sqlx.Open("postgres", connectionString)
	// if err != nil {
	// 	log.Fatal("sqlx.open failed")
	// }

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal("Unable to connect to postgres")
	// }

	// built front-end and static files get copied into the docker
	// container during the build process
	fs := http.FileServer(http.Dir("/root/front"))
	http.Handle("/", fs)

	// TODO: add special auth handler for admin js bundle after code-splitting?

	log.Println("Listening on", port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
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
			case syscall.SIGHUP:
			case syscall.SIGINT:
			case syscall.SIGTERM:
			case syscall.SIGQUIT:
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
