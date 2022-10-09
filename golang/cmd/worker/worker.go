package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/karlkeefer/pngr/golang/env"
)

func main() {
	wrk, err := New()
	if err != nil {
		log.Fatalln("Unable to initialize worker", err)
	}

	wrk.Run()
}

type Worker struct {
	env        env.Env
	wg         sync.WaitGroup
	signalChan chan os.Signal
	exitChan   chan int
}

func New() (*Worker, error) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	env, err := env.New()
	if err != nil {
		return nil, err
	}

	return &Worker{
		env,
		sync.WaitGroup{},
		signalChan,
		make(chan int),
	}, nil
}

func (w *Worker) Run() {

	// listen ticks on the timer, or exit signals
	go func() {
		for s := range w.signalChan {
			w.handleExit(s)
		}
	}()

	log.Println("Worker started")

	// TODO: Grab some tasks from a queue and process them with w.doAsync()
	// Your implementation should consider a worker pool

	// block waiting for exit signal
	code := <-w.exitChan
	os.Exit(code)
}

func (w *Worker) handleExit(s os.Signal) {
	// log.Println("Handling os signal...")

	w.wg.Wait()
	// log.Println("Goroutines finished up")

	switch s {
	case syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT:
		log.Println("Received shutdown signal:", s)
		w.exitChan <- 0

	default:
		log.Println("Unknown signal!?", s)
		w.exitChan <- 1
	}
}

func (w *Worker) doAsync(fn func() error) {
	w.wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		err := fn()
		if err != nil {
			log.Println(err)
		}
	}(&w.wg)
}
