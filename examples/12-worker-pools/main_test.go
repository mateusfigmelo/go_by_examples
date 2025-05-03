package main

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/mateusfigmelo/go_by_examples/testutil"
)

func TestBasicWorkerPool(t *testing.T) {
	tests := []struct {
		name       string
		numWorkers int
		numJobs    int
		wantJobs   []int
	}{
		{
			name:       "small pool",
			numWorkers: 2,
			numJobs:    3,
			wantJobs:   []int{1, 2, 3},
		},
		{
			name:       "larger pool",
			numWorkers: 4,
			numJobs:    5,
			wantJobs:   []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := testutil.CaptureOutput(func() {
				basicWorkerPool(tt.numWorkers, tt.numJobs)
			})

			// Verify each job was processed
			for _, job := range tt.wantJobs {
				want := fmt.Sprintf("processing job %d", job)
				if !strings.Contains(output, want) {
					t.Errorf("basicWorkerPool() missing job %d in output: %s", job, output)
				}
			}
		})
	}
}

func TestAdvancedWorkerPool(t *testing.T) {
	tests := []struct {
		name       string
		numWorkers int
		numTasks   int
	}{
		{
			name:       "small pool",
			numWorkers: 2,
			numTasks:   3,
		},
		{
			name:       "larger pool",
			numWorkers: 4,
			numTasks:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			output := testutil.CaptureOutput(func() {
				advancedWorkerPool(tt.numWorkers, tt.numTasks)
			})
			duration := time.Since(start)

			// Verify all tasks were completed
			for i := 1; i <= tt.numTasks; i++ {
				want := fmt.Sprintf("Task %d completed by worker", i)
				if !strings.Contains(output, want) {
					t.Errorf("advancedWorkerPool() missing task %d completion in output: %s", i, output)
				}
			}

			// Check reasonable execution time
			maxExpectedDuration := time.Duration(tt.numTasks*500) * time.Millisecond
			if duration > maxExpectedDuration {
				t.Errorf("advancedWorkerPool() took %v, want < %v", duration, maxExpectedDuration)
			}
		})
	}
}

func TestWaitGroupExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "concurrent workers",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			output := testutil.CaptureOutput(waitGroupExample)
			duration := time.Since(start)

			// Verify all workers completed
			completedCount := strings.Count(output, "Worker")
			if completedCount != 5 {
				t.Errorf("waitGroupExample() got %d workers, want 5", completedCount)
			}

			if !strings.Contains(output, "All workers completed") {
				t.Error("waitGroupExample() missing completion message")
			}

			// Check reasonable execution time
			if duration > 2*time.Second {
				t.Errorf("waitGroupExample() took %v, want < 2s", duration)
			}
		})
	}
}

func TestRateLimitedWorkerPool(t *testing.T) {
	tests := []struct {
		name       string
		numWorkers int
		numJobs    int
		rateLimit  time.Duration
	}{
		{
			name:       "rate limited pool",
			numWorkers: 2,
			numJobs:    4,
			rateLimit:  100 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			output := testutil.CaptureOutput(func() {
				rateLimitedWorkerPool(tt.numWorkers, tt.numJobs, tt.rateLimit)
			})
			duration := time.Since(start)

			// Verify all jobs were completed
			for i := 1; i <= tt.numJobs; i++ {
				want := fmt.Sprintf("completed job %d", i)
				if !strings.Contains(output, want) {
					t.Errorf("rateLimitedWorkerPool() missing job %d completion in output: %s", i, output)
				}
			}

			// Check rate limiting
			minExpectedDuration := time.Duration(tt.numJobs) * tt.rateLimit
			if duration < minExpectedDuration {
				t.Errorf("rateLimitedWorkerPool() took %v, want >= %v", duration, minExpectedDuration)
			}
		})
	}
}

func TestDynamicWorkerPool(t *testing.T) {
	tests := []struct {
		name           string
		initialWorkers int
		maxWorkers     int
		numJobs        int
	}{
		{
			name:           "growing pool",
			initialWorkers: 2,
			maxWorkers:     4,
			numJobs:        6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := testutil.CaptureOutput(func() {
				dynamicWorkerPool(tt.initialWorkers, tt.maxWorkers, tt.numJobs)
			})

			// Verify initial workers were started
			initialWorkerCount := 0
			for i := 1; i <= tt.initialWorkers; i++ {
				if strings.Contains(output, fmt.Sprintf("Started worker %d", i)) {
					initialWorkerCount++
				}
			}
			if initialWorkerCount != tt.initialWorkers {
				t.Errorf("dynamicWorkerPool() started %d initial workers, want %d", initialWorkerCount, tt.initialWorkers)
			}

			// Verify jobs were completed
			completedCount := 0
			for i := 1; i <= tt.numJobs; i++ {
				if strings.Contains(output, fmt.Sprintf("completed job %d with result", i)) {
					completedCount++
				}
			}
			if completedCount < tt.numJobs {
				t.Errorf("dynamicWorkerPool() completed %d jobs, want %d", completedCount, tt.numJobs)
			}
		})
	}
}

func TestMain(t *testing.T) {
	start := time.Now()
	output := testutil.CaptureOutput(main)
	duration := time.Since(start)

	// Check for section headers
	wantSections := []string{
		"=== Go Worker Pools and Rate Limiting ===",
		"1. Basic Worker Pool:",
		"2. Advanced Worker Pool:",
		"3. WaitGroup Example:",
		"4. Rate Limited Worker Pool:",
		"5. Dynamic Worker Pool:",
	}

	for _, section := range wantSections {
		if !strings.Contains(output, section) {
			t.Errorf("main() missing section %q in output", section)
		}
	}

	// Check reasonable execution time
	if duration > 10*time.Second {
		t.Errorf("main() took %v, want < 10s", duration)
	}
}

// Example outputs for documentation
func Example() {
	main()
	// Output is non-deterministic due to concurrent operations
	// We only verify the section headers
	// Output:
	// === Go Worker Pools and Rate Limiting ===
	// 1. Basic Worker Pool:
	// 2. Advanced Worker Pool:
	// 3. WaitGroup Example:
	// 4. Rate Limited Worker Pool:
	// 5. Dynamic Worker Pool:
}
