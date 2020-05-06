package main

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log"

	"git.soma.salesforce.com/jusong-chen/concurrency/pkg/demoapp"
)

const (
	numJobs    = 10
	numWorkers = 3
)

func main() {
	rand.Seed(time.Now().UnixNano())
	demo()
}

func demo() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	app, err := demoapp.New(numWorkers, numJobs)
	if err != nil {
		log.Fatalf("demo app:%v", err)
		return
	}
	select {
	case <-interrupt:
		log.Printf("\nGot SIGINT or SIGTERM, shutting down http server ...\n")
	}
	app.Close()
}
