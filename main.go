package main

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log"

	"git.soma.salesforce.com/jusong-chen/concurrency/pkg/demoapp"
	"github.com/kelseyhightower/envconfig"
)

const ()

type config struct {
	NumJobs    int `required:"true" envconfig:"NUM_JOBS"`
	NumWorkers int `required:"true" envconfig:"NUM_WORKERS"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	cfg := config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Printf("get environment variables:%v\n", err)
		return
	}
	run(cfg.NumWorkers, cfg.NumJobs)
}

func run(numWorkers, numJobs int) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	app, err := demoapp.New(numWorkers, numJobs)
	if err != nil {
		log.Fatalf("demo app:%v", err)
		return
	}
	<-interrupt
	log.Printf("\nGot SIGINT or SIGTERM, shutting down http server ...\n")
	app.Close()
}
