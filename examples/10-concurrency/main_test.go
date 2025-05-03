// Package main implements concurrency examples and their tests.
// This is part of the go_by_examples tutorial series.
package main

import (
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/mateusfigmelo/go_by_examples/testutil"
)

func TestSimpleGoroutine(t *testing.T) {
	output := testutil.CaptureOutput(simpleGoroutine)
	if !testutil.ContainsOutput(output, "Hello from goroutine!") {
		t.Errorf("Expected output to contain 'Hello from goroutine!', got %q", output)
	}
}

func TestChannelCommunication(t *testing.T) {
	output := testutil.CaptureOutput(channelCommunication)
	expected := "Received: Message from goroutine"
	if !testutil.ContainsOutput(output, expected) {
		t.Errorf("Expected output to contain %q, got %q", expected, output)
	}
}

func TestBufferedChannels(t *testing.T) {
	output := testutil.CaptureOutput(bufferedChannels)
	expected := "Buffered values: 1, 2, 3"
	if !testutil.ContainsOutput(output, expected) {
		t.Errorf("Expected output to contain %q, got %q", expected, output)
	}
}

func TestChannelSynchronization(t *testing.T) {
	start := time.Now()
	output := testutil.CaptureOutput(channelSynchronization)
	duration := time.Since(start)

	// Check output
	if !testutil.ContainsOutput(output, "Working...") || !testutil.ContainsOutput(output, "Done working!") {
		t.Errorf("Expected output to contain work messages, got %q", output)
	}

	// Check timing (should take ~500ms)
	if duration < 500*time.Millisecond {
		t.Errorf("Function returned too quickly, expected ~500ms, took %v", duration)
	}
}

func TestWorker(t *testing.T) {
	jobs := make(chan int, 2)
	results := make(chan int, 2)

	// Start a worker
	go worker(1, jobs, results)

	// Send test jobs
	testJobs := []int{1, 2}
	for _, job := range testJobs {
		jobs <- job
	}
	close(jobs)

	// Verify results
	for i := 0; i < len(testJobs); i++ {
		result := <-results
		expected := testJobs[i] * 2
		if result != expected {
			t.Errorf("Expected result %d, got %d", expected, result)
		}
	}
}

func TestChannelDirections(t *testing.T) {
	output := testutil.CaptureOutput(channelDirections)

	// Verify all results are processed
	resultCount := 0
	for _, line := range strings.Split(output, "\n") {
		if strings.Contains(line, "Result:") {
			resultCount++
		}
	}

	if resultCount != 5 {
		t.Errorf("Expected 5 results, got %d", resultCount)
	}
}

func TestWaitGroupExample(t *testing.T) {
	output := testutil.CaptureOutput(waitGroupExample)

	// Check if all goroutines executed
	executedCount := 0
	for _, line := range strings.Split(output, "\n") {
		if strings.Contains(line, "Goroutine") {
			executedCount++
		}
	}

	if executedCount != 3 {
		t.Errorf("Expected 3 goroutines to execute, got %d", executedCount)
	}

	// Check completion message
	if !testutil.ContainsOutput(output, "All goroutines completed") {
		t.Error("Missing completion message")
	}
}

func TestSelectExample(t *testing.T) {
	output := testutil.CaptureOutput(selectExample)

	// Check both channels received
	ch1Received := testutil.ContainsOutput(output, "Received from ch1")
	ch2Received := testutil.ContainsOutput(output, "Received from ch2")

	if !ch1Received || !ch2Received {
		t.Errorf("Expected messages from both channels, ch1: %v, ch2: %v", ch1Received, ch2Received)
	}
}

// Benchmark tests
func BenchmarkWorker(b *testing.B) {
	jobs := make(chan int, b.N)
	results := make(chan int, b.N)
	var wg sync.WaitGroup

	// Start worker
	go func() {
		worker(1, jobs, results)
		wg.Done()
	}()
	wg.Add(1)

	b.ResetTimer()
	// Send jobs
	for i := 0; i < b.N; i++ {
		jobs <- i
	}
	close(jobs)

	// Receive results
	for i := 0; i < b.N; i++ {
		<-results
	}
	wg.Wait()
}
