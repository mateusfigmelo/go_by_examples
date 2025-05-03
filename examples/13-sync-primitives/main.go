package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// AtomicCounter demonstrates atomic operations
type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Decrement() {
	atomic.AddInt64(&c.value, -1)
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

// atomicExample demonstrates atomic operations
func atomicExample() {
	counter := &AtomicCounter{}
	var wg sync.WaitGroup

	// Launch 100 goroutines that increment
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	// Launch 50 goroutines that decrement
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Decrement()
		}()
	}

	wg.Wait()
	fmt.Printf("    Final counter value: %d\n", counter.Value())
}

// SafeCounter demonstrates mutex usage
type SafeCounter struct {
	mu    sync.Mutex
	value map[string]int
}

func (c *SafeCounter) Increment(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value[key]
}

// mutexExample demonstrates mutex usage
func mutexExample() {
	counter := SafeCounter{value: make(map[string]int)}
	var wg sync.WaitGroup

	// Launch multiple goroutines that increment different keys
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n%3)
			counter.Increment(key)
		}(i)
	}

	wg.Wait()

	// Print results
	for i := 0; i < 3; i++ {
		key := fmt.Sprintf("key%d", i)
		fmt.Printf("    %s: %d\n", key, counter.Value(key))
	}
}

// RWMutexCounter demonstrates RWMutex usage
type RWMutexCounter struct {
	mu    sync.RWMutex
	value map[string]int
}

func (c *RWMutexCounter) Increment(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value[key]++
}

func (c *RWMutexCounter) Value(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value[key]
}

// rwMutexExample demonstrates RWMutex usage
func rwMutexExample() {
	counter := RWMutexCounter{value: make(map[string]int)}
	var wg sync.WaitGroup

	// Writers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n%3)
			counter.Increment(key)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	// Readers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n%3)
			value := counter.Value(key)
			fmt.Printf("    Read %s: %d\n", key, value)
			time.Sleep(10 * time.Millisecond)
		}(i)
	}

	wg.Wait()
}

// StatefulGoroutine represents a goroutine with state
type StatefulGoroutine struct {
	state     map[string]int
	stateChan chan func()
}

func NewStatefulGoroutine() *StatefulGoroutine {
	sg := &StatefulGoroutine{
		state:     make(map[string]int),
		stateChan: make(chan func()),
	}
	go sg.loop()
	return sg
}

func (sg *StatefulGoroutine) loop() {
	for f := range sg.stateChan {
		f()
	}
}

func (sg *StatefulGoroutine) Update(key string, value int) {
	sg.stateChan <- func() {
		sg.state[key] = value
	}
}

func (sg *StatefulGoroutine) Get(key string) int {
	resultChan := make(chan int)
	sg.stateChan <- func() {
		resultChan <- sg.state[key]
	}
	return <-resultChan
}

func (sg *StatefulGoroutine) Close() {
	close(sg.stateChan)
}

// statefulGoroutineExample demonstrates goroutine state management
func statefulGoroutineExample() {
	sg := NewStatefulGoroutine()
	var wg sync.WaitGroup

	// Update state from multiple goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n%3)
			sg.Update(key, n)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	// Read state from multiple goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n%3)
			value := sg.Get(key)
			fmt.Printf("    %s = %d\n", key, value)
			time.Sleep(50 * time.Millisecond)
		}(i)
	}

	wg.Wait()
	sg.Close()
}

// Once demonstrates sync.Once usage
func onceExample() {
	var once sync.Once
	var wg sync.WaitGroup
	onceBody := func() {
		fmt.Println("    Only once")
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(onceBody)
		}()
	}

	wg.Wait()
}

func main() {
	fmt.Println("=== Go Synchronization Primitives ===")

	fmt.Println("\n1. Atomic Operations:")
	atomicExample()

	fmt.Println("\n2. Mutex Example:")
	mutexExample()

	fmt.Println("\n3. RWMutex Example:")
	rwMutexExample()

	fmt.Println("\n4. Stateful Goroutine:")
	statefulGoroutineExample()

	fmt.Println("\n5. Once Example:")
	onceExample()
}
