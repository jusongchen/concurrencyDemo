package demoapp

import (
	"context"
	"fmt"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/jusongchen/concurrencyDemo/pkg/mandelbrot"
)

//Job describes job info
type Job struct {
	ID string
}

//Run starts a app instance
func Run(ctx context.Context, degreeOfConcurrency int, numJob int) {

	wg := &sync.WaitGroup{}

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	jobs := make(chan Job)
	go generateJobs(ctx, numJob, jobs)

	semaphore := make(chan int, degreeOfConcurrency)
	for job := range jobs {
		wg.Add(1)
		go DoJob(ctx, semaphore, wg, job)
	}

	fmt.Printf("waiting for all worker to complete their job ...\n")
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

//Run start processing and quit when all job is done or a ctx Done signal received
func DoJob(ctx context.Context, sem chan int, wg *sync.WaitGroup, job Job) {

	defer func() {
		wg.Done()
		<-sem
	}()

	sem <- 1
	select {
	case <-ctx.Done():
		fmt.Printf("quit signal detected, cancelling job%s\n", job.ID)
		return
	default:
		fmt.Printf("worker start working on %s\n", job.ID)
		// time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		mandelbrot.GenImage()

		fmt.Printf("job completed:  %s\n", job.ID)
	}
}
