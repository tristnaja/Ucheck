package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"sync"
)

type Database struct {
	Size int      `json:"size"`
	URLs []string `json:"urls"`
}

func RunCheck(filePath string) error {
	db, err := readJSON(filePath)

	if err != nil {
		return err
	}

	numJob := len(db.URLs)
	numWorker := 3
	jobs := make(chan Job, numJob)
	results := make(chan Result, numJob)
	var wg sync.WaitGroup

	for i := 1; i <= numWorker; i++ {
		wg.Add(1)
		go Worker(i, jobs, results, &wg)
	}

	for index, link := range db.URLs {
		parsedURL, err := url.Parse(link)

		if parsedURL.Scheme == "" {
			parsedURL.Scheme = "https"
		}

		if err != nil {
			continue
		}

		job := Job{
			ID:  index,
			URL: parsedURL.String(),
		}

		jobs <- job
	}

	close(jobs)
	wg.Wait()
	close(results)

	fmt.Println(db.Size)
	fmt.Println(len(db.URLs))

	fmt.Println("\n-+-+-+-Final Report-+-+-+-")
	for res := range results {
		if res.Error != nil {
			fmt.Fprintf(os.Stderr, "Job ID: %d | URL: %v | (FAILED): %v\n", res.JobID, res.URL, res.Error)
		} else {
			fmt.Printf("Job ID: %d | URL: %v | (SUCCESS) Status Code: %v\n", res.JobID, res.URL, res.StatusCode)
		}
	}

	err = CleanUp(filePath)

	if err != nil {
		return fmt.Errorf("cleaning up: %w", err)
	}

	return nil
}

func ReadDatabase(filePath string) (*Database, error) {
	db, err := readJSON(filePath)

	if err != nil {
		return nil, fmt.Errorf("opening database: %w", err)
	}

	return db, nil
}

func AddURL(url string, filePath string) error {
	dir := filepath.Dir(filePath)

	err := os.MkdirAll(dir, 0755)

	if err != nil {
		return fmt.Errorf("creating database dir: %w", err)
	}

	db, err := readJSON(filePath)

	if err != nil {
		return fmt.Errorf("opening database: %w", err)
	}

	db.URLs = append(db.URLs, url)
	db.Size++

	err = writeJSON(filePath, *db)

	if err != nil {
		return fmt.Errorf("saving url: %w", err)
	}

	return nil
}

func CleanUp(filePath string) error {
	db, err := readJSON(filePath)

	if err != nil {
		return fmt.Errorf("something happenend during cleanup: %w", err)
	}

	db.URLs = []string{`add your link with "ucheck add -u <url>"`}
	db.Size = 0

	err = writeJSON(filePath, *db)

	if err != nil {
		return fmt.Errorf("something happenend during cleanup: %w", err)
	}

	return nil
}

func writeJSON(filePath string, db Database) error {
	file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_RDWR, 0644)

	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")

	err = encoder.Encode(db)

	if err != nil {
		return fmt.Errorf("encoding data: %w", err)
	}

	return nil
}

func readJSON(filePath string) (*Database, error) {
	var result Database

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDONLY, 0644)

	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&result)

	if err != nil {
		if errors.Is(err, io.EOF) {
			result = Database{}
		} else {
			return nil, fmt.Errorf("decoding data: %w", err)
		}
	}

	return &result, nil
}
