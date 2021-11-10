package output

import (
	"sync"
)

// Job represents a single job to be executed.
type Job struct {
	// The job's operation identifier.
	OperationID string
	// The job's generator.
	Generator string
	// the job's generator arguments.
	Args []interface{}
}

// Worker represents a single worker.
type Worker = func(wg *sync.WaitGroup, jobs <-chan Job, m FileMap)

// New executes the given jobs in parallel.
func New(jobs []Job, worker Worker, m FileMap) {
	wg := &sync.WaitGroup{}

	jobsChan := make(chan Job, 8)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(wg, jobsChan, m)
	}
	for _, job := range jobs {
		jobsChan <- job
	}
	close(jobsChan)

	wg.Wait()
}
