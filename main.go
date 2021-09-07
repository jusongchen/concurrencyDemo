package main

import (
	"context"
	"math/rand"
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

	demoapp.Run(context.Background(), cfg.NumWorkers, cfg.NumJobs)
}
