package cmd

import (
	"fmt"
	"log"
	"sync"
)

func RunExecute(urls []string) {
	numJob := len(urls)
	numWorker := 3
	jobs := make(chan Job, numJob)
	results := make(chan Result, numJob)
	var wg sync.WaitGroup

	for i := 1; i <= numWorker; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for index, url := range urls {
		job := Job{
			ID:  index,
			URL: url,
		}

		jobs <- job
	}

	close(jobs)
	wg.Wait()
	close(results)

	fmt.Println("\n-+-+-+-Final Report-+-+-+-")
	for res := range results {
		if res.Error != nil {
			log.Fatalf("Job ID: %d | URL: %v | (FAILED): %v\n", res.JobID, res.URL, res.Error)
		} else {
			fmt.Printf("Job ID: %d | URL: %v | (SUCCESS) Status Code: %v\n", res.JobID, res.URL, res.StatusCode)
		}
	}
}
