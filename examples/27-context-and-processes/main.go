package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

// ProcessManager handles multiple background processes
type ProcessManager struct {
	processes map[string]*exec.Cmd
	mu        sync.RWMutex
}

// NewProcessManager creates a new process manager
func NewProcessManager() *ProcessManager {
	return &ProcessManager{
		processes: make(map[string]*exec.Cmd),
	}
}

// StartProcess starts a new background process with context
func (pm *ProcessManager) StartProcess(ctx context.Context, name string, command string, args ...string) error {
	cmd := exec.CommandContext(ctx, command, args...)

	// Set up pipes for stdout and stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to create stderr pipe: %v", err)
	}

	// Start process
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start process: %v", err)
	}

	pm.mu.Lock()
	pm.processes[name] = cmd
	pm.mu.Unlock()

	// Handle output in goroutines
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			log.Printf("[%s] %s", name, scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			log.Printf("[%s][ERROR] %s", name, scanner.Text())
		}
	}()

	// Wait for process in goroutine
	go func() {
		err := cmd.Wait()
		if err != nil {
			if ctx.Err() == context.Canceled {
				log.Printf("[%s] Process canceled by context", name)
			} else if _, ok := err.(*exec.ExitError); ok {
				log.Printf("[%s] Process exited with error", name)
			} else {
				log.Printf("[%s] Process error: %v", name, err)
			}
		} else {
			log.Printf("[%s] Process completed successfully", name)
		}

		pm.mu.Lock()
		delete(pm.processes, name)
		pm.mu.Unlock()
	}()

	return nil
}

// StopProcess stops a specific process
func (pm *ProcessManager) StopProcess(name string) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	cmd, exists := pm.processes[name]
	if !exists {
		return fmt.Errorf("process %s not found", name)
	}

	if err := cmd.Process.Signal(syscall.SIGTERM); err != nil {
		return fmt.Errorf("failed to stop process: %v", err)
	}

	return nil
}

// StopAll stops all managed processes
func (pm *ProcessManager) StopAll() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	for name, cmd := range pm.processes {
		if err := cmd.Process.Signal(syscall.SIGTERM); err != nil {
			log.Printf("Failed to stop process %s: %v", name, err)
		}
	}
}

// demonstrateContextTimeout shows context timeout functionality
func demonstrateContextTimeout() {
	log.Println("Starting context timeout demonstration...")

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Create a long-running command
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.CommandContext(ctx, "timeout", "5")
	} else {
		cmd = exec.CommandContext(ctx, "sleep", "5")
	}

	log.Println("Starting long-running command (sleep 5)...")
	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start command: %v", err)
		return
	}

	// Wait for command or context timeout
	err := cmd.Wait()
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("Command was terminated due to context timeout (after 2 seconds)")
	} else if err != nil {
		log.Printf("Command error: %v", err)
	} else {
		log.Println("Command completed successfully")
	}
	log.Println("Context timeout demonstration completed")
}

// demonstrateProcessSpawning shows how to spawn multiple processes
func demonstrateProcessSpawning(pm *ProcessManager) {
	log.Println("Starting process spawning demonstration...")

	// Create a parent context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start multiple background processes
	processes := []struct {
		name    string
		command string
		args    []string
	}{
		{
			name:    "counter1",
			command: "bash",
			args:    []string{"-c", "for i in {1..5}; do echo Counter 1: $i; sleep 1; done"},
		},
		{
			name:    "counter2",
			command: "bash",
			args:    []string{"-c", "for i in {1..5}; do echo Counter 2: $i; sleep 1.5; done"},
		},
	}

	for _, p := range processes {
		log.Printf("Starting process: %s", p.name)
		if err := pm.StartProcess(ctx, p.name, p.command, p.args...); err != nil {
			log.Printf("Failed to start %s: %v", p.name, err)
		}
	}

	// Let processes run for a while
	log.Println("Letting processes run for 5 seconds...")
	time.Sleep(5 * time.Second)

	// Stop one process
	log.Println("Stopping counter1 process...")
	if err := pm.StopProcess("counter1"); err != nil {
		log.Printf("Failed to stop counter1: %v", err)
	}

	// Let the other process continue
	log.Println("Letting counter2 run for 3 more seconds...")
	time.Sleep(3 * time.Second)

	log.Println("Process spawning demonstration completed")
}

// demonstrateProcessExec shows how to execute a process and capture its output
func demonstrateProcessExec() {
	log.Println("Starting process execution demonstration...")

	// Execute a command and capture its output
	log.Println("Running 'go version'...")
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Failed to execute command: %v", err)
		return
	}
	log.Printf("Go version output: %s", output)

	// Execute a command with pipes
	log.Println("Running directory listing...")
	cmd = exec.Command("ls", "-l")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("dir")
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("Failed to create stdout pipe: %v", err)
		return
	}

	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start command: %v", err)
		return
	}

	// Read output line by line
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		log.Printf("Directory entry: %s", scanner.Text())
	}

	if err := cmd.Wait(); err != nil {
		log.Printf("Command failed: %v", err)
		return
	}

	log.Println("Process execution demonstration completed")
}

func main() {
	log.Println("=== Context and Process Management Example ===")

	// Create process manager
	pm := NewProcessManager()

	// Setup signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Run demonstrations sequentially for clearer output
	log.Println("=== Demonstrating Context Timeout ===")
	demonstrateContextTimeout()

	log.Println("\n=== Demonstrating Process Spawning ===")
	demonstrateProcessSpawning(pm)

	log.Println("\n=== Demonstrating Process Execution ===")
	demonstrateProcessExec()

	// Check for interrupt signal
	select {
	case <-sigChan:
		log.Println("\nReceived interrupt signal")
		cancel()
		pm.StopAll()
	case <-ctx.Done():
		log.Println("\nContext canceled")
	default:
		// If no signal received, just clean up
		pm.StopAll()
	}

	log.Println("\nAll demonstrations completed")
}
