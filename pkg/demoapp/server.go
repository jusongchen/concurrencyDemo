package demoapp

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

//App implements a server

//Run starts a app instance
func Run(ctx context.Context, numWorker int, numJob int) {

	wg := &sync.WaitGroup{}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	jobs := make(chan Job)
	go generateJobs(ctx, numJob, jobs)

	for i := 0; i < numWorker; i++ {
		wg.Add(1)
		go func(workerID int) {
			w := Worker{ID: workerID}
			w.Run(ctx, wg, jobs)
		}(i)
	}

	fmt.Printf("waiting for all worker to quit ...\n")
	wg.Wait()
	fmt.Printf("server shut down.\n")
}

func generateJobs(ctx context.Context, numJob int, jobs chan Job) {

	defer close(jobs)

	for i := 0; i < numJob; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			jobs <- Job{
				ID: "job-" + strconv.Itoa(i),
			}

		}
	}

}
