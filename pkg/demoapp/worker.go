package demoapp

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Job describes job info
type Job struct {
	ID string
}

//Worker simulates a worker
type Worker struct {
	ID int
}

//Run start processing and quit when all job is done or a ctx Done signal received
func (w Worker) Run(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job) {
	defer func() {
		fmt.Printf("worker(ID=%d) quits\n", w.ID)
		wg.Done()
	}()

	fmt.Printf("worker(ID=%d) is ready ...\n", w.ID)
	for {
		select {
		case <-ctx.Done():
			return
		case j, ok := <-jobs:
			if !ok {
				time.Sleep(time.Microsecond)
				continue
			}
			w.process(&j)
		}

	}
}

func (w Worker) process(j *Job) {
	fmt.Printf("worker(ID=%d) start working on %s\n", w.ID, j.ID)
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Printf("worker(ID=%d) completed  %s\n", w.ID, j.ID)
}
