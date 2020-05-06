package demoapp

import (
	"context"
	"fmt"
	"strconv"
	"sync"
)

//App implements a server
type App struct {
	numWorker int
	cancelFn  context.CancelFunc
	wg        *sync.WaitGroup
}

//New starts a app instance
func New(numWorker int, numJob int) (*App, error) {

	app := App{
		numWorker: numWorker,
		wg:        &sync.WaitGroup{},
	}

	jobs := make(chan Job)

	ctx, cancelFn := context.WithCancel(context.Background())
	app.cancelFn = cancelFn

	go generateJobs(ctx, numJob, jobs)

	for i := 0; i < app.numWorker; i++ {
		go func(workerID int) {
			w := Worker{ID: workerID}
			app.wg.Add(1)
			w.Run(ctx, app.wg, jobs)
		}(i)
	}

	return &app, nil
}

//Close shut down the server
func (app *App) Close() {
	fmt.Printf("start to shut down the server ...\n")
	app.cancelFn()

	fmt.Printf("waiting for all worker to quit ...\n")
	app.wg.Wait()

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
