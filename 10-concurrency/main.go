package main

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"
)

// Worker represents a function that processes data from a channel
type Worker func(ctx context.Context, id int, jobs <-chan int, results chan<- int)

// simpleGoroutine demonstrates basic goroutine usage
func simpleGoroutine() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("    Hello from goroutine!")
	}()

	wg.Wait()
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

	ch <- 1
	ch <- 2
	ch <- 3

	var results []int
	for i := 0; i < 3; i++ {
		results = append(results, <-ch)
	}
	sort.Ints(results)
	fmt.Printf("    Buffered values (sorted): %v\n", results)
}

// channelSynchronization demonstrates using channels for synchronization
func channelSynchronization() {
	done := make(chan struct{})

	go func() {
		fmt.Println("    Working...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("    Done working!")
		close(done)
	}()

	<-done
}

// worker function for demonstrating channel directions and context support
func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("    Worker %d canceled\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("    Worker %d exiting\n", id)
				return
			}
			fmt.Printf("    Worker %d processing job %d\n", id, job)
			time.Sleep(100 * time.Millisecond)
			results <- job * 2
		}
	}
}

// channelDirections demonstrates channel direction constraints with context and waitgroup
func channelDirections() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, results, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for result := range results {
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

// selectExample demonstrates the select statement with proper synchronization
func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Channel 1"
	}()

	go func() {
		defer wg.Done()
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

	wg.Wait()
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
