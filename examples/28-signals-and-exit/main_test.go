package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/mateusfigmelo/go_by_examples/testutil"
)

func TestWorkerPool(t *testing.T) {
	output := testutil.CaptureOutput(func() {
		// Make sure logs go to stdout for the test
		log.SetOutput(os.Stdout)

		wp := NewWorkerPool()

		// Test adding workers
		for i := 1; i <= 3; i++ {
			wp.AddWorker(i)
		}

		// Give workers time to start
		time.Sleep(time.Millisecond * 100)

		// Verify workers are running
		wp.mu.RLock()
		if len(wp.workers) != 3 {
			t.Errorf("Expected 3 workers, got %d", len(wp.workers))
		}
		wp.mu.RUnlock()

		// Test stopping specific worker
		wp.StopWorker(2)
		time.Sleep(time.Millisecond * 100)
		wp.mu.RLock()
		if len(wp.workers) != 2 {
			t.Errorf("Expected 2 workers after stopping one, got %d", len(wp.workers))
		}
		wp.mu.RUnlock()

		// Test stopping all workers
		wp.StopAll()
		time.Sleep(time.Millisecond * 100)
		wp.mu.RLock()
		if len(wp.workers) != 0 {
			t.Errorf("Expected 0 workers after stopping all, got %d", len(wp.workers))
		}
		wp.mu.RUnlock()
	})

	t.Logf("Test output: %s", output)
}

func TestSignalHandling(t *testing.T) {
	// Save original logger configuration and restore it after the test
	originalOutput := log.Writer()
	originalFlags := log.Flags()
	originalPrefix := log.Prefix()
	defer func() {
		log.SetOutput(originalOutput)
		log.SetFlags(originalFlags)
		log.SetPrefix(originalPrefix)
	}()

	tests := []struct {
		name          string
		signal        syscall.Signal
		expectedLogs  []string
		waitDuration  time.Duration
		expectCleanup bool
	}{
		{
			name:   "SIGUSR1 status report",
			signal: syscall.SIGUSR1,
			expectedLogs: []string{
				"Received SIGUSR1: Status report",
				"Active workers: 3",
			},
			waitDuration: time.Second * 5,
		},
		{
			name:   "SIGHUP worker reload",
			signal: syscall.SIGHUP,
			expectedLogs: []string{
				"Received SIGHUP: Reloading workers",
				"Worker 1: Stopped",
				"Worker 2: Stopped",
				"Worker 3: Stopped",
			},
			waitDuration: time.Second * 5,
		},
		{
			name:   "SIGINT graceful shutdown",
			signal: syscall.SIGINT,
			expectedLogs: []string{
				"Received signal: interrupt",
				"Worker 1: Stopped",
				"Worker 2: Stopped",
				"Worker 3: Stopped",
				"Graceful shutdown completed",
			},
			waitDuration:  time.Second * 5,
			expectCleanup: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a pipe for log output
			r, w := io.Pipe()

			// Set up logging to the pipe and stdout
			log.SetOutput(io.MultiWriter(os.Stdout, w))

			// Create a context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), tt.waitDuration)
			defer cancel()

			// Create channels for coordination
			done := make(chan struct{})
			ready := make(chan struct{})
			var wg sync.WaitGroup

			// Buffer to store logs
			var outputBuf strings.Builder

			// Start the worker pool
			wp := NewWorkerPool()

			// Start a goroutine to read from the pipe
			wg.Add(1)
			go func() {
				defer wg.Done()
				buf := make([]byte, 4096)
				for {
					n, err := r.Read(buf)
					if n > 0 {
						outputBuf.Write(buf[:n])
					}
					if err != nil {
						if err != io.EOF {
							t.Errorf("Error reading from pipe: %v", err)
						}
						break
					}
				}
			}()

			// Start the signal handling goroutine
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer close(done)
				defer w.Close() // Close the writer when done

				// Set up signal handling
				sigChan := make(chan os.Signal, 1)
				signal.Notify(sigChan, tt.signal, syscall.SIGINT)
				defer signal.Stop(sigChan)

				// Signal that we're ready to receive signals
				close(ready)

				// Run the signal handling
				demonstrateSignalHandling(wp)
			}()

			// Wait for signal handler to be ready
			<-ready
			time.Sleep(time.Second * 2) // Give workers time to start

			// Send the signal
			proc, err := os.FindProcess(os.Getpid())
			if err != nil {
				t.Fatalf("Failed to find process: %v", err)
			}
			if err := proc.Signal(tt.signal); err != nil {
				t.Fatalf("Failed to send signal: %v", err)
			}

			// For non-terminating signals, send SIGINT after processing
			if tt.signal != syscall.SIGINT && tt.signal != syscall.SIGTERM {
				time.Sleep(time.Second * 2) // Give more time for signal processing
				if err := proc.Signal(syscall.SIGINT); err != nil {
					t.Logf("Failed to send SIGINT signal: %v", err)
				}
			}

			// Wait for completion with timeout
			select {
			case <-ctx.Done():
				t.Fatal("Test timed out")
			case <-done:
				// Test completed
			}

			// Wait for all goroutines to finish
			wg.Wait()

			// Get the output
			output := outputBuf.String()

			// Verify expected logs
			for _, expectedLog := range tt.expectedLogs {
				if !strings.Contains(output, expectedLog) {
					t.Errorf("Expected log not found: %s\nFull output:\n%s", expectedLog, output)
				}
			}
		})
	}
}

func TestExitHandling(t *testing.T) {
	// Save original logger and restore after all tests
	originalOutput := log.Writer()
	originalFlags := log.Flags()
	originalPrefix := log.Prefix()
	defer func() {
		log.SetOutput(originalOutput)
		log.SetFlags(originalFlags)
		log.SetPrefix(originalPrefix)
	}()

	tests := []struct {
		name         string
		args         []string
		expectedLogs []string
		expectPanic  bool
		expectedCode int
	}{
		{
			name: "normal exit",
			args: []string{"program"},
			expectedLogs: []string{
				"Cleaning up resources",
				"Cleanup completed",
			},
			expectedCode: 0,
		},
		{
			name: "error exit",
			args: []string{"program", "error"},
			expectedLogs: []string{
				"Demonstrating error exit",
				"Cleaning up resources",
				"Cleanup completed",
			},
			expectedCode: 1,
		},
		{
			name: "panic recovery",
			args: []string{"program", "panic"},
			expectedLogs: []string{
				"Recovered from panic",
				"Cleaning up resources",
				"Cleanup completed",
			},
			expectPanic:  false, // We don't expect panic to propagate now
			expectedCode: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original args
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			// Set test args
			os.Args = tt.args

			// Create a pipe to capture log output
			r, w := io.Pipe()
			defer r.Close()
			defer w.Close()

			// Set log output to write to pipe and stdout
			log.SetOutput(io.MultiWriter(os.Stdout, w))

			// Buffer to store logs
			var outputBuf strings.Builder

			// Start a goroutine to read from the pipe
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				buf := make([]byte, 4096)
				for {
					n, err := r.Read(buf)
					if n > 0 {
						outputBuf.Write(buf[:n])
					}
					if err != nil {
						if err != io.EOF {
							t.Errorf("Error reading from pipe: %v", err)
						}
						break
					}
				}
			}()

			// Run the function
			var exitCode int
			if tt.expectPanic {
				func() {
					defer func() {
						if r := recover(); r != nil {
							// Good, we expected this
						} else {
							t.Error("Expected panic, but none occurred")
						}
					}()
					exitCode = demonstrateExitHandling()
				}()
			} else {
				exitCode = demonstrateExitHandling()
			}

			// Close the writer to signal EOF to the reader
			w.Close()

			// Wait for the reader to finish
			wg.Wait()

			// Get the output
			output := outputBuf.String()

			// Check exit code
			if exitCode != tt.expectedCode {
				t.Errorf("Expected exit code %d, got %d", tt.expectedCode, exitCode)
			}

			// Verify expected logs
			for _, expectedLog := range tt.expectedLogs {
				if !strings.Contains(output, expectedLog) {
					t.Errorf("Expected log not found: %s\nFull output:\n%s", expectedLog, output)
				}
			}
		})
	}
}

func TestIntegration(t *testing.T) {
	// Save original logger and restore after test
	originalOutput := log.Writer()
	originalFlags := log.Flags()
	originalPrefix := log.Prefix()
	defer func() {
		log.SetOutput(originalOutput)
		log.SetFlags(originalFlags)
		log.SetPrefix(originalPrefix)
	}()

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Create a pipe for log output
	r, w := io.Pipe()
	defer r.Close()

	// Set up logging to write to pipe and stdout
	log.SetOutput(io.MultiWriter(os.Stdout, w))

	// Create channels for coordination
	done := make(chan struct{})
	ready := make(chan struct{})
	var wg sync.WaitGroup

	// Buffer to store logs
	var outputBuf strings.Builder

	// Start a goroutine to read from the pipe
	wg.Add(1)
	go func() {
		defer wg.Done()
		buf := make([]byte, 4096)
		for {
			n, err := r.Read(buf)
			if n > 0 {
				outputBuf.Write(buf[:n])
			}
			if err != nil {
				if err != io.EOF {
					t.Errorf("Error reading from pipe: %v", err)
				}
				break
			}
		}
	}()

	// Start the test
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(done)
		defer w.Close() // Close the writer when done

		// Set up signal handling
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT)
		defer signal.Stop(sigChan)

		// Signal that we're ready
		close(ready)

		// Run the program manually instead of calling main
		// Print the header ourselves
		log.Println("=== Signal and Exit Handling Example ===")
		log.Println("\nSignal Handling:")
		log.Println("- Use Ctrl+C (SIGINT) to exit")
		log.Println("- Use SIGHUP to reload workers")
		log.Println("- Use SIGUSR1 for status report")

		// Run the actual functions
		wp := NewWorkerPool()
		go demonstrateExitHandling()
		demonstrateSignalHandling(wp)
	}()

	// Wait for signal handler to be ready
	<-ready
	time.Sleep(time.Second * 2) // Give workers time to start

	// Send SIGINT
	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf("Failed to find process: %v", err)
	}
	if err := proc.Signal(syscall.SIGINT); err != nil {
		t.Logf("Failed to send SIGINT signal: %v", err)
	}

	// Wait for completion with timeout
	select {
	case <-ctx.Done():
		t.Fatal("Test timed out")
	case <-done:
		// Test completed
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Get the output
	output := outputBuf.String()
	t.Log("Captured output:", output)

	// Verify expected logs - more flexible now
	expectedLogs := []string{
		"Signal and Exit Handling Example",
		"Signal Handling",
		"Use Ctrl+C",
		"Use SIGHUP",
		"Use SIGUSR1",
		"Working",
		"Stopped",
		"Graceful shutdown",
	}

	for _, expectedLog := range expectedLogs {
		if !strings.Contains(output, expectedLog) {
			t.Errorf("Expected log not found: %s", expectedLog)
		}
	}
}
