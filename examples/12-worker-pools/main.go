package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task represents a unit of work
type Task struct {
	ID       int
	Duration time.Duration
}

// Result represents the output of processing a task
type Result struct {
	TaskID    int
	WorkerID  int
	Duration  time.Duration
	Timestamp time.Time
}

// basicWorkerPool demonstrates a simple worker pool pattern
func basicWorkerPool(numWorkers int, numJobs int) {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go func(id int) {
			for j := range jobs {
				fmt.Printf("    Worker %d processing job %d\n", id, j)
				time.Sleep(100 * time.Millisecond)
				results <- j * 2
			}
		}(w)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

// advancedWorkerPool demonstrates a worker pool with structured tasks and results
func advancedWorkerPool(numWorkers int, numTasks int) {
	var wg sync.WaitGroup
	tasks := make(chan Task, numTasks)
	results := make(chan Result, numTasks)

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for task := range tasks {
				// Process task
				time.Sleep(task.Duration)
				results <- Result{
					TaskID:    task.ID,
					WorkerID:  id,
					Duration:  task.Duration,
					Timestamp: time.Now(),
				}
			}
		}(w)
	}

	// Send tasks
	go func() {
		for i := 1; i <= numTasks; i++ {
			tasks <- Task{
				ID:       i,
				Duration: time.Duration(rand.Intn(500)) * time.Millisecond,
			}
		}
		close(tasks)
	}()

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("    Task %d completed by worker %d (took %v)\n",
			result.TaskID, result.WorkerID, result.Duration)
	}
}

// waitGroupExample demonstrates WaitGroup synchronization
func waitGroupExample() {
	var wg sync.WaitGroup
	workerCount := 5

	// Launch multiple workers
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Simulate work
			duration := time.Duration(rand.Intn(1000)) * time.Millisecond
			time.Sleep(duration)
			fmt.Printf("    Worker %d completed after %v\n", id, duration)
		}(i)
	}

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("    All workers completed")
}

// rateLimitedWorkerPool demonstrates a worker pool with rate limiting
func rateLimitedWorkerPool(numWorkers int, numJobs int, rateLimit time.Duration) {
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)
	limiter := time.Tick(rateLimit)

	// Start workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for job := range jobs {
				<-limiter // Rate limit
				// Process job
				time.Sleep(100 * time.Millisecond)
				results <- fmt.Sprintf("Worker %d completed job %d", id, job)
			}
		}(w)
	}

	// Send jobs
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// Wait for workers and close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("    %s\n", result)
	}
}

// dynamicWorkerPool demonstrates a worker pool that can adjust its size
func dynamicWorkerPool(initialWorkers int, maxWorkers int, numJobs int) {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	control := make(chan bool) // Channel to control worker count

	var workerCount int
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Worker factory
	startWorker := func() {
		mu.Lock()
		workerCount++
		id := workerCount
		mu.Unlock()

		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case job, ok := <-jobs:
					if !ok {
						return
					}
					// Process job
					time.Sleep(100 * time.Millisecond)
					result := job * 2
					fmt.Printf("    Worker %d completed job %d with result %d\n", id, job, result)
					results <- result
				case <-control:
					return
				}
			}
		}()
		fmt.Printf("    Started worker %d\n", id)
	}

	// Start initial workers
	for i := 0; i < initialWorkers; i++ {
		startWorker()
	}

	// Monitor and adjust worker count
	go func() {
		for i := 0; i < maxWorkers-initialWorkers; i++ {
			time.Sleep(500 * time.Millisecond)
			if len(jobs) > len(results) {
				startWorker()
			}
		}
	}()

	// Send jobs
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// Collect results
	for i := 0; i < numJobs; i++ {
		<-results
	}

	// Cleanup
	close(control)
	wg.Wait()
}

func main() {
	fmt.Println("=== Go Worker Pools and Rate Limiting ===")

	rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("\n1. Basic Worker Pool:")
	basicWorkerPool(3, 5)

	fmt.Println("\n2. Advanced Worker Pool:")
	advancedWorkerPool(3, 5)

	fmt.Println("\n3. WaitGroup Example:")
	waitGroupExample()

	fmt.Println("\n4. Rate Limited Worker Pool:")
	rateLimitedWorkerPool(3, 5, 200*time.Millisecond)

	fmt.Println("\n5. Dynamic Worker Pool:")
	dynamicWorkerPool(2, 5, 10)
}
