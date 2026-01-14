package cmd

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Job struct {
	ID  int
	URL string
}

type Result struct {
	JobID      int
	StatusCode int
	URL        string
	Latency    time.Duration
	Error      error
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		start := time.Now()

		resp, err := http.Get(job.URL)

		duration := time.Since(start)

		res := Result{
			JobID:   job.ID,
			URL:     job.URL,
			Latency: duration,
		}

		if err != nil {
			res.Error = err
		} else {
			res.StatusCode = resp.StatusCode
			resp.Body.Close()
		}

		results <- res
		fmt.Printf("Worker %d: finished processing %v with Job ID of %d\n", id, job.URL, job.ID)
	}
}
