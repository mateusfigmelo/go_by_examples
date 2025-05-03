package main

import (
	"fmt"
	"time"
)

// selectWithTimeout demonstrates select with timeout
func selectWithTimeout() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Result after 2 seconds"
	}()

	select {
	case result := <-ch:
		fmt.Printf("    Received: %s\n", result)
	case <-time.After(1 * time.Second):
		fmt.Println("    Timeout: operation took too long")
	}
}

// nonBlockingChannelOps demonstrates non-blocking channel operations
func nonBlockingChannelOps() {
	messages := make(chan string)
	signals := make(chan bool)

	// Non-blocking receive
	select {
	case msg := <-messages:
		fmt.Printf("    Received message: %s\n", msg)
	default:
		fmt.Println("    No message received")
	}

	// Non-blocking send
	select {
	case messages <- "hi":
		fmt.Println("    Sent message")
	default:
		fmt.Println("    No message sent")
	}

	// Multi-way non-blocking select
	select {
	case msg := <-messages:
		fmt.Printf("    Received message: %s\n", msg)
	case sig := <-signals:
		fmt.Printf("    Received signal: %v\n", sig)
	default:
		fmt.Println("    No activity")
	}
}

// closingChannels demonstrates channel closing and detection
func closingChannels() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Worker goroutine
	go func() {
		for {
			job, more := <-jobs
			if more {
				fmt.Printf("    Received job %d\n", job)
			} else {
				fmt.Println("    Received all jobs")
				done <- true
				return
			}
		}
	}()

	// Send jobs
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Printf("    Sent job %d\n", j)
	}
	close(jobs)
	fmt.Println("    Sent all jobs")

	<-done
}

// rangeOverChannels demonstrates iterating over channel values
func rangeOverChannels() {
	queue := make(chan string, 5)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)

	fmt.Println("    Queue contents:")
	for elem := range queue {
		fmt.Printf("    - %s\n", elem)
	}
}

// timerExample demonstrates timer usage
func timerExample() {
	// One-time timer
	timer1 := time.NewTimer(500 * time.Millisecond)
	go func() {
		<-timer1.C
		fmt.Println("    Timer 1 expired")
	}()

	// Timer cancellation
	timer2 := time.NewTimer(time.Second)
	go func() {
		if timer2.Stop() {
			fmt.Println("    Timer 2 stopped")
		}
	}()

	time.Sleep(250 * time.Millisecond)
}

// tickerExample demonstrates ticker usage for periodic operations
func tickerExample() {
	ticker := time.NewTicker(200 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Printf("    Tick at %v\n", t.Format("15:04:05.000"))
			}
		}
	}()

	// Let it tick 3 times
	time.Sleep(600 * time.Millisecond)
	ticker.Stop()
	done <- true
}

// rateLimiter demonstrates rate limiting using tickers
func rateLimiter() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Rate limit: 1 request per 200ms
	limiter := time.Tick(200 * time.Millisecond)

	fmt.Println("    Rate limited requests:")
	for req := range requests {
		<-limiter // Rate limit
		fmt.Printf("    Processing request %d\n", req)
	}
}

func main() {
	fmt.Println("=== Go Advanced Channel Operations ===")

	fmt.Println("\n1. Select with Timeout:")
	selectWithTimeout()

	fmt.Println("\n2. Non-Blocking Channel Operations:")
	nonBlockingChannelOps()

	fmt.Println("\n3. Closing Channels:")
	closingChannels()

	fmt.Println("\n4. Range over Channels:")
	rangeOverChannels()

	fmt.Println("\n5. Timer Example:")
	timerExample()

	fmt.Println("\n6. Ticker Example:")
	tickerExample()

	fmt.Println("\n7. Rate Limiter Example:")
	rateLimiter()
}
