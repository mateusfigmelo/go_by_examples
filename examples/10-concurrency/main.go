package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker represents a function that processes data from a channel
type Worker func(id int, jobs <-chan int, results chan<- int)

// simpleGoroutine demonstrates basic goroutine usage
func simpleGoroutine() {
	go func() {
		fmt.Println("    Hello from goroutine!")
	}()
	time.Sleep(100 * time.Millisecond) // Wait for goroutine to finish
}

// channelCommunication demonstrates basic channel communication
func channelCommunication() {
	ch := make(chan string)

	go func() {
		ch <- "Message from goroutine"
	}()

	msg := <-ch
	fmt.Printf("    Received: %s\n", msg)
}

// bufferedChannels demonstrates channel buffering
func bufferedChannels() {
	ch := make(chan int, 3)

	// Can send 3 values without blocking
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Printf("    Buffered values: %d, %d, %d\n", <-ch, <-ch, <-ch)
}

// channelSynchronization demonstrates using channels for synchronization
func channelSynchronization() {
	done := make(chan bool)

	go func() {
		fmt.Println("    Working...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("    Done working!")
		done <- true
	}()

	<-done // Wait for goroutine to finish
}

// worker function for demonstrating channel directions
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("    Worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond)
		results <- job * 2
	}
}

// channelDirections demonstrates channel direction constraints
func channelDirections() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 2 workers
	for w := 1; w <= 2; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for i := 1; i <= 5; i++ {
		result := <-results
		fmt.Printf("    Result: %d\n", result)
	}
}

// waitGroupExample demonstrates WaitGroup usage
func waitGroupExample() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("    Goroutine %d executing\n", id)
			time.Sleep(200 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	fmt.Println("    All goroutines completed")
}

// selectExample demonstrates the select statement
func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("    Received from ch1: %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("    Received from ch2: %s\n", msg2)
		}
	}
}

func main() {
	fmt.Println("=== Go Concurrency Examples ===")

	fmt.Println("\n1. Simple Goroutine:")
	simpleGoroutine()

	fmt.Println("\n2. Channel Communication:")
	channelCommunication()

	fmt.Println("\n3. Buffered Channels:")
	bufferedChannels()

	fmt.Println("\n4. Channel Synchronization:")
	channelSynchronization()

	fmt.Println("\n5. Channel Directions (Workers):")
	channelDirections()

	fmt.Println("\n6. WaitGroup Synchronization:")
	waitGroupExample()

	fmt.Println("\n7. Select Statement:")
	selectExample()
}
