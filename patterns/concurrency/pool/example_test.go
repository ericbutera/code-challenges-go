package pool_test

// educational purposes only
// prod code should use something like https://github.com/sourcegraph/conc

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func worker(id int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		// Simulate work
		time.Sleep(time.Second)
		// Generate unique string result
		result := fmt.Sprintf("worker%d-job%d", id, j)
		results <- result
	}
}

func TestWorkerPoolWithStrings(t *testing.T) {
	t.Parallel()
	const (
		numJobs    = 5
		numWorkers = 3
	)
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	// Launch workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	expectedResults := make([]string, numJobs)
	for j := 1; j <= numJobs; j++ {
		jobs <- j
		expectedResults[j-1] = fmt.Sprintf("worker%d-job%d", (j%numWorkers)+1, j) // Worker ID assignment can vary
	}
	close(jobs)

	// Collect results
	var actualResults []string
	for a := 1; a <= numJobs; a++ {
		actualResults = append(actualResults, <-results)
	}

	// Verify results
	assert.Len(t, actualResults, numJobs, "Number of results does not match number of jobs")
}
