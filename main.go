package main

import (
	"context"
	"math/rand"
	"time"

	"log"

	"github.com/jusongchen/concurrencyDemo/pkg/demoapp"
	"github.com/kelseyhightower/envconfig"
)

const ()

type config struct {
	NumJobs             int `required:"true" envconfig:"NUM_JOBS"`
	DegreeOfConcurrency int `required:"true" envconfig:"DEGREE_CONCURRENCY"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	cfg := config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Printf("get environment variables:%v\n", err)
		return
	}

	demoapp.Run(context.Background(), cfg.DegreeOfConcurrency, cfg.NumJobs)
}
