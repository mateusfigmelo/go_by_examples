package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// WorkerPool manages multiple workers
type WorkerPool struct {
	workers map[int]context.CancelFunc
	mu      sync.RWMutex
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool() *WorkerPool {
	return &WorkerPool{
		workers: make(map[int]context.CancelFunc),
	}
}

// AddWorker starts a new worker with the given ID
func (wp *WorkerPool) AddWorker(id int) {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	ctx, cancel := context.WithCancel(context.Background())
	wp.workers[id] = cancel

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Printf("Worker %d: Shutting down", id)
				return
			default:
				log.Printf("Worker %d: Working...", id)
				time.Sleep(time.Second)
			}
		}
	}()
}

// StopWorker stops a specific worker
func (wp *WorkerPool) StopWorker(id int) {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	if cancel, exists := wp.workers[id]; exists {
		cancel()
		delete(wp.workers, id)
		log.Printf("Worker %d: Stopped", id)
	}
}

// StopAll stops all workers
func (wp *WorkerPool) StopAll() {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	for id, cancel := range wp.workers {
		cancel()
		delete(wp.workers, id)
		log.Printf("Worker %d: Stopped", id)
	}
}

// demonstrateSignalHandling shows how to handle different signals
func demonstrateSignalHandling(wp *WorkerPool) (exitCode int) {
	// Create signal channel for multiple signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGINT,  // Ctrl+C
		syscall.SIGTERM, // Termination request
		syscall.SIGHUP,  // Terminal closed
		syscall.SIGUSR1, // User-defined signal 1
	)
	defer signal.Stop(sigChan)

	// Start some workers
	for i := 1; i <= 3; i++ {
		wp.AddWorker(i)
	}

	// Handle signals
	for sig := range sigChan {
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Printf("Received signal: %v", sig)
			wp.StopAll()
			log.Println("Graceful shutdown completed")
			return 0
		case syscall.SIGHUP:
			log.Printf("Received SIGHUP: Reloading workers")
			wp.StopAll()
			for i := 1; i <= 3; i++ {
				wp.AddWorker(i)
			}
		case syscall.SIGUSR1:
			log.Printf("Received SIGUSR1: Status report")
			wp.mu.RLock()
			log.Printf("Active workers: %d", len(wp.workers))
			wp.mu.RUnlock()
		}
	}
	return 0
}

// demonstrateExitHandling shows how to handle program exit
func demonstrateExitHandling() int {
	// Register cleanup function
	defer func() {
		log.Println("Cleaning up resources...")
		time.Sleep(time.Second) // Simulate cleanup
		log.Println("Cleanup completed")
	}()

	// Register panic recovery
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
			return
		}
	}()

	// Example of different exit scenarios
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "panic":
			panic("Demonstrating panic recovery")
		case "error":
			log.Println("Demonstrating error exit")
			return 1
		}
	}

	return 0
}

func main() {
	// Set up logging
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("=== Signal and Exit Handling Example ===")

	// Create worker pool
	wp := NewWorkerPool()

	// Start exit handling demonstration in a goroutine
	go func() {
		exitCode := demonstrateExitHandling()
		if exitCode != 0 {
			os.Exit(exitCode)
		}
	}()

	// Start signal handling demonstration
	log.Println("\nSignal Handling:")
	log.Println("- Use Ctrl+C (SIGINT) to exit")
	log.Println("- Use SIGHUP to reload workers")
	log.Println("- Use SIGUSR1 for status report")
	exitCode := demonstrateSignalHandling(wp)
	os.Exit(exitCode)
}
